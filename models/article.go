package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTotalArticle(maps interface{}) (int, error) {
	log.Println("[DEBUG] GetTotalArticle : ", maps)
	var count int
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func GetArticlesByPages(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	log.Println("[DEBUG] GetArticlesByPages : ", pageNum, pageSize, maps)
	var articles []*Article
	err := db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error //Preload會先去查Tag資料表，再去看Article資料表
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

func GetArticleById(id int) (*Article, error) {
	log.Println("[DEBUG] GetArticleById : ", id)
	var article Article
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	err = db.Model(&article).Related(&article.Tag).Error //關聯tag資料表
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}

func EditArticleById(id int, data interface{}) error {
	log.Println("[DEBUG] EditArticleById : ", id)
	if err := db.Model(&Article{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func AddArticle(data map[string]interface{}) error {
	log.Println("[DEBUG] AddArticle : ", data)
	article := Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}
	if err := db.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

func DeleteArticleById(id int) error {
	log.Println("[DEBUG] DeleteArticleById : ", id)
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}

	return nil
}

func CheckExistArticleByID(id int) (bool, error) {
	log.Println("[DEBUG] CheckExistArticleByID : ", id)
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}
