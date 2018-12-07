package main

import (
	"strings"
	"fmt"
	"net/http"
	"reflect"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"net/url"
)



type Results struct{
	Msg string
	Status string
	Obj string
}

func HttpGet(){
	url := "http://www.baidu.com"
	resp,err := http.Get(url)
	if err !=nil{
		fmt.Printf("%s",err)
		return 
	}
	defer resp.Body.Close()
	
	headers :=resp.Header
	for k ,v := range headers{
		fmt.Printf("%s:%s",k,v)
	}

	fmt.Printf("status %s,status code %d \n",resp.Status,resp.StatusCode)
	fmt.Printf("proto %s \n",resp.Proto)
	fmt.Printf("length %d \n",resp.ContentLength)
	fmt.Printf("transfer encoding %s \n",resp.TransferEncoding)
	fmt.Printf("resp Uncompressed %s \n",resp.Uncompressed)
	fmt.Println(reflect.TypeOf(resp.Body))
	buf := bytes.NewBuffer(make([]byte,0,512))
	length,_ := buf.ReadFrom(resp.Body) //ReadFrom和ioutil.Readall貌似都是作用一样
	fmt.Println(len(buf.Bytes()))
	fmt.Println(length)
	//fmt.Println(string(buf.Bytes()))
}


func HttpPost(){
	type J struct{
		status string
		message string
	}

	//url := "https://tcc.taobao.com/cc/json/mobile_tel_segment.htm?tel=18588505587"
	//url :="https://tcc.taobao.com/cc/json/mobile_tel_segment.htm?tel=18588505587"
	//http.Post url为域名时，使用strings.NewReader添加附加参数
	//或者url完整地附带参数，strings.NewReader为空 如下，不携带三个参数会报错
	url3 := "http://api.map.baidu.com/telematics/v3/weather?location=嘉兴&output=json&ak=5slgyqGDENN7Sy7pw29IUvrZ"
	resp,err := http.Post(url3,"application/x-www-form-urlencoded",strings.NewReader(""))
	if err != nil{
		fmt.Printf("resp error")
		fmt.Println(err)
		return 
	}
	//方法2
	url2 :="https://tcc.taobao.com/cc/json/mobile_tel_segment.htm"
	postParam := url.Values{
		"tel": {"18588505587"},
	}
	resp2,err2 := http.PostForm(url2,postParam)
	if err2 != nil{
		fmt.Printf("req2 error")
		fmt.Println(err2)
		return 
	}


	defer resp.Body.Close()
	defer resp2.Body.Close()

	body,err :=ioutil.ReadAll(resp.Body)
	if err !=nil{
		fmt.Printf("body error")
		fmt.Println(err)
		return 
	}
	fmt.Println(string(body))
	se := string(body)
	fmt.Println(reflect.TypeOf(se))
	//不知道为何不能解json
	j := &J{}
	json.Unmarshal(body,j)
	fmt.Println(j)

	body2,err2 := ioutil.ReadAll(resp.Body)
	if err2 !=nil{
		fmt.Print("body2 error")
		fmt.Println(err)
		return 
	}
	fmt.Println(string(body2))

}

func HttpSetArgs(){
	client :=&http.Client{}
	url := "http://10.40.0.82:9200/"
	req,err := http.NewRequest("get",url,strings.NewReader(""))
	if err !=nil{
		fmt.Printf("make Request error")
		fmt.Println(err)
		return 
	}
	req.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Content-Type","application/json")
	
	for i:=0;i<3;i++{
		resp,err :=client.Do(req)
		if err !=nil {
			fmt.Printf("req is error")
			fmt.Println(err)
			return 
		}
		
		body,err := ioutil.ReadAll(resp.Body)
		if err !=nil{
			fmt.Printf("body is error")
			fmt.Println(err)
			return
		}
		fmt.Println(resp.Request.Header)
		fmt.Println("123")
		fmt.Println(string(body))
	}

}


func main() {
	//HttpGet()
	//HttpPost()
	HttpSetArgs()
}