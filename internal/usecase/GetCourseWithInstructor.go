package usecase

import "github.com/Zhoangp/Course-Service/pb"

func (uc *coursesUseCase) GetCoursesWithInstructor(rq *pb.GetCourseWithInstructorRequest) (*pb.GetCourseWithInstructorResponse, error) {
	userIdDecoded, err := uc.h.Decode(rq.UserId)
	if err != nil {
		return nil, err
	}
	courses, err := uc.repo.FindDataWithCondition(map[string]any{"instructor_id": userIdDecoded})
	if err != nil {
		return nil, err
	}
	var res pb.GetCourseWithInstructorResponse
	for _, course := range courses {
		var img pb.Image
		if course.Thumbnail != nil {
			img = pb.Image{
				Url:    course.Thumbnail.Url,
				Width:  course.Thumbnail.Width,
				Height: course.Thumbnail.Height,
			}
		}
		res.Courses = append(res.Courses, &pb.Course{
			Id:        uc.h.Encode(course.Id),
			Title:     course.Title,
			Thumbnail: &img,
		})
	}
	return &res, nil
}
