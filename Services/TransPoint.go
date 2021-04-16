package Services

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetUser( ctx  context.Context, request *http.Request, r interface{})error{
        req :=  r.(UserRequest)
        request.URL.Path += "/user/"+strconv.Itoa(req.Uid)
	return nil
}

func GetUserResponse( ctx context.Context, res *http.Response) (response interface{}, err error){

	var UserResonsess UserResponse

	err = json.NewDecoder(res.Body).Decode(&UserResonsess)
	if err != nil {
		return nil,err
	}

	return UserResonsess,nil
}