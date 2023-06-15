package https

import (
	"context"
	"github.com/Zhoangp/Course-Service/pb"
)

func (hdl *coursesHandler) GetCourseWithInstructor(ctx context.Context, req *pb.GetCourseWithInstructorRequest) (*pb.GetCourseWithInstructorResponse, error) {

	res, err := hdl.uc.GetCoursesWithInstructor(req)
	if err != nil {

		return &pb.GetCourseWithInstructorResponse{Error: HandleError(err)}, nil
	}
	return res, nil
}
