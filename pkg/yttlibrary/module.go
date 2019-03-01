package yttlibrary

import (
	"github.com/get-ytt/ytt/pkg/template/core"
	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"
)

var (
	ModuleAPI = starlark.StringDict{
		"module": &starlarkstruct.Module{
			Name: "module",
			Members: starlark.StringDict{
				"make": starlark.NewBuiltin("module.make", core.ErrWrapper(starlarkstruct.MakeModule)),
			},
		},
	}
)
