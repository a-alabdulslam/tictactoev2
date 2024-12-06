package game

import "math/rand"

type Game struct {
	id    string
	state *GameState
	room  *Room
}

type GameState struct {
	board       [][]string
	currentTurn *Client
	winner      *Client
	isDraw      bool
}

func NewGame(room *Room) *Game {
	firstPlayer := room.players[rand.Int()%len(room.players)]

	return &Game{
		id: generateRoomID(),
		state: &GameState{
			board: [][]string{
				{"", "", ""},
				{"", "", ""},
				{"", "", ""},
			},
			currentTurn: firstPlayer,
			winner:      nil,
			isDraw:      false,
		},
	}
}

func (g Game) checkwinner() {
	// Check rows
	for i := 0; i < 3; i++ {
		if g.state.board[i][0] == g.state.board[i][1] && g.state.board[i][1] == g.state.board[i][2] && g.state.board[i][0] != "" {
			g.state.winner = g.state.currentTurn
			return
		}
	}
	// Check columns
	for i := 0; i < 3; i++ {
		if g.state.board[0][i] == g.state.board[1][i] && g.state.board[1][i] == g.state.board[2][i] && g.state.board[0][i] != "" {
			g.state.winner = g.state.currentTurn
			return
		}
	}
	// Check diagonals
	if g.state.board[0][0] == g.state.board[1][1] && g.state.board[1][1] == g.state.board[2][2] && g.state.board[0][0] != "" {
		g.state.winner = g.state.currentTurn
		return
	}
	if g.state.board[0][2] == g.state.board[1][1] && g.state.board[1][1] == g.state.board[2][0] && g.state.board[0][2] != "" {
		g.state.winner = g.state.currentTurn
		return
	}
	// Check draw
	isDraw := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.state.board[i][j] == "" {
				isDraw = false
				break
			}
		}
	}
	g.state.isDraw = isDraw
}
