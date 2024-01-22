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
func (s *RcOverride) Qscale() int {
	value := s.ptr.qscale
	return int(value)
}

// SetQscale sets the qscale field.
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
func (s *AVCodecContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
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
func (s *AVCodecContext) CodecType() AVMediaType {
	value := s.ptr.codec_type
	return AVMediaType(value)
}

// SetCodecType sets the codec_type field.
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
func (s *AVCodecContext) CodecId() AVCodecID {
	value := s.ptr.codec_id
	return AVCodecID(value)
}

// SetCodecId sets the codec_id field.
func (s *AVCodecContext) SetCodecId(value AVCodecID) {
	s.ptr.codec_id = (C.enum_AVCodecID)(value)
}

// CodecTag gets the codec_tag field.
func (s *AVCodecContext) CodecTag() uint {
	value := s.ptr.codec_tag
	return uint(value)
}

// SetCodecTag sets the codec_tag field.
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
func (s *AVCodecContext) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
func (s *AVCodecContext) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// BitRate gets the bit_rate field.
func (s *AVCodecContext) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}

// SetBitRate sets the bit_rate field.
func (s *AVCodecContext) SetBitRate(value int64) {
	s.ptr.bit_rate = (C.int64_t)(value)
}

// BitRateTolerance gets the bit_rate_tolerance field.
func (s *AVCodecContext) BitRateTolerance() int {
	value := s.ptr.bit_rate_tolerance
	return int(value)
}

// SetBitRateTolerance sets the bit_rate_tolerance field.
func (s *AVCodecContext) SetBitRateTolerance(value int) {
	s.ptr.bit_rate_tolerance = (C.int)(value)
}

// GlobalQuality gets the global_quality field.
func (s *AVCodecContext) GlobalQuality() int {
	value := s.ptr.global_quality
	return int(value)
}

// SetGlobalQuality sets the global_quality field.
func (s *AVCodecContext) SetGlobalQuality(value int) {
	s.ptr.global_quality = (C.int)(value)
}

// CompressionLevel gets the compression_level field.
func (s *AVCodecContext) CompressionLevel() int {
	value := s.ptr.compression_level
	return int(value)
}

// SetCompressionLevel sets the compression_level field.
func (s *AVCodecContext) SetCompressionLevel(value int) {
	s.ptr.compression_level = (C.int)(value)
}

// Flags gets the flags field.
func (s *AVCodecContext) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVCodecContext) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// Flags2 gets the flags2 field.
func (s *AVCodecContext) Flags2() int {
	value := s.ptr.flags2
	return int(value)
}

// SetFlags2 sets the flags2 field.
func (s *AVCodecContext) SetFlags2(value int) {
	s.ptr.flags2 = (C.int)(value)
}

// Extradata gets the extradata field.
func (s *AVCodecContext) Extradata() unsafe.Pointer {
	value := s.ptr.extradata
	return unsafe.Pointer(value)
}

// SetExtradata sets the extradata field.
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
func (s *AVCodecContext) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
func (s *AVCodecContext) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// TicksPerFrame gets the ticks_per_frame field.
func (s *AVCodecContext) TicksPerFrame() int {
	value := s.ptr.ticks_per_frame
	return int(value)
}

// SetTicksPerFrame sets the ticks_per_frame field.
func (s *AVCodecContext) SetTicksPerFrame(value int) {
	s.ptr.ticks_per_frame = (C.int)(value)
}

// Delay gets the delay field.
func (s *AVCodecContext) Delay() int {
	value := s.ptr.delay
	return int(value)
}

// SetDelay sets the delay field.
func (s *AVCodecContext) SetDelay(value int) {
	s.ptr.delay = (C.int)(value)
}

// Width gets the width field.
func (s *AVCodecContext) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
func (s *AVCodecContext) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

// Height gets the height field.
func (s *AVCodecContext) Height() int {
	value := s.ptr.height
	return int(value)
}

// SetHeight sets the height field.
func (s *AVCodecContext) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

// CodedWidth gets the coded_width field.
func (s *AVCodecContext) CodedWidth() int {
	value := s.ptr.coded_width
	return int(value)
}

// SetCodedWidth sets the coded_width field.
func (s *AVCodecContext) SetCodedWidth(value int) {
	s.ptr.coded_width = (C.int)(value)
}

// CodedHeight gets the coded_height field.
func (s *AVCodecContext) CodedHeight() int {
	value := s.ptr.coded_height
	return int(value)
}

// SetCodedHeight sets the coded_height field.
func (s *AVCodecContext) SetCodedHeight(value int) {
	s.ptr.coded_height = (C.int)(value)
}

// GopSize gets the gop_size field.
func (s *AVCodecContext) GopSize() int {
	value := s.ptr.gop_size
	return int(value)
}

// SetGopSize sets the gop_size field.
func (s *AVCodecContext) SetGopSize(value int) {
	s.ptr.gop_size = (C.int)(value)
}

// PixFmt gets the pix_fmt field.
func (s *AVCodecContext) PixFmt() AVPixelFormat {
	value := s.ptr.pix_fmt
	return AVPixelFormat(value)
}

// SetPixFmt sets the pix_fmt field.
func (s *AVCodecContext) SetPixFmt(value AVPixelFormat) {
	s.ptr.pix_fmt = (C.enum_AVPixelFormat)(value)
}

// draw_horiz_band skipped due to func ptr

// get_format skipped due to func ptr

// MaxBFrames gets the max_b_frames field.
func (s *AVCodecContext) MaxBFrames() int {
	value := s.ptr.max_b_frames
	return int(value)
}

// SetMaxBFrames sets the max_b_frames field.
func (s *AVCodecContext) SetMaxBFrames(value int) {
	s.ptr.max_b_frames = (C.int)(value)
}

// BQuantFactor gets the b_quant_factor field.
func (s *AVCodecContext) BQuantFactor() float32 {
	value := s.ptr.b_quant_factor
	return float32(value)
}

// SetBQuantFactor sets the b_quant_factor field.
func (s *AVCodecContext) SetBQuantFactor(value float32) {
	s.ptr.b_quant_factor = (C.float)(value)
}

// BQuantOffset gets the b_quant_offset field.
func (s *AVCodecContext) BQuantOffset() float32 {
	value := s.ptr.b_quant_offset
	return float32(value)
}

// SetBQuantOffset sets the b_quant_offset field.
func (s *AVCodecContext) SetBQuantOffset(value float32) {
	s.ptr.b_quant_offset = (C.float)(value)
}

// HasBFrames gets the has_b_frames field.
func (s *AVCodecContext) HasBFrames() int {
	value := s.ptr.has_b_frames
	return int(value)
}

// SetHasBFrames sets the has_b_frames field.
func (s *AVCodecContext) SetHasBFrames(value int) {
	s.ptr.has_b_frames = (C.int)(value)
}

// IQuantFactor gets the i_quant_factor field.
func (s *AVCodecContext) IQuantFactor() float32 {
	value := s.ptr.i_quant_factor
	return float32(value)
}

// SetIQuantFactor sets the i_quant_factor field.
func (s *AVCodecContext) SetIQuantFactor(value float32) {
	s.ptr.i_quant_factor = (C.float)(value)
}

// IQuantOffset gets the i_quant_offset field.
func (s *AVCodecContext) IQuantOffset() float32 {
	value := s.ptr.i_quant_offset
	return float32(value)
}

// SetIQuantOffset sets the i_quant_offset field.
func (s *AVCodecContext) SetIQuantOffset(value float32) {
	s.ptr.i_quant_offset = (C.float)(value)
}

// LumiMasking gets the lumi_masking field.
func (s *AVCodecContext) LumiMasking() float32 {
	value := s.ptr.lumi_masking
	return float32(value)
}

// SetLumiMasking sets the lumi_masking field.
func (s *AVCodecContext) SetLumiMasking(value float32) {
	s.ptr.lumi_masking = (C.float)(value)
}

// TemporalCplxMasking gets the temporal_cplx_masking field.
func (s *AVCodecContext) TemporalCplxMasking() float32 {
	value := s.ptr.temporal_cplx_masking
	return float32(value)
}

// SetTemporalCplxMasking sets the temporal_cplx_masking field.
func (s *AVCodecContext) SetTemporalCplxMasking(value float32) {
	s.ptr.temporal_cplx_masking = (C.float)(value)
}

// SpatialCplxMasking gets the spatial_cplx_masking field.
func (s *AVCodecContext) SpatialCplxMasking() float32 {
	value := s.ptr.spatial_cplx_masking
	return float32(value)
}

// SetSpatialCplxMasking sets the spatial_cplx_masking field.
func (s *AVCodecContext) SetSpatialCplxMasking(value float32) {
	s.ptr.spatial_cplx_masking = (C.float)(value)
}

// PMasking gets the p_masking field.
func (s *AVCodecContext) PMasking() float32 {
	value := s.ptr.p_masking
	return float32(value)
}

// SetPMasking sets the p_masking field.
func (s *AVCodecContext) SetPMasking(value float32) {
	s.ptr.p_masking = (C.float)(value)
}

// DarkMasking gets the dark_masking field.
func (s *AVCodecContext) DarkMasking() float32 {
	value := s.ptr.dark_masking
	return float32(value)
}

// SetDarkMasking sets the dark_masking field.
func (s *AVCodecContext) SetDarkMasking(value float32) {
	s.ptr.dark_masking = (C.float)(value)
}

// SliceCount gets the slice_count field.
func (s *AVCodecContext) SliceCount() int {
	value := s.ptr.slice_count
	return int(value)
}

// SetSliceCount sets the slice_count field.
func (s *AVCodecContext) SetSliceCount(value int) {
	s.ptr.slice_count = (C.int)(value)
}

// slice_offset skipped due to prim ptr

// SampleAspectRatio gets the sample_aspect_ratio field.
func (s *AVCodecContext) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
func (s *AVCodecContext) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

// MeCmp gets the me_cmp field.
func (s *AVCodecContext) MeCmp() int {
	value := s.ptr.me_cmp
	return int(value)
}

// SetMeCmp sets the me_cmp field.
func (s *AVCodecContext) SetMeCmp(value int) {
	s.ptr.me_cmp = (C.int)(value)
}

// MeSubCmp gets the me_sub_cmp field.
func (s *AVCodecContext) MeSubCmp() int {
	value := s.ptr.me_sub_cmp
	return int(value)
}

// SetMeSubCmp sets the me_sub_cmp field.
func (s *AVCodecContext) SetMeSubCmp(value int) {
	s.ptr.me_sub_cmp = (C.int)(value)
}

// MbCmp gets the mb_cmp field.
func (s *AVCodecContext) MbCmp() int {
	value := s.ptr.mb_cmp
	return int(value)
}

// SetMbCmp sets the mb_cmp field.
func (s *AVCodecContext) SetMbCmp(value int) {
	s.ptr.mb_cmp = (C.int)(value)
}

// IldctCmp gets the ildct_cmp field.
func (s *AVCodecContext) IldctCmp() int {
	value := s.ptr.ildct_cmp
	return int(value)
}

// SetIldctCmp sets the ildct_cmp field.
func (s *AVCodecContext) SetIldctCmp(value int) {
	s.ptr.ildct_cmp = (C.int)(value)
}

// DiaSize gets the dia_size field.
func (s *AVCodecContext) DiaSize() int {
	value := s.ptr.dia_size
	return int(value)
}

// SetDiaSize sets the dia_size field.
func (s *AVCodecContext) SetDiaSize(value int) {
	s.ptr.dia_size = (C.int)(value)
}

// LastPredictorCount gets the last_predictor_count field.
func (s *AVCodecContext) LastPredictorCount() int {
	value := s.ptr.last_predictor_count
	return int(value)
}

// SetLastPredictorCount sets the last_predictor_count field.
func (s *AVCodecContext) SetLastPredictorCount(value int) {
	s.ptr.last_predictor_count = (C.int)(value)
}

// MePreCmp gets the me_pre_cmp field.
func (s *AVCodecContext) MePreCmp() int {
	value := s.ptr.me_pre_cmp
	return int(value)
}

// SetMePreCmp sets the me_pre_cmp field.
func (s *AVCodecContext) SetMePreCmp(value int) {
	s.ptr.me_pre_cmp = (C.int)(value)
}

// PreDiaSize gets the pre_dia_size field.
func (s *AVCodecContext) PreDiaSize() int {
	value := s.ptr.pre_dia_size
	return int(value)
}

// SetPreDiaSize sets the pre_dia_size field.
func (s *AVCodecContext) SetPreDiaSize(value int) {
	s.ptr.pre_dia_size = (C.int)(value)
}

// MeSubpelQuality gets the me_subpel_quality field.
func (s *AVCodecContext) MeSubpelQuality() int {
	value := s.ptr.me_subpel_quality
	return int(value)
}

// SetMeSubpelQuality sets the me_subpel_quality field.
func (s *AVCodecContext) SetMeSubpelQuality(value int) {
	s.ptr.me_subpel_quality = (C.int)(value)
}

// MeRange gets the me_range field.
func (s *AVCodecContext) MeRange() int {
	value := s.ptr.me_range
	return int(value)
}

// SetMeRange sets the me_range field.
func (s *AVCodecContext) SetMeRange(value int) {
	s.ptr.me_range = (C.int)(value)
}

// SliceFlags gets the slice_flags field.
func (s *AVCodecContext) SliceFlags() int {
	value := s.ptr.slice_flags
	return int(value)
}

// SetSliceFlags sets the slice_flags field.
func (s *AVCodecContext) SetSliceFlags(value int) {
	s.ptr.slice_flags = (C.int)(value)
}

// MbDecision gets the mb_decision field.
func (s *AVCodecContext) MbDecision() int {
	value := s.ptr.mb_decision
	return int(value)
}

// SetMbDecision sets the mb_decision field.
func (s *AVCodecContext) SetMbDecision(value int) {
	s.ptr.mb_decision = (C.int)(value)
}

// intra_matrix skipped due to prim ptr

// inter_matrix skipped due to prim ptr

// IntraDcPrecision gets the intra_dc_precision field.
func (s *AVCodecContext) IntraDcPrecision() int {
	value := s.ptr.intra_dc_precision
	return int(value)
}

// SetIntraDcPrecision sets the intra_dc_precision field.
func (s *AVCodecContext) SetIntraDcPrecision(value int) {
	s.ptr.intra_dc_precision = (C.int)(value)
}

// SkipTop gets the skip_top field.
func (s *AVCodecContext) SkipTop() int {
	value := s.ptr.skip_top
	return int(value)
}

// SetSkipTop sets the skip_top field.
func (s *AVCodecContext) SetSkipTop(value int) {
	s.ptr.skip_top = (C.int)(value)
}

// SkipBottom gets the skip_bottom field.
func (s *AVCodecContext) SkipBottom() int {
	value := s.ptr.skip_bottom
	return int(value)
}

// SetSkipBottom sets the skip_bottom field.
func (s *AVCodecContext) SetSkipBottom(value int) {
	s.ptr.skip_bottom = (C.int)(value)
}

// MbLmin gets the mb_lmin field.
func (s *AVCodecContext) MbLmin() int {
	value := s.ptr.mb_lmin
	return int(value)
}

// SetMbLmin sets the mb_lmin field.
func (s *AVCodecContext) SetMbLmin(value int) {
	s.ptr.mb_lmin = (C.int)(value)
}

// MbLmax gets the mb_lmax field.
func (s *AVCodecContext) MbLmax() int {
	value := s.ptr.mb_lmax
	return int(value)
}

// SetMbLmax sets the mb_lmax field.
func (s *AVCodecContext) SetMbLmax(value int) {
	s.ptr.mb_lmax = (C.int)(value)
}

// BidirRefine gets the bidir_refine field.
func (s *AVCodecContext) BidirRefine() int {
	value := s.ptr.bidir_refine
	return int(value)
}

// SetBidirRefine sets the bidir_refine field.
func (s *AVCodecContext) SetBidirRefine(value int) {
	s.ptr.bidir_refine = (C.int)(value)
}

// KeyintMin gets the keyint_min field.
func (s *AVCodecContext) KeyintMin() int {
	value := s.ptr.keyint_min
	return int(value)
}

// SetKeyintMin sets the keyint_min field.
func (s *AVCodecContext) SetKeyintMin(value int) {
	s.ptr.keyint_min = (C.int)(value)
}

// Refs gets the refs field.
func (s *AVCodecContext) Refs() int {
	value := s.ptr.refs
	return int(value)
}

// SetRefs sets the refs field.
func (s *AVCodecContext) SetRefs(value int) {
	s.ptr.refs = (C.int)(value)
}

// Mv0Threshold gets the mv0_threshold field.
func (s *AVCodecContext) Mv0Threshold() int {
	value := s.ptr.mv0_threshold
	return int(value)
}

// SetMv0Threshold sets the mv0_threshold field.
func (s *AVCodecContext) SetMv0Threshold(value int) {
	s.ptr.mv0_threshold = (C.int)(value)
}

// ColorPrimaries gets the color_primaries field.
func (s *AVCodecContext) ColorPrimaries() AVColorPrimaries {
	value := s.ptr.color_primaries
	return AVColorPrimaries(value)
}

// SetColorPrimaries sets the color_primaries field.
func (s *AVCodecContext) SetColorPrimaries(value AVColorPrimaries) {
	s.ptr.color_primaries = (C.enum_AVColorPrimaries)(value)
}

// ColorTrc gets the color_trc field.
func (s *AVCodecContext) ColorTrc() AVColorTransferCharacteristic {
	value := s.ptr.color_trc
	return AVColorTransferCharacteristic(value)
}

// SetColorTrc sets the color_trc field.
func (s *AVCodecContext) SetColorTrc(value AVColorTransferCharacteristic) {
	s.ptr.color_trc = (C.enum_AVColorTransferCharacteristic)(value)
}

// Colorspace gets the colorspace field.
func (s *AVCodecContext) Colorspace() AVColorSpace {
	value := s.ptr.colorspace
	return AVColorSpace(value)
}

// SetColorspace sets the colorspace field.
func (s *AVCodecContext) SetColorspace(value AVColorSpace) {
	s.ptr.colorspace = (C.enum_AVColorSpace)(value)
}

// ColorRange gets the color_range field.
func (s *AVCodecContext) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}

// SetColorRange sets the color_range field.
func (s *AVCodecContext) SetColorRange(value AVColorRange) {
	s.ptr.color_range = (C.enum_AVColorRange)(value)
}

// ChromaSampleLocation gets the chroma_sample_location field.
func (s *AVCodecContext) ChromaSampleLocation() AVChromaLocation {
	value := s.ptr.chroma_sample_location
	return AVChromaLocation(value)
}

// SetChromaSampleLocation sets the chroma_sample_location field.
func (s *AVCodecContext) SetChromaSampleLocation(value AVChromaLocation) {
	s.ptr.chroma_sample_location = (C.enum_AVChromaLocation)(value)
}

// Slices gets the slices field.
func (s *AVCodecContext) Slices() int {
	value := s.ptr.slices
	return int(value)
}

// SetSlices sets the slices field.
func (s *AVCodecContext) SetSlices(value int) {
	s.ptr.slices = (C.int)(value)
}

// FieldOrder gets the field_order field.
func (s *AVCodecContext) FieldOrder() AVFieldOrder {
	value := s.ptr.field_order
	return AVFieldOrder(value)
}

// SetFieldOrder sets the field_order field.
func (s *AVCodecContext) SetFieldOrder(value AVFieldOrder) {
	s.ptr.field_order = (C.enum_AVFieldOrder)(value)
}

// SampleRate gets the sample_rate field.
func (s *AVCodecContext) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

// SetSampleRate sets the sample_rate field.
func (s *AVCodecContext) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

// Channels gets the channels field.
func (s *AVCodecContext) Channels() int {
	value := s.ptr.channels
	return int(value)
}

// SetChannels sets the channels field.
func (s *AVCodecContext) SetChannels(value int) {
	s.ptr.channels = (C.int)(value)
}

// SampleFmt gets the sample_fmt field.
func (s *AVCodecContext) SampleFmt() AVSampleFormat {
	value := s.ptr.sample_fmt
	return AVSampleFormat(value)
}

// SetSampleFmt sets the sample_fmt field.
func (s *AVCodecContext) SetSampleFmt(value AVSampleFormat) {
	s.ptr.sample_fmt = (C.enum_AVSampleFormat)(value)
}

// FrameSize gets the frame_size field.
func (s *AVCodecContext) FrameSize() int {
	value := s.ptr.frame_size
	return int(value)
}

// SetFrameSize sets the frame_size field.
func (s *AVCodecContext) SetFrameSize(value int) {
	s.ptr.frame_size = (C.int)(value)
}

// FrameNumber gets the frame_number field.
func (s *AVCodecContext) FrameNumber() int {
	value := s.ptr.frame_number
	return int(value)
}

