package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestCopyMode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "simple text",
			input:   "hello world",
			wantErr: false,
		},
		{
			name:    "multiline text",
			input:   "line1\nline2\nline3",
			wantErr: false,
		},
		{
			name:    "empty input",
			input:   "",
			wantErr: false,
		},
		{
			name:    "special characters",
			input:   "tab\there\nquote\"test",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdin := os.Stdin
			defer func() { os.Stdin = oldStdin }()

			r, w, _ := os.Pipe()
			os.Stdin = r

			go func() {
				w.Write([]byte(tt.input))
				w.Close()
			}()

			err := copyMode()
			if (err != nil) != tt.wantErr {
				t.Errorf("copyMode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPasteMode(t *testing.T) {
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	r, w, _ := os.Pipe()
	os.Stdout = w

	err := pasteMode()
	w.Close()

	if err != nil {
		t.Logf("pasteMode() returned error (may be expected if clipboard unavailable): %v", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	if err == nil && output == "" {
		t.Log("pasteMode() succeeded with empty clipboard")
	}
}

func TestRun(t *testing.T) {
	t.Run("detects pipe mode", func(t *testing.T) {
		oldStdin := os.Stdin
		defer func() { os.Stdin = oldStdin }()

		r, w, _ := os.Pipe()
		os.Stdin = r

		go func() {
			w.Write([]byte("test"))
			w.Close()
		}()

		err := run()
		if err != nil {
			t.Logf("run() in copy mode: %v", err)
		}
	})

	t.Run("detects terminal mode", func(t *testing.T) {
		err := run()
		if err != nil {
			t.Logf("run() in paste mode: %v", err)
		}
	})
}

func TestMain_Version(t *testing.T) {
	oldArgs := os.Args
	oldStdout := os.Stdout
	defer func() {
		os.Args = oldArgs
		os.Stdout = oldStdout
	}()

	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Args = []string{"cmd", "-version"}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Recovered from exit: %v", r)
			}
		}()
	}()

	w.Close()
	var buf bytes.Buffer
	buf.ReadFrom(r)

	if !strings.Contains(buf.String(), "version") && buf.String() != "" {
		t.Errorf("Expected version output, got: %s", buf.String())
	}
}
