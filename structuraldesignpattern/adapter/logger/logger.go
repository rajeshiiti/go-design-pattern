package main

import "fmt"

type Logger interface {
	Log(message string)
}

type ThirdPartyLogger struct{}

func (tp *ThirdPartyLogger) WriteLog(message string) {
	fmt.Println("Writing log in third party logger: ", message)
}

type ThirdPartyLoggerAdapter struct {
	thirdPartyLogger *ThirdPartyLogger
}

func (tpa *ThirdPartyLoggerAdapter) Log(message string) {
	tpa.thirdPartyLogger.WriteLog(message)
}

func main() {
	tp := &ThirdPartyLogger{}

	tpa := &ThirdPartyLoggerAdapter{
		thirdPartyLogger: tp,
	}

	message := "Log message to be logged"

	tpa.Log(message)

}
