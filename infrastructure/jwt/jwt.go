package jwt

import (
	"github.com/azkanurhuda/golang-absent/domain/entity"
	"github.com/azkanurhuda/golang-absent/domain/service"
	"github.com/azkanurhuda/golang-absent/pkg/conf"
	jwtgo "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	expireTime = time.Hour * 24 * 121
)

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwtgo.StandardClaims
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

var _ service.Jwt = (*Service)(nil)

func (j *Service) Sign(user *entity.User) (*entity.AccessToken, error) {
	now := time.Now()
	expiresAt := now.Add(expireTime)
	claims := &CustomClaims{
		UserID:         user.ID,
		StandardClaims: jwtgo.StandardClaims{ExpiresAt: expiresAt.Unix()},
	}

	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)

	signedString, err := token.SignedString(conf.JWT.SigningKey())
	if err != nil {
		return nil, err
	}

	accessToken := &entity.AccessToken{
		UserID:    user.ID,
		Token:     signedString,
		ExpiresAt: expiresAt,
	}

	return accessToken, nil
}

func (j *Service) Verify(signedToken string) (string, error) {
	token, err := jwtgo.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(token *jwtgo.Token) (interface{}, error) {
			return conf.JWT.SigningKey(), nil
		},
	)

	if err != nil {
		log.Error(err)
		return "", ErrorParseClaims
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return "", ErrorParseClaims
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return "", ErrorTokenExpired
	}

	return claims.UserID, nil
}
