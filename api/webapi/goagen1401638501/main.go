// Code generated by goagen v1.5.13, DO NOT EDIT.
//
// Code Generator
//
// Command:
// $ goagen bootstrap -d github.com/kod-source/docker-goa-next/webapi/design

package main

import (
	"github.com/shogo82148/goa-v1/goagen/gen_client"
	"fmt"
	"strings"
	"github.com/shogo82148/goa-v1/dslengine"
	_ "github.com/kod-source/docker-goa-next/webapi/design"
)


func main() {
	// Check if there were errors while running the first DSL pass
	dslengine.FailOnError(dslengine.Errors)

	// Now run the secondary DSLs
	dslengine.FailOnError(dslengine.Run())

	files, err := genclient.Generate()
	dslengine.FailOnError(err)

	// We're done
	fmt.Println(strings.Join(files, "\n"))
}