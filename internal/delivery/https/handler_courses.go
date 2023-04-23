package https

import (
	"context"
	"github.com/Zhoangp/Course-Service/internal/model"
	"github.com/Zhoangp/Course-Service/pb"
	"github.com/Zhoangp/Course-Service/pkg/common"
)

type CoursesUseCase interface {
	CreateCourse(data *model.Courses) error
	GetCourses(limit int, page int) (courses []*pb.Course, err error)
}
type coursesHandler struct {
	uc CoursesUseCase
	pb.UnimplementedCourseServiceServer
}

func HandleError(err error) *pb.ErrorResponse {
	if errors, ok := err.(*common.AppError); ok {
		return &pb.ErrorResponse{
			Code:    int64(errors.StatusCode),
			Message: errors.Message,
		}
	}
	appErr := common.ErrInternal(err.(error))
	return &pb.ErrorResponse{
		Code:    int64(appErr.StatusCode),
		Message: appErr.Message,
	}
}

func NewCoursesHandler(uc CoursesUseCase) *coursesHandler {
	return &coursesHandler{uc: uc}
}
func (hdl *coursesHandler) GetCoursesWithPagination(ctx context.Context, rq *pb.GetCoursesPaginationRequest) (*pb.GetCoursesPaginationResponse, error) {

	res, err := hdl.uc.GetCourses(int(rq.PageSize), int(rq.Page))
	if err != nil {
		return &pb.GetCoursesPaginationResponse{
			Error: HandleError(err),
		}, nil
	}
	return &pb.GetCoursesPaginationResponse{
		Courses: res,
	}, nil
}
