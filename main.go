package main

import (
	"fmt"
	"github.com/valyala/fastjson"
	"math/rand"
	"meow.tf/streamdeck/sdk"
	"os"
	"time"
)

var ticker = time.NewTicker(1 * time.Second)
var done = make(chan bool)
var debug = false
var log, _ = os.Create("log")
func main() {
	defer log.Close()
	log.WriteString(fmt.Sprintf("%v", os.Args))
	if debug {log.WriteString("In Main\n")}
	dice := []int{4, 6, 8, 10, 12, 20, 100}
	for _, sides := range dice {
		sdk.RegisterAction(fmt.Sprintf("com.hmschreck.d20.d%d", sides), HandlerGenerator(sides))
	}
	err := sdk.Open()
	if err != nil {
		sdk.Log("Died")
	}
	if debug {log.WriteString("Successfully opened up the connection to the SDK\n")}
	sdk.Wait()
}

func Roll(sides int) (result int) {
	return rand.Intn(sides) + 1
}

func HandlerGenerator(sides int) func(action, context string, payload *fastjson.Value, deviceId string) {
	return func(action, context string, payload *fastjson.Value, deviceID string) {
		sdk.SetTitle(context, fmt.Sprintf("D%d\n%d", sides, Roll(sides)), sdk.TargetBoth)
	}
}