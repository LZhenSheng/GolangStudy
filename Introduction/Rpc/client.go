package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	//http协议
	req := HttpRequest.NewRequest()
	req.SetTimeout(10 * time.Second)
	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b))
	body, _ := res.Body()
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}
func main() {
	fmt.Println(Add(2, 2))
}
