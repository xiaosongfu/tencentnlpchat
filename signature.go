package tencentnlpchat

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"strings"
)

// signature generate signature string.
func signature(appId, appKey, question, session, nonceStr, timeStamp string) string {
	// 请求参数
	params := url.Values{}
	params.Set("app_id", appId)
	params.Set("question", question)
	params.Set("session", session)
	params.Set("time_stamp", timeStamp)
	params.Set("nonce_str", nonceStr)

	// 拼接完整的参数
	// params.Encode() 会自动按照 key 排序
	plainStr := fmt.Sprintf("%s&app_key=%s", params.Encode(), appKey)
	// md5 加密为字节切片
	bytes := md5.Sum([]byte(plainStr))
	// 转换成大写的16进制
	md5Str := fmt.Sprintf("%X", bytes) // 需要大写的16进制
	return strings.ToUpper(md5Str)
}
