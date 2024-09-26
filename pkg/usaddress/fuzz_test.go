package usaddress

import (
	"testing"
)

// FuzzStandardizeAddress performs fuzz testing on the StandardizeAddress function.
func FuzzStandardizeAddress(f *testing.F) {
	// Initial corpus: Add sample addresses to the corpus.
	corpus := []string{
		"123 Main St\nAnytown, NY 12345",
		"PO BOX 789\nSomecity, CA 90210",
		"456 Elm Street Apt 5B\nBigcity, TX 75001",
		"789 O'Connor Blvd\nDublin, CA 94568",
		"℅ John Doe\n321 Maple Ave\nSpringfield, IL 62704",
		"350 Fifth Ave Empire State Building\nNew York, NY 10118",
		"1234 El Niño Ave\nSan José, CA 95112",
		"50 北京路\n上海, 200001",
		"🏠 123 Happy St\nSmile Town, CA 90210",
		"Münchner Straße 45\nMünchen, 80331",
		// Add more sample addresses as needed
		"",
		"   \n  \t",
		"!@#$%^&*()_+",
		"C/O Jane Smith\n123 Unknown Rd\nMystery, ZZ 99999",
	}

	// Add the corpus entries to the fuzzer
	for _, input := range corpus {
		f.Add(input)
	}

	// Define the fuzz function
	f.Fuzz(func(t *testing.T, input string) {
		// Call the function under test
		addr := StandardizeAddress(input)

		// Optionally, you can perform validations or invariants
		// For example, ensure that the returned Address struct is valid
		if err := addr.Validate(); err != nil {
			// It's acceptable for some random inputs to produce invalid addresses
			// But you might want to check for panics or unexpected behavior
			t.Logf("Invalid address: %v", err)
		}

		// Additional checks can be added here
		// For example, you can re-serialize the address and ensure no panics occur
		_ = addr.String()
	})
}
