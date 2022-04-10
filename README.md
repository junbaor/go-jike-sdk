# go-jike-sdk

Go 版本的 Jike SDK

## Usage
```go
// create client
client := jike.NewJike(areaCode, phone)

// passport login
loginOutput, err := client.UserService.PasswordLogin(content, areaCode, phone, password)

// follower timeline
timeline, err := client.UserService.FollowingTimeline(content, 10, jike.TimelineLoadMoreKey{})
```

## Alternatives
- [jike-cli](https://github.com/junbaor/jike-cli) - 即刻命令行客户端 (Java)
- [jike-sdk](https://github.com/open-jike/jike-sdk) - jike-sdk Ⓙ (Node.js / 浏览器 / Deno)
