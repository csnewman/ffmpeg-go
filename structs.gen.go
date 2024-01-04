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
// #include <libavutil/frame.h>
// #include <libavutil/hwcontext.h>
// #include <libavutil/log.h>
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
func (s *RcOverride) EndFrame() int {
	value := s.ptr.end_frame
	return int(value)
}
func (s *RcOverride) Qscale() int {
	value := s.ptr.qscale
	return int(value)
}
func (s *RcOverride) QualityFactor() float32 {
	value := s.ptr.quality_factor
	return float32(value)
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
func (s *AVCodecContext) LogLevelOffset() int {
	value := s.ptr.log_level_offset
	return int(value)
}
func (s *AVCodecContext) CodecType() AVMediaType {
	value := s.ptr.codec_type
	return AVMediaType(value)
}
func (s *AVCodecContext) Codec() *AVCodec {
	value := s.ptr.codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}
func (s *AVCodecContext) CodecId() AVCodecID {
	value := s.ptr.codec_id
	return AVCodecID(value)
}
func (s *AVCodecContext) CodecTag() uint {
	value := s.ptr.codec_tag
	return uint(value)
}
func (s *AVCodecContext) Internal() *AVCodecInternal {
	value := s.ptr.internal
	var valueMapped *AVCodecInternal
	if value != nil {
		valueMapped = &AVCodecInternal{ptr: value}
	}
	return valueMapped
}
func (s *AVCodecContext) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}
func (s *AVCodecContext) BitRateTolerance() int {
	value := s.ptr.bit_rate_tolerance
	return int(value)
}
func (s *AVCodecContext) GlobalQuality() int {
	value := s.ptr.global_quality
	return int(value)
}
func (s *AVCodecContext) CompressionLevel() int {
	value := s.ptr.compression_level
	return int(value)
}
func (s *AVCodecContext) Flags() int {
	value := s.ptr.flags
	return int(value)
}
func (s *AVCodecContext) Flags2() int {
	value := s.ptr.flags2
	return int(value)
}

// Extradata skipped due to ptr to uint8

func (s *AVCodecContext) ExtradataSize() int {
	value := s.ptr.extradata_size
	return int(value)
}
func (s *AVCodecContext) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}
func (s *AVCodecContext) TicksPerFrame() int {
	value := s.ptr.ticks_per_frame
	return int(value)
}
func (s *AVCodecContext) Delay() int {
	value := s.ptr.delay
	return int(value)
}
func (s *AVCodecContext) Width() int {
	value := s.ptr.width
	return int(value)
}
func (s *AVCodecContext) Height() int {
	value := s.ptr.height
	return int(value)
}
func (s *AVCodecContext) CodedWidth() int {
	value := s.ptr.coded_width
	return int(value)
}
func (s *AVCodecContext) CodedHeight() int {
	value := s.ptr.coded_height
	return int(value)
}
func (s *AVCodecContext) GopSize() int {
	value := s.ptr.gop_size
	return int(value)
}
func (s *AVCodecContext) PixFmt() AVPixelFormat {
	value := s.ptr.pix_fmt
	return AVPixelFormat(value)
}

// DrawHorizBand skipped due to func ptr

// GetFormat skipped due to func ptr

func (s *AVCodecContext) MaxBFrames() int {
	value := s.ptr.max_b_frames
	return int(value)
}
func (s *AVCodecContext) BQuantFactor() float32 {
	value := s.ptr.b_quant_factor
	return float32(value)
}
func (s *AVCodecContext) BQuantOffset() float32 {
	value := s.ptr.b_quant_offset
	return float32(value)
}
func (s *AVCodecContext) HasBFrames() int {
	value := s.ptr.has_b_frames
	return int(value)
}
func (s *AVCodecContext) IQuantFactor() float32 {
	value := s.ptr.i_quant_factor
	return float32(value)
}
func (s *AVCodecContext) IQuantOffset() float32 {
	value := s.ptr.i_quant_offset
	return float32(value)
}
func (s *AVCodecContext) LumiMasking() float32 {
	value := s.ptr.lumi_masking
	return float32(value)
}
func (s *AVCodecContext) TemporalCplxMasking() float32 {
	value := s.ptr.temporal_cplx_masking
	return float32(value)
}
func (s *AVCodecContext) SpatialCplxMasking() float32 {
	value := s.ptr.spatial_cplx_masking
	return float32(value)
}
func (s *AVCodecContext) PMasking() float32 {
	value := s.ptr.p_masking
	return float32(value)
}
func (s *AVCodecContext) DarkMasking() float32 {
	value := s.ptr.dark_masking
	return float32(value)
}
func (s *AVCodecContext) SliceCount() int {
	value := s.ptr.slice_count
	return int(value)
}

// SliceOffset skipped due to prim ptr

func (s *AVCodecContext) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}
func (s *AVCodecContext) MeCmp() int {
	value := s.ptr.me_cmp
	return int(value)
}
func (s *AVCodecContext) MeSubCmp() int {
	value := s.ptr.me_sub_cmp
	return int(value)
}
func (s *AVCodecContext) MbCmp() int {
	value := s.ptr.mb_cmp
	return int(value)
}
func (s *AVCodecContext) IldctCmp() int {
	value := s.ptr.ildct_cmp
	return int(value)
}
func (s *AVCodecContext) DiaSize() int {
	value := s.ptr.dia_size
	return int(value)
}
func (s *AVCodecContext) LastPredictorCount() int {
	value := s.ptr.last_predictor_count
	return int(value)
}
func (s *AVCodecContext) MePreCmp() int {
	value := s.ptr.me_pre_cmp
	return int(value)
}
func (s *AVCodecContext) PreDiaSize() int {
	value := s.ptr.pre_dia_size
	return int(value)
}
func (s *AVCodecContext) MeSubpelQuality() int {
	value := s.ptr.me_subpel_quality
	return int(value)
}
func (s *AVCodecContext) MeRange() int {
	value := s.ptr.me_range
	return int(value)
}
func (s *AVCodecContext) SliceFlags() int {
	value := s.ptr.slice_flags
	return int(value)
}
func (s *AVCodecContext) MbDecision() int {
	value := s.ptr.mb_decision
	return int(value)
}

// IntraMatrix skipped due to prim ptr

// InterMatrix skipped due to prim ptr

func (s *AVCodecContext) IntraDcPrecision() int {
	value := s.ptr.intra_dc_precision
	return int(value)
}
func (s *AVCodecContext) SkipTop() int {
	value := s.ptr.skip_top
	return int(value)
}
func (s *AVCodecContext) SkipBottom() int {
	value := s.ptr.skip_bottom
	return int(value)
}
func (s *AVCodecContext) MbLmin() int {
	value := s.ptr.mb_lmin
	return int(value)
}
func (s *AVCodecContext) MbLmax() int {
	value := s.ptr.mb_lmax
	return int(value)
}
func (s *AVCodecContext) BidirRefine() int {
	value := s.ptr.bidir_refine
	return int(value)
}
func (s *AVCodecContext) KeyintMin() int {
	value := s.ptr.keyint_min
	return int(value)
}
func (s *AVCodecContext) Refs() int {
	value := s.ptr.refs
	return int(value)
}
func (s *AVCodecContext) Mv0Threshold() int {
	value := s.ptr.mv0_threshold
	return int(value)
}
func (s *AVCodecContext) ColorPrimaries() AVColorPrimaries {
	value := s.ptr.color_primaries
	return AVColorPrimaries(value)
}
func (s *AVCodecContext) ColorTrc() AVColorTransferCharacteristic {
	value := s.ptr.color_trc
	return AVColorTransferCharacteristic(value)
}
func (s *AVCodecContext) Colorspace() AVColorSpace {
	value := s.ptr.colorspace
	return AVColorSpace(value)
}
func (s *AVCodecContext) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}
func (s *AVCodecContext) ChromaSampleLocation() AVChromaLocation {
	value := s.ptr.chroma_sample_location
	return AVChromaLocation(value)
}
func (s *AVCodecContext) Slices() int {
	value := s.ptr.slices
	return int(value)
}
func (s *AVCodecContext) FieldOrder() AVFieldOrder {
	value := s.ptr.field_order
	return AVFieldOrder(value)
}
func (s *AVCodecContext) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}
func (s *AVCodecContext) Channels() int {
	value := s.ptr.channels
	return int(value)
}
func (s *AVCodecContext) SampleFmt() AVSampleFormat {
	value := s.ptr.sample_fmt
	return AVSampleFormat(value)
}
func (s *AVCodecContext) FrameSize() int {
	value := s.ptr.frame_size
	return int(value)
}
func (s *AVCodecContext) FrameNumber() int {
	value := s.ptr.frame_number
	return int(value)
}
func (s *AVCodecContext) BlockAlign() int {
	value := s.ptr.block_align
	return int(value)
}
func (s *AVCodecContext) Cutoff() int {
	value := s.ptr.cutoff
	return int(value)
}
func (s *AVCodecContext) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}
func (s *AVCodecContext) RequestChannelLayout() uint64 {
	value := s.ptr.request_channel_layout
	return uint64(value)
}
func (s *AVCodecContext) AudioServiceType() AVAudioServiceType {
	value := s.ptr.audio_service_type
	return AVAudioServiceType(value)
}
func (s *AVCodecContext) RequestSampleFmt() AVSampleFormat {
	value := s.ptr.request_sample_fmt
	return AVSampleFormat(value)
}

