package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jamesmcculloch/adventofcode/utils"
	"github.com/jamesmcculloch/adventofcode/utils/queue"
	"github.com/jamesmcculloch/adventofcode/utils/stack"
)

type deckOfCards struct {
	q *queue.IntQueue
}

func new() *deckOfCards {
	return &deckOfCards{
		q: &queue.IntQueue{},
	}
}

func (d *deckOfCards) dealCard() int {
	nextCard, _ := d.q.Dequeue()
	return nextCard
}

func (d *deckOfCards) empty() bool {
	return d.q.IsEmpty()
}

func (d *deckOfCards) addCard(card int) {
	d.q.Enqueue(card)
}

func (d *deckOfCards) duplicate(numberOfCards int) *deckOfCards {
	duplicateQueue, _ := d.q.Duplicate(numberOfCards)
	return &deckOfCards{q: duplicateQueue}
}

func (d *deckOfCards) string() string {
	return d.q.String()
}

func (d *deckOfCards) length() int {
	return d.q.Length()
}

func getDeckOfCards(cards []string) (*deckOfCards, error) {
	deck := new()
	for _, card := range cards {
		cardNumber, err := strconv.Atoi(card)
		if err != nil {
			return nil, err
		}
		deck.addCard(cardNumber)
	}
	return deck, nil
}

func playCombat(player1, player2 *deckOfCards) *deckOfCards {
	var winner *deckOfCards
	for winner == nil {
		if player1.empty() {
			winner = player2
			break
		}
		if player2.empty() {
			winner = player1
			break
		}

		player1sCard := player1.dealCard()
		player2sCard := player2.dealCard()

		if player1sCard > player2sCard {
			player1.addCard(player1sCard)
			player1.addCard(player2sCard)
		} else {
			player2.addCard(player2sCard)
			player2.addCard(player1sCard)
		}
	}
	return winner
}

func determineWinnersScore(player *deckOfCards) int {
	stack := *&stack.IntStack{}
	for {
		if player.empty() {
			break
		}
		card := player.dealCard()
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

func playRecursiveCombat(player1 *deckOfCards, player2 *deckOfCards) (*deckOfCards, int) {
	previousGameStates := make(map[string]bool)
	for {
		gameState := getRoundState(player1, player2)
		if _, seen := previousGameStates[gameState]; seen {
			return player1, 1
		}
		previousGameStates[gameState] = true

		if player1.empty() {
			return player2, 2
		}
		if player2.empty() {
			return player1, 1
		}

		player1sCard := player1.dealCard()
		player2sCard := player2.dealCard()

		if player1sCard <= player1.length() && player2sCard <= player2.length() {
			player1Copy := player1.duplicate(player1sCard)
			player2Copy := player2.duplicate(player2sCard)

			_, winnerID := playRecursiveCombat(player1Copy, player2Copy)

			if winnerID == 1 {
				player1.addCard(player1sCard)
				player1.addCard(player2sCard)
			} else {
				player2.addCard(player2sCard)
				player2.addCard(player1sCard)
			}
			continue
		}

		if player1sCard > player2sCard {
			player1.addCard(player1sCard)
			player1.addCard(player2sCard)
		} else {
			player2.addCard(player2sCard)
			player2.addCard(player1sCard)
		}
	}
}

func getRoundState(player1 *deckOfCards, player2 *deckOfCards) string {
	var sb strings.Builder
	sb.WriteString(player1.string())
	sb.WriteString(" ")
	sb.WriteString(player2.string())
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
