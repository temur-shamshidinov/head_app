package token

import (
	"head_app/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	UserID   string
	UserRole string
	jwt.StandardClaims
}

var secretJWTKey = []byte("secret_key")

func GenerateJWT(claim models.Claim) (string, error) {

	expTime := time.Now().Add(5 * time.Hour)

	jwtClaim := Claim{
		UserID:         claim.UserID,
		UserRole:       claim.UserRole,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expTime.Unix()},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)

	tokenString, err := token.SignedString(secretJWTKey)
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (*Claim, error) {

	var claim = &Claim{}

	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return secretJWTKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claim, nil
}
