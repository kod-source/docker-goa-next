package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

var _ = API("docker_goa_next", func() {
	Title("The docker_goa API")
	Description("A teaser for goa")
	Host("localhost:3000")
	Scheme("http")

	Origin("/.*localhost.*/", func() {
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
