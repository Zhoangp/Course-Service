package https

import (
	"context"
	"fmt"
	"github.com/Zhoangp/Course-Service/config"
	"github.com/Zhoangp/Course-Service/internal/model"
	"github.com/Zhoangp/Course-Service/pb"
	error1 "github.com/Zhoangp/Course-Service/pb/error"
	"github.com/Zhoangp/Course-Service/pb/user"
	client2 "github.com/Zhoangp/Course-Service/pkg/client"
	"github.com/Zhoangp/Course-Service/pkg/common"
)

type CoursesUseCase interface {
	CreateCourse(data *model.Course) error
	GetCourses(limit int, page int) (coursesResponse pb.GetCoursesResponse, err error)
	GetCourse(fakeId string) (*pb.GetCourseResponse, error)
	GetAllCategories() (*pb.GetAllCategoriesResponse, error)
	NewEnrollment(userId, courseId string) error
	GetCourseContent(userId, courseId string) (*pb.GetCourseContentResponse, error)
	GetEnrollments(userId string) (*pb.GetEnrollmentsResponse, error)
	NewCourse(rq *pb.CreateCourseRequest, instructorId string) (string, error)
	GetPrices() (*pb.GetPricesResponse, error)
	UpdateCourse(rq *pb.UpdateCourseRequest) error
	PublishCourse(rq *pb.PublishCourseRequest) ([]string, error)
	GetCoursesWithInstructor(rq *pb.GetCourseWithInstructorRequest) (*pb.GetCourseWithInstructorResponse, error)
	DeleteCourse(req *pb.DeleteCourseRequest) error
}
type coursesHandler struct {
	uc CoursesUseCase
	cf *config.Config
	pb.UnimplementedCourseServiceServer
}

func HandleError(err error) *error1.ErrorResponse {
	if errors, ok := err.(*common.AppError); ok {
		return &error1.ErrorResponse{
			Code:    int64(errors.StatusCode),
			Message: errors.Message,
		}
	}
	appErr := common.ErrInternal(err.(error))
	return &error1.ErrorResponse{
		Code:    int64(appErr.StatusCode),
		Message: appErr.Message,
	}
}

func NewCoursesHandler(uc CoursesUseCase, cf *config.Config) *coursesHandler {
	return &coursesHandler{uc: uc, cf: cf}
}
func (hdl *coursesHandler) GetCourses(ctx context.Context, rq *pb.GetCoursesRequest) (*pb.GetCoursesResponse, error) {
	res, err := hdl.uc.GetCourses(int(rq.PageSize), int(rq.Page))
	if err != nil {
		return &pb.GetCoursesResponse{
			Error: HandleError(err),
		}, nil
	}
	return &res, nil
}
func (hdl *coursesHandler) GetCourse(ctx context.Context, rq *pb.GetCourseRequest) (*pb.GetCourseResponse, error) {
	res, err := hdl.uc.GetCourse(rq.Id)
	if err != nil {
		return &pb.GetCourseResponse{
			Error: HandleError(err),
		}, nil
	}
	client, err := client2.InitUserServiceClient(hdl.cf)
	if err != nil {
		return &pb.GetCourseResponse{
			Error: HandleError(err),
		}, nil
	}
	instructor, err := client.GetInstructor(ctx, &user.GetInstructorInformationRequest{
		Id:  res.Course.Instructor.Id,
		Key: "instructor",
	})

	if err != nil {
		fmt.Println(err)
		return &pb.GetCourseResponse{
			Error: HandleError(err),
		}, nil
	}
	res.Course.Instructor = instructor.Information

	return res, nil
}
func (hdl *coursesHandler) GetAllCategories(ctx context.Context, rq *pb.GetAllCategoriesRequest) (*pb.GetAllCategoriesResponse, error) {
	res, err := hdl.uc.GetAllCategories()
	if err != nil {
		return &pb.GetAllCategoriesResponse{
			Error: HandleError(err),
		}, nil
	}
	return res, nil
}
