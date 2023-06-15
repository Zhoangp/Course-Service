package https

import (
	"context"
	"github.com/Zhoangp/Course-Service/pb"
)

func (hdl *coursesHandler) UpdateCourse(ctx context.Context, rq *pb.UpdateCourseRequest) (*pb.UpdateCourseResponse, error) {
	if err := hdl.uc.UpdateCourse(rq); err != nil {
		return &pb.UpdateCourseResponse{Error: HandleError(err)}, nil
	}

	return &pb.UpdateCourseResponse{}, nil
}
