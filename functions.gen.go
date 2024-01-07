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

// --- Function avcodec_version ---

// AVCodecVersion wraps avcodec_version.
func AVCodecVersion() uint {
	ret := C.avcodec_version()
	return uint(ret)
}

// --- Function avcodec_configuration ---

// AVCodecConfiguration wraps avcodec_configuration.
func AVCodecConfiguration() *CStr {
	ret := C.avcodec_configuration()
	return wrapCStr(ret)
}

// --- Function avcodec_license ---

// AVCodecLicense wraps avcodec_license.
func AVCodecLicense() *CStr {
	ret := C.avcodec_license()
	return wrapCStr(ret)
}

// --- Function avcodec_alloc_context3 ---

// AVCodecAllocContext3 wraps avcodec_alloc_context3.
func AVCodecAllocContext3(codec *AVCodec) *AVCodecContext {
	var tmpcodec *C.AVCodec
	if codec != nil {
		tmpcodec = codec.ptr
	}
	ret := C.avcodec_alloc_context3(tmpcodec)
	var retMapped *AVCodecContext
	if ret != nil {
		retMapped = &AVCodecContext{ptr: ret}
	}
	return retMapped
}

// --- Function avcodec_free_context ---

// AVCodecFreeContext wraps avcodec_free_context.
func AVCodecFreeContext(avctx **AVCodecContext) {
	var ptravctx **C.AVCodecContext
	var tmpavctx *C.AVCodecContext
	var oldTmpavctx *C.AVCodecContext
	if avctx != nil {
		inneravctx := *avctx
		if inneravctx != nil {
			tmpavctx = inneravctx.ptr
			oldTmpavctx = tmpavctx
		}
		ptravctx = &tmpavctx
	}
	C.avcodec_free_context(ptravctx)
	if tmpavctx != oldTmpavctx && avctx != nil {
		if tmpavctx != nil {
			*avctx = &AVCodecContext{ptr: tmpavctx}
		} else {
			*avctx = nil
		}
	}
}

// --- Function avcodec_get_class ---

// AVCodecGetClass wraps avcodec_get_class.
func AVCodecGetClass() *AVClass {
	ret := C.avcodec_get_class()
	var retMapped *AVClass
	if ret != nil {
		retMapped = &AVClass{ptr: ret}
	}
	return retMapped
}

// --- Function avcodec_get_subtitle_rect_class ---

// AVCodecGetSubtitleRectClass wraps avcodec_get_subtitle_rect_class.
func AVCodecGetSubtitleRectClass() *AVClass {
	ret := C.avcodec_get_subtitle_rect_class()
	var retMapped *AVClass
	if ret != nil {
		retMapped = &AVClass{ptr: ret}
	}
	return retMapped
}

// --- Function avcodec_parameters_from_context ---

// AVCodecParametersFromContext wraps avcodec_parameters_from_context.
func AVCodecParametersFromContext(par *AVCodecParameters, codec *AVCodecContext) int {
	var tmppar *C.AVCodecParameters
	if par != nil {
		tmppar = par.ptr
	}
	var tmpcodec *C.AVCodecContext
	if codec != nil {
		tmpcodec = codec.ptr
	}
	ret := C.avcodec_parameters_from_context(tmppar, tmpcodec)
	return int(ret)
}

// --- Function avcodec_parameters_to_context ---

// AVCodecParametersToContext wraps avcodec_parameters_to_context.
func AVCodecParametersToContext(codec *AVCodecContext, par *AVCodecParameters) int {
	var tmpcodec *C.AVCodecContext
	if codec != nil {
		tmpcodec = codec.ptr
	}
	var tmppar *C.AVCodecParameters
	if par != nil {
		tmppar = par.ptr
	}
	ret := C.avcodec_parameters_to_context(tmpcodec, tmppar)
	return int(ret)
}

// --- Function avcodec_open2 ---

// AVCodecOpen2 wraps avcodec_open2.
func AVCodecOpen2(avctx *AVCodecContext, codec *AVCodec, options **AVDictionary) int {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpcodec *C.AVCodec
	if codec != nil {
		tmpcodec = codec.ptr
	}
	var ptroptions **C.AVDictionary
	var tmpoptions *C.AVDictionary
	var oldTmpoptions *C.AVDictionary
	if options != nil {
		inneroptions := *options
		if inneroptions != nil {
			tmpoptions = inneroptions.ptr
			oldTmpoptions = tmpoptions
		}
		ptroptions = &tmpoptions
	}
	ret := C.avcodec_open2(tmpavctx, tmpcodec, ptroptions)
	if tmpoptions != oldTmpoptions && options != nil {
		if tmpoptions != nil {
			*options = &AVDictionary{ptr: tmpoptions}
		} else {
			*options = nil
		}
	}
	return int(ret)
}

// --- Function avcodec_close ---

// AVCodecClose wraps avcodec_close.
func AVCodecClose(avctx *AVCodecContext) int {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	ret := C.avcodec_close(tmpavctx)
	return int(ret)
}

// --- Function avsubtitle_free ---

// AVSubtitleFree wraps avsubtitle_free.
func AVSubtitleFree(sub *AVSubtitle) {
	var tmpsub *C.AVSubtitle
	if sub != nil {
		tmpsub = sub.ptr
	}
	C.avsubtitle_free(tmpsub)
}

// --- Function avcodec_default_get_buffer2 ---

// AVCodecDefaultGetBuffer2 wraps avcodec_default_get_buffer2.
func AVCodecDefaultGetBuffer2(s *AVCodecContext, frame *AVFrame, flags int) int {
	var tmps *C.AVCodecContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.avcodec_default_get_buffer2(tmps, tmpframe, C.int(flags))
	return int(ret)
}

// --- Function avcodec_default_get_encode_buffer ---

// AVCodecDefaultGetEncodeBuffer wraps avcodec_default_get_encode_buffer.
func AVCodecDefaultGetEncodeBuffer(s *AVCodecContext, pkt *AVPacket, flags int) int {
	var tmps *C.AVCodecContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.avcodec_default_get_encode_buffer(tmps, tmppkt, C.int(flags))
	return int(ret)
}

// --- Function avcodec_align_dimensions ---

// AVCodecAlignDimensions wraps avcodec_align_dimensions.
// avcodec_align_dimensions skipped due to width

// --- Function avcodec_align_dimensions2 ---

// AVCodecAlignDimensions2 wraps avcodec_align_dimensions2.
// avcodec_align_dimensions2 skipped due to width

// --- Function avcodec_enum_to_chroma_pos ---

// AVCodecEnumToChromaPos wraps avcodec_enum_to_chroma_pos.
// avcodec_enum_to_chroma_pos skipped due to xpos

// --- Function avcodec_chroma_pos_to_enum ---

// AVCodecChromaPosToEnum wraps avcodec_chroma_pos_to_enum.
func AVCodecChromaPosToEnum(xpos int, ypos int) AVChromaLocation {
	ret := C.avcodec_chroma_pos_to_enum(C.int(xpos), C.int(ypos))
	return AVChromaLocation(ret)
}

// --- Function avcodec_decode_subtitle2 ---

// AVCodecDecodeSubtitle2 wraps avcodec_decode_subtitle2.
// avcodec_decode_subtitle2 skipped due to gotSubPtr

// --- Function avcodec_send_packet ---

// AVCodecSendPacket wraps avcodec_send_packet.
func AVCodecSendPacket(avctx *AVCodecContext, avpkt *AVPacket) int {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpavpkt *C.AVPacket
	if avpkt != nil {
		tmpavpkt = avpkt.ptr
	}
	ret := C.avcodec_send_packet(tmpavctx, tmpavpkt)
	return int(ret)
}

// --- Function avcodec_receive_frame ---

// AVCodecReceiveFrame wraps avcodec_receive_frame.
func AVCodecReceiveFrame(avctx *AVCodecContext, frame *AVFrame) int {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.avcodec_receive_frame(tmpavctx, tmpframe)
	return int(ret)
}

// --- Function avcodec_send_frame ---

// AVCodecSendFrame wraps avcodec_send_frame.
func AVCodecSendFrame(avctx *AVCodecContext, frame *AVFrame) int {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.avcodec_send_frame(tmpavctx, tmpframe)
	return int(ret)
}

// --- Function avcodec_receive_packet ---

// AVCodecReceivePacket wraps avcodec_receive_packet.
func AVCodecReceivePacket(avctx *AVCodecContext, avpkt *AVPacket) int {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpavpkt *C.AVPacket
	if avpkt != nil {
		tmpavpkt = avpkt.ptr
	}
	ret := C.avcodec_receive_packet(tmpavctx, tmpavpkt)
	return int(ret)
}

// --- Function avcodec_get_hw_frames_parameters ---

// AVCodecGetHwFramesParameters wraps avcodec_get_hw_frames_parameters.
func AVCodecGetHwFramesParameters(avctx *AVCodecContext, deviceRef *AVBufferRef, hwPixFmt AVPixelFormat, outFramesRef **AVBufferRef) int {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpdeviceRef *C.AVBufferRef
	if deviceRef != nil {
		tmpdeviceRef = deviceRef.ptr
	}
	var ptroutFramesRef **C.AVBufferRef
	var tmpoutFramesRef *C.AVBufferRef
	var oldTmpoutFramesRef *C.AVBufferRef
	if outFramesRef != nil {
		inneroutFramesRef := *outFramesRef
		if inneroutFramesRef != nil {
			tmpoutFramesRef = inneroutFramesRef.ptr
			oldTmpoutFramesRef = tmpoutFramesRef
		}
		ptroutFramesRef = &tmpoutFramesRef
	}
	ret := C.avcodec_get_hw_frames_parameters(tmpavctx, tmpdeviceRef, C.enum_AVPixelFormat(hwPixFmt), ptroutFramesRef)
	if tmpoutFramesRef != oldTmpoutFramesRef && outFramesRef != nil {
		if tmpoutFramesRef != nil {
			*outFramesRef = &AVBufferRef{ptr: tmpoutFramesRef}
		} else {
			*outFramesRef = nil
		}
	}
	return int(ret)
}

// --- Function av_parser_iterate ---

// AVParserIterate wraps av_parser_iterate.
// av_parser_iterate skipped due to opaque

// --- Function av_parser_init ---

// AVParserInit wraps av_parser_init.
func AVParserInit(codecId int) *AVCodecParserContext {
	ret := C.av_parser_init(C.int(codecId))
	var retMapped *AVCodecParserContext
	if ret != nil {
		retMapped = &AVCodecParserContext{ptr: ret}
	}
	return retMapped
}

// --- Function av_parser_parse2 ---

// AVParserParse2 wraps av_parser_parse2.
// av_parser_parse2 skipped due to poutbuf

// --- Function av_parser_close ---

// AVParserClose wraps av_parser_close.
func AVParserClose(s *AVCodecParserContext) {
	var tmps *C.AVCodecParserContext
	if s != nil {
		tmps = s.ptr
	}
	C.av_parser_close(tmps)
}

// --- Function avcodec_encode_subtitle ---

// AVCodecEncodeSubtitle wraps avcodec_encode_subtitle.
func AVCodecEncodeSubtitle(avctx *AVCodecContext, buf unsafe.Pointer, bufSize int, sub *AVSubtitle) int {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpsub *C.AVSubtitle
	if sub != nil {
		tmpsub = sub.ptr
	}
	ret := C.avcodec_encode_subtitle(tmpavctx, (*C.uint8_t)(buf), C.int(bufSize), tmpsub)
	return int(ret)
}

// --- Function avcodec_pix_fmt_to_codec_tag ---

// AVCodecPixFmtToCodecTag wraps avcodec_pix_fmt_to_codec_tag.
func AVCodecPixFmtToCodecTag(pixFmt AVPixelFormat) uint {
	ret := C.avcodec_pix_fmt_to_codec_tag(C.enum_AVPixelFormat(pixFmt))
	return uint(ret)
}

// --- Function avcodec_find_best_pix_fmt_of_list ---

// AVCodecFindBestPixFmtOfList wraps avcodec_find_best_pix_fmt_of_list.
// avcodec_find_best_pix_fmt_of_list skipped due to pixFmtList

// --- Function avcodec_default_get_format ---

// AVCodecDefaultGetFormat wraps avcodec_default_get_format.
// avcodec_default_get_format skipped due to fmt

// --- Function avcodec_string ---

// AVCodecString wraps avcodec_string.
func AVCodecString(buf *CStr, bufSize int, enc *AVCodecContext, encode int) {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	var tmpenc *C.AVCodecContext
	if enc != nil {
		tmpenc = enc.ptr
	}
	C.avcodec_string(tmpbuf, C.int(bufSize), tmpenc, C.int(encode))
}

// --- Function avcodec_default_execute ---

// avcodec_default_execute skipped due to func.

// --- Function avcodec_default_execute2 ---

// avcodec_default_execute2 skipped due to func.

// --- Function avcodec_fill_audio_frame ---

// AVCodecFillAudioFrame wraps avcodec_fill_audio_frame.
func AVCodecFillAudioFrame(frame *AVFrame, nbChannels int, sampleFmt AVSampleFormat, buf unsafe.Pointer, bufSize int, align int) int {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.avcodec_fill_audio_frame(tmpframe, C.int(nbChannels), C.enum_AVSampleFormat(sampleFmt), (*C.uint8_t)(buf), C.int(bufSize), C.int(align))
	return int(ret)
}

// --- Function avcodec_flush_buffers ---

// AVCodecFlushBuffers wraps avcodec_flush_buffers.
func AVCodecFlushBuffers(avctx *AVCodecContext) {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	C.avcodec_flush_buffers(tmpavctx)
}

// --- Function av_get_audio_frame_duration ---

// AVGetAudioFrameDuration wraps av_get_audio_frame_duration.
func AVGetAudioFrameDuration(avctx *AVCodecContext, frameBytes int) int {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	ret := C.av_get_audio_frame_duration(tmpavctx, C.int(frameBytes))
	return int(ret)
}

// --- Function av_fast_padded_malloc ---

// AVFastPaddedMalloc wraps av_fast_padded_malloc.
// av_fast_padded_malloc skipped due to size

// --- Function av_fast_padded_mallocz ---

// AVFastPaddedMallocz wraps av_fast_padded_mallocz.
// av_fast_padded_mallocz skipped due to size

// --- Function avcodec_is_open ---

// AVCodecIsOpen wraps avcodec_is_open.
func AVCodecIsOpen(s *AVCodecContext) int {
	var tmps *C.AVCodecContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avcodec_is_open(tmps)
	return int(ret)
}

// --- Function av_codec_iterate ---

// AVCodecIterate wraps av_codec_iterate.
// av_codec_iterate skipped due to opaque

// --- Function avcodec_find_decoder ---

// AVCodecFindDecoder wraps avcodec_find_decoder.
func AVCodecFindDecoder(id AVCodecID) *AVCodec {
	ret := C.avcodec_find_decoder(C.enum_AVCodecID(id))
	var retMapped *AVCodec
	if ret != nil {
		retMapped = &AVCodec{ptr: ret}
	}
	return retMapped
}

// --- Function avcodec_find_decoder_by_name ---

// AVCodecFindDecoderByName wraps avcodec_find_decoder_by_name.
func AVCodecFindDecoderByName(name *CStr) *AVCodec {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.avcodec_find_decoder_by_name(tmpname)
	var retMapped *AVCodec
	if ret != nil {
		retMapped = &AVCodec{ptr: ret}
	}
	return retMapped
}

// --- Function avcodec_find_encoder ---

// AVCodecFindEncoder wraps avcodec_find_encoder.
func AVCodecFindEncoder(id AVCodecID) *AVCodec {
	ret := C.avcodec_find_encoder(C.enum_AVCodecID(id))
	var retMapped *AVCodec
	if ret != nil {
		retMapped = &AVCodec{ptr: ret}
	}
	return retMapped
}

// --- Function avcodec_find_encoder_by_name ---

// AVCodecFindEncoderByName wraps avcodec_find_encoder_by_name.
func AVCodecFindEncoderByName(name *CStr) *AVCodec {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.avcodec_find_encoder_by_name(tmpname)
	var retMapped *AVCodec
	if ret != nil {
		retMapped = &AVCodec{ptr: ret}
	}
	return retMapped
}

// --- Function av_codec_is_encoder ---

// AVCodecIsEncoder wraps av_codec_is_encoder.
func AVCodecIsEncoder(codec *AVCodec) int {
	var tmpcodec *C.AVCodec
	if codec != nil {
		tmpcodec = codec.ptr
	}
	ret := C.av_codec_is_encoder(tmpcodec)
	return int(ret)
}

// --- Function av_codec_is_decoder ---

// AVCodecIsDecoder wraps av_codec_is_decoder.
func AVCodecIsDecoder(codec *AVCodec) int {
	var tmpcodec *C.AVCodec
	if codec != nil {
		tmpcodec = codec.ptr
	}
	ret := C.av_codec_is_decoder(tmpcodec)
	return int(ret)
}

// --- Function av_get_profile_name ---

// AVGetProfileName wraps av_get_profile_name.
func AVGetProfileName(codec *AVCodec, profile int) *CStr {
	var tmpcodec *C.AVCodec
	if codec != nil {
		tmpcodec = codec.ptr
	}
	ret := C.av_get_profile_name(tmpcodec, C.int(profile))
	return wrapCStr(ret)
}

// --- Function avcodec_get_hw_config ---

// AVCodecGetHwConfig wraps avcodec_get_hw_config.
func AVCodecGetHwConfig(codec *AVCodec, index int) *AVCodecHWConfig {
	var tmpcodec *C.AVCodec
	if codec != nil {
		tmpcodec = codec.ptr
	}
	ret := C.avcodec_get_hw_config(tmpcodec, C.int(index))
	var retMapped *AVCodecHWConfig
	if ret != nil {
		retMapped = &AVCodecHWConfig{ptr: ret}
	}
	return retMapped
}

// --- Function avcodec_descriptor_get ---

// AVCodecDescriptorGet wraps avcodec_descriptor_get.
func AVCodecDescriptorGet(id AVCodecID) *AVCodecDescriptor {
	ret := C.avcodec_descriptor_get(C.enum_AVCodecID(id))
	var retMapped *AVCodecDescriptor
	if ret != nil {
		retMapped = &AVCodecDescriptor{ptr: ret}
	}
	return retMapped
}

// --- Function avcodec_descriptor_next ---

// AVCodecDescriptorNext wraps avcodec_descriptor_next.
func AVCodecDescriptorNext(prev *AVCodecDescriptor) *AVCodecDescriptor {
	var tmpprev *C.AVCodecDescriptor
	if prev != nil {
		tmpprev = prev.ptr
	}
	ret := C.avcodec_descriptor_next(tmpprev)
	var retMapped *AVCodecDescriptor
	if ret != nil {
		retMapped = &AVCodecDescriptor{ptr: ret}
	}
	return retMapped
}

// --- Function avcodec_descriptor_get_by_name ---

// AVCodecDescriptorGetByName wraps avcodec_descriptor_get_by_name.
func AVCodecDescriptorGetByName(name *CStr) *AVCodecDescriptor {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.avcodec_descriptor_get_by_name(tmpname)
	var retMapped *AVCodecDescriptor
	if ret != nil {
		retMapped = &AVCodecDescriptor{ptr: ret}
	}
	return retMapped
}

// --- Function avcodec_get_type ---

// AVCodecGetType wraps avcodec_get_type.
func AVCodecGetType(codecId AVCodecID) AVMediaType {
	ret := C.avcodec_get_type(C.enum_AVCodecID(codecId))
	return AVMediaType(ret)
}

// --- Function avcodec_get_name ---

// AVCodecGetName wraps avcodec_get_name.
func AVCodecGetName(id AVCodecID) *CStr {
	ret := C.avcodec_get_name(C.enum_AVCodecID(id))
	return wrapCStr(ret)
}

// --- Function av_get_bits_per_sample ---

// AVGetBitsPerSample wraps av_get_bits_per_sample.
func AVGetBitsPerSample(codecId AVCodecID) int {
	ret := C.av_get_bits_per_sample(C.enum_AVCodecID(codecId))
	return int(ret)
}

// --- Function av_get_exact_bits_per_sample ---

// AVGetExactBitsPerSample wraps av_get_exact_bits_per_sample.
func AVGetExactBitsPerSample(codecId AVCodecID) int {
	ret := C.av_get_exact_bits_per_sample(C.enum_AVCodecID(codecId))
	return int(ret)
}

// --- Function avcodec_profile_name ---

// AVCodecProfileName wraps avcodec_profile_name.
func AVCodecProfileName(codecId AVCodecID, profile int) *CStr {
	ret := C.avcodec_profile_name(C.enum_AVCodecID(codecId), C.int(profile))
	return wrapCStr(ret)
}

// --- Function av_get_pcm_codec ---

// AVGetPcmCodec wraps av_get_pcm_codec.
func AVGetPcmCodec(fmt AVSampleFormat, be int) AVCodecID {
	ret := C.av_get_pcm_codec(C.enum_AVSampleFormat(fmt), C.int(be))
	return AVCodecID(ret)
}

// --- Function avcodec_parameters_alloc ---

// AVCodecParametersAlloc wraps avcodec_parameters_alloc.
func AVCodecParametersAlloc() *AVCodecParameters {
	ret := C.avcodec_parameters_alloc()
	var retMapped *AVCodecParameters
	if ret != nil {
		retMapped = &AVCodecParameters{ptr: ret}
	}
	return retMapped
}

// --- Function avcodec_parameters_free ---

// AVCodecParametersFree wraps avcodec_parameters_free.
func AVCodecParametersFree(par **AVCodecParameters) {
	var ptrpar **C.AVCodecParameters
	var tmppar *C.AVCodecParameters
	var oldTmppar *C.AVCodecParameters
	if par != nil {
		innerpar := *par
		if innerpar != nil {
			tmppar = innerpar.ptr
			oldTmppar = tmppar
		}
		ptrpar = &tmppar
	}
	C.avcodec_parameters_free(ptrpar)
	if tmppar != oldTmppar && par != nil {
		if tmppar != nil {
			*par = &AVCodecParameters{ptr: tmppar}
		} else {
			*par = nil
		}
	}
}

// --- Function avcodec_parameters_copy ---

// AVCodecParametersCopy wraps avcodec_parameters_copy.
func AVCodecParametersCopy(dst *AVCodecParameters, src *AVCodecParameters) int {
	var tmpdst *C.AVCodecParameters
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVCodecParameters
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.avcodec_parameters_copy(tmpdst, tmpsrc)
	return int(ret)
}

// --- Function av_get_audio_frame_duration2 ---

// AVGetAudioFrameDuration2 wraps av_get_audio_frame_duration2.
func AVGetAudioFrameDuration2(par *AVCodecParameters, frameBytes int) int {
	var tmppar *C.AVCodecParameters
	if par != nil {
		tmppar = par.ptr
	}
	ret := C.av_get_audio_frame_duration2(tmppar, C.int(frameBytes))
	return int(ret)
}

// --- Function av_cpb_properties_alloc ---

// AVCpbPropertiesAlloc wraps av_cpb_properties_alloc.
// av_cpb_properties_alloc skipped due to size

// --- Function av_xiphlacing ---

// AVXiphlacing wraps av_xiphlacing.
func AVXiphlacing(s unsafe.Pointer, v uint) uint {
	ret := C.av_xiphlacing((*C.uchar)(s), C.uint(v))
	return uint(ret)
}

// --- Function av_packet_alloc ---

// AVPacketAlloc wraps av_packet_alloc.
func AVPacketAlloc() *AVPacket {
	ret := C.av_packet_alloc()
	var retMapped *AVPacket
	if ret != nil {
		retMapped = &AVPacket{ptr: ret}
	}
	return retMapped
}

// --- Function av_packet_clone ---

// AVPacketClone wraps av_packet_clone.
func AVPacketClone(src *AVPacket) *AVPacket {
	var tmpsrc *C.AVPacket
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_packet_clone(tmpsrc)
	var retMapped *AVPacket
	if ret != nil {
		retMapped = &AVPacket{ptr: ret}
	}
	return retMapped
}

// --- Function av_packet_free ---

// AVPacketFree wraps av_packet_free.
func AVPacketFree(pkt **AVPacket) {
	var ptrpkt **C.AVPacket
	var tmppkt *C.AVPacket
	var oldTmppkt *C.AVPacket
	if pkt != nil {
		innerpkt := *pkt
		if innerpkt != nil {
			tmppkt = innerpkt.ptr
			oldTmppkt = tmppkt
		}
		ptrpkt = &tmppkt
	}
	C.av_packet_free(ptrpkt)
	if tmppkt != oldTmppkt && pkt != nil {
		if tmppkt != nil {
			*pkt = &AVPacket{ptr: tmppkt}
		} else {
			*pkt = nil
		}
	}
}

// --- Function av_init_packet ---

// AVInitPacket wraps av_init_packet.
func AVInitPacket(pkt *AVPacket) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	C.av_init_packet(tmppkt)
}

// --- Function av_new_packet ---

