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
	Sections          []Section     `json:"sections" gorm:"foreignKey:CourseId;references:Id"`
}
type Section struct {
	common.SQLModel
	Title            string    `gorm:"column:title"`
	CourseId         int       `gorm:"column:course_id"`
	NumberOfLectures int       `gorm:"column:numberOfLectures"`
	Lectures         []Lecture `json:"lectures" gorm:"foreignKey:SectionId;references:Id"`
}
type Lecture struct {
	common.SQLModel
	Title           string           `json:"title" gorm:"column:title"`
	Content         string           `json:"content" gorm:"column:content"`
	Status          string           `json:"status" gorm:"column:status"`
	SectionId       int              `json:"section_id" gorm:"column:section_id"`
	LectureResource LectureResources `gorm:"foreignKey:LectureId;references:Id"`
}
type LectureResources struct {
	common.SQLModel
	Url       string `gorm:"column:url"`
	Duration  string `gorm:"column:duration"`
	LectureId int    `gorm:"lecture_id"`
}

func (LectureResources) TableName() string {
	return "LectureResources"
}
func (Lecture) TableName() string {
	return "Lectures"
}
func (Section) TableName() string {
	return "Sections"
}
func (Courses) TableName() string {
	return "Courses"
}
