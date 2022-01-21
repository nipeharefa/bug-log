package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nipeharefa/bug-log/bug"
)

func lapor(ctx context.Context) error {
	a, ok := ctx.Value("bugsnag").(bug.ReportFunc)
	if !ok {
		return errors.New("bugsnag not found")
	}
	a(errors.New("S"))
	return nil
}

func main() {

	bug := bug.NewBugsnag()

	go bug.Watch()
	r := gin.Default()

	// http.Handler
	// hh := gin.WrapF(fun)
	r.Use(bug.GinHandler())
	// gin.WrapH()
	r.GET("/ping", func(c *gin.Context) {
		// a, b := c.
		a := c.Request.Context().Value("bugsnag")
		fmt.Println(a, c.Request.Context())
		lapor(c.Request.Context())
		c.JSON(http.StatusOK, nil)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
