package directives

import (
	"context"
	"famesensor/go-graphql-jwt/middlewares"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	tokenData := middlewares.GetJwtClaimFromHeader(ctx)
	if tokenData == nil {
		return nil, &gqlerror.Error{
			Message: "You're not authorization",
		}
	}

	return next(ctx)
}
