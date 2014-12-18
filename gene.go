package evogo

// Function type for the New Gene function
type newgene func(int)Gene

// Takes in a gene array and prints it to the console (optional)
type genedisplay func(*Individual)

// Just giving Gene a better type name so we can tell what our function
// signatures are supposed to do
type Gene interface {}