// GetBuffer2 skipped due to func ptr

func (s *AVCodecContext) Qcompress() float32 {
	value := s.ptr.qcompress
	return float32(value)
}
func (s *AVCodecContext) Qblur() float32 {
	value := s.ptr.qblur
	return float32(value)
}
func (s *AVCodecContext) Qmin() int {
	value := s.ptr.qmin
	return int(value)
}
func (s *AVCodecContext) Qmax() int {
	value := s.ptr.qmax
	return int(value)
}
func (s *AVCodecContext) MaxQdiff() int {
	value := s.ptr.max_qdiff
	return int(value)
}
func (s *AVCodecContext) RcBufferSize() int {
	value := s.ptr.rc_buffer_size
	return int(value)
}
func (s *AVCodecContext) RcOverrideCount() int {
	value := s.ptr.rc_override_count
	return int(value)
}
func (s *AVCodecContext) RcOverride() *RcOverride {
	value := s.ptr.rc_override
	var valueMapped *RcOverride
	if value != nil {
		valueMapped = &RcOverride{ptr: value}
	}
	return valueMapped
}
func (s *AVCodecContext) RcMaxRate() int64 {
	value := s.ptr.rc_max_rate
	return int64(value)
}
func (s *AVCodecContext) RcMinRate() int64 {
	value := s.ptr.rc_min_rate
	return int64(value)
}
func (s *AVCodecContext) RcMaxAvailableVbvUse() float32 {
	value := s.ptr.rc_max_available_vbv_use
	return float32(value)
}
func (s *AVCodecContext) RcMinVbvOverflowUse() float32 {
	value := s.ptr.rc_min_vbv_overflow_use
	return float32(value)
}
func (s *AVCodecContext) RcInitialBufferOccupancy() int {
	value := s.ptr.rc_initial_buffer_occupancy
	return int(value)
}
func (s *AVCodecContext) Trellis() int {
	value := s.ptr.trellis
	return int(value)
}
func (s *AVCodecContext) StatsOut() *CStr {
	value := s.ptr.stats_out
	return wrapCStr(value)
}
func (s *AVCodecContext) StatsIn() *CStr {
	value := s.ptr.stats_in
	return wrapCStr(value)
}
func (s *AVCodecContext) WorkaroundBugs() int {
	value := s.ptr.workaround_bugs
	return int(value)
}
func (s *AVCodecContext) StrictStdCompliance() int {
	value := s.ptr.strict_std_compliance
	return int(value)
}
func (s *AVCodecContext) ErrorConcealment() int {
	value := s.ptr.error_concealment
	return int(value)
}
func (s *AVCodecContext) Debug() int {
	value := s.ptr.debug
	return int(value)
}
func (s *AVCodecContext) ErrRecognition() int {
	value := s.ptr.err_recognition
	return int(value)
}
func (s *AVCodecContext) ReorderedOpaque() int64 {
	value := s.ptr.reordered_opaque
	return int64(value)
}
func (s *AVCodecContext) Hwaccel() *AVHWAccel {
	value := s.ptr.hwaccel
	var valueMapped *AVHWAccel
	if value != nil {
		valueMapped = &AVHWAccel{ptr: value}
	}
	return valueMapped
}

// Error skipped due to const array

func (s *AVCodecContext) DctAlgo() int {
	value := s.ptr.dct_algo
	return int(value)
}
func (s *AVCodecContext) IdctAlgo() int {
	value := s.ptr.idct_algo
	return int(value)
}
func (s *AVCodecContext) BitsPerCodedSample() int {
	value := s.ptr.bits_per_coded_sample
	return int(value)
}
func (s *AVCodecContext) BitsPerRawSample() int {
	value := s.ptr.bits_per_raw_sample
	return int(value)
}
func (s *AVCodecContext) Lowres() int {
	value := s.ptr.lowres
	return int(value)
}
func (s *AVCodecContext) ThreadCount() int {
	value := s.ptr.thread_count
	return int(value)
}
func (s *AVCodecContext) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}
func (s *AVCodecContext) ActiveThreadType() int {
	value := s.ptr.active_thread_type
	return int(value)
}

// Execute skipped due to func ptr

// Execute2 skipped due to func ptr

func (s *AVCodecContext) NsseWeight() int {
	value := s.ptr.nsse_weight
	return int(value)
}
func (s *AVCodecContext) Profile() int {
	value := s.ptr.profile
	return int(value)
}
func (s *AVCodecContext) Level() int {
	value := s.ptr.level
	return int(value)
}
func (s *AVCodecContext) SkipLoopFilter() AVDiscard {
	value := s.ptr.skip_loop_filter
	return AVDiscard(value)
}
func (s *AVCodecContext) SkipIdct() AVDiscard {
	value := s.ptr.skip_idct
	return AVDiscard(value)
}
func (s *AVCodecContext) SkipFrame() AVDiscard {
	value := s.ptr.skip_frame
	return AVDiscard(value)
}

// SubtitleHeader skipped due to ptr to uint8

func (s *AVCodecContext) SubtitleHeaderSize() int {
	value := s.ptr.subtitle_header_size
	return int(value)
}
func (s *AVCodecContext) InitialPadding() int {
	value := s.ptr.initial_padding
	return int(value)
}
func (s *AVCodecContext) Framerate() *AVRational {
	value := s.ptr.framerate
	return &AVRational{value: value}
}
func (s *AVCodecContext) SwPixFmt() AVPixelFormat {
	value := s.ptr.sw_pix_fmt
	return AVPixelFormat(value)
}
func (s *AVCodecContext) PktTimebase() *AVRational {
	value := s.ptr.pkt_timebase
	return &AVRational{value: value}
}
func (s *AVCodecContext) CodecDescriptor() *AVCodecDescriptor {
	value := s.ptr.codec_descriptor
	var valueMapped *AVCodecDescriptor
	if value != nil {
		valueMapped = &AVCodecDescriptor{ptr: value}
	}
	return valueMapped
}
func (s *AVCodecContext) PtsCorrectionNumFaultyPts() int64 {
	value := s.ptr.pts_correction_num_faulty_pts
	return int64(value)
}
func (s *AVCodecContext) PtsCorrectionNumFaultyDts() int64 {
	value := s.ptr.pts_correction_num_faulty_dts
	return int64(value)
}
func (s *AVCodecContext) PtsCorrectionLastPts() int64 {
	value := s.ptr.pts_correction_last_pts
	return int64(value)
}
func (s *AVCodecContext) PtsCorrectionLastDts() int64 {
	value := s.ptr.pts_correction_last_dts
	return int64(value)
}
func (s *AVCodecContext) SubCharenc() *CStr {
	value := s.ptr.sub_charenc
	return wrapCStr(value)
}
func (s *AVCodecContext) SubCharencMode() int {
	value := s.ptr.sub_charenc_mode
	return int(value)
}
func (s *AVCodecContext) SkipAlpha() int {
	value := s.ptr.skip_alpha
	return int(value)
}
func (s *AVCodecContext) SeekPreroll() int {
	value := s.ptr.seek_preroll
	return int(value)
}

// ChromaIntraMatrix skipped due to prim ptr

// DumpSeparator skipped due to ptr to uint8

func (s *AVCodecContext) CodecWhitelist() *CStr {
	value := s.ptr.codec_whitelist
	return wrapCStr(value)
}
func (s *AVCodecContext) Properties() uint {
	value := s.ptr.properties
	return uint(value)
}
func (s *AVCodecContext) CodedSideData() *AVPacketSideData {
	value := s.ptr.coded_side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}
func (s *AVCodecContext) NbCodedSideData() int {
	value := s.ptr.nb_coded_side_data
	return int(value)
}
func (s *AVCodecContext) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}
func (s *AVCodecContext) TrailingPadding() int {
	value := s.ptr.trailing_padding
	return int(value)
}
func (s *AVCodecContext) MaxPixels() int64 {
	value := s.ptr.max_pixels
	return int64(value)
}
func (s *AVCodecContext) HwDeviceCtx() *AVBufferRef {
	value := s.ptr.hw_device_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}
func (s *AVCodecContext) HwaccelFlags() int {
	value := s.ptr.hwaccel_flags
	return int(value)
}
func (s *AVCodecContext) ApplyCropping() int {
	value := s.ptr.apply_cropping
	return int(value)
}
func (s *AVCodecContext) ExtraHwFrames() int {
	value := s.ptr.extra_hw_frames
	return int(value)
}
func (s *AVCodecContext) DiscardDamagedPercentage() int {
	value := s.ptr.discard_damaged_percentage
	return int(value)
}
func (s *AVCodecContext) MaxSamples() int64 {
	value := s.ptr.max_samples
	return int64(value)
}
func (s *AVCodecContext) ExportSideData() int {
	value := s.ptr.export_side_data
	return int(value)
}

// GetEncodeBuffer skipped due to func ptr

