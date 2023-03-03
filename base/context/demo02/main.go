package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

/**
对第三方调用传入 context
在 Go 语言中，Context 已经是默认支持的规范了。因此我们对第三方有调用诉求的时候，可以传入 context：

一般第三方开源库都已经实现了根据 context 的超时控制，所以当程序超时时，将会中断请求。
若你发现第三方开源库没有支持 context，建议换一个，免得出现级联故障。
*/

func main() {
	req, err := http.NewRequest("GET", "https://www.baidu.com/", nil)
	if err != nil {
		fmt.Printf("http.NewRequest err: %+v", err)
		return
	}

	//5000000
	ctx, cancel := context.WithTimeout(req.Context(), 50*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("http.DefaultClient.Do err: %+v", err)
		return
	}
	defer resp.Body.Close()
}
