package pkg

import (
	"fmt"
	"strings"
)

type Player struct {
	Location  string   // Текущая локация игрока
	Inventory []string // Инвентарь игрока
}

type Location struct {
	Name        string   // Название локации
	Description string   // Описание локации
	Items       []string // Предметы в локации
}

type Command struct {
	Name      string   // Название команды
	Arguments []string // Аргументы команды
}

type GameWorld struct {
	Locations map[string]Location // Карта локаций
	Player    Player              // Игрок
}

func InitGame() GameWorld {
	locations := map[string]Location{
		"кухня": {
			Name:        "кухня",
			Description: "Вы находитесь на кухне",
			Items:       []string{"ключи"},
		},
		"коридор": {
			Name:        "коридор",
			Description: "Вы находитесь в коридоре",
			Items:       []string{"рюкзак"},
		},
		"комната": {
			Name:        "комната",
			Description: "Вы находитесь в комнате",
			Items:       []string{"конспекты"},
		},
		"улица": {
			Name:        "улица",
			Description: "Вы находитесь на улице",
			Items:       []string{},
		},
	}

	player := Player{
		Location:  "коридор",
		Inventory: []string{},
	}

	world := GameWorld{
		Locations: locations,
		Player:    player,
	}

	return world
}

func HandleCommand(world GameWorld, command Command) {
	switch command.Name {
	
	case "осмотреться":
		location := world.Locations[world.Player.Location]
		fmt.Println(location.Description)
		fmt.Println("Предметы в локации:", strings.Join(location.Items, ", "))
		fmt.Println("Ваш инвентарь:", strings.Join(world.Player.Inventory, ", "))
	
	case "идти":
		if len(command.Arguments) > 0 {
			locationName := command.Arguments[0]
			location, ok := world.Locations[locationName]
			if ok {
				world.Player.Location = location.Name
				fmt.Println("Вы перешли в локацию", location.Name)
				fmt.Println(location.Description)
			} else {
				fmt.Println("Локация", locationName, "не найдена")
			}
		} else {
			fmt.Println("Не указано имя локации")
		}
	
	case "взять":
		if len(command.Arguments) > 0 {
			itemName := command.Arguments[0]
			location := world.Locations[world.Player.Location]
			for i, item := range location.Items {
				if item == itemName {
					world.Player.Inventory = append(world.Player.Inventory, item)
					location.Items = append(location.Items[:i], location.Items[i+1:]...)
					fmt.Println("Вы взяли предмет", itemName)
					return
				}
			}
			fmt.Println("Предмет", itemName, "не найден в локации")
		} else {
			fmt.Println("Не указано имя предмета")
		}
	default:
		fmt.Println("Неизвестная команда")
	}
}
