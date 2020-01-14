package usecase

import (
	"api-pelabuhan/dock"
	"api-pelabuhan/model"
)

// DockUsecaseImpl is aaa
type DockUsecaseImpl struct {
	dockRepo dock.DockRepo
}

// CreateDockUsecase is aaa
func CreateDockUsecase(dockRepo dock.DockRepo) dock.DockUsecase {
	return &DockUsecaseImpl{dockRepo}
}

// InsertDock is aaa
func (m *DockUsecaseImpl) InsertDock(dock *model.Dock) error {
	return m.dockRepo.InsertDock(dock)
}

// UpdateDock is aaa
func (m *DockUsecaseImpl) UpdateDock(dock *model.Dock) error {
	return m.dockRepo.UpdateDock(dock)
}

// GetDockByID is aaa
func (m *DockUsecaseImpl) GetDockByID(id int) (*model.Dock, error) {
	return m.dockRepo.GetDockByID(id)
}

// DeleteDock is aaa
func (m *DockUsecaseImpl) DeleteDock(id int) error {
	return m.dockRepo.DeleteDock(id)
}

// GetAllDock is aaa
func (m *DockUsecaseImpl) GetAllDock() (*[]model.Dock, error) {
	return m.dockRepo.GetAllDock()
}
