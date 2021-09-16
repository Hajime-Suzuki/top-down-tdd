package userinput

type UserInput struct {
	x int
	y int
}

func NewUserInput(x int, y int) UserInput {
	return UserInput{x, y}
}
