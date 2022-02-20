package main

import (
	"github.com/aryadiahmad4689/test-citcall/transport"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("views/*.html")
	transport.NewServer(e).Start()
}
