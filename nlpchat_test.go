package tencentnlpchat

import "testing"

const (
	// app id
	appId = "2114827333"
	// app key
	appKey = "HpbOP2UPwQazhAZ7"
)

func TestChat(t *testing.T) {
	question := "屁"
	// 不能测试的太频繁，不然答案会不一样
	want := "屁也叫矢气，多数动物人类从肛门排放的废气。由于屁的成分与正常空气的成分比例严重不同，而有嗅觉的动物、人类的鼻子对这些异常气体能够感知。屁的产生，是食物与唾液、胃液、胰液胆汁混合后在肠道，被各种产气厌氧菌、产气好氧菌分解后生产的气体。因为屁对空气的污染也被动物作为攻击气体，而人类的屁往往是身体健康状态的提示，频繁放屁也会影响社交、个人运程，必须认真调理。"

	got, err := Chat(appId, appKey, question)

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("the result is error, want %s, got %s", want, got)
	}
}

func BenchmarkChat(b *testing.B) {
	question := "你叫啥"

	for i := 0; i < b.N; i++ {
		_, _ = Chat(appId, appKey, question)
	}
}
