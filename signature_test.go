package tencentnlpchat

import (
	"testing"
)

func TestSignature(t *testing.T) {
	appId := "2114827333"
	appKey := "HpbOP2UPwQazhAZ7"
	session := "testclient"
	nonceStr := "testclient"
	timeStamp := "1554878560"
	question := "你叫啥"
	want := "B0C6DE4B0C9D4BEED530E6542302654C"

	got := signature(appId, appKey, question, session, nonceStr, timeStamp)

	if got != want {
		t.Fatalf("the result is error, want %s, got %s", want, got)
	}
}

func BenchmarkSignature(b *testing.B) {
	timeStamp := "1554878560"
	question := "你叫啥"

	for i := 0; i < b.N; i++ {
		_ = signature(appId, appKey, question, session, nonceStr, timeStamp)
	}
}
