package ffmpeg

import "unsafe"

// #include <libavcodec/avcodec.h>
// #include <libavcodec/codec.h>
// #include <libavcodec/codec_desc.h>
// #include <libavcodec/codec_id.h>
// #include <libavcodec/codec_par.h>
// #include <libavcodec/defs.h>
// #include <libavcodec/packet.h>
// #include <libavfilter/avfilter.h>
// #include <libavfilter/buffersink.h>
// #include <libavfilter/buffersrc.h>
// #include <libavformat/avformat.h>
// #include <libavformat/avio.h>
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
import "C"

// --- Struct RcOverride ---

// RcOverride wraps RcOverride.
type RcOverride struct {
	ptr *C.RcOverride
}

func (s *RcOverride) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func (s *RcOverride) StartFrame() int {
	value := s.ptr.start_frame
	return int(value)
}

func (s *RcOverride) SetStartFrame(value int) {
	s.ptr.start_frame = (C.int)(value)
}

func (s *RcOverride) EndFrame() int {
	value := s.ptr.end_frame
	return int(value)
}

func (s *RcOverride) SetEndFrame(value int) {
	s.ptr.end_frame = (C.int)(value)
}

func (s *RcOverride) Qscale() int {
	value := s.ptr.qscale
	return int(value)
}

func (s *RcOverride) SetQscale(value int) {
	s.ptr.qscale = (C.int)(value)
}

func (s *RcOverride) QualityFactor() float32 {
	value := s.ptr.quality_factor
	return float32(value)
}

func (s *RcOverride) SetQualityFactor(value float32) {
	s.ptr.quality_factor = (C.float)(value)
}

// --- Struct AVCodecInternal ---

// AVCodecInternal wraps AVCodecInternal.
type AVCodecInternal struct {
	ptr *C.struct_AVCodecInternal
}

func (s *AVCodecInternal) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
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

func (s *AVCodecContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

func (s *AVCodecContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

func (s *AVCodecContext) LogLevelOffset() int {
	value := s.ptr.log_level_offset
	return int(value)
}

func (s *AVCodecContext) SetLogLevelOffset(value int) {
	s.ptr.log_level_offset = (C.int)(value)
}

func (s *AVCodecContext) CodecType() AVMediaType {
	value := s.ptr.codec_type
	return AVMediaType(value)
}

func (s *AVCodecContext) SetCodecType(value AVMediaType) {
	s.ptr.codec_type = (C.enum_AVMediaType)(value)
}

func (s *AVCodecContext) Codec() *AVCodec {
	value := s.ptr.codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

func (s *AVCodecContext) SetCodec(value *AVCodec) {
	if value != nil {
		s.ptr.codec = value.ptr
	} else {
		s.ptr.codec = nil
	}
}

func (s *AVCodecContext) CodecId() AVCodecID {
	value := s.ptr.codec_id
	return AVCodecID(value)
}

func (s *AVCodecContext) SetCodecId(value AVCodecID) {
	s.ptr.codec_id = (C.enum_AVCodecID)(value)
}

func (s *AVCodecContext) CodecTag() uint {
	value := s.ptr.codec_tag
	return uint(value)
}

func (s *AVCodecContext) SetCodecTag(value uint) {
	s.ptr.codec_tag = (C.uint)(value)
}

func (s *AVCodecContext) PrivData() unsafe.Pointer {
	value := s.ptr.priv_data
	return value
}

func (s *AVCodecContext) SetPrivData(value unsafe.Pointer) {
	s.ptr.priv_data = value
}

func (s *AVCodecContext) Internal() *AVCodecInternal {
	value := s.ptr.internal
	var valueMapped *AVCodecInternal
	if value != nil {
		valueMapped = &AVCodecInternal{ptr: value}
	}
	return valueMapped
}

func (s *AVCodecContext) SetInternal(value *AVCodecInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

func (s *AVCodecContext) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

func (s *AVCodecContext) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

func (s *AVCodecContext) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}

func (s *AVCodecContext) SetBitRate(value int64) {
	s.ptr.bit_rate = (C.int64_t)(value)
}

func (s *AVCodecContext) BitRateTolerance() int {
	value := s.ptr.bit_rate_tolerance
	return int(value)
}

func (s *AVCodecContext) SetBitRateTolerance(value int) {
	s.ptr.bit_rate_tolerance = (C.int)(value)
}

func (s *AVCodecContext) GlobalQuality() int {
	value := s.ptr.global_quality
	return int(value)
}

func (s *AVCodecContext) SetGlobalQuality(value int) {
	s.ptr.global_quality = (C.int)(value)
}

func (s *AVCodecContext) CompressionLevel() int {
	value := s.ptr.compression_level
	return int(value)
}

func (s *AVCodecContext) SetCompressionLevel(value int) {
	s.ptr.compression_level = (C.int)(value)
}

func (s *AVCodecContext) Flags() int {
	value := s.ptr.flags
	return int(value)
}

func (s *AVCodecContext) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

func (s *AVCodecContext) Flags2() int {
	value := s.ptr.flags2
	return int(value)
}

func (s *AVCodecContext) SetFlags2(value int) {
	s.ptr.flags2 = (C.int)(value)
}

func (s *AVCodecContext) Extradata() unsafe.Pointer {
	value := s.ptr.extradata
	return unsafe.Pointer(value)
}

func (s *AVCodecContext) SetExtradata(value unsafe.Pointer) {
	s.ptr.extradata = (*C.uint8_t)(value)
}

func (s *AVCodecContext) ExtradataSize() int {
	value := s.ptr.extradata_size
	return int(value)
}

func (s *AVCodecContext) SetExtradataSize(value int) {
	s.ptr.extradata_size = (C.int)(value)
}

func (s *AVCodecContext) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

func (s *AVCodecContext) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

func (s *AVCodecContext) TicksPerFrame() int {
	value := s.ptr.ticks_per_frame
	return int(value)
}

func (s *AVCodecContext) SetTicksPerFrame(value int) {
	s.ptr.ticks_per_frame = (C.int)(value)
}

func (s *AVCodecContext) Delay() int {
	value := s.ptr.delay
	return int(value)
}

func (s *AVCodecContext) SetDelay(value int) {
	s.ptr.delay = (C.int)(value)
}

func (s *AVCodecContext) Width() int {
	value := s.ptr.width
	return int(value)
}

func (s *AVCodecContext) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

func (s *AVCodecContext) Height() int {
	value := s.ptr.height
	return int(value)
}

func (s *AVCodecContext) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

func (s *AVCodecContext) CodedWidth() int {
	value := s.ptr.coded_width
	return int(value)
}

func (s *AVCodecContext) SetCodedWidth(value int) {
	s.ptr.coded_width = (C.int)(value)
}

func (s *AVCodecContext) CodedHeight() int {
	value := s.ptr.coded_height
	return int(value)
}

func (s *AVCodecContext) SetCodedHeight(value int) {
	s.ptr.coded_height = (C.int)(value)
}

func (s *AVCodecContext) GopSize() int {
	value := s.ptr.gop_size
	return int(value)
}

func (s *AVCodecContext) SetGopSize(value int) {
	s.ptr.gop_size = (C.int)(value)
}

func (s *AVCodecContext) PixFmt() AVPixelFormat {
	value := s.ptr.pix_fmt
	return AVPixelFormat(value)
}

func (s *AVCodecContext) SetPixFmt(value AVPixelFormat) {
	s.ptr.pix_fmt = (C.enum_AVPixelFormat)(value)
}

// DrawHorizBand skipped due to func ptr

// GetFormat skipped due to func ptr

func (s *AVCodecContext) MaxBFrames() int {
	value := s.ptr.max_b_frames
	return int(value)
}

func (s *AVCodecContext) SetMaxBFrames(value int) {
	s.ptr.max_b_frames = (C.int)(value)
}

func (s *AVCodecContext) BQuantFactor() float32 {
	value := s.ptr.b_quant_factor
	return float32(value)
}

func (s *AVCodecContext) SetBQuantFactor(value float32) {
	s.ptr.b_quant_factor = (C.float)(value)
}

func (s *AVCodecContext) BQuantOffset() float32 {
	value := s.ptr.b_quant_offset
	return float32(value)
}

func (s *AVCodecContext) SetBQuantOffset(value float32) {
	s.ptr.b_quant_offset = (C.float)(value)
}

func (s *AVCodecContext) HasBFrames() int {
	value := s.ptr.has_b_frames
	return int(value)
}

func (s *AVCodecContext) SetHasBFrames(value int) {
	s.ptr.has_b_frames = (C.int)(value)
}

func (s *AVCodecContext) IQuantFactor() float32 {
	value := s.ptr.i_quant_factor
	return float32(value)
}

func (s *AVCodecContext) SetIQuantFactor(value float32) {
	s.ptr.i_quant_factor = (C.float)(value)
}

func (s *AVCodecContext) IQuantOffset() float32 {
	value := s.ptr.i_quant_offset
	return float32(value)
}

func (s *AVCodecContext) SetIQuantOffset(value float32) {
	s.ptr.i_quant_offset = (C.float)(value)
}

func (s *AVCodecContext) LumiMasking() float32 {
	value := s.ptr.lumi_masking
	return float32(value)
}

func (s *AVCodecContext) SetLumiMasking(value float32) {
	s.ptr.lumi_masking = (C.float)(value)
}

func (s *AVCodecContext) TemporalCplxMasking() float32 {
	value := s.ptr.temporal_cplx_masking
	return float32(value)
}

func (s *AVCodecContext) SetTemporalCplxMasking(value float32) {
	s.ptr.temporal_cplx_masking = (C.float)(value)
}

func (s *AVCodecContext) SpatialCplxMasking() float32 {
	value := s.ptr.spatial_cplx_masking
	return float32(value)
}

func (s *AVCodecContext) SetSpatialCplxMasking(value float32) {
	s.ptr.spatial_cplx_masking = (C.float)(value)
}

func (s *AVCodecContext) PMasking() float32 {
	value := s.ptr.p_masking
	return float32(value)
}

func (s *AVCodecContext) SetPMasking(value float32) {
	s.ptr.p_masking = (C.float)(value)
}

func (s *AVCodecContext) DarkMasking() float32 {
	value := s.ptr.dark_masking
	return float32(value)
}

func (s *AVCodecContext) SetDarkMasking(value float32) {
	s.ptr.dark_masking = (C.float)(value)
}

func (s *AVCodecContext) SliceCount() int {
	value := s.ptr.slice_count
	return int(value)
}

func (s *AVCodecContext) SetSliceCount(value int) {
	s.ptr.slice_count = (C.int)(value)
}

// SliceOffset skipped due to prim ptr

func (s *AVCodecContext) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

func (s *AVCodecContext) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

func (s *AVCodecContext) MeCmp() int {
	value := s.ptr.me_cmp
	return int(value)
}

func (s *AVCodecContext) SetMeCmp(value int) {
	s.ptr.me_cmp = (C.int)(value)
}

func (s *AVCodecContext) MeSubCmp() int {
	value := s.ptr.me_sub_cmp
	return int(value)
}

func (s *AVCodecContext) SetMeSubCmp(value int) {
	s.ptr.me_sub_cmp = (C.int)(value)
}

func (s *AVCodecContext) MbCmp() int {
	value := s.ptr.mb_cmp
	return int(value)
}

func (s *AVCodecContext) SetMbCmp(value int) {
	s.ptr.mb_cmp = (C.int)(value)
}

func (s *AVCodecContext) IldctCmp() int {
	value := s.ptr.ildct_cmp
	return int(value)
}

func (s *AVCodecContext) SetIldctCmp(value int) {
	s.ptr.ildct_cmp = (C.int)(value)
}

func (s *AVCodecContext) DiaSize() int {
	value := s.ptr.dia_size
	return int(value)
}

func (s *AVCodecContext) SetDiaSize(value int) {
	s.ptr.dia_size = (C.int)(value)
}

func (s *AVCodecContext) LastPredictorCount() int {
	value := s.ptr.last_predictor_count
	return int(value)
}

func (s *AVCodecContext) SetLastPredictorCount(value int) {
	s.ptr.last_predictor_count = (C.int)(value)
}

func (s *AVCodecContext) MePreCmp() int {
	value := s.ptr.me_pre_cmp
	return int(value)
}

func (s *AVCodecContext) SetMePreCmp(value int) {
	s.ptr.me_pre_cmp = (C.int)(value)
}

func (s *AVCodecContext) PreDiaSize() int {
	value := s.ptr.pre_dia_size
	return int(value)
}

func (s *AVCodecContext) SetPreDiaSize(value int) {
	s.ptr.pre_dia_size = (C.int)(value)
}

func (s *AVCodecContext) MeSubpelQuality() int {
	value := s.ptr.me_subpel_quality
	return int(value)
}

func (s *AVCodecContext) SetMeSubpelQuality(value int) {
	s.ptr.me_subpel_quality = (C.int)(value)
}

func (s *AVCodecContext) MeRange() int {
	value := s.ptr.me_range
	return int(value)
}

func (s *AVCodecContext) SetMeRange(value int) {
	s.ptr.me_range = (C.int)(value)
}

func (s *AVCodecContext) SliceFlags() int {
	value := s.ptr.slice_flags
	return int(value)
}

func (s *AVCodecContext) SetSliceFlags(value int) {
	s.ptr.slice_flags = (C.int)(value)
}

func (s *AVCodecContext) MbDecision() int {
	value := s.ptr.mb_decision
	return int(value)
}

func (s *AVCodecContext) SetMbDecision(value int) {
	s.ptr.mb_decision = (C.int)(value)
}

// IntraMatrix skipped due to prim ptr

// InterMatrix skipped due to prim ptr

func (s *AVCodecContext) IntraDcPrecision() int {
	value := s.ptr.intra_dc_precision
	return int(value)
}

func (s *AVCodecContext) SetIntraDcPrecision(value int) {
	s.ptr.intra_dc_precision = (C.int)(value)
}

func (s *AVCodecContext) SkipTop() int {
	value := s.ptr.skip_top
	return int(value)
}

func (s *AVCodecContext) SetSkipTop(value int) {
	s.ptr.skip_top = (C.int)(value)
}

func (s *AVCodecContext) SkipBottom() int {
	value := s.ptr.skip_bottom
	return int(value)
}

func (s *AVCodecContext) SetSkipBottom(value int) {
	s.ptr.skip_bottom = (C.int)(value)
}

func (s *AVCodecContext) MbLmin() int {
	value := s.ptr.mb_lmin
	return int(value)
}

func (s *AVCodecContext) SetMbLmin(value int) {
	s.ptr.mb_lmin = (C.int)(value)
}

func (s *AVCodecContext) MbLmax() int {
	value := s.ptr.mb_lmax
	return int(value)
}

func (s *AVCodecContext) SetMbLmax(value int) {
	s.ptr.mb_lmax = (C.int)(value)
}

func (s *AVCodecContext) BidirRefine() int {
	value := s.ptr.bidir_refine
	return int(value)
}

func (s *AVCodecContext) SetBidirRefine(value int) {
	s.ptr.bidir_refine = (C.int)(value)
}

func (s *AVCodecContext) KeyintMin() int {
	value := s.ptr.keyint_min
	return int(value)
}

func (s *AVCodecContext) SetKeyintMin(value int) {
	s.ptr.keyint_min = (C.int)(value)
}

func (s *AVCodecContext) Refs() int {
	value := s.ptr.refs
	return int(value)
}

func (s *AVCodecContext) SetRefs(value int) {
	s.ptr.refs = (C.int)(value)
}

func (s *AVCodecContext) Mv0Threshold() int {
	value := s.ptr.mv0_threshold
	return int(value)
}

func (s *AVCodecContext) SetMv0Threshold(value int) {
	s.ptr.mv0_threshold = (C.int)(value)
}

func (s *AVCodecContext) ColorPrimaries() AVColorPrimaries {
	value := s.ptr.color_primaries
	return AVColorPrimaries(value)
}

func (s *AVCodecContext) SetColorPrimaries(value AVColorPrimaries) {
	s.ptr.color_primaries = (C.enum_AVColorPrimaries)(value)
}

func (s *AVCodecContext) ColorTrc() AVColorTransferCharacteristic {
	value := s.ptr.color_trc
	return AVColorTransferCharacteristic(value)
}

func (s *AVCodecContext) SetColorTrc(value AVColorTransferCharacteristic) {
	s.ptr.color_trc = (C.enum_AVColorTransferCharacteristic)(value)
}

func (s *AVCodecContext) Colorspace() AVColorSpace {
	value := s.ptr.colorspace
	return AVColorSpace(value)
}

func (s *AVCodecContext) SetColorspace(value AVColorSpace) {
	s.ptr.colorspace = (C.enum_AVColorSpace)(value)
}

func (s *AVCodecContext) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}

func (s *AVCodecContext) SetColorRange(value AVColorRange) {
	s.ptr.color_range = (C.enum_AVColorRange)(value)
}

func (s *AVCodecContext) ChromaSampleLocation() AVChromaLocation {
	value := s.ptr.chroma_sample_location
	return AVChromaLocation(value)
}

func (s *AVCodecContext) SetChromaSampleLocation(value AVChromaLocation) {
	s.ptr.chroma_sample_location = (C.enum_AVChromaLocation)(value)
}

func (s *AVCodecContext) Slices() int {
	value := s.ptr.slices
	return int(value)
}

func (s *AVCodecContext) SetSlices(value int) {
	s.ptr.slices = (C.int)(value)
}

func (s *AVCodecContext) FieldOrder() AVFieldOrder {
	value := s.ptr.field_order
	return AVFieldOrder(value)
}

func (s *AVCodecContext) SetFieldOrder(value AVFieldOrder) {
	s.ptr.field_order = (C.enum_AVFieldOrder)(value)
}

func (s *AVCodecContext) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

func (s *AVCodecContext) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

func (s *AVCodecContext) Channels() int {
	value := s.ptr.channels
	return int(value)
}

func (s *AVCodecContext) SetChannels(value int) {
	s.ptr.channels = (C.int)(value)
}

func (s *AVCodecContext) SampleFmt() AVSampleFormat {
	value := s.ptr.sample_fmt
	return AVSampleFormat(value)
}

func (s *AVCodecContext) SetSampleFmt(value AVSampleFormat) {
	s.ptr.sample_fmt = (C.enum_AVSampleFormat)(value)
}

func (s *AVCodecContext) FrameSize() int {
	value := s.ptr.frame_size
	return int(value)
}

func (s *AVCodecContext) SetFrameSize(value int) {
	s.ptr.frame_size = (C.int)(value)
}

func (s *AVCodecContext) FrameNumber() int {
	value := s.ptr.frame_number
	return int(value)
}

func (s *AVCodecContext) SetFrameNumber(value int) {
	s.ptr.frame_number = (C.int)(value)
}

func (s *AVCodecContext) BlockAlign() int {
	value := s.ptr.block_align
	return int(value)
}

func (s *AVCodecContext) SetBlockAlign(value int) {
	s.ptr.block_align = (C.int)(value)
}

func (s *AVCodecContext) Cutoff() int {
	value := s.ptr.cutoff
	return int(value)
}

func (s *AVCodecContext) SetCutoff(value int) {
	s.ptr.cutoff = (C.int)(value)
}

func (s *AVCodecContext) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

func (s *AVCodecContext) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

func (s *AVCodecContext) RequestChannelLayout() uint64 {
	value := s.ptr.request_channel_layout
	return uint64(value)
}

func (s *AVCodecContext) SetRequestChannelLayout(value uint64) {
	s.ptr.request_channel_layout = (C.uint64_t)(value)
}

func (s *AVCodecContext) AudioServiceType() AVAudioServiceType {
	value := s.ptr.audio_service_type
	return AVAudioServiceType(value)
}

func (s *AVCodecContext) SetAudioServiceType(value AVAudioServiceType) {
	s.ptr.audio_service_type = (C.enum_AVAudioServiceType)(value)
}

func (s *AVCodecContext) RequestSampleFmt() AVSampleFormat {
	value := s.ptr.request_sample_fmt
	return AVSampleFormat(value)
}

func (s *AVCodecContext) SetRequestSampleFmt(value AVSampleFormat) {
	s.ptr.request_sample_fmt = (C.enum_AVSampleFormat)(value)
}

// GetBuffer2 skipped due to func ptr

func (s *AVCodecContext) Qcompress() float32 {
	value := s.ptr.qcompress
	return float32(value)
}

func (s *AVCodecContext) SetQcompress(value float32) {
	s.ptr.qcompress = (C.float)(value)
}

func (s *AVCodecContext) Qblur() float32 {
	value := s.ptr.qblur
	return float32(value)
}

func (s *AVCodecContext) SetQblur(value float32) {
	s.ptr.qblur = (C.float)(value)
}

func (s *AVCodecContext) Qmin() int {
	value := s.ptr.qmin
	return int(value)
}

func (s *AVCodecContext) SetQmin(value int) {
	s.ptr.qmin = (C.int)(value)
}

func (s *AVCodecContext) Qmax() int {
	value := s.ptr.qmax
	return int(value)
}

func (s *AVCodecContext) SetQmax(value int) {
	s.ptr.qmax = (C.int)(value)
}

func (s *AVCodecContext) MaxQdiff() int {
	value := s.ptr.max_qdiff
	return int(value)
}

func (s *AVCodecContext) SetMaxQdiff(value int) {
	s.ptr.max_qdiff = (C.int)(value)
}

func (s *AVCodecContext) RcBufferSize() int {
	value := s.ptr.rc_buffer_size
	return int(value)
}

func (s *AVCodecContext) SetRcBufferSize(value int) {
	s.ptr.rc_buffer_size = (C.int)(value)
}

func (s *AVCodecContext) RcOverrideCount() int {
	value := s.ptr.rc_override_count
	return int(value)
}

func (s *AVCodecContext) SetRcOverrideCount(value int) {
	s.ptr.rc_override_count = (C.int)(value)
}

func (s *AVCodecContext) RcOverride() *RcOverride {
	value := s.ptr.rc_override
	var valueMapped *RcOverride
	if value != nil {
		valueMapped = &RcOverride{ptr: value}
	}
	return valueMapped
}

func (s *AVCodecContext) SetRcOverride(value *RcOverride) {
	if value != nil {
		s.ptr.rc_override = value.ptr
	} else {
		s.ptr.rc_override = nil
	}
}

func (s *AVCodecContext) RcMaxRate() int64 {
	value := s.ptr.rc_max_rate
	return int64(value)
}

func (s *AVCodecContext) SetRcMaxRate(value int64) {
	s.ptr.rc_max_rate = (C.int64_t)(value)
}

func (s *AVCodecContext) RcMinRate() int64 {
	value := s.ptr.rc_min_rate
	return int64(value)
}

func (s *AVCodecContext) SetRcMinRate(value int64) {
	s.ptr.rc_min_rate = (C.int64_t)(value)
}

func (s *AVCodecContext) RcMaxAvailableVbvUse() float32 {
	value := s.ptr.rc_max_available_vbv_use
	return float32(value)
}

func (s *AVCodecContext) SetRcMaxAvailableVbvUse(value float32) {
	s.ptr.rc_max_available_vbv_use = (C.float)(value)
}

func (s *AVCodecContext) RcMinVbvOverflowUse() float32 {
	value := s.ptr.rc_min_vbv_overflow_use
	return float32(value)
}

func (s *AVCodecContext) SetRcMinVbvOverflowUse(value float32) {
	s.ptr.rc_min_vbv_overflow_use = (C.float)(value)
}

func (s *AVCodecContext) RcInitialBufferOccupancy() int {
	value := s.ptr.rc_initial_buffer_occupancy
	return int(value)
}

func (s *AVCodecContext) SetRcInitialBufferOccupancy(value int) {
	s.ptr.rc_initial_buffer_occupancy = (C.int)(value)
}

func (s *AVCodecContext) Trellis() int {
	value := s.ptr.trellis
	return int(value)
}

func (s *AVCodecContext) SetTrellis(value int) {
	s.ptr.trellis = (C.int)(value)
}

func (s *AVCodecContext) StatsOut() *CStr {
	value := s.ptr.stats_out
	return wrapCStr(value)
}

func (s *AVCodecContext) SetStatsOut(value *CStr) {
	s.ptr.stats_out = value.ptr
}

func (s *AVCodecContext) StatsIn() *CStr {
	value := s.ptr.stats_in
	return wrapCStr(value)
}

func (s *AVCodecContext) SetStatsIn(value *CStr) {
	s.ptr.stats_in = value.ptr
}

func (s *AVCodecContext) WorkaroundBugs() int {
	value := s.ptr.workaround_bugs
	return int(value)
}

func (s *AVCodecContext) SetWorkaroundBugs(value int) {
	s.ptr.workaround_bugs = (C.int)(value)
}

func (s *AVCodecContext) StrictStdCompliance() int {
	value := s.ptr.strict_std_compliance
	return int(value)
}

func (s *AVCodecContext) SetStrictStdCompliance(value int) {
	s.ptr.strict_std_compliance = (C.int)(value)
}

func (s *AVCodecContext) ErrorConcealment() int {
	value := s.ptr.error_concealment
	return int(value)
}

func (s *AVCodecContext) SetErrorConcealment(value int) {
	s.ptr.error_concealment = (C.int)(value)
}

func (s *AVCodecContext) Debug() int {
	value := s.ptr.debug
	return int(value)
}

func (s *AVCodecContext) SetDebug(value int) {
	s.ptr.debug = (C.int)(value)
}

func (s *AVCodecContext) ErrRecognition() int {
	value := s.ptr.err_recognition
	return int(value)
}

func (s *AVCodecContext) SetErrRecognition(value int) {
	s.ptr.err_recognition = (C.int)(value)
}

func (s *AVCodecContext) ReorderedOpaque() int64 {
	value := s.ptr.reordered_opaque
	return int64(value)
}

func (s *AVCodecContext) SetReorderedOpaque(value int64) {
	s.ptr.reordered_opaque = (C.int64_t)(value)
}

func (s *AVCodecContext) Hwaccel() *AVHWAccel {
	value := s.ptr.hwaccel
	var valueMapped *AVHWAccel
	if value != nil {
		valueMapped = &AVHWAccel{ptr: value}
	}
	return valueMapped
}

func (s *AVCodecContext) SetHwaccel(value *AVHWAccel) {
	if value != nil {
		s.ptr.hwaccel = value.ptr
	} else {
		s.ptr.hwaccel = nil
	}
}

func (s *AVCodecContext) HwaccelContext() unsafe.Pointer {
	value := s.ptr.hwaccel_context
	return value
}

func (s *AVCodecContext) SetHwaccelContext(value unsafe.Pointer) {
	s.ptr.hwaccel_context = value
}

func (s *AVCodecContext) ErrorEntry(i uint) uint64 {
	value := s.ptr.error[i]
	return uint64(value)
}

func (s *AVCodecContext) SetErrorEntry(i uint, value uint64) {
	s.ptr.error[i] = (C.uint64_t)(value)
}

func (s *AVCodecContext) DctAlgo() int {
	value := s.ptr.dct_algo
	return int(value)
}

func (s *AVCodecContext) SetDctAlgo(value int) {
	s.ptr.dct_algo = (C.int)(value)
}

func (s *AVCodecContext) IdctAlgo() int {
	value := s.ptr.idct_algo
	return int(value)
}

func (s *AVCodecContext) SetIdctAlgo(value int) {
	s.ptr.idct_algo = (C.int)(value)
}

func (s *AVCodecContext) BitsPerCodedSample() int {
	value := s.ptr.bits_per_coded_sample
	return int(value)
}

func (s *AVCodecContext) SetBitsPerCodedSample(value int) {
	s.ptr.bits_per_coded_sample = (C.int)(value)
}

func (s *AVCodecContext) BitsPerRawSample() int {
	value := s.ptr.bits_per_raw_sample
	return int(value)
}

func (s *AVCodecContext) SetBitsPerRawSample(value int) {
	s.ptr.bits_per_raw_sample = (C.int)(value)
}

func (s *AVCodecContext) Lowres() int {
	value := s.ptr.lowres
	return int(value)
}

func (s *AVCodecContext) SetLowres(value int) {
	s.ptr.lowres = (C.int)(value)
}

func (s *AVCodecContext) ThreadCount() int {
	value := s.ptr.thread_count
	return int(value)
}

func (s *AVCodecContext) SetThreadCount(value int) {
	s.ptr.thread_count = (C.int)(value)
}

func (s *AVCodecContext) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}

func (s *AVCodecContext) SetThreadType(value int) {
	s.ptr.thread_type = (C.int)(value)
}

func (s *AVCodecContext) ActiveThreadType() int {
	value := s.ptr.active_thread_type
	return int(value)
}

func (s *AVCodecContext) SetActiveThreadType(value int) {
	s.ptr.active_thread_type = (C.int)(value)
}

// Execute skipped due to func ptr

// Execute2 skipped due to func ptr

func (s *AVCodecContext) NsseWeight() int {
	value := s.ptr.nsse_weight
	return int(value)
}

func (s *AVCodecContext) SetNsseWeight(value int) {
	s.ptr.nsse_weight = (C.int)(value)
}

func (s *AVCodecContext) Profile() int {
	value := s.ptr.profile
	return int(value)
}

func (s *AVCodecContext) SetProfile(value int) {
	s.ptr.profile = (C.int)(value)
}

func (s *AVCodecContext) Level() int {
	value := s.ptr.level
	return int(value)
}

func (s *AVCodecContext) SetLevel(value int) {
	s.ptr.level = (C.int)(value)
}

func (s *AVCodecContext) SkipLoopFilter() AVDiscard {
	value := s.ptr.skip_loop_filter
	return AVDiscard(value)
}

func (s *AVCodecContext) SetSkipLoopFilter(value AVDiscard) {
	s.ptr.skip_loop_filter = (C.enum_AVDiscard)(value)
}

func (s *AVCodecContext) SkipIdct() AVDiscard {
	value := s.ptr.skip_idct
	return AVDiscard(value)
}

func (s *AVCodecContext) SetSkipIdct(value AVDiscard) {
	s.ptr.skip_idct = (C.enum_AVDiscard)(value)
}

func (s *AVCodecContext) SkipFrame() AVDiscard {
	value := s.ptr.skip_frame
	return AVDiscard(value)
}

func (s *AVCodecContext) SetSkipFrame(value AVDiscard) {
	s.ptr.skip_frame = (C.enum_AVDiscard)(value)
}

func (s *AVCodecContext) SubtitleHeader() unsafe.Pointer {
	value := s.ptr.subtitle_header
	return unsafe.Pointer(value)
}

func (s *AVCodecContext) SetSubtitleHeader(value unsafe.Pointer) {
	s.ptr.subtitle_header = (*C.uint8_t)(value)
}

func (s *AVCodecContext) SubtitleHeaderSize() int {
	value := s.ptr.subtitle_header_size
	return int(value)
}

func (s *AVCodecContext) SetSubtitleHeaderSize(value int) {
	s.ptr.subtitle_header_size = (C.int)(value)
}

func (s *AVCodecContext) InitialPadding() int {
	value := s.ptr.initial_padding
	return int(value)
}

func (s *AVCodecContext) SetInitialPadding(value int) {
	s.ptr.initial_padding = (C.int)(value)
}

func (s *AVCodecContext) Framerate() *AVRational {
	value := s.ptr.framerate
	return &AVRational{value: value}
}

func (s *AVCodecContext) SetFramerate(value *AVRational) {
	s.ptr.framerate = value.value
}

func (s *AVCodecContext) SwPixFmt() AVPixelFormat {
	value := s.ptr.sw_pix_fmt
	return AVPixelFormat(value)
}

func (s *AVCodecContext) SetSwPixFmt(value AVPixelFormat) {
	s.ptr.sw_pix_fmt = (C.enum_AVPixelFormat)(value)
}

func (s *AVCodecContext) PktTimebase() *AVRational {
	value := s.ptr.pkt_timebase
	return &AVRational{value: value}
}

func (s *AVCodecContext) SetPktTimebase(value *AVRational) {
	s.ptr.pkt_timebase = value.value
}

func (s *AVCodecContext) CodecDescriptor() *AVCodecDescriptor {
	value := s.ptr.codec_descriptor
	var valueMapped *AVCodecDescriptor
	if value != nil {
		valueMapped = &AVCodecDescriptor{ptr: value}
	}
	return valueMapped
}

func (s *AVCodecContext) SetCodecDescriptor(value *AVCodecDescriptor) {
	if value != nil {
		s.ptr.codec_descriptor = value.ptr
	} else {
		s.ptr.codec_descriptor = nil
	}
}

func (s *AVCodecContext) PtsCorrectionNumFaultyPts() int64 {
	value := s.ptr.pts_correction_num_faulty_pts
	return int64(value)
}

func (s *AVCodecContext) SetPtsCorrectionNumFaultyPts(value int64) {
	s.ptr.pts_correction_num_faulty_pts = (C.int64_t)(value)
}

func (s *AVCodecContext) PtsCorrectionNumFaultyDts() int64 {
	value := s.ptr.pts_correction_num_faulty_dts
	return int64(value)
}

func (s *AVCodecContext) SetPtsCorrectionNumFaultyDts(value int64) {
	s.ptr.pts_correction_num_faulty_dts = (C.int64_t)(value)
}

func (s *AVCodecContext) PtsCorrectionLastPts() int64 {
	value := s.ptr.pts_correction_last_pts
	return int64(value)
}

func (s *AVCodecContext) SetPtsCorrectionLastPts(value int64) {
	s.ptr.pts_correction_last_pts = (C.int64_t)(value)
}

func (s *AVCodecContext) PtsCorrectionLastDts() int64 {
	value := s.ptr.pts_correction_last_dts
	return int64(value)
}

func (s *AVCodecContext) SetPtsCorrectionLastDts(value int64) {
	s.ptr.pts_correction_last_dts = (C.int64_t)(value)
}

func (s *AVCodecContext) SubCharenc() *CStr {
	value := s.ptr.sub_charenc
	return wrapCStr(value)
}

func (s *AVCodecContext) SetSubCharenc(value *CStr) {
	s.ptr.sub_charenc = value.ptr
}

func (s *AVCodecContext) SubCharencMode() int {
	value := s.ptr.sub_charenc_mode
	return int(value)
}

func (s *AVCodecContext) SetSubCharencMode(value int) {
	s.ptr.sub_charenc_mode = (C.int)(value)
}

func (s *AVCodecContext) SkipAlpha() int {
	value := s.ptr.skip_alpha
	return int(value)
}

func (s *AVCodecContext) SetSkipAlpha(value int) {
	s.ptr.skip_alpha = (C.int)(value)
}

func (s *AVCodecContext) SeekPreroll() int {
	value := s.ptr.seek_preroll
	return int(value)
}

func (s *AVCodecContext) SetSeekPreroll(value int) {
	s.ptr.seek_preroll = (C.int)(value)
}

// ChromaIntraMatrix skipped due to prim ptr

func (s *AVCodecContext) DumpSeparator() unsafe.Pointer {
	value := s.ptr.dump_separator
	return unsafe.Pointer(value)
}

func (s *AVCodecContext) SetDumpSeparator(value unsafe.Pointer) {
	s.ptr.dump_separator = (*C.uint8_t)(value)
}

func (s *AVCodecContext) CodecWhitelist() *CStr {
	value := s.ptr.codec_whitelist
	return wrapCStr(value)
}

func (s *AVCodecContext) SetCodecWhitelist(value *CStr) {
	s.ptr.codec_whitelist = value.ptr
}

func (s *AVCodecContext) Properties() uint {
	value := s.ptr.properties
	return uint(value)
}

func (s *AVCodecContext) SetProperties(value uint) {
	s.ptr.properties = (C.uint)(value)
}

func (s *AVCodecContext) CodedSideData() *AVPacketSideData {
	value := s.ptr.coded_side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}

func (s *AVCodecContext) SetCodedSideData(value *AVPacketSideData) {
	if value != nil {
		s.ptr.coded_side_data = value.ptr
	} else {
		s.ptr.coded_side_data = nil
	}
}

func (s *AVCodecContext) NbCodedSideData() int {
	value := s.ptr.nb_coded_side_data
	return int(value)
}

func (s *AVCodecContext) SetNbCodedSideData(value int) {
	s.ptr.nb_coded_side_data = (C.int)(value)
}

func (s *AVCodecContext) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVCodecContext) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

