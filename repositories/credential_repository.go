package repositories

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/interfaces"
	"github.com/faizalnurrozi/go-crud/domain/models"
)

type CredentialRepository struct {
	DB *sql.DB
}

func NewCredentialRepository(DB *sql.DB) interfaces.ICredentialRepository {
	return &CredentialRepository{DB: DB}
}

func (repository CredentialRepository) ReadBy(column, value, operator string) (res models.Credential, err error) {
	model := models.NewCredential()
	statement := models.CredentialSelectStatement + ` ` + models.CredentialWhereStatement + ` and ` + column + `` + operator + `$1`
	row := repository.DB.QueryRow(statement, value)
	res, err = model.ScanRow(row)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (CredentialRepository) Add(model models.Credential, tx *sql.Tx) (res string, err error) {
	statement := `insert into credentials (email,password,is_active,created_at,updated_at) values($1,$2,$3,$4,$5) returning id`
	err = tx.QueryRow(statement, model.Email, model.Password, model.IsActive, model.CreatedAt, model.UpdatedAt).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (CredentialRepository) DeleteBy(column, value, operator string, model models.Credential, tx *sql.Tx) (err error) {
	statement := `update credentials set updated_at=$1, deleted_at=$2 where ` + column + `` + operator + `$3`
	_, err = tx.Exec(statement, model.UpdatedAt, model.DeletedAt.Time, value)
	if err != nil {
		return err
	}

	return nil
}

func (repository CredentialRepository) CountBy(column, value, operator string) (res int, err error) {
	statement := `select count(id) from credentials where ` + column + `` + operator + `$1 and deleted_at is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
