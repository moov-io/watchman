package ofac

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkReader(b *testing.B) {
	b.Run("SDN", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()

			var results Results
			fd, err := os.Open(filepath.Join("testdata", "sdn.csv"))
			require.NoError(b, err)

			b.StartTimer()
			hash, err := csvSDNFile(&results, fd)
			b.StopTimer()

			require.NoError(b, err)
			require.NotEmpty(b, results.SDNs)
			require.NotEmpty(b, hash)
		}
	})

	b.Run("Addresses", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()

			results := Results{
				Addresses: make(map[string][]Address, estimatedAddresses),
			}
			fd, err := os.Open(filepath.Join("testdata", "add.csv"))
			require.NoError(b, err)

			b.StartTimer()
			hash, err := csvAddressFile(&results, fd)
			b.StopTimer()

			require.NoError(b, err)
			require.NotEmpty(b, results.Addresses)
			require.NotEmpty(b, hash)
		}
	})

	b.Run("AlternateIdentities", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()

			results := Results{
				AlternateIdentities: make(map[string][]AlternateIdentity, estimatedAlts),
			}
			fd, err := os.Open(filepath.Join("testdata", "alt.csv"))
			require.NoError(b, err)

			b.StartTimer()
			hash, err := csvAlternateIdentityFile(&results, fd)
			b.StopTimer()

			require.NoError(b, err)
			require.NotEmpty(b, results.AlternateIdentities)
			require.NotEmpty(b, hash)
		}
	})

	b.Run("Comments", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()

			results := Results{
				SDNComments: make(map[string][]SDNComments, estimatedComments),
			}
			fd, err := os.Open(filepath.Join("testdata", "sdn_comments.csv"))
			require.NoError(b, err)

			b.StartTimer()
			hash, err := csvSDNCommentsFile(&results, fd)
			b.StopTimer()

			require.NoError(b, err)
			require.NotEmpty(b, results.SDNComments)
			require.NotEmpty(b, hash)
		}
	})
}
