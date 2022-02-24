/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2022-02-17 14:55:57
 * @LastEditors: sueRimn
 * @LastEditTime: 2022-02-17 15:10:16
 */

package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

// 游戏 JWT 信息
type GameJWT struct {
	//过期时间
	Expire int
	// 生成时间
	CreateAt int
	//用户ID
	UserID int
}

// 生成游戏登录的 JWT token
//
//secretKey 秘钥
//
//iat 生成时间 (UNIX时间戳)
//
//seconds 有效时间
//
//userId 用户UID
//
//返回 (token,错误)
func GetGameJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// 解码游戏 JWT token
//
// token JWT token
func DecodeGameJwtToken(token string, secretKey string) (*GameJWT, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	info := &GameJWT{}
	c, ok := claim.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	uid, ok := c["userId"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid token(userId)")
	}

	info.UserID = int(uid)

	expire, ok := c["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid token(exp)")
	}
	info.Expire = int(expire)

	iat, ok := c["iat"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid token(iat)")
	}
	info.CreateAt = int(iat)

	return info, nil
}
