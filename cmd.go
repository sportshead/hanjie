package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: hanjie <file> <step delay>")
	}

	filename := os.Args[1]
	durationStr := os.Args[2]

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Fatal(err)
	}

	hanjie, err := parseFile(filename, duration)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Parsed Hanjie object:\n%+v\nmax row: %d\nmax col: %d", hanjie, *hanjie.Hints.MaxRowSize, *hanjie.Hints.MaxColSize)

	hanjie.Print()

	hanjie.SolveAll()
}

func parseFile(filename string, delay time.Duration) (*Hanjie, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	hanjie := &Hanjie{
		StepDelay: delay,
	}

	hanjie.Hints.MaxRowSize = new(int)
	hanjie.Hints.MaxColSize = new(int)
	*hanjie.Hints.MaxRowSize = -1
	*hanjie.Hints.MaxColSize = -1

	err = parseHints(reader, &hanjie.Hints.Row, hanjie.Hints.MaxRowSize)
	if err != nil {

	}

	err = parseHints(reader, &hanjie.Hints.Column, hanjie.Hints.MaxColSize)
	if err != nil {
		return nil, err
	}

	hanjie.Height = len(hanjie.Hints.Row)
	hanjie.Width = len(hanjie.Hints.Column)

	hanjie.InitializeGrid()

	return hanjie, nil
}

func parseHints(reader *bufio.Reader, hints *[]Hint, size *int) error {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		line = strings.TrimSpace(line)
		if line == "---" {
			break
		}

		hint, err := parseHint(line)
		if err != nil {
			return fmt.Errorf("failed to parse hint: %v", err)
		}

		if len(hint) > *size {
			*size = len(hint)
		}

		*hints = append(*hints, hint)
	}

	return nil
}

func parseHint(line string) (Hint, error) {
	parts := strings.Split(line, " ")
	hint := make(Hint, len(parts))

	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("failed to convert hint value to integer: %v", err)
		}

		hint[i] = value
	}

	return hint, nil
}
