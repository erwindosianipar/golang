package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"api-pelabuhan/kapal"
	"api-pelabuhan/model"
)

var _ kapal.KapalRepo = &KapalRepoMysqlImpl{}

// KapalRepoMysqlImpl is aaa
type KapalRepoMysqlImpl struct {
	db *sql.DB
}

// CreateKapalRepoMysqlImpl is aaa
func CreateKapalRepoMysqlImpl(db *sql.DB) kapal.KapalRepo {
	return &KapalRepoMysqlImpl{db}
}

// GetAllKapal is aaa
func (m *KapalRepoMysqlImpl) GetAllKapal() (*[]model.Kapal, error) {
	query := "SELECT id, kode, muatan, status FROM kapal WHERE is_delete = 0"
	rows, err := m.db.Query(query)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[KapalRepoMysqlImpl.GetAllKapal] Error when query get all kapal: %w", err)
	}

	defer rows.Close()

	var sliceKapal []model.Kapal

	for rows.Next() {
		var each = model.Kapal{}
		var err = rows.Scan(&each.ID, &each.Kode, &each.Muatan, &each.Status)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("[KapalRepoMysqlImpl.GetAllKapal] Error when scanning rows kapal: %w", err)
		}

		sliceKapal = append(sliceKapal, each)
	}

	return &sliceKapal, nil
}

// InsertKapal is aaaa
func (m *KapalRepoMysqlImpl) InsertKapal(kapal *model.Kapal) error {
	tx, err := m.db.Begin()
	if err != nil {
		return fmt.Errorf("[MejaRepoMysqlImpl.InsertKapal] Error when begin transaction : %w", err)
	}

	query := "INSERT INTO kapal (kode, muatan, status, is_delete) VALUES (?, ?, ?, ?)"
	_, err = m.db.Exec(query, kapal.Kode, kapal.Muatan, kapal.Status, 0)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[MejaRepoMysqlImpl.InsertKapal] Error when running query insert kapal: %w", err)
	}

	tx.Commit()
	return nil
}

// GetKapalByID is aaa
func (m *KapalRepoMysqlImpl) GetKapalByID(id int) (*model.Kapal, error) {
	kapal := model.Kapal{}

	query := "SELECT id, kode, muatan, status FROM kapal WHERE id = ? AND is_delete = 0"
	err := m.db.QueryRow(query, id).Scan(&kapal.ID, &kapal.Kode, &kapal.Muatan, &kapal.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[KapalRepoMysqlImpl.GetKapalByID] Error when query get kapal by id: %w", err)
	}

	return &kapal, nil
}

// DeleteKapal is aaa
func (m *KapalRepoMysqlImpl) DeleteKapal(id int) error {
	query := "UPDATE kapal SET is_delete = '1' WHERE id = ?"

	_, err := m.db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("[KapalRepoMysqlImpl.DeleteKapal] Error when delete kapal: %w", err)
	}

	return nil
}

// UpdateKapal is aaa
func (m *KapalRepoMysqlImpl) UpdateKapal(kapal *model.Kapal) error {
	query := "UPDATE kapal SET kode = ?, muatan = ?, status = ? WHERE id = ?"

	_, err := m.db.Exec(query, kapal.Kode, kapal.Muatan, kapal.Status, kapal.ID)

	if err != nil {
		return fmt.Errorf("[KapalRepoMysqlImpl.UpdateKapal] Error when update kapal: %w", err)
	}

	return nil
}
