package services

import (
	"errors"
	"strconv"
	"time"

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

func (s *Server) GenToken(id int32,
	name, role, nickName, phoneNum, email string) (string, error) {
	if len(phoneNum) == 0 {
		phoneNum = strconv.Itoa(int(id)) + name
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		Claims{
			id:       id,
			name:     name,
			role:     role,
			nickName: nickName,
			phoneNum: phoneNum,
			email:    email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Second * time.Duration(7200)).Unix(),
			},
		})

	return token.SignedString([]byte(phoneNum))
}

func GetRoleFromToken(c *Claims) string {
	return c.role
}
