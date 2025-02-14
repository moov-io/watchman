package postalpool_test

import (
	"embed"
	"io/fs"
	"testing"
	"time"

	"github.com/moov-io/base/config"
	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/internal/postalpool"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed testdata/configs/config.default.yml
	configFS embed.FS
)

func TestConfig(t *testing.T) {
	logger := log.NewTestLogger()
	configService := config.NewService(logger)

	// strip ./testdata/
	fsys, err := fs.Sub(configFS, "testdata")
	require.NoError(t, err)

	var config postalpool.Config
	err = configService.LoadFromFS(&config, fsys)
	require.NoError(t, err)

	require.True(t, config.Enabled)
	require.Equal(t, 2, config.Instances)
	require.Equal(t, 10000, config.StartingPort)
	require.Equal(t, 60*time.Second, config.StartupTimeout)
	require.Equal(t, 10*time.Second, config.RequestTimeout)
	require.NotNil(t, config.Dialer)
	require.NotNil(t, config.Transport)
	require.Equal(t, "/bin/postal-server", config.BinaryPath)
}
