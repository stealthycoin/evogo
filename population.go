package evogo


// This is a population of individuals upon which various operations can be performed
type Population struct {
	invertFitness  bool    // defaults to false, true makes lower fitness better
	elitism        int     // How many individuals are tossed into the next generation right away
	tournament     int     // How many per tournament in selection phase
	tProb          float64 // Proability of winning a tournament
	mProb          float64 // Proability of a mutation
	individuals    []*Individual
	crossoverMethod string //What is the crossover algorithm (2 point, cyclic)

	// Extra functions
	showGenes genedisplay // function to print a gene sequence
}


// Create a new population
func NewPopulation(count, minG, maxG int, newGene newgene) *Population {
	return &Population{
		invertFitness: false,
		elitism:       5,
		tProb:         0.75,
		mProb:         0.05,
		tournament:    20,
		individuals:   newIndividuals(count, minG, maxG, newGene),
		showGenes:     func(*Individual){},
		crossoverMethod: "twoPointCrossover",
	}
}


// Set whether negative or positive fitness is "better"
// false: higher fitness is better
// true: lower fitness is better
func (pop *Population) InvertFitness(invert bool) {
	pop.invertFitness = invert
}

//
// Set how many individuals get promoted straight to the next
//
func (pop *Population) SetElitism(e int) {
	pop.elitism = e
}


//
// Set how large tournaments are for tournament selection (default is 20)
//
func (pop *Population) SetTournamentSize(size int) {
	pop.tournament = size
}


//
// Set probability of the first individual winning the tournament
//
func (pop *Population) SetTournamentProbability(prob float64) {
	pop.tProb = prob
}


//
// Set probability of a mutation on a particular individual
//
func (pop *Population) SetMutationProbability(prob float64) {
	pop.mProb = prob
}


//
//
//
func (pop *Population) SetShowIndividual(fn genedisplay) {
	pop.showGenes = fn
}

//
// Fetch the array of individuals from the population
//
func (pop *Population) Individuals() []*Individual {
	return pop.individuals
}


//
// Sorting functions for the individuals in the population
//
func (p Population) Len() int {
	return len(p.individuals)
}


//
//
//
func (p Population) Swap(i, j int) {
	p.individuals[i], p.individuals[j] = p.individuals[j], p.individuals[i]
}


//
//
//
func (p Population) Less(i, j int) bool {
	if p.invertFitness {
		return p.individuals[i].fitness < p.individuals[j].fitness
	}
	return p.individuals[i].fitness > p.individuals[j].fitness
}
