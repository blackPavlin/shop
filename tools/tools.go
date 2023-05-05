//go:build tools
// +build tools

package tools

import (
	_ "github.com/deepmap/oapi-codegen/cmd/oapi-codegen"
	_ "github.com/dmarkham/enumer"
	_ "github.com/golang/mock/mockgen"
	_ "golang.org/x/tools/cmd/stringer"
)
