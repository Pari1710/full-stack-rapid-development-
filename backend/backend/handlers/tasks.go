package handlers

import (
	"net/http"
	"backend/db"
	"backend/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

const aiAPI = "https://api.openai.com/v1/completions"

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	aiResponse, err := getAITaskSuggestion(task.Title)
	if err == nil {
		task.Description = aiResponse
	}
	db.DB.Create(&task)
	c.JSON(http.StatusOK, task)
}

func getAITaskSuggestion(taskTitle string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	requestBody, _ := json.Marshal(map[string]string{"prompt": "Suggest steps for " + taskTitle})

	req, _ := http.NewRequest("POST", aiAPI, ioutil.NopCloser(bytes.NewReader(requestBody)))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res["choices"].([]interface{})[0].(map[string]interface{})["text"].(string), nil
}
