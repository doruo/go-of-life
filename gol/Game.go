package gol

type Game struct {
	size int
	gs   GameState
}

func NewGame(size, lag int) *Game {
	return &Game{
		size: size,
		gs:   *NewGameState(size, lag),
	}
}

func (g *Game) init() {
	g.gs.Init()
}

func (g *Game) update() {
	g.gs.Update()
}

func (g *Game) show() {
	g.gs.Show()
}

func (g *Game) Run() {
	g.init()
	for {
		g.update()
		g.show()
	}
}