func (s *AVCodecContext) TrailingPadding() int {
	value := s.ptr.trailing_padding
	return int(value)
}

func (s *AVCodecContext) SetTrailingPadding(value int) {
	s.ptr.trailing_padding = (C.int)(value)
}

func (s *AVCodecContext) MaxPixels() int64 {
	value := s.ptr.max_pixels
	return int64(value)
}

func (s *AVCodecContext) SetMaxPixels(value int64) {
	s.ptr.max_pixels = (C.int64_t)(value)
}

func (s *AVCodecContext) HwDeviceCtx() *AVBufferRef {
	value := s.ptr.hw_device_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVCodecContext) SetHwDeviceCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_device_ctx = value.ptr
	} else {
		s.ptr.hw_device_ctx = nil
	}
}

func (s *AVCodecContext) HwaccelFlags() int {
	value := s.ptr.hwaccel_flags
	return int(value)
}

func (s *AVCodecContext) SetHwaccelFlags(value int) {
	s.ptr.hwaccel_flags = (C.int)(value)
}

func (s *AVCodecContext) ApplyCropping() int {
	value := s.ptr.apply_cropping
	return int(value)
}

func (s *AVCodecContext) SetApplyCropping(value int) {
	s.ptr.apply_cropping = (C.int)(value)
}

func (s *AVCodecContext) ExtraHwFrames() int {
	value := s.ptr.extra_hw_frames
	return int(value)
}

func (s *AVCodecContext) SetExtraHwFrames(value int) {
	s.ptr.extra_hw_frames = (C.int)(value)
}

func (s *AVCodecContext) DiscardDamagedPercentage() int {
	value := s.ptr.discard_damaged_percentage
	return int(value)
}

func (s *AVCodecContext) SetDiscardDamagedPercentage(value int) {
	s.ptr.discard_damaged_percentage = (C.int)(value)
}

func (s *AVCodecContext) MaxSamples() int64 {
	value := s.ptr.max_samples
	return int64(value)
}

func (s *AVCodecContext) SetMaxSamples(value int64) {
	s.ptr.max_samples = (C.int64_t)(value)
}

func (s *AVCodecContext) ExportSideData() int {
	value := s.ptr.export_side_data
	return int(value)
}

func (s *AVCodecContext) SetExportSideData(value int) {
	s.ptr.export_side_data = (C.int)(value)
}

// GetEncodeBuffer skipped due to func ptr

// ChLayout skipped due to ident byvalue

func (s *AVCodecContext) FrameNum() int64 {
	value := s.ptr.frame_num
	return int64(value)
}

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

func (s *AVHWAccel) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

func (s *AVHWAccel) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

func (s *AVHWAccel) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}

func (s *AVHWAccel) SetType(value AVMediaType) {
	s.ptr._type = (C.enum_AVMediaType)(value)
}

func (s *AVHWAccel) Id() AVCodecID {
	value := s.ptr.id
	return AVCodecID(value)
}

func (s *AVHWAccel) SetId(value AVCodecID) {
	s.ptr.id = (C.enum_AVCodecID)(value)
}

func (s *AVHWAccel) PixFmt() AVPixelFormat {
	value := s.ptr.pix_fmt
	return AVPixelFormat(value)
}

func (s *AVHWAccel) SetPixFmt(value AVPixelFormat) {
	s.ptr.pix_fmt = (C.enum_AVPixelFormat)(value)
}

func (s *AVHWAccel) Capabilities() int {
	value := s.ptr.capabilities
	return int(value)
}

func (s *AVHWAccel) SetCapabilities(value int) {
	s.ptr.capabilities = (C.int)(value)
}

// AllocFrame skipped due to func ptr

// StartFrame skipped due to func ptr

