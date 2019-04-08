package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err:= http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	clien := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}
	//resp, err := http.Get("http://www.imooc.com")
	//resp, err := http.DefaultClient.Do(request)
	resp, err := clien.Do(request)
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(s))
}
