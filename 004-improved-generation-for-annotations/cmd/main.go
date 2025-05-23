package main

import (
	"github.com/prokhorind/go_course/004-improved-generation-for-annotations/bot"
)

const token = ""

func main() {
	engine := bot.NewEchotronBot(token)
	engine.Run()
}
