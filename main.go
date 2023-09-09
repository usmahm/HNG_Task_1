package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type submission_details struct {
	SlackName     string `json:"slack_name"`
	CurrentDay    string `json:"current_day"`
	UTCTime       string `json:"utc_time"`
	Track         string `json:"track"`
	GithubFileUrl string `json:"github_file_uri"`
	GithubRepoUrl string `json:"github_repo_url"`
	StatusCode    int    `json:"status_code"`
}

func getSubmissionDetails(c *gin.Context) {
	slack_name := c.Query("slack_name")
	track := c.Query("track")

	fmt.Println(time.Now().UTC())
	fmt.Println(time.Now().Weekday())

	curr_utc_time := time.Now().UTC()

	s_d := submission_details{
		SlackName:     slack_name,
		Track:         track,
		UTCTime:       curr_utc_time.Format("2006-01-02T15:04:05Z"),
		CurrentDay:    curr_utc_time.Weekday().String(),
		GithubFileUrl: "https://github.com/usmahm/HNG_Task_1/blob/master/main.go",
		GithubRepoUrl: "https://github.com/usmahm/HNG_Task_1",
		StatusCode:    200,
	}

	c.JSON(http.StatusOK, s_d)
}

func main() {
	router := gin.Default()

	router.GET("/submission-details", getSubmissionDetails)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