// AVNewPacket wraps av_new_packet.
func AVNewPacket(pkt *AVPacket, size int) int {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_new_packet(tmppkt, C.int(size))
	return int(ret)
}

// --- Function av_shrink_packet ---

// AVShrinkPacket wraps av_shrink_packet.
func AVShrinkPacket(pkt *AVPacket, size int) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	C.av_shrink_packet(tmppkt, C.int(size))
}

// --- Function av_grow_packet ---

// AVGrowPacket wraps av_grow_packet.
func AVGrowPacket(pkt *AVPacket, growBy int) int {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_grow_packet(tmppkt, C.int(growBy))
	return int(ret)
}

// --- Function av_packet_from_data ---

// AVPacketFromData wraps av_packet_from_data.
func AVPacketFromData(pkt *AVPacket, data unsafe.Pointer, size int) int {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_packet_from_data(tmppkt, (*C.uint8_t)(data), C.int(size))
	return int(ret)
}

// --- Function av_packet_new_side_data ---

// AVPacketNewSideData wraps av_packet_new_side_data.
func AVPacketNewSideData(pkt *AVPacket, _type AVPacketSideDataType, size uint64) unsafe.Pointer {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_packet_new_side_data(tmppkt, C.enum_AVPacketSideDataType(_type), C.size_t(size))
	return unsafe.Pointer(ret)
}

// --- Function av_packet_add_side_data ---

// AVPacketAddSideData wraps av_packet_add_side_data.
func AVPacketAddSideData(pkt *AVPacket, _type AVPacketSideDataType, data unsafe.Pointer, size uint64) int {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_packet_add_side_data(tmppkt, C.enum_AVPacketSideDataType(_type), (*C.uint8_t)(data), C.size_t(size))
	return int(ret)
}

// --- Function av_packet_shrink_side_data ---

// AVPacketShrinkSideData wraps av_packet_shrink_side_data.
func AVPacketShrinkSideData(pkt *AVPacket, _type AVPacketSideDataType, size uint64) int {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_packet_shrink_side_data(tmppkt, C.enum_AVPacketSideDataType(_type), C.size_t(size))
	return int(ret)
}

// --- Function av_packet_get_side_data ---

// AVPacketGetSideData wraps av_packet_get_side_data.
// av_packet_get_side_data skipped due to size

// --- Function av_packet_side_data_name ---

// AVPacketSideDataName wraps av_packet_side_data_name.
func AVPacketSideDataName(_type AVPacketSideDataType) *CStr {
	ret := C.av_packet_side_data_name(C.enum_AVPacketSideDataType(_type))
	return wrapCStr(ret)
}

// --- Function av_packet_pack_dictionary ---

// AVPacketPackDictionary wraps av_packet_pack_dictionary.
// av_packet_pack_dictionary skipped due to size

// --- Function av_packet_unpack_dictionary ---

// AVPacketUnpackDictionary wraps av_packet_unpack_dictionary.
func AVPacketUnpackDictionary(data unsafe.Pointer, size uint64, dict **AVDictionary) int {
	var ptrdict **C.AVDictionary
	var tmpdict *C.AVDictionary
	var oldTmpdict *C.AVDictionary
	if dict != nil {
		innerdict := *dict
		if innerdict != nil {
			tmpdict = innerdict.ptr
			oldTmpdict = tmpdict
		}
		ptrdict = &tmpdict
	}
	ret := C.av_packet_unpack_dictionary((*C.uint8_t)(data), C.size_t(size), ptrdict)
	if tmpdict != oldTmpdict && dict != nil {
		if tmpdict != nil {
			*dict = &AVDictionary{ptr: tmpdict}
		} else {
			*dict = nil
		}
	}
	return int(ret)
}

// --- Function av_packet_free_side_data ---

// AVPacketFreeSideData wraps av_packet_free_side_data.
func AVPacketFreeSideData(pkt *AVPacket) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	C.av_packet_free_side_data(tmppkt)
}

// --- Function av_packet_ref ---

// AVPacketRef wraps av_packet_ref.
func AVPacketRef(dst *AVPacket, src *AVPacket) int {
	var tmpdst *C.AVPacket
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVPacket
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_packet_ref(tmpdst, tmpsrc)
	return int(ret)
}

// --- Function av_packet_unref ---

// AVPacketUnref wraps av_packet_unref.
func AVPacketUnref(pkt *AVPacket) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	C.av_packet_unref(tmppkt)
}

// --- Function av_packet_move_ref ---

// AVPacketMoveRef wraps av_packet_move_ref.
func AVPacketMoveRef(dst *AVPacket, src *AVPacket) {
	var tmpdst *C.AVPacket
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVPacket
	if src != nil {
		tmpsrc = src.ptr
	}
	C.av_packet_move_ref(tmpdst, tmpsrc)
}

// --- Function av_packet_copy_props ---

// AVPacketCopyProps wraps av_packet_copy_props.
func AVPacketCopyProps(dst *AVPacket, src *AVPacket) int {
	var tmpdst *C.AVPacket
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVPacket
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_packet_copy_props(tmpdst, tmpsrc)
	return int(ret)
}

// --- Function av_packet_make_refcounted ---

// AVPacketMakeRefcounted wraps av_packet_make_refcounted.
func AVPacketMakeRefcounted(pkt *AVPacket) int {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_packet_make_refcounted(tmppkt)
	return int(ret)
}

// --- Function av_packet_make_writable ---

// AVPacketMakeWritable wraps av_packet_make_writable.
func AVPacketMakeWritable(pkt *AVPacket) int {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_packet_make_writable(tmppkt)
	return int(ret)
}

// --- Function av_packet_rescale_ts ---

// AVPacketRescaleTs wraps av_packet_rescale_ts.
func AVPacketRescaleTs(pkt *AVPacket, tbSrc *AVRational, tbDst *AVRational) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	C.av_packet_rescale_ts(tmppkt, tbSrc.value, tbDst.value)
}

// --- Function avfilter_version ---

// AVFilterVersion wraps avfilter_version.
func AVFilterVersion() uint {
	ret := C.avfilter_version()
	return uint(ret)
}

// --- Function avfilter_configuration ---

// AVFilterConfiguration wraps avfilter_configuration.
func AVFilterConfiguration() *CStr {
	ret := C.avfilter_configuration()
	return wrapCStr(ret)
}

// --- Function avfilter_license ---

// AVFilterLicense wraps avfilter_license.
func AVFilterLicense() *CStr {
	ret := C.avfilter_license()
	return wrapCStr(ret)
}

// --- Function avfilter_pad_get_name ---

// AVFilterPadGetName wraps avfilter_pad_get_name.
func AVFilterPadGetName(pads *AVFilterPad, padIdx int) *CStr {
	var tmppads *C.AVFilterPad
	if pads != nil {
		tmppads = pads.ptr
	}
	ret := C.avfilter_pad_get_name(tmppads, C.int(padIdx))
	return wrapCStr(ret)
}

// --- Function avfilter_pad_get_type ---

// AVFilterPadGetType wraps avfilter_pad_get_type.
func AVFilterPadGetType(pads *AVFilterPad, padIdx int) AVMediaType {
	var tmppads *C.AVFilterPad
	if pads != nil {
		tmppads = pads.ptr
	}
	ret := C.avfilter_pad_get_type(tmppads, C.int(padIdx))
	return AVMediaType(ret)
}

// --- Function avfilter_filter_pad_count ---

// AVFilterFilterPadCount wraps avfilter_filter_pad_count.
func AVFilterFilterPadCount(filter *AVFilter, isOutput int) uint {
	var tmpfilter *C.AVFilter
	if filter != nil {
		tmpfilter = filter.ptr
	}
	ret := C.avfilter_filter_pad_count(tmpfilter, C.int(isOutput))
	return uint(ret)
}

// --- Function avfilter_link ---

// AVFilterLink_ wraps avfilter_link.
func AVFilterLink_(src *AVFilterContext, srcpad uint, dst *AVFilterContext, dstpad uint) int {
	var tmpsrc *C.AVFilterContext
	if src != nil {
		tmpsrc = src.ptr
	}
	var tmpdst *C.AVFilterContext
	if dst != nil {
		tmpdst = dst.ptr
	}
	ret := C.avfilter_link(tmpsrc, C.uint(srcpad), tmpdst, C.uint(dstpad))
	return int(ret)
}

// --- Function avfilter_link_free ---

// AVFilterLinkFree wraps avfilter_link_free.
func AVFilterLinkFree(link **AVFilterLink) {
	var ptrlink **C.AVFilterLink
	var tmplink *C.AVFilterLink
	var oldTmplink *C.AVFilterLink
	if link != nil {
		innerlink := *link
		if innerlink != nil {
			tmplink = innerlink.ptr
			oldTmplink = tmplink
		}
		ptrlink = &tmplink
	}
	C.avfilter_link_free(ptrlink)
	if tmplink != oldTmplink && link != nil {
		if tmplink != nil {
			*link = &AVFilterLink{ptr: tmplink}
		} else {
			*link = nil
		}
	}
}

// --- Function avfilter_config_links ---

// AVFilterConfigLinks wraps avfilter_config_links.
func AVFilterConfigLinks(filter *AVFilterContext) int {
	var tmpfilter *C.AVFilterContext
	if filter != nil {
		tmpfilter = filter.ptr
	}
	ret := C.avfilter_config_links(tmpfilter)
	return int(ret)
}

// --- Function avfilter_process_command ---

// AVFilterProcessCommand wraps avfilter_process_command.
func AVFilterProcessCommand(filter *AVFilterContext, cmd *CStr, arg *CStr, res *CStr, resLen int, flags int) int {
	var tmpfilter *C.AVFilterContext
	if filter != nil {
		tmpfilter = filter.ptr
	}
	var tmpcmd *C.char
	if cmd != nil {
		tmpcmd = cmd.ptr
	}
	var tmparg *C.char
	if arg != nil {
		tmparg = arg.ptr
	}
	var tmpres *C.char
	if res != nil {
		tmpres = res.ptr
	}
	ret := C.avfilter_process_command(tmpfilter, tmpcmd, tmparg, tmpres, C.int(resLen), C.int(flags))
	return int(ret)
}

// --- Function av_filter_iterate ---

// AVFilterIterate wraps av_filter_iterate.
// av_filter_iterate skipped due to opaque

// --- Function avfilter_get_by_name ---

// AVFilterGetByName wraps avfilter_get_by_name.
func AVFilterGetByName(name *CStr) *AVFilter {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.avfilter_get_by_name(tmpname)
	var retMapped *AVFilter
	if ret != nil {
		retMapped = &AVFilter{ptr: ret}
	}
	return retMapped
}

// --- Function avfilter_init_str ---

// AVFilterInitStr wraps avfilter_init_str.
func AVFilterInitStr(ctx *AVFilterContext, args *CStr) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpargs *C.char
	if args != nil {
		tmpargs = args.ptr
	}
	ret := C.avfilter_init_str(tmpctx, tmpargs)
	return int(ret)
}

// --- Function avfilter_init_dict ---

// AVFilterInitDict wraps avfilter_init_dict.
func AVFilterInitDict(ctx *AVFilterContext, options **AVDictionary) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var ptroptions **C.AVDictionary
	var tmpoptions *C.AVDictionary
	var oldTmpoptions *C.AVDictionary
	if options != nil {
		inneroptions := *options
		if inneroptions != nil {
			tmpoptions = inneroptions.ptr
			oldTmpoptions = tmpoptions
		}
		ptroptions = &tmpoptions
	}
	ret := C.avfilter_init_dict(tmpctx, ptroptions)
	if tmpoptions != oldTmpoptions && options != nil {
		if tmpoptions != nil {
			*options = &AVDictionary{ptr: tmpoptions}
		} else {
			*options = nil
		}
	}
	return int(ret)
}

// --- Function avfilter_free ---

// AVFilterFree wraps avfilter_free.
func AVFilterFree(filter *AVFilterContext) {
	var tmpfilter *C.AVFilterContext
	if filter != nil {
		tmpfilter = filter.ptr
	}
	C.avfilter_free(tmpfilter)
}

// --- Function avfilter_insert_filter ---

// AVFilterInsertFilter wraps avfilter_insert_filter.
func AVFilterInsertFilter(link *AVFilterLink, filt *AVFilterContext, filtSrcpadIdx uint, filtDstpadIdx uint) int {
	var tmplink *C.AVFilterLink
	if link != nil {
		tmplink = link.ptr
	}
	var tmpfilt *C.AVFilterContext
	if filt != nil {
		tmpfilt = filt.ptr
	}
	ret := C.avfilter_insert_filter(tmplink, tmpfilt, C.uint(filtSrcpadIdx), C.uint(filtDstpadIdx))
	return int(ret)
}

// --- Function avfilter_get_class ---

// AVFilterGetClass wraps avfilter_get_class.
func AVFilterGetClass() *AVClass {
	ret := C.avfilter_get_class()
	var retMapped *AVClass
	if ret != nil {
		retMapped = &AVClass{ptr: ret}
	}
	return retMapped
}

// --- Function avfilter_graph_alloc ---

// AVFilterGraphAlloc wraps avfilter_graph_alloc.
func AVFilterGraphAlloc() *AVFilterGraph {
	ret := C.avfilter_graph_alloc()
	var retMapped *AVFilterGraph
	if ret != nil {
		retMapped = &AVFilterGraph{ptr: ret}
	}
	return retMapped
}

// --- Function avfilter_graph_alloc_filter ---

// AVFilterGraphAllocFilter wraps avfilter_graph_alloc_filter.
func AVFilterGraphAllocFilter(graph *AVFilterGraph, filter *AVFilter, name *CStr) *AVFilterContext {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	var tmpfilter *C.AVFilter
	if filter != nil {
		tmpfilter = filter.ptr
	}
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.avfilter_graph_alloc_filter(tmpgraph, tmpfilter, tmpname)
	var retMapped *AVFilterContext
	if ret != nil {
		retMapped = &AVFilterContext{ptr: ret}
	}
	return retMapped
}

// --- Function avfilter_graph_get_filter ---

// AVFilterGraphGetFilter wraps avfilter_graph_get_filter.
func AVFilterGraphGetFilter(graph *AVFilterGraph, name *CStr) *AVFilterContext {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.avfilter_graph_get_filter(tmpgraph, tmpname)
	var retMapped *AVFilterContext
	if ret != nil {
		retMapped = &AVFilterContext{ptr: ret}
	}
	return retMapped
}

// --- Function avfilter_graph_create_filter ---

// AVFilterGraphCreateFilter wraps avfilter_graph_create_filter.
func AVFilterGraphCreateFilter(filtCtx **AVFilterContext, filt *AVFilter, name *CStr, args *CStr, opaque unsafe.Pointer, graphCtx *AVFilterGraph) int {
	var ptrfiltCtx **C.AVFilterContext
	var tmpfiltCtx *C.AVFilterContext
	var oldTmpfiltCtx *C.AVFilterContext
	if filtCtx != nil {
		innerfiltCtx := *filtCtx
		if innerfiltCtx != nil {
			tmpfiltCtx = innerfiltCtx.ptr
			oldTmpfiltCtx = tmpfiltCtx
		}
		ptrfiltCtx = &tmpfiltCtx
	}
	var tmpfilt *C.AVFilter
	if filt != nil {
		tmpfilt = filt.ptr
	}
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	var tmpargs *C.char
	if args != nil {
		tmpargs = args.ptr
	}
	var tmpgraphCtx *C.AVFilterGraph
	if graphCtx != nil {
		tmpgraphCtx = graphCtx.ptr
	}
	ret := C.avfilter_graph_create_filter(ptrfiltCtx, tmpfilt, tmpname, tmpargs, opaque, tmpgraphCtx)
	if tmpfiltCtx != oldTmpfiltCtx && filtCtx != nil {
		if tmpfiltCtx != nil {
			*filtCtx = &AVFilterContext{ptr: tmpfiltCtx}
		} else {
			*filtCtx = nil
		}
	}
	return int(ret)
}

// --- Function avfilter_graph_set_auto_convert ---

// AVFilterGraphSetAutoConvert wraps avfilter_graph_set_auto_convert.
func AVFilterGraphSetAutoConvert(graph *AVFilterGraph, flags uint) {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	C.avfilter_graph_set_auto_convert(tmpgraph, C.uint(flags))
}

// --- Function avfilter_graph_config ---

// AVFilterGraphConfig wraps avfilter_graph_config.
func AVFilterGraphConfig(graphctx *AVFilterGraph, logCtx unsafe.Pointer) int {
	var tmpgraphctx *C.AVFilterGraph
	if graphctx != nil {
		tmpgraphctx = graphctx.ptr
	}
	ret := C.avfilter_graph_config(tmpgraphctx, logCtx)
	return int(ret)
}

// --- Function avfilter_graph_free ---

// AVFilterGraphFree wraps avfilter_graph_free.
func AVFilterGraphFree(graph **AVFilterGraph) {
	var ptrgraph **C.AVFilterGraph
	var tmpgraph *C.AVFilterGraph
	var oldTmpgraph *C.AVFilterGraph
	if graph != nil {
		innergraph := *graph
		if innergraph != nil {
			tmpgraph = innergraph.ptr
			oldTmpgraph = tmpgraph
		}
		ptrgraph = &tmpgraph
	}
	C.avfilter_graph_free(ptrgraph)
	if tmpgraph != oldTmpgraph && graph != nil {
		if tmpgraph != nil {
			*graph = &AVFilterGraph{ptr: tmpgraph}
		} else {
			*graph = nil
		}
	}
}

// --- Function avfilter_inout_alloc ---

// AVFilterInoutAlloc wraps avfilter_inout_alloc.
func AVFilterInoutAlloc() *AVFilterInOut {
	ret := C.avfilter_inout_alloc()
	var retMapped *AVFilterInOut
	if ret != nil {
		retMapped = &AVFilterInOut{ptr: ret}
	}
	return retMapped
}

// --- Function avfilter_inout_free ---

// AVFilterInoutFree wraps avfilter_inout_free.
func AVFilterInoutFree(inout **AVFilterInOut) {
	var ptrinout **C.AVFilterInOut
	var tmpinout *C.AVFilterInOut
	var oldTmpinout *C.AVFilterInOut
	if inout != nil {
		innerinout := *inout
		if innerinout != nil {
			tmpinout = innerinout.ptr
			oldTmpinout = tmpinout
		}
		ptrinout = &tmpinout
	}
	C.avfilter_inout_free(ptrinout)
	if tmpinout != oldTmpinout && inout != nil {
		if tmpinout != nil {
			*inout = &AVFilterInOut{ptr: tmpinout}
		} else {
			*inout = nil
		}
	}
}

// --- Function avfilter_graph_parse ---

// AVFilterGraphParse wraps avfilter_graph_parse.
func AVFilterGraphParse(graph *AVFilterGraph, filters *CStr, inputs *AVFilterInOut, outputs *AVFilterInOut, logCtx unsafe.Pointer) int {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	var tmpfilters *C.char
	if filters != nil {
		tmpfilters = filters.ptr
	}
	var tmpinputs *C.AVFilterInOut
	if inputs != nil {
		tmpinputs = inputs.ptr
	}
	var tmpoutputs *C.AVFilterInOut
	if outputs != nil {
		tmpoutputs = outputs.ptr
	}
	ret := C.avfilter_graph_parse(tmpgraph, tmpfilters, tmpinputs, tmpoutputs, logCtx)
	return int(ret)
}

// --- Function avfilter_graph_parse_ptr ---

// AVFilterGraphParsePtr wraps avfilter_graph_parse_ptr.
func AVFilterGraphParsePtr(graph *AVFilterGraph, filters *CStr, inputs **AVFilterInOut, outputs **AVFilterInOut, logCtx unsafe.Pointer) int {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	var tmpfilters *C.char
	if filters != nil {
		tmpfilters = filters.ptr
	}
	var ptrinputs **C.AVFilterInOut
	var tmpinputs *C.AVFilterInOut
	var oldTmpinputs *C.AVFilterInOut
	if inputs != nil {
		innerinputs := *inputs
		if innerinputs != nil {
			tmpinputs = innerinputs.ptr
			oldTmpinputs = tmpinputs
		}
		ptrinputs = &tmpinputs
	}
	var ptroutputs **C.AVFilterInOut
	var tmpoutputs *C.AVFilterInOut
	var oldTmpoutputs *C.AVFilterInOut
	if outputs != nil {
		inneroutputs := *outputs
		if inneroutputs != nil {
			tmpoutputs = inneroutputs.ptr
			oldTmpoutputs = tmpoutputs
		}
		ptroutputs = &tmpoutputs
	}
	ret := C.avfilter_graph_parse_ptr(tmpgraph, tmpfilters, ptrinputs, ptroutputs, logCtx)
	if tmpinputs != oldTmpinputs && inputs != nil {
		if tmpinputs != nil {
			*inputs = &AVFilterInOut{ptr: tmpinputs}
		} else {
			*inputs = nil
		}
	}
	if tmpoutputs != oldTmpoutputs && outputs != nil {
		if tmpoutputs != nil {
			*outputs = &AVFilterInOut{ptr: tmpoutputs}
		} else {
			*outputs = nil
		}
	}
	return int(ret)
}

// --- Function avfilter_graph_parse2 ---

// AVFilterGraphParse2 wraps avfilter_graph_parse2.
func AVFilterGraphParse2(graph *AVFilterGraph, filters *CStr, inputs **AVFilterInOut, outputs **AVFilterInOut) int {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	var tmpfilters *C.char
	if filters != nil {
		tmpfilters = filters.ptr
	}
	var ptrinputs **C.AVFilterInOut
	var tmpinputs *C.AVFilterInOut
	var oldTmpinputs *C.AVFilterInOut
	if inputs != nil {
		innerinputs := *inputs
		if innerinputs != nil {
			tmpinputs = innerinputs.ptr
			oldTmpinputs = tmpinputs
		}
		ptrinputs = &tmpinputs
	}
	var ptroutputs **C.AVFilterInOut
	var tmpoutputs *C.AVFilterInOut
	var oldTmpoutputs *C.AVFilterInOut
	if outputs != nil {
		inneroutputs := *outputs
		if inneroutputs != nil {
			tmpoutputs = inneroutputs.ptr
			oldTmpoutputs = tmpoutputs
		}
		ptroutputs = &tmpoutputs
	}
	ret := C.avfilter_graph_parse2(tmpgraph, tmpfilters, ptrinputs, ptroutputs)
	if tmpinputs != oldTmpinputs && inputs != nil {
		if tmpinputs != nil {
			*inputs = &AVFilterInOut{ptr: tmpinputs}
		} else {
			*inputs = nil
		}
	}
	if tmpoutputs != oldTmpoutputs && outputs != nil {
		if tmpoutputs != nil {
			*outputs = &AVFilterInOut{ptr: tmpoutputs}
		} else {
			*outputs = nil
		}
	}
	return int(ret)
}

// --- Function avfilter_graph_segment_parse ---

// AVFilterGraphSegmentParse wraps avfilter_graph_segment_parse.
func AVFilterGraphSegmentParse(graph *AVFilterGraph, graphStr *CStr, flags int, seg **AVFilterGraphSegment) int {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	var tmpgraphStr *C.char
	if graphStr != nil {
		tmpgraphStr = graphStr.ptr
	}
	var ptrseg **C.AVFilterGraphSegment
	var tmpseg *C.AVFilterGraphSegment
	var oldTmpseg *C.AVFilterGraphSegment
	if seg != nil {
		innerseg := *seg
		if innerseg != nil {
			tmpseg = innerseg.ptr
			oldTmpseg = tmpseg
		}
		ptrseg = &tmpseg
	}
	ret := C.avfilter_graph_segment_parse(tmpgraph, tmpgraphStr, C.int(flags), ptrseg)
	if tmpseg != oldTmpseg && seg != nil {
		if tmpseg != nil {
			*seg = &AVFilterGraphSegment{ptr: tmpseg}
		} else {
			*seg = nil
		}
	}
	return int(ret)
}

// --- Function avfilter_graph_segment_create_filters ---

// AVFilterGraphSegmentCreateFilters wraps avfilter_graph_segment_create_filters.
func AVFilterGraphSegmentCreateFilters(seg *AVFilterGraphSegment, flags int) int {
	var tmpseg *C.AVFilterGraphSegment
	if seg != nil {
		tmpseg = seg.ptr
	}
	ret := C.avfilter_graph_segment_create_filters(tmpseg, C.int(flags))
	return int(ret)
}

// --- Function avfilter_graph_segment_apply_opts ---

// AVFilterGraphSegmentApplyOpts wraps avfilter_graph_segment_apply_opts.
func AVFilterGraphSegmentApplyOpts(seg *AVFilterGraphSegment, flags int) int {
	var tmpseg *C.AVFilterGraphSegment
	if seg != nil {
		tmpseg = seg.ptr
	}
	ret := C.avfilter_graph_segment_apply_opts(tmpseg, C.int(flags))
	return int(ret)
}

// --- Function avfilter_graph_segment_init ---

