package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
)

func main() {
	width := 1024
	height := 1024
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
}

func drawOctagon(canvas *svg.SVG, centerX int, centerY int, maxRadius int) {

}
