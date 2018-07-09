package model

type Users struct {
	ID   string    `json:"id"`
	NAME string `json:"name"`
	AGE  string    `json:"age"`
	SEX  string    `json:"sex"`
}

type IDtype struct {
	ID string   `json:"id"`
}