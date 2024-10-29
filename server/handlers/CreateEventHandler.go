package handlers

import (
	"L2/server/middleware"
	"L2/server/models"
	"L2/server/repository"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request, c *repository.Cache) {

	if r.Method != http.MethodPost {
		middleware.ErrorLogger(w, errors.New("method not supported"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var decoded models.Event

	if err := decoder.Decode(&decoded); err != nil {
		middleware.ErrorLogger(w, err)
		return
	}

	desc := decoded.Description
	date := decoded.Date
	id := decoded.UserId
	tt := decoded.Time

	if _, errParse := time.Parse("2006-01-02", date); errParse != nil {
		middleware.ErrorLogger(w, errParse)
		return
	}

	if _, errParse := time.Parse("15:00", tt); errParse != nil {
		middleware.ErrorLogger(w, errParse)
		return
	}

	ev := middleware.NewEvent(desc, date, tt, id)
	c.Create(ev)

	middleware.ResponseLogger(w, "event succesfully created")
}
