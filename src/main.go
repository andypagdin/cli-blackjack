package main

import (
	"fmt"
	"math/rand"
	"time"
)

var deck = [...]string{
	"Ah", "2h", "3h", "4h", "5h", "6h", "7h", "8h", "9h", "10h", "Jh", "Qh", "Kh",
	"As", "2s", "3s", "4s", "5s", "6s", "7s", "8s", "9s", "10s", "Js", "Qs", "Ks",
	"Ad", "2d", "3d", "4d", "5d", "6d", "7d", "8d", "9d", "10d", "Jd", "Qd", "Kd",
	"Ac", "2c", "3c", "4c", "5c", "6c", "7c", "8c", "9c", "10c", "Jc", "Qc", "Kc"}

var value = map[string]int{
	"Ah": 11, "2h": 2, "3h": 3, "4h": 4, "5h": 5, "6h": 6, "7h": 7, "8h": 8, "9h": 9, "10h": 10, "Jh": 10, "Qh": 10, "Kh": 10,
	"As": 11, "2s": 2, "3s": 3, "4s": 4, "5s": 5, "6s": 6, "7s": 7, "8s": 8, "9s": 9, "10s": 10, "Js": 10, "Qs": 10, "Ks": 10,
	"Ad": 11, "2d": 2, "3d": 3, "4d": 4, "5d": 5, "6d": 6, "7d": 7, "8d": 8, "9d": 9, "10d": 10, "Jd": 10, "Qd": 10, "Kd": 10,
	"Ac": 11, "2c": 2, "3c": 3, "4c": 4, "5c": 5, "6c": 6, "7c": 7, "8c": 8, "9c": 9, "10c": 10, "Jc": 10, "Qc": 10, "Kc": 10}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	balance := 100.00
	newHand(balance)
}

func newHand(balance float64) {
	if balance <= 0 {
		fmt.Println("You've gone bust, gutted.")
		return
	}

	fmt.Printf("Balance £%0.2f\n", balance)

	bet := promptForBet(balance)
	fmt.Println("--------------------")

	dealer, dtotal := dealHand(2)
	player, ptotal := dealHand(2)

	fmt.Printf("Dealer shows [%s]\n", dealer[0])
	playerResult := playPlayerHand(player, ptotal)

	if playerResult == 0 {
		newHand(balance - bet)
	}

	if playerResult == 21 {
		fmt.Println("21! good for you")
		newHand(balance + (((bet * 3) / 2) + bet))
	}

	dealerResult := playDealerHand(dealer, dtotal, playerResult)

	switch dealerResult {
	case 0:
		newHand(balance + ((bet / 2) * 2))
	case 1:
		newHand(balance - bet)
	default:
		newHand(balance)
	}
}

// 0 = dealer lose
// 1 = dealer win
// 3 = push
func playDealerHand(hand []string, total int, playerResult int) int {
	var result int

	for {
		fmt.Printf("Dealer shows %v - total %d\n", hand, total)
		if total > 21 {
			fmt.Printf("Dealer bust! %d\n", total)
			result = 0
			break
		}

		if total >= 17 {
			if total > playerResult {
				fmt.Printf("Dealer wins! %d\n", total)
				result = 1
			} else if total == playerResult {
				fmt.Println("Push! bets returned")
				result = 3
			} else {
				fmt.Printf("Player wins! %d against %d\n", playerResult, total)
				result = 0
			}
			break
		}

		if total < 17 {
			card := deck[rand.Intn(len(deck))]
			hand = append(hand, card)
			// todo: count ace as 1 when hitting into
			total = total + value[card]
		}
	}

	return result
}

// 0 = bust
func playPlayerHand(hand []string, total int) int {
	var choice string

	for {
		fmt.Printf("You show %v - total %d\n", hand, total)
		fmt.Print("Hit (h) or Stand (s) ")
		_, err := fmt.Scanf("%s\n", &choice)

		if err != nil {
			fmt.Println(err)
		}

		if choice == "h" {
			// todo: if the card is already in the players hand redraw another
			card := deck[rand.Intn(len(deck))]
			// todo: count ace as 1 when hitting into
			value := value[card]
			newTotal := total + value

			if newTotal > 21 {
				fmt.Printf("Drew %s %v Total %d - Bust!\n", card, hand, newTotal)
				total = 0
				break
			}

			hand = append(hand, card)
			total = newTotal
		} else if choice == "s" {
			break
		}
	}

	return total
}

func dealHand(count int) ([]string, int) {
	var hand []string
	var total int

	for i := 0; i < count; i++ {
		card := deck[rand.Intn(len(deck))]
		hand = append(hand, card)
		total += value[card]
	}

	return hand, total
}

func promptForBet(balance float64) float64 {
	var bet float64

	fmt.Print("Enter bet amount £")
	_, scnerr := fmt.Scanf("%f\n", &bet)

	// todo: explore better solutions to handling invalid input
	if scnerr != nil || bet > balance || bet <= 0 {
		fmt.Println("Invalid bet")
		return promptForBet(balance)
	}

	return bet
}
