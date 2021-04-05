package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := newRouter()

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			// Template values
			"title":    "Demo Website",
			"version":  "1",
			"rows":     rows,
			"dateTime": time.Now().Format(time.Stamp),
		})
	})

	router.Run(fmt.Sprintf(":%d", *port))
}

var rows = [][]string{
	{
		"1,001",
		"Lorem",
		"ipsum",
		"dolor",
		"sit",
	},
	{
		"1,002",
		"amet",
		"consectetur",
		"adipiscing",
		"elit",
	},
	{
		"1,003",
		"Integer",
		"nec",
		"odio",
		"Praesent",
	},
	{
		"1,003",
		"libero",
		"Sed",
		"cursus",
		"ante",
	},
	{
		"1,004",
		"dapibus",
		"diam",
		"Sed",
		"nisi",
	},
	{
		"1,005",
		"Nulla",
		"quis",
		"sem",
		"at",
	},
	{
		"1,006",
		"nibh",
		"elementum",
		"imperdiet",
		"Duis",
	},
	{
		"1,007",
		"sagittis",
		"ipsum",
		"Praesent",
		"mauris",
	},
	{
		"1,008",
		"Fusce",
		"nec",
		"tellus",
		"sed",
	},
	{
		"1,009",
		"augue",
		"semper",
		"porta",
		"Mauris",
	},
	{
		"1,010",
		"massa",
		"Vestibulum",
		"lacinia",
		"arcu",
	},
	{
		"1,011",
		"eget",
		"nulla",
		"Class",
		"aptent",
	},
	{
		"1,012",
		"taciti",
		"sociosqu",
		"ad",
		"litora",
	},
	{
		"1,013",
		"torquent",
		"per",
		"conubia",
		"nostra",
	},
	{
		"1,014",
		"per",
		"inceptos",
		"himenaeos",
		"Curabitur",
	},
	{
		"1,015",
		"sodales",
		"ligula",
		"in",
		"libero",
	},
}

var port = flag.Int("port", 8080, "port to listen on")

func init() {
	flag.Parse()
}

func newRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/assets", "./site/assets")
	router.Static("/dist", "./site/dist")
	router.Static("/js", "./site/js")
	router.StaticFile("/dashboard.css", "./site/dashboard.css")
	router.LoadHTMLFiles("./site/index.html")
	return router
}
