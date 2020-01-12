package kapal

import "api-pelabuhan/model"

// KapalUsecase is aaaa asda
type KapalUsecase interface {
	InsertKapal(kapal *model.Kapal) error
	UpdateKapal(kapal *model.Kapal) error
	DeleteKapal(id int) error
	GetAllKapal() (*[]model.Kapal, error)
	GetKapalByID(id int) (*model.Kapal, error)
}
