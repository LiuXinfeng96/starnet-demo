package services

import (
	"errors"
	"starnet-demo/src/db"

	"github.com/dgrijalva/jwt-go"
)

// DefaultTokenSecretKey default token secret key
const DEFAULT_TOKEN_SECRET_KEY = "1045836262777123654"

type MyClaims struct {
	Id       int32
	Role     string
	Name     string
	NickName string
	PhoneNum string
	Email    string
	jwt.StandardClaims
}

func (s *Server) ParseToken(token string) (*MyClaims, error) {

	t, err := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(DEFAULT_TOKEN_SECRET_KEY), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			s.sulog.Infof("parse token failed, err: [%s], [%s]\n", err.Error(), ve.Error())
			return nil, errors.New(ve.Error())
		}
		return nil, errors.New("unknown error")
	}

	if claims, ok := t.Claims.(*MyClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func (s *Server) GenToken(user *db.User, expiresAt int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, MyClaims{
		Id:       user.Id,
		Role:     db.UserRoleTypeName[user.UserRole],
		Name:     user.UserName,
		NickName: user.UserNickName,
		PhoneNum: user.UserPhoneNum,
		Email:    user.UserEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	})

	t, err := token.SignedString([]byte(DEFAULT_TOKEN_SECRET_KEY))
	if err != nil {
		s.sulog.Infof("signed token failed, err: [%s]\n", err.Error())
		return t, err
	}
	return t, nil
}
