package https

import (
	"context"
	"github.com/Zhoangp/Course-Service/pb"
)

func (hdl *coursesHandler) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	if err := hdl.uc.DeleteCourse(req); err != nil {
		return &pb.DeleteCourseResponse{Error: HandleError(err)}, nil
	}
	return &pb.DeleteCourseResponse{}, nil
}
