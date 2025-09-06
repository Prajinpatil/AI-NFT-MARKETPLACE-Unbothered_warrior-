package main

import "gofr.dev/pkg/gofr"

func main() {
    app := gofr.New()

    app.GET("/ping", func(c *gofr.Context) (interface{}, error) {
        return "pong", nil
    })
	app.GET("/hello", func(ctx *gofr.Context) (interface{}, error) {
    return map[string]string{"message": "Hello from GoFr backend!"}, nil
})

    app.Run()
}   