package dock

import "api-pelabuhan/model"

type DockUsecase interface {
	InsertDock(dock *model.Dock) error
	UpdateDock(dock *model.Dock) error
	DeleteDock(id int) error
	GetAllDock() (*[]model.Dock, error)
	GetDockByID(id int) (*model.Dock, error)
}
