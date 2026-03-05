package clipboard

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func Write(text string) error {
	if err := clipboard.WriteAll(text); err != nil {
		return fmt.Errorf("clipboard write failed: %w", err)
	}
	return nil
}

func Read() (string, error) {
	text, err := clipboard.ReadAll()
	if err != nil {
		return "", fmt.Errorf("clipboard read failed: %w", err)
	}
	return text, nil
}
