package https

import (
	"context"
	"github.com/Zhoangp/Course-Service/pb"
)

func (hdl *coursesHandler) GetCourseContent(ctx context.Context, request *pb.GetCourseContentRequest) (*pb.GetCourseContentResponse, error) {
	res, err := hdl.uc.GetCourseContent(request.UserId, request.CourseId)
	if err != nil {
		return &pb.GetCourseContentResponse{
			Error: HandleError(err),
		}, nil
	}
	return res, nil
}
