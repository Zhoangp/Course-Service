package https

import (
	"context"
	"github.com/Zhoangp/Course-Service/pb"
)

func (hdl *coursesHandler) PublishCourse(ctx context.Context, req *pb.PublishCourseRequest) (*pb.PublishCourseResponse, error) {
	errorsValidate, err := hdl.uc.PublishCourse(req)
	if errorsValidate != nil {
		return &pb.PublishCourseResponse{ErrorsValidate: errorsValidate}, nil
	}
	if err != nil {
		return &pb.PublishCourseResponse{Error: HandleError(err)}, nil
	}
	return &pb.PublishCourseResponse{}, nil
}
