package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	url := flag.String("url", "", "The URL of the .m3u8 playlist")
	output := flag.String("output", "output.mp4", "Output filename")

	flag.Parse()

	if *url == "" {
		fmt.Println("Error: --url flag is required")
		os.Exit(1)
	}

	cmd := exec.Command(
		"ffmpeg",
		"-i", *url,
		"-c", "copy",
		"-bsf:a", "aac_adtstoasc",
		*output,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Starting download...")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running ffmpeg: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Download completed successfully!")
}
