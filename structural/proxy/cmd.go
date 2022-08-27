package proxy

import "fmt"

/*
代理是一种结构型设计模式，让你能提供真实服务对象的替代品给客户端使用。
代理接收客户端的请求并进行一些处理 （访问控制和缓存等），然后再将请求传递给服务对象。

代理对象拥有和服务对象相同的接口，这使得当其被传递给客户端时可与真实对象互换。

与decorator模式有点类似。
*/

func execute() {

	nginxServer := newNginxServer()
	appStatusURL := "/app/status"
	createuserURL := "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "POST")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

	httpCode, body = nginxServer.handleRequest(createuserURL, "GET")
	fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)
}
