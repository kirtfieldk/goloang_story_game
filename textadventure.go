package main

// Create NPC
// NPC Move around map
// items can be picked up and place
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type weapon struct {
	name       string
	weaponType string
	powerLevel int32
}
type choices struct {
	cmd        string
	desc       string
	nextNode   *storyNode
	nextChoice *choices
	weapon     *weapon
}
type storyNode struct {
	text    string
	choices *choices
	weapon  *weapon
}

func (node *storyNode) addChoice(cmd string, desc string, nextNode *storyNode, weapon *weapon) {
	choice := &choices{cmd, desc, nextNode, nil, weapon}
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
func (node *storyNode) addWepon(indiWeapon *weapon) {
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
func (node *storyNode) executeCmd(cmd string) *storyNode {
	currentChoice := node.choices
	for currentChoice != nil {
		if strings.ToLower(currentChoice.cmd) == strings.ToLower(cmd) {
			node.addWepon(currentChoice.weapon)
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	fmt.Println("Unreadable response")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
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
	weapon1 := weapon{"OmeBreaker", "sword", 3200}
	weapon2 := weapon{"World Destroyer", "axe", 48000}
	// weapon3 := weapon{"friend maker", "bow", 3}

	fmt.Println("Started")
	start := storyNode{text: `
	Deep in the abyss you wake. Its dark and forgin. Slowly you stand up, shaking, there is only one way to go and that is deeper into the void`}
	stay := storyNode{text: "Scared and afraid you stay put, paralized by the fear of the unknown you just stare into the abyss, and you wait"}
	walkIntoAbyss := storyNode{text: `Slowly you walk into the darkness. The further you go, a thic wave a cold encompasses your body. Decide left or right `}

	start.addChoice("Stay", "Stay put", &stay, &weapon1)
	start.addChoice("Go", "Walking forward", &walkIntoAbyss, &weapon2)

	start.play()
	fmt.Println()
	fmt.Println(start.weapon.name)
	fmt.Println("The End")
}
