package wecom_robot

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type wecom_Robot struct {
	url   string
	key   string
	ctype string
}

func New_WeComBot() *wecom_Robot {
	return &wecom_Robot{
		url:   "https://qyapi.weixin.qq.com/cgi-bin/webhook/send",
		key:   "",
		ctype: "application/json",
	}
}

func (w *wecom_Robot) SetUrl(u string) {
	w.url = u
}

func (w *wecom_Robot) SetKey(s string) {
	w.key = s
}

func (w *wecom_Robot) Send(b *bytes.Buffer) (resp *http.Response, err error) {
	return http.Post(w.url, w.ctype, b)
}

// Default Text message
func NewMessage(params ...string) *Message {

	switch params[0] {
	case "markdown":
		return &Message{
			Msgtype:  "markdown",
			Markdown: MarkDown{},
		}
	case "text":
		content := ""
		if len(params) > 1 {
			content = params[1]
		}
		return &Message{
			Msgtype: "text",
			Text: TEXT{
				Content: content,
			},
		}
	case "image":
		return &Message{
			Msgtype: "image",
			Img:     IMAGE{},
		}
	default:
		return &Message{
			Msgtype: "text",
			Text:    TEXT{},
		}
	}
}

func NewText(cont ...string) {

}

type Message struct {
	Msgtype string `json:"msgtype,omitempty"`

	Markdown MarkDown `json:"markdown,omitempty"`
	Text     TEXT     `json:"text,omitempty"`
	Img      IMAGE    `json:"image,omitempty"`
}

type MarkDown struct {
	Content string `json:"content,omitempty"`
}

type TEXT struct {
	Content               string   `json:"content,omitempty"`
	Mentioned_list        []string `json:"mentioned_list,omitempty"`
	Mentioned_mobile_list []string `json:"mentioned_mobile_list,omitempty"`
}

type IMAGE struct {
	Data_b64 string `json:"base64,omitempty"`
	MD5hash  string `json:"md5,omitempty"`
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func (img *IMAGE) EncodeB64FromFile(p string) {
	bytes, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatal(err)
	}
	img.MD5hash = fmt.Sprintf("%x", md5.Sum(bytes))
	img.Data_b64 = toBase64(bytes)
}

func (img *IMAGE) EncodeB64FromUrl(p string) {
	resp, err := http.Get(p)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	img.MD5hash = fmt.Sprintf("%x", md5.Sum(bytes))
	img.Data_b64 = toBase64(bytes)
}
