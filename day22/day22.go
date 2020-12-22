package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	i, err := ioutil.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}
	input := string(i)
	fmt.Printf("Part 1: %v\n", SolveDay22Part1(input))
	fmt.Printf("Part 2: %v\n", SolveDay22Part2(input))
}

//SolveDay22Part1 plays a combat game and returns the score
func SolveDay22Part1(input string) (s int) {
	cardDeck := make(map[int][]int)
	for player, cards := range strings.Split(input, "\n\n") {
		for _, card := range strings.Split(cards, "\n") {
			if strings.HasPrefix(card, "Player") {
				continue
			}
			cardNum, err := strconv.Atoi(card)
			if err != nil {
				continue
			}
			cardDeck[player] = append(cardDeck[player], cardNum)
		}
	}

	_, winnerDeck := playCombat(cardDeck[0], cardDeck[1], false)
	for i, card := range winnerDeck {
		s += card * (len(winnerDeck) - i)
	}
	return
}

//SolveDay22Part2 plays a recursive combat game and returns the score
func SolveDay22Part2(input string) (s int) {
	cardDeck := make(map[int][]int)
	for player, cards := range strings.Split(input, "\n\n") {
		for _, card := range strings.Split(cards, "\n") {
			if strings.HasPrefix(card, "Player") {
				continue
			}
			cardNum, err := strconv.Atoi(card)
			if err != nil {
				continue
			}
			cardDeck[player] = append(cardDeck[player], cardNum)
		}
	}
	_, winnerDeck := playCombat(cardDeck[0], cardDeck[1], true)
	for i, card := range winnerDeck {
		s += card * (len(winnerDeck) - i)
	}
	return
}

//playCombat plays the (recursive) combat game with the two given cards, returns true if player1 wins or false if player2 wins and return the winner deck.
func playCombat(p1, p2 []int, recursive bool) (bool, []int) {
	cacheP1 := make(map[int][]int)
	cacheP2 := make(map[int][]int)
	var round int
	for {
		if len(p1) == 0 || len(p2) == 0 {
			break
		}
		if recursive {
			for i := round - 2; i >= 0; i-- {
				if checkDecks(p1, cacheP1[i]) && checkDecks(p2, cacheP2[i]) {
					return true, p1
				}
			}
		}

		var newPlayer1, newPlayer2 []int
		for _, val := range p1[1:] {
			newPlayer1 = append(newPlayer1, val)
		}
		for _, val := range p2[1:] {
			newPlayer2 = append(newPlayer2, val)
		}

		var winner bool
		if recursive && len(newPlayer1) >= p1[0] && len(newPlayer2) >= p2[0] {
			winner, _ = playCombat(newPlayer1[:p1[0]], newPlayer2[:p2[0]], true)
		} else {
			winner = p1[0] > p2[0]
		}

		if winner {
			//player1 wins
			newPlayer1 = append(newPlayer1, p1[0], p2[0])
		} else {
			//player2 wins
			newPlayer2 = append(newPlayer2, p2[0], p1[0])
		}

		p1 = make([]int, len(newPlayer1))
		p2 = make([]int, len(newPlayer2))
		if recursive {
			cacheP1[round] = make([]int, len(newPlayer1))
			cacheP2[round] = make([]int, len(newPlayer2))
		}
		for i, val := range newPlayer1 {
			p1[i] = val
			if recursive {
				cacheP1[round][i] = val
			}
		}
		for i, val := range newPlayer2 {
			p2[i] = val
			if recursive {
				cacheP2[round][i] = val
			}
		}
		round++
	}
	if len(p1) == 0 {
		return false, p2
	} else {
		return true, p1
	}
}

func checkDecks(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

//Helper functions
//stringListToSlice converts the list of strings (each string one row) to a slice
func stringListToSlice(list string) (s []string) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		s = append(s, line)
	}
	return
}

//intListToSlice converts the list of numbers (each number one row) to a slice
func intListToSlice(list string) (i []int) {
	for _, line := range strings.Split(strings.TrimSuffix(list, "\n"), "\n") {
		lineInt, err := strconv.Atoi(line)
		if err != nil {
			return nil
		}
		i = append(i, lineInt)
	}
	return
}
