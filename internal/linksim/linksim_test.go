package linksim

import (
	"fmt"
	"testing"
)

func TestGenerateHashes(t *testing.T) {
	t.Run("John, Johnny, Johnathon", func(t *testing.T) {
		johnHashes := GenerateHashes(john)
		fmt.Println("john")
		for key, value := range johnHashes {
			fmt.Printf("  %v - %v\n", key, value)
		}
		fmt.Printf("buckets(8): %#v\n", johnHashes.GenerateBuckets(32))
		fmt.Printf("buckets(16): %#v\n", johnHashes.GenerateBuckets(32))
		fmt.Printf("buckets(32): %#v\n", johnHashes.GenerateBuckets(32))
		fmt.Printf("buckets(64): %#v\n", johnHashes.GenerateBuckets(32))

		fmt.Printf("\n\n")

		johnathonHashes := GenerateHashes(johnathon)
		fmt.Println("johnathon")
		for key, value := range johnathonHashes {
			fmt.Printf("  %v - %v\n", key, value)
		}
		fmt.Printf("buckets(8): %#v\n", johnathonHashes.GenerateBuckets(32))
		fmt.Printf("buckets(16): %#v\n", johnathonHashes.GenerateBuckets(32))
		fmt.Printf("buckets(32): %#v\n", johnathonHashes.GenerateBuckets(32))
		fmt.Printf("buckets(64): %#v\n", johnathonHashes.GenerateBuckets(32))
	})
}
