package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/jamesmcculloch/adventofcode/utils/queue"
	"github.com/jamesmcculloch/adventofcode/utils/stack"
)

func getDeckOfCards(cards []string) (*queue.IntQueue, error) {
	deck := &queue.IntQueue{}
	for _, card := range cards {
		cardNumber, err := strconv.Atoi(card)
		if err != nil {
			return nil, err
		}
		deck.Enqueue(cardNumber)
	}
	return deck, nil
}

func playCombat(player1, player2 *queue.IntQueue) *queue.IntQueue {
	var winner *queue.IntQueue
	for winner == nil {
		if player1.IsEmpty() {
			winner = player2
			break
		}
		if player2.IsEmpty() {
			winner = player1
			break
		}

		player1sCard, _ := player1.Dequeue()
		player2sCard, _ := player2.Dequeue()

		if player1sCard > player2sCard {
			player1.Enqueue(player1sCard)
			player1.Enqueue(player2sCard)
		} else {
			player2.Enqueue(player2sCard)
			player2.Enqueue(player1sCard)
		}
	}
	return winner
}

func determineWinnersScore(player *queue.IntQueue) int {
	stack := *&stack.IntStack{}
	for {
		card, err := player.Dequeue()
		if err != nil {
			break
		}
		stack.Push(card)
	}

	score := 0
	position := 1

	for {
		card, err := stack.Pop()
		if err != nil {
			break
		}
		score += (card * position)
		position++
	}
	return score
}

func playRecursiveCombat(player1 *queue.IntQueue, player2 *queue.IntQueue) (*queue.IntQueue, int) {
	previousGameStates := make(map[string]bool)
	for {
		gameState := getRoundState(player1, player2)
		if _, seen := previousGameStates[gameState]; seen {
			return player1, 1
		}
		previousGameStates[gameState] = true

		if player1.IsEmpty() {
			return player2, 2
		}
		if player2.IsEmpty() {
			return player1, 1
		}

		player1sCard, _ := player1.Dequeue()
		player2sCard, _ := player2.Dequeue()

		if player1sCard <= player1.Length() && player2sCard <= player2.Length() {
			player1Copy, _ := player1.Duplicate(player1sCard)
			player2Copy, _ := player2.Duplicate(player2sCard)

			_, winnerID := playRecursiveCombat(player1Copy, player2Copy)

			if winnerID == 1 {
				player1.Enqueue(player1sCard)
				player1.Enqueue(player2sCard)
			} else {
				player2.Enqueue(player2sCard)
				player2.Enqueue(player1sCard)
			}
			continue
		}

		if player1sCard > player2sCard {
			player1.Enqueue(player1sCard)
			player1.Enqueue(player2sCard)
		} else {
			player2.Enqueue(player2sCard)
			player2.Enqueue(player1sCard)
		}
	}
}

func getRoundState(player1 *queue.IntQueue, player2 *queue.IntQueue) string {
	var sb strings.Builder
	sb.WriteString(player1.String())
	sb.WriteString(" ")
	sb.WriteString(player2.String())
	return sb.String()
}

func main() {
	game, err := utils.LoadBlankLineSeparatedBlocksOfStringsFromFile("input")
	if err != nil {
		panic(err)
	}

	player1, err := getDeckOfCards(game[0][1:])
	if err != nil {
		panic(err)
	}
	player2, err := getDeckOfCards(game[1][1:])
	if err != nil {
		panic(err)
	}

	winner := playCombat(player1, player2)

	fmt.Printf("part 1: %d\n", determineWinnersScore(winner))

	player1, err = getDeckOfCards(game[0][1:])
	if err != nil {
		panic(err)
	}
	player2, err = getDeckOfCards(game[1][1:])
	if err != nil {
		panic(err)
	}

	winner, _ = playRecursiveCombat(player1, player2)

	fmt.Printf("part 2: %d\n", determineWinnersScore(winner))
}
