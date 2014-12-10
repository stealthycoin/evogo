package evo

type Individual struct {
	fitness int
	chromosome []Gene
}


// Create a new individual with a known gene
func newIndividualWithGenes(chrom []Gene) *Individual {
	return &Individual{
		fitness: 0,
		chromosome: chrom,
	}
}

// Create a new individual
func newIndividual(genesPerIndividual int, newGene newgene) *Individual {
	rv :=  &Individual{
		fitness: 0,
		chromosome: make([]Gene, genesPerIndividual),
	}
	for i := 0 ; i < genesPerIndividual ; i++ {
		rv.chromosome[i] = newGene()
	}
	return rv;
}

// Create a given number of new individuals
func newIndividuals(count, genesPerIndividual int, newGene newgene) []*Individual {
	rv := make([]*Individual, count)
	for i := 0 ; i < count ; i++ {
		rv[i] = newIndividual(genesPerIndividual, newGene)
	}

	return rv
}

// Get the genes from an individual
func (i *Individual) Genes() []Gene {
	return i.chromosome
}
