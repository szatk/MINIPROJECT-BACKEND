package bussiness

import (
	"MINIPROJECT-BACKEND/features/driver"
)


type PekerjaPercaBisnis struct {
	PekerjaPercaData driver.Data
}

func NewPekerjaPercaBisnis (dd driver.Data) driver.Bisnis{
	return &PekerjaPercaBisnis {dd}
}

func (percabisnis *PekerjaPercaBisnis) CreatePekerjaPerca(data driver.PekerjaPerca) error {
	if err := percabisnis.PekerjaPercaData.CreatePekerjaPerca(data); err != nil {
		return err
	}
	return nil
}

