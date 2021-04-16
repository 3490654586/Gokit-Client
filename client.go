package main

import (
	"consul-client/Services"
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	clientHttp "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/log"
	"io"
	"net/url"
	"os"
)

func main()  {
	//第一步创建consul链接
	config := consulapi.DefaultConfig()
	config.Address ="localhost:8500"
	api_client ,_ :=	consulapi.NewClient(config)
    client :=  consul.NewClient(api_client)


	logger := log.NewLogfmtLogger(os.Stdout)
	var Tag = []string{"primary"}

	//第二步穿件Consul实例
	instancer := consul.NewInstancer(client,logger,"userservice",Tag,true)
     factory := func(services_url string)(endpoint.Endpoint,io.Closer,error) { //factory定义了如何获得服务端的endpoint
     	tart,_  := url.Parse("http://"+services_url)
     	return clientHttp.NewClient("GET",tart,Services.GetUser,Services.GetUserResponse).Endpoint(),nil,nil
	 }

	endpointer := sd.NewEndpointer(instancer,factory,logger)
	endpoints,_ := endpointer.Endpoints()
	fmt.Println("服务有",len(endpoints))
	getUserInfo := endpoints[0]
    //第三步创建一个上下文对象
	ctx := context.Background()
    //第四部执行
 res ,_ :=	getUserInfo(ctx,Services.UserRequest{Uid:101})
  //第五部获得相应值
 response := res.(Services.UserResponse)
 fmt.Println(response.Resule)
}
