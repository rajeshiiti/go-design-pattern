package main

import "fmt"

type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string)
}

type AuthHandler struct {
	next Handler
}

func (a *AuthHandler) SetNext(handler Handler) Handler {
	a.next = handler
	return handler
}

func (a *AuthHandler) Handle(request string) {
	if request != "auth" {
		fmt.Println("Request is not authenticated")
		return
	}
	fmt.Println("Request is authenticated.")
	if a.next != nil {
		a.next.Handle(request)
	}

}

type LogHandler struct {
	next Handler
}

func (l *LogHandler) SetNext(handler Handler) Handler {
	l.next = handler
	return handler
}

func (l *LogHandler) Handle(request string) {
	fmt.Println("Logging the request data: ", request)

	if l.next != nil {
		l.next.Handle(request)
	}
}

type AuthzHandler struct {
	next Handler
}

func (a *AuthzHandler) SetNext(handler Handler) Handler {
	a.next = handler
	return handler
}

func (a *AuthzHandler) Handle(request string) {
	if request != "authz" {
		fmt.Println("Request authz is failed.")
		return
	}
	fmt.Println("Autz request passed.")
	if a.next != nil {
		a.next.Handle(request)
	}
}

func main() {
	auth := &AuthHandler{}
	logger := &LogHandler{}
	authz := &AuthzHandler{}

	auth.SetNext(logger).SetNext(authz)

	fmt.Println("Processing authenticate request.")
	auth.Handle("auth")

	fmt.Println("Processing unautheticate request.")
	auth.Handle("authz")

	fmt.Println("Processing authenticate request.")
	auth.Handle("auth")
}
