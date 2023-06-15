package https

import (
	"context"
	"github.com/Zhoangp/Course-Service/pb"
)

func (hdl *coursesHandler) GetPrices(ctx context.Context, request *pb.GetPricesRequest) (*pb.GetPricesResponse, error) {
	res, err := hdl.uc.GetPrices()
	if err != nil {
		return &pb.GetPricesResponse{
			Error: HandleError(err),
		}, nil
	}
	return res, nil
}
