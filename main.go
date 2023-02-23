package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	type completionRequest struct {
		Model       string  `json:"model"`
		Prompt      string  `json:"prompt"`
		MaxTokens   int     `json:"max_tokens"`
		Temperature float32 `json:"temperature"`
	}

	type completionChoice struct {
		Text         string `json:"text"`
		Index        int64  `json:"index"`
		LogProbs     string `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	}

	type completionUsage struct {
		PromptTokens     int `json:prompt_tokens`
		CompletionTokens int `json:completion_tokens`
		TotalTokens      int `json:total_tokens`
	}

	type completionResponse struct {
		ID      string             `json:"id"`
		Object  string             `json:"object"`
		Created int64              `json:"created"`
		Model   string             `json:"model"`
		Choices []completionChoice `json:"choices"`
		Usage   completionUsage    `json:"usage"`
	}

	godotenv.Load()

	authToken := os.Getenv("OPENAPI_API_KEY")
	url := "https://api.openai.com/v1/completions"

	data := completionRequest{
		Model:       "text-davinci-003",
		Prompt:      "How to become a highly competent software engineer",
		MaxTokens:   2048,
		Temperature: 0.5,
	}

	requestBody, err := json.Marshal(data)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error")
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error")
	}

	defer resp.Body.Close()

	var responseData completionResponse

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &responseData); err != nil {
		panic(err)
	}

	fmt.Println(responseData.Choices)

}
