package repo

import (
	"api-pelabuhan/docked"
	"database/sql"
)

var _ docked.DockedRepo = &DockedRepoMysqlImpl{}

// DockedRepoMysqlImpl is aaa
type DockedRepoMysqlImpl struct {
	db *sql.DB
}

// CreateDockedRepoMysqlImpl is aaa
func CreateDockedRepoMysqlImpl(db *sql.DB) docked.DockedRepo {
	return &DockedRepoMysqlImpl{db}
}

// NewDocked is aaa
func (m *DockedRepoMysqlImpl) NewDocked(docked *model.Docked) error {

}
