package middleware

import (
	"fmt"
	"strconv"
	"time"
)

var TimeChecker = "TimeCheck"

type TimeCheckerType TimeCheck

type TimeCheck struct {
	timeStamp    time.Time
	functionName string
}

// Start запускает таймер для замера времени работы функции
func (tc *TimeCheck) Start(args ...string) {
	tc.timeStamp = time.Now()
}

// Stop находит время после запуска TimeChecker и записывает в лог.
func (tc *TimeCheck) Stop(args ...string) {
	latency := strconv.FormatInt(time.Since(tc.timeStamp).Milliseconds(), 10)

	fmt.Printf(tc.functionName+": "+latency+"ms: %p\n", tc)
}

// Write устанавливает значение functionName и result. args:  functionName,
func (tc *TimeCheck) Write(args ...string) {
	if len(args) != 1 {
		tc.functionName = "name is missing"
		return
	}
	tc.functionName = args[0]

	return
}
