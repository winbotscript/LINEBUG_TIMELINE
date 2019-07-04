package main

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"os"
	"log"
	"fmt"
)

// Config
const (
	HOST         = "https://legy-jp.line.naver.jp"
	CREATE_URL   = HOST + "/mh/api/v39/post/create.json"//timelinebug＆timelineStr
	FEED_URL     = HOST + "/mh/api/v39/feed/list.json"//timelinefeed
	SEND_URL     = HOST + "/mh/api/v39/post/sendPostToTalk.json"//timelinesend
	SHARE_URL    = HOST + "/mh/api/v39/post/getShareLink.json"
	GROUP_URL    = HOST + "/mh/api/v39/post/list.json"//timelinegroup
	HASHTAG_URL  = HOST + "/mh/api/v30/hashtag/search.json"
	UPDATE_URL   = HOST + "/mh/api/v39/post/update.json"
	LIKE_URL     = HOST + "/mh/api/v39/like/create.json"//timelinelike
	COMEMENT_URL = HOST + "/mh/api/v39/comment/create.json"//timelinecomement
)

// Parameter
const (
	timelinebug   = `{"postInfo": {"readPermission": {"homeId": "Mid"}}, "sourceType": "TIMELINE", "contents": {"text": "#Apple"}}`//method >> POST
	timelinefeed  = `{"postLimit": 10, "commentLimit": 1, "likeLimit": 20, "order": "TIME"}`//method >> GET
	timelinegroup = `{"homeId": "Mid", "commentLimit": 1, "likeLimit": 20, "sourceType": "TALKLOOM"}`//method >> GET
	timelineStr   = `{"postInfo": {"readPermission": {"type": "NONE"}}, "sourceType": "TIMELINE", "contents": {"text": "#TEST"}}`//method >> POST
	timelinelike  = `{"contentId": "postId", "actorId": "Mid", "likeType": "1001", "sharable": true}`//method >> POST
	timelinecome  = `{"contentId": "postId", "commentText": "mhgfdsa", "actorId": "Mid"}`//meyhod >> POST
	timelinetest  = `{"postInfo": {"readPermission": {"type": "ALL"}}, "sourceType": "TIMELINE", "contents": {"text": ""}}`
)

func main() {
	//req, err := http.NewRequest("POST", Config_URL, bytes.NewBuffer([]byte(Parameter)))
	req, err := http.NewRequest("POST", CREATE_URL, bytes.NewBuffer([]byte(timelinebug)))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Line/9.8.0")
	req.Header.Add("X-Line-Carrier", "44010, 1")
	req.Header.Add("X-Line-Mid", "Mid")
        req.Header.Add("X-Line-Application", "IOS\t9.8.0\tiOS\t12.3.1")
        req.Header.Add("X-LAL", "ja")
	req.Header.Add("X-Line-ChannelToken", "token")

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	fmt.Printf("%s %s\n Header : %s\n", resp.Proto, resp.Status, resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
