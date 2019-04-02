package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	url2 "net/url"
	"strconv"
	"time"
)

type signature struct {
	appid int
	uin int
	daid int
	style int;
	hide_border int;
	s_url string;
	target string;
	a int;
}

const urlSuccess = "https://jiazhang.qq.com/zk/ptLoginCallBack.html"
const urlRaw = "https://xui.ptlogin2.qq.com/cgi-bin/xlogin"
const ptqrshowUrl = "https://ssl.ptlogin2.qq.com/ptqrshow"
const ptqrloginUrl = "https://ssl.ptlogin2.qq.com/ptqrlogin"

var login_sig string
var qr_sig string

//
//common structToString(obj interface{}) string {
//	var str string
//
//	for o := range obj {
//
//	}
//	return str
//}

func getCookie(resp *http.Response, name string) string {
	for _, k := range resp.Cookies() {
		fmt.Println(k.Name)
		if k.Name == name {
			return k.Value
		}
	}

	return ""
}

func hash33(qr_sig string) int {
	hash := int(0) // 001 010 100 000 101 ,hash后的分布更好一些
	//s := fmt.Sprintf("%x", md5.Sum([]byte(qr_sig)))
	s := []byte(qr_sig)
	for i := 0; i < len(s); i++ {
		hash += (hash << 5) + int(s[i])
	}

	return 0x7FFFFFFF & hash
}

func get_signature() error {
	//params := signature{
	//	appid: 11,
	//	uin:1,
	//	daid:527,
	//	style:1,
	//	hide_border:1,
	//	s_url: urlSuccess,
	//	target: "self",
	//	a: 13123131,
	//}
	url := urlRaw + "?" + "appid=" + strconv.Itoa(1600000932) + "&default_uin=0&daid=527&style=33&hide_border=1&target=self&" +
		"_=" + fmt.Sprintf("%f", rand.Float64()) + "&s_url=" + url2.QueryEscape(urlSuccess)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get_signature err=%v", err)
		return err
	}
	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	// handle error
	//}
	//
	//fmt.Println(string(body))

	login_sig = getCookie(resp, "pt_login_sig")

	if login_sig == "" {
		return errors.New("login sig error")
	}

	return nil
}

func ptqrshow() error {
	url := ptqrshowUrl + "?" + "appid=" + strconv.Itoa(1600000932) + "&e=2&l=M&s=3&d=72&v=4&pt_3rd_aid=0" + "&t=" + fmt.Sprintf("%f", rand.Float64()) +
		"&daid=527"
	req, err := http.Get(url)
	if err != nil {
		fmt.Println("ptqrshow err=%v", err)
		return err
	}
	qr_sig = getCookie(req, "qrsig")
	return nil
}

func ptqrlogin() error {
	millis := time.Now().UnixNano() / 1e6
	action := fmt.Sprintf("0-0-%d", millis)
	sig := hash33(qr_sig)
	url := ptqrloginUrl + "?u1=" + url2.QueryEscape(urlSuccess) + "&ptredirect=0&h=1&t=1&g=1&from_ui=1&ptlang=2052&"  + "&ptqrtoken=" + strconv.Itoa(sig) +
		"&action=" + action + "&js_ver=90305&login_sig=" + login_sig + "&aid=" + strconv.Itoa(1600000932) + "&daid=527"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get_signature err=%v", err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("handle error")
		return errors.New("handle error")
	}

	fmt.Println(string(body))
	return nil
}

func main()  {
	if err := get_signature(); err == nil {
		if err := ptqrshow(); err == nil {
			if err := ptqrlogin(); err == nil {

			}
		}
	}
}