// ChLayout skipped due to ident byvalue

func (s *AVCodecContext) FrameNum() int64 {
	value := s.ptr.frame_num
	return int64(value)
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
func (s *AVHWAccel) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}
func (s *AVHWAccel) Id() AVCodecID {
	value := s.ptr.id
	return AVCodecID(value)
}
func (s *AVHWAccel) PixFmt() AVPixelFormat {
	value := s.ptr.pix_fmt
	return AVPixelFormat(value)
}
func (s *AVHWAccel) Capabilities() int {
	value := s.ptr.capabilities
	return int(value)
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

// Init skipped due to func ptr

// Uninit skipped due to func ptr

func (s *AVHWAccel) PrivDataSize() int {
	value := s.ptr.priv_data_size
	return int(value)
}
func (s *AVHWAccel) CapsInternal() int {
	value := s.ptr.caps_internal
	return int(value)
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
func (s *AVSubtitleRect) Y() int {
	value := s.ptr.y
	return int(value)
}
func (s *AVSubtitleRect) W() int {
	value := s.ptr.w
	return int(value)
}
func (s *AVSubtitleRect) H() int {
	value := s.ptr.h
	return int(value)
}
func (s *AVSubtitleRect) NbColors() int {
	value := s.ptr.nb_colors
	return int(value)
}

// Data skipped due to const array

// Linesize skipped due to const array

func (s *AVSubtitleRect) Type() AVSubtitleType {
	value := s.ptr._type
	return AVSubtitleType(value)
}
func (s *AVSubtitleRect) Text() *CStr {
	value := s.ptr.text
	return wrapCStr(value)
}
func (s *AVSubtitleRect) Ass() *CStr {
	value := s.ptr.ass
	return wrapCStr(value)
}
func (s *AVSubtitleRect) Flags() int {
	value := s.ptr.flags
	return int(value)
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
func (s *AVSubtitle) StartDisplayTime() uint32 {
	value := s.ptr.start_display_time
	return uint32(value)
}
func (s *AVSubtitle) EndDisplayTime() uint32 {
	value := s.ptr.end_display_time
	return uint32(value)
}
func (s *AVSubtitle) NumRects() uint {
	value := s.ptr.num_rects
	return uint(value)
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

// --- Struct AVCodecParserContext ---

// AVCodecParserContext wraps AVCodecParserContext.
type AVCodecParserContext struct {
	ptr *C.AVCodecParserContext
}

func (s *AVCodecParserContext) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}
func (s *AVCodecParserContext) Parser() *AVCodecParser {
	value := s.ptr.parser
	var valueMapped *AVCodecParser
	if value != nil {
		valueMapped = &AVCodecParser{ptr: value}
	}
	return valueMapped
}
func (s *AVCodecParserContext) FrameOffset() int64 {
	value := s.ptr.frame_offset
	return int64(value)
}
func (s *AVCodecParserContext) CurOffset() int64 {
	value := s.ptr.cur_offset
	return int64(value)
}
func (s *AVCodecParserContext) NextFrameOffset() int64 {
	value := s.ptr.next_frame_offset
	return int64(value)
}
func (s *AVCodecParserContext) PictType() int {
	value := s.ptr.pict_type
	return int(value)
}
func (s *AVCodecParserContext) RepeatPict() int {
	value := s.ptr.repeat_pict
	return int(value)
}
func (s *AVCodecParserContext) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}
func (s *AVCodecParserContext) Dts() int64 {
	value := s.ptr.dts
	return int64(value)
}
func (s *AVCodecParserContext) LastPts() int64 {
	value := s.ptr.last_pts
	return int64(value)
}
func (s *AVCodecParserContext) LastDts() int64 {
	value := s.ptr.last_dts
	return int64(value)
}
func (s *AVCodecParserContext) FetchTimestamp() int {
	value := s.ptr.fetch_timestamp
	return int(value)
}
func (s *AVCodecParserContext) CurFrameStartIndex() int {
	value := s.ptr.cur_frame_start_index
	return int(value)
}

// CurFrameOffset skipped due to const array

// CurFramePts skipped due to const array

// CurFrameDts skipped due to const array

func (s *AVCodecParserContext) Flags() int {
	value := s.ptr.flags
	return int(value)
}
func (s *AVCodecParserContext) Offset() int64 {
	value := s.ptr.offset
	return int64(value)
}

// CurFrameEnd skipped due to const array

func (s *AVCodecParserContext) KeyFrame() int {
	value := s.ptr.key_frame
	return int(value)
}
func (s *AVCodecParserContext) DtsSyncPoint() int {
	value := s.ptr.dts_sync_point
	return int(value)
}
func (s *AVCodecParserContext) DtsRefDtsDelta() int {
	value := s.ptr.dts_ref_dts_delta
	return int(value)
}
func (s *AVCodecParserContext) PtsDtsDelta() int {
	value := s.ptr.pts_dts_delta
	return int(value)
}

// CurFramePos skipped due to const array

func (s *AVCodecParserContext) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}
func (s *AVCodecParserContext) LastPos() int64 {
	value := s.ptr.last_pos
	return int64(value)
}
func (s *AVCodecParserContext) Duration() int {
	value := s.ptr.duration
	return int(value)
}
func (s *AVCodecParserContext) FieldOrder() AVFieldOrder {
	value := s.ptr.field_order
	return AVFieldOrder(value)
}
func (s *AVCodecParserContext) PictureStructure() AVPictureStructure {
	value := s.ptr.picture_structure
	return AVPictureStructure(value)
}
func (s *AVCodecParserContext) OutputPictureNumber() int {
	value := s.ptr.output_picture_number
	return int(value)
}
func (s *AVCodecParserContext) Width() int {
	value := s.ptr.width
	return int(value)
}
func (s *AVCodecParserContext) Height() int {
	value := s.ptr.height
	return int(value)
}
func (s *AVCodecParserContext) CodedWidth() int {
	value := s.ptr.coded_width
	return int(value)
}
func (s *AVCodecParserContext) CodedHeight() int {
	value := s.ptr.coded_height
	return int(value)
}
func (s *AVCodecParserContext) Format() int {
	value := s.ptr.format
	return int(value)
}

// --- Struct AVCodecParser ---

// AVCodecParser wraps AVCodecParser.
type AVCodecParser struct {
	ptr *C.AVCodecParser
}

func (s *AVCodecParser) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

// CodecIds skipped due to const array

func (s *AVCodecParser) PrivDataSize() int {
	value := s.ptr.priv_data_size
	return int(value)
}

// ParserInit skipped due to func ptr

// ParserParse skipped due to func ptr

// ParserClose skipped due to func ptr

// Split skipped due to func ptr

// --- Struct AVProfile ---

// AVProfile wraps AVProfile.
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
func (s *AVProfile) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}

// --- Struct AVCodec ---

// AVCodec wraps AVCodec.
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
func (s *AVCodec) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}
func (s *AVCodec) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}
func (s *AVCodec) Id() AVCodecID {
	value := s.ptr.id
	return AVCodecID(value)
}
func (s *AVCodec) Capabilities() int {
	value := s.ptr.capabilities
	return int(value)
}
func (s *AVCodec) MaxLowres() uint8 {
	value := s.ptr.max_lowres
	return uint8(value)
}

// SupportedFramerates skipped due to struct value ptr

// PixFmts skipped due to enum ptr

// SupportedSamplerates skipped due to prim ptr

// SampleFmts skipped due to enum ptr

// ChannelLayouts skipped due to prim ptr

