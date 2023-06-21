package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCovidSummary(t *testing.T) {
	router := gin.Default()
	router.GET("/covid/summary", getCovidSummary)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/covid/summary", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response gin.H
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Contains(t, response, "Province")
	assert.Contains(t, response, "AgeGroup")

	provinceCounts, ok := response["Province"].(map[string]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, provinceCounts)

	ageGroupCounts, ok := response["AgeGroup"].(map[string]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, ageGroupCounts)
}
