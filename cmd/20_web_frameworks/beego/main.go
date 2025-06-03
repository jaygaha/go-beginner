package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/jaygaha/go-beginner/cmd/20_web_frameworks/beego/routers" // Import routers with blank identifier
)

func main() {
	beego.Run()
}
