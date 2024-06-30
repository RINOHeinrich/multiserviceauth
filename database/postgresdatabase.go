package database

import (
	"fmt"

	"database/sql"

	"github.com/RINOHeinrich/multiserviceauth/models"
)

type Postgres struct {
	Config models.Dbconfig
	DB     *sql.DB
}

// Connect to PostgreSQL
func (p *Postgres) Connect() error {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Config.DBHost, p.Config.DBPort, p.Config.DBUser, p.Config.DBPassword, p.Config.DBName))
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	p.DB = db
	return nil
}

// Disconnect from PostgreSQL
func (p *Postgres) Disconnect() error {
	err := p.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

// Implement the methods of the Database interface for PostgreSQL
// Insert an user
func (p *Postgres) Insert(data *models.User) error {
	stmt, err := p.DB.Prepare("INSERT INTO users (login, password) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.Login, data.Password)
	if err != nil {
		return err
	}
	return nil
}

// Delete an user
func (p *Postgres) Delete(id string) error {
	stmt, err := p.DB.Prepare("DELETE FROM users WHERE login = $1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// Update an user
func (p *Postgres) Update(id string, data *models.User) error {
	stmt, err := p.DB.Prepare("UPDATE users SET login = $1, password = $2 WHERE login = $3")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.Login, data.Password, id)
	if err != nil {
		return err
	}
	return nil
}

// Find an user by id
func (p *Postgres) Find(login string) (*models.User, error) {
	var user models.User
	stmt, err := p.DB.Prepare("SELECT * FROM users WHERE login = $1")
	if err != nil {
		return &user, err
	}

	err = stmt.QueryRow(login).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Find all users
func (p *Postgres) FindAll() ([]models.User, error) {
	rows, err := p.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Login, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (p *Postgres) LoadConfig(m *models.Dbconfig) {
	p.Config = *m
}
