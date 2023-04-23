package model

import (
	"github.com/Zhoangp/Course-Service/pkg/common"
)

type Courses struct {
	common.SQLModel
	Title             string        `json:"title" gorm:"column:title"`
	CourseDescription string        `json:"description" gorm:"column:description"`
	CourseLevel       string        `json:"level"  gorm:"column:level"`
	CourseLanguage    string        `json:"language" gorm:"column:language"`
	CoursePrice       float64       `json:"price" gorm:"column:price"`
	CourseCurrency    string        `json:"currency" gorm:"column:currency"`
	CourseDiscount    float32       `json:"discount" gorm:"column:discount"`
	CourseDuration    string        `json:"duration" gorm:"column:duration"`
	CourseStatus      string        `json:"status" gorm:"column:status"`
	CourseRating      float32       `json:"rating" gorm:"column:rating"`
	InstructorID      int           `json:"instructor_id" gorm:"column:instructor_id"`
	CourseThumbnail   *common.Image `json:"thumbnail" gorm:"column:thumbnail"`
}

func (Courses) TableName() string {
	return "Courses"
}
