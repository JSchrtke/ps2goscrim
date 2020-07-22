package main

import (
	"errors"
	"fmt"
)

type Ruleset interface {
	TeamCount() int
	PlayerCount() int
}

type Match struct {
	rules   Ruleset
	teams   []Team
	players []Player
}

type Team struct {
	id    int
	Name  string
	Score int
}

type Player struct {
	id     int
	Name   string
	Score  int
	Kills  int
	Deaths int
}

type UserInterface interface {
	GetTeams() []Team
	GetPlayers() []Player
}

func NewMatch(r Ruleset, ui UserInterface) (error, Match) {
	match := Match{
		rules:   r,
		teams:   make([]Team, r.TeamCount()),
		players: make([]Player, r.PlayerCount()),
	}

	t := ui.GetTeams()
	if len(t) != r.TeamCount() {
		return errors.New(fmt.Sprintf(
			"Invalid team count: expected %d but got %d",
			r.TeamCount(), len(t),
		)), Match{}
	}
	for i, team := range t {
		match.teams[i] = team
	}

	p := ui.GetPlayers()
	if len(p) != r.PlayerCount() {
		return errors.New(fmt.Sprintf(
			"Invalid player count: expected %d but got %d",
			r.PlayerCount(), len(p),
		)), Match{}
	}
	for i, player := range p {
		match.players[i] = player
	}

	return nil, match
}

func (m *Match) main() {
}

func (m *Match) Start() error {
	m.main()
	return nil
}

type MatchInfo struct {
	Teams   []Team
	Players []Player
}

func (m *Match) CurrentState() MatchInfo {
	return MatchInfo{
		Teams:   m.teams,
		Players: m.players,
	}
}
