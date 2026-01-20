package gol

type Game struct {
	title                     			string
	gridSize, screenWidth, screenHeight	int
	gs                        			GameState
}

func NewGame(screenWidth, screenHeigth, size, lag int) *Game {
	return &Game{
		title:         "Game of Life",
		screenWidth:   screenWidth,
		screenHeight:  screenHeigth,
		gridSize: 	   size,
		gs:            *NewGameState(size, lag),
	}
}

func (g *Game) Init() {
	g.gs.Init()
}

func (g *Game) Update()  {
	g.gs.Update()
}

func (g *Game) Show()  {
	g.gs.Show()
}

func (g *Game) Run() {
	g.Init()
	for {
		g.Update()
		g.Show()
	}
}
