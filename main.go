package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()

	r := gin.Default()
	r = collectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
