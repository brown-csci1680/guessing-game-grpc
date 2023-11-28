package main

import (
	"context"
	"fmt"
	"guessing-game/pkg/game"
	pb "guessing-game/pkg/proto"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

// NEW:  a struct to represent our server state
type Server struct {
	GameState *game.GameInfo

	pb.UnimplementedGuessingGameServer
}

// NEW:  grpc calls this each time a client makes a guess
// in is the deserialized message (struct pb.Guess)
func (s *Server) MakeGuess(ctx context.Context, in *pb.Guess) (*pb.GuessResult, error) {
	number := in.GetNumber()
	log.Printf("Received guess:  %v", number)

	// Check the input
	result := s.GameState.DoGuess(number)

	// Send a result in protobuf format
	var resp pb.GuessStatus
	switch result {
	case game.GuessTooHigh:
		resp = pb.GuessStatus_GUESS_TOO_HIGH
	case game.GuessTooLow:
		resp = pb.GuessStatus_GUESS_TOO_LOW
	case game.GuessCorrect:
		resp = pb.GuessStatus_GUESS_CORRECT
	}

	return &pb.GuessResult{Result: resp}, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s <port number>", os.Args[0])
	}
	//log.Default().SetOutput(io.Discard) //Equivalent of writing logs to /dev/null

	portNumber := os.Args[1]

	// Initialize the guessing game
	rand.Seed(time.Now().Unix())
	server := Server{
		GameState: game.InitializeGame(),
	}
	fmt.Println("Target number is:  ", server.GameState.TargetNumber)

	// Same as before:  sreate the socket
	addr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf(":%s", portNumber))
	if err != nil {
		log.Fatalln("Error translating address:  ", err)
	}

	conn, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// NEW:  Start grpc and serve our game server
	s := grpc.NewServer()
	pb.RegisterGuessingGameServer(s, &server)
	if err := s.Serve(conn); err != nil {
		log.Fatal("Failed to serve", err)
	}

}
