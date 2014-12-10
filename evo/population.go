package evo


// This is a population of individuals upon which various operations can be performed
type Population struct {
	individuals []*Individual
}

func NewPopulation(count, genesPerIndividual int, newGene newgene) *Population {
	p := &Population{
		individuals: newIndividuals(count, genesPerIndividual, newGene),
	}

	return p
}

func (p Population) Len() int {
	return len(p.individuals)
}

func (p Population) Swap(i, j int) {
	p.individuals[i], p.individuals[j] = p.individuals[j], p.individuals[i]
}

func (p Population) Less(i, j int) bool {
	return p.individuals[i].fitness < p.individuals[j].fitness
}
