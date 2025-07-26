package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRecover(t *testing.T) {
	// Create a pipe to capture output
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	// Redirect os.Stdout to the pipe writer
	stdout := os.Stdout
	os.Stdout = w

	// Call the main function
	main()

	// Restore os.Stdout and close the pipe writer
	os.Stdout = stdout
	w.Close()

	// Read the output from the pipe reader
	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	// Check the output
	output := buf.String()
	expected := "Recovered. Error:\n a problem\n"

	if output != expected {
		t.Errorf("Unexpected output:\nGot:\n%s\nExpected:\n%s", output, expected)
	}
}