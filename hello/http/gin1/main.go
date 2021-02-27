package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)
import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.LoadHTMLFiles("html/index.html")

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/api/report/:group/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		message := fmt.Sprintf("Same Path: %t", c.FullPath() == "/user/:name/*action") // true

		// https://stackoverflow.com/questions/47186741/how-to-get-the-json-from-the-body-of-a-request-on-go/47295689#47295689
		// TIL认为http.Response.Body是一个缓冲区，这意味着一旦读取它，就无法再次读取它。 就像水流一样，您可以看到它并在流过时对其进行测量，但是一旦流失，它就消失了。 但是，知道了这一点，有一种解决方法，您需要“捕获”身体并将其还原：
		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// Continue to use the Body, like Binding it to a struct:
		//order := new(models.GeaOrder)
		//error := context.Bind(order)

		c.String(http.StatusOK, message)
	})

	router.Run(":8002")
}
