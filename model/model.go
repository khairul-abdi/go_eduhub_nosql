package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Class   string `json:"class"`
}

type Course struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Schedule   string `json:"schedule"`
	Attendance int    `json:"attendance"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}
