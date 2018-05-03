package main

import (
	"net/http"
	"strings"
	"fmt"
)

func main()  {
	var r http.Request
	r.ParseForm()
	r.Form.Add("data","" )
	bodystr := strings.TrimSpace(r.Form.Encode())
	fmt.Println(bodystr)
}
