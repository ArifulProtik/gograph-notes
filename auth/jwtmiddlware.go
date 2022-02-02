package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authString string
type CustomContext struct {
	echo.Context
	ctx context.Context
}

func EchoContextFromContext(ctx context.Context) (*echo.Context, error) {
	echoContext := ctx.Value("EchoContextKey")
	if echoContext == nil {
		err := fmt.Errorf("could not retrieve echo.Context")
		return nil, err
	}

	ec, ok := echoContext.(*echo.Context)
	if !ok {
		err := fmt.Errorf("echo.Context has wrong type")
		return nil, err
	}
	return ec, nil
}
func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authheader := c.Request().Header.Get("Authorization")
			if authheader == "" {
				ctx := context.Background()
				cc := &CustomContext{c, ctx}
				return next(cc)
			}

			bearer := "Bearer "
			token := authheader[len(bearer):]
			validate, err := JwtValidate(context.Background(), token)
			if err != nil || !validate.Valid {
				return c.String(http.StatusForbidden, "Invalid token")
			}
			claims, _ := validate.Claims.(*JwtCustomClaim)
			ctx := context.WithValue(c.Request().Context(), authString("auth"), claims)
			c.SetRequest(c.Request().WithContext(ctx))
			cc := &CustomContext{c, ctx}
			return next(cc)
		}

	}
}
func CtxValue(ctx context.Context) *JwtCustomClaim {
	raw, _ := ctx.Value(authString("auth")).(*JwtCustomClaim)
	return raw
}
