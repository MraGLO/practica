package http

import "github.com/MraGLO/practica/internal/services"

type Handlers struct {
	services *services.Services
}

func MakeHandlers(apiservices *services.Services) *Handlers {
	return &Handlers{apiservices}
}

func isValidString(s string) (str string, b bool) {
	b = false
	for _, char := range s {
		if s[0] == ' ' {
			b = true
			str = "первый символ является пробелом"
			return
		}
		if char == '\t' {
			b = true
			str = "в имени присутсвует табуляция"
			return
		}
		if char == '\n' {
			b = true
			str = "в имени присутсвует несколько строк"
			return
		}
	}
	return
}
