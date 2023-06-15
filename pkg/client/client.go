package client

import (
	"fmt"
	"github.com/Zhoangp/Course-Service/config"
	"github.com/Zhoangp/Course-Service/pb/user"
	"google.golang.org/grpc"
)

// import (
//
//	"fmt"
//	"github.com/Zhoangp/User-Service/config"
//	"github.com/Zhoangp/User-Service/pb"
//	"google.golang.org/grpc"
//
// )
func InitUserServiceClient(c *config.Config) (user.UserServiceClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.OtherServices.UserUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
		return nil, err
	}
	return user.NewUserServiceClient(cc), nil
}
