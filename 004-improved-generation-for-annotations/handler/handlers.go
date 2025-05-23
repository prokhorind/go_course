package handler

// @route /start
func StartHandler(ctx Context) {
	ctx.SendMessage("Welcome!")
}

// @route /help
func HelpHandler(ctx Context) {
	ctx.SendMessage("Try /start or /help")
}

// @route /info
func InfoHandler(ctx Context) {
	ctx.SendMessage("Info handler")
}
