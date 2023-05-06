package main

import (
	"html/template"
	"log"
	"net/http"
)

func loginHandler(writer http.ResponseWriter, request *http.Request) {
	// 获取请求方式
	method := request.Method
	println("请求方式", method)

	// 如果此处读取了body数据则ParseForm()读取不到数据了
	// body := request.Body
	// length := request.ContentLength
	// println("length", length)
	// if length > 0 {
	// 	p := make([]byte, length)
	// 	body.Read(p)
	// 	println(string(p)) // 打印：userName1=faffa&password1=fafaf
	// }

	if request.Method == "GET" {
		// 如果时get请求则加载login.html，(该login.html与go文件同目录)
		t, _ := template.ParseFiles("D:\\GolandProjects\\awesomeProject\\src\\go_code\\project01\\main\\login.html")
		t.Execute(writer, nil)
	} else {
		// 解析post url 参数
		url := request.URL
		values := url.Query()
		userName := values.Get("userName")
		password := values.Get("password")
		println(userName, password)

		// 解析表单参数
		request.ParseForm()
		userName1 := request.Form.Get("userName1")
		password1 := request.Form.Get("password1")
		println(userName1, password1)
	}

}

func sukiHandler(suki http.ResponseWriter, request *http.Request) {
	htm := `<html>
<head>
    <title>登录</title>
</head>
<body>
<img src="https://img2.baidu.com/it/u=2766285767,2568602672&fm=253&fmt=auto&app=138&f=PNG?w=450&h=630">
</body>
</html>`
	// 响应报文头200成功
	suki.WriteHeader(200)
	// 响应报文体
	suki.Write([]byte(htm))
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/suki", sukiHandler)
	err := http.ListenAndServe("0.0.0.0:8090", nil)
	log.Fatal(err)
}
