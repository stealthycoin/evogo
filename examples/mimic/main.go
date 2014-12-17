package main

import (
	"os"
	"fmt"
	"math"
	"math/rand"
	"flag"
	"image"
	"strconv"
	"image/draw"
	"image/color"
	"image/png"
	_ "image/jpeg"
	_ "image/gif"
	"github.com/stealthycoin/evogo"
)

var (
	target = flag.String("target", "default.png", "Please set a target image file to load.")
	targetImage image.Image
	gen int = 1
	colorChoices []color.RGBA
)

/** Struct for circle **/
type Circle struct {
	p image.Point
	r int
	col color.RGBA
	a color.Alpha
}

func (c *Circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *Circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *Circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	// if xx*xx+yy*yy < rr*rr {
	// 	return color.Alpha{255}
	// }
	if xx*xx+yy*yy < (rr)*(rr) {
		return c.a
	}

	return color.Alpha{0}
}

func CreateGene(i int) evogo.Gene {
	return Circle{
		p: image.Point{rand.Intn(targetImage.Bounds().Max.X), rand.Intn(targetImage.Bounds().Max.Y)},
		r: rand.Intn(50) + 5,
		col: colorChoices[rand.Intn(len(colorChoices))],
		a: color.Alpha{uint8(rand.Intn(100) + 100)},
	}
}

func fitness(i *evogo.Individual, others []*evogo.Individual) int {
	dst := image.NewRGBA(image.Rect(0, 0, targetImage.Bounds().Max.X, targetImage.Bounds().Max.Y))
	draw.Draw(dst, dst.Bounds(), &image.Uniform{color.RGBA{0,0,0,255}}, image.ZP, draw.Src)
	for _, g := range i.Genes() {
		gene := g.(Circle)
		src := &image.Uniform{gene.col}
		draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &gene, image.ZP, draw.Over)
	}

	// Compare target image and dst image
	diff := 0;
	for x := 0 ; x < targetImage.Bounds().Max.X ; x++ {
		for y := 0 ; y < targetImage.Bounds().Max.Y ; y++ {
			r,g,b,_ := targetImage.At(x,y).RGBA()
			r2,g2,b2,_ := dst.At(x,y).RGBA()
			diff -= int(math.Sqrt( float64((r-r2)*(r-r2) + (g-g2)*(g-g2) + (b-b2)*(b-b2)) ))
		}
	}

	return diff
}

func ShowGenes(i *evogo.Individual) {
	// render
	dst := image.NewRGBA(image.Rect(0, 0, targetImage.Bounds().Max.X, targetImage.Bounds().Max.Y))
	draw.Draw(dst, dst.Bounds(), &image.Uniform{color.RGBA{0,0,0,255}}, image.ZP, draw.Src)
	for _, g := range i.Genes() {
		gene := g.(Circle)
		src := &image.Uniform{gene.col}
		draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &gene, image.ZP, draw.Over)
	}

	// Write to a file.
	writer, err := os.Create(strconv.Itoa(gen) + ".png")
	gen++
	defer writer.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		err = png.Encode(writer, dst)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func MutateGene(g evogo.Gene) evogo.Gene {
	return CreateGene(0)
}

func preprocess() {
	colorChoices = make([]color.RGBA, 0, targetImage.Bounds().Max.X * targetImage.Bounds().Max.Y)
	for x := 0 ; x < targetImage.Bounds().Max.X ; x++ {
		for y := 0 ; y < targetImage.Bounds().Max.Y ; y++ {
			r, g, b, _ := targetImage.At(x,y).RGBA()
			colorChoices = append(colorChoices, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
	}
}

func main() {
	flag.Parse()
	imgF, _ := os.Open(*target)
	defer imgF.Close()
	tImg , _, err := image.Decode(imgF)
	if err != nil {
		fmt.Println("decode",err)
		return
	}
	targetImage = tImg
	preprocess()

	pop := evogo.NewPopulation(1000, 100, 100, CreateGene)
	pop.SetShowIndividual(ShowGenes)
	evogo.Train(pop, 0, fitness, MutateGene)
}
