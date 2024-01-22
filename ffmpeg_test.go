package ffmpeg_test

import (
	"testing"

	"github.com/csnewman/ffmpeg-go"
	"github.com/stretchr/testify/assert"
)

func TestVersions(t *testing.T) {
	assert.Equal(t, 3940198, int(ffmpeg.AVCodecVersion()), "AVCodec version should match expected")
	assert.Equal(t, ffmpeg.LibAVCodecVersionInt, int(ffmpeg.AVCodecVersion()), "AVCodec version func and const should match")
}
