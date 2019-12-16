# Tencent NLP Chat

腾讯闲聊 go client

# 快速开始

1. 添加依赖：

```
require (
	github.com/xiaosongfu/tencentnlpchat v0.1.1
)
```

2. 调用 Chat(...) 方法：

```
ansText, err := tencentnlpchat.Chat(appId, appKey, question)
if err != nil {
    // TODO
}
```