func (s *AVCodec) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}
func (s *AVCodec) Profiles() *AVProfile {
	value := s.ptr.profiles
	var valueMapped *AVProfile
	if value != nil {
		valueMapped = &AVProfile{ptr: value}
	}
	return valueMapped
}
func (s *AVCodec) WrapperName() *CStr {
	value := s.ptr.wrapper_name
	return wrapCStr(value)
}
func (s *AVCodec) ChLayouts() *AVChannelLayout {
	value := s.ptr.ch_layouts
	var valueMapped *AVChannelLayout
	if value != nil {
		valueMapped = &AVChannelLayout{ptr: value}
	}
	return valueMapped
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
func (s *AVCodecHWConfig) Methods() int {
	value := s.ptr.methods
	return int(value)
}
func (s *AVCodecHWConfig) DeviceType() AVHWDeviceType {
	value := s.ptr.device_type
	return AVHWDeviceType(value)
}

// --- Struct AVCodecDescriptor ---

// AVCodecDescriptor wraps AVCodecDescriptor.
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
func (s *AVCodecDescriptor) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}
func (s *AVCodecDescriptor) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}
func (s *AVCodecDescriptor) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}
func (s *AVCodecDescriptor) Props() int {
	value := s.ptr.props
	return int(value)
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

// --- Struct AVCodecParameters ---

// AVCodecParameters wraps AVCodecParameters.
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
func (s *AVCodecParameters) CodecId() AVCodecID {
	value := s.ptr.codec_id
	return AVCodecID(value)
}
func (s *AVCodecParameters) CodecTag() uint32 {
	value := s.ptr.codec_tag
	return uint32(value)
}

// Extradata skipped due to ptr to uint8

func (s *AVCodecParameters) ExtradataSize() int {
	value := s.ptr.extradata_size
	return int(value)
}
func (s *AVCodecParameters) Format() int {
	value := s.ptr.format
	return int(value)
}
func (s *AVCodecParameters) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}
func (s *AVCodecParameters) BitsPerCodedSample() int {
	value := s.ptr.bits_per_coded_sample
	return int(value)
}
func (s *AVCodecParameters) BitsPerRawSample() int {
	value := s.ptr.bits_per_raw_sample
	return int(value)
}
func (s *AVCodecParameters) Profile() int {
	value := s.ptr.profile
	return int(value)
}
func (s *AVCodecParameters) Level() int {
	value := s.ptr.level
	return int(value)
}
func (s *AVCodecParameters) Width() int {
	value := s.ptr.width
	return int(value)
}
func (s *AVCodecParameters) Height() int {
	value := s.ptr.height
	return int(value)
}
func (s *AVCodecParameters) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}
func (s *AVCodecParameters) FieldOrder() AVFieldOrder {
	value := s.ptr.field_order
	return AVFieldOrder(value)
}
func (s *AVCodecParameters) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}
func (s *AVCodecParameters) ColorPrimaries() AVColorPrimaries {
	value := s.ptr.color_primaries
	return AVColorPrimaries(value)
}
func (s *AVCodecParameters) ColorTrc() AVColorTransferCharacteristic {
	value := s.ptr.color_trc
	return AVColorTransferCharacteristic(value)
}
func (s *AVCodecParameters) ColorSpace() AVColorSpace {
	value := s.ptr.color_space
	return AVColorSpace(value)
}
func (s *AVCodecParameters) ChromaLocation() AVChromaLocation {
	value := s.ptr.chroma_location
	return AVChromaLocation(value)
}
func (s *AVCodecParameters) VideoDelay() int {
	value := s.ptr.video_delay
	return int(value)
}
func (s *AVCodecParameters) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}
func (s *AVCodecParameters) Channels() int {
	value := s.ptr.channels
	return int(value)
}
func (s *AVCodecParameters) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}
func (s *AVCodecParameters) BlockAlign() int {
	value := s.ptr.block_align
	return int(value)
}
func (s *AVCodecParameters) FrameSize() int {
	value := s.ptr.frame_size
	return int(value)
}
func (s *AVCodecParameters) InitialPadding() int {
	value := s.ptr.initial_padding
	return int(value)
}
func (s *AVCodecParameters) TrailingPadding() int {
	value := s.ptr.trailing_padding
	return int(value)
}
func (s *AVCodecParameters) SeekPreroll() int {
	value := s.ptr.seek_preroll
	return int(value)
}

// ChLayout skipped due to ident byvalue

// --- Struct AVPanScan ---

// AVPanScan wraps AVPanScan.
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
func (s *AVPanScan) Width() int {
	value := s.ptr.width
	return int(value)
}
func (s *AVPanScan) Height() int {
	value := s.ptr.height
	return int(value)
}

// Position skipped due to const array

// --- Struct AVCPBProperties ---

// AVCPBProperties wraps AVCPBProperties.
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
func (s *AVCPBProperties) MinBitrate() int64 {
	value := s.ptr.min_bitrate
	return int64(value)
}
func (s *AVCPBProperties) AvgBitrate() int64 {
	value := s.ptr.avg_bitrate
	return int64(value)
}
func (s *AVCPBProperties) BufferSize() int64 {
	value := s.ptr.buffer_size
	return int64(value)
}
func (s *AVCPBProperties) VbvDelay() uint64 {
	value := s.ptr.vbv_delay
	return uint64(value)
}

// --- Struct AVProducerReferenceTime ---

// AVProducerReferenceTime wraps AVProducerReferenceTime.
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
func (s *AVProducerReferenceTime) Flags() int {
	value := s.ptr.flags
	return int(value)
}

// --- Struct AVPacketSideData ---

// AVPacketSideData wraps AVPacketSideData.
type AVPacketSideData struct {
	ptr *C.AVPacketSideData
}

func (s *AVPacketSideData) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

// Data skipped due to ptr to uint8

func (s *AVPacketSideData) Size() uint64 {
	value := s.ptr.size
	return uint64(value)
}
func (s *AVPacketSideData) Type() AVPacketSideDataType {
	value := s.ptr._type
	return AVPacketSideDataType(value)
}

// --- Struct AVPacket ---

// AVPacket wraps AVPacket.
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
func (s *AVPacket) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}
func (s *AVPacket) Dts() int64 {
	value := s.ptr.dts
	return int64(value)
}

// Data skipped due to ptr to uint8

func (s *AVPacket) Size() int {
	value := s.ptr.size
	return int(value)
}
func (s *AVPacket) StreamIndex() int {
	value := s.ptr.stream_index
	return int(value)
}
func (s *AVPacket) Flags() int {
	value := s.ptr.flags
	return int(value)
}
func (s *AVPacket) SideData() *AVPacketSideData {
	value := s.ptr.side_data
	var valueMapped *AVPacketSideData
	if value != nil {
		valueMapped = &AVPacketSideData{ptr: value}
	}
	return valueMapped
}
func (s *AVPacket) SideDataElems() int {
	value := s.ptr.side_data_elems
	return int(value)
}
func (s *AVPacket) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}
func (s *AVPacket) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}
func (s *AVPacket) OpaqueRef() *AVBufferRef {
	value := s.ptr.opaque_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}
func (s *AVPacket) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
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

// --- Struct AVFilterContext ---

// AVFilterContext wraps AVFilterContext.
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
func (s *AVFilterContext) Filter() *AVFilter {
	value := s.ptr.filter
	var valueMapped *AVFilter
	if value != nil {
		valueMapped = &AVFilter{ptr: value}
	}
	return valueMapped
}
func (s *AVFilterContext) Name() *CStr {
	value := s.ptr.name
	return wrapCStr(value)
}
func (s *AVFilterContext) InputPads() *AVFilterPad {
	value := s.ptr.input_pads
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
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
func (s *AVFilterContext) OutputPads() *AVFilterPad {
	value := s.ptr.output_pads
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
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
func (s *AVFilterContext) Graph() *AVFilterGraph {
	value := s.ptr.graph
	var valueMapped *AVFilterGraph
	if value != nil {
		valueMapped = &AVFilterGraph{ptr: value}
	}
	return valueMapped
}
func (s *AVFilterContext) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}
func (s *AVFilterContext) Internal() *AVFilterInternal {
	value := s.ptr.internal
	var valueMapped *AVFilterInternal
	if value != nil {
		valueMapped = &AVFilterInternal{ptr: value}
	}
	return valueMapped
}

// CommandQueue skipped due to ptr to ignored type

func (s *AVFilterContext) EnableStr() *CStr {
	value := s.ptr.enable_str
	return wrapCStr(value)
}

// VarValues skipped due to prim ptr

func (s *AVFilterContext) IsDisabled() int {
	value := s.ptr.is_disabled
	return int(value)
}
func (s *AVFilterContext) HwDeviceCtx() *AVBufferRef {
	value := s.ptr.hw_device_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}
func (s *AVFilterContext) NbThreads() int {
	value := s.ptr.nb_threads
	return int(value)
}
func (s *AVFilterContext) Ready() uint {
	value := s.ptr.ready
	return uint(value)
}
func (s *AVFilterContext) ExtraHwFrames() int {
	value := s.ptr.extra_hw_frames
	return int(value)
}

// --- Struct AVFilterLink ---

// AVFilterLink wraps AVFilterLink.
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
func (s *AVFilterLink) Srcpad() *AVFilterPad {
	value := s.ptr.srcpad
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}
func (s *AVFilterLink) Dst() *AVFilterContext {
	value := s.ptr.dst
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}
func (s *AVFilterLink) Dstpad() *AVFilterPad {
	value := s.ptr.dstpad
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}
func (s *AVFilterLink) Type() AVMediaType {
	value := s.ptr._type
	return AVMediaType(value)
}
func (s *AVFilterLink) W() int {
	value := s.ptr.w
	return int(value)
}
func (s *AVFilterLink) H() int {
	value := s.ptr.h
	return int(value)
}
func (s *AVFilterLink) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}
func (s *AVFilterLink) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}
func (s *AVFilterLink) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}
func (s *AVFilterLink) Format() int {
	value := s.ptr.format
	return int(value)
}
func (s *AVFilterLink) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}

// ChLayout skipped due to ident byvalue

// Incfg skipped due to ident byvalue

// Outcfg skipped due to ident byvalue

