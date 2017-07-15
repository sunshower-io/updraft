package io

import "io"

//pretty much copied from https://golang.org/src/io/multi.go



type eofReader struct{}

func (eofReader) Read([]byte) (int, error) {
    return 0, EOF
}

type MultiReader struct {
    readers []io.Reader
}

func (m *MultiReader) Append(r io.Reader) {
    m.readers = append(m.readers, r)
}

func (mr *MultiReader) Read(p []byte) (n int, err error) {
    for len(mr.readers) > 0 {
        // Optimization to flatten nested multiReaders (Issue 13558).
        if len(mr.readers) == 1 {
            if r, ok := mr.readers[0].(*MultiReader); ok {
                mr.readers = r.readers
                continue
            }
        }
        n, err = mr.readers[0].Read(p)
        if err == EOF {
            // Use eofReader instead of nil to avoid nil panic
            // after performing flatten (Issue 18232).
            mr.readers[0] = eofReader{} // permit earlier GC
            mr.readers = mr.readers[1:]
        }
        if n > 0 || err != EOF {
            if err == EOF && len(mr.readers) > 0 {
                // Don't return EOF yet. More readers remain.
                err = nil
            }
            return
        }
    }
    return 0, EOF
}

func NewMultiReader() io.Reader {
    return &MultiReader{readers: make([]io.Reader, 0)}
}
