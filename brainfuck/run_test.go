package brainfuck

import (
	"testing"
)

func TestRunFiles(t *testing.T) {
	t.Run("Beer", func(t *testing.T) {
		_, err := RunFile("./bftests/beer.b")

		if err != nil {
			t.Errorf("Got an error: %v", err)
		}

	})

	t.Run("Hello", func(t *testing.T) {
		_, err := RunFile("./bftests/hello.b")

		if err != nil {
			t.Errorf("Got an error: %v", err)
		}

	})

	t.Run("Hello2", func(t *testing.T) {
		_, err := RunFile("./bftests/hello2.b")

		if err != nil {
			t.Errorf("Got an error: %v", err)
		}

	})

	t.Run("End", func(t *testing.T) {
		_, err := RunFile("./bftests/end.b")

		if err != nil {
			t.Errorf("Got an error: %v", err)
		}

	})
}
