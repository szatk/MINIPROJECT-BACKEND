package request

import (
	"MINIPROJECT-BACKEND/features/driver"
)

type PekerjaPercaRequest struct {
	NIK            string `json:"nik"`
	NamaDepan      string `json: "namadepan"`
	NamaBelakang   string `json: "namabelakang"`
	TanggalLahir   string `json: "tanggallahir"`
	NoTelepon      string `json: "notelepon"`
	Alamat         string `json: "alamat"`
	Kabupaten_Kota string `json: "kabupaten_kota"`
	Provinsi       string `json: "provinsi"`
}

func ToCore(data PekerjaPercaRequest) driver.PekerjaPerca {
	return driver.PekerjaPerca{
		NIK:            data.NIK,
		NamaDepan:      data.NamaDepan,
		NamaBelakang:   data.NamaBelakang,
		TanggalLahir:   data.TanggalLahir,
		NoTelepon:      data.NoTelepon,
		Alamat:         data.Alamat,
		Kabupaten_Kota: data.Kabupaten_Kota,
		Provinsi:       data.Provinsi,
	}
}
