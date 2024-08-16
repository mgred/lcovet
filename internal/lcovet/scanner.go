package lcovet

import (
	"bufio"
	"bytes"
	"io"
)

type Scanner struct {
	r *bufio.Reader
}

func (s *Scanner) Scan() (tok Token, buf bytes.Buffer) {
	var key bytes.Buffer
	for {
		ch := s.read()
		if ch == ':' || ch == eof {
			break
		} else {
			key.WriteRune(ch)
		}
	}

	if tok = getIdent(key); tok == EOF {
		return
	} else {
		for {
			if ch := s.read(); ch != eof {
				buf.WriteRune(ch)
				continue
			}
			break
		}
		return
	}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func getIdent(i bytes.Buffer) Token {
	switch i.String() {
	case "TN":
		return TN
	case "SF":
		return SF
	case "FN":
		return FN
	case "FNDA":
		return FNDA
	case "FNF":
		return FNF
	case "FNH":
		return FNH
	case "BRDA":
		return BRDA
	case "BRF":
		return BRF
	case "BRH":
		return BRH
	case "DA":
		return DA
	case "LF":
		return LF
	case "LH":
		return LH
	case "end_of_record":
		return EOF
	}
	return ILLEGAL
}
