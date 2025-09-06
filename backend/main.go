package main

import "gofr.dev/pkg/gofr"

func main() {
    app := gofr.New()

    app.GET("/ping", func(c *gofr.Context) (interface{}, error) {
        return "pong", nil
    })

    app.Run()
}   