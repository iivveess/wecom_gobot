# wecom_gobot
simple library for WeCom

# Implementation
text, image, and markdown messages

# Sample code

```go
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	wecom_gobot "github.com/iivveess/wecom_gobot"
)

func main() {
	bot := wecom_gobot.New_WeComBot()
	bot.SetUrl("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=${your_key}")

	DATAs := []wecom_gobot.Message{}
	data := wecom_gobot.NewMessage("image")
	data.Img.EncodeB64FromFile("./test.png")
	// data.Img.EncodeB64FromUrl("https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRJPJqciuRfp0fx3vH87SxbcSkdsSsUCbX2C2Bxu02eOpEe2Uj93ZU92btSL71YvPi8T3A&usqp=CAU")
	DATAs = append(DATAs, *data)

	data2 := wecom_gobot.NewMessage("text", "Hello from bot")
	DATAs = append(DATAs, *data2)

	data3 := wecom_gobot.NewMessage("markdown")
	unescaped := "实时新增用户反馈<font color=\"warning\">132例</font>，请相关同事注意。\n>类型:<font color=\"comment\">用户反馈</font>\n>普通用户反馈:<font color=\"comment\">117例</font>\n>VIP用户反馈:<font color=\"comment\">15例</font>"
	data3.Markdown.Content = unescaped

	DATAs = append(DATAs, *data3)

	for _, data := range DATAs {
		postBody, _ := json.Marshal(data)
		requestBody := bytes.NewBuffer(postBody)
		resp, err := bot.Send(requestBody)
		//Handle Error
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
		//Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)
	}
}

```
