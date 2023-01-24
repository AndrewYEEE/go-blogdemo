package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model             //定義於models.go
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	log.Println("[DEBUG] GetTags : ", pageNum, pageSize, maps)
	var (
		tags []Tag
		err  error
	)

	if pageSize > 0 && pageNum > 0 { //是否有做分業設計
		err = db.Where(maps).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

func GetTagTotal(maps interface{}) (int, error) {
	log.Println("[DEBUG] GetTagTotal : ", maps)
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func CheckExistTagByName(name string) (bool, error) {
	log.Println("[DEBUG] CheckExistTagByName : ", name)
	var tag Tag
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func CheckExistTagByID(id int) (bool, error) {
	log.Println("[DEBUG] CheckExistTagByID : ", id)
	var tag Tag
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	log.Println("[DEBUG] AddTag : ", tag)
	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTag(id int) error {
	log.Println("[DEBUG] DeleteTag : ", id)
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

func EditTag(id int, data interface{}) error {
	log.Println("[DEBUG] EditTag : ", id, data)
	if err := db.Model(&Tag{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
