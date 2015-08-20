package evogo

type breedFunc func (*Population, int, int, mutate) (*Individual,*Individual)
