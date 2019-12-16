package tencentnlpchat

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	// nlp 接口地址
	nlpApi = "https://api.ai.qq.com/fcgi-bin/nlp/nlp_textchat"
	//// app id
	//appId = "2114827333"
	//// app key
	//appKey = "HpbOP2UPwQazhAZ7"
)

const (
	// 会话标识(应用内唯一)，UTF-8编码，非空且长度上限32字节
	session = "nlp-client"
	// 随机字符串，非空且长度上限32字节
	nonceStr = "nlp-client"
)

// 接口成功的返回码
const retOK = 0

// 接口返回值
type Result struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

// 接口返回值得 data 部分
type Data struct {
	Session string `json:"session"`
	Answer  string `json:"answer"`
}

// Chat 实现问答接口
func Chat(appId, appKey, question string) (answer string, err error) {
	timeStamp := strconv.Itoa(int(time.Now().Unix()))

	// 获取签名字符串
	sign := signature(appId, appKey, question, session, nonceStr, timeStamp)

	// 准备参数
	params := url.Values{}
	params.Set("app_id", appId)
	params.Set("session", session)
	params.Set("time_stamp", timeStamp)
	params.Set("nonce_str", nonceStr)
	params.Set("sign", sign)
	params.Set("question", question)

	// 发起 http 请求
	resp, err := http.PostForm(nlpApi, params)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// 解析结果
	var result Result
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	// 处理并判定结果
	if result.Ret == retOK {
		answer = result.Data.Answer
	} else {
		err = errors.New(result.Msg)
	}

	return
}
