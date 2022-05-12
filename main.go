package main

import (
	// Log items to the terminal

	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"root/configs"
	"root/controllers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

type ReverseProxy struct {
	// ModifyResponse is an optional function that
	// modifies the Response from the backend
	// If it returns an error, the proxy returns a StatusBadGateway error.
	ModifyResponse func(*http.Response) error
}

func rewriteBody(resp *http.Response) (err error) {
	b, err := ioutil.ReadAll(resp.Body) //Read html
	if err != nil {
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	b = bytes.Replace(b, []byte("server"), []byte("schmerver"), -1) // replace html
	body := ioutil.NopCloser(bytes.NewReader(b))
	resp.Body = body
	resp.ContentLength = int64(len(b))
	z := int(resp.ContentLength)
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))
	error_count := 0
	if z == 2 {
		error_count = error_count + 1

	}
	controllers.UpdateBytes(z, error_count)
	return err
}

func proxy(c *gin.Context) {
	var remote *url.URL
	var err error
	a := c.Param("proxyPath")
	i, err := strconv.Atoi(a[1:len(a)])
	v := (reflect.TypeOf(a).Kind())
	// for integer case
	if i != 0 {

		remote, err = url.Parse("https://httpstat.us/i")
		if err != nil {
			log.Println(err)
		}
	}
	if v == reflect.String && i == 0 {
		remote, err = url.Parse("https://jsonplaceholder.typicode.com/a")
		if err != nil {
			log.Println(err)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ModifyResponse = rewriteBody
	//Define the director func
	//This is a good place to log, for example
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)

}
func main() {
	log.Printf("HELLO WORLD")
	// Init gin router
	configs.Connect()
	router := gin.Default()
	// router.Use(GinBodyLogMiddleware())
	router.GET("/metrics/*proxyPath", proxy)
	router.POST("/metrics/*proxyPath", proxy)
	router.GET("/tenant/:tenantId", controllers.GetSingleTenant)
	router.GET("tenantactivity/:tenantactivityId", controllers.GetSingleTenantActivity)
	router.GET("/tenant", controllers.GetAllTenants)
	router.POST("/tenant", controllers.CreateTenant)
	// Its great to version your API's

	// Handle error response when a route is not defined
	router.NoRoute(func(c *gin.Context) {
		// In gin this is how you return a JSON response
		c.JSON(404, gin.H{"message": "Not found"})
	})

	// Init our server
	router.Run(":5000")
}
