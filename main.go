package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/user/clipboard-pipe/clipboard"
)

const version = "1.0.0"

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.Parse()

	if showVersion {
		fmt.Println("clipboard-pipe version", version)
		os.Exit(0)
	}

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat stdin: %w", err)
	}

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return copyMode()
	}
	return pasteMode()
}

func copyMode() error {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("failed to read stdin: %w", err)
	}

	if err := clipboard.Write(string(data)); err != nil {
		return fmt.Errorf("failed to write to clipboard: %w", err)
	}

	return nil
}

func pasteMode() error {
	data, err := clipboard.Read()
	if err != nil {
		return fmt.Errorf("failed to read from clipboard: %w", err)
	}

	if _, err := os.Stdout.WriteString(data); err != nil {
		return fmt.Errorf("failed to write to stdout: %w", err)
	}

	return nil
}
