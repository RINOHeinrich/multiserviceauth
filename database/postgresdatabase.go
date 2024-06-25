package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"database/sql"

	"github.com/RINOHeinrich/multiserviceauth/models"
	"github.com/joho/godotenv"
)

type Postgres struct {
	Config models.Dbconfig
	DB     *sql.DB
}

// Connect to PostgreSQL
func (m *Postgres) Connect() error {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", m.Config.DBHost, m.Config.DBPort, m.Config.DBUser, m.Config.DBPassword, m.Config.DBName))
	if err != nil {
		return err
	}
	m.DB = db
	return nil
}

// Disconnect from PostgreSQL
func (m *Postgres) Disconnect() error {
	err := m.DB.Close()
	if err != nil {
		return err
	}
	return nil
}

// Implement the methods of the Database interface for PostgreSQL
// Insert an user
func (m *Postgres) Insert(data *models.User) error {
	stmt, err := m.DB.Prepare("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.Username, data.Email, data.Password)
	if err != nil {
		return err
	}
	return nil
}

// Delete an user
func (m *Postgres) Delete(id string) error {
	stmt, err := m.DB.Prepare("DELETE FROM users WHERE id = $1")
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
func (m *Postgres) Update(id string, data *models.User) error {
	stmt, err := m.DB.Prepare("UPDATE users SET Username = $1, email = $2, password = $3 WHERE id = $4")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.Username, data.Email, data.Password, id)
	if err != nil {
		return err
	}
	return nil
}

// Find an user by id
func (m *Postgres) Find(email string) (*models.User, error) {
	stmt, err := m.DB.Prepare("SELECT * FROM users WHERE email = $1")
	if err != nil {
		return nil, err
	}
	var user models.User
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Find all users
func (m *Postgres) FindAll() ([]models.User, error) {
	rows, err := m.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (p *Postgres) LoadConfig(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		log.Default().Println(err)
	}
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return err
	}

	p.Config.DBHost = os.Getenv("DB_HOST")
	p.Config.DBPort = port
	p.Config.DBUser = os.Getenv("DB_USER")
	p.Config.DBPassword = os.Getenv("DB_PASSWORD")
	p.Config.DBName = os.Getenv("DB_NAME")
	return nil
}
