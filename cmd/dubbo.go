package main

import (
	"context"
	"time"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"github.com/zhaoyunxing/client"
	"github.com/zhaoyunxing/protobuf/search"
	"github.com/zhaoyunxing/service"
)

func init() {
	config.SetProviderService(&service.SearchService{})
	config.SetProviderService(&service.ModelService{})
}

func main() {
	var (
		err error
		res *search.SearchReply
	)

	var modelService = &client.ModelClientImpl{}

	var searchService = &client.SearchClientImpl{}

	config.SetConsumerService(modelService)
	config.SetConsumerService(searchService)

	if err = config.Load(config.WithPath("./config/application.yaml")); err != nil {
		logger.Error(err)
	}

	time.Sleep(3 * time.Second)

	if res, err = modelService.Search(context.TODO(), &search.SearchRequest{Name: "dubbo-go"}); err != nil {
		logger.Error(err)
	}
	logger.Infof("name=%v", res.Message)

	if res, err = searchService.Search(context.TODO(), &search.SearchRequest{Name: "dubbo-go"}); err != nil {
		logger.Error(err)
	}
	logger.Infof("name=%v", res.Message)

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
