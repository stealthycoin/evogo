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

// Create a randomized gene
func CreateGene(i int) evogo.Gene {
	return Circle{
		p: image.Point{rand.Intn(targetImage.Bounds().Max.X), rand.Intn(targetImage.Bounds().Max.Y)},
		r: rand.Intn(50) + 2,
		col: colorChoices[rand.Intn(len(colorChoices))],
		a:  color.Alpha{uint8(rand.Intn(100) + 100)},
	}
}

// Evaluate how closely an individual reflects our target image, TODO optimize this
func fitness(i *evogo.Individual, others []*evogo.Individual) int {
	dst := image.NewNRGBA(image.Rect(0, 0, targetImage.Bounds().Max.X, targetImage.Bounds().Max.Y))
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
			i := dst.PixOffset(x,y)
			r, g, b := targetImage.(*image.NRGBA).Pix[i], targetImage.(*image.NRGBA).Pix[i+1], targetImage.(*image.NRGBA).Pix[i+2]
			r2, g2, b2 := dst.Pix[i], dst.Pix[i+1], dst.Pix[i+2]
			diff -= int(math.Sqrt( float64((r-r2)*(r-r2) + (g-g2)*(g-g2) + (b-b2)*(b-b2)) ))
		}
	}

	return diff
}

func DumpGene(i *evogo.Individual) {
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
	// Load the image
	flag.Parse()
	imgF, _ := os.Open(*target)
	defer imgF.Close()
	tImg , _, err := image.Decode(imgF)
	if err != nil {
		fmt.Println("decode",err)
		return
	}

	// Pre process the image
	targetImage = tImg
	preprocess()

	// Generate population
	pop := evogo.NewPopulation(500, 10, 20, CreateGene)

	// Configure evolution setting
	pop.SetShowIndividual(DumpGene)
	pop.SetElitism(0)

	evogo.Train(pop, 0, fitness, MutateGenes)
}
