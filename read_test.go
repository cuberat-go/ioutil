package ioutil_test

import (
	// Built-in/core modules.
	"strings"
	"testing"

	// First-party modules.
	"github.com/cuberat-go/ioutil"
)

func TestReadStringSeq(t *testing.T) {
	input := "line1\nline2\nline3\n"
	expected := []string{"line1\n", "line2\n", "line3\n"}

	t.Run("final empty line", func(t *testing.T) {
		testReadStringSeq(t, input, expected)
	})

	input = "line1\nline2\nline3"
	expected = []string{"line1\n", "line2\n", "line3"}

	t.Run("no final empty line", func(t *testing.T) {
		testReadStringSeq(t, input, expected)
	})

}

func testReadStringSeq(
	t *testing.T,
	input string,
	expected []string,
) {
	r := strings.NewReader(input)
	seq := ioutil.ReadStringSeq(r, '\n')

	var results []string
	seq(func(s string, err error) bool {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		results = append(results, s)
		return true
	})

	if len(results) != len(expected) {
		t.Fatalf("expected %d results, got %d", len(expected), len(results))
	}

	for i, v := range expected {
		if results[i] != v {
			t.Errorf("expected %q, got %q", v, results[i])
		}
	}
}

func TestReadBytesSeq(t *testing.T) {
	input := "line1\nline2\nline3\n"
	expected := [][]byte{[]byte("line1\n"), []byte("line2\n"), []byte("line3\n")}

	t.Run("final empty line", func(t *testing.T) {
		testReadBytesSeq(t, input, expected)
	})

	input = "line1\nline2\nline3"
	expected = [][]byte{[]byte("line1\n"), []byte("line2\n"), []byte("line3")}

	t.Run("no final empty line", func(t *testing.T) {
		testReadBytesSeq(t, input, expected)
	})
}

func testReadBytesSeq(
	t *testing.T,
	input string,
	expected [][]byte,
) {
	r := strings.NewReader(input)
	seq := ioutil.ReadBytesSeq(r, '\n')

	var results [][]byte
	seq(func(b []byte, err error) bool {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		results = append(results, b)
		return true
	})

	if len(results) != len(expected) {
		t.Fatalf("expected %d results, got %d", len(expected), len(results))
	}

	for i, v := range expected {
		if string(results[i]) != string(v) {
			t.Errorf("expected %q, got %q", v, results[i])
		}
	}
}
