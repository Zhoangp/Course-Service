package usecase

import (
	"fmt"
	"github.com/Zhoangp/Course-Service/internal/model"
	"github.com/Zhoangp/Course-Service/pb"
	"github.com/Zhoangp/Course-Service/pkg/utils"
	"strconv"
)

type CoursesRepository interface {
	Create(course *model.Courses) error
	GetCourses(limit int, page int) (res []model.Courses, err error)
	GetCourse(id int) (res model.Courses, err error)
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
func (uc *coursesUseCase) GetCourse(fakeId string) (*pb.Course, error) {
	id, err := uc.h.Decode(fakeId)
	if err != nil {
		return nil, err
	}
	course, err := uc.repo.GetCourse(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(course.Sections[1].Lectures)
	course.FakeId = uc.h.Encode(course.Id)
	var sections []*pb.Section
	for _, i := range course.Sections {
		var lectures []*pb.Lecture
		for _, j := range i.Lectures {
			j.FakeId = uc.h.Encode(j.Id)
			lectures = append(lectures, &pb.Lecture{
				Id:      j.FakeId,
				Title:   j.Title,
				Content: j.Content,
				Status:  j.Status,
				Video: &pb.Resource{
					Url:      j.LectureResource.Url,
					Duration: j.LectureResource.Duration,
				},
			})

		}
		i.FakeId = uc.h.Encode(i.Id)
		sections = append(sections, &pb.Section{
			Id:               i.FakeId,
			Title:            i.Title,
			NumberOfLectures: int32(i.NumberOfLectures),
			Lectures:         lectures,
		})
	}
	res := &pb.Course{
		Id:          course.FakeId,
		Title:       course.Title,
		Description: course.CourseDescription,
		Level:       course.CourseLevel,
		Language:    course.CourseLanguage,
		Price:       strconv.FormatFloat(course.CoursePrice, 'f', -1, 64),
		Discount:    course.CourseDiscount,
		Currency:    course.CourseCurrency,
		Duration:    course.CourseDuration,
		Status:      course.CourseStatus,
		Rating:      course.CourseRating,
		Sections:    sections,
	}

	return res, nil
}
