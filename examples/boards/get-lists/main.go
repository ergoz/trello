package main

import (
	"flag"
	"fmt"

	"github.com/ttacon/pretty"
	"github.com/ttacon/trello"
)

var (
	key     = flag.String("k", "", "application key")
	token   = flag.String("t", "", "user authentication token")
	boardID = flag.String("b", "", "board to retieve")
)

func main() {
	flag.Parse()

	client := trello.NewClient(*key, *token)

	board, err := client.BoardService().GetBoard(*boardID)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	lists, err := board.Lists()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	pretty.Println(lists)
}
