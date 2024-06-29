package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/assignment-amori/internal/constant"
	"github.com/assignment-amori/internal/entity/generic"
	"github.com/assignment-amori/pkg/errorwrapper"
	"github.com/assignment-amori/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
)

type JWTMiddleware struct {
	secretKey string
}

// GetJWTAuthMiddleware represent middleware getter of JWT Auth.
func (hm *MiddlewareModule) GetJWTAuthMiddleware() *JWTMiddleware {
	return &JWTMiddleware{
		secretKey: hm.appConfig.JWTConfig.SecretKey,
	}
}

func (hm *JWTMiddleware) Filter(next generic.HTTPHandleFunc) generic.HTTPHandleFunc {
	return func(w http.ResponseWriter, r *http.Request) (data interface{}, err error) {
		ctx := r.Context()

		moduleName := utils.GetModuleKey(ctx)
		if moduleName == constant.Module(constant.DefaultString) {
			return data, errorwrapper.New(errorwrapper.ErrIDUnauthorized)
		}

		// Set the Content-Type header to indicate JSON response.
		w.Header().Set("Content-Type", "application/json")

		authHeader := r.Header.Get(constant.HTTPHeaderAuthorization)
		if authHeader == constant.DefaultString {
			return data, errorwrapper.New(errorwrapper.ErrIDUnauthorized)
		}

		// Split the token type and the token itself.
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != constant.HTTPHeaderBearer {
			return data, errorwrapper.New(errorwrapper.ErrIDUnauthorized)
		}

		// Parse the token.
		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return data, errorwrapper.New(errorwrapper.ErrIDUnauthorized)
			}
			return []byte(hm.secretKey), nil
		})

		if err != nil || !token.Valid {
			return data, errorwrapper.Wrap(err, errorwrapper.ErrIDUnauthorized)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return data, errorwrapper.New(errorwrapper.ErrIDUnauthorized)
		}

		ctx = context.WithValue(ctx, constant.ContextUser, claims)
		r = r.WithContext(ctx)
		return next(w, r)
	}
}
