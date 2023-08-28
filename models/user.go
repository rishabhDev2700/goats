package models

import (
    "gorm.io/gorm"
)


type User struct{
    gorm.Model
    Name string `gorm:"size : 60" json:"name"`
    Email string `gorm:"size:70;unique;not null;" json:"email"`
    Username string `gorm:"unique;not null;"`
    Password string `gorm:"not null"`
    IsAdmin bool
    Location string 
    GoogleID string `gorm:"not null" json:"id"`
}