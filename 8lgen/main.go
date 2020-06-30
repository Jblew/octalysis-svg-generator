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
	octalysisPoints := []int{
		10, 5, 2, 8, 1, 0, 1, 4,
	}
	descriptions := []string{
		"-A point\n-next long text\n-status bullets\n-a long text description of some fancy core drive use\n-sth else\n",
		"-A point\n-next long text\n-status bullets\n-a long text description of some fancy core drive use\n-sth else\n",
		"-A point\n-next long text\n-status bullets\n-a long text description of some fancy core drive use\n-sth else\n",
		"-A point\n-next long text\n-status bullets\n-a long text description of some fancy core drive use\n-sth else\n",
		"-A point\n-next long text\n-status bullets\n-a long text description of some fancy core drive use\n-sth else\n",
		"-A point\n-next long text\n-status bullets\n-a long text description of some fancy core drive use\n-sth else\n",
		"-A point\n-next long text\n-status bullets\n-a long text description of some fancy core drive use\n-sth else\n",
		"-A point\n-next long text\n-status bullets\n-a long text description of some fancy core drive use\n-sth else\n",
	}
	drawOctagon(canvas, width/2, height/2, width*2/3/2, octalysisPoints)
	drawDescriptions(canvas, width/2, height/2, width*3/4/2, descriptions)
	canvas.Text(width/2, height/2, "Octalysis", "text-anchor:middle;font-size:50px;fill:white")
	canvas.End()
}

func drawOctagon(canvas *svg.SVG, centerX int, centerY int, maxRadius int, octalysisPoints []int) {
	minRadius := float64(maxRadius) * 0.66
	radiusSpan := float64(maxRadius) - float64(minRadius)

	pointsX := []int{}
	pointsY := []int{}
	for i := 0; i < 8; i++ {
		r := minRadius + radiusSpan*float64(octalysisPoints[i])/10.0
		x := float64(centerX) + r*math.Cos(float64(2)*math.Pi*float64(i)/float64(8))
		y := float64(centerY) + r*math.Sin(float64(2)*math.Pi*float64(i)/float64(8))

		pointsX = append(pointsX, int(x))
		pointsY = append(pointsY, int(y))
	}
	canvas.Polygon(pointsX, pointsY, "fill:blue;stroke:none;")
	canvas.Circle(centerX, centerY, int(minRadius*0.9), "fill:none;stroke:white;stroke-width:2px;")
}

func drawDescriptions(canvas *svg.SVG, centerX int, centerY int, radius int, descriptions []string) {
	for i := 0; i < 8; i++ {
		r := float64(radius)
		angle := float64(2)*math.Pi*float64(i)/float64(8) - 0.1
		x := float64(centerX) + r*math.Cos(angle)
		y := float64(centerY) + r*math.Sin(angle)

		description := descriptions[i]
		canvas.Text(int(x), int(y), description, "text-anchor:middle;font-size:25px;fill:black;inline-size: 200px;shape-inside: circle(120px at 150px 150px);")
	}
}
