package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type COVIDData struct {
	Data []struct {
		Province string `json:"Province"`
		Age      int    `json:"Age"`
	} `json:"Data"`
}

func getCovidSummary(c *gin.Context) {
	resp, err := http.Get("https://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch data",
		})
		return
	}

	defer resp.Body.Close()

	var data COVIDData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to decode data",
		})
		return
	}

	provinceCounts := make(map[string]int)
	ageGroupCounts := make(map[string]int)

	for _, caseData := range data.Data {
		// count cases by province
		provinceCounts[caseData.Province]++

		// count cases by age group
		ageGroup := getAgeGroup(caseData.Age)
		ageGroupCounts[ageGroup]++
	}

	c.JSON(http.StatusOK, gin.H{
		"Province": provinceCounts,
		"AgeGroup": ageGroupCounts,
	})
}

func getAgeGroup(age int) string {
	if age >= 0 && age <= 30 {
		return "0-30"
	} else if age >= 31 && age <= 60 {
		return "31-60"
	} else if age > 60 {
		return "61+"
	}
	return "N/A"
}

func main() {
	r := gin.Default()
	r.GET("/covid/summary", getCovidSummary)
	r.Run(":2566")
}