// AVFilterGraphSegmentInit wraps avfilter_graph_segment_init.
func AVFilterGraphSegmentInit(seg *AVFilterGraphSegment, flags int) int {
	var tmpseg *C.AVFilterGraphSegment
	if seg != nil {
		tmpseg = seg.ptr
	}
	ret := C.avfilter_graph_segment_init(tmpseg, C.int(flags))
	return int(ret)
}

// --- Function avfilter_graph_segment_link ---

// AVFilterGraphSegmentLink wraps avfilter_graph_segment_link.
func AVFilterGraphSegmentLink(seg *AVFilterGraphSegment, flags int, inputs **AVFilterInOut, outputs **AVFilterInOut) int {
	var tmpseg *C.AVFilterGraphSegment
	if seg != nil {
		tmpseg = seg.ptr
	}
	var ptrinputs **C.AVFilterInOut
	var tmpinputs *C.AVFilterInOut
	var oldTmpinputs *C.AVFilterInOut
	if inputs != nil {
		innerinputs := *inputs
		if innerinputs != nil {
			tmpinputs = innerinputs.ptr
			oldTmpinputs = tmpinputs
		}
		ptrinputs = &tmpinputs
	}
	var ptroutputs **C.AVFilterInOut
	var tmpoutputs *C.AVFilterInOut
	var oldTmpoutputs *C.AVFilterInOut
	if outputs != nil {
		inneroutputs := *outputs
		if inneroutputs != nil {
			tmpoutputs = inneroutputs.ptr
			oldTmpoutputs = tmpoutputs
		}
		ptroutputs = &tmpoutputs
	}
	ret := C.avfilter_graph_segment_link(tmpseg, C.int(flags), ptrinputs, ptroutputs)
	if tmpinputs != oldTmpinputs && inputs != nil {
		if tmpinputs != nil {
			*inputs = &AVFilterInOut{ptr: tmpinputs}
		} else {
			*inputs = nil
		}
	}
	if tmpoutputs != oldTmpoutputs && outputs != nil {
		if tmpoutputs != nil {
			*outputs = &AVFilterInOut{ptr: tmpoutputs}
		} else {
			*outputs = nil
		}
	}
	return int(ret)
}

// --- Function avfilter_graph_segment_apply ---

// AVFilterGraphSegmentApply wraps avfilter_graph_segment_apply.
func AVFilterGraphSegmentApply(seg *AVFilterGraphSegment, flags int, inputs **AVFilterInOut, outputs **AVFilterInOut) int {
	var tmpseg *C.AVFilterGraphSegment
	if seg != nil {
		tmpseg = seg.ptr
	}
	var ptrinputs **C.AVFilterInOut
	var tmpinputs *C.AVFilterInOut
	var oldTmpinputs *C.AVFilterInOut
	if inputs != nil {
		innerinputs := *inputs
		if innerinputs != nil {
			tmpinputs = innerinputs.ptr
			oldTmpinputs = tmpinputs
		}
		ptrinputs = &tmpinputs
	}
	var ptroutputs **C.AVFilterInOut
	var tmpoutputs *C.AVFilterInOut
	var oldTmpoutputs *C.AVFilterInOut
	if outputs != nil {
		inneroutputs := *outputs
		if inneroutputs != nil {
			tmpoutputs = inneroutputs.ptr
			oldTmpoutputs = tmpoutputs
		}
		ptroutputs = &tmpoutputs
	}
	ret := C.avfilter_graph_segment_apply(tmpseg, C.int(flags), ptrinputs, ptroutputs)
	if tmpinputs != oldTmpinputs && inputs != nil {
		if tmpinputs != nil {
			*inputs = &AVFilterInOut{ptr: tmpinputs}
		} else {
			*inputs = nil
		}
	}
	if tmpoutputs != oldTmpoutputs && outputs != nil {
		if tmpoutputs != nil {
			*outputs = &AVFilterInOut{ptr: tmpoutputs}
		} else {
			*outputs = nil
		}
	}
	return int(ret)
}

// --- Function avfilter_graph_segment_free ---

// AVFilterGraphSegmentFree wraps avfilter_graph_segment_free.
func AVFilterGraphSegmentFree(seg **AVFilterGraphSegment) {
	var ptrseg **C.AVFilterGraphSegment
	var tmpseg *C.AVFilterGraphSegment
	var oldTmpseg *C.AVFilterGraphSegment
	if seg != nil {
		innerseg := *seg
		if innerseg != nil {
			tmpseg = innerseg.ptr
			oldTmpseg = tmpseg
		}
		ptrseg = &tmpseg
	}
	C.avfilter_graph_segment_free(ptrseg)
	if tmpseg != oldTmpseg && seg != nil {
		if tmpseg != nil {
			*seg = &AVFilterGraphSegment{ptr: tmpseg}
		} else {
			*seg = nil
		}
	}
}

// --- Function avfilter_graph_send_command ---

// AVFilterGraphSendCommand wraps avfilter_graph_send_command.
func AVFilterGraphSendCommand(graph *AVFilterGraph, target *CStr, cmd *CStr, arg *CStr, res *CStr, resLen int, flags int) int {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	var tmptarget *C.char
	if target != nil {
		tmptarget = target.ptr
	}
	var tmpcmd *C.char
	if cmd != nil {
		tmpcmd = cmd.ptr
	}
	var tmparg *C.char
	if arg != nil {
		tmparg = arg.ptr
	}
	var tmpres *C.char
	if res != nil {
		tmpres = res.ptr
	}
	ret := C.avfilter_graph_send_command(tmpgraph, tmptarget, tmpcmd, tmparg, tmpres, C.int(resLen), C.int(flags))
	return int(ret)
}

// --- Function avfilter_graph_queue_command ---

// AVFilterGraphQueueCommand wraps avfilter_graph_queue_command.
func AVFilterGraphQueueCommand(graph *AVFilterGraph, target *CStr, cmd *CStr, arg *CStr, flags int, ts float64) int {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	var tmptarget *C.char
	if target != nil {
		tmptarget = target.ptr
	}
	var tmpcmd *C.char
	if cmd != nil {
		tmpcmd = cmd.ptr
	}
	var tmparg *C.char
	if arg != nil {
		tmparg = arg.ptr
	}
	ret := C.avfilter_graph_queue_command(tmpgraph, tmptarget, tmpcmd, tmparg, C.int(flags), C.double(ts))
	return int(ret)
}

// --- Function avfilter_graph_dump ---

// AVFilterGraphDump wraps avfilter_graph_dump.
func AVFilterGraphDump(graph *AVFilterGraph, options *CStr) *CStr {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	var tmpoptions *C.char
	if options != nil {
		tmpoptions = options.ptr
	}
	ret := C.avfilter_graph_dump(tmpgraph, tmpoptions)
	return wrapCStr(ret)
}

// --- Function avfilter_graph_request_oldest ---

// AVFilterGraphRequestOldest wraps avfilter_graph_request_oldest.
func AVFilterGraphRequestOldest(graph *AVFilterGraph) int {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	ret := C.avfilter_graph_request_oldest(tmpgraph)
	return int(ret)
}

// --- Function av_buffersink_get_frame_flags ---

// AVBuffersinkGetFrameFlags wraps av_buffersink_get_frame_flags.
func AVBuffersinkGetFrameFlags(ctx *AVFilterContext, frame *AVFrame, flags int) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersink_get_frame_flags(tmpctx, tmpframe, C.int(flags))
	return int(ret)
}

// --- Function av_buffersink_set_frame_size ---

// AVBuffersinkSetFrameSize wraps av_buffersink_set_frame_size.
func AVBuffersinkSetFrameSize(ctx *AVFilterContext, frameSize uint) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	C.av_buffersink_set_frame_size(tmpctx, C.uint(frameSize))
}

// --- Function av_buffersink_get_type ---

// AVBuffersinkGetType wraps av_buffersink_get_type.
func AVBuffersinkGetType(ctx *AVFilterContext) AVMediaType {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_type(tmpctx)
	return AVMediaType(ret)
}

// --- Function av_buffersink_get_time_base ---

// AVBuffersinkGetTimeBase wraps av_buffersink_get_time_base.
func AVBuffersinkGetTimeBase(ctx *AVFilterContext) *AVRational {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_time_base(tmpctx)
	return &AVRational{value: ret}
}

// --- Function av_buffersink_get_format ---

// AVBuffersinkGetFormat wraps av_buffersink_get_format.
func AVBuffersinkGetFormat(ctx *AVFilterContext) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_format(tmpctx)
	return int(ret)
}

// --- Function av_buffersink_get_frame_rate ---

// AVBuffersinkGetFrameRate wraps av_buffersink_get_frame_rate.
func AVBuffersinkGetFrameRate(ctx *AVFilterContext) *AVRational {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_frame_rate(tmpctx)
	return &AVRational{value: ret}
}

// --- Function av_buffersink_get_w ---

// AVBuffersinkGetW wraps av_buffersink_get_w.
func AVBuffersinkGetW(ctx *AVFilterContext) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_w(tmpctx)
	return int(ret)
}

// --- Function av_buffersink_get_h ---

// AVBuffersinkGetH wraps av_buffersink_get_h.
func AVBuffersinkGetH(ctx *AVFilterContext) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_h(tmpctx)
	return int(ret)
}

// --- Function av_buffersink_get_sample_aspect_ratio ---

// AVBuffersinkGetSampleAspectRatio wraps av_buffersink_get_sample_aspect_ratio.
func AVBuffersinkGetSampleAspectRatio(ctx *AVFilterContext) *AVRational {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_sample_aspect_ratio(tmpctx)
	return &AVRational{value: ret}
}

// --- Function av_buffersink_get_channels ---

// AVBuffersinkGetChannels wraps av_buffersink_get_channels.
func AVBuffersinkGetChannels(ctx *AVFilterContext) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_channels(tmpctx)
	return int(ret)
}

// --- Function av_buffersink_get_channel_layout ---

// AVBuffersinkGetChannelLayout wraps av_buffersink_get_channel_layout.
func AVBuffersinkGetChannelLayout(ctx *AVFilterContext) uint64 {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_channel_layout(tmpctx)
	return uint64(ret)
}

// --- Function av_buffersink_get_ch_layout ---

// AVBuffersinkGetChLayout wraps av_buffersink_get_ch_layout.
func AVBuffersinkGetChLayout(ctx *AVFilterContext, chLayout *AVChannelLayout) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpchLayout *C.AVChannelLayout
	if chLayout != nil {
		tmpchLayout = chLayout.ptr
	}
	ret := C.av_buffersink_get_ch_layout(tmpctx, tmpchLayout)
	return int(ret)
}

// --- Function av_buffersink_get_sample_rate ---

// AVBuffersinkGetSampleRate wraps av_buffersink_get_sample_rate.
func AVBuffersinkGetSampleRate(ctx *AVFilterContext) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_sample_rate(tmpctx)
	return int(ret)
}

// --- Function av_buffersink_get_hw_frames_ctx ---

// AVBuffersinkGetHwFramesCtx wraps av_buffersink_get_hw_frames_ctx.
func AVBuffersinkGetHwFramesCtx(ctx *AVFilterContext) *AVBufferRef {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_hw_frames_ctx(tmpctx)
	var retMapped *AVBufferRef
	if ret != nil {
		retMapped = &AVBufferRef{ptr: ret}
	}
	return retMapped
}

// --- Function av_buffersink_get_frame ---

// AVBuffersinkGetFrame wraps av_buffersink_get_frame.
func AVBuffersinkGetFrame(ctx *AVFilterContext, frame *AVFrame) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersink_get_frame(tmpctx, tmpframe)
	return int(ret)
}

// --- Function av_buffersink_get_samples ---

// AVBuffersinkGetSamples wraps av_buffersink_get_samples.
func AVBuffersinkGetSamples(ctx *AVFilterContext, frame *AVFrame, nbSamples int) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersink_get_samples(tmpctx, tmpframe, C.int(nbSamples))
	return int(ret)
}

// --- Function av_buffersrc_get_nb_failed_requests ---

// AVBuffersrcGetNbFailedRequests wraps av_buffersrc_get_nb_failed_requests.
func AVBuffersrcGetNbFailedRequests(bufferSrc *AVFilterContext) uint {
	var tmpbufferSrc *C.AVFilterContext
	if bufferSrc != nil {
		tmpbufferSrc = bufferSrc.ptr
	}
	ret := C.av_buffersrc_get_nb_failed_requests(tmpbufferSrc)
	return uint(ret)
}

// --- Function av_buffersrc_parameters_alloc ---

// AVBuffersrcParametersAlloc wraps av_buffersrc_parameters_alloc.
func AVBuffersrcParametersAlloc() *AVBufferSrcParameters {
	ret := C.av_buffersrc_parameters_alloc()
	var retMapped *AVBufferSrcParameters
	if ret != nil {
		retMapped = &AVBufferSrcParameters{ptr: ret}
	}
	return retMapped
}

// --- Function av_buffersrc_parameters_set ---

// AVBuffersrcParametersSet wraps av_buffersrc_parameters_set.
func AVBuffersrcParametersSet(ctx *AVFilterContext, param *AVBufferSrcParameters) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpparam *C.AVBufferSrcParameters
	if param != nil {
		tmpparam = param.ptr
	}
	ret := C.av_buffersrc_parameters_set(tmpctx, tmpparam)
	return int(ret)
}

// --- Function av_buffersrc_write_frame ---

// AVBuffersrcWriteFrame wraps av_buffersrc_write_frame.
func AVBuffersrcWriteFrame(ctx *AVFilterContext, frame *AVFrame) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersrc_write_frame(tmpctx, tmpframe)
	return int(ret)
}

// --- Function av_buffersrc_add_frame ---

// AVBuffersrcAddFrame wraps av_buffersrc_add_frame.
func AVBuffersrcAddFrame(ctx *AVFilterContext, frame *AVFrame) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersrc_add_frame(tmpctx, tmpframe)
	return int(ret)
}

// --- Function av_buffersrc_add_frame_flags ---

// AVBuffersrcAddFrameFlags wraps av_buffersrc_add_frame_flags.
func AVBuffersrcAddFrameFlags(bufferSrc *AVFilterContext, frame *AVFrame, flags int) int {
	var tmpbufferSrc *C.AVFilterContext
	if bufferSrc != nil {
		tmpbufferSrc = bufferSrc.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersrc_add_frame_flags(tmpbufferSrc, tmpframe, C.int(flags))
	return int(ret)
}

// --- Function av_buffersrc_close ---

// AVBuffersrcClose wraps av_buffersrc_close.
func AVBuffersrcClose(ctx *AVFilterContext, pts int64, flags uint) int {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersrc_close(tmpctx, C.int64_t(pts), C.uint(flags))
	return int(ret)
}

// --- Function av_get_packet ---

// AVGetPacket wraps av_get_packet.
func AVGetPacket(s *AVIOContext, pkt *AVPacket, size int) int {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_get_packet(tmps, tmppkt, C.int(size))
	return int(ret)
}

// --- Function av_append_packet ---

// AVAppendPacket wraps av_append_packet.
func AVAppendPacket(s *AVIOContext, pkt *AVPacket, size int) int {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_append_packet(tmps, tmppkt, C.int(size))
	return int(ret)
}

// --- Function av_disposition_from_string ---

// AVDispositionFromString wraps av_disposition_from_string.
func AVDispositionFromString(disp *CStr) int {
	var tmpdisp *C.char
	if disp != nil {
		tmpdisp = disp.ptr
	}
	ret := C.av_disposition_from_string(tmpdisp)
	return int(ret)
}

// --- Function av_disposition_to_string ---

// AVDispositionToString wraps av_disposition_to_string.
func AVDispositionToString(disposition int) *CStr {
	ret := C.av_disposition_to_string(C.int(disposition))
	return wrapCStr(ret)
}

// --- Function av_stream_get_parser ---

// AVStreamGetParser wraps av_stream_get_parser.
func AVStreamGetParser(s *AVStream) *AVCodecParserContext {
	var tmps *C.AVStream
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_stream_get_parser(tmps)
	var retMapped *AVCodecParserContext
	if ret != nil {
		retMapped = &AVCodecParserContext{ptr: ret}
	}
	return retMapped
}

// --- Function av_stream_get_end_pts ---

// AVStreamGetEndPts wraps av_stream_get_end_pts.
func AVStreamGetEndPts(st *AVStream) int64 {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.av_stream_get_end_pts(tmpst)
	return int64(ret)
}

// --- Function av_stream_get_first_dts ---

// AVStreamGetFirstDts wraps av_stream_get_first_dts.
func AVStreamGetFirstDts(st *AVStream) int64 {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.av_stream_get_first_dts(tmpst)
	return int64(ret)
}

// --- Function av_format_inject_global_side_data ---

// AVFormatInjectGlobalSideData wraps av_format_inject_global_side_data.
func AVFormatInjectGlobalSideData(s *AVFormatContext) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	C.av_format_inject_global_side_data(tmps)
}

// --- Function av_fmt_ctx_get_duration_estimation_method ---

// AVFmtCtxGetDurationEstimationMethod wraps av_fmt_ctx_get_duration_estimation_method.
func AVFmtCtxGetDurationEstimationMethod(ctx *AVFormatContext) AVDurationEstimationMethod {
	var tmpctx *C.AVFormatContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_fmt_ctx_get_duration_estimation_method(tmpctx)
	return AVDurationEstimationMethod(ret)
}

// --- Function avformat_version ---

// AVFormatVersion wraps avformat_version.
func AVFormatVersion() uint {
	ret := C.avformat_version()
	return uint(ret)
}

// --- Function avformat_configuration ---

// AVFormatConfiguration wraps avformat_configuration.
func AVFormatConfiguration() *CStr {
	ret := C.avformat_configuration()
	return wrapCStr(ret)
}

// --- Function avformat_license ---

// AVFormatLicense wraps avformat_license.
func AVFormatLicense() *CStr {
	ret := C.avformat_license()
	return wrapCStr(ret)
}

// --- Function avformat_network_init ---

// AVFormatNetworkInit wraps avformat_network_init.
func AVFormatNetworkInit() int {
	ret := C.avformat_network_init()
	return int(ret)
}

// --- Function avformat_network_deinit ---

// AVFormatNetworkDeinit wraps avformat_network_deinit.
func AVFormatNetworkDeinit() int {
	ret := C.avformat_network_deinit()
	return int(ret)
}

// --- Function av_muxer_iterate ---

// AVMuxerIterate wraps av_muxer_iterate.
// av_muxer_iterate skipped due to opaque

// --- Function av_demuxer_iterate ---

// AVDemuxerIterate wraps av_demuxer_iterate.
// av_demuxer_iterate skipped due to opaque

// --- Function avformat_alloc_context ---

// AVFormatAllocContext wraps avformat_alloc_context.
func AVFormatAllocContext() *AVFormatContext {
	ret := C.avformat_alloc_context()
	var retMapped *AVFormatContext
	if ret != nil {
		retMapped = &AVFormatContext{ptr: ret}
	}
	return retMapped
}

// --- Function avformat_free_context ---

// AVFormatFreeContext wraps avformat_free_context.
func AVFormatFreeContext(s *AVFormatContext) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	C.avformat_free_context(tmps)
}

// --- Function avformat_get_class ---

// AVFormatGetClass wraps avformat_get_class.
func AVFormatGetClass() *AVClass {
	ret := C.avformat_get_class()
	var retMapped *AVClass
	if ret != nil {
		retMapped = &AVClass{ptr: ret}
	}
	return retMapped
}

// --- Function av_stream_get_class ---

// AVStreamGetClass wraps av_stream_get_class.
func AVStreamGetClass() *AVClass {
	ret := C.av_stream_get_class()
	var retMapped *AVClass
	if ret != nil {
		retMapped = &AVClass{ptr: ret}
	}
	return retMapped
}

// --- Function avformat_new_stream ---

// AVFormatNewStream wraps avformat_new_stream.
func AVFormatNewStream(s *AVFormatContext, c *AVCodec) *AVStream {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpc *C.AVCodec
	if c != nil {
		tmpc = c.ptr
	}
	ret := C.avformat_new_stream(tmps, tmpc)
	var retMapped *AVStream
	if ret != nil {
		retMapped = &AVStream{ptr: ret}
	}
	return retMapped
}

// --- Function av_stream_add_side_data ---

// AVStreamAddSideData wraps av_stream_add_side_data.
func AVStreamAddSideData(st *AVStream, _type AVPacketSideDataType, data unsafe.Pointer, size uint64) int {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.av_stream_add_side_data(tmpst, C.enum_AVPacketSideDataType(_type), (*C.uint8_t)(data), C.size_t(size))
	return int(ret)
}

// --- Function av_stream_new_side_data ---

// AVStreamNewSideData wraps av_stream_new_side_data.
func AVStreamNewSideData(stream *AVStream, _type AVPacketSideDataType, size uint64) unsafe.Pointer {
	var tmpstream *C.AVStream
	if stream != nil {
		tmpstream = stream.ptr
	}
	ret := C.av_stream_new_side_data(tmpstream, C.enum_AVPacketSideDataType(_type), C.size_t(size))
	return unsafe.Pointer(ret)
}

// --- Function av_stream_get_side_data ---

// AVStreamGetSideData wraps av_stream_get_side_data.
// av_stream_get_side_data skipped due to size

// --- Function av_new_program ---

// AVNewProgram wraps av_new_program.
func AVNewProgram(s *AVFormatContext, id int) *AVProgram {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_new_program(tmps, C.int(id))
	var retMapped *AVProgram
	if ret != nil {
		retMapped = &AVProgram{ptr: ret}
	}
	return retMapped
}

// --- Function avformat_alloc_output_context2 ---

// AVFormatAllocOutputContext2 wraps avformat_alloc_output_context2.
func AVFormatAllocOutputContext2(ctx **AVFormatContext, oformat *AVOutputFormat, formatName *CStr, filename *CStr) int {
	var ptrctx **C.AVFormatContext
	var tmpctx *C.AVFormatContext
	var oldTmpctx *C.AVFormatContext
	if ctx != nil {
		innerctx := *ctx
		if innerctx != nil {
			tmpctx = innerctx.ptr
			oldTmpctx = tmpctx
		}
		ptrctx = &tmpctx
	}
	var tmpoformat *C.AVOutputFormat
	if oformat != nil {
		tmpoformat = oformat.ptr
	}
	var tmpformatName *C.char
	if formatName != nil {
		tmpformatName = formatName.ptr
	}
	var tmpfilename *C.char
	if filename != nil {
		tmpfilename = filename.ptr
	}
	ret := C.avformat_alloc_output_context2(ptrctx, tmpoformat, tmpformatName, tmpfilename)
	if tmpctx != oldTmpctx && ctx != nil {
		if tmpctx != nil {
			*ctx = &AVFormatContext{ptr: tmpctx}
		} else {
			*ctx = nil
		}
	}
	return int(ret)
}

// --- Function av_find_input_format ---

// AVFindInputFormat wraps av_find_input_format.
func AVFindInputFormat(shortName *CStr) *AVInputFormat {
	var tmpshortName *C.char
	if shortName != nil {
		tmpshortName = shortName.ptr
	}
	ret := C.av_find_input_format(tmpshortName)
	var retMapped *AVInputFormat
	if ret != nil {
		retMapped = &AVInputFormat{ptr: ret}
	}
	return retMapped
}

// --- Function av_probe_input_format ---

// AVProbeInputFormat wraps av_probe_input_format.
func AVProbeInputFormat(pd *AVProbeData, isOpened int) *AVInputFormat {
	var tmppd *C.AVProbeData
	if pd != nil {
		tmppd = pd.ptr
	}
	ret := C.av_probe_input_format(tmppd, C.int(isOpened))
	var retMapped *AVInputFormat
	if ret != nil {
		retMapped = &AVInputFormat{ptr: ret}
	}
	return retMapped
}

// --- Function av_probe_input_format2 ---

// AVProbeInputFormat2 wraps av_probe_input_format2.
// av_probe_input_format2 skipped due to scoreMax

// --- Function av_probe_input_format3 ---

// AVProbeInputFormat3 wraps av_probe_input_format3.
// av_probe_input_format3 skipped due to scoreRet

// --- Function av_probe_input_buffer2 ---

// AVProbeInputBuffer2 wraps av_probe_input_buffer2.
func AVProbeInputBuffer2(pb *AVIOContext, fmt **AVInputFormat, url *CStr, logctx unsafe.Pointer, offset uint, maxProbeSize uint) int {
	var tmppb *C.AVIOContext
	if pb != nil {
		tmppb = pb.ptr
	}
	var ptrfmt **C.AVInputFormat
	var tmpfmt *C.AVInputFormat
	var oldTmpfmt *C.AVInputFormat
	if fmt != nil {
		innerfmt := *fmt
		if innerfmt != nil {
			tmpfmt = innerfmt.ptr
			oldTmpfmt = tmpfmt
		}
		ptrfmt = &tmpfmt
	}
	var tmpurl *C.char
	if url != nil {
		tmpurl = url.ptr
	}
	ret := C.av_probe_input_buffer2(tmppb, ptrfmt, tmpurl, logctx, C.uint(offset), C.uint(maxProbeSize))
	if tmpfmt != oldTmpfmt && fmt != nil {
		if tmpfmt != nil {
			*fmt = &AVInputFormat{ptr: tmpfmt}
		} else {
			*fmt = nil
		}
	}
	return int(ret)
}

