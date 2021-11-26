package driver

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

type Bisnis interface {
	CreatePekerjaPerca(PekerjaPerca) error
}

type Data interface {
	CreatePekerjaPerca(PekerjaPerca) error
}
