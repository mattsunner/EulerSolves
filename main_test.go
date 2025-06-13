package main

import (
	"io"
	"os"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Test1", "1\n", "1"},
		{"Test2", "hello\n", "hello"},
		{"Test3", "world\n", "world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}

			os.Stdin = r

			go func() {
				defer w.Close()
				io.WriteString(w, tt.input)
			}()

			got := getUserInput("test prompt")
			if got != tt.want {
				t.Errorf("getUserInput() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPrintMenu(t *testing.T) {
	// This test checks if the printMenu function runs without errors.
	// Since printMenu does not return any value, we cannot check its output directly.
	// Instead, we can check if it runs without panicking.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("printMenu function panicked: %v", r)
		}
	}()

	printMenu()
}

func TestMain(t *testing.T) {
	// This is a placeholder test for the main function.
	// Since main does not return any value, we cannot test it directly.
	// Instead, we can check if it runs without panicking.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main function panicked: %v", r)
		}
	}()

	main()
}