// SetFrameNumber sets the frame_number field.
func (s *AVCodecContext) SetFrameNumber(value int) {
	s.ptr.frame_number = (C.int)(value)
}

// BlockAlign gets the block_align field.
func (s *AVCodecContext) BlockAlign() int {
	value := s.ptr.block_align
	return int(value)
}

// SetBlockAlign sets the block_align field.
func (s *AVCodecContext) SetBlockAlign(value int) {
	s.ptr.block_align = (C.int)(value)
}

// Cutoff gets the cutoff field.
func (s *AVCodecContext) Cutoff() int {
	value := s.ptr.cutoff
	return int(value)
}

// SetCutoff sets the cutoff field.
func (s *AVCodecContext) SetCutoff(value int) {
	s.ptr.cutoff = (C.int)(value)
}

// ChannelLayout gets the channel_layout field.
func (s *AVCodecContext) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

// SetChannelLayout sets the channel_layout field.
func (s *AVCodecContext) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

// RequestChannelLayout gets the request_channel_layout field.
func (s *AVCodecContext) RequestChannelLayout() uint64 {
	value := s.ptr.request_channel_layout
	return uint64(value)
}

// SetRequestChannelLayout sets the request_channel_layout field.
func (s *AVCodecContext) SetRequestChannelLayout(value uint64) {
	s.ptr.request_channel_layout = (C.uint64_t)(value)
}

// AudioServiceType gets the audio_service_type field.
func (s *AVCodecContext) AudioServiceType() AVAudioServiceType {
	value := s.ptr.audio_service_type
	return AVAudioServiceType(value)
}

// SetAudioServiceType sets the audio_service_type field.
func (s *AVCodecContext) SetAudioServiceType(value AVAudioServiceType) {
	s.ptr.audio_service_type = (C.enum_AVAudioServiceType)(value)
}

// RequestSampleFmt gets the request_sample_fmt field.
func (s *AVCodecContext) RequestSampleFmt() AVSampleFormat {
	value := s.ptr.request_sample_fmt
	return AVSampleFormat(value)
}

// SetRequestSampleFmt sets the request_sample_fmt field.
func (s *AVCodecContext) SetRequestSampleFmt(value AVSampleFormat) {
	s.ptr.request_sample_fmt = (C.enum_AVSampleFormat)(value)
}

// get_buffer2 skipped due to func ptr

// Qcompress gets the qcompress field.
func (s *AVCodecContext) Qcompress() float32 {
	value := s.ptr.qcompress
	return float32(value)
}

// SetQcompress sets the qcompress field.
func (s *AVCodecContext) SetQcompress(value float32) {
	s.ptr.qcompress = (C.float)(value)
}

// Qblur gets the qblur field.
func (s *AVCodecContext) Qblur() float32 {
	value := s.ptr.qblur
	return float32(value)
}

// SetQblur sets the qblur field.
func (s *AVCodecContext) SetQblur(value float32) {
	s.ptr.qblur = (C.float)(value)
}

// Qmin gets the qmin field.
func (s *AVCodecContext) Qmin() int {
	value := s.ptr.qmin
	return int(value)
}

// SetQmin sets the qmin field.
func (s *AVCodecContext) SetQmin(value int) {
	s.ptr.qmin = (C.int)(value)
}

// Qmax gets the qmax field.
func (s *AVCodecContext) Qmax() int {
	value := s.ptr.qmax
	return int(value)
}

// SetQmax sets the qmax field.
func (s *AVCodecContext) SetQmax(value int) {
	s.ptr.qmax = (C.int)(value)
}

// MaxQdiff gets the max_qdiff field.
func (s *AVCodecContext) MaxQdiff() int {
	value := s.ptr.max_qdiff
	return int(value)
}

// SetMaxQdiff sets the max_qdiff field.
func (s *AVCodecContext) SetMaxQdiff(value int) {
	s.ptr.max_qdiff = (C.int)(value)
}

// RcBufferSize gets the rc_buffer_size field.
func (s *AVCodecContext) RcBufferSize() int {
	value := s.ptr.rc_buffer_size
	return int(value)
}

// SetRcBufferSize sets the rc_buffer_size field.
func (s *AVCodecContext) SetRcBufferSize(value int) {
	s.ptr.rc_buffer_size = (C.int)(value)
}

// RcOverrideCount gets the rc_override_count field.
func (s *AVCodecContext) RcOverrideCount() int {
	value := s.ptr.rc_override_count
	return int(value)
}

// SetRcOverrideCount sets the rc_override_count field.
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
func (s *AVCodecContext) RcMaxRate() int64 {
	value := s.ptr.rc_max_rate
	return int64(value)
}

// SetRcMaxRate sets the rc_max_rate field.
func (s *AVCodecContext) SetRcMaxRate(value int64) {
	s.ptr.rc_max_rate = (C.int64_t)(value)
}

// RcMinRate gets the rc_min_rate field.
func (s *AVCodecContext) RcMinRate() int64 {
	value := s.ptr.rc_min_rate
	return int64(value)
}

// SetRcMinRate sets the rc_min_rate field.
func (s *AVCodecContext) SetRcMinRate(value int64) {
	s.ptr.rc_min_rate = (C.int64_t)(value)
}

// RcMaxAvailableVbvUse gets the rc_max_available_vbv_use field.
func (s *AVCodecContext) RcMaxAvailableVbvUse() float32 {
	value := s.ptr.rc_max_available_vbv_use
	return float32(value)
}

// SetRcMaxAvailableVbvUse sets the rc_max_available_vbv_use field.
func (s *AVCodecContext) SetRcMaxAvailableVbvUse(value float32) {
	s.ptr.rc_max_available_vbv_use = (C.float)(value)
}

// RcMinVbvOverflowUse gets the rc_min_vbv_overflow_use field.
func (s *AVCodecContext) RcMinVbvOverflowUse() float32 {
	value := s.ptr.rc_min_vbv_overflow_use
	return float32(value)
}

// SetRcMinVbvOverflowUse sets the rc_min_vbv_overflow_use field.
func (s *AVCodecContext) SetRcMinVbvOverflowUse(value float32) {
	s.ptr.rc_min_vbv_overflow_use = (C.float)(value)
}

// RcInitialBufferOccupancy gets the rc_initial_buffer_occupancy field.
func (s *AVCodecContext) RcInitialBufferOccupancy() int {
	value := s.ptr.rc_initial_buffer_occupancy
	return int(value)
}

// SetRcInitialBufferOccupancy sets the rc_initial_buffer_occupancy field.
func (s *AVCodecContext) SetRcInitialBufferOccupancy(value int) {
	s.ptr.rc_initial_buffer_occupancy = (C.int)(value)
}

// Trellis gets the trellis field.
func (s *AVCodecContext) Trellis() int {
	value := s.ptr.trellis
	return int(value)
}

// SetTrellis sets the trellis field.
func (s *AVCodecContext) SetTrellis(value int) {
	s.ptr.trellis = (C.int)(value)
}

// StatsOut gets the stats_out field.
func (s *AVCodecContext) StatsOut() *CStr {
	value := s.ptr.stats_out
	return wrapCStr(value)
}

// SetStatsOut sets the stats_out field.
func (s *AVCodecContext) SetStatsOut(value *CStr) {
	s.ptr.stats_out = value.ptr
}

// StatsIn gets the stats_in field.
func (s *AVCodecContext) StatsIn() *CStr {
	value := s.ptr.stats_in
	return wrapCStr(value)
}

// SetStatsIn sets the stats_in field.
func (s *AVCodecContext) SetStatsIn(value *CStr) {
	s.ptr.stats_in = value.ptr
}

// WorkaroundBugs gets the workaround_bugs field.
func (s *AVCodecContext) WorkaroundBugs() int {
	value := s.ptr.workaround_bugs
	return int(value)
}

// SetWorkaroundBugs sets the workaround_bugs field.
func (s *AVCodecContext) SetWorkaroundBugs(value int) {
	s.ptr.workaround_bugs = (C.int)(value)
}

// StrictStdCompliance gets the strict_std_compliance field.
func (s *AVCodecContext) StrictStdCompliance() int {
	value := s.ptr.strict_std_compliance
	return int(value)
}

// SetStrictStdCompliance sets the strict_std_compliance field.
func (s *AVCodecContext) SetStrictStdCompliance(value int) {
	s.ptr.strict_std_compliance = (C.int)(value)
}

// ErrorConcealment gets the error_concealment field.
func (s *AVCodecContext) ErrorConcealment() int {
	value := s.ptr.error_concealment
	return int(value)
}

// SetErrorConcealment sets the error_concealment field.
func (s *AVCodecContext) SetErrorConcealment(value int) {
	s.ptr.error_concealment = (C.int)(value)
}

// Debug gets the debug field.
func (s *AVCodecContext) Debug() int {
	value := s.ptr.debug
	return int(value)
}

// SetDebug sets the debug field.
func (s *AVCodecContext) SetDebug(value int) {
	s.ptr.debug = (C.int)(value)
}

// ErrRecognition gets the err_recognition field.
func (s *AVCodecContext) ErrRecognition() int {
	value := s.ptr.err_recognition
	return int(value)
}

// SetErrRecognition sets the err_recognition field.
func (s *AVCodecContext) SetErrRecognition(value int) {
	s.ptr.err_recognition = (C.int)(value)
}

// ReorderedOpaque gets the reordered_opaque field.
func (s *AVCodecContext) ReorderedOpaque() int64 {
	value := s.ptr.reordered_opaque
	return int64(value)
}

// SetReorderedOpaque sets the reordered_opaque field.
func (s *AVCodecContext) SetReorderedOpaque(value int64) {
	s.ptr.reordered_opaque = (C.int64_t)(value)
}

// Hwaccel gets the hwaccel field.
func (s *AVCodecContext) Hwaccel() *AVHWAccel {
	value := s.ptr.hwaccel
	var valueMapped *AVHWAccel
	if value != nil {
		valueMapped = &AVHWAccel{ptr: value}
	}
	return valueMapped
}

// SetHwaccel sets the hwaccel field.
func (s *AVCodecContext) SetHwaccel(value *AVHWAccel) {
	if value != nil {
		s.ptr.hwaccel = value.ptr
	} else {
		s.ptr.hwaccel = nil
	}
}

// HwaccelContext gets the hwaccel_context field.
func (s *AVCodecContext) HwaccelContext() unsafe.Pointer {
	value := s.ptr.hwaccel_context
	return value
}

// SetHwaccelContext sets the hwaccel_context field.
func (s *AVCodecContext) SetHwaccelContext(value unsafe.Pointer) {
	s.ptr.hwaccel_context = value
}

// Error gets the error field.
func (s *AVCodecContext) Error() *Array[uint64] {
	value := &s.ptr.error
	return ToUint64Array(unsafe.Pointer(value))
}

// DctAlgo gets the dct_algo field.
func (s *AVCodecContext) DctAlgo() int {
	value := s.ptr.dct_algo
	return int(value)
}

// SetDctAlgo sets the dct_algo field.
func (s *AVCodecContext) SetDctAlgo(value int) {
	s.ptr.dct_algo = (C.int)(value)
}

// IdctAlgo gets the idct_algo field.
func (s *AVCodecContext) IdctAlgo() int {
	value := s.ptr.idct_algo
	return int(value)
}

// SetIdctAlgo sets the idct_algo field.
func (s *AVCodecContext) SetIdctAlgo(value int) {
	s.ptr.idct_algo = (C.int)(value)
}

// BitsPerCodedSample gets the bits_per_coded_sample field.
func (s *AVCodecContext) BitsPerCodedSample() int {
	value := s.ptr.bits_per_coded_sample
	return int(value)
}

// SetBitsPerCodedSample sets the bits_per_coded_sample field.
func (s *AVCodecContext) SetBitsPerCodedSample(value int) {
	s.ptr.bits_per_coded_sample = (C.int)(value)
}

// BitsPerRawSample gets the bits_per_raw_sample field.
func (s *AVCodecContext) BitsPerRawSample() int {
	value := s.ptr.bits_per_raw_sample
	return int(value)
}

// SetBitsPerRawSample sets the bits_per_raw_sample field.
func (s *AVCodecContext) SetBitsPerRawSample(value int) {
	s.ptr.bits_per_raw_sample = (C.int)(value)
}

// Lowres gets the lowres field.
func (s *AVCodecContext) Lowres() int {
	value := s.ptr.lowres
	return int(value)
}

// SetLowres sets the lowres field.
func (s *AVCodecContext) SetLowres(value int) {
	s.ptr.lowres = (C.int)(value)
}

// ThreadCount gets the thread_count field.
func (s *AVCodecContext) ThreadCount() int {
	value := s.ptr.thread_count
	return int(value)
}

// SetThreadCount sets the thread_count field.
func (s *AVCodecContext) SetThreadCount(value int) {
	s.ptr.thread_count = (C.int)(value)
}

// ThreadType gets the thread_type field.
func (s *AVCodecContext) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}

// SetThreadType sets the thread_type field.
func (s *AVCodecContext) SetThreadType(value int) {
	s.ptr.thread_type = (C.int)(value)
}

// ActiveThreadType gets the active_thread_type field.
func (s *AVCodecContext) ActiveThreadType() int {
	value := s.ptr.active_thread_type
	return int(value)
}

// SetActiveThreadType sets the active_thread_type field.
func (s *AVCodecContext) SetActiveThreadType(value int) {
	s.ptr.active_thread_type = (C.int)(value)
}

// execute skipped due to func ptr

// execute2 skipped due to func ptr

// NsseWeight gets the nsse_weight field.
func (s *AVCodecContext) NsseWeight() int {
	value := s.ptr.nsse_weight
	return int(value)
}

// SetNsseWeight sets the nsse_weight field.
func (s *AVCodecContext) SetNsseWeight(value int) {
	s.ptr.nsse_weight = (C.int)(value)
}

// Profile gets the profile field.
func (s *AVCodecContext) Profile() int {
	value := s.ptr.profile
	return int(value)
}

// SetProfile sets the profile field.
func (s *AVCodecContext) SetProfile(value int) {
	s.ptr.profile = (C.int)(value)
}

// Level gets the level field.
func (s *AVCodecContext) Level() int {
	value := s.ptr.level
	return int(value)
}

// SetLevel sets the level field.
func (s *AVCodecContext) SetLevel(value int) {
	s.ptr.level = (C.int)(value)
}

// SkipLoopFilter gets the skip_loop_filter field.
func (s *AVCodecContext) SkipLoopFilter() AVDiscard {
	value := s.ptr.skip_loop_filter
	return AVDiscard(value)
}

// SetSkipLoopFilter sets the skip_loop_filter field.
func (s *AVCodecContext) SetSkipLoopFilter(value AVDiscard) {
	s.ptr.skip_loop_filter = (C.enum_AVDiscard)(value)
}

// SkipIdct gets the skip_idct field.
func (s *AVCodecContext) SkipIdct() AVDiscard {
	value := s.ptr.skip_idct
	return AVDiscard(value)
}

// SetSkipIdct sets the skip_idct field.
func (s *AVCodecContext) SetSkipIdct(value AVDiscard) {
	s.ptr.skip_idct = (C.enum_AVDiscard)(value)
}

// SkipFrame gets the skip_frame field.
func (s *AVCodecContext) SkipFrame() AVDiscard {
	value := s.ptr.skip_frame
	return AVDiscard(value)
}

// SetSkipFrame sets the skip_frame field.
func (s *AVCodecContext) SetSkipFrame(value AVDiscard) {
	s.ptr.skip_frame = (C.enum_AVDiscard)(value)
}

// SubtitleHeader gets the subtitle_header field.
func (s *AVCodecContext) SubtitleHeader() unsafe.Pointer {
	value := s.ptr.subtitle_header
	return unsafe.Pointer(value)
}

// SetSubtitleHeader sets the subtitle_header field.
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
func (s *AVCodecContext) InitialPadding() int {
	value := s.ptr.initial_padding
	return int(value)
}

// SetInitialPadding sets the initial_padding field.
func (s *AVCodecContext) SetInitialPadding(value int) {
	s.ptr.initial_padding = (C.int)(value)
}

// Framerate gets the framerate field.
func (s *AVCodecContext) Framerate() *AVRational {
	value := s.ptr.framerate
	return &AVRational{value: value}
}

// SetFramerate sets the framerate field.
func (s *AVCodecContext) SetFramerate(value *AVRational) {
	s.ptr.framerate = value.value
}

// SwPixFmt gets the sw_pix_fmt field.
func (s *AVCodecContext) SwPixFmt() AVPixelFormat {
	value := s.ptr.sw_pix_fmt
	return AVPixelFormat(value)
}

// SetSwPixFmt sets the sw_pix_fmt field.
func (s *AVCodecContext) SetSwPixFmt(value AVPixelFormat) {
	s.ptr.sw_pix_fmt = (C.enum_AVPixelFormat)(value)
}

// PktTimebase gets the pkt_timebase field.
func (s *AVCodecContext) PktTimebase() *AVRational {
	value := s.ptr.pkt_timebase
	return &AVRational{value: value}
}

// SetPktTimebase sets the pkt_timebase field.
func (s *AVCodecContext) SetPktTimebase(value *AVRational) {
	s.ptr.pkt_timebase = value.value
}

// CodecDescriptor gets the codec_descriptor field.
func (s *AVCodecContext) CodecDescriptor() *AVCodecDescriptor {
	value := s.ptr.codec_descriptor
	var valueMapped *AVCodecDescriptor
	if value != nil {
		valueMapped = &AVCodecDescriptor{ptr: value}
	}
	return valueMapped
}

// SetCodecDescriptor sets the codec_descriptor field.
func (s *AVCodecContext) SetCodecDescriptor(value *AVCodecDescriptor) {
	if value != nil {
		s.ptr.codec_descriptor = value.ptr
	} else {
		s.ptr.codec_descriptor = nil
	}
}

// PtsCorrectionNumFaultyPts gets the pts_correction_num_faulty_pts field.
func (s *AVCodecContext) PtsCorrectionNumFaultyPts() int64 {
	value := s.ptr.pts_correction_num_faulty_pts
	return int64(value)
}

// SetPtsCorrectionNumFaultyPts sets the pts_correction_num_faulty_pts field.
func (s *AVCodecContext) SetPtsCorrectionNumFaultyPts(value int64) {
	s.ptr.pts_correction_num_faulty_pts = (C.int64_t)(value)
}

// PtsCorrectionNumFaultyDts gets the pts_correction_num_faulty_dts field.
func (s *AVCodecContext) PtsCorrectionNumFaultyDts() int64 {
	value := s.ptr.pts_correction_num_faulty_dts
	return int64(value)
}

// SetPtsCorrectionNumFaultyDts sets the pts_correction_num_faulty_dts field.
func (s *AVCodecContext) SetPtsCorrectionNumFaultyDts(value int64) {
	s.ptr.pts_correction_num_faulty_dts = (C.int64_t)(value)
}

// PtsCorrectionLastPts gets the pts_correction_last_pts field.
func (s *AVCodecContext) PtsCorrectionLastPts() int64 {
	value := s.ptr.pts_correction_last_pts
	return int64(value)
}

// SetPtsCorrectionLastPts sets the pts_correction_last_pts field.
func (s *AVCodecContext) SetPtsCorrectionLastPts(value int64) {
	s.ptr.pts_correction_last_pts = (C.int64_t)(value)
}

// PtsCorrectionLastDts gets the pts_correction_last_dts field.
func (s *AVCodecContext) PtsCorrectionLastDts() int64 {
	value := s.ptr.pts_correction_last_dts
	return int64(value)
}

// SetPtsCorrectionLastDts sets the pts_correction_last_dts field.
func (s *AVCodecContext) SetPtsCorrectionLastDts(value int64) {
	s.ptr.pts_correction_last_dts = (C.int64_t)(value)
}

// SubCharenc gets the sub_charenc field.
func (s *AVCodecContext) SubCharenc() *CStr {
	value := s.ptr.sub_charenc
	return wrapCStr(value)
}

// SetSubCharenc sets the sub_charenc field.
func (s *AVCodecContext) SetSubCharenc(value *CStr) {
	s.ptr.sub_charenc = value.ptr
}

// SubCharencMode gets the sub_charenc_mode field.
func (s *AVCodecContext) SubCharencMode() int {
	value := s.ptr.sub_charenc_mode
	return int(value)
}

// SetSubCharencMode sets the sub_charenc_mode field.
func (s *AVCodecContext) SetSubCharencMode(value int) {
	s.ptr.sub_charenc_mode = (C.int)(value)
}

// SkipAlpha gets the skip_alpha field.
func (s *AVCodecContext) SkipAlpha() int {
	value := s.ptr.skip_alpha
	return int(value)
}

// SetSkipAlpha sets the skip_alpha field.
func (s *AVCodecContext) SetSkipAlpha(value int) {
	s.ptr.skip_alpha = (C.int)(value)
}

// SeekPreroll gets the seek_preroll field.
func (s *AVCodecContext) SeekPreroll() int {
	value := s.ptr.seek_preroll
	return int(value)
}

