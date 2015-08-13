package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
)

func ExampleNewHandler() {
	s := httptest.NewServer(New(func(content string) {
		fmt.Println(content)
	}))
	defer s.Close()
	c := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return errors.New("fake error")
	}}
	url := s.URL + "/log?url=http%3A%2F%2Fredirect_url&content=test_content"
	resp, err := c.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
	fmt.Println(resp.Header["Location"])
	// Output:
	// test_content
	// Get http://redirect_url: fake error
	// 303 See Other
	// [http://redirect_url]
}
