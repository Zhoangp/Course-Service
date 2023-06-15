package usecase

import (
	"github.com/Zhoangp/Course-Service/internal/model"
	"github.com/Zhoangp/Course-Service/pb"
	"github.com/Zhoangp/Course-Service/pb/user"
	"strconv"
)

func (uc *coursesUseCase) GetCourseContent(userId, courseId string) (*pb.GetCourseContentResponse, error) {
	userIdDecoded, err := uc.h.Decode(userId)
	if err != nil {
		return nil, err
	}
	courseIdDecoded, err := uc.h.Decode(courseId)
	if err != nil {
		return nil, err
	}
	course, err := uc.repo.GetCourseContent(&model.Enrollment{
		UserId:   userIdDecoded,
		CourseId: courseIdDecoded,
	})
	if err != nil {
		return nil, err
	}
	course.FakeId = uc.h.Encode(course.Id)
	var sections []*pb.Section
	for _, i := range course.Sections {
		var lectures []*pb.Lecture
		for _, j := range i.Lectures {
			j.FakeId = uc.h.Encode(j.Id)
			lectures = append(lectures, &pb.Lecture{
				Id:      j.FakeId,
				Title:   j.Title,
				Content: j.Content,
				Status:  j.Status,
				Resource: &pb.Resource{
					Url:      j.Resource.Url,
					Duration: j.Resource.Duration,
				},
			})
		}
		i.FakeId = uc.h.Encode(i.Id)
		sections = append(sections, &pb.Section{
			Id:               i.FakeId,
			Title:            i.Title,
			NumberOfLectures: int32(i.NumberOfLectures),
			Lectures:         lectures,
		})
	}

	res := &pb.Course{
		Id:    course.FakeId,
		Title: course.Title,
		Instructor: &user.Instructor{
			Id: uc.h.Encode(course.InstructorID),
		},
		Sections:   sections,
		NumReviews: strconv.Itoa(course.NumReviews),
	}

	return &pb.GetCourseContentResponse{
		Course: res,
	}, nil
}
