package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

/*
	Go语言中提供了一个完善的net/http包，通过这个http包可以很方便的就搭建起来一个可运行的Web服务
	go默认的http提供的路由规则不支持请求方法的限制
	下面演示了get post的请求
*/
func main() {
	//设置访问的路由
	http.HandleFunc("/say-hello", SayHello)
	//设置监听的端口
	err := http.ListenAndServe(":8080", http.DefaultServeMux)
	if nil != err {
		log.Fatal("http.ListenAndServe error: ", err)
	}
}

// 定义一个方法，用于处理业务逻辑，这里称为一个handle function，这个函数定义需要两个参数：
// 第一个参数为http.ResponseWriter类型，负责数据的响应
// 第二个参数为*http.Request类型，负责请求体的读取
// 这里需要注意：第一参数是interface，本身就是引用类型，第二个参数是struct，在go中是值传递，因此这里使用指针类型
func SayHello(w http.ResponseWriter, r *http.Request) {
	// 获取请求体
	body, _ := ioutil.ReadAll(r.Body)
	log.Println("request body:", string(body))
	req := struct {
		RequestId int `json:"request_id"`
	}{}
	//Json Unmarshal
	//将json字符串解码到相应的数据结构。
	// Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构.第二个参数必须是指针，否则无法接收解析的数据，
	_ = json.Unmarshal(body, &req)
	log.Println("req id is:", req)
	log.Println("req type is:", reflect.TypeOf(req))
	data := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{Name: "jack", Age: 18}

	//json.Marshal
	//将数据编码成json字符串。
	resp, _ := json.Marshal(data)
	w.Write(resp)
}

/*
curl --location --request POST 'localhost:8080/say-hello' \
--header 'Content-Type: application/json' \
--data-raw '{"request_id":11111111}'


curl --location --request GET 'localhost:8080/say-hello' \
--data-raw ''
*/