// SetSeekPreroll sets the seek_preroll field.
func (s *AVCodecContext) SetSeekPreroll(value int) {
	s.ptr.seek_preroll = (C.int)(value)
}

// chroma_intra_matrix skipped due to prim ptr

// DumpSeparator gets the dump_separator field.
func (s *AVCodecContext) DumpSeparator() unsafe.Pointer {
	value := s.ptr.dump_separator
	return unsafe.Pointer(value)
}

// SetDumpSeparator sets the dump_separator field.
func (s *AVCodecContext) SetDumpSeparator(value unsafe.Pointer) {
	s.ptr.dump_separator = (*C.uint8_t)(value)
}

// CodecWhitelist gets the codec_whitelist field.
func (s *AVCodecContext) CodecWhitelist() *CStr {
	value := s.ptr.codec_whitelist
	return wrapCStr(value)
}

// SetCodecWhitelist sets the codec_whitelist field.
func (s *AVCodecContext) SetCodecWhitelist(value *CStr) {
	s.ptr.codec_whitelist = value.ptr
}

// Properties gets the properties field.
func (s *AVCodecContext) Properties() uint {
	value := s.ptr.properties
	return uint(value)
}

// SetProperties sets the properties field.
func (s *AVCodecContext) SetProperties(value uint) {
	s.ptr.properties = (C.uint)(value)
}

// CodedSideData gets the coded_side_data field.
func (s *AVCodecContext) CodedSideData() *AVPacketSideData {
	value := s.ptr.coded_side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}

// SetCodedSideData sets the coded_side_data field.
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
func (s *AVCodecContext) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwFramesCtx sets the hw_frames_ctx field.
func (s *AVCodecContext) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

// TrailingPadding gets the trailing_padding field.
func (s *AVCodecContext) TrailingPadding() int {
	value := s.ptr.trailing_padding
	return int(value)
}

// SetTrailingPadding sets the trailing_padding field.
func (s *AVCodecContext) SetTrailingPadding(value int) {
	s.ptr.trailing_padding = (C.int)(value)
}

// MaxPixels gets the max_pixels field.
func (s *AVCodecContext) MaxPixels() int64 {
	value := s.ptr.max_pixels
	return int64(value)
}

// SetMaxPixels sets the max_pixels field.
func (s *AVCodecContext) SetMaxPixels(value int64) {
	s.ptr.max_pixels = (C.int64_t)(value)
}

// HwDeviceCtx gets the hw_device_ctx field.
func (s *AVCodecContext) HwDeviceCtx() *AVBufferRef {
	value := s.ptr.hw_device_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwDeviceCtx sets the hw_device_ctx field.
func (s *AVCodecContext) SetHwDeviceCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_device_ctx = value.ptr
	} else {
		s.ptr.hw_device_ctx = nil
	}
}

// HwaccelFlags gets the hwaccel_flags field.
func (s *AVCodecContext) HwaccelFlags() int {
	value := s.ptr.hwaccel_flags
	return int(value)
}

// SetHwaccelFlags sets the hwaccel_flags field.
func (s *AVCodecContext) SetHwaccelFlags(value int) {
	s.ptr.hwaccel_flags = (C.int)(value)
}

// ApplyCropping gets the apply_cropping field.
func (s *AVCodecContext) ApplyCropping() int {
	value := s.ptr.apply_cropping
	return int(value)
}

// SetApplyCropping sets the apply_cropping field.
func (s *AVCodecContext) SetApplyCropping(value int) {
	s.ptr.apply_cropping = (C.int)(value)
}

// ExtraHwFrames gets the extra_hw_frames field.
func (s *AVCodecContext) ExtraHwFrames() int {
	value := s.ptr.extra_hw_frames
	return int(value)
}

// SetExtraHwFrames sets the extra_hw_frames field.
func (s *AVCodecContext) SetExtraHwFrames(value int) {
	s.ptr.extra_hw_frames = (C.int)(value)
}

// DiscardDamagedPercentage gets the discard_damaged_percentage field.
func (s *AVCodecContext) DiscardDamagedPercentage() int {
	value := s.ptr.discard_damaged_percentage
	return int(value)
}

// SetDiscardDamagedPercentage sets the discard_damaged_percentage field.
func (s *AVCodecContext) SetDiscardDamagedPercentage(value int) {
	s.ptr.discard_damaged_percentage = (C.int)(value)
}

// MaxSamples gets the max_samples field.
func (s *AVCodecContext) MaxSamples() int64 {
	value := s.ptr.max_samples
	return int64(value)
}

// SetMaxSamples sets the max_samples field.
func (s *AVCodecContext) SetMaxSamples(value int64) {
	s.ptr.max_samples = (C.int64_t)(value)
}

// ExportSideData gets the export_side_data field.
func (s *AVCodecContext) ExportSideData() int {
	value := s.ptr.export_side_data
	return int(value)
}

// SetExportSideData sets the export_side_data field.
func (s *AVCodecContext) SetExportSideData(value int) {
	s.ptr.export_side_data = (C.int)(value)
}

// get_encode_buffer skipped due to func ptr

// ChLayout gets the ch_layout field.
func (s *AVCodecContext) ChLayout() *AVChannelLayout {
	value := &s.ptr.ch_layout
	return &AVChannelLayout{ptr: value}
}

// FrameNum gets the frame_num field.
func (s *AVCodecContext) FrameNum() int64 {
	value := s.ptr.frame_num
	return int64(value)
}

// SetFrameNum sets the frame_num field.
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
func (s *AVHWAccel) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
func (s *AVHWAccel) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// Type gets the type field.
func (s *AVHWAccel) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}

// SetType sets the type field.
func (s *AVHWAccel) SetType(value AVMediaType) {
	s.ptr._type = (C.enum_AVMediaType)(value)
}

// Id gets the id field.
func (s *AVHWAccel) Id() AVCodecID {
	value := s.ptr.id
	return AVCodecID(value)
}

// SetId sets the id field.
func (s *AVHWAccel) SetId(value AVCodecID) {
	s.ptr.id = (C.enum_AVCodecID)(value)
}

// PixFmt gets the pix_fmt field.
func (s *AVHWAccel) PixFmt() AVPixelFormat {
	value := s.ptr.pix_fmt
	return AVPixelFormat(value)
}

// SetPixFmt sets the pix_fmt field.
func (s *AVHWAccel) SetPixFmt(value AVPixelFormat) {
	s.ptr.pix_fmt = (C.enum_AVPixelFormat)(value)
}

// Capabilities gets the capabilities field.
func (s *AVHWAccel) Capabilities() int {
	value := s.ptr.capabilities
	return int(value)
}

// SetCapabilities sets the capabilities field.
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
func (s *AVSubtitleRect) X() int {
	value := s.ptr.x
	return int(value)
}

// SetX sets the x field.
func (s *AVSubtitleRect) SetX(value int) {
	s.ptr.x = (C.int)(value)
}

// Y gets the y field.
func (s *AVSubtitleRect) Y() int {
	value := s.ptr.y
	return int(value)
}

// SetY sets the y field.
func (s *AVSubtitleRect) SetY(value int) {
	s.ptr.y = (C.int)(value)
}

// W gets the w field.
func (s *AVSubtitleRect) W() int {
	value := s.ptr.w
	return int(value)
}

// SetW sets the w field.
func (s *AVSubtitleRect) SetW(value int) {
	s.ptr.w = (C.int)(value)
}

// H gets the h field.
func (s *AVSubtitleRect) H() int {
	value := s.ptr.h
	return int(value)
}

// SetH sets the h field.
func (s *AVSubtitleRect) SetH(value int) {
	s.ptr.h = (C.int)(value)
}

// NbColors gets the nb_colors field.
func (s *AVSubtitleRect) NbColors() int {
	value := s.ptr.nb_colors
	return int(value)
}

// SetNbColors sets the nb_colors field.
func (s *AVSubtitleRect) SetNbColors(value int) {
	s.ptr.nb_colors = (C.int)(value)
}

// Data gets the data field.
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
func (s *AVSubtitleRect) Text() *CStr {
	value := s.ptr.text
	return wrapCStr(value)
}

// SetText sets the text field.
func (s *AVSubtitleRect) SetText(value *CStr) {
	s.ptr.text = value.ptr
}

// Ass gets the ass field.
func (s *AVSubtitleRect) Ass() *CStr {
	value := s.ptr.ass
	return wrapCStr(value)
}

// SetAss sets the ass field.
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
func (s *AVSubtitle) Format() uint16 {
	value := s.ptr.format
	return uint16(value)
}

// SetFormat sets the format field.
func (s *AVSubtitle) SetFormat(value uint16) {
	s.ptr.format = (C.uint16_t)(value)
}

// StartDisplayTime gets the start_display_time field.
func (s *AVSubtitle) StartDisplayTime() uint32 {
	value := s.ptr.start_display_time
	return uint32(value)
}

// SetStartDisplayTime sets the start_display_time field.
func (s *AVSubtitle) SetStartDisplayTime(value uint32) {
	s.ptr.start_display_time = (C.uint32_t)(value)
}

// EndDisplayTime gets the end_display_time field.
func (s *AVSubtitle) EndDisplayTime() uint32 {
	value := s.ptr.end_display_time
	return uint32(value)
}

// SetEndDisplayTime sets the end_display_time field.
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
func (s *AVSubtitle) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

// SetPts sets the pts field.
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
func (s *AVCodecParserContext) FrameOffset() int64 {
	value := s.ptr.frame_offset
	return int64(value)
}

// SetFrameOffset sets the frame_offset field.
func (s *AVCodecParserContext) SetFrameOffset(value int64) {
	s.ptr.frame_offset = (C.int64_t)(value)
}

// CurOffset gets the cur_offset field.
func (s *AVCodecParserContext) CurOffset() int64 {
	value := s.ptr.cur_offset
	return int64(value)
}

// SetCurOffset sets the cur_offset field.
func (s *AVCodecParserContext) SetCurOffset(value int64) {
	s.ptr.cur_offset = (C.int64_t)(value)
}

// NextFrameOffset gets the next_frame_offset field.
func (s *AVCodecParserContext) NextFrameOffset() int64 {
	value := s.ptr.next_frame_offset
	return int64(value)
}

// SetNextFrameOffset sets the next_frame_offset field.
func (s *AVCodecParserContext) SetNextFrameOffset(value int64) {
	s.ptr.next_frame_offset = (C.int64_t)(value)
}

// PictType gets the pict_type field.
func (s *AVCodecParserContext) PictType() int {
	value := s.ptr.pict_type
	return int(value)
}

// SetPictType sets the pict_type field.
func (s *AVCodecParserContext) SetPictType(value int) {
	s.ptr.pict_type = (C.int)(value)
}

// RepeatPict gets the repeat_pict field.
func (s *AVCodecParserContext) RepeatPict() int {
	value := s.ptr.repeat_pict
	return int(value)
}

// SetRepeatPict sets the repeat_pict field.
func (s *AVCodecParserContext) SetRepeatPict(value int) {
	s.ptr.repeat_pict = (C.int)(value)
}

// Pts gets the pts field.
func (s *AVCodecParserContext) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

// SetPts sets the pts field.
func (s *AVCodecParserContext) SetPts(value int64) {
	s.ptr.pts = (C.int64_t)(value)
}

// Dts gets the dts field.
func (s *AVCodecParserContext) Dts() int64 {
	value := s.ptr.dts
	return int64(value)
}

// SetDts sets the dts field.
func (s *AVCodecParserContext) SetDts(value int64) {
	s.ptr.dts = (C.int64_t)(value)
}

// LastPts gets the last_pts field.
func (s *AVCodecParserContext) LastPts() int64 {
	value := s.ptr.last_pts
	return int64(value)
}

// SetLastPts sets the last_pts field.
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
func (s *AVCodecParserContext) Offset() int64 {
	value := s.ptr.offset
	return int64(value)
}

// SetOffset sets the offset field.
func (s *AVCodecParserContext) SetOffset(value int64) {
	s.ptr.offset = (C.int64_t)(value)
}

// CurFrameEnd gets the cur_frame_end field.
func (s *AVCodecParserContext) CurFrameEnd() *Array[int64] {
	value := &s.ptr.cur_frame_end
	return ToInt64Array(unsafe.Pointer(value))
}

// KeyFrame gets the key_frame field.
func (s *AVCodecParserContext) KeyFrame() int {
	value := s.ptr.key_frame
	return int(value)
}

// SetKeyFrame sets the key_frame field.
func (s *AVCodecParserContext) SetKeyFrame(value int) {
	s.ptr.key_frame = (C.int)(value)
}

// DtsSyncPoint gets the dts_sync_point field.
func (s *AVCodecParserContext) DtsSyncPoint() int {
	value := s.ptr.dts_sync_point
	return int(value)
}

// SetDtsSyncPoint sets the dts_sync_point field.
func (s *AVCodecParserContext) SetDtsSyncPoint(value int) {
	s.ptr.dts_sync_point = (C.int)(value)
}

// DtsRefDtsDelta gets the dts_ref_dts_delta field.
func (s *AVCodecParserContext) DtsRefDtsDelta() int {
	value := s.ptr.dts_ref_dts_delta
	return int(value)
}

// SetDtsRefDtsDelta sets the dts_ref_dts_delta field.
func (s *AVCodecParserContext) SetDtsRefDtsDelta(value int) {
	s.ptr.dts_ref_dts_delta = (C.int)(value)
}

// PtsDtsDelta gets the pts_dts_delta field.
func (s *AVCodecParserContext) PtsDtsDelta() int {
	value := s.ptr.pts_dts_delta
	return int(value)
}

// SetPtsDtsDelta sets the pts_dts_delta field.
func (s *AVCodecParserContext) SetPtsDtsDelta(value int) {
	s.ptr.pts_dts_delta = (C.int)(value)
}

// CurFramePos gets the cur_frame_pos field.
func (s *AVCodecParserContext) CurFramePos() *Array[int64] {
	value := &s.ptr.cur_frame_pos
	return ToInt64Array(unsafe.Pointer(value))
}

// Pos gets the pos field.
func (s *AVCodecParserContext) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}

// SetPos sets the pos field.
func (s *AVCodecParserContext) SetPos(value int64) {
	s.ptr.pos = (C.int64_t)(value)
}

// LastPos gets the last_pos field.
func (s *AVCodecParserContext) LastPos() int64 {
	value := s.ptr.last_pos
	return int64(value)
}

// SetLastPos sets the last_pos field.
func (s *AVCodecParserContext) SetLastPos(value int64) {
	s.ptr.last_pos = (C.int64_t)(value)
}

// Duration gets the duration field.
func (s *AVCodecParserContext) Duration() int {
	value := s.ptr.duration
	return int(value)
}

// SetDuration sets the duration field.
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
func (s *AVCodecParserContext) PictureStructure() AVPictureStructure {
	value := s.ptr.picture_structure
	return AVPictureStructure(value)
}

// SetPictureStructure sets the picture_structure field.
func (s *AVCodecParserContext) SetPictureStructure(value AVPictureStructure) {
	s.ptr.picture_structure = (C.enum_AVPictureStructure)(value)
}

// OutputPictureNumber gets the output_picture_number field.
func (s *AVCodecParserContext) OutputPictureNumber() int {
	value := s.ptr.output_picture_number
	return int(value)
}

// SetOutputPictureNumber sets the output_picture_number field.
func (s *AVCodecParserContext) SetOutputPictureNumber(value int) {
	s.ptr.output_picture_number = (C.int)(value)
}

// Width gets the width field.
func (s *AVCodecParserContext) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
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
func (s *AVCodecParserContext) CodedWidth() int {
	value := s.ptr.coded_width
	return int(value)
}

// SetCodedWidth sets the coded_width field.
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
func (s *AVCodecParserContext) Format() int {
	value := s.ptr.format
	return int(value)
}

// SetFormat sets the format field.
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
func (s *AVProfile) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
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
func (s *AVCodec) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
func (s *AVCodec) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// LongName gets the long_name field.
func (s *AVCodec) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

// SetLongName sets the long_name field.
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
func (s *AVCodec) Capabilities() int {
	value := s.ptr.capabilities
	return int(value)
}

// SetCapabilities sets the capabilities field.
func (s *AVCodec) SetCapabilities(value int) {
	s.ptr.capabilities = (C.int)(value)
}

// MaxLowres gets the max_lowres field.
func (s *AVCodec) MaxLowres() uint8 {
	value := s.ptr.max_lowres
	return uint8(value)
}

// SetMaxLowres sets the max_lowres field.
func (s *AVCodec) SetMaxLowres(value uint8) {
	s.ptr.max_lowres = (C.uint8_t)(value)
}

// supported_framerates skipped due to struct value ptr

// PixFmts gets the pix_fmts field.
func (s *AVCodec) PixFmts() *Array[AVPixelFormat] {
	value := s.ptr.pix_fmts
	return ToAVPixelFormatArray(unsafe.Pointer(value))
}

// SetPixFmts sets the pix_fmts field.
func (s *AVCodec) SetPixFmts(value *Array[AVPixelFormat]) {
	if value != nil {
		s.ptr.pix_fmts = (*C.enum_AVPixelFormat)(value.ptr)
	} else {
		s.ptr.pix_fmts = nil
	}
}

// supported_samplerates skipped due to prim ptr

// SampleFmts gets the sample_fmts field.
func (s *AVCodec) SampleFmts() *Array[AVSampleFormat] {
	value := s.ptr.sample_fmts
	return ToAVSampleFormatArray(unsafe.Pointer(value))
}

// SetSampleFmts sets the sample_fmts field.
func (s *AVCodec) SetSampleFmts(value *Array[AVSampleFormat]) {
	if value != nil {
		s.ptr.sample_fmts = (*C.enum_AVSampleFormat)(value.ptr)
	} else {
		s.ptr.sample_fmts = nil
	}
}

// channel_layouts skipped due to prim ptr

// PrivClass gets the priv_class field.
func (s *AVCodec) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetPrivClass sets the priv_class field.
func (s *AVCodec) SetPrivClass(value *AVClass) {
	if value != nil {
		s.ptr.priv_class = value.ptr
	} else {
		s.ptr.priv_class = nil
	}
}

// Profiles gets the profiles field.
func (s *AVCodec) Profiles() *AVProfile {
	value := s.ptr.profiles
	var valueMapped *AVProfile
	if value != nil {
		valueMapped = &AVProfile{ptr: value}
	}
	return valueMapped
}

// SetProfiles sets the profiles field.
func (s *AVCodec) SetProfiles(value *AVProfile) {
	if value != nil {
		s.ptr.profiles = value.ptr
	} else {
		s.ptr.profiles = nil
	}
}

// WrapperName gets the wrapper_name field.
func (s *AVCodec) WrapperName() *CStr {
	value := s.ptr.wrapper_name
	return wrapCStr(value)
}

// SetWrapperName sets the wrapper_name field.
func (s *AVCodec) SetWrapperName(value *CStr) {
	s.ptr.wrapper_name = value.ptr
}

// ChLayouts gets the ch_layouts field.
func (s *AVCodec) ChLayouts() *AVChannelLayout {
	value := s.ptr.ch_layouts
	var valueMapped *AVChannelLayout
	if value != nil {
		valueMapped = &AVChannelLayout{ptr: value}
	}
	return valueMapped
}

// SetChLayouts sets the ch_layouts field.
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
func (s *AVCodecHWConfig) PixFmt() AVPixelFormat {
	value := s.ptr.pix_fmt
	return AVPixelFormat(value)
}

// SetPixFmt sets the pix_fmt field.
func (s *AVCodecHWConfig) SetPixFmt(value AVPixelFormat) {
	s.ptr.pix_fmt = (C.enum_AVPixelFormat)(value)
}

// Methods gets the methods field.
func (s *AVCodecHWConfig) Methods() int {
	value := s.ptr.methods
	return int(value)
}

// SetMethods sets the methods field.
func (s *AVCodecHWConfig) SetMethods(value int) {
	s.ptr.methods = (C.int)(value)
}

// DeviceType gets the device_type field.
func (s *AVCodecHWConfig) DeviceType() AVHWDeviceType {
	value := s.ptr.device_type
	return AVHWDeviceType(value)
}

// SetDeviceType sets the device_type field.
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
func (s *AVCodecDescriptor) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
func (s *AVCodecDescriptor) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// LongName gets the long_name field.
func (s *AVCodecDescriptor) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

// SetLongName sets the long_name field.
func (s *AVCodecDescriptor) SetLongName(value *CStr) {
	s.ptr.long_name = value.ptr
}

// Props gets the props field.
func (s *AVCodecDescriptor) Props() int {
	value := s.ptr.props
	return int(value)
}

// SetProps sets the props field.
func (s *AVCodecDescriptor) SetProps(value int) {
	s.ptr.props = (C.int)(value)
}

// mime_types skipped due to unknown ptr ptr

