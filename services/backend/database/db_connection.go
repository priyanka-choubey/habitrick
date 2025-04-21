package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

// Database collections
type LoginDetails struct {
	AuthToken string
	Username  string
}

type HabitDetails struct {
	Name      string
	Intention string
	StartDate time.Time
}

type MySqlDatabase struct {
	Db *sql.DB
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) (LoginDetails, error)
	SetupDatabase() error
}

func NewDatabase() (*MySqlDatabase, error) {
	var database MySqlDatabase

	err := database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}

// USER

func (d *MySqlDatabase) CreateUserLoginDetails(username string, token string) (*LoginDetails, error) {

	var clientData LoginDetails

	d.Db.QueryRow("INSERT INTO user (username,token) VALUES (?,?)", username, token)
	row := d.Db.QueryRow("SELECT username,token FROM user WHERE username = ?", username)
	if err := row.Scan(&clientData.Username, &clientData.AuthToken); err != nil {
		return &clientData, fmt.Errorf("User %d: Cannot create user: %v", username, err)
	}
	return &clientData, nil
}

func (d *MySqlDatabase) DeleteUserLoginDetails(username string) error {

	var clientData LoginDetails

	d.Db.QueryRow("DELETE FROM user WHERE username = ?", username)
	row := d.Db.QueryRow("SELECT (username,token) FROM user WHERE username = ?", username)
	if err := row.Scan(&clientData.Username, &clientData.AuthToken); err != nil {
		return nil
	}
	return fmt.Errorf("User %d: Cannot delete user", clientData.Username)
}

func (d *MySqlDatabase) UpdateUserLoginDetails(username string, token string) error {

	var clientData LoginDetails

	d.Db.QueryRow("UPDATE user SET token = ? WHERE username = ?", token, username)
	row := d.Db.QueryRow("SELECT username,token FROM user WHERE username = ?", username)
	if err := row.Scan(&clientData.Username, &clientData.AuthToken); err != nil {
		return fmt.Errorf("Unexpected Error: %v", err)
	}

	if clientData.AuthToken != token {
		return fmt.Errorf("User %d: Cannot update user token", clientData.Username)
	}
	return nil
}

func (d *MySqlDatabase) GetUserLoginDetails(username string) (*LoginDetails, error) {

	var clientData LoginDetails

	row := d.Db.QueryRow("SELECT username,token FROM user WHERE username = ?", username)
	if err := row.Scan(&clientData.Username, &clientData.AuthToken); err != nil {
		if err == sql.ErrNoRows {
			return &clientData, fmt.Errorf("User %d: no such user", username)
		}
		return &clientData, fmt.Errorf("User %d: %v", username, err)
	}
	return &clientData, nil
}

// HABIT

func (d *MySqlDatabase) CreateHabit(name string, intention string) (*HabitDetails, error) {

	var habitData HabitDetails

	d.Db.QueryRow("INSERT INTO habit (name,intention) VALUES (?,?)", name, intention)
	row := d.Db.QueryRow("SELECT name,intention FROM habit WHERE name = ?", name)
	if err := row.Scan(&habitData.Name, &habitData.Intention); err != nil {
		return &habitData, fmt.Errorf("User %d: Cannot create habit: %v", name, err)
	}
	return &habitData, nil
}

func (d *MySqlDatabase) GetHabitDetails(name string, intention string) (*HabitDetails, error) {

	var habitData HabitDetails

	row := d.Db.QueryRow("SELECT name,intention,start_time FROM habit WHERE name = ?", name)
	if err := row.Scan(&habitData.Name, &habitData.Intention, &habitData.StartDate); err != nil {
		if err == sql.ErrNoRows {
			return &habitData, fmt.Errorf("Habit %d: no such habit under user", name)
		}
		return &habitData, fmt.Errorf("Habit %d: %v", name, err)
	}
	return &habitData, nil
}

func (d *MySqlDatabase) SetupDatabase() error {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "dev",
	}
	// Get a database handle.
	var err error
	d.Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := d.Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Debug("Connected to the databse!")
	return nil
}
