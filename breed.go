package evogo

import (
	
	"math/rand"
)



var (
	crossoverMap map[string]crossoverFunc
)


//
// Breed two individuals i and j together, returns a pair of new individual pointers
//
func twoPointCrossover(pop *Population, i, j int, m mutate) (*Individual,*Individual){
	pa := pop.individuals[i]
	pb := pop.individuals[j]

	// Two point crossover (assume lengths are the same for now)
	firstPoint := rand.Intn(len(pa.chromosome)-1)
	secondPoint := rand.Intn(len(pa.chromosome)-1)
	if firstPoint > secondPoint {
		firstPoint, secondPoint = secondPoint, firstPoint
	}

	// Make first child
	ca := make([]Gene,0)
	ca = append(ca, pa.chromosome[:firstPoint]...)
	ca = append(ca, pb.chromosome[firstPoint:secondPoint]...)
	ca = append(ca, pa.chromosome[secondPoint:]...)
	// Mutation for first child
	for i, _ := range ca {
		if rand.Float64() < pop.mProb {
			ca[i] = m(ca[i])
		}
	}


	// Make second child
	cb := make([]Gene,0)
	cb = append(cb, pb.chromosome[:firstPoint]...)
	cb = append(cb, pa.chromosome[firstPoint:secondPoint]...)
	cb = append(cb, pb.chromosome[secondPoint:]...)
	for i, _ := range cb {
		if rand.Float64() < pop.mProb {
			ca[i] = m(cb[i])
		}
	}

	return newIndividualWithGenes(ca), newIndividualWithGenes(cb)

}