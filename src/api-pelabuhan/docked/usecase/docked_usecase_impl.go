package usecase

import (
	"api-pelabuhan/docked"
)

// DockedUsecaseImpl is aaa
type DockedUsecaseImpl struct {
	dockedRepo docked.DockedRepo
}

// CreateDockedUsecase is aaa
func CreateDockedUsecase(dockedRepo docked.DockedRepo) docked.DockedUsecase {
	return &DockedUsecaseImpl{dockedRepo}
}

// NewDocked is aaa
func (m *DockedUsecaseImpl) NewDocked(docked *model.Docked) error {
	return m.dockedRepo.NewDock(dock)
}
