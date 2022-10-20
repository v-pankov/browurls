package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/pkg/browser"
)

func main() {
	if err := exec(); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
	fmt.Println("DONE")
}

func exec() error {
	args, err := parseArgs()
	if err != nil {
		return fmt.Errorf("invalid arguments: %w", err)
	}

	file, err := os.Open(args.FilePath)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var (
			nextLine = scanner.Text()
			url, err = url.Parse(nextLine)
		)

		if err != nil {
			return fmt.Errorf("parse url [%s]: %w", nextLine, err)
		}

		if err := browser.OpenURL(url.String()); err != nil {
			return fmt.Errorf("open url [%s]: %w", url.String(), err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scan: %w", err)
	}

	return nil
}

func parseArgs() (*Args, error) {
	cmdArgs := os.Args[1:]

	if len(cmdArgs) < 1 {
		return nil, errors.New("not enough arguments")
	}

	if len(cmdArgs) > 1 {
		return nil, errors.New("more than enough arguments")
	}

	if _, err := os.Stat(cmdArgs[0]); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("file does not exist")
	}

	return &Args{
		FilePath: cmdArgs[0],
	}, nil
}

type Args struct {
	FilePath string
}
