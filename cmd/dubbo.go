package main

import (
	"context"
	"time"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	hessian "github.com/apache/dubbo-go-hessian2"

	"github.com/zhaoyunxing/domain"
	"github.com/zhaoyunxing/reference"
	"github.com/zhaoyunxing/service"
)

func init() {
	config.SetProviderService(&service.UserService{})
	config.SetProviderService(&service.OrderService{})
	hessian.RegisterPOJO(&domain.User{})
}

func main() {
	var (
		err error
		res *domain.User
	)

	var rpcService = &reference.SearchService{}

	var modelService = &reference.ModelService{}
	config.SetConsumerService(rpcService)
	config.SetConsumerService(modelService)

	if err = config.Load(config.WithPath("./config/application.yaml")); err != nil {
		logger.Error(err)
	}

	time.Sleep(3 * time.Second)

	if res, err = rpcService.GetUser(context.TODO(), &domain.User{Name: "zhaoyunxing"}); err != nil {
		logger.Error(err)
	}

	if res, err = modelService.GetUser(context.TODO(), &domain.User{Name: "zhaoyunxing"}); err != nil {
		logger.Error(err)
	}
	logger.Infof("name=%v", res.Name)

	//if err = config.Load(config.WithPath("./config/application.yaml")); err != nil {
	//	fmt.Printf("dubbo go config load err: %v \n", err)
	//}
	//
	//services[0].SetService(rpcService)
	//services[1].SetService(rpcService)
	//
	//if res, err = services[0].Get(context.TODO(), &domain.User{Name: "zhaoyunxing"}); err != nil {
	//	logger.Error(err)
	//}
	//
	//logger.Infof("name=%v", res.Name)
	//
	//if res, err = services[1].Get(context.TODO(), &domain.User{Name: "Alex001"}); err != nil {
	//	logger.Error(err)
	//}
}
