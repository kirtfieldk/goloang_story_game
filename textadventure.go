package main

// Create NPC
// NPC Move around map
// Randomize enemy encounter
// wepons can be picked up
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type monster struct {
	name       string
	powerLevel int64
	element    string
}
type weapon struct {
	name       string
	weaponType string
	powerLevel int32
	element    string
}
type choices struct {
	cmd        string
	desc       string
	nextNode   *storyNode
	nextChoice *choices
	weapon     *weapon
	monster    *monster
}
type storyNode struct {
	text    string
	choices *choices
}
type player struct {
	weapon *weapon
	name   string
}

func (node *storyNode) addChoice(cmd string, desc string, nextNode *storyNode, weapon *weapon, monster *monster) {
	choice := &choices{cmd, desc, nextNode, nil, weapon, monster}
	fmt.Println(weapon)
	if node.choices == nil {
		node.choices = choice
	} else {
		currentChoice := node.choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}
func (node *player) addWepon(indiWeapon *weapon) {
	if node.weapon == nil {
		node.weapon = indiWeapon
	} else {
		if node.weapon.powerLevel < indiWeapon.powerLevel {
			node.weapon = indiWeapon
		} else {
			fmt.Println("Weapon is weak")
		}
	}
}
func (node *storyNode) executeCmd(cmd string, player *player) *storyNode {
	currentChoice := node.choices
	for currentChoice != nil {
		if strings.ToLower(currentChoice.cmd) == strings.ToLower(cmd) {
			// Once user makes a command the wepon is picked up
			player.addWepon(currentChoice.weapon)
			if getRan() {
				fmt.Printf("\nENEMY ENCOUNTER WITH: %v\n", currentChoice.monster.name)
			}
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	fmt.Println("Unreadable response")
	return node
}
func getRan() bool {
	num := rand.Intn(10)
	fmt.Println(num)
	if num > 1 {
		return true
	}
	return false
}

var scanner *bufio.Scanner

func (node *storyNode) play(player *player) {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text(), player).play(player)
	}
}
func (node *storyNode) render() {
	fmt.Println(node.text)
	currentChoice := node.choices
	for currentChoice != nil {
		fmt.Println(currentChoice.cmd, ":", currentChoice.desc)
		currentChoice = currentChoice.nextChoice
	}
}
func main() {
	scanner = bufio.NewScanner(os.Stdin)
	player := player{name: "Keith Kirtfield"}
	// Wepons
	weapon1 := weapon{"OmeBreaker", "sword", 3200, "fire"}
	weapon2 := weapon{"World Destroyer", "axe", 48000, "water"}
	weapon3 := weapon{"friend maker", "bow", 3, "air"}
	//Monsters
	odin := monster{"Odin", 22000, "Earth"}
	alexander := monster{"Alexander", 18000, "Air"}
	// shiva := monster{"Shiva", 32000, "Water"}
	// bahamat := monster{"Bahamut", 15000, "Fire"}
	fmt.Println("Started")
	start := storyNode{text: `
	Deep in the abyss you wake. Its dark and forgin. Slowly you stand up, shaking, there is only one way to go and that is deeper into the void`}
	stay := storyNode{text: "Scared and afraid you stay put, paralized by the fear of the unknown you just stare into the abyss, and you wait. Should you stay or leave"}
	walkIntoAbyss := storyNode{text: `Slowly you walk into the darkness. The further you go, a thic wave a cold encompasses your body. Decide left or right `}
	right := storyNode{text: `You choose to go right and you suddenly wake up`}
	start.addChoice("Stay", "Stay put", &stay, &weapon1, &odin)
	start.addChoice("Go", "Walking forward", &walkIntoAbyss, &weapon2, &odin)
	walkIntoAbyss.addChoice("right", "Walking roght", &right, &weapon3, &alexander)

	start.play(&player)
	fmt.Println()
	fmt.Println(player.weapon.name)
	fmt.Println("The End")
}
