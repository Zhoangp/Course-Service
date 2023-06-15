package usecase

import (
	"github.com/Zhoangp/Course-Service/pb"
	"github.com/Zhoangp/Course-Service/pkg/common"
	"github.com/asaskevich/govalidator"
	"net/http"
)

func (uc *coursesUseCase) PublishCourse(rq *pb.PublishCourseRequest) ([]string, error) {
	courseIdDecoded, err := uc.h.Decode(rq.CourseId)
	userIdDecoded, err := uc.h.Decode(rq.UserId)
	if err != nil {
		return nil, err
	}
	course, err := uc.repo.GetCourse(courseIdDecoded)
	if err != nil {
		return nil, err
	}
	if course.InstructorID != userIdDecoded {
		return nil, common.NewCustomError(err, http.StatusNotFound, "Your are not an instructor of this course!")
	}
	if _, err = govalidator.ValidateStruct(course); err != nil {
		var errors []string
		for _, e := range err.(govalidator.Errors) {
			errors = append(errors, e.Error())
		}
		return errors, nil
	}
	if err = uc.repo.PublishCourse(&course); err != nil {
		return nil, err
	}
	return nil, nil
}
