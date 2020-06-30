package main

import (
	"math"
	"os"

	svg "github.com/ajstarks/svgo"
)

func main() {
	width := 1024
	height := 1024
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	drawOctagon(canvas, width/2, height/2, width*3/4/2)
	canvas.Text(width/2, height/2, "Octalysis", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
}

func drawOctagon(canvas *svg.SVG, centerX int, centerY int, maxRadius int) {
	canvas.Circle(centerX, centerY, maxRadius, "fill:none;stroke:black;stroke-width:2px;")
	pointsX := []int{}
	pointsY := []int{}
	for i := 0; i < 8; i++ {
		r := float64(maxRadius)
		x := float64(centerX) + r*math.Cos(float64(2)*math.Pi*float64(i)/float64(8))
		y := float64(centerY) + r*math.Sin(float64(2)*math.Pi*float64(i)/float64(8))

		pointsX = append(pointsX, int(x))
		pointsY = append(pointsY, int(y))
	}
	canvas.Polygon(pointsX, pointsY, "fill:blue;stroke:none;")

}
