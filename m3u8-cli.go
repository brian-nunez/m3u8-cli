package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const version = "v0.0.1"

type Config struct {
	URL         string `flag:"url" default:"" help:"The URL of the .m3u8 playlist"`
	Output      string `flag:"output" default:"output.mp4" help:"Output filename"`
	OutputDir   string `flag:"output-dir" default:"." help:"Directory to save the output file"`
	Timeout     int    `flag:"timeout" default:"30" help:"Timeout in seconds for stalled downloads"`
	Quiet       bool   `flag:"quiet" default:"false" help:"Suppress ffmpeg output logs"`
	ShowVersion bool   `flag:"version" default:"false" help:"Print the version and exit"`
}

func main() {
	fmt.Printf("🎬 m3u8-download %s\n\n", version)

	config := &Config{}
	BindFlags(config)
	flag.Parse()

	if config.ShowVersion {
		fmt.Printf("m3u8-download version: %s\n", version)
		os.Exit(0)
	}

	if config.URL == "" {
		fmt.Println("Error: --url flag is required")
		os.Exit(1)
	}

	if stat, err := os.Stat(config.OutputDir); err != nil {
		fmt.Printf("Error: output directory does not exist: %s\n", config.OutputDir)
		os.Exit(1)
	} else if !stat.IsDir() {
		fmt.Printf("Error: output path is not a directory: %s\n", config.OutputDir)
		os.Exit(1)
	}

	outputPath := filepath.Join(config.OutputDir, config.Output)

	cmd := exec.Command(
		"ffmpeg",
		"-rw_timeout", fmt.Sprintf("%d", config.Timeout*1000000),
		"-i", config.URL,
		"-c", "copy",
		"-bsf:a", "aac_adtstoasc",
		outputPath,
	)

	if config.Quiet {
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
