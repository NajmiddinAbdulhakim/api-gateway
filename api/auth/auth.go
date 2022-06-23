package auth

import (
	"github.com/NajmiddinAbdulhakim/api-gateway/pkg/logger"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWTHendler struct {
	Sub string
	Iss string
	Exp string
	Iat string
	Aud []string
	Role string
	SigniKey string
	Log logger.Logger
	Token string
}

func (jwtHendler *JWTHendler) GenerateAuthJWT() (access, refresh string, err error) {
	var(
		accessToken,refreshToken *jwt.Token
		claims jwt.MapClaims
	)
	accessToken = jwt.New(jwt.SigningMethodES256)
	refreshToken = jwt.New(jwt.SigningMethodES256)
	
	claims = accessToken.Claims.(jwt.MapClaims)

	claims["sub"] = jwtHendler.Sub
	claims["iss"] = jwtHendler.Iss
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = jwtHendler.Role
	claims["aud"] = jwtHendler.Aud

	access, err = accessToken.SignedString([]byte(jwtHendler.SigniKey))
	if err != nil {
		jwtHendler.Log.Error("Error generate access token", logger.Error(err))
		return
	}
	refresh, err = refreshToken.SignedString([]byte(jwtHendler.SigniKey))
	if err != nil {
		jwtHendler.Log.Error("Error generate refresh token", logger.Error(err))
		return
	}
	return access, refresh, nil
}