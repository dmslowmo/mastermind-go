package main

// /root page -- welcome, show link to go to /mastermind
// /mastermind -- welcome, mastermind home page, instruction
// /mastermind/play -- start playing, sessionId, secret code generated, board init
// /mastermind/play/{sessionId}?guess={guess} -- send guess to the server to crack the code
// /mastermind/endgame -- game ended, link to /play again, or quit
// /mastermind/

// tables
// 1. secret code table, cols:
//		- client_id	     - string, pk
//		- session_id     - string (client_id + '_' + timestamp)
//		- secret         - string (4 digits)
//		- session_active - boolean
// 2. guess table, pk = sessionId_attempt
//		- session_id_attempt - string, pk
//		- session_id         - string
//		- attempt_no         - int
//		- guess		         - string
//		- hint		         - string

func main() {
}
