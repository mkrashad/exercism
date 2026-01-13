package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    totalCount := 0
    fileData, exists := cb[file]
    if !exists {
		return 0
	}
	for _, occupied := range fileData {
		if occupied {
			totalCount++
		}
	}
	return totalCount
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    totalCount := 0
	if rank < 1 || rank > 8 {
        return 0
    }
    for _, rows := range cb {
        for index, occupied := range rows{
            if index + 1 == rank && occupied{
                totalCount++
            }
        }
 
	}
    return totalCount
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
  	total := 0
 	for _, ranks := range cb {
      total += len(ranks)
	}
	return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	total := 0
 	for _, ranks := range cb {
      for _, occupied := range ranks{
            if occupied{
                total++
            }
		}
	}
	return total
}
