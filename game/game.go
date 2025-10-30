package game

type Game struct {
	Board         [][]int
	CurrentPlayer int
	Winner        int
}

// Fonction appelée par défaut
func NewGame() *Game {
	return NewGameEasy()
}

// ---- EASY ----
func NewGameEasy() *Game {
	return newCustomGame(6, 7)
}

// ---- NORMAL ----
func NewGameNormal() *Game {
	return newCustomGame(6, 9)
}

// ---- HARD ----
func NewGameHard() *Game {
	return newCustomGame(7, 8)
}

// Fonction commune pour générer la grille
func newCustomGame(rows, cols int) *Game {
	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, cols)
	}
	return &Game{
		Board:         board,
		CurrentPlayer: 1,
		Winner:        0,
	}
}

func (g *Game) PlayMove(col int) bool {
	if col < 0 || col >= len(g.Board[0]) || g.Winner != 0 {
		return false
	}
	for row := len(g.Board) - 1; row >= 0; row-- {
		if g.Board[row][col] == 0 {
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
	if g.CurrentPlayer == 1 {
		g.CurrentPlayer = 2
	} else {
		g.CurrentPlayer = 1
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
