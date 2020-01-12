package usecase

import (
	"errors"
	"fmt"
	"goclean/meja"
	"goclean/model"
)

type MejaUsecaseImpl struct {
	mejaRepo meja.MejaRepo
}

func CreateMejaUsecase(mejaRepo meja.MejaRepo) meja.MejaUsecase {
	return &MejaUsecaseImpl{mejaRepo}
}

func (m *MejaUsecaseImpl) GetById(id int) (*model.Meja, error) {
	return m.mejaRepo.GetById(id)
}

func (m *MejaUsecaseImpl) Insert(meja *model.Meja) error {

	mejaVal, err := m.mejaRepo.GetById(meja.ID)
	if err != nil {
		return fmt.Errorf("[CreateMejaUsecase.Insert] Error when get meja by id' : %w", err)
	}

	fmt.Println(mejaVal)

	if mejaVal != nil {
		return errors.New("ID Meja sudah ada, silahkan masukkan id lain")
	}

	return m.mejaRepo.Insert(meja)
}
