package https

import (
	"context"
	"github.com/Zhoangp/Course-Service/pb"
)

func (hdl *coursesHandler) CreateCourse(ctx context.Context, rq *pb.CreateCourseRequest) (*pb.CreateCourseResponse, error) {
	courseId, err := hdl.uc.NewCourse(rq, rq.InstructorId)
	if err != nil {
		return &pb.CreateCourseResponse{
			Error: HandleError(err),
		}, nil
	}
	return &pb.CreateCourseResponse{
		CourseId: courseId,
	}, nil
}
