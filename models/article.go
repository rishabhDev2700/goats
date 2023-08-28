package models

import "gorm.io/gorm"


type Article struct{
    gorm.Model
    Name string `gorm:"unqiue"`
    Description string `gorm:"not null"`
    Price uint
    InStock bool `gorm:"default:false"`
    CategoryID uint
    Tags []*Tag `gorm:"many2many:article_tags;"`
}