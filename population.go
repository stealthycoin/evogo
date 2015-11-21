package evogo


// This is a population of individuals upon which various operations can be performed
type Population struct {
	InvertFitness  bool    // defaults to false, true makes lower fitness better
	Elitism        int     // How many individuals are tossed into the next generation right away
	Hoboism        int     // How many hobos (random individuals) are tossed into the next generation right away
	Tournament     int     // How many per tournament in selection phase
	TProb          float64 // Proability of winning a tournament
	MProb          float64 // Proability of a mutation
	Individuals    []*Individual
	BreedMethod    string  //What is the crossover algorithm (2 point, cyclic)
	DiversityFunc func(*Individual, []*Individual) int //optional diversity function
	LoopFunc func(*Population, int, int)
	MinG        int // minimum number of genes a chromosome is allowed to have
	MaxG        int // maximum number of genes a chromosome is allowed to have
	NewGene func(int)Gene //create gene function

	// Extra functions
	ShowGenes GeneDisplay // function to print a gene sequence

	//tranform variables
	MinFitness int
	MaxFitness int
	MinDiversity int
	MaxDiversity int

	//weighting metrics used when normalizing
	FitnessWeight float64
	DiversityWeight float64


	
	//store these during the fitness / diversity calc
	//perform a linear tranform (normalize values)
	//create a combined fitness / diversity metric, which is distance from max div, max fitness
	//less will compare the combined fitness
}

type ByFitness Population
type ByCombinedFitness Population


// Create a new population
func NewPopulation(count, minG, maxG int, newGene NewGene) *Population {
	return &Population{
		InvertFitness: false,
		Elitism:       5,
		Hoboism:       0,
		TProb:         0.75,
		MProb:         0.05,
		Tournament:    20,
		Individuals:   newIndividuals(count, minG, maxG, newGene),
		ShowGenes:     func(*Individual){},
		LoopFunc:     func(*Population, int, int){},
		BreedMethod: "twoPointCrossover",
		DiversityFunc: nil,
		FitnessWeight: 1.0,
		DiversityWeight: 1.0,
		NewGene: newGene,
		MinG: minG,
		MaxG: maxG,
	}
}


// Set whether negative or positive fitness is "better"
// false: higher fitness is better
// true: lower fitness is better
func (pop *Population) SetInvertFitness(invert bool) {
	pop.InvertFitness = invert
}

//
// Set how many individuals get promoted straight to the next generation
//
func (pop *Population) SetElitism(e int) {
	pop.Elitism = e
}

//
// Set how many random individuals get promoted straight to the next generation
//
func (pop *Population) SetHoboism(e int) {
	pop.Hoboism = e
}


//
//  Set breed function, options: ['cyclicCrossover', 'twoPointCrossover']
//
func (pop *Population) SetBreedMethod(b string) {
	pop.BreedMethod = b
}


//
// Set how large tournaments are for tournament selection (default is 20)
//
func (pop *Population) SetTournamentSize(size int) {
	pop.Tournament = size
}


//
// Set probability of the first individual winning the tournament
//
func (pop *Population) SetTournamentProbability(prob float64) {
	pop.TProb = prob
}


//
// Set probability of a mutation on a particular individual
//
func (pop *Population) SetMutationProbability(prob float64) {
	pop.MProb = prob
}

//
// Get probability of a mutation on a particular individual
//
func (pop *Population) GetMutationProbability() float64{
	return pop.MProb
}


//
// Set the printing function
//
func (pop *Population) SetShowIndividual(fn GeneDisplay) {
	pop.ShowGenes = fn
}
//
// Set LoopFunc, a function ran every generation
//
func (pop *Population) SetLoopFunction(fn LoopFunc){
	pop.LoopFunc = fn
}

//
// Set our optional diveristy function
//
func (pop *Population) SetDiversityFunction(fn DiversityFunc) {
	pop.DiversityFunc = fn
}

//
// Set Diversity Weight
//
func (pop *Population) SetDiversityWeight(w float64){
	pop.DiversityWeight = w
} 

//
// Get Diversity Weight
//
func (pop *Population) GetDiversityWeight() float64{
	return pop.DiversityWeight
} 


//
// Set Fitness Weight
//
func (pop *Population) SetFitnessWeight(w float64){
	pop.FitnessWeight = w
} 



//
// Fetch the array of individuals from the population
//
func (pop *Population) GetIndividuals() []*Individual {
	return pop.Individuals
}


//
// Sorting functions for the individuals in the population
//
func (p ByFitness) Len() int {
	return len(p.Individuals)
}


//
// Swapping function for the individuals in the population
//
func (p ByFitness) Swap(i, j int) {
	p.Individuals[i], p.Individuals[j] = p.Individuals[j], p.Individuals[i]
}


//
// 
//
func (p ByFitness) Less(i, j int) bool {
	if p.InvertFitness {
		return p.Individuals[i].Fitness < p.Individuals[j].Fitness
	}
	return p.Individuals[i].Fitness > p.Individuals[j].Fitness
}


//
// Sort by Combined Fitness
//
func (p ByCombinedFitness) Len() int {
	return len(p.Individuals)
}


//
// Swapping function for the individuals in the population (Combined Fitness)
//
func (p ByCombinedFitness) Swap(i, j int) {
	p.Individuals[i], p.Individuals[j] = p.Individuals[j], p.Individuals[i]
}


//
// Less function for individuals in the population (Combined Fitness)
//
func (p ByCombinedFitness) Less(i, j int) bool {
	if p.InvertFitness {
		return p.Individuals[i].CombinedFitness < p.Individuals[j].CombinedFitness
	}
	return p.Individuals[i].CombinedFitness > p.Individuals[j].CombinedFitness
}
