package mp4

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoxTypeString(t *testing.T) {
	assert.Equal(t, "1234", BoxType{'1', '2', '3', '4'}.String())
}

func TestIsSupported(t *testing.T) {
	assert.True(t, StrToBoxType("pssh").IsSupported())
	assert.False(t, StrToBoxType("1234").IsSupported())
}

func TestGetSupportedVersions(t *testing.T) {
	vers, err := BoxTypePssh().GetSupportedVersions()
	require.NoError(t, err)
	assert.Equal(t, []uint8{0, 1}, vers)
}

func TestIsSupportedVersion(t *testing.T) {
	assert.True(t, BoxTypePssh().IsSupportedVersion(0))
	assert.True(t, BoxTypePssh().IsSupportedVersion(1))
	assert.False(t, BoxTypePssh().IsSupportedVersion(2))
}