func (s *AVFilterLink) InitState() int {
	value := s.ptr.init_state
	return int(value)
}
func (s *AVFilterLink) Graph() *AVFilterGraph {
	value := s.ptr.graph
	var valueMapped *AVFilterGraph
	if value != nil {
		valueMapped = &AVFilterGraph{ptr: value}
	}
	return valueMapped
}
func (s *AVFilterLink) CurrentPts() int64 {
	value := s.ptr.current_pts
	return int64(value)
}
func (s *AVFilterLink) CurrentPtsUs() int64 {
	value := s.ptr.current_pts_us
	return int64(value)
}
func (s *AVFilterLink) AgeIndex() int {
	value := s.ptr.age_index
	return int(value)
}
func (s *AVFilterLink) FrameRate() *AVRational {
	value := s.ptr.frame_rate
	return &AVRational{value: value}
}
func (s *AVFilterLink) MinSamples() int {
	value := s.ptr.min_samples
	return int(value)
}
func (s *AVFilterLink) MaxSamples() int {
	value := s.ptr.max_samples
	return int(value)
}
func (s *AVFilterLink) FrameCountIn() int64 {
	value := s.ptr.frame_count_in
	return int64(value)
}
func (s *AVFilterLink) FrameCountOut() int64 {
	value := s.ptr.frame_count_out
	return int64(value)
}
func (s *AVFilterLink) SampleCountIn() int64 {
	value := s.ptr.sample_count_in
	return int64(value)
}
func (s *AVFilterLink) SampleCountOut() int64 {
	value := s.ptr.sample_count_out
	return int64(value)
}
func (s *AVFilterLink) FrameWantedOut() int {
	value := s.ptr.frame_wanted_out
	return int(value)
}
func (s *AVFilterLink) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// Reserved skipped due to const array

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
func (s *AVFilter) Description() *CStr {
	value := s.ptr.description
	return wrapCStr(value)
}
func (s *AVFilter) Inputs() *AVFilterPad {
	value := s.ptr.inputs
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}
func (s *AVFilter) Outputs() *AVFilterPad {
	value := s.ptr.outputs
	var valueMapped *AVFilterPad
	if value != nil {
		valueMapped = &AVFilterPad{ptr: value}
	}
	return valueMapped
}
func (s *AVFilter) PrivClass() *AVClass {
	value := s.ptr.priv_class
	var valueMapped *AVClass
	if value != nil {
		valueMapped = &AVClass{ptr: value}
	}
	return valueMapped
}
func (s *AVFilter) Flags() int {
	value := s.ptr.flags
	return int(value)
}
func (s *AVFilter) NbInputs() uint8 {
	value := s.ptr.nb_inputs
	return uint8(value)
}
func (s *AVFilter) NbOutputs() uint8 {
	value := s.ptr.nb_outputs
	return uint8(value)
}
func (s *AVFilter) FormatsState() uint8 {
	value := s.ptr.formats_state
	return uint8(value)
}

// Preinit skipped due to func ptr

// Init skipped due to func ptr

// Uninit skipped due to func ptr

// Formats skipped due to union type

func (s *AVFilter) PrivSize() int {
	value := s.ptr.priv_size
	return int(value)
}
func (s *AVFilter) FlagsInternal() int {
	value := s.ptr.flags_internal
	return int(value)
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
func (s *AVFilterFormatsConfig) Samplerates() *AVFilterFormats {
	value := s.ptr.samplerates
	var valueMapped *AVFilterFormats
	if value != nil {
		valueMapped = &AVFilterFormats{ptr: value}
	}
	return valueMapped
}
func (s *AVFilterFormatsConfig) ChannelLayouts() *AVFilterChannelLayouts {
	value := s.ptr.channel_layouts
	var valueMapped *AVFilterChannelLayouts
	if value != nil {
		valueMapped = &AVFilterChannelLayouts{ptr: value}
	}
	return valueMapped
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
func (s *AVFilterGraph) ScaleSwsOpts() *CStr {
	value := s.ptr.scale_sws_opts
	return wrapCStr(value)
}
func (s *AVFilterGraph) ThreadType() int {
	value := s.ptr.thread_type
	return int(value)
}
func (s *AVFilterGraph) NbThreads() int {
	value := s.ptr.nb_threads
	return int(value)
}
func (s *AVFilterGraph) Internal() *AVFilterGraphInternal {
	value := s.ptr.internal
	var valueMapped *AVFilterGraphInternal
	if value != nil {
		valueMapped = &AVFilterGraphInternal{ptr: value}
	}
	return valueMapped
}

// Execute skipped due to callback ptr

func (s *AVFilterGraph) AresampleSwrOpts() *CStr {
	value := s.ptr.aresample_swr_opts
	return wrapCStr(value)
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
func (s *AVFilterGraph) DisableAutoConvert() uint {
	value := s.ptr.disable_auto_convert
	return uint(value)
}

// --- Struct AVFilterInOut ---

// AVFilterInOut wraps AVFilterInOut.
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
func (s *AVFilterInOut) FilterCtx() *AVFilterContext {
	value := s.ptr.filter_ctx
	var valueMapped *AVFilterContext
	if value != nil {
		valueMapped = &AVFilterContext{ptr: value}
	}
	return valueMapped
}
func (s *AVFilterInOut) PadIdx() int {
	value := s.ptr.pad_idx
	return int(value)
}
func (s *AVFilterInOut) Next() *AVFilterInOut {
	value := s.ptr.next
	var valueMapped *AVFilterInOut
	if value != nil {
		valueMapped = &AVFilterInOut{ptr: value}
	}
	return valueMapped
}

// --- Struct AVFilterPadParams ---

// AVFilterPadParams wraps AVFilterPadParams.
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

// --- Struct AVFilterParams ---

// AVFilterParams wraps AVFilterParams.
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
func (s *AVFilterParams) FilterName() *CStr {
	value := s.ptr.filter_name
	return wrapCStr(value)
}
func (s *AVFilterParams) InstanceName() *CStr {
	value := s.ptr.instance_name
	return wrapCStr(value)
}
func (s *AVFilterParams) Opts() *AVDictionary {
	value := s.ptr.opts
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
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

// --- Struct AVFilterChain ---

// AVFilterChain wraps AVFilterChain.
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

// --- Struct AVFilterGraphSegment ---

// AVFilterGraphSegment wraps AVFilterGraphSegment.
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
func (s *AVFilterGraphSegment) ScaleSwsOpts() *CStr {
	value := s.ptr.scale_sws_opts
	return wrapCStr(value)
}

// --- Struct AVBufferSrcParameters ---

// AVBufferSrcParameters wraps AVBufferSrcParameters.
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
func (s *AVBufferSrcParameters) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}
func (s *AVBufferSrcParameters) Width() int {
	value := s.ptr.width
	return int(value)
}
func (s *AVBufferSrcParameters) Height() int {
	value := s.ptr.height
	return int(value)
}
func (s *AVBufferSrcParameters) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}
func (s *AVBufferSrcParameters) FrameRate() *AVRational {
	value := s.ptr.frame_rate
	return &AVRational{value: value}
}
func (s *AVBufferSrcParameters) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}
func (s *AVBufferSrcParameters) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}
func (s *AVBufferSrcParameters) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
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
type AVCodecTag struct {
	ptr *C.struct_AVCodecTag
}

func (s *AVCodecTag) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

// --- Struct AVProbeData ---

// AVProbeData wraps AVProbeData.
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

// Buf skipped due to prim ptr

