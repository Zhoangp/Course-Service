package repo

import (
	"github.com/Zhoangp/Course-Service/internal/model"
)

func (c *coursesRepository) UpdateCourse(course *model.Course) error {
	db := c.db.Begin()
	if err := db.Save(course).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
