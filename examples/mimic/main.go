package main

import (
	"os"
	"fmt"
	"flag"
	"image"
	"image/color"
	_ "image/png"
	"github.com/stealthycoin/evogo"
	"code.google.com/p/draw2d/draw2d"
)

var (
	target = flag.String("target", "default.png", "Please set a target image file to load.")
	tImg *image.Image
)

/** Struct for circle **/
type Circle struct {
	p image.Point
	r int
}

/**
 * The following is taken from:
 * http://blog.golang.org/go-imagedraw-package
 * Excellent use of masking!!!
 */
func (c *Circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *Circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *Circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}


type Gene struct {
	C color.Color // Color
	P image.Point // Center
	R int         // Radius
}

func CreateGene(i int) evogo.Gene {
	return Gene{
		C: color.White,
		P: image.Point{0, 0},
		R: 4,
	}
}

func fitness(i *evogo.Individual, others []*evogo.Individual) int {
	portrait := image.NewRGBA(image.Rect(0, 0, tImg.Width, tImg.Height))
	gc := image.NewGraphicContext(portrait)
	for i, g := range i.GetIndividuals() {
		gene := g.(Gene)
		gc.Circle
	}
}

func MutateGene(g evogo.Gene) evogo.Gene {
	return g
}


func main() {
	flag.Parse()
	imgF, _ := os.Open(*target)
	defer imgF.Close()
	tImg , _, err := image.Decode(imgF)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tImg.At(10,10))

	pop := evogo.NewPopulation(1000, 10, 10, CreateGene)
	evogo.Train(pop, 0, fitness, MutateGene)
}