// Profiles gets the profiles field.
func (s *AVCodecDescriptor) Profiles() *AVProfile {
	value := s.ptr.profiles
	var valueMapped *AVProfile
	if value != nil {
		valueMapped = &AVProfile{ptr: value}
	}
	return valueMapped
}

// SetProfiles sets the profiles field.
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
func (s *AVCodecParameters) CodecType() AVMediaType {
	value := s.ptr.codec_type
	return AVMediaType(value)
}

// SetCodecType sets the codec_type field.
func (s *AVCodecParameters) SetCodecType(value AVMediaType) {
	s.ptr.codec_type = (C.enum_AVMediaType)(value)
}

// CodecId gets the codec_id field.
func (s *AVCodecParameters) CodecId() AVCodecID {
	value := s.ptr.codec_id
	return AVCodecID(value)
}

// SetCodecId sets the codec_id field.
func (s *AVCodecParameters) SetCodecId(value AVCodecID) {
	s.ptr.codec_id = (C.enum_AVCodecID)(value)
}

// CodecTag gets the codec_tag field.
func (s *AVCodecParameters) CodecTag() uint32 {
	value := s.ptr.codec_tag
	return uint32(value)
}

// SetCodecTag sets the codec_tag field.
func (s *AVCodecParameters) SetCodecTag(value uint32) {
	s.ptr.codec_tag = (C.uint32_t)(value)
}

// Extradata gets the extradata field.
func (s *AVCodecParameters) Extradata() unsafe.Pointer {
	value := s.ptr.extradata
	return unsafe.Pointer(value)
}

// SetExtradata sets the extradata field.
func (s *AVCodecParameters) SetExtradata(value unsafe.Pointer) {
	s.ptr.extradata = (*C.uint8_t)(value)
}

// ExtradataSize gets the extradata_size field.
func (s *AVCodecParameters) ExtradataSize() int {
	value := s.ptr.extradata_size
	return int(value)
}

// SetExtradataSize sets the extradata_size field.
func (s *AVCodecParameters) SetExtradataSize(value int) {
	s.ptr.extradata_size = (C.int)(value)
}

// Format gets the format field.
func (s *AVCodecParameters) Format() int {
	value := s.ptr.format
	return int(value)
}

// SetFormat sets the format field.
func (s *AVCodecParameters) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

// BitRate gets the bit_rate field.
func (s *AVCodecParameters) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}

// SetBitRate sets the bit_rate field.
func (s *AVCodecParameters) SetBitRate(value int64) {
	s.ptr.bit_rate = (C.int64_t)(value)
}

// BitsPerCodedSample gets the bits_per_coded_sample field.
func (s *AVCodecParameters) BitsPerCodedSample() int {
	value := s.ptr.bits_per_coded_sample
	return int(value)
}

// SetBitsPerCodedSample sets the bits_per_coded_sample field.
func (s *AVCodecParameters) SetBitsPerCodedSample(value int) {
	s.ptr.bits_per_coded_sample = (C.int)(value)
}

// BitsPerRawSample gets the bits_per_raw_sample field.
func (s *AVCodecParameters) BitsPerRawSample() int {
	value := s.ptr.bits_per_raw_sample
	return int(value)
}

// SetBitsPerRawSample sets the bits_per_raw_sample field.
func (s *AVCodecParameters) SetBitsPerRawSample(value int) {
	s.ptr.bits_per_raw_sample = (C.int)(value)
}

// Profile gets the profile field.
func (s *AVCodecParameters) Profile() int {
	value := s.ptr.profile
	return int(value)
}

// SetProfile sets the profile field.
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
func (s *AVCodecParameters) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
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
func (s *AVCodecParameters) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
func (s *AVCodecParameters) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

// FieldOrder gets the field_order field.
func (s *AVCodecParameters) FieldOrder() AVFieldOrder {
	value := s.ptr.field_order
	return AVFieldOrder(value)
}

// SetFieldOrder sets the field_order field.
func (s *AVCodecParameters) SetFieldOrder(value AVFieldOrder) {
	s.ptr.field_order = (C.enum_AVFieldOrder)(value)
}

// ColorRange gets the color_range field.
func (s *AVCodecParameters) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}

// SetColorRange sets the color_range field.
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
func (s *AVCodecParameters) VideoDelay() int {
	value := s.ptr.video_delay
	return int(value)
}

// SetVideoDelay sets the video_delay field.
func (s *AVCodecParameters) SetVideoDelay(value int) {
	s.ptr.video_delay = (C.int)(value)
}

// ChannelLayout gets the channel_layout field.
func (s *AVCodecParameters) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

// SetChannelLayout sets the channel_layout field.
func (s *AVCodecParameters) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

// Channels gets the channels field.
func (s *AVCodecParameters) Channels() int {
	value := s.ptr.channels
	return int(value)
}

// SetChannels sets the channels field.
func (s *AVCodecParameters) SetChannels(value int) {
	s.ptr.channels = (C.int)(value)
}

// SampleRate gets the sample_rate field.
func (s *AVCodecParameters) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

// SetSampleRate sets the sample_rate field.
func (s *AVCodecParameters) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

// BlockAlign gets the block_align field.
func (s *AVCodecParameters) BlockAlign() int {
	value := s.ptr.block_align
	return int(value)
}

// SetBlockAlign sets the block_align field.
func (s *AVCodecParameters) SetBlockAlign(value int) {
	s.ptr.block_align = (C.int)(value)
}

// FrameSize gets the frame_size field.
func (s *AVCodecParameters) FrameSize() int {
	value := s.ptr.frame_size
	return int(value)
}

// SetFrameSize sets the frame_size field.
func (s *AVCodecParameters) SetFrameSize(value int) {
	s.ptr.frame_size = (C.int)(value)
}

// InitialPadding gets the initial_padding field.
func (s *AVCodecParameters) InitialPadding() int {
	value := s.ptr.initial_padding
	return int(value)
}

// SetInitialPadding sets the initial_padding field.
func (s *AVCodecParameters) SetInitialPadding(value int) {
	s.ptr.initial_padding = (C.int)(value)
}

// TrailingPadding gets the trailing_padding field.
func (s *AVCodecParameters) TrailingPadding() int {
	value := s.ptr.trailing_padding
	return int(value)
}

// SetTrailingPadding sets the trailing_padding field.
func (s *AVCodecParameters) SetTrailingPadding(value int) {
	s.ptr.trailing_padding = (C.int)(value)
}

// SeekPreroll gets the seek_preroll field.
func (s *AVCodecParameters) SeekPreroll() int {
	value := s.ptr.seek_preroll
	return int(value)
}

// SetSeekPreroll sets the seek_preroll field.
func (s *AVCodecParameters) SetSeekPreroll(value int) {
	s.ptr.seek_preroll = (C.int)(value)
}

// ChLayout gets the ch_layout field.
func (s *AVCodecParameters) ChLayout() *AVChannelLayout {
	value := &s.ptr.ch_layout
	return &AVChannelLayout{ptr: value}
}

// Framerate gets the framerate field.
func (s *AVCodecParameters) Framerate() *AVRational {
	value := s.ptr.framerate
	return &AVRational{value: value}
}

// SetFramerate sets the framerate field.
func (s *AVCodecParameters) SetFramerate(value *AVRational) {
	s.ptr.framerate = value.value
}

// CodedSideData gets the coded_side_data field.
func (s *AVCodecParameters) CodedSideData() *AVPacketSideData {
	value := s.ptr.coded_side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}

// SetCodedSideData sets the coded_side_data field.
func (s *AVCodecParameters) SetCodedSideData(value *AVPacketSideData) {
	if value != nil {
		s.ptr.coded_side_data = value.ptr
	} else {
		s.ptr.coded_side_data = nil
	}
}

// NbCodedSideData gets the nb_coded_side_data field.
func (s *AVCodecParameters) NbCodedSideData() int {
	value := s.ptr.nb_coded_side_data
	return int(value)
}

// SetNbCodedSideData sets the nb_coded_side_data field.
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
func (s *AVPanScan) Id() int {
	value := s.ptr.id
	return int(value)
}

// SetId sets the id field.
func (s *AVPanScan) SetId(value int) {
	s.ptr.id = (C.int)(value)
}

// Width gets the width field.
func (s *AVPanScan) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
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
func (s *AVCPBProperties) MaxBitrate() int64 {
	value := s.ptr.max_bitrate
	return int64(value)
}

// SetMaxBitrate sets the max_bitrate field.
func (s *AVCPBProperties) SetMaxBitrate(value int64) {
	s.ptr.max_bitrate = (C.int64_t)(value)
}

// MinBitrate gets the min_bitrate field.
func (s *AVCPBProperties) MinBitrate() int64 {
	value := s.ptr.min_bitrate
	return int64(value)
}

// SetMinBitrate sets the min_bitrate field.
func (s *AVCPBProperties) SetMinBitrate(value int64) {
	s.ptr.min_bitrate = (C.int64_t)(value)
}

// AvgBitrate gets the avg_bitrate field.
func (s *AVCPBProperties) AvgBitrate() int64 {
	value := s.ptr.avg_bitrate
	return int64(value)
}

// SetAvgBitrate sets the avg_bitrate field.
func (s *AVCPBProperties) SetAvgBitrate(value int64) {
	s.ptr.avg_bitrate = (C.int64_t)(value)
}

// BufferSize gets the buffer_size field.
func (s *AVCPBProperties) BufferSize() int64 {
	value := s.ptr.buffer_size
	return int64(value)
}

// SetBufferSize sets the buffer_size field.
func (s *AVCPBProperties) SetBufferSize(value int64) {
	s.ptr.buffer_size = (C.int64_t)(value)
}

// VbvDelay gets the vbv_delay field.
func (s *AVCPBProperties) VbvDelay() uint64 {
	value := s.ptr.vbv_delay
	return uint64(value)
}

// SetVbvDelay sets the vbv_delay field.
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
func (s *AVProducerReferenceTime) Wallclock() int64 {
	value := s.ptr.wallclock
	return int64(value)
}

// SetWallclock sets the wallclock field.
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
func (s *AVPacket) Buf() *AVBufferRef {
	value := s.ptr.buf
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetBuf sets the buf field.
func (s *AVPacket) SetBuf(value *AVBufferRef) {
	if value != nil {
		s.ptr.buf = value.ptr
	} else {
		s.ptr.buf = nil
	}
}

// Pts gets the pts field.
func (s *AVPacket) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

// SetPts sets the pts field.
func (s *AVPacket) SetPts(value int64) {
	s.ptr.pts = (C.int64_t)(value)
}

// Dts gets the dts field.
func (s *AVPacket) Dts() int64 {
	value := s.ptr.dts
	return int64(value)
}

// SetDts sets the dts field.
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
func (s *AVPacket) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVPacket) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// SideData gets the side_data field.
func (s *AVPacket) SideData() *AVPacketSideData {
	value := s.ptr.side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}

// SetSideData sets the side_data field.
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
func (s *AVPacket) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

// SetDuration sets the duration field.
func (s *AVPacket) SetDuration(value int64) {
	s.ptr.duration = (C.int64_t)(value)
}

// Pos gets the pos field.
func (s *AVPacket) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}

// SetPos sets the pos field.
func (s *AVPacket) SetPos(value int64) {
	s.ptr.pos = (C.int64_t)(value)
}

