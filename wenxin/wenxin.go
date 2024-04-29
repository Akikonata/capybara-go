package wenxin

import (
	"capybara-go/config"
	"time"
)

var String access_token = ""
var int expires_in = 0
apiKey := config.GlobalConfig.Qianfan.ApiKey
secretKey := config.GlobalConfig.Qianfan.SecretKey
tokenUrl := config.GlobalConfig.Qianfan.TokenUrl

func isAccessTokenExpired() bool {
	timestamp := time.Now().Unix()
	return timestamp >= expires_in
}


func getAccessToken() {
	if !isAccessTokenExpired() {
		return access_token
	}
	// 1. 构造请求参数
	params := url.Values{}
	params.Set("grant_type", "client_credentials")
	params.Set("client_id", apiKey)
	params.Set("client_secret", secretKey)

	// 2. 发送请求
	resp, err := http.PostForm("https://aip.baidubce.com/oauth/2.0/token", params)
	if err != nil {
		// 处理请求错误
		return
	}
	defer resp.Body.Close()

	// 3. 解析返回值
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 处理读取返回值错误
		return
	}

	// 4. 解析access_token和expires_in
	type AccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	var accessTokenResponse AccessTokenResponse
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		// 处理解析返回值错误
		return
	}

	// 5. 更新access_token和expires_in
	access_token = accessTokenResponse.AccessToken
	timestamp := time.Now().Unix()
	expires_in = timestamp + accessTokenResponse.ExpiresIn - 3000
	return access_token
}

func chat(String query) {

}

func init() {
	GetAccessToken()
}
