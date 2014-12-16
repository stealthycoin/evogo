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
)

/** Struct for circle **/
type Circle struct {
	p image.Point
	r int
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
		return color.Alpha{100}
	}

	return color.Alpha{0}
}


type Gene struct {
	circ Circle
	col color.RGBA
}

func CreateGene(i int) evogo.Gene {
	return Gene{
		circ: Circle{
			p: image.Point{rand.Intn(targetImage.Bounds().Max.X), rand.Intn(targetImage.Bounds().Max.Y)},
			r: rand.Intn(targetImage.Bounds().Max.X / 2 - 5) + 5,
		},
		col: color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255},
	}
}

func fitness(i *evogo.Individual, others []*evogo.Individual) int {
	dst := image.NewRGBA(image.Rect(0, 0, targetImage.Bounds().Max.X, targetImage.Bounds().Max.Y))
	draw.Draw(dst, dst.Bounds(), &image.Uniform{color.RGBA{0,0,0,255}}, image.ZP, draw.Src)
	for _, g := range i.Genes() {
		gene := g.(Gene)
		src := &image.Uniform{gene.col}
		draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &gene.circ, image.ZP, draw.Over)
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
		gene := g.(Gene)
		src := &image.Uniform{gene.col}
		draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &gene.circ, image.ZP, draw.Over)
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


func main() {
	flag.Parse()
	imgF, _ := os.Open(*target)
	defer imgF.Close()
	tImg , _, err := image.Decode(imgF)
	if err != nil {
		fmt.Println("decode",err)
		return
	}
	fmt.Println(tImg.Bounds())
	fmt.Println(tImg.At(10,10))

	targetImage = tImg

	pop := evogo.NewPopulation(1000, 50, 50, CreateGene)
	pop.SetShowIndividual(ShowGenes)
	evogo.Train(pop, 0, fitness, MutateGene)
}
