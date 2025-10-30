package game

import (
	"math/rand"
	"time"
)

const (
	Empty   = 0
	Player1 = 1
	Player2 = 2
	Blocked = 3 // case pré-remplie / obstacle (ne peut pas être remplacée)
)

type Game struct {
	Board         [][]int // 0 = vide, 1 = joueur1, 2 = joueur2
	CurrentPlayer int
	Winner        int
}

func NewGame() *Game {
	board := make([][]int, 6)
	for i := range board {
		board[i] = make([]int, 7)
	}
	g := &Game{
		Board:         board,
		CurrentPlayer: Player1,
		Winner:        0,
	}
	rand.Seed(time.Now().UnixNano())
	g.placeRandomBlocks(3)
	return g
}

func (g *Game) PlayMove(col int) bool {
	if col < 0 || col >= len(g.Board[0]) || g.Winner != 0 {
		return false
	}
	for row := len(g.Board) - 1; row >= 0; row-- {
		if g.Board[row][col] == Empty {
			g.Board[row][col] = g.CurrentPlayer
			if g.checkWin(row, col) {
				g.Winner = g.CurrentPlayer
			}
			g.switchPlayer()
			return true
		}
	}
	return false
}

func (g *Game) switchPlayer() {
	if g.CurrentPlayer == Player1 {
		g.CurrentPlayer = Player2
	} else {
		g.CurrentPlayer = Player1
	}
}
func (g *Game) checkWin(r, c int) bool {
	player := g.Board[r][c]
	directions := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}
	for _, d := range directions {
		count := 1
		count += g.countDirection(r, c, d[0], d[1], player)
		count += g.countDirection(r, c, -d[0], -d[1], player)
		if count >= 4 {
			return true
		}
	}
	return false
}

func (g *Game) countDirection(r, c, dr, dc, player int) int {
	count := 0
	for {
		r += dr
		c += dc
		if r < 0 || r >= len(g.Board) || c < 0 || c >= len(g.Board[0]) {
			break
		}
		if g.Board[r][c] != player {
			break
		}
		count++
	}
	return count
}
func (g *Game) placeRandomBlocks(n int) {
	placed := 0
	rows := len(g.Board)
	cols := len(g.Board[0])
	for placed < n {
		r := rand.Intn(rows-1) + 1 // de 1..rows-1 (évite la ligne 0)
		c := rand.Intn(cols)
		if g.Board[r][c] == Empty {
			g.Board[r][c] = Blocked
			placed++
		}
	}
}
func (g *Game) IsAIsTurn() bool {
	return g.CurrentPlayer == Player2 && g.Winner == 0
}
func (g *Game) AIPlay() {
	valid := g.validMoves()
	if len(valid) == 0 {
		return
	}
	col := valid[rand.Intn(len(valid))]
	_ = g.PlayMove(col)
}
func (g *Game) validMoves() []int {
	res := []int{}
	cols := len(g.Board[0])
	for c := 0; c < cols; c++ {
		if g.Board[0][c] == Empty {
			res = append(res, c)
		}
	}
	return res
}
