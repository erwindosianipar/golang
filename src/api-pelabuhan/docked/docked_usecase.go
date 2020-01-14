package docked

import "api-pelabuhan/model"

type DockedUsecase interface {
	NewDocked(dock *model.Dock) error
}
