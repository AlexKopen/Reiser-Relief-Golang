package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ResponseData struct {
	Data ResponseDetails `json:"data"`
	Meta json.RawMessage `json:"meta"`
}

type ResponseDetails struct {
	ID         int        `json:"id"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	Text        string    `json:"text"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	PublishedAt time.Time `json:"publishedAt"`
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Define a route for the main page
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"Title":            "Home",
			"MissionStatement": callAPI("mission-statement"),
		})
	})

	// Define a route for the about page
	router.GET("/about", func(c *gin.Context) {
		c.HTML(200, "about.tmpl", gin.H{
			"Title": "About",
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}

func callAPI(key string) string {
	apiURL := fmt.Sprintf("http://127.0.0.1:1337/api/%s", key)
	bearerToken := "Bearer c7ca6965ed3d9542dec6de192c30986c5f5ace7dbc8f2b3f0ace3344cdee1e6f81417a18385f146e4682d771a1d02505f9ec93bbb51529883f7a4908161248e02028d9c763e6c3c6061931a95b3a3f30a006d275ca7964d25ec13c51bec3cc367db6c24307aabea063361462a0fce406b701cced72b9cfed3433f251c9e2bf4f"

	req, err := http.NewRequest("GET", apiURL, nil)

	if err != nil {
		log.Fatal("http.NewRequest error:", err)
		return err.Error()
	}

	req.Header.Add("Authorization", bearerToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("client.Do error:", err)
		return err.Error()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll error:", err)
		return err.Error()
	}

	var data ResponseData

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("json.Unmarshal error:", err)
		return err.Error()
	}

	return data.Data.Attributes.Text
}
