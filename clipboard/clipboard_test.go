package clipboard

import (
	"testing"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		wantErr bool
	}{
		{
			name:    "simple text",
			text:    "hello",
			wantErr: false,
		},
		{
			name:    "empty string",
			text:    "",
			wantErr: false,
		},
		{
			name:    "multiline",
			text:    "line1\nline2\nline3",
			wantErr: false,
		},
		{
			name:    "unicode",
			text:    "Hello 世界 🌍",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Write(tt.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				got, readErr := Read()
				if readErr != nil {
					t.Fatalf("Read() after Write() failed: %v", readErr)
				}
				if got != tt.text {
					t.Errorf("Read() = %q, want %q", got, tt.text)
				}
			}
		})
	}
}

func TestRead(t *testing.T) {
	testData := "test clipboard content"
	err := Write(testData)
	if err != nil {
		t.Fatalf("Setup Write() failed: %v", err)
	}

	got, err := Read()
	if err != nil {
		t.Errorf("Read() error = %v", err)
		return
	}

	if got != testData {
		t.Errorf("Read() = %q, want %q", got, testData)
	}
}

func TestWriteRead_RoundTrip(t *testing.T) {
	tests := []string{
		"simple",
		"with\nnewlines\n",
		"tabs\there",
		"quotes\"and'apostrophes",
	}

	for _, text := range tests {
		t.Run(text, func(t *testing.T) {
			if err := Write(text); err != nil {
				t.Fatalf("Write() failed: %v", err)
			}

			got, err := Read()
			if err != nil {
				t.Fatalf("Read() failed: %v", err)
			}

			if got != text {
				t.Errorf("Round trip failed: got %q, want %q", got, text)
			}
		})
	}
}
