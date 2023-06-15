package https

import (
	"context"
	"github.com/Zhoangp/Course-Service/pb"
)

func (hdl *coursesHandler) Enrollment(ctx context.Context, request *pb.EnrollmentRequest) (*pb.EnrollmentResponse, error) {
	if err := hdl.uc.NewEnrollment(request.UserId, request.CourseId); err != nil {
		return &pb.EnrollmentResponse{Error: HandleError(err)}, nil
	}
	return &pb.EnrollmentResponse{}, nil
}
func (hdl *coursesHandler) GetEnrollments(ctx context.Context, rq *pb.GetEnrollmentsRequest) (*pb.GetEnrollmentsResponse, error) {
	res, err := hdl.uc.GetEnrollments(rq.UserId)
	if err != nil {
		return &pb.GetEnrollmentsResponse{
			Error: HandleError(err),
		}, nil
	}
	return res, nil
}
