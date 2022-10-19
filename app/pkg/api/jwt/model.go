package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

const (
	AccessTokenDuration  = 5
	RefreshTokenDuration = 24 * 31
	AccessTokenName      = "Access-Token"
	RefreshTokenName     = "Refresh-Token"
)

type Pair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CustomClaims struct {
	ExpireAt   time.Time
	UserID     string
	IssuedAt   time.Time
	IssuerName string
	RoleID     uint64
}

func newAccessTokenClaims(userID, issuerName string, roleID uint64) *CustomClaims {
	claims := newClaims(userID, issuerName, roleID)
	claims.ExpireAt = claims.ExpireAt.Add(AccessTokenDuration * time.Minute)
	return claims
}

func newRefreshTokenClaims(userID, issuerName string, roleID uint64) *CustomClaims {
	claims := newClaims(userID, issuerName, roleID)
	claims.ExpireAt = claims.ExpireAt.Add(RefreshTokenDuration * time.Hour)
	return claims
}

func newClaims(userID, issuerName string, roleID uint64) *CustomClaims {
	return &CustomClaims{
		ExpireAt:   time.Now(),
		UserID:     userID,
		IssuedAt:   time.Now(),
		IssuerName: issuerName,
		RoleID:     roleID,
	}
}

func (c *CustomClaims) ToMapClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"exp":     jwt.At(c.ExpireAt),
		"id":      c.UserID,
		"iss_at":  jwt.At(time.Now()),
		"iss":     c.IssuerName,
		"role_id": c.RoleID,
	}
}
