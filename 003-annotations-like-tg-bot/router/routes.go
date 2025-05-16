package router

//go:generate go run generate_routes.go
import (
	"github.com/NicoNex/echotron/v3"
)

var routes = make(map[string]func(*echotron.API, *echotron.Message))

func registerRoute(cmd string, handler func(*echotron.API, *echotron.Message)) {
	routes[cmd] = handler
}

func GetHandler(cmd string) (func(*echotron.API, *echotron.Message), bool) {
	h, ok := routes[cmd]
	return h, ok
}
