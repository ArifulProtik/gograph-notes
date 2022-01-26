package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ArifulProtik/gograph-notes/auth"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	tokenData := auth.CtxValue(ctx)
	if tokenData == nil {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}

	return next(ctx)
}
