package utils

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
)

var jwtKey = "userId"

func GetJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[jwtKey] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	jsonUid := ctx.Value(jwtKey)
	i, err := strconv.Atoi(fmt.Sprintf("%v", jsonUid))
	if err != nil {
		i = 0
	}
	uid = int64(i)
	return uid
}
