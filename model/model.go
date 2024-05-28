package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type Student struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Email   string             `bson:"email"`
	Address string             `bson:"address"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}
