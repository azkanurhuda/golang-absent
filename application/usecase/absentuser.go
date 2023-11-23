package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/azkanurhuda/golang-absent/domain/entity"
	"github.com/azkanurhuda/golang-absent/domain/repository"
	"github.com/azkanurhuda/golang-absent/domain/service"
	"github.com/azkanurhuda/golang-absent/pkg/conf"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type AbsentUserUseCase interface {
	CheckInAbsentUser(ctx context.Context, absentUser *entity.AbsentUser) (*entity.AbsentUser, error)
	CheckOutAbsentUser(ctx context.Context, absentUser *entity.AbsentUser) (*entity.AbsentUser, error)
}

type absentUserUseCase struct {
	absentUserRepository repository.AbsentUser
	jwtService           service.Jwt
}

func NewAbsentUserUseCase(absentUserRepository repository.AbsentUser, jwtService service.Jwt) AbsentUserUseCase {
	return &absentUserUseCase{
		absentUserRepository: absentUserRepository,
		jwtService:           jwtService,
	}
}

func (u *absentUserUseCase) CheckInAbsentUser(ctx context.Context, payload *entity.AbsentUser) (*entity.AbsentUser, error) {
	ipAddr, err := getPublicIP()
	if err != nil {
		return nil, err
	}

	location, err := getLocationInfo(ipAddr)
	if err != nil {
		return nil, err
	}

	input := &entity.AbsentUser{
		ID:        payload.ID,
		UserID:    payload.UserID,
		IPAddress: ipAddr,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Status:    payload.Status,
	}

	if err := u.absentUserRepository.CheckInAbsentUser(ctx, input); err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	res := &entity.AbsentUser{
		ID:        input.ID,
		UserID:    input.UserID,
		IPAddress: input.IPAddress,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Status:    input.Status,
	}

	return res, nil
}

func (u *absentUserUseCase) CheckOutAbsentUser(ctx context.Context, payload *entity.AbsentUser) (*entity.AbsentUser, error) {
	ipAddr, err := getPublicIP()
	if err != nil {
		return nil, err
	}

	location, err := getLocationInfo(ipAddr)
	if err != nil {
		return nil, err
	}

	input := &entity.AbsentUser{
		ID:        payload.ID,
		UserID:    payload.UserID,
		IPAddress: ipAddr,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Status:    payload.Status,
	}

	if err := u.absentUserRepository.CheckOutAbsentUser(ctx, input); err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	res := &entity.AbsentUser{
		ID:        input.ID,
		UserID:    input.UserID,
		IPAddress: input.IPAddress,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		Status:    input.Status,
	}

	return res, nil
}

func getLocationInfo(ipAddress string) (*entity.Location, error) {
	apiKey := conf.APIStackKey.AccessKey

	apiURL := fmt.Sprintf("http://api.ipstack.com/%s?access_key=%s", ipAddress, apiKey)
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response
	var location entity.Location
	err = json.Unmarshal(body, &location)
	if err != nil {
		return nil, err
	}

	return &location, nil
}

func getPublicIP() (string, error) {
	response, err := http.Get("https://api64.ipify.org?format=text")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	ip, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}