func (s *AVProbeData) BufSize() int {
	value := s.ptr.buf_size
	return int(value)
}
func (s *AVProbeData) MimeType() *CStr {
	value := s.ptr.mime_type
	return wrapCStr(value)
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
func (s *AVOutputFormat) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}
func (s *AVOutputFormat) MimeType() *CStr {
	value := s.ptr.mime_type
	return wrapCStr(value)
}
func (s *AVOutputFormat) Extensions() *CStr {
	value := s.ptr.extensions
	return wrapCStr(value)
}
func (s *AVOutputFormat) AudioCodec() AVCodecID {
	value := s.ptr.audio_codec
	return AVCodecID(value)
}
func (s *AVOutputFormat) VideoCodec() AVCodecID {
	value := s.ptr.video_codec
	return AVCodecID(value)
}
func (s *AVOutputFormat) SubtitleCodec() AVCodecID {
	value := s.ptr.subtitle_codec
	return AVCodecID(value)
}
func (s *AVOutputFormat) Flags() int {
	value := s.ptr.flags
	return int(value)
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
func (s *AVInputFormat) LongName() *CStr {
	value := s.ptr.long_name
	return wrapCStr(value)
}
func (s *AVInputFormat) Flags() int {
	value := s.ptr.flags
	return int(value)
}
func (s *AVInputFormat) Extensions() *CStr {
	value := s.ptr.extensions
	return wrapCStr(value)
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
func (s *AVInputFormat) MimeType() *CStr {
	value := s.ptr.mime_type
	return wrapCStr(value)
}
func (s *AVInputFormat) RawCodecId() int {
	value := s.ptr.raw_codec_id
	return int(value)
}
func (s *AVInputFormat) PrivDataSize() int {
	value := s.ptr.priv_data_size
	return int(value)
}
func (s *AVInputFormat) FlagsInternal() int {
	value := s.ptr.flags_internal
	return int(value)
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
func (s *AVIndexEntry) Timestamp() int64 {
	value := s.ptr.timestamp
	return int64(value)
}

// Flags skipped due to bitfield

// Size skipped due to bitfield

func (s *AVIndexEntry) MinDistance() int {
	value := s.ptr.min_distance
	return int(value)
}

// --- Struct AVStream ---

// AVStream wraps AVStream.
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
func (s *AVStream) Index() int {
	value := s.ptr.index
	return int(value)
}
func (s *AVStream) Id() int {
	value := s.ptr.id
	return int(value)
}
func (s *AVStream) Codecpar() *AVCodecParameters {
	value := s.ptr.codecpar
	var valueMapped *AVCodecParameters
	if value != nil {
		valueMapped = &AVCodecParameters{ptr: value}
	}
	return valueMapped
}
func (s *AVStream) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}
func (s *AVStream) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}
func (s *AVStream) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}
func (s *AVStream) NbFrames() int64 {
	value := s.ptr.nb_frames
	return int64(value)
}
func (s *AVStream) Disposition() int {
	value := s.ptr.disposition
	return int(value)
}
func (s *AVStream) Discard() AVDiscard {
	value := s.ptr.discard
	return AVDiscard(value)
}
func (s *AVStream) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}
func (s *AVStream) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}
func (s *AVStream) AvgFrameRate() *AVRational {
	value := s.ptr.avg_frame_rate
	return &AVRational{value: value}
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
func (s *AVStream) NbSideData() int {
	value := s.ptr.nb_side_data
	return int(value)
}
func (s *AVStream) EventFlags() int {
	value := s.ptr.event_flags
	return int(value)
}
func (s *AVStream) RFrameRate() *AVRational {
	value := s.ptr.r_frame_rate
	return &AVRational{value: value}
}
func (s *AVStream) PtsWrapBits() int {
	value := s.ptr.pts_wrap_bits
	return int(value)
}

// --- Struct AVProgram ---

// AVProgram wraps AVProgram.
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
func (s *AVProgram) Flags() int {
	value := s.ptr.flags
	return int(value)
}
func (s *AVProgram) Discard() AVDiscard {
	value := s.ptr.discard
	return AVDiscard(value)
}

// StreamIndex skipped due to prim ptr

func (s *AVProgram) NbStreamIndexes() uint {
	value := s.ptr.nb_stream_indexes
	return uint(value)
}
func (s *AVProgram) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}
func (s *AVProgram) ProgramNum() int {
	value := s.ptr.program_num
	return int(value)
}
func (s *AVProgram) PmtPid() int {
	value := s.ptr.pmt_pid
	return int(value)
}
func (s *AVProgram) PcrPid() int {
	value := s.ptr.pcr_pid
	return int(value)
}
func (s *AVProgram) PmtVersion() int {
	value := s.ptr.pmt_version
	return int(value)
}
func (s *AVProgram) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}
func (s *AVProgram) EndTime() int64 {
	value := s.ptr.end_time
	return int64(value)
}
func (s *AVProgram) PtsWrapReference() int64 {
	value := s.ptr.pts_wrap_reference
	return int64(value)
}
func (s *AVProgram) PtsWrapBehavior() int {
	value := s.ptr.pts_wrap_behavior
	return int(value)
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
func (s *AVChapter) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}
func (s *AVChapter) Start() int64 {
	value := s.ptr.start
	return int64(value)
}
func (s *AVChapter) End() int64 {
	value := s.ptr.end
	return int64(value)
}
func (s *AVChapter) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}

// --- Struct AVFormatContext ---

// AVFormatContext wraps AVFormatContext.
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
func (s *AVFormatContext) Iformat() *AVInputFormat {
	value := s.ptr.iformat
	var valueMapped *AVInputFormat
	if value != nil {
		valueMapped = &AVInputFormat{ptr: value}
	}
	return valueMapped
}
func (s *AVFormatContext) Oformat() *AVOutputFormat {
	value := s.ptr.oformat
	var valueMapped *AVOutputFormat
	if value != nil {
		valueMapped = &AVOutputFormat{ptr: value}
	}
	return valueMapped
}
func (s *AVFormatContext) Pb() *AVIOContext {
	value := s.ptr.pb
	var valueMapped *AVIOContext
	if value != nil {
		valueMapped = &AVIOContext{ptr: value}
	}
	return valueMapped
}
func (s *AVFormatContext) CtxFlags() int {
	value := s.ptr.ctx_flags
	return int(value)
}
func (s *AVFormatContext) NbStreams() uint {
	value := s.ptr.nb_streams
	return uint(value)
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
func (s *AVFormatContext) StartTime() int64 {
	value := s.ptr.start_time
	return int64(value)
}
func (s *AVFormatContext) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
}
func (s *AVFormatContext) BitRate() int64 {
	value := s.ptr.bit_rate
	return int64(value)
}
func (s *AVFormatContext) PacketSize() uint {
	value := s.ptr.packet_size
	return uint(value)
}
func (s *AVFormatContext) MaxDelay() int {
	value := s.ptr.max_delay
	return int(value)
}
func (s *AVFormatContext) Flags() int {
	value := s.ptr.flags
	return int(value)
}
func (s *AVFormatContext) Probesize() int64 {
	value := s.ptr.probesize
	return int64(value)
}
func (s *AVFormatContext) MaxAnalyzeDuration() int64 {
	value := s.ptr.max_analyze_duration
	return int64(value)
}

// Key skipped due to ptr to uint8

