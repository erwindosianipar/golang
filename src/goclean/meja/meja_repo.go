package meja

import "goclean/model"

type MejaRepo interface {
	GetById(id int) (*model.Meja, error)
	Insert(meja *model.Meja) error
}
