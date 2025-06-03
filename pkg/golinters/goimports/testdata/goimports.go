//golangcitest:args -Egoimports
package testdata

import (
	"fmt" // want "File is not properly formatted"

	"github.com/doitforme/golangci-lint/pkg/config"
)

func Bar() {
	fmt.Print("x")
	_ = config.Config{}
}
