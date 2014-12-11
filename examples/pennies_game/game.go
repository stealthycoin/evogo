package main

import (
	"fmt"
	"github.com/stealthycoin/evogo"
)


// Players must respond to a prompt with
// c,t the count and turn it is of the game
// and respond with an integer that is a valid
// number of pennies to remove from the board
type player interface {
	prompt(int,int)int
}

// Utility function
func most(c int) int {
	if c < 4 {
		return c
	}
	return 4
}


/*
*
* Human structure and its functions
*
*/
type human struct {}


func (h human) prompt(c, t int) int {
	fmt.Println("\n******* Player ", t, "*******")
	fmt.Print("There are ", c, " pennies on the table\nYou may remove 1-", most(c)," how many will you remove: ")
	var result int

	done := false
	for !done {
		_, err := fmt.Scanf("%d", &result)
		if err != nil || result < 1 || result > most(c){
			fmt.Print("Please enter a valid number: ")
		} else {
			done = true
		}
	}
	return result
}


/*
*
* Computer structure and its functions
*
*/
type computer struct {
	i *evogo.Individual
	talkative bool
}

func NewComputer(strategy *evogo.Individual, t bool) computer {
	return computer{
		i: strategy,
		talkative: t,
	}
}

func (com computer) prompt(c, t int) (rv int) {
	if com.talkative {
		defer func() {
			fmt.Println("\n******* Player ", t, "*******\nTakes ", rv, " pennies.")
		}()
	}
	rv = 1
	for _, g := range com.i.Genes() {
		gene := g.(Gene)
		if gene.state == c {
			rv = gene.choice
			return
		}
	}
	return // Should never happen but its always a safe move
}

/*
*
* Game structure and its functions
*
*/
type game struct {
	turn, total int
	a, b player
}

// Construct a new game with the given players
func NewGame(pa, pb player) game {
	return game{
		turn: 1,
		total: 20,
		a: pa,
		b: pb,
	}
}

// Plays game and returns winner
func (g game) play() int {
	for g.total > 0 {
		if g.turn == 1 {
			g.total -= g.a.prompt(g.total, g.turn);
			g.turn = 2
		} else {
			g.total -= g.b.prompt(g.total, g.turn);
			g.turn = 1
		}
		if g.total < 0 { // Cheating is an instant loss
			break
		}
	}
	return g.turn
}
