package meja

import "goclean/model"

type MejaUsecase interface {
	GetById(id int) (*model.Meja, error)
	Insert(meja *model.Meja) error
}
