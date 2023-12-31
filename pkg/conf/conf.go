package conf

import (
	"fmt"
	"github.com/azkanurhuda/golang-absent/pkg/env"
)

var (
	Server      *server
	Postgres    *postgres
	JWT         *jWT
	APIStackKey *apiStackKey
)

func init() {
	v := env.NewVariables()

	Server = &server{
		Host: v.ServerHost.String(),
		Port: v.ServerPort.String(),
	}

	Postgres = &postgres{
		Host:     v.PostgresHost.String(),
		Port:     v.PostgresPort.Int(),
		Username: v.PostgresUser.String(),
		Password: v.PostgresPassword.String(),
		DB:       v.PostgresDB.String(),
	}

	JWT = &jWT{
		Secret: v.JWTSecret.String(),
	}

	APIStackKey = &apiStackKey{
		AccessKey: v.APIStackKey.String(),
	}
}

type server struct {
	Host string
	Port string
}

func (s *server) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}

type postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
}

func (p *postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		p.Host,
		p.Port,
		p.Username,
		p.DB,
		p.Password)
}

type jWT struct {
	Secret string
}

func (j *jWT) SigningKey() []byte {
	return []byte(j.Secret)
}

type apiStackKey struct {
	AccessKey string
}

func (a *apiStackKey) APIStackKey() string {
	return a.AccessKey
}
