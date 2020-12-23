package main


import(
	"fmt"
)

func print(board [9]int) {					// Used to print the board 
		
	fmt.Println("\n")

	 for count := 0; count < 9; count++ {			// Iterate through the loop
                    if (count % 3 == 0) && (count != 0){
			if (board[count] == 0 ) {		// If the data in that particular index is 0, then place an x (player 1)
				fmt.Printf("\n X ")
			} else if (board[count] == 10) {	// If the data in that particular index is 10, then place an O (player 2)
				fmt.Printf("\n O ")
			} else { 
                            fmt.Printf("\n %v ", board[count])	// else print out the remaining moves 
			}
                    }else{
			if (board[count] == 0 ) {		// If the data in that particular index is 0, then place an x (player 1)
                                fmt.Printf(" X ")
	
                        } else if (board[count] == 10) {	// If the data in that particular index is 10, then place a O (player 2)
				fmt.Printf(" O " )
			} else {
  
                            fmt.Printf(" %v ",board[count])	// Print out the remaining moves
                    	}
		     }
	}
	
	fmt.Println("\n")
	return	
}

func moves(player int, board [9]int)int {			// Function that allows players to input the moves that they like to make
	
	duplicateMoves := true
	for duplicateMoves == true {				// It will run to see if the user enters a move that is already been taken

		fmt.Println("\nPlease enter a move ")		
		fmt.Scan(&player)				// Entering a number
		
		if board[player - 1] == player {		// Check to see if the moves is already taken
			duplicateMoves = false
		} else {					// else, the move has already been taken 
			continue
		}
	}
	return player
}
func winner(board [9] int) bool {				// Uses short-hand arithmatic to detect a winner, since player 1 == 0
	win := [9]int { 1, 1, 1, 1, 1, 1, 1, 1 }		// and player 2 == 10
								// if one of these index == 0 or 30, then we have a winner
		// the sum of the moves
		
		// horizontal "lines " 
	win[0] = board[0] + board[1] + board[2]
	win[1] = board[3] + board[4] + board[5]
	win[2] = board[6] + board[7] + board[8]

		// diagonal "lines "
	win[3] = board[0] + board[4] + board[8]
	win[4] = board[6] + board[4] + board[2]

		// Vertical "lines" 
	win[5] = board[0] + board[3] + board[6]
	win[6] = board[1] + board[4] + board[7]
	win[7] = board[2] + board[5] + board[8]
	
		// check for a winner
	for count := 0; count < 8; count++ {			// A loop to see if one of the player had already won 
		if (win[count] == 0) {
			return true
		} else if (win[count] == 30) {
			return true
		}
	}
	return false 
}

// **************************** The main **************************
func main() {
	board := [9]int{ 1, 2, 3, 4, 5, 6, 7, 8, 9}		// initial board 
	won := [9]int{ 0, 0, 0, 0, 0, 0, 0, 0}			// for the winnder
	total_moves := 9
	win := false						// to see if there is a winner, 
	print(board)
	for win == false {
		player1 := 0					// Player 1 move
		move := moves(player1, board)
		//fmt.Println(move)
		total_moves = total_moves - 1
		board[move - 1] = 0				// mark the board
		print(board)
		win = winner(board)
		if win == true {
			break
		} else if total_moves == 0 {
			break
		}
		player2 := 0					// Player 2's move
		move = moves(player2, board)
		board[move - 1] = 10				// Mark the board
		total_moves = total_moves- 1
		print(board)
		if total_moves == 0 {
			break
		}

		win = winner(board) 				// see if there is a winner
	}
							// Check to see which player won
							// if one of the index == 0, then player 1 won
	won[0] = board[0] + board[1] + board[2]		// if one of the index == 30, then player 2 won
        won[1] = board[3] + board[4] + board[5]		// There should aready be a winner considering that it broke out of the loop
        won[2] = board[6] + board[7] + board[8]

                // diagonal "lines "
        won[3] = board[0] + board[4] + board[8]
        won[4] = board[6] + board[4] + board[2]

                // Vertical "lines"
        won[5] = board[0] + board[3] + board[6]
        won[6] = board[1] + board[4] + board[7]
        won[7] = board[2] + board[5] + board[8]
	     


	for count := 0; count < 8; count++ {		// See the winnder
		if (won[count] < 1) {
			fmt.Println("\nPLAYER 1 WiN ")
			break
		} else if (won[count] == 30) {
			fmt.Println("PLAYER 2 WiN ")
			break
		}
	}
	
	 if total_moves == 0 {
                fmt.Println("Nobody wins")
        }

}
