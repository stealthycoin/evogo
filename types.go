package evogo

// Function type for the New Gene function
type NewGene func(int)Gene

// Function type for a mutation function on a gene
type Mutate func(Gene)Gene

// Takes in a gene array and prints it to the console (optional)
type GeneDisplay func(*Individual)

// Just giving Gene a better type name so we can tell what our function
// signatures are supposed to do
type Gene interface {}

type DiversityFunc func (*Individual, []*Individual) int

type LoopFunc func (*Population, int, int)

// Fitness evaluation function type
type fitness func(*Individual,[]*Individual) int

type BreedFunc func (*Population, int, int, Mutate) (*Individual,*Individual)
