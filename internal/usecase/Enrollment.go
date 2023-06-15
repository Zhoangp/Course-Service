package usecase

import (
	"github.com/Zhoangp/Course-Service/internal/model"
	"github.com/Zhoangp/Course-Service/pb"
	"github.com/Zhoangp/Course-Service/pkg/common"
)

func (uc *coursesUseCase) NewEnrollment(userId, courseId string) error {
	userIdDecode, err := uc.h.Decode(userId)
	if err != nil {
		return err
	}
	courseIdDecode, err := uc.h.Decode(courseId)
	if err != nil {
		return err
	}
	if err = uc.repo.NewEnrollment(&model.Enrollment{
		UserId:   userIdDecode,
		CourseId: courseIdDecode,
	}); err != nil {

		return err
	}
	return nil
}
func (uc *coursesUseCase) GetEnrollments(userId string) (*pb.GetEnrollmentsResponse, error) {
	userIdDecode, err := uc.h.Decode(userId)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		SQLModel: common.SQLModel{
			Id: userIdDecode,
		},
	}
	var res pb.GetEnrollmentsResponse
	courses, err := uc.repo.GetEnrollments(user)
	if err != nil {
		return nil, err
	}
	for _, i := range courses {
		res.Enrollments = append(res.Enrollments, &pb.Enrollment{
			CourseId: uc.h.Encode(i.Id),
			Title:    i.Title,
			Thumbnail: &pb.Image{
				Url:    i.Thumbnail.Url,
				Width:  i.Thumbnail.Width,
				Height: i.Thumbnail.Height,
			},
		})
	}
	return &res, nil
}