// DecodeParams skipped due to func ptr

// DecodeSlice skipped due to func ptr

// EndFrame skipped due to func ptr

func (s *AVHWAccel) FramePrivDataSize() int {
	value := s.ptr.frame_priv_data_size
	return int(value)
}

func (s *AVHWAccel) SetFramePrivDataSize(value int) {
	s.ptr.frame_priv_data_size = (C.int)(value)
}

// Init skipped due to func ptr

// Uninit skipped due to func ptr

func (s *AVHWAccel) PrivDataSize() int {
	value := s.ptr.priv_data_size
	return int(value)
}

func (s *AVHWAccel) SetPrivDataSize(value int) {
	s.ptr.priv_data_size = (C.int)(value)
}

func (s *AVHWAccel) CapsInternal() int {
	value := s.ptr.caps_internal
	return int(value)
}

func (s *AVHWAccel) SetCapsInternal(value int) {
	s.ptr.caps_internal = (C.int)(value)
}

// FrameParams skipped due to func ptr

// --- Struct AVSubtitleRect ---

// AVSubtitleRect wraps AVSubtitleRect.
type AVSubtitleRect struct {
	ptr *C.AVSubtitleRect
}

func (s *AVSubtitleRect) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func (s *AVSubtitleRect) X() int {
	value := s.ptr.x
	return int(value)
}

func (s *AVSubtitleRect) SetX(value int) {
	s.ptr.x = (C.int)(value)
}

func (s *AVSubtitleRect) Y() int {
	value := s.ptr.y
	return int(value)
}

func (s *AVSubtitleRect) SetY(value int) {
	s.ptr.y = (C.int)(value)
}

func (s *AVSubtitleRect) W() int {
	value := s.ptr.w
	return int(value)
}

func (s *AVSubtitleRect) SetW(value int) {
	s.ptr.w = (C.int)(value)
}

func (s *AVSubtitleRect) H() int {
	value := s.ptr.h
	return int(value)
}

func (s *AVSubtitleRect) SetH(value int) {
	s.ptr.h = (C.int)(value)
}

func (s *AVSubtitleRect) NbColors() int {
	value := s.ptr.nb_colors
	return int(value)
}

func (s *AVSubtitleRect) SetNbColors(value int) {
	s.ptr.nb_colors = (C.int)(value)
}

func (s *AVSubtitleRect) DataEntry(i uint) unsafe.Pointer {
	value := s.ptr.data[i]
	return unsafe.Pointer(value)
}

func (s *AVSubtitleRect) SetDataEntry(i uint, value unsafe.Pointer) {
	s.ptr.data[i] = (*C.uint8_t)(value)
}

func (s *AVSubtitleRect) LinesizeEntry(i uint) int {
	value := s.ptr.linesize[i]
	return int(value)
}

func (s *AVSubtitleRect) SetLinesizeEntry(i uint, value int) {
	s.ptr.linesize[i] = (C.int)(value)
}

func (s *AVSubtitleRect) Type() AVSubtitleType {
	value := s.ptr._type
	return AVSubtitleType(value)
}

func (s *AVSubtitleRect) SetType(value AVSubtitleType) {
	s.ptr._type = (C.enum_AVSubtitleType)(value)
}

func (s *AVSubtitleRect) Text() *CStr {
	value := s.ptr.text
	return wrapCStr(value)
}

func (s *AVSubtitleRect) SetText(value *CStr) {
	s.ptr.text = value.ptr
}

func (s *AVSubtitleRect) Ass() *CStr {
	value := s.ptr.ass
	return wrapCStr(value)
}

func (s *AVSubtitleRect) SetAss(value *CStr) {
	s.ptr.ass = value.ptr
}

func (s *AVSubtitleRect) Flags() int {
	value := s.ptr.flags
	return int(value)
}

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

func (s *AVSubtitle) Format() uint16 {
	value := s.ptr.format
	return uint16(value)
}

func (s *AVSubtitle) SetFormat(value uint16) {
	s.ptr.format = (C.uint16_t)(value)
}

func (s *AVSubtitle) StartDisplayTime() uint32 {
	value := s.ptr.start_display_time
	return uint32(value)
}

func (s *AVSubtitle) SetStartDisplayTime(value uint32) {
	s.ptr.start_display_time = (C.uint32_t)(value)
}

func (s *AVSubtitle) EndDisplayTime() uint32 {
	value := s.ptr.end_display_time
	return uint32(value)
}

func (s *AVSubtitle) SetEndDisplayTime(value uint32) {
	s.ptr.end_display_time = (C.uint32_t)(value)
}

func (s *AVSubtitle) NumRects() uint {
	value := s.ptr.num_rects
	return uint(value)
}

func (s *AVSubtitle) SetNumRects(value uint) {
	s.ptr.num_rects = (C.uint)(value)
}

func (s *AVSubtitle) RectsEntry(i uint) *AVSubtitleRect {
	value := s.ptr.rects
	ptr := arrayGet[*C.AVSubtitleRect](value, uintptr(i))
	var valueMapped *AVSubtitleRect
	if ptr != nil {
		valueMapped = &AVSubtitleRect{ptr: ptr}
	}
	return valueMapped
}

func (s *AVSubtitle) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

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

func (s *AVCodecParserContext) PrivData() unsafe.Pointer {
	value := s.ptr.priv_data
	return value
}

func (s *AVCodecParserContext) SetPrivData(value unsafe.Pointer) {
	s.ptr.priv_data = value
}

func (s *AVCodecParserContext) Parser() *AVCodecParser {
	value := s.ptr.parser
	var valueMapped *AVCodecParser
	if value != nil {
		valueMapped = &AVCodecParser{ptr: value}
	}
	return valueMapped
}

func (s *AVCodecParserContext) SetParser(value *AVCodecParser) {
	if value != nil {
		s.ptr.parser = value.ptr
	} else {
		s.ptr.parser = nil
	}
}

func (s *AVCodecParserContext) FrameOffset() int64 {
	value := s.ptr.frame_offset
	return int64(value)
}

func (s *AVCodecParserContext) SetFrameOffset(value int64) {
	s.ptr.frame_offset = (C.int64_t)(value)
}

func (s *AVCodecParserContext) CurOffset() int64 {
	value := s.ptr.cur_offset
	return int64(value)
}

func (s *AVCodecParserContext) SetCurOffset(value int64) {
	s.ptr.cur_offset = (C.int64_t)(value)
}

func (s *AVCodecParserContext) NextFrameOffset() int64 {
	value := s.ptr.next_frame_offset
	return int64(value)
}

func (s *AVCodecParserContext) SetNextFrameOffset(value int64) {
	s.ptr.next_frame_offset = (C.int64_t)(value)
}

func (s *AVCodecParserContext) PictType() int {
	value := s.ptr.pict_type
	return int(value)
}

func (s *AVCodecParserContext) SetPictType(value int) {
	s.ptr.pict_type = (C.int)(value)
}

func (s *AVCodecParserContext) RepeatPict() int {
	value := s.ptr.repeat_pict
	return int(value)
}

func (s *AVCodecParserContext) SetRepeatPict(value int) {
	s.ptr.repeat_pict = (C.int)(value)
}

func (s *AVCodecParserContext) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

func (s *AVCodecParserContext) SetPts(value int64) {
	s.ptr.pts = (C.int64_t)(value)
}

func (s *AVCodecParserContext) Dts() int64 {
	value := s.ptr.dts
	return int64(value)
}

func (s *AVCodecParserContext) SetDts(value int64) {
	s.ptr.dts = (C.int64_t)(value)
}

func (s *AVCodecParserContext) LastPts() int64 {
	value := s.ptr.last_pts
	return int64(value)
}

func (s *AVCodecParserContext) SetLastPts(value int64) {
	s.ptr.last_pts = (C.int64_t)(value)
}

func (s *AVCodecParserContext) LastDts() int64 {
	value := s.ptr.last_dts
	return int64(value)
}

func (s *AVCodecParserContext) SetLastDts(value int64) {
	s.ptr.last_dts = (C.int64_t)(value)
}

func (s *AVCodecParserContext) FetchTimestamp() int {
	value := s.ptr.fetch_timestamp
	return int(value)
}

func (s *AVCodecParserContext) SetFetchTimestamp(value int) {
	s.ptr.fetch_timestamp = (C.int)(value)
}

func (s *AVCodecParserContext) CurFrameStartIndex() int {
	value := s.ptr.cur_frame_start_index
	return int(value)
}

func (s *AVCodecParserContext) SetCurFrameStartIndex(value int) {
	s.ptr.cur_frame_start_index = (C.int)(value)
}

func (s *AVCodecParserContext) CurFrameOffsetEntry(i uint) int64 {
	value := s.ptr.cur_frame_offset[i]
	return int64(value)
}

func (s *AVCodecParserContext) SetCurFrameOffsetEntry(i uint, value int64) {
	s.ptr.cur_frame_offset[i] = (C.int64_t)(value)
}

func (s *AVCodecParserContext) CurFramePtsEntry(i uint) int64 {
	value := s.ptr.cur_frame_pts[i]
	return int64(value)
}

func (s *AVCodecParserContext) SetCurFramePtsEntry(i uint, value int64) {
	s.ptr.cur_frame_pts[i] = (C.int64_t)(value)
}

func (s *AVCodecParserContext) CurFrameDtsEntry(i uint) int64 {
	value := s.ptr.cur_frame_dts[i]
	return int64(value)
}

func (s *AVCodecParserContext) SetCurFrameDtsEntry(i uint, value int64) {
	s.ptr.cur_frame_dts[i] = (C.int64_t)(value)
}

func (s *AVCodecParserContext) Flags() int {
	value := s.ptr.flags
	return int(value)
}

func (s *AVCodecParserContext) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

func (s *AVCodecParserContext) Offset() int64 {
	value := s.ptr.offset
	return int64(value)
}

func (s *AVCodecParserContext) SetOffset(value int64) {
	s.ptr.offset = (C.int64_t)(value)
}

func (s *AVCodecParserContext) CurFrameEndEntry(i uint) int64 {
	value := s.ptr.cur_frame_end[i]
	return int64(value)
}

func (s *AVCodecParserContext) SetCurFrameEndEntry(i uint, value int64) {
	s.ptr.cur_frame_end[i] = (C.int64_t)(value)
}

func (s *AVCodecParserContext) KeyFrame() int {
	value := s.ptr.key_frame
	return int(value)
}

func (s *AVCodecParserContext) SetKeyFrame(value int) {
	s.ptr.key_frame = (C.int)(value)
}

func (s *AVCodecParserContext) DtsSyncPoint() int {
	value := s.ptr.dts_sync_point
	return int(value)
}

func (s *AVCodecParserContext) SetDtsSyncPoint(value int) {
	s.ptr.dts_sync_point = (C.int)(value)
}

func (s *AVCodecParserContext) DtsRefDtsDelta() int {
	value := s.ptr.dts_ref_dts_delta
	return int(value)
}

func (s *AVCodecParserContext) SetDtsRefDtsDelta(value int) {
	s.ptr.dts_ref_dts_delta = (C.int)(value)
}

func (s *AVCodecParserContext) PtsDtsDelta() int {
	value := s.ptr.pts_dts_delta
	return int(value)
}

func (s *AVCodecParserContext) SetPtsDtsDelta(value int) {
	s.ptr.pts_dts_delta = (C.int)(value)
}

func (s *AVCodecParserContext) CurFramePosEntry(i uint) int64 {
	value := s.ptr.cur_frame_pos[i]
	return int64(value)
}

func (s *AVCodecParserContext) SetCurFramePosEntry(i uint, value int64) {
	s.ptr.cur_frame_pos[i] = (C.int64_t)(value)
}

func (s *AVCodecParserContext) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}

func (s *AVCodecParserContext) SetPos(value int64) {
	s.ptr.pos = (C.int64_t)(value)
}

func (s *AVCodecParserContext) LastPos() int64 {
	value := s.ptr.last_pos
	return int64(value)
}

func (s *AVCodecParserContext) SetLastPos(value int64) {
	s.ptr.last_pos = (C.int64_t)(value)
}

func (s *AVCodecParserContext) Duration() int {
	value := s.ptr.duration
	return int(value)
}

func (s *AVCodecParserContext) SetDuration(value int) {
	s.ptr.duration = (C.int)(value)
}

func (s *AVCodecParserContext) FieldOrder() AVFieldOrder {
	value := s.ptr.field_order
	return AVFieldOrder(value)
}

func (s *AVCodecParserContext) SetFieldOrder(value AVFieldOrder) {
	s.ptr.field_order = (C.enum_AVFieldOrder)(value)
}

func (s *AVCodecParserContext) PictureStructure() AVPictureStructure {
	value := s.ptr.picture_structure
	return AVPictureStructure(value)
}

func (s *AVCodecParserContext) SetPictureStructure(value AVPictureStructure) {
	s.ptr.picture_structure = (C.enum_AVPictureStructure)(value)
}

func (s *AVCodecParserContext) OutputPictureNumber() int {
	value := s.ptr.output_picture_number
	return int(value)
}

func (s *AVCodecParserContext) SetOutputPictureNumber(value int) {
	s.ptr.output_picture_number = (C.int)(value)
}

func (s *AVCodecParserContext) Width() int {
	value := s.ptr.width
	return int(value)
}

func (s *AVCodecParserContext) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

func (s *AVCodecParserContext) Height() int {
	value := s.ptr.height
	return int(value)
}

func (s *AVCodecParserContext) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

func (s *AVCodecParserContext) CodedWidth() int {
	value := s.ptr.coded_width
	return int(value)
}

func (s *AVCodecParserContext) SetCodedWidth(value int) {
	s.ptr.coded_width = (C.int)(value)
}

func (s *AVCodecParserContext) CodedHeight() int {
	value := s.ptr.coded_height
	return int(value)
}

func (s *AVCodecParserContext) SetCodedHeight(value int) {
	s.ptr.coded_height = (C.int)(value)
}

func (s *AVCodecParserContext) Format() int {
	value := s.ptr.format
	return int(value)
}

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

func (s *AVCodecParser) CodecIdsEntry(i uint) int {
	value := s.ptr.codec_ids[i]
	return int(value)
}

func (s *AVCodecParser) SetCodecIdsEntry(i uint, value int) {
	s.ptr.codec_ids[i] = (C.int)(value)
}

func (s *AVCodecParser) PrivDataSize() int {
	value := s.ptr.priv_data_size
	return int(value)
}

func (s *AVCodecParser) SetPrivDataSize(value int) {
	s.ptr.priv_data_size = (C.int)(value)
}

// ParserInit skipped due to func ptr

// ParserParse skipped due to func ptr

// ParserClose skipped due to func ptr

// Split skipped due to func ptr

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

func (s *AVProfile) Profile() int {
	value := s.ptr.profile
	return int(value)
}

func (s *AVProfile) SetProfile(value int) {
	s.ptr.profile = (C.int)(value)
}

func (s *AVProfile) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

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

func (s *AVCodec) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

func (s *AVCodec) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

func (s *AVCodec) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

func (s *AVCodec) SetLongName(value *CStr) {
	s.ptr.long_name = value.ptr
}

func (s *AVCodec) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}

func (s *AVCodec) SetType(value AVMediaType) {
	s.ptr._type = (C.enum_AVMediaType)(value)
}

func (s *AVCodec) Id() AVCodecID {
	value := s.ptr.id
	return AVCodecID(value)
}

func (s *AVCodec) SetId(value AVCodecID) {
	s.ptr.id = (C.enum_AVCodecID)(value)
}

func (s *AVCodec) Capabilities() int {
	value := s.ptr.capabilities
	return int(value)
}

func (s *AVCodec) SetCapabilities(value int) {
	s.ptr.capabilities = (C.int)(value)
}

func (s *AVCodec) MaxLowres() uint8 {
	value := s.ptr.max_lowres
	return uint8(value)
}

func (s *AVCodec) SetMaxLowres(value uint8) {
	s.ptr.max_lowres = (C.uint8_t)(value)
}

// SupportedFramerates skipped due to struct value ptr

func (s *AVCodec) PixFmts() *Array[AVPixelFormat] {
	value := s.ptr.pix_fmts
	return ToAVPixelFormatArray(unsafe.Pointer(value))
}

func (s *AVCodec) SetPixFmts(value *Array[AVPixelFormat]) {
	if value != nil {
		s.ptr.pix_fmts = (*C.enum_AVPixelFormat)(value.ptr)
	} else {
		s.ptr.pix_fmts = nil
	}
}

// SupportedSamplerates skipped due to prim ptr

func (s *AVCodec) SampleFmts() *Array[AVSampleFormat] {
	value := s.ptr.sample_fmts
	return ToAVSampleFormatArray(unsafe.Pointer(value))
}

func (s *AVCodec) SetSampleFmts(value *Array[AVSampleFormat]) {
	if value != nil {
		s.ptr.sample_fmts = (*C.enum_AVSampleFormat)(value.ptr)
	} else {
		s.ptr.sample_fmts = nil
	}
}

// ChannelLayouts skipped due to prim ptr

func (s *AVCodec) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

func (s *AVCodec) SetPrivClass(value *AVClass) {
	if value != nil {
		s.ptr.priv_class = value.ptr
	} else {
		s.ptr.priv_class = nil
	}
}

func (s *AVCodec) Profiles() *AVProfile {
	value := s.ptr.profiles
	var valueMapped *AVProfile
	if value != nil {
		valueMapped = &AVProfile{ptr: value}
	}
	return valueMapped
}

func (s *AVCodec) SetProfiles(value *AVProfile) {
	if value != nil {
		s.ptr.profiles = value.ptr
	} else {
		s.ptr.profiles = nil
	}
}

func (s *AVCodec) WrapperName() *CStr {
	value := s.ptr.wrapper_name
	return wrapCStr(value)
}

func (s *AVCodec) SetWrapperName(value *CStr) {
	s.ptr.wrapper_name = value.ptr
}

func (s *AVCodec) ChLayouts() *AVChannelLayout {
	value := s.ptr.ch_layouts
	var valueMapped *AVChannelLayout
	if value != nil {
		valueMapped = &AVChannelLayout{ptr: value}
	}
	return valueMapped
}

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

func (s *AVCodecHWConfig) PixFmt() AVPixelFormat {
	value := s.ptr.pix_fmt
	return AVPixelFormat(value)
}

func (s *AVCodecHWConfig) SetPixFmt(value AVPixelFormat) {
	s.ptr.pix_fmt = (C.enum_AVPixelFormat)(value)
}

func (s *AVCodecHWConfig) Methods() int {
	value := s.ptr.methods
	return int(value)
}

func (s *AVCodecHWConfig) SetMethods(value int) {
	s.ptr.methods = (C.int)(value)
}

func (s *AVCodecHWConfig) DeviceType() AVHWDeviceType {
	value := s.ptr.device_type
	return AVHWDeviceType(value)
}

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

func (s *AVCodecDescriptor) Id() AVCodecID {
	value := s.ptr.id
	return AVCodecID(value)
}

func (s *AVCodecDescriptor) SetId(value AVCodecID) {
	s.ptr.id = (C.enum_AVCodecID)(value)
}

func (s *AVCodecDescriptor) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}

func (s *AVCodecDescriptor) SetType(value AVMediaType) {
	s.ptr._type = (C.enum_AVMediaType)(value)
}

func (s *AVCodecDescriptor) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

func (s *AVCodecDescriptor) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

func (s *AVCodecDescriptor) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

func (s *AVCodecDescriptor) SetLongName(value *CStr) {
	s.ptr.long_name = value.ptr
}

func (s *AVCodecDescriptor) Props() int {
	value := s.ptr.props
	return int(value)
}

func (s *AVCodecDescriptor) SetProps(value int) {
	s.ptr.props = (C.int)(value)
}

// MimeTypes skipped due to unknown ptr ptr

func (s *AVCodecDescriptor) Profiles() *AVProfile {
	value := s.ptr.profiles
	var valueMapped *AVProfile
	if value != nil {
		valueMapped = &AVProfile{ptr: value}
	}
	return valueMapped
}

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

