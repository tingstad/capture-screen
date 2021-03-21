package main

import (
	"io"
	"strings"
	"testing"
)

type TestReader struct {
	text  string
	count int
}

func (r *TestReader) ReadString(delim byte) (string, error) {
	if r.count < 1 {
		r.count++
		return r.text, nil
	}
	return "", io.EOF
}

func TestOneLine(t *testing.T) {
	lines := CaptureReader(strings.NewReader("hello\n"))

	got := strings.Join(lines, "")
	if got != "hello" {
		t.Errorf("Want \"hello\", got \"%s\"", got)
	}
}

func TestTwoLines(t *testing.T) {
	lines := CaptureReader(strings.NewReader("hello\nworld\n"))

	got := len(lines)
	if got != 2 {
		t.Errorf("Want 2, got %d", got)
	}
	one := lines[0]
	if one != "hello" {
		t.Errorf("Want \"hello\", got %s", one)
	}
	two := lines[1]
	if two != "world" {
		t.Errorf("Want \"world\", got %s", two)
	}
}

func TestPrint(t *testing.T) {
	screen := make([]string, 0)
	lines := Print(screen, "hello", 0, 0)

	got := strings.Join(lines, "")
	if got != "hello" {
		t.Errorf("Want \"hello\", got %s", got)
	}
}

func TestPrintDown(t *testing.T) {
	screen := make([]string, 0)
	lines := Print(screen, "hello", 0, 2)

	got := strings.Join(lines, ",")
	if got != ",,hello" {
		t.Errorf("Want \",,hello\", got %s", got)
	}
}

func TestPrintOver(t *testing.T) {
	screen := []string{"hello"}
	lines := Print(screen, "world", 0, 0)

	got := strings.Join(lines, "")
	if got != "world" {
		t.Errorf("Want \"world\", got %s", got)
	}
}

func TestPrintOverPartly(t *testing.T) {
	screen := []string{"hello"}
	lines := Print(screen, "world", 4, 0)

	got := strings.Join(lines, "")
	if got != "hellworld" {
		t.Errorf("Want \"hellworld\", got %s", got)
	}
	got = Print(lines, "hi, ", 0, 0)[0]
	if got != "hi, world" {
		t.Errorf("Want \"hi, world\", got %s", got)
	}
	got = Print([]string{"hello world"}, "owdy ", 1, 0)[0]
	if got != "howdy world" {
		t.Errorf("Want \"howdy world\", got %s", got)
	}
	got = Print([]string{"hello"}, "world", 10, 0)[0]
	if got != "hello     world" {
		t.Errorf("Want \"hello     world\", got %s", got)
	}
}

func TestDown(t *testing.T) {
	lines := CaptureReader(strings.NewReader("hello\x1b[Bhi\n"))

	got := strings.Join(lines, ",")
	if got != "hello,     hi" {
		t.Errorf("Want \"hello,     hi\", got \"%s\"", got)
	}
}

func TestUp(t *testing.T) {
	lines := CaptureReader(strings.NewReader("hello\n\x1b[Aansi\n"))

	got := strings.Join(lines, "")
	if got != "ansio" {
		t.Errorf("Want \"ansio\", got \"%s\"", got)
	}
}

func TestUpDown(t *testing.T) {
	lines := CaptureReader(strings.NewReader("one \x1b[2B two \x1b[2A three\n"))

	got := strings.Join(lines, ":")
	want := "one       three::     two "
	if got != want {
		t.Errorf("Want:\n%s\ngot:\n%s", want, got)
	}
}

func TestLeftRight(t *testing.T) {
	lines := CaptureReader(strings.NewReader("\x1b[10C world \x1b[14D hello,\n"))

	got := strings.Join(lines, "")
	want := "    hello, world "
	if got != want {
		t.Errorf("Want:\n%s\ngot:\n%s", want, got)
	}
}
