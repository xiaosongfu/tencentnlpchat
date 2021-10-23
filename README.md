## 1. Tencent NLP Chat

腾讯闲聊 go client

## 2. 使用

2.1 添加依赖：

```
$ go get github.com/xiaosongfu/tencentnlpchat v0.1.1
```

2.2 调用 `Chat(...)` 方法：

```
ansText, err := tencentnlpchat.Chat(appId, appKey, question)
if err != nil {
    // TODO
}
```