func (s *AVCodecParameters) CodecType() AVMediaType {
	value := s.ptr.codec_type
	return AVMediaType(value)
}

func (s *AVCodecParameters) SetCodecType(value AVMediaType) {
	s.ptr.codec_type = (C.enum_AVMediaType)(value)
}

func (s *AVCodecParameters) CodecId() AVCodecID {
	value := s.ptr.codec_id
	return AVCodecID(value)
}

func (s *AVCodecParameters) SetCodecId(value AVCodecID) {
	s.ptr.codec_id = (C.enum_AVCodecID)(value)
}

func (s *AVCodecParameters) CodecTag() uint32 {
	value := s.ptr.codec_tag
	return uint32(value)
}

func (s *AVCodecParameters) SetCodecTag(value uint32) {
	s.ptr.codec_tag = (C.uint32_t)(value)
}

func (s *AVCodecParameters) Extradata() unsafe.Pointer {
	value := s.ptr.extradata
	return unsafe.Pointer(value)
}

func (s *AVCodecParameters) SetExtradata(value unsafe.Pointer) {
	s.ptr.extradata = (*C.uint8_t)(value)
}

func (s *AVCodecParameters) ExtradataSize() int {
	value := s.ptr.extradata_size
	return int(value)
}

func (s *AVCodecParameters) SetExtradataSize(value int) {
	s.ptr.extradata_size = (C.int)(value)
}

func (s *AVCodecParameters) Format() int {
	value := s.ptr.format
	return int(value)
}

func (s *AVCodecParameters) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

func (s *AVCodecParameters) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}

func (s *AVCodecParameters) SetBitRate(value int64) {
	s.ptr.bit_rate = (C.int64_t)(value)
}

func (s *AVCodecParameters) BitsPerCodedSample() int {
	value := s.ptr.bits_per_coded_sample
	return int(value)
}

func (s *AVCodecParameters) SetBitsPerCodedSample(value int) {
	s.ptr.bits_per_coded_sample = (C.int)(value)
}

func (s *AVCodecParameters) BitsPerRawSample() int {
	value := s.ptr.bits_per_raw_sample
	return int(value)
}

func (s *AVCodecParameters) SetBitsPerRawSample(value int) {
	s.ptr.bits_per_raw_sample = (C.int)(value)
}

func (s *AVCodecParameters) Profile() int {
	value := s.ptr.profile
	return int(value)
}

func (s *AVCodecParameters) SetProfile(value int) {
	s.ptr.profile = (C.int)(value)
}

func (s *AVCodecParameters) Level() int {
	value := s.ptr.level
	return int(value)
}

func (s *AVCodecParameters) SetLevel(value int) {
	s.ptr.level = (C.int)(value)
}

func (s *AVCodecParameters) Width() int {
	value := s.ptr.width
	return int(value)
}

func (s *AVCodecParameters) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

func (s *AVCodecParameters) Height() int {
	value := s.ptr.height
	return int(value)
}

func (s *AVCodecParameters) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

func (s *AVCodecParameters) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

func (s *AVCodecParameters) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

func (s *AVCodecParameters) FieldOrder() AVFieldOrder {
	value := s.ptr.field_order
	return AVFieldOrder(value)
}

func (s *AVCodecParameters) SetFieldOrder(value AVFieldOrder) {
	s.ptr.field_order = (C.enum_AVFieldOrder)(value)
}

func (s *AVCodecParameters) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}

func (s *AVCodecParameters) SetColorRange(value AVColorRange) {
	s.ptr.color_range = (C.enum_AVColorRange)(value)
}

func (s *AVCodecParameters) ColorPrimaries() AVColorPrimaries {
	value := s.ptr.color_primaries
	return AVColorPrimaries(value)
}

func (s *AVCodecParameters) SetColorPrimaries(value AVColorPrimaries) {
	s.ptr.color_primaries = (C.enum_AVColorPrimaries)(value)
}

func (s *AVCodecParameters) ColorTrc() AVColorTransferCharacteristic {
	value := s.ptr.color_trc
	return AVColorTransferCharacteristic(value)
}

func (s *AVCodecParameters) SetColorTrc(value AVColorTransferCharacteristic) {
	s.ptr.color_trc = (C.enum_AVColorTransferCharacteristic)(value)
}

func (s *AVCodecParameters) ColorSpace() AVColorSpace {
	value := s.ptr.color_space
	return AVColorSpace(value)
}

func (s *AVCodecParameters) SetColorSpace(value AVColorSpace) {
	s.ptr.color_space = (C.enum_AVColorSpace)(value)
}

func (s *AVCodecParameters) ChromaLocation() AVChromaLocation {
	value := s.ptr.chroma_location
	return AVChromaLocation(value)
}

func (s *AVCodecParameters) SetChromaLocation(value AVChromaLocation) {
	s.ptr.chroma_location = (C.enum_AVChromaLocation)(value)
}

func (s *AVCodecParameters) VideoDelay() int {
	value := s.ptr.video_delay
	return int(value)
}

func (s *AVCodecParameters) SetVideoDelay(value int) {
	s.ptr.video_delay = (C.int)(value)
}

func (s *AVCodecParameters) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

func (s *AVCodecParameters) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

func (s *AVCodecParameters) Channels() int {
	value := s.ptr.channels
	return int(value)
}

func (s *AVCodecParameters) SetChannels(value int) {
	s.ptr.channels = (C.int)(value)
}

func (s *AVCodecParameters) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

func (s *AVCodecParameters) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

func (s *AVCodecParameters) BlockAlign() int {
	value := s.ptr.block_align
	return int(value)
}

func (s *AVCodecParameters) SetBlockAlign(value int) {
	s.ptr.block_align = (C.int)(value)
}

func (s *AVCodecParameters) FrameSize() int {
	value := s.ptr.frame_size
	return int(value)
}

func (s *AVCodecParameters) SetFrameSize(value int) {
	s.ptr.frame_size = (C.int)(value)
}

func (s *AVCodecParameters) InitialPadding() int {
	value := s.ptr.initial_padding
	return int(value)
}

func (s *AVCodecParameters) SetInitialPadding(value int) {
	s.ptr.initial_padding = (C.int)(value)
}

func (s *AVCodecParameters) TrailingPadding() int {
	value := s.ptr.trailing_padding
	return int(value)
}

func (s *AVCodecParameters) SetTrailingPadding(value int) {
	s.ptr.trailing_padding = (C.int)(value)
}

func (s *AVCodecParameters) SeekPreroll() int {
	value := s.ptr.seek_preroll
	return int(value)
}

func (s *AVCodecParameters) SetSeekPreroll(value int) {
	s.ptr.seek_preroll = (C.int)(value)
}

// ChLayout skipped due to ident byvalue

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

func (s *AVPanScan) Id() int {
	value := s.ptr.id
	return int(value)
}

func (s *AVPanScan) SetId(value int) {
	s.ptr.id = (C.int)(value)
}

func (s *AVPanScan) Width() int {
	value := s.ptr.width
	return int(value)
}

func (s *AVPanScan) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

func (s *AVPanScan) Height() int {
	value := s.ptr.height
	return int(value)
}

func (s *AVPanScan) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

// PositionEntry skipped due to const array

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

func (s *AVCPBProperties) MaxBitrate() int64 {
	value := s.ptr.max_bitrate
	return int64(value)
}

func (s *AVCPBProperties) SetMaxBitrate(value int64) {
	s.ptr.max_bitrate = (C.int64_t)(value)
}

func (s *AVCPBProperties) MinBitrate() int64 {
	value := s.ptr.min_bitrate
	return int64(value)
}

func (s *AVCPBProperties) SetMinBitrate(value int64) {
	s.ptr.min_bitrate = (C.int64_t)(value)
}

func (s *AVCPBProperties) AvgBitrate() int64 {
	value := s.ptr.avg_bitrate
	return int64(value)
}

func (s *AVCPBProperties) SetAvgBitrate(value int64) {
	s.ptr.avg_bitrate = (C.int64_t)(value)
}

func (s *AVCPBProperties) BufferSize() int64 {
	value := s.ptr.buffer_size
	return int64(value)
}

func (s *AVCPBProperties) SetBufferSize(value int64) {
	s.ptr.buffer_size = (C.int64_t)(value)
}

func (s *AVCPBProperties) VbvDelay() uint64 {
	value := s.ptr.vbv_delay
	return uint64(value)
}

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

func (s *AVProducerReferenceTime) Wallclock() int64 {
	value := s.ptr.wallclock
	return int64(value)
}

func (s *AVProducerReferenceTime) SetWallclock(value int64) {
	s.ptr.wallclock = (C.int64_t)(value)
}

func (s *AVProducerReferenceTime) Flags() int {
	value := s.ptr.flags
	return int(value)
}

func (s *AVProducerReferenceTime) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// --- Struct AVPacketSideData ---

// AVPacketSideData wraps AVPacketSideData.
type AVPacketSideData struct {
	ptr *C.AVPacketSideData
}

func (s *AVPacketSideData) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func (s *AVPacketSideData) Data() unsafe.Pointer {
	value := s.ptr.data
	return unsafe.Pointer(value)
}

func (s *AVPacketSideData) SetData(value unsafe.Pointer) {
	s.ptr.data = (*C.uint8_t)(value)
}

func (s *AVPacketSideData) Size() uint64 {
	value := s.ptr.size
	return uint64(value)
}

func (s *AVPacketSideData) SetSize(value uint64) {
	s.ptr.size = (C.size_t)(value)
}

func (s *AVPacketSideData) Type() AVPacketSideDataType {
	value := s.ptr._type
	return AVPacketSideDataType(value)
}

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

func (s *AVPacket) Buf() *AVBufferRef {
	value := s.ptr.buf
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVPacket) SetBuf(value *AVBufferRef) {
	if value != nil {
		s.ptr.buf = value.ptr
	} else {
		s.ptr.buf = nil
	}
}

func (s *AVPacket) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

func (s *AVPacket) SetPts(value int64) {
	s.ptr.pts = (C.int64_t)(value)
}

func (s *AVPacket) Dts() int64 {
	value := s.ptr.dts
	return int64(value)
}

func (s *AVPacket) SetDts(value int64) {
	s.ptr.dts = (C.int64_t)(value)
}

func (s *AVPacket) Data() unsafe.Pointer {
	value := s.ptr.data
	return unsafe.Pointer(value)
}

func (s *AVPacket) SetData(value unsafe.Pointer) {
	s.ptr.data = (*C.uint8_t)(value)
}

func (s *AVPacket) Size() int {
	value := s.ptr.size
	return int(value)
}

func (s *AVPacket) SetSize(value int) {
	s.ptr.size = (C.int)(value)
}

func (s *AVPacket) StreamIndex() int {
	value := s.ptr.stream_index
	return int(value)
}

func (s *AVPacket) SetStreamIndex(value int) {
	s.ptr.stream_index = (C.int)(value)
}

func (s *AVPacket) Flags() int {
	value := s.ptr.flags
	return int(value)
}

func (s *AVPacket) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

func (s *AVPacket) SideData() *AVPacketSideData {
	value := s.ptr.side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}

func (s *AVPacket) SetSideData(value *AVPacketSideData) {
	if value != nil {
		s.ptr.side_data = value.ptr
	} else {
		s.ptr.side_data = nil
	}
}

func (s *AVPacket) SideDataElems() int {
	value := s.ptr.side_data_elems
	return int(value)
}

func (s *AVPacket) SetSideDataElems(value int) {
	s.ptr.side_data_elems = (C.int)(value)
}

func (s *AVPacket) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

func (s *AVPacket) SetDuration(value int64) {
	s.ptr.duration = (C.int64_t)(value)
}

func (s *AVPacket) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}

func (s *AVPacket) SetPos(value int64) {
	s.ptr.pos = (C.int64_t)(value)
}

func (s *AVPacket) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

func (s *AVPacket) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

func (s *AVPacket) OpaqueRef() *AVBufferRef {
	value := s.ptr.opaque_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVPacket) SetOpaqueRef(value *AVBufferRef) {
	if value != nil {
		s.ptr.opaque_ref = value.ptr
	} else {
		s.ptr.opaque_ref = nil
	}
}

func (s *AVPacket) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

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

// Pkt skipped due to ident byvalue

func (s *AVPacketList) Next() *AVPacketList {
	value := s.ptr.next
	var valueMapped *AVPacketList
	if value != nil {
		valueMapped = &AVPacketList{ptr: value}
	}
	return valueMapped
}

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

