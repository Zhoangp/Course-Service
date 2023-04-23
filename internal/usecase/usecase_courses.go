package usecase

import (
	"github.com/Zhoangp/Course-Service/internal/model"
	"github.com/Zhoangp/Course-Service/pb"
	"github.com/Zhoangp/Course-Service/pkg/utils"
	"strconv"
)

type CoursesRepository interface {
	Create(course *model.Courses) error
	GetCourses(limit int, page int) (res []model.Courses, err error)
}
type coursesUseCase struct {
	repo CoursesRepository
	h    *utils.Hasher
}

func NewCoursesUseCase(repo CoursesRepository, h *utils.Hasher) *coursesUseCase {
	return &coursesUseCase{repo, h}
}
func (uc *coursesUseCase) CreateCourse(data *model.Courses) error {

	if err := uc.repo.Create(data); err != nil {
		return err
	}

	return nil
}
func (uc *coursesUseCase) GetCourses(limit int, page int) (courses []*pb.Course, err error) {
	res, err := uc.repo.GetCourses(limit, page)
	for i, _ := range res {
		res[i].FakeId = uc.h.Encode(res[i].Id)
		courses = append(courses, &pb.Course{
			Id:          res[i].FakeId,
			Title:       res[i].Title,
			Description: res[i].CourseDescription,
			Level:       res[i].CourseLevel,
			Language:    res[i].CourseLanguage,
			Price:       strconv.FormatFloat(res[i].CoursePrice, 'f', -1, 64),
			Discount:    res[i].CourseDiscount,
			Currency:    res[i].CourseCurrency,
			Duration:    res[i].CourseDuration,
			Status:      res[i].CourseStatus,
			Rating:      res[i].CourseRating,
			Thumbnail: &pb.Image{
				Url:    res[i].CourseThumbnail.Url,
				Width:  res[i].CourseThumbnail.Width,
				Height: res[i].CourseThumbnail.Height,
			},
		})
	}
	return

}
