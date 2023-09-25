package internal

import (
	"fmt"
	"log"
	"strings"

	gohook "github.com/robotn/gohook"
)

func Run() {
	config, err := loadConfigFromEnviron()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("--- Please press %-20s to stop hook ---\n", strings.Join(config.Base.StopKey, "+"))
	gohook.Register(gohook.KeyDown, config.Base.StopKey, func(e gohook.Event) {
		fmt.Println("ctrl-shift-q")
		gohook.End()
	})

	fmt.Printf("--- Please press %-20s to fire repeated-click ---\n", strings.Join(config.RepeatedClick.FireKey, "+"))
	fmt.Printf("--- Please press %-20s to stop repeated-click ---\n", strings.Join(config.RepeatedClick.StopKey, "+"))
	clickCtxCancelStore := newCtxCancelStore()
	gohook.Register(gohook.KeyDown, config.RepeatedClick.FireKey, func(e gohook.Event) {
		go fireRepeatedClick(clickCtxCancelStore, config.RepeatedClick.Interval, config.RepeatedClick.Duration)
	})
	gohook.Register(gohook.KeyDown, config.RepeatedClick.StopKey, func(e gohook.Event) {
		stopRepeatedClick(clickCtxCancelStore)
	})
	defer stopRepeatedClick(clickCtxCancelStore)

	s := gohook.Start()
	<-gohook.Process(s)
}
