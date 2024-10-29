package middleware

import (
	"L2/server/models"
	"encoding/json"
	"log"
	"net/http"
	time2 "time"
)

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time2.Now()
		handler.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time2.Since(start))
	})
}

func NewEvent(des, date, time string, userId int) *models.Event {
	uid := time2.Now().Unix()
	return &models.Event{
		UserId:      userId,
		Date:        date,
		Uid:         uid,
		Description: des,
		Time:        time,
	}
}

func ErrorLogger(w http.ResponseWriter, errorInput error) {
	log.Println(errorInput)
	w.WriteHeader(http.StatusInternalServerError)
	details := models.Details{ErrCode: http.StatusInternalServerError, ErrMessage: errorInput.Error()}
	response, _ := json.MarshalIndent(models.Error{Err: details}, "", "\t")
	w.Write(response)
}

func ResponseLogger(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	result := models.Result{StatusCode: http.StatusOK, Message: message}
	response, _ := json.MarshalIndent(models.Response{Body: result}, "", "\t")
	w.Write(response)
}
