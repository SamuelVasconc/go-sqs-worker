package repositories_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SamuelVasconc/go-sqs-worker/models"
	"github.com/SamuelVasconc/go-sqs-worker/repositories"
	"github.com/stretchr/testify/assert"
)

func TestGenerateProtocol(t *testing.T) {
	db, mockSql, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	statement := `SELECT (replace(uuid(),'-',''))`

	rows := sqlmock.NewRows([]string{
		"(replace(uuid(),'-',''))",
	}).
		AddRow(
			1,
		)

	defer db.Close()

	t.Run("success", func(t *testing.T) {

		mockSql.ExpectQuery(statement).WithArgs().WillReturnRows(rows)

		repository := repositories.NewMySqlRepository(db)
		_, err := repository.GenerateProtocol()

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {

		mockSql.ExpectQuery(statement).WithArgs().WillReturnError(errors.New(""))

		repository := repositories.NewMySqlRepository(db)
		_, err := repository.GenerateProtocol()

		assert.Error(t, err)
	})
}

func TestSetProtocol(t *testing.T) {
	mockProtocol := "TESTE"
	mockstatus := "S"
	mockdate := "2020-09-01"
	mocklimit := 1000

	db, mockSql, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	statement := `UPDATE t_transactions
	SET protocol = ?
  WHERE status = ?
		  AND date = ?
		  LIMIT ?`

	defer db.Close()

	t.Run("success", func(t *testing.T) {

		mockSql.ExpectPrepare(statement).ExpectExec().WithArgs(mockProtocol, mockstatus, mockdate, mocklimit).WillReturnResult(sqlmock.NewResult(1, 1))

		repository := repositories.NewMySqlRepository(db)
		err := repository.SetProtocol(mockstatus, mockdate, mockProtocol, mocklimit)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {

		mockSql.ExpectPrepare(statement).ExpectExec().WithArgs(mockProtocol, mockstatus, mockdate, mocklimit).WillReturnError(errors.New(""))

		repository := repositories.NewMySqlRepository(db)
		err := repository.SetProtocol(mockstatus, mockdate, mockProtocol, mocklimit)

		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {

		mockSql.ExpectPrepare(statement).ExpectExec().WithArgs(mockProtocol, mockstatus, mockdate, mocklimit).WillReturnResult(sqlmock.NewResult(1, 1))

		db.Close()
		repository := repositories.NewMySqlRepository(db)
		err := repository.SetProtocol(mockstatus, mockdate, mockProtocol, mocklimit)

		assert.Error(t, err)
	})
}

func TestUpdateLine(t *testing.T) {
	mockstatus := "Z"
	mockobservacao := "registro liberado"
	var mockid int64
	mockid = 321123321

	db, mockSql, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	statement := `UPDATE t_transactions
					SET observation = ?,
						status = ?
					WHERE id = ?`

	defer db.Close()

	t.Run("success", func(t *testing.T) {
		mockSql.ExpectPrepare(statement).ExpectExec().WithArgs(mockobservacao, mockstatus, mockid).WillReturnResult(sqlmock.NewResult(1, 1))

		repository := repositories.NewMySqlRepository(db)
		err := repository.UpdateLine(mockobservacao, mockstatus, mockid)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {

		mockSql.ExpectPrepare(statement).ExpectExec().WithArgs(mockobservacao, mockstatus, mockid).WillReturnError(errors.New(""))

		repository := repositories.NewMySqlRepository(db)
		err := repository.UpdateLine(mockobservacao, mockstatus, mockid)

		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {

		mockSql.ExpectPrepare(statement).ExpectExec().WithArgs(mockobservacao, mockstatus, mockid).WillReturnResult(sqlmock.NewResult(1, 1))

		db.Close()
		repository := repositories.NewMySqlRepository(db)
		err := repository.UpdateLine(mockobservacao, mockstatus, mockid)

		assert.Error(t, err)
	})
}

func TestGetLines(t *testing.T) {
	mocktransaction := new(models.Transaction)
	protocol := "TESTE"

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{
		"id",
		"date",
		"amount",
		"obs",
		"status",
	}).
		AddRow(
			mocktransaction.ID,
			mocktransaction.Date,
			mocktransaction.Amount,
			mocktransaction.Observation,
			mocktransaction.Status,
		)

	query := `SELECT id, 
					date, 
					amount, 
					observation, 
					status 
				FROM t_transactions
				WHERE protocol = ?`

	t.Run("success", func(t *testing.T) {
		mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(rows)

		repository := repositories.NewMySqlRepository(db)

		_, err = repository.GetLines(protocol)
		assert.NoError(t, err)
	})

	t.Run("error-scan", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id",
			"date",
			"amount",
			"obs",
			"status",
			"error",
		}).
			AddRow(
				mocktransaction.ID,
				mocktransaction.Date,
				mocktransaction.Amount,
				mocktransaction.Observation,
				mocktransaction.Status,
				"",
			)
		mock.ExpectPrepare(query).ExpectExec().WithArgs(protocol).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)
		repository := repositories.NewMySqlRepository(db)

		_, err = repository.GetLines(protocol)
		assert.Error(t, err)
	})

	t.Run("error-database", func(t *testing.T) {
		mock.ExpectPrepare(query).ExpectExec().WithArgs(protocol).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery(query).WithArgs(protocol).WillReturnRows(rows)
		db.Close()
		repository := repositories.NewMySqlRepository(db)

		_, err = repository.GetLines(protocol)
		assert.Error(t, err)
	})
}
