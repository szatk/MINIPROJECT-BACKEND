package data

import (
	"MINIPROJECT-BACKEND/features/driver"
)

type PekerjaPerca struct {
	Id             uint
	NIK            string
	HousePercaId   uint
	NamaDepan      string
	NamaBelakang   string
	TanggalLahir   string
	NoTelepon      string
	Alamat         string
	Kabupaten_Kota string
	Provinsi       string
}

func fromCore(data driver.PekerjaPerca) PekerjaPerca {
	return PekerjaPerca{
		NIK: data.NIK,
		HousePercaId: data.HousePercaId,
		NamaDepan: data.NamaDepan,
		NamaBelakang: data.NamaBelakang,
		TanggalLahir: data.TanggalLahir,
		NoTelepon: data.NoTelepon,
		Alamat: data.Alamat,
		Kabupaten_Kota: data.Kabupaten_Kota,
		Provinsi: data.Provinsi,
	}
}