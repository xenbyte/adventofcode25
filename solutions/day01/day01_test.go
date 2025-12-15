package day01

import (
	"os"
	"path/filepath"
	"testing"
)

func mustReadFile(t *testing.T, name string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(".", name))
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return string(b)
}

func TestPart1Example(t *testing.T) {
	input := mustReadFile(t, "example.txt")
	got, err := sol{}.Part1(input)
	if err != nil {
		t.Fatalf("Part1 error: %v", err)
	}
	_ = got
	// TODO: set expected
	// if got != 123 { t.Fatalf("got %v", got) }
}

func TestPart2Example(t *testing.T) {
	input := mustReadFile(t, "example.txt")
	got, err := sol{}.Part2(input)
	if err != nil {
		t.Fatalf("Part2 error: %v", err)
	}
	_ = got
	// TODO: set expected
	// if got != 456 { t.Fatalf("got %v", got) }
}
