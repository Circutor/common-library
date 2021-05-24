// Copyright (c) 2021 Circutor S.A. All rights reserved.

package device

import "github.com/circutor/thingsboard-methods/pkg/controller"

// CommonService fields to make calls to thingsBoard.
type CommonService struct {
	ThingsBoard    controller.ThingsBoardController
	DeviceID       string
	DeviceName     string
	AccessToken    string
	UserID         string
	UserToken      string
	AdminToken     string
	SecretKey      string
	ExpirationTime int
}

// NewCommonService creates a new CommonService.
func NewCommonService(tb controller.ThingsBoardController, deviceID, deviceName, accessToken, userID, userToken,
	adminToken, secretKey string, expirationTime int) CommonService {
	return CommonService{
		ThingsBoard:    tb,
		DeviceID:       deviceID,
		DeviceName:     deviceName,
		AccessToken:    accessToken,
		UserID:         userID,
		UserToken:      userToken,
		AdminToken:     adminToken,
		SecretKey:      secretKey,
		ExpirationTime: expirationTime,
	}
}

// VerticalDevice Operations device vertical.
type VerticalDevice interface {
	SetAttrServer(c CommonService, microservice string) (int, map[string]interface{}, error)
	SetAttrClient(c CommonService, microservice string) (int, map[string]interface{}, error)
	SetAttrServerClient(c CommonService, microservice string) (int, map[string]interface{}, error)
}

// TransversalDevice Operations device transversal.
type TransversalDevice interface {
	SetAttrTransversal(c CommonService, microservice string) (int, map[string]interface{}, error)
	DeleteDevice(c CommonService, microservice string) (int, map[string]interface{}, error)
}
