// Copyright (c) 2021 Circutor S.A. All rights reserved.

package core

// AttrTransversal transversal attributes to send to ThingsBoard.
type AttrTransversal struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Company     string `json:"company"`
	Zone        string `json:"zone"`
	Typology    string `json:"typology"`
	Use         string `json:"use"`
}

// NewAttrTransversal created a struct AttrTransversal.
func NewAttrTransversal() AttrTransversal {
	return AttrTransversal{
		Name:        "",
		Description: "",
		Company:     "",
		Zone:        "",
		Typology:    "",
		Use:         "",
	}
}

// AttrComputerC6C12 vertical attributes device type ComputerC6C12.
type AttrComputerC6C12 struct {
	TargetCosPhi     float32 `json:"targetCosPhi"`
	Power            float32 `json:"power"`
	Voltage          float32 `json:"voltage"`
	RejectionFilters bool    `json:"rejectionFilters"`
}

// NewAttrComputerC6C12 created a struct AttrComputerC6C12.
func NewAttrComputerC6C12() AttrComputerC6C12 {
	return AttrComputerC6C12{
		TargetCosPhi:     0,
		Power:            0,
		Voltage:          0,
		RejectionFilters: false,
	}
}
