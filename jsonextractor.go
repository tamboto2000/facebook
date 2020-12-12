package facebook

import (
	"bytes"
	"io"
)

func extractJSONBytes(str []byte) ([][]byte, error) {
	reader := bytes.NewReader(str)
	r := newReader(reader)
	return extractJSON(r)
}

func extractJSONFromReader(reader io.Reader) ([][]byte, error) {
	r := newReader(reader)
	return extractJSON(r)
}

func extractJSON(r *reader) ([][]byte, error) {
	jsons := make([][]byte, 0)
	byts := make([]byte, 0)
	brackOpen := 0
	brackClose := 0
	var prevChar byte
	for r.next() {
		//r.byt == "{"
		if r.byt[0] == 123 {
			brackOpen++
		}

		//r.byt == "}"
		if brackOpen > 0 && r.byt[0] == 125 {
			brackClose++
		}

		//"{"
		if prevChar == 123 {
			if !isAllowedChar(r.byt[0]) && !isNumber(r.byt[0]) {
				brackOpen = 0
				brackClose = 0
				prevChar = 0
				byts = make([]byte, 0)

				continue
			}
		}

		if brackOpen > 0 {
			byts = append(byts, r.byt[0])
			prevChar = r.byt[0]

			if brackOpen == brackClose {
				jsons = append(jsons, byts)
				brackOpen = 0
				brackClose = 0
				prevChar = 0
				byts = make([]byte, 0)
				continue
			}
		}
	}

	if r.err != nil {
		return nil, r.err
	}

	r = nil
	return jsons, nil
}

type reader struct {
	r   io.Reader
	byt []byte
	err error
}

func newReader(r io.Reader) *reader {
	return &reader{
		byt: make([]byte, 1),
		r:   r,
	}
}

func (r *reader) next() bool {
	_, err := r.r.Read(r.byt)
	if err != nil {
		if err != io.EOF {
			r.err = err
		}

		return false
	}

	return true
}

func isNumber(char byte) bool {
	// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
	if char != 48 && char != 49 && char != 50 &&
		char != 51 && char != 52 && char != 53 &&
		char != 54 && char != 55 && char != 56 && char != 57 {
		return false
	}

	return true
}

func isAllowedChar(char byte) bool {
	// {, }, [,  , \", \t, \v, \n, \
	if char != 123 && char != 125 && char != 91 &&
		char != 32 && char != 34 && char != 9 &&
		char != 11 && char != 10 && char != 92 {
		return false
	}

	return true
}
