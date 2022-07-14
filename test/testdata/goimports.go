//args: -Egoimports
package testdata

import (
	"fmt" // ERROR "File is not `goimports`-ed"
	"github.com/lgvital/golangci-lint/pkg/config"
)

func Bar() {
	fmt.Print("x")
	_ = config.Config{}
}
