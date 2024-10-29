package models

type Event struct {
	Uid         int64  `json:"uid"`
	UserId      int    `json:"user_id"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Time        string `json:"time"`
}

type Error struct {
	Err Details `json:"error"`
}

type Details struct {
	ErrCode    int    `json:"code"`
	ErrMessage string `json:"message"`
}

type Response struct {
	Body Result `json:"result"`
}

type Result struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}
