package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const version = "v0.0.1"

func main() {
	fmt.Printf("🎬 m3u8-download %s\n\n", version)

	url := flag.String("url", "", "The URL of the .m3u8 playlist")
	output := flag.String("output", "output.mp4", "Output filename")
	outputDir := flag.String("output-dir", ".", "Directory to save the output file")
	timeout := flag.Int("timeout", 30, "Timeout in seconds for stalled downloads")
	quiet := flag.Bool("quiet", false, "Suppress ffmpeg output logs")

	flag.Parse()

	if *url == "" {
		fmt.Println("Error: --url flag is required")
		os.Exit(1)
	}

	if stat, err := os.Stat(*outputDir); err != nil {
		fmt.Printf("Error: output directory does not exist: %s\n", *outputDir)
		os.Exit(1)
	} else if !stat.IsDir() {
		fmt.Printf("Error: output path is not a directory: %s\n", *outputDir)
		os.Exit(1)
	}

	outputPath := filepath.Join(*outputDir, *output)

	cmd := exec.Command(
		"ffmpeg",
		"-rw_timeout", fmt.Sprintf("%d", *timeout*1000000),
		"-i", *url,
		"-c", "copy",
		"-bsf:a", "aac_adtstoasc",
		outputPath,
	)

	if *quiet {
		cmd.Stdout = nil
		cmd.Stderr = nil
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	fmt.Println("Starting download...")

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running ffmpeg: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Download completed successfully!")
}
