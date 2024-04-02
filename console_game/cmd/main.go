package main

import "github.com/TursunovImran/mini-projects_go/console_game/pkg"

func main() {
	world := initGame()

	command1 := pkg.Command{Name: "осмотреться"}
	handleCommand(world, command1)

	command2 := pkg.Command{Name: "идти", Arguments: []string{"кухня"}}
	handleCommand(world, command2)

	command3 := pkg.Command{Name: "взять", Arguments: []string{"ключи"}}
	handleCommand(world, command3)

	command4 := pkg.Command{Name: "идти", Arguments: []string{"комната"}}
	handleCommand(world, command4)

	command5 := pkg.Command{Name: "взять", Arguments: []string{"конспекты"}}
	handleCommand(world, command5)

	command6 := pkg.Command{Name: "осмотреться"}
	handleCommand(world, command6)

	command7 := pkg.Command{Name: "взять", Arguments: []string{"рюкзак"}}
	handleCommand(world, command7)
}

func initGame() pkg.GameWorld {
	return pkg.InitGame()
}

func handleCommand(world pkg.GameWorld, command pkg.Command) {
	pkg.HandleCommand(world, command)
}
