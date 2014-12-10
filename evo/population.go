package evo


// This is a population of individuals upon which various operations can be performed
type Population struct {
	invertFitness bool // defaults to false, true makes lower fitness better
	elitism int // How many individuals are tossed into the next generation right away
	tournament int // How many per tournament in selection phase
	tProb int // Proability of winning a tournament
	mProb float32 // Proability of a mutation
	individuals []*Individual

	// Extra functions
	showGenes genedisplay // function to print a gene sequence
}


// Create a new population
func NewPopulation(count, genesPerIndividual int, newGene newgene) *Population {
	return &Population{
		invertFitness: false,
		elitism: 0,
		tProb: 5,
		mProb: .05,
		tournament: int(count / 2),
		individuals: newIndividuals(count, genesPerIndividual, newGene),
		showGenes: func(*Individual){},
	}
}

func (pop *Population) InvertFitness(invert bool) {
	pop.invertFitness = invert
}

func (pop *Population) SetElitism(e int) {
	pop.elitism = e
}

func (pop *Population) SetShowIndividual(fn genedisplay) {
	pop.showGenes = fn
}

// Sorting functions for the individuals in the population
func (p Population) Len() int {
	return len(p.individuals)
}

func (p Population) Swap(i, j int) {
	p.individuals[i], p.individuals[j] = p.individuals[j], p.individuals[i]
}

func (p Population) Less(i, j int) bool {
	if p.invertFitness {
		return p.individuals[i].fitness > p.individuals[j].fitness
	}
	return p.individuals[i].fitness < p.individuals[j].fitness
}
