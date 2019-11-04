package structs

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
}
type storyNode struct {
	text    string
	choices *choices
	weapon  *weapon
}
