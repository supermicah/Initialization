package main

import (
	"fmt"

	v20190614 "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/asr/v20190614"
)

func main() {
	fmt.Println(1)
	test()
}

func test() {
	cli, err := v20190614.NewClient(nil, "", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cli)
}
