// Copyright (c) 2021 Circutor S.A. All rights reserved.

package device

import "github.com/circutor/thingsboard-methods/pkg/controller"

// CommonService fields to make calls to thingsBoard.
type CommonService struct {
	thingsBoard    controller.ThingsBoardController
	deviceID       string
	deviceName     string
	accessToken    string
	userID         string
	userToken      string
	adminToken     string
	secretKey      string
	expirationTime int
}

// NewCommonService creates a new CommonService.
func NewCommonService(tb controller.ThingsBoardController, deviceID, deviceName, accessToken, userID, userToken,
	adminToken, secretKey string, expirationTime int) CommonService {
	return CommonService{
		thingsBoard:    tb,
		deviceID:       deviceID,
		deviceName:     deviceName,
		accessToken:    accessToken,
		userID:         userID,
		userToken:      userToken,
		adminToken:     adminToken,
		secretKey:      secretKey,
		expirationTime: expirationTime,
	}
}

// VerticalDevice Operations device vertical.
type VerticalDevice interface {
	SetAttrServer(c CommonService, microservice string) (int, map[string]interface{}, error)
	SetAttrClient(c CommonService, microservice string) (int, map[string]interface{}, error)
}

// TransversalDevice Operations device transversal.
type TransversalDevice interface {
	SetAttrTransversal(c CommonService, microservice string) (int, map[string]interface{}, error)
	DeleteDevice(c CommonService, microservice string) (int, map[string]interface{}, error)
}
