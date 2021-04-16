package main

import (
	"consul-client/Services"
	"context"
	"fmt"
	clientHttp "github.com/go-kit/kit/transport/http"
	"net/url"
)
func main()  {
	tat,_:= url.Parse("http://127.0.0.1:8050")
	client := clientHttp.NewClient("GET",tat,Services.GetUser,Services.GetUserResponse)
   getUserInfo := client.Endpoint()

   ctx := context.Background()

  res ,err := getUserInfo(ctx,Services.UserRequest{Uid:100})
	if err != nil {
		fmt.Println(err)
	}

	str := res.(Services.UserResponse)
	fmt.Println(str.Resule)

	fmt.Println(12)
}