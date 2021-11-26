package presentation

import (
	"MINIPROJECT-BACKEND/features/driver"
	"MINIPROJECT-BACKEND/features/driver/presentation/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PekerjaPercaPresentasi struct {
	PekerjaPercaBisnis driver.Bisnis
}

func NewPekerjaPercaPresentasi(db driver.Bisnis) *PekerjaPercaPresentasi {
	return &PekerjaPercaPresentasi{
		PekerjaPercaBisnis: db,
	}
}

func (percapresntasi *PekerjaPercaPresentasi) CreatePekerjaPerca(e echo.Context) error {
	var requestdata request.PekerjaPercaRequest
	err := e.Bind(&requestdata)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	if err := percapresntasi.PekerjaPercaBisnis.CreatePekerjaPerca(request.ToCore(requestdata)); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Sukses",
	})
}