func (s *AVFormatContext) Keylen() int {
	value := s.ptr.keylen
	return int(value)
}
func (s *AVFormatContext) NbPrograms() uint {
	value := s.ptr.nb_programs
	return uint(value)
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
func (s *AVFormatContext) AudioCodecId() AVCodecID {
	value := s.ptr.audio_codec_id
	return AVCodecID(value)
}
func (s *AVFormatContext) SubtitleCodecId() AVCodecID {
	value := s.ptr.subtitle_codec_id
	return AVCodecID(value)
}
func (s *AVFormatContext) MaxIndexSize() uint {
	value := s.ptr.max_index_size
	return uint(value)
}
func (s *AVFormatContext) MaxPictureBuffer() uint {
	value := s.ptr.max_picture_buffer
	return uint(value)
}
func (s *AVFormatContext) NbChapters() uint {
	value := s.ptr.nb_chapters
	return uint(value)
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
func (s *AVFormatContext) StartTimeRealtime() int64 {
	value := s.ptr.start_time_realtime
	return int64(value)
}
func (s *AVFormatContext) FpsProbeSize() int {
	value := s.ptr.fps_probe_size
	return int(value)
}
func (s *AVFormatContext) ErrorRecognition() int {
	value := s.ptr.error_recognition
	return int(value)
}

// InterruptCallback skipped due to ident byvalue

func (s *AVFormatContext) Debug() int {
	value := s.ptr.debug
	return int(value)
}
func (s *AVFormatContext) MaxInterleaveDelta() int64 {
	value := s.ptr.max_interleave_delta
	return int64(value)
}
func (s *AVFormatContext) StrictStdCompliance() int {
	value := s.ptr.strict_std_compliance
	return int(value)
}
func (s *AVFormatContext) EventFlags() int {
	value := s.ptr.event_flags
	return int(value)
}
func (s *AVFormatContext) MaxTsProbe() int {
	value := s.ptr.max_ts_probe
	return int(value)
}
func (s *AVFormatContext) AvoidNegativeTs() int {
	value := s.ptr.avoid_negative_ts
	return int(value)
}
func (s *AVFormatContext) TsId() int {
	value := s.ptr.ts_id
	return int(value)
}
func (s *AVFormatContext) AudioPreload() int {
	value := s.ptr.audio_preload
	return int(value)
}
func (s *AVFormatContext) MaxChunkDuration() int {
	value := s.ptr.max_chunk_duration
	return int(value)
}
func (s *AVFormatContext) MaxChunkSize() int {
	value := s.ptr.max_chunk_size
	return int(value)
}
func (s *AVFormatContext) UseWallclockAsTimestamps() int {
	value := s.ptr.use_wallclock_as_timestamps
	return int(value)
}
func (s *AVFormatContext) AvioFlags() int {
	value := s.ptr.avio_flags
	return int(value)
}
func (s *AVFormatContext) DurationEstimationMethod() AVDurationEstimationMethod {
	value := s.ptr.duration_estimation_method
	return AVDurationEstimationMethod(value)
}
func (s *AVFormatContext) SkipInitialBytes() int64 {
	value := s.ptr.skip_initial_bytes
	return int64(value)
}
func (s *AVFormatContext) CorrectTsOverflow() uint {
	value := s.ptr.correct_ts_overflow
	return uint(value)
}
func (s *AVFormatContext) Seek2Any() int {
	value := s.ptr.seek2any
	return int(value)
}
func (s *AVFormatContext) FlushPackets() int {
	value := s.ptr.flush_packets
	return int(value)
}
func (s *AVFormatContext) ProbeScore() int {
	value := s.ptr.probe_score
	return int(value)
}
func (s *AVFormatContext) FormatProbesize() int {
	value := s.ptr.format_probesize
	return int(value)
}
func (s *AVFormatContext) CodecWhitelist() *CStr {
	value := s.ptr.codec_whitelist
	return wrapCStr(value)
}
func (s *AVFormatContext) FormatWhitelist() *CStr {
	value := s.ptr.format_whitelist
	return wrapCStr(value)
}
func (s *AVFormatContext) IoRepositioned() int {
	value := s.ptr.io_repositioned
	return int(value)
}
func (s *AVFormatContext) VideoCodec() *AVCodec {
	value := s.ptr.video_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}
func (s *AVFormatContext) AudioCodec() *AVCodec {
	value := s.ptr.audio_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}
func (s *AVFormatContext) SubtitleCodec() *AVCodec {
	value := s.ptr.subtitle_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}
func (s *AVFormatContext) DataCodec() *AVCodec {
	value := s.ptr.data_codec
	var valueMapped *AVCodec
	if value != nil {
		valueMapped = &AVCodec{ptr: value}
	}
	return valueMapped
}
func (s *AVFormatContext) MetadataHeaderPadding() int {
	value := s.ptr.metadata_header_padding
	return int(value)
}

// ControlMessageCb skipped due to ident callback

func (s *AVFormatContext) OutputTsOffset() int64 {
	value := s.ptr.output_ts_offset
	return int64(value)
}

// DumpSeparator skipped due to ptr to uint8

func (s *AVFormatContext) DataCodecId() AVCodecID {
	value := s.ptr.data_codec_id
	return AVCodecID(value)
}
func (s *AVFormatContext) ProtocolWhitelist() *CStr {
	value := s.ptr.protocol_whitelist
	return wrapCStr(value)
}

// IoOpen skipped due to func ptr

// IoClose skipped due to func ptr

func (s *AVFormatContext) ProtocolBlacklist() *CStr {
	value := s.ptr.protocol_blacklist
	return wrapCStr(value)
}
func (s *AVFormatContext) MaxStreams() int {
	value := s.ptr.max_streams
	return int(value)
}
func (s *AVFormatContext) SkipEstimateDurationFromPts() int {
	value := s.ptr.skip_estimate_duration_from_pts
	return int(value)
}
func (s *AVFormatContext) MaxProbePackets() int {
	value := s.ptr.max_probe_packets
	return int(value)
}

// IoClose2 skipped due to func ptr

// --- Struct AVIOInterruptCB ---

// AVIOInterruptCB wraps AVIOInterruptCB.
type AVIOInterruptCB struct {
	ptr *C.AVIOInterruptCB
}

func (s *AVIOInterruptCB) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

// Callback skipped due to func ptr

// --- Struct AVIODirEntry ---

// AVIODirEntry wraps AVIODirEntry.
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
func (s *AVIODirEntry) Type() int {
	value := s.ptr._type
	return int(value)
}
func (s *AVIODirEntry) Utf8() int {
	value := s.ptr.utf8
	return int(value)
}
func (s *AVIODirEntry) Size() int64 {
	value := s.ptr.size
	return int64(value)
}
func (s *AVIODirEntry) ModificationTimestamp() int64 {
	value := s.ptr.modification_timestamp
	return int64(value)
}
func (s *AVIODirEntry) AccessTimestamp() int64 {
	value := s.ptr.access_timestamp
	return int64(value)
}
func (s *AVIODirEntry) StatusChangeTimestamp() int64 {
	value := s.ptr.status_change_timestamp
	return int64(value)
}
func (s *AVIODirEntry) UserId() int64 {
	value := s.ptr.user_id
	return int64(value)
}
func (s *AVIODirEntry) GroupId() int64 {
	value := s.ptr.group_id
	return int64(value)
}
func (s *AVIODirEntry) Filemode() int64 {
	value := s.ptr.filemode
	return int64(value)
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

// Buffer skipped due to prim ptr

func (s *AVIOContext) BufferSize() int {
	value := s.ptr.buffer_size
	return int(value)
}

// BufPtr skipped due to prim ptr

// BufEnd skipped due to prim ptr

// ReadPacket skipped due to func ptr

// WritePacket skipped due to func ptr

// Seek skipped due to func ptr

func (s *AVIOContext) Pos() int64 {
	value := s.ptr.pos
	return int64(value)
}
func (s *AVIOContext) EofReached() int {
	value := s.ptr.eof_reached
	return int(value)
}
func (s *AVIOContext) Error() int {
	value := s.ptr.error
	return int(value)
}
func (s *AVIOContext) WriteFlag() int {
	value := s.ptr.write_flag
	return int(value)
}
func (s *AVIOContext) MaxPacketSize() int {
	value := s.ptr.max_packet_size
	return int(value)
}
func (s *AVIOContext) MinPacketSize() int {
	value := s.ptr.min_packet_size
	return int(value)
}
func (s *AVIOContext) Checksum() uint32 {
	value := s.ptr.checksum
	return uint32(value)
}

// ChecksumPtr skipped due to prim ptr

// UpdateChecksum skipped due to func ptr

// ReadPause skipped due to func ptr

// ReadSeek skipped due to func ptr

func (s *AVIOContext) Seekable() int {
	value := s.ptr.seekable
	return int(value)
}
func (s *AVIOContext) Direct() int {
	value := s.ptr.direct
	return int(value)
}
func (s *AVIOContext) ProtocolWhitelist() *CStr {
	value := s.ptr.protocol_whitelist
	return wrapCStr(value)
}
func (s *AVIOContext) ProtocolBlacklist() *CStr {
	value := s.ptr.protocol_blacklist
	return wrapCStr(value)
}

// WriteDataType skipped due to func ptr

func (s *AVIOContext) IgnoreBoundaryPoint() int {
	value := s.ptr.ignore_boundary_point
	return int(value)
}

// BufPtrMax skipped due to prim ptr

func (s *AVIOContext) BytesRead() int64 {
	value := s.ptr.bytes_read
	return int64(value)
}
func (s *AVIOContext) BytesWritten() int64 {
	value := s.ptr.bytes_written
	return int64(value)
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

// Data skipped due to ptr to uint8

func (s *AVBufferRef) Size() uint64 {
	value := s.ptr.size
	return uint64(value)
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

// Name skipped due to const array

// --- Struct AVChannelLayout ---

// AVChannelLayout wraps AVChannelLayout.
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
func (s *AVChannelLayout) NbChannels() int {
	value := s.ptr.nb_channels
	return int(value)
}

// U skipped due to union type

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
func (s *AVDictionaryEntry) Value() *CStr {
	value := s.ptr.value
	return wrapCStr(value)
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

// Data skipped due to ptr to uint8

func (s *AVFrameSideData) Size() uint64 {
	value := s.ptr.size
	return uint64(value)
}
func (s *AVFrameSideData) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}
func (s *AVFrameSideData) Buf() *AVBufferRef {
	value := s.ptr.buf
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// --- Struct AVRegionOfInterest ---

// AVRegionOfInterest wraps AVRegionOfInterest.
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
func (s *AVRegionOfInterest) Top() int {
	value := s.ptr.top
	return int(value)
}
func (s *AVRegionOfInterest) Bottom() int {
	value := s.ptr.bottom
	return int(value)
}
func (s *AVRegionOfInterest) Left() int {
	value := s.ptr.left
	return int(value)
}
func (s *AVRegionOfInterest) Right() int {
	value := s.ptr.right
	return int(value)
}
func (s *AVRegionOfInterest) Qoffset() *AVRational {
	value := s.ptr.qoffset
	return &AVRational{value: value}
}

// --- Struct AVFrame ---

// AVFrame wraps AVFrame.
type AVFrame struct {
	ptr *C.AVFrame
}

func (s *AVFrame) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

// Data skipped due to const array

// Linesize skipped due to const array

// ExtendedData skipped due to unknown ptr ptr

func (s *AVFrame) Width() int {
	value := s.ptr.width
	return int(value)
}
func (s *AVFrame) Height() int {
	value := s.ptr.height
	return int(value)
}
func (s *AVFrame) NbSamples() int {
	value := s.ptr.nb_samples
	return int(value)
}
func (s *AVFrame) Format() int {
	value := s.ptr.format
	return int(value)
}
func (s *AVFrame) KeyFrame() int {
	value := s.ptr.key_frame
	return int(value)
}
func (s *AVFrame) PictType() AVPictureType {
	value := s.ptr.pict_type
	return AVPictureType(value)
}
func (s *AVFrame) SampleAspectRatio() *AVRational {
	value := s.ptr.sample_aspect_ratio
	return &AVRational{value: value}
}
func (s *AVFrame) Pts() int64 {
	value := s.ptr.pts
	return int64(value)
}
func (s *AVFrame) PktDts() int64 {
	value := s.ptr.pkt_dts
	return int64(value)
}
func (s *AVFrame) TimeBase() *AVRational {
	value := s.ptr.time_base
	return &AVRational{value: value}
}
func (s *AVFrame) CodedPictureNumber() int {
	value := s.ptr.coded_picture_number
	return int(value)
}
func (s *AVFrame) DisplayPictureNumber() int {
	value := s.ptr.display_picture_number
	return int(value)
}
func (s *AVFrame) Quality() int {
	value := s.ptr.quality
	return int(value)
}
func (s *AVFrame) RepeatPict() int {
	value := s.ptr.repeat_pict
	return int(value)
}
func (s *AVFrame) InterlacedFrame() int {
	value := s.ptr.interlaced_frame
	return int(value)
}
func (s *AVFrame) TopFieldFirst() int {
	value := s.ptr.top_field_first
	return int(value)
}
func (s *AVFrame) PaletteHasChanged() int {
	value := s.ptr.palette_has_changed
	return int(value)
}
func (s *AVFrame) ReorderedOpaque() int64 {
	value := s.ptr.reordered_opaque
	return int64(value)
}
func (s *AVFrame) SampleRate() int {
	value := s.ptr.sample_rate
	return int(value)
}
func (s *AVFrame) ChannelLayout() uint64 {
	value := s.ptr.channel_layout
	return uint64(value)
}

// Buf skipped due to const array

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
func (s *AVFrame) Flags() int {
	value := s.ptr.flags
	return int(value)
}
func (s *AVFrame) ColorRange() AVColorRange {
	value := s.ptr.color_range
	return AVColorRange(value)
}
func (s *AVFrame) ColorPrimaries() AVColorPrimaries {
	value := s.ptr.color_primaries
	return AVColorPrimaries(value)
}
func (s *AVFrame) ColorTrc() AVColorTransferCharacteristic {
	value := s.ptr.color_trc
	return AVColorTransferCharacteristic(value)
}
func (s *AVFrame) Colorspace() AVColorSpace {
	value := s.ptr.colorspace
	return AVColorSpace(value)
}
func (s *AVFrame) ChromaLocation() AVChromaLocation {
	value := s.ptr.chroma_location
	return AVChromaLocation(value)
}
func (s *AVFrame) BestEffortTimestamp() int64 {
	value := s.ptr.best_effort_timestamp
	return int64(value)
}
func (s *AVFrame) PktPos() int64 {
	value := s.ptr.pkt_pos
	return int64(value)
}
func (s *AVFrame) PktDuration() int64 {
	value := s.ptr.pkt_duration
	return int64(value)
}
func (s *AVFrame) Metadata() *AVDictionary {
	value := s.ptr.metadata
	var valueMapped *AVDictionary
	if value != nil {
		valueMapped = &AVDictionary{ptr: value}
	}
	return valueMapped
}
func (s *AVFrame) DecodeErrorFlags() int {
	value := s.ptr.decode_error_flags
	return int(value)
}
func (s *AVFrame) Channels() int {
	value := s.ptr.channels
	return int(value)
}
func (s *AVFrame) PktSize() int {
	value := s.ptr.pkt_size
	return int(value)
}
func (s *AVFrame) HwFramesCtx() *AVBufferRef {
	value := s.ptr.hw_frames_ctx
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}
func (s *AVFrame) OpaqueRef() *AVBufferRef {
	value := s.ptr.opaque_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}
func (s *AVFrame) CropTop() uint64 {
	value := s.ptr.crop_top
	return uint64(value)
}
func (s *AVFrame) CropBottom() uint64 {
	value := s.ptr.crop_bottom
	return uint64(value)
}
func (s *AVFrame) CropLeft() uint64 {
	value := s.ptr.crop_left
	return uint64(value)
}
func (s *AVFrame) CropRight() uint64 {
	value := s.ptr.crop_right
	return uint64(value)
}
func (s *AVFrame) PrivateRef() *AVBufferRef {
	value := s.ptr.private_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}

// ChLayout skipped due to ident byvalue

func (s *AVFrame) Duration() int64 {
	value := s.ptr.duration
	return int64(value)
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
func (s *AVHWDeviceContext) Internal() *AVHWDeviceInternal {
	value := s.ptr.internal
	var valueMapped *AVHWDeviceInternal
	if value != nil {
		valueMapped = &AVHWDeviceInternal{ptr: value}
	}
	return valueMapped
}
func (s *AVHWDeviceContext) Type() AVHWDeviceType {
	value := s.ptr._type
	return AVHWDeviceType(value)
}

// Free skipped due to func ptr

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
func (s *AVHWFramesContext) Internal() *AVHWFramesInternal {
	value := s.ptr.internal
	var valueMapped *AVHWFramesInternal
	if value != nil {
		valueMapped = &AVHWFramesInternal{ptr: value}
	}
	return valueMapped
}
func (s *AVHWFramesContext) DeviceRef() *AVBufferRef {
	value := s.ptr.device_ref
	var valueMapped *AVBufferRef
	if value != nil {
		valueMapped = &AVBufferRef{ptr: value}
	}
	return valueMapped
}
func (s *AVHWFramesContext) DeviceCtx() *AVHWDeviceContext {
	value := s.ptr.device_ctx
	var valueMapped *AVHWDeviceContext
	if value != nil {
		valueMapped = &AVHWDeviceContext{ptr: value}
	}
	return valueMapped
}

// Free skipped due to func ptr

func (s *AVHWFramesContext) Pool() *AVBufferPool {
	value := s.ptr.pool
	var valueMapped *AVBufferPool
	if value != nil {
		valueMapped = &AVBufferPool{ptr: value}
	}
	return valueMapped
}
func (s *AVHWFramesContext) InitialPoolSize() int {
	value := s.ptr.initial_pool_size
	return int(value)
}
func (s *AVHWFramesContext) Format() AVPixelFormat {
	value := s.ptr.format
	return AVPixelFormat(value)
}
func (s *AVHWFramesContext) SwFormat() AVPixelFormat {
	value := s.ptr.sw_format
	return AVPixelFormat(value)
}
func (s *AVHWFramesContext) Width() int {
	value := s.ptr.width
	return int(value)
}
func (s *AVHWFramesContext) Height() int {
	value := s.ptr.height
	return int(value)
}

// --- Struct AVHWFramesConstraints ---

// AVHWFramesConstraints wraps AVHWFramesConstraints.
type AVHWFramesConstraints struct {
	ptr *C.AVHWFramesConstraints
}

func (s *AVHWFramesConstraints) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

// ValidHwFormats skipped due to enum ptr

// ValidSwFormats skipped due to enum ptr

func (s *AVHWFramesConstraints) MinWidth() int {
	value := s.ptr.min_width
	return int(value)
}
func (s *AVHWFramesConstraints) MinHeight() int {
	value := s.ptr.min_height
	return int(value)
}
func (s *AVHWFramesConstraints) MaxWidth() int {
	value := s.ptr.max_width
	return int(value)
}
func (s *AVHWFramesConstraints) MaxHeight() int {
	value := s.ptr.max_height
	return int(value)
}

// --- Struct AVClass ---

// AVClass wraps AVClass.
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

// ItemName skipped due to func ptr

func (s *AVClass) Option() *AVOption {
	value := s.ptr.option
	var valueMapped *AVOption
	if value != nil {
		valueMapped = &AVOption{ptr: value}
	}
	return valueMapped
}
func (s *AVClass) Version() int {
	value := s.ptr.version
	return int(value)
}
func (s *AVClass) LogLevelOffsetOffset() int {
	value := s.ptr.log_level_offset_offset
	return int(value)
}
func (s *AVClass) ParentLogContextOffset() int {
	value := s.ptr.parent_log_context_offset
	return int(value)
}
func (s *AVClass) Category() AVClassCategory {
	value := s.ptr.category
	return AVClassCategory(value)
}

// GetCategory skipped due to func ptr

// QueryRanges skipped due to func ptr

// ChildNext skipped due to func ptr

// ChildClassIterate skipped due to func ptr

// --- Struct AVOption ---

// AVOption wraps AVOption.
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
func (s *AVOption) Help() *CStr {
	value := s.ptr.help
	return wrapCStr(value)
}
func (s *AVOption) Offset() int {
	value := s.ptr.offset
	return int(value)
}
func (s *AVOption) Type() AVOptionType {
	value := s.ptr._type
	return AVOptionType(value)
}

// DefaultVal skipped due to union type

func (s *AVOption) Min() float64 {
	value := s.ptr.min
	return float64(value)
}
func (s *AVOption) Max() float64 {
	value := s.ptr.max
	return float64(value)
}
func (s *AVOption) Flags() int {
	value := s.ptr.flags
	return int(value)
}
func (s *AVOption) Unit() *CStr {
	value := s.ptr.unit
	return wrapCStr(value)
}

// --- Struct AVOptionRange ---

// AVOptionRange wraps AVOptionRange.
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
func (s *AVOptionRange) ValueMin() float64 {
	value := s.ptr.value_min
	return float64(value)
}
func (s *AVOptionRange) ValueMax() float64 {
	value := s.ptr.value_max
	return float64(value)
}
func (s *AVOptionRange) ComponentMin() float64 {
	value := s.ptr.component_min
	return float64(value)
}
func (s *AVOptionRange) ComponentMax() float64 {
	value := s.ptr.component_max
	return float64(value)
}
func (s *AVOptionRange) IsRange() int {
	value := s.ptr.is_range
	return int(value)
}

// --- Struct AVOptionRanges ---

// AVOptionRanges wraps AVOptionRanges.
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
func (s *AVOptionRanges) NbComponents() int {
	value := s.ptr.nb_components
	return int(value)
}

// --- Struct AVRational ---

// AVRational wraps AVRational.
type AVRational struct {
	value C.AVRational
}

func (s *AVRational) Num() int {
	value := s.value.num
	return int(value)
}
func (s *AVRational) Den() int {
	value := s.value.den
	return int(value)
}
