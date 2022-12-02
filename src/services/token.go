package services

import (
	"errors"
	"starnet-demo/src/db"
	"starnet-demo/src/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	id       int32
	role     string
	name     string
	nickName string
	phoneNum string
	email    string
	jwt.StandardClaims
}

func (s *Server) ParseToken(token string) (*Claims, error) {

	t, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		c, ok := t.Claims.(*Claims)
		if !ok {
			return nil, errors.New("the claim type error")
		}

		if len(c.phoneNum) == 0 {
			return strconv.Itoa(int(c.id)) + c.name, nil
		}
		return c.phoneNum, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			return nil, errors.New(ve.Error())
		}
		return nil, errors.New("unknown error")
	}

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func (s *Server) GenToken(user *db.User, expiresAt int64) (string, error) {

	phoneNum := user.UserPhoneNum
	if len(user.UserPhoneNum) == 0 {
		phoneNum = strconv.Itoa(int(user.Id)) + user.UserName
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		Claims{
			id:       user.Id,
			name:     user.UserName,
			role:     db.UserRoleTypeName[user.UserRole],
			nickName: user.UserNickName,
			phoneNum: phoneNum,
			email:    user.UserEmail,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expiresAt,
			},
		})

	return token.SignedString([]byte(phoneNum))
}

func (c *Claims) GetRoleFromToken() string {
	return c.role
}

func (c *Claims) GetUserInfo() *models.UserInfo {
	return &models.UserInfo{
		Id:           c.id,
		UserName:     c.name,
		UserRole:     c.role,
		UserNickName: c.nickName,
		UserPhoneNum: c.phoneNum,
		UserEmail:    c.email,
		Expires:      c.ExpiresAt,
	}
}