// Opaque gets the opaque field.
func (s *AVPacket) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
func (s *AVPacket) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// OpaqueRef gets the opaque_ref field.
func (s *AVPacket) OpaqueRef() *AVBufferRef {
	value := s.ptr.opaque_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetOpaqueRef sets the opaque_ref field.
func (s *AVPacket) SetOpaqueRef(value *AVBufferRef) {
	if value != nil {
		s.ptr.opaque_ref = value.ptr
	} else {
		s.ptr.opaque_ref = nil
	}
}

// TimeBase gets the time_base field.
func (s *AVPacket) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
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
func (s *AVFilterContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
func (s *AVFilterContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Filter gets the filter field.
func (s *AVFilterContext) Filter() *AVFilter {
	value := s.ptr.filter
	var valueMapped *AVFilter
	if value != nil {
		valueMapped = &AVFilter{ptr: value}
	}
	return valueMapped
}

// SetFilter sets the filter field.
func (s *AVFilterContext) SetFilter(value *AVFilter) {
	if value != nil {
		s.ptr.filter = value.ptr
	} else {
		s.ptr.filter = nil
	}
}

// Name gets the name field.
func (s *AVFilterContext) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
func (s *AVFilterContext) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// InputPads gets the input_pads field.
func (s *AVFilterContext) InputPads() *AVFilterPad {
	value := s.ptr.input_pads
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetInputPads sets the input_pads field.
func (s *AVFilterContext) SetInputPads(value *AVFilterPad) {
	if value != nil {
		s.ptr.input_pads = value.ptr
	} else {
		s.ptr.input_pads = nil
	}
}

// Inputs gets the inputs field.
func (s *AVFilterContext) Inputs() *Array[*AVFilterLink] {
	value := s.ptr.inputs
	return ToAVFilterLinkArray(unsafe.Pointer(value))
}

// SetInputs sets the inputs field.
func (s *AVFilterContext) SetInputs(value *Array[AVFilterLink]) {
	if value != nil {
		s.ptr.inputs = (**C.AVFilterLink)(value.ptr)
	} else {
		s.ptr.inputs = nil
	}
}

// NbInputs gets the nb_inputs field.
func (s *AVFilterContext) NbInputs() uint {
	value := s.ptr.nb_inputs
	return uint(value)
}

// SetNbInputs sets the nb_inputs field.
func (s *AVFilterContext) SetNbInputs(value uint) {
	s.ptr.nb_inputs = (C.uint)(value)
}

// OutputPads gets the output_pads field.
func (s *AVFilterContext) OutputPads() *AVFilterPad {
	value := s.ptr.output_pads
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetOutputPads sets the output_pads field.
func (s *AVFilterContext) SetOutputPads(value *AVFilterPad) {
	if value != nil {
		s.ptr.output_pads = value.ptr
	} else {
		s.ptr.output_pads = nil
	}
}

// Outputs gets the outputs field.
func (s *AVFilterContext) Outputs() *Array[*AVFilterLink] {
	value := s.ptr.outputs
	return ToAVFilterLinkArray(unsafe.Pointer(value))
}

// SetOutputs sets the outputs field.
func (s *AVFilterContext) SetOutputs(value *Array[AVFilterLink]) {
	if value != nil {
		s.ptr.outputs = (**C.AVFilterLink)(value.ptr)
	} else {
		s.ptr.outputs = nil
	}
}

// NbOutputs gets the nb_outputs field.
func (s *AVFilterContext) NbOutputs() uint {
	value := s.ptr.nb_outputs
	return uint(value)
}

// SetNbOutputs sets the nb_outputs field.
func (s *AVFilterContext) SetNbOutputs(value uint) {
	s.ptr.nb_outputs = (C.uint)(value)
}

// Priv gets the priv field.
func (s *AVFilterContext) Priv() unsafe.Pointer {
	value := s.ptr.priv
	return value
}

// SetPriv sets the priv field.
func (s *AVFilterContext) SetPriv(value unsafe.Pointer) {
	s.ptr.priv = value
}

// Graph gets the graph field.
func (s *AVFilterContext) Graph() *AVFilterGraph {
	value := s.ptr.graph
	var valueMapped *AVFilterGraph
	if value != nil {
		valueMapped = &AVFilterGraph{ptr: value}
	}
	return valueMapped
}

// SetGraph sets the graph field.
func (s *AVFilterContext) SetGraph(value *AVFilterGraph) {
	if value != nil {
		s.ptr.graph = value.ptr
	} else {
		s.ptr.graph = nil
	}
}

// ThreadType gets the thread_type field.
func (s *AVFilterContext) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}

// SetThreadType sets the thread_type field.
func (s *AVFilterContext) SetThreadType(value int) {
	s.ptr.thread_type = (C.int)(value)
}

// Internal gets the internal field.
func (s *AVFilterContext) Internal() *AVFilterInternal {
	value := s.ptr.internal
	var valueMapped *AVFilterInternal
	if value != nil {
		valueMapped = &AVFilterInternal{ptr: value}
	}
	return valueMapped
}

// SetInternal sets the internal field.
func (s *AVFilterContext) SetInternal(value *AVFilterInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

// command_queue skipped due to ptr to ignored type

// EnableStr gets the enable_str field.
func (s *AVFilterContext) EnableStr() *CStr {
	value := s.ptr.enable_str
	return wrapCStr(value)
}

// SetEnableStr sets the enable_str field.
func (s *AVFilterContext) SetEnableStr(value *CStr) {
	s.ptr.enable_str = value.ptr
}

// Enable gets the enable field.
func (s *AVFilterContext) Enable() unsafe.Pointer {
	value := s.ptr.enable
	return value
}

// SetEnable sets the enable field.
func (s *AVFilterContext) SetEnable(value unsafe.Pointer) {
	s.ptr.enable = value
}

// var_values skipped due to prim ptr

// IsDisabled gets the is_disabled field.
func (s *AVFilterContext) IsDisabled() int {
	value := s.ptr.is_disabled
	return int(value)
}

// SetIsDisabled sets the is_disabled field.
func (s *AVFilterContext) SetIsDisabled(value int) {
	s.ptr.is_disabled = (C.int)(value)
}

// HwDeviceCtx gets the hw_device_ctx field.
func (s *AVFilterContext) HwDeviceCtx() *AVBufferRef {
	value := s.ptr.hw_device_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwDeviceCtx sets the hw_device_ctx field.
func (s *AVFilterContext) SetHwDeviceCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_device_ctx = value.ptr
	} else {
		s.ptr.hw_device_ctx = nil
	}
}

// NbThreads gets the nb_threads field.
func (s *AVFilterContext) NbThreads() int {
	value := s.ptr.nb_threads
	return int(value)
}

// SetNbThreads sets the nb_threads field.
func (s *AVFilterContext) SetNbThreads(value int) {
	s.ptr.nb_threads = (C.int)(value)
}

// Ready gets the ready field.
func (s *AVFilterContext) Ready() uint {
	value := s.ptr.ready
	return uint(value)
}

// SetReady sets the ready field.
func (s *AVFilterContext) SetReady(value uint) {
	s.ptr.ready = (C.uint)(value)
}

// ExtraHwFrames gets the extra_hw_frames field.
func (s *AVFilterContext) ExtraHwFrames() int {
	value := s.ptr.extra_hw_frames
	return int(value)
}

// SetExtraHwFrames sets the extra_hw_frames field.
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
func (s *AVFilterLink) Src() *AVFilterContext {
	value := s.ptr.src
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

// SetSrc sets the src field.
func (s *AVFilterLink) SetSrc(value *AVFilterContext) {
	if value != nil {
		s.ptr.src = value.ptr
	} else {
		s.ptr.src = nil
	}
}

// Srcpad gets the srcpad field.
func (s *AVFilterLink) Srcpad() *AVFilterPad {
	value := s.ptr.srcpad
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetSrcpad sets the srcpad field.
func (s *AVFilterLink) SetSrcpad(value *AVFilterPad) {
	if value != nil {
		s.ptr.srcpad = value.ptr
	} else {
		s.ptr.srcpad = nil
	}
}

// Dst gets the dst field.
func (s *AVFilterLink) Dst() *AVFilterContext {
	value := s.ptr.dst
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

// SetDst sets the dst field.
func (s *AVFilterLink) SetDst(value *AVFilterContext) {
	if value != nil {
		s.ptr.dst = value.ptr
	} else {
		s.ptr.dst = nil
	}
}

// Dstpad gets the dstpad field.
func (s *AVFilterLink) Dstpad() *AVFilterPad {
	value := s.ptr.dstpad
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetDstpad sets the dstpad field.
func (s *AVFilterLink) SetDstpad(value *AVFilterPad) {
	if value != nil {
		s.ptr.dstpad = value.ptr
	} else {
		s.ptr.dstpad = nil
	}
}

// Type gets the type field.
func (s *AVFilterLink) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}

// SetType sets the type field.
func (s *AVFilterLink) SetType(value AVMediaType) {
	s.ptr._type = (C.enum_AVMediaType)(value)
}

// W gets the w field.
func (s *AVFilterLink) W() int {
	value := s.ptr.w
	return int(value)
}

// SetW sets the w field.
func (s *AVFilterLink) SetW(value int) {
	s.ptr.w = (C.int)(value)
}

// H gets the h field.
func (s *AVFilterLink) H() int {
	value := s.ptr.h
	return int(value)
}

// SetH sets the h field.
func (s *AVFilterLink) SetH(value int) {
	s.ptr.h = (C.int)(value)
}

// SampleAspectRatio gets the sample_aspect_ratio field.
func (s *AVFilterLink) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
func (s *AVFilterLink) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

// ChannelLayout gets the channel_layout field.
func (s *AVFilterLink) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

// SetChannelLayout sets the channel_layout field.
func (s *AVFilterLink) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

// SampleRate gets the sample_rate field.
func (s *AVFilterLink) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

// SetSampleRate sets the sample_rate field.
func (s *AVFilterLink) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

// Format gets the format field.
func (s *AVFilterLink) Format() int {
	value := s.ptr.format
	return int(value)
}

// SetFormat sets the format field.
func (s *AVFilterLink) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

// TimeBase gets the time_base field.
func (s *AVFilterLink) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
func (s *AVFilterLink) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// ChLayout gets the ch_layout field.
func (s *AVFilterLink) ChLayout() *AVChannelLayout {
	value := &s.ptr.ch_layout
	return &AVChannelLayout{ptr: value}
}

// Incfg gets the incfg field.
func (s *AVFilterLink) Incfg() *AVFilterFormatsConfig {
	value := &s.ptr.incfg
	return &AVFilterFormatsConfig{ptr: value}
}

// Outcfg gets the outcfg field.
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
func (s *AVFilterLink) Graph() *AVFilterGraph {
	value := s.ptr.graph
	var valueMapped *AVFilterGraph
	if value != nil {
		valueMapped = &AVFilterGraph{ptr: value}
	}
	return valueMapped
}

// SetGraph sets the graph field.
func (s *AVFilterLink) SetGraph(value *AVFilterGraph) {
	if value != nil {
		s.ptr.graph = value.ptr
	} else {
		s.ptr.graph = nil
	}
}

// CurrentPts gets the current_pts field.
func (s *AVFilterLink) CurrentPts() int64 {
	value := s.ptr.current_pts
	return int64(value)
}

// SetCurrentPts sets the current_pts field.
func (s *AVFilterLink) SetCurrentPts(value int64) {
	s.ptr.current_pts = (C.int64_t)(value)
}

// CurrentPtsUs gets the current_pts_us field.
func (s *AVFilterLink) CurrentPtsUs() int64 {
	value := s.ptr.current_pts_us
	return int64(value)
}

// SetCurrentPtsUs sets the current_pts_us field.
func (s *AVFilterLink) SetCurrentPtsUs(value int64) {
	s.ptr.current_pts_us = (C.int64_t)(value)
}

// AgeIndex gets the age_index field.
func (s *AVFilterLink) AgeIndex() int {
	value := s.ptr.age_index
	return int(value)
}

// SetAgeIndex sets the age_index field.
func (s *AVFilterLink) SetAgeIndex(value int) {
	s.ptr.age_index = (C.int)(value)
}

// FrameRate gets the frame_rate field.
func (s *AVFilterLink) FrameRate() *AVRational {
	value := s.ptr.frame_rate
	return &AVRational{value: value}
}

// SetFrameRate sets the frame_rate field.
func (s *AVFilterLink) SetFrameRate(value *AVRational) {
	s.ptr.frame_rate = value.value
}

// MinSamples gets the min_samples field.
func (s *AVFilterLink) MinSamples() int {
	value := s.ptr.min_samples
	return int(value)
}

// SetMinSamples sets the min_samples field.
func (s *AVFilterLink) SetMinSamples(value int) {
	s.ptr.min_samples = (C.int)(value)
}

// MaxSamples gets the max_samples field.
func (s *AVFilterLink) MaxSamples() int {
	value := s.ptr.max_samples
	return int(value)
}

// SetMaxSamples sets the max_samples field.
func (s *AVFilterLink) SetMaxSamples(value int) {
	s.ptr.max_samples = (C.int)(value)
}

// FrameCountIn gets the frame_count_in field.
func (s *AVFilterLink) FrameCountIn() int64 {
	value := s.ptr.frame_count_in
	return int64(value)
}

// SetFrameCountIn sets the frame_count_in field.
func (s *AVFilterLink) SetFrameCountIn(value int64) {
	s.ptr.frame_count_in = (C.int64_t)(value)
}

// FrameCountOut gets the frame_count_out field.
func (s *AVFilterLink) FrameCountOut() int64 {
	value := s.ptr.frame_count_out
	return int64(value)
}

// SetFrameCountOut sets the frame_count_out field.
func (s *AVFilterLink) SetFrameCountOut(value int64) {
	s.ptr.frame_count_out = (C.int64_t)(value)
}

// SampleCountIn gets the sample_count_in field.
func (s *AVFilterLink) SampleCountIn() int64 {
	value := s.ptr.sample_count_in
	return int64(value)
}

// SetSampleCountIn sets the sample_count_in field.
func (s *AVFilterLink) SetSampleCountIn(value int64) {
	s.ptr.sample_count_in = (C.int64_t)(value)
}

// SampleCountOut gets the sample_count_out field.
func (s *AVFilterLink) SampleCountOut() int64 {
	value := s.ptr.sample_count_out
	return int64(value)
}

// SetSampleCountOut sets the sample_count_out field.
func (s *AVFilterLink) SetSampleCountOut(value int64) {
	s.ptr.sample_count_out = (C.int64_t)(value)
}

// FramePool gets the frame_pool field.
func (s *AVFilterLink) FramePool() unsafe.Pointer {
	value := s.ptr.frame_pool
	return value
}

// SetFramePool sets the frame_pool field.
func (s *AVFilterLink) SetFramePool(value unsafe.Pointer) {
	s.ptr.frame_pool = value
}

// FrameWantedOut gets the frame_wanted_out field.
func (s *AVFilterLink) FrameWantedOut() int {
	value := s.ptr.frame_wanted_out
	return int(value)
}

// SetFrameWantedOut sets the frame_wanted_out field.
func (s *AVFilterLink) SetFrameWantedOut(value int) {
	s.ptr.frame_wanted_out = (C.int)(value)
}

// HwFramesCtx gets the hw_frames_ctx field.
func (s *AVFilterLink) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwFramesCtx sets the hw_frames_ctx field.
func (s *AVFilterLink) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

// Reserved gets the reserved field.
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
func (s *AVFilter) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
func (s *AVFilter) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// Description gets the description field.
func (s *AVFilter) Description() *CStr {
	value := s.ptr.description
	return wrapCStr(value)
}

// SetDescription sets the description field.
func (s *AVFilter) SetDescription(value *CStr) {
	s.ptr.description = value.ptr
}

// Inputs gets the inputs field.
func (s *AVFilter) Inputs() *AVFilterPad {
	value := s.ptr.inputs
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetInputs sets the inputs field.
func (s *AVFilter) SetInputs(value *AVFilterPad) {
	if value != nil {
		s.ptr.inputs = value.ptr
	} else {
		s.ptr.inputs = nil
	}
}

// Outputs gets the outputs field.
func (s *AVFilter) Outputs() *AVFilterPad {
	value := s.ptr.outputs
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}

// SetOutputs sets the outputs field.
func (s *AVFilter) SetOutputs(value *AVFilterPad) {
	if value != nil {
		s.ptr.outputs = value.ptr
	} else {
		s.ptr.outputs = nil
	}
}

// PrivClass gets the priv_class field.
func (s *AVFilter) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetPrivClass sets the priv_class field.
func (s *AVFilter) SetPrivClass(value *AVClass) {
	if value != nil {
		s.ptr.priv_class = value.ptr
	} else {
		s.ptr.priv_class = nil
	}
}

// Flags gets the flags field.
func (s *AVFilter) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVFilter) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// NbInputs gets the nb_inputs field.
func (s *AVFilter) NbInputs() uint8 {
	value := s.ptr.nb_inputs
	return uint8(value)
}

// SetNbInputs sets the nb_inputs field.
func (s *AVFilter) SetNbInputs(value uint8) {
	s.ptr.nb_inputs = (C.uint8_t)(value)
}

// NbOutputs gets the nb_outputs field.
func (s *AVFilter) NbOutputs() uint8 {
	value := s.ptr.nb_outputs
	return uint8(value)
}

// SetNbOutputs sets the nb_outputs field.
func (s *AVFilter) SetNbOutputs(value uint8) {
	s.ptr.nb_outputs = (C.uint8_t)(value)
}

// FormatsState gets the formats_state field.
func (s *AVFilter) FormatsState() uint8 {
	value := s.ptr.formats_state
	return uint8(value)
}

// SetFormatsState sets the formats_state field.
func (s *AVFilter) SetFormatsState(value uint8) {
	s.ptr.formats_state = (C.uint8_t)(value)
}

// preinit skipped due to func ptr

// init skipped due to func ptr

// uninit skipped due to func ptr

// formats skipped due to union type

// PrivSize gets the priv_size field.
func (s *AVFilter) PrivSize() int {
	value := s.ptr.priv_size
	return int(value)
}

// SetPrivSize sets the priv_size field.
func (s *AVFilter) SetPrivSize(value int) {
	s.ptr.priv_size = (C.int)(value)
}

// FlagsInternal gets the flags_internal field.
func (s *AVFilter) FlagsInternal() int {
	value := s.ptr.flags_internal
	return int(value)
}

// SetFlagsInternal sets the flags_internal field.
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
func (s *AVFilterFormatsConfig) Formats() *AVFilterFormats {
	value := s.ptr.formats
	var valueMapped *AVFilterFormats
	if value != nil {
		valueMapped = &AVFilterFormats{ptr: value}
	}
	return valueMapped
}

// SetFormats sets the formats field.
func (s *AVFilterFormatsConfig) SetFormats(value *AVFilterFormats) {
	if value != nil {
		s.ptr.formats = value.ptr
	} else {
		s.ptr.formats = nil
	}
}

// Samplerates gets the samplerates field.
func (s *AVFilterFormatsConfig) Samplerates() *AVFilterFormats {
	value := s.ptr.samplerates
	var valueMapped *AVFilterFormats
	if value != nil {
		valueMapped = &AVFilterFormats{ptr: value}
	}
	return valueMapped
}

// SetSamplerates sets the samplerates field.
func (s *AVFilterFormatsConfig) SetSamplerates(value *AVFilterFormats) {
	if value != nil {
		s.ptr.samplerates = value.ptr
	} else {
		s.ptr.samplerates = nil
	}
}

// ChannelLayouts gets the channel_layouts field.
func (s *AVFilterFormatsConfig) ChannelLayouts() *AVFilterChannelLayouts {
	value := s.ptr.channel_layouts
	var valueMapped *AVFilterChannelLayouts
	if value != nil {
		valueMapped = &AVFilterChannelLayouts{ptr: value}
	}
	return valueMapped
}

// SetChannelLayouts sets the channel_layouts field.
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
func (s *AVFilterGraph) ScaleSwsOpts() *CStr {
	value := s.ptr.scale_sws_opts
	return wrapCStr(value)
}

// SetScaleSwsOpts sets the scale_sws_opts field.
func (s *AVFilterGraph) SetScaleSwsOpts(value *CStr) {
	s.ptr.scale_sws_opts = value.ptr
}

// ThreadType gets the thread_type field.
func (s *AVFilterGraph) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}

// SetThreadType sets the thread_type field.
func (s *AVFilterGraph) SetThreadType(value int) {
	s.ptr.thread_type = (C.int)(value)
}

// NbThreads gets the nb_threads field.
func (s *AVFilterGraph) NbThreads() int {
	value := s.ptr.nb_threads
	return int(value)
}

// SetNbThreads sets the nb_threads field.
func (s *AVFilterGraph) SetNbThreads(value int) {
	s.ptr.nb_threads = (C.int)(value)
}

// Internal gets the internal field.
func (s *AVFilterGraph) Internal() *AVFilterGraphInternal {
	value := s.ptr.internal
	var valueMapped *AVFilterGraphInternal
	if value != nil {
		valueMapped = &AVFilterGraphInternal{ptr: value}
	}
	return valueMapped
}

// SetInternal sets the internal field.
func (s *AVFilterGraph) SetInternal(value *AVFilterGraphInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

// Opaque gets the opaque field.
func (s *AVFilterGraph) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
func (s *AVFilterGraph) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// execute skipped due to callback ptr

// AresampleSwrOpts gets the aresample_swr_opts field.
func (s *AVFilterGraph) AresampleSwrOpts() *CStr {
	value := s.ptr.aresample_swr_opts
	return wrapCStr(value)
}

// SetAresampleSwrOpts sets the aresample_swr_opts field.
func (s *AVFilterGraph) SetAresampleSwrOpts(value *CStr) {
	s.ptr.aresample_swr_opts = value.ptr
}

// SinkLinks gets the sink_links field.
func (s *AVFilterGraph) SinkLinks() *Array[*AVFilterLink] {
	value := s.ptr.sink_links
	return ToAVFilterLinkArray(unsafe.Pointer(value))
}

// SetSinkLinks sets the sink_links field.
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
func (s *AVFilterInOut) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
func (s *AVFilterInOut) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// FilterCtx gets the filter_ctx field.
func (s *AVFilterInOut) FilterCtx() *AVFilterContext {
	value := s.ptr.filter_ctx
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

// SetFilterCtx sets the filter_ctx field.
func (s *AVFilterInOut) SetFilterCtx(value *AVFilterContext) {
	if value != nil {
		s.ptr.filter_ctx = value.ptr
	} else {
		s.ptr.filter_ctx = nil
	}
}

// PadIdx gets the pad_idx field.
func (s *AVFilterInOut) PadIdx() int {
	value := s.ptr.pad_idx
	return int(value)
}

// SetPadIdx sets the pad_idx field.
func (s *AVFilterInOut) SetPadIdx(value int) {
	s.ptr.pad_idx = (C.int)(value)
}

// Next gets the next field.
func (s *AVFilterInOut) Next() *AVFilterInOut {
	value := s.ptr.next
	var valueMapped *AVFilterInOut
	if value != nil {
		valueMapped = &AVFilterInOut{ptr: value}
	}
	return valueMapped
}

// SetNext sets the next field.
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
func (s *AVFilterPadParams) Label() *CStr {
	value := s.ptr.label
	return wrapCStr(value)
}

// SetLabel sets the label field.
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
func (s *AVFilterParams) Filter() *AVFilterContext {
	value := s.ptr.filter
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}

// SetFilter sets the filter field.
func (s *AVFilterParams) SetFilter(value *AVFilterContext) {
	if value != nil {
		s.ptr.filter = value.ptr
	} else {
		s.ptr.filter = nil
	}
}

// FilterName gets the filter_name field.
func (s *AVFilterParams) FilterName() *CStr {
	value := s.ptr.filter_name
	return wrapCStr(value)
}

// SetFilterName sets the filter_name field.
func (s *AVFilterParams) SetFilterName(value *CStr) {
	s.ptr.filter_name = value.ptr
}

// InstanceName gets the instance_name field.
func (s *AVFilterParams) InstanceName() *CStr {
	value := s.ptr.instance_name
	return wrapCStr(value)
}

// SetInstanceName sets the instance_name field.
func (s *AVFilterParams) SetInstanceName(value *CStr) {
	s.ptr.instance_name = value.ptr
}

// Opts gets the opts field.
func (s *AVFilterParams) Opts() *AVDictionary {
	value := s.ptr.opts
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

// SetOpts sets the opts field.
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
func (s *AVFilterGraphSegment) Graph() *AVFilterGraph {
	value := s.ptr.graph
	var valueMapped *AVFilterGraph
	if value != nil {
		valueMapped = &AVFilterGraph{ptr: value}
	}
	return valueMapped
}

// SetGraph sets the graph field.
func (s *AVFilterGraphSegment) SetGraph(value *AVFilterGraph) {
	if value != nil {
		s.ptr.graph = value.ptr
	} else {
		s.ptr.graph = nil
	}
}

// Chains gets the chains field.
func (s *AVFilterGraphSegment) Chains() *Array[*AVFilterChain] {
	value := s.ptr.chains
	return ToAVFilterChainArray(unsafe.Pointer(value))
}

// SetChains sets the chains field.
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
func (s *AVFilterGraphSegment) ScaleSwsOpts() *CStr {
	value := s.ptr.scale_sws_opts
	return wrapCStr(value)
}

// SetScaleSwsOpts sets the scale_sws_opts field.
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
func (s *AVBufferSrcParameters) Format() int {
	value := s.ptr.format
	return int(value)
}

// SetFormat sets the format field.
func (s *AVBufferSrcParameters) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

// TimeBase gets the time_base field.
func (s *AVBufferSrcParameters) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
func (s *AVBufferSrcParameters) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// Width gets the width field.
func (s *AVBufferSrcParameters) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
func (s *AVBufferSrcParameters) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

// Height gets the height field.
func (s *AVBufferSrcParameters) Height() int {
	value := s.ptr.height
	return int(value)
}

// SetHeight sets the height field.
func (s *AVBufferSrcParameters) SetHeight(value int) {
	s.ptr.height = (C.int)(value)
}

// SampleAspectRatio gets the sample_aspect_ratio field.
func (s *AVBufferSrcParameters) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
func (s *AVBufferSrcParameters) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

// FrameRate gets the frame_rate field.
func (s *AVBufferSrcParameters) FrameRate() *AVRational {
	value := s.ptr.frame_rate
	return &AVRational{value: value}
}

// SetFrameRate sets the frame_rate field.
func (s *AVBufferSrcParameters) SetFrameRate(value *AVRational) {
	s.ptr.frame_rate = value.value
}

// HwFramesCtx gets the hw_frames_ctx field.
func (s *AVBufferSrcParameters) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwFramesCtx sets the hw_frames_ctx field.
func (s *AVBufferSrcParameters) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

// SampleRate gets the sample_rate field.
func (s *AVBufferSrcParameters) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

// SetSampleRate sets the sample_rate field.
func (s *AVBufferSrcParameters) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

// ChannelLayout gets the channel_layout field.
func (s *AVBufferSrcParameters) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

// SetChannelLayout sets the channel_layout field.
func (s *AVBufferSrcParameters) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

// ChLayout gets the ch_layout field.
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
func (s *AVProbeData) BufSize() int {
	value := s.ptr.buf_size
	return int(value)
}

// SetBufSize sets the buf_size field.
func (s *AVProbeData) SetBufSize(value int) {
	s.ptr.buf_size = (C.int)(value)
}

// MimeType gets the mime_type field.
func (s *AVProbeData) MimeType() *CStr {
	value := s.ptr.mime_type
	return wrapCStr(value)
}

// SetMimeType sets the mime_type field.
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
func (s *AVOutputFormat) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

// SetLongName sets the long_name field.
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
func (s *AVOutputFormat) Extensions() *CStr {
	value := s.ptr.extensions
	return wrapCStr(value)
}

// SetExtensions sets the extensions field.
func (s *AVOutputFormat) SetExtensions(value *CStr) {
	s.ptr.extensions = value.ptr
}

// AudioCodec gets the audio_codec field.
func (s *AVOutputFormat) AudioCodec() AVCodecID {
	value := s.ptr.audio_codec
	return AVCodecID(value)
}

// SetAudioCodec sets the audio_codec field.
func (s *AVOutputFormat) SetAudioCodec(value AVCodecID) {
	s.ptr.audio_codec = (C.enum_AVCodecID)(value)
}

// VideoCodec gets the video_codec field.
func (s *AVOutputFormat) VideoCodec() AVCodecID {
	value := s.ptr.video_codec
	return AVCodecID(value)
}

// SetVideoCodec sets the video_codec field.
func (s *AVOutputFormat) SetVideoCodec(value AVCodecID) {
	s.ptr.video_codec = (C.enum_AVCodecID)(value)
}

// SubtitleCodec gets the subtitle_codec field.
func (s *AVOutputFormat) SubtitleCodec() AVCodecID {
	value := s.ptr.subtitle_codec
	return AVCodecID(value)
}

// SetSubtitleCodec sets the subtitle_codec field.
func (s *AVOutputFormat) SetSubtitleCodec(value AVCodecID) {
	s.ptr.subtitle_codec = (C.enum_AVCodecID)(value)
}

// Flags gets the flags field.
func (s *AVOutputFormat) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVOutputFormat) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// CodecTag gets the codec_tag field.
func (s *AVOutputFormat) CodecTag() *Array[*AVCodecTag] {
	value := s.ptr.codec_tag
	return ToAVCodecTagArray(unsafe.Pointer(value))
}

// SetCodecTag sets the codec_tag field.
func (s *AVOutputFormat) SetCodecTag(value *Array[AVCodecTag]) {
	if value != nil {
		s.ptr.codec_tag = (**C.struct_AVCodecTag)(value.ptr)
	} else {
		s.ptr.codec_tag = nil
	}
}

// PrivClass gets the priv_class field.
func (s *AVOutputFormat) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetPrivClass sets the priv_class field.
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
func (s *AVInputFormat) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
func (s *AVInputFormat) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// LongName gets the long_name field.
func (s *AVInputFormat) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}

// SetLongName sets the long_name field.
func (s *AVInputFormat) SetLongName(value *CStr) {
	s.ptr.long_name = value.ptr
}

// Flags gets the flags field.
func (s *AVInputFormat) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVInputFormat) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// Extensions gets the extensions field.
func (s *AVInputFormat) Extensions() *CStr {
	value := s.ptr.extensions
	return wrapCStr(value)
}

// SetExtensions sets the extensions field.
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
func (s *AVInputFormat) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetPrivClass sets the priv_class field.
func (s *AVInputFormat) SetPrivClass(value *AVClass) {
	if value != nil {
		s.ptr.priv_class = value.ptr
	} else {
		s.ptr.priv_class = nil
	}
}

// MimeType gets the mime_type field.
func (s *AVInputFormat) MimeType() *CStr {
	value := s.ptr.mime_type
	return wrapCStr(value)
}

// SetMimeType sets the mime_type field.
func (s *AVInputFormat) SetMimeType(value *CStr) {
	s.ptr.mime_type = value.ptr
}

// RawCodecId gets the raw_codec_id field.
func (s *AVInputFormat) RawCodecId() int {
	value := s.ptr.raw_codec_id
	return int(value)
}

// SetRawCodecId sets the raw_codec_id field.
func (s *AVInputFormat) SetRawCodecId(value int) {
	s.ptr.raw_codec_id = (C.int)(value)
}

// PrivDataSize gets the priv_data_size field.
func (s *AVInputFormat) PrivDataSize() int {
	value := s.ptr.priv_data_size
	return int(value)
}

// SetPrivDataSize sets the priv_data_size field.
func (s *AVInputFormat) SetPrivDataSize(value int) {
	s.ptr.priv_data_size = (C.int)(value)
}

// FlagsInternal gets the flags_internal field.
func (s *AVInputFormat) FlagsInternal() int {
	value := s.ptr.flags_internal
	return int(value)
}

// SetFlagsInternal sets the flags_internal field.
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
func (s *AVIndexEntry) Timestamp() int64 {
	value := s.ptr.timestamp
	return int64(value)
}

// SetTimestamp sets the timestamp field.
func (s *AVIndexEntry) SetTimestamp(value int64) {
	s.ptr.timestamp = (C.int64_t)(value)
}

// flags skipped due to bitfield

// size skipped due to bitfield

// MinDistance gets the min_distance field.
func (s *AVIndexEntry) MinDistance() int {
	value := s.ptr.min_distance
	return int(value)
}

// SetMinDistance sets the min_distance field.
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
func (s *AVStream) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
func (s *AVStream) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Index gets the index field.
func (s *AVStream) Index() int {
	value := s.ptr.index
	return int(value)
}

// SetIndex sets the index field.
func (s *AVStream) SetIndex(value int) {
	s.ptr.index = (C.int)(value)
}

// Id gets the id field.
func (s *AVStream) Id() int {
	value := s.ptr.id
	return int(value)
}

// SetId sets the id field.
func (s *AVStream) SetId(value int) {
	s.ptr.id = (C.int)(value)
}

// Codecpar gets the codecpar field.
func (s *AVStream) Codecpar() *AVCodecParameters {
	value := s.ptr.codecpar
	var valueMapped *AVCodecParameters
	if value != nil {
		valueMapped = &AVCodecParameters{ptr: value}
	}
	return valueMapped
}

// SetCodecpar sets the codecpar field.
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
func (s *AVStream) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
func (s *AVStream) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// StartTime gets the start_time field.
func (s *AVStream) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}

// SetStartTime sets the start_time field.
func (s *AVStream) SetStartTime(value int64) {
	s.ptr.start_time = (C.int64_t)(value)
}

// Duration gets the duration field.
func (s *AVStream) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

// SetDuration sets the duration field.
func (s *AVStream) SetDuration(value int64) {
	s.ptr.duration = (C.int64_t)(value)
}

// NbFrames gets the nb_frames field.
func (s *AVStream) NbFrames() int64 {
	value := s.ptr.nb_frames
	return int64(value)
}

// SetNbFrames sets the nb_frames field.
func (s *AVStream) SetNbFrames(value int64) {
	s.ptr.nb_frames = (C.int64_t)(value)
}

// Disposition gets the disposition field.
func (s *AVStream) Disposition() int {
	value := s.ptr.disposition
	return int(value)
}

// SetDisposition sets the disposition field.
func (s *AVStream) SetDisposition(value int) {
	s.ptr.disposition = (C.int)(value)
}

// Discard gets the discard field.
func (s *AVStream) Discard() AVDiscard {
	value := s.ptr.discard
	return AVDiscard(value)
}

// SetDiscard sets the discard field.
func (s *AVStream) SetDiscard(value AVDiscard) {
	s.ptr.discard = (C.enum_AVDiscard)(value)
}

// SampleAspectRatio gets the sample_aspect_ratio field.
func (s *AVStream) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
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
func (s *AVStream) AvgFrameRate() *AVRational {
	value := s.ptr.avg_frame_rate
	return &AVRational{value: value}
}

// SetAvgFrameRate sets the avg_frame_rate field.
func (s *AVStream) SetAvgFrameRate(value *AVRational) {
	s.ptr.avg_frame_rate = value.value
}

// AttachedPic gets the attached_pic field.
func (s *AVStream) AttachedPic() *AVPacket {
	value := &s.ptr.attached_pic
	return &AVPacket{ptr: value}
}

// SideData gets the side_data field.
func (s *AVStream) SideData() *AVPacketSideData {
	value := s.ptr.side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}

// SetSideData sets the side_data field.
func (s *AVStream) SetSideData(value *AVPacketSideData) {
	if value != nil {
		s.ptr.side_data = value.ptr
	} else {
		s.ptr.side_data = nil
	}
}

// NbSideData gets the nb_side_data field.
func (s *AVStream) NbSideData() int {
	value := s.ptr.nb_side_data
	return int(value)
}

// SetNbSideData sets the nb_side_data field.
func (s *AVStream) SetNbSideData(value int) {
	s.ptr.nb_side_data = (C.int)(value)
}

// EventFlags gets the event_flags field.
func (s *AVStream) EventFlags() int {
	value := s.ptr.event_flags
	return int(value)
}

// SetEventFlags sets the event_flags field.
func (s *AVStream) SetEventFlags(value int) {
	s.ptr.event_flags = (C.int)(value)
}

// RFrameRate gets the r_frame_rate field.
func (s *AVStream) RFrameRate() *AVRational {
	value := s.ptr.r_frame_rate
	return &AVRational{value: value}
}

// SetRFrameRate sets the r_frame_rate field.
func (s *AVStream) SetRFrameRate(value *AVRational) {
	s.ptr.r_frame_rate = value.value
}

// PtsWrapBits gets the pts_wrap_bits field.
func (s *AVStream) PtsWrapBits() int {
	value := s.ptr.pts_wrap_bits
	return int(value)
}

// SetPtsWrapBits sets the pts_wrap_bits field.
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
func (s *AVProgram) Discard() AVDiscard {
	value := s.ptr.discard
	return AVDiscard(value)
}

// SetDiscard sets the discard field.
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
func (s *AVProgram) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}

// SetStartTime sets the start_time field.
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
func (s *AVProgram) PtsWrapReference() int64 {
	value := s.ptr.pts_wrap_reference
	return int64(value)
}

// SetPtsWrapReference sets the pts_wrap_reference field.
func (s *AVProgram) SetPtsWrapReference(value int64) {
	s.ptr.pts_wrap_reference = (C.int64_t)(value)
}

// PtsWrapBehavior gets the pts_wrap_behavior field.
func (s *AVProgram) PtsWrapBehavior() int {
	value := s.ptr.pts_wrap_behavior
	return int(value)
}

// SetPtsWrapBehavior sets the pts_wrap_behavior field.
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
func (s *AVChapter) Id() int64 {
	value := s.ptr.id
	return int64(value)
}

// SetId sets the id field.
func (s *AVChapter) SetId(value int64) {
	s.ptr.id = (C.int64_t)(value)
}

// TimeBase gets the time_base field.
func (s *AVChapter) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
func (s *AVChapter) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// Start gets the start field.
func (s *AVChapter) Start() int64 {
	value := s.ptr.start
	return int64(value)
}

// SetStart sets the start field.
func (s *AVChapter) SetStart(value int64) {
	s.ptr.start = (C.int64_t)(value)
}

// End gets the end field.
func (s *AVChapter) End() int64 {
	value := s.ptr.end
	return int64(value)
}

// SetEnd sets the end field.
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
func (s *AVFormatContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
func (s *AVFormatContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Iformat gets the iformat field.
func (s *AVFormatContext) Iformat() *AVInputFormat {
	value := s.ptr.iformat
	var valueMapped *AVInputFormat
	if value != nil {
		valueMapped = &AVInputFormat{ptr: value}
	}
	return valueMapped
}

// SetIformat sets the iformat field.
func (s *AVFormatContext) SetIformat(value *AVInputFormat) {
	if value != nil {
		s.ptr.iformat = value.ptr
	} else {
		s.ptr.iformat = nil
	}
}

// Oformat gets the oformat field.
func (s *AVFormatContext) Oformat() *AVOutputFormat {
	value := s.ptr.oformat
	var valueMapped *AVOutputFormat
	if value != nil {
		valueMapped = &AVOutputFormat{ptr: value}
	}
	return valueMapped
}

// SetOformat sets the oformat field.
func (s *AVFormatContext) SetOformat(value *AVOutputFormat) {
	if value != nil {
		s.ptr.oformat = value.ptr
	} else {
		s.ptr.oformat = nil
	}
}

// PrivData gets the priv_data field.
func (s *AVFormatContext) PrivData() unsafe.Pointer {
	value := s.ptr.priv_data
	return value
}

// SetPrivData sets the priv_data field.
func (s *AVFormatContext) SetPrivData(value unsafe.Pointer) {
	s.ptr.priv_data = value
}

// Pb gets the pb field.
func (s *AVFormatContext) Pb() *AVIOContext {
	value := s.ptr.pb
	var valueMapped *AVIOContext
	if value != nil {
		valueMapped = &AVIOContext{ptr: value}
	}
	return valueMapped
}

// SetPb sets the pb field.
func (s *AVFormatContext) SetPb(value *AVIOContext) {
	if value != nil {
		s.ptr.pb = value.ptr
	} else {
		s.ptr.pb = nil
	}
}

// CtxFlags gets the ctx_flags field.
func (s *AVFormatContext) CtxFlags() int {
	value := s.ptr.ctx_flags
	return int(value)
}

// SetCtxFlags sets the ctx_flags field.
func (s *AVFormatContext) SetCtxFlags(value int) {
	s.ptr.ctx_flags = (C.int)(value)
}

// NbStreams gets the nb_streams field.
func (s *AVFormatContext) NbStreams() uint {
	value := s.ptr.nb_streams
	return uint(value)
}

// SetNbStreams sets the nb_streams field.
func (s *AVFormatContext) SetNbStreams(value uint) {
	s.ptr.nb_streams = (C.uint)(value)
}

// Streams gets the streams field.
func (s *AVFormatContext) Streams() *Array[*AVStream] {
	value := s.ptr.streams
	return ToAVStreamArray(unsafe.Pointer(value))
}

// SetStreams sets the streams field.
func (s *AVFormatContext) SetStreams(value *Array[AVStream]) {
	if value != nil {
		s.ptr.streams = (**C.AVStream)(value.ptr)
	} else {
		s.ptr.streams = nil
	}
}

// Url gets the url field.
func (s *AVFormatContext) Url() *CStr {
	value := s.ptr.url
	return wrapCStr(value)
}

// SetUrl sets the url field.
func (s *AVFormatContext) SetUrl(value *CStr) {
	s.ptr.url = value.ptr
}

// StartTime gets the start_time field.
func (s *AVFormatContext) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}

// SetStartTime sets the start_time field.
func (s *AVFormatContext) SetStartTime(value int64) {
	s.ptr.start_time = (C.int64_t)(value)
}

// Duration gets the duration field.
func (s *AVFormatContext) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

// SetDuration sets the duration field.
func (s *AVFormatContext) SetDuration(value int64) {
	s.ptr.duration = (C.int64_t)(value)
}

// BitRate gets the bit_rate field.
func (s *AVFormatContext) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}

// SetBitRate sets the bit_rate field.
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
func (s *AVFormatContext) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVFormatContext) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// Probesize gets the probesize field.
func (s *AVFormatContext) Probesize() int64 {
	value := s.ptr.probesize
	return int64(value)
}

// SetProbesize sets the probesize field.
func (s *AVFormatContext) SetProbesize(value int64) {
	s.ptr.probesize = (C.int64_t)(value)
}

// MaxAnalyzeDuration gets the max_analyze_duration field.
func (s *AVFormatContext) MaxAnalyzeDuration() int64 {
	value := s.ptr.max_analyze_duration
	return int64(value)
}

// SetMaxAnalyzeDuration sets the max_analyze_duration field.
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
func (s *AVFormatContext) VideoCodecId() AVCodecID {
	value := s.ptr.video_codec_id
	return AVCodecID(value)
}

// SetVideoCodecId sets the video_codec_id field.
func (s *AVFormatContext) SetVideoCodecId(value AVCodecID) {
	s.ptr.video_codec_id = (C.enum_AVCodecID)(value)
}

// AudioCodecId gets the audio_codec_id field.
func (s *AVFormatContext) AudioCodecId() AVCodecID {
	value := s.ptr.audio_codec_id
	return AVCodecID(value)
}

// SetAudioCodecId sets the audio_codec_id field.
func (s *AVFormatContext) SetAudioCodecId(value AVCodecID) {
	s.ptr.audio_codec_id = (C.enum_AVCodecID)(value)
}

// SubtitleCodecId gets the subtitle_codec_id field.
func (s *AVFormatContext) SubtitleCodecId() AVCodecID {
	value := s.ptr.subtitle_codec_id
	return AVCodecID(value)
}

// SetSubtitleCodecId sets the subtitle_codec_id field.
func (s *AVFormatContext) SetSubtitleCodecId(value AVCodecID) {
	s.ptr.subtitle_codec_id = (C.enum_AVCodecID)(value)
}

// MaxIndexSize gets the max_index_size field.
func (s *AVFormatContext) MaxIndexSize() uint {
	value := s.ptr.max_index_size
	return uint(value)
}

// SetMaxIndexSize sets the max_index_size field.
func (s *AVFormatContext) SetMaxIndexSize(value uint) {
	s.ptr.max_index_size = (C.uint)(value)
}

// MaxPictureBuffer gets the max_picture_buffer field.
func (s *AVFormatContext) MaxPictureBuffer() uint {
	value := s.ptr.max_picture_buffer
	return uint(value)
}

// SetMaxPictureBuffer sets the max_picture_buffer field.
func (s *AVFormatContext) SetMaxPictureBuffer(value uint) {
	s.ptr.max_picture_buffer = (C.uint)(value)
}

// NbChapters gets the nb_chapters field.
func (s *AVFormatContext) NbChapters() uint {
	value := s.ptr.nb_chapters
	return uint(value)
}

// SetNbChapters sets the nb_chapters field.
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
func (s *AVFormatContext) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

// SetMetadata sets the metadata field.
func (s *AVFormatContext) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

// StartTimeRealtime gets the start_time_realtime field.
func (s *AVFormatContext) StartTimeRealtime() int64 {
	value := s.ptr.start_time_realtime
	return int64(value)
}

// SetStartTimeRealtime sets the start_time_realtime field.
func (s *AVFormatContext) SetStartTimeRealtime(value int64) {
	s.ptr.start_time_realtime = (C.int64_t)(value)
}

// FpsProbeSize gets the fps_probe_size field.
func (s *AVFormatContext) FpsProbeSize() int {
	value := s.ptr.fps_probe_size
	return int(value)
}

// SetFpsProbeSize sets the fps_probe_size field.
func (s *AVFormatContext) SetFpsProbeSize(value int) {
	s.ptr.fps_probe_size = (C.int)(value)
}

// ErrorRecognition gets the error_recognition field.
func (s *AVFormatContext) ErrorRecognition() int {
	value := s.ptr.error_recognition
	return int(value)
}

// SetErrorRecognition sets the error_recognition field.
func (s *AVFormatContext) SetErrorRecognition(value int) {
	s.ptr.error_recognition = (C.int)(value)
}

// InterruptCallback gets the interrupt_callback field.
func (s *AVFormatContext) InterruptCallback() *AVIOInterruptCB {
	value := &s.ptr.interrupt_callback
	return &AVIOInterruptCB{ptr: value}
}

// Debug gets the debug field.
func (s *AVFormatContext) Debug() int {
	value := s.ptr.debug
	return int(value)
}

// SetDebug sets the debug field.
func (s *AVFormatContext) SetDebug(value int) {
	s.ptr.debug = (C.int)(value)
}

// MaxInterleaveDelta gets the max_interleave_delta field.
func (s *AVFormatContext) MaxInterleaveDelta() int64 {
	value := s.ptr.max_interleave_delta
	return int64(value)
}

// SetMaxInterleaveDelta sets the max_interleave_delta field.
func (s *AVFormatContext) SetMaxInterleaveDelta(value int64) {
	s.ptr.max_interleave_delta = (C.int64_t)(value)
}

// StrictStdCompliance gets the strict_std_compliance field.
func (s *AVFormatContext) StrictStdCompliance() int {
	value := s.ptr.strict_std_compliance
	return int(value)
}

// SetStrictStdCompliance sets the strict_std_compliance field.
func (s *AVFormatContext) SetStrictStdCompliance(value int) {
	s.ptr.strict_std_compliance = (C.int)(value)
}

// EventFlags gets the event_flags field.
func (s *AVFormatContext) EventFlags() int {
	value := s.ptr.event_flags
	return int(value)
}

// SetEventFlags sets the event_flags field.
func (s *AVFormatContext) SetEventFlags(value int) {
	s.ptr.event_flags = (C.int)(value)
}

// MaxTsProbe gets the max_ts_probe field.
func (s *AVFormatContext) MaxTsProbe() int {
	value := s.ptr.max_ts_probe
	return int(value)
}

// SetMaxTsProbe sets the max_ts_probe field.
func (s *AVFormatContext) SetMaxTsProbe(value int) {
	s.ptr.max_ts_probe = (C.int)(value)
}

// AvoidNegativeTs gets the avoid_negative_ts field.
func (s *AVFormatContext) AvoidNegativeTs() int {
	value := s.ptr.avoid_negative_ts
	return int(value)
}

// SetAvoidNegativeTs sets the avoid_negative_ts field.
func (s *AVFormatContext) SetAvoidNegativeTs(value int) {
	s.ptr.avoid_negative_ts = (C.int)(value)
}

// TsId gets the ts_id field.
func (s *AVFormatContext) TsId() int {
	value := s.ptr.ts_id
	return int(value)
}

// SetTsId sets the ts_id field.
func (s *AVFormatContext) SetTsId(value int) {
	s.ptr.ts_id = (C.int)(value)
}

// AudioPreload gets the audio_preload field.
func (s *AVFormatContext) AudioPreload() int {
	value := s.ptr.audio_preload
	return int(value)
}

// SetAudioPreload sets the audio_preload field.
func (s *AVFormatContext) SetAudioPreload(value int) {
	s.ptr.audio_preload = (C.int)(value)
}

// MaxChunkDuration gets the max_chunk_duration field.
func (s *AVFormatContext) MaxChunkDuration() int {
	value := s.ptr.max_chunk_duration
	return int(value)
}

// SetMaxChunkDuration sets the max_chunk_duration field.
func (s *AVFormatContext) SetMaxChunkDuration(value int) {
	s.ptr.max_chunk_duration = (C.int)(value)
}

// MaxChunkSize gets the max_chunk_size field.
func (s *AVFormatContext) MaxChunkSize() int {
	value := s.ptr.max_chunk_size
	return int(value)
}

// SetMaxChunkSize sets the max_chunk_size field.
func (s *AVFormatContext) SetMaxChunkSize(value int) {
	s.ptr.max_chunk_size = (C.int)(value)
}

// UseWallclockAsTimestamps gets the use_wallclock_as_timestamps field.
func (s *AVFormatContext) UseWallclockAsTimestamps() int {
	value := s.ptr.use_wallclock_as_timestamps
	return int(value)
}

// SetUseWallclockAsTimestamps sets the use_wallclock_as_timestamps field.
func (s *AVFormatContext) SetUseWallclockAsTimestamps(value int) {
	s.ptr.use_wallclock_as_timestamps = (C.int)(value)
}

// AvioFlags gets the avio_flags field.
func (s *AVFormatContext) AvioFlags() int {
	value := s.ptr.avio_flags
	return int(value)
}

// SetAvioFlags sets the avio_flags field.
func (s *AVFormatContext) SetAvioFlags(value int) {
	s.ptr.avio_flags = (C.int)(value)
}

// DurationEstimationMethod gets the duration_estimation_method field.
func (s *AVFormatContext) DurationEstimationMethod() AVDurationEstimationMethod {
	value := s.ptr.duration_estimation_method
	return AVDurationEstimationMethod(value)
}

// SetDurationEstimationMethod sets the duration_estimation_method field.
func (s *AVFormatContext) SetDurationEstimationMethod(value AVDurationEstimationMethod) {
	s.ptr.duration_estimation_method = (C.enum_AVDurationEstimationMethod)(value)
}

// SkipInitialBytes gets the skip_initial_bytes field.
func (s *AVFormatContext) SkipInitialBytes() int64 {
	value := s.ptr.skip_initial_bytes
	return int64(value)
}

// SetSkipInitialBytes sets the skip_initial_bytes field.
func (s *AVFormatContext) SetSkipInitialBytes(value int64) {
	s.ptr.skip_initial_bytes = (C.int64_t)(value)
}

// CorrectTsOverflow gets the correct_ts_overflow field.
func (s *AVFormatContext) CorrectTsOverflow() uint {
	value := s.ptr.correct_ts_overflow
	return uint(value)
}

// SetCorrectTsOverflow sets the correct_ts_overflow field.
func (s *AVFormatContext) SetCorrectTsOverflow(value uint) {
	s.ptr.correct_ts_overflow = (C.uint)(value)
}

// Seek2Any gets the seek2any field.
func (s *AVFormatContext) Seek2Any() int {
	value := s.ptr.seek2any
	return int(value)
}

// SetSeek2Any sets the seek2any field.
func (s *AVFormatContext) SetSeek2Any(value int) {
	s.ptr.seek2any = (C.int)(value)
}

// FlushPackets gets the flush_packets field.
func (s *AVFormatContext) FlushPackets() int {
	value := s.ptr.flush_packets
	return int(value)
}

// SetFlushPackets sets the flush_packets field.
func (s *AVFormatContext) SetFlushPackets(value int) {
	s.ptr.flush_packets = (C.int)(value)
}

// ProbeScore gets the probe_score field.
func (s *AVFormatContext) ProbeScore() int {
	value := s.ptr.probe_score
	return int(value)
}

// SetProbeScore sets the probe_score field.
func (s *AVFormatContext) SetProbeScore(value int) {
	s.ptr.probe_score = (C.int)(value)
}

// FormatProbesize gets the format_probesize field.
func (s *AVFormatContext) FormatProbesize() int {
	value := s.ptr.format_probesize
	return int(value)
}

// SetFormatProbesize sets the format_probesize field.
func (s *AVFormatContext) SetFormatProbesize(value int) {
	s.ptr.format_probesize = (C.int)(value)
}

// CodecWhitelist gets the codec_whitelist field.
func (s *AVFormatContext) CodecWhitelist() *CStr {
	value := s.ptr.codec_whitelist
	return wrapCStr(value)
}

// SetCodecWhitelist sets the codec_whitelist field.
func (s *AVFormatContext) SetCodecWhitelist(value *CStr) {
	s.ptr.codec_whitelist = value.ptr
}

// FormatWhitelist gets the format_whitelist field.
func (s *AVFormatContext) FormatWhitelist() *CStr {
	value := s.ptr.format_whitelist
	return wrapCStr(value)
}

// SetFormatWhitelist sets the format_whitelist field.
func (s *AVFormatContext) SetFormatWhitelist(value *CStr) {
	s.ptr.format_whitelist = value.ptr
}

// IoRepositioned gets the io_repositioned field.
func (s *AVFormatContext) IoRepositioned() int {
	value := s.ptr.io_repositioned
	return int(value)
}

// SetIoRepositioned sets the io_repositioned field.
func (s *AVFormatContext) SetIoRepositioned(value int) {
	s.ptr.io_repositioned = (C.int)(value)
}

// VideoCodec gets the video_codec field.
func (s *AVFormatContext) VideoCodec() *AVCodec {
	value := s.ptr.video_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

// SetVideoCodec sets the video_codec field.
func (s *AVFormatContext) SetVideoCodec(value *AVCodec) {
	if value != nil {
		s.ptr.video_codec = value.ptr
	} else {
		s.ptr.video_codec = nil
	}
}

// AudioCodec gets the audio_codec field.
func (s *AVFormatContext) AudioCodec() *AVCodec {
	value := s.ptr.audio_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

// SetAudioCodec sets the audio_codec field.
func (s *AVFormatContext) SetAudioCodec(value *AVCodec) {
	if value != nil {
		s.ptr.audio_codec = value.ptr
	} else {
		s.ptr.audio_codec = nil
	}
}

// SubtitleCodec gets the subtitle_codec field.
func (s *AVFormatContext) SubtitleCodec() *AVCodec {
	value := s.ptr.subtitle_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

// SetSubtitleCodec sets the subtitle_codec field.
func (s *AVFormatContext) SetSubtitleCodec(value *AVCodec) {
	if value != nil {
		s.ptr.subtitle_codec = value.ptr
	} else {
		s.ptr.subtitle_codec = nil
	}
}

// DataCodec gets the data_codec field.
func (s *AVFormatContext) DataCodec() *AVCodec {
	value := s.ptr.data_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}

// SetDataCodec sets the data_codec field.
func (s *AVFormatContext) SetDataCodec(value *AVCodec) {
	if value != nil {
		s.ptr.data_codec = value.ptr
	} else {
		s.ptr.data_codec = nil
	}
}

// MetadataHeaderPadding gets the metadata_header_padding field.
func (s *AVFormatContext) MetadataHeaderPadding() int {
	value := s.ptr.metadata_header_padding
	return int(value)
}

// SetMetadataHeaderPadding sets the metadata_header_padding field.
func (s *AVFormatContext) SetMetadataHeaderPadding(value int) {
	s.ptr.metadata_header_padding = (C.int)(value)
}

// Opaque gets the opaque field.
func (s *AVFormatContext) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
func (s *AVFormatContext) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// control_message_cb skipped due to ident callback

// OutputTsOffset gets the output_ts_offset field.
func (s *AVFormatContext) OutputTsOffset() int64 {
	value := s.ptr.output_ts_offset
	return int64(value)
}

// SetOutputTsOffset sets the output_ts_offset field.
func (s *AVFormatContext) SetOutputTsOffset(value int64) {
	s.ptr.output_ts_offset = (C.int64_t)(value)
}

// DumpSeparator gets the dump_separator field.
func (s *AVFormatContext) DumpSeparator() unsafe.Pointer {
	value := s.ptr.dump_separator
	return unsafe.Pointer(value)
}

// SetDumpSeparator sets the dump_separator field.
func (s *AVFormatContext) SetDumpSeparator(value unsafe.Pointer) {
	s.ptr.dump_separator = (*C.uint8_t)(value)
}

// DataCodecId gets the data_codec_id field.
func (s *AVFormatContext) DataCodecId() AVCodecID {
	value := s.ptr.data_codec_id
	return AVCodecID(value)
}

// SetDataCodecId sets the data_codec_id field.
func (s *AVFormatContext) SetDataCodecId(value AVCodecID) {
	s.ptr.data_codec_id = (C.enum_AVCodecID)(value)
}

// ProtocolWhitelist gets the protocol_whitelist field.
func (s *AVFormatContext) ProtocolWhitelist() *CStr {
	value := s.ptr.protocol_whitelist
	return wrapCStr(value)
}

// SetProtocolWhitelist sets the protocol_whitelist field.
func (s *AVFormatContext) SetProtocolWhitelist(value *CStr) {
	s.ptr.protocol_whitelist = value.ptr
}

// io_open skipped due to func ptr

// io_close skipped due to func ptr

// ProtocolBlacklist gets the protocol_blacklist field.
func (s *AVFormatContext) ProtocolBlacklist() *CStr {
	value := s.ptr.protocol_blacklist
	return wrapCStr(value)
}

// SetProtocolBlacklist sets the protocol_blacklist field.
func (s *AVFormatContext) SetProtocolBlacklist(value *CStr) {
	s.ptr.protocol_blacklist = value.ptr
}

// MaxStreams gets the max_streams field.
func (s *AVFormatContext) MaxStreams() int {
	value := s.ptr.max_streams
	return int(value)
}

// SetMaxStreams sets the max_streams field.
func (s *AVFormatContext) SetMaxStreams(value int) {
	s.ptr.max_streams = (C.int)(value)
}

// SkipEstimateDurationFromPts gets the skip_estimate_duration_from_pts field.
func (s *AVFormatContext) SkipEstimateDurationFromPts() int {
	value := s.ptr.skip_estimate_duration_from_pts
	return int(value)
}

// SetSkipEstimateDurationFromPts sets the skip_estimate_duration_from_pts field.
func (s *AVFormatContext) SetSkipEstimateDurationFromPts(value int) {
	s.ptr.skip_estimate_duration_from_pts = (C.int)(value)
}

// MaxProbePackets gets the max_probe_packets field.
func (s *AVFormatContext) MaxProbePackets() int {
	value := s.ptr.max_probe_packets
	return int(value)
}

// SetMaxProbePackets sets the max_probe_packets field.
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
func (s *AVIODirEntry) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// SetName sets the name field.
func (s *AVIODirEntry) SetName(value *CStr) {
	s.ptr.name = value.ptr
}

// Type gets the type field.
func (s *AVIODirEntry) Type() int {
	value := s.ptr._type
	return int(value)
}

// SetType sets the type field.
func (s *AVIODirEntry) SetType(value int) {
	s.ptr._type = (C.int)(value)
}

// Utf8 gets the utf8 field.
func (s *AVIODirEntry) Utf8() int {
	value := s.ptr.utf8
	return int(value)
}

// SetUtf8 sets the utf8 field.
func (s *AVIODirEntry) SetUtf8(value int) {
	s.ptr.utf8 = (C.int)(value)
}

// Size gets the size field.
func (s *AVIODirEntry) Size() int64 {
	value := s.ptr.size
	return int64(value)
}

// SetSize sets the size field.
func (s *AVIODirEntry) SetSize(value int64) {
	s.ptr.size = (C.int64_t)(value)
}

// ModificationTimestamp gets the modification_timestamp field.
func (s *AVIODirEntry) ModificationTimestamp() int64 {
	value := s.ptr.modification_timestamp
	return int64(value)
}

// SetModificationTimestamp sets the modification_timestamp field.
func (s *AVIODirEntry) SetModificationTimestamp(value int64) {
	s.ptr.modification_timestamp = (C.int64_t)(value)
}

// AccessTimestamp gets the access_timestamp field.
func (s *AVIODirEntry) AccessTimestamp() int64 {
	value := s.ptr.access_timestamp
	return int64(value)
}

// SetAccessTimestamp sets the access_timestamp field.
func (s *AVIODirEntry) SetAccessTimestamp(value int64) {
	s.ptr.access_timestamp = (C.int64_t)(value)
}

// StatusChangeTimestamp gets the status_change_timestamp field.
func (s *AVIODirEntry) StatusChangeTimestamp() int64 {
	value := s.ptr.status_change_timestamp
	return int64(value)
}

// SetStatusChangeTimestamp sets the status_change_timestamp field.
func (s *AVIODirEntry) SetStatusChangeTimestamp(value int64) {
	s.ptr.status_change_timestamp = (C.int64_t)(value)
}

// UserId gets the user_id field.
func (s *AVIODirEntry) UserId() int64 {
	value := s.ptr.user_id
	return int64(value)
}

// SetUserId sets the user_id field.
func (s *AVIODirEntry) SetUserId(value int64) {
	s.ptr.user_id = (C.int64_t)(value)
}

// GroupId gets the group_id field.
func (s *AVIODirEntry) GroupId() int64 {
	value := s.ptr.group_id
	return int64(value)
}

// SetGroupId sets the group_id field.
func (s *AVIODirEntry) SetGroupId(value int64) {
	s.ptr.group_id = (C.int64_t)(value)
}

// Filemode gets the filemode field.
func (s *AVIODirEntry) Filemode() int64 {
	value := s.ptr.filemode
	return int64(value)
}

// SetFilemode sets the filemode field.
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
func (s *AVIOContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
func (s *AVIOContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// buffer skipped due to prim ptr

// BufferSize gets the buffer_size field.
func (s *AVIOContext) BufferSize() int {
	value := s.ptr.buffer_size
	return int(value)
}

// SetBufferSize sets the buffer_size field.
func (s *AVIOContext) SetBufferSize(value int) {
	s.ptr.buffer_size = (C.int)(value)
}

// buf_ptr skipped due to prim ptr

// buf_end skipped due to prim ptr

// Opaque gets the opaque field.
func (s *AVIOContext) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
func (s *AVIOContext) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// read_packet skipped due to func ptr

// write_packet skipped due to func ptr

// seek skipped due to func ptr

// Pos gets the pos field.
func (s *AVIOContext) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}

// SetPos sets the pos field.
func (s *AVIOContext) SetPos(value int64) {
	s.ptr.pos = (C.int64_t)(value)
}

// EofReached gets the eof_reached field.
func (s *AVIOContext) EofReached() int {
	value := s.ptr.eof_reached
	return int(value)
}

// SetEofReached sets the eof_reached field.
func (s *AVIOContext) SetEofReached(value int) {
	s.ptr.eof_reached = (C.int)(value)
}

// Error gets the error field.
func (s *AVIOContext) Error() int {
	value := s.ptr.error
	return int(value)
}

// SetError sets the error field.
func (s *AVIOContext) SetError(value int) {
	s.ptr.error = (C.int)(value)
}

// WriteFlag gets the write_flag field.
func (s *AVIOContext) WriteFlag() int {
	value := s.ptr.write_flag
	return int(value)
}

// SetWriteFlag sets the write_flag field.
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
func (s *AVIOContext) MinPacketSize() int {
	value := s.ptr.min_packet_size
	return int(value)
}

// SetMinPacketSize sets the min_packet_size field.
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
func (s *AVIOContext) Seekable() int {
	value := s.ptr.seekable
	return int(value)
}

// SetSeekable sets the seekable field.
func (s *AVIOContext) SetSeekable(value int) {
	s.ptr.seekable = (C.int)(value)
}

// Direct gets the direct field.
func (s *AVIOContext) Direct() int {
	value := s.ptr.direct
	return int(value)
}

// SetDirect sets the direct field.
func (s *AVIOContext) SetDirect(value int) {
	s.ptr.direct = (C.int)(value)
}

// ProtocolWhitelist gets the protocol_whitelist field.
func (s *AVIOContext) ProtocolWhitelist() *CStr {
	value := s.ptr.protocol_whitelist
	return wrapCStr(value)
}

// SetProtocolWhitelist sets the protocol_whitelist field.
func (s *AVIOContext) SetProtocolWhitelist(value *CStr) {
	s.ptr.protocol_whitelist = value.ptr
}

// ProtocolBlacklist gets the protocol_blacklist field.
func (s *AVIOContext) ProtocolBlacklist() *CStr {
	value := s.ptr.protocol_blacklist
	return wrapCStr(value)
}

// SetProtocolBlacklist sets the protocol_blacklist field.
func (s *AVIOContext) SetProtocolBlacklist(value *CStr) {
	s.ptr.protocol_blacklist = value.ptr
}

// write_data_type skipped due to func ptr

// IgnoreBoundaryPoint gets the ignore_boundary_point field.
func (s *AVIOContext) IgnoreBoundaryPoint() int {
	value := s.ptr.ignore_boundary_point
	return int(value)
}

// SetIgnoreBoundaryPoint sets the ignore_boundary_point field.
func (s *AVIOContext) SetIgnoreBoundaryPoint(value int) {
	s.ptr.ignore_boundary_point = (C.int)(value)
}

// buf_ptr_max skipped due to prim ptr

// BytesRead gets the bytes_read field.
func (s *AVIOContext) BytesRead() int64 {
	value := s.ptr.bytes_read
	return int64(value)
}

// SetBytesRead sets the bytes_read field.
func (s *AVIOContext) SetBytesRead(value int64) {
	s.ptr.bytes_read = (C.int64_t)(value)
}

// BytesWritten gets the bytes_written field.
func (s *AVIOContext) BytesWritten() int64 {
	value := s.ptr.bytes_written
	return int64(value)
}

// SetBytesWritten sets the bytes_written field.
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
func (s *AVBufferRef) Data() unsafe.Pointer {
	value := s.ptr.data
	return unsafe.Pointer(value)
}

// SetData sets the data field.
func (s *AVBufferRef) SetData(value unsafe.Pointer) {
	s.ptr.data = (*C.uint8_t)(value)
}

// Size gets the size field.
func (s *AVBufferRef) Size() uint64 {
	value := s.ptr.size
	return uint64(value)
}

// SetSize sets the size field.
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
func (s *AVChannelLayout) Order() AVChannelOrder {
	value := s.ptr.order
	return AVChannelOrder(value)
}

// SetOrder sets the order field.
func (s *AVChannelLayout) SetOrder(value AVChannelOrder) {
	s.ptr.order = (C.enum_AVChannelOrder)(value)
}

// NbChannels gets the nb_channels field.
func (s *AVChannelLayout) NbChannels() int {
	value := s.ptr.nb_channels
	return int(value)
}

// SetNbChannels sets the nb_channels field.
func (s *AVChannelLayout) SetNbChannels(value int) {
	s.ptr.nb_channels = (C.int)(value)
}

// u skipped due to union type

// Opaque gets the opaque field.
func (s *AVChannelLayout) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
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
func (s *AVRegionOfInterest) SelfSize() uint32 {
	value := s.ptr.self_size
	return uint32(value)
}

// SetSelfSize sets the self_size field.
func (s *AVRegionOfInterest) SetSelfSize(value uint32) {
	s.ptr.self_size = (C.uint32_t)(value)
}

// Top gets the top field.
func (s *AVRegionOfInterest) Top() int {
	value := s.ptr.top
	return int(value)
}

// SetTop sets the top field.
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
func (s *AVRegionOfInterest) Qoffset() *AVRational {
	value := s.ptr.qoffset
	return &AVRational{value: value}
}

// SetQoffset sets the qoffset field.
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
func (s *AVFrame) Data() *Array[unsafe.Pointer] {
	value := &s.ptr.data
	return ToUint8PtrArray(unsafe.Pointer(value))
}

// Linesize gets the linesize field.
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
func (s *AVFrame) NbSamples() int {
	value := s.ptr.nb_samples
	return int(value)
}

// SetNbSamples sets the nb_samples field.
func (s *AVFrame) SetNbSamples(value int) {
	s.ptr.nb_samples = (C.int)(value)
}

// Format gets the format field.
func (s *AVFrame) Format() int {
	value := s.ptr.format
	return int(value)
}

// SetFormat sets the format field.
func (s *AVFrame) SetFormat(value int) {
	s.ptr.format = (C.int)(value)
}

// KeyFrame gets the key_frame field.
func (s *AVFrame) KeyFrame() int {
	value := s.ptr.key_frame
	return int(value)
}

// SetKeyFrame sets the key_frame field.
func (s *AVFrame) SetKeyFrame(value int) {
	s.ptr.key_frame = (C.int)(value)
}

// PictType gets the pict_type field.
func (s *AVFrame) PictType() AVPictureType {
	value := s.ptr.pict_type
	return AVPictureType(value)
}

// SetPictType sets the pict_type field.
func (s *AVFrame) SetPictType(value AVPictureType) {
	s.ptr.pict_type = (C.enum_AVPictureType)(value)
}

// SampleAspectRatio gets the sample_aspect_ratio field.
func (s *AVFrame) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}

// SetSampleAspectRatio sets the sample_aspect_ratio field.
func (s *AVFrame) SetSampleAspectRatio(value *AVRational) {
	s.ptr.sample_aspect_ratio = value.value
}

// Pts gets the pts field.
func (s *AVFrame) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}

// SetPts sets the pts field.
func (s *AVFrame) SetPts(value int64) {
	s.ptr.pts = (C.int64_t)(value)
}

// PktDts gets the pkt_dts field.
func (s *AVFrame) PktDts() int64 {
	value := s.ptr.pkt_dts
	return int64(value)
}

// SetPktDts sets the pkt_dts field.
func (s *AVFrame) SetPktDts(value int64) {
	s.ptr.pkt_dts = (C.int64_t)(value)
}

// TimeBase gets the time_base field.
func (s *AVFrame) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// SetTimeBase sets the time_base field.
func (s *AVFrame) SetTimeBase(value *AVRational) {
	s.ptr.time_base = value.value
}

// CodedPictureNumber gets the coded_picture_number field.
func (s *AVFrame) CodedPictureNumber() int {
	value := s.ptr.coded_picture_number
	return int(value)
}

// SetCodedPictureNumber sets the coded_picture_number field.
func (s *AVFrame) SetCodedPictureNumber(value int) {
	s.ptr.coded_picture_number = (C.int)(value)
}

// DisplayPictureNumber gets the display_picture_number field.
func (s *AVFrame) DisplayPictureNumber() int {
	value := s.ptr.display_picture_number
	return int(value)
}

// SetDisplayPictureNumber sets the display_picture_number field.
func (s *AVFrame) SetDisplayPictureNumber(value int) {
	s.ptr.display_picture_number = (C.int)(value)
}

// Quality gets the quality field.
func (s *AVFrame) Quality() int {
	value := s.ptr.quality
	return int(value)
}

// SetQuality sets the quality field.
func (s *AVFrame) SetQuality(value int) {
	s.ptr.quality = (C.int)(value)
}

// Opaque gets the opaque field.
func (s *AVFrame) Opaque() unsafe.Pointer {
	value := s.ptr.opaque
	return value
}

// SetOpaque sets the opaque field.
func (s *AVFrame) SetOpaque(value unsafe.Pointer) {
	s.ptr.opaque = value
}

// RepeatPict gets the repeat_pict field.
func (s *AVFrame) RepeatPict() int {
	value := s.ptr.repeat_pict
	return int(value)
}

// SetRepeatPict sets the repeat_pict field.
func (s *AVFrame) SetRepeatPict(value int) {
	s.ptr.repeat_pict = (C.int)(value)
}

// InterlacedFrame gets the interlaced_frame field.
func (s *AVFrame) InterlacedFrame() int {
	value := s.ptr.interlaced_frame
	return int(value)
}

// SetInterlacedFrame sets the interlaced_frame field.
func (s *AVFrame) SetInterlacedFrame(value int) {
	s.ptr.interlaced_frame = (C.int)(value)
}

// TopFieldFirst gets the top_field_first field.
func (s *AVFrame) TopFieldFirst() int {
	value := s.ptr.top_field_first
	return int(value)
}

// SetTopFieldFirst sets the top_field_first field.
func (s *AVFrame) SetTopFieldFirst(value int) {
	s.ptr.top_field_first = (C.int)(value)
}

// PaletteHasChanged gets the palette_has_changed field.
func (s *AVFrame) PaletteHasChanged() int {
	value := s.ptr.palette_has_changed
	return int(value)
}

// SetPaletteHasChanged sets the palette_has_changed field.
func (s *AVFrame) SetPaletteHasChanged(value int) {
	s.ptr.palette_has_changed = (C.int)(value)
}

// ReorderedOpaque gets the reordered_opaque field.
func (s *AVFrame) ReorderedOpaque() int64 {
	value := s.ptr.reordered_opaque
	return int64(value)
}

// SetReorderedOpaque sets the reordered_opaque field.
func (s *AVFrame) SetReorderedOpaque(value int64) {
	s.ptr.reordered_opaque = (C.int64_t)(value)
}

// SampleRate gets the sample_rate field.
func (s *AVFrame) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}

// SetSampleRate sets the sample_rate field.
func (s *AVFrame) SetSampleRate(value int) {
	s.ptr.sample_rate = (C.int)(value)
}

// ChannelLayout gets the channel_layout field.
func (s *AVFrame) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

// SetChannelLayout sets the channel_layout field.
func (s *AVFrame) SetChannelLayout(value uint64) {
	s.ptr.channel_layout = (C.uint64_t)(value)
}

// Buf gets the buf field.
func (s *AVFrame) Buf() *Array[*AVBufferRef] {
	value := &s.ptr.buf
	return ToAVBufferRefArray(unsafe.Pointer(value))
}

// ExtendedBuf gets the extended_buf field.
func (s *AVFrame) ExtendedBuf() *Array[*AVBufferRef] {
	value := s.ptr.extended_buf
	return ToAVBufferRefArray(unsafe.Pointer(value))
}

// SetExtendedBuf sets the extended_buf field.
func (s *AVFrame) SetExtendedBuf(value *Array[AVBufferRef]) {
	if value != nil {
		s.ptr.extended_buf = (**C.AVBufferRef)(value.ptr)
	} else {
		s.ptr.extended_buf = nil
	}
}

// NbExtendedBuf gets the nb_extended_buf field.
func (s *AVFrame) NbExtendedBuf() int {
	value := s.ptr.nb_extended_buf
	return int(value)
}

// SetNbExtendedBuf sets the nb_extended_buf field.
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
func (s *AVFrame) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// SetFlags sets the flags field.
func (s *AVFrame) SetFlags(value int) {
	s.ptr.flags = (C.int)(value)
}

// ColorRange gets the color_range field.
func (s *AVFrame) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}

// SetColorRange sets the color_range field.
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
func (s *AVFrame) Colorspace() AVColorSpace {
	value := s.ptr.colorspace
	return AVColorSpace(value)
}

// SetColorspace sets the colorspace field.
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
func (s *AVFrame) BestEffortTimestamp() int64 {
	value := s.ptr.best_effort_timestamp
	return int64(value)
}

// SetBestEffortTimestamp sets the best_effort_timestamp field.
func (s *AVFrame) SetBestEffortTimestamp(value int64) {
	s.ptr.best_effort_timestamp = (C.int64_t)(value)
}

// PktPos gets the pkt_pos field.
func (s *AVFrame) PktPos() int64 {
	value := s.ptr.pkt_pos
	return int64(value)
}

// SetPktPos sets the pkt_pos field.
func (s *AVFrame) SetPktPos(value int64) {
	s.ptr.pkt_pos = (C.int64_t)(value)
}

// PktDuration gets the pkt_duration field.
func (s *AVFrame) PktDuration() int64 {
	value := s.ptr.pkt_duration
	return int64(value)
}

// SetPktDuration sets the pkt_duration field.
func (s *AVFrame) SetPktDuration(value int64) {
	s.ptr.pkt_duration = (C.int64_t)(value)
}

// Metadata gets the metadata field.
func (s *AVFrame) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

// SetMetadata sets the metadata field.
func (s *AVFrame) SetMetadata(value *AVDictionary) {
	if value != nil {
		s.ptr.metadata = value.ptr
	} else {
		s.ptr.metadata = nil
	}
}

// DecodeErrorFlags gets the decode_error_flags field.
func (s *AVFrame) DecodeErrorFlags() int {
	value := s.ptr.decode_error_flags
	return int(value)
}

// SetDecodeErrorFlags sets the decode_error_flags field.
func (s *AVFrame) SetDecodeErrorFlags(value int) {
	s.ptr.decode_error_flags = (C.int)(value)
}

// Channels gets the channels field.
func (s *AVFrame) Channels() int {
	value := s.ptr.channels
	return int(value)
}

// SetChannels sets the channels field.
func (s *AVFrame) SetChannels(value int) {
	s.ptr.channels = (C.int)(value)
}

// PktSize gets the pkt_size field.
func (s *AVFrame) PktSize() int {
	value := s.ptr.pkt_size
	return int(value)
}

// SetPktSize sets the pkt_size field.
func (s *AVFrame) SetPktSize(value int) {
	s.ptr.pkt_size = (C.int)(value)
}

// HwFramesCtx gets the hw_frames_ctx field.
func (s *AVFrame) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetHwFramesCtx sets the hw_frames_ctx field.
func (s *AVFrame) SetHwFramesCtx(value *AVBufferRef) {
	if value != nil {
		s.ptr.hw_frames_ctx = value.ptr
	} else {
		s.ptr.hw_frames_ctx = nil
	}
}

// OpaqueRef gets the opaque_ref field.
func (s *AVFrame) OpaqueRef() *AVBufferRef {
	value := s.ptr.opaque_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetOpaqueRef sets the opaque_ref field.
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
func (s *AVFrame) PrivateRef() *AVBufferRef {
	value := s.ptr.private_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetPrivateRef sets the private_ref field.
func (s *AVFrame) SetPrivateRef(value *AVBufferRef) {
	if value != nil {
		s.ptr.private_ref = value.ptr
	} else {
		s.ptr.private_ref = nil
	}
}

// ChLayout gets the ch_layout field.
func (s *AVFrame) ChLayout() *AVChannelLayout {
	value := &s.ptr.ch_layout
	return &AVChannelLayout{ptr: value}
}

// Duration gets the duration field.
func (s *AVFrame) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}

// SetDuration sets the duration field.
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
func (s *AVHWDeviceContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
func (s *AVHWDeviceContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Internal gets the internal field.
func (s *AVHWDeviceContext) Internal() *AVHWDeviceInternal {
	value := s.ptr.internal
	var valueMapped *AVHWDeviceInternal
	if value != nil {
		valueMapped = &AVHWDeviceInternal{ptr: value}
	}
	return valueMapped
}

// SetInternal sets the internal field.
func (s *AVHWDeviceContext) SetInternal(value *AVHWDeviceInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

// Type gets the type field.
func (s *AVHWDeviceContext) Type() AVHWDeviceType {
	value := s.ptr._type
	return AVHWDeviceType(value)
}

// SetType sets the type field.
func (s *AVHWDeviceContext) SetType(value AVHWDeviceType) {
	s.ptr._type = (C.enum_AVHWDeviceType)(value)
}

// Hwctx gets the hwctx field.
func (s *AVHWDeviceContext) Hwctx() unsafe.Pointer {
	value := s.ptr.hwctx
	return value
}

// SetHwctx sets the hwctx field.
func (s *AVHWDeviceContext) SetHwctx(value unsafe.Pointer) {
	s.ptr.hwctx = value
}

// free skipped due to func ptr

// UserOpaque gets the user_opaque field.
func (s *AVHWDeviceContext) UserOpaque() unsafe.Pointer {
	value := s.ptr.user_opaque
	return value
}

// SetUserOpaque sets the user_opaque field.
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
func (s *AVHWFramesContext) AvClass() *AVClass {
	value := s.ptr.av_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}

// SetAvClass sets the av_class field.
func (s *AVHWFramesContext) SetAvClass(value *AVClass) {
	if value != nil {
		s.ptr.av_class = value.ptr
	} else {
		s.ptr.av_class = nil
	}
}

// Internal gets the internal field.
func (s *AVHWFramesContext) Internal() *AVHWFramesInternal {
	value := s.ptr.internal
	var valueMapped *AVHWFramesInternal
	if value != nil {
		valueMapped = &AVHWFramesInternal{ptr: value}
	}
	return valueMapped
}

// SetInternal sets the internal field.
func (s *AVHWFramesContext) SetInternal(value *AVHWFramesInternal) {
	if value != nil {
		s.ptr.internal = value.ptr
	} else {
		s.ptr.internal = nil
	}
}

// DeviceRef gets the device_ref field.
func (s *AVHWFramesContext) DeviceRef() *AVBufferRef {
	value := s.ptr.device_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// SetDeviceRef sets the device_ref field.
func (s *AVHWFramesContext) SetDeviceRef(value *AVBufferRef) {
	if value != nil {
		s.ptr.device_ref = value.ptr
	} else {
		s.ptr.device_ref = nil
	}
}

// DeviceCtx gets the device_ctx field.
func (s *AVHWFramesContext) DeviceCtx() *AVHWDeviceContext {
	value := s.ptr.device_ctx
	var valueMapped *AVHWDeviceContext
	if value != nil {
		valueMapped = &AVHWDeviceContext{ptr: value}
	}
	return valueMapped
}

// SetDeviceCtx sets the device_ctx field.
func (s *AVHWFramesContext) SetDeviceCtx(value *AVHWDeviceContext) {
	if value != nil {
		s.ptr.device_ctx = value.ptr
	} else {
		s.ptr.device_ctx = nil
	}
}

// Hwctx gets the hwctx field.
func (s *AVHWFramesContext) Hwctx() unsafe.Pointer {
	value := s.ptr.hwctx
	return value
}

// SetHwctx sets the hwctx field.
func (s *AVHWFramesContext) SetHwctx(value unsafe.Pointer) {
	s.ptr.hwctx = value
}

// free skipped due to func ptr

// UserOpaque gets the user_opaque field.
func (s *AVHWFramesContext) UserOpaque() unsafe.Pointer {
	value := s.ptr.user_opaque
	return value
}

// SetUserOpaque sets the user_opaque field.
func (s *AVHWFramesContext) SetUserOpaque(value unsafe.Pointer) {
	s.ptr.user_opaque = value
}

// Pool gets the pool field.
func (s *AVHWFramesContext) Pool() *AVBufferPool {
	value := s.ptr.pool
	var valueMapped *AVBufferPool
	if value != nil {
		valueMapped = &AVBufferPool{ptr: value}
	}
	return valueMapped
}

// SetPool sets the pool field.
func (s *AVHWFramesContext) SetPool(value *AVBufferPool) {
	if value != nil {
		s.ptr.pool = value.ptr
	} else {
		s.ptr.pool = nil
	}
}

// InitialPoolSize gets the initial_pool_size field.
func (s *AVHWFramesContext) InitialPoolSize() int {
	value := s.ptr.initial_pool_size
	return int(value)
}

// SetInitialPoolSize sets the initial_pool_size field.
func (s *AVHWFramesContext) SetInitialPoolSize(value int) {
	s.ptr.initial_pool_size = (C.int)(value)
}

// Format gets the format field.
func (s *AVHWFramesContext) Format() AVPixelFormat {
	value := s.ptr.format
	return AVPixelFormat(value)
}

// SetFormat sets the format field.
func (s *AVHWFramesContext) SetFormat(value AVPixelFormat) {
	s.ptr.format = (C.enum_AVPixelFormat)(value)
}

// SwFormat gets the sw_format field.
func (s *AVHWFramesContext) SwFormat() AVPixelFormat {
	value := s.ptr.sw_format
	return AVPixelFormat(value)
}

// SetSwFormat sets the sw_format field.
func (s *AVHWFramesContext) SetSwFormat(value AVPixelFormat) {
	s.ptr.sw_format = (C.enum_AVPixelFormat)(value)
}

// Width gets the width field.
func (s *AVHWFramesContext) Width() int {
	value := s.ptr.width
	return int(value)
}

// SetWidth sets the width field.
func (s *AVHWFramesContext) SetWidth(value int) {
	s.ptr.width = (C.int)(value)
}

// Height gets the height field.
func (s *AVHWFramesContext) Height() int {
	value := s.ptr.height
	return int(value)
}

// SetHeight sets the height field.
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
func (s *AVHWFramesConstraints) ValidHwFormats() *Array[AVPixelFormat] {
	value := s.ptr.valid_hw_formats
	return ToAVPixelFormatArray(unsafe.Pointer(value))
}

// SetValidHwFormats sets the valid_hw_formats field.
func (s *AVHWFramesConstraints) SetValidHwFormats(value *Array[AVPixelFormat]) {
	if value != nil {
		s.ptr.valid_hw_formats = (*C.enum_AVPixelFormat)(value.ptr)
	} else {
		s.ptr.valid_hw_formats = nil
	}
}

// ValidSwFormats gets the valid_sw_formats field.
func (s *AVHWFramesConstraints) ValidSwFormats() *Array[AVPixelFormat] {
	value := s.ptr.valid_sw_formats
	return ToAVPixelFormatArray(unsafe.Pointer(value))
}

// SetValidSwFormats sets the valid_sw_formats field.
func (s *AVHWFramesConstraints) SetValidSwFormats(value *Array[AVPixelFormat]) {
	if value != nil {
		s.ptr.valid_sw_formats = (*C.enum_AVPixelFormat)(value.ptr)
	} else {
		s.ptr.valid_sw_formats = nil
	}
}

// MinWidth gets the min_width field.
func (s *AVHWFramesConstraints) MinWidth() int {
	value := s.ptr.min_width
	return int(value)
}

// SetMinWidth sets the min_width field.
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
func (s *AVHWFramesConstraints) MaxWidth() int {
	value := s.ptr.max_width
	return int(value)
}

// SetMaxWidth sets the max_width field.
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
func (s *AVClass) ClassName() *CStr {
	value := s.ptr.class_name
	return wrapCStr(value)
}

// SetClassName sets the class_name field.
func (s *AVClass) SetClassName(value *CStr) {
	s.ptr.class_name = value.ptr
}

// item_name skipped due to func ptr

// Option gets the option field.
func (s *AVClass) Option() *AVOption {
	value := s.ptr.option
	var valueMapped *AVOption
	if value != nil {
		valueMapped = &AVOption{ptr: value}
	}
	return valueMapped
}

// SetOption sets the option field.
func (s *AVClass) SetOption(value *AVOption) {
	if value != nil {
		s.ptr.option = value.ptr
	} else {
		s.ptr.option = nil
	}
}

// Version gets the version field.
func (s *AVClass) Version() int {
	value := s.ptr.version
	return int(value)
}

// SetVersion sets the version field.
func (s *AVClass) SetVersion(value int) {
	s.ptr.version = (C.int)(value)
}

// LogLevelOffsetOffset gets the log_level_offset_offset field.
func (s *AVClass) LogLevelOffsetOffset() int {
	value := s.ptr.log_level_offset_offset
	return int(value)
}

// SetLogLevelOffsetOffset sets the log_level_offset_offset field.
func (s *AVClass) SetLogLevelOffsetOffset(value int) {
	s.ptr.log_level_offset_offset = (C.int)(value)
}

// ParentLogContextOffset gets the parent_log_context_offset field.
func (s *AVClass) ParentLogContextOffset() int {
	value := s.ptr.parent_log_context_offset
	return int(value)
}

// SetParentLogContextOffset sets the parent_log_context_offset field.
func (s *AVClass) SetParentLogContextOffset(value int) {
	s.ptr.parent_log_context_offset = (C.int)(value)
}

// Category gets the category field.
func (s *AVClass) Category() AVClassCategory {
	value := s.ptr.category
	return AVClassCategory(value)
}

// SetCategory sets the category field.
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
func (s *AVOption) Help() *CStr {
	value := s.ptr.help
	return wrapCStr(value)
}

// SetHelp sets the help field.
func (s *AVOption) SetHelp(value *CStr) {
	s.ptr.help = value.ptr
}

// Offset gets the offset field.
func (s *AVOption) Offset() int {
	value := s.ptr.offset
	return int(value)
}

// SetOffset sets the offset field.
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
func (s *AVOption) Min() float64 {
	value := s.ptr.min
	return float64(value)
}

// SetMin sets the min field.
func (s *AVOption) SetMin(value float64) {
	s.ptr.min = (C.double)(value)
}

// Max gets the max field.
func (s *AVOption) Max() float64 {
	value := s.ptr.max
	return float64(value)
}

// SetMax sets the max field.
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
func (s *AVOption) Unit() *CStr {
	value := s.ptr.unit
	return wrapCStr(value)
}

// SetUnit sets the unit field.
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
func (s *AVOptionRange) ValueMin() float64 {
	value := s.ptr.value_min
	return float64(value)
}

// SetValueMin sets the value_min field.
func (s *AVOptionRange) SetValueMin(value float64) {
	s.ptr.value_min = (C.double)(value)
}

// ValueMax gets the value_max field.
func (s *AVOptionRange) ValueMax() float64 {
	value := s.ptr.value_max
	return float64(value)
}

// SetValueMax sets the value_max field.
func (s *AVOptionRange) SetValueMax(value float64) {
	s.ptr.value_max = (C.double)(value)
}

// ComponentMin gets the component_min field.
func (s *AVOptionRange) ComponentMin() float64 {
	value := s.ptr.component_min
	return float64(value)
}

// SetComponentMin sets the component_min field.
func (s *AVOptionRange) SetComponentMin(value float64) {
	s.ptr.component_min = (C.double)(value)
}

// ComponentMax gets the component_max field.
func (s *AVOptionRange) ComponentMax() float64 {
	value := s.ptr.component_max
	return float64(value)
}

// SetComponentMax sets the component_max field.
func (s *AVOptionRange) SetComponentMax(value float64) {
	s.ptr.component_max = (C.double)(value)
}

// IsRange gets the is_range field.
func (s *AVOptionRange) IsRange() int {
	value := s.ptr.is_range
	return int(value)
}

// SetIsRange sets the is_range field.
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
func (s *AVOptionRanges) Range() *Array[*AVOptionRange] {
	value := s.ptr._range
	return ToAVOptionRangeArray(unsafe.Pointer(value))
}

// SetRange sets the range field.
func (s *AVOptionRanges) SetRange(value *Array[AVOptionRange]) {
	if value != nil {
		s.ptr._range = (**C.AVOptionRange)(value.ptr)
	} else {
		s.ptr._range = nil
	}
}

// NbRanges gets the nb_ranges field.
func (s *AVOptionRanges) NbRanges() int {
	value := s.ptr.nb_ranges
	return int(value)
}

// SetNbRanges sets the nb_ranges field.
func (s *AVOptionRanges) SetNbRanges(value int) {
	s.ptr.nb_ranges = (C.int)(value)
}

// NbComponents gets the nb_components field.
func (s *AVOptionRanges) NbComponents() int {
	value := s.ptr.nb_components
	return int(value)
}

// SetNbComponents sets the nb_components field.
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
func (s *AVRational) Num() int {
	value := s.value.num
	return int(value)
}

// SetNum sets the num field.
func (s *AVRational) SetNum(value int) {
	s.value.num = (C.int)(value)
}

// Den gets the den field.
func (s *AVRational) Den() int {
	value := s.value.den
	return int(value)
}

// SetDen sets the den field.
func (s *AVRational) SetDen(value int) {
	s.value.den = (C.int)(value)
}