// --- Function av_probe_input_buffer ---

// AVProbeInputBuffer wraps av_probe_input_buffer.
func AVProbeInputBuffer(pb *AVIOContext, fmt **AVInputFormat, url *CStr, logctx unsafe.Pointer, offset uint, maxProbeSize uint) int {
	var tmppb *C.AVIOContext
	if pb != nil {
		tmppb = pb.ptr
	}
	var ptrfmt **C.AVInputFormat
	var tmpfmt *C.AVInputFormat
	var oldTmpfmt *C.AVInputFormat
	if fmt != nil {
		innerfmt := *fmt
		if innerfmt != nil {
			tmpfmt = innerfmt.ptr
			oldTmpfmt = tmpfmt
		}
		ptrfmt = &tmpfmt
	}
	var tmpurl *C.char
	if url != nil {
		tmpurl = url.ptr
	}
	ret := C.av_probe_input_buffer(tmppb, ptrfmt, tmpurl, logctx, C.uint(offset), C.uint(maxProbeSize))
	if tmpfmt != oldTmpfmt && fmt != nil {
		if tmpfmt != nil {
			*fmt = &AVInputFormat{ptr: tmpfmt}
		} else {
			*fmt = nil
		}
	}
	return int(ret)
}

// --- Function avformat_open_input ---

// AVFormatOpenInput wraps avformat_open_input.
func AVFormatOpenInput(ps **AVFormatContext, url *CStr, fmt *AVInputFormat, options **AVDictionary) int {
	var ptrps **C.AVFormatContext
	var tmpps *C.AVFormatContext
	var oldTmpps *C.AVFormatContext
	if ps != nil {
		innerps := *ps
		if innerps != nil {
			tmpps = innerps.ptr
			oldTmpps = tmpps
		}
		ptrps = &tmpps
	}
	var tmpurl *C.char
	if url != nil {
		tmpurl = url.ptr
	}
	var tmpfmt *C.AVInputFormat
	if fmt != nil {
		tmpfmt = fmt.ptr
	}
	var ptroptions **C.AVDictionary
	var tmpoptions *C.AVDictionary
	var oldTmpoptions *C.AVDictionary
	if options != nil {
		inneroptions := *options
		if inneroptions != nil {
			tmpoptions = inneroptions.ptr
			oldTmpoptions = tmpoptions
		}
		ptroptions = &tmpoptions
	}
	ret := C.avformat_open_input(ptrps, tmpurl, tmpfmt, ptroptions)
	if tmpps != oldTmpps && ps != nil {
		if tmpps != nil {
			*ps = &AVFormatContext{ptr: tmpps}
		} else {
			*ps = nil
		}
	}
	if tmpoptions != oldTmpoptions && options != nil {
		if tmpoptions != nil {
			*options = &AVDictionary{ptr: tmpoptions}
		} else {
			*options = nil
		}
	}
	return int(ret)
}

// --- Function avformat_find_stream_info ---

// AVFormatFindStreamInfo wraps avformat_find_stream_info.
func AVFormatFindStreamInfo(ic *AVFormatContext, options **AVDictionary) int {
	var tmpic *C.AVFormatContext
	if ic != nil {
		tmpic = ic.ptr
	}
	var ptroptions **C.AVDictionary
	var tmpoptions *C.AVDictionary
	var oldTmpoptions *C.AVDictionary
	if options != nil {
		inneroptions := *options
		if inneroptions != nil {
			tmpoptions = inneroptions.ptr
			oldTmpoptions = tmpoptions
		}
		ptroptions = &tmpoptions
	}
	ret := C.avformat_find_stream_info(tmpic, ptroptions)
	if tmpoptions != oldTmpoptions && options != nil {
		if tmpoptions != nil {
			*options = &AVDictionary{ptr: tmpoptions}
		} else {
			*options = nil
		}
	}
	return int(ret)
}

// --- Function av_find_program_from_stream ---

// AVFindProgramFromStream wraps av_find_program_from_stream.
func AVFindProgramFromStream(ic *AVFormatContext, last *AVProgram, s int) *AVProgram {
	var tmpic *C.AVFormatContext
	if ic != nil {
		tmpic = ic.ptr
	}
	var tmplast *C.AVProgram
	if last != nil {
		tmplast = last.ptr
	}
	ret := C.av_find_program_from_stream(tmpic, tmplast, C.int(s))
	var retMapped *AVProgram
	if ret != nil {
		retMapped = &AVProgram{ptr: ret}
	}
	return retMapped
}

// --- Function av_program_add_stream_index ---

// AVProgramAddStreamIndex wraps av_program_add_stream_index.
func AVProgramAddStreamIndex(ac *AVFormatContext, progid int, idx uint) {
	var tmpac *C.AVFormatContext
	if ac != nil {
		tmpac = ac.ptr
	}
	C.av_program_add_stream_index(tmpac, C.int(progid), C.uint(idx))
}

// --- Function av_find_best_stream ---

// AVFindBestStream wraps av_find_best_stream.
func AVFindBestStream(ic *AVFormatContext, _type AVMediaType, wantedStreamNb int, relatedStream int, decoderRet **AVCodec, flags int) int {
	var tmpic *C.AVFormatContext
	if ic != nil {
		tmpic = ic.ptr
	}
	var ptrdecoderRet **C.AVCodec
	var tmpdecoderRet *C.AVCodec
	var oldTmpdecoderRet *C.AVCodec
	if decoderRet != nil {
		innerdecoderRet := *decoderRet
		if innerdecoderRet != nil {
			tmpdecoderRet = innerdecoderRet.ptr
			oldTmpdecoderRet = tmpdecoderRet
		}
		ptrdecoderRet = &tmpdecoderRet
	}
	ret := C.av_find_best_stream(tmpic, C.enum_AVMediaType(_type), C.int(wantedStreamNb), C.int(relatedStream), ptrdecoderRet, C.int(flags))
	if tmpdecoderRet != oldTmpdecoderRet && decoderRet != nil {
		if tmpdecoderRet != nil {
			*decoderRet = &AVCodec{ptr: tmpdecoderRet}
		} else {
			*decoderRet = nil
		}
	}
	return int(ret)
}

// --- Function av_read_frame ---

// AVReadFrame wraps av_read_frame.
func AVReadFrame(s *AVFormatContext, pkt *AVPacket) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_read_frame(tmps, tmppkt)
	return int(ret)
}

// --- Function av_seek_frame ---

// AVSeekFrame wraps av_seek_frame.
func AVSeekFrame(s *AVFormatContext, streamIndex int, timestamp int64, flags int) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_seek_frame(tmps, C.int(streamIndex), C.int64_t(timestamp), C.int(flags))
	return int(ret)
}

// --- Function avformat_seek_file ---

// AVFormatSeekFile wraps avformat_seek_file.
func AVFormatSeekFile(s *AVFormatContext, streamIndex int, minTs int64, ts int64, maxTs int64, flags int) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avformat_seek_file(tmps, C.int(streamIndex), C.int64_t(minTs), C.int64_t(ts), C.int64_t(maxTs), C.int(flags))
	return int(ret)
}

// --- Function avformat_flush ---

// AVFormatFlush wraps avformat_flush.
func AVFormatFlush(s *AVFormatContext) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avformat_flush(tmps)
	return int(ret)
}

// --- Function av_read_play ---

// AVReadPlay wraps av_read_play.
func AVReadPlay(s *AVFormatContext) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_read_play(tmps)
	return int(ret)
}

// --- Function av_read_pause ---

// AVReadPause wraps av_read_pause.
func AVReadPause(s *AVFormatContext) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_read_pause(tmps)
	return int(ret)
}

// --- Function avformat_close_input ---

// AVFormatCloseInput wraps avformat_close_input.
func AVFormatCloseInput(s **AVFormatContext) {
	var ptrs **C.AVFormatContext
	var tmps *C.AVFormatContext
	var oldTmps *C.AVFormatContext
	if s != nil {
		inners := *s
		if inners != nil {
			tmps = inners.ptr
			oldTmps = tmps
		}
		ptrs = &tmps
	}
	C.avformat_close_input(ptrs)
	if tmps != oldTmps && s != nil {
		if tmps != nil {
			*s = &AVFormatContext{ptr: tmps}
		} else {
			*s = nil
		}
	}
}

// --- Function avformat_write_header ---

// AVFormatWriteHeader wraps avformat_write_header.
func AVFormatWriteHeader(s *AVFormatContext, options **AVDictionary) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var ptroptions **C.AVDictionary
	var tmpoptions *C.AVDictionary
	var oldTmpoptions *C.AVDictionary
	if options != nil {
		inneroptions := *options
		if inneroptions != nil {
			tmpoptions = inneroptions.ptr
			oldTmpoptions = tmpoptions
		}
		ptroptions = &tmpoptions
	}
	ret := C.avformat_write_header(tmps, ptroptions)
	if tmpoptions != oldTmpoptions && options != nil {
		if tmpoptions != nil {
			*options = &AVDictionary{ptr: tmpoptions}
		} else {
			*options = nil
		}
	}
	return int(ret)
}

// --- Function avformat_init_output ---

// AVFormatInitOutput wraps avformat_init_output.
func AVFormatInitOutput(s *AVFormatContext, options **AVDictionary) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var ptroptions **C.AVDictionary
	var tmpoptions *C.AVDictionary
	var oldTmpoptions *C.AVDictionary
	if options != nil {
		inneroptions := *options
		if inneroptions != nil {
			tmpoptions = inneroptions.ptr
			oldTmpoptions = tmpoptions
		}
		ptroptions = &tmpoptions
	}
	ret := C.avformat_init_output(tmps, ptroptions)
	if tmpoptions != oldTmpoptions && options != nil {
		if tmpoptions != nil {
			*options = &AVDictionary{ptr: tmpoptions}
		} else {
			*options = nil
		}
	}
	return int(ret)
}

// --- Function av_write_frame ---

// AVWriteFrame wraps av_write_frame.
func AVWriteFrame(s *AVFormatContext, pkt *AVPacket) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_write_frame(tmps, tmppkt)
	return int(ret)
}

// --- Function av_interleaved_write_frame ---

// AVInterleavedWriteFrame wraps av_interleaved_write_frame.
func AVInterleavedWriteFrame(s *AVFormatContext, pkt *AVPacket) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_interleaved_write_frame(tmps, tmppkt)
	return int(ret)
}

// --- Function av_write_uncoded_frame ---

// AVWriteUncodedFrame wraps av_write_uncoded_frame.
func AVWriteUncodedFrame(s *AVFormatContext, streamIndex int, frame *AVFrame) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_write_uncoded_frame(tmps, C.int(streamIndex), tmpframe)
	return int(ret)
}

// --- Function av_interleaved_write_uncoded_frame ---

// AVInterleavedWriteUncodedFrame wraps av_interleaved_write_uncoded_frame.
func AVInterleavedWriteUncodedFrame(s *AVFormatContext, streamIndex int, frame *AVFrame) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_interleaved_write_uncoded_frame(tmps, C.int(streamIndex), tmpframe)
	return int(ret)
}

// --- Function av_write_uncoded_frame_query ---

// AVWriteUncodedFrameQuery wraps av_write_uncoded_frame_query.
func AVWriteUncodedFrameQuery(s *AVFormatContext, streamIndex int) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_write_uncoded_frame_query(tmps, C.int(streamIndex))
	return int(ret)
}

// --- Function av_write_trailer ---

// AVWriteTrailer wraps av_write_trailer.
func AVWriteTrailer(s *AVFormatContext) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_write_trailer(tmps)
	return int(ret)
}

// --- Function av_guess_format ---

// AVGuessFormat wraps av_guess_format.
func AVGuessFormat(shortName *CStr, filename *CStr, mimeType *CStr) *AVOutputFormat {
	var tmpshortName *C.char
	if shortName != nil {
		tmpshortName = shortName.ptr
	}
	var tmpfilename *C.char
	if filename != nil {
		tmpfilename = filename.ptr
	}
	var tmpmimeType *C.char
	if mimeType != nil {
		tmpmimeType = mimeType.ptr
	}
	ret := C.av_guess_format(tmpshortName, tmpfilename, tmpmimeType)
	var retMapped *AVOutputFormat
	if ret != nil {
		retMapped = &AVOutputFormat{ptr: ret}
	}
	return retMapped
}

// --- Function av_guess_codec ---

// AVGuessCodec wraps av_guess_codec.
func AVGuessCodec(fmt *AVOutputFormat, shortName *CStr, filename *CStr, mimeType *CStr, _type AVMediaType) AVCodecID {
	var tmpfmt *C.AVOutputFormat
	if fmt != nil {
		tmpfmt = fmt.ptr
	}
	var tmpshortName *C.char
	if shortName != nil {
		tmpshortName = shortName.ptr
	}
	var tmpfilename *C.char
	if filename != nil {
		tmpfilename = filename.ptr
	}
	var tmpmimeType *C.char
	if mimeType != nil {
		tmpmimeType = mimeType.ptr
	}
	ret := C.av_guess_codec(tmpfmt, tmpshortName, tmpfilename, tmpmimeType, C.enum_AVMediaType(_type))
	return AVCodecID(ret)
}

// --- Function av_get_output_timestamp ---

// AVGetOutputTimestamp wraps av_get_output_timestamp.
// av_get_output_timestamp skipped due to dts

// --- Function av_hex_dump ---

// av_hex_dump skipped due to f.

// --- Function av_hex_dump_log ---

// AVHexDumpLog wraps av_hex_dump_log.
func AVHexDumpLog(avcl unsafe.Pointer, level int, buf unsafe.Pointer, size int) {
	C.av_hex_dump_log(avcl, C.int(level), (*C.uint8_t)(buf), C.int(size))
}

// --- Function av_pkt_dump2 ---

// av_pkt_dump2 skipped due to f.

// --- Function av_pkt_dump_log2 ---

// AVPktDumpLog2 wraps av_pkt_dump_log2.
func AVPktDumpLog2(avcl unsafe.Pointer, level int, pkt *AVPacket, dumpPayload int, st *AVStream) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	C.av_pkt_dump_log2(avcl, C.int(level), tmppkt, C.int(dumpPayload), tmpst)
}

// --- Function av_codec_get_id ---

// AVCodecGetId wraps av_codec_get_id.
func AVCodecGetId(tags **AVCodecTag, tag uint) AVCodecID {
	var ptrtags **C.struct_AVCodecTag
	var tmptags *C.struct_AVCodecTag
	var oldTmptags *C.struct_AVCodecTag
	if tags != nil {
		innertags := *tags
		if innertags != nil {
			tmptags = innertags.ptr
			oldTmptags = tmptags
		}
		ptrtags = &tmptags
	}
	ret := C.av_codec_get_id(ptrtags, C.uint(tag))
	if tmptags != oldTmptags && tags != nil {
		if tmptags != nil {
			*tags = &AVCodecTag{ptr: tmptags}
		} else {
			*tags = nil
		}
	}
	return AVCodecID(ret)
}

// --- Function av_codec_get_tag ---

// AVCodecGetTag wraps av_codec_get_tag.
func AVCodecGetTag(tags **AVCodecTag, id AVCodecID) uint {
	var ptrtags **C.struct_AVCodecTag
	var tmptags *C.struct_AVCodecTag
	var oldTmptags *C.struct_AVCodecTag
	if tags != nil {
		innertags := *tags
		if innertags != nil {
			tmptags = innertags.ptr
			oldTmptags = tmptags
		}
		ptrtags = &tmptags
	}
	ret := C.av_codec_get_tag(ptrtags, C.enum_AVCodecID(id))
	if tmptags != oldTmptags && tags != nil {
		if tmptags != nil {
			*tags = &AVCodecTag{ptr: tmptags}
		} else {
			*tags = nil
		}
	}
	return uint(ret)
}

// --- Function av_codec_get_tag2 ---

// AVCodecGetTag2 wraps av_codec_get_tag2.
// av_codec_get_tag2 skipped due to tag

// --- Function av_find_default_stream_index ---

// AVFindDefaultStreamIndex wraps av_find_default_stream_index.
func AVFindDefaultStreamIndex(s *AVFormatContext) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_find_default_stream_index(tmps)
	return int(ret)
}

// --- Function av_index_search_timestamp ---

// AVIndexSearchTimestamp wraps av_index_search_timestamp.
func AVIndexSearchTimestamp(st *AVStream, timestamp int64, flags int) int {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.av_index_search_timestamp(tmpst, C.int64_t(timestamp), C.int(flags))
	return int(ret)
}

// --- Function avformat_index_get_entries_count ---

// AVFormatIndexGetEntriesCount wraps avformat_index_get_entries_count.
func AVFormatIndexGetEntriesCount(st *AVStream) int {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.avformat_index_get_entries_count(tmpst)
	return int(ret)
}

// --- Function avformat_index_get_entry ---

// AVFormatIndexGetEntry wraps avformat_index_get_entry.
func AVFormatIndexGetEntry(st *AVStream, idx int) *AVIndexEntry {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.avformat_index_get_entry(tmpst, C.int(idx))
	var retMapped *AVIndexEntry
	if ret != nil {
		retMapped = &AVIndexEntry{ptr: ret}
	}
	return retMapped
}

// --- Function avformat_index_get_entry_from_timestamp ---

// AVFormatIndexGetEntryFromTimestamp wraps avformat_index_get_entry_from_timestamp.
func AVFormatIndexGetEntryFromTimestamp(st *AVStream, wantedTimestamp int64, flags int) *AVIndexEntry {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.avformat_index_get_entry_from_timestamp(tmpst, C.int64_t(wantedTimestamp), C.int(flags))
	var retMapped *AVIndexEntry
	if ret != nil {
		retMapped = &AVIndexEntry{ptr: ret}
	}
	return retMapped
}

// --- Function av_add_index_entry ---

// AVAddIndexEntry wraps av_add_index_entry.
func AVAddIndexEntry(st *AVStream, pos int64, timestamp int64, size int, distance int, flags int) int {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.av_add_index_entry(tmpst, C.int64_t(pos), C.int64_t(timestamp), C.int(size), C.int(distance), C.int(flags))
	return int(ret)
}

// --- Function av_url_split ---

// AVUrlSplit wraps av_url_split.
// av_url_split skipped due to portPtr

// --- Function av_dump_format ---

// AVDumpFormat wraps av_dump_format.
func AVDumpFormat(ic *AVFormatContext, index int, url *CStr, isOutput int) {
	var tmpic *C.AVFormatContext
	if ic != nil {
		tmpic = ic.ptr
	}
	var tmpurl *C.char
	if url != nil {
		tmpurl = url.ptr
	}
	C.av_dump_format(tmpic, C.int(index), tmpurl, C.int(isOutput))
}

// --- Function av_get_frame_filename2 ---

// AVGetFrameFilename2 wraps av_get_frame_filename2.
func AVGetFrameFilename2(buf *CStr, bufSize int, path *CStr, number int, flags int) int {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	var tmppath *C.char
	if path != nil {
		tmppath = path.ptr
	}
	ret := C.av_get_frame_filename2(tmpbuf, C.int(bufSize), tmppath, C.int(number), C.int(flags))
	return int(ret)
}

// --- Function av_get_frame_filename ---

// AVGetFrameFilename wraps av_get_frame_filename.
func AVGetFrameFilename(buf *CStr, bufSize int, path *CStr, number int) int {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	var tmppath *C.char
	if path != nil {
		tmppath = path.ptr
	}
	ret := C.av_get_frame_filename(tmpbuf, C.int(bufSize), tmppath, C.int(number))
	return int(ret)
}

// --- Function av_filename_number_test ---

// AVFilenameNumberTest wraps av_filename_number_test.
func AVFilenameNumberTest(filename *CStr) int {
	var tmpfilename *C.char
	if filename != nil {
		tmpfilename = filename.ptr
	}
	ret := C.av_filename_number_test(tmpfilename)
	return int(ret)
}

// --- Function av_sdp_create ---

// AVSdpCreate wraps av_sdp_create.
// av_sdp_create skipped due to ac

// --- Function av_match_ext ---

// AVMatchExt wraps av_match_ext.
func AVMatchExt(filename *CStr, extensions *CStr) int {
	var tmpfilename *C.char
	if filename != nil {
		tmpfilename = filename.ptr
	}
	var tmpextensions *C.char
	if extensions != nil {
		tmpextensions = extensions.ptr
	}
	ret := C.av_match_ext(tmpfilename, tmpextensions)
	return int(ret)
}

// --- Function avformat_query_codec ---

// AVFormatQueryCodec wraps avformat_query_codec.
func AVFormatQueryCodec(ofmt *AVOutputFormat, codecId AVCodecID, stdCompliance int) int {
	var tmpofmt *C.AVOutputFormat
	if ofmt != nil {
		tmpofmt = ofmt.ptr
	}
	ret := C.avformat_query_codec(tmpofmt, C.enum_AVCodecID(codecId), C.int(stdCompliance))
	return int(ret)
}

// --- Function avformat_get_riff_video_tags ---

// AVFormatGetRiffVideoTags wraps avformat_get_riff_video_tags.
func AVFormatGetRiffVideoTags() *AVCodecTag {
	ret := C.avformat_get_riff_video_tags()
	var retMapped *AVCodecTag
	if ret != nil {
		retMapped = &AVCodecTag{ptr: ret}
	}
	return retMapped
}

// --- Function avformat_get_riff_audio_tags ---

// AVFormatGetRiffAudioTags wraps avformat_get_riff_audio_tags.
func AVFormatGetRiffAudioTags() *AVCodecTag {
	ret := C.avformat_get_riff_audio_tags()
	var retMapped *AVCodecTag
	if ret != nil {
		retMapped = &AVCodecTag{ptr: ret}
	}
	return retMapped
}

// --- Function avformat_get_mov_video_tags ---

// AVFormatGetMovVideoTags wraps avformat_get_mov_video_tags.
func AVFormatGetMovVideoTags() *AVCodecTag {
	ret := C.avformat_get_mov_video_tags()
	var retMapped *AVCodecTag
	if ret != nil {
		retMapped = &AVCodecTag{ptr: ret}
	}
	return retMapped
}

// --- Function avformat_get_mov_audio_tags ---

// AVFormatGetMovAudioTags wraps avformat_get_mov_audio_tags.
func AVFormatGetMovAudioTags() *AVCodecTag {
	ret := C.avformat_get_mov_audio_tags()
	var retMapped *AVCodecTag
	if ret != nil {
		retMapped = &AVCodecTag{ptr: ret}
	}
	return retMapped
}

// --- Function av_guess_sample_aspect_ratio ---

