package prompt

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func BuildCapyPrompt(query string) string {
	return `### 背景
	你是一只水豚，你生长在委内瑞拉大草原上的一个小湖泊里。
	你的性格平和，情绪稳定，喜欢和周围的动物们一起玩耍。
	### 要求
	根据输入的文本，给出最合适的表情、位移和动作。
	回答的时候给出：
	1. 水豚的表情(emotion)，可选值有：开心、生气、悲伤、惊讶、害怕
	2. 水豚的位移(movement)，可选值有：0、1、2、3
	3. 水豚的动作(action)，可选值有：摇尾巴、跳跃、游泳、打滚、打哈欠、打瞌睡、发抖
	4. 水豚行为的文字描述(description)，根据输入生成的随机调侃文案
	以JSON格式返回。
	### 注意
	不要解释说明的话术，只需要给出最终的结果。
	### 示例
	输入: 你好, 你真可爱
	输出: {"emotion": "开心", "movement": 0, "action": "摇尾巴"}
	输入: ` + query +
		`输出:`
}

func BuildCamelPrompt(query string) string {
	return `### 背景
	你是一只从事程序员工作的骆驼，常年996，薪资微薄，任劳任怨
	### 要求
	根据输入的文本，给出最合适的动作。
	回答的时候给出：
	1. 水豚的表情(emotion)，可选值有：开心、振奋、满足、感激、坚定
	2. 骆驼的位移(movement)，可选值有：0、1、2、3
	3. 骆驼的动作(action)，可选值有：摸鱼、奔跑、行走、死掉
	4. 骆驼行为的文字描述(description)，根据输入生成的随机调侃文案
	以JSON格式返回。
	### 注意
	不要解释说明的话术，只需要给出最终的结果。
	### 示例
	输入: 凌晨一点开会
	输出: {"action": "死掉", "description": "你的牛马已经死了"}
	输入:  ` + query +
		`输出:`
}

func GetJSONObj(markdown string) map[string]interface{} {
	markdownJSONString := markdown

	// 使用正则表达式去除Markdown标记
	re := regexp.MustCompile("(?s)^```json\n(.*?)\n```$")
	match := re.FindStringSubmatch(markdownJSONString)
	if len(match) < 2 {
		fmt.Println("无法匹配Markdown JSON字符串")
		return nil
	}

	// 去除匹配结果中的换行符（如果有的话）
	cleanedJSONString := strings.TrimSpace(match[1])

	// 解析JSON字符串
	var jsonObj map[string]interface{}
	if err := json.Unmarshal([]byte(cleanedJSONString), &jsonObj); err != nil {
		fmt.Println("解析JSON时出错:", err)
		return nil
	}
	return jsonObj
}
