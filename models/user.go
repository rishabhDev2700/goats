package models

import (
    "gorm.io/gorm"
)


type User struct{
    gorm.Model
    Name string `gorm:"size : 60"`
    Email string `gorm:"size:70;unique;not null;`
    Username string
    Password string
    Type int
    Location string
}