// AVGuessSampleAspectRatio wraps av_guess_sample_aspect_ratio.
func AVGuessSampleAspectRatio(format *AVFormatContext, stream *AVStream, frame *AVFrame) *AVRational {
	var tmpformat *C.AVFormatContext
	if format != nil {
		tmpformat = format.ptr
	}
	var tmpstream *C.AVStream
	if stream != nil {
		tmpstream = stream.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_guess_sample_aspect_ratio(tmpformat, tmpstream, tmpframe)
	return &AVRational{value: ret}
}

// --- Function av_guess_frame_rate ---

// AVGuessFrameRate wraps av_guess_frame_rate.
func AVGuessFrameRate(ctx *AVFormatContext, stream *AVStream, frame *AVFrame) *AVRational {
	var tmpctx *C.AVFormatContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpstream *C.AVStream
	if stream != nil {
		tmpstream = stream.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_guess_frame_rate(tmpctx, tmpstream, tmpframe)
	return &AVRational{value: ret}
}

// --- Function avformat_match_stream_specifier ---

// AVFormatMatchStreamSpecifier wraps avformat_match_stream_specifier.
func AVFormatMatchStreamSpecifier(s *AVFormatContext, st *AVStream, spec *CStr) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	var tmpspec *C.char
	if spec != nil {
		tmpspec = spec.ptr
	}
	ret := C.avformat_match_stream_specifier(tmps, tmpst, tmpspec)
	return int(ret)
}

// --- Function avformat_queue_attached_pictures ---

// AVFormatQueueAttachedPictures wraps avformat_queue_attached_pictures.
func AVFormatQueueAttachedPictures(s *AVFormatContext) int {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avformat_queue_attached_pictures(tmps)
	return int(ret)
}

// --- Function avformat_transfer_internal_stream_timing_info ---

// AVFormatTransferInternalStreamTimingInfo wraps avformat_transfer_internal_stream_timing_info.
func AVFormatTransferInternalStreamTimingInfo(ofmt *AVOutputFormat, ost *AVStream, ist *AVStream, copyTb AVTimebaseSource) int {
	var tmpofmt *C.AVOutputFormat
	if ofmt != nil {
		tmpofmt = ofmt.ptr
	}
	var tmpost *C.AVStream
	if ost != nil {
		tmpost = ost.ptr
	}
	var tmpist *C.AVStream
	if ist != nil {
		tmpist = ist.ptr
	}
	ret := C.avformat_transfer_internal_stream_timing_info(tmpofmt, tmpost, tmpist, C.enum_AVTimebaseSource(copyTb))
	return int(ret)
}

// --- Function av_stream_get_codec_timebase ---

// AVStreamGetCodecTimebase wraps av_stream_get_codec_timebase.
func AVStreamGetCodecTimebase(st *AVStream) *AVRational {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.av_stream_get_codec_timebase(tmpst)
	return &AVRational{value: ret}
}

// --- Function avio_find_protocol_name ---

// AVIOFindProtocolName wraps avio_find_protocol_name.
func AVIOFindProtocolName(url *CStr) *CStr {
	var tmpurl *C.char
	if url != nil {
		tmpurl = url.ptr
	}
	ret := C.avio_find_protocol_name(tmpurl)
	return wrapCStr(ret)
}

// --- Function avio_check ---

// AVIOCheck wraps avio_check.
func AVIOCheck(url *CStr, flags int) int {
	var tmpurl *C.char
	if url != nil {
		tmpurl = url.ptr
	}
	ret := C.avio_check(tmpurl, C.int(flags))
	return int(ret)
}

// --- Function avio_open_dir ---

// AVIOOpenDir wraps avio_open_dir.
func AVIOOpenDir(s **AVIODirContext, url *CStr, options **AVDictionary) int {
	var ptrs **C.AVIODirContext
	var tmps *C.AVIODirContext
	var oldTmps *C.AVIODirContext
	if s != nil {
		inners := *s
		if inners != nil {
			tmps = inners.ptr
			oldTmps = tmps
		}
		ptrs = &tmps
	}
	var tmpurl *C.char
	if url != nil {
		tmpurl = url.ptr
	}
	var ptroptions **C.AVDictionary
	var tmpoptions *C.AVDictionary
	var oldTmpoptions *C.AVDictionary
	if options != nil {
		inneroptions := *options
		if inneroptions != nil {
			tmpoptions = inneroptions.ptr
			oldTmpoptions = tmpoptions
		}
		ptroptions = &tmpoptions
	}
	ret := C.avio_open_dir(ptrs, tmpurl, ptroptions)
	if tmps != oldTmps && s != nil {
		if tmps != nil {
			*s = &AVIODirContext{ptr: tmps}
		} else {
			*s = nil
		}
	}
	if tmpoptions != oldTmpoptions && options != nil {
		if tmpoptions != nil {
			*options = &AVDictionary{ptr: tmpoptions}
		} else {
			*options = nil
		}
	}
	return int(ret)
}

// --- Function avio_read_dir ---

// AVIOReadDir wraps avio_read_dir.
func AVIOReadDir(s *AVIODirContext, next **AVIODirEntry) int {
	var tmps *C.AVIODirContext
	if s != nil {
		tmps = s.ptr
	}
	var ptrnext **C.AVIODirEntry
	var tmpnext *C.AVIODirEntry
	var oldTmpnext *C.AVIODirEntry
	if next != nil {
		innernext := *next
		if innernext != nil {
			tmpnext = innernext.ptr
			oldTmpnext = tmpnext
		}
		ptrnext = &tmpnext
	}
	ret := C.avio_read_dir(tmps, ptrnext)
	if tmpnext != oldTmpnext && next != nil {
		if tmpnext != nil {
			*next = &AVIODirEntry{ptr: tmpnext}
		} else {
			*next = nil
		}
	}
	return int(ret)
}

// --- Function avio_close_dir ---

// AVIOCloseDir wraps avio_close_dir.
func AVIOCloseDir(s **AVIODirContext) int {
	var ptrs **C.AVIODirContext
	var tmps *C.AVIODirContext
	var oldTmps *C.AVIODirContext
	if s != nil {
		inners := *s
		if inners != nil {
			tmps = inners.ptr
			oldTmps = tmps
		}
		ptrs = &tmps
	}
	ret := C.avio_close_dir(ptrs)
	if tmps != oldTmps && s != nil {
		if tmps != nil {
			*s = &AVIODirContext{ptr: tmps}
		} else {
			*s = nil
		}
	}
	return int(ret)
}

// --- Function avio_free_directory_entry ---

// AVIOFreeDirectoryEntry wraps avio_free_directory_entry.
func AVIOFreeDirectoryEntry(entry **AVIODirEntry) {
	var ptrentry **C.AVIODirEntry
	var tmpentry *C.AVIODirEntry
	var oldTmpentry *C.AVIODirEntry
	if entry != nil {
		innerentry := *entry
		if innerentry != nil {
			tmpentry = innerentry.ptr
			oldTmpentry = tmpentry
		}
		ptrentry = &tmpentry
	}
	C.avio_free_directory_entry(ptrentry)
	if tmpentry != oldTmpentry && entry != nil {
		if tmpentry != nil {
			*entry = &AVIODirEntry{ptr: tmpentry}
		} else {
			*entry = nil
		}
	}
}

// --- Function avio_alloc_context ---

// avio_alloc_context skipped due to read_packet.

// --- Function avio_context_free ---

// AVIOContextFree wraps avio_context_free.
func AVIOContextFree(s **AVIOContext) {
	var ptrs **C.AVIOContext
	var tmps *C.AVIOContext
	var oldTmps *C.AVIOContext
	if s != nil {
		inners := *s
		if inners != nil {
			tmps = inners.ptr
			oldTmps = tmps
		}
		ptrs = &tmps
	}
	C.avio_context_free(ptrs)
	if tmps != oldTmps && s != nil {
		if tmps != nil {
			*s = &AVIOContext{ptr: tmps}
		} else {
			*s = nil
		}
	}
}

// --- Function avio_w8 ---

// AVIOW8 wraps avio_w8.
func AVIOW8(s *AVIOContext, b int) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_w8(tmps, C.int(b))
}

// --- Function avio_write ---

// AVIOWrite wraps avio_write.
func AVIOWrite(s *AVIOContext, buf unsafe.Pointer, size int) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_write(tmps, (*C.uchar)(buf), C.int(size))
}

// --- Function avio_wl64 ---

// AVIOWl64 wraps avio_wl64.
func AVIOWl64(s *AVIOContext, val uint64) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_wl64(tmps, C.uint64_t(val))
}

// --- Function avio_wb64 ---

// AVIOWb64 wraps avio_wb64.
func AVIOWb64(s *AVIOContext, val uint64) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_wb64(tmps, C.uint64_t(val))
}

// --- Function avio_wl32 ---

// AVIOWl32 wraps avio_wl32.
func AVIOWl32(s *AVIOContext, val uint) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_wl32(tmps, C.uint(val))
}

// --- Function avio_wb32 ---

// AVIOWb32 wraps avio_wb32.
func AVIOWb32(s *AVIOContext, val uint) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_wb32(tmps, C.uint(val))
}

// --- Function avio_wl24 ---

// AVIOWl24 wraps avio_wl24.
func AVIOWl24(s *AVIOContext, val uint) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_wl24(tmps, C.uint(val))
}

// --- Function avio_wb24 ---

// AVIOWb24 wraps avio_wb24.
func AVIOWb24(s *AVIOContext, val uint) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_wb24(tmps, C.uint(val))
}

// --- Function avio_wl16 ---

// AVIOWl16 wraps avio_wl16.
func AVIOWl16(s *AVIOContext, val uint) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_wl16(tmps, C.uint(val))
}

// --- Function avio_wb16 ---

// AVIOWb16 wraps avio_wb16.
func AVIOWb16(s *AVIOContext, val uint) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_wb16(tmps, C.uint(val))
}

// --- Function avio_put_str ---

// AVIOPutStr wraps avio_put_str.
func AVIOPutStr(s *AVIOContext, str *CStr) int {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpstr *C.char
	if str != nil {
		tmpstr = str.ptr
	}
	ret := C.avio_put_str(tmps, tmpstr)
	return int(ret)
}

// --- Function avio_put_str16le ---

// AVIOPutStr16Le wraps avio_put_str16le.
func AVIOPutStr16Le(s *AVIOContext, str *CStr) int {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpstr *C.char
	if str != nil {
		tmpstr = str.ptr
	}
	ret := C.avio_put_str16le(tmps, tmpstr)
	return int(ret)
}

// --- Function avio_put_str16be ---

// AVIOPutStr16Be wraps avio_put_str16be.
func AVIOPutStr16Be(s *AVIOContext, str *CStr) int {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpstr *C.char
	if str != nil {
		tmpstr = str.ptr
	}
	ret := C.avio_put_str16be(tmps, tmpstr)
	return int(ret)
}

// --- Function avio_write_marker ---

// AVIOWriteMarker wraps avio_write_marker.
func AVIOWriteMarker(s *AVIOContext, time int64, _type AVIODataMarkerType) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_write_marker(tmps, C.int64_t(time), C.enum_AVIODataMarkerType(_type))
}

// --- Function avio_seek ---

// AVIOSeek wraps avio_seek.
func AVIOSeek(s *AVIOContext, offset int64, whence int) int64 {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_seek(tmps, C.int64_t(offset), C.int(whence))
	return int64(ret)
}

// --- Function avio_skip ---

// AVIOSkip wraps avio_skip.
func AVIOSkip(s *AVIOContext, offset int64) int64 {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_skip(tmps, C.int64_t(offset))
	return int64(ret)
}

// --- Function avio_tell ---

// AVIOTell wraps avio_tell.
func AVIOTell(s *AVIOContext) int64 {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_tell(tmps)
	return int64(ret)
}

// --- Function avio_size ---

// AVIOSize wraps avio_size.
func AVIOSize(s *AVIOContext) int64 {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_size(tmps)
	return int64(ret)
}

// --- Function avio_feof ---

// AVIOFeof wraps avio_feof.
func AVIOFeof(s *AVIOContext) int {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_feof(tmps)
	return int(ret)
}

// --- Function avio_vprintf ---

// avio_vprintf skipped due to ap.

// --- Function avio_printf ---

// avio_printf skipped due to variadic arg.

// --- Function avio_print_string_array ---

// AVIOPrintStringArray wraps avio_print_string_array.
// avio_print_string_array skipped due to strings

// --- Function avio_flush ---

// AVIOFlush wraps avio_flush.
func AVIOFlush(s *AVIOContext) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_flush(tmps)
}

// --- Function avio_read ---

// AVIORead wraps avio_read.
func AVIORead(s *AVIOContext, buf unsafe.Pointer, size int) int {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_read(tmps, (*C.uchar)(buf), C.int(size))
	return int(ret)
}

// --- Function avio_read_partial ---

// AVIOReadPartial wraps avio_read_partial.
func AVIOReadPartial(s *AVIOContext, buf unsafe.Pointer, size int) int {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_read_partial(tmps, (*C.uchar)(buf), C.int(size))
	return int(ret)
}

// --- Function avio_r8 ---

// AVIOR8 wraps avio_r8.
func AVIOR8(s *AVIOContext) int {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_r8(tmps)
	return int(ret)
}

// --- Function avio_rl16 ---

// AVIORl16 wraps avio_rl16.
func AVIORl16(s *AVIOContext) uint {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_rl16(tmps)
	return uint(ret)
}

// --- Function avio_rl24 ---

// AVIORl24 wraps avio_rl24.
func AVIORl24(s *AVIOContext) uint {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_rl24(tmps)
	return uint(ret)
}

// --- Function avio_rl32 ---

// AVIORl32 wraps avio_rl32.
func AVIORl32(s *AVIOContext) uint {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_rl32(tmps)
	return uint(ret)
}

// --- Function avio_rl64 ---

// AVIORl64 wraps avio_rl64.
func AVIORl64(s *AVIOContext) uint64 {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_rl64(tmps)
	return uint64(ret)
}

// --- Function avio_rb16 ---

// AVIORb16 wraps avio_rb16.
func AVIORb16(s *AVIOContext) uint {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_rb16(tmps)
	return uint(ret)
}

// --- Function avio_rb24 ---

// AVIORb24 wraps avio_rb24.
func AVIORb24(s *AVIOContext) uint {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_rb24(tmps)
	return uint(ret)
}

// --- Function avio_rb32 ---

// AVIORb32 wraps avio_rb32.
func AVIORb32(s *AVIOContext) uint {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_rb32(tmps)
	return uint(ret)
}

// --- Function avio_rb64 ---

// AVIORb64 wraps avio_rb64.
func AVIORb64(s *AVIOContext) uint64 {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_rb64(tmps)
	return uint64(ret)
}

// --- Function avio_get_str ---

// AVIOGetStr wraps avio_get_str.
func AVIOGetStr(pb *AVIOContext, maxlen int, buf *CStr, buflen int) int {
	var tmppb *C.AVIOContext
	if pb != nil {
		tmppb = pb.ptr
	}
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.avio_get_str(tmppb, C.int(maxlen), tmpbuf, C.int(buflen))
	return int(ret)
}

// --- Function avio_get_str16le ---

// AVIOGetStr16Le wraps avio_get_str16le.
func AVIOGetStr16Le(pb *AVIOContext, maxlen int, buf *CStr, buflen int) int {
	var tmppb *C.AVIOContext
	if pb != nil {
		tmppb = pb.ptr
	}
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.avio_get_str16le(tmppb, C.int(maxlen), tmpbuf, C.int(buflen))
	return int(ret)
}

// --- Function avio_get_str16be ---

// AVIOGetStr16Be wraps avio_get_str16be.
func AVIOGetStr16Be(pb *AVIOContext, maxlen int, buf *CStr, buflen int) int {
	var tmppb *C.AVIOContext
	if pb != nil {
		tmppb = pb.ptr
	}
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.avio_get_str16be(tmppb, C.int(maxlen), tmpbuf, C.int(buflen))
	return int(ret)
}

// --- Function avio_open ---

// AVIOOpen wraps avio_open.
func AVIOOpen(s **AVIOContext, url *CStr, flags int) int {
	var ptrs **C.AVIOContext
	var tmps *C.AVIOContext
	var oldTmps *C.AVIOContext
	if s != nil {
		inners := *s
		if inners != nil {
			tmps = inners.ptr
			oldTmps = tmps
		}
		ptrs = &tmps
	}
	var tmpurl *C.char
	if url != nil {
		tmpurl = url.ptr
	}
	ret := C.avio_open(ptrs, tmpurl, C.int(flags))
	if tmps != oldTmps && s != nil {
		if tmps != nil {
			*s = &AVIOContext{ptr: tmps}
		} else {
			*s = nil
		}
	}
	return int(ret)
}

// --- Function avio_open2 ---

// AVIOOpen2 wraps avio_open2.
func AVIOOpen2(s **AVIOContext, url *CStr, flags int, intCb *AVIOInterruptCB, options **AVDictionary) int {
	var ptrs **C.AVIOContext
	var tmps *C.AVIOContext
	var oldTmps *C.AVIOContext
	if s != nil {
		inners := *s
		if inners != nil {
			tmps = inners.ptr
			oldTmps = tmps
		}
		ptrs = &tmps
	}
	var tmpurl *C.char
	if url != nil {
		tmpurl = url.ptr
	}
	var tmpintCb *C.AVIOInterruptCB
	if intCb != nil {
		tmpintCb = intCb.ptr
	}
	var ptroptions **C.AVDictionary
	var tmpoptions *C.AVDictionary
	var oldTmpoptions *C.AVDictionary
	if options != nil {
		inneroptions := *options
		if inneroptions != nil {
			tmpoptions = inneroptions.ptr
			oldTmpoptions = tmpoptions
		}
		ptroptions = &tmpoptions
	}
	ret := C.avio_open2(ptrs, tmpurl, C.int(flags), tmpintCb, ptroptions)
	if tmps != oldTmps && s != nil {
		if tmps != nil {
			*s = &AVIOContext{ptr: tmps}
		} else {
			*s = nil
		}
	}
	if tmpoptions != oldTmpoptions && options != nil {
		if tmpoptions != nil {
			*options = &AVDictionary{ptr: tmpoptions}
		} else {
			*options = nil
		}
	}
	return int(ret)
}

// --- Function avio_close ---

// AVIOClose wraps avio_close.
func AVIOClose(s *AVIOContext) int {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_close(tmps)
	return int(ret)
}

// --- Function avio_closep ---

// AVIOClosep wraps avio_closep.
func AVIOClosep(s **AVIOContext) int {
	var ptrs **C.AVIOContext
	var tmps *C.AVIOContext
	var oldTmps *C.AVIOContext
	if s != nil {
		inners := *s
		if inners != nil {
			tmps = inners.ptr
			oldTmps = tmps
		}
		ptrs = &tmps
	}
	ret := C.avio_closep(ptrs)
	if tmps != oldTmps && s != nil {
		if tmps != nil {
			*s = &AVIOContext{ptr: tmps}
		} else {
			*s = nil
		}
	}
	return int(ret)
}

// --- Function avio_open_dyn_buf ---

// AVIOOpenDynBuf wraps avio_open_dyn_buf.
func AVIOOpenDynBuf(s **AVIOContext) int {
	var ptrs **C.AVIOContext
	var tmps *C.AVIOContext
	var oldTmps *C.AVIOContext
	if s != nil {
		inners := *s
		if inners != nil {
			tmps = inners.ptr
			oldTmps = tmps
		}
		ptrs = &tmps
	}
	ret := C.avio_open_dyn_buf(ptrs)
	if tmps != oldTmps && s != nil {
		if tmps != nil {
			*s = &AVIOContext{ptr: tmps}
		} else {
			*s = nil
		}
	}
	return int(ret)
}

// --- Function avio_get_dyn_buf ---

// AVIOGetDynBuf wraps avio_get_dyn_buf.
// avio_get_dyn_buf skipped due to pbuffer

// --- Function avio_close_dyn_buf ---

// AVIOCloseDynBuf wraps avio_close_dyn_buf.
// avio_close_dyn_buf skipped due to pbuffer

// --- Function avio_enum_protocols ---

// AVIOEnumProtocols wraps avio_enum_protocols.
// avio_enum_protocols skipped due to opaque

// --- Function avio_protocol_get_class ---

// AVIOProtocolGetClass wraps avio_protocol_get_class.
func AVIOProtocolGetClass(name *CStr) *AVClass {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.avio_protocol_get_class(tmpname)
	var retMapped *AVClass
	if ret != nil {
		retMapped = &AVClass{ptr: ret}
	}
	return retMapped
}

// --- Function avio_pause ---

// AVIOPause wraps avio_pause.
func AVIOPause(h *AVIOContext, pause int) int {
	var tmph *C.AVIOContext
	if h != nil {
		tmph = h.ptr
	}
	ret := C.avio_pause(tmph, C.int(pause))
	return int(ret)
}

// --- Function avio_seek_time ---

// AVIOSeekTime wraps avio_seek_time.
func AVIOSeekTime(h *AVIOContext, streamIndex int, timestamp int64, flags int) int64 {
	var tmph *C.AVIOContext
	if h != nil {
		tmph = h.ptr
	}
	ret := C.avio_seek_time(tmph, C.int(streamIndex), C.int64_t(timestamp), C.int(flags))
	return int64(ret)
}

// --- Function avio_read_to_bprint ---

// AVIOReadToBprint wraps avio_read_to_bprint.
func AVIOReadToBprint(h *AVIOContext, pb *AVBPrint, maxSize uint64) int {
	var tmph *C.AVIOContext
	if h != nil {
		tmph = h.ptr
	}
	var tmppb *C.struct_AVBPrint
	if pb != nil {
		tmppb = pb.ptr
	}
	ret := C.avio_read_to_bprint(tmph, tmppb, C.size_t(maxSize))
	return int(ret)
}

// --- Function avio_accept ---

// AVIOAccept wraps avio_accept.
func AVIOAccept(s *AVIOContext, c **AVIOContext) int {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	var ptrc **C.AVIOContext
	var tmpc *C.AVIOContext
	var oldTmpc *C.AVIOContext
	if c != nil {
		innerc := *c
		if innerc != nil {
			tmpc = innerc.ptr
			oldTmpc = tmpc
		}
		ptrc = &tmpc
	}
	ret := C.avio_accept(tmps, ptrc)
	if tmpc != oldTmpc && c != nil {
		if tmpc != nil {
			*c = &AVIOContext{ptr: tmpc}
		} else {
			*c = nil
		}
	}
	return int(ret)
}

// --- Function avio_handshake ---

// AVIOHandshake wraps avio_handshake.
func AVIOHandshake(c *AVIOContext) int {
	var tmpc *C.AVIOContext
	if c != nil {
		tmpc = c.ptr
	}
	ret := C.avio_handshake(tmpc)
	return int(ret)
}

// --- Function avutil_version ---

// AVUtilVersion wraps avutil_version.
func AVUtilVersion() uint {
	ret := C.avutil_version()
	return uint(ret)
}

// --- Function av_version_info ---

// AVVersionInfo wraps av_version_info.
func AVVersionInfo() *CStr {
	ret := C.av_version_info()
	return wrapCStr(ret)
}

// --- Function avutil_configuration ---

// AVUtilConfiguration wraps avutil_configuration.
func AVUtilConfiguration() *CStr {
	ret := C.avutil_configuration()
	return wrapCStr(ret)
}

// --- Function avutil_license ---

// AVUtilLicense wraps avutil_license.
func AVUtilLicense() *CStr {
	ret := C.avutil_license()
	return wrapCStr(ret)
}

// --- Function av_get_media_type_string ---

// AVGetMediaTypeString wraps av_get_media_type_string.
func AVGetMediaTypeString(mediaType AVMediaType) *CStr {
	ret := C.av_get_media_type_string(C.enum_AVMediaType(mediaType))
	return wrapCStr(ret)
}

// --- Function av_get_picture_type_char ---

// AVGetPictureTypeChar wraps av_get_picture_type_char.
func AVGetPictureTypeChar(pictType AVPictureType) uint8 {
	ret := C.av_get_picture_type_char(C.enum_AVPictureType(pictType))
	return uint8(ret)
}

// --- Function av_x_if_null ---

// AVXIfNull wraps av_x_if_null.
func AVXIfNull(p unsafe.Pointer, x unsafe.Pointer) unsafe.Pointer {
	ret := C.av_x_if_null(p, x)
	return ret
}

// --- Function av_int_list_length_for_size ---

// AVIntListLengthForSize wraps av_int_list_length_for_size.
func AVIntListLengthForSize(elsize uint, list unsafe.Pointer, term uint64) uint {
	ret := C.av_int_list_length_for_size(C.uint(elsize), list, C.uint64_t(term))
	return uint(ret)
}

// --- Function av_fopen_utf8 ---

// av_fopen_utf8 skipped due to return.

// --- Function av_get_time_base_q ---

// AVGetTimeBaseQ wraps av_get_time_base_q.
func AVGetTimeBaseQ() *AVRational {
	ret := C.av_get_time_base_q()
	return &AVRational{value: ret}
}

// --- Function av_fourcc_make_string ---

// AVFourccMakeString wraps av_fourcc_make_string.
func AVFourccMakeString(buf *CStr, fourcc uint32) *CStr {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_fourcc_make_string(tmpbuf, C.uint32_t(fourcc))
	return wrapCStr(ret)
}

// --- Function av_buffer_alloc ---

// AVBufferAlloc wraps av_buffer_alloc.
func AVBufferAlloc(size uint64) *AVBufferRef {
	ret := C.av_buffer_alloc(C.size_t(size))
	var retMapped *AVBufferRef
	if ret != nil {
		retMapped = &AVBufferRef{ptr: ret}
	}
	return retMapped
}

// --- Function av_buffer_allocz ---

// AVBufferAllocz wraps av_buffer_allocz.
func AVBufferAllocz(size uint64) *AVBufferRef {
	ret := C.av_buffer_allocz(C.size_t(size))
	var retMapped *AVBufferRef
	if ret != nil {
		retMapped = &AVBufferRef{ptr: ret}
	}
	return retMapped
}

// --- Function av_buffer_create ---

// av_buffer_create skipped due to free.

// --- Function av_buffer_default_free ---

// AVBufferDefaultFree wraps av_buffer_default_free.
func AVBufferDefaultFree(opaque unsafe.Pointer, data unsafe.Pointer) {
	C.av_buffer_default_free(opaque, (*C.uint8_t)(data))
}

// --- Function av_buffer_ref ---

// AVBufferRef_ wraps av_buffer_ref.
func AVBufferRef_(buf *AVBufferRef) *AVBufferRef {
	var tmpbuf *C.AVBufferRef
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_buffer_ref(tmpbuf)
	var retMapped *AVBufferRef
	if ret != nil {
		retMapped = &AVBufferRef{ptr: ret}
	}
	return retMapped
}

// --- Function av_buffer_unref ---

// AVBufferUnref wraps av_buffer_unref.
func AVBufferUnref(buf **AVBufferRef) {
	var ptrbuf **C.AVBufferRef
	var tmpbuf *C.AVBufferRef
	var oldTmpbuf *C.AVBufferRef
	if buf != nil {
		innerbuf := *buf
		if innerbuf != nil {
			tmpbuf = innerbuf.ptr
			oldTmpbuf = tmpbuf
		}
		ptrbuf = &tmpbuf
	}
	C.av_buffer_unref(ptrbuf)
	if tmpbuf != oldTmpbuf && buf != nil {
		if tmpbuf != nil {
			*buf = &AVBufferRef{ptr: tmpbuf}
		} else {
			*buf = nil
		}
	}
}

