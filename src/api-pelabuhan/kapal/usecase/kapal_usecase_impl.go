package usecase

import (
	"api-pelabuhan/kapal"
	"api-pelabuhan/model"
)

// KapalUsecaseImpl is aaa
type KapalUsecaseImpl struct {
	kapalRepo kapal.KapalRepo
}

// CreateKapalUsecase is aaa
func CreateKapalUsecase(kapalRepo kapal.KapalRepo) kapal.KapalUsecase {
	return &KapalUsecaseImpl{kapalRepo}
}

// InsertKapal is aaa
func (m *KapalUsecaseImpl) InsertKapal(kapal *model.Kapal) error {
	return m.kapalRepo.InsertKapal(kapal)
}

// UpdateKapal is aaa
func (m *KapalUsecaseImpl) UpdateKapal(kapal *model.Kapal) error {
	return m.kapalRepo.UpdateKapal(kapal)
}

// GetAllKapal is aaa
func (m *KapalUsecaseImpl) GetAllKapal() (*[]model.Kapal, error) {
	return m.kapalRepo.GetAllKapal()
}

// GetKapalByID is aaa
func (m *KapalUsecaseImpl) GetKapalByID(id int) (*model.Kapal, error) {
	return m.kapalRepo.GetKapalByID(id)
}

// DeleteKapal is aaa
func (m *KapalUsecaseImpl) DeleteKapal(id int) error {
	return m.kapalRepo.DeleteKapal(id)
}
