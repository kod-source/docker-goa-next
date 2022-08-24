package main

import (
	_ "github.com/kod-source/docker-goa-next/webapi/design"

	"github.com/shogo82148/goa-v1/design"
	"github.com/shogo82148/goa-v1/goagen/codegen"
	genapp "github.com/shogo82148/goa-v1/goagen/gen_app"
	genswagger "github.com/shogo82148/goa-v1/goagen/gen_swagger"
)

func main() {
	gengoa()
}

func gengoa() {
	codegen.ParseDSL()
	codegen.Run(
		genswagger.NewGenerator(
			genswagger.API(design.Design),
		),
		genapp.NewGenerator(
			genapp.API(design.Design),
			genapp.OutDir("app"),
			genapp.Target("app"),
		),
	)
}