// --- Function av_buffer_is_writable ---

// AVBufferIsWritable wraps av_buffer_is_writable.
func AVBufferIsWritable(buf *AVBufferRef) int {
	var tmpbuf *C.AVBufferRef
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_buffer_is_writable(tmpbuf)
	return int(ret)
}

// --- Function av_buffer_get_opaque ---

// AVBufferGetOpaque wraps av_buffer_get_opaque.
func AVBufferGetOpaque(buf *AVBufferRef) unsafe.Pointer {
	var tmpbuf *C.AVBufferRef
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_buffer_get_opaque(tmpbuf)
	return ret
}

// --- Function av_buffer_get_ref_count ---

// AVBufferGetRefCount wraps av_buffer_get_ref_count.
func AVBufferGetRefCount(buf *AVBufferRef) int {
	var tmpbuf *C.AVBufferRef
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_buffer_get_ref_count(tmpbuf)
	return int(ret)
}

// --- Function av_buffer_make_writable ---

// AVBufferMakeWritable wraps av_buffer_make_writable.
func AVBufferMakeWritable(buf **AVBufferRef) int {
	var ptrbuf **C.AVBufferRef
	var tmpbuf *C.AVBufferRef
	var oldTmpbuf *C.AVBufferRef
	if buf != nil {
		innerbuf := *buf
		if innerbuf != nil {
			tmpbuf = innerbuf.ptr
			oldTmpbuf = tmpbuf
		}
		ptrbuf = &tmpbuf
	}
	ret := C.av_buffer_make_writable(ptrbuf)
	if tmpbuf != oldTmpbuf && buf != nil {
		if tmpbuf != nil {
			*buf = &AVBufferRef{ptr: tmpbuf}
		} else {
			*buf = nil
		}
	}
	return int(ret)
}

// --- Function av_buffer_realloc ---

// AVBufferRealloc wraps av_buffer_realloc.
func AVBufferRealloc(buf **AVBufferRef, size uint64) int {
	var ptrbuf **C.AVBufferRef
	var tmpbuf *C.AVBufferRef
	var oldTmpbuf *C.AVBufferRef
	if buf != nil {
		innerbuf := *buf
		if innerbuf != nil {
			tmpbuf = innerbuf.ptr
			oldTmpbuf = tmpbuf
		}
		ptrbuf = &tmpbuf
	}
	ret := C.av_buffer_realloc(ptrbuf, C.size_t(size))
	if tmpbuf != oldTmpbuf && buf != nil {
		if tmpbuf != nil {
			*buf = &AVBufferRef{ptr: tmpbuf}
		} else {
			*buf = nil
		}
	}
	return int(ret)
}

// --- Function av_buffer_replace ---

// AVBufferReplace wraps av_buffer_replace.
func AVBufferReplace(dst **AVBufferRef, src *AVBufferRef) int {
	var ptrdst **C.AVBufferRef
	var tmpdst *C.AVBufferRef
	var oldTmpdst *C.AVBufferRef
	if dst != nil {
		innerdst := *dst
		if innerdst != nil {
			tmpdst = innerdst.ptr
			oldTmpdst = tmpdst
		}
		ptrdst = &tmpdst
	}
	var tmpsrc *C.AVBufferRef
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_buffer_replace(ptrdst, tmpsrc)
	if tmpdst != oldTmpdst && dst != nil {
		if tmpdst != nil {
			*dst = &AVBufferRef{ptr: tmpdst}
		} else {
			*dst = nil
		}
	}
	return int(ret)
}

// --- Function av_buffer_pool_init ---

// av_buffer_pool_init skipped due to alloc.

// --- Function av_buffer_pool_init2 ---

// av_buffer_pool_init2 skipped due to alloc.

// --- Function av_buffer_pool_uninit ---

// AVBufferPoolUninit wraps av_buffer_pool_uninit.
func AVBufferPoolUninit(pool **AVBufferPool) {
	var ptrpool **C.AVBufferPool
	var tmppool *C.AVBufferPool
	var oldTmppool *C.AVBufferPool
	if pool != nil {
		innerpool := *pool
		if innerpool != nil {
			tmppool = innerpool.ptr
			oldTmppool = tmppool
		}
		ptrpool = &tmppool
	}
	C.av_buffer_pool_uninit(ptrpool)
	if tmppool != oldTmppool && pool != nil {
		if tmppool != nil {
			*pool = &AVBufferPool{ptr: tmppool}
		} else {
			*pool = nil
		}
	}
}

// --- Function av_buffer_pool_get ---

// AVBufferPoolGet wraps av_buffer_pool_get.
func AVBufferPoolGet(pool *AVBufferPool) *AVBufferRef {
	var tmppool *C.AVBufferPool
	if pool != nil {
		tmppool = pool.ptr
	}
	ret := C.av_buffer_pool_get(tmppool)
	var retMapped *AVBufferRef
	if ret != nil {
		retMapped = &AVBufferRef{ptr: ret}
	}
	return retMapped
}

// --- Function av_buffer_pool_buffer_get_opaque ---

// AVBufferPoolBufferGetOpaque wraps av_buffer_pool_buffer_get_opaque.
func AVBufferPoolBufferGetOpaque(ref *AVBufferRef) unsafe.Pointer {
	var tmpref *C.AVBufferRef
	if ref != nil {
		tmpref = ref.ptr
	}
	ret := C.av_buffer_pool_buffer_get_opaque(tmpref)
	return ret
}

// --- Function av_get_channel_layout ---

// AVGetChannelLayout wraps av_get_channel_layout.
func AVGetChannelLayout(name *CStr) uint64 {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_get_channel_layout(tmpname)
	return uint64(ret)
}

// --- Function av_get_extended_channel_layout ---

// AVGetExtendedChannelLayout wraps av_get_extended_channel_layout.
// av_get_extended_channel_layout skipped due to channelLayout

// --- Function av_get_channel_layout_string ---

// AVGetChannelLayoutString wraps av_get_channel_layout_string.
func AVGetChannelLayoutString(buf *CStr, bufSize int, nbChannels int, channelLayout uint64) {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	C.av_get_channel_layout_string(tmpbuf, C.int(bufSize), C.int(nbChannels), C.uint64_t(channelLayout))
}

// --- Function av_bprint_channel_layout ---

// AVBprintChannelLayout wraps av_bprint_channel_layout.
func AVBprintChannelLayout(bp *AVBPrint, nbChannels int, channelLayout uint64) {
	var tmpbp *C.struct_AVBPrint
	if bp != nil {
		tmpbp = bp.ptr
	}
	C.av_bprint_channel_layout(tmpbp, C.int(nbChannels), C.uint64_t(channelLayout))
}

// --- Function av_get_channel_layout_nb_channels ---

// AVGetChannelLayoutNbChannels wraps av_get_channel_layout_nb_channels.
func AVGetChannelLayoutNbChannels(channelLayout uint64) int {
	ret := C.av_get_channel_layout_nb_channels(C.uint64_t(channelLayout))
	return int(ret)
}

// --- Function av_get_default_channel_layout ---

// AVGetDefaultChannelLayout wraps av_get_default_channel_layout.
func AVGetDefaultChannelLayout(nbChannels int) int64 {
	ret := C.av_get_default_channel_layout(C.int(nbChannels))
	return int64(ret)
}

// --- Function av_get_channel_layout_channel_index ---

// AVGetChannelLayoutChannelIndex wraps av_get_channel_layout_channel_index.
func AVGetChannelLayoutChannelIndex(channelLayout uint64, channel uint64) int {
	ret := C.av_get_channel_layout_channel_index(C.uint64_t(channelLayout), C.uint64_t(channel))
	return int(ret)
}

// --- Function av_channel_layout_extract_channel ---

// AVChannelLayoutExtractChannel wraps av_channel_layout_extract_channel.
func AVChannelLayoutExtractChannel(channelLayout uint64, index int) uint64 {
	ret := C.av_channel_layout_extract_channel(C.uint64_t(channelLayout), C.int(index))
	return uint64(ret)
}

// --- Function av_get_channel_name ---

// AVGetChannelName wraps av_get_channel_name.
func AVGetChannelName(channel uint64) *CStr {
	ret := C.av_get_channel_name(C.uint64_t(channel))
	return wrapCStr(ret)
}

// --- Function av_get_channel_description ---

// AVGetChannelDescription wraps av_get_channel_description.
func AVGetChannelDescription(channel uint64) *CStr {
	ret := C.av_get_channel_description(C.uint64_t(channel))
	return wrapCStr(ret)
}

// --- Function av_get_standard_channel_layout ---

// AVGetStandardChannelLayout wraps av_get_standard_channel_layout.
// av_get_standard_channel_layout skipped due to layout

// --- Function av_channel_name ---

// AVChannelName wraps av_channel_name.
func AVChannelName(buf *CStr, bufSize uint64, channel AVChannel) int {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_channel_name(tmpbuf, C.size_t(bufSize), C.enum_AVChannel(channel))
	return int(ret)
}

// --- Function av_channel_name_bprint ---

// AVChannelNameBprint wraps av_channel_name_bprint.
func AVChannelNameBprint(bp *AVBPrint, channelId AVChannel) {
	var tmpbp *C.struct_AVBPrint
	if bp != nil {
		tmpbp = bp.ptr
	}
	C.av_channel_name_bprint(tmpbp, C.enum_AVChannel(channelId))
}

// --- Function av_channel_description ---

// AVChannelDescription wraps av_channel_description.
func AVChannelDescription(buf *CStr, bufSize uint64, channel AVChannel) int {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_channel_description(tmpbuf, C.size_t(bufSize), C.enum_AVChannel(channel))
	return int(ret)
}

// --- Function av_channel_description_bprint ---

// AVChannelDescriptionBprint wraps av_channel_description_bprint.
func AVChannelDescriptionBprint(bp *AVBPrint, channelId AVChannel) {
	var tmpbp *C.struct_AVBPrint
	if bp != nil {
		tmpbp = bp.ptr
	}
	C.av_channel_description_bprint(tmpbp, C.enum_AVChannel(channelId))
}

// --- Function av_channel_from_string ---

// AVChannelFromString wraps av_channel_from_string.
func AVChannelFromString(name *CStr) AVChannel {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_channel_from_string(tmpname)
	return AVChannel(ret)
}

// --- Function av_channel_layout_from_mask ---

// AVChannelLayoutFromMask wraps av_channel_layout_from_mask.
func AVChannelLayoutFromMask(channelLayout *AVChannelLayout, mask uint64) int {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	ret := C.av_channel_layout_from_mask(tmpchannelLayout, C.uint64_t(mask))
	return int(ret)
}

// --- Function av_channel_layout_from_string ---

// AVChannelLayoutFromString wraps av_channel_layout_from_string.
func AVChannelLayoutFromString(channelLayout *AVChannelLayout, str *CStr) int {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	var tmpstr *C.char
	if str != nil {
		tmpstr = str.ptr
	}
	ret := C.av_channel_layout_from_string(tmpchannelLayout, tmpstr)
	return int(ret)
}

// --- Function av_channel_layout_default ---

// AVChannelLayoutDefault wraps av_channel_layout_default.
func AVChannelLayoutDefault(chLayout *AVChannelLayout, nbChannels int) {
	var tmpchLayout *C.AVChannelLayout
	if chLayout != nil {
		tmpchLayout = chLayout.ptr
	}
	C.av_channel_layout_default(tmpchLayout, C.int(nbChannels))
}

// --- Function av_channel_layout_standard ---

// AVChannelLayoutStandard wraps av_channel_layout_standard.
// av_channel_layout_standard skipped due to opaque

// --- Function av_channel_layout_uninit ---

// AVChannelLayoutUninit wraps av_channel_layout_uninit.
func AVChannelLayoutUninit(channelLayout *AVChannelLayout) {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	C.av_channel_layout_uninit(tmpchannelLayout)
}

// --- Function av_channel_layout_copy ---

// AVChannelLayoutCopy wraps av_channel_layout_copy.
func AVChannelLayoutCopy(dst *AVChannelLayout, src *AVChannelLayout) int {
	var tmpdst *C.AVChannelLayout
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVChannelLayout
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_channel_layout_copy(tmpdst, tmpsrc)
	return int(ret)
}

// --- Function av_channel_layout_describe ---

// AVChannelLayoutDescribe wraps av_channel_layout_describe.
func AVChannelLayoutDescribe(channelLayout *AVChannelLayout, buf *CStr, bufSize uint64) int {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_channel_layout_describe(tmpchannelLayout, tmpbuf, C.size_t(bufSize))
	return int(ret)
}

// --- Function av_channel_layout_describe_bprint ---

// AVChannelLayoutDescribeBprint wraps av_channel_layout_describe_bprint.
func AVChannelLayoutDescribeBprint(channelLayout *AVChannelLayout, bp *AVBPrint) int {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	var tmpbp *C.struct_AVBPrint
	if bp != nil {
		tmpbp = bp.ptr
	}
	ret := C.av_channel_layout_describe_bprint(tmpchannelLayout, tmpbp)
	return int(ret)
}

// --- Function av_channel_layout_channel_from_index ---

// AVChannelLayoutChannelFromIndex wraps av_channel_layout_channel_from_index.
func AVChannelLayoutChannelFromIndex(channelLayout *AVChannelLayout, idx uint) AVChannel {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	ret := C.av_channel_layout_channel_from_index(tmpchannelLayout, C.uint(idx))
	return AVChannel(ret)
}

// --- Function av_channel_layout_index_from_channel ---

// AVChannelLayoutIndexFromChannel wraps av_channel_layout_index_from_channel.
func AVChannelLayoutIndexFromChannel(channelLayout *AVChannelLayout, channel AVChannel) int {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	ret := C.av_channel_layout_index_from_channel(tmpchannelLayout, C.enum_AVChannel(channel))
	return int(ret)
}

// --- Function av_channel_layout_index_from_string ---

// AVChannelLayoutIndexFromString wraps av_channel_layout_index_from_string.
func AVChannelLayoutIndexFromString(channelLayout *AVChannelLayout, name *CStr) int {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_channel_layout_index_from_string(tmpchannelLayout, tmpname)
	return int(ret)
}

// --- Function av_channel_layout_channel_from_string ---

// AVChannelLayoutChannelFromString wraps av_channel_layout_channel_from_string.
func AVChannelLayoutChannelFromString(channelLayout *AVChannelLayout, name *CStr) AVChannel {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_channel_layout_channel_from_string(tmpchannelLayout, tmpname)
	return AVChannel(ret)
}

// --- Function av_channel_layout_subset ---

// AVChannelLayoutSubset wraps av_channel_layout_subset.
func AVChannelLayoutSubset(channelLayout *AVChannelLayout, mask uint64) uint64 {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	ret := C.av_channel_layout_subset(tmpchannelLayout, C.uint64_t(mask))
	return uint64(ret)
}

// --- Function av_channel_layout_check ---

// AVChannelLayoutCheck wraps av_channel_layout_check.
func AVChannelLayoutCheck(channelLayout *AVChannelLayout) int {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	ret := C.av_channel_layout_check(tmpchannelLayout)
	return int(ret)
}

// --- Function av_channel_layout_compare ---

// AVChannelLayoutCompare wraps av_channel_layout_compare.
func AVChannelLayoutCompare(chl *AVChannelLayout, chl1 *AVChannelLayout) int {
	var tmpchl *C.AVChannelLayout
	if chl != nil {
		tmpchl = chl.ptr
	}
	var tmpchl1 *C.AVChannelLayout
	if chl1 != nil {
		tmpchl1 = chl1.ptr
	}
	ret := C.av_channel_layout_compare(tmpchl, tmpchl1)
	return int(ret)
}

// --- Function av_dict_get ---

// AVDictGet wraps av_dict_get.
func AVDictGet(m *AVDictionary, key *CStr, prev *AVDictionaryEntry, flags int) *AVDictionaryEntry {
	var tmpm *C.AVDictionary
	if m != nil {
		tmpm = m.ptr
	}
	var tmpkey *C.char
	if key != nil {
		tmpkey = key.ptr
	}
	var tmpprev *C.AVDictionaryEntry
	if prev != nil {
		tmpprev = prev.ptr
	}
	ret := C.av_dict_get(tmpm, tmpkey, tmpprev, C.int(flags))
	var retMapped *AVDictionaryEntry
	if ret != nil {
		retMapped = &AVDictionaryEntry{ptr: ret}
	}
	return retMapped
}

// --- Function av_dict_iterate ---

// AVDictIterate wraps av_dict_iterate.
func AVDictIterate(m *AVDictionary, prev *AVDictionaryEntry) *AVDictionaryEntry {
	var tmpm *C.AVDictionary
	if m != nil {
		tmpm = m.ptr
	}
	var tmpprev *C.AVDictionaryEntry
	if prev != nil {
		tmpprev = prev.ptr
	}
	ret := C.av_dict_iterate(tmpm, tmpprev)
	var retMapped *AVDictionaryEntry
	if ret != nil {
		retMapped = &AVDictionaryEntry{ptr: ret}
	}
	return retMapped
}

// --- Function av_dict_count ---

// AVDictCount wraps av_dict_count.
func AVDictCount(m *AVDictionary) int {
	var tmpm *C.AVDictionary
	if m != nil {
		tmpm = m.ptr
	}
	ret := C.av_dict_count(tmpm)
	return int(ret)
}

// --- Function av_dict_set ---

// AVDictSet wraps av_dict_set.
func AVDictSet(pm **AVDictionary, key *CStr, value *CStr, flags int) int {
	var ptrpm **C.AVDictionary
	var tmppm *C.AVDictionary
	var oldTmppm *C.AVDictionary
	if pm != nil {
		innerpm := *pm
		if innerpm != nil {
			tmppm = innerpm.ptr
			oldTmppm = tmppm
		}
		ptrpm = &tmppm
	}
	var tmpkey *C.char
	if key != nil {
		tmpkey = key.ptr
	}
	var tmpvalue *C.char
	if value != nil {
		tmpvalue = value.ptr
	}
	ret := C.av_dict_set(ptrpm, tmpkey, tmpvalue, C.int(flags))
	if tmppm != oldTmppm && pm != nil {
		if tmppm != nil {
			*pm = &AVDictionary{ptr: tmppm}
		} else {
			*pm = nil
		}
	}
	return int(ret)
}

// --- Function av_dict_set_int ---

// AVDictSetInt wraps av_dict_set_int.
func AVDictSetInt(pm **AVDictionary, key *CStr, value int64, flags int) int {
	var ptrpm **C.AVDictionary
	var tmppm *C.AVDictionary
	var oldTmppm *C.AVDictionary
	if pm != nil {
		innerpm := *pm
		if innerpm != nil {
			tmppm = innerpm.ptr
			oldTmppm = tmppm
		}
		ptrpm = &tmppm
	}
	var tmpkey *C.char
	if key != nil {
		tmpkey = key.ptr
	}
	ret := C.av_dict_set_int(ptrpm, tmpkey, C.int64_t(value), C.int(flags))
	if tmppm != oldTmppm && pm != nil {
		if tmppm != nil {
			*pm = &AVDictionary{ptr: tmppm}
		} else {
			*pm = nil
		}
	}
	return int(ret)
}

// --- Function av_dict_parse_string ---

// AVDictParseString wraps av_dict_parse_string.
func AVDictParseString(pm **AVDictionary, str *CStr, keyValSep *CStr, pairsSep *CStr, flags int) int {
	var ptrpm **C.AVDictionary
	var tmppm *C.AVDictionary
	var oldTmppm *C.AVDictionary
	if pm != nil {
		innerpm := *pm
		if innerpm != nil {
			tmppm = innerpm.ptr
			oldTmppm = tmppm
		}
		ptrpm = &tmppm
	}
	var tmpstr *C.char
	if str != nil {
		tmpstr = str.ptr
	}
	var tmpkeyValSep *C.char
	if keyValSep != nil {
		tmpkeyValSep = keyValSep.ptr
	}
	var tmppairsSep *C.char
	if pairsSep != nil {
		tmppairsSep = pairsSep.ptr
	}
	ret := C.av_dict_parse_string(ptrpm, tmpstr, tmpkeyValSep, tmppairsSep, C.int(flags))
	if tmppm != oldTmppm && pm != nil {
		if tmppm != nil {
			*pm = &AVDictionary{ptr: tmppm}
		} else {
			*pm = nil
		}
	}
	return int(ret)
}

// --- Function av_dict_copy ---

// AVDictCopy wraps av_dict_copy.
func AVDictCopy(dst **AVDictionary, src *AVDictionary, flags int) int {
	var ptrdst **C.AVDictionary
	var tmpdst *C.AVDictionary
	var oldTmpdst *C.AVDictionary
	if dst != nil {
		innerdst := *dst
		if innerdst != nil {
			tmpdst = innerdst.ptr
			oldTmpdst = tmpdst
		}
		ptrdst = &tmpdst
	}
	var tmpsrc *C.AVDictionary
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_dict_copy(ptrdst, tmpsrc, C.int(flags))
	if tmpdst != oldTmpdst && dst != nil {
		if tmpdst != nil {
			*dst = &AVDictionary{ptr: tmpdst}
		} else {
			*dst = nil
		}
	}
	return int(ret)
}

// --- Function av_dict_free ---

// AVDictFree wraps av_dict_free.
func AVDictFree(m **AVDictionary) {
	var ptrm **C.AVDictionary
	var tmpm *C.AVDictionary
	var oldTmpm *C.AVDictionary
	if m != nil {
		innerm := *m
		if innerm != nil {
			tmpm = innerm.ptr
			oldTmpm = tmpm
		}
		ptrm = &tmpm
	}
	C.av_dict_free(ptrm)
	if tmpm != oldTmpm && m != nil {
		if tmpm != nil {
			*m = &AVDictionary{ptr: tmpm}
		} else {
			*m = nil
		}
	}
}

// --- Function av_dict_get_string ---

// AVDictGetString wraps av_dict_get_string.
// av_dict_get_string skipped due to buffer

// --- Function av_strerror ---

// AVStrerror wraps av_strerror.
func AVStrerror(errnum int, errbuf *CStr, errbufSize uint64) int {
	var tmperrbuf *C.char
	if errbuf != nil {
		tmperrbuf = errbuf.ptr
	}
	ret := C.av_strerror(C.int(errnum), tmperrbuf, C.size_t(errbufSize))
	return int(ret)
}

// --- Function av_make_error_string ---

// AVMakeErrorString wraps av_make_error_string.
func AVMakeErrorString(errbuf *CStr, errbufSize uint64, errnum int) *CStr {
	var tmperrbuf *C.char
	if errbuf != nil {
		tmperrbuf = errbuf.ptr
	}
	ret := C.av_make_error_string(tmperrbuf, C.size_t(errbufSize), C.int(errnum))
	return wrapCStr(ret)
}

// --- Function av_frame_alloc ---

// AVFrameAlloc wraps av_frame_alloc.
func AVFrameAlloc() *AVFrame {
	ret := C.av_frame_alloc()
	var retMapped *AVFrame
	if ret != nil {
		retMapped = &AVFrame{ptr: ret}
	}
	return retMapped
}

// --- Function av_frame_free ---

// AVFrameFree wraps av_frame_free.
func AVFrameFree(frame **AVFrame) {
	var ptrframe **C.AVFrame
	var tmpframe *C.AVFrame
	var oldTmpframe *C.AVFrame
	if frame != nil {
		innerframe := *frame
		if innerframe != nil {
			tmpframe = innerframe.ptr
			oldTmpframe = tmpframe
		}
		ptrframe = &tmpframe
	}
	C.av_frame_free(ptrframe)
	if tmpframe != oldTmpframe && frame != nil {
		if tmpframe != nil {
			*frame = &AVFrame{ptr: tmpframe}
		} else {
			*frame = nil
		}
	}
}

// --- Function av_frame_ref ---

// AVFrameRef wraps av_frame_ref.
func AVFrameRef(dst *AVFrame, src *AVFrame) int {
	var tmpdst *C.AVFrame
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_frame_ref(tmpdst, tmpsrc)
	return int(ret)
}

// --- Function av_frame_clone ---

// AVFrameClone wraps av_frame_clone.
func AVFrameClone(src *AVFrame) *AVFrame {
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_frame_clone(tmpsrc)
	var retMapped *AVFrame
	if ret != nil {
		retMapped = &AVFrame{ptr: ret}
	}
	return retMapped
}

// --- Function av_frame_unref ---

// AVFrameUnref wraps av_frame_unref.
func AVFrameUnref(frame *AVFrame) {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	C.av_frame_unref(tmpframe)
}

// --- Function av_frame_move_ref ---

// AVFrameMoveRef wraps av_frame_move_ref.
func AVFrameMoveRef(dst *AVFrame, src *AVFrame) {
	var tmpdst *C.AVFrame
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	C.av_frame_move_ref(tmpdst, tmpsrc)
}

