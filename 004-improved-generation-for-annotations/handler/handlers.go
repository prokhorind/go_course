package handler

// @route /start
func StartHandler(ctx Telegram) {
	ctx.SendMessage("Welcome!")
}

// @route /help
func HelpHandler(ctx Telegram) {
	ctx.SendMessage("Try /start or /help")
}

// @route /info
func InfoHandler(ctx Telegram) {
	ctx.SendMessage("Info handler")
}
