package jwt

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
)

type AuthInterceptor struct {
	jwtHelper Helper
	roles     map[string][]uint64
}

func NewAuthInterceptor(jwtHelper Helper, roles map[string][]uint64) *AuthInterceptor {
	return &AuthInterceptor{jwtHelper: jwtHelper, roles: roles}
}

func (i *AuthInterceptor) AuthorizeHandler(ctx context.Context) (context.Context, error) {
	fromContext := grpc.ServerTransportStreamFromContext(ctx)
	method := fromContext.Method()

	accessibleRoles, ok := i.roles[method]
	if !ok {
		// everyone can access
		return ctx, nil
	}

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	tokenMC, err := i.jwtHelper.ParseToken(token)
	if err != nil {
		return nil, ErrBadToken
	}

	claims := i.jwtHelper.ParseMapClaims(tokenMC)

	grpc_ctxtags.Extract(ctx).Set("role_id", claims.RoleID)
	grpc_ctxtags.Extract(ctx).Set("user_id", claims.UserID)

	for _, role := range accessibleRoles {
		if role == claims.RoleID {
			return ctx, nil
		}
	}

	ctx = context.WithValue(ctx, "role_id", claims.RoleID)
	ctx = context.WithValue(ctx, "user_id", claims.UserID)

	return ctx, nil
}
