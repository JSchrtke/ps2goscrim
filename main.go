package main

import (
	"fmt"
	"log"
	"time"
)

type MockRules struct {
}

func (r MockRules) TeamCount() int {
	return 2
}

func (r MockRules) PlayerCount() int {
	return 12
}

type MockUi struct {
}

func (u MockUi) GetTeams() []Team {
	teams := make([]Team, 2)
	for i := 0; i < len(teams); i++ {
		teams[i] = Team{
			id:   i,
			Name: fmt.Sprintf("Template team %d", i),
		}
	}
	return teams
}

func (u MockUi) GetPlayers() []Player {
	players := make([]Player, 12)
	for i := 0; i < len(players); i++ {
		players[i] = Player{
			id:   i,
			Name: fmt.Sprintf("player%d", i),
		}
	}
	return players
}

func main() {
	rules := MockRules{}
	ui := MockUi{}

	err, match := NewMatch(rules, ui)
	if err != nil {
		log.Printf("Error when creating match: %s", err)
	}

	err = match.Start()
	if err != nil {
		log.Printf("Error when starting match: %s", err)
	}

	for {
		log.Println("Update:")
		info := match.CurrentState()
		for _, team := range info.Teams {
			log.Printf("%s:\n\tscore: %d\n\n", team.Name, team.Score)
		}
		for _, player := range info.Players {
			log.Printf(
				"%s:\n\tscore: %d\n\tkills: %d\n\tdeaths: %d\n\n",
				player.Name, player.Score, player.Kills, player.Deaths,
			)
		}

		time.Sleep(1 * time.Second)
	}
}
