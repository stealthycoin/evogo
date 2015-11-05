package evogo

import "math/rand"

type Individual struct {
	fitness int
	chromosome []Gene
	diversity int
}


// Create a new individual with a known gene
func newIndividualWithGenes(chrom []Gene) *Individual {
	return &Individual{
		fitness: 0,
		chromosome: chrom,
		diversity: 0,
	}
}

// Create a new individual
func newIndividual(minG, maxG int, newGene newgene) *Individual {
	var genes int
	if minG == maxG {
		genes = minG
	} else {
		genes = minG + rand.Intn(maxG - minG)
	}

	rv :=  &Individual{
		fitness: 0,
		chromosome: make([]Gene, genes),
	}
	for i := 0 ; i < len(rv.chromosome) ; i++ {
		rv.chromosome[i] = newGene(i)
	}
	return rv;
}

// Create a given number of new individuals
func newIndividuals(count, minG, maxG int, newGene newgene) []*Individual {
	rv := make([]*Individual, count)
	for i := 0 ; i < count ; i++ {
		rv[i] = newIndividual(minG, maxG, newGene)
	}

	return rv
}

// Get the genes from an individual
func (i *Individual) Genes() []Gene {
	return i.chromosome
}

// Get the fitness from an individual
func (i *Individual) Fitness() int {
	return i.fitness
}

// Get the fitness from an individual
func (i *Individual) Diversity() int {
	return i.diversity
}

// 
// Add diversity to this individual
//
func (i *Individual) IncreaseDiversity(div int){
	i.diversity += div
}

