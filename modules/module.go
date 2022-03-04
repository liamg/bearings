package modules

import "github.com/liamg/bearings/powerline"

type Module interface {
	Render(w *powerline.Writer)
}
