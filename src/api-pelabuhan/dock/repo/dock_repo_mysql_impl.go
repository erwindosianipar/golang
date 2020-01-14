package repo

import (
	"api-pelabuhan/dock"
	"api-pelabuhan/model"
	"database/sql"
	"errors"
	"fmt"
)

var _ dock.DockRepo = &DockRepoMysqlImpl{}

// DockRepoMysqlImpl is aaa
type DockRepoMysqlImpl struct {
	db *sql.DB
}

// CreateDockRepoMysqlImpl is aaa
func CreateDockRepoMysqlImpl(db *sql.DB) dock.DockRepo {
	return &DockRepoMysqlImpl{db}
}

// InsertDock is aaa
func (m *DockRepoMysqlImpl) InsertDock(dock *model.Dock) error {
	tx, err := m.db.Begin()

	if err != nil {
		return fmt.Errorf("[DockRepoMysqlImpl.InsertDock] Error when begin transaction : %w", err)
	}

	query := "INSERT INTO dock (kode) VALUES (?)"
	_, err = m.db.Exec(query, dock.Kode)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[DockRepoMysqlImpl.InsertDock] Error when running query insert dock: %w", err)
	}

	tx.Commit()
	return nil

}

// UpdateDock is aaa
func (m *DockRepoMysqlImpl) UpdateDock(dock *model.Dock) error {
	query := "UPDATE dock SET kode = ? WHERE id = ?"

	_, err := m.db.Exec(query, dock.Kode, dock.ID)

	if err != nil {
		return fmt.Errorf("[DockRepoMysqlImpl.UpdateDock] Error when update dock: %w", err)
	}

	return nil
}

// GetDockByID is aaa
func (m *DockRepoMysqlImpl) GetDockByID(id int) (*model.Dock, error) {
	dock := model.Dock{}

	query := "SELECT id, kode, status FROM dock WHERE id = ? AND is_delete = 0"
	err := m.db.QueryRow(query, id).Scan(&dock.ID, &dock.Kode, &dock.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[DockRepoMysqlImpl.GetDockByID] Error when query get dock by id: %w", err)
	}

	return &dock, nil
}

// DeleteDock is aaa
func (m *DockRepoMysqlImpl) DeleteDock(id int) error {
	query := "UPDATE dock SET is_delete = '1' WHERE id = ?"

	_, err := m.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("[DockRepoMysqlImpl.DeleteDock] Error when delete dock: %w", err)
	}

	return nil
}

// GetAllDock is aaa
func (m *DockRepoMysqlImpl) GetAllDock() (*[]model.Dock, error) {
	query := "SELECT id, kode, status FROM dock WHERE is_delete = 0"
	rows, err := m.db.Query(query)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[DockRepoMysqlImpl.GetAllDock] Error when query get all dock: %w", err)
	}

	defer rows.Close()

	var sliceDock []model.Dock

	for rows.Next() {
		var each = model.Dock{}
		var err = rows.Scan(&each.ID, &each.Kode, &each.Status)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("[DockRepoMysqlImpl.GetAllDock] Error when scanning rows dock: %w", err)
		}

		sliceDock = append(sliceDock, each)
	}

	return &sliceDock, nil
}
