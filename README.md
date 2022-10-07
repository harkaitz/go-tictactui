# TICTACTUI

The TicTacToe game in Go. I wrote this to learn the Go language.

- The original here [c-tictactui](https://github.com/harkaitz/c-tictactui).
- [tictactoe.go](./tictactoe.go) : The rules of the TicTacToe game.
- [board.go](./board.go) : Generic 8x8 or 3x3 board, can be used for chess too.
- [minimax.go](./minimax.go) : Generic Minimax algorithm.
- [cmd/tictactoe/main.go](./cmd/tictactoe/main.go) : CLI to the TicTacToe game that follows the UNIX philosophy.
- [bin/tictactui](bin/tictactui) : Shell script TUI to the TicTacToe game. Uses the CLI.

Things missing:

- Chess, checkers.
- Web interface.
- Multiplayer, maybe with Redis.
