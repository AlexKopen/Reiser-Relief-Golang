package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

var useLocal bool

func main() {
	//cmd := exec.Command("sh", "-c", "lsof -t -i :8080 | xargs -r kill -9")
	//
	//_, err := cmd.Output()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	useLocal = os.Getenv("USELOCAL") == "true"

	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {

		data := map[string]interface{}{
			"Title":            "Home",
			"MissionStatement": callAPI("mission-statement"),
		}

		render(c, "home.tmpl", data)
	})

	router.GET("/about", func(c *gin.Context) {
		data := map[string]interface{}{
			"Title": "About",
		}

		render(c, "about.tmpl", data)
	})

	router.GET("/donate", func(c *gin.Context) {
		data := map[string]interface{}{
			"Title": "Donate",
		}

		render(c, "donate.tmpl", data)
	})

	router.GET("/contact", func(c *gin.Context) {
		data := map[string]interface{}{
			"Title": "Contact",
		}

		render(c, "contact.tmpl", data)
	})

	router.Run(":8080")
}

func render(c *gin.Context, templateName string, data map[string]interface{}) {
	commonData := map[string]interface{}{
		"Stylesheet": getCSSFile(),
	}

	for k, v := range data {
		commonData[k] = v
	}

	c.HTML(http.StatusOK, templateName, commonData)
}

func getCSSFile() string {
	files, err := ioutil.ReadDir("static")
	if err != nil {
		return ""
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".css") {
			return "/static/" + file.Name()
		}
	}

	return ""
}

func callAPI(key string) string {
	var hostURL string

	if useLocal {
		hostURL = "http://localhost:1337/api/"
	} else {
		hostURL = "https://goldfish-app-4dxk2.ondigitalocean.app/api/"
	}

	apiURL := fmt.Sprintf("%s%s", hostURL, key)
	bearerToken := "Bearer db7dcd0fe57639244ee844a6493294ff7859e7c5e7ba08cb5a87f6b8c9773cfe0b2aaea926240b801b633b1d2d159aa19834be0cadfb7e2ee86b1942c2b2a2b89fffe1bbc8b51920ec14cdc9f477f609c1cb10ece3a87719e24af2b0d91b38c3e3e1eac8bfaad89f9fc3ea54a2543272810844bf0730f6af2dbe10fad34a1cb2"

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
