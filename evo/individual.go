package evo

type Individual struct {
	fitness int
	chromosome []Gene
}

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

func newIndividuals(count, genesPerIndividual int, newGene newgene) []*Individual {
	rv := make([]*Individual, count)
	for i := 0 ; i < count ; i++ {
		rv[i] = newIndividual(genesPerIndividual, newGene)
	}

	return rv
}


func (i *Individual) Genes() []Gene {
	return i.chromosome
}
