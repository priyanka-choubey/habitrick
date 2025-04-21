package handlers

import (
	"fmt"

	dbmanager "github.com/priyanka-choubey/habitrick/services/backend/database"
)

type Milestone struct {
	goal   string
	date   string
	reward string
}

type Habit struct {
	id         string
	name       string
	intention  string
	start_date string // start date could be any future date from now
	end_date   string // Optional
	milestones []Milestone
}

func CreateHabit(habit_args map[string]string) status {

	var resp status

	db, err := dbmanager.SetupDatabase()
	if err != nil {
		return resp.getStatus(500, "Internal error")
	}

	err = db.CreateUserLoginDetails(username, token)
	if err != nil {
		return resp.getStatus(403, fmt.Sprintf("%v", err))
	}

	resp = resp.getStatus(200, "OK")
	defer db.Db.Close()
	return resp
}
