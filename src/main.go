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
	for {
		if balance <= 0 {
			fmt.Println("You've gone bust, gutted.")
			break
		}

		fmt.Printf("Balance £%0.2f\n", balance)
		bet := promptForBet(balance)

		dealer, dtotal := dealHand(2)
		player, ptotal := dealHand(2)

		fmt.Printf("Dealer shows [%s]\n", dealer[0])
		playerResult, isBust := playPlayerHand(player, ptotal)

		if isBust {
			balance -= bet
			continue
		}

		if playerResult == 21 {
			fmt.Println("Player wins with 21!")
			balance += (((bet * 3) / 2) + bet)
			continue
		}

		dealerResult := playDealerHand(dealer, dtotal, playerResult)

		switch dealerResult {
		case "player":
			balance += ((bet / 2) * 2)
		case "dealer":
			balance -= bet
		}
	}
}

func playDealerHand(hand []string, total int, playerResult int) string {
	var winner string
	var value int

	for {
		fmt.Printf("Dealer shows %v - total %d\n", hand, total)
		if total > 21 {
			fmt.Printf("Dealer bust! %d\n", total)
			winner = "player"
			break
		}

		if total >= 17 {
			if total > playerResult {
				fmt.Printf("Dealer wins! %d\n", total)
				winner = "dealer"
			} else if total == playerResult {
				fmt.Println("Push! bets returned")
			} else {
				fmt.Printf("Player wins! %d against %d\n", playerResult, total)
				winner = "player"
			}
			break
		}

		hand, _, value = dealCard(hand)
		total += value
	}

	return winner
}

func playPlayerHand(hand []string, total int) (int, bool) {
	var card, choice string
	var value int

	for {
		fmt.Printf("You show %v - total %d\n", hand, total)
		fmt.Print("Hit or Stand (h/s) ")
		_, err := fmt.Scanf("%s\n", &choice)

		if err != nil {
			fmt.Println(err)
		}

		if choice == "h" {
			hand, card, value = dealCard(hand)
			total += value

			if total > 21 {
				fmt.Printf("Drew %s %v Total %d - Bust!\n", card, hand, total)
				return 0, true
			}
		} else if choice == "s" {
			break
		}
	}

	return total, false
}

func dealHand(count int) ([]string, int) {
	var hand []string
	var total, value int

	for i := 0; i < count; i++ {
		hand, _, value = dealCard(hand)
		total += value
	}

	return hand, total
}

func dealCard(hand []string) ([]string, string, int) {
	// todo: if the card is already in the players hand redraw another
	// todo: count ace as 1 when hitting into
	card := deck[rand.Intn(len(deck))]
	hand = append(hand, card)
	value := value[card]
	return hand, card, value
}

func promptForBet(balance float64) float64 {
	var bet float64

	for {
		fmt.Print("Enter bet amount £")
		_, scnerr := fmt.Scanf("%f\n", &bet)

		if scnerr != nil || bet > balance || bet <= 0 {
			fmt.Println("Invalid bet")
		} else {
			break
		}
	}

	return bet
}
