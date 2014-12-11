package main

import (
	"fmt"
	"time"
	"flag"
	"github.com/stealthycoin/evogo"
	"math/rand"
)

var (
	dummies *evogo.Population
	train_com = flag.Bool("t", false, "true to train computer")
)


// Genetic training stuff
type Gene struct {
	// 20 possible board states so we just map each to a number of pennies to remove
	state, choice int
}

func CreateGene(i int) evogo.Gene {
	return Gene{
		state: i+1,
		choice: rand.Intn(4) + 1,
	}
}

func MutateGene(g evogo.Gene) evogo.Gene {
	gene := g.(Gene)
	old := gene.choice
	for old == gene.choice {
		gene.choice = rand.Intn(4) + 1
	}
	return gene
}

func ShowGenes(i *evogo.Individual) {
	for _, g := range i.Genes() {
		gene := g.(Gene)
		fmt.Printf("%d ", gene.choice)
	}
}

func Fitness(indiv *evogo.Individual, others []*evogo.Individual) int {
	wins := 0
	// Battle other evolving members
	for i, enemy := range others {
		if indiv != enemy {
			if NewGame(NewComputer(indiv, false), NewComputer(others[i], false)).play() == 1 {
				wins++
			}
			if NewGame(NewComputer(others[i], false), NewComputer(indiv, false)).play() == 2 {
				wins++
			}
		}
	}

	// Battle dummies that don't evolve
	for _, enemy := range dummies.Individuals() {
		if NewGame(NewComputer(indiv, false), NewComputer(enemy, false)).play() == 1 {
			wins++
		}
		if NewGame(NewComputer(enemy, false), NewComputer(indiv, false)).play() == 2 {
			wins++
		}
	}

	return wins
}


func main() {
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	// Make the dummy population to train against
	dummies = evogo.NewPopulation(1000, 20, 20, CreateGene)

	var strongest *evogo.Individual
	if (*train_com) {
		pop := evogo.NewPopulation(100, 20, 20, CreateGene)
		evogo.Train(pop, 2198, Fitness, MutateGene)
		strongest = pop.Individuals()[0] // Get strongest individual
	} else {
		strongest = dummies.Individuals()[0]
	}

	for {
		// Start games
		var a, b string
		var pa, pb player
		fmt.Print("Player 1 is human or computer (h, c, q to quit): ")
		fmt.Scanf("%v", &a)
		if a == "q" {
			return
		}
		fmt.Print("Player 2 is human or computer (h, c, q to quit): ")
		fmt.Scanf("%v", &b)
		if b == "q" {
			return
		}

		if a == "h" {
			pa = human{}
		} else {
			pa = NewComputer(strongest, true)
		}
		if b == "h" {
			pb = human{}
		} else {
			pb = NewComputer(strongest, true)
		}

		winner := NewGame(pa, pb).play()
		fmt.Println("The winner is player", winner, "\n\n")
	}
}
