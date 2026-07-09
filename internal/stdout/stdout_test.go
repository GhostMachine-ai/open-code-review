package stdout

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestWriter_Default(t *testing.T) {
	w := Writer()
	if w != os.Stdout {
		t.Error("expected default Writer to be os.Stdout")
	}
}

func TestQuiet(t *testing.T) {
	restore := Quiet()

	w := Writer()
	if w != io.Discard {
		t.Error("expected Writer to be io.Discard after Quiet()")
	}

	restore()

	w = Writer()
	if w != os.Stdout {
		t.Error("expected Writer to be os.Stdout after restore")
	}
}

func TestRedirect(t *testing.T) {
	var buf bytes.Buffer
	restore := Redirect(&buf)
	defer restore()

	if Writer() != &buf {
		t.Fatalf("Writer() did not return the redirect target")
	}
	if _, err := Writer().Write([]byte("hello")); err != nil {
		t.Fatalf("write: %v", err)
	}
	if got := buf.String(); got != "hello" {
		t.Errorf("buf = %q, want %q", got, "hello")
	}
}

func TestRedirectRestore(t *testing.T) {
	var buf bytes.Buffer
	restore := Redirect(&buf)
	restore()

	if Writer() != os.Stdout {
		t.Error("expected Writer to be os.Stdout after redirect restore")
	}
}
