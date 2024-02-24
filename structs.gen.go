package ffmpeg

import "unsafe"

// #include <libavcodec/avcodec.h>
// #include <libavcodec/codec.h>
// #include <libavcodec/codec_desc.h>
// #include <libavcodec/codec_id.h>
// #include <libavcodec/codec_par.h>
// #include <libavcodec/defs.h>
// #include <libavcodec/packet.h>
// #include <libavcodec/version.h>
// #include <libavcodec/version_major.h>
// #include <libavdevice/version.h>
// #include <libavdevice/version_major.h>
// #include <libavfilter/avfilter.h>
// #include <libavfilter/buffersink.h>
// #include <libavfilter/buffersrc.h>
// #include <libavfilter/version.h>
// #include <libavfilter/version_major.h>
// #include <libavformat/avformat.h>
// #include <libavformat/avio.h>
// #include <libavformat/version.h>
// #include <libavformat/version_major.h>
// #include <libavutil/avutil.h>
// #include <libavutil/buffer.h>
// #include <libavutil/channel_layout.h>
// #include <libavutil/dict.h>
// #include <libavutil/error.h>
// #include <libavutil/frame.h>
// #include <libavutil/hwcontext.h>
// #include <libavutil/log.h>
// #include <libavutil/mathematics.h>
// #include <libavutil/mem.h>
// #include <libavutil/opt.h>
// #include <libavutil/pixfmt.h>
// #include <libavutil/rational.h>
// #include <libavutil/samplefmt.h>
// #include <libavutil/version.h>
// #include <libpostproc/version.h>
// #include <libpostproc/version_major.h>
// #include <libswresample/version.h>
// #include <libswresample/version_major.h>
// #include <libswscale/version.h>
// #include <libswscale/version_major.h>
import "C"

// --- Struct RcOverride ---

// RcOverride wraps RcOverride.
type RcOverride struct {
	ptr *C.RcOverride
}

func (s *RcOverride) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToRcOverrideArray(ptr unsafe.Pointer) *Array[*RcOverride] {
	if ptr == nil {
		return nil
	}

	return &Array[*RcOverride]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *RcOverride {
			ptr := (**C.RcOverride)(pointer)
			value := *ptr
			var valueMapped *RcOverride
			if value != nil {
				valueMapped = &RcOverride{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *RcOverride) {
			ptr := (**C.RcOverride)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// StartFrame gets the start_frame field.
func (s *RcOverride) StartFrame() int {
	value := s.ptr.start_frame
	return int(value)
}

// SetStartFrame sets the start_frame field.
func (s *RcOverride) SetStartFrame(value int) {
	s.ptr.start_frame = (C.int)(value)
}

// EndFrame gets the end_frame field.
func (s *RcOverride) EndFrame() int {
	value := s.ptr.end_frame
	return int(value)
}

// SetEndFrame sets the end_frame field.
func (s *RcOverride) SetEndFrame(value int) {
	s.ptr.end_frame = (C.int)(value)
}

// Qscale gets the qscale field.
//
//	If this is 0 then quality_factor will be used instead.
func (s *RcOverride) Qscale() int {
	value := s.ptr.qscale
	return int(value)
}

// SetQscale sets the qscale field.
//
//	If this is 0 then quality_factor will be used instead.
func (s *RcOverride) SetQscale(value int) {
	s.ptr.qscale = (C.int)(value)
}

// QualityFactor gets the quality_factor field.
func (s *RcOverride) QualityFactor() float32 {
	value := s.ptr.quality_factor
	return float32(value)
}

// SetQualityFactor sets the quality_factor field.
func (s *RcOverride) SetQualityFactor(value float32) {
	s.ptr.quality_factor = (C.float)(value)
}

// --- Struct AVCodecContext ---

// AVCodecContext wraps AVCodecContext.
/*
  main external API structure.
  New fields can be added to the end with minor version bumps.
  Removal, reordering and changes to existing fields require a major
  version bump.
  You can use AVOptions (av_opt* / av_set/get*()) to access these fields from user
  applications.
  The name string for AVOptions options matches the associated command line
  parameter name and can be found in libavcodec/options_table.h
  The AVOption/command line parameter names differ in some cases from the C
  structure field names for historic reasons or brevity.
  sizeof(AVCodecContext) must not be used outside libav*.
*/
type AVCodecContext struct {
	ptr *C.AVCodecContext
}

func (s *AVCodecContext) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVCodecContextArray(ptr unsafe.Pointer) *Array[*AVCodecContext] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVCodecContext]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVCodecContext {
			ptr := (**C.AVCodecContext)(pointer)
			value := *ptr
			var valueMapped *AVCodecContext
			if value != nil {
				valueMapped = &AVCodecContext{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVCodecContext) {
			ptr := (**C.AVCodecContext)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// AvClass gets the av_class field.
/*
  information on struct for av_log
  - set by avcodec_alloc_context3
*/
func (s *AVCodecContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
/*
  information on struct for av_log
  - set by avcodec_alloc_context3
*/
func (s *AVCodecContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// LogLevelOffset gets the log_level_offset field.
func (s *AVCodecContext) LogLevelOffset() int {
	value := s.ptr.log_level_offset
	return int(value)
}

// SetLogLevelOffset sets the log_level_offset field.
func (s *AVCodecContext) SetLogLevelOffset(value int) {
	s.ptr.log_level_offset = (C.int)(value)
}

// CodecType gets the codec_type field.
//
//	see AVMEDIA_TYPE_xxx
func (s *AVCodecContext) CodecType() AVMediaType {
	value := s.ptr.codec_type
	return AVMediaType(value)
}

// SetCodecType sets the codec_type field.
//
//	see AVMEDIA_TYPE_xxx
func (s *AVCodecContext) SetCodecType(value AVMediaType) {
	s.ptr.codec_type = (C.enum_AVMediaType)(value)
}

// Codec gets the codec field.
func (s *AVCodecContext) Codec() *AVCodec {
	value := s.ptr.codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

// SetCodec sets the codec field.
func (s *AVCodecContext) SetCodec(value *AVCodec) {
	if value != nil {
		s.ptr.codec = value.ptr
	} else {
		s.ptr.codec = nil
	}
}

// CodecId gets the codec_id field.
//
//	see AV_CODEC_ID_xxx
func (s *AVCodecContext) CodecId() AVCodecID {
	value := s.ptr.codec_id
	return AVCodecID(value)
}

// SetCodecId sets the codec_id field.
//
//	see AV_CODEC_ID_xxx
func (s *AVCodecContext) SetCodecId(value AVCodecID) {
	s.ptr.codec_id = (C.enum_AVCodecID)(value)
}

// CodecTag gets the codec_tag field.
/*
  fourcc (LSB first, so "ABCD" -> ('D'<<24) + ('C'<<16) + ('B'<<8) + 'A').
  This is used to work around some encoder bugs.
  A demuxer should set this to what is stored in the field used to identify the codec.
  If there are multiple such fields in a container then the demuxer should choose the one
  which maximizes the information about the used codec.
  If the codec tag field in a container is larger than 32 bits then the demuxer should
  remap the longer ID to 32 bits with a table or other structure. Alternatively a new
  extra_codec_tag + size could be added but for this a clear advantage must be demonstrated
  first.
  - encoding: Set by user, if not then the default based on codec_id will be used.
  - decoding: Set by user, will be converted to uppercase by libavcodec during init.
*/
func (s *AVCodecContext) CodecTag() uint {
	value := s.ptr.codec_tag
	return uint(value)
}

// SetCodecTag sets the codec_tag field.
/*
  fourcc (LSB first, so "ABCD" -> ('D'<<24) + ('C'<<16) + ('B'<<8) + 'A').
  This is used to work around some encoder bugs.
  A demuxer should set this to what is stored in the field used to identify the codec.
  If there are multiple such fields in a container then the demuxer should choose the one
  which maximizes the information about the used codec.
  If the codec tag field in a container is larger than 32 bits then the demuxer should
  remap the longer ID to 32 bits with a table or other structure. Alternatively a new
  extra_codec_tag + size could be added but for this a clear advantage must be demonstrated
  first.
  - encoding: Set by user, if not then the default based on codec_id will be used.
  - decoding: Set by user, will be converted to uppercase by libavcodec during init.
*/
func (s *AVCodecContext) SetCodecTag(value uint) {
	s.ptr.codec_tag = (C.uint)(value)
}

// PrivData gets the priv_data field.
func (s *AVCodecContext) PrivData() unsafe.Pointer {
	value := s.ptr.priv_data
	return value
}

// SetPrivData sets the priv_data field.
func (s *AVCodecContext) SetPrivData(value unsafe.Pointer) {
	s.ptr.priv_data = value
}

// internal skipped due to ptr to ignored type

// Opaque gets the opaque field.
/*
  Private data of the user, can be used to carry app specific stuff.
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
/*
  Private data of the user, can be used to carry app specific stuff.
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// BitRate gets the bit_rate field.
/*
  the average bitrate
  - encoding: Set by user; unused for constant quantizer encoding.
  - decoding: Set by user, may be overwritten by libavcodec
              if this info is available in the stream
*/
func (s *AVCodecContext) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}

// SetBitRate sets the bit_rate field.
/*
  the average bitrate
  - encoding: Set by user; unused for constant quantizer encoding.
  - decoding: Set by user, may be overwritten by libavcodec
              if this info is available in the stream
*/
func (s *AVCodecContext) SetBitRate(value int64) {
	s.ptr.bit_rate = (C.int64_t)(value)
}

// BitRateTolerance gets the bit_rate_tolerance field.
/*
  number of bits the bitstream is allowed to diverge from the reference.
            the reference can be CBR (for CBR pass1) or VBR (for pass2)
  - encoding: Set by user; unused for constant quantizer encoding.
  - decoding: unused
*/
func (s *AVCodecContext) BitRateTolerance() int {
	value := s.ptr.bit_rate_tolerance
	return int(value)
}

// SetBitRateTolerance sets the bit_rate_tolerance field.
/*
  number of bits the bitstream is allowed to diverge from the reference.
            the reference can be CBR (for CBR pass1) or VBR (for pass2)
  - encoding: Set by user; unused for constant quantizer encoding.
  - decoding: unused
*/
func (s *AVCodecContext) SetBitRateTolerance(value int) {
	s.ptr.bit_rate_tolerance = (C.int)(value)
}

// GlobalQuality gets the global_quality field.
/*
  Global quality for codecs which cannot change it per frame.
  This should be proportional to MPEG-1/2/4 qscale.
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) GlobalQuality() int {
	value := s.ptr.global_quality
	return int(value)
}

// SetGlobalQuality sets the global_quality field.
/*
  Global quality for codecs which cannot change it per frame.
  This should be proportional to MPEG-1/2/4 qscale.
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetGlobalQuality(value int) {
	s.ptr.global_quality = (C.int)(value)
}

// CompressionLevel gets the compression_level field.
/*
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) CompressionLevel() int {
	value := s.ptr.compression_level
	return int(value)
}

// SetCompressionLevel sets the compression_level field.
/*
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetCompressionLevel(value int) {
	s.ptr.compression_level = (C.int)(value)
}

// Flags gets the flags field.
/*
  AV_CODEC_FLAG_*.
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
/*
  AV_CODEC_FLAG_*.
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// Flags2 gets the flags2 field.
/*
  AV_CODEC_FLAG2_*
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) Flags2() int {
	value := s.ptr.flags2
	return int(value)
}

// SetFlags2 sets the flags2 field.
/*
  AV_CODEC_FLAG2_*
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetFlags2(value int) {
	s.ptr.flags2 = (C.int)(value)
}

// Extradata gets the extradata field.
/*
  some codecs need / can use extradata like Huffman tables.
  MJPEG: Huffman tables
  rv10: additional flags
  MPEG-4: global headers (they can be in the bitstream or here)
  The allocated memory should be AV_INPUT_BUFFER_PADDING_SIZE bytes larger
  than extradata_size to avoid problems if it is read with the bitstream reader.
  The bytewise contents of extradata must not depend on the architecture or CPU endianness.
  Must be allocated with the av_malloc() family of functions.
  - encoding: Set/allocated/freed by libavcodec.
  - decoding: Set/allocated/freed by user.
*/
func (s *AVCodecContext) Extradata() unsafe.Pointer {
	value := s.ptr.extradata
	return unsafe.Pointer(value)
}

// SetExtradata sets the extradata field.
/*
  some codecs need / can use extradata like Huffman tables.
  MJPEG: Huffman tables
  rv10: additional flags
  MPEG-4: global headers (they can be in the bitstream or here)
  The allocated memory should be AV_INPUT_BUFFER_PADDING_SIZE bytes larger
  than extradata_size to avoid problems if it is read with the bitstream reader.
  The bytewise contents of extradata must not depend on the architecture or CPU endianness.
  Must be allocated with the av_malloc() family of functions.
  - encoding: Set/allocated/freed by libavcodec.
  - decoding: Set/allocated/freed by user.
*/
func (s *AVCodecContext) SetExtradata(value unsafe.Pointer) {
	s.ptr.extradata = (*C.uint8_t)(value)
}

// ExtradataSize gets the extradata_size field.
func (s *AVCodecContext) ExtradataSize() int {
	value := s.ptr.extradata_size
	return int(value)
}

// SetExtradataSize sets the extradata_size field.
func (s *AVCodecContext) SetExtradataSize(value int) {
	s.ptr.extradata_size = (C.int)(value)
}

// TimeBase gets the time_base field.
/*
  This is the fundamental unit of time (in seconds) in terms
  of which frame timestamps are represented. For fixed-fps content,
  timebase should be 1/framerate and timestamp increments should be
  identically 1.
  This often, but not always is the inverse of the frame rate or field rate
  for video. 1/time_base is not the average frame rate if the frame rate is not
  constant.

  Like containers, elementary streams also can store timestamps, 1/time_base
  is the unit in which these timestamps are specified.
  As example of such codec time base see ISO/IEC 14496-2:2001(E)
  vop_time_increment_resolution and fixed_vop_rate
  (fixed_vop_rate == 0 implies that it is different from the framerate)

  - encoding: MUST be set by user.
  - decoding: unused.
*/
func (s *AVCodecContext) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
/*
  This is the fundamental unit of time (in seconds) in terms
  of which frame timestamps are represented. For fixed-fps content,
  timebase should be 1/framerate and timestamp increments should be
  identically 1.
  This often, but not always is the inverse of the frame rate or field rate
  for video. 1/time_base is not the average frame rate if the frame rate is not
  constant.

  Like containers, elementary streams also can store timestamps, 1/time_base
  is the unit in which these timestamps are specified.
  As example of such codec time base see ISO/IEC 14496-2:2001(E)
  vop_time_increment_resolution and fixed_vop_rate
  (fixed_vop_rate == 0 implies that it is different from the framerate)

  - encoding: MUST be set by user.
  - decoding: unused.
*/
func (s *AVCodecContext) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// TicksPerFrame gets the ticks_per_frame field.
/*
  For some codecs, the time base is closer to the field rate than the frame rate.
  Most notably, H.264 and MPEG-2 specify time_base as half of frame duration
  if no telecine is used ...

  Set to time_base ticks per frame. Default 1, e.g., H.264/MPEG-2 set it to 2.

  @deprecated
  - decoding: Use AVCodecDescriptor.props & AV_CODEC_PROP_FIELDS
  - encoding: Set AVCodecContext.framerate instead
*/
func (s *AVCodecContext) TicksPerFrame() int {
	value := s.ptr.ticks_per_frame
	return int(value)
}

// SetTicksPerFrame sets the ticks_per_frame field.
/*
  For some codecs, the time base is closer to the field rate than the frame rate.
  Most notably, H.264 and MPEG-2 specify time_base as half of frame duration
  if no telecine is used ...

  Set to time_base ticks per frame. Default 1, e.g., H.264/MPEG-2 set it to 2.

  @deprecated
  - decoding: Use AVCodecDescriptor.props & AV_CODEC_PROP_FIELDS
  - encoding: Set AVCodecContext.framerate instead
*/
func (s *AVCodecContext) SetTicksPerFrame(value int) {
	s.ptr.ticks_per_frame = (C.int)(value)
}

// Delay gets the delay field.
/*
  Codec delay.

  Encoding: Number of frames delay there will be from the encoder input to
            the decoder output. (we assume the decoder matches the spec)
  Decoding: Number of frames delay in addition to what a standard decoder
            as specified in the spec would produce.

  Video:
    Number of frames the decoded output will be delayed relative to the
    encoded input.

  Audio:
    For encoding, this field is unused (see initial_padding).

    For decoding, this is the number of samples the decoder needs to
    output before the decoder's output is valid. When seeking, you should
    start decoding this many samples prior to your desired seek point.

  - encoding: Set by libavcodec.
  - decoding: Set by libavcodec.
*/
func (s *AVCodecContext) Delay() int {
	value := s.ptr.delay
	return int(value)
}

// SetDelay sets the delay field.
/*
  Codec delay.

  Encoding: Number of frames delay there will be from the encoder input to
            the decoder output. (we assume the decoder matches the spec)
  Decoding: Number of frames delay in addition to what a standard decoder
            as specified in the spec would produce.

  Video:
    Number of frames the decoded output will be delayed relative to the
    encoded input.

  Audio:
    For encoding, this field is unused (see initial_padding).

    For decoding, this is the number of samples the decoder needs to
    output before the decoder's output is valid. When seeking, you should
    start decoding this many samples prior to your desired seek point.

  - encoding: Set by libavcodec.
  - decoding: Set by libavcodec.
*/
func (s *AVCodecContext) SetDelay(value int) {
	s.ptr.delay = (C.int)(value)
}

// Width gets the width field.
/*
  picture width / height.

  @note Those fields may not match the values of the last
  AVFrame output by avcodec_receive_frame() due frame
  reordering.

  - encoding: MUST be set by user.
  - decoding: May be set by the user before opening the decoder if known e.g.
              from the container. Some decoders will require the dimensions
              to be set by the caller. During decoding, the decoder may
              overwrite those values as required while parsing the data.
*/
func (s *AVCodecContext) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
/*
  picture width / height.

  @note Those fields may not match the values of the last
  AVFrame output by avcodec_receive_frame() due frame
  reordering.

  - encoding: MUST be set by user.
  - decoding: May be set by the user before opening the decoder if known e.g.
              from the container. Some decoders will require the dimensions
              to be set by the caller. During decoding, the decoder may
              overwrite those values as required while parsing the data.
*/
func (s *AVCodecContext) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

// Height gets the height field.
/*
  picture width / height.

  @note Those fields may not match the values of the last
  AVFrame output by avcodec_receive_frame() due frame
  reordering.

  - encoding: MUST be set by user.
  - decoding: May be set by the user before opening the decoder if known e.g.
              from the container. Some decoders will require the dimensions
              to be set by the caller. During decoding, the decoder may
              overwrite those values as required while parsing the data.
*/
func (s *AVCodecContext) Height() int {
	value := s.ptr.height
	return int(value)
}

// SetHeight sets the height field.
/*
  picture width / height.

  @note Those fields may not match the values of the last
  AVFrame output by avcodec_receive_frame() due frame
  reordering.

  - encoding: MUST be set by user.
  - decoding: May be set by the user before opening the decoder if known e.g.
              from the container. Some decoders will require the dimensions
              to be set by the caller. During decoding, the decoder may
              overwrite those values as required while parsing the data.
*/
func (s *AVCodecContext) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

// CodedWidth gets the coded_width field.
/*
  Bitstream width / height, may be different from width/height e.g. when
  the decoded frame is cropped before being output or lowres is enabled.

  @note Those field may not match the value of the last
  AVFrame output by avcodec_receive_frame() due frame
  reordering.

  - encoding: unused
  - decoding: May be set by the user before opening the decoder if known
              e.g. from the container. During decoding, the decoder may
              overwrite those values as required while parsing the data.
*/
func (s *AVCodecContext) CodedWidth() int {
	value := s.ptr.coded_width
	return int(value)
}

// SetCodedWidth sets the coded_width field.
/*
  Bitstream width / height, may be different from width/height e.g. when
  the decoded frame is cropped before being output or lowres is enabled.

  @note Those field may not match the value of the last
  AVFrame output by avcodec_receive_frame() due frame
  reordering.

  - encoding: unused
  - decoding: May be set by the user before opening the decoder if known
              e.g. from the container. During decoding, the decoder may
              overwrite those values as required while parsing the data.
*/
func (s *AVCodecContext) SetCodedWidth(value int) {
	s.ptr.coded_width = (C.int)(value)
}

// CodedHeight gets the coded_height field.
/*
  Bitstream width / height, may be different from width/height e.g. when
  the decoded frame is cropped before being output or lowres is enabled.

  @note Those field may not match the value of the last
  AVFrame output by avcodec_receive_frame() due frame
  reordering.

  - encoding: unused
  - decoding: May be set by the user before opening the decoder if known
              e.g. from the container. During decoding, the decoder may
              overwrite those values as required while parsing the data.
*/
func (s *AVCodecContext) CodedHeight() int {
	value := s.ptr.coded_height
	return int(value)
}

// SetCodedHeight sets the coded_height field.
/*
  Bitstream width / height, may be different from width/height e.g. when
  the decoded frame is cropped before being output or lowres is enabled.

  @note Those field may not match the value of the last
  AVFrame output by avcodec_receive_frame() due frame
  reordering.

  - encoding: unused
  - decoding: May be set by the user before opening the decoder if known
              e.g. from the container. During decoding, the decoder may
              overwrite those values as required while parsing the data.
*/
func (s *AVCodecContext) SetCodedHeight(value int) {
	s.ptr.coded_height = (C.int)(value)
}

// GopSize gets the gop_size field.
/*
  the number of pictures in a group of pictures, or 0 for intra_only
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) GopSize() int {
	value := s.ptr.gop_size
	return int(value)
}

// SetGopSize sets the gop_size field.
/*
  the number of pictures in a group of pictures, or 0 for intra_only
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetGopSize(value int) {
	s.ptr.gop_size = (C.int)(value)
}

// PixFmt gets the pix_fmt field.
/*
  Pixel format, see AV_PIX_FMT_xxx.
  May be set by the demuxer if known from headers.
  May be overridden by the decoder if it knows better.

  @note This field may not match the value of the last
  AVFrame output by avcodec_receive_frame() due frame
  reordering.

  - encoding: Set by user.
  - decoding: Set by user if known, overridden by libavcodec while
              parsing the data.
*/
func (s *AVCodecContext) PixFmt() AVPixelFormat {
	value := s.ptr.pix_fmt
	return AVPixelFormat(value)
}

// SetPixFmt sets the pix_fmt field.
/*
  Pixel format, see AV_PIX_FMT_xxx.
  May be set by the demuxer if known from headers.
  May be overridden by the decoder if it knows better.

  @note This field may not match the value of the last
  AVFrame output by avcodec_receive_frame() due frame
  reordering.

  - encoding: Set by user.
  - decoding: Set by user if known, overridden by libavcodec while
              parsing the data.
*/
func (s *AVCodecContext) SetPixFmt(value AVPixelFormat) {
	s.ptr.pix_fmt = (C.enum_AVPixelFormat)(value)
}

// draw_horiz_band skipped due to func ptr

// get_format skipped due to func ptr

// MaxBFrames gets the max_b_frames field.
/*
  maximum number of B-frames between non-B-frames
  Note: The output will be delayed by max_b_frames+1 relative to the input.
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) MaxBFrames() int {
	value := s.ptr.max_b_frames
	return int(value)
}

// SetMaxBFrames sets the max_b_frames field.
/*
  maximum number of B-frames between non-B-frames
  Note: The output will be delayed by max_b_frames+1 relative to the input.
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMaxBFrames(value int) {
	s.ptr.max_b_frames = (C.int)(value)
}

// BQuantFactor gets the b_quant_factor field.
/*
  qscale factor between IP and B-frames
  If > 0 then the last P-frame quantizer will be used (q= lastp_q*factor+offset).
  If < 0 then normal ratecontrol will be done (q= -normal_q*factor+offset).
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) BQuantFactor() float32 {
	value := s.ptr.b_quant_factor
	return float32(value)
}

// SetBQuantFactor sets the b_quant_factor field.
/*
  qscale factor between IP and B-frames
  If > 0 then the last P-frame quantizer will be used (q= lastp_q*factor+offset).
  If < 0 then normal ratecontrol will be done (q= -normal_q*factor+offset).
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetBQuantFactor(value float32) {
	s.ptr.b_quant_factor = (C.float)(value)
}

// BQuantOffset gets the b_quant_offset field.
/*
  qscale offset between IP and B-frames
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) BQuantOffset() float32 {
	value := s.ptr.b_quant_offset
	return float32(value)
}

// SetBQuantOffset sets the b_quant_offset field.
/*
  qscale offset between IP and B-frames
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetBQuantOffset(value float32) {
	s.ptr.b_quant_offset = (C.float)(value)
}

// HasBFrames gets the has_b_frames field.
/*
  Size of the frame reordering buffer in the decoder.
  For MPEG-2 it is 1 IPB or 0 low delay IP.
  - encoding: Set by libavcodec.
  - decoding: Set by libavcodec.
*/
func (s *AVCodecContext) HasBFrames() int {
	value := s.ptr.has_b_frames
	return int(value)
}

// SetHasBFrames sets the has_b_frames field.
/*
  Size of the frame reordering buffer in the decoder.
  For MPEG-2 it is 1 IPB or 0 low delay IP.
  - encoding: Set by libavcodec.
  - decoding: Set by libavcodec.
*/
func (s *AVCodecContext) SetHasBFrames(value int) {
	s.ptr.has_b_frames = (C.int)(value)
}

// IQuantFactor gets the i_quant_factor field.
/*
  qscale factor between P- and I-frames
  If > 0 then the last P-frame quantizer will be used (q = lastp_q * factor + offset).
  If < 0 then normal ratecontrol will be done (q= -normal_q*factor+offset).
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) IQuantFactor() float32 {
	value := s.ptr.i_quant_factor
	return float32(value)
}

// SetIQuantFactor sets the i_quant_factor field.
/*
  qscale factor between P- and I-frames
  If > 0 then the last P-frame quantizer will be used (q = lastp_q * factor + offset).
  If < 0 then normal ratecontrol will be done (q= -normal_q*factor+offset).
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetIQuantFactor(value float32) {
	s.ptr.i_quant_factor = (C.float)(value)
}

// IQuantOffset gets the i_quant_offset field.
/*
  qscale offset between P and I-frames
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) IQuantOffset() float32 {
	value := s.ptr.i_quant_offset
	return float32(value)
}

// SetIQuantOffset sets the i_quant_offset field.
/*
  qscale offset between P and I-frames
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetIQuantOffset(value float32) {
	s.ptr.i_quant_offset = (C.float)(value)
}

// LumiMasking gets the lumi_masking field.
/*
  luminance masking (0-> disabled)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) LumiMasking() float32 {
	value := s.ptr.lumi_masking
	return float32(value)
}

// SetLumiMasking sets the lumi_masking field.
/*
  luminance masking (0-> disabled)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetLumiMasking(value float32) {
	s.ptr.lumi_masking = (C.float)(value)
}

// TemporalCplxMasking gets the temporal_cplx_masking field.
/*
  temporary complexity masking (0-> disabled)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) TemporalCplxMasking() float32 {
	value := s.ptr.temporal_cplx_masking
	return float32(value)
}

// SetTemporalCplxMasking sets the temporal_cplx_masking field.
/*
  temporary complexity masking (0-> disabled)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetTemporalCplxMasking(value float32) {
	s.ptr.temporal_cplx_masking = (C.float)(value)
}

// SpatialCplxMasking gets the spatial_cplx_masking field.
/*
  spatial complexity masking (0-> disabled)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SpatialCplxMasking() float32 {
	value := s.ptr.spatial_cplx_masking
	return float32(value)
}

// SetSpatialCplxMasking sets the spatial_cplx_masking field.
/*
  spatial complexity masking (0-> disabled)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetSpatialCplxMasking(value float32) {
	s.ptr.spatial_cplx_masking = (C.float)(value)
}

// PMasking gets the p_masking field.
/*
  p block masking (0-> disabled)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) PMasking() float32 {
	value := s.ptr.p_masking
	return float32(value)
}

// SetPMasking sets the p_masking field.
/*
  p block masking (0-> disabled)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetPMasking(value float32) {
	s.ptr.p_masking = (C.float)(value)
}

// DarkMasking gets the dark_masking field.
/*
  darkness masking (0-> disabled)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) DarkMasking() float32 {
	value := s.ptr.dark_masking
	return float32(value)
}

// SetDarkMasking sets the dark_masking field.
/*
  darkness masking (0-> disabled)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetDarkMasking(value float32) {
	s.ptr.dark_masking = (C.float)(value)
}

// SliceCount gets the slice_count field.
/*
  slice count
  - encoding: Set by libavcodec.
  - decoding: Set by user (or 0).
*/
func (s *AVCodecContext) SliceCount() int {
	value := s.ptr.slice_count
	return int(value)
}

// SetSliceCount sets the slice_count field.
/*
  slice count
  - encoding: Set by libavcodec.
  - decoding: Set by user (or 0).
*/
func (s *AVCodecContext) SetSliceCount(value int) {
	s.ptr.slice_count = (C.int)(value)
}

// slice_offset skipped due to prim ptr

// SampleAspectRatio gets the sample_aspect_ratio field.
/*
  sample aspect ratio (0 if unknown)
  That is the width of a pixel divided by the height of the pixel.
  Numerator and denominator must be relatively prime and smaller than 256 for some video standards.
  - encoding: Set by user.
  - decoding: Set by libavcodec.
*/
func (s *AVCodecContext) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
/*
  sample aspect ratio (0 if unknown)
  That is the width of a pixel divided by the height of the pixel.
  Numerator and denominator must be relatively prime and smaller than 256 for some video standards.
  - encoding: Set by user.
  - decoding: Set by libavcodec.
*/
func (s *AVCodecContext) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

// MeCmp gets the me_cmp field.
/*
  motion estimation comparison function
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) MeCmp() int {
	value := s.ptr.me_cmp
	return int(value)
}

// SetMeCmp sets the me_cmp field.
/*
  motion estimation comparison function
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMeCmp(value int) {
	s.ptr.me_cmp = (C.int)(value)
}

// MeSubCmp gets the me_sub_cmp field.
/*
  subpixel motion estimation comparison function
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) MeSubCmp() int {
	value := s.ptr.me_sub_cmp
	return int(value)
}

// SetMeSubCmp sets the me_sub_cmp field.
/*
  subpixel motion estimation comparison function
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMeSubCmp(value int) {
	s.ptr.me_sub_cmp = (C.int)(value)
}

// MbCmp gets the mb_cmp field.
/*
  macroblock comparison function (not supported yet)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) MbCmp() int {
	value := s.ptr.mb_cmp
	return int(value)
}

// SetMbCmp sets the mb_cmp field.
/*
  macroblock comparison function (not supported yet)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMbCmp(value int) {
	s.ptr.mb_cmp = (C.int)(value)
}

// IldctCmp gets the ildct_cmp field.
/*
  interlaced DCT comparison function
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) IldctCmp() int {
	value := s.ptr.ildct_cmp
	return int(value)
}

// SetIldctCmp sets the ildct_cmp field.
/*
  interlaced DCT comparison function
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetIldctCmp(value int) {
	s.ptr.ildct_cmp = (C.int)(value)
}

// DiaSize gets the dia_size field.
/*
  ME diamond size & shape
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) DiaSize() int {
	value := s.ptr.dia_size
	return int(value)
}

// SetDiaSize sets the dia_size field.
/*
  ME diamond size & shape
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetDiaSize(value int) {
	s.ptr.dia_size = (C.int)(value)
}

// LastPredictorCount gets the last_predictor_count field.
/*
  amount of previous MV predictors (2a+1 x 2a+1 square)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) LastPredictorCount() int {
	value := s.ptr.last_predictor_count
	return int(value)
}

// SetLastPredictorCount sets the last_predictor_count field.
/*
  amount of previous MV predictors (2a+1 x 2a+1 square)
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetLastPredictorCount(value int) {
	s.ptr.last_predictor_count = (C.int)(value)
}

// MePreCmp gets the me_pre_cmp field.
/*
  motion estimation prepass comparison function
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) MePreCmp() int {
	value := s.ptr.me_pre_cmp
	return int(value)
}

// SetMePreCmp sets the me_pre_cmp field.
/*
  motion estimation prepass comparison function
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMePreCmp(value int) {
	s.ptr.me_pre_cmp = (C.int)(value)
}

// PreDiaSize gets the pre_dia_size field.
/*
  ME prepass diamond size & shape
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) PreDiaSize() int {
	value := s.ptr.pre_dia_size
	return int(value)
}

// SetPreDiaSize sets the pre_dia_size field.
/*
  ME prepass diamond size & shape
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetPreDiaSize(value int) {
	s.ptr.pre_dia_size = (C.int)(value)
}

// MeSubpelQuality gets the me_subpel_quality field.
/*
  subpel ME quality
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) MeSubpelQuality() int {
	value := s.ptr.me_subpel_quality
	return int(value)
}

// SetMeSubpelQuality sets the me_subpel_quality field.
/*
  subpel ME quality
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMeSubpelQuality(value int) {
	s.ptr.me_subpel_quality = (C.int)(value)
}

// MeRange gets the me_range field.
/*
  maximum motion estimation search range in subpel units
  If 0 then no limit.

  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) MeRange() int {
	value := s.ptr.me_range
	return int(value)
}

// SetMeRange sets the me_range field.
/*
  maximum motion estimation search range in subpel units
  If 0 then no limit.

  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMeRange(value int) {
	s.ptr.me_range = (C.int)(value)
}

// SliceFlags gets the slice_flags field.
/*
  slice flags
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SliceFlags() int {
	value := s.ptr.slice_flags
	return int(value)
}

// SetSliceFlags sets the slice_flags field.
/*
  slice flags
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetSliceFlags(value int) {
	s.ptr.slice_flags = (C.int)(value)
}

// MbDecision gets the mb_decision field.
/*
  macroblock decision mode
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) MbDecision() int {
	value := s.ptr.mb_decision
	return int(value)
}

// SetMbDecision sets the mb_decision field.
/*
  macroblock decision mode
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMbDecision(value int) {
	s.ptr.mb_decision = (C.int)(value)
}

// intra_matrix skipped due to prim ptr

// inter_matrix skipped due to prim ptr

// IntraDcPrecision gets the intra_dc_precision field.
/*
  precision of the intra DC coefficient - 8
  - encoding: Set by user.
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) IntraDcPrecision() int {
	value := s.ptr.intra_dc_precision
	return int(value)
}

// SetIntraDcPrecision sets the intra_dc_precision field.
/*
  precision of the intra DC coefficient - 8
  - encoding: Set by user.
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) SetIntraDcPrecision(value int) {
	s.ptr.intra_dc_precision = (C.int)(value)
}

// SkipTop gets the skip_top field.
/*
  Number of macroblock rows at the top which are skipped.
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SkipTop() int {
	value := s.ptr.skip_top
	return int(value)
}

// SetSkipTop sets the skip_top field.
/*
  Number of macroblock rows at the top which are skipped.
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetSkipTop(value int) {
	s.ptr.skip_top = (C.int)(value)
}

// SkipBottom gets the skip_bottom field.
/*
  Number of macroblock rows at the bottom which are skipped.
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SkipBottom() int {
	value := s.ptr.skip_bottom
	return int(value)
}

// SetSkipBottom sets the skip_bottom field.
/*
  Number of macroblock rows at the bottom which are skipped.
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetSkipBottom(value int) {
	s.ptr.skip_bottom = (C.int)(value)
}

// MbLmin gets the mb_lmin field.
/*
  minimum MB Lagrange multiplier
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) MbLmin() int {
	value := s.ptr.mb_lmin
	return int(value)
}

// SetMbLmin sets the mb_lmin field.
/*
  minimum MB Lagrange multiplier
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMbLmin(value int) {
	s.ptr.mb_lmin = (C.int)(value)
}

// MbLmax gets the mb_lmax field.
/*
  maximum MB Lagrange multiplier
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) MbLmax() int {
	value := s.ptr.mb_lmax
	return int(value)
}

// SetMbLmax sets the mb_lmax field.
/*
  maximum MB Lagrange multiplier
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMbLmax(value int) {
	s.ptr.mb_lmax = (C.int)(value)
}

// BidirRefine gets the bidir_refine field.
/*
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) BidirRefine() int {
	value := s.ptr.bidir_refine
	return int(value)
}

// SetBidirRefine sets the bidir_refine field.
/*
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetBidirRefine(value int) {
	s.ptr.bidir_refine = (C.int)(value)
}

// KeyintMin gets the keyint_min field.
/*
  minimum GOP size
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) KeyintMin() int {
	value := s.ptr.keyint_min
	return int(value)
}

// SetKeyintMin sets the keyint_min field.
/*
  minimum GOP size
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetKeyintMin(value int) {
	s.ptr.keyint_min = (C.int)(value)
}

// Refs gets the refs field.
/*
  number of reference frames
  - encoding: Set by user.
  - decoding: Set by lavc.
*/
func (s *AVCodecContext) Refs() int {
	value := s.ptr.refs
	return int(value)
}

// SetRefs sets the refs field.
/*
  number of reference frames
  - encoding: Set by user.
  - decoding: Set by lavc.
*/
func (s *AVCodecContext) SetRefs(value int) {
	s.ptr.refs = (C.int)(value)
}

// Mv0Threshold gets the mv0_threshold field.
/*
  Note: Value depends upon the compare function used for fullpel ME.
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) Mv0Threshold() int {
	value := s.ptr.mv0_threshold
	return int(value)
}

// SetMv0Threshold sets the mv0_threshold field.
/*
  Note: Value depends upon the compare function used for fullpel ME.
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMv0Threshold(value int) {
	s.ptr.mv0_threshold = (C.int)(value)
}

// ColorPrimaries gets the color_primaries field.
/*
  Chromaticity coordinates of the source primaries.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) ColorPrimaries() AVColorPrimaries {
	value := s.ptr.color_primaries
	return AVColorPrimaries(value)
}

// SetColorPrimaries sets the color_primaries field.
/*
  Chromaticity coordinates of the source primaries.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) SetColorPrimaries(value AVColorPrimaries) {
	s.ptr.color_primaries = (C.enum_AVColorPrimaries)(value)
}

// ColorTrc gets the color_trc field.
/*
  Color Transfer Characteristic.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) ColorTrc() AVColorTransferCharacteristic {
	value := s.ptr.color_trc
	return AVColorTransferCharacteristic(value)
}

// SetColorTrc sets the color_trc field.
/*
  Color Transfer Characteristic.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) SetColorTrc(value AVColorTransferCharacteristic) {
	s.ptr.color_trc = (C.enum_AVColorTransferCharacteristic)(value)
}

// Colorspace gets the colorspace field.
/*
  YUV colorspace type.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) Colorspace() AVColorSpace {
	value := s.ptr.colorspace
	return AVColorSpace(value)
}

// SetColorspace sets the colorspace field.
/*
  YUV colorspace type.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) SetColorspace(value AVColorSpace) {
	s.ptr.colorspace = (C.enum_AVColorSpace)(value)
}

// ColorRange gets the color_range field.
/*
  MPEG vs JPEG YUV range.
  - encoding: Set by user to override the default output color range value,
    If not specified, libavcodec sets the color range depending on the
    output format.
  - decoding: Set by libavcodec, can be set by the user to propagate the
    color range to components reading from the decoder context.
*/
func (s *AVCodecContext) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}

// SetColorRange sets the color_range field.
/*
  MPEG vs JPEG YUV range.
  - encoding: Set by user to override the default output color range value,
    If not specified, libavcodec sets the color range depending on the
    output format.
  - decoding: Set by libavcodec, can be set by the user to propagate the
    color range to components reading from the decoder context.
*/
func (s *AVCodecContext) SetColorRange(value AVColorRange) {
	s.ptr.color_range = (C.enum_AVColorRange)(value)
}

// ChromaSampleLocation gets the chroma_sample_location field.
/*
  This defines the location of chroma samples.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) ChromaSampleLocation() AVChromaLocation {
	value := s.ptr.chroma_sample_location
	return AVChromaLocation(value)
}

// SetChromaSampleLocation sets the chroma_sample_location field.
/*
  This defines the location of chroma samples.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) SetChromaSampleLocation(value AVChromaLocation) {
	s.ptr.chroma_sample_location = (C.enum_AVChromaLocation)(value)
}

// Slices gets the slices field.
/*
  Number of slices.
  Indicates number of picture subdivisions. Used for parallelized
  decoding.
  - encoding: Set by user
  - decoding: unused
*/
func (s *AVCodecContext) Slices() int {
	value := s.ptr.slices
	return int(value)
}

// SetSlices sets the slices field.
/*
  Number of slices.
  Indicates number of picture subdivisions. Used for parallelized
  decoding.
  - encoding: Set by user
  - decoding: unused
*/
func (s *AVCodecContext) SetSlices(value int) {
	s.ptr.slices = (C.int)(value)
}

// FieldOrder gets the field_order field.
/*
  Field order
  - encoding: set by libavcodec
  - decoding: Set by user.
*/
func (s *AVCodecContext) FieldOrder() AVFieldOrder {
	value := s.ptr.field_order
	return AVFieldOrder(value)
}

// SetFieldOrder sets the field_order field.
/*
  Field order
  - encoding: set by libavcodec
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetFieldOrder(value AVFieldOrder) {
	s.ptr.field_order = (C.enum_AVFieldOrder)(value)
}

// SampleRate gets the sample_rate field.
//
//	samples per second
func (s *AVCodecContext) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

// SetSampleRate sets the sample_rate field.
//
//	samples per second
func (s *AVCodecContext) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

// Channels gets the channels field.
/*
  number of audio channels
  @deprecated use ch_layout.nb_channels
*/
func (s *AVCodecContext) Channels() int {
	value := s.ptr.channels
	return int(value)
}

// SetChannels sets the channels field.
/*
  number of audio channels
  @deprecated use ch_layout.nb_channels
*/
func (s *AVCodecContext) SetChannels(value int) {
	s.ptr.channels = (C.int)(value)
}

// SampleFmt gets the sample_fmt field.
//
//	sample format
func (s *AVCodecContext) SampleFmt() AVSampleFormat {
	value := s.ptr.sample_fmt
	return AVSampleFormat(value)
}

// SetSampleFmt sets the sample_fmt field.
//
//	sample format
func (s *AVCodecContext) SetSampleFmt(value AVSampleFormat) {
	s.ptr.sample_fmt = (C.enum_AVSampleFormat)(value)
}

// FrameSize gets the frame_size field.
/*
  Number of samples per channel in an audio frame.

  - encoding: set by libavcodec in avcodec_open2(). Each submitted frame
    except the last must contain exactly frame_size samples per channel.
    May be 0 when the codec has AV_CODEC_CAP_VARIABLE_FRAME_SIZE set, then the
    frame size is not restricted.
  - decoding: may be set by some decoders to indicate constant frame size
*/
func (s *AVCodecContext) FrameSize() int {
	value := s.ptr.frame_size
	return int(value)
}

// SetFrameSize sets the frame_size field.
/*
  Number of samples per channel in an audio frame.

  - encoding: set by libavcodec in avcodec_open2(). Each submitted frame
    except the last must contain exactly frame_size samples per channel.
    May be 0 when the codec has AV_CODEC_CAP_VARIABLE_FRAME_SIZE set, then the
    frame size is not restricted.
  - decoding: may be set by some decoders to indicate constant frame size
*/
func (s *AVCodecContext) SetFrameSize(value int) {
	s.ptr.frame_size = (C.int)(value)
}

// FrameNumber gets the frame_number field.
/*
  Frame counter, set by libavcodec.

  - decoding: total number of frames returned from the decoder so far.
  - encoding: total number of frames passed to the encoder so far.

    @note the counter is not incremented if encoding/decoding resulted in
    an error.
    @deprecated use frame_num instead
*/
func (s *AVCodecContext) FrameNumber() int {
	value := s.ptr.frame_number
	return int(value)
}

// SetFrameNumber sets the frame_number field.
/*
  Frame counter, set by libavcodec.

  - decoding: total number of frames returned from the decoder so far.
  - encoding: total number of frames passed to the encoder so far.

    @note the counter is not incremented if encoding/decoding resulted in
    an error.
    @deprecated use frame_num instead
*/
func (s *AVCodecContext) SetFrameNumber(value int) {
	s.ptr.frame_number = (C.int)(value)
}

// BlockAlign gets the block_align field.
/*
  number of bytes per packet if constant and known or 0
  Used by some WAV based audio codecs.
*/
func (s *AVCodecContext) BlockAlign() int {
	value := s.ptr.block_align
	return int(value)
}

// SetBlockAlign sets the block_align field.
/*
  number of bytes per packet if constant and known or 0
  Used by some WAV based audio codecs.
*/
func (s *AVCodecContext) SetBlockAlign(value int) {
	s.ptr.block_align = (C.int)(value)
}

// Cutoff gets the cutoff field.
/*
  Audio cutoff bandwidth (0 means "automatic")
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) Cutoff() int {
	value := s.ptr.cutoff
	return int(value)
}

// SetCutoff sets the cutoff field.
/*
  Audio cutoff bandwidth (0 means "automatic")
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetCutoff(value int) {
	s.ptr.cutoff = (C.int)(value)
}

// ChannelLayout gets the channel_layout field.
/*
  Audio channel layout.
  - encoding: set by user.
  - decoding: set by user, may be overwritten by libavcodec.
  @deprecated use ch_layout
*/
func (s *AVCodecContext) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

// SetChannelLayout sets the channel_layout field.
/*
  Audio channel layout.
  - encoding: set by user.
  - decoding: set by user, may be overwritten by libavcodec.
  @deprecated use ch_layout
*/
func (s *AVCodecContext) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

// RequestChannelLayout gets the request_channel_layout field.
/*
  Request decoder to use this channel layout if it can (0 for default)
  - encoding: unused
  - decoding: Set by user.
  @deprecated use "downmix" codec private option
*/
func (s *AVCodecContext) RequestChannelLayout() uint64 {
	value := s.ptr.request_channel_layout
	return uint64(value)
}

// SetRequestChannelLayout sets the request_channel_layout field.
/*
  Request decoder to use this channel layout if it can (0 for default)
  - encoding: unused
  - decoding: Set by user.
  @deprecated use "downmix" codec private option
*/
func (s *AVCodecContext) SetRequestChannelLayout(value uint64) {
	s.ptr.request_channel_layout = (C.uint64_t)(value)
}

// AudioServiceType gets the audio_service_type field.
/*
  Type of service that the audio stream conveys.
  - encoding: Set by user.
  - decoding: Set by libavcodec.
*/
func (s *AVCodecContext) AudioServiceType() AVAudioServiceType {
	value := s.ptr.audio_service_type
	return AVAudioServiceType(value)
}

// SetAudioServiceType sets the audio_service_type field.
/*
  Type of service that the audio stream conveys.
  - encoding: Set by user.
  - decoding: Set by libavcodec.
*/
func (s *AVCodecContext) SetAudioServiceType(value AVAudioServiceType) {
	s.ptr.audio_service_type = (C.enum_AVAudioServiceType)(value)
}

// RequestSampleFmt gets the request_sample_fmt field.
/*
  desired sample format
  - encoding: Not used.
  - decoding: Set by user.
  Decoder will decode to this format if it can.
*/
func (s *AVCodecContext) RequestSampleFmt() AVSampleFormat {
	value := s.ptr.request_sample_fmt
	return AVSampleFormat(value)
}

// SetRequestSampleFmt sets the request_sample_fmt field.
/*
  desired sample format
  - encoding: Not used.
  - decoding: Set by user.
  Decoder will decode to this format if it can.
*/
func (s *AVCodecContext) SetRequestSampleFmt(value AVSampleFormat) {
	s.ptr.request_sample_fmt = (C.enum_AVSampleFormat)(value)
}

// get_buffer2 skipped due to func ptr

// Qcompress gets the qcompress field.
//
//	amount of qscale change between easy & hard scenes (0.0-1.0)
func (s *AVCodecContext) Qcompress() float32 {
	value := s.ptr.qcompress
	return float32(value)
}

// SetQcompress sets the qcompress field.
//
//	amount of qscale change between easy & hard scenes (0.0-1.0)
func (s *AVCodecContext) SetQcompress(value float32) {
	s.ptr.qcompress = (C.float)(value)
}

// Qblur gets the qblur field.
//
//	amount of qscale smoothing over time (0.0-1.0)
func (s *AVCodecContext) Qblur() float32 {
	value := s.ptr.qblur
	return float32(value)
}

// SetQblur sets the qblur field.
//
//	amount of qscale smoothing over time (0.0-1.0)
func (s *AVCodecContext) SetQblur(value float32) {
	s.ptr.qblur = (C.float)(value)
}

// Qmin gets the qmin field.
/*
  minimum quantizer
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) Qmin() int {
	value := s.ptr.qmin
	return int(value)
}

// SetQmin sets the qmin field.
/*
  minimum quantizer
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetQmin(value int) {
	s.ptr.qmin = (C.int)(value)
}

// Qmax gets the qmax field.
/*
  maximum quantizer
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) Qmax() int {
	value := s.ptr.qmax
	return int(value)
}

// SetQmax sets the qmax field.
/*
  maximum quantizer
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetQmax(value int) {
	s.ptr.qmax = (C.int)(value)
}

// MaxQdiff gets the max_qdiff field.
/*
  maximum quantizer difference between frames
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) MaxQdiff() int {
	value := s.ptr.max_qdiff
	return int(value)
}

// SetMaxQdiff sets the max_qdiff field.
/*
  maximum quantizer difference between frames
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetMaxQdiff(value int) {
	s.ptr.max_qdiff = (C.int)(value)
}

// RcBufferSize gets the rc_buffer_size field.
/*
  decoder bitstream buffer size
  - encoding: Set by user.
  - decoding: May be set by libavcodec.
*/
func (s *AVCodecContext) RcBufferSize() int {
	value := s.ptr.rc_buffer_size
	return int(value)
}

// SetRcBufferSize sets the rc_buffer_size field.
/*
  decoder bitstream buffer size
  - encoding: Set by user.
  - decoding: May be set by libavcodec.
*/
func (s *AVCodecContext) SetRcBufferSize(value int) {
	s.ptr.rc_buffer_size = (C.int)(value)
}

// RcOverrideCount gets the rc_override_count field.
/*
  ratecontrol override, see RcOverride
  - encoding: Allocated/set/freed by user.
  - decoding: unused
*/
func (s *AVCodecContext) RcOverrideCount() int {
	value := s.ptr.rc_override_count
	return int(value)
}

// SetRcOverrideCount sets the rc_override_count field.
/*
  ratecontrol override, see RcOverride
  - encoding: Allocated/set/freed by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetRcOverrideCount(value int) {
	s.ptr.rc_override_count = (C.int)(value)
}

// RcOverride gets the rc_override field.
func (s *AVCodecContext) RcOverride() *RcOverride {
	value := s.ptr.rc_override
	var valueMapped *RcOverride
	if value != nil {
		valueMapped = &RcOverride{ptr: value}
	}
	return valueMapped
}

// SetRcOverride sets the rc_override field.
func (s *AVCodecContext) SetRcOverride(value *RcOverride) {
	if value != nil {
		s.ptr.rc_override = value.ptr
	} else {
		s.ptr.rc_override = nil
	}
}

// RcMaxRate gets the rc_max_rate field.
/*
  maximum bitrate
  - encoding: Set by user.
  - decoding: Set by user, may be overwritten by libavcodec.
*/
func (s *AVCodecContext) RcMaxRate() int64 {
	value := s.ptr.rc_max_rate
	return int64(value)
}

// SetRcMaxRate sets the rc_max_rate field.
/*
  maximum bitrate
  - encoding: Set by user.
  - decoding: Set by user, may be overwritten by libavcodec.
*/
func (s *AVCodecContext) SetRcMaxRate(value int64) {
	s.ptr.rc_max_rate = (C.int64_t)(value)
}

// RcMinRate gets the rc_min_rate field.
/*
  minimum bitrate
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) RcMinRate() int64 {
	value := s.ptr.rc_min_rate
	return int64(value)
}

// SetRcMinRate sets the rc_min_rate field.
/*
  minimum bitrate
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetRcMinRate(value int64) {
	s.ptr.rc_min_rate = (C.int64_t)(value)
}

// RcMaxAvailableVbvUse gets the rc_max_available_vbv_use field.
/*
  Ratecontrol attempt to use, at maximum, <value> of what can be used without an underflow.
  - encoding: Set by user.
  - decoding: unused.
*/
func (s *AVCodecContext) RcMaxAvailableVbvUse() float32 {
	value := s.ptr.rc_max_available_vbv_use
	return float32(value)
}

// SetRcMaxAvailableVbvUse sets the rc_max_available_vbv_use field.
/*
  Ratecontrol attempt to use, at maximum, <value> of what can be used without an underflow.
  - encoding: Set by user.
  - decoding: unused.
*/
func (s *AVCodecContext) SetRcMaxAvailableVbvUse(value float32) {
	s.ptr.rc_max_available_vbv_use = (C.float)(value)
}

// RcMinVbvOverflowUse gets the rc_min_vbv_overflow_use field.
/*
  Ratecontrol attempt to use, at least, <value> times the amount needed to prevent a vbv overflow.
  - encoding: Set by user.
  - decoding: unused.
*/
func (s *AVCodecContext) RcMinVbvOverflowUse() float32 {
	value := s.ptr.rc_min_vbv_overflow_use
	return float32(value)
}

// SetRcMinVbvOverflowUse sets the rc_min_vbv_overflow_use field.
/*
  Ratecontrol attempt to use, at least, <value> times the amount needed to prevent a vbv overflow.
  - encoding: Set by user.
  - decoding: unused.
*/
func (s *AVCodecContext) SetRcMinVbvOverflowUse(value float32) {
	s.ptr.rc_min_vbv_overflow_use = (C.float)(value)
}

// RcInitialBufferOccupancy gets the rc_initial_buffer_occupancy field.
/*
  Number of bits which should be loaded into the rc buffer before decoding starts.
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) RcInitialBufferOccupancy() int {
	value := s.ptr.rc_initial_buffer_occupancy
	return int(value)
}

// SetRcInitialBufferOccupancy sets the rc_initial_buffer_occupancy field.
/*
  Number of bits which should be loaded into the rc buffer before decoding starts.
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetRcInitialBufferOccupancy(value int) {
	s.ptr.rc_initial_buffer_occupancy = (C.int)(value)
}

// Trellis gets the trellis field.
/*
  trellis RD quantization
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) Trellis() int {
	value := s.ptr.trellis
	return int(value)
}

// SetTrellis sets the trellis field.
/*
  trellis RD quantization
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetTrellis(value int) {
	s.ptr.trellis = (C.int)(value)
}

// StatsOut gets the stats_out field.
/*
  pass1 encoding statistics output buffer
  - encoding: Set by libavcodec.
  - decoding: unused
*/
func (s *AVCodecContext) StatsOut() *CStr {
	value := s.ptr.stats_out
	return wrapCStr(value)
}

// SetStatsOut sets the stats_out field.
/*
  pass1 encoding statistics output buffer
  - encoding: Set by libavcodec.
  - decoding: unused
*/
func (s *AVCodecContext) SetStatsOut(value *CStr) {
	s.ptr.stats_out = value.ptr
}

// StatsIn gets the stats_in field.
/*
  pass2 encoding statistics input buffer
  Concatenated stuff from stats_out of pass1 should be placed here.
  - encoding: Allocated/set/freed by user.
  - decoding: unused
*/
func (s *AVCodecContext) StatsIn() *CStr {
	value := s.ptr.stats_in
	return wrapCStr(value)
}

// SetStatsIn sets the stats_in field.
/*
  pass2 encoding statistics input buffer
  Concatenated stuff from stats_out of pass1 should be placed here.
  - encoding: Allocated/set/freed by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetStatsIn(value *CStr) {
	s.ptr.stats_in = value.ptr
}

// WorkaroundBugs gets the workaround_bugs field.
/*
  Work around bugs in encoders which sometimes cannot be detected automatically.
  - encoding: Set by user
  - decoding: Set by user
*/
func (s *AVCodecContext) WorkaroundBugs() int {
	value := s.ptr.workaround_bugs
	return int(value)
}

// SetWorkaroundBugs sets the workaround_bugs field.
/*
  Work around bugs in encoders which sometimes cannot be detected automatically.
  - encoding: Set by user
  - decoding: Set by user
*/
func (s *AVCodecContext) SetWorkaroundBugs(value int) {
	s.ptr.workaround_bugs = (C.int)(value)
}

// StrictStdCompliance gets the strict_std_compliance field.
/*
  strictly follow the standard (MPEG-4, ...).
  - encoding: Set by user.
  - decoding: Set by user.
  Setting this to STRICT or higher means the encoder and decoder will
  generally do stupid things, whereas setting it to unofficial or lower
  will mean the encoder might produce output that is not supported by all
  spec-compliant decoders. Decoders don't differentiate between normal,
  unofficial and experimental (that is, they always try to decode things
  when they can) unless they are explicitly asked to behave stupidly
  (=strictly conform to the specs)
  This may only be set to one of the FF_COMPLIANCE_* values in defs.h.
*/
func (s *AVCodecContext) StrictStdCompliance() int {
	value := s.ptr.strict_std_compliance
	return int(value)
}

// SetStrictStdCompliance sets the strict_std_compliance field.
/*
  strictly follow the standard (MPEG-4, ...).
  - encoding: Set by user.
  - decoding: Set by user.
  Setting this to STRICT or higher means the encoder and decoder will
  generally do stupid things, whereas setting it to unofficial or lower
  will mean the encoder might produce output that is not supported by all
  spec-compliant decoders. Decoders don't differentiate between normal,
  unofficial and experimental (that is, they always try to decode things
  when they can) unless they are explicitly asked to behave stupidly
  (=strictly conform to the specs)
  This may only be set to one of the FF_COMPLIANCE_* values in defs.h.
*/
func (s *AVCodecContext) SetStrictStdCompliance(value int) {
	s.ptr.strict_std_compliance = (C.int)(value)
}

// ErrorConcealment gets the error_concealment field.
/*
  error concealment flags
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) ErrorConcealment() int {
	value := s.ptr.error_concealment
	return int(value)
}

// SetErrorConcealment sets the error_concealment field.
/*
  error concealment flags
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetErrorConcealment(value int) {
	s.ptr.error_concealment = (C.int)(value)
}

// Debug gets the debug field.
/*
  debug
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) Debug() int {
	value := s.ptr.debug
	return int(value)
}

// SetDebug sets the debug field.
/*
  debug
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetDebug(value int) {
	s.ptr.debug = (C.int)(value)
}

// ErrRecognition gets the err_recognition field.
/*
  Error recognition; may misdetect some more or less valid parts as errors.
  This is a bitfield of the AV_EF_* values defined in defs.h.

  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) ErrRecognition() int {
	value := s.ptr.err_recognition
	return int(value)
}

// SetErrRecognition sets the err_recognition field.
/*
  Error recognition; may misdetect some more or less valid parts as errors.
  This is a bitfield of the AV_EF_* values defined in defs.h.

  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetErrRecognition(value int) {
	s.ptr.err_recognition = (C.int)(value)
}

// ReorderedOpaque gets the reordered_opaque field.
/*
  opaque 64-bit number (generally a PTS) that will be reordered and
  output in AVFrame.reordered_opaque
  - encoding: Set by libavcodec to the reordered_opaque of the input
              frame corresponding to the last returned packet. Only
              supported by encoders with the
              AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE capability.
  - decoding: Set by user.

  @deprecated Use AV_CODEC_FLAG_COPY_OPAQUE instead
*/
func (s *AVCodecContext) ReorderedOpaque() int64 {
	value := s.ptr.reordered_opaque
	return int64(value)
}

// SetReorderedOpaque sets the reordered_opaque field.
/*
  opaque 64-bit number (generally a PTS) that will be reordered and
  output in AVFrame.reordered_opaque
  - encoding: Set by libavcodec to the reordered_opaque of the input
              frame corresponding to the last returned packet. Only
              supported by encoders with the
              AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE capability.
  - decoding: Set by user.

  @deprecated Use AV_CODEC_FLAG_COPY_OPAQUE instead
*/
func (s *AVCodecContext) SetReorderedOpaque(value int64) {
	s.ptr.reordered_opaque = (C.int64_t)(value)
}

// Hwaccel gets the hwaccel field.
/*
  Hardware accelerator in use
  - encoding: unused.
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) Hwaccel() *AVHWAccel {
	value := s.ptr.hwaccel
	var valueMapped *AVHWAccel
	if value != nil {
		valueMapped = &AVHWAccel{ptr: value}
	}
	return valueMapped
}

// SetHwaccel sets the hwaccel field.
/*
  Hardware accelerator in use
  - encoding: unused.
  - decoding: Set by libavcodec
*/
func (s *AVCodecContext) SetHwaccel(value *AVHWAccel) {
	if value != nil {
		s.ptr.hwaccel = value.ptr
	} else {
		s.ptr.hwaccel = nil
	}
}

// HwaccelContext gets the hwaccel_context field.
/*
  Legacy hardware accelerator context.

  For some hardware acceleration methods, the caller may use this field to
  signal hwaccel-specific data to the codec. The struct pointed to by this
  pointer is hwaccel-dependent and defined in the respective header. Please
  refer to the FFmpeg HW accelerator documentation to know how to fill
  this.

  In most cases this field is optional - the necessary information may also
  be provided to libavcodec through @ref hw_frames_ctx or @ref
  hw_device_ctx (see avcodec_get_hw_config()). However, in some cases it
  may be the only method of signalling some (optional) information.

  The struct and its contents are owned by the caller.

  - encoding: May be set by the caller before avcodec_open2(). Must remain
              valid until avcodec_free_context().
  - decoding: May be set by the caller in the get_format() callback.
              Must remain valid until the next get_format() call,
              or avcodec_free_context() (whichever comes first).
*/
func (s *AVCodecContext) HwaccelContext() unsafe.Pointer {
	value := s.ptr.hwaccel_context
	return value
}

// SetHwaccelContext sets the hwaccel_context field.
/*
  Legacy hardware accelerator context.

  For some hardware acceleration methods, the caller may use this field to
  signal hwaccel-specific data to the codec. The struct pointed to by this
  pointer is hwaccel-dependent and defined in the respective header. Please
  refer to the FFmpeg HW accelerator documentation to know how to fill
  this.

  In most cases this field is optional - the necessary information may also
  be provided to libavcodec through @ref hw_frames_ctx or @ref
  hw_device_ctx (see avcodec_get_hw_config()). However, in some cases it
  may be the only method of signalling some (optional) information.

  The struct and its contents are owned by the caller.

  - encoding: May be set by the caller before avcodec_open2(). Must remain
              valid until avcodec_free_context().
  - decoding: May be set by the caller in the get_format() callback.
              Must remain valid until the next get_format() call,
              or avcodec_free_context() (whichever comes first).
*/
func (s *AVCodecContext) SetHwaccelContext(value unsafe.Pointer) {
	s.ptr.hwaccel_context = value
}

// Error gets the error field.
/*
  error
  - encoding: Set by libavcodec if flags & AV_CODEC_FLAG_PSNR.
  - decoding: unused
*/
func (s *AVCodecContext) Error() *Array[uint64] {
	value := &s.ptr.error
	return ToUint64Array(unsafe.Pointer(value))
}

// DctAlgo gets the dct_algo field.
/*
  DCT algorithm, see FF_DCT_* below
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) DctAlgo() int {
	value := s.ptr.dct_algo
	return int(value)
}

// SetDctAlgo sets the dct_algo field.
/*
  DCT algorithm, see FF_DCT_* below
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetDctAlgo(value int) {
	s.ptr.dct_algo = (C.int)(value)
}

// IdctAlgo gets the idct_algo field.
/*
  IDCT algorithm, see FF_IDCT_* below.
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) IdctAlgo() int {
	value := s.ptr.idct_algo
	return int(value)
}

// SetIdctAlgo sets the idct_algo field.
/*
  IDCT algorithm, see FF_IDCT_* below.
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetIdctAlgo(value int) {
	s.ptr.idct_algo = (C.int)(value)
}

// BitsPerCodedSample gets the bits_per_coded_sample field.
/*
  bits per sample/pixel from the demuxer (needed for huffyuv).
  - encoding: Set by libavcodec.
  - decoding: Set by user.
*/
func (s *AVCodecContext) BitsPerCodedSample() int {
	value := s.ptr.bits_per_coded_sample
	return int(value)
}

// SetBitsPerCodedSample sets the bits_per_coded_sample field.
/*
  bits per sample/pixel from the demuxer (needed for huffyuv).
  - encoding: Set by libavcodec.
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetBitsPerCodedSample(value int) {
	s.ptr.bits_per_coded_sample = (C.int)(value)
}

// BitsPerRawSample gets the bits_per_raw_sample field.
/*
  Bits per sample/pixel of internal libavcodec pixel/sample format.
  - encoding: set by user.
  - decoding: set by libavcodec.
*/
func (s *AVCodecContext) BitsPerRawSample() int {
	value := s.ptr.bits_per_raw_sample
	return int(value)
}

// SetBitsPerRawSample sets the bits_per_raw_sample field.
/*
  Bits per sample/pixel of internal libavcodec pixel/sample format.
  - encoding: set by user.
  - decoding: set by libavcodec.
*/
func (s *AVCodecContext) SetBitsPerRawSample(value int) {
	s.ptr.bits_per_raw_sample = (C.int)(value)
}

// Lowres gets the lowres field.
/*
  low resolution decoding, 1-> 1/2 size, 2->1/4 size
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) Lowres() int {
	value := s.ptr.lowres
	return int(value)
}

// SetLowres sets the lowres field.
/*
  low resolution decoding, 1-> 1/2 size, 2->1/4 size
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetLowres(value int) {
	s.ptr.lowres = (C.int)(value)
}

// ThreadCount gets the thread_count field.
/*
  thread count
  is used to decide how many independent tasks should be passed to execute()
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) ThreadCount() int {
	value := s.ptr.thread_count
	return int(value)
}

// SetThreadCount sets the thread_count field.
/*
  thread count
  is used to decide how many independent tasks should be passed to execute()
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetThreadCount(value int) {
	s.ptr.thread_count = (C.int)(value)
}

// ThreadType gets the thread_type field.
/*
  Which multithreading methods to use.
  Use of FF_THREAD_FRAME will increase decoding delay by one frame per thread,
  so clients which cannot provide future frames should not use it.

  - encoding: Set by user, otherwise the default is used.
  - decoding: Set by user, otherwise the default is used.
*/
func (s *AVCodecContext) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}

// SetThreadType sets the thread_type field.
/*
  Which multithreading methods to use.
  Use of FF_THREAD_FRAME will increase decoding delay by one frame per thread,
  so clients which cannot provide future frames should not use it.

  - encoding: Set by user, otherwise the default is used.
  - decoding: Set by user, otherwise the default is used.
*/
func (s *AVCodecContext) SetThreadType(value int) {
	s.ptr.thread_type = (C.int)(value)
}

// ActiveThreadType gets the active_thread_type field.
/*
  Which multithreading methods are in use by the codec.
  - encoding: Set by libavcodec.
  - decoding: Set by libavcodec.
*/
func (s *AVCodecContext) ActiveThreadType() int {
	value := s.ptr.active_thread_type
	return int(value)
}

// SetActiveThreadType sets the active_thread_type field.
/*
  Which multithreading methods are in use by the codec.
  - encoding: Set by libavcodec.
  - decoding: Set by libavcodec.
*/
func (s *AVCodecContext) SetActiveThreadType(value int) {
	s.ptr.active_thread_type = (C.int)(value)
}

// execute skipped due to func ptr

// execute2 skipped due to func ptr

// NsseWeight gets the nsse_weight field.
/*
  noise vs. sse weight for the nsse comparison function
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) NsseWeight() int {
	value := s.ptr.nsse_weight
	return int(value)
}

// SetNsseWeight sets the nsse_weight field.
/*
  noise vs. sse weight for the nsse comparison function
  - encoding: Set by user.
  - decoding: unused
*/
func (s *AVCodecContext) SetNsseWeight(value int) {
	s.ptr.nsse_weight = (C.int)(value)
}

// Profile gets the profile field.
/*
  profile
  - encoding: Set by user.
  - decoding: Set by libavcodec.
  See the AV_PROFILE_* defines in defs.h.
*/
func (s *AVCodecContext) Profile() int {
	value := s.ptr.profile
	return int(value)
}

// SetProfile sets the profile field.
/*
  profile
  - encoding: Set by user.
  - decoding: Set by libavcodec.
  See the AV_PROFILE_* defines in defs.h.
*/
func (s *AVCodecContext) SetProfile(value int) {
	s.ptr.profile = (C.int)(value)
}

// Level gets the level field.
/*
  Encoding level descriptor.
  - encoding: Set by user, corresponds to a specific level defined by the
    codec, usually corresponding to the profile level, if not specified it
    is set to FF_LEVEL_UNKNOWN.
  - decoding: Set by libavcodec.
  See AV_LEVEL_* in defs.h.
*/
func (s *AVCodecContext) Level() int {
	value := s.ptr.level
	return int(value)
}

// SetLevel sets the level field.
/*
  Encoding level descriptor.
  - encoding: Set by user, corresponds to a specific level defined by the
    codec, usually corresponding to the profile level, if not specified it
    is set to FF_LEVEL_UNKNOWN.
  - decoding: Set by libavcodec.
  See AV_LEVEL_* in defs.h.
*/
func (s *AVCodecContext) SetLevel(value int) {
	s.ptr.level = (C.int)(value)
}

// SkipLoopFilter gets the skip_loop_filter field.
/*
  Skip loop filtering for selected frames.
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SkipLoopFilter() AVDiscard {
	value := s.ptr.skip_loop_filter
	return AVDiscard(value)
}

// SetSkipLoopFilter sets the skip_loop_filter field.
/*
  Skip loop filtering for selected frames.
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetSkipLoopFilter(value AVDiscard) {
	s.ptr.skip_loop_filter = (C.enum_AVDiscard)(value)
}

// SkipIdct gets the skip_idct field.
/*
  Skip IDCT/dequantization for selected frames.
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SkipIdct() AVDiscard {
	value := s.ptr.skip_idct
	return AVDiscard(value)
}

// SetSkipIdct sets the skip_idct field.
/*
  Skip IDCT/dequantization for selected frames.
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetSkipIdct(value AVDiscard) {
	s.ptr.skip_idct = (C.enum_AVDiscard)(value)
}

// SkipFrame gets the skip_frame field.
/*
  Skip decoding for selected frames.
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SkipFrame() AVDiscard {
	value := s.ptr.skip_frame
	return AVDiscard(value)
}

// SetSkipFrame sets the skip_frame field.
/*
  Skip decoding for selected frames.
  - encoding: unused
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetSkipFrame(value AVDiscard) {
	s.ptr.skip_frame = (C.enum_AVDiscard)(value)
}

// SubtitleHeader gets the subtitle_header field.
/*
  Header containing style information for text subtitles.
  For SUBTITLE_ASS subtitle type, it should contain the whole ASS
  [Script Info] and [V4+ Styles] section, plus the [Events] line and
  the Format line following. It shouldn't include any Dialogue line.
  - encoding: Set/allocated/freed by user (before avcodec_open2())
  - decoding: Set/allocated/freed by libavcodec (by avcodec_open2())
*/
func (s *AVCodecContext) SubtitleHeader() unsafe.Pointer {
	value := s.ptr.subtitle_header
	return unsafe.Pointer(value)
}

// SetSubtitleHeader sets the subtitle_header field.
/*
  Header containing style information for text subtitles.
  For SUBTITLE_ASS subtitle type, it should contain the whole ASS
  [Script Info] and [V4+ Styles] section, plus the [Events] line and
  the Format line following. It shouldn't include any Dialogue line.
  - encoding: Set/allocated/freed by user (before avcodec_open2())
  - decoding: Set/allocated/freed by libavcodec (by avcodec_open2())
*/
func (s *AVCodecContext) SetSubtitleHeader(value unsafe.Pointer) {
	s.ptr.subtitle_header = (*C.uint8_t)(value)
}

// SubtitleHeaderSize gets the subtitle_header_size field.
func (s *AVCodecContext) SubtitleHeaderSize() int {
	value := s.ptr.subtitle_header_size
	return int(value)
}

// SetSubtitleHeaderSize sets the subtitle_header_size field.
func (s *AVCodecContext) SetSubtitleHeaderSize(value int) {
	s.ptr.subtitle_header_size = (C.int)(value)
}

// InitialPadding gets the initial_padding field.
/*
  Audio only. The number of "priming" samples (padding) inserted by the
  encoder at the beginning of the audio. I.e. this number of leading
  decoded samples must be discarded by the caller to get the original audio
  without leading padding.

  - decoding: unused
  - encoding: Set by libavcodec. The timestamps on the output packets are
              adjusted by the encoder so that they always refer to the
              first sample of the data actually contained in the packet,
              including any added padding.  E.g. if the timebase is
              1/samplerate and the timestamp of the first input sample is
              0, the timestamp of the first output packet will be
              -initial_padding.
*/
func (s *AVCodecContext) InitialPadding() int {
	value := s.ptr.initial_padding
	return int(value)
}

// SetInitialPadding sets the initial_padding field.
/*
  Audio only. The number of "priming" samples (padding) inserted by the
  encoder at the beginning of the audio. I.e. this number of leading
  decoded samples must be discarded by the caller to get the original audio
  without leading padding.

  - decoding: unused
  - encoding: Set by libavcodec. The timestamps on the output packets are
              adjusted by the encoder so that they always refer to the
              first sample of the data actually contained in the packet,
              including any added padding.  E.g. if the timebase is
              1/samplerate and the timestamp of the first input sample is
              0, the timestamp of the first output packet will be
              -initial_padding.
*/
func (s *AVCodecContext) SetInitialPadding(value int) {
	s.ptr.initial_padding = (C.int)(value)
}

// Framerate gets the framerate field.
/*
  - decoding: For codecs that store a framerate value in the compressed
              bitstream, the decoder may export it here. { 0, 1} when
              unknown.
  - encoding: May be used to signal the framerate of CFR content to an
              encoder.
*/
func (s *AVCodecContext) Framerate() *AVRational {
	value := s.ptr.framerate
	return &AVRational{value: value}
}

// SetFramerate sets the framerate field.
/*
  - decoding: For codecs that store a framerate value in the compressed
              bitstream, the decoder may export it here. { 0, 1} when
              unknown.
  - encoding: May be used to signal the framerate of CFR content to an
              encoder.
*/
func (s *AVCodecContext) SetFramerate(value *AVRational) {
	s.ptr.framerate = value.value
}

// SwPixFmt gets the sw_pix_fmt field.
/*
  Nominal unaccelerated pixel format, see AV_PIX_FMT_xxx.
  - encoding: unused.
  - decoding: Set by libavcodec before calling get_format()
*/
func (s *AVCodecContext) SwPixFmt() AVPixelFormat {
	value := s.ptr.sw_pix_fmt
	return AVPixelFormat(value)
}

// SetSwPixFmt sets the sw_pix_fmt field.
/*
  Nominal unaccelerated pixel format, see AV_PIX_FMT_xxx.
  - encoding: unused.
  - decoding: Set by libavcodec before calling get_format()
*/
func (s *AVCodecContext) SetSwPixFmt(value AVPixelFormat) {
	s.ptr.sw_pix_fmt = (C.enum_AVPixelFormat)(value)
}

// PktTimebase gets the pkt_timebase field.
/*
  Timebase in which pkt_dts/pts and AVPacket.dts/pts are expressed.
  - encoding: unused.
  - decoding: set by user.
*/
func (s *AVCodecContext) PktTimebase() *AVRational {
	value := s.ptr.pkt_timebase
	return &AVRational{value: value}
}

// SetPktTimebase sets the pkt_timebase field.
/*
  Timebase in which pkt_dts/pts and AVPacket.dts/pts are expressed.
  - encoding: unused.
  - decoding: set by user.
*/
func (s *AVCodecContext) SetPktTimebase(value *AVRational) {
	s.ptr.pkt_timebase = value.value
}

// CodecDescriptor gets the codec_descriptor field.
/*
  AVCodecDescriptor
  - encoding: unused.
  - decoding: set by libavcodec.
*/
func (s *AVCodecContext) CodecDescriptor() *AVCodecDescriptor {
	value := s.ptr.codec_descriptor
	var valueMapped *AVCodecDescriptor
	if value != nil {
		valueMapped = &AVCodecDescriptor{ptr: value}
	}
	return valueMapped
}

// SetCodecDescriptor sets the codec_descriptor field.
/*
  AVCodecDescriptor
  - encoding: unused.
  - decoding: set by libavcodec.
*/
func (s *AVCodecContext) SetCodecDescriptor(value *AVCodecDescriptor) {
	if value != nil {
		s.ptr.codec_descriptor = value.ptr
	} else {
		s.ptr.codec_descriptor = nil
	}
}

// PtsCorrectionNumFaultyPts gets the pts_correction_num_faulty_pts field.
/*
  Current statistics for PTS correction.
  - decoding: maintained and used by libavcodec, not intended to be used by user apps
  - encoding: unused
*/
func (s *AVCodecContext) PtsCorrectionNumFaultyPts() int64 {
	value := s.ptr.pts_correction_num_faulty_pts
	return int64(value)
}

// SetPtsCorrectionNumFaultyPts sets the pts_correction_num_faulty_pts field.
/*
  Current statistics for PTS correction.
  - decoding: maintained and used by libavcodec, not intended to be used by user apps
  - encoding: unused
*/
func (s *AVCodecContext) SetPtsCorrectionNumFaultyPts(value int64) {
	s.ptr.pts_correction_num_faulty_pts = (C.int64_t)(value)
}

// PtsCorrectionNumFaultyDts gets the pts_correction_num_faulty_dts field.
//
//	Number of incorrect PTS values so far
func (s *AVCodecContext) PtsCorrectionNumFaultyDts() int64 {
	value := s.ptr.pts_correction_num_faulty_dts
	return int64(value)
}

// SetPtsCorrectionNumFaultyDts sets the pts_correction_num_faulty_dts field.
//
//	Number of incorrect PTS values so far
func (s *AVCodecContext) SetPtsCorrectionNumFaultyDts(value int64) {
	s.ptr.pts_correction_num_faulty_dts = (C.int64_t)(value)
}

// PtsCorrectionLastPts gets the pts_correction_last_pts field.
//
//	Number of incorrect DTS values so far
func (s *AVCodecContext) PtsCorrectionLastPts() int64 {
	value := s.ptr.pts_correction_last_pts
	return int64(value)
}

// SetPtsCorrectionLastPts sets the pts_correction_last_pts field.
//
//	Number of incorrect DTS values so far
func (s *AVCodecContext) SetPtsCorrectionLastPts(value int64) {
	s.ptr.pts_correction_last_pts = (C.int64_t)(value)
}

// PtsCorrectionLastDts gets the pts_correction_last_dts field.
//
//	PTS of the last frame
func (s *AVCodecContext) PtsCorrectionLastDts() int64 {
	value := s.ptr.pts_correction_last_dts
	return int64(value)
}

// SetPtsCorrectionLastDts sets the pts_correction_last_dts field.
//
//	PTS of the last frame
func (s *AVCodecContext) SetPtsCorrectionLastDts(value int64) {
	s.ptr.pts_correction_last_dts = (C.int64_t)(value)
}

// SubCharenc gets the sub_charenc field.
/*
  Character encoding of the input subtitles file.
  - decoding: set by user
  - encoding: unused
*/
func (s *AVCodecContext) SubCharenc() *CStr {
	value := s.ptr.sub_charenc
	return wrapCStr(value)
}

// SetSubCharenc sets the sub_charenc field.
/*
  Character encoding of the input subtitles file.
  - decoding: set by user
  - encoding: unused
*/
func (s *AVCodecContext) SetSubCharenc(value *CStr) {
	s.ptr.sub_charenc = value.ptr
}

// SubCharencMode gets the sub_charenc_mode field.
/*
  Subtitles character encoding mode. Formats or codecs might be adjusting
  this setting (if they are doing the conversion themselves for instance).
  - decoding: set by libavcodec
  - encoding: unused
*/
func (s *AVCodecContext) SubCharencMode() int {
	value := s.ptr.sub_charenc_mode
	return int(value)
}

// SetSubCharencMode sets the sub_charenc_mode field.
/*
  Subtitles character encoding mode. Formats or codecs might be adjusting
  this setting (if they are doing the conversion themselves for instance).
  - decoding: set by libavcodec
  - encoding: unused
*/
func (s *AVCodecContext) SetSubCharencMode(value int) {
	s.ptr.sub_charenc_mode = (C.int)(value)
}

// SkipAlpha gets the skip_alpha field.
/*
  Skip processing alpha if supported by codec.
  Note that if the format uses pre-multiplied alpha (common with VP6,
  and recommended due to better video quality/compression)
  the image will look as if alpha-blended onto a black background.
  However for formats that do not use pre-multiplied alpha
  there might be serious artefacts (though e.g. libswscale currently
  assumes pre-multiplied alpha anyway).

  - decoding: set by user
  - encoding: unused
*/
func (s *AVCodecContext) SkipAlpha() int {
	value := s.ptr.skip_alpha
	return int(value)
}

// SetSkipAlpha sets the skip_alpha field.
/*
  Skip processing alpha if supported by codec.
  Note that if the format uses pre-multiplied alpha (common with VP6,
  and recommended due to better video quality/compression)
  the image will look as if alpha-blended onto a black background.
  However for formats that do not use pre-multiplied alpha
  there might be serious artefacts (though e.g. libswscale currently
  assumes pre-multiplied alpha anyway).

  - decoding: set by user
  - encoding: unused
*/
func (s *AVCodecContext) SetSkipAlpha(value int) {
	s.ptr.skip_alpha = (C.int)(value)
}

// SeekPreroll gets the seek_preroll field.
/*
  Number of samples to skip after a discontinuity
  - decoding: unused
  - encoding: set by libavcodec
*/
func (s *AVCodecContext) SeekPreroll() int {
	value := s.ptr.seek_preroll
	return int(value)
}

// SetSeekPreroll sets the seek_preroll field.
/*
  Number of samples to skip after a discontinuity
  - decoding: unused
  - encoding: set by libavcodec
*/
func (s *AVCodecContext) SetSeekPreroll(value int) {
	s.ptr.seek_preroll = (C.int)(value)
}

// chroma_intra_matrix skipped due to prim ptr

// DumpSeparator gets the dump_separator field.
/*
  dump format separator.
  can be ", " or "\n      " or anything else
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) DumpSeparator() unsafe.Pointer {
	value := s.ptr.dump_separator
	return unsafe.Pointer(value)
}

// SetDumpSeparator sets the dump_separator field.
/*
  dump format separator.
  can be ", " or "\n      " or anything else
  - encoding: Set by user.
  - decoding: Set by user.
*/
func (s *AVCodecContext) SetDumpSeparator(value unsafe.Pointer) {
	s.ptr.dump_separator = (*C.uint8_t)(value)
}

// CodecWhitelist gets the codec_whitelist field.
/*
  ',' separated list of allowed decoders.
  If NULL then all are allowed
  - encoding: unused
  - decoding: set by user
*/
func (s *AVCodecContext) CodecWhitelist() *CStr {
	value := s.ptr.codec_whitelist
	return wrapCStr(value)
}

// SetCodecWhitelist sets the codec_whitelist field.
/*
  ',' separated list of allowed decoders.
  If NULL then all are allowed
  - encoding: unused
  - decoding: set by user
*/
func (s *AVCodecContext) SetCodecWhitelist(value *CStr) {
	s.ptr.codec_whitelist = value.ptr
}

// Properties gets the properties field.
/*
  Properties of the stream that gets decoded
  - encoding: unused
  - decoding: set by libavcodec
*/
func (s *AVCodecContext) Properties() uint {
	value := s.ptr.properties
	return uint(value)
}

// SetProperties sets the properties field.
/*
  Properties of the stream that gets decoded
  - encoding: unused
  - decoding: set by libavcodec
*/
func (s *AVCodecContext) SetProperties(value uint) {
	s.ptr.properties = (C.uint)(value)
}

// CodedSideData gets the coded_side_data field.
/*
  Additional data associated with the entire coded stream.

  - decoding: may be set by user before calling avcodec_open2().
  - encoding: may be set by libavcodec after avcodec_open2().
*/
func (s *AVCodecContext) CodedSideData() *AVPacketSideData {
	value := s.ptr.coded_side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}

// SetCodedSideData sets the coded_side_data field.
/*
  Additional data associated with the entire coded stream.

  - decoding: may be set by user before calling avcodec_open2().
  - encoding: may be set by libavcodec after avcodec_open2().
*/
func (s *AVCodecContext) SetCodedSideData(value *AVPacketSideData) {
	if value != nil {
		s.ptr.coded_side_data = value.ptr
	} else {
		s.ptr.coded_side_data = nil
	}
}

// NbCodedSideData gets the nb_coded_side_data field.
func (s *AVCodecContext) NbCodedSideData() int {
	value := s.ptr.nb_coded_side_data
	return int(value)
}

// SetNbCodedSideData sets the nb_coded_side_data field.
func (s *AVCodecContext) SetNbCodedSideData(value int) {
	s.ptr.nb_coded_side_data = (C.int)(value)
}

// HwFramesCtx gets the hw_frames_ctx field.
/*
  A reference to the AVHWFramesContext describing the input (for encoding)
  or output (decoding) frames. The reference is set by the caller and
  afterwards owned (and freed) by libavcodec - it should never be read by
  the caller after being set.

  - decoding: This field should be set by the caller from the get_format()
              callback. The previous reference (if any) will always be
              unreffed by libavcodec before the get_format() call.

              If the default get_buffer2() is used with a hwaccel pixel
              format, then this AVHWFramesContext will be used for
              allocating the frame buffers.

  - encoding: For hardware encoders configured to use a hwaccel pixel
              format, this field should be set by the caller to a reference
              to the AVHWFramesContext describing input frames.
              AVHWFramesContext.format must be equal to
              AVCodecContext.pix_fmt.

              This field should be set before avcodec_open2() is called.
*/
func (s *AVCodecContext) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwFramesCtx sets the hw_frames_ctx field.
/*
  A reference to the AVHWFramesContext describing the input (for encoding)
  or output (decoding) frames. The reference is set by the caller and
  afterwards owned (and freed) by libavcodec - it should never be read by
  the caller after being set.

  - decoding: This field should be set by the caller from the get_format()
              callback. The previous reference (if any) will always be
              unreffed by libavcodec before the get_format() call.

              If the default get_buffer2() is used with a hwaccel pixel
              format, then this AVHWFramesContext will be used for
              allocating the frame buffers.

  - encoding: For hardware encoders configured to use a hwaccel pixel
              format, this field should be set by the caller to a reference
              to the AVHWFramesContext describing input frames.
              AVHWFramesContext.format must be equal to
              AVCodecContext.pix_fmt.

              This field should be set before avcodec_open2() is called.
*/
func (s *AVCodecContext) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

// TrailingPadding gets the trailing_padding field.
/*
  Audio only. The amount of padding (in samples) appended by the encoder to
  the end of the audio. I.e. this number of decoded samples must be
  discarded by the caller from the end of the stream to get the original
  audio without any trailing padding.

  - decoding: unused
  - encoding: unused
*/
func (s *AVCodecContext) TrailingPadding() int {
	value := s.ptr.trailing_padding
	return int(value)
}

// SetTrailingPadding sets the trailing_padding field.
/*
  Audio only. The amount of padding (in samples) appended by the encoder to
  the end of the audio. I.e. this number of decoded samples must be
  discarded by the caller from the end of the stream to get the original
  audio without any trailing padding.

  - decoding: unused
  - encoding: unused
*/
func (s *AVCodecContext) SetTrailingPadding(value int) {
	s.ptr.trailing_padding = (C.int)(value)
}

// MaxPixels gets the max_pixels field.
/*
  The number of pixels per image to maximally accept.

  - decoding: set by user
  - encoding: set by user
*/
func (s *AVCodecContext) MaxPixels() int64 {
	value := s.ptr.max_pixels
	return int64(value)
}

// SetMaxPixels sets the max_pixels field.
/*
  The number of pixels per image to maximally accept.

  - decoding: set by user
  - encoding: set by user
*/
func (s *AVCodecContext) SetMaxPixels(value int64) {
	s.ptr.max_pixels = (C.int64_t)(value)
}

// HwDeviceCtx gets the hw_device_ctx field.
/*
  A reference to the AVHWDeviceContext describing the device which will
  be used by a hardware encoder/decoder.  The reference is set by the
  caller and afterwards owned (and freed) by libavcodec.

  This should be used if either the codec device does not require
  hardware frames or any that are used are to be allocated internally by
  libavcodec.  If the user wishes to supply any of the frames used as
  encoder input or decoder output then hw_frames_ctx should be used
  instead.  When hw_frames_ctx is set in get_format() for a decoder, this
  field will be ignored while decoding the associated stream segment, but
  may again be used on a following one after another get_format() call.

  For both encoders and decoders this field should be set before
  avcodec_open2() is called and must not be written to thereafter.

  Note that some decoders may require this field to be set initially in
  order to support hw_frames_ctx at all - in that case, all frames
  contexts used must be created on the same device.
*/
func (s *AVCodecContext) HwDeviceCtx() *AVBufferRef {
	value := s.ptr.hw_device_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwDeviceCtx sets the hw_device_ctx field.
/*
  A reference to the AVHWDeviceContext describing the device which will
  be used by a hardware encoder/decoder.  The reference is set by the
  caller and afterwards owned (and freed) by libavcodec.

  This should be used if either the codec device does not require
  hardware frames or any that are used are to be allocated internally by
  libavcodec.  If the user wishes to supply any of the frames used as
  encoder input or decoder output then hw_frames_ctx should be used
  instead.  When hw_frames_ctx is set in get_format() for a decoder, this
  field will be ignored while decoding the associated stream segment, but
  may again be used on a following one after another get_format() call.

  For both encoders and decoders this field should be set before
  avcodec_open2() is called and must not be written to thereafter.

  Note that some decoders may require this field to be set initially in
  order to support hw_frames_ctx at all - in that case, all frames
  contexts used must be created on the same device.
*/
func (s *AVCodecContext) SetHwDeviceCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_device_ctx = value.ptr
	} else {
		s.ptr.hw_device_ctx = nil
	}
}

// HwaccelFlags gets the hwaccel_flags field.
/*
  Bit set of AV_HWACCEL_FLAG_* flags, which affect hardware accelerated
  decoding (if active).
  - encoding: unused
  - decoding: Set by user (either before avcodec_open2(), or in the
              AVCodecContext.get_format callback)
*/
func (s *AVCodecContext) HwaccelFlags() int {
	value := s.ptr.hwaccel_flags
	return int(value)
}

// SetHwaccelFlags sets the hwaccel_flags field.
/*
  Bit set of AV_HWACCEL_FLAG_* flags, which affect hardware accelerated
  decoding (if active).
  - encoding: unused
  - decoding: Set by user (either before avcodec_open2(), or in the
              AVCodecContext.get_format callback)
*/
func (s *AVCodecContext) SetHwaccelFlags(value int) {
	s.ptr.hwaccel_flags = (C.int)(value)
}

// ApplyCropping gets the apply_cropping field.
/*
  Video decoding only. Certain video codecs support cropping, meaning that
  only a sub-rectangle of the decoded frame is intended for display.  This
  option controls how cropping is handled by libavcodec.

  When set to 1 (the default), libavcodec will apply cropping internally.
  I.e. it will modify the output frame width/height fields and offset the
  data pointers (only by as much as possible while preserving alignment, or
  by the full amount if the AV_CODEC_FLAG_UNALIGNED flag is set) so that
  the frames output by the decoder refer only to the cropped area. The
  crop_* fields of the output frames will be zero.

  When set to 0, the width/height fields of the output frames will be set
  to the coded dimensions and the crop_* fields will describe the cropping
  rectangle. Applying the cropping is left to the caller.

  @warning When hardware acceleration with opaque output frames is used,
  libavcodec is unable to apply cropping from the top/left border.

  @note when this option is set to zero, the width/height fields of the
  AVCodecContext and output AVFrames have different meanings. The codec
  context fields store display dimensions (with the coded dimensions in
  coded_width/height), while the frame fields store the coded dimensions
  (with the display dimensions being determined by the crop_* fields).
*/
func (s *AVCodecContext) ApplyCropping() int {
	value := s.ptr.apply_cropping
	return int(value)
}

// SetApplyCropping sets the apply_cropping field.
/*
  Video decoding only. Certain video codecs support cropping, meaning that
  only a sub-rectangle of the decoded frame is intended for display.  This
  option controls how cropping is handled by libavcodec.

  When set to 1 (the default), libavcodec will apply cropping internally.
  I.e. it will modify the output frame width/height fields and offset the
  data pointers (only by as much as possible while preserving alignment, or
  by the full amount if the AV_CODEC_FLAG_UNALIGNED flag is set) so that
  the frames output by the decoder refer only to the cropped area. The
  crop_* fields of the output frames will be zero.

  When set to 0, the width/height fields of the output frames will be set
  to the coded dimensions and the crop_* fields will describe the cropping
  rectangle. Applying the cropping is left to the caller.

  @warning When hardware acceleration with opaque output frames is used,
  libavcodec is unable to apply cropping from the top/left border.

  @note when this option is set to zero, the width/height fields of the
  AVCodecContext and output AVFrames have different meanings. The codec
  context fields store display dimensions (with the coded dimensions in
  coded_width/height), while the frame fields store the coded dimensions
  (with the display dimensions being determined by the crop_* fields).
*/
func (s *AVCodecContext) SetApplyCropping(value int) {
	s.ptr.apply_cropping = (C.int)(value)
}

// ExtraHwFrames gets the extra_hw_frames field.
/*
  Video decoding only.  Sets the number of extra hardware frames which
  the decoder will allocate for use by the caller.  This must be set
  before avcodec_open2() is called.

  Some hardware decoders require all frames that they will use for
  output to be defined in advance before decoding starts.  For such
  decoders, the hardware frame pool must therefore be of a fixed size.
  The extra frames set here are on top of any number that the decoder
  needs internally in order to operate normally (for example, frames
  used as reference pictures).
*/
func (s *AVCodecContext) ExtraHwFrames() int {
	value := s.ptr.extra_hw_frames
	return int(value)
}

// SetExtraHwFrames sets the extra_hw_frames field.
/*
  Video decoding only.  Sets the number of extra hardware frames which
  the decoder will allocate for use by the caller.  This must be set
  before avcodec_open2() is called.

  Some hardware decoders require all frames that they will use for
  output to be defined in advance before decoding starts.  For such
  decoders, the hardware frame pool must therefore be of a fixed size.
  The extra frames set here are on top of any number that the decoder
  needs internally in order to operate normally (for example, frames
  used as reference pictures).
*/
func (s *AVCodecContext) SetExtraHwFrames(value int) {
	s.ptr.extra_hw_frames = (C.int)(value)
}

// DiscardDamagedPercentage gets the discard_damaged_percentage field.
/*
  The percentage of damaged samples to discard a frame.

  - decoding: set by user
  - encoding: unused
*/
func (s *AVCodecContext) DiscardDamagedPercentage() int {
	value := s.ptr.discard_damaged_percentage
	return int(value)
}

// SetDiscardDamagedPercentage sets the discard_damaged_percentage field.
/*
  The percentage of damaged samples to discard a frame.

  - decoding: set by user
  - encoding: unused
*/
func (s *AVCodecContext) SetDiscardDamagedPercentage(value int) {
	s.ptr.discard_damaged_percentage = (C.int)(value)
}

// MaxSamples gets the max_samples field.
/*
  The number of samples per frame to maximally accept.

  - decoding: set by user
  - encoding: set by user
*/
func (s *AVCodecContext) MaxSamples() int64 {
	value := s.ptr.max_samples
	return int64(value)
}

// SetMaxSamples sets the max_samples field.
/*
  The number of samples per frame to maximally accept.

  - decoding: set by user
  - encoding: set by user
*/
func (s *AVCodecContext) SetMaxSamples(value int64) {
	s.ptr.max_samples = (C.int64_t)(value)
}

// ExportSideData gets the export_side_data field.
/*
  Bit set of AV_CODEC_EXPORT_DATA_* flags, which affects the kind of
  metadata exported in frame, packet, or coded stream side data by
  decoders and encoders.

  - decoding: set by user
  - encoding: set by user
*/
func (s *AVCodecContext) ExportSideData() int {
	value := s.ptr.export_side_data
	return int(value)
}

// SetExportSideData sets the export_side_data field.
/*
  Bit set of AV_CODEC_EXPORT_DATA_* flags, which affects the kind of
  metadata exported in frame, packet, or coded stream side data by
  decoders and encoders.

  - decoding: set by user
  - encoding: set by user
*/
func (s *AVCodecContext) SetExportSideData(value int) {
	s.ptr.export_side_data = (C.int)(value)
}

// get_encode_buffer skipped due to func ptr

// ChLayout gets the ch_layout field.
/*
  Audio channel layout.
  - encoding: must be set by the caller, to one of AVCodec.ch_layouts.
  - decoding: may be set by the caller if known e.g. from the container.
              The decoder can then override during decoding as needed.
*/
func (s *AVCodecContext) ChLayout() *AVChannelLayout {
	value := &s.ptr.ch_layout
	return &AVChannelLayout{ptr: value}
}

// FrameNum gets the frame_num field.
/*
  Frame counter, set by libavcodec.

  - decoding: total number of frames returned from the decoder so far.
  - encoding: total number of frames passed to the encoder so far.

    @note the counter is not incremented if encoding/decoding resulted in
    an error.
*/
func (s *AVCodecContext) FrameNum() int64 {
	value := s.ptr.frame_num
	return int64(value)
}

// SetFrameNum sets the frame_num field.
/*
  Frame counter, set by libavcodec.

  - decoding: total number of frames returned from the decoder so far.
  - encoding: total number of frames passed to the encoder so far.

    @note the counter is not incremented if encoding/decoding resulted in
    an error.
*/
func (s *AVCodecContext) SetFrameNum(value int64) {
	s.ptr.frame_num = (C.int64_t)(value)
}

// --- Struct AVHWAccel ---

// AVHWAccel wraps AVHWAccel.
type AVHWAccel struct {
	ptr *C.AVHWAccel
}

func (s *AVHWAccel) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVHWAccelArray(ptr unsafe.Pointer) *Array[*AVHWAccel] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVHWAccel]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVHWAccel {
			ptr := (**C.AVHWAccel)(pointer)
			value := *ptr
			var valueMapped *AVHWAccel
			if value != nil {
				valueMapped = &AVHWAccel{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVHWAccel) {
			ptr := (**C.AVHWAccel)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Name gets the name field.
/*
  Name of the hardware accelerated codec.
  The name is globally unique among encoders and among decoders (but an
  encoder and a decoder can share the same name).
*/
func (s *AVHWAccel) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
/*
  Name of the hardware accelerated codec.
  The name is globally unique among encoders and among decoders (but an
  encoder and a decoder can share the same name).
*/
func (s *AVHWAccel) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// Type gets the type field.
/*
  Type of codec implemented by the hardware accelerator.

  See AVMEDIA_TYPE_xxx
*/
func (s *AVHWAccel) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}

// SetType sets the type field.
/*
  Type of codec implemented by the hardware accelerator.

  See AVMEDIA_TYPE_xxx
*/
func (s *AVHWAccel) SetType(value AVMediaType) {
	s.ptr._type = (C.enum_AVMediaType)(value)
}

// Id gets the id field.
/*
  Codec implemented by the hardware accelerator.

  See AV_CODEC_ID_xxx
*/
func (s *AVHWAccel) Id() AVCodecID {
	value := s.ptr.id
	return AVCodecID(value)
}

// SetId sets the id field.
/*
  Codec implemented by the hardware accelerator.

  See AV_CODEC_ID_xxx
*/
func (s *AVHWAccel) SetId(value AVCodecID) {
	s.ptr.id = (C.enum_AVCodecID)(value)
}

// PixFmt gets the pix_fmt field.
/*
  Supported pixel format.

  Only hardware accelerated formats are supported here.
*/
func (s *AVHWAccel) PixFmt() AVPixelFormat {
	value := s.ptr.pix_fmt
	return AVPixelFormat(value)
}

// SetPixFmt sets the pix_fmt field.
/*
  Supported pixel format.

  Only hardware accelerated formats are supported here.
*/
func (s *AVHWAccel) SetPixFmt(value AVPixelFormat) {
	s.ptr.pix_fmt = (C.enum_AVPixelFormat)(value)
}

// Capabilities gets the capabilities field.
/*
  Hardware accelerated codec capabilities.
  see AV_HWACCEL_CODEC_CAP_*
*/
func (s *AVHWAccel) Capabilities() int {
	value := s.ptr.capabilities
	return int(value)
}

// SetCapabilities sets the capabilities field.
/*
  Hardware accelerated codec capabilities.
  see AV_HWACCEL_CODEC_CAP_*
*/
func (s *AVHWAccel) SetCapabilities(value int) {
	s.ptr.capabilities = (C.int)(value)
}

// --- Struct AVSubtitleRect ---

// AVSubtitleRect wraps AVSubtitleRect.
type AVSubtitleRect struct {
	ptr *C.AVSubtitleRect
}

func (s *AVSubtitleRect) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVSubtitleRectArray(ptr unsafe.Pointer) *Array[*AVSubtitleRect] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVSubtitleRect]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVSubtitleRect {
			ptr := (**C.AVSubtitleRect)(pointer)
			value := *ptr
			var valueMapped *AVSubtitleRect
			if value != nil {
				valueMapped = &AVSubtitleRect{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVSubtitleRect) {
			ptr := (**C.AVSubtitleRect)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// X gets the x field.
//
//	top left corner  of pict, undefined when pict is not set
func (s *AVSubtitleRect) X() int {
	value := s.ptr.x
	return int(value)
}

// SetX sets the x field.
//
//	top left corner  of pict, undefined when pict is not set
func (s *AVSubtitleRect) SetX(value int) {
	s.ptr.x = (C.int)(value)
}

// Y gets the y field.
//
//	top left corner  of pict, undefined when pict is not set
func (s *AVSubtitleRect) Y() int {
	value := s.ptr.y
	return int(value)
}

// SetY sets the y field.
//
//	top left corner  of pict, undefined when pict is not set
func (s *AVSubtitleRect) SetY(value int) {
	s.ptr.y = (C.int)(value)
}

// W gets the w field.
//
//	width            of pict, undefined when pict is not set
func (s *AVSubtitleRect) W() int {
	value := s.ptr.w
	return int(value)
}

// SetW sets the w field.
//
//	width            of pict, undefined when pict is not set
func (s *AVSubtitleRect) SetW(value int) {
	s.ptr.w = (C.int)(value)
}

// H gets the h field.
//
//	height           of pict, undefined when pict is not set
func (s *AVSubtitleRect) H() int {
	value := s.ptr.h
	return int(value)
}

// SetH sets the h field.
//
//	height           of pict, undefined when pict is not set
func (s *AVSubtitleRect) SetH(value int) {
	s.ptr.h = (C.int)(value)
}

// NbColors gets the nb_colors field.
//
//	number of colors in pict, undefined when pict is not set
func (s *AVSubtitleRect) NbColors() int {
	value := s.ptr.nb_colors
	return int(value)
}

// SetNbColors sets the nb_colors field.
//
//	number of colors in pict, undefined when pict is not set
func (s *AVSubtitleRect) SetNbColors(value int) {
	s.ptr.nb_colors = (C.int)(value)
}

// Data gets the data field.
/*
  data+linesize for the bitmap of this subtitle.
  Can be set for text/ass as well once they are rendered.
*/
func (s *AVSubtitleRect) Data() *Array[unsafe.Pointer] {
	value := &s.ptr.data
	return ToUint8PtrArray(unsafe.Pointer(value))
}

// Linesize gets the linesize field.
func (s *AVSubtitleRect) Linesize() *Array[int] {
	value := &s.ptr.linesize
	return ToIntArray(unsafe.Pointer(value))
}

// Type gets the type field.
func (s *AVSubtitleRect) Type() AVSubtitleType {
	value := s.ptr._type
	return AVSubtitleType(value)
}

// SetType sets the type field.
func (s *AVSubtitleRect) SetType(value AVSubtitleType) {
	s.ptr._type = (C.enum_AVSubtitleType)(value)
}

// Text gets the text field.
//
//	0 terminated plain UTF-8 text
func (s *AVSubtitleRect) Text() *CStr {
	value := s.ptr.text
	return wrapCStr(value)
}

// SetText sets the text field.
//
//	0 terminated plain UTF-8 text
func (s *AVSubtitleRect) SetText(value *CStr) {
	s.ptr.text = value.ptr
}

// Ass gets the ass field.
/*
  0 terminated ASS/SSA compatible event line.
  The presentation of this is unaffected by the other values in this
  struct.
*/
func (s *AVSubtitleRect) Ass() *CStr {
	value := s.ptr.ass
	return wrapCStr(value)
}

// SetAss sets the ass field.
/*
  0 terminated ASS/SSA compatible event line.
  The presentation of this is unaffected by the other values in this
  struct.
*/
func (s *AVSubtitleRect) SetAss(value *CStr) {
	s.ptr.ass = value.ptr
}

// Flags gets the flags field.
func (s *AVSubtitleRect) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVSubtitleRect) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// --- Struct AVSubtitle ---

// AVSubtitle wraps AVSubtitle.
type AVSubtitle struct {
	ptr *C.AVSubtitle
}

func (s *AVSubtitle) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVSubtitleArray(ptr unsafe.Pointer) *Array[*AVSubtitle] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVSubtitle]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVSubtitle {
			ptr := (**C.AVSubtitle)(pointer)
			value := *ptr
			var valueMapped *AVSubtitle
			if value != nil {
				valueMapped = &AVSubtitle{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVSubtitle) {
			ptr := (**C.AVSubtitle)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Format gets the format field.
//
//	0 = graphics
func (s *AVSubtitle) Format() uint16 {
	value := s.ptr.format
	return uint16(value)
}

// SetFormat sets the format field.
//
//	0 = graphics
func (s *AVSubtitle) SetFormat(value uint16) {
	s.ptr.format = (C.uint16_t)(value)
}

// StartDisplayTime gets the start_display_time field.
//
//	relative to packet pts, in ms
func (s *AVSubtitle) StartDisplayTime() uint32 {
	value := s.ptr.start_display_time
	return uint32(value)
}

// SetStartDisplayTime sets the start_display_time field.
//
//	relative to packet pts, in ms
func (s *AVSubtitle) SetStartDisplayTime(value uint32) {
	s.ptr.start_display_time = (C.uint32_t)(value)
}

// EndDisplayTime gets the end_display_time field.
//
//	relative to packet pts, in ms
func (s *AVSubtitle) EndDisplayTime() uint32 {
	value := s.ptr.end_display_time
	return uint32(value)
}

// SetEndDisplayTime sets the end_display_time field.
//
//	relative to packet pts, in ms
func (s *AVSubtitle) SetEndDisplayTime(value uint32) {
	s.ptr.end_display_time = (C.uint32_t)(value)
}

// NumRects gets the num_rects field.
func (s *AVSubtitle) NumRects() uint {
	value := s.ptr.num_rects
	return uint(value)
}

// SetNumRects sets the num_rects field.
func (s *AVSubtitle) SetNumRects(value uint) {
	s.ptr.num_rects = (C.uint)(value)
}

// Rects gets the rects field.
func (s *AVSubtitle) Rects() *Array[*AVSubtitleRect] {
	value := s.ptr.rects
	return ToAVSubtitleRectArray(unsafe.Pointer(value))
}

// SetRects sets the rects field.
func (s *AVSubtitle) SetRects(value *Array[AVSubtitleRect]) {
	if value != nil {
		s.ptr.rects = (**C.AVSubtitleRect)(value.ptr)
	} else {
		s.ptr.rects = nil
	}
}

// Pts gets the pts field.
//
//	Same as packet pts, in AV_TIME_BASE
func (s *AVSubtitle) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

// SetPts sets the pts field.
//
//	Same as packet pts, in AV_TIME_BASE
func (s *AVSubtitle) SetPts(value int64) {
	s.ptr.pts = (C.int64_t)(value)
}

// --- Struct AVCodecParserContext ---

// AVCodecParserContext wraps AVCodecParserContext.
type AVCodecParserContext struct {
	ptr *C.AVCodecParserContext
}

func (s *AVCodecParserContext) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVCodecParserContextArray(ptr unsafe.Pointer) *Array[*AVCodecParserContext] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVCodecParserContext]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVCodecParserContext {
			ptr := (**C.AVCodecParserContext)(pointer)
			value := *ptr
			var valueMapped *AVCodecParserContext
			if value != nil {
				valueMapped = &AVCodecParserContext{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVCodecParserContext) {
			ptr := (**C.AVCodecParserContext)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// PrivData gets the priv_data field.
func (s *AVCodecParserContext) PrivData() unsafe.Pointer {
	value := s.ptr.priv_data
	return value
}

// SetPrivData sets the priv_data field.
func (s *AVCodecParserContext) SetPrivData(value unsafe.Pointer) {
	s.ptr.priv_data = value
}

// Parser gets the parser field.
func (s *AVCodecParserContext) Parser() *AVCodecParser {
	value := s.ptr.parser
	var valueMapped *AVCodecParser
	if value != nil {
		valueMapped = &AVCodecParser{ptr: value}
	}
	return valueMapped
}

// SetParser sets the parser field.
func (s *AVCodecParserContext) SetParser(value *AVCodecParser) {
	if value != nil {
		s.ptr.parser = value.ptr
	} else {
		s.ptr.parser = nil
	}
}

// FrameOffset gets the frame_offset field.
//
//	offset of the current frame
func (s *AVCodecParserContext) FrameOffset() int64 {
	value := s.ptr.frame_offset
	return int64(value)
}

// SetFrameOffset sets the frame_offset field.
//
//	offset of the current frame
func (s *AVCodecParserContext) SetFrameOffset(value int64) {
	s.ptr.frame_offset = (C.int64_t)(value)
}

// CurOffset gets the cur_offset field.
/*
  current offset
  (incremented by each av_parser_parse())
*/
func (s *AVCodecParserContext) CurOffset() int64 {
	value := s.ptr.cur_offset
	return int64(value)
}

// SetCurOffset sets the cur_offset field.
/*
  current offset
  (incremented by each av_parser_parse())
*/
func (s *AVCodecParserContext) SetCurOffset(value int64) {
	s.ptr.cur_offset = (C.int64_t)(value)
}

// NextFrameOffset gets the next_frame_offset field.
//
//	offset of the next frame
func (s *AVCodecParserContext) NextFrameOffset() int64 {
	value := s.ptr.next_frame_offset
	return int64(value)
}

// SetNextFrameOffset sets the next_frame_offset field.
//
//	offset of the next frame
func (s *AVCodecParserContext) SetNextFrameOffset(value int64) {
	s.ptr.next_frame_offset = (C.int64_t)(value)
}

// PictType gets the pict_type field.
//
//	XXX: Put it back in AVCodecContext.
func (s *AVCodecParserContext) PictType() int {
	value := s.ptr.pict_type
	return int(value)
}

// SetPictType sets the pict_type field.
//
//	XXX: Put it back in AVCodecContext.
func (s *AVCodecParserContext) SetPictType(value int) {
	s.ptr.pict_type = (C.int)(value)
}

// RepeatPict gets the repeat_pict field.
//
//	XXX: Put it back in AVCodecContext.
func (s *AVCodecParserContext) RepeatPict() int {
	value := s.ptr.repeat_pict
	return int(value)
}

// SetRepeatPict sets the repeat_pict field.
//
//	XXX: Put it back in AVCodecContext.
func (s *AVCodecParserContext) SetRepeatPict(value int) {
	s.ptr.repeat_pict = (C.int)(value)
}

// Pts gets the pts field.
//
//	pts of the current frame
func (s *AVCodecParserContext) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

// SetPts sets the pts field.
//
//	pts of the current frame
func (s *AVCodecParserContext) SetPts(value int64) {
	s.ptr.pts = (C.int64_t)(value)
}

// Dts gets the dts field.
//
//	dts of the current frame
func (s *AVCodecParserContext) Dts() int64 {
	value := s.ptr.dts
	return int64(value)
}

// SetDts sets the dts field.
//
//	dts of the current frame
func (s *AVCodecParserContext) SetDts(value int64) {
	s.ptr.dts = (C.int64_t)(value)
}

// LastPts gets the last_pts field.
//
//	private data
func (s *AVCodecParserContext) LastPts() int64 {
	value := s.ptr.last_pts
	return int64(value)
}

// SetLastPts sets the last_pts field.
//
//	private data
func (s *AVCodecParserContext) SetLastPts(value int64) {
	s.ptr.last_pts = (C.int64_t)(value)
}

// LastDts gets the last_dts field.
func (s *AVCodecParserContext) LastDts() int64 {
	value := s.ptr.last_dts
	return int64(value)
}

// SetLastDts sets the last_dts field.
func (s *AVCodecParserContext) SetLastDts(value int64) {
	s.ptr.last_dts = (C.int64_t)(value)
}

// FetchTimestamp gets the fetch_timestamp field.
func (s *AVCodecParserContext) FetchTimestamp() int {
	value := s.ptr.fetch_timestamp
	return int(value)
}

// SetFetchTimestamp sets the fetch_timestamp field.
func (s *AVCodecParserContext) SetFetchTimestamp(value int) {
	s.ptr.fetch_timestamp = (C.int)(value)
}

// CurFrameStartIndex gets the cur_frame_start_index field.
func (s *AVCodecParserContext) CurFrameStartIndex() int {
	value := s.ptr.cur_frame_start_index
	return int(value)
}

// SetCurFrameStartIndex sets the cur_frame_start_index field.
func (s *AVCodecParserContext) SetCurFrameStartIndex(value int) {
	s.ptr.cur_frame_start_index = (C.int)(value)
}

// CurFrameOffset gets the cur_frame_offset field.
func (s *AVCodecParserContext) CurFrameOffset() *Array[int64] {
	value := &s.ptr.cur_frame_offset
	return ToInt64Array(unsafe.Pointer(value))
}

// CurFramePts gets the cur_frame_pts field.
func (s *AVCodecParserContext) CurFramePts() *Array[int64] {
	value := &s.ptr.cur_frame_pts
	return ToInt64Array(unsafe.Pointer(value))
}

// CurFrameDts gets the cur_frame_dts field.
func (s *AVCodecParserContext) CurFrameDts() *Array[int64] {
	value := &s.ptr.cur_frame_dts
	return ToInt64Array(unsafe.Pointer(value))
}

// Flags gets the flags field.
func (s *AVCodecParserContext) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVCodecParserContext) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// Offset gets the offset field.
//
//	byte offset from starting packet start
func (s *AVCodecParserContext) Offset() int64 {
	value := s.ptr.offset
	return int64(value)
}

// SetOffset sets the offset field.
//
//	byte offset from starting packet start
func (s *AVCodecParserContext) SetOffset(value int64) {
	s.ptr.offset = (C.int64_t)(value)
}

// CurFrameEnd gets the cur_frame_end field.
func (s *AVCodecParserContext) CurFrameEnd() *Array[int64] {
	value := &s.ptr.cur_frame_end
	return ToInt64Array(unsafe.Pointer(value))
}

// KeyFrame gets the key_frame field.
/*
  Set by parser to 1 for key frames and 0 for non-key frames.
  It is initialized to -1, so if the parser doesn't set this flag,
  old-style fallback using AV_PICTURE_TYPE_I picture type as key frames
  will be used.
*/
func (s *AVCodecParserContext) KeyFrame() int {
	value := s.ptr.key_frame
	return int(value)
}

// SetKeyFrame sets the key_frame field.
/*
  Set by parser to 1 for key frames and 0 for non-key frames.
  It is initialized to -1, so if the parser doesn't set this flag,
  old-style fallback using AV_PICTURE_TYPE_I picture type as key frames
  will be used.
*/
func (s *AVCodecParserContext) SetKeyFrame(value int) {
	s.ptr.key_frame = (C.int)(value)
}

// DtsSyncPoint gets the dts_sync_point field.
/*
  Synchronization point for start of timestamp generation.

  Set to >0 for sync point, 0 for no sync point and <0 for undefined
  (default).

  For example, this corresponds to presence of H.264 buffering period
  SEI message.
*/
func (s *AVCodecParserContext) DtsSyncPoint() int {
	value := s.ptr.dts_sync_point
	return int(value)
}

// SetDtsSyncPoint sets the dts_sync_point field.
/*
  Synchronization point for start of timestamp generation.

  Set to >0 for sync point, 0 for no sync point and <0 for undefined
  (default).

  For example, this corresponds to presence of H.264 buffering period
  SEI message.
*/
func (s *AVCodecParserContext) SetDtsSyncPoint(value int) {
	s.ptr.dts_sync_point = (C.int)(value)
}

// DtsRefDtsDelta gets the dts_ref_dts_delta field.
/*
  Offset of the current timestamp against last timestamp sync point in
  units of AVCodecContext.time_base.

  Set to INT_MIN when dts_sync_point unused. Otherwise, it must
  contain a valid timestamp offset.

  Note that the timestamp of sync point has usually a nonzero
  dts_ref_dts_delta, which refers to the previous sync point. Offset of
  the next frame after timestamp sync point will be usually 1.

  For example, this corresponds to H.264 cpb_removal_delay.
*/
func (s *AVCodecParserContext) DtsRefDtsDelta() int {
	value := s.ptr.dts_ref_dts_delta
	return int(value)
}

// SetDtsRefDtsDelta sets the dts_ref_dts_delta field.
/*
  Offset of the current timestamp against last timestamp sync point in
  units of AVCodecContext.time_base.

  Set to INT_MIN when dts_sync_point unused. Otherwise, it must
  contain a valid timestamp offset.

  Note that the timestamp of sync point has usually a nonzero
  dts_ref_dts_delta, which refers to the previous sync point. Offset of
  the next frame after timestamp sync point will be usually 1.

  For example, this corresponds to H.264 cpb_removal_delay.
*/
func (s *AVCodecParserContext) SetDtsRefDtsDelta(value int) {
	s.ptr.dts_ref_dts_delta = (C.int)(value)
}

// PtsDtsDelta gets the pts_dts_delta field.
/*
  Presentation delay of current frame in units of AVCodecContext.time_base.

  Set to INT_MIN when dts_sync_point unused. Otherwise, it must
  contain valid non-negative timestamp delta (presentation time of a frame
  must not lie in the past).

  This delay represents the difference between decoding and presentation
  time of the frame.

  For example, this corresponds to H.264 dpb_output_delay.
*/
func (s *AVCodecParserContext) PtsDtsDelta() int {
	value := s.ptr.pts_dts_delta
	return int(value)
}

// SetPtsDtsDelta sets the pts_dts_delta field.
/*
  Presentation delay of current frame in units of AVCodecContext.time_base.

  Set to INT_MIN when dts_sync_point unused. Otherwise, it must
  contain valid non-negative timestamp delta (presentation time of a frame
  must not lie in the past).

  This delay represents the difference between decoding and presentation
  time of the frame.

  For example, this corresponds to H.264 dpb_output_delay.
*/
func (s *AVCodecParserContext) SetPtsDtsDelta(value int) {
	s.ptr.pts_dts_delta = (C.int)(value)
}

// CurFramePos gets the cur_frame_pos field.
/*
  Position of the packet in file.

  Analogous to cur_frame_pts/dts
*/
func (s *AVCodecParserContext) CurFramePos() *Array[int64] {
	value := &s.ptr.cur_frame_pos
	return ToInt64Array(unsafe.Pointer(value))
}

// Pos gets the pos field.
//
//	Byte position of currently parsed frame in stream.
func (s *AVCodecParserContext) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}

// SetPos sets the pos field.
//
//	Byte position of currently parsed frame in stream.
func (s *AVCodecParserContext) SetPos(value int64) {
	s.ptr.pos = (C.int64_t)(value)
}

// LastPos gets the last_pos field.
//
//	Previous frame byte position.
func (s *AVCodecParserContext) LastPos() int64 {
	value := s.ptr.last_pos
	return int64(value)
}

// SetLastPos sets the last_pos field.
//
//	Previous frame byte position.
func (s *AVCodecParserContext) SetLastPos(value int64) {
	s.ptr.last_pos = (C.int64_t)(value)
}

// Duration gets the duration field.
/*
  Duration of the current frame.
  For audio, this is in units of 1 / AVCodecContext.sample_rate.
  For all other types, this is in units of AVCodecContext.time_base.
*/
func (s *AVCodecParserContext) Duration() int {
	value := s.ptr.duration
	return int(value)
}

// SetDuration sets the duration field.
/*
  Duration of the current frame.
  For audio, this is in units of 1 / AVCodecContext.sample_rate.
  For all other types, this is in units of AVCodecContext.time_base.
*/
func (s *AVCodecParserContext) SetDuration(value int) {
	s.ptr.duration = (C.int)(value)
}

// FieldOrder gets the field_order field.
func (s *AVCodecParserContext) FieldOrder() AVFieldOrder {
	value := s.ptr.field_order
	return AVFieldOrder(value)
}

// SetFieldOrder sets the field_order field.
func (s *AVCodecParserContext) SetFieldOrder(value AVFieldOrder) {
	s.ptr.field_order = (C.enum_AVFieldOrder)(value)
}

// PictureStructure gets the picture_structure field.
/*
  Indicate whether a picture is coded as a frame, top field or bottom field.

  For example, H.264 field_pic_flag equal to 0 corresponds to
  AV_PICTURE_STRUCTURE_FRAME. An H.264 picture with field_pic_flag
  equal to 1 and bottom_field_flag equal to 0 corresponds to
  AV_PICTURE_STRUCTURE_TOP_FIELD.
*/
func (s *AVCodecParserContext) PictureStructure() AVPictureStructure {
	value := s.ptr.picture_structure
	return AVPictureStructure(value)
}

// SetPictureStructure sets the picture_structure field.
/*
  Indicate whether a picture is coded as a frame, top field or bottom field.

  For example, H.264 field_pic_flag equal to 0 corresponds to
  AV_PICTURE_STRUCTURE_FRAME. An H.264 picture with field_pic_flag
  equal to 1 and bottom_field_flag equal to 0 corresponds to
  AV_PICTURE_STRUCTURE_TOP_FIELD.
*/
func (s *AVCodecParserContext) SetPictureStructure(value AVPictureStructure) {
	s.ptr.picture_structure = (C.enum_AVPictureStructure)(value)
}

// OutputPictureNumber gets the output_picture_number field.
/*
  Picture number incremented in presentation or output order.
  This field may be reinitialized at the first picture of a new sequence.

  For example, this corresponds to H.264 PicOrderCnt.
*/
func (s *AVCodecParserContext) OutputPictureNumber() int {
	value := s.ptr.output_picture_number
	return int(value)
}

// SetOutputPictureNumber sets the output_picture_number field.
/*
  Picture number incremented in presentation or output order.
  This field may be reinitialized at the first picture of a new sequence.

  For example, this corresponds to H.264 PicOrderCnt.
*/
func (s *AVCodecParserContext) SetOutputPictureNumber(value int) {
	s.ptr.output_picture_number = (C.int)(value)
}

// Width gets the width field.
//
//	Dimensions of the decoded video intended for presentation.
func (s *AVCodecParserContext) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
//
//	Dimensions of the decoded video intended for presentation.
func (s *AVCodecParserContext) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

// Height gets the height field.
func (s *AVCodecParserContext) Height() int {
	value := s.ptr.height
	return int(value)
}

// SetHeight sets the height field.
func (s *AVCodecParserContext) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

// CodedWidth gets the coded_width field.
//
//	Dimensions of the coded video.
func (s *AVCodecParserContext) CodedWidth() int {
	value := s.ptr.coded_width
	return int(value)
}

// SetCodedWidth sets the coded_width field.
//
//	Dimensions of the coded video.
func (s *AVCodecParserContext) SetCodedWidth(value int) {
	s.ptr.coded_width = (C.int)(value)
}

// CodedHeight gets the coded_height field.
func (s *AVCodecParserContext) CodedHeight() int {
	value := s.ptr.coded_height
	return int(value)
}

// SetCodedHeight sets the coded_height field.
func (s *AVCodecParserContext) SetCodedHeight(value int) {
	s.ptr.coded_height = (C.int)(value)
}

// Format gets the format field.
/*
  The format of the coded data, corresponds to enum AVPixelFormat for video
  and for enum AVSampleFormat for audio.

  Note that a decoder can have considerable freedom in how exactly it
  decodes the data, so the format reported here might be different from the
  one returned by a decoder.
*/
func (s *AVCodecParserContext) Format() int {
	value := s.ptr.format
	return int(value)
}

// SetFormat sets the format field.
/*
  The format of the coded data, corresponds to enum AVPixelFormat for video
  and for enum AVSampleFormat for audio.

  Note that a decoder can have considerable freedom in how exactly it
  decodes the data, so the format reported here might be different from the
  one returned by a decoder.
*/
func (s *AVCodecParserContext) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

// --- Struct AVCodecParser ---

// AVCodecParser wraps AVCodecParser.
type AVCodecParser struct {
	ptr *C.AVCodecParser
}

func (s *AVCodecParser) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVCodecParserArray(ptr unsafe.Pointer) *Array[*AVCodecParser] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVCodecParser]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVCodecParser {
			ptr := (**C.AVCodecParser)(pointer)
			value := *ptr
			var valueMapped *AVCodecParser
			if value != nil {
				valueMapped = &AVCodecParser{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVCodecParser) {
			ptr := (**C.AVCodecParser)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// CodecIds gets the codec_ids field.
//
//	several codec IDs are permitted
func (s *AVCodecParser) CodecIds() *Array[int] {
	value := &s.ptr.codec_ids
	return ToIntArray(unsafe.Pointer(value))
}

// PrivDataSize gets the priv_data_size field.
func (s *AVCodecParser) PrivDataSize() int {
	value := s.ptr.priv_data_size
	return int(value)
}

// SetPrivDataSize sets the priv_data_size field.
func (s *AVCodecParser) SetPrivDataSize(value int) {
	s.ptr.priv_data_size = (C.int)(value)
}

// parser_init skipped due to func ptr

// parser_parse skipped due to func ptr

// parser_close skipped due to func ptr

// split skipped due to func ptr

// --- Struct AVProfile ---

// AVProfile wraps AVProfile.
//
//	AVProfile.
type AVProfile struct {
	ptr *C.AVProfile
}

func (s *AVProfile) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVProfileArray(ptr unsafe.Pointer) *Array[*AVProfile] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVProfile]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVProfile {
			ptr := (**C.AVProfile)(pointer)
			value := *ptr
			var valueMapped *AVProfile
			if value != nil {
				valueMapped = &AVProfile{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVProfile) {
			ptr := (**C.AVProfile)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Profile gets the profile field.
func (s *AVProfile) Profile() int {
	value := s.ptr.profile
	return int(value)
}

// SetProfile sets the profile field.
func (s *AVProfile) SetProfile(value int) {
	s.ptr.profile = (C.int)(value)
}

// Name gets the name field.
//
//	short name for the profile
func (s *AVProfile) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
//
//	short name for the profile
func (s *AVProfile) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// --- Struct AVCodec ---

// AVCodec wraps AVCodec.
//
//	AVCodec.
type AVCodec struct {
	ptr *C.AVCodec
}

func (s *AVCodec) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVCodecArray(ptr unsafe.Pointer) *Array[*AVCodec] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVCodec]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVCodec {
			ptr := (**C.AVCodec)(pointer)
			value := *ptr
			var valueMapped *AVCodec
			if value != nil {
				valueMapped = &AVCodec{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVCodec) {
			ptr := (**C.AVCodec)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Name gets the name field.
/*
  Name of the codec implementation.
  The name is globally unique among encoders and among decoders (but an
  encoder and a decoder can share the same name).
  This is the primary way to find a codec from the user perspective.
*/
func (s *AVCodec) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
/*
  Name of the codec implementation.
  The name is globally unique among encoders and among decoders (but an
  encoder and a decoder can share the same name).
  This is the primary way to find a codec from the user perspective.
*/
func (s *AVCodec) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// LongName gets the long_name field.
/*
  Descriptive name for the codec, meant to be more human readable than name.
  You should use the NULL_IF_CONFIG_SMALL() macro to define it.
*/
func (s *AVCodec) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

// SetLongName sets the long_name field.
/*
  Descriptive name for the codec, meant to be more human readable than name.
  You should use the NULL_IF_CONFIG_SMALL() macro to define it.
*/
func (s *AVCodec) SetLongName(value *CStr) {
	s.ptr.long_name = value.ptr
}

// Type gets the type field.
func (s *AVCodec) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}

// SetType sets the type field.
func (s *AVCodec) SetType(value AVMediaType) {
	s.ptr._type = (C.enum_AVMediaType)(value)
}

// Id gets the id field.
func (s *AVCodec) Id() AVCodecID {
	value := s.ptr.id
	return AVCodecID(value)
}

// SetId sets the id field.
func (s *AVCodec) SetId(value AVCodecID) {
	s.ptr.id = (C.enum_AVCodecID)(value)
}

// Capabilities gets the capabilities field.
/*
  Codec capabilities.
  see AV_CODEC_CAP_*
*/
func (s *AVCodec) Capabilities() int {
	value := s.ptr.capabilities
	return int(value)
}

// SetCapabilities sets the capabilities field.
/*
  Codec capabilities.
  see AV_CODEC_CAP_*
*/
func (s *AVCodec) SetCapabilities(value int) {
	s.ptr.capabilities = (C.int)(value)
}

// MaxLowres gets the max_lowres field.
//
//	maximum value for lowres supported by the decoder
func (s *AVCodec) MaxLowres() uint8 {
	value := s.ptr.max_lowres
	return uint8(value)
}

// SetMaxLowres sets the max_lowres field.
//
//	maximum value for lowres supported by the decoder
func (s *AVCodec) SetMaxLowres(value uint8) {
	s.ptr.max_lowres = (C.uint8_t)(value)
}

// supported_framerates skipped due to struct value ptr

// PixFmts gets the pix_fmts field.
//
//	array of supported pixel formats, or NULL if unknown, array is terminated by -1
func (s *AVCodec) PixFmts() *Array[AVPixelFormat] {
	value := s.ptr.pix_fmts
	return ToAVPixelFormatArray(unsafe.Pointer(value))
}

// SetPixFmts sets the pix_fmts field.
//
//	array of supported pixel formats, or NULL if unknown, array is terminated by -1
func (s *AVCodec) SetPixFmts(value *Array[AVPixelFormat]) {
	if value != nil {
		s.ptr.pix_fmts = (*C.enum_AVPixelFormat)(value.ptr)
	} else {
		s.ptr.pix_fmts = nil
	}
}

// supported_samplerates skipped due to prim ptr

// SampleFmts gets the sample_fmts field.
//
//	array of supported sample formats, or NULL if unknown, array is terminated by -1
func (s *AVCodec) SampleFmts() *Array[AVSampleFormat] {
	value := s.ptr.sample_fmts
	return ToAVSampleFormatArray(unsafe.Pointer(value))
}

// SetSampleFmts sets the sample_fmts field.
//
//	array of supported sample formats, or NULL if unknown, array is terminated by -1
func (s *AVCodec) SetSampleFmts(value *Array[AVSampleFormat]) {
	if value != nil {
		s.ptr.sample_fmts = (*C.enum_AVSampleFormat)(value.ptr)
	} else {
		s.ptr.sample_fmts = nil
	}
}

// channel_layouts skipped due to prim ptr

// PrivClass gets the priv_class field.
//
//	AVClass for the private context
func (s *AVCodec) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetPrivClass sets the priv_class field.
//
//	AVClass for the private context
func (s *AVCodec) SetPrivClass(value *AVClass) {
	if value != nil {
		s.ptr.priv_class = value.ptr
	} else {
		s.ptr.priv_class = nil
	}
}

// Profiles gets the profiles field.
//
//	array of recognized profiles, or NULL if unknown, array is terminated by {AV_PROFILE_UNKNOWN}
func (s *AVCodec) Profiles() *AVProfile {
	value := s.ptr.profiles
	var valueMapped *AVProfile
	if value != nil {
		valueMapped = &AVProfile{ptr: value}
	}
	return valueMapped
}

// SetProfiles sets the profiles field.
//
//	array of recognized profiles, or NULL if unknown, array is terminated by {AV_PROFILE_UNKNOWN}
func (s *AVCodec) SetProfiles(value *AVProfile) {
	if value != nil {
		s.ptr.profiles = value.ptr
	} else {
		s.ptr.profiles = nil
	}
}

// WrapperName gets the wrapper_name field.
/*
  Group name of the codec implementation.
  This is a short symbolic name of the wrapper backing this codec. A
  wrapper uses some kind of external implementation for the codec, such
  as an external library, or a codec implementation provided by the OS or
  the hardware.
  If this field is NULL, this is a builtin, libavcodec native codec.
  If non-NULL, this will be the suffix in AVCodec.name in most cases
  (usually AVCodec.name will be of the form "<codec_name>_<wrapper_name>").
*/
func (s *AVCodec) WrapperName() *CStr {
	value := s.ptr.wrapper_name
	return wrapCStr(value)
}

// SetWrapperName sets the wrapper_name field.
/*
  Group name of the codec implementation.
  This is a short symbolic name of the wrapper backing this codec. A
  wrapper uses some kind of external implementation for the codec, such
  as an external library, or a codec implementation provided by the OS or
  the hardware.
  If this field is NULL, this is a builtin, libavcodec native codec.
  If non-NULL, this will be the suffix in AVCodec.name in most cases
  (usually AVCodec.name will be of the form "<codec_name>_<wrapper_name>").
*/
func (s *AVCodec) SetWrapperName(value *CStr) {
	s.ptr.wrapper_name = value.ptr
}

// ChLayouts gets the ch_layouts field.
//
//	Array of supported channel layouts, terminated with a zeroed layout.
func (s *AVCodec) ChLayouts() *AVChannelLayout {
	value := s.ptr.ch_layouts
	var valueMapped *AVChannelLayout
	if value != nil {
		valueMapped = &AVChannelLayout{ptr: value}
	}
	return valueMapped
}

// SetChLayouts sets the ch_layouts field.
//
//	Array of supported channel layouts, terminated with a zeroed layout.
func (s *AVCodec) SetChLayouts(value *AVChannelLayout) {
	if value != nil {
		s.ptr.ch_layouts = value.ptr
	} else {
		s.ptr.ch_layouts = nil
	}
}

// --- Struct AVCodecHWConfig ---

// AVCodecHWConfig wraps AVCodecHWConfig.
type AVCodecHWConfig struct {
	ptr *C.AVCodecHWConfig
}

func (s *AVCodecHWConfig) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVCodecHWConfigArray(ptr unsafe.Pointer) *Array[*AVCodecHWConfig] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVCodecHWConfig]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVCodecHWConfig {
			ptr := (**C.AVCodecHWConfig)(pointer)
			value := *ptr
			var valueMapped *AVCodecHWConfig
			if value != nil {
				valueMapped = &AVCodecHWConfig{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVCodecHWConfig) {
			ptr := (**C.AVCodecHWConfig)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// PixFmt gets the pix_fmt field.
/*
  For decoders, a hardware pixel format which that decoder may be
  able to decode to if suitable hardware is available.

  For encoders, a pixel format which the encoder may be able to
  accept.  If set to AV_PIX_FMT_NONE, this applies to all pixel
  formats supported by the codec.
*/
func (s *AVCodecHWConfig) PixFmt() AVPixelFormat {
	value := s.ptr.pix_fmt
	return AVPixelFormat(value)
}

// SetPixFmt sets the pix_fmt field.
/*
  For decoders, a hardware pixel format which that decoder may be
  able to decode to if suitable hardware is available.

  For encoders, a pixel format which the encoder may be able to
  accept.  If set to AV_PIX_FMT_NONE, this applies to all pixel
  formats supported by the codec.
*/
func (s *AVCodecHWConfig) SetPixFmt(value AVPixelFormat) {
	s.ptr.pix_fmt = (C.enum_AVPixelFormat)(value)
}

// Methods gets the methods field.
/*
  Bit set of AV_CODEC_HW_CONFIG_METHOD_* flags, describing the possible
  setup methods which can be used with this configuration.
*/
func (s *AVCodecHWConfig) Methods() int {
	value := s.ptr.methods
	return int(value)
}

// SetMethods sets the methods field.
/*
  Bit set of AV_CODEC_HW_CONFIG_METHOD_* flags, describing the possible
  setup methods which can be used with this configuration.
*/
func (s *AVCodecHWConfig) SetMethods(value int) {
	s.ptr.methods = (C.int)(value)
}

// DeviceType gets the device_type field.
/*
  The device type associated with the configuration.

  Must be set for AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX and
  AV_CODEC_HW_CONFIG_METHOD_HW_FRAMES_CTX, otherwise unused.
*/
func (s *AVCodecHWConfig) DeviceType() AVHWDeviceType {
	value := s.ptr.device_type
	return AVHWDeviceType(value)
}

// SetDeviceType sets the device_type field.
/*
  The device type associated with the configuration.

  Must be set for AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX and
  AV_CODEC_HW_CONFIG_METHOD_HW_FRAMES_CTX, otherwise unused.
*/
func (s *AVCodecHWConfig) SetDeviceType(value AVHWDeviceType) {
	s.ptr.device_type = (C.enum_AVHWDeviceType)(value)
}

// --- Struct AVCodecDescriptor ---

// AVCodecDescriptor wraps AVCodecDescriptor.
/*
  This struct describes the properties of a single codec described by an
  AVCodecID.
  @see avcodec_descriptor_get()
*/
type AVCodecDescriptor struct {
	ptr *C.AVCodecDescriptor
}

func (s *AVCodecDescriptor) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVCodecDescriptorArray(ptr unsafe.Pointer) *Array[*AVCodecDescriptor] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVCodecDescriptor]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVCodecDescriptor {
			ptr := (**C.AVCodecDescriptor)(pointer)
			value := *ptr
			var valueMapped *AVCodecDescriptor
			if value != nil {
				valueMapped = &AVCodecDescriptor{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVCodecDescriptor) {
			ptr := (**C.AVCodecDescriptor)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Id gets the id field.
func (s *AVCodecDescriptor) Id() AVCodecID {
	value := s.ptr.id
	return AVCodecID(value)
}

// SetId sets the id field.
func (s *AVCodecDescriptor) SetId(value AVCodecID) {
	s.ptr.id = (C.enum_AVCodecID)(value)
}

// Type gets the type field.
func (s *AVCodecDescriptor) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}

// SetType sets the type field.
func (s *AVCodecDescriptor) SetType(value AVMediaType) {
	s.ptr._type = (C.enum_AVMediaType)(value)
}

// Name gets the name field.
/*
  Name of the codec described by this descriptor. It is non-empty and
  unique for each codec descriptor. It should contain alphanumeric
  characters and '_' only.
*/
func (s *AVCodecDescriptor) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
/*
  Name of the codec described by this descriptor. It is non-empty and
  unique for each codec descriptor. It should contain alphanumeric
  characters and '_' only.
*/
func (s *AVCodecDescriptor) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// LongName gets the long_name field.
//
//	A more descriptive name for this codec. May be NULL.
func (s *AVCodecDescriptor) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

// SetLongName sets the long_name field.
//
//	A more descriptive name for this codec. May be NULL.
func (s *AVCodecDescriptor) SetLongName(value *CStr) {
	s.ptr.long_name = value.ptr
}

// Props gets the props field.
//
//	Codec properties, a combination of AV_CODEC_PROP_* flags.
func (s *AVCodecDescriptor) Props() int {
	value := s.ptr.props
	return int(value)
}

// SetProps sets the props field.
//
//	Codec properties, a combination of AV_CODEC_PROP_* flags.
func (s *AVCodecDescriptor) SetProps(value int) {
	s.ptr.props = (C.int)(value)
}

// mime_types skipped due to unknown ptr ptr

// Profiles gets the profiles field.
/*
  If non-NULL, an array of profiles recognized for this codec.
  Terminated with AV_PROFILE_UNKNOWN.
*/
func (s *AVCodecDescriptor) Profiles() *AVProfile {
	value := s.ptr.profiles
	var valueMapped *AVProfile
	if value != nil {
		valueMapped = &AVProfile{ptr: value}
	}
	return valueMapped
}

// SetProfiles sets the profiles field.
/*
  If non-NULL, an array of profiles recognized for this codec.
  Terminated with AV_PROFILE_UNKNOWN.
*/
func (s *AVCodecDescriptor) SetProfiles(value *AVProfile) {
	if value != nil {
		s.ptr.profiles = value.ptr
	} else {
		s.ptr.profiles = nil
	}
}

// --- Struct AVCodecParameters ---

// AVCodecParameters wraps AVCodecParameters.
/*
  This struct describes the properties of an encoded stream.

  sizeof(AVCodecParameters) is not a part of the public ABI, this struct must
  be allocated with avcodec_parameters_alloc() and freed with
  avcodec_parameters_free().
*/
type AVCodecParameters struct {
	ptr *C.AVCodecParameters
}

func (s *AVCodecParameters) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVCodecParametersArray(ptr unsafe.Pointer) *Array[*AVCodecParameters] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVCodecParameters]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVCodecParameters {
			ptr := (**C.AVCodecParameters)(pointer)
			value := *ptr
			var valueMapped *AVCodecParameters
			if value != nil {
				valueMapped = &AVCodecParameters{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVCodecParameters) {
			ptr := (**C.AVCodecParameters)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// CodecType gets the codec_type field.
//
//	General type of the encoded data.
func (s *AVCodecParameters) CodecType() AVMediaType {
	value := s.ptr.codec_type
	return AVMediaType(value)
}

// SetCodecType sets the codec_type field.
//
//	General type of the encoded data.
func (s *AVCodecParameters) SetCodecType(value AVMediaType) {
	s.ptr.codec_type = (C.enum_AVMediaType)(value)
}

// CodecId gets the codec_id field.
//
//	Specific type of the encoded data (the codec used).
func (s *AVCodecParameters) CodecId() AVCodecID {
	value := s.ptr.codec_id
	return AVCodecID(value)
}

// SetCodecId sets the codec_id field.
//
//	Specific type of the encoded data (the codec used).
func (s *AVCodecParameters) SetCodecId(value AVCodecID) {
	s.ptr.codec_id = (C.enum_AVCodecID)(value)
}

// CodecTag gets the codec_tag field.
//
//	Additional information about the codec (corresponds to the AVI FOURCC).
func (s *AVCodecParameters) CodecTag() uint32 {
	value := s.ptr.codec_tag
	return uint32(value)
}

// SetCodecTag sets the codec_tag field.
//
//	Additional information about the codec (corresponds to the AVI FOURCC).
func (s *AVCodecParameters) SetCodecTag(value uint32) {
	s.ptr.codec_tag = (C.uint32_t)(value)
}

// Extradata gets the extradata field.
/*
  Extra binary data needed for initializing the decoder, codec-dependent.

  Must be allocated with av_malloc() and will be freed by
  avcodec_parameters_free(). The allocated size of extradata must be at
  least extradata_size + AV_INPUT_BUFFER_PADDING_SIZE, with the padding
  bytes zeroed.
*/
func (s *AVCodecParameters) Extradata() unsafe.Pointer {
	value := s.ptr.extradata
	return unsafe.Pointer(value)
}

// SetExtradata sets the extradata field.
/*
  Extra binary data needed for initializing the decoder, codec-dependent.

  Must be allocated with av_malloc() and will be freed by
  avcodec_parameters_free(). The allocated size of extradata must be at
  least extradata_size + AV_INPUT_BUFFER_PADDING_SIZE, with the padding
  bytes zeroed.
*/
func (s *AVCodecParameters) SetExtradata(value unsafe.Pointer) {
	s.ptr.extradata = (*C.uint8_t)(value)
}

// ExtradataSize gets the extradata_size field.
//
//	Size of the extradata content in bytes.
func (s *AVCodecParameters) ExtradataSize() int {
	value := s.ptr.extradata_size
	return int(value)
}

// SetExtradataSize sets the extradata_size field.
//
//	Size of the extradata content in bytes.
func (s *AVCodecParameters) SetExtradataSize(value int) {
	s.ptr.extradata_size = (C.int)(value)
}

// Format gets the format field.
/*
  - video: the pixel format, the value corresponds to enum AVPixelFormat.
  - audio: the sample format, the value corresponds to enum AVSampleFormat.
*/
func (s *AVCodecParameters) Format() int {
	value := s.ptr.format
	return int(value)
}

// SetFormat sets the format field.
/*
  - video: the pixel format, the value corresponds to enum AVPixelFormat.
  - audio: the sample format, the value corresponds to enum AVSampleFormat.
*/
func (s *AVCodecParameters) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

// BitRate gets the bit_rate field.
//
//	The average bitrate of the encoded data (in bits per second).
func (s *AVCodecParameters) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}

// SetBitRate sets the bit_rate field.
//
//	The average bitrate of the encoded data (in bits per second).
func (s *AVCodecParameters) SetBitRate(value int64) {
	s.ptr.bit_rate = (C.int64_t)(value)
}

// BitsPerCodedSample gets the bits_per_coded_sample field.
/*
  The number of bits per sample in the codedwords.

  This is basically the bitrate per sample. It is mandatory for a bunch of
  formats to actually decode them. It's the number of bits for one sample in
  the actual coded bitstream.

  This could be for example 4 for ADPCM
  For PCM formats this matches bits_per_raw_sample
  Can be 0
*/
func (s *AVCodecParameters) BitsPerCodedSample() int {
	value := s.ptr.bits_per_coded_sample
	return int(value)
}

// SetBitsPerCodedSample sets the bits_per_coded_sample field.
/*
  The number of bits per sample in the codedwords.

  This is basically the bitrate per sample. It is mandatory for a bunch of
  formats to actually decode them. It's the number of bits for one sample in
  the actual coded bitstream.

  This could be for example 4 for ADPCM
  For PCM formats this matches bits_per_raw_sample
  Can be 0
*/
func (s *AVCodecParameters) SetBitsPerCodedSample(value int) {
	s.ptr.bits_per_coded_sample = (C.int)(value)
}

// BitsPerRawSample gets the bits_per_raw_sample field.
/*
  This is the number of valid bits in each output sample. If the
  sample format has more bits, the least significant bits are additional
  padding bits, which are always 0. Use right shifts to reduce the sample
  to its actual size. For example, audio formats with 24 bit samples will
  have bits_per_raw_sample set to 24, and format set to AV_SAMPLE_FMT_S32.
  To get the original sample use "(int32_t)sample >> 8"."

  For ADPCM this might be 12 or 16 or similar
  Can be 0
*/
func (s *AVCodecParameters) BitsPerRawSample() int {
	value := s.ptr.bits_per_raw_sample
	return int(value)
}

// SetBitsPerRawSample sets the bits_per_raw_sample field.
/*
  This is the number of valid bits in each output sample. If the
  sample format has more bits, the least significant bits are additional
  padding bits, which are always 0. Use right shifts to reduce the sample
  to its actual size. For example, audio formats with 24 bit samples will
  have bits_per_raw_sample set to 24, and format set to AV_SAMPLE_FMT_S32.
  To get the original sample use "(int32_t)sample >> 8"."

  For ADPCM this might be 12 or 16 or similar
  Can be 0
*/
func (s *AVCodecParameters) SetBitsPerRawSample(value int) {
	s.ptr.bits_per_raw_sample = (C.int)(value)
}

// Profile gets the profile field.
//
//	Codec-specific bitstream restrictions that the stream conforms to.
func (s *AVCodecParameters) Profile() int {
	value := s.ptr.profile
	return int(value)
}

// SetProfile sets the profile field.
//
//	Codec-specific bitstream restrictions that the stream conforms to.
func (s *AVCodecParameters) SetProfile(value int) {
	s.ptr.profile = (C.int)(value)
}

// Level gets the level field.
func (s *AVCodecParameters) Level() int {
	value := s.ptr.level
	return int(value)
}

// SetLevel sets the level field.
func (s *AVCodecParameters) SetLevel(value int) {
	s.ptr.level = (C.int)(value)
}

// Width gets the width field.
//
//	Video only. The dimensions of the video frame in pixels.
func (s *AVCodecParameters) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
//
//	Video only. The dimensions of the video frame in pixels.
func (s *AVCodecParameters) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

// Height gets the height field.
func (s *AVCodecParameters) Height() int {
	value := s.ptr.height
	return int(value)
}

// SetHeight sets the height field.
func (s *AVCodecParameters) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

// SampleAspectRatio gets the sample_aspect_ratio field.
/*
  Video only. The aspect ratio (width / height) which a single pixel
  should have when displayed.

  When the aspect ratio is unknown / undefined, the numerator should be
  set to 0 (the denominator may have any value).
*/
func (s *AVCodecParameters) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
/*
  Video only. The aspect ratio (width / height) which a single pixel
  should have when displayed.

  When the aspect ratio is unknown / undefined, the numerator should be
  set to 0 (the denominator may have any value).
*/
func (s *AVCodecParameters) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

// FieldOrder gets the field_order field.
//
//	Video only. The order of the fields in interlaced video.
func (s *AVCodecParameters) FieldOrder() AVFieldOrder {
	value := s.ptr.field_order
	return AVFieldOrder(value)
}

// SetFieldOrder sets the field_order field.
//
//	Video only. The order of the fields in interlaced video.
func (s *AVCodecParameters) SetFieldOrder(value AVFieldOrder) {
	s.ptr.field_order = (C.enum_AVFieldOrder)(value)
}

// ColorRange gets the color_range field.
//
//	Video only. Additional colorspace characteristics.
func (s *AVCodecParameters) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}

// SetColorRange sets the color_range field.
//
//	Video only. Additional colorspace characteristics.
func (s *AVCodecParameters) SetColorRange(value AVColorRange) {
	s.ptr.color_range = (C.enum_AVColorRange)(value)
}

// ColorPrimaries gets the color_primaries field.
func (s *AVCodecParameters) ColorPrimaries() AVColorPrimaries {
	value := s.ptr.color_primaries
	return AVColorPrimaries(value)
}

// SetColorPrimaries sets the color_primaries field.
func (s *AVCodecParameters) SetColorPrimaries(value AVColorPrimaries) {
	s.ptr.color_primaries = (C.enum_AVColorPrimaries)(value)
}

// ColorTrc gets the color_trc field.
func (s *AVCodecParameters) ColorTrc() AVColorTransferCharacteristic {
	value := s.ptr.color_trc
	return AVColorTransferCharacteristic(value)
}

// SetColorTrc sets the color_trc field.
func (s *AVCodecParameters) SetColorTrc(value AVColorTransferCharacteristic) {
	s.ptr.color_trc = (C.enum_AVColorTransferCharacteristic)(value)
}

// ColorSpace gets the color_space field.
func (s *AVCodecParameters) ColorSpace() AVColorSpace {
	value := s.ptr.color_space
	return AVColorSpace(value)
}

// SetColorSpace sets the color_space field.
func (s *AVCodecParameters) SetColorSpace(value AVColorSpace) {
	s.ptr.color_space = (C.enum_AVColorSpace)(value)
}

// ChromaLocation gets the chroma_location field.
func (s *AVCodecParameters) ChromaLocation() AVChromaLocation {
	value := s.ptr.chroma_location
	return AVChromaLocation(value)
}

// SetChromaLocation sets the chroma_location field.
func (s *AVCodecParameters) SetChromaLocation(value AVChromaLocation) {
	s.ptr.chroma_location = (C.enum_AVChromaLocation)(value)
}

// VideoDelay gets the video_delay field.
//
//	Video only. Number of delayed frames.
func (s *AVCodecParameters) VideoDelay() int {
	value := s.ptr.video_delay
	return int(value)
}

// SetVideoDelay sets the video_delay field.
//
//	Video only. Number of delayed frames.
func (s *AVCodecParameters) SetVideoDelay(value int) {
	s.ptr.video_delay = (C.int)(value)
}

// ChannelLayout gets the channel_layout field.
/*
  Audio only. The channel layout bitmask. May be 0 if the channel layout is
  unknown or unspecified, otherwise the number of bits set must be equal to
  the channels field.
  @deprecated use ch_layout
*/
func (s *AVCodecParameters) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

// SetChannelLayout sets the channel_layout field.
/*
  Audio only. The channel layout bitmask. May be 0 if the channel layout is
  unknown or unspecified, otherwise the number of bits set must be equal to
  the channels field.
  @deprecated use ch_layout
*/
func (s *AVCodecParameters) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

// Channels gets the channels field.
/*
  Audio only. The number of audio channels.
  @deprecated use ch_layout.nb_channels
*/
func (s *AVCodecParameters) Channels() int {
	value := s.ptr.channels
	return int(value)
}

// SetChannels sets the channels field.
/*
  Audio only. The number of audio channels.
  @deprecated use ch_layout.nb_channels
*/
func (s *AVCodecParameters) SetChannels(value int) {
	s.ptr.channels = (C.int)(value)
}

// SampleRate gets the sample_rate field.
//
//	Audio only. The number of audio samples per second.
func (s *AVCodecParameters) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

// SetSampleRate sets the sample_rate field.
//
//	Audio only. The number of audio samples per second.
func (s *AVCodecParameters) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

// BlockAlign gets the block_align field.
/*
  Audio only. The number of bytes per coded audio frame, required by some
  formats.

  Corresponds to nBlockAlign in WAVEFORMATEX.
*/
func (s *AVCodecParameters) BlockAlign() int {
	value := s.ptr.block_align
	return int(value)
}

// SetBlockAlign sets the block_align field.
/*
  Audio only. The number of bytes per coded audio frame, required by some
  formats.

  Corresponds to nBlockAlign in WAVEFORMATEX.
*/
func (s *AVCodecParameters) SetBlockAlign(value int) {
	s.ptr.block_align = (C.int)(value)
}

// FrameSize gets the frame_size field.
//
//	Audio only. Audio frame size, if known. Required by some formats to be static.
func (s *AVCodecParameters) FrameSize() int {
	value := s.ptr.frame_size
	return int(value)
}

// SetFrameSize sets the frame_size field.
//
//	Audio only. Audio frame size, if known. Required by some formats to be static.
func (s *AVCodecParameters) SetFrameSize(value int) {
	s.ptr.frame_size = (C.int)(value)
}

// InitialPadding gets the initial_padding field.
/*
  Audio only. The amount of padding (in samples) inserted by the encoder at
  the beginning of the audio. I.e. this number of leading decoded samples
  must be discarded by the caller to get the original audio without leading
  padding.
*/
func (s *AVCodecParameters) InitialPadding() int {
	value := s.ptr.initial_padding
	return int(value)
}

// SetInitialPadding sets the initial_padding field.
/*
  Audio only. The amount of padding (in samples) inserted by the encoder at
  the beginning of the audio. I.e. this number of leading decoded samples
  must be discarded by the caller to get the original audio without leading
  padding.
*/
func (s *AVCodecParameters) SetInitialPadding(value int) {
	s.ptr.initial_padding = (C.int)(value)
}

// TrailingPadding gets the trailing_padding field.
/*
  Audio only. The amount of padding (in samples) appended by the encoder to
  the end of the audio. I.e. this number of decoded samples must be
  discarded by the caller from the end of the stream to get the original
  audio without any trailing padding.
*/
func (s *AVCodecParameters) TrailingPadding() int {
	value := s.ptr.trailing_padding
	return int(value)
}

// SetTrailingPadding sets the trailing_padding field.
/*
  Audio only. The amount of padding (in samples) appended by the encoder to
  the end of the audio. I.e. this number of decoded samples must be
  discarded by the caller from the end of the stream to get the original
  audio without any trailing padding.
*/
func (s *AVCodecParameters) SetTrailingPadding(value int) {
	s.ptr.trailing_padding = (C.int)(value)
}

// SeekPreroll gets the seek_preroll field.
//
//	Audio only. Number of samples to skip after a discontinuity.
func (s *AVCodecParameters) SeekPreroll() int {
	value := s.ptr.seek_preroll
	return int(value)
}

// SetSeekPreroll sets the seek_preroll field.
//
//	Audio only. Number of samples to skip after a discontinuity.
func (s *AVCodecParameters) SetSeekPreroll(value int) {
	s.ptr.seek_preroll = (C.int)(value)
}

// ChLayout gets the ch_layout field.
//
//	Audio only. The channel layout and number of channels.
func (s *AVCodecParameters) ChLayout() *AVChannelLayout {
	value := &s.ptr.ch_layout
	return &AVChannelLayout{ptr: value}
}

// Framerate gets the framerate field.
/*
  Video only. Number of frames per second, for streams with constant frame
  durations. Should be set to { 0, 1 } when some frames have differing
  durations or if the value is not known.

  @note This field correponds to values that are stored in codec-level
  headers and is typically overridden by container/transport-layer
  timestamps, when available. It should thus be used only as a last resort,
  when no higher-level timing information is available.
*/
func (s *AVCodecParameters) Framerate() *AVRational {
	value := s.ptr.framerate
	return &AVRational{value: value}
}

// SetFramerate sets the framerate field.
/*
  Video only. Number of frames per second, for streams with constant frame
  durations. Should be set to { 0, 1 } when some frames have differing
  durations or if the value is not known.

  @note This field correponds to values that are stored in codec-level
  headers and is typically overridden by container/transport-layer
  timestamps, when available. It should thus be used only as a last resort,
  when no higher-level timing information is available.
*/
func (s *AVCodecParameters) SetFramerate(value *AVRational) {
	s.ptr.framerate = value.value
}

// CodedSideData gets the coded_side_data field.
//
//	Additional data associated with the entire stream.
func (s *AVCodecParameters) CodedSideData() *AVPacketSideData {
	value := s.ptr.coded_side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}

// SetCodedSideData sets the coded_side_data field.
//
//	Additional data associated with the entire stream.
func (s *AVCodecParameters) SetCodedSideData(value *AVPacketSideData) {
	if value != nil {
		s.ptr.coded_side_data = value.ptr
	} else {
		s.ptr.coded_side_data = nil
	}
}

// NbCodedSideData gets the nb_coded_side_data field.
//
//	Amount of entries in @ref coded_side_data.
func (s *AVCodecParameters) NbCodedSideData() int {
	value := s.ptr.nb_coded_side_data
	return int(value)
}

// SetNbCodedSideData sets the nb_coded_side_data field.
//
//	Amount of entries in @ref coded_side_data.
func (s *AVCodecParameters) SetNbCodedSideData(value int) {
	s.ptr.nb_coded_side_data = (C.int)(value)
}

// --- Struct AVPanScan ---

// AVPanScan wraps AVPanScan.
/*
  Pan Scan area.
  This specifies the area which should be displayed.
  Note there may be multiple such areas for one frame.
*/
type AVPanScan struct {
	ptr *C.AVPanScan
}

func (s *AVPanScan) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVPanScanArray(ptr unsafe.Pointer) *Array[*AVPanScan] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVPanScan]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVPanScan {
			ptr := (**C.AVPanScan)(pointer)
			value := *ptr
			var valueMapped *AVPanScan
			if value != nil {
				valueMapped = &AVPanScan{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVPanScan) {
			ptr := (**C.AVPanScan)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Id gets the id field.
/*
  id
  - encoding: Set by user.
  - decoding: Set by libavcodec.
*/
func (s *AVPanScan) Id() int {
	value := s.ptr.id
	return int(value)
}

// SetId sets the id field.
/*
  id
  - encoding: Set by user.
  - decoding: Set by libavcodec.
*/
func (s *AVPanScan) SetId(value int) {
	s.ptr.id = (C.int)(value)
}

// Width gets the width field.
/*
  width and height in 1/16 pel
  - encoding: Set by user.
  - decoding: Set by libavcodec.
*/
func (s *AVPanScan) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
/*
  width and height in 1/16 pel
  - encoding: Set by user.
  - decoding: Set by libavcodec.
*/
func (s *AVPanScan) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

// Height gets the height field.
func (s *AVPanScan) Height() int {
	value := s.ptr.height
	return int(value)
}

// SetHeight sets the height field.
func (s *AVPanScan) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

// position skipped due to multi dim const array

// --- Struct AVCPBProperties ---

// AVCPBProperties wraps AVCPBProperties.
/*
  This structure describes the bitrate properties of an encoded bitstream. It
  roughly corresponds to a subset the VBV parameters for MPEG-2 or HRD
  parameters for H.264/HEVC.
*/
type AVCPBProperties struct {
	ptr *C.AVCPBProperties
}

func (s *AVCPBProperties) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVCPBPropertiesArray(ptr unsafe.Pointer) *Array[*AVCPBProperties] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVCPBProperties]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVCPBProperties {
			ptr := (**C.AVCPBProperties)(pointer)
			value := *ptr
			var valueMapped *AVCPBProperties
			if value != nil {
				valueMapped = &AVCPBProperties{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVCPBProperties) {
			ptr := (**C.AVCPBProperties)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// MaxBitrate gets the max_bitrate field.
/*
  Maximum bitrate of the stream, in bits per second.
  Zero if unknown or unspecified.
*/
func (s *AVCPBProperties) MaxBitrate() int64 {
	value := s.ptr.max_bitrate
	return int64(value)
}

// SetMaxBitrate sets the max_bitrate field.
/*
  Maximum bitrate of the stream, in bits per second.
  Zero if unknown or unspecified.
*/
func (s *AVCPBProperties) SetMaxBitrate(value int64) {
	s.ptr.max_bitrate = (C.int64_t)(value)
}

// MinBitrate gets the min_bitrate field.
/*
  Minimum bitrate of the stream, in bits per second.
  Zero if unknown or unspecified.
*/
func (s *AVCPBProperties) MinBitrate() int64 {
	value := s.ptr.min_bitrate
	return int64(value)
}

// SetMinBitrate sets the min_bitrate field.
/*
  Minimum bitrate of the stream, in bits per second.
  Zero if unknown or unspecified.
*/
func (s *AVCPBProperties) SetMinBitrate(value int64) {
	s.ptr.min_bitrate = (C.int64_t)(value)
}

// AvgBitrate gets the avg_bitrate field.
/*
  Average bitrate of the stream, in bits per second.
  Zero if unknown or unspecified.
*/
func (s *AVCPBProperties) AvgBitrate() int64 {
	value := s.ptr.avg_bitrate
	return int64(value)
}

// SetAvgBitrate sets the avg_bitrate field.
/*
  Average bitrate of the stream, in bits per second.
  Zero if unknown or unspecified.
*/
func (s *AVCPBProperties) SetAvgBitrate(value int64) {
	s.ptr.avg_bitrate = (C.int64_t)(value)
}

// BufferSize gets the buffer_size field.
/*
  The size of the buffer to which the ratecontrol is applied, in bits.
  Zero if unknown or unspecified.
*/
func (s *AVCPBProperties) BufferSize() int64 {
	value := s.ptr.buffer_size
	return int64(value)
}

// SetBufferSize sets the buffer_size field.
/*
  The size of the buffer to which the ratecontrol is applied, in bits.
  Zero if unknown or unspecified.
*/
func (s *AVCPBProperties) SetBufferSize(value int64) {
	s.ptr.buffer_size = (C.int64_t)(value)
}

// VbvDelay gets the vbv_delay field.
/*
  The delay between the time the packet this structure is associated with
  is received and the time when it should be decoded, in periods of a 27MHz
  clock.

  UINT64_MAX when unknown or unspecified.
*/
func (s *AVCPBProperties) VbvDelay() uint64 {
	value := s.ptr.vbv_delay
	return uint64(value)
}

// SetVbvDelay sets the vbv_delay field.
/*
  The delay between the time the packet this structure is associated with
  is received and the time when it should be decoded, in periods of a 27MHz
  clock.

  UINT64_MAX when unknown or unspecified.
*/
func (s *AVCPBProperties) SetVbvDelay(value uint64) {
	s.ptr.vbv_delay = (C.uint64_t)(value)
}

// --- Struct AVProducerReferenceTime ---

// AVProducerReferenceTime wraps AVProducerReferenceTime.
/*
  This structure supplies correlation between a packet timestamp and a wall clock
  production time. The definition follows the Producer Reference Time ('prft')
  as defined in ISO/IEC 14496-12
*/
type AVProducerReferenceTime struct {
	ptr *C.AVProducerReferenceTime
}

func (s *AVProducerReferenceTime) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVProducerReferenceTimeArray(ptr unsafe.Pointer) *Array[*AVProducerReferenceTime] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVProducerReferenceTime]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVProducerReferenceTime {
			ptr := (**C.AVProducerReferenceTime)(pointer)
			value := *ptr
			var valueMapped *AVProducerReferenceTime
			if value != nil {
				valueMapped = &AVProducerReferenceTime{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVProducerReferenceTime) {
			ptr := (**C.AVProducerReferenceTime)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Wallclock gets the wallclock field.
//
//	A UTC timestamp, in microseconds, since Unix epoch (e.g, av_gettime()).
func (s *AVProducerReferenceTime) Wallclock() int64 {
	value := s.ptr.wallclock
	return int64(value)
}

// SetWallclock sets the wallclock field.
//
//	A UTC timestamp, in microseconds, since Unix epoch (e.g, av_gettime()).
func (s *AVProducerReferenceTime) SetWallclock(value int64) {
	s.ptr.wallclock = (C.int64_t)(value)
}

// Flags gets the flags field.
func (s *AVProducerReferenceTime) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVProducerReferenceTime) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// --- Struct AVPacketSideData ---

// AVPacketSideData wraps AVPacketSideData.
/*
  This structure stores auxiliary information for decoding, presenting, or
  otherwise processing the coded stream. It is typically exported by demuxers
  and encoders and can be fed to decoders and muxers either in a per packet
  basis, or as global side data (applying to the entire coded stream).

  Global side data is handled as follows:
  - During demuxing, it may be exported through
    @ref AVStream.codecpar.side_data "AVStream's codec parameters", which can
    then be passed as input to decoders through the
    @ref AVCodecContext.coded_side_data "decoder context's side data", for
    initialization.
  - For muxing, it can be fed through @ref AVStream.codecpar.side_data
    "AVStream's codec parameters", typically  the output of encoders through
    the @ref AVCodecContext.coded_side_data "encoder context's side data", for
    initialization.

  Packet specific side data is handled as follows:
  - During demuxing, it may be exported through @ref AVPacket.side_data
    "AVPacket's side data", which can then be passed as input to decoders.
  - For muxing, it can be fed through @ref AVPacket.side_data "AVPacket's
    side data", typically the output of encoders.

  Different modules may accept or export different types of side data
  depending on media type and codec. Refer to @ref AVPacketSideDataType for a
  list of defined types and where they may be found or used.
*/
type AVPacketSideData struct {
	ptr *C.AVPacketSideData
}

func (s *AVPacketSideData) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVPacketSideDataArray(ptr unsafe.Pointer) *Array[*AVPacketSideData] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVPacketSideData]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVPacketSideData {
			ptr := (**C.AVPacketSideData)(pointer)
			value := *ptr
			var valueMapped *AVPacketSideData
			if value != nil {
				valueMapped = &AVPacketSideData{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVPacketSideData) {
			ptr := (**C.AVPacketSideData)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Data gets the data field.
func (s *AVPacketSideData) Data() unsafe.Pointer {
	value := s.ptr.data
	return unsafe.Pointer(value)
}

// SetData sets the data field.
func (s *AVPacketSideData) SetData(value unsafe.Pointer) {
	s.ptr.data = (*C.uint8_t)(value)
}

// Size gets the size field.
func (s *AVPacketSideData) Size() uint64 {
	value := s.ptr.size
	return uint64(value)
}

// SetSize sets the size field.
func (s *AVPacketSideData) SetSize(value uint64) {
	s.ptr.size = (C.size_t)(value)
}

// Type gets the type field.
func (s *AVPacketSideData) Type() AVPacketSideDataType {
	value := s.ptr._type
	return AVPacketSideDataType(value)
}

// SetType sets the type field.
func (s *AVPacketSideData) SetType(value AVPacketSideDataType) {
	s.ptr._type = (C.enum_AVPacketSideDataType)(value)
}

// --- Struct AVPacket ---

// AVPacket wraps AVPacket.
/*
  This structure stores compressed data. It is typically exported by demuxers
  and then passed as input to decoders, or received as output from encoders and
  then passed to muxers.

  For video, it should typically contain one compressed frame. For audio it may
  contain several compressed frames. Encoders are allowed to output empty
  packets, with no compressed data, containing only side data
  (e.g. to update some stream parameters at the end of encoding).

  The semantics of data ownership depends on the buf field.
  If it is set, the packet data is dynamically allocated and is
  valid indefinitely until a call to av_packet_unref() reduces the
  reference count to 0.

  If the buf field is not set av_packet_ref() would make a copy instead
  of increasing the reference count.

  The side data is always allocated with av_malloc(), copied by
  av_packet_ref() and freed by av_packet_unref().

  sizeof(AVPacket) being a part of the public ABI is deprecated. once
  av_init_packet() is removed, new packets will only be able to be allocated
  with av_packet_alloc(), and new fields may be added to the end of the struct
  with a minor bump.

  @see av_packet_alloc
  @see av_packet_ref
  @see av_packet_unref
*/
type AVPacket struct {
	ptr *C.AVPacket
}

func (s *AVPacket) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVPacketArray(ptr unsafe.Pointer) *Array[*AVPacket] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVPacket]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVPacket {
			ptr := (**C.AVPacket)(pointer)
			value := *ptr
			var valueMapped *AVPacket
			if value != nil {
				valueMapped = &AVPacket{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVPacket) {
			ptr := (**C.AVPacket)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Buf gets the buf field.
/*
  A reference to the reference-counted buffer where the packet data is
  stored.
  May be NULL, then the packet data is not reference-counted.
*/
func (s *AVPacket) Buf() *AVBufferRef {
	value := s.ptr.buf
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetBuf sets the buf field.
/*
  A reference to the reference-counted buffer where the packet data is
  stored.
  May be NULL, then the packet data is not reference-counted.
*/
func (s *AVPacket) SetBuf(value *AVBufferRef) {
	if value != nil {
		s.ptr.buf = value.ptr
	} else {
		s.ptr.buf = nil
	}
}

// Pts gets the pts field.
/*
  Presentation timestamp in AVStream->time_base units; the time at which
  the decompressed packet will be presented to the user.
  Can be AV_NOPTS_VALUE if it is not stored in the file.
  pts MUST be larger or equal to dts as presentation cannot happen before
  decompression, unless one wants to view hex dumps. Some formats misuse
  the terms dts and pts/cts to mean something different. Such timestamps
  must be converted to true pts/dts before they are stored in AVPacket.
*/
func (s *AVPacket) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

// SetPts sets the pts field.
/*
  Presentation timestamp in AVStream->time_base units; the time at which
  the decompressed packet will be presented to the user.
  Can be AV_NOPTS_VALUE if it is not stored in the file.
  pts MUST be larger or equal to dts as presentation cannot happen before
  decompression, unless one wants to view hex dumps. Some formats misuse
  the terms dts and pts/cts to mean something different. Such timestamps
  must be converted to true pts/dts before they are stored in AVPacket.
*/
func (s *AVPacket) SetPts(value int64) {
	s.ptr.pts = (C.int64_t)(value)
}

// Dts gets the dts field.
/*
  Decompression timestamp in AVStream->time_base units; the time at which
  the packet is decompressed.
  Can be AV_NOPTS_VALUE if it is not stored in the file.
*/
func (s *AVPacket) Dts() int64 {
	value := s.ptr.dts
	return int64(value)
}

// SetDts sets the dts field.
/*
  Decompression timestamp in AVStream->time_base units; the time at which
  the packet is decompressed.
  Can be AV_NOPTS_VALUE if it is not stored in the file.
*/
func (s *AVPacket) SetDts(value int64) {
	s.ptr.dts = (C.int64_t)(value)
}

// Data gets the data field.
func (s *AVPacket) Data() unsafe.Pointer {
	value := s.ptr.data
	return unsafe.Pointer(value)
}

// SetData sets the data field.
func (s *AVPacket) SetData(value unsafe.Pointer) {
	s.ptr.data = (*C.uint8_t)(value)
}

// Size gets the size field.
func (s *AVPacket) Size() int {
	value := s.ptr.size
	return int(value)
}

// SetSize sets the size field.
func (s *AVPacket) SetSize(value int) {
	s.ptr.size = (C.int)(value)
}

// StreamIndex gets the stream_index field.
func (s *AVPacket) StreamIndex() int {
	value := s.ptr.stream_index
	return int(value)
}

// SetStreamIndex sets the stream_index field.
func (s *AVPacket) SetStreamIndex(value int) {
	s.ptr.stream_index = (C.int)(value)
}

// Flags gets the flags field.
//
//	A combination of AV_PKT_FLAG values
func (s *AVPacket) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
//
//	A combination of AV_PKT_FLAG values
func (s *AVPacket) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// SideData gets the side_data field.
/*
  Additional packet data that can be provided by the container.
  Packet can contain several types of side information.
*/
func (s *AVPacket) SideData() *AVPacketSideData {
	value := s.ptr.side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}

// SetSideData sets the side_data field.
/*
  Additional packet data that can be provided by the container.
  Packet can contain several types of side information.
*/
func (s *AVPacket) SetSideData(value *AVPacketSideData) {
	if value != nil {
		s.ptr.side_data = value.ptr
	} else {
		s.ptr.side_data = nil
	}
}

// SideDataElems gets the side_data_elems field.
func (s *AVPacket) SideDataElems() int {
	value := s.ptr.side_data_elems
	return int(value)
}

// SetSideDataElems sets the side_data_elems field.
func (s *AVPacket) SetSideDataElems(value int) {
	s.ptr.side_data_elems = (C.int)(value)
}

// Duration gets the duration field.
/*
  Duration of this packet in AVStream->time_base units, 0 if unknown.
  Equals next_pts - this_pts in presentation order.
*/
func (s *AVPacket) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

// SetDuration sets the duration field.
/*
  Duration of this packet in AVStream->time_base units, 0 if unknown.
  Equals next_pts - this_pts in presentation order.
*/
func (s *AVPacket) SetDuration(value int64) {
	s.ptr.duration = (C.int64_t)(value)
}

// Pos gets the pos field.
//
//	byte position in stream, -1 if unknown
func (s *AVPacket) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}

// SetPos sets the pos field.
//
//	byte position in stream, -1 if unknown
func (s *AVPacket) SetPos(value int64) {
	s.ptr.pos = (C.int64_t)(value)
}

// Opaque gets the opaque field.
//
//	for some private data of the user
func (s *AVPacket) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
//
//	for some private data of the user
func (s *AVPacket) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// OpaqueRef gets the opaque_ref field.
/*
  AVBufferRef for free use by the API user. FFmpeg will never check the
  contents of the buffer ref. FFmpeg calls av_buffer_unref() on it when
  the packet is unreferenced. av_packet_copy_props() calls create a new
  reference with av_buffer_ref() for the target packet's opaque_ref field.

  This is unrelated to the opaque field, although it serves a similar
  purpose.
*/
func (s *AVPacket) OpaqueRef() *AVBufferRef {
	value := s.ptr.opaque_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetOpaqueRef sets the opaque_ref field.
/*
  AVBufferRef for free use by the API user. FFmpeg will never check the
  contents of the buffer ref. FFmpeg calls av_buffer_unref() on it when
  the packet is unreferenced. av_packet_copy_props() calls create a new
  reference with av_buffer_ref() for the target packet's opaque_ref field.

  This is unrelated to the opaque field, although it serves a similar
  purpose.
*/
func (s *AVPacket) SetOpaqueRef(value *AVBufferRef) {
	if value != nil {
		s.ptr.opaque_ref = value.ptr
	} else {
		s.ptr.opaque_ref = nil
	}
}

// TimeBase gets the time_base field.
/*
  Time base of the packet's timestamps.
  In the future, this field may be set on packets output by encoders or
  demuxers, but its value will be by default ignored on input to decoders
  or muxers.
*/
func (s *AVPacket) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
/*
  Time base of the packet's timestamps.
  In the future, this field may be set on packets output by encoders or
  demuxers, but its value will be by default ignored on input to decoders
  or muxers.
*/
func (s *AVPacket) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// --- Struct AVPacketList ---

// AVPacketList wraps AVPacketList.
type AVPacketList struct {
	ptr *C.AVPacketList
}

func (s *AVPacketList) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVPacketListArray(ptr unsafe.Pointer) *Array[*AVPacketList] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVPacketList]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVPacketList {
			ptr := (**C.AVPacketList)(pointer)
			value := *ptr
			var valueMapped *AVPacketList
			if value != nil {
				valueMapped = &AVPacketList{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVPacketList) {
			ptr := (**C.AVPacketList)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Pkt gets the pkt field.
func (s *AVPacketList) Pkt() *AVPacket {
	value := &s.ptr.pkt
	return &AVPacket{ptr: value}
}

// Next gets the next field.
func (s *AVPacketList) Next() *AVPacketList {
	value := s.ptr.next
	var valueMapped *AVPacketList
	if value != nil {
		valueMapped = &AVPacketList{ptr: value}
	}
	return valueMapped
}

// SetNext sets the next field.
func (s *AVPacketList) SetNext(value *AVPacketList) {
	if value != nil {
		s.ptr.next = value.ptr
	} else {
		s.ptr.next = nil
	}
}

// --- Struct AVFilterContext ---

// AVFilterContext wraps AVFilterContext.
//
//	An instance of a filter
type AVFilterContext struct {
	ptr *C.AVFilterContext
}

func (s *AVFilterContext) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterContextArray(ptr unsafe.Pointer) *Array[*AVFilterContext] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterContext]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterContext {
			ptr := (**C.AVFilterContext)(pointer)
			value := *ptr
			var valueMapped *AVFilterContext
			if value != nil {
				valueMapped = &AVFilterContext{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterContext) {
			ptr := (**C.AVFilterContext)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// AvClass gets the av_class field.
//
//	needed for av_log() and filters common options
func (s *AVFilterContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
//
//	needed for av_log() and filters common options
func (s *AVFilterContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Filter gets the filter field.
//
//	the AVFilter of which this is an instance
func (s *AVFilterContext) Filter() *AVFilter {
	value := s.ptr.filter
	var valueMapped *AVFilter
	if value != nil {
		valueMapped = &AVFilter{ptr: value}
	}
	return valueMapped
}

// SetFilter sets the filter field.
//
//	the AVFilter of which this is an instance
func (s *AVFilterContext) SetFilter(value *AVFilter) {
	if value != nil {
		s.ptr.filter = value.ptr
	} else {
		s.ptr.filter = nil
	}
}

// Name gets the name field.
//
//	name of this filter instance
func (s *AVFilterContext) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
//
//	name of this filter instance
func (s *AVFilterContext) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// InputPads gets the input_pads field.
//
//	array of input pads
func (s *AVFilterContext) InputPads() *AVFilterPad {
	value := s.ptr.input_pads
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetInputPads sets the input_pads field.
//
//	array of input pads
func (s *AVFilterContext) SetInputPads(value *AVFilterPad) {
	if value != nil {
		s.ptr.input_pads = value.ptr
	} else {
		s.ptr.input_pads = nil
	}
}

// Inputs gets the inputs field.
//
//	array of pointers to input links
func (s *AVFilterContext) Inputs() *Array[*AVFilterLink] {
	value := s.ptr.inputs
	return ToAVFilterLinkArray(unsafe.Pointer(value))
}

// SetInputs sets the inputs field.
//
//	array of pointers to input links
func (s *AVFilterContext) SetInputs(value *Array[AVFilterLink]) {
	if value != nil {
		s.ptr.inputs = (**C.AVFilterLink)(value.ptr)
	} else {
		s.ptr.inputs = nil
	}
}

// NbInputs gets the nb_inputs field.
//
//	number of input pads
func (s *AVFilterContext) NbInputs() uint {
	value := s.ptr.nb_inputs
	return uint(value)
}

// SetNbInputs sets the nb_inputs field.
//
//	number of input pads
func (s *AVFilterContext) SetNbInputs(value uint) {
	s.ptr.nb_inputs = (C.uint)(value)
}

// OutputPads gets the output_pads field.
//
//	array of output pads
func (s *AVFilterContext) OutputPads() *AVFilterPad {
	value := s.ptr.output_pads
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetOutputPads sets the output_pads field.
//
//	array of output pads
func (s *AVFilterContext) SetOutputPads(value *AVFilterPad) {
	if value != nil {
		s.ptr.output_pads = value.ptr
	} else {
		s.ptr.output_pads = nil
	}
}

// Outputs gets the outputs field.
//
//	array of pointers to output links
func (s *AVFilterContext) Outputs() *Array[*AVFilterLink] {
	value := s.ptr.outputs
	return ToAVFilterLinkArray(unsafe.Pointer(value))
}

// SetOutputs sets the outputs field.
//
//	array of pointers to output links
func (s *AVFilterContext) SetOutputs(value *Array[AVFilterLink]) {
	if value != nil {
		s.ptr.outputs = (**C.AVFilterLink)(value.ptr)
	} else {
		s.ptr.outputs = nil
	}
}

// NbOutputs gets the nb_outputs field.
//
//	number of output pads
func (s *AVFilterContext) NbOutputs() uint {
	value := s.ptr.nb_outputs
	return uint(value)
}

// SetNbOutputs sets the nb_outputs field.
//
//	number of output pads
func (s *AVFilterContext) SetNbOutputs(value uint) {
	s.ptr.nb_outputs = (C.uint)(value)
}

// Priv gets the priv field.
//
//	private data for use by the filter
func (s *AVFilterContext) Priv() unsafe.Pointer {
	value := s.ptr.priv
	return value
}

// SetPriv sets the priv field.
//
//	private data for use by the filter
func (s *AVFilterContext) SetPriv(value unsafe.Pointer) {
	s.ptr.priv = value
}

// Graph gets the graph field.
//
//	filtergraph this filter belongs to
func (s *AVFilterContext) Graph() *AVFilterGraph {
	value := s.ptr.graph
	var valueMapped *AVFilterGraph
	if value != nil {
		valueMapped = &AVFilterGraph{ptr: value}
	}
	return valueMapped
}

// SetGraph sets the graph field.
//
//	filtergraph this filter belongs to
func (s *AVFilterContext) SetGraph(value *AVFilterGraph) {
	if value != nil {
		s.ptr.graph = value.ptr
	} else {
		s.ptr.graph = nil
	}
}

// ThreadType gets the thread_type field.
/*
  Type of multithreading being allowed/used. A combination of
  AVFILTER_THREAD_* flags.

  May be set by the caller before initializing the filter to forbid some
  or all kinds of multithreading for this filter. The default is allowing
  everything.

  When the filter is initialized, this field is combined using bit AND with
  AVFilterGraph.thread_type to get the final mask used for determining
  allowed threading types. I.e. a threading type needs to be set in both
  to be allowed.

  After the filter is initialized, libavfilter sets this field to the
  threading type that is actually used (0 for no multithreading).
*/
func (s *AVFilterContext) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}

// SetThreadType sets the thread_type field.
/*
  Type of multithreading being allowed/used. A combination of
  AVFILTER_THREAD_* flags.

  May be set by the caller before initializing the filter to forbid some
  or all kinds of multithreading for this filter. The default is allowing
  everything.

  When the filter is initialized, this field is combined using bit AND with
  AVFilterGraph.thread_type to get the final mask used for determining
  allowed threading types. I.e. a threading type needs to be set in both
  to be allowed.

  After the filter is initialized, libavfilter sets this field to the
  threading type that is actually used (0 for no multithreading).
*/
func (s *AVFilterContext) SetThreadType(value int) {
	s.ptr.thread_type = (C.int)(value)
}

// Internal gets the internal field.
//
//	An opaque struct for libavfilter internal use.
func (s *AVFilterContext) Internal() *AVFilterInternal {
	value := s.ptr.internal
	var valueMapped *AVFilterInternal
	if value != nil {
		valueMapped = &AVFilterInternal{ptr: value}
	}
	return valueMapped
}

// SetInternal sets the internal field.
//
//	An opaque struct for libavfilter internal use.
func (s *AVFilterContext) SetInternal(value *AVFilterInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

// command_queue skipped due to ptr to ignored type

// EnableStr gets the enable_str field.
//
//	enable expression string
func (s *AVFilterContext) EnableStr() *CStr {
	value := s.ptr.enable_str
	return wrapCStr(value)
}

// SetEnableStr sets the enable_str field.
//
//	enable expression string
func (s *AVFilterContext) SetEnableStr(value *CStr) {
	s.ptr.enable_str = value.ptr
}

// Enable gets the enable field.
//
//	parsed expression (AVExpr*)
func (s *AVFilterContext) Enable() unsafe.Pointer {
	value := s.ptr.enable
	return value
}

// SetEnable sets the enable field.
//
//	parsed expression (AVExpr*)
func (s *AVFilterContext) SetEnable(value unsafe.Pointer) {
	s.ptr.enable = value
}

// var_values skipped due to prim ptr

// IsDisabled gets the is_disabled field.
//
//	the enabled state from the last expression evaluation
func (s *AVFilterContext) IsDisabled() int {
	value := s.ptr.is_disabled
	return int(value)
}

// SetIsDisabled sets the is_disabled field.
//
//	the enabled state from the last expression evaluation
func (s *AVFilterContext) SetIsDisabled(value int) {
	s.ptr.is_disabled = (C.int)(value)
}

// HwDeviceCtx gets the hw_device_ctx field.
/*
  For filters which will create hardware frames, sets the device the
  filter should create them in.  All other filters will ignore this field:
  in particular, a filter which consumes or processes hardware frames will
  instead use the hw_frames_ctx field in AVFilterLink to carry the
  hardware context information.

  May be set by the caller on filters flagged with AVFILTER_FLAG_HWDEVICE
  before initializing the filter with avfilter_init_str() or
  avfilter_init_dict().
*/
func (s *AVFilterContext) HwDeviceCtx() *AVBufferRef {
	value := s.ptr.hw_device_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwDeviceCtx sets the hw_device_ctx field.
/*
  For filters which will create hardware frames, sets the device the
  filter should create them in.  All other filters will ignore this field:
  in particular, a filter which consumes or processes hardware frames will
  instead use the hw_frames_ctx field in AVFilterLink to carry the
  hardware context information.

  May be set by the caller on filters flagged with AVFILTER_FLAG_HWDEVICE
  before initializing the filter with avfilter_init_str() or
  avfilter_init_dict().
*/
func (s *AVFilterContext) SetHwDeviceCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_device_ctx = value.ptr
	} else {
		s.ptr.hw_device_ctx = nil
	}
}

// NbThreads gets the nb_threads field.
/*
  Max number of threads allowed in this filter instance.
  If <= 0, its value is ignored.
  Overrides global number of threads set per filter graph.
*/
func (s *AVFilterContext) NbThreads() int {
	value := s.ptr.nb_threads
	return int(value)
}

// SetNbThreads sets the nb_threads field.
/*
  Max number of threads allowed in this filter instance.
  If <= 0, its value is ignored.
  Overrides global number of threads set per filter graph.
*/
func (s *AVFilterContext) SetNbThreads(value int) {
	s.ptr.nb_threads = (C.int)(value)
}

// Ready gets the ready field.
/*
  Ready status of the filter.
  A non-0 value means that the filter needs activating;
  a higher value suggests a more urgent activation.
*/
func (s *AVFilterContext) Ready() uint {
	value := s.ptr.ready
	return uint(value)
}

// SetReady sets the ready field.
/*
  Ready status of the filter.
  A non-0 value means that the filter needs activating;
  a higher value suggests a more urgent activation.
*/
func (s *AVFilterContext) SetReady(value uint) {
	s.ptr.ready = (C.uint)(value)
}

// ExtraHwFrames gets the extra_hw_frames field.
/*
  Sets the number of extra hardware frames which the filter will
  allocate on its output links for use in following filters or by
  the caller.

  Some hardware filters require all frames that they will use for
  output to be defined in advance before filtering starts.  For such
  filters, any hardware frame pools used for output must therefore be
  of fixed size.  The extra frames set here are on top of any number
  that the filter needs internally in order to operate normally.

  This field must be set before the graph containing this filter is
  configured.
*/
func (s *AVFilterContext) ExtraHwFrames() int {
	value := s.ptr.extra_hw_frames
	return int(value)
}

// SetExtraHwFrames sets the extra_hw_frames field.
/*
  Sets the number of extra hardware frames which the filter will
  allocate on its output links for use in following filters or by
  the caller.

  Some hardware filters require all frames that they will use for
  output to be defined in advance before filtering starts.  For such
  filters, any hardware frame pools used for output must therefore be
  of fixed size.  The extra frames set here are on top of any number
  that the filter needs internally in order to operate normally.

  This field must be set before the graph containing this filter is
  configured.
*/
func (s *AVFilterContext) SetExtraHwFrames(value int) {
	s.ptr.extra_hw_frames = (C.int)(value)
}

// --- Struct AVFilterLink ---

// AVFilterLink wraps AVFilterLink.
/*
  A link between two filters. This contains pointers to the source and
  destination filters between which this link exists, and the indexes of
  the pads involved. In addition, this link also contains the parameters
  which have been negotiated and agreed upon between the filter, such as
  image dimensions, format, etc.

  Applications must not normally access the link structure directly.
  Use the buffersrc and buffersink API instead.
  In the future, access to the header may be reserved for filters
  implementation.
*/
type AVFilterLink struct {
	ptr *C.AVFilterLink
}

func (s *AVFilterLink) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterLinkArray(ptr unsafe.Pointer) *Array[*AVFilterLink] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterLink]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterLink {
			ptr := (**C.AVFilterLink)(pointer)
			value := *ptr
			var valueMapped *AVFilterLink
			if value != nil {
				valueMapped = &AVFilterLink{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterLink) {
			ptr := (**C.AVFilterLink)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Src gets the src field.
//
//	source filter
func (s *AVFilterLink) Src() *AVFilterContext {
	value := s.ptr.src
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

// SetSrc sets the src field.
//
//	source filter
func (s *AVFilterLink) SetSrc(value *AVFilterContext) {
	if value != nil {
		s.ptr.src = value.ptr
	} else {
		s.ptr.src = nil
	}
}

// Srcpad gets the srcpad field.
//
//	output pad on the source filter
func (s *AVFilterLink) Srcpad() *AVFilterPad {
	value := s.ptr.srcpad
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetSrcpad sets the srcpad field.
//
//	output pad on the source filter
func (s *AVFilterLink) SetSrcpad(value *AVFilterPad) {
	if value != nil {
		s.ptr.srcpad = value.ptr
	} else {
		s.ptr.srcpad = nil
	}
}

// Dst gets the dst field.
//
//	dest filter
func (s *AVFilterLink) Dst() *AVFilterContext {
	value := s.ptr.dst
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

// SetDst sets the dst field.
//
//	dest filter
func (s *AVFilterLink) SetDst(value *AVFilterContext) {
	if value != nil {
		s.ptr.dst = value.ptr
	} else {
		s.ptr.dst = nil
	}
}

// Dstpad gets the dstpad field.
//
//	input pad on the dest filter
func (s *AVFilterLink) Dstpad() *AVFilterPad {
	value := s.ptr.dstpad
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetDstpad sets the dstpad field.
//
//	input pad on the dest filter
func (s *AVFilterLink) SetDstpad(value *AVFilterPad) {
	if value != nil {
		s.ptr.dstpad = value.ptr
	} else {
		s.ptr.dstpad = nil
	}
}

// Type gets the type field.
//
//	filter media type
func (s *AVFilterLink) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}

// SetType sets the type field.
//
//	filter media type
func (s *AVFilterLink) SetType(value AVMediaType) {
	s.ptr._type = (C.enum_AVMediaType)(value)
}

// W gets the w field.
//
//	agreed upon image width
func (s *AVFilterLink) W() int {
	value := s.ptr.w
	return int(value)
}

// SetW sets the w field.
//
//	agreed upon image width
func (s *AVFilterLink) SetW(value int) {
	s.ptr.w = (C.int)(value)
}

// H gets the h field.
//
//	agreed upon image height
func (s *AVFilterLink) H() int {
	value := s.ptr.h
	return int(value)
}

// SetH sets the h field.
//
//	agreed upon image height
func (s *AVFilterLink) SetH(value int) {
	s.ptr.h = (C.int)(value)
}

// SampleAspectRatio gets the sample_aspect_ratio field.
//
//	agreed upon sample aspect ratio
func (s *AVFilterLink) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
//
//	agreed upon sample aspect ratio
func (s *AVFilterLink) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

// ChannelLayout gets the channel_layout field.
/*
  channel layout of current buffer (see libavutil/channel_layout.h)
  @deprecated use ch_layout
*/
func (s *AVFilterLink) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

// SetChannelLayout sets the channel_layout field.
/*
  channel layout of current buffer (see libavutil/channel_layout.h)
  @deprecated use ch_layout
*/
func (s *AVFilterLink) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

// SampleRate gets the sample_rate field.
//
//	samples per second
func (s *AVFilterLink) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

// SetSampleRate sets the sample_rate field.
//
//	samples per second
func (s *AVFilterLink) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

// Format gets the format field.
//
//	agreed upon media format
func (s *AVFilterLink) Format() int {
	value := s.ptr.format
	return int(value)
}

// SetFormat sets the format field.
//
//	agreed upon media format
func (s *AVFilterLink) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

// TimeBase gets the time_base field.
/*
  Define the time base used by the PTS of the frames/samples
  which will pass through this link.
  During the configuration stage, each filter is supposed to
  change only the output timebase, while the timebase of the
  input link is assumed to be an unchangeable property.
*/
func (s *AVFilterLink) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
/*
  Define the time base used by the PTS of the frames/samples
  which will pass through this link.
  During the configuration stage, each filter is supposed to
  change only the output timebase, while the timebase of the
  input link is assumed to be an unchangeable property.
*/
func (s *AVFilterLink) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// ChLayout gets the ch_layout field.
//
//	channel layout of current buffer (see libavutil/channel_layout.h)
func (s *AVFilterLink) ChLayout() *AVChannelLayout {
	value := &s.ptr.ch_layout
	return &AVChannelLayout{ptr: value}
}

// Incfg gets the incfg field.
//
//	Lists of supported formats / etc. supported by the input filter.
func (s *AVFilterLink) Incfg() *AVFilterFormatsConfig {
	value := &s.ptr.incfg
	return &AVFilterFormatsConfig{ptr: value}
}

// Outcfg gets the outcfg field.
//
//	Lists of supported formats / etc. supported by the output filter.
func (s *AVFilterLink) Outcfg() *AVFilterFormatsConfig {
	value := &s.ptr.outcfg
	return &AVFilterFormatsConfig{ptr: value}
}

// InitState gets the init_state field.
func (s *AVFilterLink) InitState() uint32 {
	value := s.ptr.init_state
	return uint32(value)
}

// SetInitState sets the init_state field.
func (s *AVFilterLink) SetInitState(value uint32) {
	s.ptr.init_state = value
}

// Graph gets the graph field.
//
//	Graph the filter belongs to.
func (s *AVFilterLink) Graph() *AVFilterGraph {
	value := s.ptr.graph
	var valueMapped *AVFilterGraph
	if value != nil {
		valueMapped = &AVFilterGraph{ptr: value}
	}
	return valueMapped
}

// SetGraph sets the graph field.
//
//	Graph the filter belongs to.
func (s *AVFilterLink) SetGraph(value *AVFilterGraph) {
	if value != nil {
		s.ptr.graph = value.ptr
	} else {
		s.ptr.graph = nil
	}
}

// CurrentPts gets the current_pts field.
/*
  Current timestamp of the link, as defined by the most recent
  frame(s), in link time_base units.
*/
func (s *AVFilterLink) CurrentPts() int64 {
	value := s.ptr.current_pts
	return int64(value)
}

// SetCurrentPts sets the current_pts field.
/*
  Current timestamp of the link, as defined by the most recent
  frame(s), in link time_base units.
*/
func (s *AVFilterLink) SetCurrentPts(value int64) {
	s.ptr.current_pts = (C.int64_t)(value)
}

// CurrentPtsUs gets the current_pts_us field.
/*
  Current timestamp of the link, as defined by the most recent
  frame(s), in AV_TIME_BASE units.
*/
func (s *AVFilterLink) CurrentPtsUs() int64 {
	value := s.ptr.current_pts_us
	return int64(value)
}

// SetCurrentPtsUs sets the current_pts_us field.
/*
  Current timestamp of the link, as defined by the most recent
  frame(s), in AV_TIME_BASE units.
*/
func (s *AVFilterLink) SetCurrentPtsUs(value int64) {
	s.ptr.current_pts_us = (C.int64_t)(value)
}

// AgeIndex gets the age_index field.
//
//	Index in the age array.
func (s *AVFilterLink) AgeIndex() int {
	value := s.ptr.age_index
	return int(value)
}

// SetAgeIndex sets the age_index field.
//
//	Index in the age array.
func (s *AVFilterLink) SetAgeIndex(value int) {
	s.ptr.age_index = (C.int)(value)
}

// FrameRate gets the frame_rate field.
/*
  Frame rate of the stream on the link, or 1/0 if unknown or variable;
  if left to 0/0, will be automatically copied from the first input
  of the source filter if it exists.

  Sources should set it to the best estimation of the real frame rate.
  If the source frame rate is unknown or variable, set this to 1/0.
  Filters should update it if necessary depending on their function.
  Sinks can use it to set a default output frame rate.
  It is similar to the r_frame_rate field in AVStream.
*/
func (s *AVFilterLink) FrameRate() *AVRational {
	value := s.ptr.frame_rate
	return &AVRational{value: value}
}

// SetFrameRate sets the frame_rate field.
/*
  Frame rate of the stream on the link, or 1/0 if unknown or variable;
  if left to 0/0, will be automatically copied from the first input
  of the source filter if it exists.

  Sources should set it to the best estimation of the real frame rate.
  If the source frame rate is unknown or variable, set this to 1/0.
  Filters should update it if necessary depending on their function.
  Sinks can use it to set a default output frame rate.
  It is similar to the r_frame_rate field in AVStream.
*/
func (s *AVFilterLink) SetFrameRate(value *AVRational) {
	s.ptr.frame_rate = value.value
}

// MinSamples gets the min_samples field.
/*
  Minimum number of samples to filter at once. If filter_frame() is
  called with fewer samples, it will accumulate them in fifo.
  This field and the related ones must not be changed after filtering
  has started.
  If 0, all related fields are ignored.
*/
func (s *AVFilterLink) MinSamples() int {
	value := s.ptr.min_samples
	return int(value)
}

// SetMinSamples sets the min_samples field.
/*
  Minimum number of samples to filter at once. If filter_frame() is
  called with fewer samples, it will accumulate them in fifo.
  This field and the related ones must not be changed after filtering
  has started.
  If 0, all related fields are ignored.
*/
func (s *AVFilterLink) SetMinSamples(value int) {
	s.ptr.min_samples = (C.int)(value)
}

// MaxSamples gets the max_samples field.
/*
  Maximum number of samples to filter at once. If filter_frame() is
  called with more samples, it will split them.
*/
func (s *AVFilterLink) MaxSamples() int {
	value := s.ptr.max_samples
	return int(value)
}

// SetMaxSamples sets the max_samples field.
/*
  Maximum number of samples to filter at once. If filter_frame() is
  called with more samples, it will split them.
*/
func (s *AVFilterLink) SetMaxSamples(value int) {
	s.ptr.max_samples = (C.int)(value)
}

// FrameCountIn gets the frame_count_in field.
//
//	Number of past frames sent through the link.
func (s *AVFilterLink) FrameCountIn() int64 {
	value := s.ptr.frame_count_in
	return int64(value)
}

// SetFrameCountIn sets the frame_count_in field.
//
//	Number of past frames sent through the link.
func (s *AVFilterLink) SetFrameCountIn(value int64) {
	s.ptr.frame_count_in = (C.int64_t)(value)
}

// FrameCountOut gets the frame_count_out field.
//
//	Number of past frames sent through the link.
func (s *AVFilterLink) FrameCountOut() int64 {
	value := s.ptr.frame_count_out
	return int64(value)
}

// SetFrameCountOut sets the frame_count_out field.
//
//	Number of past frames sent through the link.
func (s *AVFilterLink) SetFrameCountOut(value int64) {
	s.ptr.frame_count_out = (C.int64_t)(value)
}

// SampleCountIn gets the sample_count_in field.
//
//	Number of past samples sent through the link.
func (s *AVFilterLink) SampleCountIn() int64 {
	value := s.ptr.sample_count_in
	return int64(value)
}

// SetSampleCountIn sets the sample_count_in field.
//
//	Number of past samples sent through the link.
func (s *AVFilterLink) SetSampleCountIn(value int64) {
	s.ptr.sample_count_in = (C.int64_t)(value)
}

// SampleCountOut gets the sample_count_out field.
//
//	Number of past samples sent through the link.
func (s *AVFilterLink) SampleCountOut() int64 {
	value := s.ptr.sample_count_out
	return int64(value)
}

// SetSampleCountOut sets the sample_count_out field.
//
//	Number of past samples sent through the link.
func (s *AVFilterLink) SetSampleCountOut(value int64) {
	s.ptr.sample_count_out = (C.int64_t)(value)
}

// FramePool gets the frame_pool field.
//
//	A pointer to a FFFramePool struct.
func (s *AVFilterLink) FramePool() unsafe.Pointer {
	value := s.ptr.frame_pool
	return value
}

// SetFramePool sets the frame_pool field.
//
//	A pointer to a FFFramePool struct.
func (s *AVFilterLink) SetFramePool(value unsafe.Pointer) {
	s.ptr.frame_pool = value
}

// FrameWantedOut gets the frame_wanted_out field.
/*
  True if a frame is currently wanted on the output of this filter.
  Set when ff_request_frame() is called by the output,
  cleared when a frame is filtered.
*/
func (s *AVFilterLink) FrameWantedOut() int {
	value := s.ptr.frame_wanted_out
	return int(value)
}

// SetFrameWantedOut sets the frame_wanted_out field.
/*
  True if a frame is currently wanted on the output of this filter.
  Set when ff_request_frame() is called by the output,
  cleared when a frame is filtered.
*/
func (s *AVFilterLink) SetFrameWantedOut(value int) {
	s.ptr.frame_wanted_out = (C.int)(value)
}

// HwFramesCtx gets the hw_frames_ctx field.
/*
  For hwaccel pixel formats, this should be a reference to the
  AVHWFramesContext describing the frames.
*/
func (s *AVFilterLink) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwFramesCtx sets the hw_frames_ctx field.
/*
  For hwaccel pixel formats, this should be a reference to the
  AVHWFramesContext describing the frames.
*/
func (s *AVFilterLink) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

// Reserved gets the reserved field.
/*
  Internal structure members.
  The fields below this limit are internal for libavfilter's use
  and must in no way be accessed by applications.
*/
func (s *AVFilterLink) Reserved() *Array[uint8] {
	value := &s.ptr.reserved
	return ToUint8Array(unsafe.Pointer(value))
}

// --- Struct AVFilterPad ---

// AVFilterPad wraps AVFilterPad.
type AVFilterPad struct {
	ptr *C.AVFilterPad
}

func (s *AVFilterPad) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterPadArray(ptr unsafe.Pointer) *Array[*AVFilterPad] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterPad]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterPad {
			ptr := (**C.AVFilterPad)(pointer)
			value := *ptr
			var valueMapped *AVFilterPad
			if value != nil {
				valueMapped = &AVFilterPad{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterPad) {
			ptr := (**C.AVFilterPad)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVFilterFormats ---

// AVFilterFormats wraps AVFilterFormats.
type AVFilterFormats struct {
	ptr *C.AVFilterFormats
}

func (s *AVFilterFormats) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterFormatsArray(ptr unsafe.Pointer) *Array[*AVFilterFormats] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterFormats]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterFormats {
			ptr := (**C.AVFilterFormats)(pointer)
			value := *ptr
			var valueMapped *AVFilterFormats
			if value != nil {
				valueMapped = &AVFilterFormats{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterFormats) {
			ptr := (**C.AVFilterFormats)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVFilterChannelLayouts ---

// AVFilterChannelLayouts wraps AVFilterChannelLayouts.
type AVFilterChannelLayouts struct {
	ptr *C.AVFilterChannelLayouts
}

func (s *AVFilterChannelLayouts) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterChannelLayoutsArray(ptr unsafe.Pointer) *Array[*AVFilterChannelLayouts] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterChannelLayouts]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterChannelLayouts {
			ptr := (**C.AVFilterChannelLayouts)(pointer)
			value := *ptr
			var valueMapped *AVFilterChannelLayouts
			if value != nil {
				valueMapped = &AVFilterChannelLayouts{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterChannelLayouts) {
			ptr := (**C.AVFilterChannelLayouts)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVFilter ---

// AVFilter wraps AVFilter.
/*
  Filter definition. This defines the pads a filter contains, and all the
  callback functions used to interact with the filter.
*/
type AVFilter struct {
	ptr *C.AVFilter
}

func (s *AVFilter) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterArray(ptr unsafe.Pointer) *Array[*AVFilter] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilter]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilter {
			ptr := (**C.AVFilter)(pointer)
			value := *ptr
			var valueMapped *AVFilter
			if value != nil {
				valueMapped = &AVFilter{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilter) {
			ptr := (**C.AVFilter)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Name gets the name field.
//
//	Filter name. Must be non-NULL and unique among filters.
func (s *AVFilter) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
//
//	Filter name. Must be non-NULL and unique among filters.
func (s *AVFilter) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// Description gets the description field.
/*
  A description of the filter. May be NULL.

  You should use the NULL_IF_CONFIG_SMALL() macro to define it.
*/
func (s *AVFilter) Description() *CStr {
	value := s.ptr.description
	return wrapCStr(value)
}

// SetDescription sets the description field.
/*
  A description of the filter. May be NULL.

  You should use the NULL_IF_CONFIG_SMALL() macro to define it.
*/
func (s *AVFilter) SetDescription(value *CStr) {
	s.ptr.description = value.ptr
}

// Inputs gets the inputs field.
/*
  List of static inputs.

  NULL if there are no (static) inputs. Instances of filters with
  AVFILTER_FLAG_DYNAMIC_INPUTS set may have more inputs than present in
  this list.
*/
func (s *AVFilter) Inputs() *AVFilterPad {
	value := s.ptr.inputs
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetInputs sets the inputs field.
/*
  List of static inputs.

  NULL if there are no (static) inputs. Instances of filters with
  AVFILTER_FLAG_DYNAMIC_INPUTS set may have more inputs than present in
  this list.
*/
func (s *AVFilter) SetInputs(value *AVFilterPad) {
	if value != nil {
		s.ptr.inputs = value.ptr
	} else {
		s.ptr.inputs = nil
	}
}

// Outputs gets the outputs field.
/*
  List of static outputs.

  NULL if there are no (static) outputs. Instances of filters with
  AVFILTER_FLAG_DYNAMIC_OUTPUTS set may have more outputs than present in
  this list.
*/
func (s *AVFilter) Outputs() *AVFilterPad {
	value := s.ptr.outputs
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetOutputs sets the outputs field.
/*
  List of static outputs.

  NULL if there are no (static) outputs. Instances of filters with
  AVFILTER_FLAG_DYNAMIC_OUTPUTS set may have more outputs than present in
  this list.
*/
func (s *AVFilter) SetOutputs(value *AVFilterPad) {
	if value != nil {
		s.ptr.outputs = value.ptr
	} else {
		s.ptr.outputs = nil
	}
}

// PrivClass gets the priv_class field.
/*
  A class for the private data, used to declare filter private AVOptions.
  This field is NULL for filters that do not declare any options.

  If this field is non-NULL, the first member of the filter private data
  must be a pointer to AVClass, which will be set by libavfilter generic
  code to this class.
*/
func (s *AVFilter) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetPrivClass sets the priv_class field.
/*
  A class for the private data, used to declare filter private AVOptions.
  This field is NULL for filters that do not declare any options.

  If this field is non-NULL, the first member of the filter private data
  must be a pointer to AVClass, which will be set by libavfilter generic
  code to this class.
*/
func (s *AVFilter) SetPrivClass(value *AVClass) {
	if value != nil {
		s.ptr.priv_class = value.ptr
	} else {
		s.ptr.priv_class = nil
	}
}

// Flags gets the flags field.
//
//	A combination of AVFILTER_FLAG_*
func (s *AVFilter) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
//
//	A combination of AVFILTER_FLAG_*
func (s *AVFilter) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// NbInputs gets the nb_inputs field.
//
//	The number of entries in the list of inputs.
func (s *AVFilter) NbInputs() uint8 {
	value := s.ptr.nb_inputs
	return uint8(value)
}

// SetNbInputs sets the nb_inputs field.
//
//	The number of entries in the list of inputs.
func (s *AVFilter) SetNbInputs(value uint8) {
	s.ptr.nb_inputs = (C.uint8_t)(value)
}

// NbOutputs gets the nb_outputs field.
//
//	The number of entries in the list of outputs.
func (s *AVFilter) NbOutputs() uint8 {
	value := s.ptr.nb_outputs
	return uint8(value)
}

// SetNbOutputs sets the nb_outputs field.
//
//	The number of entries in the list of outputs.
func (s *AVFilter) SetNbOutputs(value uint8) {
	s.ptr.nb_outputs = (C.uint8_t)(value)
}

// FormatsState gets the formats_state field.
/*
  This field determines the state of the formats union.
  It is an enum FilterFormatsState value.
*/
func (s *AVFilter) FormatsState() uint8 {
	value := s.ptr.formats_state
	return uint8(value)
}

// SetFormatsState sets the formats_state field.
/*
  This field determines the state of the formats union.
  It is an enum FilterFormatsState value.
*/
func (s *AVFilter) SetFormatsState(value uint8) {
	s.ptr.formats_state = (C.uint8_t)(value)
}

// preinit skipped due to func ptr

// init skipped due to func ptr

// uninit skipped due to func ptr

// formats skipped due to union type

// PrivSize gets the priv_size field.
//
//	size of private data to allocate for the filter
func (s *AVFilter) PrivSize() int {
	value := s.ptr.priv_size
	return int(value)
}

// SetPrivSize sets the priv_size field.
//
//	size of private data to allocate for the filter
func (s *AVFilter) SetPrivSize(value int) {
	s.ptr.priv_size = (C.int)(value)
}

// FlagsInternal gets the flags_internal field.
//
//	Additional flags for avfilter internal use only.
func (s *AVFilter) FlagsInternal() int {
	value := s.ptr.flags_internal
	return int(value)
}

// SetFlagsInternal sets the flags_internal field.
//
//	Additional flags for avfilter internal use only.
func (s *AVFilter) SetFlagsInternal(value int) {
	s.ptr.flags_internal = (C.int)(value)
}

// process_command skipped due to func ptr

// activate skipped due to func ptr

// --- Struct AVFilterInternal ---

// AVFilterInternal wraps AVFilterInternal.
type AVFilterInternal struct {
	ptr *C.AVFilterInternal
}

func (s *AVFilterInternal) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterInternalArray(ptr unsafe.Pointer) *Array[*AVFilterInternal] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterInternal]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterInternal {
			ptr := (**C.AVFilterInternal)(pointer)
			value := *ptr
			var valueMapped *AVFilterInternal
			if value != nil {
				valueMapped = &AVFilterInternal{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterInternal) {
			ptr := (**C.AVFilterInternal)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVFilterFormatsConfig ---

// AVFilterFormatsConfig wraps AVFilterFormatsConfig.
/*
  Lists of formats / etc. supported by an end of a link.

  This structure is directly part of AVFilterLink, in two copies:
  one for the source filter, one for the destination filter.

  These lists are used for negotiating the format to actually be used,
  which will be loaded into the format and channel_layout members of
  AVFilterLink, when chosen.
*/
type AVFilterFormatsConfig struct {
	ptr *C.AVFilterFormatsConfig
}

func (s *AVFilterFormatsConfig) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterFormatsConfigArray(ptr unsafe.Pointer) *Array[*AVFilterFormatsConfig] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterFormatsConfig]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterFormatsConfig {
			ptr := (**C.AVFilterFormatsConfig)(pointer)
			value := *ptr
			var valueMapped *AVFilterFormatsConfig
			if value != nil {
				valueMapped = &AVFilterFormatsConfig{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterFormatsConfig) {
			ptr := (**C.AVFilterFormatsConfig)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Formats gets the formats field.
//
//	List of supported formats (pixel or sample).
func (s *AVFilterFormatsConfig) Formats() *AVFilterFormats {
	value := s.ptr.formats
	var valueMapped *AVFilterFormats
	if value != nil {
		valueMapped = &AVFilterFormats{ptr: value}
	}
	return valueMapped
}

// SetFormats sets the formats field.
//
//	List of supported formats (pixel or sample).
func (s *AVFilterFormatsConfig) SetFormats(value *AVFilterFormats) {
	if value != nil {
		s.ptr.formats = value.ptr
	} else {
		s.ptr.formats = nil
	}
}

// Samplerates gets the samplerates field.
//
//	Lists of supported sample rates, only for audio.
func (s *AVFilterFormatsConfig) Samplerates() *AVFilterFormats {
	value := s.ptr.samplerates
	var valueMapped *AVFilterFormats
	if value != nil {
		valueMapped = &AVFilterFormats{ptr: value}
	}
	return valueMapped
}

// SetSamplerates sets the samplerates field.
//
//	Lists of supported sample rates, only for audio.
func (s *AVFilterFormatsConfig) SetSamplerates(value *AVFilterFormats) {
	if value != nil {
		s.ptr.samplerates = value.ptr
	} else {
		s.ptr.samplerates = nil
	}
}

// ChannelLayouts gets the channel_layouts field.
//
//	Lists of supported channel layouts, only for audio.
func (s *AVFilterFormatsConfig) ChannelLayouts() *AVFilterChannelLayouts {
	value := s.ptr.channel_layouts
	var valueMapped *AVFilterChannelLayouts
	if value != nil {
		valueMapped = &AVFilterChannelLayouts{ptr: value}
	}
	return valueMapped
}

// SetChannelLayouts sets the channel_layouts field.
//
//	Lists of supported channel layouts, only for audio.
func (s *AVFilterFormatsConfig) SetChannelLayouts(value *AVFilterChannelLayouts) {
	if value != nil {
		s.ptr.channel_layouts = value.ptr
	} else {
		s.ptr.channel_layouts = nil
	}
}

// --- Struct AVFilterGraphInternal ---

// AVFilterGraphInternal wraps AVFilterGraphInternal.
type AVFilterGraphInternal struct {
	ptr *C.AVFilterGraphInternal
}

func (s *AVFilterGraphInternal) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterGraphInternalArray(ptr unsafe.Pointer) *Array[*AVFilterGraphInternal] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterGraphInternal]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterGraphInternal {
			ptr := (**C.AVFilterGraphInternal)(pointer)
			value := *ptr
			var valueMapped *AVFilterGraphInternal
			if value != nil {
				valueMapped = &AVFilterGraphInternal{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterGraphInternal) {
			ptr := (**C.AVFilterGraphInternal)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVFilterGraph ---

// AVFilterGraph wraps AVFilterGraph.
type AVFilterGraph struct {
	ptr *C.AVFilterGraph
}

func (s *AVFilterGraph) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterGraphArray(ptr unsafe.Pointer) *Array[*AVFilterGraph] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterGraph]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterGraph {
			ptr := (**C.AVFilterGraph)(pointer)
			value := *ptr
			var valueMapped *AVFilterGraph
			if value != nil {
				valueMapped = &AVFilterGraph{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterGraph) {
			ptr := (**C.AVFilterGraph)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// AvClass gets the av_class field.
func (s *AVFilterGraph) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
func (s *AVFilterGraph) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Filters gets the filters field.
func (s *AVFilterGraph) Filters() *Array[*AVFilterContext] {
	value := s.ptr.filters
	return ToAVFilterContextArray(unsafe.Pointer(value))
}

// SetFilters sets the filters field.
func (s *AVFilterGraph) SetFilters(value *Array[AVFilterContext]) {
	if value != nil {
		s.ptr.filters = (**C.AVFilterContext)(value.ptr)
	} else {
		s.ptr.filters = nil
	}
}

// NbFilters gets the nb_filters field.
func (s *AVFilterGraph) NbFilters() uint {
	value := s.ptr.nb_filters
	return uint(value)
}

// SetNbFilters sets the nb_filters field.
func (s *AVFilterGraph) SetNbFilters(value uint) {
	s.ptr.nb_filters = (C.uint)(value)
}

// ScaleSwsOpts gets the scale_sws_opts field.
//
//	sws options to use for the auto-inserted scale filters
func (s *AVFilterGraph) ScaleSwsOpts() *CStr {
	value := s.ptr.scale_sws_opts
	return wrapCStr(value)
}

// SetScaleSwsOpts sets the scale_sws_opts field.
//
//	sws options to use for the auto-inserted scale filters
func (s *AVFilterGraph) SetScaleSwsOpts(value *CStr) {
	s.ptr.scale_sws_opts = value.ptr
}

// ThreadType gets the thread_type field.
/*
  Type of multithreading allowed for filters in this graph. A combination
  of AVFILTER_THREAD_* flags.

  May be set by the caller at any point, the setting will apply to all
  filters initialized after that. The default is allowing everything.

  When a filter in this graph is initialized, this field is combined using
  bit AND with AVFilterContext.thread_type to get the final mask used for
  determining allowed threading types. I.e. a threading type needs to be
  set in both to be allowed.
*/
func (s *AVFilterGraph) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}

// SetThreadType sets the thread_type field.
/*
  Type of multithreading allowed for filters in this graph. A combination
  of AVFILTER_THREAD_* flags.

  May be set by the caller at any point, the setting will apply to all
  filters initialized after that. The default is allowing everything.

  When a filter in this graph is initialized, this field is combined using
  bit AND with AVFilterContext.thread_type to get the final mask used for
  determining allowed threading types. I.e. a threading type needs to be
  set in both to be allowed.
*/
func (s *AVFilterGraph) SetThreadType(value int) {
	s.ptr.thread_type = (C.int)(value)
}

// NbThreads gets the nb_threads field.
/*
  Maximum number of threads used by filters in this graph. May be set by
  the caller before adding any filters to the filtergraph. Zero (the
  default) means that the number of threads is determined automatically.
*/
func (s *AVFilterGraph) NbThreads() int {
	value := s.ptr.nb_threads
	return int(value)
}

// SetNbThreads sets the nb_threads field.
/*
  Maximum number of threads used by filters in this graph. May be set by
  the caller before adding any filters to the filtergraph. Zero (the
  default) means that the number of threads is determined automatically.
*/
func (s *AVFilterGraph) SetNbThreads(value int) {
	s.ptr.nb_threads = (C.int)(value)
}

// Internal gets the internal field.
//
//	Opaque object for libavfilter internal use.
func (s *AVFilterGraph) Internal() *AVFilterGraphInternal {
	value := s.ptr.internal
	var valueMapped *AVFilterGraphInternal
	if value != nil {
		valueMapped = &AVFilterGraphInternal{ptr: value}
	}
	return valueMapped
}

// SetInternal sets the internal field.
//
//	Opaque object for libavfilter internal use.
func (s *AVFilterGraph) SetInternal(value *AVFilterGraphInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

// Opaque gets the opaque field.
/*
  Opaque user data. May be set by the caller to an arbitrary value, e.g. to
  be used from callbacks like @ref AVFilterGraph.execute.
  Libavfilter will not touch this field in any way.
*/
func (s *AVFilterGraph) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
/*
  Opaque user data. May be set by the caller to an arbitrary value, e.g. to
  be used from callbacks like @ref AVFilterGraph.execute.
  Libavfilter will not touch this field in any way.
*/
func (s *AVFilterGraph) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// execute skipped due to callback ptr

// AresampleSwrOpts gets the aresample_swr_opts field.
//
//	swr options to use for the auto-inserted aresample filters, Access ONLY through AVOptions
func (s *AVFilterGraph) AresampleSwrOpts() *CStr {
	value := s.ptr.aresample_swr_opts
	return wrapCStr(value)
}

// SetAresampleSwrOpts sets the aresample_swr_opts field.
//
//	swr options to use for the auto-inserted aresample filters, Access ONLY through AVOptions
func (s *AVFilterGraph) SetAresampleSwrOpts(value *CStr) {
	s.ptr.aresample_swr_opts = value.ptr
}

// SinkLinks gets the sink_links field.
/*
  Private fields

  The following fields are for internal use only.
  Their type, offset, number and semantic can change without notice.
*/
func (s *AVFilterGraph) SinkLinks() *Array[*AVFilterLink] {
	value := s.ptr.sink_links
	return ToAVFilterLinkArray(unsafe.Pointer(value))
}

// SetSinkLinks sets the sink_links field.
/*
  Private fields

  The following fields are for internal use only.
  Their type, offset, number and semantic can change without notice.
*/
func (s *AVFilterGraph) SetSinkLinks(value *Array[AVFilterLink]) {
	if value != nil {
		s.ptr.sink_links = (**C.AVFilterLink)(value.ptr)
	} else {
		s.ptr.sink_links = nil
	}
}

// SinkLinksCount gets the sink_links_count field.
func (s *AVFilterGraph) SinkLinksCount() int {
	value := s.ptr.sink_links_count
	return int(value)
}

// SetSinkLinksCount sets the sink_links_count field.
func (s *AVFilterGraph) SetSinkLinksCount(value int) {
	s.ptr.sink_links_count = (C.int)(value)
}

// DisableAutoConvert gets the disable_auto_convert field.
func (s *AVFilterGraph) DisableAutoConvert() uint {
	value := s.ptr.disable_auto_convert
	return uint(value)
}

// SetDisableAutoConvert sets the disable_auto_convert field.
func (s *AVFilterGraph) SetDisableAutoConvert(value uint) {
	s.ptr.disable_auto_convert = (C.uint)(value)
}

// --- Struct AVFilterInOut ---

// AVFilterInOut wraps AVFilterInOut.
/*
  A linked-list of the inputs/outputs of the filter chain.

  This is mainly useful for avfilter_graph_parse() / avfilter_graph_parse2(),
  where it is used to communicate open (unlinked) inputs and outputs from and
  to the caller.
  This struct specifies, per each not connected pad contained in the graph, the
  filter context and the pad index required for establishing a link.
*/
type AVFilterInOut struct {
	ptr *C.AVFilterInOut
}

func (s *AVFilterInOut) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterInOutArray(ptr unsafe.Pointer) *Array[*AVFilterInOut] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterInOut]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterInOut {
			ptr := (**C.AVFilterInOut)(pointer)
			value := *ptr
			var valueMapped *AVFilterInOut
			if value != nil {
				valueMapped = &AVFilterInOut{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterInOut) {
			ptr := (**C.AVFilterInOut)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Name gets the name field.
//
//	unique name for this input/output in the list
func (s *AVFilterInOut) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
//
//	unique name for this input/output in the list
func (s *AVFilterInOut) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// FilterCtx gets the filter_ctx field.
//
//	filter context associated to this input/output
func (s *AVFilterInOut) FilterCtx() *AVFilterContext {
	value := s.ptr.filter_ctx
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

// SetFilterCtx sets the filter_ctx field.
//
//	filter context associated to this input/output
func (s *AVFilterInOut) SetFilterCtx(value *AVFilterContext) {
	if value != nil {
		s.ptr.filter_ctx = value.ptr
	} else {
		s.ptr.filter_ctx = nil
	}
}

// PadIdx gets the pad_idx field.
//
//	index of the filt_ctx pad to use for linking
func (s *AVFilterInOut) PadIdx() int {
	value := s.ptr.pad_idx
	return int(value)
}

// SetPadIdx sets the pad_idx field.
//
//	index of the filt_ctx pad to use for linking
func (s *AVFilterInOut) SetPadIdx(value int) {
	s.ptr.pad_idx = (C.int)(value)
}

// Next gets the next field.
//
//	next input/input in the list, NULL if this is the last
func (s *AVFilterInOut) Next() *AVFilterInOut {
	value := s.ptr.next
	var valueMapped *AVFilterInOut
	if value != nil {
		valueMapped = &AVFilterInOut{ptr: value}
	}
	return valueMapped
}

// SetNext sets the next field.
//
//	next input/input in the list, NULL if this is the last
func (s *AVFilterInOut) SetNext(value *AVFilterInOut) {
	if value != nil {
		s.ptr.next = value.ptr
	} else {
		s.ptr.next = nil
	}
}

// --- Struct AVFilterPadParams ---

// AVFilterPadParams wraps AVFilterPadParams.
/*
  Parameters of a filter's input or output pad.

  Created as a child of AVFilterParams by avfilter_graph_segment_parse().
  Freed in avfilter_graph_segment_free().
*/
type AVFilterPadParams struct {
	ptr *C.AVFilterPadParams
}

func (s *AVFilterPadParams) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterPadParamsArray(ptr unsafe.Pointer) *Array[*AVFilterPadParams] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterPadParams]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterPadParams {
			ptr := (**C.AVFilterPadParams)(pointer)
			value := *ptr
			var valueMapped *AVFilterPadParams
			if value != nil {
				valueMapped = &AVFilterPadParams{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterPadParams) {
			ptr := (**C.AVFilterPadParams)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Label gets the label field.
/*
  An av_malloc()'ed string containing the pad label.

  May be av_free()'d and set to NULL by the caller, in which case this pad
  will be treated as unlabeled for linking.
  May also be replaced by another av_malloc()'ed string.
*/
func (s *AVFilterPadParams) Label() *CStr {
	value := s.ptr.label
	return wrapCStr(value)
}

// SetLabel sets the label field.
/*
  An av_malloc()'ed string containing the pad label.

  May be av_free()'d and set to NULL by the caller, in which case this pad
  will be treated as unlabeled for linking.
  May also be replaced by another av_malloc()'ed string.
*/
func (s *AVFilterPadParams) SetLabel(value *CStr) {
	s.ptr.label = value.ptr
}

// --- Struct AVFilterParams ---

// AVFilterParams wraps AVFilterParams.
/*
  Parameters describing a filter to be created in a filtergraph.

  Created as a child of AVFilterGraphSegment by avfilter_graph_segment_parse().
  Freed in avfilter_graph_segment_free().
*/
type AVFilterParams struct {
	ptr *C.AVFilterParams
}

func (s *AVFilterParams) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterParamsArray(ptr unsafe.Pointer) *Array[*AVFilterParams] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterParams]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterParams {
			ptr := (**C.AVFilterParams)(pointer)
			value := *ptr
			var valueMapped *AVFilterParams
			if value != nil {
				valueMapped = &AVFilterParams{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterParams) {
			ptr := (**C.AVFilterParams)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Filter gets the filter field.
/*
  The filter context.

  Created by avfilter_graph_segment_create_filters() based on
  AVFilterParams.filter_name and instance_name.

  Callers may also create the filter context manually, then they should
  av_free() filter_name and set it to NULL. Such AVFilterParams instances
  are then skipped by avfilter_graph_segment_create_filters().
*/
func (s *AVFilterParams) Filter() *AVFilterContext {
	value := s.ptr.filter
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

// SetFilter sets the filter field.
/*
  The filter context.

  Created by avfilter_graph_segment_create_filters() based on
  AVFilterParams.filter_name and instance_name.

  Callers may also create the filter context manually, then they should
  av_free() filter_name and set it to NULL. Such AVFilterParams instances
  are then skipped by avfilter_graph_segment_create_filters().
*/
func (s *AVFilterParams) SetFilter(value *AVFilterContext) {
	if value != nil {
		s.ptr.filter = value.ptr
	} else {
		s.ptr.filter = nil
	}
}

// FilterName gets the filter_name field.
/*
  Name of the AVFilter to be used.

  An av_malloc()'ed string, set by avfilter_graph_segment_parse(). Will be
  passed to avfilter_get_by_name() by
  avfilter_graph_segment_create_filters().

  Callers may av_free() this string and replace it with another one or
  NULL. If the caller creates the filter instance manually, this string
  MUST be set to NULL.

  When both AVFilterParams.filter an AVFilterParams.filter_name are NULL,
  this AVFilterParams instance is skipped by avfilter_graph_segment_*()
  functions.
*/
func (s *AVFilterParams) FilterName() *CStr {
	value := s.ptr.filter_name
	return wrapCStr(value)
}

// SetFilterName sets the filter_name field.
/*
  Name of the AVFilter to be used.

  An av_malloc()'ed string, set by avfilter_graph_segment_parse(). Will be
  passed to avfilter_get_by_name() by
  avfilter_graph_segment_create_filters().

  Callers may av_free() this string and replace it with another one or
  NULL. If the caller creates the filter instance manually, this string
  MUST be set to NULL.

  When both AVFilterParams.filter an AVFilterParams.filter_name are NULL,
  this AVFilterParams instance is skipped by avfilter_graph_segment_*()
  functions.
*/
func (s *AVFilterParams) SetFilterName(value *CStr) {
	s.ptr.filter_name = value.ptr
}

// InstanceName gets the instance_name field.
/*
  Name to be used for this filter instance.

  An av_malloc()'ed string, may be set by avfilter_graph_segment_parse() or
  left NULL. The caller may av_free() this string and replace with another
  one or NULL.

  Will be used by avfilter_graph_segment_create_filters() - passed as the
  third argument to avfilter_graph_alloc_filter(), then freed and set to
  NULL.
*/
func (s *AVFilterParams) InstanceName() *CStr {
	value := s.ptr.instance_name
	return wrapCStr(value)
}

// SetInstanceName sets the instance_name field.
/*
  Name to be used for this filter instance.

  An av_malloc()'ed string, may be set by avfilter_graph_segment_parse() or
  left NULL. The caller may av_free() this string and replace with another
  one or NULL.

  Will be used by avfilter_graph_segment_create_filters() - passed as the
  third argument to avfilter_graph_alloc_filter(), then freed and set to
  NULL.
*/
func (s *AVFilterParams) SetInstanceName(value *CStr) {
	s.ptr.instance_name = value.ptr
}

// Opts gets the opts field.
/*
  Options to be apllied to the filter.

  Filled by avfilter_graph_segment_parse(). Afterwards may be freely
  modified by the caller.

  Will be applied to the filter by avfilter_graph_segment_apply_opts()
  with an equivalent of av_opt_set_dict2(filter, &opts, AV_OPT_SEARCH_CHILDREN),
  i.e. any unapplied options will be left in this dictionary.
*/
func (s *AVFilterParams) Opts() *AVDictionary {
	value := s.ptr.opts
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

// SetOpts sets the opts field.
/*
  Options to be apllied to the filter.

  Filled by avfilter_graph_segment_parse(). Afterwards may be freely
  modified by the caller.

  Will be applied to the filter by avfilter_graph_segment_apply_opts()
  with an equivalent of av_opt_set_dict2(filter, &opts, AV_OPT_SEARCH_CHILDREN),
  i.e. any unapplied options will be left in this dictionary.
*/
func (s *AVFilterParams) SetOpts(value *AVDictionary) {
	if value != nil {
		s.ptr.opts = value.ptr
	} else {
		s.ptr.opts = nil
	}
}

// Inputs gets the inputs field.
func (s *AVFilterParams) Inputs() *Array[*AVFilterPadParams] {
	value := s.ptr.inputs
	return ToAVFilterPadParamsArray(unsafe.Pointer(value))
}

// SetInputs sets the inputs field.
func (s *AVFilterParams) SetInputs(value *Array[AVFilterPadParams]) {
	if value != nil {
		s.ptr.inputs = (**C.AVFilterPadParams)(value.ptr)
	} else {
		s.ptr.inputs = nil
	}
}

// NbInputs gets the nb_inputs field.
func (s *AVFilterParams) NbInputs() uint {
	value := s.ptr.nb_inputs
	return uint(value)
}

// SetNbInputs sets the nb_inputs field.
func (s *AVFilterParams) SetNbInputs(value uint) {
	s.ptr.nb_inputs = (C.uint)(value)
}

// Outputs gets the outputs field.
func (s *AVFilterParams) Outputs() *Array[*AVFilterPadParams] {
	value := s.ptr.outputs
	return ToAVFilterPadParamsArray(unsafe.Pointer(value))
}

// SetOutputs sets the outputs field.
func (s *AVFilterParams) SetOutputs(value *Array[AVFilterPadParams]) {
	if value != nil {
		s.ptr.outputs = (**C.AVFilterPadParams)(value.ptr)
	} else {
		s.ptr.outputs = nil
	}
}

// NbOutputs gets the nb_outputs field.
func (s *AVFilterParams) NbOutputs() uint {
	value := s.ptr.nb_outputs
	return uint(value)
}

// SetNbOutputs sets the nb_outputs field.
func (s *AVFilterParams) SetNbOutputs(value uint) {
	s.ptr.nb_outputs = (C.uint)(value)
}

// --- Struct AVFilterChain ---

// AVFilterChain wraps AVFilterChain.
/*
  A filterchain is a list of filter specifications.

  Created as a child of AVFilterGraphSegment by avfilter_graph_segment_parse().
  Freed in avfilter_graph_segment_free().
*/
type AVFilterChain struct {
	ptr *C.AVFilterChain
}

func (s *AVFilterChain) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterChainArray(ptr unsafe.Pointer) *Array[*AVFilterChain] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterChain]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterChain {
			ptr := (**C.AVFilterChain)(pointer)
			value := *ptr
			var valueMapped *AVFilterChain
			if value != nil {
				valueMapped = &AVFilterChain{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterChain) {
			ptr := (**C.AVFilterChain)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Filters gets the filters field.
func (s *AVFilterChain) Filters() *Array[*AVFilterParams] {
	value := s.ptr.filters
	return ToAVFilterParamsArray(unsafe.Pointer(value))
}

// SetFilters sets the filters field.
func (s *AVFilterChain) SetFilters(value *Array[AVFilterParams]) {
	if value != nil {
		s.ptr.filters = (**C.AVFilterParams)(value.ptr)
	} else {
		s.ptr.filters = nil
	}
}

// NbFilters gets the nb_filters field.
func (s *AVFilterChain) NbFilters() uint64 {
	value := s.ptr.nb_filters
	return uint64(value)
}

// SetNbFilters sets the nb_filters field.
func (s *AVFilterChain) SetNbFilters(value uint64) {
	s.ptr.nb_filters = (C.size_t)(value)
}

// --- Struct AVFilterGraphSegment ---

// AVFilterGraphSegment wraps AVFilterGraphSegment.
/*
  A parsed representation of a filtergraph segment.

  A filtergraph segment is conceptually a list of filterchains, with some
  supplementary information (e.g. format conversion flags).

  Created by avfilter_graph_segment_parse(). Must be freed with
  avfilter_graph_segment_free().
*/
type AVFilterGraphSegment struct {
	ptr *C.AVFilterGraphSegment
}

func (s *AVFilterGraphSegment) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFilterGraphSegmentArray(ptr unsafe.Pointer) *Array[*AVFilterGraphSegment] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFilterGraphSegment]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFilterGraphSegment {
			ptr := (**C.AVFilterGraphSegment)(pointer)
			value := *ptr
			var valueMapped *AVFilterGraphSegment
			if value != nil {
				valueMapped = &AVFilterGraphSegment{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFilterGraphSegment) {
			ptr := (**C.AVFilterGraphSegment)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Graph gets the graph field.
/*
  The filtergraph this segment is associated with.
  Set by avfilter_graph_segment_parse().
*/
func (s *AVFilterGraphSegment) Graph() *AVFilterGraph {
	value := s.ptr.graph
	var valueMapped *AVFilterGraph
	if value != nil {
		valueMapped = &AVFilterGraph{ptr: value}
	}
	return valueMapped
}

// SetGraph sets the graph field.
/*
  The filtergraph this segment is associated with.
  Set by avfilter_graph_segment_parse().
*/
func (s *AVFilterGraphSegment) SetGraph(value *AVFilterGraph) {
	if value != nil {
		s.ptr.graph = value.ptr
	} else {
		s.ptr.graph = nil
	}
}

// Chains gets the chains field.
/*
  A list of filter chain contained in this segment.
  Set in avfilter_graph_segment_parse().
*/
func (s *AVFilterGraphSegment) Chains() *Array[*AVFilterChain] {
	value := s.ptr.chains
	return ToAVFilterChainArray(unsafe.Pointer(value))
}

// SetChains sets the chains field.
/*
  A list of filter chain contained in this segment.
  Set in avfilter_graph_segment_parse().
*/
func (s *AVFilterGraphSegment) SetChains(value *Array[AVFilterChain]) {
	if value != nil {
		s.ptr.chains = (**C.AVFilterChain)(value.ptr)
	} else {
		s.ptr.chains = nil
	}
}

// NbChains gets the nb_chains field.
func (s *AVFilterGraphSegment) NbChains() uint64 {
	value := s.ptr.nb_chains
	return uint64(value)
}

// SetNbChains sets the nb_chains field.
func (s *AVFilterGraphSegment) SetNbChains(value uint64) {
	s.ptr.nb_chains = (C.size_t)(value)
}

// ScaleSwsOpts gets the scale_sws_opts field.
/*
  A string containing a colon-separated list of key=value options applied
  to all scale filters in this segment.

  May be set by avfilter_graph_segment_parse().
  The caller may free this string with av_free() and replace it with a
  different av_malloc()'ed string.
*/
func (s *AVFilterGraphSegment) ScaleSwsOpts() *CStr {
	value := s.ptr.scale_sws_opts
	return wrapCStr(value)
}

// SetScaleSwsOpts sets the scale_sws_opts field.
/*
  A string containing a colon-separated list of key=value options applied
  to all scale filters in this segment.

  May be set by avfilter_graph_segment_parse().
  The caller may free this string with av_free() and replace it with a
  different av_malloc()'ed string.
*/
func (s *AVFilterGraphSegment) SetScaleSwsOpts(value *CStr) {
	s.ptr.scale_sws_opts = value.ptr
}

// --- Struct AVBufferSrcParameters ---

// AVBufferSrcParameters wraps AVBufferSrcParameters.
/*
  This structure contains the parameters describing the frames that will be
  passed to this filter.

  It should be allocated with av_buffersrc_parameters_alloc() and freed with
  av_free(). All the allocated fields in it remain owned by the caller.
*/
type AVBufferSrcParameters struct {
	ptr *C.AVBufferSrcParameters
}

func (s *AVBufferSrcParameters) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVBufferSrcParametersArray(ptr unsafe.Pointer) *Array[*AVBufferSrcParameters] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVBufferSrcParameters]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVBufferSrcParameters {
			ptr := (**C.AVBufferSrcParameters)(pointer)
			value := *ptr
			var valueMapped *AVBufferSrcParameters
			if value != nil {
				valueMapped = &AVBufferSrcParameters{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVBufferSrcParameters) {
			ptr := (**C.AVBufferSrcParameters)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Format gets the format field.
/*
  video: the pixel format, value corresponds to enum AVPixelFormat
  audio: the sample format, value corresponds to enum AVSampleFormat
*/
func (s *AVBufferSrcParameters) Format() int {
	value := s.ptr.format
	return int(value)
}

// SetFormat sets the format field.
/*
  video: the pixel format, value corresponds to enum AVPixelFormat
  audio: the sample format, value corresponds to enum AVSampleFormat
*/
func (s *AVBufferSrcParameters) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

// TimeBase gets the time_base field.
//
//	The timebase to be used for the timestamps on the input frames.
func (s *AVBufferSrcParameters) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
//
//	The timebase to be used for the timestamps on the input frames.
func (s *AVBufferSrcParameters) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// Width gets the width field.
//
//	Video only, the display dimensions of the input frames.
func (s *AVBufferSrcParameters) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
//
//	Video only, the display dimensions of the input frames.
func (s *AVBufferSrcParameters) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

// Height gets the height field.
//
//	Video only, the display dimensions of the input frames.
func (s *AVBufferSrcParameters) Height() int {
	value := s.ptr.height
	return int(value)
}

// SetHeight sets the height field.
//
//	Video only, the display dimensions of the input frames.
func (s *AVBufferSrcParameters) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

// SampleAspectRatio gets the sample_aspect_ratio field.
//
//	Video only, the sample (pixel) aspect ratio.
func (s *AVBufferSrcParameters) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
//
//	Video only, the sample (pixel) aspect ratio.
func (s *AVBufferSrcParameters) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

// FrameRate gets the frame_rate field.
/*
  Video only, the frame rate of the input video. This field must only be
  set to a non-zero value if input stream has a known constant framerate
  and should be left at its initial value if the framerate is variable or
  unknown.
*/
func (s *AVBufferSrcParameters) FrameRate() *AVRational {
	value := s.ptr.frame_rate
	return &AVRational{value: value}
}

// SetFrameRate sets the frame_rate field.
/*
  Video only, the frame rate of the input video. This field must only be
  set to a non-zero value if input stream has a known constant framerate
  and should be left at its initial value if the framerate is variable or
  unknown.
*/
func (s *AVBufferSrcParameters) SetFrameRate(value *AVRational) {
	s.ptr.frame_rate = value.value
}

// HwFramesCtx gets the hw_frames_ctx field.
/*
  Video with a hwaccel pixel format only. This should be a reference to an
  AVHWFramesContext instance describing the input frames.
*/
func (s *AVBufferSrcParameters) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwFramesCtx sets the hw_frames_ctx field.
/*
  Video with a hwaccel pixel format only. This should be a reference to an
  AVHWFramesContext instance describing the input frames.
*/
func (s *AVBufferSrcParameters) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

// SampleRate gets the sample_rate field.
//
//	Audio only, the audio sampling rate in samples per second.
func (s *AVBufferSrcParameters) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

// SetSampleRate sets the sample_rate field.
//
//	Audio only, the audio sampling rate in samples per second.
func (s *AVBufferSrcParameters) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

// ChannelLayout gets the channel_layout field.
/*
  Audio only, the audio channel layout
  @deprecated use ch_layout
*/
func (s *AVBufferSrcParameters) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

// SetChannelLayout sets the channel_layout field.
/*
  Audio only, the audio channel layout
  @deprecated use ch_layout
*/
func (s *AVBufferSrcParameters) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

// ChLayout gets the ch_layout field.
//
//	Audio only, the audio channel layout
func (s *AVBufferSrcParameters) ChLayout() *AVChannelLayout {
	value := &s.ptr.ch_layout
	return &AVChannelLayout{ptr: value}
}

// --- Struct AVDeviceInfoList ---

// AVDeviceInfoList wraps AVDeviceInfoList.
type AVDeviceInfoList struct {
	ptr *C.struct_AVDeviceInfoList
}

func (s *AVDeviceInfoList) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVDeviceInfoListArray(ptr unsafe.Pointer) *Array[*AVDeviceInfoList] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVDeviceInfoList]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVDeviceInfoList {
			ptr := (**C.struct_AVDeviceInfoList)(pointer)
			value := *ptr
			var valueMapped *AVDeviceInfoList
			if value != nil {
				valueMapped = &AVDeviceInfoList{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVDeviceInfoList) {
			ptr := (**C.struct_AVDeviceInfoList)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVCodecTag ---

// AVCodecTag wraps AVCodecTag.
//
//	input/output formats
type AVCodecTag struct {
	ptr *C.struct_AVCodecTag
}

func (s *AVCodecTag) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVCodecTagArray(ptr unsafe.Pointer) *Array[*AVCodecTag] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVCodecTag]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVCodecTag {
			ptr := (**C.struct_AVCodecTag)(pointer)
			value := *ptr
			var valueMapped *AVCodecTag
			if value != nil {
				valueMapped = &AVCodecTag{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVCodecTag) {
			ptr := (**C.struct_AVCodecTag)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVProbeData ---

// AVProbeData wraps AVProbeData.
//
//	This structure contains the data a format has to probe a file.
type AVProbeData struct {
	ptr *C.AVProbeData
}

func (s *AVProbeData) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVProbeDataArray(ptr unsafe.Pointer) *Array[*AVProbeData] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVProbeData]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVProbeData {
			ptr := (**C.AVProbeData)(pointer)
			value := *ptr
			var valueMapped *AVProbeData
			if value != nil {
				valueMapped = &AVProbeData{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVProbeData) {
			ptr := (**C.AVProbeData)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Filename gets the filename field.
func (s *AVProbeData) Filename() *CStr {
	value := s.ptr.filename
	return wrapCStr(value)
}

// SetFilename sets the filename field.
func (s *AVProbeData) SetFilename(value *CStr) {
	s.ptr.filename = value.ptr
}

// buf skipped due to prim ptr

// BufSize gets the buf_size field.
//
//	Size of buf except extra allocated bytes
func (s *AVProbeData) BufSize() int {
	value := s.ptr.buf_size
	return int(value)
}

// SetBufSize sets the buf_size field.
//
//	Size of buf except extra allocated bytes
func (s *AVProbeData) SetBufSize(value int) {
	s.ptr.buf_size = (C.int)(value)
}

// MimeType gets the mime_type field.
//
//	mime_type, when known.
func (s *AVProbeData) MimeType() *CStr {
	value := s.ptr.mime_type
	return wrapCStr(value)
}

// SetMimeType sets the mime_type field.
//
//	mime_type, when known.
func (s *AVProbeData) SetMimeType(value *CStr) {
	s.ptr.mime_type = value.ptr
}

// --- Struct AVOutputFormat ---

// AVOutputFormat wraps AVOutputFormat.
type AVOutputFormat struct {
	ptr *C.AVOutputFormat
}

func (s *AVOutputFormat) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVOutputFormatArray(ptr unsafe.Pointer) *Array[*AVOutputFormat] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVOutputFormat]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVOutputFormat {
			ptr := (**C.AVOutputFormat)(pointer)
			value := *ptr
			var valueMapped *AVOutputFormat
			if value != nil {
				valueMapped = &AVOutputFormat{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVOutputFormat) {
			ptr := (**C.AVOutputFormat)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Name gets the name field.
func (s *AVOutputFormat) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
func (s *AVOutputFormat) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// LongName gets the long_name field.
/*
  Descriptive name for the format, meant to be more human-readable
  than name. You should use the NULL_IF_CONFIG_SMALL() macro
  to define it.
*/
func (s *AVOutputFormat) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

// SetLongName sets the long_name field.
/*
  Descriptive name for the format, meant to be more human-readable
  than name. You should use the NULL_IF_CONFIG_SMALL() macro
  to define it.
*/
func (s *AVOutputFormat) SetLongName(value *CStr) {
	s.ptr.long_name = value.ptr
}

// MimeType gets the mime_type field.
func (s *AVOutputFormat) MimeType() *CStr {
	value := s.ptr.mime_type
	return wrapCStr(value)
}

// SetMimeType sets the mime_type field.
func (s *AVOutputFormat) SetMimeType(value *CStr) {
	s.ptr.mime_type = value.ptr
}

// Extensions gets the extensions field.
//
//	comma-separated filename extensions
func (s *AVOutputFormat) Extensions() *CStr {
	value := s.ptr.extensions
	return wrapCStr(value)
}

// SetExtensions sets the extensions field.
//
//	comma-separated filename extensions
func (s *AVOutputFormat) SetExtensions(value *CStr) {
	s.ptr.extensions = value.ptr
}

// AudioCodec gets the audio_codec field.
//
//	default audio codec
func (s *AVOutputFormat) AudioCodec() AVCodecID {
	value := s.ptr.audio_codec
	return AVCodecID(value)
}

// SetAudioCodec sets the audio_codec field.
//
//	default audio codec
func (s *AVOutputFormat) SetAudioCodec(value AVCodecID) {
	s.ptr.audio_codec = (C.enum_AVCodecID)(value)
}

// VideoCodec gets the video_codec field.
//
//	default video codec
func (s *AVOutputFormat) VideoCodec() AVCodecID {
	value := s.ptr.video_codec
	return AVCodecID(value)
}

// SetVideoCodec sets the video_codec field.
//
//	default video codec
func (s *AVOutputFormat) SetVideoCodec(value AVCodecID) {
	s.ptr.video_codec = (C.enum_AVCodecID)(value)
}

// SubtitleCodec gets the subtitle_codec field.
//
//	default subtitle codec
func (s *AVOutputFormat) SubtitleCodec() AVCodecID {
	value := s.ptr.subtitle_codec
	return AVCodecID(value)
}

// SetSubtitleCodec sets the subtitle_codec field.
//
//	default subtitle codec
func (s *AVOutputFormat) SetSubtitleCodec(value AVCodecID) {
	s.ptr.subtitle_codec = (C.enum_AVCodecID)(value)
}

// Flags gets the flags field.
/*
  can use flags: AVFMT_NOFILE, AVFMT_NEEDNUMBER,
  AVFMT_GLOBALHEADER, AVFMT_NOTIMESTAMPS, AVFMT_VARIABLE_FPS,
  AVFMT_NODIMENSIONS, AVFMT_NOSTREAMS,
  AVFMT_TS_NONSTRICT, AVFMT_TS_NEGATIVE
*/
func (s *AVOutputFormat) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
/*
  can use flags: AVFMT_NOFILE, AVFMT_NEEDNUMBER,
  AVFMT_GLOBALHEADER, AVFMT_NOTIMESTAMPS, AVFMT_VARIABLE_FPS,
  AVFMT_NODIMENSIONS, AVFMT_NOSTREAMS,
  AVFMT_TS_NONSTRICT, AVFMT_TS_NEGATIVE
*/
func (s *AVOutputFormat) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// CodecTag gets the codec_tag field.
/*
  List of supported codec_id-codec_tag pairs, ordered by "better
  choice first". The arrays are all terminated by AV_CODEC_ID_NONE.
*/
func (s *AVOutputFormat) CodecTag() *Array[*AVCodecTag] {
	value := s.ptr.codec_tag
	return ToAVCodecTagArray(unsafe.Pointer(value))
}

// SetCodecTag sets the codec_tag field.
/*
  List of supported codec_id-codec_tag pairs, ordered by "better
  choice first". The arrays are all terminated by AV_CODEC_ID_NONE.
*/
func (s *AVOutputFormat) SetCodecTag(value *Array[AVCodecTag]) {
	if value != nil {
		s.ptr.codec_tag = (**C.struct_AVCodecTag)(value.ptr)
	} else {
		s.ptr.codec_tag = nil
	}
}

// PrivClass gets the priv_class field.
//
//	AVClass for the private context
func (s *AVOutputFormat) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetPrivClass sets the priv_class field.
//
//	AVClass for the private context
func (s *AVOutputFormat) SetPrivClass(value *AVClass) {
	if value != nil {
		s.ptr.priv_class = value.ptr
	} else {
		s.ptr.priv_class = nil
	}
}

// --- Struct AVInputFormat ---

// AVInputFormat wraps AVInputFormat.
type AVInputFormat struct {
	ptr *C.AVInputFormat
}

func (s *AVInputFormat) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVInputFormatArray(ptr unsafe.Pointer) *Array[*AVInputFormat] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVInputFormat]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVInputFormat {
			ptr := (**C.AVInputFormat)(pointer)
			value := *ptr
			var valueMapped *AVInputFormat
			if value != nil {
				valueMapped = &AVInputFormat{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVInputFormat) {
			ptr := (**C.AVInputFormat)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Name gets the name field.
/*
  A comma separated list of short names for the format. New names
  may be appended with a minor bump.
*/
func (s *AVInputFormat) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
/*
  A comma separated list of short names for the format. New names
  may be appended with a minor bump.
*/
func (s *AVInputFormat) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// LongName gets the long_name field.
/*
  Descriptive name for the format, meant to be more human-readable
  than name. You should use the NULL_IF_CONFIG_SMALL() macro
  to define it.
*/
func (s *AVInputFormat) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

// SetLongName sets the long_name field.
/*
  Descriptive name for the format, meant to be more human-readable
  than name. You should use the NULL_IF_CONFIG_SMALL() macro
  to define it.
*/
func (s *AVInputFormat) SetLongName(value *CStr) {
	s.ptr.long_name = value.ptr
}

// Flags gets the flags field.
/*
  Can use flags: AVFMT_NOFILE, AVFMT_NEEDNUMBER, AVFMT_SHOW_IDS,
  AVFMT_NOTIMESTAMPS, AVFMT_GENERIC_INDEX, AVFMT_TS_DISCONT, AVFMT_NOBINSEARCH,
  AVFMT_NOGENSEARCH, AVFMT_NO_BYTE_SEEK, AVFMT_SEEK_TO_PTS.
*/
func (s *AVInputFormat) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
/*
  Can use flags: AVFMT_NOFILE, AVFMT_NEEDNUMBER, AVFMT_SHOW_IDS,
  AVFMT_NOTIMESTAMPS, AVFMT_GENERIC_INDEX, AVFMT_TS_DISCONT, AVFMT_NOBINSEARCH,
  AVFMT_NOGENSEARCH, AVFMT_NO_BYTE_SEEK, AVFMT_SEEK_TO_PTS.
*/
func (s *AVInputFormat) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// Extensions gets the extensions field.
/*
  If extensions are defined, then no probe is done. You should
  usually not use extension format guessing because it is not
  reliable enough
*/
func (s *AVInputFormat) Extensions() *CStr {
	value := s.ptr.extensions
	return wrapCStr(value)
}

// SetExtensions sets the extensions field.
/*
  If extensions are defined, then no probe is done. You should
  usually not use extension format guessing because it is not
  reliable enough
*/
func (s *AVInputFormat) SetExtensions(value *CStr) {
	s.ptr.extensions = value.ptr
}

// CodecTag gets the codec_tag field.
func (s *AVInputFormat) CodecTag() *Array[*AVCodecTag] {
	value := s.ptr.codec_tag
	return ToAVCodecTagArray(unsafe.Pointer(value))
}

// SetCodecTag sets the codec_tag field.
func (s *AVInputFormat) SetCodecTag(value *Array[AVCodecTag]) {
	if value != nil {
		s.ptr.codec_tag = (**C.struct_AVCodecTag)(value.ptr)
	} else {
		s.ptr.codec_tag = nil
	}
}

// PrivClass gets the priv_class field.
//
//	AVClass for the private context
func (s *AVInputFormat) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetPrivClass sets the priv_class field.
//
//	AVClass for the private context
func (s *AVInputFormat) SetPrivClass(value *AVClass) {
	if value != nil {
		s.ptr.priv_class = value.ptr
	} else {
		s.ptr.priv_class = nil
	}
}

// MimeType gets the mime_type field.
/*
  Comma-separated list of mime types.
  It is used check for matching mime types while probing.
  @see av_probe_input_format2
*/
func (s *AVInputFormat) MimeType() *CStr {
	value := s.ptr.mime_type
	return wrapCStr(value)
}

// SetMimeType sets the mime_type field.
/*
  Comma-separated list of mime types.
  It is used check for matching mime types while probing.
  @see av_probe_input_format2
*/
func (s *AVInputFormat) SetMimeType(value *CStr) {
	s.ptr.mime_type = value.ptr
}

// RawCodecId gets the raw_codec_id field.
//
//	Raw demuxers store their codec ID here.
func (s *AVInputFormat) RawCodecId() int {
	value := s.ptr.raw_codec_id
	return int(value)
}

// SetRawCodecId sets the raw_codec_id field.
//
//	Raw demuxers store their codec ID here.
func (s *AVInputFormat) SetRawCodecId(value int) {
	s.ptr.raw_codec_id = (C.int)(value)
}

// PrivDataSize gets the priv_data_size field.
//
//	Size of private data so that it can be allocated in the wrapper.
func (s *AVInputFormat) PrivDataSize() int {
	value := s.ptr.priv_data_size
	return int(value)
}

// SetPrivDataSize sets the priv_data_size field.
//
//	Size of private data so that it can be allocated in the wrapper.
func (s *AVInputFormat) SetPrivDataSize(value int) {
	s.ptr.priv_data_size = (C.int)(value)
}

// FlagsInternal gets the flags_internal field.
//
//	Internal flags. See FF_FMT_FLAG_* in internal.h.
func (s *AVInputFormat) FlagsInternal() int {
	value := s.ptr.flags_internal
	return int(value)
}

// SetFlagsInternal sets the flags_internal field.
//
//	Internal flags. See FF_FMT_FLAG_* in internal.h.
func (s *AVInputFormat) SetFlagsInternal(value int) {
	s.ptr.flags_internal = (C.int)(value)
}

// read_probe skipped due to func ptr

// read_header skipped due to func ptr

// read_packet skipped due to func ptr

// read_close skipped due to func ptr

// read_seek skipped due to func ptr

// read_timestamp skipped due to func ptr

// read_play skipped due to func ptr

// read_pause skipped due to func ptr

// read_seek2 skipped due to func ptr

// get_device_list skipped due to func ptr

// --- Struct AVIndexEntry ---

// AVIndexEntry wraps AVIndexEntry.
type AVIndexEntry struct {
	ptr *C.AVIndexEntry
}

func (s *AVIndexEntry) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVIndexEntryArray(ptr unsafe.Pointer) *Array[*AVIndexEntry] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVIndexEntry]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVIndexEntry {
			ptr := (**C.AVIndexEntry)(pointer)
			value := *ptr
			var valueMapped *AVIndexEntry
			if value != nil {
				valueMapped = &AVIndexEntry{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVIndexEntry) {
			ptr := (**C.AVIndexEntry)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Pos gets the pos field.
func (s *AVIndexEntry) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}

// SetPos sets the pos field.
func (s *AVIndexEntry) SetPos(value int64) {
	s.ptr.pos = (C.int64_t)(value)
}

// Timestamp gets the timestamp field.
/*
  <
  Timestamp in AVStream.time_base units, preferably the time from which on correctly decoded frames are available
  when seeking to this entry. That means preferable PTS on keyframe based formats.
  But demuxers can choose to store a different timestamp, if it is more convenient for the implementation or nothing better
  is known
*/
func (s *AVIndexEntry) Timestamp() int64 {
	value := s.ptr.timestamp
	return int64(value)
}

// SetTimestamp sets the timestamp field.
/*
  <
  Timestamp in AVStream.time_base units, preferably the time from which on correctly decoded frames are available
  when seeking to this entry. That means preferable PTS on keyframe based formats.
  But demuxers can choose to store a different timestamp, if it is more convenient for the implementation or nothing better
  is known
*/
func (s *AVIndexEntry) SetTimestamp(value int64) {
	s.ptr.timestamp = (C.int64_t)(value)
}

// flags skipped due to bitfield

// size skipped due to bitfield

// MinDistance gets the min_distance field.
//
//	Minimum distance between this and the previous keyframe, used to avoid unneeded searching.
func (s *AVIndexEntry) MinDistance() int {
	value := s.ptr.min_distance
	return int(value)
}

// SetMinDistance sets the min_distance field.
//
//	Minimum distance between this and the previous keyframe, used to avoid unneeded searching.
func (s *AVIndexEntry) SetMinDistance(value int) {
	s.ptr.min_distance = (C.int)(value)
}

// --- Struct AVStream ---

// AVStream wraps AVStream.
/*
  Stream structure.
  New fields can be added to the end with minor version bumps.
  Removal, reordering and changes to existing fields require a major
  version bump.
  sizeof(AVStream) must not be used outside libav*.
*/
type AVStream struct {
	ptr *C.AVStream
}

func (s *AVStream) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVStreamArray(ptr unsafe.Pointer) *Array[*AVStream] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVStream]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVStream {
			ptr := (**C.AVStream)(pointer)
			value := *ptr
			var valueMapped *AVStream
			if value != nil {
				valueMapped = &AVStream{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVStream) {
			ptr := (**C.AVStream)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// AvClass gets the av_class field.
//
//	A class for @ref avoptions. Set on stream creation.
func (s *AVStream) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
//
//	A class for @ref avoptions. Set on stream creation.
func (s *AVStream) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Index gets the index field.
//
//	stream index in AVFormatContext
func (s *AVStream) Index() int {
	value := s.ptr.index
	return int(value)
}

// SetIndex sets the index field.
//
//	stream index in AVFormatContext
func (s *AVStream) SetIndex(value int) {
	s.ptr.index = (C.int)(value)
}

// Id gets the id field.
/*
  Format-specific stream ID.
  decoding: set by libavformat
  encoding: set by the user, replaced by libavformat if left unset
*/
func (s *AVStream) Id() int {
	value := s.ptr.id
	return int(value)
}

// SetId sets the id field.
/*
  Format-specific stream ID.
  decoding: set by libavformat
  encoding: set by the user, replaced by libavformat if left unset
*/
func (s *AVStream) SetId(value int) {
	s.ptr.id = (C.int)(value)
}

// Codecpar gets the codecpar field.
/*
  Codec parameters associated with this stream. Allocated and freed by
  libavformat in avformat_new_stream() and avformat_free_context()
  respectively.

  - demuxing: filled by libavformat on stream creation or in
              avformat_find_stream_info()
  - muxing: filled by the caller before avformat_write_header()
*/
func (s *AVStream) Codecpar() *AVCodecParameters {
	value := s.ptr.codecpar
	var valueMapped *AVCodecParameters
	if value != nil {
		valueMapped = &AVCodecParameters{ptr: value}
	}
	return valueMapped
}

// SetCodecpar sets the codecpar field.
/*
  Codec parameters associated with this stream. Allocated and freed by
  libavformat in avformat_new_stream() and avformat_free_context()
  respectively.

  - demuxing: filled by libavformat on stream creation or in
              avformat_find_stream_info()
  - muxing: filled by the caller before avformat_write_header()
*/
func (s *AVStream) SetCodecpar(value *AVCodecParameters) {
	if value != nil {
		s.ptr.codecpar = value.ptr
	} else {
		s.ptr.codecpar = nil
	}
}

// PrivData gets the priv_data field.
func (s *AVStream) PrivData() unsafe.Pointer {
	value := s.ptr.priv_data
	return value
}

// SetPrivData sets the priv_data field.
func (s *AVStream) SetPrivData(value unsafe.Pointer) {
	s.ptr.priv_data = value
}

// TimeBase gets the time_base field.
/*
  This is the fundamental unit of time (in seconds) in terms
  of which frame timestamps are represented.

  decoding: set by libavformat
  encoding: May be set by the caller before avformat_write_header() to
            provide a hint to the muxer about the desired timebase. In
            avformat_write_header(), the muxer will overwrite this field
            with the timebase that will actually be used for the timestamps
            written into the file (which may or may not be related to the
            user-provided one, depending on the format).
*/
func (s *AVStream) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
/*
  This is the fundamental unit of time (in seconds) in terms
  of which frame timestamps are represented.

  decoding: set by libavformat
  encoding: May be set by the caller before avformat_write_header() to
            provide a hint to the muxer about the desired timebase. In
            avformat_write_header(), the muxer will overwrite this field
            with the timebase that will actually be used for the timestamps
            written into the file (which may or may not be related to the
            user-provided one, depending on the format).
*/
func (s *AVStream) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// StartTime gets the start_time field.
/*
  Decoding: pts of the first frame of the stream in presentation order, in stream time base.
  Only set this if you are absolutely 100% sure that the value you set
  it to really is the pts of the first frame.
  This may be undefined (AV_NOPTS_VALUE).
  @note The ASF header does NOT contain a correct start_time the ASF
  demuxer must NOT set this.
*/
func (s *AVStream) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}

// SetStartTime sets the start_time field.
/*
  Decoding: pts of the first frame of the stream in presentation order, in stream time base.
  Only set this if you are absolutely 100% sure that the value you set
  it to really is the pts of the first frame.
  This may be undefined (AV_NOPTS_VALUE).
  @note The ASF header does NOT contain a correct start_time the ASF
  demuxer must NOT set this.
*/
func (s *AVStream) SetStartTime(value int64) {
	s.ptr.start_time = (C.int64_t)(value)
}

// Duration gets the duration field.
/*
  Decoding: duration of the stream, in stream time base.
  If a source file does not specify a duration, but does specify
  a bitrate, this value will be estimated from bitrate and file size.

  Encoding: May be set by the caller before avformat_write_header() to
  provide a hint to the muxer about the estimated duration.
*/
func (s *AVStream) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

// SetDuration sets the duration field.
/*
  Decoding: duration of the stream, in stream time base.
  If a source file does not specify a duration, but does specify
  a bitrate, this value will be estimated from bitrate and file size.

  Encoding: May be set by the caller before avformat_write_header() to
  provide a hint to the muxer about the estimated duration.
*/
func (s *AVStream) SetDuration(value int64) {
	s.ptr.duration = (C.int64_t)(value)
}

// NbFrames gets the nb_frames field.
//
//	number of frames in this stream if known or 0
func (s *AVStream) NbFrames() int64 {
	value := s.ptr.nb_frames
	return int64(value)
}

// SetNbFrames sets the nb_frames field.
//
//	number of frames in this stream if known or 0
func (s *AVStream) SetNbFrames(value int64) {
	s.ptr.nb_frames = (C.int64_t)(value)
}

// Disposition gets the disposition field.
/*
  Stream disposition - a combination of AV_DISPOSITION_* flags.
  - demuxing: set by libavformat when creating the stream or in
              avformat_find_stream_info().
  - muxing: may be set by the caller before avformat_write_header().
*/
func (s *AVStream) Disposition() int {
	value := s.ptr.disposition
	return int(value)
}

// SetDisposition sets the disposition field.
/*
  Stream disposition - a combination of AV_DISPOSITION_* flags.
  - demuxing: set by libavformat when creating the stream or in
              avformat_find_stream_info().
  - muxing: may be set by the caller before avformat_write_header().
*/
func (s *AVStream) SetDisposition(value int) {
	s.ptr.disposition = (C.int)(value)
}

// Discard gets the discard field.
//
//	Selects which packets can be discarded at will and do not need to be demuxed.
func (s *AVStream) Discard() AVDiscard {
	value := s.ptr.discard
	return AVDiscard(value)
}

// SetDiscard sets the discard field.
//
//	Selects which packets can be discarded at will and do not need to be demuxed.
func (s *AVStream) SetDiscard(value AVDiscard) {
	s.ptr.discard = (C.enum_AVDiscard)(value)
}

// SampleAspectRatio gets the sample_aspect_ratio field.
/*
  sample aspect ratio (0 if unknown)
  - encoding: Set by user.
  - decoding: Set by libavformat.
*/
func (s *AVStream) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
/*
  sample aspect ratio (0 if unknown)
  - encoding: Set by user.
  - decoding: Set by libavformat.
*/
func (s *AVStream) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

// Metadata gets the metadata field.
func (s *AVStream) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

// SetMetadata sets the metadata field.
func (s *AVStream) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

// AvgFrameRate gets the avg_frame_rate field.
/*
  Average framerate

  - demuxing: May be set by libavformat when creating the stream or in
              avformat_find_stream_info().
  - muxing: May be set by the caller before avformat_write_header().
*/
func (s *AVStream) AvgFrameRate() *AVRational {
	value := s.ptr.avg_frame_rate
	return &AVRational{value: value}
}

// SetAvgFrameRate sets the avg_frame_rate field.
/*
  Average framerate

  - demuxing: May be set by libavformat when creating the stream or in
              avformat_find_stream_info().
  - muxing: May be set by the caller before avformat_write_header().
*/
func (s *AVStream) SetAvgFrameRate(value *AVRational) {
	s.ptr.avg_frame_rate = value.value
}

// AttachedPic gets the attached_pic field.
/*
  For streams with AV_DISPOSITION_ATTACHED_PIC disposition, this packet
  will contain the attached picture.

  decoding: set by libavformat, must not be modified by the caller.
  encoding: unused
*/
func (s *AVStream) AttachedPic() *AVPacket {
	value := &s.ptr.attached_pic
	return &AVPacket{ptr: value}
}

// SideData gets the side_data field.
/*
  An array of side data that applies to the whole stream (i.e. the
  container does not allow it to change between packets).

  There may be no overlap between the side data in this array and side data
  in the packets. I.e. a given side data is either exported by the muxer
  (demuxing) / set by the caller (muxing) in this array, then it never
  appears in the packets, or the side data is exported / sent through
  the packets (always in the first packet where the value becomes known or
  changes), then it does not appear in this array.

  - demuxing: Set by libavformat when the stream is created.
  - muxing: May be set by the caller before avformat_write_header().

  Freed by libavformat in avformat_free_context().

  @deprecated use AVStream's @ref AVCodecParameters.coded_side_data
              "codecpar side data".
*/
func (s *AVStream) SideData() *AVPacketSideData {
	value := s.ptr.side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}

// SetSideData sets the side_data field.
/*
  An array of side data that applies to the whole stream (i.e. the
  container does not allow it to change between packets).

  There may be no overlap between the side data in this array and side data
  in the packets. I.e. a given side data is either exported by the muxer
  (demuxing) / set by the caller (muxing) in this array, then it never
  appears in the packets, or the side data is exported / sent through
  the packets (always in the first packet where the value becomes known or
  changes), then it does not appear in this array.

  - demuxing: Set by libavformat when the stream is created.
  - muxing: May be set by the caller before avformat_write_header().

  Freed by libavformat in avformat_free_context().

  @deprecated use AVStream's @ref AVCodecParameters.coded_side_data
              "codecpar side data".
*/
func (s *AVStream) SetSideData(value *AVPacketSideData) {
	if value != nil {
		s.ptr.side_data = value.ptr
	} else {
		s.ptr.side_data = nil
	}
}

// NbSideData gets the nb_side_data field.
/*
  The number of elements in the AVStream.side_data array.

  @deprecated use AVStream's @ref AVCodecParameters.nb_coded_side_data
              "codecpar side data".
*/
func (s *AVStream) NbSideData() int {
	value := s.ptr.nb_side_data
	return int(value)
}

// SetNbSideData sets the nb_side_data field.
/*
  The number of elements in the AVStream.side_data array.

  @deprecated use AVStream's @ref AVCodecParameters.nb_coded_side_data
              "codecpar side data".
*/
func (s *AVStream) SetNbSideData(value int) {
	s.ptr.nb_side_data = (C.int)(value)
}

// EventFlags gets the event_flags field.
/*
  Flags indicating events happening on the stream, a combination of
  AVSTREAM_EVENT_FLAG_*.

  - demuxing: may be set by the demuxer in avformat_open_input(),
    avformat_find_stream_info() and av_read_frame(). Flags must be cleared
    by the user once the event has been handled.
  - muxing: may be set by the user after avformat_write_header(). to
    indicate a user-triggered event.  The muxer will clear the flags for
    events it has handled in av_[interleaved]_write_frame().
*/
func (s *AVStream) EventFlags() int {
	value := s.ptr.event_flags
	return int(value)
}

// SetEventFlags sets the event_flags field.
/*
  Flags indicating events happening on the stream, a combination of
  AVSTREAM_EVENT_FLAG_*.

  - demuxing: may be set by the demuxer in avformat_open_input(),
    avformat_find_stream_info() and av_read_frame(). Flags must be cleared
    by the user once the event has been handled.
  - muxing: may be set by the user after avformat_write_header(). to
    indicate a user-triggered event.  The muxer will clear the flags for
    events it has handled in av_[interleaved]_write_frame().
*/
func (s *AVStream) SetEventFlags(value int) {
	s.ptr.event_flags = (C.int)(value)
}

// RFrameRate gets the r_frame_rate field.
/*
  Real base framerate of the stream.
  This is the lowest framerate with which all timestamps can be
  represented accurately (it is the least common multiple of all
  framerates in the stream). Note, this value is just a guess!
  For example, if the time base is 1/90000 and all frames have either
  approximately 3600 or 1800 timer ticks, then r_frame_rate will be 50/1.
*/
func (s *AVStream) RFrameRate() *AVRational {
	value := s.ptr.r_frame_rate
	return &AVRational{value: value}
}

// SetRFrameRate sets the r_frame_rate field.
/*
  Real base framerate of the stream.
  This is the lowest framerate with which all timestamps can be
  represented accurately (it is the least common multiple of all
  framerates in the stream). Note, this value is just a guess!
  For example, if the time base is 1/90000 and all frames have either
  approximately 3600 or 1800 timer ticks, then r_frame_rate will be 50/1.
*/
func (s *AVStream) SetRFrameRate(value *AVRational) {
	s.ptr.r_frame_rate = value.value
}

// PtsWrapBits gets the pts_wrap_bits field.
/*
  Number of bits in timestamps. Used for wrapping control.

  - demuxing: set by libavformat
  - muxing: set by libavformat
*/
func (s *AVStream) PtsWrapBits() int {
	value := s.ptr.pts_wrap_bits
	return int(value)
}

// SetPtsWrapBits sets the pts_wrap_bits field.
/*
  Number of bits in timestamps. Used for wrapping control.

  - demuxing: set by libavformat
  - muxing: set by libavformat
*/
func (s *AVStream) SetPtsWrapBits(value int) {
	s.ptr.pts_wrap_bits = (C.int)(value)
}

// --- Struct AVProgram ---

// AVProgram wraps AVProgram.
/*
  New fields can be added to the end with minor version bumps.
  Removal, reordering and changes to existing fields require a major
  version bump.
  sizeof(AVProgram) must not be used outside libav*.
*/
type AVProgram struct {
	ptr *C.AVProgram
}

func (s *AVProgram) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVProgramArray(ptr unsafe.Pointer) *Array[*AVProgram] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVProgram]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVProgram {
			ptr := (**C.AVProgram)(pointer)
			value := *ptr
			var valueMapped *AVProgram
			if value != nil {
				valueMapped = &AVProgram{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVProgram) {
			ptr := (**C.AVProgram)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Id gets the id field.
func (s *AVProgram) Id() int {
	value := s.ptr.id
	return int(value)
}

// SetId sets the id field.
func (s *AVProgram) SetId(value int) {
	s.ptr.id = (C.int)(value)
}

// Flags gets the flags field.
func (s *AVProgram) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVProgram) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// Discard gets the discard field.
//
//	selects which program to discard and which to feed to the caller
func (s *AVProgram) Discard() AVDiscard {
	value := s.ptr.discard
	return AVDiscard(value)
}

// SetDiscard sets the discard field.
//
//	selects which program to discard and which to feed to the caller
func (s *AVProgram) SetDiscard(value AVDiscard) {
	s.ptr.discard = (C.enum_AVDiscard)(value)
}

// stream_index skipped due to prim ptr

// NbStreamIndexes gets the nb_stream_indexes field.
func (s *AVProgram) NbStreamIndexes() uint {
	value := s.ptr.nb_stream_indexes
	return uint(value)
}

// SetNbStreamIndexes sets the nb_stream_indexes field.
func (s *AVProgram) SetNbStreamIndexes(value uint) {
	s.ptr.nb_stream_indexes = (C.uint)(value)
}

// Metadata gets the metadata field.
func (s *AVProgram) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

// SetMetadata sets the metadata field.
func (s *AVProgram) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

// ProgramNum gets the program_num field.
func (s *AVProgram) ProgramNum() int {
	value := s.ptr.program_num
	return int(value)
}

// SetProgramNum sets the program_num field.
func (s *AVProgram) SetProgramNum(value int) {
	s.ptr.program_num = (C.int)(value)
}

// PmtPid gets the pmt_pid field.
func (s *AVProgram) PmtPid() int {
	value := s.ptr.pmt_pid
	return int(value)
}

// SetPmtPid sets the pmt_pid field.
func (s *AVProgram) SetPmtPid(value int) {
	s.ptr.pmt_pid = (C.int)(value)
}

// PcrPid gets the pcr_pid field.
func (s *AVProgram) PcrPid() int {
	value := s.ptr.pcr_pid
	return int(value)
}

// SetPcrPid sets the pcr_pid field.
func (s *AVProgram) SetPcrPid(value int) {
	s.ptr.pcr_pid = (C.int)(value)
}

// PmtVersion gets the pmt_version field.
func (s *AVProgram) PmtVersion() int {
	value := s.ptr.pmt_version
	return int(value)
}

// SetPmtVersion sets the pmt_version field.
func (s *AVProgram) SetPmtVersion(value int) {
	s.ptr.pmt_version = (C.int)(value)
}

// StartTime gets the start_time field.
/*
  ***************************************************************
  All fields below this line are not part of the public API. They
  may not be used outside of libavformat and can be changed and
  removed at will.
  New public fields should be added right above.
  *****************************************************************
*/
func (s *AVProgram) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}

// SetStartTime sets the start_time field.
/*
  ***************************************************************
  All fields below this line are not part of the public API. They
  may not be used outside of libavformat and can be changed and
  removed at will.
  New public fields should be added right above.
  *****************************************************************
*/
func (s *AVProgram) SetStartTime(value int64) {
	s.ptr.start_time = (C.int64_t)(value)
}

// EndTime gets the end_time field.
func (s *AVProgram) EndTime() int64 {
	value := s.ptr.end_time
	return int64(value)
}

// SetEndTime sets the end_time field.
func (s *AVProgram) SetEndTime(value int64) {
	s.ptr.end_time = (C.int64_t)(value)
}

// PtsWrapReference gets the pts_wrap_reference field.
//
//	reference dts for wrap detection
func (s *AVProgram) PtsWrapReference() int64 {
	value := s.ptr.pts_wrap_reference
	return int64(value)
}

// SetPtsWrapReference sets the pts_wrap_reference field.
//
//	reference dts for wrap detection
func (s *AVProgram) SetPtsWrapReference(value int64) {
	s.ptr.pts_wrap_reference = (C.int64_t)(value)
}

// PtsWrapBehavior gets the pts_wrap_behavior field.
//
//	behavior on wrap detection
func (s *AVProgram) PtsWrapBehavior() int {
	value := s.ptr.pts_wrap_behavior
	return int(value)
}

// SetPtsWrapBehavior sets the pts_wrap_behavior field.
//
//	behavior on wrap detection
func (s *AVProgram) SetPtsWrapBehavior(value int) {
	s.ptr.pts_wrap_behavior = (C.int)(value)
}

// --- Struct AVChapter ---

// AVChapter wraps AVChapter.
type AVChapter struct {
	ptr *C.AVChapter
}

func (s *AVChapter) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVChapterArray(ptr unsafe.Pointer) *Array[*AVChapter] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVChapter]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVChapter {
			ptr := (**C.AVChapter)(pointer)
			value := *ptr
			var valueMapped *AVChapter
			if value != nil {
				valueMapped = &AVChapter{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVChapter) {
			ptr := (**C.AVChapter)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Id gets the id field.
//
//	unique ID to identify the chapter
func (s *AVChapter) Id() int64 {
	value := s.ptr.id
	return int64(value)
}

// SetId sets the id field.
//
//	unique ID to identify the chapter
func (s *AVChapter) SetId(value int64) {
	s.ptr.id = (C.int64_t)(value)
}

// TimeBase gets the time_base field.
//
//	time base in which the start/end timestamps are specified
func (s *AVChapter) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
//
//	time base in which the start/end timestamps are specified
func (s *AVChapter) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// Start gets the start field.
//
//	chapter start/end time in time_base units
func (s *AVChapter) Start() int64 {
	value := s.ptr.start
	return int64(value)
}

// SetStart sets the start field.
//
//	chapter start/end time in time_base units
func (s *AVChapter) SetStart(value int64) {
	s.ptr.start = (C.int64_t)(value)
}

// End gets the end field.
//
//	chapter start/end time in time_base units
func (s *AVChapter) End() int64 {
	value := s.ptr.end
	return int64(value)
}

// SetEnd sets the end field.
//
//	chapter start/end time in time_base units
func (s *AVChapter) SetEnd(value int64) {
	s.ptr.end = (C.int64_t)(value)
}

// Metadata gets the metadata field.
func (s *AVChapter) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

// SetMetadata sets the metadata field.
func (s *AVChapter) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

// --- Struct AVFormatContext ---

// AVFormatContext wraps AVFormatContext.
/*
  Format I/O context.
  New fields can be added to the end with minor version bumps.
  Removal, reordering and changes to existing fields require a major
  version bump.
  sizeof(AVFormatContext) must not be used outside libav*, use
  avformat_alloc_context() to create an AVFormatContext.

  Fields can be accessed through AVOptions (av_opt*),
  the name string used matches the associated command line parameter name and
  can be found in libavformat/options_table.h.
  The AVOption/command line parameter names differ in some cases from the C
  structure field names for historic reasons or brevity.
*/
type AVFormatContext struct {
	ptr *C.AVFormatContext
}

func (s *AVFormatContext) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFormatContextArray(ptr unsafe.Pointer) *Array[*AVFormatContext] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFormatContext]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFormatContext {
			ptr := (**C.AVFormatContext)(pointer)
			value := *ptr
			var valueMapped *AVFormatContext
			if value != nil {
				valueMapped = &AVFormatContext{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFormatContext) {
			ptr := (**C.AVFormatContext)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// AvClass gets the av_class field.
/*
  A class for logging and @ref avoptions. Set by avformat_alloc_context().
  Exports (de)muxer private options if they exist.
*/
func (s *AVFormatContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
/*
  A class for logging and @ref avoptions. Set by avformat_alloc_context().
  Exports (de)muxer private options if they exist.
*/
func (s *AVFormatContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Iformat gets the iformat field.
/*
  The input container format.

  Demuxing only, set by avformat_open_input().
*/
func (s *AVFormatContext) Iformat() *AVInputFormat {
	value := s.ptr.iformat
	var valueMapped *AVInputFormat
	if value != nil {
		valueMapped = &AVInputFormat{ptr: value}
	}
	return valueMapped
}

// SetIformat sets the iformat field.
/*
  The input container format.

  Demuxing only, set by avformat_open_input().
*/
func (s *AVFormatContext) SetIformat(value *AVInputFormat) {
	if value != nil {
		s.ptr.iformat = value.ptr
	} else {
		s.ptr.iformat = nil
	}
}

// Oformat gets the oformat field.
/*
  The output container format.

  Muxing only, must be set by the caller before avformat_write_header().
*/
func (s *AVFormatContext) Oformat() *AVOutputFormat {
	value := s.ptr.oformat
	var valueMapped *AVOutputFormat
	if value != nil {
		valueMapped = &AVOutputFormat{ptr: value}
	}
	return valueMapped
}

// SetOformat sets the oformat field.
/*
  The output container format.

  Muxing only, must be set by the caller before avformat_write_header().
*/
func (s *AVFormatContext) SetOformat(value *AVOutputFormat) {
	if value != nil {
		s.ptr.oformat = value.ptr
	} else {
		s.ptr.oformat = nil
	}
}

// PrivData gets the priv_data field.
/*
  Format private data. This is an AVOptions-enabled struct
  if and only if iformat/oformat.priv_class is not NULL.

  - muxing: set by avformat_write_header()
  - demuxing: set by avformat_open_input()
*/
func (s *AVFormatContext) PrivData() unsafe.Pointer {
	value := s.ptr.priv_data
	return value
}

// SetPrivData sets the priv_data field.
/*
  Format private data. This is an AVOptions-enabled struct
  if and only if iformat/oformat.priv_class is not NULL.

  - muxing: set by avformat_write_header()
  - demuxing: set by avformat_open_input()
*/
func (s *AVFormatContext) SetPrivData(value unsafe.Pointer) {
	s.ptr.priv_data = value
}

// Pb gets the pb field.
/*
  I/O context.

  - demuxing: either set by the user before avformat_open_input() (then
              the user must close it manually) or set by avformat_open_input().
  - muxing: set by the user before avformat_write_header(). The caller must
            take care of closing / freeing the IO context.

  Do NOT set this field if AVFMT_NOFILE flag is set in
  iformat/oformat.flags. In such a case, the (de)muxer will handle
  I/O in some other way and this field will be NULL.
*/
func (s *AVFormatContext) Pb() *AVIOContext {
	value := s.ptr.pb
	var valueMapped *AVIOContext
	if value != nil {
		valueMapped = &AVIOContext{ptr: value}
	}
	return valueMapped
}

// SetPb sets the pb field.
/*
  I/O context.

  - demuxing: either set by the user before avformat_open_input() (then
              the user must close it manually) or set by avformat_open_input().
  - muxing: set by the user before avformat_write_header(). The caller must
            take care of closing / freeing the IO context.

  Do NOT set this field if AVFMT_NOFILE flag is set in
  iformat/oformat.flags. In such a case, the (de)muxer will handle
  I/O in some other way and this field will be NULL.
*/
func (s *AVFormatContext) SetPb(value *AVIOContext) {
	if value != nil {
		s.ptr.pb = value.ptr
	} else {
		s.ptr.pb = nil
	}
}

// CtxFlags gets the ctx_flags field.
/*
  Flags signalling stream properties. A combination of AVFMTCTX_*.
  Set by libavformat.
*/
func (s *AVFormatContext) CtxFlags() int {
	value := s.ptr.ctx_flags
	return int(value)
}

// SetCtxFlags sets the ctx_flags field.
/*
  Flags signalling stream properties. A combination of AVFMTCTX_*.
  Set by libavformat.
*/
func (s *AVFormatContext) SetCtxFlags(value int) {
	s.ptr.ctx_flags = (C.int)(value)
}

// NbStreams gets the nb_streams field.
/*
  Number of elements in AVFormatContext.streams.

  Set by avformat_new_stream(), must not be modified by any other code.
*/
func (s *AVFormatContext) NbStreams() uint {
	value := s.ptr.nb_streams
	return uint(value)
}

// SetNbStreams sets the nb_streams field.
/*
  Number of elements in AVFormatContext.streams.

  Set by avformat_new_stream(), must not be modified by any other code.
*/
func (s *AVFormatContext) SetNbStreams(value uint) {
	s.ptr.nb_streams = (C.uint)(value)
}

// Streams gets the streams field.
/*
  A list of all streams in the file. New streams are created with
  avformat_new_stream().

  - demuxing: streams are created by libavformat in avformat_open_input().
              If AVFMTCTX_NOHEADER is set in ctx_flags, then new streams may also
              appear in av_read_frame().
  - muxing: streams are created by the user before avformat_write_header().

  Freed by libavformat in avformat_free_context().
*/
func (s *AVFormatContext) Streams() *Array[*AVStream] {
	value := s.ptr.streams
	return ToAVStreamArray(unsafe.Pointer(value))
}

// SetStreams sets the streams field.
/*
  A list of all streams in the file. New streams are created with
  avformat_new_stream().

  - demuxing: streams are created by libavformat in avformat_open_input().
              If AVFMTCTX_NOHEADER is set in ctx_flags, then new streams may also
              appear in av_read_frame().
  - muxing: streams are created by the user before avformat_write_header().

  Freed by libavformat in avformat_free_context().
*/
func (s *AVFormatContext) SetStreams(value *Array[AVStream]) {
	if value != nil {
		s.ptr.streams = (**C.AVStream)(value.ptr)
	} else {
		s.ptr.streams = nil
	}
}

// Url gets the url field.
/*
  input or output URL. Unlike the old filename field, this field has no
  length restriction.

  - demuxing: set by avformat_open_input(), initialized to an empty
              string if url parameter was NULL in avformat_open_input().
  - muxing: may be set by the caller before calling avformat_write_header()
            (or avformat_init_output() if that is called first) to a string
            which is freeable by av_free(). Set to an empty string if it
            was NULL in avformat_init_output().

  Freed by libavformat in avformat_free_context().
*/
func (s *AVFormatContext) Url() *CStr {
	value := s.ptr.url
	return wrapCStr(value)
}

// SetUrl sets the url field.
/*
  input or output URL. Unlike the old filename field, this field has no
  length restriction.

  - demuxing: set by avformat_open_input(), initialized to an empty
              string if url parameter was NULL in avformat_open_input().
  - muxing: may be set by the caller before calling avformat_write_header()
            (or avformat_init_output() if that is called first) to a string
            which is freeable by av_free(). Set to an empty string if it
            was NULL in avformat_init_output().

  Freed by libavformat in avformat_free_context().
*/
func (s *AVFormatContext) SetUrl(value *CStr) {
	s.ptr.url = value.ptr
}

// StartTime gets the start_time field.
/*
  Position of the first frame of the component, in
  AV_TIME_BASE fractional seconds. NEVER set this value directly:
  It is deduced from the AVStream values.

  Demuxing only, set by libavformat.
*/
func (s *AVFormatContext) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}

// SetStartTime sets the start_time field.
/*
  Position of the first frame of the component, in
  AV_TIME_BASE fractional seconds. NEVER set this value directly:
  It is deduced from the AVStream values.

  Demuxing only, set by libavformat.
*/
func (s *AVFormatContext) SetStartTime(value int64) {
	s.ptr.start_time = (C.int64_t)(value)
}

// Duration gets the duration field.
/*
  Duration of the stream, in AV_TIME_BASE fractional
  seconds. Only set this value if you know none of the individual stream
  durations and also do not set any of them. This is deduced from the
  AVStream values if not set.

  Demuxing only, set by libavformat.
*/
func (s *AVFormatContext) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

// SetDuration sets the duration field.
/*
  Duration of the stream, in AV_TIME_BASE fractional
  seconds. Only set this value if you know none of the individual stream
  durations and also do not set any of them. This is deduced from the
  AVStream values if not set.

  Demuxing only, set by libavformat.
*/
func (s *AVFormatContext) SetDuration(value int64) {
	s.ptr.duration = (C.int64_t)(value)
}

// BitRate gets the bit_rate field.
/*
  Total stream bitrate in bit/s, 0 if not
  available. Never set it directly if the file_size and the
  duration are known as FFmpeg can compute it automatically.
*/
func (s *AVFormatContext) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}

// SetBitRate sets the bit_rate field.
/*
  Total stream bitrate in bit/s, 0 if not
  available. Never set it directly if the file_size and the
  duration are known as FFmpeg can compute it automatically.
*/
func (s *AVFormatContext) SetBitRate(value int64) {
	s.ptr.bit_rate = (C.int64_t)(value)
}

// PacketSize gets the packet_size field.
func (s *AVFormatContext) PacketSize() uint {
	value := s.ptr.packet_size
	return uint(value)
}

// SetPacketSize sets the packet_size field.
func (s *AVFormatContext) SetPacketSize(value uint) {
	s.ptr.packet_size = (C.uint)(value)
}

// MaxDelay gets the max_delay field.
func (s *AVFormatContext) MaxDelay() int {
	value := s.ptr.max_delay
	return int(value)
}

// SetMaxDelay sets the max_delay field.
func (s *AVFormatContext) SetMaxDelay(value int) {
	s.ptr.max_delay = (C.int)(value)
}

// Flags gets the flags field.
/*
  Flags modifying the (de)muxer behaviour. A combination of AVFMT_FLAG_*.
  Set by the user before avformat_open_input() / avformat_write_header().
*/
func (s *AVFormatContext) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
/*
  Flags modifying the (de)muxer behaviour. A combination of AVFMT_FLAG_*.
  Set by the user before avformat_open_input() / avformat_write_header().
*/
func (s *AVFormatContext) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// Probesize gets the probesize field.
/*
  Maximum number of bytes read from input in order to determine stream
  properties. Used when reading the global header and in
  avformat_find_stream_info().

  Demuxing only, set by the caller before avformat_open_input().

  @note this is \e not  used for determining the \ref AVInputFormat
        "input format"
  @sa format_probesize
*/
func (s *AVFormatContext) Probesize() int64 {
	value := s.ptr.probesize
	return int64(value)
}

// SetProbesize sets the probesize field.
/*
  Maximum number of bytes read from input in order to determine stream
  properties. Used when reading the global header and in
  avformat_find_stream_info().

  Demuxing only, set by the caller before avformat_open_input().

  @note this is \e not  used for determining the \ref AVInputFormat
        "input format"
  @sa format_probesize
*/
func (s *AVFormatContext) SetProbesize(value int64) {
	s.ptr.probesize = (C.int64_t)(value)
}

// MaxAnalyzeDuration gets the max_analyze_duration field.
/*
  Maximum duration (in AV_TIME_BASE units) of the data read
  from input in avformat_find_stream_info().
  Demuxing only, set by the caller before avformat_find_stream_info().
  Can be set to 0 to let avformat choose using a heuristic.
*/
func (s *AVFormatContext) MaxAnalyzeDuration() int64 {
	value := s.ptr.max_analyze_duration
	return int64(value)
}

// SetMaxAnalyzeDuration sets the max_analyze_duration field.
/*
  Maximum duration (in AV_TIME_BASE units) of the data read
  from input in avformat_find_stream_info().
  Demuxing only, set by the caller before avformat_find_stream_info().
  Can be set to 0 to let avformat choose using a heuristic.
*/
func (s *AVFormatContext) SetMaxAnalyzeDuration(value int64) {
	s.ptr.max_analyze_duration = (C.int64_t)(value)
}

// Key gets the key field.
func (s *AVFormatContext) Key() unsafe.Pointer {
	value := s.ptr.key
	return unsafe.Pointer(value)
}

// SetKey sets the key field.
func (s *AVFormatContext) SetKey(value unsafe.Pointer) {
	s.ptr.key = (*C.uint8_t)(value)
}

// Keylen gets the keylen field.
func (s *AVFormatContext) Keylen() int {
	value := s.ptr.keylen
	return int(value)
}

// SetKeylen sets the keylen field.
func (s *AVFormatContext) SetKeylen(value int) {
	s.ptr.keylen = (C.int)(value)
}

// NbPrograms gets the nb_programs field.
func (s *AVFormatContext) NbPrograms() uint {
	value := s.ptr.nb_programs
	return uint(value)
}

// SetNbPrograms sets the nb_programs field.
func (s *AVFormatContext) SetNbPrograms(value uint) {
	s.ptr.nb_programs = (C.uint)(value)
}

// Programs gets the programs field.
func (s *AVFormatContext) Programs() *Array[*AVProgram] {
	value := s.ptr.programs
	return ToAVProgramArray(unsafe.Pointer(value))
}

// SetPrograms sets the programs field.
func (s *AVFormatContext) SetPrograms(value *Array[AVProgram]) {
	if value != nil {
		s.ptr.programs = (**C.AVProgram)(value.ptr)
	} else {
		s.ptr.programs = nil
	}
}

// VideoCodecId gets the video_codec_id field.
/*
  Forced video codec_id.
  Demuxing: Set by user.
*/
func (s *AVFormatContext) VideoCodecId() AVCodecID {
	value := s.ptr.video_codec_id
	return AVCodecID(value)
}

// SetVideoCodecId sets the video_codec_id field.
/*
  Forced video codec_id.
  Demuxing: Set by user.
*/
func (s *AVFormatContext) SetVideoCodecId(value AVCodecID) {
	s.ptr.video_codec_id = (C.enum_AVCodecID)(value)
}

// AudioCodecId gets the audio_codec_id field.
/*
  Forced audio codec_id.
  Demuxing: Set by user.
*/
func (s *AVFormatContext) AudioCodecId() AVCodecID {
	value := s.ptr.audio_codec_id
	return AVCodecID(value)
}

// SetAudioCodecId sets the audio_codec_id field.
/*
  Forced audio codec_id.
  Demuxing: Set by user.
*/
func (s *AVFormatContext) SetAudioCodecId(value AVCodecID) {
	s.ptr.audio_codec_id = (C.enum_AVCodecID)(value)
}

// SubtitleCodecId gets the subtitle_codec_id field.
/*
  Forced subtitle codec_id.
  Demuxing: Set by user.
*/
func (s *AVFormatContext) SubtitleCodecId() AVCodecID {
	value := s.ptr.subtitle_codec_id
	return AVCodecID(value)
}

// SetSubtitleCodecId sets the subtitle_codec_id field.
/*
  Forced subtitle codec_id.
  Demuxing: Set by user.
*/
func (s *AVFormatContext) SetSubtitleCodecId(value AVCodecID) {
	s.ptr.subtitle_codec_id = (C.enum_AVCodecID)(value)
}

// MaxIndexSize gets the max_index_size field.
/*
  Maximum amount of memory in bytes to use for the index of each stream.
  If the index exceeds this size, entries will be discarded as
  needed to maintain a smaller size. This can lead to slower or less
  accurate seeking (depends on demuxer).
  Demuxers for which a full in-memory index is mandatory will ignore
  this.
  - muxing: unused
  - demuxing: set by user
*/
func (s *AVFormatContext) MaxIndexSize() uint {
	value := s.ptr.max_index_size
	return uint(value)
}

// SetMaxIndexSize sets the max_index_size field.
/*
  Maximum amount of memory in bytes to use for the index of each stream.
  If the index exceeds this size, entries will be discarded as
  needed to maintain a smaller size. This can lead to slower or less
  accurate seeking (depends on demuxer).
  Demuxers for which a full in-memory index is mandatory will ignore
  this.
  - muxing: unused
  - demuxing: set by user
*/
func (s *AVFormatContext) SetMaxIndexSize(value uint) {
	s.ptr.max_index_size = (C.uint)(value)
}

// MaxPictureBuffer gets the max_picture_buffer field.
/*
  Maximum amount of memory in bytes to use for buffering frames
  obtained from realtime capture devices.
*/
func (s *AVFormatContext) MaxPictureBuffer() uint {
	value := s.ptr.max_picture_buffer
	return uint(value)
}

// SetMaxPictureBuffer sets the max_picture_buffer field.
/*
  Maximum amount of memory in bytes to use for buffering frames
  obtained from realtime capture devices.
*/
func (s *AVFormatContext) SetMaxPictureBuffer(value uint) {
	s.ptr.max_picture_buffer = (C.uint)(value)
}

// NbChapters gets the nb_chapters field.
/*
  Number of chapters in AVChapter array.
  When muxing, chapters are normally written in the file header,
  so nb_chapters should normally be initialized before write_header
  is called. Some muxers (e.g. mov and mkv) can also write chapters
  in the trailer.  To write chapters in the trailer, nb_chapters
  must be zero when write_header is called and non-zero when
  write_trailer is called.
  - muxing: set by user
  - demuxing: set by libavformat
*/
func (s *AVFormatContext) NbChapters() uint {
	value := s.ptr.nb_chapters
	return uint(value)
}

// SetNbChapters sets the nb_chapters field.
/*
  Number of chapters in AVChapter array.
  When muxing, chapters are normally written in the file header,
  so nb_chapters should normally be initialized before write_header
  is called. Some muxers (e.g. mov and mkv) can also write chapters
  in the trailer.  To write chapters in the trailer, nb_chapters
  must be zero when write_header is called and non-zero when
  write_trailer is called.
  - muxing: set by user
  - demuxing: set by libavformat
*/
func (s *AVFormatContext) SetNbChapters(value uint) {
	s.ptr.nb_chapters = (C.uint)(value)
}

// Chapters gets the chapters field.
func (s *AVFormatContext) Chapters() *Array[*AVChapter] {
	value := s.ptr.chapters
	return ToAVChapterArray(unsafe.Pointer(value))
}

// SetChapters sets the chapters field.
func (s *AVFormatContext) SetChapters(value *Array[AVChapter]) {
	if value != nil {
		s.ptr.chapters = (**C.AVChapter)(value.ptr)
	} else {
		s.ptr.chapters = nil
	}
}

// Metadata gets the metadata field.
/*
  Metadata that applies to the whole file.

  - demuxing: set by libavformat in avformat_open_input()
  - muxing: may be set by the caller before avformat_write_header()

  Freed by libavformat in avformat_free_context().
*/
func (s *AVFormatContext) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

// SetMetadata sets the metadata field.
/*
  Metadata that applies to the whole file.

  - demuxing: set by libavformat in avformat_open_input()
  - muxing: may be set by the caller before avformat_write_header()

  Freed by libavformat in avformat_free_context().
*/
func (s *AVFormatContext) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

// StartTimeRealtime gets the start_time_realtime field.
/*
  Start time of the stream in real world time, in microseconds
  since the Unix epoch (00:00 1st January 1970). That is, pts=0 in the
  stream was captured at this real world time.
  - muxing: Set by the caller before avformat_write_header(). If set to
            either 0 or AV_NOPTS_VALUE, then the current wall-time will
            be used.
  - demuxing: Set by libavformat. AV_NOPTS_VALUE if unknown. Note that
              the value may become known after some number of frames
              have been received.
*/
func (s *AVFormatContext) StartTimeRealtime() int64 {
	value := s.ptr.start_time_realtime
	return int64(value)
}

// SetStartTimeRealtime sets the start_time_realtime field.
/*
  Start time of the stream in real world time, in microseconds
  since the Unix epoch (00:00 1st January 1970). That is, pts=0 in the
  stream was captured at this real world time.
  - muxing: Set by the caller before avformat_write_header(). If set to
            either 0 or AV_NOPTS_VALUE, then the current wall-time will
            be used.
  - demuxing: Set by libavformat. AV_NOPTS_VALUE if unknown. Note that
              the value may become known after some number of frames
              have been received.
*/
func (s *AVFormatContext) SetStartTimeRealtime(value int64) {
	s.ptr.start_time_realtime = (C.int64_t)(value)
}

// FpsProbeSize gets the fps_probe_size field.
/*
  The number of frames used for determining the framerate in
  avformat_find_stream_info().
  Demuxing only, set by the caller before avformat_find_stream_info().
*/
func (s *AVFormatContext) FpsProbeSize() int {
	value := s.ptr.fps_probe_size
	return int(value)
}

// SetFpsProbeSize sets the fps_probe_size field.
/*
  The number of frames used for determining the framerate in
  avformat_find_stream_info().
  Demuxing only, set by the caller before avformat_find_stream_info().
*/
func (s *AVFormatContext) SetFpsProbeSize(value int) {
	s.ptr.fps_probe_size = (C.int)(value)
}

// ErrorRecognition gets the error_recognition field.
/*
  Error recognition; higher values will detect more errors but may
  misdetect some more or less valid parts as errors.
  Demuxing only, set by the caller before avformat_open_input().
*/
func (s *AVFormatContext) ErrorRecognition() int {
	value := s.ptr.error_recognition
	return int(value)
}

// SetErrorRecognition sets the error_recognition field.
/*
  Error recognition; higher values will detect more errors but may
  misdetect some more or less valid parts as errors.
  Demuxing only, set by the caller before avformat_open_input().
*/
func (s *AVFormatContext) SetErrorRecognition(value int) {
	s.ptr.error_recognition = (C.int)(value)
}

// InterruptCallback gets the interrupt_callback field.
/*
  Custom interrupt callbacks for the I/O layer.

  demuxing: set by the user before avformat_open_input().
  muxing: set by the user before avformat_write_header()
  (mainly useful for AVFMT_NOFILE formats). The callback
  should also be passed to avio_open2() if it's used to
  open the file.
*/
func (s *AVFormatContext) InterruptCallback() *AVIOInterruptCB {
	value := &s.ptr.interrupt_callback
	return &AVIOInterruptCB{ptr: value}
}

// Debug gets the debug field.
//
//	Flags to enable debugging.
func (s *AVFormatContext) Debug() int {
	value := s.ptr.debug
	return int(value)
}

// SetDebug sets the debug field.
//
//	Flags to enable debugging.
func (s *AVFormatContext) SetDebug(value int) {
	s.ptr.debug = (C.int)(value)
}

// MaxInterleaveDelta gets the max_interleave_delta field.
/*
  Maximum buffering duration for interleaving.

  To ensure all the streams are interleaved correctly,
  av_interleaved_write_frame() will wait until it has at least one packet
  for each stream before actually writing any packets to the output file.
  When some streams are "sparse" (i.e. there are large gaps between
  successive packets), this can result in excessive buffering.

  This field specifies the maximum difference between the timestamps of the
  first and the last packet in the muxing queue, above which libavformat
  will output a packet regardless of whether it has queued a packet for all
  the streams.

  Muxing only, set by the caller before avformat_write_header().
*/
func (s *AVFormatContext) MaxInterleaveDelta() int64 {
	value := s.ptr.max_interleave_delta
	return int64(value)
}

// SetMaxInterleaveDelta sets the max_interleave_delta field.
/*
  Maximum buffering duration for interleaving.

  To ensure all the streams are interleaved correctly,
  av_interleaved_write_frame() will wait until it has at least one packet
  for each stream before actually writing any packets to the output file.
  When some streams are "sparse" (i.e. there are large gaps between
  successive packets), this can result in excessive buffering.

  This field specifies the maximum difference between the timestamps of the
  first and the last packet in the muxing queue, above which libavformat
  will output a packet regardless of whether it has queued a packet for all
  the streams.

  Muxing only, set by the caller before avformat_write_header().
*/
func (s *AVFormatContext) SetMaxInterleaveDelta(value int64) {
	s.ptr.max_interleave_delta = (C.int64_t)(value)
}

// StrictStdCompliance gets the strict_std_compliance field.
/*
  Allow non-standard and experimental extension
  @see AVCodecContext.strict_std_compliance
*/
func (s *AVFormatContext) StrictStdCompliance() int {
	value := s.ptr.strict_std_compliance
	return int(value)
}

// SetStrictStdCompliance sets the strict_std_compliance field.
/*
  Allow non-standard and experimental extension
  @see AVCodecContext.strict_std_compliance
*/
func (s *AVFormatContext) SetStrictStdCompliance(value int) {
	s.ptr.strict_std_compliance = (C.int)(value)
}

// EventFlags gets the event_flags field.
/*
  Flags indicating events happening on the file, a combination of
  AVFMT_EVENT_FLAG_*.

  - demuxing: may be set by the demuxer in avformat_open_input(),
    avformat_find_stream_info() and av_read_frame(). Flags must be cleared
    by the user once the event has been handled.
  - muxing: may be set by the user after avformat_write_header() to
    indicate a user-triggered event.  The muxer will clear the flags for
    events it has handled in av_[interleaved]_write_frame().
*/
func (s *AVFormatContext) EventFlags() int {
	value := s.ptr.event_flags
	return int(value)
}

// SetEventFlags sets the event_flags field.
/*
  Flags indicating events happening on the file, a combination of
  AVFMT_EVENT_FLAG_*.

  - demuxing: may be set by the demuxer in avformat_open_input(),
    avformat_find_stream_info() and av_read_frame(). Flags must be cleared
    by the user once the event has been handled.
  - muxing: may be set by the user after avformat_write_header() to
    indicate a user-triggered event.  The muxer will clear the flags for
    events it has handled in av_[interleaved]_write_frame().
*/
func (s *AVFormatContext) SetEventFlags(value int) {
	s.ptr.event_flags = (C.int)(value)
}

// MaxTsProbe gets the max_ts_probe field.
/*
  Maximum number of packets to read while waiting for the first timestamp.
  Decoding only.
*/
func (s *AVFormatContext) MaxTsProbe() int {
	value := s.ptr.max_ts_probe
	return int(value)
}

// SetMaxTsProbe sets the max_ts_probe field.
/*
  Maximum number of packets to read while waiting for the first timestamp.
  Decoding only.
*/
func (s *AVFormatContext) SetMaxTsProbe(value int) {
	s.ptr.max_ts_probe = (C.int)(value)
}

// AvoidNegativeTs gets the avoid_negative_ts field.
/*
  Avoid negative timestamps during muxing.
  Any value of the AVFMT_AVOID_NEG_TS_* constants.
  Note, this works better when using av_interleaved_write_frame().
  - muxing: Set by user
  - demuxing: unused
*/
func (s *AVFormatContext) AvoidNegativeTs() int {
	value := s.ptr.avoid_negative_ts
	return int(value)
}

// SetAvoidNegativeTs sets the avoid_negative_ts field.
/*
  Avoid negative timestamps during muxing.
  Any value of the AVFMT_AVOID_NEG_TS_* constants.
  Note, this works better when using av_interleaved_write_frame().
  - muxing: Set by user
  - demuxing: unused
*/
func (s *AVFormatContext) SetAvoidNegativeTs(value int) {
	s.ptr.avoid_negative_ts = (C.int)(value)
}

// TsId gets the ts_id field.
/*
  Transport stream id.
  This will be moved into demuxer private options. Thus no API/ABI compatibility
*/
func (s *AVFormatContext) TsId() int {
	value := s.ptr.ts_id
	return int(value)
}

// SetTsId sets the ts_id field.
/*
  Transport stream id.
  This will be moved into demuxer private options. Thus no API/ABI compatibility
*/
func (s *AVFormatContext) SetTsId(value int) {
	s.ptr.ts_id = (C.int)(value)
}

// AudioPreload gets the audio_preload field.
/*
  Audio preload in microseconds.
  Note, not all formats support this and unpredictable things may happen if it is used when not supported.
  - encoding: Set by user
  - decoding: unused
*/
func (s *AVFormatContext) AudioPreload() int {
	value := s.ptr.audio_preload
	return int(value)
}

// SetAudioPreload sets the audio_preload field.
/*
  Audio preload in microseconds.
  Note, not all formats support this and unpredictable things may happen if it is used when not supported.
  - encoding: Set by user
  - decoding: unused
*/
func (s *AVFormatContext) SetAudioPreload(value int) {
	s.ptr.audio_preload = (C.int)(value)
}

// MaxChunkDuration gets the max_chunk_duration field.
/*
  Max chunk time in microseconds.
  Note, not all formats support this and unpredictable things may happen if it is used when not supported.
  - encoding: Set by user
  - decoding: unused
*/
func (s *AVFormatContext) MaxChunkDuration() int {
	value := s.ptr.max_chunk_duration
	return int(value)
}

// SetMaxChunkDuration sets the max_chunk_duration field.
/*
  Max chunk time in microseconds.
  Note, not all formats support this and unpredictable things may happen if it is used when not supported.
  - encoding: Set by user
  - decoding: unused
*/
func (s *AVFormatContext) SetMaxChunkDuration(value int) {
	s.ptr.max_chunk_duration = (C.int)(value)
}

// MaxChunkSize gets the max_chunk_size field.
/*
  Max chunk size in bytes
  Note, not all formats support this and unpredictable things may happen if it is used when not supported.
  - encoding: Set by user
  - decoding: unused
*/
func (s *AVFormatContext) MaxChunkSize() int {
	value := s.ptr.max_chunk_size
	return int(value)
}

// SetMaxChunkSize sets the max_chunk_size field.
/*
  Max chunk size in bytes
  Note, not all formats support this and unpredictable things may happen if it is used when not supported.
  - encoding: Set by user
  - decoding: unused
*/
func (s *AVFormatContext) SetMaxChunkSize(value int) {
	s.ptr.max_chunk_size = (C.int)(value)
}

// UseWallclockAsTimestamps gets the use_wallclock_as_timestamps field.
/*
  forces the use of wallclock timestamps as pts/dts of packets
  This has undefined results in the presence of B frames.
  - encoding: unused
  - decoding: Set by user
*/
func (s *AVFormatContext) UseWallclockAsTimestamps() int {
	value := s.ptr.use_wallclock_as_timestamps
	return int(value)
}

// SetUseWallclockAsTimestamps sets the use_wallclock_as_timestamps field.
/*
  forces the use of wallclock timestamps as pts/dts of packets
  This has undefined results in the presence of B frames.
  - encoding: unused
  - decoding: Set by user
*/
func (s *AVFormatContext) SetUseWallclockAsTimestamps(value int) {
	s.ptr.use_wallclock_as_timestamps = (C.int)(value)
}

// AvioFlags gets the avio_flags field.
/*
  avio flags, used to force AVIO_FLAG_DIRECT.
  - encoding: unused
  - decoding: Set by user
*/
func (s *AVFormatContext) AvioFlags() int {
	value := s.ptr.avio_flags
	return int(value)
}

// SetAvioFlags sets the avio_flags field.
/*
  avio flags, used to force AVIO_FLAG_DIRECT.
  - encoding: unused
  - decoding: Set by user
*/
func (s *AVFormatContext) SetAvioFlags(value int) {
	s.ptr.avio_flags = (C.int)(value)
}

// DurationEstimationMethod gets the duration_estimation_method field.
/*
  The duration field can be estimated through various ways, and this field can be used
  to know how the duration was estimated.
  - encoding: unused
  - decoding: Read by user
*/
func (s *AVFormatContext) DurationEstimationMethod() AVDurationEstimationMethod {
	value := s.ptr.duration_estimation_method
	return AVDurationEstimationMethod(value)
}

// SetDurationEstimationMethod sets the duration_estimation_method field.
/*
  The duration field can be estimated through various ways, and this field can be used
  to know how the duration was estimated.
  - encoding: unused
  - decoding: Read by user
*/
func (s *AVFormatContext) SetDurationEstimationMethod(value AVDurationEstimationMethod) {
	s.ptr.duration_estimation_method = (C.enum_AVDurationEstimationMethod)(value)
}

// SkipInitialBytes gets the skip_initial_bytes field.
/*
  Skip initial bytes when opening stream
  - encoding: unused
  - decoding: Set by user
*/
func (s *AVFormatContext) SkipInitialBytes() int64 {
	value := s.ptr.skip_initial_bytes
	return int64(value)
}

// SetSkipInitialBytes sets the skip_initial_bytes field.
/*
  Skip initial bytes when opening stream
  - encoding: unused
  - decoding: Set by user
*/
func (s *AVFormatContext) SetSkipInitialBytes(value int64) {
	s.ptr.skip_initial_bytes = (C.int64_t)(value)
}

// CorrectTsOverflow gets the correct_ts_overflow field.
/*
  Correct single timestamp overflows
  - encoding: unused
  - decoding: Set by user
*/
func (s *AVFormatContext) CorrectTsOverflow() uint {
	value := s.ptr.correct_ts_overflow
	return uint(value)
}

// SetCorrectTsOverflow sets the correct_ts_overflow field.
/*
  Correct single timestamp overflows
  - encoding: unused
  - decoding: Set by user
*/
func (s *AVFormatContext) SetCorrectTsOverflow(value uint) {
	s.ptr.correct_ts_overflow = (C.uint)(value)
}

// Seek2Any gets the seek2any field.
/*
  Force seeking to any (also non key) frames.
  - encoding: unused
  - decoding: Set by user
*/
func (s *AVFormatContext) Seek2Any() int {
	value := s.ptr.seek2any
	return int(value)
}

// SetSeek2Any sets the seek2any field.
/*
  Force seeking to any (also non key) frames.
  - encoding: unused
  - decoding: Set by user
*/
func (s *AVFormatContext) SetSeek2Any(value int) {
	s.ptr.seek2any = (C.int)(value)
}

// FlushPackets gets the flush_packets field.
/*
  Flush the I/O context after each packet.
  - encoding: Set by user
  - decoding: unused
*/
func (s *AVFormatContext) FlushPackets() int {
	value := s.ptr.flush_packets
	return int(value)
}

// SetFlushPackets sets the flush_packets field.
/*
  Flush the I/O context after each packet.
  - encoding: Set by user
  - decoding: unused
*/
func (s *AVFormatContext) SetFlushPackets(value int) {
	s.ptr.flush_packets = (C.int)(value)
}

// ProbeScore gets the probe_score field.
/*
  format probing score.
  The maximal score is AVPROBE_SCORE_MAX, its set when the demuxer probes
  the format.
  - encoding: unused
  - decoding: set by avformat, read by user
*/
func (s *AVFormatContext) ProbeScore() int {
	value := s.ptr.probe_score
	return int(value)
}

// SetProbeScore sets the probe_score field.
/*
  format probing score.
  The maximal score is AVPROBE_SCORE_MAX, its set when the demuxer probes
  the format.
  - encoding: unused
  - decoding: set by avformat, read by user
*/
func (s *AVFormatContext) SetProbeScore(value int) {
	s.ptr.probe_score = (C.int)(value)
}

// FormatProbesize gets the format_probesize field.
/*
  Maximum number of bytes read from input in order to identify the
  \ref AVInputFormat "input format". Only used when the format is not set
  explicitly by the caller.

  Demuxing only, set by the caller before avformat_open_input().

  @sa probesize
*/
func (s *AVFormatContext) FormatProbesize() int {
	value := s.ptr.format_probesize
	return int(value)
}

// SetFormatProbesize sets the format_probesize field.
/*
  Maximum number of bytes read from input in order to identify the
  \ref AVInputFormat "input format". Only used when the format is not set
  explicitly by the caller.

  Demuxing only, set by the caller before avformat_open_input().

  @sa probesize
*/
func (s *AVFormatContext) SetFormatProbesize(value int) {
	s.ptr.format_probesize = (C.int)(value)
}

// CodecWhitelist gets the codec_whitelist field.
/*
  ',' separated list of allowed decoders.
  If NULL then all are allowed
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) CodecWhitelist() *CStr {
	value := s.ptr.codec_whitelist
	return wrapCStr(value)
}

// SetCodecWhitelist sets the codec_whitelist field.
/*
  ',' separated list of allowed decoders.
  If NULL then all are allowed
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) SetCodecWhitelist(value *CStr) {
	s.ptr.codec_whitelist = value.ptr
}

// FormatWhitelist gets the format_whitelist field.
/*
  ',' separated list of allowed demuxers.
  If NULL then all are allowed
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) FormatWhitelist() *CStr {
	value := s.ptr.format_whitelist
	return wrapCStr(value)
}

// SetFormatWhitelist sets the format_whitelist field.
/*
  ',' separated list of allowed demuxers.
  If NULL then all are allowed
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) SetFormatWhitelist(value *CStr) {
	s.ptr.format_whitelist = value.ptr
}

// IoRepositioned gets the io_repositioned field.
/*
  IO repositioned flag.
  This is set by avformat when the underlaying IO context read pointer
  is repositioned, for example when doing byte based seeking.
  Demuxers can use the flag to detect such changes.
*/
func (s *AVFormatContext) IoRepositioned() int {
	value := s.ptr.io_repositioned
	return int(value)
}

// SetIoRepositioned sets the io_repositioned field.
/*
  IO repositioned flag.
  This is set by avformat when the underlaying IO context read pointer
  is repositioned, for example when doing byte based seeking.
  Demuxers can use the flag to detect such changes.
*/
func (s *AVFormatContext) SetIoRepositioned(value int) {
	s.ptr.io_repositioned = (C.int)(value)
}

// VideoCodec gets the video_codec field.
/*
  Forced video codec.
  This allows forcing a specific decoder, even when there are multiple with
  the same codec_id.
  Demuxing: Set by user
*/
func (s *AVFormatContext) VideoCodec() *AVCodec {
	value := s.ptr.video_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

// SetVideoCodec sets the video_codec field.
/*
  Forced video codec.
  This allows forcing a specific decoder, even when there are multiple with
  the same codec_id.
  Demuxing: Set by user
*/
func (s *AVFormatContext) SetVideoCodec(value *AVCodec) {
	if value != nil {
		s.ptr.video_codec = value.ptr
	} else {
		s.ptr.video_codec = nil
	}
}

// AudioCodec gets the audio_codec field.
/*
  Forced audio codec.
  This allows forcing a specific decoder, even when there are multiple with
  the same codec_id.
  Demuxing: Set by user
*/
func (s *AVFormatContext) AudioCodec() *AVCodec {
	value := s.ptr.audio_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

// SetAudioCodec sets the audio_codec field.
/*
  Forced audio codec.
  This allows forcing a specific decoder, even when there are multiple with
  the same codec_id.
  Demuxing: Set by user
*/
func (s *AVFormatContext) SetAudioCodec(value *AVCodec) {
	if value != nil {
		s.ptr.audio_codec = value.ptr
	} else {
		s.ptr.audio_codec = nil
	}
}

// SubtitleCodec gets the subtitle_codec field.
/*
  Forced subtitle codec.
  This allows forcing a specific decoder, even when there are multiple with
  the same codec_id.
  Demuxing: Set by user
*/
func (s *AVFormatContext) SubtitleCodec() *AVCodec {
	value := s.ptr.subtitle_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

// SetSubtitleCodec sets the subtitle_codec field.
/*
  Forced subtitle codec.
  This allows forcing a specific decoder, even when there are multiple with
  the same codec_id.
  Demuxing: Set by user
*/
func (s *AVFormatContext) SetSubtitleCodec(value *AVCodec) {
	if value != nil {
		s.ptr.subtitle_codec = value.ptr
	} else {
		s.ptr.subtitle_codec = nil
	}
}

// DataCodec gets the data_codec field.
/*
  Forced data codec.
  This allows forcing a specific decoder, even when there are multiple with
  the same codec_id.
  Demuxing: Set by user
*/
func (s *AVFormatContext) DataCodec() *AVCodec {
	value := s.ptr.data_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

// SetDataCodec sets the data_codec field.
/*
  Forced data codec.
  This allows forcing a specific decoder, even when there are multiple with
  the same codec_id.
  Demuxing: Set by user
*/
func (s *AVFormatContext) SetDataCodec(value *AVCodec) {
	if value != nil {
		s.ptr.data_codec = value.ptr
	} else {
		s.ptr.data_codec = nil
	}
}

// MetadataHeaderPadding gets the metadata_header_padding field.
/*
  Number of bytes to be written as padding in a metadata header.
  Demuxing: Unused.
  Muxing: Set by user via av_format_set_metadata_header_padding.
*/
func (s *AVFormatContext) MetadataHeaderPadding() int {
	value := s.ptr.metadata_header_padding
	return int(value)
}

// SetMetadataHeaderPadding sets the metadata_header_padding field.
/*
  Number of bytes to be written as padding in a metadata header.
  Demuxing: Unused.
  Muxing: Set by user via av_format_set_metadata_header_padding.
*/
func (s *AVFormatContext) SetMetadataHeaderPadding(value int) {
	s.ptr.metadata_header_padding = (C.int)(value)
}

// Opaque gets the opaque field.
/*
  User data.
  This is a place for some private data of the user.
*/
func (s *AVFormatContext) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
/*
  User data.
  This is a place for some private data of the user.
*/
func (s *AVFormatContext) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// control_message_cb skipped due to ident callback

// OutputTsOffset gets the output_ts_offset field.
/*
  Output timestamp offset, in microseconds.
  Muxing: set by user
*/
func (s *AVFormatContext) OutputTsOffset() int64 {
	value := s.ptr.output_ts_offset
	return int64(value)
}

// SetOutputTsOffset sets the output_ts_offset field.
/*
  Output timestamp offset, in microseconds.
  Muxing: set by user
*/
func (s *AVFormatContext) SetOutputTsOffset(value int64) {
	s.ptr.output_ts_offset = (C.int64_t)(value)
}

// DumpSeparator gets the dump_separator field.
/*
  dump format separator.
  can be ", " or "\n      " or anything else
  - muxing: Set by user.
  - demuxing: Set by user.
*/
func (s *AVFormatContext) DumpSeparator() unsafe.Pointer {
	value := s.ptr.dump_separator
	return unsafe.Pointer(value)
}

// SetDumpSeparator sets the dump_separator field.
/*
  dump format separator.
  can be ", " or "\n      " or anything else
  - muxing: Set by user.
  - demuxing: Set by user.
*/
func (s *AVFormatContext) SetDumpSeparator(value unsafe.Pointer) {
	s.ptr.dump_separator = (*C.uint8_t)(value)
}

// DataCodecId gets the data_codec_id field.
/*
  Forced Data codec_id.
  Demuxing: Set by user.
*/
func (s *AVFormatContext) DataCodecId() AVCodecID {
	value := s.ptr.data_codec_id
	return AVCodecID(value)
}

// SetDataCodecId sets the data_codec_id field.
/*
  Forced Data codec_id.
  Demuxing: Set by user.
*/
func (s *AVFormatContext) SetDataCodecId(value AVCodecID) {
	s.ptr.data_codec_id = (C.enum_AVCodecID)(value)
}

// ProtocolWhitelist gets the protocol_whitelist field.
/*
  ',' separated list of allowed protocols.
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) ProtocolWhitelist() *CStr {
	value := s.ptr.protocol_whitelist
	return wrapCStr(value)
}

// SetProtocolWhitelist sets the protocol_whitelist field.
/*
  ',' separated list of allowed protocols.
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) SetProtocolWhitelist(value *CStr) {
	s.ptr.protocol_whitelist = value.ptr
}

// io_open skipped due to func ptr

// io_close skipped due to func ptr

// ProtocolBlacklist gets the protocol_blacklist field.
/*
  ',' separated list of disallowed protocols.
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) ProtocolBlacklist() *CStr {
	value := s.ptr.protocol_blacklist
	return wrapCStr(value)
}

// SetProtocolBlacklist sets the protocol_blacklist field.
/*
  ',' separated list of disallowed protocols.
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) SetProtocolBlacklist(value *CStr) {
	s.ptr.protocol_blacklist = value.ptr
}

// MaxStreams gets the max_streams field.
/*
  The maximum number of streams.
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) MaxStreams() int {
	value := s.ptr.max_streams
	return int(value)
}

// SetMaxStreams sets the max_streams field.
/*
  The maximum number of streams.
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) SetMaxStreams(value int) {
	s.ptr.max_streams = (C.int)(value)
}

// SkipEstimateDurationFromPts gets the skip_estimate_duration_from_pts field.
/*
  Skip duration calcuation in estimate_timings_from_pts.
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) SkipEstimateDurationFromPts() int {
	value := s.ptr.skip_estimate_duration_from_pts
	return int(value)
}

// SetSkipEstimateDurationFromPts sets the skip_estimate_duration_from_pts field.
/*
  Skip duration calcuation in estimate_timings_from_pts.
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) SetSkipEstimateDurationFromPts(value int) {
	s.ptr.skip_estimate_duration_from_pts = (C.int)(value)
}

// MaxProbePackets gets the max_probe_packets field.
/*
  Maximum number of packets that can be probed
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) MaxProbePackets() int {
	value := s.ptr.max_probe_packets
	return int(value)
}

// SetMaxProbePackets sets the max_probe_packets field.
/*
  Maximum number of packets that can be probed
  - encoding: unused
  - decoding: set by user
*/
func (s *AVFormatContext) SetMaxProbePackets(value int) {
	s.ptr.max_probe_packets = (C.int)(value)
}

// io_close2 skipped due to func ptr

// --- Struct AVIOInterruptCB ---

// AVIOInterruptCB wraps AVIOInterruptCB.
/*
  Callback for checking whether to abort blocking functions.
  AVERROR_EXIT is returned in this case by the interrupted
  function. During blocking operations, callback is called with
  opaque as parameter. If the callback returns 1, the
  blocking operation will be aborted.

  No members can be added to this struct without a major bump, if
  new elements have been added after this struct in AVFormatContext
  or AVIOContext.
*/
type AVIOInterruptCB struct {
	ptr *C.AVIOInterruptCB
}

func (s *AVIOInterruptCB) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVIOInterruptCBArray(ptr unsafe.Pointer) *Array[*AVIOInterruptCB] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVIOInterruptCB]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVIOInterruptCB {
			ptr := (**C.AVIOInterruptCB)(pointer)
			value := *ptr
			var valueMapped *AVIOInterruptCB
			if value != nil {
				valueMapped = &AVIOInterruptCB{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVIOInterruptCB) {
			ptr := (**C.AVIOInterruptCB)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// callback skipped due to func ptr

// Opaque gets the opaque field.
func (s *AVIOInterruptCB) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
func (s *AVIOInterruptCB) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// --- Struct AVIODirEntry ---

// AVIODirEntry wraps AVIODirEntry.
/*
  Describes single entry of the directory.

  Only name and type fields are guaranteed be set.
  Rest of fields are protocol or/and platform dependent and might be unknown.
*/
type AVIODirEntry struct {
	ptr *C.AVIODirEntry
}

func (s *AVIODirEntry) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVIODirEntryArray(ptr unsafe.Pointer) *Array[*AVIODirEntry] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVIODirEntry]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVIODirEntry {
			ptr := (**C.AVIODirEntry)(pointer)
			value := *ptr
			var valueMapped *AVIODirEntry
			if value != nil {
				valueMapped = &AVIODirEntry{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVIODirEntry) {
			ptr := (**C.AVIODirEntry)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Name gets the name field.
//
//	Filename
func (s *AVIODirEntry) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
//
//	Filename
func (s *AVIODirEntry) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// Type gets the type field.
//
//	Type of the entry
func (s *AVIODirEntry) Type() int {
	value := s.ptr._type
	return int(value)
}

// SetType sets the type field.
//
//	Type of the entry
func (s *AVIODirEntry) SetType(value int) {
	s.ptr._type = (C.int)(value)
}

// Utf8 gets the utf8 field.
/*
  < Set to 1 when name is encoded with UTF-8, 0 otherwise.
  Name can be encoded with UTF-8 even though 0 is set.
*/
func (s *AVIODirEntry) Utf8() int {
	value := s.ptr.utf8
	return int(value)
}

// SetUtf8 sets the utf8 field.
/*
  < Set to 1 when name is encoded with UTF-8, 0 otherwise.
  Name can be encoded with UTF-8 even though 0 is set.
*/
func (s *AVIODirEntry) SetUtf8(value int) {
	s.ptr.utf8 = (C.int)(value)
}

// Size gets the size field.
//
//	File size in bytes, -1 if unknown.
func (s *AVIODirEntry) Size() int64 {
	value := s.ptr.size
	return int64(value)
}

// SetSize sets the size field.
//
//	File size in bytes, -1 if unknown.
func (s *AVIODirEntry) SetSize(value int64) {
	s.ptr.size = (C.int64_t)(value)
}

// ModificationTimestamp gets the modification_timestamp field.
/*
  < Time of last modification in microseconds since unix
  epoch, -1 if unknown.
*/
func (s *AVIODirEntry) ModificationTimestamp() int64 {
	value := s.ptr.modification_timestamp
	return int64(value)
}

// SetModificationTimestamp sets the modification_timestamp field.
/*
  < Time of last modification in microseconds since unix
  epoch, -1 if unknown.
*/
func (s *AVIODirEntry) SetModificationTimestamp(value int64) {
	s.ptr.modification_timestamp = (C.int64_t)(value)
}

// AccessTimestamp gets the access_timestamp field.
/*
  < Time of last access in microseconds since unix epoch,
  -1 if unknown.
*/
func (s *AVIODirEntry) AccessTimestamp() int64 {
	value := s.ptr.access_timestamp
	return int64(value)
}

// SetAccessTimestamp sets the access_timestamp field.
/*
  < Time of last access in microseconds since unix epoch,
  -1 if unknown.
*/
func (s *AVIODirEntry) SetAccessTimestamp(value int64) {
	s.ptr.access_timestamp = (C.int64_t)(value)
}

// StatusChangeTimestamp gets the status_change_timestamp field.
/*
  < Time of last status change in microseconds since unix
  epoch, -1 if unknown.
*/
func (s *AVIODirEntry) StatusChangeTimestamp() int64 {
	value := s.ptr.status_change_timestamp
	return int64(value)
}

// SetStatusChangeTimestamp sets the status_change_timestamp field.
/*
  < Time of last status change in microseconds since unix
  epoch, -1 if unknown.
*/
func (s *AVIODirEntry) SetStatusChangeTimestamp(value int64) {
	s.ptr.status_change_timestamp = (C.int64_t)(value)
}

// UserId gets the user_id field.
//
//	User ID of owner, -1 if unknown.
func (s *AVIODirEntry) UserId() int64 {
	value := s.ptr.user_id
	return int64(value)
}

// SetUserId sets the user_id field.
//
//	User ID of owner, -1 if unknown.
func (s *AVIODirEntry) SetUserId(value int64) {
	s.ptr.user_id = (C.int64_t)(value)
}

// GroupId gets the group_id field.
//
//	Group ID of owner, -1 if unknown.
func (s *AVIODirEntry) GroupId() int64 {
	value := s.ptr.group_id
	return int64(value)
}

// SetGroupId sets the group_id field.
//
//	Group ID of owner, -1 if unknown.
func (s *AVIODirEntry) SetGroupId(value int64) {
	s.ptr.group_id = (C.int64_t)(value)
}

// Filemode gets the filemode field.
//
//	Unix file mode, -1 if unknown.
func (s *AVIODirEntry) Filemode() int64 {
	value := s.ptr.filemode
	return int64(value)
}

// SetFilemode sets the filemode field.
//
//	Unix file mode, -1 if unknown.
func (s *AVIODirEntry) SetFilemode(value int64) {
	s.ptr.filemode = (C.int64_t)(value)
}

// --- Struct AVIODirContext ---

// AVIODirContext wraps AVIODirContext.
type AVIODirContext struct {
	ptr *C.AVIODirContext
}

func (s *AVIODirContext) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVIODirContextArray(ptr unsafe.Pointer) *Array[*AVIODirContext] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVIODirContext]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVIODirContext {
			ptr := (**C.AVIODirContext)(pointer)
			value := *ptr
			var valueMapped *AVIODirContext
			if value != nil {
				valueMapped = &AVIODirContext{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVIODirContext) {
			ptr := (**C.AVIODirContext)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// url_context skipped due to ptr to ignored type

// --- Struct AVIOContext ---

// AVIOContext wraps AVIOContext.
/*
  Bytestream IO Context.
  New public fields can be added with minor version bumps.
  Removal, reordering and changes to existing public fields require
  a major version bump.
  sizeof(AVIOContext) must not be used outside libav*.

  @note None of the function pointers in AVIOContext should be called
        directly, they should only be set by the client application
        when implementing custom I/O. Normally these are set to the
        function pointers specified in avio_alloc_context()
*/
type AVIOContext struct {
	ptr *C.AVIOContext
}

func (s *AVIOContext) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVIOContextArray(ptr unsafe.Pointer) *Array[*AVIOContext] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVIOContext]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVIOContext {
			ptr := (**C.AVIOContext)(pointer)
			value := *ptr
			var valueMapped *AVIOContext
			if value != nil {
				valueMapped = &AVIOContext{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVIOContext) {
			ptr := (**C.AVIOContext)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// AvClass gets the av_class field.
/*
  A class for private options.

  If this AVIOContext is created by avio_open2(), av_class is set and
  passes the options down to protocols.

  If this AVIOContext is manually allocated, then av_class may be set by
  the caller.

  warning -- this field can be NULL, be sure to not pass this AVIOContext
  to any av_opt_* functions in that case.
*/
func (s *AVIOContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
/*
  A class for private options.

  If this AVIOContext is created by avio_open2(), av_class is set and
  passes the options down to protocols.

  If this AVIOContext is manually allocated, then av_class may be set by
  the caller.

  warning -- this field can be NULL, be sure to not pass this AVIOContext
  to any av_opt_* functions in that case.
*/
func (s *AVIOContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// buffer skipped due to prim ptr

// BufferSize gets the buffer_size field.
//
//	Maximum buffer size
func (s *AVIOContext) BufferSize() int {
	value := s.ptr.buffer_size
	return int(value)
}

// SetBufferSize sets the buffer_size field.
//
//	Maximum buffer size
func (s *AVIOContext) SetBufferSize(value int) {
	s.ptr.buffer_size = (C.int)(value)
}

// buf_ptr skipped due to prim ptr

// buf_end skipped due to prim ptr

// Opaque gets the opaque field.
/*
  < A private pointer, passed to the read/write/seek/...
  functions.
*/
func (s *AVIOContext) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
/*
  < A private pointer, passed to the read/write/seek/...
  functions.
*/
func (s *AVIOContext) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// read_packet skipped due to func ptr

// write_packet skipped due to func ptr

// seek skipped due to func ptr

// Pos gets the pos field.
//
//	position in the file of the current buffer
func (s *AVIOContext) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}

// SetPos sets the pos field.
//
//	position in the file of the current buffer
func (s *AVIOContext) SetPos(value int64) {
	s.ptr.pos = (C.int64_t)(value)
}

// EofReached gets the eof_reached field.
//
//	true if was unable to read due to error or eof
func (s *AVIOContext) EofReached() int {
	value := s.ptr.eof_reached
	return int(value)
}

// SetEofReached sets the eof_reached field.
//
//	true if was unable to read due to error or eof
func (s *AVIOContext) SetEofReached(value int) {
	s.ptr.eof_reached = (C.int)(value)
}

// Error gets the error field.
//
//	contains the error code or 0 if no error happened
func (s *AVIOContext) Error() int {
	value := s.ptr.error
	return int(value)
}

// SetError sets the error field.
//
//	contains the error code or 0 if no error happened
func (s *AVIOContext) SetError(value int) {
	s.ptr.error = (C.int)(value)
}

// WriteFlag gets the write_flag field.
//
//	true if open for writing
func (s *AVIOContext) WriteFlag() int {
	value := s.ptr.write_flag
	return int(value)
}

// SetWriteFlag sets the write_flag field.
//
//	true if open for writing
func (s *AVIOContext) SetWriteFlag(value int) {
	s.ptr.write_flag = (C.int)(value)
}

// MaxPacketSize gets the max_packet_size field.
func (s *AVIOContext) MaxPacketSize() int {
	value := s.ptr.max_packet_size
	return int(value)
}

// SetMaxPacketSize sets the max_packet_size field.
func (s *AVIOContext) SetMaxPacketSize(value int) {
	s.ptr.max_packet_size = (C.int)(value)
}

// MinPacketSize gets the min_packet_size field.
/*
  < Try to buffer at least this amount of data
  before flushing it.
*/
func (s *AVIOContext) MinPacketSize() int {
	value := s.ptr.min_packet_size
	return int(value)
}

// SetMinPacketSize sets the min_packet_size field.
/*
  < Try to buffer at least this amount of data
  before flushing it.
*/
func (s *AVIOContext) SetMinPacketSize(value int) {
	s.ptr.min_packet_size = (C.int)(value)
}

// Checksum gets the checksum field.
func (s *AVIOContext) Checksum() uint32 {
	value := s.ptr.checksum
	return uint32(value)
}

// SetChecksum sets the checksum field.
func (s *AVIOContext) SetChecksum(value uint32) {
	s.ptr.checksum = (C.ulong)(value)
}

// checksum_ptr skipped due to prim ptr

// update_checksum skipped due to func ptr

// read_pause skipped due to func ptr

// read_seek skipped due to func ptr

// Seekable gets the seekable field.
//
//	A combination of AVIO_SEEKABLE_ flags or 0 when the stream is not seekable.
func (s *AVIOContext) Seekable() int {
	value := s.ptr.seekable
	return int(value)
}

// SetSeekable sets the seekable field.
//
//	A combination of AVIO_SEEKABLE_ flags or 0 when the stream is not seekable.
func (s *AVIOContext) SetSeekable(value int) {
	s.ptr.seekable = (C.int)(value)
}

// Direct gets the direct field.
/*
  avio_read and avio_write should if possible be satisfied directly
  instead of going through a buffer, and avio_seek will always
  call the underlying seek function directly.
*/
func (s *AVIOContext) Direct() int {
	value := s.ptr.direct
	return int(value)
}

// SetDirect sets the direct field.
/*
  avio_read and avio_write should if possible be satisfied directly
  instead of going through a buffer, and avio_seek will always
  call the underlying seek function directly.
*/
func (s *AVIOContext) SetDirect(value int) {
	s.ptr.direct = (C.int)(value)
}

// ProtocolWhitelist gets the protocol_whitelist field.
//
//	',' separated list of allowed protocols.
func (s *AVIOContext) ProtocolWhitelist() *CStr {
	value := s.ptr.protocol_whitelist
	return wrapCStr(value)
}

// SetProtocolWhitelist sets the protocol_whitelist field.
//
//	',' separated list of allowed protocols.
func (s *AVIOContext) SetProtocolWhitelist(value *CStr) {
	s.ptr.protocol_whitelist = value.ptr
}

// ProtocolBlacklist gets the protocol_blacklist field.
//
//	',' separated list of disallowed protocols.
func (s *AVIOContext) ProtocolBlacklist() *CStr {
	value := s.ptr.protocol_blacklist
	return wrapCStr(value)
}

// SetProtocolBlacklist sets the protocol_blacklist field.
//
//	',' separated list of disallowed protocols.
func (s *AVIOContext) SetProtocolBlacklist(value *CStr) {
	s.ptr.protocol_blacklist = value.ptr
}

// write_data_type skipped due to func ptr

// IgnoreBoundaryPoint gets the ignore_boundary_point field.
/*
  If set, don't call write_data_type separately for AVIO_DATA_MARKER_BOUNDARY_POINT,
  but ignore them and treat them as AVIO_DATA_MARKER_UNKNOWN (to avoid needlessly
  small chunks of data returned from the callback).
*/
func (s *AVIOContext) IgnoreBoundaryPoint() int {
	value := s.ptr.ignore_boundary_point
	return int(value)
}

// SetIgnoreBoundaryPoint sets the ignore_boundary_point field.
/*
  If set, don't call write_data_type separately for AVIO_DATA_MARKER_BOUNDARY_POINT,
  but ignore them and treat them as AVIO_DATA_MARKER_UNKNOWN (to avoid needlessly
  small chunks of data returned from the callback).
*/
func (s *AVIOContext) SetIgnoreBoundaryPoint(value int) {
	s.ptr.ignore_boundary_point = (C.int)(value)
}

// buf_ptr_max skipped due to prim ptr

// BytesRead gets the bytes_read field.
//
//	Read-only statistic of bytes read for this AVIOContext.
func (s *AVIOContext) BytesRead() int64 {
	value := s.ptr.bytes_read
	return int64(value)
}

// SetBytesRead sets the bytes_read field.
//
//	Read-only statistic of bytes read for this AVIOContext.
func (s *AVIOContext) SetBytesRead(value int64) {
	s.ptr.bytes_read = (C.int64_t)(value)
}

// BytesWritten gets the bytes_written field.
//
//	Read-only statistic of bytes written for this AVIOContext.
func (s *AVIOContext) BytesWritten() int64 {
	value := s.ptr.bytes_written
	return int64(value)
}

// SetBytesWritten sets the bytes_written field.
//
//	Read-only statistic of bytes written for this AVIOContext.
func (s *AVIOContext) SetBytesWritten(value int64) {
	s.ptr.bytes_written = (C.int64_t)(value)
}

// --- Struct AVBuffer ---

// AVBuffer wraps AVBuffer.
type AVBuffer struct {
	ptr *C.AVBuffer
}

func (s *AVBuffer) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVBufferArray(ptr unsafe.Pointer) *Array[*AVBuffer] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVBuffer]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVBuffer {
			ptr := (**C.AVBuffer)(pointer)
			value := *ptr
			var valueMapped *AVBuffer
			if value != nil {
				valueMapped = &AVBuffer{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVBuffer) {
			ptr := (**C.AVBuffer)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVBufferRef ---

// AVBufferRef wraps AVBufferRef.
/*
  A reference to a data buffer.

  The size of this struct is not a part of the public ABI and it is not meant
  to be allocated directly.
*/
type AVBufferRef struct {
	ptr *C.AVBufferRef
}

func (s *AVBufferRef) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVBufferRefArray(ptr unsafe.Pointer) *Array[*AVBufferRef] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVBufferRef]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVBufferRef {
			ptr := (**C.AVBufferRef)(pointer)
			value := *ptr
			var valueMapped *AVBufferRef
			if value != nil {
				valueMapped = &AVBufferRef{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVBufferRef) {
			ptr := (**C.AVBufferRef)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Buffer gets the buffer field.
func (s *AVBufferRef) Buffer() *AVBuffer {
	value := s.ptr.buffer
	var valueMapped *AVBuffer
	if value != nil {
		valueMapped = &AVBuffer{ptr: value}
	}
	return valueMapped
}

// SetBuffer sets the buffer field.
func (s *AVBufferRef) SetBuffer(value *AVBuffer) {
	if value != nil {
		s.ptr.buffer = value.ptr
	} else {
		s.ptr.buffer = nil
	}
}

// Data gets the data field.
/*
  The data buffer. It is considered writable if and only if
  this is the only reference to the buffer, in which case
  av_buffer_is_writable() returns 1.
*/
func (s *AVBufferRef) Data() unsafe.Pointer {
	value := s.ptr.data
	return unsafe.Pointer(value)
}

// SetData sets the data field.
/*
  The data buffer. It is considered writable if and only if
  this is the only reference to the buffer, in which case
  av_buffer_is_writable() returns 1.
*/
func (s *AVBufferRef) SetData(value unsafe.Pointer) {
	s.ptr.data = (*C.uint8_t)(value)
}

// Size gets the size field.
//
//	Size of data in bytes.
func (s *AVBufferRef) Size() uint64 {
	value := s.ptr.size
	return uint64(value)
}

// SetSize sets the size field.
//
//	Size of data in bytes.
func (s *AVBufferRef) SetSize(value uint64) {
	s.ptr.size = (C.size_t)(value)
}

// --- Struct AVBufferPool ---

// AVBufferPool wraps AVBufferPool.
type AVBufferPool struct {
	ptr *C.AVBufferPool
}

func (s *AVBufferPool) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVBufferPoolArray(ptr unsafe.Pointer) *Array[*AVBufferPool] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVBufferPool]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVBufferPool {
			ptr := (**C.AVBufferPool)(pointer)
			value := *ptr
			var valueMapped *AVBufferPool
			if value != nil {
				valueMapped = &AVBufferPool{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVBufferPool) {
			ptr := (**C.AVBufferPool)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVChannelCustom ---

// AVChannelCustom wraps AVChannelCustom.
/*
  An AVChannelCustom defines a single channel within a custom order layout

  Unlike most structures in FFmpeg, sizeof(AVChannelCustom) is a part of the
  public ABI.

  No new fields may be added to it without a major version bump.
*/
type AVChannelCustom struct {
	ptr *C.AVChannelCustom
}

func (s *AVChannelCustom) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVChannelCustomArray(ptr unsafe.Pointer) *Array[*AVChannelCustom] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVChannelCustom]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVChannelCustom {
			ptr := (**C.AVChannelCustom)(pointer)
			value := *ptr
			var valueMapped *AVChannelCustom
			if value != nil {
				valueMapped = &AVChannelCustom{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVChannelCustom) {
			ptr := (**C.AVChannelCustom)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Id gets the id field.
func (s *AVChannelCustom) Id() AVChannel {
	value := s.ptr.id
	return AVChannel(value)
}

// SetId sets the id field.
func (s *AVChannelCustom) SetId(value AVChannel) {
	s.ptr.id = (C.enum_AVChannel)(value)
}

// Name gets the name field.
func (s *AVChannelCustom) Name() *Array[uint8] {
	value := &s.ptr.name
	return ToUint8Array(unsafe.Pointer(value))
}

// Opaque gets the opaque field.
func (s *AVChannelCustom) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
func (s *AVChannelCustom) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// --- Struct AVChannelLayout ---

// AVChannelLayout wraps AVChannelLayout.
/*
  An AVChannelLayout holds information about the channel layout of audio data.

  A channel layout here is defined as a set of channels ordered in a specific
  way (unless the channel order is AV_CHANNEL_ORDER_UNSPEC, in which case an
  AVChannelLayout carries only the channel count).
  All orders may be treated as if they were AV_CHANNEL_ORDER_UNSPEC by
  ignoring everything but the channel count, as long as av_channel_layout_check()
  considers they are valid.

  Unlike most structures in FFmpeg, sizeof(AVChannelLayout) is a part of the
  public ABI and may be used by the caller. E.g. it may be allocated on stack
  or embedded in caller-defined structs.

  AVChannelLayout can be initialized as follows:
  - default initialization with {0}, followed by setting all used fields
    correctly;
  - by assigning one of the predefined AV_CHANNEL_LAYOUT_* initializers;
  - with a constructor function, such as av_channel_layout_default(),
    av_channel_layout_from_mask() or av_channel_layout_from_string().

  The channel layout must be unitialized with av_channel_layout_uninit()

  Copying an AVChannelLayout via assigning is forbidden,
  av_channel_layout_copy() must be used instead (and its return value should
  be checked)

  No new fields may be added to it without a major version bump, except for
  new elements of the union fitting in sizeof(uint64_t).
*/
type AVChannelLayout struct {
	ptr *C.AVChannelLayout
}

func (s *AVChannelLayout) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVChannelLayoutArray(ptr unsafe.Pointer) *Array[*AVChannelLayout] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVChannelLayout]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVChannelLayout {
			ptr := (**C.AVChannelLayout)(pointer)
			value := *ptr
			var valueMapped *AVChannelLayout
			if value != nil {
				valueMapped = &AVChannelLayout{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVChannelLayout) {
			ptr := (**C.AVChannelLayout)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Order gets the order field.
/*
  Channel order used in this layout.
  This is a mandatory field.
*/
func (s *AVChannelLayout) Order() AVChannelOrder {
	value := s.ptr.order
	return AVChannelOrder(value)
}

// SetOrder sets the order field.
/*
  Channel order used in this layout.
  This is a mandatory field.
*/
func (s *AVChannelLayout) SetOrder(value AVChannelOrder) {
	s.ptr.order = (C.enum_AVChannelOrder)(value)
}

// NbChannels gets the nb_channels field.
//
//	Number of channels in this layout. Mandatory field.
func (s *AVChannelLayout) NbChannels() int {
	value := s.ptr.nb_channels
	return int(value)
}

// SetNbChannels sets the nb_channels field.
//
//	Number of channels in this layout. Mandatory field.
func (s *AVChannelLayout) SetNbChannels(value int) {
	s.ptr.nb_channels = (C.int)(value)
}

// u skipped due to union type

// Opaque gets the opaque field.
//
//	For some private data of the user.
func (s *AVChannelLayout) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
//
//	For some private data of the user.
func (s *AVChannelLayout) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// --- Struct AVBPrint ---

// AVBPrint wraps AVBPrint.
type AVBPrint struct {
	ptr *C.struct_AVBPrint
}

func (s *AVBPrint) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVBPrintArray(ptr unsafe.Pointer) *Array[*AVBPrint] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVBPrint]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVBPrint {
			ptr := (**C.struct_AVBPrint)(pointer)
			value := *ptr
			var valueMapped *AVBPrint
			if value != nil {
				valueMapped = &AVBPrint{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVBPrint) {
			ptr := (**C.struct_AVBPrint)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVDictionaryEntry ---

// AVDictionaryEntry wraps AVDictionaryEntry.
type AVDictionaryEntry struct {
	ptr *C.AVDictionaryEntry
}

func (s *AVDictionaryEntry) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVDictionaryEntryArray(ptr unsafe.Pointer) *Array[*AVDictionaryEntry] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVDictionaryEntry]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVDictionaryEntry {
			ptr := (**C.AVDictionaryEntry)(pointer)
			value := *ptr
			var valueMapped *AVDictionaryEntry
			if value != nil {
				valueMapped = &AVDictionaryEntry{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVDictionaryEntry) {
			ptr := (**C.AVDictionaryEntry)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Key gets the key field.
func (s *AVDictionaryEntry) Key() *CStr {
	value := s.ptr.key
	return wrapCStr(value)
}

// SetKey sets the key field.
func (s *AVDictionaryEntry) SetKey(value *CStr) {
	s.ptr.key = value.ptr
}

// Value gets the value field.
func (s *AVDictionaryEntry) Value() *CStr {
	value := s.ptr.value
	return wrapCStr(value)
}

// SetValue sets the value field.
func (s *AVDictionaryEntry) SetValue(value *CStr) {
	s.ptr.value = value.ptr
}

// --- Struct AVDictionary ---

// AVDictionary wraps AVDictionary.
type AVDictionary struct {
	ptr *C.AVDictionary
}

func (s *AVDictionary) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVDictionaryArray(ptr unsafe.Pointer) *Array[*AVDictionary] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVDictionary]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVDictionary {
			ptr := (**C.AVDictionary)(pointer)
			value := *ptr
			var valueMapped *AVDictionary
			if value != nil {
				valueMapped = &AVDictionary{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVDictionary) {
			ptr := (**C.AVDictionary)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVFrameSideData ---

// AVFrameSideData wraps AVFrameSideData.
/*
  Structure to hold side data for an AVFrame.

  sizeof(AVFrameSideData) is not a part of the public ABI, so new fields may be added
  to the end with a minor bump.
*/
type AVFrameSideData struct {
	ptr *C.AVFrameSideData
}

func (s *AVFrameSideData) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFrameSideDataArray(ptr unsafe.Pointer) *Array[*AVFrameSideData] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFrameSideData]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFrameSideData {
			ptr := (**C.AVFrameSideData)(pointer)
			value := *ptr
			var valueMapped *AVFrameSideData
			if value != nil {
				valueMapped = &AVFrameSideData{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFrameSideData) {
			ptr := (**C.AVFrameSideData)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Type gets the type field.
func (s *AVFrameSideData) Type() AVFrameSideDataType {
	value := s.ptr._type
	return AVFrameSideDataType(value)
}

// SetType sets the type field.
func (s *AVFrameSideData) SetType(value AVFrameSideDataType) {
	s.ptr._type = (C.enum_AVFrameSideDataType)(value)
}

// Data gets the data field.
func (s *AVFrameSideData) Data() unsafe.Pointer {
	value := s.ptr.data
	return unsafe.Pointer(value)
}

// SetData sets the data field.
func (s *AVFrameSideData) SetData(value unsafe.Pointer) {
	s.ptr.data = (*C.uint8_t)(value)
}

// Size gets the size field.
func (s *AVFrameSideData) Size() uint64 {
	value := s.ptr.size
	return uint64(value)
}

// SetSize sets the size field.
func (s *AVFrameSideData) SetSize(value uint64) {
	s.ptr.size = (C.size_t)(value)
}

// Metadata gets the metadata field.
func (s *AVFrameSideData) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

// SetMetadata sets the metadata field.
func (s *AVFrameSideData) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

// Buf gets the buf field.
func (s *AVFrameSideData) Buf() *AVBufferRef {
	value := s.ptr.buf
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetBuf sets the buf field.
func (s *AVFrameSideData) SetBuf(value *AVBufferRef) {
	if value != nil {
		s.ptr.buf = value.ptr
	} else {
		s.ptr.buf = nil
	}
}

// --- Struct AVRegionOfInterest ---

// AVRegionOfInterest wraps AVRegionOfInterest.
/*
  Structure describing a single Region Of Interest.

  When multiple regions are defined in a single side-data block, they
  should be ordered from most to least important - some encoders are only
  capable of supporting a limited number of distinct regions, so will have
  to truncate the list.

  When overlapping regions are defined, the first region containing a given
  area of the frame applies.
*/
type AVRegionOfInterest struct {
	ptr *C.AVRegionOfInterest
}

func (s *AVRegionOfInterest) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVRegionOfInterestArray(ptr unsafe.Pointer) *Array[*AVRegionOfInterest] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVRegionOfInterest]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVRegionOfInterest {
			ptr := (**C.AVRegionOfInterest)(pointer)
			value := *ptr
			var valueMapped *AVRegionOfInterest
			if value != nil {
				valueMapped = &AVRegionOfInterest{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVRegionOfInterest) {
			ptr := (**C.AVRegionOfInterest)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// SelfSize gets the self_size field.
/*
  Must be set to the size of this data structure (that is,
  sizeof(AVRegionOfInterest)).
*/
func (s *AVRegionOfInterest) SelfSize() uint32 {
	value := s.ptr.self_size
	return uint32(value)
}

// SetSelfSize sets the self_size field.
/*
  Must be set to the size of this data structure (that is,
  sizeof(AVRegionOfInterest)).
*/
func (s *AVRegionOfInterest) SetSelfSize(value uint32) {
	s.ptr.self_size = (C.uint32_t)(value)
}

// Top gets the top field.
/*
  Distance in pixels from the top edge of the frame to the top and
  bottom edges and from the left edge of the frame to the left and
  right edges of the rectangle defining this region of interest.

  The constraints on a region are encoder dependent, so the region
  actually affected may be slightly larger for alignment or other
  reasons.
*/
func (s *AVRegionOfInterest) Top() int {
	value := s.ptr.top
	return int(value)
}

// SetTop sets the top field.
/*
  Distance in pixels from the top edge of the frame to the top and
  bottom edges and from the left edge of the frame to the left and
  right edges of the rectangle defining this region of interest.

  The constraints on a region are encoder dependent, so the region
  actually affected may be slightly larger for alignment or other
  reasons.
*/
func (s *AVRegionOfInterest) SetTop(value int) {
	s.ptr.top = (C.int)(value)
}

// Bottom gets the bottom field.
func (s *AVRegionOfInterest) Bottom() int {
	value := s.ptr.bottom
	return int(value)
}

// SetBottom sets the bottom field.
func (s *AVRegionOfInterest) SetBottom(value int) {
	s.ptr.bottom = (C.int)(value)
}

// Left gets the left field.
func (s *AVRegionOfInterest) Left() int {
	value := s.ptr.left
	return int(value)
}

// SetLeft sets the left field.
func (s *AVRegionOfInterest) SetLeft(value int) {
	s.ptr.left = (C.int)(value)
}

// Right gets the right field.
func (s *AVRegionOfInterest) Right() int {
	value := s.ptr.right
	return int(value)
}

// SetRight sets the right field.
func (s *AVRegionOfInterest) SetRight(value int) {
	s.ptr.right = (C.int)(value)
}

// Qoffset gets the qoffset field.
/*
  Quantisation offset.

  Must be in the range -1 to +1.  A value of zero indicates no quality
  change.  A negative value asks for better quality (less quantisation),
  while a positive value asks for worse quality (greater quantisation).

  The range is calibrated so that the extreme values indicate the
  largest possible offset - if the rest of the frame is encoded with the
  worst possible quality, an offset of -1 indicates that this region
  should be encoded with the best possible quality anyway.  Intermediate
  values are then interpolated in some codec-dependent way.

  For example, in 10-bit H.264 the quantisation parameter varies between
  -12 and 51.  A typical qoffset value of -1/10 therefore indicates that
  this region should be encoded with a QP around one-tenth of the full
  range better than the rest of the frame.  So, if most of the frame
  were to be encoded with a QP of around 30, this region would get a QP
  of around 24 (an offset of approximately -1/10 * (51 - -12) = -6.3).
  An extreme value of -1 would indicate that this region should be
  encoded with the best possible quality regardless of the treatment of
  the rest of the frame - that is, should be encoded at a QP of -12.
*/
func (s *AVRegionOfInterest) Qoffset() *AVRational {
	value := s.ptr.qoffset
	return &AVRational{value: value}
}

// SetQoffset sets the qoffset field.
/*
  Quantisation offset.

  Must be in the range -1 to +1.  A value of zero indicates no quality
  change.  A negative value asks for better quality (less quantisation),
  while a positive value asks for worse quality (greater quantisation).

  The range is calibrated so that the extreme values indicate the
  largest possible offset - if the rest of the frame is encoded with the
  worst possible quality, an offset of -1 indicates that this region
  should be encoded with the best possible quality anyway.  Intermediate
  values are then interpolated in some codec-dependent way.

  For example, in 10-bit H.264 the quantisation parameter varies between
  -12 and 51.  A typical qoffset value of -1/10 therefore indicates that
  this region should be encoded with a QP around one-tenth of the full
  range better than the rest of the frame.  So, if most of the frame
  were to be encoded with a QP of around 30, this region would get a QP
  of around 24 (an offset of approximately -1/10 * (51 - -12) = -6.3).
  An extreme value of -1 would indicate that this region should be
  encoded with the best possible quality regardless of the treatment of
  the rest of the frame - that is, should be encoded at a QP of -12.
*/
func (s *AVRegionOfInterest) SetQoffset(value *AVRational) {
	s.ptr.qoffset = value.value
}

// --- Struct AVFrame ---

// AVFrame wraps AVFrame.
/*
  This structure describes decoded (raw) audio or video data.

  AVFrame must be allocated using av_frame_alloc(). Note that this only
  allocates the AVFrame itself, the buffers for the data must be managed
  through other means (see below).
  AVFrame must be freed with av_frame_free().

  AVFrame is typically allocated once and then reused multiple times to hold
  different data (e.g. a single AVFrame to hold frames received from a
  decoder). In such a case, av_frame_unref() will free any references held by
  the frame and reset it to its original clean state before it
  is reused again.

  The data described by an AVFrame is usually reference counted through the
  AVBuffer API. The underlying buffer references are stored in AVFrame.buf /
  AVFrame.extended_buf. An AVFrame is considered to be reference counted if at
  least one reference is set, i.e. if AVFrame.buf[0] != NULL. In such a case,
  every single data plane must be contained in one of the buffers in
  AVFrame.buf or AVFrame.extended_buf.
  There may be a single buffer for all the data, or one separate buffer for
  each plane, or anything in between.

  sizeof(AVFrame) is not a part of the public ABI, so new fields may be added
  to the end with a minor bump.

  Fields can be accessed through AVOptions, the name string used, matches the
  C structure field name for fields accessible through AVOptions. The AVClass
  for AVFrame can be obtained from avcodec_get_frame_class()
*/
type AVFrame struct {
	ptr *C.AVFrame
}

func (s *AVFrame) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVFrameArray(ptr unsafe.Pointer) *Array[*AVFrame] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVFrame]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVFrame {
			ptr := (**C.AVFrame)(pointer)
			value := *ptr
			var valueMapped *AVFrame
			if value != nil {
				valueMapped = &AVFrame{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVFrame) {
			ptr := (**C.AVFrame)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Data gets the data field.
/*
  pointer to the picture/channel planes.
  This might be different from the first allocated byte. For video,
  it could even point to the end of the image data.

  All pointers in data and extended_data must point into one of the
  AVBufferRef in buf or extended_buf.

  Some decoders access areas outside 0,0 - width,height, please
  see avcodec_align_dimensions2(). Some filters and swscale can read
  up to 16 bytes beyond the planes, if these filters are to be used,
  then 16 extra bytes must be allocated.

  NOTE: Pointers not needed by the format MUST be set to NULL.

  @attention In case of video, the data[] pointers can point to the
  end of image data in order to reverse line order, when used in
  combination with negative values in the linesize[] array.
*/
func (s *AVFrame) Data() *Array[unsafe.Pointer] {
	value := &s.ptr.data
	return ToUint8PtrArray(unsafe.Pointer(value))
}

// Linesize gets the linesize field.
/*
  For video, a positive or negative value, which is typically indicating
  the size in bytes of each picture line, but it can also be:
  - the negative byte size of lines for vertical flipping
    (with data[n] pointing to the end of the data
  - a positive or negative multiple of the byte size as for accessing
    even and odd fields of a frame (possibly flipped)

  For audio, only linesize[0] may be set. For planar audio, each channel
  plane must be the same size.

  For video the linesizes should be multiples of the CPUs alignment
  preference, this is 16 or 32 for modern desktop CPUs.
  Some code requires such alignment other code can be slower without
  correct alignment, for yet other it makes no difference.

  @note The linesize may be larger than the size of usable data -- there
  may be extra padding present for performance reasons.

  @attention In case of video, line size values can be negative to achieve
  a vertically inverted iteration over image lines.
*/
func (s *AVFrame) Linesize() *Array[int] {
	value := &s.ptr.linesize
	return ToIntArray(unsafe.Pointer(value))
}

// extended_data skipped due to unknown ptr ptr

// Width gets the width field.
func (s *AVFrame) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
func (s *AVFrame) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

// Height gets the height field.
func (s *AVFrame) Height() int {
	value := s.ptr.height
	return int(value)
}

// SetHeight sets the height field.
func (s *AVFrame) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

// NbSamples gets the nb_samples field.
//
//	number of audio samples (per channel) described by this frame
func (s *AVFrame) NbSamples() int {
	value := s.ptr.nb_samples
	return int(value)
}

// SetNbSamples sets the nb_samples field.
//
//	number of audio samples (per channel) described by this frame
func (s *AVFrame) SetNbSamples(value int) {
	s.ptr.nb_samples = (C.int)(value)
}

// Format gets the format field.
/*
  format of the frame, -1 if unknown or unset
  Values correspond to enum AVPixelFormat for video frames,
  enum AVSampleFormat for audio)
*/
func (s *AVFrame) Format() int {
	value := s.ptr.format
	return int(value)
}

// SetFormat sets the format field.
/*
  format of the frame, -1 if unknown or unset
  Values correspond to enum AVPixelFormat for video frames,
  enum AVSampleFormat for audio)
*/
func (s *AVFrame) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

// KeyFrame gets the key_frame field.
/*
  1 -> keyframe, 0-> not

  @deprecated Use AV_FRAME_FLAG_KEY instead
*/
func (s *AVFrame) KeyFrame() int {
	value := s.ptr.key_frame
	return int(value)
}

// SetKeyFrame sets the key_frame field.
/*
  1 -> keyframe, 0-> not

  @deprecated Use AV_FRAME_FLAG_KEY instead
*/
func (s *AVFrame) SetKeyFrame(value int) {
	s.ptr.key_frame = (C.int)(value)
}

// PictType gets the pict_type field.
//
//	Picture type of the frame.
func (s *AVFrame) PictType() AVPictureType {
	value := s.ptr.pict_type
	return AVPictureType(value)
}

// SetPictType sets the pict_type field.
//
//	Picture type of the frame.
func (s *AVFrame) SetPictType(value AVPictureType) {
	s.ptr.pict_type = (C.enum_AVPictureType)(value)
}

// SampleAspectRatio gets the sample_aspect_ratio field.
//
//	Sample aspect ratio for the video frame, 0/1 if unknown/unspecified.
func (s *AVFrame) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
//
//	Sample aspect ratio for the video frame, 0/1 if unknown/unspecified.
func (s *AVFrame) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

// Pts gets the pts field.
//
//	Presentation timestamp in time_base units (time when frame should be shown to user).
func (s *AVFrame) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

// SetPts sets the pts field.
//
//	Presentation timestamp in time_base units (time when frame should be shown to user).
func (s *AVFrame) SetPts(value int64) {
	s.ptr.pts = (C.int64_t)(value)
}

// PktDts gets the pkt_dts field.
/*
  DTS copied from the AVPacket that triggered returning this frame. (if frame threading isn't used)
  This is also the Presentation time of this AVFrame calculated from
  only AVPacket.dts values without pts values.
*/
func (s *AVFrame) PktDts() int64 {
	value := s.ptr.pkt_dts
	return int64(value)
}

// SetPktDts sets the pkt_dts field.
/*
  DTS copied from the AVPacket that triggered returning this frame. (if frame threading isn't used)
  This is also the Presentation time of this AVFrame calculated from
  only AVPacket.dts values without pts values.
*/
func (s *AVFrame) SetPktDts(value int64) {
	s.ptr.pkt_dts = (C.int64_t)(value)
}

// TimeBase gets the time_base field.
/*
  Time base for the timestamps in this frame.
  In the future, this field may be set on frames output by decoders or
  filters, but its value will be by default ignored on input to encoders
  or filters.
*/
func (s *AVFrame) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
/*
  Time base for the timestamps in this frame.
  In the future, this field may be set on frames output by decoders or
  filters, but its value will be by default ignored on input to encoders
  or filters.
*/
func (s *AVFrame) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// CodedPictureNumber gets the coded_picture_number field.
//
//	picture number in bitstream order
func (s *AVFrame) CodedPictureNumber() int {
	value := s.ptr.coded_picture_number
	return int(value)
}

// SetCodedPictureNumber sets the coded_picture_number field.
//
//	picture number in bitstream order
func (s *AVFrame) SetCodedPictureNumber(value int) {
	s.ptr.coded_picture_number = (C.int)(value)
}

// DisplayPictureNumber gets the display_picture_number field.
//
//	picture number in display order
func (s *AVFrame) DisplayPictureNumber() int {
	value := s.ptr.display_picture_number
	return int(value)
}

// SetDisplayPictureNumber sets the display_picture_number field.
//
//	picture number in display order
func (s *AVFrame) SetDisplayPictureNumber(value int) {
	s.ptr.display_picture_number = (C.int)(value)
}

// Quality gets the quality field.
//
//	quality (between 1 (good) and FF_LAMBDA_MAX (bad))
func (s *AVFrame) Quality() int {
	value := s.ptr.quality
	return int(value)
}

// SetQuality sets the quality field.
//
//	quality (between 1 (good) and FF_LAMBDA_MAX (bad))
func (s *AVFrame) SetQuality(value int) {
	s.ptr.quality = (C.int)(value)
}

// Opaque gets the opaque field.
/*
  Frame owner's private data.

  This field may be set by the code that allocates/owns the frame data.
  It is then not touched by any library functions, except:
  - it is copied to other references by av_frame_copy_props() (and hence by
    av_frame_ref());
  - it is set to NULL when the frame is cleared by av_frame_unref()
  - on the caller's explicit request. E.g. libavcodec encoders/decoders
    will copy this field to/from @ref AVPacket "AVPackets" if the caller sets
    @ref AV_CODEC_FLAG_COPY_OPAQUE.

  @see opaque_ref the reference-counted analogue
*/
func (s *AVFrame) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
/*
  Frame owner's private data.

  This field may be set by the code that allocates/owns the frame data.
  It is then not touched by any library functions, except:
  - it is copied to other references by av_frame_copy_props() (and hence by
    av_frame_ref());
  - it is set to NULL when the frame is cleared by av_frame_unref()
  - on the caller's explicit request. E.g. libavcodec encoders/decoders
    will copy this field to/from @ref AVPacket "AVPackets" if the caller sets
    @ref AV_CODEC_FLAG_COPY_OPAQUE.

  @see opaque_ref the reference-counted analogue
*/
func (s *AVFrame) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// RepeatPict gets the repeat_pict field.
/*
  Number of fields in this frame which should be repeated, i.e. the total
  duration of this frame should be repeat_pict + 2 normal field durations.

  For interlaced frames this field may be set to 1, which signals that this
  frame should be presented as 3 fields: beginning with the first field (as
  determined by AV_FRAME_FLAG_TOP_FIELD_FIRST being set or not), followed
  by the second field, and then the first field again.

  For progressive frames this field may be set to a multiple of 2, which
  signals that this frame's duration should be (repeat_pict + 2) / 2
  normal frame durations.

  @note This field is computed from MPEG2 repeat_first_field flag and its
  associated flags, H.264 pic_struct from picture timing SEI, and
  their analogues in other codecs. Typically it should only be used when
  higher-layer timing information is not available.
*/
func (s *AVFrame) RepeatPict() int {
	value := s.ptr.repeat_pict
	return int(value)
}

// SetRepeatPict sets the repeat_pict field.
/*
  Number of fields in this frame which should be repeated, i.e. the total
  duration of this frame should be repeat_pict + 2 normal field durations.

  For interlaced frames this field may be set to 1, which signals that this
  frame should be presented as 3 fields: beginning with the first field (as
  determined by AV_FRAME_FLAG_TOP_FIELD_FIRST being set or not), followed
  by the second field, and then the first field again.

  For progressive frames this field may be set to a multiple of 2, which
  signals that this frame's duration should be (repeat_pict + 2) / 2
  normal frame durations.

  @note This field is computed from MPEG2 repeat_first_field flag and its
  associated flags, H.264 pic_struct from picture timing SEI, and
  their analogues in other codecs. Typically it should only be used when
  higher-layer timing information is not available.
*/
func (s *AVFrame) SetRepeatPict(value int) {
	s.ptr.repeat_pict = (C.int)(value)
}

// InterlacedFrame gets the interlaced_frame field.
/*
  The content of the picture is interlaced.

  @deprecated Use AV_FRAME_FLAG_INTERLACED instead
*/
func (s *AVFrame) InterlacedFrame() int {
	value := s.ptr.interlaced_frame
	return int(value)
}

// SetInterlacedFrame sets the interlaced_frame field.
/*
  The content of the picture is interlaced.

  @deprecated Use AV_FRAME_FLAG_INTERLACED instead
*/
func (s *AVFrame) SetInterlacedFrame(value int) {
	s.ptr.interlaced_frame = (C.int)(value)
}

// TopFieldFirst gets the top_field_first field.
/*
  If the content is interlaced, is top field displayed first.

  @deprecated Use AV_FRAME_FLAG_TOP_FIELD_FIRST instead
*/
func (s *AVFrame) TopFieldFirst() int {
	value := s.ptr.top_field_first
	return int(value)
}

// SetTopFieldFirst sets the top_field_first field.
/*
  If the content is interlaced, is top field displayed first.

  @deprecated Use AV_FRAME_FLAG_TOP_FIELD_FIRST instead
*/
func (s *AVFrame) SetTopFieldFirst(value int) {
	s.ptr.top_field_first = (C.int)(value)
}

// PaletteHasChanged gets the palette_has_changed field.
//
//	Tell user application that palette has changed from previous frame.
func (s *AVFrame) PaletteHasChanged() int {
	value := s.ptr.palette_has_changed
	return int(value)
}

// SetPaletteHasChanged sets the palette_has_changed field.
//
//	Tell user application that palette has changed from previous frame.
func (s *AVFrame) SetPaletteHasChanged(value int) {
	s.ptr.palette_has_changed = (C.int)(value)
}

// ReorderedOpaque gets the reordered_opaque field.
/*
  reordered opaque 64 bits (generally an integer or a double precision float
  PTS but can be anything).
  The user sets AVCodecContext.reordered_opaque to represent the input at
  that time,
  the decoder reorders values as needed and sets AVFrame.reordered_opaque
  to exactly one of the values provided by the user through AVCodecContext.reordered_opaque

  @deprecated Use AV_CODEC_FLAG_COPY_OPAQUE instead
*/
func (s *AVFrame) ReorderedOpaque() int64 {
	value := s.ptr.reordered_opaque
	return int64(value)
}

// SetReorderedOpaque sets the reordered_opaque field.
/*
  reordered opaque 64 bits (generally an integer or a double precision float
  PTS but can be anything).
  The user sets AVCodecContext.reordered_opaque to represent the input at
  that time,
  the decoder reorders values as needed and sets AVFrame.reordered_opaque
  to exactly one of the values provided by the user through AVCodecContext.reordered_opaque

  @deprecated Use AV_CODEC_FLAG_COPY_OPAQUE instead
*/
func (s *AVFrame) SetReorderedOpaque(value int64) {
	s.ptr.reordered_opaque = (C.int64_t)(value)
}

// SampleRate gets the sample_rate field.
//
//	Sample rate of the audio data.
func (s *AVFrame) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

// SetSampleRate sets the sample_rate field.
//
//	Sample rate of the audio data.
func (s *AVFrame) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

// ChannelLayout gets the channel_layout field.
/*
  Channel layout of the audio data.
  @deprecated use ch_layout instead
*/
func (s *AVFrame) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

// SetChannelLayout sets the channel_layout field.
/*
  Channel layout of the audio data.
  @deprecated use ch_layout instead
*/
func (s *AVFrame) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

// Buf gets the buf field.
/*
  AVBuffer references backing the data for this frame. All the pointers in
  data and extended_data must point inside one of the buffers in buf or
  extended_buf. This array must be filled contiguously -- if buf[i] is
  non-NULL then buf[j] must also be non-NULL for all j < i.

  There may be at most one AVBuffer per data plane, so for video this array
  always contains all the references. For planar audio with more than
  AV_NUM_DATA_POINTERS channels, there may be more buffers than can fit in
  this array. Then the extra AVBufferRef pointers are stored in the
  extended_buf array.
*/
func (s *AVFrame) Buf() *Array[*AVBufferRef] {
	value := &s.ptr.buf
	return ToAVBufferRefArray(unsafe.Pointer(value))
}

// ExtendedBuf gets the extended_buf field.
/*
  For planar audio which requires more than AV_NUM_DATA_POINTERS
  AVBufferRef pointers, this array will hold all the references which
  cannot fit into AVFrame.buf.

  Note that this is different from AVFrame.extended_data, which always
  contains all the pointers. This array only contains the extra pointers,
  which cannot fit into AVFrame.buf.

  This array is always allocated using av_malloc() by whoever constructs
  the frame. It is freed in av_frame_unref().
*/
func (s *AVFrame) ExtendedBuf() *Array[*AVBufferRef] {
	value := s.ptr.extended_buf
	return ToAVBufferRefArray(unsafe.Pointer(value))
}

// SetExtendedBuf sets the extended_buf field.
/*
  For planar audio which requires more than AV_NUM_DATA_POINTERS
  AVBufferRef pointers, this array will hold all the references which
  cannot fit into AVFrame.buf.

  Note that this is different from AVFrame.extended_data, which always
  contains all the pointers. This array only contains the extra pointers,
  which cannot fit into AVFrame.buf.

  This array is always allocated using av_malloc() by whoever constructs
  the frame. It is freed in av_frame_unref().
*/
func (s *AVFrame) SetExtendedBuf(value *Array[AVBufferRef]) {
	if value != nil {
		s.ptr.extended_buf = (**C.AVBufferRef)(value.ptr)
	} else {
		s.ptr.extended_buf = nil
	}
}

// NbExtendedBuf gets the nb_extended_buf field.
//
//	Number of elements in extended_buf.
func (s *AVFrame) NbExtendedBuf() int {
	value := s.ptr.nb_extended_buf
	return int(value)
}

// SetNbExtendedBuf sets the nb_extended_buf field.
//
//	Number of elements in extended_buf.
func (s *AVFrame) SetNbExtendedBuf(value int) {
	s.ptr.nb_extended_buf = (C.int)(value)
}

// SideData gets the side_data field.
func (s *AVFrame) SideData() *Array[*AVFrameSideData] {
	value := s.ptr.side_data
	return ToAVFrameSideDataArray(unsafe.Pointer(value))
}

// SetSideData sets the side_data field.
func (s *AVFrame) SetSideData(value *Array[AVFrameSideData]) {
	if value != nil {
		s.ptr.side_data = (**C.AVFrameSideData)(value.ptr)
	} else {
		s.ptr.side_data = nil
	}
}

// NbSideData gets the nb_side_data field.
func (s *AVFrame) NbSideData() int {
	value := s.ptr.nb_side_data
	return int(value)
}

// SetNbSideData sets the nb_side_data field.
func (s *AVFrame) SetNbSideData(value int) {
	s.ptr.nb_side_data = (C.int)(value)
}

// Flags gets the flags field.
//
//	Frame flags, a combination of @ref lavu_frame_flags
func (s *AVFrame) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
//
//	Frame flags, a combination of @ref lavu_frame_flags
func (s *AVFrame) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// ColorRange gets the color_range field.
/*
  MPEG vs JPEG YUV range.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVFrame) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}

// SetColorRange sets the color_range field.
/*
  MPEG vs JPEG YUV range.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVFrame) SetColorRange(value AVColorRange) {
	s.ptr.color_range = (C.enum_AVColorRange)(value)
}

// ColorPrimaries gets the color_primaries field.
func (s *AVFrame) ColorPrimaries() AVColorPrimaries {
	value := s.ptr.color_primaries
	return AVColorPrimaries(value)
}

// SetColorPrimaries sets the color_primaries field.
func (s *AVFrame) SetColorPrimaries(value AVColorPrimaries) {
	s.ptr.color_primaries = (C.enum_AVColorPrimaries)(value)
}

// ColorTrc gets the color_trc field.
func (s *AVFrame) ColorTrc() AVColorTransferCharacteristic {
	value := s.ptr.color_trc
	return AVColorTransferCharacteristic(value)
}

// SetColorTrc sets the color_trc field.
func (s *AVFrame) SetColorTrc(value AVColorTransferCharacteristic) {
	s.ptr.color_trc = (C.enum_AVColorTransferCharacteristic)(value)
}

// Colorspace gets the colorspace field.
/*
  YUV colorspace type.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVFrame) Colorspace() AVColorSpace {
	value := s.ptr.colorspace
	return AVColorSpace(value)
}

// SetColorspace sets the colorspace field.
/*
  YUV colorspace type.
  - encoding: Set by user
  - decoding: Set by libavcodec
*/
func (s *AVFrame) SetColorspace(value AVColorSpace) {
	s.ptr.colorspace = (C.enum_AVColorSpace)(value)
}

// ChromaLocation gets the chroma_location field.
func (s *AVFrame) ChromaLocation() AVChromaLocation {
	value := s.ptr.chroma_location
	return AVChromaLocation(value)
}

// SetChromaLocation sets the chroma_location field.
func (s *AVFrame) SetChromaLocation(value AVChromaLocation) {
	s.ptr.chroma_location = (C.enum_AVChromaLocation)(value)
}

// BestEffortTimestamp gets the best_effort_timestamp field.
/*
  frame timestamp estimated using various heuristics, in stream time base
  - encoding: unused
  - decoding: set by libavcodec, read by user.
*/
func (s *AVFrame) BestEffortTimestamp() int64 {
	value := s.ptr.best_effort_timestamp
	return int64(value)
}

// SetBestEffortTimestamp sets the best_effort_timestamp field.
/*
  frame timestamp estimated using various heuristics, in stream time base
  - encoding: unused
  - decoding: set by libavcodec, read by user.
*/
func (s *AVFrame) SetBestEffortTimestamp(value int64) {
	s.ptr.best_effort_timestamp = (C.int64_t)(value)
}

// PktPos gets the pkt_pos field.
/*
  reordered pos from the last AVPacket that has been input into the decoder
  - encoding: unused
  - decoding: Read by user.
  @deprecated use AV_CODEC_FLAG_COPY_OPAQUE to pass through arbitrary user
              data from packets to frames
*/
func (s *AVFrame) PktPos() int64 {
	value := s.ptr.pkt_pos
	return int64(value)
}

// SetPktPos sets the pkt_pos field.
/*
  reordered pos from the last AVPacket that has been input into the decoder
  - encoding: unused
  - decoding: Read by user.
  @deprecated use AV_CODEC_FLAG_COPY_OPAQUE to pass through arbitrary user
              data from packets to frames
*/
func (s *AVFrame) SetPktPos(value int64) {
	s.ptr.pkt_pos = (C.int64_t)(value)
}

// PktDuration gets the pkt_duration field.
/*
  duration of the corresponding packet, expressed in
  AVStream->time_base units, 0 if unknown.
  - encoding: unused
  - decoding: Read by user.

  @deprecated use duration instead
*/
func (s *AVFrame) PktDuration() int64 {
	value := s.ptr.pkt_duration
	return int64(value)
}

// SetPktDuration sets the pkt_duration field.
/*
  duration of the corresponding packet, expressed in
  AVStream->time_base units, 0 if unknown.
  - encoding: unused
  - decoding: Read by user.

  @deprecated use duration instead
*/
func (s *AVFrame) SetPktDuration(value int64) {
	s.ptr.pkt_duration = (C.int64_t)(value)
}

// Metadata gets the metadata field.
/*
  metadata.
  - encoding: Set by user.
  - decoding: Set by libavcodec.
*/
func (s *AVFrame) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

// SetMetadata sets the metadata field.
/*
  metadata.
  - encoding: Set by user.
  - decoding: Set by libavcodec.
*/
func (s *AVFrame) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

// DecodeErrorFlags gets the decode_error_flags field.
/*
  decode error flags of the frame, set to a combination of
  FF_DECODE_ERROR_xxx flags if the decoder produced a frame, but there
  were errors during the decoding.
  - encoding: unused
  - decoding: set by libavcodec, read by user.
*/
func (s *AVFrame) DecodeErrorFlags() int {
	value := s.ptr.decode_error_flags
	return int(value)
}

// SetDecodeErrorFlags sets the decode_error_flags field.
/*
  decode error flags of the frame, set to a combination of
  FF_DECODE_ERROR_xxx flags if the decoder produced a frame, but there
  were errors during the decoding.
  - encoding: unused
  - decoding: set by libavcodec, read by user.
*/
func (s *AVFrame) SetDecodeErrorFlags(value int) {
	s.ptr.decode_error_flags = (C.int)(value)
}

// Channels gets the channels field.
/*
  number of audio channels, only used for audio.
  - encoding: unused
  - decoding: Read by user.
  @deprecated use ch_layout instead
*/
func (s *AVFrame) Channels() int {
	value := s.ptr.channels
	return int(value)
}

// SetChannels sets the channels field.
/*
  number of audio channels, only used for audio.
  - encoding: unused
  - decoding: Read by user.
  @deprecated use ch_layout instead
*/
func (s *AVFrame) SetChannels(value int) {
	s.ptr.channels = (C.int)(value)
}

// PktSize gets the pkt_size field.
/*
  size of the corresponding packet containing the compressed
  frame.
  It is set to a negative value if unknown.
  - encoding: unused
  - decoding: set by libavcodec, read by user.
  @deprecated use AV_CODEC_FLAG_COPY_OPAQUE to pass through arbitrary user
              data from packets to frames
*/
func (s *AVFrame) PktSize() int {
	value := s.ptr.pkt_size
	return int(value)
}

// SetPktSize sets the pkt_size field.
/*
  size of the corresponding packet containing the compressed
  frame.
  It is set to a negative value if unknown.
  - encoding: unused
  - decoding: set by libavcodec, read by user.
  @deprecated use AV_CODEC_FLAG_COPY_OPAQUE to pass through arbitrary user
              data from packets to frames
*/
func (s *AVFrame) SetPktSize(value int) {
	s.ptr.pkt_size = (C.int)(value)
}

// HwFramesCtx gets the hw_frames_ctx field.
/*
  For hwaccel-format frames, this should be a reference to the
  AVHWFramesContext describing the frame.
*/
func (s *AVFrame) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwFramesCtx sets the hw_frames_ctx field.
/*
  For hwaccel-format frames, this should be a reference to the
  AVHWFramesContext describing the frame.
*/
func (s *AVFrame) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

// OpaqueRef gets the opaque_ref field.
/*
  Frame owner's private data.

  This field may be set by the code that allocates/owns the frame data.
  It is then not touched by any library functions, except:
  - a new reference to the underlying buffer is propagated by
    av_frame_copy_props() (and hence by av_frame_ref());
  - it is unreferenced in av_frame_unref();
  - on the caller's explicit request. E.g. libavcodec encoders/decoders
    will propagate a new reference to/from @ref AVPacket "AVPackets" if the
    caller sets @ref AV_CODEC_FLAG_COPY_OPAQUE.

  @see opaque the plain pointer analogue
*/
func (s *AVFrame) OpaqueRef() *AVBufferRef {
	value := s.ptr.opaque_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetOpaqueRef sets the opaque_ref field.
/*
  Frame owner's private data.

  This field may be set by the code that allocates/owns the frame data.
  It is then not touched by any library functions, except:
  - a new reference to the underlying buffer is propagated by
    av_frame_copy_props() (and hence by av_frame_ref());
  - it is unreferenced in av_frame_unref();
  - on the caller's explicit request. E.g. libavcodec encoders/decoders
    will propagate a new reference to/from @ref AVPacket "AVPackets" if the
    caller sets @ref AV_CODEC_FLAG_COPY_OPAQUE.

  @see opaque the plain pointer analogue
*/
func (s *AVFrame) SetOpaqueRef(value *AVBufferRef) {
	if value != nil {
		s.ptr.opaque_ref = value.ptr
	} else {
		s.ptr.opaque_ref = nil
	}
}

// CropTop gets the crop_top field.
func (s *AVFrame) CropTop() uint64 {
	value := s.ptr.crop_top
	return uint64(value)
}

// SetCropTop sets the crop_top field.
func (s *AVFrame) SetCropTop(value uint64) {
	s.ptr.crop_top = (C.size_t)(value)
}

// CropBottom gets the crop_bottom field.
func (s *AVFrame) CropBottom() uint64 {
	value := s.ptr.crop_bottom
	return uint64(value)
}

// SetCropBottom sets the crop_bottom field.
func (s *AVFrame) SetCropBottom(value uint64) {
	s.ptr.crop_bottom = (C.size_t)(value)
}

// CropLeft gets the crop_left field.
func (s *AVFrame) CropLeft() uint64 {
	value := s.ptr.crop_left
	return uint64(value)
}

// SetCropLeft sets the crop_left field.
func (s *AVFrame) SetCropLeft(value uint64) {
	s.ptr.crop_left = (C.size_t)(value)
}

// CropRight gets the crop_right field.
func (s *AVFrame) CropRight() uint64 {
	value := s.ptr.crop_right
	return uint64(value)
}

// SetCropRight sets the crop_right field.
func (s *AVFrame) SetCropRight(value uint64) {
	s.ptr.crop_right = (C.size_t)(value)
}

// PrivateRef gets the private_ref field.
/*
  AVBufferRef for internal use by a single libav* library.
  Must not be used to transfer data between libraries.
  Has to be NULL when ownership of the frame leaves the respective library.

  Code outside the FFmpeg libs should never check or change the contents of the buffer ref.

  FFmpeg calls av_buffer_unref() on it when the frame is unreferenced.
  av_frame_copy_props() calls create a new reference with av_buffer_ref()
  for the target frame's private_ref field.
*/
func (s *AVFrame) PrivateRef() *AVBufferRef {
	value := s.ptr.private_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetPrivateRef sets the private_ref field.
/*
  AVBufferRef for internal use by a single libav* library.
  Must not be used to transfer data between libraries.
  Has to be NULL when ownership of the frame leaves the respective library.

  Code outside the FFmpeg libs should never check or change the contents of the buffer ref.

  FFmpeg calls av_buffer_unref() on it when the frame is unreferenced.
  av_frame_copy_props() calls create a new reference with av_buffer_ref()
  for the target frame's private_ref field.
*/
func (s *AVFrame) SetPrivateRef(value *AVBufferRef) {
	if value != nil {
		s.ptr.private_ref = value.ptr
	} else {
		s.ptr.private_ref = nil
	}
}

// ChLayout gets the ch_layout field.
//
//	Channel layout of the audio data.
func (s *AVFrame) ChLayout() *AVChannelLayout {
	value := &s.ptr.ch_layout
	return &AVChannelLayout{ptr: value}
}

// Duration gets the duration field.
//
//	Duration of the frame, in the same units as pts. 0 if unknown.
func (s *AVFrame) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

// SetDuration sets the duration field.
//
//	Duration of the frame, in the same units as pts. 0 if unknown.
func (s *AVFrame) SetDuration(value int64) {
	s.ptr.duration = (C.int64_t)(value)
}

// --- Struct AVHWDeviceInternal ---

// AVHWDeviceInternal wraps AVHWDeviceInternal.
type AVHWDeviceInternal struct {
	ptr *C.AVHWDeviceInternal
}

func (s *AVHWDeviceInternal) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVHWDeviceInternalArray(ptr unsafe.Pointer) *Array[*AVHWDeviceInternal] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVHWDeviceInternal]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVHWDeviceInternal {
			ptr := (**C.AVHWDeviceInternal)(pointer)
			value := *ptr
			var valueMapped *AVHWDeviceInternal
			if value != nil {
				valueMapped = &AVHWDeviceInternal{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVHWDeviceInternal) {
			ptr := (**C.AVHWDeviceInternal)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVHWDeviceContext ---

// AVHWDeviceContext wraps AVHWDeviceContext.
/*
  This struct aggregates all the (hardware/vendor-specific) "high-level" state,
  i.e. state that is not tied to a concrete processing configuration.
  E.g., in an API that supports hardware-accelerated encoding and decoding,
  this struct will (if possible) wrap the state that is common to both encoding
  and decoding and from which specific instances of encoders or decoders can be
  derived.

  This struct is reference-counted with the AVBuffer mechanism. The
  av_hwdevice_ctx_alloc() constructor yields a reference, whose data field
  points to the actual AVHWDeviceContext. Further objects derived from
  AVHWDeviceContext (such as AVHWFramesContext, describing a frame pool with
  specific properties) will hold an internal reference to it. After all the
  references are released, the AVHWDeviceContext itself will be freed,
  optionally invoking a user-specified callback for uninitializing the hardware
  state.
*/
type AVHWDeviceContext struct {
	ptr *C.AVHWDeviceContext
}

func (s *AVHWDeviceContext) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVHWDeviceContextArray(ptr unsafe.Pointer) *Array[*AVHWDeviceContext] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVHWDeviceContext]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVHWDeviceContext {
			ptr := (**C.AVHWDeviceContext)(pointer)
			value := *ptr
			var valueMapped *AVHWDeviceContext
			if value != nil {
				valueMapped = &AVHWDeviceContext{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVHWDeviceContext) {
			ptr := (**C.AVHWDeviceContext)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// AvClass gets the av_class field.
//
//	A class for logging. Set by av_hwdevice_ctx_alloc().
func (s *AVHWDeviceContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
//
//	A class for logging. Set by av_hwdevice_ctx_alloc().
func (s *AVHWDeviceContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Internal gets the internal field.
/*
  Private data used internally by libavutil. Must not be accessed in any
  way by the caller.
*/
func (s *AVHWDeviceContext) Internal() *AVHWDeviceInternal {
	value := s.ptr.internal
	var valueMapped *AVHWDeviceInternal
	if value != nil {
		valueMapped = &AVHWDeviceInternal{ptr: value}
	}
	return valueMapped
}

// SetInternal sets the internal field.
/*
  Private data used internally by libavutil. Must not be accessed in any
  way by the caller.
*/
func (s *AVHWDeviceContext) SetInternal(value *AVHWDeviceInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

// Type gets the type field.
/*
  This field identifies the underlying API used for hardware access.

  This field is set when this struct is allocated and never changed
  afterwards.
*/
func (s *AVHWDeviceContext) Type() AVHWDeviceType {
	value := s.ptr._type
	return AVHWDeviceType(value)
}

// SetType sets the type field.
/*
  This field identifies the underlying API used for hardware access.

  This field is set when this struct is allocated and never changed
  afterwards.
*/
func (s *AVHWDeviceContext) SetType(value AVHWDeviceType) {
	s.ptr._type = (C.enum_AVHWDeviceType)(value)
}

// Hwctx gets the hwctx field.
/*
  The format-specific data, allocated and freed by libavutil along with
  this context.

  Should be cast by the user to the format-specific context defined in the
  corresponding header (hwcontext_*.h) and filled as described in the
  documentation before calling av_hwdevice_ctx_init().

  After calling av_hwdevice_ctx_init() this struct should not be modified
  by the caller.
*/
func (s *AVHWDeviceContext) Hwctx() unsafe.Pointer {
	value := s.ptr.hwctx
	return value
}

// SetHwctx sets the hwctx field.
/*
  The format-specific data, allocated and freed by libavutil along with
  this context.

  Should be cast by the user to the format-specific context defined in the
  corresponding header (hwcontext_*.h) and filled as described in the
  documentation before calling av_hwdevice_ctx_init().

  After calling av_hwdevice_ctx_init() this struct should not be modified
  by the caller.
*/
func (s *AVHWDeviceContext) SetHwctx(value unsafe.Pointer) {
	s.ptr.hwctx = value
}

// free skipped due to func ptr

// UserOpaque gets the user_opaque field.
//
//	Arbitrary user data, to be used e.g. by the free() callback.
func (s *AVHWDeviceContext) UserOpaque() unsafe.Pointer {
	value := s.ptr.user_opaque
	return value
}

// SetUserOpaque sets the user_opaque field.
//
//	Arbitrary user data, to be used e.g. by the free() callback.
func (s *AVHWDeviceContext) SetUserOpaque(value unsafe.Pointer) {
	s.ptr.user_opaque = value
}

// --- Struct AVHWFramesInternal ---

// AVHWFramesInternal wraps AVHWFramesInternal.
type AVHWFramesInternal struct {
	ptr *C.AVHWFramesInternal
}

func (s *AVHWFramesInternal) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVHWFramesInternalArray(ptr unsafe.Pointer) *Array[*AVHWFramesInternal] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVHWFramesInternal]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVHWFramesInternal {
			ptr := (**C.AVHWFramesInternal)(pointer)
			value := *ptr
			var valueMapped *AVHWFramesInternal
			if value != nil {
				valueMapped = &AVHWFramesInternal{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVHWFramesInternal) {
			ptr := (**C.AVHWFramesInternal)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// --- Struct AVHWFramesContext ---

// AVHWFramesContext wraps AVHWFramesContext.
/*
  This struct describes a set or pool of "hardware" frames (i.e. those with
  data not located in normal system memory). All the frames in the pool are
  assumed to be allocated in the same way and interchangeable.

  This struct is reference-counted with the AVBuffer mechanism and tied to a
  given AVHWDeviceContext instance. The av_hwframe_ctx_alloc() constructor
  yields a reference, whose data field points to the actual AVHWFramesContext
  struct.
*/
type AVHWFramesContext struct {
	ptr *C.AVHWFramesContext
}

func (s *AVHWFramesContext) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVHWFramesContextArray(ptr unsafe.Pointer) *Array[*AVHWFramesContext] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVHWFramesContext]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVHWFramesContext {
			ptr := (**C.AVHWFramesContext)(pointer)
			value := *ptr
			var valueMapped *AVHWFramesContext
			if value != nil {
				valueMapped = &AVHWFramesContext{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVHWFramesContext) {
			ptr := (**C.AVHWFramesContext)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// AvClass gets the av_class field.
//
//	A class for logging.
func (s *AVHWFramesContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
//
//	A class for logging.
func (s *AVHWFramesContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Internal gets the internal field.
/*
  Private data used internally by libavutil. Must not be accessed in any
  way by the caller.
*/
func (s *AVHWFramesContext) Internal() *AVHWFramesInternal {
	value := s.ptr.internal
	var valueMapped *AVHWFramesInternal
	if value != nil {
		valueMapped = &AVHWFramesInternal{ptr: value}
	}
	return valueMapped
}

// SetInternal sets the internal field.
/*
  Private data used internally by libavutil. Must not be accessed in any
  way by the caller.
*/
func (s *AVHWFramesContext) SetInternal(value *AVHWFramesInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

// DeviceRef gets the device_ref field.
/*
  A reference to the parent AVHWDeviceContext. This reference is owned and
  managed by the enclosing AVHWFramesContext, but the caller may derive
  additional references from it.
*/
func (s *AVHWFramesContext) DeviceRef() *AVBufferRef {
	value := s.ptr.device_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetDeviceRef sets the device_ref field.
/*
  A reference to the parent AVHWDeviceContext. This reference is owned and
  managed by the enclosing AVHWFramesContext, but the caller may derive
  additional references from it.
*/
func (s *AVHWFramesContext) SetDeviceRef(value *AVBufferRef) {
	if value != nil {
		s.ptr.device_ref = value.ptr
	} else {
		s.ptr.device_ref = nil
	}
}

// DeviceCtx gets the device_ctx field.
/*
  The parent AVHWDeviceContext. This is simply a pointer to
  device_ref->data provided for convenience.

  Set by libavutil in av_hwframe_ctx_init().
*/
func (s *AVHWFramesContext) DeviceCtx() *AVHWDeviceContext {
	value := s.ptr.device_ctx
	var valueMapped *AVHWDeviceContext
	if value != nil {
		valueMapped = &AVHWDeviceContext{ptr: value}
	}
	return valueMapped
}

// SetDeviceCtx sets the device_ctx field.
/*
  The parent AVHWDeviceContext. This is simply a pointer to
  device_ref->data provided for convenience.

  Set by libavutil in av_hwframe_ctx_init().
*/
func (s *AVHWFramesContext) SetDeviceCtx(value *AVHWDeviceContext) {
	if value != nil {
		s.ptr.device_ctx = value.ptr
	} else {
		s.ptr.device_ctx = nil
	}
}

// Hwctx gets the hwctx field.
/*
  The format-specific data, allocated and freed automatically along with
  this context.

  Should be cast by the user to the format-specific context defined in the
  corresponding header (hwframe_*.h) and filled as described in the
  documentation before calling av_hwframe_ctx_init().

  After any frames using this context are created, the contents of this
  struct should not be modified by the caller.
*/
func (s *AVHWFramesContext) Hwctx() unsafe.Pointer {
	value := s.ptr.hwctx
	return value
}

// SetHwctx sets the hwctx field.
/*
  The format-specific data, allocated and freed automatically along with
  this context.

  Should be cast by the user to the format-specific context defined in the
  corresponding header (hwframe_*.h) and filled as described in the
  documentation before calling av_hwframe_ctx_init().

  After any frames using this context are created, the contents of this
  struct should not be modified by the caller.
*/
func (s *AVHWFramesContext) SetHwctx(value unsafe.Pointer) {
	s.ptr.hwctx = value
}

// free skipped due to func ptr

// UserOpaque gets the user_opaque field.
//
//	Arbitrary user data, to be used e.g. by the free() callback.
func (s *AVHWFramesContext) UserOpaque() unsafe.Pointer {
	value := s.ptr.user_opaque
	return value
}

// SetUserOpaque sets the user_opaque field.
//
//	Arbitrary user data, to be used e.g. by the free() callback.
func (s *AVHWFramesContext) SetUserOpaque(value unsafe.Pointer) {
	s.ptr.user_opaque = value
}

// Pool gets the pool field.
/*
  A pool from which the frames are allocated by av_hwframe_get_buffer().
  This field may be set by the caller before calling av_hwframe_ctx_init().
  The buffers returned by calling av_buffer_pool_get() on this pool must
  have the properties described in the documentation in the corresponding hw
  type's header (hwcontext_*.h). The pool will be freed strictly before
  this struct's free() callback is invoked.

  This field may be NULL, then libavutil will attempt to allocate a pool
  internally. Note that certain device types enforce pools allocated at
  fixed size (frame count), which cannot be extended dynamically. In such a
  case, initial_pool_size must be set appropriately.
*/
func (s *AVHWFramesContext) Pool() *AVBufferPool {
	value := s.ptr.pool
	var valueMapped *AVBufferPool
	if value != nil {
		valueMapped = &AVBufferPool{ptr: value}
	}
	return valueMapped
}

// SetPool sets the pool field.
/*
  A pool from which the frames are allocated by av_hwframe_get_buffer().
  This field may be set by the caller before calling av_hwframe_ctx_init().
  The buffers returned by calling av_buffer_pool_get() on this pool must
  have the properties described in the documentation in the corresponding hw
  type's header (hwcontext_*.h). The pool will be freed strictly before
  this struct's free() callback is invoked.

  This field may be NULL, then libavutil will attempt to allocate a pool
  internally. Note that certain device types enforce pools allocated at
  fixed size (frame count), which cannot be extended dynamically. In such a
  case, initial_pool_size must be set appropriately.
*/
func (s *AVHWFramesContext) SetPool(value *AVBufferPool) {
	if value != nil {
		s.ptr.pool = value.ptr
	} else {
		s.ptr.pool = nil
	}
}

// InitialPoolSize gets the initial_pool_size field.
/*
  Initial size of the frame pool. If a device type does not support
  dynamically resizing the pool, then this is also the maximum pool size.

  May be set by the caller before calling av_hwframe_ctx_init(). Must be
  set if pool is NULL and the device type does not support dynamic pools.
*/
func (s *AVHWFramesContext) InitialPoolSize() int {
	value := s.ptr.initial_pool_size
	return int(value)
}

// SetInitialPoolSize sets the initial_pool_size field.
/*
  Initial size of the frame pool. If a device type does not support
  dynamically resizing the pool, then this is also the maximum pool size.

  May be set by the caller before calling av_hwframe_ctx_init(). Must be
  set if pool is NULL and the device type does not support dynamic pools.
*/
func (s *AVHWFramesContext) SetInitialPoolSize(value int) {
	s.ptr.initial_pool_size = (C.int)(value)
}

// Format gets the format field.
/*
  The pixel format identifying the underlying HW surface type.

  Must be a hwaccel format, i.e. the corresponding descriptor must have the
  AV_PIX_FMT_FLAG_HWACCEL flag set.

  Must be set by the user before calling av_hwframe_ctx_init().
*/
func (s *AVHWFramesContext) Format() AVPixelFormat {
	value := s.ptr.format
	return AVPixelFormat(value)
}

// SetFormat sets the format field.
/*
  The pixel format identifying the underlying HW surface type.

  Must be a hwaccel format, i.e. the corresponding descriptor must have the
  AV_PIX_FMT_FLAG_HWACCEL flag set.

  Must be set by the user before calling av_hwframe_ctx_init().
*/
func (s *AVHWFramesContext) SetFormat(value AVPixelFormat) {
	s.ptr.format = (C.enum_AVPixelFormat)(value)
}

// SwFormat gets the sw_format field.
/*
  The pixel format identifying the actual data layout of the hardware
  frames.

  Must be set by the caller before calling av_hwframe_ctx_init().

  @note when the underlying API does not provide the exact data layout, but
  only the colorspace/bit depth, this field should be set to the fully
  planar version of that format (e.g. for 8-bit 420 YUV it should be
  AV_PIX_FMT_YUV420P, not AV_PIX_FMT_NV12 or anything else).
*/
func (s *AVHWFramesContext) SwFormat() AVPixelFormat {
	value := s.ptr.sw_format
	return AVPixelFormat(value)
}

// SetSwFormat sets the sw_format field.
/*
  The pixel format identifying the actual data layout of the hardware
  frames.

  Must be set by the caller before calling av_hwframe_ctx_init().

  @note when the underlying API does not provide the exact data layout, but
  only the colorspace/bit depth, this field should be set to the fully
  planar version of that format (e.g. for 8-bit 420 YUV it should be
  AV_PIX_FMT_YUV420P, not AV_PIX_FMT_NV12 or anything else).
*/
func (s *AVHWFramesContext) SetSwFormat(value AVPixelFormat) {
	s.ptr.sw_format = (C.enum_AVPixelFormat)(value)
}

// Width gets the width field.
/*
  The allocated dimensions of the frames in this pool.

  Must be set by the user before calling av_hwframe_ctx_init().
*/
func (s *AVHWFramesContext) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
/*
  The allocated dimensions of the frames in this pool.

  Must be set by the user before calling av_hwframe_ctx_init().
*/
func (s *AVHWFramesContext) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

// Height gets the height field.
/*
  The allocated dimensions of the frames in this pool.

  Must be set by the user before calling av_hwframe_ctx_init().
*/
func (s *AVHWFramesContext) Height() int {
	value := s.ptr.height
	return int(value)
}

// SetHeight sets the height field.
/*
  The allocated dimensions of the frames in this pool.

  Must be set by the user before calling av_hwframe_ctx_init().
*/
func (s *AVHWFramesContext) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

// --- Struct AVHWFramesConstraints ---

// AVHWFramesConstraints wraps AVHWFramesConstraints.
/*
  This struct describes the constraints on hardware frames attached to
  a given device with a hardware-specific configuration.  This is returned
  by av_hwdevice_get_hwframe_constraints() and must be freed by
  av_hwframe_constraints_free() after use.
*/
type AVHWFramesConstraints struct {
	ptr *C.AVHWFramesConstraints
}

func (s *AVHWFramesConstraints) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVHWFramesConstraintsArray(ptr unsafe.Pointer) *Array[*AVHWFramesConstraints] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVHWFramesConstraints]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVHWFramesConstraints {
			ptr := (**C.AVHWFramesConstraints)(pointer)
			value := *ptr
			var valueMapped *AVHWFramesConstraints
			if value != nil {
				valueMapped = &AVHWFramesConstraints{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVHWFramesConstraints) {
			ptr := (**C.AVHWFramesConstraints)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// ValidHwFormats gets the valid_hw_formats field.
/*
  A list of possible values for format in the hw_frames_ctx,
  terminated by AV_PIX_FMT_NONE.  This member will always be filled.
*/
func (s *AVHWFramesConstraints) ValidHwFormats() *Array[AVPixelFormat] {
	value := s.ptr.valid_hw_formats
	return ToAVPixelFormatArray(unsafe.Pointer(value))
}

// SetValidHwFormats sets the valid_hw_formats field.
/*
  A list of possible values for format in the hw_frames_ctx,
  terminated by AV_PIX_FMT_NONE.  This member will always be filled.
*/
func (s *AVHWFramesConstraints) SetValidHwFormats(value *Array[AVPixelFormat]) {
	if value != nil {
		s.ptr.valid_hw_formats = (*C.enum_AVPixelFormat)(value.ptr)
	} else {
		s.ptr.valid_hw_formats = nil
	}
}

// ValidSwFormats gets the valid_sw_formats field.
/*
  A list of possible values for sw_format in the hw_frames_ctx,
  terminated by AV_PIX_FMT_NONE.  Can be NULL if this information is
  not known.
*/
func (s *AVHWFramesConstraints) ValidSwFormats() *Array[AVPixelFormat] {
	value := s.ptr.valid_sw_formats
	return ToAVPixelFormatArray(unsafe.Pointer(value))
}

// SetValidSwFormats sets the valid_sw_formats field.
/*
  A list of possible values for sw_format in the hw_frames_ctx,
  terminated by AV_PIX_FMT_NONE.  Can be NULL if this information is
  not known.
*/
func (s *AVHWFramesConstraints) SetValidSwFormats(value *Array[AVPixelFormat]) {
	if value != nil {
		s.ptr.valid_sw_formats = (*C.enum_AVPixelFormat)(value.ptr)
	} else {
		s.ptr.valid_sw_formats = nil
	}
}

// MinWidth gets the min_width field.
/*
  The minimum size of frames in this hw_frames_ctx.
  (Zero if not known.)
*/
func (s *AVHWFramesConstraints) MinWidth() int {
	value := s.ptr.min_width
	return int(value)
}

// SetMinWidth sets the min_width field.
/*
  The minimum size of frames in this hw_frames_ctx.
  (Zero if not known.)
*/
func (s *AVHWFramesConstraints) SetMinWidth(value int) {
	s.ptr.min_width = (C.int)(value)
}

// MinHeight gets the min_height field.
func (s *AVHWFramesConstraints) MinHeight() int {
	value := s.ptr.min_height
	return int(value)
}

// SetMinHeight sets the min_height field.
func (s *AVHWFramesConstraints) SetMinHeight(value int) {
	s.ptr.min_height = (C.int)(value)
}

// MaxWidth gets the max_width field.
/*
  The maximum size of frames in this hw_frames_ctx.
  (INT_MAX if not known / no limit.)
*/
func (s *AVHWFramesConstraints) MaxWidth() int {
	value := s.ptr.max_width
	return int(value)
}

// SetMaxWidth sets the max_width field.
/*
  The maximum size of frames in this hw_frames_ctx.
  (INT_MAX if not known / no limit.)
*/
func (s *AVHWFramesConstraints) SetMaxWidth(value int) {
	s.ptr.max_width = (C.int)(value)
}

// MaxHeight gets the max_height field.
func (s *AVHWFramesConstraints) MaxHeight() int {
	value := s.ptr.max_height
	return int(value)
}

// SetMaxHeight sets the max_height field.
func (s *AVHWFramesConstraints) SetMaxHeight(value int) {
	s.ptr.max_height = (C.int)(value)
}

// --- Struct AVClass ---

// AVClass wraps AVClass.
/*
  Describe the class of an AVClass context structure. That is an
  arbitrary struct of which the first field is a pointer to an
  AVClass struct (e.g. AVCodecContext, AVFormatContext etc.).
*/
type AVClass struct {
	ptr *C.AVClass
}

func (s *AVClass) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVClassArray(ptr unsafe.Pointer) *Array[*AVClass] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVClass]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVClass {
			ptr := (**C.AVClass)(pointer)
			value := *ptr
			var valueMapped *AVClass
			if value != nil {
				valueMapped = &AVClass{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVClass) {
			ptr := (**C.AVClass)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// ClassName gets the class_name field.
/*
  The name of the class; usually it is the same name as the
  context structure type to which the AVClass is associated.
*/
func (s *AVClass) ClassName() *CStr {
	value := s.ptr.class_name
	return wrapCStr(value)
}

// SetClassName sets the class_name field.
/*
  The name of the class; usually it is the same name as the
  context structure type to which the AVClass is associated.
*/
func (s *AVClass) SetClassName(value *CStr) {
	s.ptr.class_name = value.ptr
}

// item_name skipped due to func ptr

// Option gets the option field.
/*
  a pointer to the first option specified in the class if any or NULL

  @see av_set_default_options()
*/
func (s *AVClass) Option() *AVOption {
	value := s.ptr.option
	var valueMapped *AVOption
	if value != nil {
		valueMapped = &AVOption{ptr: value}
	}
	return valueMapped
}

// SetOption sets the option field.
/*
  a pointer to the first option specified in the class if any or NULL

  @see av_set_default_options()
*/
func (s *AVClass) SetOption(value *AVOption) {
	if value != nil {
		s.ptr.option = value.ptr
	} else {
		s.ptr.option = nil
	}
}

// Version gets the version field.
/*
  LIBAVUTIL_VERSION with which this structure was created.
  This is used to allow fields to be added without requiring major
  version bumps everywhere.
*/
func (s *AVClass) Version() int {
	value := s.ptr.version
	return int(value)
}

// SetVersion sets the version field.
/*
  LIBAVUTIL_VERSION with which this structure was created.
  This is used to allow fields to be added without requiring major
  version bumps everywhere.
*/
func (s *AVClass) SetVersion(value int) {
	s.ptr.version = (C.int)(value)
}

// LogLevelOffsetOffset gets the log_level_offset_offset field.
/*
  Offset in the structure where log_level_offset is stored.
  0 means there is no such variable
*/
func (s *AVClass) LogLevelOffsetOffset() int {
	value := s.ptr.log_level_offset_offset
	return int(value)
}

// SetLogLevelOffsetOffset sets the log_level_offset_offset field.
/*
  Offset in the structure where log_level_offset is stored.
  0 means there is no such variable
*/
func (s *AVClass) SetLogLevelOffsetOffset(value int) {
	s.ptr.log_level_offset_offset = (C.int)(value)
}

// ParentLogContextOffset gets the parent_log_context_offset field.
/*
  Offset in the structure where a pointer to the parent context for
  logging is stored. For example a decoder could pass its AVCodecContext
  to eval as such a parent context, which an av_log() implementation
  could then leverage to display the parent context.
  The offset can be NULL.
*/
func (s *AVClass) ParentLogContextOffset() int {
	value := s.ptr.parent_log_context_offset
	return int(value)
}

// SetParentLogContextOffset sets the parent_log_context_offset field.
/*
  Offset in the structure where a pointer to the parent context for
  logging is stored. For example a decoder could pass its AVCodecContext
  to eval as such a parent context, which an av_log() implementation
  could then leverage to display the parent context.
  The offset can be NULL.
*/
func (s *AVClass) SetParentLogContextOffset(value int) {
	s.ptr.parent_log_context_offset = (C.int)(value)
}

// Category gets the category field.
/*
  Category used for visualization (like color)
  This is only set if the category is equal for all objects using this class.
  available since version (51 << 16 | 56 << 8 | 100)
*/
func (s *AVClass) Category() AVClassCategory {
	value := s.ptr.category
	return AVClassCategory(value)
}

// SetCategory sets the category field.
/*
  Category used for visualization (like color)
  This is only set if the category is equal for all objects using this class.
  available since version (51 << 16 | 56 << 8 | 100)
*/
func (s *AVClass) SetCategory(value AVClassCategory) {
	s.ptr.category = (C.AVClassCategory)(value)
}

// get_category skipped due to func ptr

// query_ranges skipped due to func ptr

// child_next skipped due to func ptr

// child_class_iterate skipped due to func ptr

// --- Struct AVOption ---

// AVOption wraps AVOption.
//
//	AVOption
type AVOption struct {
	ptr *C.AVOption
}

func (s *AVOption) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVOptionArray(ptr unsafe.Pointer) *Array[*AVOption] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVOption]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVOption {
			ptr := (**C.AVOption)(pointer)
			value := *ptr
			var valueMapped *AVOption
			if value != nil {
				valueMapped = &AVOption{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVOption) {
			ptr := (**C.AVOption)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Name gets the name field.
func (s *AVOption) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
func (s *AVOption) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// Help gets the help field.
/*
  short English help text
  @todo What about other languages?
*/
func (s *AVOption) Help() *CStr {
	value := s.ptr.help
	return wrapCStr(value)
}

// SetHelp sets the help field.
/*
  short English help text
  @todo What about other languages?
*/
func (s *AVOption) SetHelp(value *CStr) {
	s.ptr.help = value.ptr
}

// Offset gets the offset field.
/*
  The offset relative to the context structure where the option
  value is stored. It should be 0 for named constants.
*/
func (s *AVOption) Offset() int {
	value := s.ptr.offset
	return int(value)
}

// SetOffset sets the offset field.
/*
  The offset relative to the context structure where the option
  value is stored. It should be 0 for named constants.
*/
func (s *AVOption) SetOffset(value int) {
	s.ptr.offset = (C.int)(value)
}

// Type gets the type field.
func (s *AVOption) Type() AVOptionType {
	value := s.ptr._type
	return AVOptionType(value)
}

// SetType sets the type field.
func (s *AVOption) SetType(value AVOptionType) {
	s.ptr._type = (C.enum_AVOptionType)(value)
}

// default_val skipped due to union type

// Min gets the min field.
//
//	minimum valid value for the option
func (s *AVOption) Min() float64 {
	value := s.ptr.min
	return float64(value)
}

// SetMin sets the min field.
//
//	minimum valid value for the option
func (s *AVOption) SetMin(value float64) {
	s.ptr.min = (C.double)(value)
}

// Max gets the max field.
//
//	maximum valid value for the option
func (s *AVOption) Max() float64 {
	value := s.ptr.max
	return float64(value)
}

// SetMax sets the max field.
//
//	maximum valid value for the option
func (s *AVOption) SetMax(value float64) {
	s.ptr.max = (C.double)(value)
}

// Flags gets the flags field.
func (s *AVOption) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVOption) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// Unit gets the unit field.
/*
  The logical unit to which the option belongs. Non-constant
  options and corresponding named constants share the same
  unit. May be NULL.
*/
func (s *AVOption) Unit() *CStr {
	value := s.ptr.unit
	return wrapCStr(value)
}

// SetUnit sets the unit field.
/*
  The logical unit to which the option belongs. Non-constant
  options and corresponding named constants share the same
  unit. May be NULL.
*/
func (s *AVOption) SetUnit(value *CStr) {
	s.ptr.unit = value.ptr
}

// --- Struct AVOptionRange ---

// AVOptionRange wraps AVOptionRange.
//
//	A single allowed range of values, or a single allowed value.
type AVOptionRange struct {
	ptr *C.AVOptionRange
}

func (s *AVOptionRange) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVOptionRangeArray(ptr unsafe.Pointer) *Array[*AVOptionRange] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVOptionRange]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVOptionRange {
			ptr := (**C.AVOptionRange)(pointer)
			value := *ptr
			var valueMapped *AVOptionRange
			if value != nil {
				valueMapped = &AVOptionRange{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVOptionRange) {
			ptr := (**C.AVOptionRange)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Str gets the str field.
func (s *AVOptionRange) Str() *CStr {
	value := s.ptr.str
	return wrapCStr(value)
}

// SetStr sets the str field.
func (s *AVOptionRange) SetStr(value *CStr) {
	s.ptr.str = value.ptr
}

// ValueMin gets the value_min field.
/*
  Value range.
  For string ranges this represents the min/max length.
  For dimensions this represents the min/max pixel count or width/height in multi-component case.
*/
func (s *AVOptionRange) ValueMin() float64 {
	value := s.ptr.value_min
	return float64(value)
}

// SetValueMin sets the value_min field.
/*
  Value range.
  For string ranges this represents the min/max length.
  For dimensions this represents the min/max pixel count or width/height in multi-component case.
*/
func (s *AVOptionRange) SetValueMin(value float64) {
	s.ptr.value_min = (C.double)(value)
}

// ValueMax gets the value_max field.
/*
  Value range.
  For string ranges this represents the min/max length.
  For dimensions this represents the min/max pixel count or width/height in multi-component case.
*/
func (s *AVOptionRange) ValueMax() float64 {
	value := s.ptr.value_max
	return float64(value)
}

// SetValueMax sets the value_max field.
/*
  Value range.
  For string ranges this represents the min/max length.
  For dimensions this represents the min/max pixel count or width/height in multi-component case.
*/
func (s *AVOptionRange) SetValueMax(value float64) {
	s.ptr.value_max = (C.double)(value)
}

// ComponentMin gets the component_min field.
/*
  Value's component range.
  For string this represents the unicode range for chars, 0-127 limits to ASCII.
*/
func (s *AVOptionRange) ComponentMin() float64 {
	value := s.ptr.component_min
	return float64(value)
}

// SetComponentMin sets the component_min field.
/*
  Value's component range.
  For string this represents the unicode range for chars, 0-127 limits to ASCII.
*/
func (s *AVOptionRange) SetComponentMin(value float64) {
	s.ptr.component_min = (C.double)(value)
}

// ComponentMax gets the component_max field.
/*
  Value's component range.
  For string this represents the unicode range for chars, 0-127 limits to ASCII.
*/
func (s *AVOptionRange) ComponentMax() float64 {
	value := s.ptr.component_max
	return float64(value)
}

// SetComponentMax sets the component_max field.
/*
  Value's component range.
  For string this represents the unicode range for chars, 0-127 limits to ASCII.
*/
func (s *AVOptionRange) SetComponentMax(value float64) {
	s.ptr.component_max = (C.double)(value)
}

// IsRange gets the is_range field.
/*
  Range flag.
  If set to 1 the struct encodes a range, if set to 0 a single value.
*/
func (s *AVOptionRange) IsRange() int {
	value := s.ptr.is_range
	return int(value)
}

// SetIsRange sets the is_range field.
/*
  Range flag.
  If set to 1 the struct encodes a range, if set to 0 a single value.
*/
func (s *AVOptionRange) SetIsRange(value int) {
	s.ptr.is_range = (C.int)(value)
}

// --- Struct AVOptionRanges ---

// AVOptionRanges wraps AVOptionRanges.
//
//	List of AVOptionRange structs.
type AVOptionRanges struct {
	ptr *C.AVOptionRanges
}

func (s *AVOptionRanges) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func ToAVOptionRangesArray(ptr unsafe.Pointer) *Array[*AVOptionRanges] {
	if ptr == nil {
		return nil
	}

	return &Array[*AVOptionRanges]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) *AVOptionRanges {
			ptr := (**C.AVOptionRanges)(pointer)
			value := *ptr
			var valueMapped *AVOptionRanges
			if value != nil {
				valueMapped = &AVOptionRanges{ptr: value}
			}
			return valueMapped
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value *AVOptionRanges) {
			ptr := (**C.AVOptionRanges)(pointer)
			if value != nil {
				*ptr = value.ptr
			} else {
				*ptr = nil
			}
		},
	}
}

// Range gets the range field.
/*
  Array of option ranges.

  Most of option types use just one component.
  Following describes multi-component option types:

  AV_OPT_TYPE_IMAGE_SIZE:
  component index 0: range of pixel count (width * height).
  component index 1: range of width.
  component index 2: range of height.

  @note To obtain multi-component version of this structure, user must
        provide AV_OPT_MULTI_COMPONENT_RANGE to av_opt_query_ranges or
        av_opt_query_ranges_default function.

  Multi-component range can be read as in following example:

  @code
  int range_index, component_index;
  AVOptionRanges *ranges;
  AVOptionRange *range[3]; //may require more than 3 in the future.
  av_opt_query_ranges(&ranges, obj, key, AV_OPT_MULTI_COMPONENT_RANGE);
  for (range_index = 0; range_index < ranges->nb_ranges; range_index++) {
      for (component_index = 0; component_index < ranges->nb_components; component_index++)
          range[component_index] = ranges->range[ranges->nb_ranges * component_index + range_index];
      //do something with range here.
  }
  av_opt_freep_ranges(&ranges);
  @endcode
*/
func (s *AVOptionRanges) Range() *Array[*AVOptionRange] {
	value := s.ptr._range
	return ToAVOptionRangeArray(unsafe.Pointer(value))
}

// SetRange sets the range field.
/*
  Array of option ranges.

  Most of option types use just one component.
  Following describes multi-component option types:

  AV_OPT_TYPE_IMAGE_SIZE:
  component index 0: range of pixel count (width * height).
  component index 1: range of width.
  component index 2: range of height.

  @note To obtain multi-component version of this structure, user must
        provide AV_OPT_MULTI_COMPONENT_RANGE to av_opt_query_ranges or
        av_opt_query_ranges_default function.

  Multi-component range can be read as in following example:

  @code
  int range_index, component_index;
  AVOptionRanges *ranges;
  AVOptionRange *range[3]; //may require more than 3 in the future.
  av_opt_query_ranges(&ranges, obj, key, AV_OPT_MULTI_COMPONENT_RANGE);
  for (range_index = 0; range_index < ranges->nb_ranges; range_index++) {
      for (component_index = 0; component_index < ranges->nb_components; component_index++)
          range[component_index] = ranges->range[ranges->nb_ranges * component_index + range_index];
      //do something with range here.
  }
  av_opt_freep_ranges(&ranges);
  @endcode
*/
func (s *AVOptionRanges) SetRange(value *Array[AVOptionRange]) {
	if value != nil {
		s.ptr._range = (**C.AVOptionRange)(value.ptr)
	} else {
		s.ptr._range = nil
	}
}

// NbRanges gets the nb_ranges field.
//
//	Number of ranges per component.
func (s *AVOptionRanges) NbRanges() int {
	value := s.ptr.nb_ranges
	return int(value)
}

// SetNbRanges sets the nb_ranges field.
//
//	Number of ranges per component.
func (s *AVOptionRanges) SetNbRanges(value int) {
	s.ptr.nb_ranges = (C.int)(value)
}

// NbComponents gets the nb_components field.
//
//	Number of componentes.
func (s *AVOptionRanges) NbComponents() int {
	value := s.ptr.nb_components
	return int(value)
}

// SetNbComponents sets the nb_components field.
//
//	Number of componentes.
func (s *AVOptionRanges) SetNbComponents(value int) {
	s.ptr.nb_components = (C.int)(value)
}

// --- Struct AVRational ---

// AVRational wraps AVRational.
//
//	Rational number (pair of numerator and denominator).
type AVRational struct {
	value C.AVRational
}

// Num gets the num field.
//
//	Numerator
func (s *AVRational) Num() int {
	value := s.value.num
	return int(value)
}

// SetNum sets the num field.
//
//	Numerator
func (s *AVRational) SetNum(value int) {
	s.value.num = (C.int)(value)
}

// Den gets the den field.
//
//	Denominator
func (s *AVRational) Den() int {
	value := s.value.den
	return int(value)
}

// SetDen sets the den field.
//
//	Denominator
func (s *AVRational) SetDen(value int) {
	s.value.den = (C.int)(value)
}
