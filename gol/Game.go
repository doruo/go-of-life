package gol

type Game struct {
	title                     string
	screenWidth, screenHeight int
	gs                        GameState
}

func NewGame(width, height, lag int) *Game {
	return &Game{
		title:        "Game of life",
		screenWidth:  width,
		screenHeight: height,
		gs:           *NewGameState(width * height, lag),
	}
}

func (g *Game) Update()  {
	g.gs.Update()
}

func (g *Game) Run() {
	for {
		g.Update()
	}
}

/*
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.title)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
*/