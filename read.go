package ioutil

import (
	// Built-in/core modules.
	"bufio"
	"io"
	"iter"
)

// Returns a sequence of strings read from the given reader, split by the
// specified delimiter (and including the delimiter in the result). Each
// string is yielded along with any error encountered during reading. EOF is
// not considered an error and will yield the last line read before EOF. If
// the reader is empty, it will yield an empty string and nil error.
func ReadStringSeq(r io.Reader, delim byte) iter.Seq2[string, error] {
	b := bufio.NewReader(r)
	return func(yield func(string, error) bool) {
		for {
			line, err := b.ReadString(delim)
			if err != nil {
				if err == io.EOF {
					if line != "" {
						yield(line, nil)
					}
					return
				}
				yield(line, err)
				return
			}
			if !yield(line, nil) {
				return
			}
		}
	}
}

// Returns a sequence of byte slices read from the given reader, split by the
// specified delimiter (and including the delimiter in the result). Each
// byte slice is yielded along with any error encountered during reading. EOF
// is not considered an error and will yield the last line read before EOF.
// If the reader is empty, it will yield an empty byte slice and nil error.
func ReadBytesSeq(r io.Reader, delim byte) iter.Seq2[[]byte, error] {
	b := bufio.NewReader(r)
	return func(yield func([]byte, error) bool) {
		for {
			line, err := b.ReadBytes(delim)
			if err != nil {
				if err == io.EOF {
					if len(line) > 0 {
						yield(line, nil)
					}
					return
				}
				yield(line, err)
				return
			}
			if !yield(line, nil) {
				return
			}
		}
	}
}
