package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/priyanka-choubey/habitrick/services/backend/api"
	db "github.com/priyanka-choubey/habitrick/services/backend/database"
	log "github.com/sirupsen/logrus"
)

type Milestone struct {
	goal   string
	date   string
	reward string
}

type Habit struct {
	Username   string
	Name       string
	Intention  string
	Start_date string // start date could be any future date from now
	End_date   string // Optional
	Milestones []Milestone
}

func CreateHabit(w http.ResponseWriter, r *http.Request) {
	var params = Habit{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	if params.Name == "" || params.Intention == "" {
		log.Errorf("Habit params not as expected. Please check and update them name= %d and intention= %d.", params.Name, params.Intention)
		return
	}

	var database *db.MySqlDatabase
	database, err = db.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	_, err = (*database).GetHabitDetails(params.Name, params.Intention)
	if err == nil {
		log.Error("Habit was not created")
		return
	} else {

		habitDetails, err := (*database).CreateHabit(params.Name, params.Intention)
		if err != nil {
			api.RequestErrorHandler(w, err)
			return
		}

		if habitDetails == nil {
			api.InternalErrorHandler(w)
			return
		}

		var response = api.HabitResponse{
			Name:      habitDetails.Name,
			Intention: habitDetails.Intention,
			StartDate: habitDetails.StartDate,
			Code:      http.StatusOK,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return
		}
	}

}
