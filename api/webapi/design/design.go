package design

import (
	"errors"
	"os"

	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = API("docker_goa_next", func() {
	Title("The docker_goa API")
	Description("A teaser for goa")
	host := os.Getenv("SWAGGER_HOST")
	if host == "" {
		panic(errors.New("SWAGGER_HOST is not set"))
	}
	Host(host)
	Scheme("http")

	Version("v1")
	BasePath("/api/v1")

	Origin("*", func() {
		Headers("Authorization, Content-Type")
		Methods("GET", "POST", "PUT", "DELETE")
	})
})

var _ = Resource("operands", func() {
	Action("add", func() {
		Routing(GET("add/:left/:right"))
		Description("add returns the sum of the left and right parameters in the response body")
		Params(func() {
			Param("left", Integer, "Left operand")
			Param("right", Integer, "Right operand")
		})
		Response(OK, "text/plain")
	})
})
