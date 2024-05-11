package main

import (
	"strconv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"torrentassignment/generateTorrent"
	"torrentassignment/leecher"
	"torrentassignment/seeder"
)

func main() {
	args := os.Args

	fmt.Println(len(args))

	command := args[1]

	switch command {
	case "generateTorrent":
		if len(args) == 3 {
			fmt.Println("Generating torrent ")
			generateTorrent.GenerateTorrent(os.Args[2], "http://104.28.16.69/announce", 1024)
		} else {
			fmt.Println("No")
		}

	case "seed":
		if len(args) == 3 {
			fmt.Println("Seeding torrent:", args[2])
			seeder.Seed(args[2])
		} else {
			fmt.Println("No")
		}

	case "download":
		if len(args) == 3 {
			fmt.Println("Download the given torrent")
			leecher.Leech(args[2])
		} else {
			fmt.Println("No")
		}

	default:
		fmt.Println("Unknown command:", command)
	}

	http.HandleFunc("/newrequiements.txt", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("newrequiements.txt")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		defer file.Close()

		// Get file information
		fileInfo, err := file.Stat()
		if err != nil {
			http.Error(w, "Failed to get file information", http.StatusInternalServerError)
			return
		}

		// Set the appropriate headers for downloading the file
		w.Header().Set("Content-Disposition", "attachment; filename=newrequiements.txt")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

		// Copy the file content to the response writer
		io.Copy(w, file)
	})

	log.Println("Server started on http://192.168.244.141:2700")
	err := http.ListenAndServe("192.168.244.141:2700", nil)
	if err != nil {
		log.Fatal(err)
	}

}
// The provided code is a command-line utility that supports generating torrents, seeding files, leeching files, and serving a specific file for downloading via an HTTP server. The `main` function parses command-line arguments to determine the desired operation, including "generateTorrent" for creating torrent files, "seed" for sharing files, and "download" for downloading files using the leeching process. Additionally, the code sets up an HTTP server to serve a specific file for downloading, handling requests for the file path "/newrequirements.txt" and responding with the file's content. This code facilitates torrent operations, file sharing, and file downloading through torrent and HTTP protocols.