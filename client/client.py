import sys
import logging
import argparse

import grpc
import proto.guessing_game_pb2     as pb2_game
import proto.guessing_game_pb2_grpc as pb2_grpc


def main(input_args):
    parser = argparse.ArgumentParser()
    parser.add_argument("port")
    parser.add_argument("guess")

    args = parser.parse_args(input_args)

    addr = "localhost:{}".format(args.port)
    print("Connecting")
    with grpc.insecure_channel(addr) as channel:
        stub = pb2_grpc.GuessingGameStub(channel)

        guess = int(args.guess)
        response = stub.MakeGuess(pb2_game.Guess(number=guess))
        print("Received  {}".format(pb2_game.GuessStatus.Name(response.result)))


if __name__ == "__main__":
    main(sys.argv[1:])


