// Copyright (c) 2021 Circutor S.A. All rights reserved.

package transversal

import (
	"github.com/circutor/common-library/pkg/data"
	"github.com/circutor/common-library/pkg/errors"
	"github.com/circutor/common-library/pkg/translations"
	"github.com/circutor/thingsboard-methods/pkg/controller"
	"github.com/go-playground/validator/v10"
)

// InterfaceTransversal Operations device transversal.
type InterfaceTransversal interface {
	SetAttrDevice(deviceID, token, msg string, tb controller.ThingsBoardController,
		data data.InterfaceData) (int, map[string]interface{}, error)
	ReClaimDevice(deviceName, token, msg string, tb controller.ThingsBoardController) (int, map[string]interface{}, error)
}

//nolint:lll
//go:generate mockery --name InterfaceTransversal --structname InterfaceTransversalMock --filename InterfaceTransversalMock.go

// Transversal transversal attributes to send to ThingsBoard.
type Transversal struct {
	Name        string `json:"name" validate:"required,max=32"`
	Description string `json:"description" validate:"required,max=64"`
	Company     string `json:"company" validate:"required,max=32"`
	Zone        string `json:"zone" validate:"required,max=32"`
	Typology    string `json:"typology" validate:"required,max=32"`
	Use         string `json:"use" validate:"required,max=32"`
}

// NewTransversal created a struct AttrTransversal implement InterfaceTransversal.
func NewTransversal(name, description, company, zone, typology, use string) Transversal {
	return Transversal{
		Name:        name,
		Description: description,
		Company:     company,
		Zone:        zone,
		Typology:    typology,
		Use:         use,
	}
}

// IsValid validates if a Transversal struct is valid or not.
func (t Transversal) IsValid() ([]string, error) {
	validate := validator.New()

	err := validate.Struct(t)
	if err != nil {
		messageError := translations.TranslateEnError(err, validate)

		return messageError, errors.WrapErrFound(err, err.Error())
	}

	return nil, nil
}
