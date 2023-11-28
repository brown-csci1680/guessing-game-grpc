package game

import (
	"math/rand"
	"net"
	"sync"
)

type ClientInfo struct {
	Id              int
	Conn            net.Conn
	GameResetChan   chan bool
	ServerCloseChan chan bool
}

type GameInfo struct {
	GameLock     sync.Mutex
	TotalGuesses int
	TargetNumber int32

	// TODO:  The current version doesn't track clients currently playing the game
	//nextClientIdx int // Counter to increment each time we add a new client
}

const (
	GuessTooHigh = 1
	GuessCorrect = 0
	GuessTooLow  = -1
)

func InitializeGame() *GameInfo {
	return &GameInfo{
		// Other fields initialized to zero
		TargetNumber: rand.Int31n(8192),
	}
}

func (g *GameInfo) ResetGame() {
	g.GameLock.Lock()
	defer g.GameLock.Unlock()

	g.TargetNumber = rand.Int31()
	g.TotalGuesses = 0
}

func (g *GameInfo) DoGuess(n int32) int32 {
	g.GameLock.Lock()
	defer g.GameLock.Unlock()

	g.TotalGuesses++

	if n < g.TargetNumber {
		return GuessTooLow
	} else if n > g.TargetNumber {
		return GuessTooHigh
	} else {
		return GuessCorrect
	}
}
