/**
 * @Author: Resynz
 * @Date: 2020/12/30 11:55
 */
package lib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/rosbit/go-wget"
	"log"
	"net/http"
	"net/url"
	"robot-service/util"
	"time"
)

type DingTalkRobot struct {
	Name    string `json:"name"`
	Webhook string `json:"webhook"`
	Secret  string `json:"secret"`
}

func (dtR *DingTalkRobot) genSign(ts int64) string {
	h := hmac.New(sha256.New, []byte(dtR.Secret))
	needSign := fmt.Sprintf("%d\n%s", ts, dtR.Secret)
	h.Write([]byte(needSign))
	hs := h.Sum(nil)
	bs := base64.StdEncoding.EncodeToString(hs)
	return url.QueryEscape(bs)
}

func (dtR *DingTalkRobot) SayHello() error {
	data := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": fmt.Sprintf("Hello，我是机器人【%s】，来自于DingTalk。", dtR.Name),
		},
	}
	ts := time.Now().UnixNano() / 1e6
	reqUrl := fmt.Sprintf("%s&timestamp=%d&sign=%s", dtR.Webhook, ts, dtR.genSign(ts))
	method := "POST"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	status, content, _, err := wget.PostJson(reqUrl, method, data, header)
	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return fmt.Errorf("invalid http status:%d\n", status)
	}

	log.Printf("dingtalk robot response:%s\n", string(content))
	return nil
}

func (dtR *DingTalkRobot) genRemindDrinkContent(mobile string) string {
	cons := []string{
		fmt.Sprintf("@%s 药补不如食补,食补不如水补。你又到该补水的时候咯～", mobile),
		fmt.Sprintf("元气满满的一天，怎么少得了生命源泉Buff加持！@%s 补充水分后再战吧！", mobile),
		fmt.Sprintf("@%s 又被Bug折磨得无法自拔了？来杯水缓解下心情吧～", mobile),
		fmt.Sprintf("@%s 春风十里不如你，你不喝水可不行～", mobile),
		fmt.Sprintf("@%s 盼望着，盼望着，又到了喝水摸鱼的时间了，快快快～", mobile),
		fmt.Sprintf("@%s 吾虽浪迹天涯，却未迷失本心。喝杯水，想想自己的本心是啥吧", mobile),
		fmt.Sprintf("@%s 我曾踏足山巅，也曾进入低谷，二者都使我受益良多。工作无涯，生命有限，补补水，Bug最怕打不到的心～", mobile),
		fmt.Sprintf("@%s 好运不会眷顾傻瓜。不喝水就是傻瓜～", mobile),
		fmt.Sprintf("@%s 你们知道最强的武器是什么?没错就是———水!", mobile),
		fmt.Sprintf("@%s 断剑重铸之日，水杯归来之时。", mobile),
	}
	return cons[util.GenRandom(int64(len(cons)))]
}

func (dtR *DingTalkRobot) remindDrink(mobile string) error {

	data := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": dtR.genRemindDrinkContent(mobile),
		},
		"at": map[string]interface{}{
			"atMobiles": []string{mobile},
			"isAtAll":   false,
		},
	}
	ts := time.Now().UnixNano() / 1e6
	reqUrl := fmt.Sprintf("%s&timestamp=%d&sign=%s", dtR.Webhook, ts, dtR.genSign(ts))
	method := "POST"
	header := map[string]string{
		"Content-Type": "application/json",
	}
	status, content, _, err := wget.PostJson(reqUrl, method, data, header)
	if err != nil {
		return err
	}

	if status != http.StatusOK {
		return fmt.Errorf("invalid http status:%d\n", status)
	}

	log.Printf("dingtalk robot response:%s\n", string(content))
	return nil
}

func (dtR *DingTalkRobot) Remind(_type RemindType, mobile string) error {
	if mobile == "" {
		mobile = "13811213603"
	}
	var err error
	// todo more
	switch _type {
	case RemindDrink:
		err = dtR.remindDrink(mobile)
		break
	default:
		err = nil
	}
	return err
}
