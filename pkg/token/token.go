package token

import (
	log "head_app/pkg/logger"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/saidamir98/udevs_pkg/logger"
)

type JWTHandler struct {
	Sub       string
	Exp       string
	Iat       string
	Aud       []string
	Role      string
	SignedKey string
	Log       log.Log
	Token     string
	Timeout   int
}

type CustomClaims struct {
	*jwt.Token
	Sub  string   `json:"sub"`
	Exp  string   `json:"exp"`
	Iat  string   `json:"iat"`
	Aud  []string `json:"aud"`
	Role string   `json:"role"`
}

func (j *JWTHandler) GenerateToken() (string, error) {
	var (
		token  *jwt.Token
		claims jwt.MapClaims
	)

	token = jwt.New(jwt.SigningMethodHS256)

	claims = token.Claims.(jwt.MapClaims)
	claims["sub"] = j.Sub
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(j.Timeout)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["aud"] = j.Aud  
	claims["role"] = j.Role

	accessToken, err := token.SignedString([]byte(j.SignedKey))
	if err != nil {
		j.Log.Error("error on genereting token", logger.Error(err))
		return "", err
	}

	return accessToken, nil

}