func (s *AVFilterContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

func (s *AVFilterContext) Filter() *AVFilter {
	value := s.ptr.filter
	var valueMapped *AVFilter
	if value != nil {
		valueMapped = &AVFilter{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterContext) SetFilter(value *AVFilter) {
	if value != nil {
		s.ptr.filter = value.ptr
	} else {
		s.ptr.filter = nil
	}
}

func (s *AVFilterContext) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

func (s *AVFilterContext) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

func (s *AVFilterContext) InputPads() *AVFilterPad {
	value := s.ptr.input_pads
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterContext) SetInputPads(value *AVFilterPad) {
	if value != nil {
		s.ptr.input_pads = value.ptr
	} else {
		s.ptr.input_pads = nil
	}
}

func (s *AVFilterContext) InputsEntry(i uint) *AVFilterLink {
	value := s.ptr.inputs
	ptr := arrayGet[*C.AVFilterLink](value, uintptr(i))
	var valueMapped *AVFilterLink
	if ptr != nil {
		valueMapped = &AVFilterLink{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFilterContext) NbInputs() uint {
	value := s.ptr.nb_inputs
	return uint(value)
}

func (s *AVFilterContext) SetNbInputs(value uint) {
	s.ptr.nb_inputs = (C.uint)(value)
}

func (s *AVFilterContext) OutputPads() *AVFilterPad {
	value := s.ptr.output_pads
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterContext) SetOutputPads(value *AVFilterPad) {
	if value != nil {
		s.ptr.output_pads = value.ptr
	} else {
		s.ptr.output_pads = nil
	}
}

func (s *AVFilterContext) OutputsEntry(i uint) *AVFilterLink {
	value := s.ptr.outputs
	ptr := arrayGet[*C.AVFilterLink](value, uintptr(i))
	var valueMapped *AVFilterLink
	if ptr != nil {
		valueMapped = &AVFilterLink{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFilterContext) NbOutputs() uint {
	value := s.ptr.nb_outputs
	return uint(value)
}

func (s *AVFilterContext) SetNbOutputs(value uint) {
	s.ptr.nb_outputs = (C.uint)(value)
}

func (s *AVFilterContext) Priv() unsafe.Pointer {
	value := s.ptr.priv
	return value
}

func (s *AVFilterContext) SetPriv(value unsafe.Pointer) {
	s.ptr.priv = value
}

func (s *AVFilterContext) Graph() *AVFilterGraph {
	value := s.ptr.graph
	var valueMapped *AVFilterGraph
	if value != nil {
		valueMapped = &AVFilterGraph{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterContext) SetGraph(value *AVFilterGraph) {
	if value != nil {
		s.ptr.graph = value.ptr
	} else {
		s.ptr.graph = nil
	}
}

func (s *AVFilterContext) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}

func (s *AVFilterContext) SetThreadType(value int) {
	s.ptr.thread_type = (C.int)(value)
}

func (s *AVFilterContext) Internal() *AVFilterInternal {
	value := s.ptr.internal
	var valueMapped *AVFilterInternal
	if value != nil {
		valueMapped = &AVFilterInternal{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterContext) SetInternal(value *AVFilterInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

// CommandQueue skipped due to ptr to ignored type

func (s *AVFilterContext) EnableStr() *CStr {
	value := s.ptr.enable_str
	return wrapCStr(value)
}

func (s *AVFilterContext) SetEnableStr(value *CStr) {
	s.ptr.enable_str = value.ptr
}

func (s *AVFilterContext) Enable() unsafe.Pointer {
	value := s.ptr.enable
	return value
}

func (s *AVFilterContext) SetEnable(value unsafe.Pointer) {
	s.ptr.enable = value
}

// VarValues skipped due to prim ptr

func (s *AVFilterContext) IsDisabled() int {
	value := s.ptr.is_disabled
	return int(value)
}

func (s *AVFilterContext) SetIsDisabled(value int) {
	s.ptr.is_disabled = (C.int)(value)
}

func (s *AVFilterContext) HwDeviceCtx() *AVBufferRef {
	value := s.ptr.hw_device_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterContext) SetHwDeviceCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_device_ctx = value.ptr
	} else {
		s.ptr.hw_device_ctx = nil
	}
}

func (s *AVFilterContext) NbThreads() int {
	value := s.ptr.nb_threads
	return int(value)
}

func (s *AVFilterContext) SetNbThreads(value int) {
	s.ptr.nb_threads = (C.int)(value)
}

func (s *AVFilterContext) Ready() uint {
	value := s.ptr.ready
	return uint(value)
}

func (s *AVFilterContext) SetReady(value uint) {
	s.ptr.ready = (C.uint)(value)
}

func (s *AVFilterContext) ExtraHwFrames() int {
	value := s.ptr.extra_hw_frames
	return int(value)
}

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

func (s *AVFilterLink) Src() *AVFilterContext {
	value := s.ptr.src
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterLink) SetSrc(value *AVFilterContext) {
	if value != nil {
		s.ptr.src = value.ptr
	} else {
		s.ptr.src = nil
	}
}

func (s *AVFilterLink) Srcpad() *AVFilterPad {
	value := s.ptr.srcpad
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterLink) SetSrcpad(value *AVFilterPad) {
	if value != nil {
		s.ptr.srcpad = value.ptr
	} else {
		s.ptr.srcpad = nil
	}
}

func (s *AVFilterLink) Dst() *AVFilterContext {
	value := s.ptr.dst
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterLink) SetDst(value *AVFilterContext) {
	if value != nil {
		s.ptr.dst = value.ptr
	} else {
		s.ptr.dst = nil
	}
}

func (s *AVFilterLink) Dstpad() *AVFilterPad {
	value := s.ptr.dstpad
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterLink) SetDstpad(value *AVFilterPad) {
	if value != nil {
		s.ptr.dstpad = value.ptr
	} else {
		s.ptr.dstpad = nil
	}
}

func (s *AVFilterLink) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}

func (s *AVFilterLink) SetType(value AVMediaType) {
	s.ptr._type = (C.enum_AVMediaType)(value)
}

func (s *AVFilterLink) W() int {
	value := s.ptr.w
	return int(value)
}

func (s *AVFilterLink) SetW(value int) {
	s.ptr.w = (C.int)(value)
}

func (s *AVFilterLink) H() int {
	value := s.ptr.h
	return int(value)
}

func (s *AVFilterLink) SetH(value int) {
	s.ptr.h = (C.int)(value)
}

func (s *AVFilterLink) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

func (s *AVFilterLink) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

func (s *AVFilterLink) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

func (s *AVFilterLink) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

func (s *AVFilterLink) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

func (s *AVFilterLink) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

func (s *AVFilterLink) Format() int {
	value := s.ptr.format
	return int(value)
}

func (s *AVFilterLink) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

func (s *AVFilterLink) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

func (s *AVFilterLink) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// ChLayout skipped due to ident byvalue

// Incfg skipped due to ident byvalue

// Outcfg skipped due to ident byvalue

func (s *AVFilterLink) InitState() uint32 {
	value := s.ptr.init_state
	return uint32(value)
}

func (s *AVFilterLink) SetInitState(value uint32) {
	s.ptr.init_state = value
}

func (s *AVFilterLink) Graph() *AVFilterGraph {
	value := s.ptr.graph
	var valueMapped *AVFilterGraph
	if value != nil {
		valueMapped = &AVFilterGraph{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterLink) SetGraph(value *AVFilterGraph) {
	if value != nil {
		s.ptr.graph = value.ptr
	} else {
		s.ptr.graph = nil
	}
}

func (s *AVFilterLink) CurrentPts() int64 {
	value := s.ptr.current_pts
	return int64(value)
}

func (s *AVFilterLink) SetCurrentPts(value int64) {
	s.ptr.current_pts = (C.int64_t)(value)
}

func (s *AVFilterLink) CurrentPtsUs() int64 {
	value := s.ptr.current_pts_us
	return int64(value)
}

func (s *AVFilterLink) SetCurrentPtsUs(value int64) {
	s.ptr.current_pts_us = (C.int64_t)(value)
}

func (s *AVFilterLink) AgeIndex() int {
	value := s.ptr.age_index
	return int(value)
}

func (s *AVFilterLink) SetAgeIndex(value int) {
	s.ptr.age_index = (C.int)(value)
}

func (s *AVFilterLink) FrameRate() *AVRational {
	value := s.ptr.frame_rate
	return &AVRational{value: value}
}

func (s *AVFilterLink) SetFrameRate(value *AVRational) {
	s.ptr.frame_rate = value.value
}

func (s *AVFilterLink) MinSamples() int {
	value := s.ptr.min_samples
	return int(value)
}

func (s *AVFilterLink) SetMinSamples(value int) {
	s.ptr.min_samples = (C.int)(value)
}

func (s *AVFilterLink) MaxSamples() int {
	value := s.ptr.max_samples
	return int(value)
}

func (s *AVFilterLink) SetMaxSamples(value int) {
	s.ptr.max_samples = (C.int)(value)
}

func (s *AVFilterLink) FrameCountIn() int64 {
	value := s.ptr.frame_count_in
	return int64(value)
}

func (s *AVFilterLink) SetFrameCountIn(value int64) {
	s.ptr.frame_count_in = (C.int64_t)(value)
}

func (s *AVFilterLink) FrameCountOut() int64 {
	value := s.ptr.frame_count_out
	return int64(value)
}

func (s *AVFilterLink) SetFrameCountOut(value int64) {
	s.ptr.frame_count_out = (C.int64_t)(value)
}

func (s *AVFilterLink) SampleCountIn() int64 {
	value := s.ptr.sample_count_in
	return int64(value)
}

func (s *AVFilterLink) SetSampleCountIn(value int64) {
	s.ptr.sample_count_in = (C.int64_t)(value)
}

func (s *AVFilterLink) SampleCountOut() int64 {
	value := s.ptr.sample_count_out
	return int64(value)
}

func (s *AVFilterLink) SetSampleCountOut(value int64) {
	s.ptr.sample_count_out = (C.int64_t)(value)
}

func (s *AVFilterLink) FramePool() unsafe.Pointer {
	value := s.ptr.frame_pool
	return value
}

func (s *AVFilterLink) SetFramePool(value unsafe.Pointer) {
	s.ptr.frame_pool = value
}

func (s *AVFilterLink) FrameWantedOut() int {
	value := s.ptr.frame_wanted_out
	return int(value)
}

func (s *AVFilterLink) SetFrameWantedOut(value int) {
	s.ptr.frame_wanted_out = (C.int)(value)
}

func (s *AVFilterLink) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterLink) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

func (s *AVFilterLink) ReservedEntry(i uint) uint8 {
	value := s.ptr.reserved[i]
	return uint8(value)
}

func (s *AVFilterLink) SetReservedEntry(i uint, value uint8) {
	s.ptr.reserved[i] = (C.char)(value)
}

// --- Struct AVFilterPad ---

// AVFilterPad wraps AVFilterPad.
type AVFilterPad struct {
	ptr *C.AVFilterPad
}

func (s *AVFilterPad) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

// --- Struct AVFilterFormats ---

// AVFilterFormats wraps AVFilterFormats.
type AVFilterFormats struct {
	ptr *C.AVFilterFormats
}

func (s *AVFilterFormats) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

// --- Struct AVFilterChannelLayouts ---

// AVFilterChannelLayouts wraps AVFilterChannelLayouts.
type AVFilterChannelLayouts struct {
	ptr *C.AVFilterChannelLayouts
}

func (s *AVFilterChannelLayouts) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
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

func (s *AVFilter) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

func (s *AVFilter) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

func (s *AVFilter) Description() *CStr {
	value := s.ptr.description
	return wrapCStr(value)
}

func (s *AVFilter) SetDescription(value *CStr) {
	s.ptr.description = value.ptr
}

func (s *AVFilter) Inputs() *AVFilterPad {
	value := s.ptr.inputs
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

func (s *AVFilter) SetInputs(value *AVFilterPad) {
	if value != nil {
		s.ptr.inputs = value.ptr
	} else {
		s.ptr.inputs = nil
	}
}

func (s *AVFilter) Outputs() *AVFilterPad {
	value := s.ptr.outputs
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

func (s *AVFilter) SetOutputs(value *AVFilterPad) {
	if value != nil {
		s.ptr.outputs = value.ptr
	} else {
		s.ptr.outputs = nil
	}
}

func (s *AVFilter) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

func (s *AVFilter) SetPrivClass(value *AVClass) {
	if value != nil {
		s.ptr.priv_class = value.ptr
	} else {
		s.ptr.priv_class = nil
	}
}

func (s *AVFilter) Flags() int {
	value := s.ptr.flags
	return int(value)
}

func (s *AVFilter) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

func (s *AVFilter) NbInputs() uint8 {
	value := s.ptr.nb_inputs
	return uint8(value)
}

func (s *AVFilter) SetNbInputs(value uint8) {
	s.ptr.nb_inputs = (C.uint8_t)(value)
}

func (s *AVFilter) NbOutputs() uint8 {
	value := s.ptr.nb_outputs
	return uint8(value)
}

func (s *AVFilter) SetNbOutputs(value uint8) {
	s.ptr.nb_outputs = (C.uint8_t)(value)
}

func (s *AVFilter) FormatsState() uint8 {
	value := s.ptr.formats_state
	return uint8(value)
}

func (s *AVFilter) SetFormatsState(value uint8) {
	s.ptr.formats_state = (C.uint8_t)(value)
}

// Preinit skipped due to func ptr

// Init skipped due to func ptr

// Uninit skipped due to func ptr

// Formats skipped due to union type

func (s *AVFilter) PrivSize() int {
	value := s.ptr.priv_size
	return int(value)
}

func (s *AVFilter) SetPrivSize(value int) {
	s.ptr.priv_size = (C.int)(value)
}

func (s *AVFilter) FlagsInternal() int {
	value := s.ptr.flags_internal
	return int(value)
}

func (s *AVFilter) SetFlagsInternal(value int) {
	s.ptr.flags_internal = (C.int)(value)
}

// ProcessCommand skipped due to func ptr

// Activate skipped due to func ptr

// --- Struct AVFilterInternal ---

// AVFilterInternal wraps AVFilterInternal.
type AVFilterInternal struct {
	ptr *C.AVFilterInternal
}

func (s *AVFilterInternal) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
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

func (s *AVFilterFormatsConfig) Formats() *AVFilterFormats {
	value := s.ptr.formats
	var valueMapped *AVFilterFormats
	if value != nil {
		valueMapped = &AVFilterFormats{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterFormatsConfig) SetFormats(value *AVFilterFormats) {
	if value != nil {
		s.ptr.formats = value.ptr
	} else {
		s.ptr.formats = nil
	}
}

func (s *AVFilterFormatsConfig) Samplerates() *AVFilterFormats {
	value := s.ptr.samplerates
	var valueMapped *AVFilterFormats
	if value != nil {
		valueMapped = &AVFilterFormats{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterFormatsConfig) SetSamplerates(value *AVFilterFormats) {
	if value != nil {
		s.ptr.samplerates = value.ptr
	} else {
		s.ptr.samplerates = nil
	}
}

func (s *AVFilterFormatsConfig) ChannelLayouts() *AVFilterChannelLayouts {
	value := s.ptr.channel_layouts
	var valueMapped *AVFilterChannelLayouts
	if value != nil {
		valueMapped = &AVFilterChannelLayouts{ptr: value}
	}
	return valueMapped
}

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

// --- Struct AVFilterGraph ---

// AVFilterGraph wraps AVFilterGraph.
type AVFilterGraph struct {
	ptr *C.AVFilterGraph
}

func (s *AVFilterGraph) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func (s *AVFilterGraph) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterGraph) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

func (s *AVFilterGraph) FiltersEntry(i uint) *AVFilterContext {
	value := s.ptr.filters
	ptr := arrayGet[*C.AVFilterContext](value, uintptr(i))
	var valueMapped *AVFilterContext
	if ptr != nil {
		valueMapped = &AVFilterContext{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFilterGraph) NbFilters() uint {
	value := s.ptr.nb_filters
	return uint(value)
}

func (s *AVFilterGraph) SetNbFilters(value uint) {
	s.ptr.nb_filters = (C.uint)(value)
}

func (s *AVFilterGraph) ScaleSwsOpts() *CStr {
	value := s.ptr.scale_sws_opts
	return wrapCStr(value)
}

func (s *AVFilterGraph) SetScaleSwsOpts(value *CStr) {
	s.ptr.scale_sws_opts = value.ptr
}

func (s *AVFilterGraph) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}

func (s *AVFilterGraph) SetThreadType(value int) {
	s.ptr.thread_type = (C.int)(value)
}

func (s *AVFilterGraph) NbThreads() int {
	value := s.ptr.nb_threads
	return int(value)
}

func (s *AVFilterGraph) SetNbThreads(value int) {
	s.ptr.nb_threads = (C.int)(value)
}

func (s *AVFilterGraph) Internal() *AVFilterGraphInternal {
	value := s.ptr.internal
	var valueMapped *AVFilterGraphInternal
	if value != nil {
		valueMapped = &AVFilterGraphInternal{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterGraph) SetInternal(value *AVFilterGraphInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

func (s *AVFilterGraph) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

func (s *AVFilterGraph) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// Execute skipped due to callback ptr

func (s *AVFilterGraph) AresampleSwrOpts() *CStr {
	value := s.ptr.aresample_swr_opts
	return wrapCStr(value)
}

func (s *AVFilterGraph) SetAresampleSwrOpts(value *CStr) {
	s.ptr.aresample_swr_opts = value.ptr
}

func (s *AVFilterGraph) SinkLinksEntry(i uint) *AVFilterLink {
	value := s.ptr.sink_links
	ptr := arrayGet[*C.AVFilterLink](value, uintptr(i))
	var valueMapped *AVFilterLink
	if ptr != nil {
		valueMapped = &AVFilterLink{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFilterGraph) SinkLinksCount() int {
	value := s.ptr.sink_links_count
	return int(value)
}

func (s *AVFilterGraph) SetSinkLinksCount(value int) {
	s.ptr.sink_links_count = (C.int)(value)
}

func (s *AVFilterGraph) DisableAutoConvert() uint {
	value := s.ptr.disable_auto_convert
	return uint(value)
}

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

func (s *AVFilterInOut) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

func (s *AVFilterInOut) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

func (s *AVFilterInOut) FilterCtx() *AVFilterContext {
	value := s.ptr.filter_ctx
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterInOut) SetFilterCtx(value *AVFilterContext) {
	if value != nil {
		s.ptr.filter_ctx = value.ptr
	} else {
		s.ptr.filter_ctx = nil
	}
}

func (s *AVFilterInOut) PadIdx() int {
	value := s.ptr.pad_idx
	return int(value)
}

func (s *AVFilterInOut) SetPadIdx(value int) {
	s.ptr.pad_idx = (C.int)(value)
}

func (s *AVFilterInOut) Next() *AVFilterInOut {
	value := s.ptr.next
	var valueMapped *AVFilterInOut
	if value != nil {
		valueMapped = &AVFilterInOut{ptr: value}
	}
	return valueMapped
}

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

func (s *AVFilterPadParams) Label() *CStr {
	value := s.ptr.label
	return wrapCStr(value)
}

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

func (s *AVFilterParams) Filter() *AVFilterContext {
	value := s.ptr.filter
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterParams) SetFilter(value *AVFilterContext) {
	if value != nil {
		s.ptr.filter = value.ptr
	} else {
		s.ptr.filter = nil
	}
}

func (s *AVFilterParams) FilterName() *CStr {
	value := s.ptr.filter_name
	return wrapCStr(value)
}

func (s *AVFilterParams) SetFilterName(value *CStr) {
	s.ptr.filter_name = value.ptr
}

func (s *AVFilterParams) InstanceName() *CStr {
	value := s.ptr.instance_name
	return wrapCStr(value)
}

func (s *AVFilterParams) SetInstanceName(value *CStr) {
	s.ptr.instance_name = value.ptr
}

func (s *AVFilterParams) Opts() *AVDictionary {
	value := s.ptr.opts
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterParams) SetOpts(value *AVDictionary) {
	if value != nil {
		s.ptr.opts = value.ptr
	} else {
		s.ptr.opts = nil
	}
}

func (s *AVFilterParams) InputsEntry(i uint) *AVFilterPadParams {
	value := s.ptr.inputs
	ptr := arrayGet[*C.AVFilterPadParams](value, uintptr(i))
	var valueMapped *AVFilterPadParams
	if ptr != nil {
		valueMapped = &AVFilterPadParams{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFilterParams) NbInputs() uint {
	value := s.ptr.nb_inputs
	return uint(value)
}

func (s *AVFilterParams) SetNbInputs(value uint) {
	s.ptr.nb_inputs = (C.uint)(value)
}

func (s *AVFilterParams) OutputsEntry(i uint) *AVFilterPadParams {
	value := s.ptr.outputs
	ptr := arrayGet[*C.AVFilterPadParams](value, uintptr(i))
	var valueMapped *AVFilterPadParams
	if ptr != nil {
		valueMapped = &AVFilterPadParams{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFilterParams) NbOutputs() uint {
	value := s.ptr.nb_outputs
	return uint(value)
}

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

func (s *AVFilterChain) FiltersEntry(i uint) *AVFilterParams {
	value := s.ptr.filters
	ptr := arrayGet[*C.AVFilterParams](value, uintptr(i))
	var valueMapped *AVFilterParams
	if ptr != nil {
		valueMapped = &AVFilterParams{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFilterChain) NbFilters() uint64 {
	value := s.ptr.nb_filters
	return uint64(value)
}

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

func (s *AVFilterGraphSegment) Graph() *AVFilterGraph {
	value := s.ptr.graph
	var valueMapped *AVFilterGraph
	if value != nil {
		valueMapped = &AVFilterGraph{ptr: value}
	}
	return valueMapped
}

func (s *AVFilterGraphSegment) SetGraph(value *AVFilterGraph) {
	if value != nil {
		s.ptr.graph = value.ptr
	} else {
		s.ptr.graph = nil
	}
}

func (s *AVFilterGraphSegment) ChainsEntry(i uint) *AVFilterChain {
	value := s.ptr.chains
	ptr := arrayGet[*C.AVFilterChain](value, uintptr(i))
	var valueMapped *AVFilterChain
	if ptr != nil {
		valueMapped = &AVFilterChain{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFilterGraphSegment) NbChains() uint64 {
	value := s.ptr.nb_chains
	return uint64(value)
}

func (s *AVFilterGraphSegment) SetNbChains(value uint64) {
	s.ptr.nb_chains = (C.size_t)(value)
}

func (s *AVFilterGraphSegment) ScaleSwsOpts() *CStr {
	value := s.ptr.scale_sws_opts
	return wrapCStr(value)
}

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

func (s *AVBufferSrcParameters) Format() int {
	value := s.ptr.format
	return int(value)
}

func (s *AVBufferSrcParameters) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

func (s *AVBufferSrcParameters) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

func (s *AVBufferSrcParameters) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

func (s *AVBufferSrcParameters) Width() int {
	value := s.ptr.width
	return int(value)
}

func (s *AVBufferSrcParameters) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

func (s *AVBufferSrcParameters) Height() int {
	value := s.ptr.height
	return int(value)
}

func (s *AVBufferSrcParameters) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

func (s *AVBufferSrcParameters) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

func (s *AVBufferSrcParameters) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

func (s *AVBufferSrcParameters) FrameRate() *AVRational {
	value := s.ptr.frame_rate
	return &AVRational{value: value}
}

func (s *AVBufferSrcParameters) SetFrameRate(value *AVRational) {
	s.ptr.frame_rate = value.value
}

func (s *AVBufferSrcParameters) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVBufferSrcParameters) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

func (s *AVBufferSrcParameters) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

func (s *AVBufferSrcParameters) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

func (s *AVBufferSrcParameters) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

func (s *AVBufferSrcParameters) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

// ChLayout skipped due to ident byvalue

// --- Struct AVDeviceInfoList ---

// AVDeviceInfoList wraps AVDeviceInfoList.
type AVDeviceInfoList struct {
	ptr *C.struct_AVDeviceInfoList
}

func (s *AVDeviceInfoList) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
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

func (s *AVProbeData) Filename() *CStr {
	value := s.ptr.filename
	return wrapCStr(value)
}

func (s *AVProbeData) SetFilename(value *CStr) {
	s.ptr.filename = value.ptr
}

// Buf skipped due to prim ptr

func (s *AVProbeData) BufSize() int {
	value := s.ptr.buf_size
	return int(value)
}

func (s *AVProbeData) SetBufSize(value int) {
	s.ptr.buf_size = (C.int)(value)
}

func (s *AVProbeData) MimeType() *CStr {
	value := s.ptr.mime_type
	return wrapCStr(value)
}

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

func (s *AVOutputFormat) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

func (s *AVOutputFormat) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

func (s *AVOutputFormat) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

func (s *AVOutputFormat) SetLongName(value *CStr) {
	s.ptr.long_name = value.ptr
}

func (s *AVOutputFormat) MimeType() *CStr {
	value := s.ptr.mime_type
	return wrapCStr(value)
}

func (s *AVOutputFormat) SetMimeType(value *CStr) {
	s.ptr.mime_type = value.ptr
}

func (s *AVOutputFormat) Extensions() *CStr {
	value := s.ptr.extensions
	return wrapCStr(value)
}

func (s *AVOutputFormat) SetExtensions(value *CStr) {
	s.ptr.extensions = value.ptr
}

func (s *AVOutputFormat) AudioCodec() AVCodecID {
	value := s.ptr.audio_codec
	return AVCodecID(value)
}

func (s *AVOutputFormat) SetAudioCodec(value AVCodecID) {
	s.ptr.audio_codec = (C.enum_AVCodecID)(value)
}

func (s *AVOutputFormat) VideoCodec() AVCodecID {
	value := s.ptr.video_codec
	return AVCodecID(value)
}

func (s *AVOutputFormat) SetVideoCodec(value AVCodecID) {
	s.ptr.video_codec = (C.enum_AVCodecID)(value)
}

func (s *AVOutputFormat) SubtitleCodec() AVCodecID {
	value := s.ptr.subtitle_codec
	return AVCodecID(value)
}

func (s *AVOutputFormat) SetSubtitleCodec(value AVCodecID) {
	s.ptr.subtitle_codec = (C.enum_AVCodecID)(value)
}

func (s *AVOutputFormat) Flags() int {
	value := s.ptr.flags
	return int(value)
}

func (s *AVOutputFormat) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

func (s *AVOutputFormat) CodecTagEntry(i uint) *AVCodecTag {
	value := s.ptr.codec_tag
	ptr := arrayGet[*C.struct_AVCodecTag](value, uintptr(i))
	var valueMapped *AVCodecTag
	if ptr != nil {
		valueMapped = &AVCodecTag{ptr: ptr}
	}
	return valueMapped
}

func (s *AVOutputFormat) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

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

func (s *AVInputFormat) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

func (s *AVInputFormat) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

func (s *AVInputFormat) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

func (s *AVInputFormat) SetLongName(value *CStr) {
	s.ptr.long_name = value.ptr
}

func (s *AVInputFormat) Flags() int {
	value := s.ptr.flags
	return int(value)
}

func (s *AVInputFormat) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

func (s *AVInputFormat) Extensions() *CStr {
	value := s.ptr.extensions
	return wrapCStr(value)
}

func (s *AVInputFormat) SetExtensions(value *CStr) {
	s.ptr.extensions = value.ptr
}

func (s *AVInputFormat) CodecTagEntry(i uint) *AVCodecTag {
	value := s.ptr.codec_tag
	ptr := arrayGet[*C.struct_AVCodecTag](value, uintptr(i))
	var valueMapped *AVCodecTag
	if ptr != nil {
		valueMapped = &AVCodecTag{ptr: ptr}
	}
	return valueMapped
}

func (s *AVInputFormat) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

func (s *AVInputFormat) SetPrivClass(value *AVClass) {
	if value != nil {
		s.ptr.priv_class = value.ptr
	} else {
		s.ptr.priv_class = nil
	}
}

func (s *AVInputFormat) MimeType() *CStr {
	value := s.ptr.mime_type
	return wrapCStr(value)
}

func (s *AVInputFormat) SetMimeType(value *CStr) {
	s.ptr.mime_type = value.ptr
}

func (s *AVInputFormat) RawCodecId() int {
	value := s.ptr.raw_codec_id
	return int(value)
}

func (s *AVInputFormat) SetRawCodecId(value int) {
	s.ptr.raw_codec_id = (C.int)(value)
}

func (s *AVInputFormat) PrivDataSize() int {
	value := s.ptr.priv_data_size
	return int(value)
}

func (s *AVInputFormat) SetPrivDataSize(value int) {
	s.ptr.priv_data_size = (C.int)(value)
}

func (s *AVInputFormat) FlagsInternal() int {
	value := s.ptr.flags_internal
	return int(value)
}

func (s *AVInputFormat) SetFlagsInternal(value int) {
	s.ptr.flags_internal = (C.int)(value)
}

// ReadProbe skipped due to func ptr

// ReadHeader skipped due to func ptr

// ReadPacket skipped due to func ptr

// ReadClose skipped due to func ptr

// ReadSeek skipped due to func ptr

// ReadTimestamp skipped due to func ptr

// ReadPlay skipped due to func ptr

// ReadPause skipped due to func ptr

// ReadSeek2 skipped due to func ptr

// GetDeviceList skipped due to func ptr

// --- Struct AVIndexEntry ---

// AVIndexEntry wraps AVIndexEntry.
type AVIndexEntry struct {
	ptr *C.AVIndexEntry
}

func (s *AVIndexEntry) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func (s *AVIndexEntry) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}

func (s *AVIndexEntry) SetPos(value int64) {
	s.ptr.pos = (C.int64_t)(value)
}

func (s *AVIndexEntry) Timestamp() int64 {
	value := s.ptr.timestamp
	return int64(value)
}

func (s *AVIndexEntry) SetTimestamp(value int64) {
	s.ptr.timestamp = (C.int64_t)(value)
}

// Flags skipped due to bitfield

// Size skipped due to bitfield

func (s *AVIndexEntry) MinDistance() int {
	value := s.ptr.min_distance
	return int(value)
}

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

func (s *AVStream) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

func (s *AVStream) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

func (s *AVStream) Index() int {
	value := s.ptr.index
	return int(value)
}

func (s *AVStream) SetIndex(value int) {
	s.ptr.index = (C.int)(value)
}

func (s *AVStream) Id() int {
	value := s.ptr.id
	return int(value)
}

func (s *AVStream) SetId(value int) {
	s.ptr.id = (C.int)(value)
}

func (s *AVStream) Codecpar() *AVCodecParameters {
	value := s.ptr.codecpar
	var valueMapped *AVCodecParameters
	if value != nil {
		valueMapped = &AVCodecParameters{ptr: value}
	}
	return valueMapped
}

func (s *AVStream) SetCodecpar(value *AVCodecParameters) {
	if value != nil {
		s.ptr.codecpar = value.ptr
	} else {
		s.ptr.codecpar = nil
	}
}

func (s *AVStream) PrivData() unsafe.Pointer {
	value := s.ptr.priv_data
	return value
}

func (s *AVStream) SetPrivData(value unsafe.Pointer) {
	s.ptr.priv_data = value
}

func (s *AVStream) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

func (s *AVStream) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

func (s *AVStream) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}

func (s *AVStream) SetStartTime(value int64) {
	s.ptr.start_time = (C.int64_t)(value)
}

func (s *AVStream) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

func (s *AVStream) SetDuration(value int64) {
	s.ptr.duration = (C.int64_t)(value)
}

func (s *AVStream) NbFrames() int64 {
	value := s.ptr.nb_frames
	return int64(value)
}

func (s *AVStream) SetNbFrames(value int64) {
	s.ptr.nb_frames = (C.int64_t)(value)
}

func (s *AVStream) Disposition() int {
	value := s.ptr.disposition
	return int(value)
}

func (s *AVStream) SetDisposition(value int) {
	s.ptr.disposition = (C.int)(value)
}

func (s *AVStream) Discard() AVDiscard {
	value := s.ptr.discard
	return AVDiscard(value)
}

func (s *AVStream) SetDiscard(value AVDiscard) {
	s.ptr.discard = (C.enum_AVDiscard)(value)
}

func (s *AVStream) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

func (s *AVStream) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

func (s *AVStream) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

func (s *AVStream) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

func (s *AVStream) AvgFrameRate() *AVRational {
	value := s.ptr.avg_frame_rate
	return &AVRational{value: value}
}

func (s *AVStream) SetAvgFrameRate(value *AVRational) {
	s.ptr.avg_frame_rate = value.value
}

// AttachedPic skipped due to ident byvalue

func (s *AVStream) SideData() *AVPacketSideData {
	value := s.ptr.side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}

func (s *AVStream) SetSideData(value *AVPacketSideData) {
	if value != nil {
		s.ptr.side_data = value.ptr
	} else {
		s.ptr.side_data = nil
	}
}

func (s *AVStream) NbSideData() int {
	value := s.ptr.nb_side_data
	return int(value)
}

func (s *AVStream) SetNbSideData(value int) {
	s.ptr.nb_side_data = (C.int)(value)
}

func (s *AVStream) EventFlags() int {
	value := s.ptr.event_flags
	return int(value)
}

func (s *AVStream) SetEventFlags(value int) {
	s.ptr.event_flags = (C.int)(value)
}

func (s *AVStream) RFrameRate() *AVRational {
	value := s.ptr.r_frame_rate
	return &AVRational{value: value}
}

func (s *AVStream) SetRFrameRate(value *AVRational) {
	s.ptr.r_frame_rate = value.value
}

func (s *AVStream) PtsWrapBits() int {
	value := s.ptr.pts_wrap_bits
	return int(value)
}

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

func (s *AVProgram) Id() int {
	value := s.ptr.id
	return int(value)
}

func (s *AVProgram) SetId(value int) {
	s.ptr.id = (C.int)(value)
}

func (s *AVProgram) Flags() int {
	value := s.ptr.flags
	return int(value)
}

func (s *AVProgram) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

func (s *AVProgram) Discard() AVDiscard {
	value := s.ptr.discard
	return AVDiscard(value)
}

func (s *AVProgram) SetDiscard(value AVDiscard) {
	s.ptr.discard = (C.enum_AVDiscard)(value)
}

// StreamIndex skipped due to prim ptr

func (s *AVProgram) NbStreamIndexes() uint {
	value := s.ptr.nb_stream_indexes
	return uint(value)
}

func (s *AVProgram) SetNbStreamIndexes(value uint) {
	s.ptr.nb_stream_indexes = (C.uint)(value)
}

func (s *AVProgram) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

func (s *AVProgram) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

func (s *AVProgram) ProgramNum() int {
	value := s.ptr.program_num
	return int(value)
}

func (s *AVProgram) SetProgramNum(value int) {
	s.ptr.program_num = (C.int)(value)
}

func (s *AVProgram) PmtPid() int {
	value := s.ptr.pmt_pid
	return int(value)
}

func (s *AVProgram) SetPmtPid(value int) {
	s.ptr.pmt_pid = (C.int)(value)
}

func (s *AVProgram) PcrPid() int {
	value := s.ptr.pcr_pid
	return int(value)
}

func (s *AVProgram) SetPcrPid(value int) {
	s.ptr.pcr_pid = (C.int)(value)
}

func (s *AVProgram) PmtVersion() int {
	value := s.ptr.pmt_version
	return int(value)
}

func (s *AVProgram) SetPmtVersion(value int) {
	s.ptr.pmt_version = (C.int)(value)
}

func (s *AVProgram) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}

func (s *AVProgram) SetStartTime(value int64) {
	s.ptr.start_time = (C.int64_t)(value)
}

func (s *AVProgram) EndTime() int64 {
	value := s.ptr.end_time
	return int64(value)
}

func (s *AVProgram) SetEndTime(value int64) {
	s.ptr.end_time = (C.int64_t)(value)
}

func (s *AVProgram) PtsWrapReference() int64 {
	value := s.ptr.pts_wrap_reference
	return int64(value)
}

func (s *AVProgram) SetPtsWrapReference(value int64) {
	s.ptr.pts_wrap_reference = (C.int64_t)(value)
}

func (s *AVProgram) PtsWrapBehavior() int {
	value := s.ptr.pts_wrap_behavior
	return int(value)
}

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

func (s *AVChapter) Id() int64 {
	value := s.ptr.id
	return int64(value)
}

func (s *AVChapter) SetId(value int64) {
	s.ptr.id = (C.int64_t)(value)
}

func (s *AVChapter) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

func (s *AVChapter) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

func (s *AVChapter) Start() int64 {
	value := s.ptr.start
	return int64(value)
}

func (s *AVChapter) SetStart(value int64) {
	s.ptr.start = (C.int64_t)(value)
}

func (s *AVChapter) End() int64 {
	value := s.ptr.end
	return int64(value)
}

func (s *AVChapter) SetEnd(value int64) {
	s.ptr.end = (C.int64_t)(value)
}

func (s *AVChapter) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

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

func (s *AVFormatContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

func (s *AVFormatContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

func (s *AVFormatContext) Iformat() *AVInputFormat {
	value := s.ptr.iformat
	var valueMapped *AVInputFormat
	if value != nil {
		valueMapped = &AVInputFormat{ptr: value}
	}
	return valueMapped
}

func (s *AVFormatContext) SetIformat(value *AVInputFormat) {
	if value != nil {
		s.ptr.iformat = value.ptr
	} else {
		s.ptr.iformat = nil
	}
}

func (s *AVFormatContext) Oformat() *AVOutputFormat {
	value := s.ptr.oformat
	var valueMapped *AVOutputFormat
	if value != nil {
		valueMapped = &AVOutputFormat{ptr: value}
	}
	return valueMapped
}

func (s *AVFormatContext) SetOformat(value *AVOutputFormat) {
	if value != nil {
		s.ptr.oformat = value.ptr
	} else {
		s.ptr.oformat = nil
	}
}

func (s *AVFormatContext) PrivData() unsafe.Pointer {
	value := s.ptr.priv_data
	return value
}

func (s *AVFormatContext) SetPrivData(value unsafe.Pointer) {
	s.ptr.priv_data = value
}

func (s *AVFormatContext) Pb() *AVIOContext {
	value := s.ptr.pb
	var valueMapped *AVIOContext
	if value != nil {
		valueMapped = &AVIOContext{ptr: value}
	}
	return valueMapped
}

func (s *AVFormatContext) SetPb(value *AVIOContext) {
	if value != nil {
		s.ptr.pb = value.ptr
	} else {
		s.ptr.pb = nil
	}
}

func (s *AVFormatContext) CtxFlags() int {
	value := s.ptr.ctx_flags
	return int(value)
}

func (s *AVFormatContext) SetCtxFlags(value int) {
	s.ptr.ctx_flags = (C.int)(value)
}

func (s *AVFormatContext) NbStreams() uint {
	value := s.ptr.nb_streams
	return uint(value)
}

func (s *AVFormatContext) SetNbStreams(value uint) {
	s.ptr.nb_streams = (C.uint)(value)
}

func (s *AVFormatContext) StreamsEntry(i uint) *AVStream {
	value := s.ptr.streams
	ptr := arrayGet[*C.AVStream](value, uintptr(i))
	var valueMapped *AVStream
	if ptr != nil {
		valueMapped = &AVStream{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFormatContext) Url() *CStr {
	value := s.ptr.url
	return wrapCStr(value)
}

func (s *AVFormatContext) SetUrl(value *CStr) {
	s.ptr.url = value.ptr
}

func (s *AVFormatContext) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}

func (s *AVFormatContext) SetStartTime(value int64) {
	s.ptr.start_time = (C.int64_t)(value)
}

func (s *AVFormatContext) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

func (s *AVFormatContext) SetDuration(value int64) {
	s.ptr.duration = (C.int64_t)(value)
}

func (s *AVFormatContext) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}

func (s *AVFormatContext) SetBitRate(value int64) {
	s.ptr.bit_rate = (C.int64_t)(value)
}

func (s *AVFormatContext) PacketSize() uint {
	value := s.ptr.packet_size
	return uint(value)
}

func (s *AVFormatContext) SetPacketSize(value uint) {
	s.ptr.packet_size = (C.uint)(value)
}

func (s *AVFormatContext) MaxDelay() int {
	value := s.ptr.max_delay
	return int(value)
}

func (s *AVFormatContext) SetMaxDelay(value int) {
	s.ptr.max_delay = (C.int)(value)
}

func (s *AVFormatContext) Flags() int {
	value := s.ptr.flags
	return int(value)
}

func (s *AVFormatContext) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

func (s *AVFormatContext) Probesize() int64 {
	value := s.ptr.probesize
	return int64(value)
}

func (s *AVFormatContext) SetProbesize(value int64) {
	s.ptr.probesize = (C.int64_t)(value)
}

func (s *AVFormatContext) MaxAnalyzeDuration() int64 {
	value := s.ptr.max_analyze_duration
	return int64(value)
}

func (s *AVFormatContext) SetMaxAnalyzeDuration(value int64) {
	s.ptr.max_analyze_duration = (C.int64_t)(value)
}

func (s *AVFormatContext) Key() unsafe.Pointer {
	value := s.ptr.key
	return unsafe.Pointer(value)
}

func (s *AVFormatContext) SetKey(value unsafe.Pointer) {
	s.ptr.key = (*C.uint8_t)(value)
}

func (s *AVFormatContext) Keylen() int {
	value := s.ptr.keylen
	return int(value)
}

func (s *AVFormatContext) SetKeylen(value int) {
	s.ptr.keylen = (C.int)(value)
}

func (s *AVFormatContext) NbPrograms() uint {
	value := s.ptr.nb_programs
	return uint(value)
}

func (s *AVFormatContext) SetNbPrograms(value uint) {
	s.ptr.nb_programs = (C.uint)(value)
}

func (s *AVFormatContext) ProgramsEntry(i uint) *AVProgram {
	value := s.ptr.programs
	ptr := arrayGet[*C.AVProgram](value, uintptr(i))
	var valueMapped *AVProgram
	if ptr != nil {
		valueMapped = &AVProgram{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFormatContext) VideoCodecId() AVCodecID {
	value := s.ptr.video_codec_id
	return AVCodecID(value)
}

func (s *AVFormatContext) SetVideoCodecId(value AVCodecID) {
	s.ptr.video_codec_id = (C.enum_AVCodecID)(value)
}

func (s *AVFormatContext) AudioCodecId() AVCodecID {
	value := s.ptr.audio_codec_id
	return AVCodecID(value)
}

func (s *AVFormatContext) SetAudioCodecId(value AVCodecID) {
	s.ptr.audio_codec_id = (C.enum_AVCodecID)(value)
}

func (s *AVFormatContext) SubtitleCodecId() AVCodecID {
	value := s.ptr.subtitle_codec_id
	return AVCodecID(value)
}

func (s *AVFormatContext) SetSubtitleCodecId(value AVCodecID) {
	s.ptr.subtitle_codec_id = (C.enum_AVCodecID)(value)
}

func (s *AVFormatContext) MaxIndexSize() uint {
	value := s.ptr.max_index_size
	return uint(value)
}

func (s *AVFormatContext) SetMaxIndexSize(value uint) {
	s.ptr.max_index_size = (C.uint)(value)
}

func (s *AVFormatContext) MaxPictureBuffer() uint {
	value := s.ptr.max_picture_buffer
	return uint(value)
}

func (s *AVFormatContext) SetMaxPictureBuffer(value uint) {
	s.ptr.max_picture_buffer = (C.uint)(value)
}

func (s *AVFormatContext) NbChapters() uint {
	value := s.ptr.nb_chapters
	return uint(value)
}

func (s *AVFormatContext) SetNbChapters(value uint) {
	s.ptr.nb_chapters = (C.uint)(value)
}

func (s *AVFormatContext) ChaptersEntry(i uint) *AVChapter {
	value := s.ptr.chapters
	ptr := arrayGet[*C.AVChapter](value, uintptr(i))
	var valueMapped *AVChapter
	if ptr != nil {
		valueMapped = &AVChapter{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFormatContext) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

func (s *AVFormatContext) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

func (s *AVFormatContext) StartTimeRealtime() int64 {
	value := s.ptr.start_time_realtime
	return int64(value)
}

func (s *AVFormatContext) SetStartTimeRealtime(value int64) {
	s.ptr.start_time_realtime = (C.int64_t)(value)
}

func (s *AVFormatContext) FpsProbeSize() int {
	value := s.ptr.fps_probe_size
	return int(value)
}

func (s *AVFormatContext) SetFpsProbeSize(value int) {
	s.ptr.fps_probe_size = (C.int)(value)
}

func (s *AVFormatContext) ErrorRecognition() int {
	value := s.ptr.error_recognition
	return int(value)
}

func (s *AVFormatContext) SetErrorRecognition(value int) {
	s.ptr.error_recognition = (C.int)(value)
}

// InterruptCallback skipped due to ident byvalue

func (s *AVFormatContext) Debug() int {
	value := s.ptr.debug
	return int(value)
}

func (s *AVFormatContext) SetDebug(value int) {
	s.ptr.debug = (C.int)(value)
}

func (s *AVFormatContext) MaxInterleaveDelta() int64 {
	value := s.ptr.max_interleave_delta
	return int64(value)
}

func (s *AVFormatContext) SetMaxInterleaveDelta(value int64) {
	s.ptr.max_interleave_delta = (C.int64_t)(value)
}

func (s *AVFormatContext) StrictStdCompliance() int {
	value := s.ptr.strict_std_compliance
	return int(value)
}

func (s *AVFormatContext) SetStrictStdCompliance(value int) {
	s.ptr.strict_std_compliance = (C.int)(value)
}

func (s *AVFormatContext) EventFlags() int {
	value := s.ptr.event_flags
	return int(value)
}

func (s *AVFormatContext) SetEventFlags(value int) {
	s.ptr.event_flags = (C.int)(value)
}

func (s *AVFormatContext) MaxTsProbe() int {
	value := s.ptr.max_ts_probe
	return int(value)
}

func (s *AVFormatContext) SetMaxTsProbe(value int) {
	s.ptr.max_ts_probe = (C.int)(value)
}

func (s *AVFormatContext) AvoidNegativeTs() int {
	value := s.ptr.avoid_negative_ts
	return int(value)
}

func (s *AVFormatContext) SetAvoidNegativeTs(value int) {
	s.ptr.avoid_negative_ts = (C.int)(value)
}

func (s *AVFormatContext) TsId() int {
	value := s.ptr.ts_id
	return int(value)
}

func (s *AVFormatContext) SetTsId(value int) {
	s.ptr.ts_id = (C.int)(value)
}

func (s *AVFormatContext) AudioPreload() int {
	value := s.ptr.audio_preload
	return int(value)
}

func (s *AVFormatContext) SetAudioPreload(value int) {
	s.ptr.audio_preload = (C.int)(value)
}

func (s *AVFormatContext) MaxChunkDuration() int {
	value := s.ptr.max_chunk_duration
	return int(value)
}

func (s *AVFormatContext) SetMaxChunkDuration(value int) {
	s.ptr.max_chunk_duration = (C.int)(value)
}

func (s *AVFormatContext) MaxChunkSize() int {
	value := s.ptr.max_chunk_size
	return int(value)
}

func (s *AVFormatContext) SetMaxChunkSize(value int) {
	s.ptr.max_chunk_size = (C.int)(value)
}

func (s *AVFormatContext) UseWallclockAsTimestamps() int {
	value := s.ptr.use_wallclock_as_timestamps
	return int(value)
}

func (s *AVFormatContext) SetUseWallclockAsTimestamps(value int) {
	s.ptr.use_wallclock_as_timestamps = (C.int)(value)
}

func (s *AVFormatContext) AvioFlags() int {
	value := s.ptr.avio_flags
	return int(value)
}

func (s *AVFormatContext) SetAvioFlags(value int) {
	s.ptr.avio_flags = (C.int)(value)
}

func (s *AVFormatContext) DurationEstimationMethod() AVDurationEstimationMethod {
	value := s.ptr.duration_estimation_method
	return AVDurationEstimationMethod(value)
}

func (s *AVFormatContext) SetDurationEstimationMethod(value AVDurationEstimationMethod) {
	s.ptr.duration_estimation_method = (C.enum_AVDurationEstimationMethod)(value)
}

func (s *AVFormatContext) SkipInitialBytes() int64 {
	value := s.ptr.skip_initial_bytes
	return int64(value)
}

func (s *AVFormatContext) SetSkipInitialBytes(value int64) {
	s.ptr.skip_initial_bytes = (C.int64_t)(value)
}

func (s *AVFormatContext) CorrectTsOverflow() uint {
	value := s.ptr.correct_ts_overflow
	return uint(value)
}

func (s *AVFormatContext) SetCorrectTsOverflow(value uint) {
	s.ptr.correct_ts_overflow = (C.uint)(value)
}

func (s *AVFormatContext) Seek2Any() int {
	value := s.ptr.seek2any
	return int(value)
}

func (s *AVFormatContext) SetSeek2Any(value int) {
	s.ptr.seek2any = (C.int)(value)
}

func (s *AVFormatContext) FlushPackets() int {
	value := s.ptr.flush_packets
	return int(value)
}

func (s *AVFormatContext) SetFlushPackets(value int) {
	s.ptr.flush_packets = (C.int)(value)
}

func (s *AVFormatContext) ProbeScore() int {
	value := s.ptr.probe_score
	return int(value)
}

func (s *AVFormatContext) SetProbeScore(value int) {
	s.ptr.probe_score = (C.int)(value)
}

func (s *AVFormatContext) FormatProbesize() int {
	value := s.ptr.format_probesize
	return int(value)
}

func (s *AVFormatContext) SetFormatProbesize(value int) {
	s.ptr.format_probesize = (C.int)(value)
}

func (s *AVFormatContext) CodecWhitelist() *CStr {
	value := s.ptr.codec_whitelist
	return wrapCStr(value)
}

func (s *AVFormatContext) SetCodecWhitelist(value *CStr) {
	s.ptr.codec_whitelist = value.ptr
}

func (s *AVFormatContext) FormatWhitelist() *CStr {
	value := s.ptr.format_whitelist
	return wrapCStr(value)
}

func (s *AVFormatContext) SetFormatWhitelist(value *CStr) {
	s.ptr.format_whitelist = value.ptr
}

func (s *AVFormatContext) IoRepositioned() int {
	value := s.ptr.io_repositioned
	return int(value)
}

func (s *AVFormatContext) SetIoRepositioned(value int) {
	s.ptr.io_repositioned = (C.int)(value)
}

func (s *AVFormatContext) VideoCodec() *AVCodec {
	value := s.ptr.video_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

func (s *AVFormatContext) SetVideoCodec(value *AVCodec) {
	if value != nil {
		s.ptr.video_codec = value.ptr
	} else {
		s.ptr.video_codec = nil
	}
}

func (s *AVFormatContext) AudioCodec() *AVCodec {
	value := s.ptr.audio_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

func (s *AVFormatContext) SetAudioCodec(value *AVCodec) {
	if value != nil {
		s.ptr.audio_codec = value.ptr
	} else {
		s.ptr.audio_codec = nil
	}
}

func (s *AVFormatContext) SubtitleCodec() *AVCodec {
	value := s.ptr.subtitle_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

func (s *AVFormatContext) SetSubtitleCodec(value *AVCodec) {
	if value != nil {
		s.ptr.subtitle_codec = value.ptr
	} else {
		s.ptr.subtitle_codec = nil
	}
}

func (s *AVFormatContext) DataCodec() *AVCodec {
	value := s.ptr.data_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

func (s *AVFormatContext) SetDataCodec(value *AVCodec) {
	if value != nil {
		s.ptr.data_codec = value.ptr
	} else {
		s.ptr.data_codec = nil
	}
}

func (s *AVFormatContext) MetadataHeaderPadding() int {
	value := s.ptr.metadata_header_padding
	return int(value)
}

func (s *AVFormatContext) SetMetadataHeaderPadding(value int) {
	s.ptr.metadata_header_padding = (C.int)(value)
}

func (s *AVFormatContext) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

func (s *AVFormatContext) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// ControlMessageCb skipped due to ident callback

func (s *AVFormatContext) OutputTsOffset() int64 {
	value := s.ptr.output_ts_offset
	return int64(value)
}

func (s *AVFormatContext) SetOutputTsOffset(value int64) {
	s.ptr.output_ts_offset = (C.int64_t)(value)
}

func (s *AVFormatContext) DumpSeparator() unsafe.Pointer {
	value := s.ptr.dump_separator
	return unsafe.Pointer(value)
}

func (s *AVFormatContext) SetDumpSeparator(value unsafe.Pointer) {
	s.ptr.dump_separator = (*C.uint8_t)(value)
}

func (s *AVFormatContext) DataCodecId() AVCodecID {
	value := s.ptr.data_codec_id
	return AVCodecID(value)
}

func (s *AVFormatContext) SetDataCodecId(value AVCodecID) {
	s.ptr.data_codec_id = (C.enum_AVCodecID)(value)
}

func (s *AVFormatContext) ProtocolWhitelist() *CStr {
	value := s.ptr.protocol_whitelist
	return wrapCStr(value)
}

func (s *AVFormatContext) SetProtocolWhitelist(value *CStr) {
	s.ptr.protocol_whitelist = value.ptr
}

// IoOpen skipped due to func ptr

// IoClose skipped due to func ptr

func (s *AVFormatContext) ProtocolBlacklist() *CStr {
	value := s.ptr.protocol_blacklist
	return wrapCStr(value)
}

func (s *AVFormatContext) SetProtocolBlacklist(value *CStr) {
	s.ptr.protocol_blacklist = value.ptr
}

func (s *AVFormatContext) MaxStreams() int {
	value := s.ptr.max_streams
	return int(value)
}

func (s *AVFormatContext) SetMaxStreams(value int) {
	s.ptr.max_streams = (C.int)(value)
}

func (s *AVFormatContext) SkipEstimateDurationFromPts() int {
	value := s.ptr.skip_estimate_duration_from_pts
	return int(value)
}

func (s *AVFormatContext) SetSkipEstimateDurationFromPts(value int) {
	s.ptr.skip_estimate_duration_from_pts = (C.int)(value)
}

func (s *AVFormatContext) MaxProbePackets() int {
	value := s.ptr.max_probe_packets
	return int(value)
}

func (s *AVFormatContext) SetMaxProbePackets(value int) {
	s.ptr.max_probe_packets = (C.int)(value)
}

// IoClose2 skipped due to func ptr

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

// Callback skipped due to func ptr

func (s *AVIOInterruptCB) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

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

func (s *AVIODirEntry) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

func (s *AVIODirEntry) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

func (s *AVIODirEntry) Type() int {
	value := s.ptr._type
	return int(value)
}

func (s *AVIODirEntry) SetType(value int) {
	s.ptr._type = (C.int)(value)
}

func (s *AVIODirEntry) Utf8() int {
	value := s.ptr.utf8
	return int(value)
}

func (s *AVIODirEntry) SetUtf8(value int) {
	s.ptr.utf8 = (C.int)(value)
}

func (s *AVIODirEntry) Size() int64 {
	value := s.ptr.size
	return int64(value)
}

func (s *AVIODirEntry) SetSize(value int64) {
	s.ptr.size = (C.int64_t)(value)
}

func (s *AVIODirEntry) ModificationTimestamp() int64 {
	value := s.ptr.modification_timestamp
	return int64(value)
}

func (s *AVIODirEntry) SetModificationTimestamp(value int64) {
	s.ptr.modification_timestamp = (C.int64_t)(value)
}

func (s *AVIODirEntry) AccessTimestamp() int64 {
	value := s.ptr.access_timestamp
	return int64(value)
}

func (s *AVIODirEntry) SetAccessTimestamp(value int64) {
	s.ptr.access_timestamp = (C.int64_t)(value)
}

func (s *AVIODirEntry) StatusChangeTimestamp() int64 {
	value := s.ptr.status_change_timestamp
	return int64(value)
}

func (s *AVIODirEntry) SetStatusChangeTimestamp(value int64) {
	s.ptr.status_change_timestamp = (C.int64_t)(value)
}

func (s *AVIODirEntry) UserId() int64 {
	value := s.ptr.user_id
	return int64(value)
}

func (s *AVIODirEntry) SetUserId(value int64) {
	s.ptr.user_id = (C.int64_t)(value)
}

func (s *AVIODirEntry) GroupId() int64 {
	value := s.ptr.group_id
	return int64(value)
}

func (s *AVIODirEntry) SetGroupId(value int64) {
	s.ptr.group_id = (C.int64_t)(value)
}

func (s *AVIODirEntry) Filemode() int64 {
	value := s.ptr.filemode
	return int64(value)
}

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

// UrlContext skipped due to ptr to ignored type

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

func (s *AVIOContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

func (s *AVIOContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Buffer skipped due to prim ptr

func (s *AVIOContext) BufferSize() int {
	value := s.ptr.buffer_size
	return int(value)
}

func (s *AVIOContext) SetBufferSize(value int) {
	s.ptr.buffer_size = (C.int)(value)
}

// BufPtr skipped due to prim ptr

// BufEnd skipped due to prim ptr

func (s *AVIOContext) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

func (s *AVIOContext) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// ReadPacket skipped due to func ptr

// WritePacket skipped due to func ptr

// Seek skipped due to func ptr

func (s *AVIOContext) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}

func (s *AVIOContext) SetPos(value int64) {
	s.ptr.pos = (C.int64_t)(value)
}

func (s *AVIOContext) EofReached() int {
	value := s.ptr.eof_reached
	return int(value)
}

func (s *AVIOContext) SetEofReached(value int) {
	s.ptr.eof_reached = (C.int)(value)
}

func (s *AVIOContext) Error() int {
	value := s.ptr.error
	return int(value)
}

func (s *AVIOContext) SetError(value int) {
	s.ptr.error = (C.int)(value)
}

func (s *AVIOContext) WriteFlag() int {
	value := s.ptr.write_flag
	return int(value)
}

func (s *AVIOContext) SetWriteFlag(value int) {
	s.ptr.write_flag = (C.int)(value)
}

func (s *AVIOContext) MaxPacketSize() int {
	value := s.ptr.max_packet_size
	return int(value)
}

func (s *AVIOContext) SetMaxPacketSize(value int) {
	s.ptr.max_packet_size = (C.int)(value)
}

func (s *AVIOContext) MinPacketSize() int {
	value := s.ptr.min_packet_size
	return int(value)
}

func (s *AVIOContext) SetMinPacketSize(value int) {
	s.ptr.min_packet_size = (C.int)(value)
}

func (s *AVIOContext) Checksum() uint32 {
	value := s.ptr.checksum
	return uint32(value)
}

func (s *AVIOContext) SetChecksum(value uint32) {
	s.ptr.checksum = (C.ulong)(value)
}

// ChecksumPtr skipped due to prim ptr

// UpdateChecksum skipped due to func ptr

// ReadPause skipped due to func ptr

// ReadSeek skipped due to func ptr

func (s *AVIOContext) Seekable() int {
	value := s.ptr.seekable
	return int(value)
}

func (s *AVIOContext) SetSeekable(value int) {
	s.ptr.seekable = (C.int)(value)
}

func (s *AVIOContext) Direct() int {
	value := s.ptr.direct
	return int(value)
}

func (s *AVIOContext) SetDirect(value int) {
	s.ptr.direct = (C.int)(value)
}

func (s *AVIOContext) ProtocolWhitelist() *CStr {
	value := s.ptr.protocol_whitelist
	return wrapCStr(value)
}

func (s *AVIOContext) SetProtocolWhitelist(value *CStr) {
	s.ptr.protocol_whitelist = value.ptr
}

func (s *AVIOContext) ProtocolBlacklist() *CStr {
	value := s.ptr.protocol_blacklist
	return wrapCStr(value)
}

func (s *AVIOContext) SetProtocolBlacklist(value *CStr) {
	s.ptr.protocol_blacklist = value.ptr
}

// WriteDataType skipped due to func ptr

func (s *AVIOContext) IgnoreBoundaryPoint() int {
	value := s.ptr.ignore_boundary_point
	return int(value)
}

func (s *AVIOContext) SetIgnoreBoundaryPoint(value int) {
	s.ptr.ignore_boundary_point = (C.int)(value)
}

// BufPtrMax skipped due to prim ptr

func (s *AVIOContext) BytesRead() int64 {
	value := s.ptr.bytes_read
	return int64(value)
}

func (s *AVIOContext) SetBytesRead(value int64) {
	s.ptr.bytes_read = (C.int64_t)(value)
}

func (s *AVIOContext) BytesWritten() int64 {
	value := s.ptr.bytes_written
	return int64(value)
}

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

func (s *AVBufferRef) Buffer() *AVBuffer {
	value := s.ptr.buffer
	var valueMapped *AVBuffer
	if value != nil {
		valueMapped = &AVBuffer{ptr: value}
	}
	return valueMapped
}

func (s *AVBufferRef) SetBuffer(value *AVBuffer) {
	if value != nil {
		s.ptr.buffer = value.ptr
	} else {
		s.ptr.buffer = nil
	}
}

func (s *AVBufferRef) Data() unsafe.Pointer {
	value := s.ptr.data
	return unsafe.Pointer(value)
}

func (s *AVBufferRef) SetData(value unsafe.Pointer) {
	s.ptr.data = (*C.uint8_t)(value)
}

func (s *AVBufferRef) Size() uint64 {
	value := s.ptr.size
	return uint64(value)
}

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

func (s *AVChannelCustom) Id() AVChannel {
	value := s.ptr.id
	return AVChannel(value)
}

func (s *AVChannelCustom) SetId(value AVChannel) {
	s.ptr.id = (C.enum_AVChannel)(value)
}

func (s *AVChannelCustom) NameEntry(i uint) uint8 {
	value := s.ptr.name[i]
	return uint8(value)
}

func (s *AVChannelCustom) SetNameEntry(i uint, value uint8) {
	s.ptr.name[i] = (C.char)(value)
}

func (s *AVChannelCustom) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

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

func (s *AVChannelLayout) Order() AVChannelOrder {
	value := s.ptr.order
	return AVChannelOrder(value)
}

func (s *AVChannelLayout) SetOrder(value AVChannelOrder) {
	s.ptr.order = (C.enum_AVChannelOrder)(value)
}

func (s *AVChannelLayout) NbChannels() int {
	value := s.ptr.nb_channels
	return int(value)
}

func (s *AVChannelLayout) SetNbChannels(value int) {
	s.ptr.nb_channels = (C.int)(value)
}

// U skipped due to union type

func (s *AVChannelLayout) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

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

// --- Struct AVDictionaryEntry ---

// AVDictionaryEntry wraps AVDictionaryEntry.
type AVDictionaryEntry struct {
	ptr *C.AVDictionaryEntry
}

func (s *AVDictionaryEntry) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

func (s *AVDictionaryEntry) Key() *CStr {
	value := s.ptr.key
	return wrapCStr(value)
}

func (s *AVDictionaryEntry) SetKey(value *CStr) {
	s.ptr.key = value.ptr
}

func (s *AVDictionaryEntry) Value() *CStr {
	value := s.ptr.value
	return wrapCStr(value)
}

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

func (s *AVFrameSideData) Type() AVFrameSideDataType {
	value := s.ptr._type
	return AVFrameSideDataType(value)
}

func (s *AVFrameSideData) SetType(value AVFrameSideDataType) {
	s.ptr._type = (C.enum_AVFrameSideDataType)(value)
}

func (s *AVFrameSideData) Data() unsafe.Pointer {
	value := s.ptr.data
	return unsafe.Pointer(value)
}

func (s *AVFrameSideData) SetData(value unsafe.Pointer) {
	s.ptr.data = (*C.uint8_t)(value)
}

func (s *AVFrameSideData) Size() uint64 {
	value := s.ptr.size
	return uint64(value)
}

func (s *AVFrameSideData) SetSize(value uint64) {
	s.ptr.size = (C.size_t)(value)
}

func (s *AVFrameSideData) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

func (s *AVFrameSideData) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

func (s *AVFrameSideData) Buf() *AVBufferRef {
	value := s.ptr.buf
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

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

func (s *AVRegionOfInterest) SelfSize() uint32 {
	value := s.ptr.self_size
	return uint32(value)
}

func (s *AVRegionOfInterest) SetSelfSize(value uint32) {
	s.ptr.self_size = (C.uint32_t)(value)
}

func (s *AVRegionOfInterest) Top() int {
	value := s.ptr.top
	return int(value)
}

func (s *AVRegionOfInterest) SetTop(value int) {
	s.ptr.top = (C.int)(value)
}

func (s *AVRegionOfInterest) Bottom() int {
	value := s.ptr.bottom
	return int(value)
}

func (s *AVRegionOfInterest) SetBottom(value int) {
	s.ptr.bottom = (C.int)(value)
}

func (s *AVRegionOfInterest) Left() int {
	value := s.ptr.left
	return int(value)
}

func (s *AVRegionOfInterest) SetLeft(value int) {
	s.ptr.left = (C.int)(value)
}

func (s *AVRegionOfInterest) Right() int {
	value := s.ptr.right
	return int(value)
}

func (s *AVRegionOfInterest) SetRight(value int) {
	s.ptr.right = (C.int)(value)
}

func (s *AVRegionOfInterest) Qoffset() *AVRational {
	value := s.ptr.qoffset
	return &AVRational{value: value}
}

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

func (s *AVFrame) DataEntry(i uint) unsafe.Pointer {
	value := s.ptr.data[i]
	return unsafe.Pointer(value)
}

func (s *AVFrame) SetDataEntry(i uint, value unsafe.Pointer) {
	s.ptr.data[i] = (*C.uint8_t)(value)
}

func (s *AVFrame) LinesizeEntry(i uint) int {
	value := s.ptr.linesize[i]
	return int(value)
}

func (s *AVFrame) SetLinesizeEntry(i uint, value int) {
	s.ptr.linesize[i] = (C.int)(value)
}

// ExtendedData skipped due to unknown ptr ptr

func (s *AVFrame) Width() int {
	value := s.ptr.width
	return int(value)
}

func (s *AVFrame) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

func (s *AVFrame) Height() int {
	value := s.ptr.height
	return int(value)
}

func (s *AVFrame) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

func (s *AVFrame) NbSamples() int {
	value := s.ptr.nb_samples
	return int(value)
}

func (s *AVFrame) SetNbSamples(value int) {
	s.ptr.nb_samples = (C.int)(value)
}

func (s *AVFrame) Format() int {
	value := s.ptr.format
	return int(value)
}

func (s *AVFrame) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

func (s *AVFrame) KeyFrame() int {
	value := s.ptr.key_frame
	return int(value)
}

func (s *AVFrame) SetKeyFrame(value int) {
	s.ptr.key_frame = (C.int)(value)
}

func (s *AVFrame) PictType() AVPictureType {
	value := s.ptr.pict_type
	return AVPictureType(value)
}

func (s *AVFrame) SetPictType(value AVPictureType) {
	s.ptr.pict_type = (C.enum_AVPictureType)(value)
}

func (s *AVFrame) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

func (s *AVFrame) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

func (s *AVFrame) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

func (s *AVFrame) SetPts(value int64) {
	s.ptr.pts = (C.int64_t)(value)
}

func (s *AVFrame) PktDts() int64 {
	value := s.ptr.pkt_dts
	return int64(value)
}

func (s *AVFrame) SetPktDts(value int64) {
	s.ptr.pkt_dts = (C.int64_t)(value)
}

func (s *AVFrame) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

func (s *AVFrame) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

func (s *AVFrame) CodedPictureNumber() int {
	value := s.ptr.coded_picture_number
	return int(value)
}

func (s *AVFrame) SetCodedPictureNumber(value int) {
	s.ptr.coded_picture_number = (C.int)(value)
}

func (s *AVFrame) DisplayPictureNumber() int {
	value := s.ptr.display_picture_number
	return int(value)
}

func (s *AVFrame) SetDisplayPictureNumber(value int) {
	s.ptr.display_picture_number = (C.int)(value)
}

func (s *AVFrame) Quality() int {
	value := s.ptr.quality
	return int(value)
}

func (s *AVFrame) SetQuality(value int) {
	s.ptr.quality = (C.int)(value)
}

func (s *AVFrame) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

func (s *AVFrame) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

func (s *AVFrame) RepeatPict() int {
	value := s.ptr.repeat_pict
	return int(value)
}

func (s *AVFrame) SetRepeatPict(value int) {
	s.ptr.repeat_pict = (C.int)(value)
}

func (s *AVFrame) InterlacedFrame() int {
	value := s.ptr.interlaced_frame
	return int(value)
}

func (s *AVFrame) SetInterlacedFrame(value int) {
	s.ptr.interlaced_frame = (C.int)(value)
}

func (s *AVFrame) TopFieldFirst() int {
	value := s.ptr.top_field_first
	return int(value)
}

func (s *AVFrame) SetTopFieldFirst(value int) {
	s.ptr.top_field_first = (C.int)(value)
}

func (s *AVFrame) PaletteHasChanged() int {
	value := s.ptr.palette_has_changed
	return int(value)
}

func (s *AVFrame) SetPaletteHasChanged(value int) {
	s.ptr.palette_has_changed = (C.int)(value)
}

func (s *AVFrame) ReorderedOpaque() int64 {
	value := s.ptr.reordered_opaque
	return int64(value)
}

func (s *AVFrame) SetReorderedOpaque(value int64) {
	s.ptr.reordered_opaque = (C.int64_t)(value)
}

func (s *AVFrame) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

func (s *AVFrame) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

func (s *AVFrame) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

func (s *AVFrame) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

func (s *AVFrame) BufEntry(i uint) *AVBufferRef {
	value := s.ptr.buf[i]
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVFrame) SetBufEntry(i uint, value *AVBufferRef) {
	if value != nil {
		s.ptr.buf[i] = value.ptr
	} else {
		s.ptr.buf[i] = nil
	}
}

func (s *AVFrame) ExtendedBufEntry(i uint) *AVBufferRef {
	value := s.ptr.extended_buf
	ptr := arrayGet[*C.AVBufferRef](value, uintptr(i))
	var valueMapped *AVBufferRef
	if ptr != nil {
		valueMapped = &AVBufferRef{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFrame) NbExtendedBuf() int {
	value := s.ptr.nb_extended_buf
	return int(value)
}

func (s *AVFrame) SetNbExtendedBuf(value int) {
	s.ptr.nb_extended_buf = (C.int)(value)
}

func (s *AVFrame) SideDataEntry(i uint) *AVFrameSideData {
	value := s.ptr.side_data
	ptr := arrayGet[*C.AVFrameSideData](value, uintptr(i))
	var valueMapped *AVFrameSideData
	if ptr != nil {
		valueMapped = &AVFrameSideData{ptr: ptr}
	}
	return valueMapped
}

func (s *AVFrame) NbSideData() int {
	value := s.ptr.nb_side_data
	return int(value)
}

func (s *AVFrame) SetNbSideData(value int) {
	s.ptr.nb_side_data = (C.int)(value)
}

func (s *AVFrame) Flags() int {
	value := s.ptr.flags
	return int(value)
}

func (s *AVFrame) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

func (s *AVFrame) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}

func (s *AVFrame) SetColorRange(value AVColorRange) {
	s.ptr.color_range = (C.enum_AVColorRange)(value)
}

func (s *AVFrame) ColorPrimaries() AVColorPrimaries {
	value := s.ptr.color_primaries
	return AVColorPrimaries(value)
}

func (s *AVFrame) SetColorPrimaries(value AVColorPrimaries) {
	s.ptr.color_primaries = (C.enum_AVColorPrimaries)(value)
}

func (s *AVFrame) ColorTrc() AVColorTransferCharacteristic {
	value := s.ptr.color_trc
	return AVColorTransferCharacteristic(value)
}

func (s *AVFrame) SetColorTrc(value AVColorTransferCharacteristic) {
	s.ptr.color_trc = (C.enum_AVColorTransferCharacteristic)(value)
}

func (s *AVFrame) Colorspace() AVColorSpace {
	value := s.ptr.colorspace
	return AVColorSpace(value)
}

func (s *AVFrame) SetColorspace(value AVColorSpace) {
	s.ptr.colorspace = (C.enum_AVColorSpace)(value)
}

func (s *AVFrame) ChromaLocation() AVChromaLocation {
	value := s.ptr.chroma_location
	return AVChromaLocation(value)
}

func (s *AVFrame) SetChromaLocation(value AVChromaLocation) {
	s.ptr.chroma_location = (C.enum_AVChromaLocation)(value)
}

func (s *AVFrame) BestEffortTimestamp() int64 {
	value := s.ptr.best_effort_timestamp
	return int64(value)
}

func (s *AVFrame) SetBestEffortTimestamp(value int64) {
	s.ptr.best_effort_timestamp = (C.int64_t)(value)
}

func (s *AVFrame) PktPos() int64 {
	value := s.ptr.pkt_pos
	return int64(value)
}

func (s *AVFrame) SetPktPos(value int64) {
	s.ptr.pkt_pos = (C.int64_t)(value)
}

func (s *AVFrame) PktDuration() int64 {
	value := s.ptr.pkt_duration
	return int64(value)
}

func (s *AVFrame) SetPktDuration(value int64) {
	s.ptr.pkt_duration = (C.int64_t)(value)
}

func (s *AVFrame) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

func (s *AVFrame) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

func (s *AVFrame) DecodeErrorFlags() int {
	value := s.ptr.decode_error_flags
	return int(value)
}

func (s *AVFrame) SetDecodeErrorFlags(value int) {
	s.ptr.decode_error_flags = (C.int)(value)
}

func (s *AVFrame) Channels() int {
	value := s.ptr.channels
	return int(value)
}

func (s *AVFrame) SetChannels(value int) {
	s.ptr.channels = (C.int)(value)
}

func (s *AVFrame) PktSize() int {
	value := s.ptr.pkt_size
	return int(value)
}

func (s *AVFrame) SetPktSize(value int) {
	s.ptr.pkt_size = (C.int)(value)
}

func (s *AVFrame) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVFrame) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

func (s *AVFrame) OpaqueRef() *AVBufferRef {
	value := s.ptr.opaque_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVFrame) SetOpaqueRef(value *AVBufferRef) {
	if value != nil {
		s.ptr.opaque_ref = value.ptr
	} else {
		s.ptr.opaque_ref = nil
	}
}

func (s *AVFrame) CropTop() uint64 {
	value := s.ptr.crop_top
	return uint64(value)
}

func (s *AVFrame) SetCropTop(value uint64) {
	s.ptr.crop_top = (C.size_t)(value)
}

func (s *AVFrame) CropBottom() uint64 {
	value := s.ptr.crop_bottom
	return uint64(value)
}

func (s *AVFrame) SetCropBottom(value uint64) {
	s.ptr.crop_bottom = (C.size_t)(value)
}

func (s *AVFrame) CropLeft() uint64 {
	value := s.ptr.crop_left
	return uint64(value)
}

func (s *AVFrame) SetCropLeft(value uint64) {
	s.ptr.crop_left = (C.size_t)(value)
}

func (s *AVFrame) CropRight() uint64 {
	value := s.ptr.crop_right
	return uint64(value)
}

func (s *AVFrame) SetCropRight(value uint64) {
	s.ptr.crop_right = (C.size_t)(value)
}

func (s *AVFrame) PrivateRef() *AVBufferRef {
	value := s.ptr.private_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVFrame) SetPrivateRef(value *AVBufferRef) {
	if value != nil {
		s.ptr.private_ref = value.ptr
	} else {
		s.ptr.private_ref = nil
	}
}

// ChLayout skipped due to ident byvalue

func (s *AVFrame) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

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

func (s *AVHWDeviceContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

func (s *AVHWDeviceContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

func (s *AVHWDeviceContext) Internal() *AVHWDeviceInternal {
	value := s.ptr.internal
	var valueMapped *AVHWDeviceInternal
	if value != nil {
		valueMapped = &AVHWDeviceInternal{ptr: value}
	}
	return valueMapped
}

func (s *AVHWDeviceContext) SetInternal(value *AVHWDeviceInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

func (s *AVHWDeviceContext) Type() AVHWDeviceType {
	value := s.ptr._type
	return AVHWDeviceType(value)
}

func (s *AVHWDeviceContext) SetType(value AVHWDeviceType) {
	s.ptr._type = (C.enum_AVHWDeviceType)(value)
}

func (s *AVHWDeviceContext) Hwctx() unsafe.Pointer {
	value := s.ptr.hwctx
	return value
}

func (s *AVHWDeviceContext) SetHwctx(value unsafe.Pointer) {
	s.ptr.hwctx = value
}

// Free skipped due to func ptr

func (s *AVHWDeviceContext) UserOpaque() unsafe.Pointer {
	value := s.ptr.user_opaque
	return value
}

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

func (s *AVHWFramesContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

func (s *AVHWFramesContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

func (s *AVHWFramesContext) Internal() *AVHWFramesInternal {
	value := s.ptr.internal
	var valueMapped *AVHWFramesInternal
	if value != nil {
		valueMapped = &AVHWFramesInternal{ptr: value}
	}
	return valueMapped
}

func (s *AVHWFramesContext) SetInternal(value *AVHWFramesInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

func (s *AVHWFramesContext) DeviceRef() *AVBufferRef {
	value := s.ptr.device_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

func (s *AVHWFramesContext) SetDeviceRef(value *AVBufferRef) {
	if value != nil {
		s.ptr.device_ref = value.ptr
	} else {
		s.ptr.device_ref = nil
	}
}

func (s *AVHWFramesContext) DeviceCtx() *AVHWDeviceContext {
	value := s.ptr.device_ctx
	var valueMapped *AVHWDeviceContext
	if value != nil {
		valueMapped = &AVHWDeviceContext{ptr: value}
	}
	return valueMapped
}

func (s *AVHWFramesContext) SetDeviceCtx(value *AVHWDeviceContext) {
	if value != nil {
		s.ptr.device_ctx = value.ptr
	} else {
		s.ptr.device_ctx = nil
	}
}

func (s *AVHWFramesContext) Hwctx() unsafe.Pointer {
	value := s.ptr.hwctx
	return value
}

func (s *AVHWFramesContext) SetHwctx(value unsafe.Pointer) {
	s.ptr.hwctx = value
}

// Free skipped due to func ptr

func (s *AVHWFramesContext) UserOpaque() unsafe.Pointer {
	value := s.ptr.user_opaque
	return value
}

func (s *AVHWFramesContext) SetUserOpaque(value unsafe.Pointer) {
	s.ptr.user_opaque = value
}

func (s *AVHWFramesContext) Pool() *AVBufferPool {
	value := s.ptr.pool
	var valueMapped *AVBufferPool
	if value != nil {
		valueMapped = &AVBufferPool{ptr: value}
	}
	return valueMapped
}

func (s *AVHWFramesContext) SetPool(value *AVBufferPool) {
	if value != nil {
		s.ptr.pool = value.ptr
	} else {
		s.ptr.pool = nil
	}
}

func (s *AVHWFramesContext) InitialPoolSize() int {
	value := s.ptr.initial_pool_size
	return int(value)
}

func (s *AVHWFramesContext) SetInitialPoolSize(value int) {
	s.ptr.initial_pool_size = (C.int)(value)
}

func (s *AVHWFramesContext) Format() AVPixelFormat {
	value := s.ptr.format
	return AVPixelFormat(value)
}

func (s *AVHWFramesContext) SetFormat(value AVPixelFormat) {
	s.ptr.format = (C.enum_AVPixelFormat)(value)
}

func (s *AVHWFramesContext) SwFormat() AVPixelFormat {
	value := s.ptr.sw_format
	return AVPixelFormat(value)
}

func (s *AVHWFramesContext) SetSwFormat(value AVPixelFormat) {
	s.ptr.sw_format = (C.enum_AVPixelFormat)(value)
}

func (s *AVHWFramesContext) Width() int {
	value := s.ptr.width
	return int(value)
}

func (s *AVHWFramesContext) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

func (s *AVHWFramesContext) Height() int {
	value := s.ptr.height
	return int(value)
}

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

func (s *AVHWFramesConstraints) ValidHwFormats() *Array[AVPixelFormat] {
	value := s.ptr.valid_hw_formats
	return ToAVPixelFormatArray(unsafe.Pointer(value))
}

func (s *AVHWFramesConstraints) SetValidHwFormats(value *Array[AVPixelFormat]) {
	if value != nil {
		s.ptr.valid_hw_formats = (*C.enum_AVPixelFormat)(value.ptr)
	} else {
		s.ptr.valid_hw_formats = nil
	}
}

func (s *AVHWFramesConstraints) ValidSwFormats() *Array[AVPixelFormat] {
	value := s.ptr.valid_sw_formats
	return ToAVPixelFormatArray(unsafe.Pointer(value))
}

func (s *AVHWFramesConstraints) SetValidSwFormats(value *Array[AVPixelFormat]) {
	if value != nil {
		s.ptr.valid_sw_formats = (*C.enum_AVPixelFormat)(value.ptr)
	} else {
		s.ptr.valid_sw_formats = nil
	}
}

func (s *AVHWFramesConstraints) MinWidth() int {
	value := s.ptr.min_width
	return int(value)
}

func (s *AVHWFramesConstraints) SetMinWidth(value int) {
	s.ptr.min_width = (C.int)(value)
}

func (s *AVHWFramesConstraints) MinHeight() int {
	value := s.ptr.min_height
	return int(value)
}

func (s *AVHWFramesConstraints) SetMinHeight(value int) {
	s.ptr.min_height = (C.int)(value)
}

func (s *AVHWFramesConstraints) MaxWidth() int {
	value := s.ptr.max_width
	return int(value)
}

func (s *AVHWFramesConstraints) SetMaxWidth(value int) {
	s.ptr.max_width = (C.int)(value)
}

func (s *AVHWFramesConstraints) MaxHeight() int {
	value := s.ptr.max_height
	return int(value)
}

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

func (s *AVClass) ClassName() *CStr {
	value := s.ptr.class_name
	return wrapCStr(value)
}

func (s *AVClass) SetClassName(value *CStr) {
	s.ptr.class_name = value.ptr
}

// ItemName skipped due to func ptr

func (s *AVClass) Option() *AVOption {
	value := s.ptr.option
	var valueMapped *AVOption
	if value != nil {
		valueMapped = &AVOption{ptr: value}
	}
	return valueMapped
}

func (s *AVClass) SetOption(value *AVOption) {
	if value != nil {
		s.ptr.option = value.ptr
	} else {
		s.ptr.option = nil
	}
}

func (s *AVClass) Version() int {
	value := s.ptr.version
	return int(value)
}

func (s *AVClass) SetVersion(value int) {
	s.ptr.version = (C.int)(value)
}

func (s *AVClass) LogLevelOffsetOffset() int {
	value := s.ptr.log_level_offset_offset
	return int(value)
}

func (s *AVClass) SetLogLevelOffsetOffset(value int) {
	s.ptr.log_level_offset_offset = (C.int)(value)
}

func (s *AVClass) ParentLogContextOffset() int {
	value := s.ptr.parent_log_context_offset
	return int(value)
}

func (s *AVClass) SetParentLogContextOffset(value int) {
	s.ptr.parent_log_context_offset = (C.int)(value)
}

func (s *AVClass) Category() AVClassCategory {
	value := s.ptr.category
	return AVClassCategory(value)
}

func (s *AVClass) SetCategory(value AVClassCategory) {
	s.ptr.category = (C.AVClassCategory)(value)
}

// GetCategory skipped due to func ptr

// QueryRanges skipped due to func ptr

// ChildNext skipped due to func ptr

// ChildClassIterate skipped due to func ptr

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

func (s *AVOption) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

func (s *AVOption) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

func (s *AVOption) Help() *CStr {
	value := s.ptr.help
	return wrapCStr(value)
}

func (s *AVOption) SetHelp(value *CStr) {
	s.ptr.help = value.ptr
}

func (s *AVOption) Offset() int {
	value := s.ptr.offset
	return int(value)
}

func (s *AVOption) SetOffset(value int) {
	s.ptr.offset = (C.int)(value)
}

func (s *AVOption) Type() AVOptionType {
	value := s.ptr._type
	return AVOptionType(value)
}

func (s *AVOption) SetType(value AVOptionType) {
	s.ptr._type = (C.enum_AVOptionType)(value)
}

// DefaultVal skipped due to union type

func (s *AVOption) Min() float64 {
	value := s.ptr.min
	return float64(value)
}

func (s *AVOption) SetMin(value float64) {
	s.ptr.min = (C.double)(value)
}

func (s *AVOption) Max() float64 {
	value := s.ptr.max
	return float64(value)
}

func (s *AVOption) SetMax(value float64) {
	s.ptr.max = (C.double)(value)
}

func (s *AVOption) Flags() int {
	value := s.ptr.flags
	return int(value)
}

func (s *AVOption) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

func (s *AVOption) Unit() *CStr {
	value := s.ptr.unit
	return wrapCStr(value)
}

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

func (s *AVOptionRange) Str() *CStr {
	value := s.ptr.str
	return wrapCStr(value)
}

func (s *AVOptionRange) SetStr(value *CStr) {
	s.ptr.str = value.ptr
}

func (s *AVOptionRange) ValueMin() float64 {
	value := s.ptr.value_min
	return float64(value)
}

func (s *AVOptionRange) SetValueMin(value float64) {
	s.ptr.value_min = (C.double)(value)
}

func (s *AVOptionRange) ValueMax() float64 {
	value := s.ptr.value_max
	return float64(value)
}

func (s *AVOptionRange) SetValueMax(value float64) {
	s.ptr.value_max = (C.double)(value)
}

func (s *AVOptionRange) ComponentMin() float64 {
	value := s.ptr.component_min
	return float64(value)
}

func (s *AVOptionRange) SetComponentMin(value float64) {
	s.ptr.component_min = (C.double)(value)
}

func (s *AVOptionRange) ComponentMax() float64 {
	value := s.ptr.component_max
	return float64(value)
}

func (s *AVOptionRange) SetComponentMax(value float64) {
	s.ptr.component_max = (C.double)(value)
}

func (s *AVOptionRange) IsRange() int {
	value := s.ptr.is_range
	return int(value)
}

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

func (s *AVOptionRanges) RangeEntry(i uint) *AVOptionRange {
	value := s.ptr._range
	ptr := arrayGet[*C.AVOptionRange](value, uintptr(i))
	var valueMapped *AVOptionRange
	if ptr != nil {
		valueMapped = &AVOptionRange{ptr: ptr}
	}
	return valueMapped
}

func (s *AVOptionRanges) NbRanges() int {
	value := s.ptr.nb_ranges
	return int(value)
}

func (s *AVOptionRanges) SetNbRanges(value int) {
	s.ptr.nb_ranges = (C.int)(value)
}

func (s *AVOptionRanges) NbComponents() int {
	value := s.ptr.nb_components
	return int(value)
}

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

func (s *AVRational) Num() int {
	value := s.value.num
	return int(value)
}

func (s *AVRational) SetNum(value int) {
	s.value.num = (C.int)(value)
}

func (s *AVRational) Den() int {
	value := s.value.den
	return int(value)
}

func (s *AVRational) SetDen(value int) {
	s.value.den = (C.int)(value)
}
