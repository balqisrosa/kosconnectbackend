package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FullName string             `bson:"fullname" json:"fullname"`
	Email    string             `bson:"email" json:"email"`
	Role     string             `bson:"role" json:"role"`
	Password string             `bson:"password" json:"password"`
}

<<<<<<< HEAD
type Category struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string             `bson:"name" json:"name"`
	Slug string             `bson:"slug" json:"slug"`
}
=======
type CustomFacility struct {
    ID    string  `json:"id" bson:"_id"`
    Name  string  `json:"name" bson:"name"`
    Price float64 `json:"price" bson:"price"`
}
>>>>>>> f2fce078d6b6d3e4d4f618c2160ccd72a7082cb8
