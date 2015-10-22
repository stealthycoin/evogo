package evogo

import (
	
	"math/rand"
	_"log"

)

var (
	breedMap map[string]breedFunc = map[string]breedFunc{
		"twoPointCrossover":twoPointCrossover,
		"cyclicCrossover":cyclicCrossover,
	}
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

//
// Breed two individuals i and j together using random cycles, returns a pair of new individual pointers
//
func cyclicCrossover(pop *Population, i, j int, m mutate) (*Individual,*Individual){
	pa := pop.individuals[i]
	pb := pop.individuals[j]

	

	if len(pa.chromosome) != len(pb.chromosome){
		panic("Cyclic crossover requires same length chromosomes")
	}

	//PARENT A
	// make a random array the length of the chromosome for cyclic crossover
	list := rand.Perm(len(pa.chromosome))
	cycleArray := make([]int, 0)

	last_num := 0

	for true {
	  val := list[last_num]
	  cycleArray = append(cycleArray, val)
	  if val == 0 {
	  	break
	  }
	  last_num = val
	}

	// Make first child
	ca := make([]Gene,0)
	//now mix the parents based off our cycleArray
	cycle_index := 0
	for i, _ := range pb.chromosome {
		

		if cycle_index >= len(cycleArray){
			//we ran out of values in our cycle array
			ca = append(ca, pb.chromosome[i])
			
		} else if i == cycleArray[cycle_index] {
			//grabbing from cycle array
			ca = append(ca, pa.chromosome[i])
			cycle_index++

		} else {
			//grabbing from other parent array
			ca = append(ca, pb.chromosome[i])
		}
		
	
	}

	
	//PARENT B
	// make a random array the length of the chromosome for cyclic crossover
	list = rand.Perm(len(pb.chromosome))
	cycleArray = make([]int, 0)

	last_num = 0

	for true {
	  val := list[last_num]
	  cycleArray = append(cycleArray, val)
	  if val == 0 {
	  	break
	  }
	  last_num = val
	}

	// Make first child
	cb := make([]Gene,0)
	//now mix the parents based off our cycleArray
	cycle_index = 0
	for i, _ := range pb.chromosome {
		
		if cycle_index >= len(cycleArray){
			//we ran out of values in our cycle array
			cb = append(cb, pa.chromosome[i])
			
		} else if i == cycleArray[cycle_index] {
			//grabbing from cycle array
			cb = append(cb, pb.chromosome[i])
			cycle_index++

		} else {
			//grabbing from other parent array
			cb = append(cb, pa.chromosome[i])
		}
		
	
	}

	return newIndividualWithGenes(ca), newIndividualWithGenes(cb)

}














