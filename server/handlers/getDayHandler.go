package handlers

import (
	"L2/server/middleware"
	"L2/server/repository"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func GetEventDayHandler(w http.ResponseWriter, r *http.Request, c *repository.Cache) {
	if r.Method != http.MethodGet {
		middleware.ErrorLogger(w, errors.New("method not supported"))
		return
	}

	dateQuery := r.URL.Query().Get("date")

	if _, errParse := time.Parse("2006-01-02", dateQuery); errParse != nil {
		middleware.ErrorLogger(w, errParse)
		return
	}

	events := c.GetDay(dateQuery)
	resp, err := json.MarshalIndent(events, "", "\t")
	if err != nil {
		middleware.ErrorLogger(w, err)
		return
	}
	w.Write(resp)
}
