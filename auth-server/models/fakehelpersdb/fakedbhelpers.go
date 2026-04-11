package fakedbhelpers

import "github.com/golang-jwt/jwt/v5"

var refreshTokens []*jwt.Token

func init() {
	refreshTokens = []*jwt.Token{}
}

func StoreRefreshToken(token *jwt.Token) error {
	refreshTokens = append(refreshTokens, token)
	return nil
}
