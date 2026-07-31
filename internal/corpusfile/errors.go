package corpusfile

import "fmt"

type formatError struct {
	msg string
}

func (e formatError) Error() string { return "corpusfile: " + e.msg }

func errShort(what string, want, got int) error {
	return formatError{msg: fmt.Sprintf("short %s: need %d bytes, got %d", what, want, got)}
}

func errBadMagic(got string) error {
	return formatError{msg: fmt.Sprintf("bad magic %q, want %q", got, Magic)}
}

func errCodec(c uint8) error {
	return formatError{msg: fmt.Sprintf("unsupported cold codec %d", c)}
}

func errEntityCount(want, got int) error {
	return formatError{msg: fmt.Sprintf("entity count mismatch: header %d, got %d", want, got)}
}
