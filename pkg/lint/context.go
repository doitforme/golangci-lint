package lint

import (
	"context"
	"fmt"

	"github.com/doitforme/golangci-lint/internal/cache"
	"github.com/doitforme/golangci-lint/pkg/config"
	"github.com/doitforme/golangci-lint/pkg/exitcodes"
	"github.com/doitforme/golangci-lint/pkg/goanalysis/load"
	"github.com/doitforme/golangci-lint/pkg/lint/linter"
	"github.com/doitforme/golangci-lint/pkg/logutils"
)

type ContextBuilder struct {
	cfg *config.Config

	pkgLoader *PackageLoader

	pkgCache *cache.Cache

	loadGuard *load.Guard
}

func NewContextBuilder(cfg *config.Config, pkgLoader *PackageLoader,
	pkgCache *cache.Cache, loadGuard *load.Guard,
) *ContextBuilder {
	return &ContextBuilder{
		cfg:       cfg,
		pkgLoader: pkgLoader,
		pkgCache:  pkgCache,
		loadGuard: loadGuard,
	}
}

func (cl *ContextBuilder) Build(ctx context.Context, log logutils.Log, linters []*linter.Config) (*linter.Context, error) {
	pkgs, deduplicatedPkgs, err := cl.pkgLoader.Load(ctx, linters)
	if err != nil {
		return nil, fmt.Errorf("failed to load packages: %w", err)
	}

	if len(deduplicatedPkgs) == 0 {
		return nil, fmt.Errorf("%w: running `go mod tidy` may solve the problem", exitcodes.ErrNoGoFiles)
	}

	ret := &linter.Context{
		Packages: deduplicatedPkgs,

		// At least `unused` linters works properly only on original (not deduplicated) packages,
		// see https://github.com/doitforme/golangci-lint/pull/585.
		OriginalPackages: pkgs,

		Cfg:       cl.cfg,
		Log:       log,
		PkgCache:  cl.pkgCache,
		LoadGuard: cl.loadGuard,
	}

	return ret, nil
}
