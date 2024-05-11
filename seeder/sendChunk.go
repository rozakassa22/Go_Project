package seeder

import(
	"net"
	"os"
	"io"
	// "strconv"
	// "fmt"
	"encoding/gob"
	// "crypto/sha1"
)

func sendChunk(buf int, filepath string, req net.Conn){

	f, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer f.Close()

	// Get the file info for the file to be shared
	info, err := f.Stat()
	if err != nil {
		return
	}

	// Calculate the number of pieces
	numPieces := info.Size() / int64(1024)
	if info.Size()%int64(1024) != 0 {
		numPieces++
	}

	// Create the torrent file data structure
	

	// Read the file piece by piece and hash each piece
	pieceBuf := make([]byte, 1024)

	for i := int64(0); i < numPieces; i++ {
		n, err := io.ReadFull(f, pieceBuf)
		if err != nil && err != io.ErrUnexpectedEOF {
			return 
		}
		if n == 0 {
			break
		}
		if int(i) == buf{
		// fmt.Println(pieceBuf[:n])
		encoder := gob.NewEncoder(req)
		err = encoder.Encode(pieceBuf[:n])
		}
	}


	
}