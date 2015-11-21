package evogo

import "math/rand"

type Individual struct {
	Fitness int 
	Chromosome []Gene
	Diversity int
	CombinedFitness float64
}


// Create a new individual with a known gene
func newIndividualWithGenes(chrom []Gene) *Individual {
	return &Individual{
		Fitness: 0,
		Chromosome: chrom,
		Diversity: 0,
	}
}

// Create a new individual
func newIndividual(minG, maxG int, newGene NewGene) *Individual {
	var genes int
	if minG == maxG {
		genes = minG
	} else {
		genes = minG + rand.Intn(maxG - minG)
	}

	rv :=  &Individual{
		Fitness: 0,
		Chromosome: make([]Gene, genes),
	}
	for i := 0 ; i < len(rv.Chromosome) ; i++ {
		rv.Chromosome[i] = newGene(i)
	}
	return rv;
}

// Create a given number of new individuals
func newIndividuals(count, minG, maxG int, newGene NewGene) []*Individual {
	rv := make([]*Individual, count)
	for i := 0 ; i < count ; i++ {
		rv[i] = newIndividual(minG, maxG, newGene)
	}

	return rv
}

// Get the genes from an individual
func (i *Individual) Genes() []Gene {
	return i.Chromosome
}

// Get the Fitness from an individual
func (i *Individual) GetFitness() int {
	return i.Fitness
}

// Get the Diversity from an individual
func (i *Individual) GetDiversity() int {
	return i.Diversity
}

// 
// Get individual CombinedFitness
//
func (i *Individual) GetCombinedFitness() float64{
	return i.CombinedFitness
}

// 
// Add Diversity to this individual
//
func (i *Individual) IncreaseDiversity(div int){
	i.Diversity += div
}











