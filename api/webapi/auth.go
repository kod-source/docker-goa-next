package webapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/usecase"
	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
	uu usecase.UserUseCase
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service, uu usecase.UserUseCase) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController"), uu: uu}
}

// Login runs the login action.
func (c *AuthController) Login(ctx *app.LoginAuthContext) error {
	email := ctx.Payload.Email
	password := ctx.Payload.Password
	user, err := c.uu.GetUserByEmail(ctx, email, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.NotFound(&app.ServiceVerror{
				Code:    404,
				Message: "メールアドレスが間違っています。\n再度ご確認ください。",
				Status:  "404 Not Found",
			})
		}
		if errors.Is(err, myerrors.ErrPasswordWorng) {
			return ctx.BadRequest(&app.ServiceVerror{
				Code:    400,
				Details: myerrors.ErrPasswordWorng,
				Message: "パスワードが間違っています。",
				Status:  myerrors.ErrPasswordWorng.Error(),
			})
		}
		return ctx.InternalServerError()
	}
	token, err := c.uu.CreateJWTToken(ctx, user.ID, user.Name)
	if err != nil {
		return ctx.InternalServerError()
	}
	res := &app.Token{
		User: &app.User{
			ID:        int(user.ID),
			Email:     &user.Email,
			Name:      &user.Name,
			Password:  &user.Password,
			CreatedAt: &user.CreatedAt,
			Avatar:    user.Avatar,
		},
		Token: *token,
	}
	return ctx.OK(res)
}

func (c *AuthController) SignUp(ctx *app.SignUpAuthContext) error {
	user, err := c.uu.SignUp(ctx, ctx.Payload.Name, ctx.Payload.Email, ctx.Payload.Password, ctx.Payload.Avatar)
	if err != nil {
		if number := myerrors.GetMySQLErrorNumber(err); number == myerrors.MySQLErrorDuplicate.Number {
			return ctx.BadRequest(&app.ServiceVerror{
				Code:    400,
				Message: "そのメールアドレスは既に使用されています",
				Status:  "your email address is already in use",
			})
		}
		return ctx.InternalServerError()
	}
	token, err := c.uu.CreateJWTToken(ctx, user.ID, user.Name)
	if err != nil {
		return ctx.InternalServerError()
	}
	return ctx.Created(&app.Token{
		Token: *token,
		User: &app.User{
			CreatedAt: &user.CreatedAt,
			Email:     &user.Email,
			ID:        int(user.ID),
			Name:      &user.Name,
			Password:  &user.Password,
			Avatar:    user.Avatar,
		},
	})
}

// newAuthMiddleware JWT 発行す（APIGatewayにおいて検証済）からユーザ情報を取り出して Context に紐付ける
func newAuthMiddleware() goa.Middleware {
	return func(nextHandler goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			tokenString := req.Header.Get("Authorization")
			if len(tokenString) < 7 || !strings.EqualFold(tokenString[:7], "Bearer ") {
				return goa.ErrBadRequest(fmt.Errorf("unexpected authorization header %s", tokenString))
			}
			tokenString = tokenString[7:]

			// tokenの認証
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil // CreateTokenにて指定した文字列を使います
			})
			if err != nil {
				return goa.ErrBadRequest(fmt.Errorf("failed to parse jwt token %s", tokenString))
			}

			// ユーザーIDを取り出し ctx に詰める
			claims := token.Claims.(jwt.MapClaims)
			userID, ok := claims["user_id"].(float64)
			if !ok {
				return goa.ErrBadRequest(fmt.Errorf("cognito:username not found in token claims"))
			}

			ctx = context.WithValue(ctx, userIDCodeKey, int(userID))

			return nextHandler(ctx, rw, req)
		}
	}
}

type userIDKeyType struct{}

var userIDCodeKey userIDKeyType

func getUserIDCode(ctx context.Context) int {
	v, ok := ctx.Value(userIDCodeKey).(int)
	if !ok {
		return 0
	}
	return v
}
