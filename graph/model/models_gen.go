// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Location struct {
    State   string `json:"state" bson:"state"`
    City    string `json:"city" bson:"city"` 
    ZipCode int    `json:"zip_code" bson:"zip_code"`
}

type Name struct {
    FirstName string `json:"first_name" bson:"first_name"`
    LastName  string `json:"last_name" bson:"last_name"`
}

type NewProfile struct {
    FirstName   *string `json:"first_name" bson:"first_name"`
    LastName    *string `json:"last_name" bson:"last_name"`
    State       *string `json:"state" bson:"state"`
    City        *string `json:"city" bson:"city"`
    ZipCode     *int    `json:"zip_code" bson:"zip_code"`
    Email       *string `json:"email" bson:"email"`
    Description *string `json:"description" bson:"career_description"`
}

type ProfileQl struct {
    ID          string    `json:"id" bson:"id,omitempty"`
    Name        *Name     `json:"name" bson:"full_name"`
    Location    *Location `json:"location" bson:"location"`
    Email       string    `json:"email" bson:"email"`
    Skills      []*Skill  `json:"skills" bson:"skills_list"`
    Description string    `json:"description" bson:"career_description"`
}

type Skill struct {
    ID          string `json:"id" bson:"id"`
    Name        string `json:"name" bson:"string_name"`
    Description string `json:"Description" bson:"skill_description"`
}