package repositories

import (
	"database/sql"

	"github.com/SamuelVasconc/go-sqs-worker/interfaces"
	"github.com/SamuelVasconc/go-sqs-worker/models"
)

type mysqlRepository struct {
	Conn *sql.DB
}

func NewMySqlRepository(Conn *sql.DB) interfaces.MySqlRepository {
	return &mysqlRepository{Conn}
}

func (m *mysqlRepository) GenerateProtocol() (string, error) {

	var uuid string
	err := m.Conn.QueryRow(`SELECT (replace(uuid(),'-',''))`).Scan(&uuid)
	if err != nil {
		return "", err
	}

	return uuid, err
}

func (m *mysqlRepository) SetProtocol(status, date, protocol string, limit int) error {
	query := `UPDATE t_transactions
							SET protocol = ?
						  WHERE status = ?
								  AND date = ?
								  LIMIT ?`

	smt, err := m.Conn.Prepare(query)
	if err != nil {
		return err
	}

	defer smt.Close()
	_, err = smt.Exec(protocol, status, date, limit)
	if err != nil {
		return err
	}

	return err
}

func (m *mysqlRepository) GetLines(protocol string) ([]*models.Transaction, error) {

	query := `SELECT id, 
					date, 
					amount, 
					obs, 
					status 
				FROM t_transactions
				WHERE protocol = ?`

	smt, err := m.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := smt.Query(protocol)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*models.Transaction, 0)

	for rows.Next() {
		line := &models.Transaction{}

		var (
			id     sql.NullInt64
			date   sql.NullString
			amount sql.NullFloat64
			obs    sql.NullString
			status sql.NullString
		)

		if err := rows.Scan(&id, &date, &amount, &obs, &status); err != nil {
			return nil, err
		}

		line.ID = id.Int64
		line.Date = date.String
		line.Amount = amount.Float64
		line.Observation = obs.String
		line.Status = status.String

		result = append(result, line)
	}

	return result, nil
}

func (m *mysqlRepository) UpdateLine(obs, status string, id int64) error {
	query := `UPDATE t_transactions
				SET observation = ?,
					status = ?
				WHERE id = ?`

	smt, err := m.Conn.Prepare(query)
	if err != nil {
		return err
	}

	defer smt.Close()

	_, err = smt.Exec(obs, status, id)
	if err != nil {
		return err
	}

	return err
}
