package usecase

import (
	"github.com/Zhoangp/Course-Service/pb"
	"strconv"
)

func (uc *coursesUseCase) GetPrices() (*pb.GetPricesResponse, error) {
	res, err := uc.repo.GetPrices()
	if err != nil {
		return nil, err
	}
	var list pb.GetPricesResponse
	for _, item := range res {
		list.Prices = append(list.Prices, &pb.Price{
			Id:       strconv.Itoa(item.Id),
			Value:    item.Value,
			Currency: item.Currency,
		})
	}
	return &list, nil

}
