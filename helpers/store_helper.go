package helpers

import "rishabhdev2700/goats/models"


func GetCategories() []models.Category{
    var categories []models.Category
    models.DB.Find(&categories)
    return categories
}
func GetArticles()[]models.Article{
    var articles []models.Article
    models.DB.Find(&articles)
    return articles
}

func GetTags()[]models.Tag{
    var tags []models.Tag
    models.DB.Find(&tags)
    return tags
}

func GetCategoryArticles(category models.Category)[]models.Article{
    var articles []models.Article
    models.DB.Where("category_id=?",category.ID).Find(&articles)
    return articles
}

func GetArticleTags(article models.Article)[]models.Tag{
    var tags []models.Tag
    models.DB.Model(&article).Association("Tags").Find(&tags)
    return tags 
}