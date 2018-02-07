package main

import "net/url"

func main()  {
	print(url.QueryEscape("中文"))
}
