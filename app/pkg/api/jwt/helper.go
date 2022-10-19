package jwt

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type Helper struct {
	secret []byte
}

func NewHelper(secret string) Helper {
	return Helper{
		secret: []byte(secret),
	}
}

func (h *Helper) ParseToken(tokenS string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenS, func(token *jwt.Token) (interface{}, error) {
		return h.secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, err
	} else {
		if err = claims.Valid(jwt.DefaultValidationHelper); err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		return claims, nil
	}
}

func (h *Helper) GeneratePair(userID, issuerName string, roleID uint64) (*Pair, error) {
	claims := newAccessTokenClaims(userID, issuerName, roleID)
	accessToken, err := h.generateToken(claims)
	if err != nil {
		return nil, err
	}

	claims = newRefreshTokenClaims(userID, issuerName, roleID)
	refreshToken, err := h.generateToken(claims)
	if err != nil {
		return nil, err
	}

	return &Pair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (h *Helper) generateToken(claims *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.ToMapClaims())
	return token.SignedString(h.secret)
}

func (h *Helper) ParseMapClaims(mapClaims jwt.MapClaims) *CustomClaims {
	sec, dec := math.Modf(mapClaims["exp"].(float64))
	ExpireAt := time.Unix(int64(sec), int64(dec*(1e9)))

	sec, dec = math.Modf(mapClaims["iss_at"].(float64))
	IssuedAt := time.Unix(int64(sec), int64(dec*(1e9)))

	return &CustomClaims{
		ExpireAt:   ExpireAt,
		UserID:     mapClaims["id"].(string),
		IssuedAt:   IssuedAt,
		IssuerName: mapClaims["iss"].(string),
		RoleID:     uint64(mapClaims["role_id"].(float64)),
	}
}

func (h *Helper) PrepareCookies(pair *Pair) (*http.Cookie, *http.Cookie) {
	accessTokenCookie := new(http.Cookie)
	accessTokenCookie.Name = AccessTokenName
	accessTokenCookie.Path = "/"
	accessTokenCookie.Value = pair.AccessToken
	accessTokenCookie.Secure = true
	accessTokenCookie.HttpOnly = true

	refreshTokenCookie := new(http.Cookie)
	refreshTokenCookie.Name = RefreshTokenName
	refreshTokenCookie.Path = "/"
	refreshTokenCookie.Value = pair.RefreshToken
	refreshTokenCookie.Secure = true
	refreshTokenCookie.HttpOnly = true

	return accessTokenCookie, refreshTokenCookie
}
