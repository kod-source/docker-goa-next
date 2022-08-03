package design

import (
	. "github.com/shogo82148/goa-v1/design"
	. "github.com/shogo82148/goa-v1/design/apidsl"
)

// JWT jwtでログインしていないと叩けないAPI
var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access")
})

var _ = Resource("auth", func() {
	Action("login", func() {
		Routing(POST("login"))
		Description("jwtでのログイン処理")
		Payload(func() {
			Attribute("email", String, "メール", func() {
				Example("sample@goa-sample.test.com")
			})
			Attribute("password", String, "パスワード", func() {
				Example("test1234")
			})
			Required("email", "password")
		})
		Response(OK, token)
		Response(NotFound)
		Response(BadRequest)
		Response(InternalServerError)
	})
	Action("sign_up", func() {
		Routing(POST("sign_up"))
		Description("サインアップ")
		Payload(func() {
			Attribute("name", String, "名前", func() {
				Example("田中　太郎")
			})
			Attribute("email", String, "メール", func() {
				Example("sample@goa-sample.test.com")
			})
			Attribute("password", String, "パスワード", func() {
				Example("test1234")
			})
			Required("name", "email", "password")
		})
		Response(Created, token)
		Response(BadRequest)
		Response(InternalServerError)
	})
})

var token = MediaType("application/vnd.token+json", func() {
	Description("token")
	Attribute("token", String, "token value", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
	})
	Attribute("user", user, "user value")
	View("default", func() {
		Attribute("token")
		Attribute("user")
	})
	Required("token")
	Required("user")
})
