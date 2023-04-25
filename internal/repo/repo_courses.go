package repo

import (
	"fmt"
	"github.com/Zhoangp/Course-Service/internal/model"
	"gorm.io/gorm"
)

type coursesRepository struct {
	db *gorm.DB
}

func NewCoursesRepository(db *gorm.DB) *coursesRepository {
	return &coursesRepository{db}
}
func (c *coursesRepository) Create(course *model.Courses) error {
	db := c.db.Begin()
	if err := db.Table(model.Courses{}.TableName()).Create(course).Error; err != nil {
		fmt.Println(err)
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
func (c *coursesRepository) GetCourses(limit int, page int) (res []model.Courses, err error) {
	db := c.db.Model(&res)
	var total int64
	db.Count(&total)
	if limit <= 0 {
		db.Limit((int)(total)).Find(&res)
	} else {
		db.Limit(limit).Offset(limit * (page - 1)).Find(&res)
	}
	db.Preload("Images")

	err = nil
	return
}
func (c *coursesRepository) GetCourse(id int) (res model.Courses, err error) {
	db := c.db.Model(&res)

	if err = db.Where("id = ?", id).Preload("Sections.Lectures").First(&res).Error; err != nil {
		return
	}

	return
}