// --- Function av_frame_get_buffer ---

// AVFrameGetBuffer wraps av_frame_get_buffer.
func AVFrameGetBuffer(frame *AVFrame, align int) int {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_frame_get_buffer(tmpframe, C.int(align))
	return int(ret)
}

// --- Function av_frame_is_writable ---

// AVFrameIsWritable wraps av_frame_is_writable.
func AVFrameIsWritable(frame *AVFrame) int {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_frame_is_writable(tmpframe)
	return int(ret)
}

// --- Function av_frame_make_writable ---

// AVFrameMakeWritable wraps av_frame_make_writable.
func AVFrameMakeWritable(frame *AVFrame) int {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_frame_make_writable(tmpframe)
	return int(ret)
}

// --- Function av_frame_copy ---

// AVFrameCopy wraps av_frame_copy.
func AVFrameCopy(dst *AVFrame, src *AVFrame) int {
	var tmpdst *C.AVFrame
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_frame_copy(tmpdst, tmpsrc)
	return int(ret)
}

// --- Function av_frame_copy_props ---

// AVFrameCopyProps wraps av_frame_copy_props.
func AVFrameCopyProps(dst *AVFrame, src *AVFrame) int {
	var tmpdst *C.AVFrame
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_frame_copy_props(tmpdst, tmpsrc)
	return int(ret)
}

// --- Function av_frame_get_plane_buffer ---

// AVFrameGetPlaneBuffer wraps av_frame_get_plane_buffer.
func AVFrameGetPlaneBuffer(frame *AVFrame, plane int) *AVBufferRef {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_frame_get_plane_buffer(tmpframe, C.int(plane))
	var retMapped *AVBufferRef
	if ret != nil {
		retMapped = &AVBufferRef{ptr: ret}
	}
	return retMapped
}

// --- Function av_frame_new_side_data ---

// AVFrameNewSideData wraps av_frame_new_side_data.
func AVFrameNewSideData(frame *AVFrame, _type AVFrameSideDataType, size uint64) *AVFrameSideData {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_frame_new_side_data(tmpframe, C.enum_AVFrameSideDataType(_type), C.size_t(size))
	var retMapped *AVFrameSideData
	if ret != nil {
		retMapped = &AVFrameSideData{ptr: ret}
	}
	return retMapped
}

// --- Function av_frame_new_side_data_from_buf ---

// AVFrameNewSideDataFromBuf wraps av_frame_new_side_data_from_buf.
func AVFrameNewSideDataFromBuf(frame *AVFrame, _type AVFrameSideDataType, buf *AVBufferRef) *AVFrameSideData {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	var tmpbuf *C.AVBufferRef
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_frame_new_side_data_from_buf(tmpframe, C.enum_AVFrameSideDataType(_type), tmpbuf)
	var retMapped *AVFrameSideData
	if ret != nil {
		retMapped = &AVFrameSideData{ptr: ret}
	}
	return retMapped
}

// --- Function av_frame_get_side_data ---

// AVFrameGetSideData wraps av_frame_get_side_data.
func AVFrameGetSideData(frame *AVFrame, _type AVFrameSideDataType) *AVFrameSideData {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_frame_get_side_data(tmpframe, C.enum_AVFrameSideDataType(_type))
	var retMapped *AVFrameSideData
	if ret != nil {
		retMapped = &AVFrameSideData{ptr: ret}
	}
	return retMapped
}

// --- Function av_frame_remove_side_data ---

// AVFrameRemoveSideData wraps av_frame_remove_side_data.
func AVFrameRemoveSideData(frame *AVFrame, _type AVFrameSideDataType) {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	C.av_frame_remove_side_data(tmpframe, C.enum_AVFrameSideDataType(_type))
}

// --- Function av_frame_apply_cropping ---

// AVFrameApplyCropping wraps av_frame_apply_cropping.
func AVFrameApplyCropping(frame *AVFrame, flags int) int {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_frame_apply_cropping(tmpframe, C.int(flags))
	return int(ret)
}

// --- Function av_frame_side_data_name ---

// AVFrameSideDataName wraps av_frame_side_data_name.
func AVFrameSideDataName(_type AVFrameSideDataType) *CStr {
	ret := C.av_frame_side_data_name(C.enum_AVFrameSideDataType(_type))
	return wrapCStr(ret)
}

// --- Function av_hwdevice_find_type_by_name ---

// AVHwdeviceFindTypeByName wraps av_hwdevice_find_type_by_name.
func AVHwdeviceFindTypeByName(name *CStr) AVHWDeviceType {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_hwdevice_find_type_by_name(tmpname)
	return AVHWDeviceType(ret)
}

// --- Function av_hwdevice_get_type_name ---

// AVHwdeviceGetTypeName wraps av_hwdevice_get_type_name.
func AVHwdeviceGetTypeName(_type AVHWDeviceType) *CStr {
	ret := C.av_hwdevice_get_type_name(C.enum_AVHWDeviceType(_type))
	return wrapCStr(ret)
}

// --- Function av_hwdevice_iterate_types ---

// AVHwdeviceIterateTypes wraps av_hwdevice_iterate_types.
func AVHwdeviceIterateTypes(prev AVHWDeviceType) AVHWDeviceType {
	ret := C.av_hwdevice_iterate_types(C.enum_AVHWDeviceType(prev))
	return AVHWDeviceType(ret)
}

// --- Function av_hwdevice_ctx_alloc ---

// AVHwdeviceCtxAlloc wraps av_hwdevice_ctx_alloc.
func AVHwdeviceCtxAlloc(_type AVHWDeviceType) *AVBufferRef {
	ret := C.av_hwdevice_ctx_alloc(C.enum_AVHWDeviceType(_type))
	var retMapped *AVBufferRef
	if ret != nil {
		retMapped = &AVBufferRef{ptr: ret}
	}
	return retMapped
}

// --- Function av_hwdevice_ctx_init ---

// AVHwdeviceCtxInit wraps av_hwdevice_ctx_init.
func AVHwdeviceCtxInit(ref *AVBufferRef) int {
	var tmpref *C.AVBufferRef
	if ref != nil {
		tmpref = ref.ptr
	}
	ret := C.av_hwdevice_ctx_init(tmpref)
	return int(ret)
}

// --- Function av_hwdevice_ctx_create ---

// AVHwdeviceCtxCreate wraps av_hwdevice_ctx_create.
func AVHwdeviceCtxCreate(deviceCtx **AVBufferRef, _type AVHWDeviceType, device *CStr, opts *AVDictionary, flags int) int {
	var ptrdeviceCtx **C.AVBufferRef
	var tmpdeviceCtx *C.AVBufferRef
	var oldTmpdeviceCtx *C.AVBufferRef
	if deviceCtx != nil {
		innerdeviceCtx := *deviceCtx
		if innerdeviceCtx != nil {
			tmpdeviceCtx = innerdeviceCtx.ptr
			oldTmpdeviceCtx = tmpdeviceCtx
		}
		ptrdeviceCtx = &tmpdeviceCtx
	}
	var tmpdevice *C.char
	if device != nil {
		tmpdevice = device.ptr
	}
	var tmpopts *C.AVDictionary
	if opts != nil {
		tmpopts = opts.ptr
	}
	ret := C.av_hwdevice_ctx_create(ptrdeviceCtx, C.enum_AVHWDeviceType(_type), tmpdevice, tmpopts, C.int(flags))
	if tmpdeviceCtx != oldTmpdeviceCtx && deviceCtx != nil {
		if tmpdeviceCtx != nil {
			*deviceCtx = &AVBufferRef{ptr: tmpdeviceCtx}
		} else {
			*deviceCtx = nil
		}
	}
	return int(ret)
}

// --- Function av_hwdevice_ctx_create_derived ---

// AVHwdeviceCtxCreateDerived wraps av_hwdevice_ctx_create_derived.
func AVHwdeviceCtxCreateDerived(dstCtx **AVBufferRef, _type AVHWDeviceType, srcCtx *AVBufferRef, flags int) int {
	var ptrdstCtx **C.AVBufferRef
	var tmpdstCtx *C.AVBufferRef
	var oldTmpdstCtx *C.AVBufferRef
	if dstCtx != nil {
		innerdstCtx := *dstCtx
		if innerdstCtx != nil {
			tmpdstCtx = innerdstCtx.ptr
			oldTmpdstCtx = tmpdstCtx
		}
		ptrdstCtx = &tmpdstCtx
	}
	var tmpsrcCtx *C.AVBufferRef
	if srcCtx != nil {
		tmpsrcCtx = srcCtx.ptr
	}
	ret := C.av_hwdevice_ctx_create_derived(ptrdstCtx, C.enum_AVHWDeviceType(_type), tmpsrcCtx, C.int(flags))
	if tmpdstCtx != oldTmpdstCtx && dstCtx != nil {
		if tmpdstCtx != nil {
			*dstCtx = &AVBufferRef{ptr: tmpdstCtx}
		} else {
			*dstCtx = nil
		}
	}
	return int(ret)
}

// --- Function av_hwdevice_ctx_create_derived_opts ---

// AVHwdeviceCtxCreateDerivedOpts wraps av_hwdevice_ctx_create_derived_opts.
func AVHwdeviceCtxCreateDerivedOpts(dstCtx **AVBufferRef, _type AVHWDeviceType, srcCtx *AVBufferRef, options *AVDictionary, flags int) int {
	var ptrdstCtx **C.AVBufferRef
	var tmpdstCtx *C.AVBufferRef
	var oldTmpdstCtx *C.AVBufferRef
	if dstCtx != nil {
		innerdstCtx := *dstCtx
		if innerdstCtx != nil {
			tmpdstCtx = innerdstCtx.ptr
			oldTmpdstCtx = tmpdstCtx
		}
		ptrdstCtx = &tmpdstCtx
	}
	var tmpsrcCtx *C.AVBufferRef
	if srcCtx != nil {
		tmpsrcCtx = srcCtx.ptr
	}
	var tmpoptions *C.AVDictionary
	if options != nil {
		tmpoptions = options.ptr
	}
	ret := C.av_hwdevice_ctx_create_derived_opts(ptrdstCtx, C.enum_AVHWDeviceType(_type), tmpsrcCtx, tmpoptions, C.int(flags))
	if tmpdstCtx != oldTmpdstCtx && dstCtx != nil {
		if tmpdstCtx != nil {
			*dstCtx = &AVBufferRef{ptr: tmpdstCtx}
		} else {
			*dstCtx = nil
		}
	}
	return int(ret)
}

// --- Function av_hwframe_ctx_alloc ---

// AVHwframeCtxAlloc wraps av_hwframe_ctx_alloc.
func AVHwframeCtxAlloc(deviceCtx *AVBufferRef) *AVBufferRef {
	var tmpdeviceCtx *C.AVBufferRef
	if deviceCtx != nil {
		tmpdeviceCtx = deviceCtx.ptr
	}
	ret := C.av_hwframe_ctx_alloc(tmpdeviceCtx)
	var retMapped *AVBufferRef
	if ret != nil {
		retMapped = &AVBufferRef{ptr: ret}
	}
	return retMapped
}

// --- Function av_hwframe_ctx_init ---

// AVHwframeCtxInit wraps av_hwframe_ctx_init.
func AVHwframeCtxInit(ref *AVBufferRef) int {
	var tmpref *C.AVBufferRef
	if ref != nil {
		tmpref = ref.ptr
	}
	ret := C.av_hwframe_ctx_init(tmpref)
	return int(ret)
}

// --- Function av_hwframe_get_buffer ---

