package repositories

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mtti/board/server/models"
	"github.com/mtti/board/server/utils"

)

type MySQL struct {
	DB *sqlx.DB
	Parent Repository
}

// Create a new PostgreSQL repository instance.
func NewMySQL(connectionString string) (*MySQL, error) {

	if connectionString == "" {
		return nil, errors.New("Database connection string is required")
	}

	db, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	repo := &MySQL{ DB: db }
	repo.Parent = repo
	return repo, nil

}

func (repo *MySQL) SetParent(parent Repository) {
	repo.Parent = parent
}

func (repo *MySQL) LoadCardByID(id int) (*models.Card, error) {
	obj := &models.Card{}
	err := repo.DB.Get(obj, "SELECT * FROM card WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (repo *MySQL) LoadCardByName(name string) (*models.Card, error) {
	
	key, number, err := utils.ParseCardName(name)
	if err != nil {
		return nil, err
	}

	project, err := repo.Parent.LoadProjectByKey(key)
	if err != nil {
		return nil, err
	}

	obj := &models.Card{}
	err = repo.DB.Get(obj,
		"SELECT * FROM card WHERE `project` = ? AND `number` = ?",
		project.Id, number)
	if err != nil {
		return nil, err
	}

	return obj, nil

}

func (repo *MySQL) LoadProjectByKey(key string) (*models.Project, error) {
	obj := &models.Project{}
	err := repo.DB.Get(obj, "SELECT * FROM project WHERE `key` = ?", key)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
