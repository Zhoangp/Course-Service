package usecase

import (
	"github.com/Zhoangp/Course-Service/internal/model"
	"github.com/Zhoangp/Course-Service/pb"
	"github.com/Zhoangp/Course-Service/pkg/common"
	"net/http"
)

func (uc *coursesUseCase) UpdateCourse(rq *pb.UpdateCourseRequest) error {
	courseIdDecoded, err := uc.h.Decode(rq.CourseId)
	if err != nil {
		return err
	}
	instructorIdDecoded, err := uc.h.Decode(rq.InstructorId)
	if err != nil {
		return err
	}
	course, err := uc.repo.GetCourse(courseIdDecoded)
	if err != nil {
		return err
	}
	if course.InstructorID != instructorIdDecoded {
		return common.NewCustomError(err, http.StatusNotFound, "Your are not an instructor of this course!")
	}
	var img common.Image
	if rq.Thumbnail != nil {
		img.Url = rq.Thumbnail.Url
		img.Width = rq.Thumbnail.Width
		img.Height = rq.Thumbnail.Height
	}
	course.Title = rq.Title
	course.Description = rq.Description
	course.Level = rq.Level
	course.Language = rq.Language
	course.PriceID = int(rq.Price)
	course.IsPublish = false
	course.Thumbnail = &img
	course.IsPaid = rq.IsPaid
	course.Goals = rq.Goals
	course.Requirement = rq.Requirement
	var sections []model.Section
	for _, section := range rq.Sections {
		var lectures []model.Lecture
		for _, lecture := range section.Lectures {
			lectures = append(lectures, model.Lecture{
				Title:   lecture.Title,
				Content: lecture.Content,
				Status:  lecture.Status,
				Resource: model.Resource{
					Url:      lecture.Resource.Url,
					Duration: lecture.Resource.Duration,
				},
			})
		}
		sections = append(sections, model.Section{
			Title:    section.Title,
			Lectures: lectures,
		})
	}
	course.Sections = sections
	if err := uc.repo.UpdateCourse(&course); err != nil {
		return err
	}
	return nil
}