// AVHwframeGetBuffer wraps av_hwframe_get_buffer.
func AVHwframeGetBuffer(hwframeCtx *AVBufferRef, frame *AVFrame, flags int) int {
	var tmphwframeCtx *C.AVBufferRef
	if hwframeCtx != nil {
		tmphwframeCtx = hwframeCtx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_hwframe_get_buffer(tmphwframeCtx, tmpframe, C.int(flags))
	return int(ret)
}

// --- Function av_hwframe_transfer_data ---

// AVHwframeTransferData wraps av_hwframe_transfer_data.
func AVHwframeTransferData(dst *AVFrame, src *AVFrame, flags int) int {
	var tmpdst *C.AVFrame
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_hwframe_transfer_data(tmpdst, tmpsrc, C.int(flags))
	return int(ret)
}

// --- Function av_hwframe_transfer_get_formats ---

// AVHwframeTransferGetFormats wraps av_hwframe_transfer_get_formats.
// av_hwframe_transfer_get_formats skipped due to formats

// --- Function av_hwdevice_hwconfig_alloc ---

// AVHwdeviceHwconfigAlloc wraps av_hwdevice_hwconfig_alloc.
func AVHwdeviceHwconfigAlloc(deviceCtx *AVBufferRef) unsafe.Pointer {
	var tmpdeviceCtx *C.AVBufferRef
	if deviceCtx != nil {
		tmpdeviceCtx = deviceCtx.ptr
	}
	ret := C.av_hwdevice_hwconfig_alloc(tmpdeviceCtx)
	return ret
}

// --- Function av_hwdevice_get_hwframe_constraints ---

// AVHwdeviceGetHwframeConstraints wraps av_hwdevice_get_hwframe_constraints.
func AVHwdeviceGetHwframeConstraints(ref *AVBufferRef, hwconfig unsafe.Pointer) *AVHWFramesConstraints {
	var tmpref *C.AVBufferRef
	if ref != nil {
		tmpref = ref.ptr
	}
	ret := C.av_hwdevice_get_hwframe_constraints(tmpref, hwconfig)
	var retMapped *AVHWFramesConstraints
	if ret != nil {
		retMapped = &AVHWFramesConstraints{ptr: ret}
	}
	return retMapped
}

// --- Function av_hwframe_constraints_free ---

// AVHwframeConstraintsFree wraps av_hwframe_constraints_free.
func AVHwframeConstraintsFree(constraints **AVHWFramesConstraints) {
	var ptrconstraints **C.AVHWFramesConstraints
	var tmpconstraints *C.AVHWFramesConstraints
	var oldTmpconstraints *C.AVHWFramesConstraints
	if constraints != nil {
		innerconstraints := *constraints
		if innerconstraints != nil {
			tmpconstraints = innerconstraints.ptr
			oldTmpconstraints = tmpconstraints
		}
		ptrconstraints = &tmpconstraints
	}
	C.av_hwframe_constraints_free(ptrconstraints)
	if tmpconstraints != oldTmpconstraints && constraints != nil {
		if tmpconstraints != nil {
			*constraints = &AVHWFramesConstraints{ptr: tmpconstraints}
		} else {
			*constraints = nil
		}
	}
}

// --- Function av_hwframe_map ---

// AVHwframeMap wraps av_hwframe_map.
func AVHwframeMap(dst *AVFrame, src *AVFrame, flags int) int {
	var tmpdst *C.AVFrame
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_hwframe_map(tmpdst, tmpsrc, C.int(flags))
	return int(ret)
}

// --- Function av_hwframe_ctx_create_derived ---

// AVHwframeCtxCreateDerived wraps av_hwframe_ctx_create_derived.
func AVHwframeCtxCreateDerived(derivedFrameCtx **AVBufferRef, format AVPixelFormat, derivedDeviceCtx *AVBufferRef, sourceFrameCtx *AVBufferRef, flags int) int {
	var ptrderivedFrameCtx **C.AVBufferRef
	var tmpderivedFrameCtx *C.AVBufferRef
	var oldTmpderivedFrameCtx *C.AVBufferRef
	if derivedFrameCtx != nil {
		innerderivedFrameCtx := *derivedFrameCtx
		if innerderivedFrameCtx != nil {
			tmpderivedFrameCtx = innerderivedFrameCtx.ptr
			oldTmpderivedFrameCtx = tmpderivedFrameCtx
		}
		ptrderivedFrameCtx = &tmpderivedFrameCtx
	}
	var tmpderivedDeviceCtx *C.AVBufferRef
	if derivedDeviceCtx != nil {
		tmpderivedDeviceCtx = derivedDeviceCtx.ptr
	}
	var tmpsourceFrameCtx *C.AVBufferRef
	if sourceFrameCtx != nil {
		tmpsourceFrameCtx = sourceFrameCtx.ptr
	}
	ret := C.av_hwframe_ctx_create_derived(ptrderivedFrameCtx, C.enum_AVPixelFormat(format), tmpderivedDeviceCtx, tmpsourceFrameCtx, C.int(flags))
	if tmpderivedFrameCtx != oldTmpderivedFrameCtx && derivedFrameCtx != nil {
		if tmpderivedFrameCtx != nil {
			*derivedFrameCtx = &AVBufferRef{ptr: tmpderivedFrameCtx}
		} else {
			*derivedFrameCtx = nil
		}
	}
	return int(ret)
}

// --- Function av_log ---

// av_log skipped due to variadic arg.

// --- Function av_log_once ---

// av_log_once skipped due to variadic arg.

// --- Function av_vlog ---

// av_vlog skipped due to vl.

// --- Function av_log_get_level ---

// AVLogGetLevel wraps av_log_get_level.
func AVLogGetLevel() int {
	ret := C.av_log_get_level()
	return int(ret)
}

// --- Function av_log_set_level ---

// AVLogSetLevel wraps av_log_set_level.
func AVLogSetLevel(level int) {
	C.av_log_set_level(C.int(level))
}

// --- Function av_log_set_callback ---

// av_log_set_callback skipped due to callback.

// --- Function av_log_default_callback ---

// av_log_default_callback skipped due to vl.

// --- Function av_default_item_name ---

// AVDefaultItemName wraps av_default_item_name.
func AVDefaultItemName(ctx unsafe.Pointer) *CStr {
	ret := C.av_default_item_name(ctx)
	return wrapCStr(ret)
}

// --- Function av_default_get_category ---

// AVDefaultGetCategory wraps av_default_get_category.
func AVDefaultGetCategory(ptr unsafe.Pointer) AVClassCategory {
	ret := C.av_default_get_category(ptr)
	return AVClassCategory(ret)
}

// --- Function av_log_format_line ---

// av_log_format_line skipped due to vl.

// --- Function av_log_format_line2 ---

// av_log_format_line2 skipped due to vl.

// --- Function av_log_set_flags ---

// AVLogSetFlags wraps av_log_set_flags.
func AVLogSetFlags(arg int) {
	C.av_log_set_flags(C.int(arg))
}

// --- Function av_log_get_flags ---

// AVLogGetFlags wraps av_log_get_flags.
func AVLogGetFlags() int {
	ret := C.av_log_get_flags()
	return int(ret)
}

// --- Function av_gcd ---

// AVGcd wraps av_gcd.
func AVGcd(a int64, b int64) int64 {
	ret := C.av_gcd(C.int64_t(a), C.int64_t(b))
	return int64(ret)
}

// --- Function av_rescale ---

// AVRescale wraps av_rescale.
func AVRescale(a int64, b int64, c int64) int64 {
	ret := C.av_rescale(C.int64_t(a), C.int64_t(b), C.int64_t(c))
	return int64(ret)
}

// --- Function av_rescale_rnd ---

// AVRescaleRnd wraps av_rescale_rnd.
func AVRescaleRnd(a int64, b int64, c int64, rnd AVRounding) int64 {
	ret := C.av_rescale_rnd(C.int64_t(a), C.int64_t(b), C.int64_t(c), C.enum_AVRounding(rnd))
	return int64(ret)
}

// --- Function av_rescale_q ---

// AVRescaleQ wraps av_rescale_q.
func AVRescaleQ(a int64, bq *AVRational, cq *AVRational) int64 {
	ret := C.av_rescale_q(C.int64_t(a), bq.value, cq.value)
	return int64(ret)
}

// --- Function av_rescale_q_rnd ---

// AVRescaleQRnd wraps av_rescale_q_rnd.
func AVRescaleQRnd(a int64, bq *AVRational, cq *AVRational, rnd AVRounding) int64 {
	ret := C.av_rescale_q_rnd(C.int64_t(a), bq.value, cq.value, C.enum_AVRounding(rnd))
	return int64(ret)
}

// --- Function av_compare_ts ---

// AVCompareTs wraps av_compare_ts.
func AVCompareTs(tsA int64, tbA *AVRational, tsB int64, tbB *AVRational) int {
	ret := C.av_compare_ts(C.int64_t(tsA), tbA.value, C.int64_t(tsB), tbB.value)
	return int(ret)
}

// --- Function av_compare_mod ---

// AVCompareMod wraps av_compare_mod.
func AVCompareMod(a uint64, b uint64, mod uint64) int64 {
	ret := C.av_compare_mod(C.uint64_t(a), C.uint64_t(b), C.uint64_t(mod))
	return int64(ret)
}

// --- Function av_rescale_delta ---

// AVRescaleDelta wraps av_rescale_delta.
// av_rescale_delta skipped due to last

// --- Function av_add_stable ---

// AVAddStable wraps av_add_stable.
func AVAddStable(tsTb *AVRational, ts int64, incTb *AVRational, inc int64) int64 {
	ret := C.av_add_stable(tsTb.value, C.int64_t(ts), incTb.value, C.int64_t(inc))
	return int64(ret)
}

// --- Function av_malloc ---

// AVMalloc wraps av_malloc.
func AVMalloc(size uint64) unsafe.Pointer {
	ret := C.av_malloc(C.size_t(size))
	return ret
}

// --- Function av_mallocz ---

// AVMallocz wraps av_mallocz.
func AVMallocz(size uint64) unsafe.Pointer {
	ret := C.av_mallocz(C.size_t(size))
	return ret
}

// --- Function av_malloc_array ---

// AVMallocArray wraps av_malloc_array.
func AVMallocArray(nmemb uint64, size uint64) unsafe.Pointer {
	ret := C.av_malloc_array(C.size_t(nmemb), C.size_t(size))
	return ret
}

// --- Function av_calloc ---

// AVCalloc wraps av_calloc.
func AVCalloc(nmemb uint64, size uint64) unsafe.Pointer {
	ret := C.av_calloc(C.size_t(nmemb), C.size_t(size))
	return ret
}

// --- Function av_realloc ---

// AVRealloc wraps av_realloc.
func AVRealloc(ptr unsafe.Pointer, size uint64) unsafe.Pointer {
	ret := C.av_realloc(ptr, C.size_t(size))
	return ret
}

// --- Function av_reallocp ---

// AVReallocp wraps av_reallocp.
func AVReallocp(ptr unsafe.Pointer, size uint64) int {
	ret := C.av_reallocp(ptr, C.size_t(size))
	return int(ret)
}

// --- Function av_realloc_f ---

// AVReallocF wraps av_realloc_f.
func AVReallocF(ptr unsafe.Pointer, nelem uint64, elsize uint64) unsafe.Pointer {
	ret := C.av_realloc_f(ptr, C.size_t(nelem), C.size_t(elsize))
	return ret
}

// --- Function av_realloc_array ---

// AVReallocArray wraps av_realloc_array.
func AVReallocArray(ptr unsafe.Pointer, nmemb uint64, size uint64) unsafe.Pointer {
	ret := C.av_realloc_array(ptr, C.size_t(nmemb), C.size_t(size))
	return ret
}

// --- Function av_reallocp_array ---

// AVReallocpArray wraps av_reallocp_array.
func AVReallocpArray(ptr unsafe.Pointer, nmemb uint64, size uint64) int {
	ret := C.av_reallocp_array(ptr, C.size_t(nmemb), C.size_t(size))
	return int(ret)
}

// --- Function av_fast_realloc ---

// AVFastRealloc wraps av_fast_realloc.
// av_fast_realloc skipped due to size

// --- Function av_fast_malloc ---

// AVFastMalloc wraps av_fast_malloc.
// av_fast_malloc skipped due to size

// --- Function av_fast_mallocz ---

// AVFastMallocz wraps av_fast_mallocz.
// av_fast_mallocz skipped due to size

// --- Function av_free ---

// AVFree wraps av_free.
func AVFree(ptr unsafe.Pointer) {
	C.av_free(ptr)
}

// --- Function av_freep ---

// AVFreep wraps av_freep.
func AVFreep(ptr unsafe.Pointer) {
	C.av_freep(ptr)
}

// --- Function av_strdup ---

// AVStrdup wraps av_strdup.
func AVStrdup(s *CStr) *CStr {
	var tmps *C.char
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_strdup(tmps)
	return wrapCStr(ret)
}

// --- Function av_strndup ---

// AVStrndup wraps av_strndup.
func AVStrndup(s *CStr, len uint64) *CStr {
	var tmps *C.char
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_strndup(tmps, C.size_t(len))
	return wrapCStr(ret)
}

// --- Function av_memdup ---

// AVMemdup wraps av_memdup.
func AVMemdup(p unsafe.Pointer, size uint64) unsafe.Pointer {
	ret := C.av_memdup(p, C.size_t(size))
	return ret
}

// --- Function av_memcpy_backptr ---

// AVMemcpyBackptr wraps av_memcpy_backptr.
func AVMemcpyBackptr(dst unsafe.Pointer, back int, cnt int) {
	C.av_memcpy_backptr((*C.uint8_t)(dst), C.int(back), C.int(cnt))
}

// --- Function av_dynarray_add ---

// AVDynarrayAdd wraps av_dynarray_add.
// av_dynarray_add skipped due to nbPtr

// --- Function av_dynarray_add_nofree ---

// AVDynarrayAddNofree wraps av_dynarray_add_nofree.
// av_dynarray_add_nofree skipped due to nbPtr

// --- Function av_dynarray2_add ---

// AVDynarray2Add wraps av_dynarray2_add.
// av_dynarray2_add skipped due to tabPtr

// --- Function av_size_mult ---

// AVSizeMult wraps av_size_mult.
// av_size_mult skipped due to r

// --- Function av_max_alloc ---

// AVMaxAlloc wraps av_max_alloc.
func AVMaxAlloc(max uint64) {
	C.av_max_alloc(C.size_t(max))
}

// --- Function av_opt_show2 ---

// AVOptShow2 wraps av_opt_show2.
func AVOptShow2(obj unsafe.Pointer, avLogObj unsafe.Pointer, reqFlags int, rejFlags int) int {
	ret := C.av_opt_show2(obj, avLogObj, C.int(reqFlags), C.int(rejFlags))
	return int(ret)
}

// --- Function av_opt_set_defaults ---

// AVOptSetDefaults wraps av_opt_set_defaults.
func AVOptSetDefaults(s unsafe.Pointer) {
	C.av_opt_set_defaults(s)
}

// --- Function av_opt_set_defaults2 ---

// AVOptSetDefaults2 wraps av_opt_set_defaults2.
func AVOptSetDefaults2(s unsafe.Pointer, mask int, flags int) {
	C.av_opt_set_defaults2(s, C.int(mask), C.int(flags))
}

// --- Function av_set_options_string ---

// AVSetOptionsString wraps av_set_options_string.
func AVSetOptionsString(ctx unsafe.Pointer, opts *CStr, keyValSep *CStr, pairsSep *CStr) int {
	var tmpopts *C.char
	if opts != nil {
		tmpopts = opts.ptr
	}
	var tmpkeyValSep *C.char
	if keyValSep != nil {
		tmpkeyValSep = keyValSep.ptr
	}
	var tmppairsSep *C.char
	if pairsSep != nil {
		tmppairsSep = pairsSep.ptr
	}
	ret := C.av_set_options_string(ctx, tmpopts, tmpkeyValSep, tmppairsSep)
	return int(ret)
}

// --- Function av_opt_set_from_string ---

// AVOptSetFromString wraps av_opt_set_from_string.
// av_opt_set_from_string skipped due to shorthand

// --- Function av_opt_free ---

// AVOptFree wraps av_opt_free.
func AVOptFree(obj unsafe.Pointer) {
	C.av_opt_free(obj)
}

// --- Function av_opt_flag_is_set ---

// AVOptFlagIsSet wraps av_opt_flag_is_set.
func AVOptFlagIsSet(obj unsafe.Pointer, fieldName *CStr, flagName *CStr) int {
	var tmpfieldName *C.char
	if fieldName != nil {
		tmpfieldName = fieldName.ptr
	}
	var tmpflagName *C.char
	if flagName != nil {
		tmpflagName = flagName.ptr
	}
	ret := C.av_opt_flag_is_set(obj, tmpfieldName, tmpflagName)
	return int(ret)
}

// --- Function av_opt_set_dict ---

// AVOptSetDict wraps av_opt_set_dict.
func AVOptSetDict(obj unsafe.Pointer, options **AVDictionary) int {
	var ptroptions **C.AVDictionary
	var tmpoptions *C.AVDictionary
	var oldTmpoptions *C.AVDictionary
	if options != nil {
		inneroptions := *options
		if inneroptions != nil {
			tmpoptions = inneroptions.ptr
			oldTmpoptions = tmpoptions
		}
		ptroptions = &tmpoptions
	}
	ret := C.av_opt_set_dict(obj, ptroptions)
	if tmpoptions != oldTmpoptions && options != nil {
		if tmpoptions != nil {
			*options = &AVDictionary{ptr: tmpoptions}
		} else {
			*options = nil
		}
	}
	return int(ret)
}

// --- Function av_opt_set_dict2 ---

// AVOptSetDict2 wraps av_opt_set_dict2.
func AVOptSetDict2(obj unsafe.Pointer, options **AVDictionary, searchFlags int) int {
	var ptroptions **C.AVDictionary
	var tmpoptions *C.AVDictionary
	var oldTmpoptions *C.AVDictionary
	if options != nil {
		inneroptions := *options
		if inneroptions != nil {
			tmpoptions = inneroptions.ptr
			oldTmpoptions = tmpoptions
		}
		ptroptions = &tmpoptions
	}
	ret := C.av_opt_set_dict2(obj, ptroptions, C.int(searchFlags))
	if tmpoptions != oldTmpoptions && options != nil {
		if tmpoptions != nil {
			*options = &AVDictionary{ptr: tmpoptions}
		} else {
			*options = nil
		}
	}
	return int(ret)
}

// --- Function av_opt_get_key_value ---

// AVOptGetKeyValue wraps av_opt_get_key_value.
// av_opt_get_key_value skipped due to ropts

// --- Function av_opt_eval_flags ---

// AVOptEvalFlags wraps av_opt_eval_flags.
// av_opt_eval_flags skipped due to flagsOut

// --- Function av_opt_eval_int ---

// AVOptEvalInt wraps av_opt_eval_int.
// av_opt_eval_int skipped due to intOut

// --- Function av_opt_eval_int64 ---

// AVOptEvalInt64 wraps av_opt_eval_int64.
// av_opt_eval_int64 skipped due to int64Out

// --- Function av_opt_eval_float ---

// AVOptEvalFloat wraps av_opt_eval_float.
// av_opt_eval_float skipped due to floatOut

// --- Function av_opt_eval_double ---

// AVOptEvalDouble wraps av_opt_eval_double.
// av_opt_eval_double skipped due to doubleOut

// --- Function av_opt_eval_q ---

// AVOptEvalQ wraps av_opt_eval_q.
// av_opt_eval_q skipped due to qOut

// --- Function av_opt_find ---

// AVOptFind wraps av_opt_find.
func AVOptFind(obj unsafe.Pointer, name *CStr, unit *CStr, optFlags int, searchFlags int) *AVOption {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	var tmpunit *C.char
	if unit != nil {
		tmpunit = unit.ptr
	}
	ret := C.av_opt_find(obj, tmpname, tmpunit, C.int(optFlags), C.int(searchFlags))
	var retMapped *AVOption
	if ret != nil {
		retMapped = &AVOption{ptr: ret}
	}
	return retMapped
}

// --- Function av_opt_find2 ---

// AVOptFind2 wraps av_opt_find2.
// av_opt_find2 skipped due to targetObj

// --- Function av_opt_next ---

// AVOptNext wraps av_opt_next.
func AVOptNext(obj unsafe.Pointer, prev *AVOption) *AVOption {
	var tmpprev *C.AVOption
	if prev != nil {
		tmpprev = prev.ptr
	}
	ret := C.av_opt_next(obj, tmpprev)
	var retMapped *AVOption
	if ret != nil {
		retMapped = &AVOption{ptr: ret}
	}
	return retMapped
}

// --- Function av_opt_child_next ---

// AVOptChildNext wraps av_opt_child_next.
func AVOptChildNext(obj unsafe.Pointer, prev unsafe.Pointer) unsafe.Pointer {
	ret := C.av_opt_child_next(obj, prev)
	return ret
}

// --- Function av_opt_child_class_iterate ---

// AVOptChildClassIterate wraps av_opt_child_class_iterate.
// av_opt_child_class_iterate skipped due to iter

// --- Function av_opt_set ---

// AVOptSet wraps av_opt_set.
func AVOptSet(obj unsafe.Pointer, name *CStr, val *CStr, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	var tmpval *C.char
	if val != nil {
		tmpval = val.ptr
	}
	ret := C.av_opt_set(obj, tmpname, tmpval, C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_set_int ---

// AVOptSetInt wraps av_opt_set_int.
func AVOptSetInt(obj unsafe.Pointer, name *CStr, val int64, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_int(obj, tmpname, C.int64_t(val), C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_set_double ---

// AVOptSetDouble wraps av_opt_set_double.
func AVOptSetDouble(obj unsafe.Pointer, name *CStr, val float64, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_double(obj, tmpname, C.double(val), C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_set_q ---

// AVOptSetQ wraps av_opt_set_q.
func AVOptSetQ(obj unsafe.Pointer, name *CStr, val *AVRational, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_q(obj, tmpname, val.value, C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_set_bin ---

// AVOptSetBin wraps av_opt_set_bin.
func AVOptSetBin(obj unsafe.Pointer, name *CStr, val unsafe.Pointer, size int, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_bin(obj, tmpname, (*C.uint8_t)(val), C.int(size), C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_set_image_size ---

// AVOptSetImageSize wraps av_opt_set_image_size.
func AVOptSetImageSize(obj unsafe.Pointer, name *CStr, w int, h int, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_image_size(obj, tmpname, C.int(w), C.int(h), C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_set_pixel_fmt ---

// AVOptSetPixelFmt wraps av_opt_set_pixel_fmt.
func AVOptSetPixelFmt(obj unsafe.Pointer, name *CStr, fmt AVPixelFormat, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_pixel_fmt(obj, tmpname, C.enum_AVPixelFormat(fmt), C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_set_sample_fmt ---

// AVOptSetSampleFmt wraps av_opt_set_sample_fmt.
func AVOptSetSampleFmt(obj unsafe.Pointer, name *CStr, fmt AVSampleFormat, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_sample_fmt(obj, tmpname, C.enum_AVSampleFormat(fmt), C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_set_video_rate ---

// AVOptSetVideoRate wraps av_opt_set_video_rate.
func AVOptSetVideoRate(obj unsafe.Pointer, name *CStr, val *AVRational, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_video_rate(obj, tmpname, val.value, C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_set_channel_layout ---

// AVOptSetChannelLayout wraps av_opt_set_channel_layout.
func AVOptSetChannelLayout(obj unsafe.Pointer, name *CStr, chLayout int64, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_channel_layout(obj, tmpname, C.int64_t(chLayout), C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_set_chlayout ---

// AVOptSetChlayout wraps av_opt_set_chlayout.
func AVOptSetChlayout(obj unsafe.Pointer, name *CStr, layout *AVChannelLayout, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	var tmplayout *C.AVChannelLayout
	if layout != nil {
		tmplayout = layout.ptr
	}
	ret := C.av_opt_set_chlayout(obj, tmpname, tmplayout, C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_set_dict_val ---

// AVOptSetDictVal wraps av_opt_set_dict_val.
func AVOptSetDictVal(obj unsafe.Pointer, name *CStr, val *AVDictionary, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	var tmpval *C.AVDictionary
	if val != nil {
		tmpval = val.ptr
	}
	ret := C.av_opt_set_dict_val(obj, tmpname, tmpval, C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_get ---

// AVOptGet wraps av_opt_get.
// av_opt_get skipped due to outVal

// --- Function av_opt_get_int ---

// AVOptGetInt wraps av_opt_get_int.
// av_opt_get_int skipped due to outVal

// --- Function av_opt_get_double ---

// AVOptGetDouble wraps av_opt_get_double.
// av_opt_get_double skipped due to outVal

// --- Function av_opt_get_q ---

// AVOptGetQ wraps av_opt_get_q.
// av_opt_get_q skipped due to outVal

// --- Function av_opt_get_image_size ---

// AVOptGetImageSize wraps av_opt_get_image_size.
// av_opt_get_image_size skipped due to wOut

// --- Function av_opt_get_pixel_fmt ---

// AVOptGetPixelFmt wraps av_opt_get_pixel_fmt.
// av_opt_get_pixel_fmt skipped due to outFmt

// --- Function av_opt_get_sample_fmt ---

// AVOptGetSampleFmt wraps av_opt_get_sample_fmt.
// av_opt_get_sample_fmt skipped due to outFmt

// --- Function av_opt_get_video_rate ---

// AVOptGetVideoRate wraps av_opt_get_video_rate.
// av_opt_get_video_rate skipped due to outVal

// --- Function av_opt_get_channel_layout ---

// AVOptGetChannelLayout wraps av_opt_get_channel_layout.
// av_opt_get_channel_layout skipped due to chLayout

// --- Function av_opt_get_chlayout ---

// AVOptGetChlayout wraps av_opt_get_chlayout.
func AVOptGetChlayout(obj unsafe.Pointer, name *CStr, searchFlags int, layout *AVChannelLayout) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	var tmplayout *C.AVChannelLayout
	if layout != nil {
		tmplayout = layout.ptr
	}
	ret := C.av_opt_get_chlayout(obj, tmpname, C.int(searchFlags), tmplayout)
	return int(ret)
}

// --- Function av_opt_get_dict_val ---

// AVOptGetDictVal wraps av_opt_get_dict_val.
func AVOptGetDictVal(obj unsafe.Pointer, name *CStr, searchFlags int, outVal **AVDictionary) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	var ptroutVal **C.AVDictionary
	var tmpoutVal *C.AVDictionary
	var oldTmpoutVal *C.AVDictionary
	if outVal != nil {
		inneroutVal := *outVal
		if inneroutVal != nil {
			tmpoutVal = inneroutVal.ptr
			oldTmpoutVal = tmpoutVal
		}
		ptroutVal = &tmpoutVal
	}
	ret := C.av_opt_get_dict_val(obj, tmpname, C.int(searchFlags), ptroutVal)
	if tmpoutVal != oldTmpoutVal && outVal != nil {
		if tmpoutVal != nil {
			*outVal = &AVDictionary{ptr: tmpoutVal}
		} else {
			*outVal = nil
		}
	}
	return int(ret)
}

// --- Function av_opt_ptr ---

// AVOptPtr wraps av_opt_ptr.
func AVOptPtr(avclass *AVClass, obj unsafe.Pointer, name *CStr) unsafe.Pointer {
	var tmpavclass *C.AVClass
	if avclass != nil {
		tmpavclass = avclass.ptr
	}
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_ptr(tmpavclass, obj, tmpname)
	return ret
}

// --- Function av_opt_freep_ranges ---

// AVOptFreepRanges wraps av_opt_freep_ranges.
func AVOptFreepRanges(ranges **AVOptionRanges) {
	var ptrranges **C.AVOptionRanges
	var tmpranges *C.AVOptionRanges
	var oldTmpranges *C.AVOptionRanges
	if ranges != nil {
		innerranges := *ranges
		if innerranges != nil {
			tmpranges = innerranges.ptr
			oldTmpranges = tmpranges
		}
		ptrranges = &tmpranges
	}
	C.av_opt_freep_ranges(ptrranges)
	if tmpranges != oldTmpranges && ranges != nil {
		if tmpranges != nil {
			*ranges = &AVOptionRanges{ptr: tmpranges}
		} else {
			*ranges = nil
		}
	}
}

// --- Function av_opt_query_ranges ---

// AVOptQueryRanges wraps av_opt_query_ranges.
func AVOptQueryRanges(param0 **AVOptionRanges, obj unsafe.Pointer, key *CStr, flags int) int {
	var ptrparam0 **C.AVOptionRanges
	var tmpparam0 *C.AVOptionRanges
	var oldTmpparam0 *C.AVOptionRanges
	if param0 != nil {
		innerparam0 := *param0
		if innerparam0 != nil {
			tmpparam0 = innerparam0.ptr
			oldTmpparam0 = tmpparam0
		}
		ptrparam0 = &tmpparam0
	}
	var tmpkey *C.char
	if key != nil {
		tmpkey = key.ptr
	}
	ret := C.av_opt_query_ranges(ptrparam0, obj, tmpkey, C.int(flags))
	if tmpparam0 != oldTmpparam0 && param0 != nil {
		if tmpparam0 != nil {
			*param0 = &AVOptionRanges{ptr: tmpparam0}
		} else {
			*param0 = nil
		}
	}
	return int(ret)
}

// --- Function av_opt_copy ---

// AVOptCopy wraps av_opt_copy.
func AVOptCopy(dest unsafe.Pointer, src unsafe.Pointer) int {
	ret := C.av_opt_copy(dest, src)
	return int(ret)
}

// --- Function av_opt_query_ranges_default ---

// AVOptQueryRangesDefault wraps av_opt_query_ranges_default.
func AVOptQueryRangesDefault(param0 **AVOptionRanges, obj unsafe.Pointer, key *CStr, flags int) int {
	var ptrparam0 **C.AVOptionRanges
	var tmpparam0 *C.AVOptionRanges
	var oldTmpparam0 *C.AVOptionRanges
	if param0 != nil {
		innerparam0 := *param0
		if innerparam0 != nil {
			tmpparam0 = innerparam0.ptr
			oldTmpparam0 = tmpparam0
		}
		ptrparam0 = &tmpparam0
	}
	var tmpkey *C.char
	if key != nil {
		tmpkey = key.ptr
	}
	ret := C.av_opt_query_ranges_default(ptrparam0, obj, tmpkey, C.int(flags))
	if tmpparam0 != oldTmpparam0 && param0 != nil {
		if tmpparam0 != nil {
			*param0 = &AVOptionRanges{ptr: tmpparam0}
		} else {
			*param0 = nil
		}
	}
	return int(ret)
}

// --- Function av_opt_is_set_to_default ---

// AVOptIsSetToDefault wraps av_opt_is_set_to_default.
func AVOptIsSetToDefault(obj unsafe.Pointer, o *AVOption) int {
	var tmpo *C.AVOption
	if o != nil {
		tmpo = o.ptr
	}
	ret := C.av_opt_is_set_to_default(obj, tmpo)
	return int(ret)
}

// --- Function av_opt_is_set_to_default_by_name ---

// AVOptIsSetToDefaultByName wraps av_opt_is_set_to_default_by_name.
func AVOptIsSetToDefaultByName(obj unsafe.Pointer, name *CStr, searchFlags int) int {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_is_set_to_default_by_name(obj, tmpname, C.int(searchFlags))
	return int(ret)
}

// --- Function av_opt_serialize ---

// AVOptSerialize wraps av_opt_serialize.
// av_opt_serialize skipped due to buffer

// --- Function av_make_q ---

// AVMakeQ wraps av_make_q.
func AVMakeQ(num int, den int) *AVRational {
	ret := C.av_make_q(C.int(num), C.int(den))
	return &AVRational{value: ret}
}

// --- Function av_cmp_q ---

// AVCmpQ wraps av_cmp_q.
func AVCmpQ(a *AVRational, b *AVRational) int {
	ret := C.av_cmp_q(a.value, b.value)
	return int(ret)
}

// --- Function av_q2d ---

// AVQ2D wraps av_q2d.
func AVQ2D(a *AVRational) float64 {
	ret := C.av_q2d(a.value)
	return float64(ret)
}

// --- Function av_reduce ---

// AVReduce wraps av_reduce.
// av_reduce skipped due to dstNum

// --- Function av_mul_q ---

// AVMulQ wraps av_mul_q.
func AVMulQ(b *AVRational, c *AVRational) *AVRational {
	ret := C.av_mul_q(b.value, c.value)
	return &AVRational{value: ret}
}

// --- Function av_div_q ---

// AVDivQ wraps av_div_q.
func AVDivQ(b *AVRational, c *AVRational) *AVRational {
	ret := C.av_div_q(b.value, c.value)
	return &AVRational{value: ret}
}

// --- Function av_add_q ---

// AVAddQ wraps av_add_q.
func AVAddQ(b *AVRational, c *AVRational) *AVRational {
	ret := C.av_add_q(b.value, c.value)
	return &AVRational{value: ret}
}

// --- Function av_sub_q ---

// AVSubQ wraps av_sub_q.
func AVSubQ(b *AVRational, c *AVRational) *AVRational {
	ret := C.av_sub_q(b.value, c.value)
	return &AVRational{value: ret}
}

// --- Function av_inv_q ---

// AVInvQ wraps av_inv_q.
func AVInvQ(q *AVRational) *AVRational {
	ret := C.av_inv_q(q.value)
	return &AVRational{value: ret}
}

// --- Function av_d2q ---

// AVD2Q wraps av_d2q.
func AVD2Q(d float64, max int) *AVRational {
	ret := C.av_d2q(C.double(d), C.int(max))
	return &AVRational{value: ret}
}

// --- Function av_nearer_q ---

// AVNearerQ wraps av_nearer_q.
func AVNearerQ(q *AVRational, q1 *AVRational, q2 *AVRational) int {
	ret := C.av_nearer_q(q.value, q1.value, q2.value)
	return int(ret)
}

// --- Function av_find_nearest_q_idx ---

// AVFindNearestQIdx wraps av_find_nearest_q_idx.
// av_find_nearest_q_idx skipped due to qList

// --- Function av_q2intfloat ---

// AVQ2Intfloat wraps av_q2intfloat.
func AVQ2Intfloat(q *AVRational) uint32 {
	ret := C.av_q2intfloat(q.value)
	return uint32(ret)
}

// --- Function av_gcd_q ---

// AVGcdQ wraps av_gcd_q.
func AVGcdQ(a *AVRational, b *AVRational, maxDen int, def *AVRational) *AVRational {
	ret := C.av_gcd_q(a.value, b.value, C.int(maxDen), def.value)
	return &AVRational{value: ret}
}

// --- Function av_get_sample_fmt_name ---

// AVGetSampleFmtName wraps av_get_sample_fmt_name.
func AVGetSampleFmtName(sampleFmt AVSampleFormat) *CStr {
	ret := C.av_get_sample_fmt_name(C.enum_AVSampleFormat(sampleFmt))
	return wrapCStr(ret)
}

// --- Function av_get_sample_fmt ---

// AVGetSampleFmt wraps av_get_sample_fmt.
func AVGetSampleFmt(name *CStr) AVSampleFormat {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_get_sample_fmt(tmpname)
	return AVSampleFormat(ret)
}

// --- Function av_get_alt_sample_fmt ---

// AVGetAltSampleFmt wraps av_get_alt_sample_fmt.
func AVGetAltSampleFmt(sampleFmt AVSampleFormat, planar int) AVSampleFormat {
	ret := C.av_get_alt_sample_fmt(C.enum_AVSampleFormat(sampleFmt), C.int(planar))
	return AVSampleFormat(ret)
}

// --- Function av_get_packed_sample_fmt ---

// AVGetPackedSampleFmt wraps av_get_packed_sample_fmt.
func AVGetPackedSampleFmt(sampleFmt AVSampleFormat) AVSampleFormat {
	ret := C.av_get_packed_sample_fmt(C.enum_AVSampleFormat(sampleFmt))
	return AVSampleFormat(ret)
}

// --- Function av_get_planar_sample_fmt ---

// AVGetPlanarSampleFmt wraps av_get_planar_sample_fmt.
func AVGetPlanarSampleFmt(sampleFmt AVSampleFormat) AVSampleFormat {
	ret := C.av_get_planar_sample_fmt(C.enum_AVSampleFormat(sampleFmt))
	return AVSampleFormat(ret)
}

// --- Function av_get_sample_fmt_string ---

// AVGetSampleFmtString wraps av_get_sample_fmt_string.
func AVGetSampleFmtString(buf *CStr, bufSize int, sampleFmt AVSampleFormat) *CStr {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_get_sample_fmt_string(tmpbuf, C.int(bufSize), C.enum_AVSampleFormat(sampleFmt))
	return wrapCStr(ret)
}

// --- Function av_get_bytes_per_sample ---

// AVGetBytesPerSample wraps av_get_bytes_per_sample.
func AVGetBytesPerSample(sampleFmt AVSampleFormat) int {
	ret := C.av_get_bytes_per_sample(C.enum_AVSampleFormat(sampleFmt))
	return int(ret)
}

// --- Function av_sample_fmt_is_planar ---

// AVSampleFmtIsPlanar wraps av_sample_fmt_is_planar.
func AVSampleFmtIsPlanar(sampleFmt AVSampleFormat) int {
	ret := C.av_sample_fmt_is_planar(C.enum_AVSampleFormat(sampleFmt))
	return int(ret)
}

// --- Function av_samples_get_buffer_size ---

// AVSamplesGetBufferSize wraps av_samples_get_buffer_size.
// av_samples_get_buffer_size skipped due to linesize

// --- Function av_samples_fill_arrays ---

// AVSamplesFillArrays wraps av_samples_fill_arrays.
// av_samples_fill_arrays skipped due to audioData

// --- Function av_samples_alloc ---

// AVSamplesAlloc wraps av_samples_alloc.
// av_samples_alloc skipped due to audioData

// --- Function av_samples_alloc_array_and_samples ---

// AVSamplesAllocArrayAndSamples wraps av_samples_alloc_array_and_samples.
// av_samples_alloc_array_and_samples skipped due to audioData

// --- Function av_samples_copy ---

// AVSamplesCopy wraps av_samples_copy.
// av_samples_copy skipped due to dst

// --- Function av_samples_set_silence ---

// AVSamplesSetSilence wraps av_samples_set_silence.
// av_samples_set_silence skipped due to audioData
