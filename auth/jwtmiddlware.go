package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authString string
type CustomContext struct {
	echo.Context
	ctx context.Context
}

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authheader := c.Request().Header.Get("Authorization")
			if authheader == "" {
				return next(c)
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
