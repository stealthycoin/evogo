package evo

// Function type for the New Gene function
type newgene func()Gene

// Just giving Gene a better type name so we can tell what our function
// signatures are supposed to do
type Gene interface {}
