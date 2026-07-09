package us_tel

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

type TELProperties struct {
	Program []string `json:"program"`
	Topics  []string `json:"topics"`
	Name    []string `json:"name"`
}

type Reader struct {
	source io.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{source: r}
}

func (r *Reader) Read(onRecord func(TELRecord)) error {
	if r.source == nil {
		return fmt.Errorf("source reader is nil")
	}

	br := bufio.NewReader(r.source)

	// Peek at the first non-whitespace character to determine format
	var firstByte byte
	for {
		b, err := br.ReadByte()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if b == ' ' || b == '\t' || b == '\r' || b == '\n' {
			continue
		}
		firstByte = b
		err = br.UnreadByte()
		if err != nil {
			return err
		}
		break
	}

	decoder := json.NewDecoder(br)

	if firstByte == '[' {
		// Case 1: JSON Array
		token, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("failed to read initial array token: %w", err)
		}
		if delim, ok := token.(json.Delim); !ok || delim != '[' {
			return fmt.Errorf("expected JSON array starting with '['")
		}

		for decoder.More() {
			var record TELRecord
			if err := decoder.Decode(&record); err != nil {
				return fmt.Errorf("failed to decode array record: %w", err)
			}
			onRecord(record)
		}

		_, err = decoder.Token()
		return err
	} else if firstByte == '{' {
		// Case 2: JSON Lines (newline-delimited JSON objects)
		for {
			var record TELRecord
			if err := decoder.Decode(&record); err != nil {
				if err == io.EOF {
					return nil
				}
				return fmt.Errorf("failed to decode JSON lines record: %w", err)
			}
			onRecord(record)
		}
	}

	return fmt.Errorf("expected JSON array starting with '[' or JSON Lines starting with '{', got %q", firstByte)
}
