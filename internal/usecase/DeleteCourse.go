package usecase

import (
	"github.com/Zhoangp/Course-Service/pb"
	"github.com/Zhoangp/Course-Service/pkg/common"
	"net/http"
)

func (uc *coursesUseCase) DeleteCourse(req *pb.DeleteCourseRequest) error {
	userIdDecoded, err := uc.h.Decode(req.UserId)
	if err != nil {
		return err
	}
	courseIdDecoded, err := uc.h.Decode(req.CourseId)
	if err != nil {
		return err
	}
	course, err := uc.repo.GetCourse(courseIdDecoded)
	if err != nil {
		return err
	}
	if course.InstructorID != userIdDecoded {
		return common.NewCustomError(err, http.StatusNotFound, "Your are not allow!")
	}
	if err := uc.repo.DeleteCourse(course); err != nil {
		return common.NewCustomError(err, http.StatusBadRequest, "Unable delete course!")
	}
	return nil
}
