package test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	req, err := http.NewRequest("get", "http://www.baidu.com", nil)
	fmt.Println(err.Error())
	resp, err := http.DefaultClient.Do(req)
	fmt.Println(resp)
}
