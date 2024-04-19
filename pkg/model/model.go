package model

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StatusEnum string

const (
	Pending  StatusEnum = "pending"
	Approved StatusEnum = "approved"
	Rejected StatusEnum = "rejected"
)

type Complain struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" validate:"required"`
	ExhibitionID    primitive.ObjectID `bson:"exhibitionID,omitempty" json:"exhibitionID,omitempty"`
	ComplainMessage string             `bson:"message,omitempty" json:"message,omitempty" validate:"required"`
	User            User               `bson:"user" json:"user"`
	CreateDateAt    primitive.DateTime `bson:"createDateAt" json:"createDateAt" validate:"required"`
}

type ResponseComplain struct {
	ExhibitionID        primitive.ObjectID `bson:"exhibitionID,omitempty" json:"exhibitionID,omitempty" validate:"required"`
	ExhibitionName      string             `bson:"exhibitionName,omitempty" json:"exhibitionName,omitempty" validate:"required"`
	ExhibitionOrganizer UserExhibition     `bson:"exhibitionOrganizer,omitempty" json:"exhibitionOrganizer,omitempty" validate:"required"`
	ComplainMessage     []Complain         `bson:"complainmessage,omitempty" json:"complainmessage,omitempty" validate:"required"`
	Status              StatusEnum         `bson:"status,omitempty" json:"status,omitempty"`
	BanRequests         string             `bson:"banRequests,omitempty" json:"banRequests,omitempty"`
}

type RequestCreateComplain struct {
	ExhibitionID    primitive.ObjectID `bson:"exhibitionID,omitempty" json:"exhibitionID,omitempty" validate:"required"`
	ComplainMessage string             `bson:"message,omitempty" json:"message,omitempty" validate:"required"`
	CreateDateAt    primitive.DateTime `bson:"createDateAt" json:"createDateAt"`
	Status          StatusEnum         `bson:"status,omitempty" json:"status,omitempty"`
	User            User               `bson:"user" json:"user"`
}

type JwtCustomClaims struct {
	ID           primitive.ObjectID `json:"userID" bson:"userID"`
	Role         string             `json:"role"`
	UserName     string             `json:"username" bson:"username"`
	FirstName    string             `json:"firstname" bson:"firstname"`
	LastName     string             `json:"lastname" bson:"lastname"`
	ProfileImage string             `json:"profile,omitempty" bson:"profile,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	jwt.StandardClaims
}

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName    string             `json:"firstname" bson:"firstname"`
	LastName     string             `json:"lastname" bson:"lastname"`
	ProfileImage string             `json:"profile,omitempty" bson:"profile,omitempty"`
}

type UserExhibition struct {
	ID           primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
	FirstName    string             `json:"firstname" bson:"firstname"`
	LastName     string             `json:"lastname" bson:"lastname"`
	ProfileImage string             `json:"profile,omitempty" bson:"profile,omitempty"`
}

type Exhibition struct {
	ExhibitionID        primitive.ObjectID `bson:"exhibitionID,omitempty" json:"exhibitionID,omitempty"`
	ExhibitionName      string             `bson:"exhibitionName,omitempty" json:"exhibitionName,omitempty"`
	ExhibitionOrganizer UserExhibition     `bson:"userId,omitempty" json:"exhibitionOrganizer,omitempty"`
	Status              string             `bson:"status" json:"status" validate:"required" error:"status is required"`
}
