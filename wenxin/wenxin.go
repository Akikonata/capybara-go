package wenxin

import (
	"bytes"
	"capybara-go/config"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

var access_token string = ""
var expires_in int64 = 0

func isAccessTokenExpired() bool {
	timestamp := time.Now().Unix()
	return timestamp >= expires_in
}

type AccessTokenResponse struct {
	RefreshToken  string `json:"refresh_token"`
	AccessToken   string `json:"access_token"`
	ExpiresIn     int64  `json:"expires_in"`
	SessionKey    string `json:"session_key"`
	Scope         string `json:"scope"`
	SessionSecret string `json:"session"`
}

func getAccessToken() string {
	if !isAccessTokenExpired() {
		return access_token
	}
	requestUrl := config.GlobalConfig.Qianfan.TokenUrl + "?grant_type=client_credentials&client_id=" + config.GlobalConfig.Qianfan.ApiKey + "&client_secret=" + config.GlobalConfig.Qianfan.SecretKey
	resp, err := http.Post(requestUrl, "application/json", nil)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	var accessTokenResponse AccessTokenResponse
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		return err.Error()
	}
	access_token = accessTokenResponse.AccessToken
	timestamp := time.Now().Unix()
	expires_in = timestamp + accessTokenResponse.ExpiresIn - 3000
	return access_token
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages []Message `json:"messages"`
}

type ChatResponse struct {
	Result string `json:"result"`
}

func Chat(query string) string {
	accessToken := getAccessToken()
	requestUrl := config.GlobalConfig.Qianfan.ChatbotUrl + "?access_token=" + accessToken
	chatRequest := ChatRequest{
		Messages: []Message{{
			Role:    "user",
			Content: query,
		}},
	}
	requestBody, err := json.Marshal(chatRequest)
	println("请求体：", string(requestBody))
	if err != nil {
		return err.Error()
	}
	resp, err := http.Post(requestUrl, "application/json", bytes.NewReader(requestBody))

	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	println(string(body))
	if err != nil {
		return err.Error()
	}
	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		return err.Error()
	}
	return chatResponse.Result
}
func init() {
	getAccessToken()
}
