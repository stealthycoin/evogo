package evogo

type crossoverFunc func (*Population, int, int, mutate) (*Individual,*Individual)
