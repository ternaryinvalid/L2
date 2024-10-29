package handlers

import (
	"L2/server/middleware"
	"L2/server/models"
	"L2/server/repository"
	"encoding/json"
	"errors"
	"net/http"
)

func DeleteEventHandler(w http.ResponseWriter, r *http.Request, c *repository.Cache) {

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

	date := decoded.Date
	time := decoded.Time

	c.Delete(date, time)

	middleware.ResponseLogger(w, "event succesfully deleted")

}
