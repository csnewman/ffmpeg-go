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
//
//	Return the LIBAVCODEC_VERSION_INT constant.
func AVCodecVersion() uint {
	ret := C.avcodec_version()
	return uint(ret)
}

// --- Function avcodec_configuration ---

// AVCodecConfiguration wraps avcodec_configuration.
//
//	Return the libavcodec build-time configuration.
func AVCodecConfiguration() *CStr {
	ret := C.avcodec_configuration()
	return wrapCStr(ret)
}

// --- Function avcodec_license ---

// AVCodecLicense wraps avcodec_license.
//
//	Return the libavcodec license.
func AVCodecLicense() *CStr {
	ret := C.avcodec_license()
	return wrapCStr(ret)
}

// --- Function avcodec_alloc_context3 ---

// AVCodecAllocContext3 wraps avcodec_alloc_context3.
/*
  Allocate an AVCodecContext and set its fields to default values. The
  resulting struct should be freed with avcodec_free_context().

  @param codec if non-NULL, allocate private data and initialize defaults
               for the given codec. It is illegal to then call avcodec_open2()
               with a different codec.
               If NULL, then the codec-specific defaults won't be initialized,
               which may result in suboptimal default settings (this is
               important mainly for encoders, e.g. libx264).

  @return An AVCodecContext filled with default values or NULL on failure.
*/
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
/*
  Free the codec context and everything associated with it and write NULL to
  the provided pointer.
*/
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
/*
  Get the AVClass for AVCodecContext. It can be used in combination with
  AV_OPT_SEARCH_FAKE_OBJ for examining options.

  @see av_opt_find().
*/
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
/*
  Get the AVClass for AVSubtitleRect. It can be used in combination with
  AV_OPT_SEARCH_FAKE_OBJ for examining options.

  @see av_opt_find().
*/
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
/*
  Fill the parameters struct based on the values from the supplied codec
  context. Any allocated fields in par are freed and replaced with duplicates
  of the corresponding fields in codec.

  @return >= 0 on success, a negative AVERROR code on failure
*/
func AVCodecParametersFromContext(par *AVCodecParameters, codec *AVCodecContext) (int, error) {
	var tmppar *C.AVCodecParameters
	if par != nil {
		tmppar = par.ptr
	}
	var tmpcodec *C.AVCodecContext
	if codec != nil {
		tmpcodec = codec.ptr
	}
	ret := C.avcodec_parameters_from_context(tmppar, tmpcodec)
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_parameters_to_context ---

// AVCodecParametersToContext wraps avcodec_parameters_to_context.
/*
  Fill the codec context based on the values from the supplied codec
  parameters. Any allocated fields in codec that have a corresponding field in
  par are freed and replaced with duplicates of the corresponding field in par.
  Fields in codec that do not have a counterpart in par are not touched.

  @return >= 0 on success, a negative AVERROR code on failure.
*/
func AVCodecParametersToContext(codec *AVCodecContext, par *AVCodecParameters) (int, error) {
	var tmpcodec *C.AVCodecContext
	if codec != nil {
		tmpcodec = codec.ptr
	}
	var tmppar *C.AVCodecParameters
	if par != nil {
		tmppar = par.ptr
	}
	ret := C.avcodec_parameters_to_context(tmpcodec, tmppar)
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_open2 ---

// AVCodecOpen2 wraps avcodec_open2.
/*
  Initialize the AVCodecContext to use the given AVCodec. Prior to using this
  function the context has to be allocated with avcodec_alloc_context3().

  The functions avcodec_find_decoder_by_name(), avcodec_find_encoder_by_name(),
  avcodec_find_decoder() and avcodec_find_encoder() provide an easy way for
  retrieving a codec.

  @note Always call this function before using decoding routines (such as
  @ref avcodec_receive_frame()).

  @code
  av_dict_set(&opts, "b", "2.5M", 0);
  codec = avcodec_find_decoder(AV_CODEC_ID_H264);
  if (!codec)
      exit(1);

  context = avcodec_alloc_context3(codec);

  if (avcodec_open2(context, codec, opts) < 0)
      exit(1);
  @endcode

  @param avctx The context to initialize.
  @param codec The codec to open this context for. If a non-NULL codec has been
               previously passed to avcodec_alloc_context3() or
               for this context, then this parameter MUST be either NULL or
               equal to the previously passed codec.
  @param options A dictionary filled with AVCodecContext and codec-private options.
                 On return this object will be filled with options that were not found.

  @return zero on success, a negative value on error
  @see avcodec_alloc_context3(), avcodec_find_decoder(), avcodec_find_encoder(),
       av_dict_set(), av_opt_find().
*/
func AVCodecOpen2(avctx *AVCodecContext, codec *AVCodec, options **AVDictionary) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_close ---

// AVCodecClose wraps avcodec_close.
/*
  Close a given AVCodecContext and free all the data associated with it
  (but not the AVCodecContext itself).

  Calling this function on an AVCodecContext that hasn't been opened will free
  the codec-specific data allocated in avcodec_alloc_context3() with a non-NULL
  codec. Subsequent calls will do nothing.

  @note Do not use this function. Use avcodec_free_context() to destroy a
  codec context (either open or closed). Opening and closing a codec context
  multiple times is not supported anymore -- use multiple codec contexts
  instead.
*/
func AVCodecClose(avctx *AVCodecContext) (int, error) {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	ret := C.avcodec_close(tmpavctx)
	return int(ret), WrapErr(int(ret))
}

// --- Function avsubtitle_free ---

// AVSubtitleFree wraps avsubtitle_free.
/*
  Free all allocated data in the given subtitle struct.

  @param sub AVSubtitle to free.
*/
func AVSubtitleFree(sub *AVSubtitle) {
	var tmpsub *C.AVSubtitle
	if sub != nil {
		tmpsub = sub.ptr
	}
	C.avsubtitle_free(tmpsub)
}

// --- Function avcodec_default_get_buffer2 ---

// AVCodecDefaultGetBuffer2 wraps avcodec_default_get_buffer2.
/*
  The default callback for AVCodecContext.get_buffer2(). It is made public so
  it can be called by custom get_buffer2() implementations for decoders without
  AV_CODEC_CAP_DR1 set.
*/
func AVCodecDefaultGetBuffer2(s *AVCodecContext, frame *AVFrame, flags int) (int, error) {
	var tmps *C.AVCodecContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.avcodec_default_get_buffer2(tmps, tmpframe, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_default_get_encode_buffer ---

// AVCodecDefaultGetEncodeBuffer wraps avcodec_default_get_encode_buffer.
/*
  The default callback for AVCodecContext.get_encode_buffer(). It is made public so
  it can be called by custom get_encode_buffer() implementations for encoders without
  AV_CODEC_CAP_DR1 set.
*/
func AVCodecDefaultGetEncodeBuffer(s *AVCodecContext, pkt *AVPacket, flags int) (int, error) {
	var tmps *C.AVCodecContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.avcodec_default_get_encode_buffer(tmps, tmppkt, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_align_dimensions ---

// avcodec_align_dimensions skipped due to width

// --- Function avcodec_align_dimensions2 ---

// avcodec_align_dimensions2 skipped due to width

// --- Function avcodec_enum_to_chroma_pos ---

// avcodec_enum_to_chroma_pos skipped due to xpos

// --- Function avcodec_chroma_pos_to_enum ---

// AVCodecChromaPosToEnum wraps avcodec_chroma_pos_to_enum.
/*
  Converts swscale x/y chroma position to AVChromaLocation.

  The positions represent the chroma (0,0) position in a coordinates system
  with luma (0,0) representing the origin and luma(1,1) representing 256,256

  @param xpos  horizontal chroma sample position
  @param ypos  vertical   chroma sample position
  @deprecated Use av_chroma_location_pos_to_enum() instead.
*/
func AVCodecChromaPosToEnum(xpos int, ypos int) AVChromaLocation {
	ret := C.avcodec_chroma_pos_to_enum(C.int(xpos), C.int(ypos))
	return AVChromaLocation(ret)
}

// --- Function avcodec_decode_subtitle2 ---

// avcodec_decode_subtitle2 skipped due to gotSubPtr

// --- Function avcodec_send_packet ---

// AVCodecSendPacket wraps avcodec_send_packet.
/*
  Supply raw packet data as input to a decoder.

  Internally, this call will copy relevant AVCodecContext fields, which can
  influence decoding per-packet, and apply them when the packet is actually
  decoded. (For example AVCodecContext.skip_frame, which might direct the
  decoder to drop the frame contained by the packet sent with this function.)

  @warning The input buffer, avpkt->data must be AV_INPUT_BUFFER_PADDING_SIZE
           larger than the actual read bytes because some optimized bitstream
           readers read 32 or 64 bits at once and could read over the end.

  @note The AVCodecContext MUST have been opened with @ref avcodec_open2()
        before packets may be fed to the decoder.

  @param avctx codec context
  @param[in] avpkt The input AVPacket. Usually, this will be a single video
                   frame, or several complete audio frames.
                   Ownership of the packet remains with the caller, and the
                   decoder will not write to the packet. The decoder may create
                   a reference to the packet data (or copy it if the packet is
                   not reference-counted).
                   Unlike with older APIs, the packet is always fully consumed,
                   and if it contains multiple frames (e.g. some audio codecs),
                   will require you to call avcodec_receive_frame() multiple
                   times afterwards before you can send a new packet.
                   It can be NULL (or an AVPacket with data set to NULL and
                   size set to 0); in this case, it is considered a flush
                   packet, which signals the end of the stream. Sending the
                   first flush packet will return success. Subsequent ones are
                   unnecessary and will return AVERROR_EOF. If the decoder
                   still has frames buffered, it will return them after sending
                   a flush packet.

  @retval 0                 success
  @retval AVERROR(EAGAIN)   input is not accepted in the current state - user
                            must read output with avcodec_receive_frame() (once
                            all output is read, the packet should be resent,
                            and the call will not fail with EAGAIN).
  @retval AVERROR_EOF       the decoder has been flushed, and no new packets can be
                            sent to it (also returned if more than 1 flush
                            packet is sent)
  @retval AVERROR(EINVAL)   codec not opened, it is an encoder, or requires flush
  @retval AVERROR(ENOMEM)   failed to add packet to internal queue, or similar
  @retval "another negative error code" legitimate decoding errors
*/
func AVCodecSendPacket(avctx *AVCodecContext, avpkt *AVPacket) (int, error) {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpavpkt *C.AVPacket
	if avpkt != nil {
		tmpavpkt = avpkt.ptr
	}
	ret := C.avcodec_send_packet(tmpavctx, tmpavpkt)
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_receive_frame ---

// AVCodecReceiveFrame wraps avcodec_receive_frame.
/*
  Return decoded output data from a decoder or encoder (when the
  AV_CODEC_FLAG_RECON_FRAME flag is used).

  @param avctx codec context
  @param frame This will be set to a reference-counted video or audio
               frame (depending on the decoder type) allocated by the
               codec. Note that the function will always call
               av_frame_unref(frame) before doing anything else.

  @retval 0                success, a frame was returned
  @retval AVERROR(EAGAIN)  output is not available in this state - user must
                           try to send new input
  @retval AVERROR_EOF      the codec has been fully flushed, and there will be
                           no more output frames
  @retval AVERROR(EINVAL)  codec not opened, or it is an encoder without the
                           AV_CODEC_FLAG_RECON_FRAME flag enabled
  @retval AVERROR_INPUT_CHANGED current decoded frame has changed parameters with
                           respect to first decoded frame. Applicable when flag
                           AV_CODEC_FLAG_DROPCHANGED is set.
  @retval "other negative error code" legitimate decoding errors
*/
func AVCodecReceiveFrame(avctx *AVCodecContext, frame *AVFrame) (int, error) {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.avcodec_receive_frame(tmpavctx, tmpframe)
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_send_frame ---

// AVCodecSendFrame wraps avcodec_send_frame.
/*
  Supply a raw video or audio frame to the encoder. Use avcodec_receive_packet()
  to retrieve buffered output packets.

  @param avctx     codec context
  @param[in] frame AVFrame containing the raw audio or video frame to be encoded.
                   Ownership of the frame remains with the caller, and the
                   encoder will not write to the frame. The encoder may create
                   a reference to the frame data (or copy it if the frame is
                   not reference-counted).
                   It can be NULL, in which case it is considered a flush
                   packet.  This signals the end of the stream. If the encoder
                   still has packets buffered, it will return them after this
                   call. Once flushing mode has been entered, additional flush
                   packets are ignored, and sending frames will return
                   AVERROR_EOF.

                   For audio:
                   If AV_CODEC_CAP_VARIABLE_FRAME_SIZE is set, then each frame
                   can have any number of samples.
                   If it is not set, frame->nb_samples must be equal to
                   avctx->frame_size for all frames except the last.
                   The final frame may be smaller than avctx->frame_size.
  @retval 0                 success
  @retval AVERROR(EAGAIN)   input is not accepted in the current state - user must
                            read output with avcodec_receive_packet() (once all
                            output is read, the packet should be resent, and the
                            call will not fail with EAGAIN).
  @retval AVERROR_EOF       the encoder has been flushed, and no new frames can
                            be sent to it
  @retval AVERROR(EINVAL)   codec not opened, it is a decoder, or requires flush
  @retval AVERROR(ENOMEM)   failed to add packet to internal queue, or similar
  @retval "another negative error code" legitimate encoding errors
*/
func AVCodecSendFrame(avctx *AVCodecContext, frame *AVFrame) (int, error) {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.avcodec_send_frame(tmpavctx, tmpframe)
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_receive_packet ---

// AVCodecReceivePacket wraps avcodec_receive_packet.
/*
  Read encoded data from the encoder.

  @param avctx codec context
  @param avpkt This will be set to a reference-counted packet allocated by the
               encoder. Note that the function will always call
               av_packet_unref(avpkt) before doing anything else.
  @retval 0               success
  @retval AVERROR(EAGAIN) output is not available in the current state - user must
                          try to send input
  @retval AVERROR_EOF     the encoder has been fully flushed, and there will be no
                          more output packets
  @retval AVERROR(EINVAL) codec not opened, or it is a decoder
  @retval "another negative error code" legitimate encoding errors
*/
func AVCodecReceivePacket(avctx *AVCodecContext, avpkt *AVPacket) (int, error) {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpavpkt *C.AVPacket
	if avpkt != nil {
		tmpavpkt = avpkt.ptr
	}
	ret := C.avcodec_receive_packet(tmpavctx, tmpavpkt)
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_get_hw_frames_parameters ---

// AVCodecGetHwFramesParameters wraps avcodec_get_hw_frames_parameters.
/*
  Create and return a AVHWFramesContext with values adequate for hardware
  decoding. This is meant to get called from the get_format callback, and is
  a helper for preparing a AVHWFramesContext for AVCodecContext.hw_frames_ctx.
  This API is for decoding with certain hardware acceleration modes/APIs only.

  The returned AVHWFramesContext is not initialized. The caller must do this
  with av_hwframe_ctx_init().

  Calling this function is not a requirement, but makes it simpler to avoid
  codec or hardware API specific details when manually allocating frames.

  Alternatively to this, an API user can set AVCodecContext.hw_device_ctx,
  which sets up AVCodecContext.hw_frames_ctx fully automatically, and makes
  it unnecessary to call this function or having to care about
  AVHWFramesContext initialization at all.

  There are a number of requirements for calling this function:

  - It must be called from get_format with the same avctx parameter that was
    passed to get_format. Calling it outside of get_format is not allowed, and
    can trigger undefined behavior.
  - The function is not always supported (see description of return values).
    Even if this function returns successfully, hwaccel initialization could
    fail later. (The degree to which implementations check whether the stream
    is actually supported varies. Some do this check only after the user's
    get_format callback returns.)
  - The hw_pix_fmt must be one of the choices suggested by get_format. If the
    user decides to use a AVHWFramesContext prepared with this API function,
    the user must return the same hw_pix_fmt from get_format.
  - The device_ref passed to this function must support the given hw_pix_fmt.
  - After calling this API function, it is the user's responsibility to
    initialize the AVHWFramesContext (returned by the out_frames_ref parameter),
    and to set AVCodecContext.hw_frames_ctx to it. If done, this must be done
    before returning from get_format (this is implied by the normal
    AVCodecContext.hw_frames_ctx API rules).
  - The AVHWFramesContext parameters may change every time time get_format is
    called. Also, AVCodecContext.hw_frames_ctx is reset before get_format. So
    you are inherently required to go through this process again on every
    get_format call.
  - It is perfectly possible to call this function without actually using
    the resulting AVHWFramesContext. One use-case might be trying to reuse a
    previously initialized AVHWFramesContext, and calling this API function
    only to test whether the required frame parameters have changed.
  - Fields that use dynamically allocated values of any kind must not be set
    by the user unless setting them is explicitly allowed by the documentation.
    If the user sets AVHWFramesContext.free and AVHWFramesContext.user_opaque,
    the new free callback must call the potentially set previous free callback.
    This API call may set any dynamically allocated fields, including the free
    callback.

  The function will set at least the following fields on AVHWFramesContext
  (potentially more, depending on hwaccel API):

  - All fields set by av_hwframe_ctx_alloc().
  - Set the format field to hw_pix_fmt.
  - Set the sw_format field to the most suited and most versatile format. (An
    implication is that this will prefer generic formats over opaque formats
    with arbitrary restrictions, if possible.)
  - Set the width/height fields to the coded frame size, rounded up to the
    API-specific minimum alignment.
  - Only _if_ the hwaccel requires a pre-allocated pool: set the initial_pool_size
    field to the number of maximum reference surfaces possible with the codec,
    plus 1 surface for the user to work (meaning the user can safely reference
    at most 1 decoded surface at a time), plus additional buffering introduced
    by frame threading. If the hwaccel does not require pre-allocation, the
    field is left to 0, and the decoder will allocate new surfaces on demand
    during decoding.
  - Possibly AVHWFramesContext.hwctx fields, depending on the underlying
    hardware API.

  Essentially, out_frames_ref returns the same as av_hwframe_ctx_alloc(), but
  with basic frame parameters set.

  The function is stateless, and does not change the AVCodecContext or the
  device_ref AVHWDeviceContext.

  @param avctx The context which is currently calling get_format, and which
               implicitly contains all state needed for filling the returned
               AVHWFramesContext properly.
  @param device_ref A reference to the AVHWDeviceContext describing the device
                    which will be used by the hardware decoder.
  @param hw_pix_fmt The hwaccel format you are going to return from get_format.
  @param out_frames_ref On success, set to a reference to an _uninitialized_
                        AVHWFramesContext, created from the given device_ref.
                        Fields will be set to values required for decoding.
                        Not changed if an error is returned.
  @return zero on success, a negative value on error. The following error codes
          have special semantics:
       AVERROR(ENOENT): the decoder does not support this functionality. Setup
                        is always manual, or it is a decoder which does not
                        support setting AVCodecContext.hw_frames_ctx at all,
                        or it is a software format.
       AVERROR(EINVAL): it is known that hardware decoding is not supported for
                        this configuration, or the device_ref is not supported
                        for the hwaccel referenced by hw_pix_fmt.
*/
func AVCodecGetHwFramesParameters(avctx *AVCodecContext, deviceRef *AVBufferRef, hwPixFmt AVPixelFormat, outFramesRef **AVBufferRef) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_parser_iterate ---

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
func AVCodecEncodeSubtitle(avctx *AVCodecContext, buf unsafe.Pointer, bufSize int, sub *AVSubtitle) (int, error) {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	var tmpsub *C.AVSubtitle
	if sub != nil {
		tmpsub = sub.ptr
	}
	ret := C.avcodec_encode_subtitle(tmpavctx, (*C.uint8_t)(buf), C.int(bufSize), tmpsub)
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_pix_fmt_to_codec_tag ---

// AVCodecPixFmtToCodecTag wraps avcodec_pix_fmt_to_codec_tag.
/*
  Return a value representing the fourCC code associated to the
  pixel format pix_fmt, or 0 if no associated fourCC code can be
  found.
*/
func AVCodecPixFmtToCodecTag(pixFmt AVPixelFormat) uint {
	ret := C.avcodec_pix_fmt_to_codec_tag(C.enum_AVPixelFormat(pixFmt))
	return uint(ret)
}

// --- Function avcodec_find_best_pix_fmt_of_list ---

// avcodec_find_best_pix_fmt_of_list skipped due to pixFmtList

// --- Function avcodec_default_get_format ---

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
/*
  Fill AVFrame audio data and linesize pointers.

  The buffer buf must be a preallocated buffer with a size big enough
  to contain the specified samples amount. The filled AVFrame data
  pointers will point to this buffer.

  AVFrame extended_data channel pointers are allocated if necessary for
  planar audio.

  @param frame       the AVFrame
                     frame->nb_samples must be set prior to calling the
                     function. This function fills in frame->data,
                     frame->extended_data, frame->linesize[0].
  @param nb_channels channel count
  @param sample_fmt  sample format
  @param buf         buffer to use for frame data
  @param buf_size    size of buffer
  @param align       plane size sample alignment (0 = default)
  @return            >=0 on success, negative error code on failure
  @todo return the size in bytes required to store the samples in
  case of success, at the next libavutil bump
*/
func AVCodecFillAudioFrame(frame *AVFrame, nbChannels int, sampleFmt AVSampleFormat, buf unsafe.Pointer, bufSize int, align int) (int, error) {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.avcodec_fill_audio_frame(tmpframe, C.int(nbChannels), C.enum_AVSampleFormat(sampleFmt), (*C.uint8_t)(buf), C.int(bufSize), C.int(align))
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_flush_buffers ---

// AVCodecFlushBuffers wraps avcodec_flush_buffers.
/*
  Reset the internal codec state / flush internal buffers. Should be called
  e.g. when seeking or when switching to a different stream.

  @note for decoders, this function just releases any references the decoder
  might keep internally, but the caller's references remain valid.

  @note for encoders, this function will only do something if the encoder
  declares support for AV_CODEC_CAP_ENCODER_FLUSH. When called, the encoder
  will drain any remaining packets, and can then be re-used for a different
  stream (as opposed to sending a null frame which will leave the encoder
  in a permanent EOF state after draining). This can be desirable if the
  cost of tearing down and replacing the encoder instance is high.
*/
func AVCodecFlushBuffers(avctx *AVCodecContext) {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	C.avcodec_flush_buffers(tmpavctx)
}

// --- Function av_get_audio_frame_duration ---

// AVGetAudioFrameDuration wraps av_get_audio_frame_duration.
/*
  Return audio frame duration.

  @param avctx        codec context
  @param frame_bytes  size of the frame, or 0 if unknown
  @return             frame duration, in samples, if known. 0 if not able to
                      determine.
*/
func AVGetAudioFrameDuration(avctx *AVCodecContext, frameBytes int) (int, error) {
	var tmpavctx *C.AVCodecContext
	if avctx != nil {
		tmpavctx = avctx.ptr
	}
	ret := C.av_get_audio_frame_duration(tmpavctx, C.int(frameBytes))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_fast_padded_malloc ---

// av_fast_padded_malloc skipped due to size

// --- Function av_fast_padded_mallocz ---

// av_fast_padded_mallocz skipped due to size

// --- Function avcodec_is_open ---

// AVCodecIsOpen wraps avcodec_is_open.
/*
  @return a positive value if s is open (i.e. avcodec_open2() was called on it
  with no corresponding avcodec_close()), 0 otherwise.
*/
func AVCodecIsOpen(s *AVCodecContext) (int, error) {
	var tmps *C.AVCodecContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avcodec_is_open(tmps)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_codec_iterate ---

// av_codec_iterate skipped due to opaque

// --- Function avcodec_find_decoder ---

// AVCodecFindDecoder wraps avcodec_find_decoder.
/*
  Find a registered decoder with a matching codec ID.

  @param id AVCodecID of the requested decoder
  @return A decoder if one was found, NULL otherwise.
*/
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
/*
  Find a registered decoder with the specified name.

  @param name name of the requested decoder
  @return A decoder if one was found, NULL otherwise.
*/
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
/*
  Find a registered encoder with a matching codec ID.

  @param id AVCodecID of the requested encoder
  @return An encoder if one was found, NULL otherwise.
*/
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
/*
  Find a registered encoder with the specified name.

  @param name name of the requested encoder
  @return An encoder if one was found, NULL otherwise.
*/
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
//
//	@return a non-zero number if codec is an encoder, zero otherwise
func AVCodecIsEncoder(codec *AVCodec) (int, error) {
	var tmpcodec *C.AVCodec
	if codec != nil {
		tmpcodec = codec.ptr
	}
	ret := C.av_codec_is_encoder(tmpcodec)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_codec_is_decoder ---

// AVCodecIsDecoder wraps av_codec_is_decoder.
//
//	@return a non-zero number if codec is a decoder, zero otherwise
func AVCodecIsDecoder(codec *AVCodec) (int, error) {
	var tmpcodec *C.AVCodec
	if codec != nil {
		tmpcodec = codec.ptr
	}
	ret := C.av_codec_is_decoder(tmpcodec)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_get_profile_name ---

// AVGetProfileName wraps av_get_profile_name.
/*
  Return a name for the specified profile, if available.

  @param codec the codec that is searched for the given profile
  @param profile the profile value for which a name is requested
  @return A name for the profile if found, NULL otherwise.
*/
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
/*
  Retrieve supported hardware configurations for a codec.

  Values of index from zero to some maximum return the indexed configuration
  descriptor; all other values return NULL.  If the codec does not support
  any hardware configurations then it will always return NULL.
*/
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
//
//	@return descriptor for given codec ID or NULL if no descriptor exists.
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
/*
  Iterate over all codec descriptors known to libavcodec.

  @param prev previous descriptor. NULL to get the first descriptor.

  @return next descriptor or NULL after the last descriptor
*/
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
/*
  @return codec descriptor with the given name or NULL if no such descriptor
          exists.
*/
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
//
//	Get the type of the given codec.
func AVCodecGetType(codecId AVCodecID) AVMediaType {
	ret := C.avcodec_get_type(C.enum_AVCodecID(codecId))
	return AVMediaType(ret)
}

// --- Function avcodec_get_name ---

// AVCodecGetName wraps avcodec_get_name.
/*
  Get the name of a codec.
  @return  a static string identifying the codec; never NULL
*/
func AVCodecGetName(id AVCodecID) *CStr {
	ret := C.avcodec_get_name(C.enum_AVCodecID(id))
	return wrapCStr(ret)
}

// --- Function av_get_bits_per_sample ---

// AVGetBitsPerSample wraps av_get_bits_per_sample.
/*
  Return codec bits per sample.

  @param[in] codec_id the codec
  @return Number of bits per sample or zero if unknown for the given codec.
*/
func AVGetBitsPerSample(codecId AVCodecID) (int, error) {
	ret := C.av_get_bits_per_sample(C.enum_AVCodecID(codecId))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_get_exact_bits_per_sample ---

// AVGetExactBitsPerSample wraps av_get_exact_bits_per_sample.
/*
  Return codec bits per sample.
  Only return non-zero if the bits per sample is exactly correct, not an
  approximation.

  @param[in] codec_id the codec
  @return Number of bits per sample or zero if unknown for the given codec.
*/
func AVGetExactBitsPerSample(codecId AVCodecID) (int, error) {
	ret := C.av_get_exact_bits_per_sample(C.enum_AVCodecID(codecId))
	return int(ret), WrapErr(int(ret))
}

// --- Function avcodec_profile_name ---

// AVCodecProfileName wraps avcodec_profile_name.
/*
  Return a name for the specified profile, if available.

  @param codec_id the ID of the codec to which the requested profile belongs
  @param profile the profile value for which a name is requested
  @return A name for the profile if found, NULL otherwise.

  @note unlike av_get_profile_name(), which searches a list of profiles
        supported by a specific decoder or encoder implementation, this
        function searches the list of profiles from the AVCodecDescriptor
*/
func AVCodecProfileName(codecId AVCodecID, profile int) *CStr {
	ret := C.avcodec_profile_name(C.enum_AVCodecID(codecId), C.int(profile))
	return wrapCStr(ret)
}

// --- Function av_get_pcm_codec ---

// AVGetPcmCodec wraps av_get_pcm_codec.
/*
  Return the PCM codec associated with a sample format.
  @param be  endianness, 0 for little, 1 for big,
             -1 (or anything else) for native
  @return  AV_CODEC_ID_PCM_* or AV_CODEC_ID_NONE
*/
func AVGetPcmCodec(fmt AVSampleFormat, be int) AVCodecID {
	ret := C.av_get_pcm_codec(C.enum_AVSampleFormat(fmt), C.int(be))
	return AVCodecID(ret)
}

// --- Function avcodec_parameters_alloc ---

// AVCodecParametersAlloc wraps avcodec_parameters_alloc.
/*
  Allocate a new AVCodecParameters and set its fields to default values
  (unknown/invalid/0). The returned struct must be freed with
  avcodec_parameters_free().
*/
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
/*
  Free an AVCodecParameters instance and everything associated with it and
  write NULL to the supplied pointer.
*/
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
/*
  Copy the contents of src to dst. Any allocated fields in dst are freed and
  replaced with newly allocated duplicates of the corresponding fields in src.

  @return >= 0 on success, a negative AVERROR code on failure.
*/
func AVCodecParametersCopy(dst *AVCodecParameters, src *AVCodecParameters) (int, error) {
	var tmpdst *C.AVCodecParameters
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVCodecParameters
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.avcodec_parameters_copy(tmpdst, tmpsrc)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_get_audio_frame_duration2 ---

// AVGetAudioFrameDuration2 wraps av_get_audio_frame_duration2.
/*
  This function is the same as av_get_audio_frame_duration(), except it works
  with AVCodecParameters instead of an AVCodecContext.
*/
func AVGetAudioFrameDuration2(par *AVCodecParameters, frameBytes int) (int, error) {
	var tmppar *C.AVCodecParameters
	if par != nil {
		tmppar = par.ptr
	}
	ret := C.av_get_audio_frame_duration2(tmppar, C.int(frameBytes))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_cpb_properties_alloc ---

// av_cpb_properties_alloc skipped due to size

// --- Function av_xiphlacing ---

// AVXiphlacing wraps av_xiphlacing.
/*
  Encode extradata length to a buffer. Used by xiph codecs.

  @param s buffer to write to; must be at least (v/255+1) bytes long
  @param v size of extradata in bytes
  @return number of bytes written to the buffer.
*/
func AVXiphlacing(s unsafe.Pointer, v uint) uint {
	ret := C.av_xiphlacing((*C.uchar)(s), C.uint(v))
	return uint(ret)
}

// --- Function av_packet_alloc ---

// AVPacketAlloc wraps av_packet_alloc.
/*
  Allocate an AVPacket and set its fields to default values.  The resulting
  struct must be freed using av_packet_free().

  @return An AVPacket filled with default values or NULL on failure.

  @note this only allocates the AVPacket itself, not the data buffers. Those
  must be allocated through other means such as av_new_packet.

  @see av_new_packet
*/
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
/*
  Create a new packet that references the same data as src.

  This is a shortcut for av_packet_alloc()+av_packet_ref().

  @return newly created AVPacket on success, NULL on error.

  @see av_packet_alloc
  @see av_packet_ref
*/
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
/*
  Free the packet, if the packet is reference counted, it will be
  unreferenced first.

  @param pkt packet to be freed. The pointer will be set to NULL.
  @note passing NULL is a no-op.
*/
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
/*
  Initialize optional fields of a packet with default values.

  Note, this does not touch the data and size members, which have to be
  initialized separately.

  @param pkt packet

  @see av_packet_alloc
  @see av_packet_unref

  @deprecated This function is deprecated. Once it's removed,
  sizeof(AVPacket) will not be a part of the ABI anymore.
*/
func AVInitPacket(pkt *AVPacket) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	C.av_init_packet(tmppkt)
}

// --- Function av_new_packet ---

// AVNewPacket wraps av_new_packet.
/*
  Allocate the payload of a packet and initialize its fields with
  default values.

  @param pkt packet
  @param size wanted payload size
  @return 0 if OK, AVERROR_xxx otherwise
*/
func AVNewPacket(pkt *AVPacket, size int) (int, error) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_new_packet(tmppkt, C.int(size))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_shrink_packet ---

// AVShrinkPacket wraps av_shrink_packet.
/*
  Reduce packet size, correctly zeroing padding

  @param pkt packet
  @param size new size
*/
func AVShrinkPacket(pkt *AVPacket, size int) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	C.av_shrink_packet(tmppkt, C.int(size))
}

// --- Function av_grow_packet ---

// AVGrowPacket wraps av_grow_packet.
/*
  Increase packet size, correctly zeroing padding

  @param pkt packet
  @param grow_by number of bytes by which to increase the size of the packet
*/
func AVGrowPacket(pkt *AVPacket, growBy int) (int, error) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_grow_packet(tmppkt, C.int(growBy))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_packet_from_data ---

// AVPacketFromData wraps av_packet_from_data.
/*
  Initialize a reference-counted packet from av_malloc()ed data.

  @param pkt packet to be initialized. This function will set the data, size,
         and buf fields, all others are left untouched.
  @param data Data allocated by av_malloc() to be used as packet data. If this
         function returns successfully, the data is owned by the underlying AVBuffer.
         The caller may not access the data through other means.
  @param size size of data in bytes, without the padding. I.e. the full buffer
         size is assumed to be size + AV_INPUT_BUFFER_PADDING_SIZE.

  @return 0 on success, a negative AVERROR on error
*/
func AVPacketFromData(pkt *AVPacket, data unsafe.Pointer, size int) (int, error) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_packet_from_data(tmppkt, (*C.uint8_t)(data), C.int(size))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_packet_new_side_data ---

// AVPacketNewSideData wraps av_packet_new_side_data.
/*
  Allocate new information of a packet.

  @param pkt packet
  @param type side information type
  @param size side information size
  @return pointer to fresh allocated data or NULL otherwise
*/
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
/*
  Wrap an existing array as a packet side data.

  @param pkt packet
  @param type side information type
  @param data the side data array. It must be allocated with the av_malloc()
              family of functions. The ownership of the data is transferred to
              pkt.
  @param size side information size
  @return a non-negative number on success, a negative AVERROR code on
          failure. On failure, the packet is unchanged and the data remains
          owned by the caller.
*/
func AVPacketAddSideData(pkt *AVPacket, _type AVPacketSideDataType, data unsafe.Pointer, size uint64) (int, error) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_packet_add_side_data(tmppkt, C.enum_AVPacketSideDataType(_type), (*C.uint8_t)(data), C.size_t(size))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_packet_shrink_side_data ---

// AVPacketShrinkSideData wraps av_packet_shrink_side_data.
/*
  Shrink the already allocated side data buffer

  @param pkt packet
  @param type side information type
  @param size new side information size
  @return 0 on success, < 0 on failure
*/
func AVPacketShrinkSideData(pkt *AVPacket, _type AVPacketSideDataType, size uint64) (int, error) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_packet_shrink_side_data(tmppkt, C.enum_AVPacketSideDataType(_type), C.size_t(size))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_packet_get_side_data ---

// av_packet_get_side_data skipped due to size

// --- Function av_packet_side_data_name ---

// AVPacketSideDataName wraps av_packet_side_data_name.
func AVPacketSideDataName(_type AVPacketSideDataType) *CStr {
	ret := C.av_packet_side_data_name(C.enum_AVPacketSideDataType(_type))
	return wrapCStr(ret)
}

// --- Function av_packet_pack_dictionary ---

// av_packet_pack_dictionary skipped due to size

// --- Function av_packet_unpack_dictionary ---

// AVPacketUnpackDictionary wraps av_packet_unpack_dictionary.
/*
  Unpack a dictionary from side_data.

  @param data data from side_data
  @param size size of the data
  @param dict the metadata storage dictionary
  @return 0 on success, < 0 on failure
*/
func AVPacketUnpackDictionary(data unsafe.Pointer, size uint64, dict **AVDictionary) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_packet_free_side_data ---

// AVPacketFreeSideData wraps av_packet_free_side_data.
/*
  Convenience function to free all the side data stored.
  All the other fields stay untouched.

  @param pkt packet
*/
func AVPacketFreeSideData(pkt *AVPacket) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	C.av_packet_free_side_data(tmppkt)
}

// --- Function av_packet_ref ---

// AVPacketRef wraps av_packet_ref.
/*
  Setup a new reference to the data described by a given packet

  If src is reference-counted, setup dst as a new reference to the
  buffer in src. Otherwise allocate a new buffer in dst and copy the
  data from src into it.

  All the other fields are copied from src.

  @see av_packet_unref

  @param dst Destination packet. Will be completely overwritten.
  @param src Source packet

  @return 0 on success, a negative AVERROR on error. On error, dst
          will be blank (as if returned by av_packet_alloc()).
*/
func AVPacketRef(dst *AVPacket, src *AVPacket) (int, error) {
	var tmpdst *C.AVPacket
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVPacket
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_packet_ref(tmpdst, tmpsrc)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_packet_unref ---

// AVPacketUnref wraps av_packet_unref.
/*
  Wipe the packet.

  Unreference the buffer referenced by the packet and reset the
  remaining packet fields to their default values.

  @param pkt The packet to be unreferenced.
*/
func AVPacketUnref(pkt *AVPacket) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	C.av_packet_unref(tmppkt)
}

// --- Function av_packet_move_ref ---

// AVPacketMoveRef wraps av_packet_move_ref.
/*
  Move every field in src to dst and reset src.

  @see av_packet_unref

  @param src Source packet, will be reset
  @param dst Destination packet
*/
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
/*
  Copy only "properties" fields from src to dst.

  Properties for the purpose of this function are all the fields
  beside those related to the packet data (buf, data, size)

  @param dst Destination packet
  @param src Source packet

  @return 0 on success AVERROR on failure.
*/
func AVPacketCopyProps(dst *AVPacket, src *AVPacket) (int, error) {
	var tmpdst *C.AVPacket
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVPacket
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_packet_copy_props(tmpdst, tmpsrc)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_packet_make_refcounted ---

// AVPacketMakeRefcounted wraps av_packet_make_refcounted.
/*
  Ensure the data described by a given packet is reference counted.

  @note This function does not ensure that the reference will be writable.
        Use av_packet_make_writable instead for that purpose.

  @see av_packet_ref
  @see av_packet_make_writable

  @param pkt packet whose data should be made reference counted.

  @return 0 on success, a negative AVERROR on error. On failure, the
          packet is unchanged.
*/
func AVPacketMakeRefcounted(pkt *AVPacket) (int, error) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_packet_make_refcounted(tmppkt)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_packet_make_writable ---

// AVPacketMakeWritable wraps av_packet_make_writable.
/*
  Create a writable reference for the data described by a given packet,
  avoiding data copy if possible.

  @param pkt Packet whose data should be made writable.

  @return 0 on success, a negative AVERROR on failure. On failure, the
          packet is unchanged.
*/
func AVPacketMakeWritable(pkt *AVPacket) (int, error) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_packet_make_writable(tmppkt)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_packet_rescale_ts ---

// AVPacketRescaleTs wraps av_packet_rescale_ts.
/*
  Convert valid timing fields (timestamps / durations) in a packet from one
  timebase to another. Timestamps with unknown values (AV_NOPTS_VALUE) will be
  ignored.

  @param pkt packet on which the conversion will be performed
  @param tb_src source timebase, in which the timing fields in pkt are
                expressed
  @param tb_dst destination timebase, to which the timing fields will be
                converted
*/
func AVPacketRescaleTs(pkt *AVPacket, tbSrc *AVRational, tbDst *AVRational) {
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	C.av_packet_rescale_ts(tmppkt, tbSrc.value, tbDst.value)
}

// --- Function avfilter_version ---

// AVFilterVersion wraps avfilter_version.
//
//	Return the LIBAVFILTER_VERSION_INT constant.
func AVFilterVersion() uint {
	ret := C.avfilter_version()
	return uint(ret)
}

// --- Function avfilter_configuration ---

// AVFilterConfiguration wraps avfilter_configuration.
//
//	Return the libavfilter build-time configuration.
func AVFilterConfiguration() *CStr {
	ret := C.avfilter_configuration()
	return wrapCStr(ret)
}

// --- Function avfilter_license ---

// AVFilterLicense wraps avfilter_license.
//
//	Return the libavfilter license.
func AVFilterLicense() *CStr {
	ret := C.avfilter_license()
	return wrapCStr(ret)
}

// --- Function avfilter_pad_get_name ---

// AVFilterPadGetName wraps avfilter_pad_get_name.
/*
  Get the name of an AVFilterPad.

  @param pads an array of AVFilterPads
  @param pad_idx index of the pad in the array; it is the caller's
                 responsibility to ensure the index is valid

  @return name of the pad_idx'th pad in pads
*/
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
/*
  Get the type of an AVFilterPad.

  @param pads an array of AVFilterPads
  @param pad_idx index of the pad in the array; it is the caller's
                 responsibility to ensure the index is valid

  @return type of the pad_idx'th pad in pads
*/
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
//
//	Get the number of elements in an AVFilter's inputs or outputs array.
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
/*
  Link two filters together.

  @param src    the source filter
  @param srcpad index of the output pad on the source filter
  @param dst    the destination filter
  @param dstpad index of the input pad on the destination filter
  @return       zero on success
*/
func AVFilterLink_(src *AVFilterContext, srcpad uint, dst *AVFilterContext, dstpad uint) (int, error) {
	var tmpsrc *C.AVFilterContext
	if src != nil {
		tmpsrc = src.ptr
	}
	var tmpdst *C.AVFilterContext
	if dst != nil {
		tmpdst = dst.ptr
	}
	ret := C.avfilter_link(tmpsrc, C.uint(srcpad), tmpdst, C.uint(dstpad))
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_link_free ---

// AVFilterLinkFree wraps avfilter_link_free.
//
//	Free the link in *link, and set its pointer to NULL.
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
/*
  Negotiate the media format, dimensions, etc of all inputs to a filter.

  @param filter the filter to negotiate the properties for its inputs
  @return       zero on successful negotiation
*/
func AVFilterConfigLinks(filter *AVFilterContext) (int, error) {
	var tmpfilter *C.AVFilterContext
	if filter != nil {
		tmpfilter = filter.ptr
	}
	ret := C.avfilter_config_links(tmpfilter)
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_process_command ---

// AVFilterProcessCommand wraps avfilter_process_command.
/*
  Make the filter instance process a command.
  It is recommended to use avfilter_graph_send_command().
*/
func AVFilterProcessCommand(filter *AVFilterContext, cmd *CStr, arg *CStr, res *CStr, resLen int, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_filter_iterate ---

// av_filter_iterate skipped due to opaque

// --- Function avfilter_get_by_name ---

// AVFilterGetByName wraps avfilter_get_by_name.
/*
  Get a filter definition matching the given name.

  @param name the filter name to find
  @return     the filter definition, if any matching one is registered.
              NULL if none found.
*/
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
/*
  Initialize a filter with the supplied parameters.

  @param ctx  uninitialized filter context to initialize
  @param args Options to initialize the filter with. This must be a
              ':'-separated list of options in the 'key=value' form.
              May be NULL if the options have been set directly using the
              AVOptions API or there are no options that need to be set.
  @return 0 on success, a negative AVERROR on failure
*/
func AVFilterInitStr(ctx *AVFilterContext, args *CStr) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpargs *C.char
	if args != nil {
		tmpargs = args.ptr
	}
	ret := C.avfilter_init_str(tmpctx, tmpargs)
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_init_dict ---

// AVFilterInitDict wraps avfilter_init_dict.
/*
  Initialize a filter with the supplied dictionary of options.

  @param ctx     uninitialized filter context to initialize
  @param options An AVDictionary filled with options for this filter. On
                 return this parameter will be destroyed and replaced with
                 a dict containing options that were not found. This dictionary
                 must be freed by the caller.
                 May be NULL, then this function is equivalent to
                 avfilter_init_str() with the second parameter set to NULL.
  @return 0 on success, a negative AVERROR on failure

  @note This function and avfilter_init_str() do essentially the same thing,
  the difference is in manner in which the options are passed. It is up to the
  calling code to choose whichever is more preferable. The two functions also
  behave differently when some of the provided options are not declared as
  supported by the filter. In such a case, avfilter_init_str() will fail, but
  this function will leave those extra options in the options AVDictionary and
  continue as usual.
*/
func AVFilterInitDict(ctx *AVFilterContext, options **AVDictionary) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_free ---

// AVFilterFree wraps avfilter_free.
/*
  Free a filter context. This will also remove the filter from its
  filtergraph's list of filters.

  @param filter the filter to free
*/
func AVFilterFree(filter *AVFilterContext) {
	var tmpfilter *C.AVFilterContext
	if filter != nil {
		tmpfilter = filter.ptr
	}
	C.avfilter_free(tmpfilter)
}

// --- Function avfilter_insert_filter ---

// AVFilterInsertFilter wraps avfilter_insert_filter.
/*
  Insert a filter in the middle of an existing link.

  @param link the link into which the filter should be inserted
  @param filt the filter to be inserted
  @param filt_srcpad_idx the input pad on the filter to connect
  @param filt_dstpad_idx the output pad on the filter to connect
  @return     zero on success
*/
func AVFilterInsertFilter(link *AVFilterLink, filt *AVFilterContext, filtSrcpadIdx uint, filtDstpadIdx uint) (int, error) {
	var tmplink *C.AVFilterLink
	if link != nil {
		tmplink = link.ptr
	}
	var tmpfilt *C.AVFilterContext
	if filt != nil {
		tmpfilt = filt.ptr
	}
	ret := C.avfilter_insert_filter(tmplink, tmpfilt, C.uint(filtSrcpadIdx), C.uint(filtDstpadIdx))
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_get_class ---

// AVFilterGetClass wraps avfilter_get_class.
/*
  @return AVClass for AVFilterContext.

  @see av_opt_find().
*/
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
/*
  Allocate a filter graph.

  @return the allocated filter graph on success or NULL.
*/
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
/*
  Create a new filter instance in a filter graph.

  @param graph graph in which the new filter will be used
  @param filter the filter to create an instance of
  @param name Name to give to the new instance (will be copied to
              AVFilterContext.name). This may be used by the caller to identify
              different filters, libavfilter itself assigns no semantics to
              this parameter. May be NULL.

  @return the context of the newly created filter instance (note that it is
          also retrievable directly through AVFilterGraph.filters or with
          avfilter_graph_get_filter()) on success or NULL on failure.
*/
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
/*
  Get a filter instance identified by instance name from graph.

  @param graph filter graph to search through.
  @param name filter instance name (should be unique in the graph).
  @return the pointer to the found filter instance or NULL if it
  cannot be found.
*/
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
/*
  Create and add a filter instance into an existing graph.
  The filter instance is created from the filter filt and inited
  with the parameter args. opaque is currently ignored.

  In case of success put in *filt_ctx the pointer to the created
  filter instance, otherwise set *filt_ctx to NULL.

  @param name the instance name to give to the created filter instance
  @param graph_ctx the filter graph
  @return a negative AVERROR error code in case of failure, a non
  negative value otherwise
*/
func AVFilterGraphCreateFilter(filtCtx **AVFilterContext, filt *AVFilter, name *CStr, args *CStr, opaque unsafe.Pointer, graphCtx *AVFilterGraph) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_set_auto_convert ---

// AVFilterGraphSetAutoConvert wraps avfilter_graph_set_auto_convert.
/*
  Enable or disable automatic format conversion inside the graph.

  Note that format conversion can still happen inside explicitly inserted
  scale and aresample filters.

  @param flags  any of the AVFILTER_AUTO_CONVERT_* constants
*/
func AVFilterGraphSetAutoConvert(graph *AVFilterGraph, flags uint) {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	C.avfilter_graph_set_auto_convert(tmpgraph, C.uint(flags))
}

// --- Function avfilter_graph_config ---

// AVFilterGraphConfig wraps avfilter_graph_config.
/*
  Check validity and configure all the links and formats in the graph.

  @param graphctx the filter graph
  @param log_ctx context used for logging
  @return >= 0 in case of success, a negative AVERROR code otherwise
*/
func AVFilterGraphConfig(graphctx *AVFilterGraph, logCtx unsafe.Pointer) (int, error) {
	var tmpgraphctx *C.AVFilterGraph
	if graphctx != nil {
		tmpgraphctx = graphctx.ptr
	}
	ret := C.avfilter_graph_config(tmpgraphctx, logCtx)
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_free ---

// AVFilterGraphFree wraps avfilter_graph_free.
/*
  Free a graph, destroy its links, and set *graph to NULL.
  If *graph is NULL, do nothing.
*/
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
/*
  Allocate a single AVFilterInOut entry.
  Must be freed with avfilter_inout_free().
  @return allocated AVFilterInOut on success, NULL on failure.
*/
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
/*
  Free the supplied list of AVFilterInOut and set *inout to NULL.
  If *inout is NULL, do nothing.
*/
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
/*
  Add a graph described by a string to a graph.

  @note The caller must provide the lists of inputs and outputs,
  which therefore must be known before calling the function.

  @note The inputs parameter describes inputs of the already existing
  part of the graph; i.e. from the point of view of the newly created
  part, they are outputs. Similarly the outputs parameter describes
  outputs of the already existing filters, which are provided as
  inputs to the parsed filters.

  @param graph   the filter graph where to link the parsed graph context
  @param filters string to be parsed
  @param inputs  linked list to the inputs of the graph
  @param outputs linked list to the outputs of the graph
  @return zero on success, a negative AVERROR code on error
*/
func AVFilterGraphParse(graph *AVFilterGraph, filters *CStr, inputs *AVFilterInOut, outputs *AVFilterInOut, logCtx unsafe.Pointer) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_parse_ptr ---

// AVFilterGraphParsePtr wraps avfilter_graph_parse_ptr.
/*
  Add a graph described by a string to a graph.

  In the graph filters description, if the input label of the first
  filter is not specified, "in" is assumed; if the output label of
  the last filter is not specified, "out" is assumed.

  @param graph   the filter graph where to link the parsed graph context
  @param filters string to be parsed
  @param inputs  pointer to a linked list to the inputs of the graph, may be NULL.
                 If non-NULL, *inputs is updated to contain the list of open inputs
                 after the parsing, should be freed with avfilter_inout_free().
  @param outputs pointer to a linked list to the outputs of the graph, may be NULL.
                 If non-NULL, *outputs is updated to contain the list of open outputs
                 after the parsing, should be freed with avfilter_inout_free().
  @return non negative on success, a negative AVERROR code on error
*/
func AVFilterGraphParsePtr(graph *AVFilterGraph, filters *CStr, inputs **AVFilterInOut, outputs **AVFilterInOut, logCtx unsafe.Pointer) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_parse2 ---

// AVFilterGraphParse2 wraps avfilter_graph_parse2.
/*
  Add a graph described by a string to a graph.

  @param[in]  graph   the filter graph where to link the parsed graph context
  @param[in]  filters string to be parsed
  @param[out] inputs  a linked list of all free (unlinked) inputs of the
                      parsed graph will be returned here. It is to be freed
                      by the caller using avfilter_inout_free().
  @param[out] outputs a linked list of all free (unlinked) outputs of the
                      parsed graph will be returned here. It is to be freed by the
                      caller using avfilter_inout_free().
  @return zero on success, a negative AVERROR code on error

  @note This function returns the inputs and outputs that are left
  unlinked after parsing the graph and the caller then deals with
  them.
  @note This function makes no reference whatsoever to already
  existing parts of the graph and the inputs parameter will on return
  contain inputs of the newly parsed part of the graph.  Analogously
  the outputs parameter will contain outputs of the newly created
  filters.
*/
func AVFilterGraphParse2(graph *AVFilterGraph, filters *CStr, inputs **AVFilterInOut, outputs **AVFilterInOut) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_segment_parse ---

// AVFilterGraphSegmentParse wraps avfilter_graph_segment_parse.
/*
  Parse a textual filtergraph description into an intermediate form.

  This intermediate representation is intended to be modified by the caller as
  described in the documentation of AVFilterGraphSegment and its children, and
  then applied to the graph either manually or with other
  avfilter_graph_segment_*() functions. See the documentation for
  avfilter_graph_segment_apply() for the canonical way to apply
  AVFilterGraphSegment.

  @param graph Filter graph the parsed segment is associated with. Will only be
               used for logging and similar auxiliary purposes. The graph will
               not be actually modified by this function - the parsing results
               are instead stored in seg for further processing.
  @param graph_str a string describing the filtergraph segment
  @param flags reserved for future use, caller must set to 0 for now
  @param seg A pointer to the newly-created AVFilterGraphSegment is written
             here on success. The graph segment is owned by the caller and must
             be freed with avfilter_graph_segment_free() before graph itself is
             freed.

  @retval "non-negative number" success
  @retval "negative error code" failure
*/
func AVFilterGraphSegmentParse(graph *AVFilterGraph, graphStr *CStr, flags int, seg **AVFilterGraphSegment) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_segment_create_filters ---

// AVFilterGraphSegmentCreateFilters wraps avfilter_graph_segment_create_filters.
/*
  Create filters specified in a graph segment.

  Walk through the creation-pending AVFilterParams in the segment and create
  new filter instances for them.
  Creation-pending params are those where AVFilterParams.filter_name is
  non-NULL (and hence AVFilterParams.filter is NULL). All other AVFilterParams
  instances are ignored.

  For any filter created by this function, the corresponding
  AVFilterParams.filter is set to the newly-created filter context,
  AVFilterParams.filter_name and AVFilterParams.instance_name are freed and set
  to NULL.

  @param seg the filtergraph segment to process
  @param flags reserved for future use, caller must set to 0 for now

  @retval "non-negative number" Success, all creation-pending filters were
                                successfully created
  @retval AVERROR_FILTER_NOT_FOUND some filter's name did not correspond to a
                                   known filter
  @retval "another negative error code" other failures

  @note Calling this function multiple times is safe, as it is idempotent.
*/
func AVFilterGraphSegmentCreateFilters(seg *AVFilterGraphSegment, flags int) (int, error) {
	var tmpseg *C.AVFilterGraphSegment
	if seg != nil {
		tmpseg = seg.ptr
	}
	ret := C.avfilter_graph_segment_create_filters(tmpseg, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_segment_apply_opts ---

// AVFilterGraphSegmentApplyOpts wraps avfilter_graph_segment_apply_opts.
/*
  Apply parsed options to filter instances in a graph segment.

  Walk through all filter instances in the graph segment that have option
  dictionaries associated with them and apply those options with
  av_opt_set_dict2(..., AV_OPT_SEARCH_CHILDREN). AVFilterParams.opts is
  replaced by the dictionary output by av_opt_set_dict2(), which should be
  empty (NULL) if all options were successfully applied.

  If any options could not be found, this function will continue processing all
  other filters and finally return AVERROR_OPTION_NOT_FOUND (unless another
  error happens). The calling program may then deal with unapplied options as
  it wishes.

  Any creation-pending filters (see avfilter_graph_segment_create_filters())
  present in the segment will cause this function to fail. AVFilterParams with
  no associated filter context are simply skipped.

  @param seg the filtergraph segment to process
  @param flags reserved for future use, caller must set to 0 for now

  @retval "non-negative number" Success, all options were successfully applied.
  @retval AVERROR_OPTION_NOT_FOUND some options were not found in a filter
  @retval "another negative error code" other failures

  @note Calling this function multiple times is safe, as it is idempotent.
*/
func AVFilterGraphSegmentApplyOpts(seg *AVFilterGraphSegment, flags int) (int, error) {
	var tmpseg *C.AVFilterGraphSegment
	if seg != nil {
		tmpseg = seg.ptr
	}
	ret := C.avfilter_graph_segment_apply_opts(tmpseg, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_segment_init ---

// AVFilterGraphSegmentInit wraps avfilter_graph_segment_init.
/*
  Initialize all filter instances in a graph segment.

  Walk through all filter instances in the graph segment and call
  avfilter_init_dict(..., NULL) on those that have not been initialized yet.

  Any creation-pending filters (see avfilter_graph_segment_create_filters())
  present in the segment will cause this function to fail. AVFilterParams with
  no associated filter context or whose filter context is already initialized,
  are simply skipped.

  @param seg the filtergraph segment to process
  @param flags reserved for future use, caller must set to 0 for now

  @retval "non-negative number" Success, all filter instances were successfully
                                initialized
  @retval "negative error code" failure

  @note Calling this function multiple times is safe, as it is idempotent.
*/
func AVFilterGraphSegmentInit(seg *AVFilterGraphSegment, flags int) (int, error) {
	var tmpseg *C.AVFilterGraphSegment
	if seg != nil {
		tmpseg = seg.ptr
	}
	ret := C.avfilter_graph_segment_init(tmpseg, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_segment_link ---

// AVFilterGraphSegmentLink wraps avfilter_graph_segment_link.
/*
  Link filters in a graph segment.

  Walk through all filter instances in the graph segment and try to link all
  unlinked input and output pads. Any creation-pending filters (see
  avfilter_graph_segment_create_filters()) present in the segment will cause
  this function to fail. Disabled filters and already linked pads are skipped.

  Every filter output pad that has a corresponding AVFilterPadParams with a
  non-NULL label is
  - linked to the input with the matching label, if one exists;
  - exported in the outputs linked list otherwise, with the label preserved.
  Unlabeled outputs are
  - linked to the first unlinked unlabeled input in the next non-disabled
    filter in the chain, if one exists
  - exported in the ouputs linked list otherwise, with NULL label

  Similarly, unlinked input pads are exported in the inputs linked list.

  @param seg the filtergraph segment to process
  @param flags reserved for future use, caller must set to 0 for now
  @param[out] inputs  a linked list of all free (unlinked) inputs of the
                      filters in this graph segment will be returned here. It
                      is to be freed by the caller using avfilter_inout_free().
  @param[out] outputs a linked list of all free (unlinked) outputs of the
                      filters in this graph segment will be returned here. It
                      is to be freed by the caller using avfilter_inout_free().

  @retval "non-negative number" success
  @retval "negative error code" failure

  @note Calling this function multiple times is safe, as it is idempotent.
*/
func AVFilterGraphSegmentLink(seg *AVFilterGraphSegment, flags int, inputs **AVFilterInOut, outputs **AVFilterInOut) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_segment_apply ---

// AVFilterGraphSegmentApply wraps avfilter_graph_segment_apply.
/*
  Apply all filter/link descriptions from a graph segment to the associated filtergraph.

  This functions is currently equivalent to calling the following in sequence:
  - avfilter_graph_segment_create_filters();
  - avfilter_graph_segment_apply_opts();
  - avfilter_graph_segment_init();
  - avfilter_graph_segment_link();
  failing if any of them fails. This list may be extended in the future.

  Since the above functions are idempotent, the caller may call some of them
  manually, then do some custom processing on the filtergraph, then call this
  function to do the rest.

  @param seg the filtergraph segment to process
  @param flags reserved for future use, caller must set to 0 for now
  @param[out] inputs passed to avfilter_graph_segment_link()
  @param[out] outputs passed to avfilter_graph_segment_link()

  @retval "non-negative number" success
  @retval "negative error code" failure

  @note Calling this function multiple times is safe, as it is idempotent.
*/
func AVFilterGraphSegmentApply(seg *AVFilterGraphSegment, flags int, inputs **AVFilterInOut, outputs **AVFilterInOut) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_segment_free ---

// AVFilterGraphSegmentFree wraps avfilter_graph_segment_free.
/*
  Free the provided AVFilterGraphSegment and everything associated with it.

  @param seg double pointer to the AVFilterGraphSegment to be freed. NULL will
  be written to this pointer on exit from this function.

  @note
  The filter contexts (AVFilterParams.filter) are owned by AVFilterGraph rather
  than AVFilterGraphSegment, so they are not freed.
*/
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
/*
  Send a command to one or more filter instances.

  @param graph  the filter graph
  @param target the filter(s) to which the command should be sent
                "all" sends to all filters
                otherwise it can be a filter or filter instance name
                which will send the command to all matching filters.
  @param cmd    the command to send, for handling simplicity all commands must be alphanumeric only
  @param arg    the argument for the command
  @param res    a buffer with size res_size where the filter(s) can return a response.

  @returns >=0 on success otherwise an error code.
               AVERROR(ENOSYS) on unsupported commands
*/
func AVFilterGraphSendCommand(graph *AVFilterGraph, target *CStr, cmd *CStr, arg *CStr, res *CStr, resLen int, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_queue_command ---

// AVFilterGraphQueueCommand wraps avfilter_graph_queue_command.
/*
  Queue a command for one or more filter instances.

  @param graph  the filter graph
  @param target the filter(s) to which the command should be sent
                "all" sends to all filters
                otherwise it can be a filter or filter instance name
                which will send the command to all matching filters.
  @param cmd    the command to sent, for handling simplicity all commands must be alphanumeric only
  @param arg    the argument for the command
  @param ts     time at which the command should be sent to the filter

  @note As this executes commands after this function returns, no return code
        from the filter is provided, also AVFILTER_CMD_FLAG_ONE is not supported.
*/
func AVFilterGraphQueueCommand(graph *AVFilterGraph, target *CStr, cmd *CStr, arg *CStr, flags int, ts float64) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avfilter_graph_dump ---

// AVFilterGraphDump wraps avfilter_graph_dump.
/*
  Dump a graph into a human-readable string representation.

  @param graph    the graph to dump
  @param options  formatting options; currently ignored
  @return  a string, or NULL in case of memory allocation failure;
           the string must be freed using av_free
*/
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
/*
  Request a frame on the oldest sink link.

  If the request returns AVERROR_EOF, try the next.

  Note that this function is not meant to be the sole scheduling mechanism
  of a filtergraph, only a convenience function to help drain a filtergraph
  in a balanced way under normal circumstances.

  Also note that AVERROR_EOF does not mean that frames did not arrive on
  some of the sinks during the process.
  When there are multiple sink links, in case the requested link
  returns an EOF, this may cause a filter to flush pending frames
  which are sent to another sink link, although unrequested.

  @return  the return value of ff_request_frame(),
           or AVERROR_EOF if all links returned AVERROR_EOF
*/
func AVFilterGraphRequestOldest(graph *AVFilterGraph) (int, error) {
	var tmpgraph *C.AVFilterGraph
	if graph != nil {
		tmpgraph = graph.ptr
	}
	ret := C.avfilter_graph_request_oldest(tmpgraph)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffersink_get_frame_flags ---

// AVBuffersinkGetFrameFlags wraps av_buffersink_get_frame_flags.
/*
  Get a frame with filtered data from sink and put it in frame.

  @param ctx    pointer to a buffersink or abuffersink filter context.
  @param frame  pointer to an allocated frame that will be filled with data.
                The data must be freed using av_frame_unref() / av_frame_free()
  @param flags  a combination of AV_BUFFERSINK_FLAG_* flags

  @return  >= 0 in for success, a negative AVERROR code for failure.
*/
func AVBuffersinkGetFrameFlags(ctx *AVFilterContext, frame *AVFrame, flags int) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersink_get_frame_flags(tmpctx, tmpframe, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffersink_set_frame_size ---

// AVBuffersinkSetFrameSize wraps av_buffersink_set_frame_size.
/*
  Set the frame size for an audio buffer sink.

  All calls to av_buffersink_get_buffer_ref will return a buffer with
  exactly the specified number of samples, or AVERROR(EAGAIN) if there is
  not enough. The last buffer at EOF will be padded with 0.
*/
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
func AVBuffersinkGetFormat(ctx *AVFilterContext) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_format(tmpctx)
	return int(ret), WrapErr(int(ret))
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
func AVBuffersinkGetW(ctx *AVFilterContext) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_w(tmpctx)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffersink_get_h ---

// AVBuffersinkGetH wraps av_buffersink_get_h.
func AVBuffersinkGetH(ctx *AVFilterContext) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_h(tmpctx)
	return int(ret), WrapErr(int(ret))
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
func AVBuffersinkGetChannels(ctx *AVFilterContext) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_channels(tmpctx)
	return int(ret), WrapErr(int(ret))
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
func AVBuffersinkGetChLayout(ctx *AVFilterContext, chLayout *AVChannelLayout) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpchLayout *C.AVChannelLayout
	if chLayout != nil {
		tmpchLayout = chLayout.ptr
	}
	ret := C.av_buffersink_get_ch_layout(tmpctx, tmpchLayout)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffersink_get_sample_rate ---

// AVBuffersinkGetSampleRate wraps av_buffersink_get_sample_rate.
func AVBuffersinkGetSampleRate(ctx *AVFilterContext) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersink_get_sample_rate(tmpctx)
	return int(ret), WrapErr(int(ret))
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
/*
  Get a frame with filtered data from sink and put it in frame.

  @param ctx pointer to a context of a buffersink or abuffersink AVFilter.
  @param frame pointer to an allocated frame that will be filled with data.
               The data must be freed using av_frame_unref() / av_frame_free()

  @return
          - >= 0 if a frame was successfully returned.
          - AVERROR(EAGAIN) if no frames are available at this point; more
            input frames must be added to the filtergraph to get more output.
          - AVERROR_EOF if there will be no more output frames on this sink.
          - A different negative AVERROR code in other failure cases.
*/
func AVBuffersinkGetFrame(ctx *AVFilterContext, frame *AVFrame) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersink_get_frame(tmpctx, tmpframe)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffersink_get_samples ---

// AVBuffersinkGetSamples wraps av_buffersink_get_samples.
/*
  Same as av_buffersink_get_frame(), but with the ability to specify the number
  of samples read. This function is less efficient than
  av_buffersink_get_frame(), because it copies the data around.

  @param ctx pointer to a context of the abuffersink AVFilter.
  @param frame pointer to an allocated frame that will be filled with data.
               The data must be freed using av_frame_unref() / av_frame_free()
               frame will contain exactly nb_samples audio samples, except at
               the end of stream, when it can contain less than nb_samples.

  @return The return codes have the same meaning as for
          av_buffersink_get_frame().

  @warning do not mix this function with av_buffersink_get_frame(). Use only one or
  the other with a single sink, not both.
*/
func AVBuffersinkGetSamples(ctx *AVFilterContext, frame *AVFrame, nbSamples int) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersink_get_samples(tmpctx, tmpframe, C.int(nbSamples))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffersrc_get_nb_failed_requests ---

// AVBuffersrcGetNbFailedRequests wraps av_buffersrc_get_nb_failed_requests.
/*
  Get the number of failed requests.

  A failed request is when the request_frame method is called while no
  frame is present in the buffer.
  The number is reset when a frame is added.
*/
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
/*
  Allocate a new AVBufferSrcParameters instance. It should be freed by the
  caller with av_free().
*/
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
/*
  Initialize the buffersrc or abuffersrc filter with the provided parameters.
  This function may be called multiple times, the later calls override the
  previous ones. Some of the parameters may also be set through AVOptions, then
  whatever method is used last takes precedence.

  @param ctx an instance of the buffersrc or abuffersrc filter
  @param param the stream parameters. The frames later passed to this filter
               must conform to those parameters. All the allocated fields in
               param remain owned by the caller, libavfilter will make internal
               copies or references when necessary.
  @return 0 on success, a negative AVERROR code on failure.
*/
func AVBuffersrcParametersSet(ctx *AVFilterContext, param *AVBufferSrcParameters) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpparam *C.AVBufferSrcParameters
	if param != nil {
		tmpparam = param.ptr
	}
	ret := C.av_buffersrc_parameters_set(tmpctx, tmpparam)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffersrc_write_frame ---

// AVBuffersrcWriteFrame wraps av_buffersrc_write_frame.
/*
  Add a frame to the buffer source.

  @param ctx   an instance of the buffersrc filter
  @param frame frame to be added. If the frame is reference counted, this
  function will make a new reference to it. Otherwise the frame data will be
  copied.

  @return 0 on success, a negative AVERROR on error

  This function is equivalent to av_buffersrc_add_frame_flags() with the
  AV_BUFFERSRC_FLAG_KEEP_REF flag.
*/
func AVBuffersrcWriteFrame(ctx *AVFilterContext, frame *AVFrame) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersrc_write_frame(tmpctx, tmpframe)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffersrc_add_frame ---

// AVBuffersrcAddFrame wraps av_buffersrc_add_frame.
/*
  Add a frame to the buffer source.

  @param ctx   an instance of the buffersrc filter
  @param frame frame to be added. If the frame is reference counted, this
  function will take ownership of the reference(s) and reset the frame.
  Otherwise the frame data will be copied. If this function returns an error,
  the input frame is not touched.

  @return 0 on success, a negative AVERROR on error.

  @note the difference between this function and av_buffersrc_write_frame() is
  that av_buffersrc_write_frame() creates a new reference to the input frame,
  while this function takes ownership of the reference passed to it.

  This function is equivalent to av_buffersrc_add_frame_flags() without the
  AV_BUFFERSRC_FLAG_KEEP_REF flag.
*/
func AVBuffersrcAddFrame(ctx *AVFilterContext, frame *AVFrame) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersrc_add_frame(tmpctx, tmpframe)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffersrc_add_frame_flags ---

// AVBuffersrcAddFrameFlags wraps av_buffersrc_add_frame_flags.
/*
  Add a frame to the buffer source.

  By default, if the frame is reference-counted, this function will take
  ownership of the reference(s) and reset the frame. This can be controlled
  using the flags.

  If this function returns an error, the input frame is not touched.

  @param buffer_src  pointer to a buffer source context
  @param frame       a frame, or NULL to mark EOF
  @param flags       a combination of AV_BUFFERSRC_FLAG_*
  @return            >= 0 in case of success, a negative AVERROR code
                     in case of failure
*/
func AVBuffersrcAddFrameFlags(bufferSrc *AVFilterContext, frame *AVFrame, flags int) (int, error) {
	var tmpbufferSrc *C.AVFilterContext
	if bufferSrc != nil {
		tmpbufferSrc = bufferSrc.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_buffersrc_add_frame_flags(tmpbufferSrc, tmpframe, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffersrc_close ---

// AVBuffersrcClose wraps av_buffersrc_close.
/*
  Close the buffer source after EOF.

  This is similar to passing NULL to av_buffersrc_add_frame_flags()
  except it takes the timestamp of the EOF, i.e. the timestamp of the end
  of the last frame.
*/
func AVBuffersrcClose(ctx *AVFilterContext, pts int64, flags uint) (int, error) {
	var tmpctx *C.AVFilterContext
	if ctx != nil {
		tmpctx = ctx.ptr
	}
	ret := C.av_buffersrc_close(tmpctx, C.int64_t(pts), C.uint(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_get_packet ---

// AVGetPacket wraps av_get_packet.
/*
  Allocate and read the payload of a packet and initialize its
  fields with default values.

  @param s    associated IO context
  @param pkt packet
  @param size desired payload size
  @return >0 (read size) if OK, AVERROR_xxx otherwise
*/
func AVGetPacket(s *AVIOContext, pkt *AVPacket, size int) (int, error) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_get_packet(tmps, tmppkt, C.int(size))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_append_packet ---

// AVAppendPacket wraps av_append_packet.
/*
  Read data and append it to the current content of the AVPacket.
  If pkt->size is 0 this is identical to av_get_packet.
  Note that this uses av_grow_packet and thus involves a realloc
  which is inefficient. Thus this function should only be used
  when there is no reasonable way to know (an upper bound of)
  the final size.

  @param s    associated IO context
  @param pkt packet
  @param size amount of data to read
  @return >0 (read size) if OK, AVERROR_xxx otherwise, previous data
          will not be lost even if an error occurs.
*/
func AVAppendPacket(s *AVIOContext, pkt *AVPacket, size int) (int, error) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_append_packet(tmps, tmppkt, C.int(size))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_disposition_from_string ---

// AVDispositionFromString wraps av_disposition_from_string.
/*
  @return The AV_DISPOSITION_* flag corresponding to disp or a negative error
          code if disp does not correspond to a known stream disposition.
*/
func AVDispositionFromString(disp *CStr) (int, error) {
	var tmpdisp *C.char
	if disp != nil {
		tmpdisp = disp.ptr
	}
	ret := C.av_disposition_from_string(tmpdisp)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_disposition_to_string ---

// AVDispositionToString wraps av_disposition_to_string.
/*
  @param disposition a combination of AV_DISPOSITION_* values
  @return The string description corresponding to the lowest set bit in
          disposition. NULL when the lowest set bit does not correspond
          to a known disposition or when disposition is 0.
*/
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
/*
  Returns the pts of the last muxed packet + its duration

  the retuned value is undefined when used with a demuxer.
*/
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
//
//	// Chromium: We use the internal field first_dts vvv
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
/*
  This function will cause global side data to be injected in the next packet
  of each stream as well as after any subsequent seek.
*/
func AVFormatInjectGlobalSideData(s *AVFormatContext) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	C.av_format_inject_global_side_data(tmps)
}

// --- Function av_fmt_ctx_get_duration_estimation_method ---

// AVFmtCtxGetDurationEstimationMethod wraps av_fmt_ctx_get_duration_estimation_method.
/*
  Returns the method used to set ctx->duration.

  @return AVFMT_DURATION_FROM_PTS, AVFMT_DURATION_FROM_STREAM, or AVFMT_DURATION_FROM_BITRATE.
*/
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
//
//	Return the LIBAVFORMAT_VERSION_INT constant.
func AVFormatVersion() uint {
	ret := C.avformat_version()
	return uint(ret)
}

// --- Function avformat_configuration ---

// AVFormatConfiguration wraps avformat_configuration.
//
//	Return the libavformat build-time configuration.
func AVFormatConfiguration() *CStr {
	ret := C.avformat_configuration()
	return wrapCStr(ret)
}

// --- Function avformat_license ---

// AVFormatLicense wraps avformat_license.
//
//	Return the libavformat license.
func AVFormatLicense() *CStr {
	ret := C.avformat_license()
	return wrapCStr(ret)
}

// --- Function avformat_network_init ---

// AVFormatNetworkInit wraps avformat_network_init.
/*
  Do global initialization of network libraries. This is optional,
  and not recommended anymore.

  This functions only exists to work around thread-safety issues
  with older GnuTLS or OpenSSL libraries. If libavformat is linked
  to newer versions of those libraries, or if you do not use them,
  calling this function is unnecessary. Otherwise, you need to call
  this function before any other threads using them are started.

  This function will be deprecated once support for older GnuTLS and
  OpenSSL libraries is removed, and this function has no purpose
  anymore.
*/
func AVFormatNetworkInit() (int, error) {
	ret := C.avformat_network_init()
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_network_deinit ---

// AVFormatNetworkDeinit wraps avformat_network_deinit.
/*
  Undo the initialization done by avformat_network_init. Call it only
  once for each time you called avformat_network_init.
*/
func AVFormatNetworkDeinit() (int, error) {
	ret := C.avformat_network_deinit()
	return int(ret), WrapErr(int(ret))
}

// --- Function av_muxer_iterate ---

// av_muxer_iterate skipped due to opaque

// --- Function av_demuxer_iterate ---

// av_demuxer_iterate skipped due to opaque

// --- Function avformat_alloc_context ---

// AVFormatAllocContext wraps avformat_alloc_context.
/*
  Allocate an AVFormatContext.
  avformat_free_context() can be used to free the context and everything
  allocated by the framework within it.
*/
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
/*
  Free an AVFormatContext and all its streams.
  @param s context to free
*/
func AVFormatFreeContext(s *AVFormatContext) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	C.avformat_free_context(tmps)
}

// --- Function avformat_get_class ---

// AVFormatGetClass wraps avformat_get_class.
/*
  Get the AVClass for AVFormatContext. It can be used in combination with
  AV_OPT_SEARCH_FAKE_OBJ for examining options.

  @see av_opt_find().
*/
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
/*
  Get the AVClass for AVStream. It can be used in combination with
  AV_OPT_SEARCH_FAKE_OBJ for examining options.

  @see av_opt_find().
*/
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
/*
  Add a new stream to a media file.

  When demuxing, it is called by the demuxer in read_header(). If the
  flag AVFMTCTX_NOHEADER is set in s.ctx_flags, then it may also
  be called in read_packet().

  When muxing, should be called by the user before avformat_write_header().

  User is required to call avformat_free_context() to clean up the allocation
  by avformat_new_stream().

  @param s media file handle
  @param c unused, does nothing

  @return newly created stream or NULL on error.
*/
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
/*
  Wrap an existing array as stream side data.

  @param st   stream
  @param type side information type
  @param data the side data array. It must be allocated with the av_malloc()
              family of functions. The ownership of the data is transferred to
              st.
  @param size side information size

  @return zero on success, a negative AVERROR code on failure. On failure,
          the stream is unchanged and the data remains owned by the caller.
*/
func AVStreamAddSideData(st *AVStream, _type AVPacketSideDataType, data unsafe.Pointer, size uint64) (int, error) {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.av_stream_add_side_data(tmpst, C.enum_AVPacketSideDataType(_type), (*C.uint8_t)(data), C.size_t(size))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_stream_new_side_data ---

// AVStreamNewSideData wraps av_stream_new_side_data.
/*
  Allocate new information from stream.

  @param stream stream
  @param type   desired side information type
  @param size   side information size

  @return pointer to fresh allocated data or NULL otherwise
*/
func AVStreamNewSideData(stream *AVStream, _type AVPacketSideDataType, size uint64) unsafe.Pointer {
	var tmpstream *C.AVStream
	if stream != nil {
		tmpstream = stream.ptr
	}
	ret := C.av_stream_new_side_data(tmpstream, C.enum_AVPacketSideDataType(_type), C.size_t(size))
	return unsafe.Pointer(ret)
}

// --- Function av_stream_get_side_data ---

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
/*
  Allocate an AVFormatContext for an output format.
  avformat_free_context() can be used to free the context and
  everything allocated by the framework within it.

  @param ctx           pointee is set to the created format context,
                       or to NULL in case of failure
  @param oformat       format to use for allocating the context, if NULL
                       format_name and filename are used instead
  @param format_name   the name of output format to use for allocating the
                       context, if NULL filename is used instead
  @param filename      the name of the filename to use for allocating the
                       context, may be NULL

  @return  >= 0 in case of success, a negative AVERROR code in case of
           failure
*/
func AVFormatAllocOutputContext2(ctx **AVFormatContext, oformat *AVOutputFormat, formatName *CStr, filename *CStr) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_find_input_format ---

// AVFindInputFormat wraps av_find_input_format.
//
//	Find AVInputFormat based on the short name of the input format.
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
/*
  Guess the file format.

  @param pd        data to be probed
  @param is_opened Whether the file is already opened; determines whether
                   demuxers with or without AVFMT_NOFILE are probed.
*/
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

// av_probe_input_format2 skipped due to scoreMax

// --- Function av_probe_input_format3 ---

// av_probe_input_format3 skipped due to scoreRet

// --- Function av_probe_input_buffer2 ---

// AVProbeInputBuffer2 wraps av_probe_input_buffer2.
/*
  Probe a bytestream to determine the input format. Each time a probe returns
  with a score that is too low, the probe buffer size is increased and another
  attempt is made. When the maximum probe size is reached, the input format
  with the highest score is returned.

  @param pb             the bytestream to probe
  @param fmt            the input format is put here
  @param url            the url of the stream
  @param logctx         the log context
  @param offset         the offset within the bytestream to probe from
  @param max_probe_size the maximum probe buffer size (zero for default)

  @return the score in case of success, a negative value corresponding to an
          the maximal score is AVPROBE_SCORE_MAX
          AVERROR code otherwise
*/
func AVProbeInputBuffer2(pb *AVIOContext, fmt **AVInputFormat, url *CStr, logctx unsafe.Pointer, offset uint, maxProbeSize uint) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_probe_input_buffer ---

// AVProbeInputBuffer wraps av_probe_input_buffer.
//
//	Like av_probe_input_buffer2() but returns 0 on success
func AVProbeInputBuffer(pb *AVIOContext, fmt **AVInputFormat, url *CStr, logctx unsafe.Pointer, offset uint, maxProbeSize uint) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_open_input ---

// AVFormatOpenInput wraps avformat_open_input.
/*
  Open an input stream and read the header. The codecs are not opened.
  The stream must be closed with avformat_close_input().

  @param ps       Pointer to user-supplied AVFormatContext (allocated by
                  avformat_alloc_context). May be a pointer to NULL, in
                  which case an AVFormatContext is allocated by this
                  function and written into ps.
                  Note that a user-supplied AVFormatContext will be freed
                  on failure.
  @param url      URL of the stream to open.
  @param fmt      If non-NULL, this parameter forces a specific input format.
                  Otherwise the format is autodetected.
  @param options  A dictionary filled with AVFormatContext and demuxer-private
                  options.
                  On return this parameter will be destroyed and replaced with
                  a dict containing options that were not found. May be NULL.

  @return 0 on success, a negative AVERROR on failure.

  @note If you want to use custom IO, preallocate the format context and set its pb field.
*/
func AVFormatOpenInput(ps **AVFormatContext, url *CStr, fmt *AVInputFormat, options **AVDictionary) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_find_stream_info ---

// AVFormatFindStreamInfo wraps avformat_find_stream_info.
/*
  Read packets of a media file to get stream information. This
  is useful for file formats with no headers such as MPEG. This
  function also computes the real framerate in case of MPEG-2 repeat
  frame mode.
  The logical file position is not changed by this function;
  examined packets may be buffered for later processing.

  @param ic media file handle
  @param options  If non-NULL, an ic.nb_streams long array of pointers to
                  dictionaries, where i-th member contains options for
                  codec corresponding to i-th stream.
                  On return each dictionary will be filled with options that were not found.
  @return >=0 if OK, AVERROR_xxx on error

  @note this function isn't guaranteed to open all the codecs, so
        options being non-empty at return is a perfectly normal behavior.

  @todo Let the user decide somehow what information is needed so that
        we do not waste time getting stuff the user does not need.
*/
func AVFormatFindStreamInfo(ic *AVFormatContext, options **AVDictionary) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_find_program_from_stream ---

// AVFindProgramFromStream wraps av_find_program_from_stream.
/*
  Find the programs which belong to a given stream.

  @param ic    media file handle
  @param last  the last found program, the search will start after this
               program, or from the beginning if it is NULL
  @param s     stream index

  @return the next program which belongs to s, NULL if no program is found or
          the last program is not among the programs of ic.
*/
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
/*
  Find the "best" stream in the file.
  The best stream is determined according to various heuristics as the most
  likely to be what the user expects.
  If the decoder parameter is non-NULL, av_find_best_stream will find the
  default decoder for the stream's codec; streams for which no decoder can
  be found are ignored.

  @param ic                media file handle
  @param type              stream type: video, audio, subtitles, etc.
  @param wanted_stream_nb  user-requested stream number,
                           or -1 for automatic selection
  @param related_stream    try to find a stream related (eg. in the same
                           program) to this one, or -1 if none
  @param decoder_ret       if non-NULL, returns the decoder for the
                           selected stream
  @param flags             flags; none are currently defined

  @return  the non-negative stream number in case of success,
           AVERROR_STREAM_NOT_FOUND if no stream with the requested type
           could be found,
           AVERROR_DECODER_NOT_FOUND if streams were found but no decoder

  @note  If av_find_best_stream returns successfully and decoder_ret is not
         NULL, then *decoder_ret is guaranteed to be set to a valid AVCodec.
*/
func AVFindBestStream(ic *AVFormatContext, _type AVMediaType, wantedStreamNb int, relatedStream int, decoderRet **AVCodec, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_read_frame ---

// AVReadFrame wraps av_read_frame.
/*
  Return the next frame of a stream.
  This function returns what is stored in the file, and does not validate
  that what is there are valid frames for the decoder. It will split what is
  stored in the file into frames and return one for each call. It will not
  omit invalid data between valid frames so as to give the decoder the maximum
  information possible for decoding.

  On success, the returned packet is reference-counted (pkt->buf is set) and
  valid indefinitely. The packet must be freed with av_packet_unref() when
  it is no longer needed. For video, the packet contains exactly one frame.
  For audio, it contains an integer number of frames if each frame has
  a known fixed size (e.g. PCM or ADPCM data). If the audio frames have
  a variable size (e.g. MPEG audio), then it contains one frame.

  pkt->pts, pkt->dts and pkt->duration are always set to correct
  values in AVStream.time_base units (and guessed if the format cannot
  provide them). pkt->pts can be AV_NOPTS_VALUE if the video format
  has B-frames, so it is better to rely on pkt->dts if you do not
  decompress the payload.

  @return 0 if OK, < 0 on error or end of file. On error, pkt will be blank
          (as if it came from av_packet_alloc()).

  @note pkt will be initialized, so it may be uninitialized, but it must not
        contain data that needs to be freed.
*/
func AVReadFrame(s *AVFormatContext, pkt *AVPacket) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_read_frame(tmps, tmppkt)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_seek_frame ---

// AVSeekFrame wraps av_seek_frame.
/*
  Seek to the keyframe at timestamp.
  'timestamp' in 'stream_index'.

  @param s            media file handle
  @param stream_index If stream_index is (-1), a default stream is selected,
                      and timestamp is automatically converted from
                      AV_TIME_BASE units to the stream specific time_base.
  @param timestamp    Timestamp in AVStream.time_base units or, if no stream
                      is specified, in AV_TIME_BASE units.
  @param flags        flags which select direction and seeking mode

  @return >= 0 on success
*/
func AVSeekFrame(s *AVFormatContext, streamIndex int, timestamp int64, flags int) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_seek_frame(tmps, C.int(streamIndex), C.int64_t(timestamp), C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_seek_file ---

// AVFormatSeekFile wraps avformat_seek_file.
/*
  Seek to timestamp ts.
  Seeking will be done so that the point from which all active streams
  can be presented successfully will be closest to ts and within min/max_ts.
  Active streams are all streams that have AVStream.discard < AVDISCARD_ALL.

  If flags contain AVSEEK_FLAG_BYTE, then all timestamps are in bytes and
  are the file position (this may not be supported by all demuxers).
  If flags contain AVSEEK_FLAG_FRAME, then all timestamps are in frames
  in the stream with stream_index (this may not be supported by all demuxers).
  Otherwise all timestamps are in units of the stream selected by stream_index
  or if stream_index is -1, in AV_TIME_BASE units.
  If flags contain AVSEEK_FLAG_ANY, then non-keyframes are treated as
  keyframes (this may not be supported by all demuxers).
  If flags contain AVSEEK_FLAG_BACKWARD, it is ignored.

  @param s            media file handle
  @param stream_index index of the stream which is used as time base reference
  @param min_ts       smallest acceptable timestamp
  @param ts           target timestamp
  @param max_ts       largest acceptable timestamp
  @param flags        flags
  @return >=0 on success, error code otherwise

  @note This is part of the new seek API which is still under construction.
*/
func AVFormatSeekFile(s *AVFormatContext, streamIndex int, minTs int64, ts int64, maxTs int64, flags int) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avformat_seek_file(tmps, C.int(streamIndex), C.int64_t(minTs), C.int64_t(ts), C.int64_t(maxTs), C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_flush ---

// AVFormatFlush wraps avformat_flush.
/*
  Discard all internally buffered data. This can be useful when dealing with
  discontinuities in the byte stream. Generally works only with formats that
  can resync. This includes headerless formats like MPEG-TS/TS but should also
  work with NUT, Ogg and in a limited way AVI for example.

  The set of streams, the detected duration, stream parameters and codecs do
  not change when calling this function. If you want a complete reset, it's
  better to open a new AVFormatContext.

  This does not flush the AVIOContext (s->pb). If necessary, call
  avio_flush(s->pb) before calling this function.

  @param s media file handle
  @return >=0 on success, error code otherwise
*/
func AVFormatFlush(s *AVFormatContext) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avformat_flush(tmps)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_read_play ---

// AVReadPlay wraps av_read_play.
/*
  Start playing a network-based stream (e.g. RTSP stream) at the
  current position.
*/
func AVReadPlay(s *AVFormatContext) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_read_play(tmps)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_read_pause ---

// AVReadPause wraps av_read_pause.
/*
  Pause a network-based stream (e.g. RTSP stream).

  Use av_read_play() to resume it.
*/
func AVReadPause(s *AVFormatContext) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_read_pause(tmps)
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_close_input ---

// AVFormatCloseInput wraps avformat_close_input.
/*
  Close an opened input AVFormatContext. Free it and all its contents
  and set *s to NULL.
*/
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
/*
  Allocate the stream private data and write the stream header to
  an output media file.

  @param s        Media file handle, must be allocated with
                  avformat_alloc_context().
                  Its \ref AVFormatContext.oformat "oformat" field must be set
                  to the desired output format;
                  Its \ref AVFormatContext.pb "pb" field must be set to an
                  already opened ::AVIOContext.
  @param options  An ::AVDictionary filled with AVFormatContext and
                  muxer-private options.
                  On return this parameter will be destroyed and replaced with
                  a dict containing options that were not found. May be NULL.

  @retval AVSTREAM_INIT_IN_WRITE_HEADER On success, if the codec had not already been
                                        fully initialized in avformat_init_output().
  @retval AVSTREAM_INIT_IN_INIT_OUTPUT  On success, if the codec had already been fully
                                        initialized in avformat_init_output().
  @retval AVERROR                       A negative AVERROR on failure.

  @see av_opt_find, av_dict_set, avio_open, av_oformat_next, avformat_init_output.
*/
func AVFormatWriteHeader(s *AVFormatContext, options **AVDictionary) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_init_output ---

// AVFormatInitOutput wraps avformat_init_output.
/*
  Allocate the stream private data and initialize the codec, but do not write the header.
  May optionally be used before avformat_write_header() to initialize stream parameters
  before actually writing the header.
  If using this function, do not pass the same options to avformat_write_header().

  @param s        Media file handle, must be allocated with
                  avformat_alloc_context().
                  Its \ref AVFormatContext.oformat "oformat" field must be set
                  to the desired output format;
                  Its \ref AVFormatContext.pb "pb" field must be set to an
                  already opened ::AVIOContext.
  @param options  An ::AVDictionary filled with AVFormatContext and
                  muxer-private options.
                  On return this parameter will be destroyed and replaced with
                  a dict containing options that were not found. May be NULL.

  @retval AVSTREAM_INIT_IN_WRITE_HEADER On success, if the codec requires
                                        avformat_write_header to fully initialize.
  @retval AVSTREAM_INIT_IN_INIT_OUTPUT  On success, if the codec has been fully
                                        initialized.
  @retval AVERROR                       Anegative AVERROR on failure.

  @see av_opt_find, av_dict_set, avio_open, av_oformat_next, avformat_write_header.
*/
func AVFormatInitOutput(s *AVFormatContext, options **AVDictionary) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_write_frame ---

// AVWriteFrame wraps av_write_frame.
/*
  Write a packet to an output media file.

  This function passes the packet directly to the muxer, without any buffering
  or reordering. The caller is responsible for correctly interleaving the
  packets if the format requires it. Callers that want libavformat to handle
  the interleaving should call av_interleaved_write_frame() instead of this
  function.

  @param s media file handle
  @param pkt The packet containing the data to be written. Note that unlike
             av_interleaved_write_frame(), this function does not take
             ownership of the packet passed to it (though some muxers may make
             an internal reference to the input packet).
             <br>
             This parameter can be NULL (at any time, not just at the end), in
             order to immediately flush data buffered within the muxer, for
             muxers that buffer up data internally before writing it to the
             output.
             <br>
             Packet's @ref AVPacket.stream_index "stream_index" field must be
             set to the index of the corresponding stream in @ref
             AVFormatContext.streams "s->streams".
             <br>
             The timestamps (@ref AVPacket.pts "pts", @ref AVPacket.dts "dts")
             must be set to correct values in the stream's timebase (unless the
             output format is flagged with the AVFMT_NOTIMESTAMPS flag, then
             they can be set to AV_NOPTS_VALUE).
             The dts for subsequent packets passed to this function must be strictly
             increasing when compared in their respective timebases (unless the
             output format is flagged with the AVFMT_TS_NONSTRICT, then they
             merely have to be nondecreasing).  @ref AVPacket.duration
             "duration") should also be set if known.
  @return < 0 on error, = 0 if OK, 1 if flushed and there is no more data to flush

  @see av_interleaved_write_frame()
*/
func AVWriteFrame(s *AVFormatContext, pkt *AVPacket) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_write_frame(tmps, tmppkt)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_interleaved_write_frame ---

// AVInterleavedWriteFrame wraps av_interleaved_write_frame.
/*
  Write a packet to an output media file ensuring correct interleaving.

  This function will buffer the packets internally as needed to make sure the
  packets in the output file are properly interleaved, usually ordered by
  increasing dts. Callers doing their own interleaving should call
  av_write_frame() instead of this function.

  Using this function instead of av_write_frame() can give muxers advance
  knowledge of future packets, improving e.g. the behaviour of the mp4
  muxer for VFR content in fragmenting mode.

  @param s media file handle
  @param pkt The packet containing the data to be written.
             <br>
             If the packet is reference-counted, this function will take
             ownership of this reference and unreference it later when it sees
             fit. If the packet is not reference-counted, libavformat will
             make a copy.
             The returned packet will be blank (as if returned from
             av_packet_alloc()), even on error.
             <br>
             This parameter can be NULL (at any time, not just at the end), to
             flush the interleaving queues.
             <br>
             Packet's @ref AVPacket.stream_index "stream_index" field must be
             set to the index of the corresponding stream in @ref
             AVFormatContext.streams "s->streams".
             <br>
             The timestamps (@ref AVPacket.pts "pts", @ref AVPacket.dts "dts")
             must be set to correct values in the stream's timebase (unless the
             output format is flagged with the AVFMT_NOTIMESTAMPS flag, then
             they can be set to AV_NOPTS_VALUE).
             The dts for subsequent packets in one stream must be strictly
             increasing (unless the output format is flagged with the
             AVFMT_TS_NONSTRICT, then they merely have to be nondecreasing).
             @ref AVPacket.duration "duration" should also be set if known.

  @return 0 on success, a negative AVERROR on error.

  @see av_write_frame(), AVFormatContext.max_interleave_delta
*/
func AVInterleavedWriteFrame(s *AVFormatContext, pkt *AVPacket) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmppkt *C.AVPacket
	if pkt != nil {
		tmppkt = pkt.ptr
	}
	ret := C.av_interleaved_write_frame(tmps, tmppkt)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_write_uncoded_frame ---

// AVWriteUncodedFrame wraps av_write_uncoded_frame.
/*
  Write an uncoded frame to an output media file.

  The frame must be correctly interleaved according to the container
  specification; if not, av_interleaved_write_uncoded_frame() must be used.

  See av_interleaved_write_uncoded_frame() for details.
*/
func AVWriteUncodedFrame(s *AVFormatContext, streamIndex int, frame *AVFrame) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_write_uncoded_frame(tmps, C.int(streamIndex), tmpframe)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_interleaved_write_uncoded_frame ---

// AVInterleavedWriteUncodedFrame wraps av_interleaved_write_uncoded_frame.
/*
  Write an uncoded frame to an output media file.

  If the muxer supports it, this function makes it possible to write an AVFrame
  structure directly, without encoding it into a packet.
  It is mostly useful for devices and similar special muxers that use raw
  video or PCM data and will not serialize it into a byte stream.

  To test whether it is possible to use it with a given muxer and stream,
  use av_write_uncoded_frame_query().

  The caller gives up ownership of the frame and must not access it
  afterwards.

  @return  >=0 for success, a negative code on error
*/
func AVInterleavedWriteUncodedFrame(s *AVFormatContext, streamIndex int, frame *AVFrame) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_interleaved_write_uncoded_frame(tmps, C.int(streamIndex), tmpframe)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_write_uncoded_frame_query ---

// AVWriteUncodedFrameQuery wraps av_write_uncoded_frame_query.
/*
  Test whether a muxer supports uncoded frame.

  @return  >=0 if an uncoded frame can be written to that muxer and stream,
           <0 if not
*/
func AVWriteUncodedFrameQuery(s *AVFormatContext, streamIndex int) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_write_uncoded_frame_query(tmps, C.int(streamIndex))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_write_trailer ---

// AVWriteTrailer wraps av_write_trailer.
/*
  Write the stream trailer to an output media file and free the
  file private data.

  May only be called after a successful call to avformat_write_header.

  @param s media file handle
  @return 0 if OK, AVERROR_xxx on error
*/
func AVWriteTrailer(s *AVFormatContext) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_write_trailer(tmps)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_guess_format ---

// AVGuessFormat wraps av_guess_format.
/*
  Return the output format in the list of registered output formats
  which best matches the provided parameters, or return NULL if
  there is no match.

  @param short_name if non-NULL checks if short_name matches with the
                    names of the registered formats
  @param filename   if non-NULL checks if filename terminates with the
                    extensions of the registered formats
  @param mime_type  if non-NULL checks if mime_type matches with the
                    MIME type of the registered formats
*/
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
//
//	Guess the codec ID based upon muxer and filename.
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

// av_get_output_timestamp skipped due to dts

// --- Function av_hex_dump ---

// av_hex_dump skipped due to f.

// --- Function av_hex_dump_log ---

// AVHexDumpLog wraps av_hex_dump_log.
/*
  Send a nice hexadecimal dump of a buffer to the log.

  @param avcl A pointer to an arbitrary struct of which the first field is a
  pointer to an AVClass struct.
  @param level The importance level of the message, lower values signifying
  higher importance.
  @param buf buffer
  @param size buffer size

  @see av_hex_dump, av_pkt_dump2, av_pkt_dump_log2
*/
func AVHexDumpLog(avcl unsafe.Pointer, level int, buf unsafe.Pointer, size int) {
	C.av_hex_dump_log(avcl, C.int(level), (*C.uint8_t)(buf), C.int(size))
}

// --- Function av_pkt_dump2 ---

// av_pkt_dump2 skipped due to f.

// --- Function av_pkt_dump_log2 ---

// AVPktDumpLog2 wraps av_pkt_dump_log2.
/*
  Send a nice dump of a packet to the log.

  @param avcl A pointer to an arbitrary struct of which the first field is a
  pointer to an AVClass struct.
  @param level The importance level of the message, lower values signifying
  higher importance.
  @param pkt packet to dump
  @param dump_payload True if the payload must be displayed, too.
  @param st AVStream that the packet belongs to
*/
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
/*
  Get the AVCodecID for the given codec tag tag.
  If no codec id is found returns AV_CODEC_ID_NONE.

  @param tags list of supported codec_id-codec_tag pairs, as stored
  in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
  @param tag  codec tag to match to a codec ID
*/
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
/*
  Get the codec tag for the given codec id id.
  If no codec tag is found returns 0.

  @param tags list of supported codec_id-codec_tag pairs, as stored
  in AVInputFormat.codec_tag and AVOutputFormat.codec_tag
  @param id   codec ID to match to a codec tag
*/
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

// av_codec_get_tag2 skipped due to tag

// --- Function av_find_default_stream_index ---

// AVFindDefaultStreamIndex wraps av_find_default_stream_index.
func AVFindDefaultStreamIndex(s *AVFormatContext) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.av_find_default_stream_index(tmps)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_index_search_timestamp ---

// AVIndexSearchTimestamp wraps av_index_search_timestamp.
/*
  Get the index for a specific timestamp.

  @param st        stream that the timestamp belongs to
  @param timestamp timestamp to retrieve the index for
  @param flags if AVSEEK_FLAG_BACKWARD then the returned index will correspond
                  to the timestamp which is <= the requested one, if backward
                  is 0, then it will be >=
               if AVSEEK_FLAG_ANY seek to any frame, only keyframes otherwise
  @return < 0 if no such timestamp could be found
*/
func AVIndexSearchTimestamp(st *AVStream, timestamp int64, flags int) (int, error) {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.av_index_search_timestamp(tmpst, C.int64_t(timestamp), C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_index_get_entries_count ---

// AVFormatIndexGetEntriesCount wraps avformat_index_get_entries_count.
/*
  Get the index entry count for the given AVStream.

  @param st stream
  @return the number of index entries in the stream
*/
func AVFormatIndexGetEntriesCount(st *AVStream) (int, error) {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.avformat_index_get_entries_count(tmpst)
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_index_get_entry ---

// AVFormatIndexGetEntry wraps avformat_index_get_entry.
/*
  Get the AVIndexEntry corresponding to the given index.

  @param st          Stream containing the requested AVIndexEntry.
  @param idx         The desired index.
  @return A pointer to the requested AVIndexEntry if it exists, NULL otherwise.

  @note The pointer returned by this function is only guaranteed to be valid
        until any function that takes the stream or the parent AVFormatContext
        as input argument is called.
*/
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
/*
  Get the AVIndexEntry corresponding to the given timestamp.

  @param st          Stream containing the requested AVIndexEntry.
  @param wanted_timestamp   Timestamp to retrieve the index entry for.
  @param flags       If AVSEEK_FLAG_BACKWARD then the returned entry will correspond
                     to the timestamp which is <= the requested one, if backward
                     is 0, then it will be >=
                     if AVSEEK_FLAG_ANY seek to any frame, only keyframes otherwise.
  @return A pointer to the requested AVIndexEntry if it exists, NULL otherwise.

  @note The pointer returned by this function is only guaranteed to be valid
        until any function that takes the stream or the parent AVFormatContext
        as input argument is called.
*/
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
/*
  Add an index entry into a sorted list. Update the entry if the list
  already contains it.

  @param timestamp timestamp in the time base of the given stream
*/
func AVAddIndexEntry(st *AVStream, pos int64, timestamp int64, size int, distance int, flags int) (int, error) {
	var tmpst *C.AVStream
	if st != nil {
		tmpst = st.ptr
	}
	ret := C.av_add_index_entry(tmpst, C.int64_t(pos), C.int64_t(timestamp), C.int(size), C.int(distance), C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_url_split ---

// av_url_split skipped due to portPtr

// --- Function av_dump_format ---

// AVDumpFormat wraps av_dump_format.
/*
  Print detailed information about the input or output format, such as
  duration, bitrate, streams, container, programs, metadata, side data,
  codec and time base.

  @param ic        the context to analyze
  @param index     index of the stream to dump information about
  @param url       the URL to print, such as source or destination file
  @param is_output Select whether the specified context is an input(0) or output(1)
*/
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
/*
  Return in 'buf' the path with '%d' replaced by a number.

  Also handles the '%0nd' format where 'n' is the total number
  of digits and '%%'.

  @param buf destination buffer
  @param buf_size destination buffer size
  @param path numbered sequence string
  @param number frame number
  @param flags AV_FRAME_FILENAME_FLAGS_*
  @return 0 if OK, -1 on format error
*/
func AVGetFrameFilename2(buf *CStr, bufSize int, path *CStr, number int, flags int) (int, error) {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	var tmppath *C.char
	if path != nil {
		tmppath = path.ptr
	}
	ret := C.av_get_frame_filename2(tmpbuf, C.int(bufSize), tmppath, C.int(number), C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_get_frame_filename ---

// AVGetFrameFilename wraps av_get_frame_filename.
func AVGetFrameFilename(buf *CStr, bufSize int, path *CStr, number int) (int, error) {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	var tmppath *C.char
	if path != nil {
		tmppath = path.ptr
	}
	ret := C.av_get_frame_filename(tmpbuf, C.int(bufSize), tmppath, C.int(number))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_filename_number_test ---

// AVFilenameNumberTest wraps av_filename_number_test.
/*
  Check whether filename actually is a numbered sequence generator.

  @param filename possible numbered sequence string
  @return 1 if a valid numbered sequence string, 0 otherwise
*/
func AVFilenameNumberTest(filename *CStr) (int, error) {
	var tmpfilename *C.char
	if filename != nil {
		tmpfilename = filename.ptr
	}
	ret := C.av_filename_number_test(tmpfilename)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_sdp_create ---

// av_sdp_create skipped due to ac

// --- Function av_match_ext ---

// AVMatchExt wraps av_match_ext.
/*
  Return a positive value if the given filename has one of the given
  extensions, 0 otherwise.

  @param filename   file name to check against the given extensions
  @param extensions a comma-separated list of filename extensions
*/
func AVMatchExt(filename *CStr, extensions *CStr) (int, error) {
	var tmpfilename *C.char
	if filename != nil {
		tmpfilename = filename.ptr
	}
	var tmpextensions *C.char
	if extensions != nil {
		tmpextensions = extensions.ptr
	}
	ret := C.av_match_ext(tmpfilename, tmpextensions)
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_query_codec ---

// AVFormatQueryCodec wraps avformat_query_codec.
/*
  Test if the given container can store a codec.

  @param ofmt           container to check for compatibility
  @param codec_id       codec to potentially store in container
  @param std_compliance standards compliance level, one of FF_COMPLIANCE_*

  @return 1 if codec with ID codec_id can be stored in ofmt, 0 if it cannot.
          A negative number if this information is not available.
*/
func AVFormatQueryCodec(ofmt *AVOutputFormat, codecId AVCodecID, stdCompliance int) (int, error) {
	var tmpofmt *C.AVOutputFormat
	if ofmt != nil {
		tmpofmt = ofmt.ptr
	}
	ret := C.avformat_query_codec(tmpofmt, C.enum_AVCodecID(codecId), C.int(stdCompliance))
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_get_riff_video_tags ---

// AVFormatGetRiffVideoTags wraps avformat_get_riff_video_tags.
//
//	@return the table mapping RIFF FourCCs for video to libavcodec AVCodecID.
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
//
//	@return the table mapping RIFF FourCCs for audio to AVCodecID.
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
//
//	@return the table mapping MOV FourCCs for video to libavcodec AVCodecID.
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
//
//	@return the table mapping MOV FourCCs for audio to AVCodecID.
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
/*
  Guess the sample aspect ratio of a frame, based on both the stream and the
  frame aspect ratio.

  Since the frame aspect ratio is set by the codec but the stream aspect ratio
  is set by the demuxer, these two may not be equal. This function tries to
  return the value that you should use if you would like to display the frame.

  Basic logic is to use the stream aspect ratio if it is set to something sane
  otherwise use the frame aspect ratio. This way a container setting, which is
  usually easy to modify can override the coded value in the frames.

  @param format the format context which the stream is part of
  @param stream the stream which the frame is part of
  @param frame the frame with the aspect ratio to be determined
  @return the guessed (valid) sample_aspect_ratio, 0/1 if no idea
*/
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
/*
  Guess the frame rate, based on both the container and codec information.

  @param ctx the format context which the stream is part of
  @param stream the stream which the frame is part of
  @param frame the frame for which the frame rate should be determined, may be NULL
  @return the guessed (valid) frame rate, 0/1 if no idea
*/
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
/*
  Check if the stream st contained in s is matched by the stream specifier
  spec.

  See the "stream specifiers" chapter in the documentation for the syntax
  of spec.

  @return  >0 if st is matched by spec;
           0  if st is not matched by spec;
           AVERROR code if spec is invalid

  @note  A stream specifier can match several streams in the format.
*/
func AVFormatMatchStreamSpecifier(s *AVFormatContext, st *AVStream, spec *CStr) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_queue_attached_pictures ---

// AVFormatQueueAttachedPictures wraps avformat_queue_attached_pictures.
func AVFormatQueueAttachedPictures(s *AVFormatContext) (int, error) {
	var tmps *C.AVFormatContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avformat_queue_attached_pictures(tmps)
	return int(ret), WrapErr(int(ret))
}

// --- Function avformat_transfer_internal_stream_timing_info ---

// AVFormatTransferInternalStreamTimingInfo wraps avformat_transfer_internal_stream_timing_info.
/*
  Transfer internal timing information from one stream to another.

  This function is useful when doing stream copy.

  @param ofmt     target output format for ost
  @param ost      output stream which needs timings copy and adjustments
  @param ist      reference input stream to copy timings from
  @param copy_tb  define from where the stream codec timebase needs to be imported
*/
func AVFormatTransferInternalStreamTimingInfo(ofmt *AVOutputFormat, ost *AVStream, ist *AVStream, copyTb AVTimebaseSource) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_stream_get_codec_timebase ---

// AVStreamGetCodecTimebase wraps av_stream_get_codec_timebase.
/*
  Get the internal codec timebase from a stream.

  @param st  input stream to extract the timebase from
*/
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
/*
  Return the name of the protocol that will handle the passed URL.

  NULL is returned if no protocol could be found for the given URL.

  @return Name of the protocol or NULL.
*/
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
/*
  Return AVIO_FLAG_* access flags corresponding to the access permissions
  of the resource in url, or a negative value corresponding to an
  AVERROR code in case of failure. The returned access flags are
  masked by the value in flags.

  @note This function is intrinsically unsafe, in the sense that the
  checked resource may change its existence or permission status from
  one call to another. Thus you should not trust the returned value,
  unless you are sure that no other processes are accessing the
  checked resource.
*/
func AVIOCheck(url *CStr, flags int) (int, error) {
	var tmpurl *C.char
	if url != nil {
		tmpurl = url.ptr
	}
	ret := C.avio_check(tmpurl, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_open_dir ---

// AVIOOpenDir wraps avio_open_dir.
/*
  Open directory for reading.

  @param s       directory read context. Pointer to a NULL pointer must be passed.
  @param url     directory to be listed.
  @param options A dictionary filled with protocol-private options. On return
                 this parameter will be destroyed and replaced with a dictionary
                 containing options that were not found. May be NULL.
  @return >=0 on success or negative on error.
*/
func AVIOOpenDir(s **AVIODirContext, url *CStr, options **AVDictionary) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_read_dir ---

// AVIOReadDir wraps avio_read_dir.
/*
  Get next directory entry.

  Returned entry must be freed with avio_free_directory_entry(). In particular
  it may outlive AVIODirContext.

  @param s         directory read context.
  @param[out] next next entry or NULL when no more entries.
  @return >=0 on success or negative on error. End of list is not considered an
              error.
*/
func AVIOReadDir(s *AVIODirContext, next **AVIODirEntry) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_close_dir ---

// AVIOCloseDir wraps avio_close_dir.
/*
  Close directory.

  @note Entries created using avio_read_dir() are not deleted and must be
  freeded with avio_free_directory_entry().

  @param s         directory read context.
  @return >=0 on success or negative on error.
*/
func AVIOCloseDir(s **AVIODirContext) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_free_directory_entry ---

// AVIOFreeDirectoryEntry wraps avio_free_directory_entry.
/*
  Free entry allocated by avio_read_dir().

  @param entry entry to be freed.
*/
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
/*
  Free the supplied IO context and everything associated with it.

  @param s Double pointer to the IO context. This function will write NULL
  into s.
*/
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
/*
  Write a NULL-terminated string.
  @return number of bytes written.
*/
func AVIOPutStr(s *AVIOContext, str *CStr) (int, error) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpstr *C.char
	if str != nil {
		tmpstr = str.ptr
	}
	ret := C.avio_put_str(tmps, tmpstr)
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_put_str16le ---

// AVIOPutStr16Le wraps avio_put_str16le.
/*
  Convert an UTF-8 string to UTF-16LE and write it.
  @param s the AVIOContext
  @param str NULL-terminated UTF-8 string

  @return number of bytes written.
*/
func AVIOPutStr16Le(s *AVIOContext, str *CStr) (int, error) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpstr *C.char
	if str != nil {
		tmpstr = str.ptr
	}
	ret := C.avio_put_str16le(tmps, tmpstr)
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_put_str16be ---

// AVIOPutStr16Be wraps avio_put_str16be.
/*
  Convert an UTF-8 string to UTF-16BE and write it.
  @param s the AVIOContext
  @param str NULL-terminated UTF-8 string

  @return number of bytes written.
*/
func AVIOPutStr16Be(s *AVIOContext, str *CStr) (int, error) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	var tmpstr *C.char
	if str != nil {
		tmpstr = str.ptr
	}
	ret := C.avio_put_str16be(tmps, tmpstr)
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_write_marker ---

// AVIOWriteMarker wraps avio_write_marker.
/*
  Mark the written bytestream as a specific type.

  Zero-length ranges are omitted from the output.

  @param s    the AVIOContext
  @param time the stream time the current bytestream pos corresponds to
              (in AV_TIME_BASE units), or AV_NOPTS_VALUE if unknown or not
              applicable
  @param type the kind of data written starting at the current pos
*/
func AVIOWriteMarker(s *AVIOContext, time int64, _type AVIODataMarkerType) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_write_marker(tmps, C.int64_t(time), C.enum_AVIODataMarkerType(_type))
}

// --- Function avio_seek ---

// AVIOSeek wraps avio_seek.
/*
  fseek() equivalent for AVIOContext.
  @return new position or AVERROR.
*/
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
/*
  Skip given number of bytes forward
  @return new position or AVERROR.
*/
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
/*
  ftell() equivalent for AVIOContext.
  @return position or AVERROR.
*/
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
/*
  Get the filesize.
  @return filesize or AVERROR
*/
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
/*
  Similar to feof() but also returns nonzero on read errors.
  @return non zero if and only if at end of file or a read error happened when reading.
*/
func AVIOFeof(s *AVIOContext) (int, error) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_feof(tmps)
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_vprintf ---

// avio_vprintf skipped due to ap.

// --- Function avio_printf ---

// avio_printf skipped due to variadic arg.

// --- Function avio_print_string_array ---

// avio_print_string_array skipped due to strings

// --- Function avio_flush ---

// AVIOFlush wraps avio_flush.
/*
  Force flushing of buffered data.

  For write streams, force the buffered data to be immediately written to the output,
  without to wait to fill the internal buffer.

  For read streams, discard all currently buffered data, and advance the
  reported file position to that of the underlying stream. This does not
  read new data, and does not perform any seeks.
*/
func AVIOFlush(s *AVIOContext) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	C.avio_flush(tmps)
}

// --- Function avio_read ---

// AVIORead wraps avio_read.
/*
  Read size bytes from AVIOContext into buf.
  @return number of bytes read or AVERROR
*/
func AVIORead(s *AVIOContext, buf unsafe.Pointer, size int) (int, error) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_read(tmps, (*C.uchar)(buf), C.int(size))
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_read_partial ---

// AVIOReadPartial wraps avio_read_partial.
/*
  Read size bytes from AVIOContext into buf. Unlike avio_read(), this is allowed
  to read fewer bytes than requested. The missing bytes can be read in the next
  call. This always tries to read at least 1 byte.
  Useful to reduce latency in certain cases.
  @return number of bytes read or AVERROR
*/
func AVIOReadPartial(s *AVIOContext, buf unsafe.Pointer, size int) (int, error) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_read_partial(tmps, (*C.uchar)(buf), C.int(size))
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_r8 ---

// AVIOR8 wraps avio_r8.
/*

  @note return 0 if EOF, so you cannot use it if EOF handling is
        necessary
*/
func AVIOR8(s *AVIOContext) (int, error) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_r8(tmps)
	return int(ret), WrapErr(int(ret))
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
/*
  Read a string from pb into buf. The reading will terminate when either
  a NULL character was encountered, maxlen bytes have been read, or nothing
  more can be read from pb. The result is guaranteed to be NULL-terminated, it
  will be truncated if buf is too small.
  Note that the string is not interpreted or validated in any way, it
  might get truncated in the middle of a sequence for multi-byte encodings.

  @return number of bytes read (is always <= maxlen).
  If reading ends on EOF or error, the return value will be one more than
  bytes actually read.
*/
func AVIOGetStr(pb *AVIOContext, maxlen int, buf *CStr, buflen int) (int, error) {
	var tmppb *C.AVIOContext
	if pb != nil {
		tmppb = pb.ptr
	}
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.avio_get_str(tmppb, C.int(maxlen), tmpbuf, C.int(buflen))
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_get_str16le ---

// AVIOGetStr16Le wraps avio_get_str16le.
/*
  Read a UTF-16 string from pb and convert it to UTF-8.
  The reading will terminate when either a null or invalid character was
  encountered or maxlen bytes have been read.
  @return number of bytes read (is always <= maxlen)
*/
func AVIOGetStr16Le(pb *AVIOContext, maxlen int, buf *CStr, buflen int) (int, error) {
	var tmppb *C.AVIOContext
	if pb != nil {
		tmppb = pb.ptr
	}
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.avio_get_str16le(tmppb, C.int(maxlen), tmpbuf, C.int(buflen))
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_get_str16be ---

// AVIOGetStr16Be wraps avio_get_str16be.
func AVIOGetStr16Be(pb *AVIOContext, maxlen int, buf *CStr, buflen int) (int, error) {
	var tmppb *C.AVIOContext
	if pb != nil {
		tmppb = pb.ptr
	}
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.avio_get_str16be(tmppb, C.int(maxlen), tmpbuf, C.int(buflen))
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_open ---

// AVIOOpen wraps avio_open.
/*
  Create and initialize a AVIOContext for accessing the
  resource indicated by url.
  @note When the resource indicated by url has been opened in
  read+write mode, the AVIOContext can be used only for writing.

  @param s Used to return the pointer to the created AVIOContext.
  In case of failure the pointed to value is set to NULL.
  @param url resource to access
  @param flags flags which control how the resource indicated by url
  is to be opened
  @return >= 0 in case of success, a negative value corresponding to an
  AVERROR code in case of failure
*/
func AVIOOpen(s **AVIOContext, url *CStr, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_open2 ---

// AVIOOpen2 wraps avio_open2.
/*
  Create and initialize a AVIOContext for accessing the
  resource indicated by url.
  @note When the resource indicated by url has been opened in
  read+write mode, the AVIOContext can be used only for writing.

  @param s Used to return the pointer to the created AVIOContext.
  In case of failure the pointed to value is set to NULL.
  @param url resource to access
  @param flags flags which control how the resource indicated by url
  is to be opened
  @param int_cb an interrupt callback to be used at the protocols level
  @param options  A dictionary filled with protocol-private options. On return
  this parameter will be destroyed and replaced with a dict containing options
  that were not found. May be NULL.
  @return >= 0 in case of success, a negative value corresponding to an
  AVERROR code in case of failure
*/
func AVIOOpen2(s **AVIOContext, url *CStr, flags int, intCb *AVIOInterruptCB, options **AVDictionary) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_close ---

// AVIOClose wraps avio_close.
/*
  Close the resource accessed by the AVIOContext s and free it.
  This function can only be used if s was opened by avio_open().

  The internal buffer is automatically flushed before closing the
  resource.

  @return 0 on success, an AVERROR < 0 on error.
  @see avio_closep
*/
func AVIOClose(s *AVIOContext) (int, error) {
	var tmps *C.AVIOContext
	if s != nil {
		tmps = s.ptr
	}
	ret := C.avio_close(tmps)
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_closep ---

// AVIOClosep wraps avio_closep.
/*
  Close the resource accessed by the AVIOContext *s, free it
  and set the pointer pointing to it to NULL.
  This function can only be used if s was opened by avio_open().

  The internal buffer is automatically flushed before closing the
  resource.

  @return 0 on success, an AVERROR < 0 on error.
  @see avio_close
*/
func AVIOClosep(s **AVIOContext) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_open_dyn_buf ---

// AVIOOpenDynBuf wraps avio_open_dyn_buf.
/*
  Open a write only memory stream.

  @param s new IO context
  @return zero if no error.
*/
func AVIOOpenDynBuf(s **AVIOContext) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_get_dyn_buf ---

// avio_get_dyn_buf skipped due to pbuffer

// --- Function avio_close_dyn_buf ---

// avio_close_dyn_buf skipped due to pbuffer

// --- Function avio_enum_protocols ---

// avio_enum_protocols skipped due to opaque

// --- Function avio_protocol_get_class ---

// AVIOProtocolGetClass wraps avio_protocol_get_class.
/*
  Get AVClass by names of available protocols.

  @return A AVClass of input protocol name or NULL
*/
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
/*
  Pause and resume playing - only meaningful if using a network streaming
  protocol (e.g. MMS).

  @param h     IO context from which to call the read_pause function pointer
  @param pause 1 for pause, 0 for resume
*/
func AVIOPause(h *AVIOContext, pause int) (int, error) {
	var tmph *C.AVIOContext
	if h != nil {
		tmph = h.ptr
	}
	ret := C.avio_pause(tmph, C.int(pause))
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_seek_time ---

// AVIOSeekTime wraps avio_seek_time.
/*
  Seek to a given timestamp relative to some component stream.
  Only meaningful if using a network streaming protocol (e.g. MMS.).

  @param h IO context from which to call the seek function pointers
  @param stream_index The stream index that the timestamp is relative to.
         If stream_index is (-1) the timestamp should be in AV_TIME_BASE
         units from the beginning of the presentation.
         If a stream_index >= 0 is used and the protocol does not support
         seeking based on component streams, the call will fail.
  @param timestamp timestamp in AVStream.time_base units
         or if there is no stream specified then in AV_TIME_BASE units.
  @param flags Optional combination of AVSEEK_FLAG_BACKWARD, AVSEEK_FLAG_BYTE
         and AVSEEK_FLAG_ANY. The protocol may silently ignore
         AVSEEK_FLAG_BACKWARD and AVSEEK_FLAG_ANY, but AVSEEK_FLAG_BYTE will
         fail if used and not supported.
  @return >= 0 on success
  @see AVInputFormat::read_seek
*/
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
/*
  Read contents of h into print buffer, up to max_size bytes, or up to EOF.

  @return 0 for success (max_size bytes read or EOF reached), negative error
  code otherwise
*/
func AVIOReadToBprint(h *AVIOContext, pb *AVBPrint, maxSize uint64) (int, error) {
	var tmph *C.AVIOContext
	if h != nil {
		tmph = h.ptr
	}
	var tmppb *C.struct_AVBPrint
	if pb != nil {
		tmppb = pb.ptr
	}
	ret := C.avio_read_to_bprint(tmph, tmppb, C.size_t(maxSize))
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_accept ---

// AVIOAccept wraps avio_accept.
/*
  Accept and allocate a client context on a server context.
  @param  s the server context
  @param  c the client context, must be unallocated
  @return   >= 0 on success or a negative value corresponding
            to an AVERROR on failure
*/
func AVIOAccept(s *AVIOContext, c **AVIOContext) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function avio_handshake ---

// AVIOHandshake wraps avio_handshake.
/*
  Perform one step of the protocol handshake to accept a new client.
  This function must be called on a client returned by avio_accept() before
  using it as a read/write context.
  It is separate from avio_accept() because it may block.
  A step of the handshake is defined by places where the application may
  decide to change the proceedings.
  For example, on a protocol with a request header and a reply header, each
  one can constitute a step because the application may use the parameters
  from the request to change parameters in the reply; or each individual
  chunk of the request can constitute a step.
  If the handshake is already finished, avio_handshake() does nothing and
  returns 0 immediately.

  @param  c the client context to perform the handshake on
  @return   0   on a complete and successful handshake
            > 0 if the handshake progressed, but is not complete
            < 0 for an AVERROR code
*/
func AVIOHandshake(c *AVIOContext) (int, error) {
	var tmpc *C.AVIOContext
	if c != nil {
		tmpc = c.ptr
	}
	ret := C.avio_handshake(tmpc)
	return int(ret), WrapErr(int(ret))
}

// --- Function avutil_version ---

// AVUtilVersion wraps avutil_version.
//
//	Return the LIBAVUTIL_VERSION_INT constant.
func AVUtilVersion() uint {
	ret := C.avutil_version()
	return uint(ret)
}

// --- Function av_version_info ---

// AVVersionInfo wraps av_version_info.
/*
  Return an informative version string. This usually is the actual release
  version number or a git commit description. This string has no fixed format
  and can change any time. It should never be parsed by code.
*/
func AVVersionInfo() *CStr {
	ret := C.av_version_info()
	return wrapCStr(ret)
}

// --- Function avutil_configuration ---

// AVUtilConfiguration wraps avutil_configuration.
//
//	Return the libavutil build-time configuration.
func AVUtilConfiguration() *CStr {
	ret := C.avutil_configuration()
	return wrapCStr(ret)
}

// --- Function avutil_license ---

// AVUtilLicense wraps avutil_license.
//
//	Return the libavutil license.
func AVUtilLicense() *CStr {
	ret := C.avutil_license()
	return wrapCStr(ret)
}

// --- Function av_get_media_type_string ---

// AVGetMediaTypeString wraps av_get_media_type_string.
/*
  Return a string describing the media_type enum, NULL if media_type
  is unknown.
*/
func AVGetMediaTypeString(mediaType AVMediaType) *CStr {
	ret := C.av_get_media_type_string(C.enum_AVMediaType(mediaType))
	return wrapCStr(ret)
}

// --- Function av_get_picture_type_char ---

// AVGetPictureTypeChar wraps av_get_picture_type_char.
/*
  Return a single letter to describe the given picture type
  pict_type.

  @param[in] pict_type the picture type @return a single character
  representing the picture type, '?' if pict_type is unknown
*/
func AVGetPictureTypeChar(pictType AVPictureType) uint8 {
	ret := C.av_get_picture_type_char(C.enum_AVPictureType(pictType))
	return uint8(ret)
}

// --- Function av_x_if_null ---

// AVXIfNull wraps av_x_if_null.
//
//	Return x default pointer in case p is NULL.
func AVXIfNull(p unsafe.Pointer, x unsafe.Pointer) unsafe.Pointer {
	ret := C.av_x_if_null(p, x)
	return ret
}

// --- Function av_int_list_length_for_size ---

// AVIntListLengthForSize wraps av_int_list_length_for_size.
/*
  Compute the length of an integer list.

  @param elsize  size in bytes of each list element (only 1, 2, 4 or 8)
  @param term    list terminator (usually 0 or -1)
  @param list    pointer to the list
  @return  length of the list, in elements, not counting the terminator
*/
func AVIntListLengthForSize(elsize uint, list unsafe.Pointer, term uint64) uint {
	ret := C.av_int_list_length_for_size(C.uint(elsize), list, C.uint64_t(term))
	return uint(ret)
}

// --- Function av_fopen_utf8 ---

// av_fopen_utf8 skipped due to return.

// --- Function av_get_time_base_q ---

// AVGetTimeBaseQ wraps av_get_time_base_q.
//
//	Return the fractional representation of the internal time base.
func AVGetTimeBaseQ() *AVRational {
	ret := C.av_get_time_base_q()
	return &AVRational{value: ret}
}

// --- Function av_fourcc_make_string ---

// AVFourccMakeString wraps av_fourcc_make_string.
/*
  Fill the provided buffer with a string containing a FourCC (four-character
  code) representation.

  @param buf    a buffer with size in bytes of at least AV_FOURCC_MAX_STRING_SIZE
  @param fourcc the fourcc to represent
  @return the buffer in input
*/
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
/*
  Allocate an AVBuffer of the given size using av_malloc().

  @return an AVBufferRef of given size or NULL when out of memory
*/
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
/*
  Same as av_buffer_alloc(), except the returned buffer will be initialized
  to zero.
*/
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
/*
  Default free callback, which calls av_free() on the buffer data.
  This function is meant to be passed to av_buffer_create(), not called
  directly.
*/
func AVBufferDefaultFree(opaque unsafe.Pointer, data unsafe.Pointer) {
	C.av_buffer_default_free(opaque, (*C.uint8_t)(data))
}

// --- Function av_buffer_ref ---

// AVBufferRef_ wraps av_buffer_ref.
/*
  Create a new reference to an AVBuffer.

  @return a new AVBufferRef referring to the same AVBuffer as buf or NULL on
  failure.
*/
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
/*
  Free a given reference and automatically free the buffer if there are no more
  references to it.

  @param buf the reference to be freed. The pointer is set to NULL on return.
*/
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
/*
  @return 1 if the caller may write to the data referred to by buf (which is
  true if and only if buf is the only reference to the underlying AVBuffer).
  Return 0 otherwise.
  A positive answer is valid until av_buffer_ref() is called on buf.
*/
func AVBufferIsWritable(buf *AVBufferRef) (int, error) {
	var tmpbuf *C.AVBufferRef
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_buffer_is_writable(tmpbuf)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffer_get_opaque ---

// AVBufferGetOpaque wraps av_buffer_get_opaque.
//
//	@return the opaque parameter set by av_buffer_create.
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
func AVBufferGetRefCount(buf *AVBufferRef) (int, error) {
	var tmpbuf *C.AVBufferRef
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_buffer_get_ref_count(tmpbuf)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffer_make_writable ---

// AVBufferMakeWritable wraps av_buffer_make_writable.
/*
  Create a writable reference from a given buffer reference, avoiding data copy
  if possible.

  @param buf buffer reference to make writable. On success, buf is either left
             untouched, or it is unreferenced and a new writable AVBufferRef is
             written in its place. On failure, buf is left untouched.
  @return 0 on success, a negative AVERROR on failure.
*/
func AVBufferMakeWritable(buf **AVBufferRef) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffer_realloc ---

// AVBufferRealloc wraps av_buffer_realloc.
/*
  Reallocate a given buffer.

  @param buf  a buffer reference to reallocate. On success, buf will be
              unreferenced and a new reference with the required size will be
              written in its place. On failure buf will be left untouched. *buf
              may be NULL, then a new buffer is allocated.
  @param size required new buffer size.
  @return 0 on success, a negative AVERROR on failure.

  @note the buffer is actually reallocated with av_realloc() only if it was
  initially allocated through av_buffer_realloc(NULL) and there is only one
  reference to it (i.e. the one passed to this function). In all other cases
  a new buffer is allocated and the data is copied.
*/
func AVBufferRealloc(buf **AVBufferRef, size uint64) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffer_replace ---

// AVBufferReplace wraps av_buffer_replace.
/*
  Ensure dst refers to the same data as src.

  When *dst is already equivalent to src, do nothing. Otherwise unreference dst
  and replace it with a new reference to src.

  @param dst Pointer to either a valid buffer reference or NULL. On success,
             this will point to a buffer reference equivalent to src. On
             failure, dst will be left untouched.
  @param src A buffer reference to replace dst with. May be NULL, then this
             function is equivalent to av_buffer_unref(dst).
  @return 0 on success
          AVERROR(ENOMEM) on memory allocation failure.
*/
func AVBufferReplace(dst **AVBufferRef, src *AVBufferRef) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_buffer_pool_init ---

// av_buffer_pool_init skipped due to alloc.

// --- Function av_buffer_pool_init2 ---

// av_buffer_pool_init2 skipped due to alloc.

// --- Function av_buffer_pool_uninit ---

// AVBufferPoolUninit wraps av_buffer_pool_uninit.
/*
  Mark the pool as being available for freeing. It will actually be freed only
  once all the allocated buffers associated with the pool are released. Thus it
  is safe to call this function while some of the allocated buffers are still
  in use.

  @param pool pointer to the pool to be freed. It will be set to NULL.
*/
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
/*
  Allocate a new AVBuffer, reusing an old buffer from the pool when available.
  This function may be called simultaneously from multiple threads.

  @return a reference to the new buffer on success, NULL on error.
*/
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
/*
  Query the original opaque parameter of an allocated buffer in the pool.

  @param ref a buffer reference to a buffer returned by av_buffer_pool_get.
  @return the opaque parameter set by the buffer allocator function of the
          buffer pool.

  @note the opaque parameter of ref is used by the buffer pool implementation,
  therefore you have to use this function to access the original opaque
  parameter of an allocated buffer.
*/
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
/*
  Return a channel layout id that matches name, or 0 if no match is found.

  name can be one or several of the following notations,
  separated by '+' or '|':
  - the name of an usual channel layout (mono, stereo, 4.0, quad, 5.0,
    5.0(side), 5.1, 5.1(side), 7.1, 7.1(wide), downmix);
  - the name of a single channel (FL, FR, FC, LFE, BL, BR, FLC, FRC, BC,
    SL, SR, TC, TFL, TFC, TFR, TBL, TBC, TBR, DL, DR);
  - a number of channels, in decimal, followed by 'c', yielding
    the default channel layout for that number of channels (@see
    av_get_default_channel_layout);
  - a channel layout mask, in hexadecimal starting with "0x" (see the
    AV_CH_* macros).

  Example: "stereo+FC" = "2c+FC" = "2c+1c" = "0x7"

  @deprecated use av_channel_layout_from_string()
*/
func AVGetChannelLayout(name *CStr) uint64 {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_get_channel_layout(tmpname)
	return uint64(ret)
}

// --- Function av_get_extended_channel_layout ---

// av_get_extended_channel_layout skipped due to channelLayout

// --- Function av_get_channel_layout_string ---

// AVGetChannelLayoutString wraps av_get_channel_layout_string.
/*
  Return a description of a channel layout.
  If nb_channels is <= 0, it is guessed from the channel_layout.

  @param buf put here the string containing the channel layout
  @param buf_size size in bytes of the buffer
  @param nb_channels number of channels
  @param channel_layout channel layout bitset
  @deprecated use av_channel_layout_describe()
*/
func AVGetChannelLayoutString(buf *CStr, bufSize int, nbChannels int, channelLayout uint64) {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	C.av_get_channel_layout_string(tmpbuf, C.int(bufSize), C.int(nbChannels), C.uint64_t(channelLayout))
}

// --- Function av_bprint_channel_layout ---

// AVBprintChannelLayout wraps av_bprint_channel_layout.
/*
  Append a description of a channel layout to a bprint buffer.
  @deprecated use av_channel_layout_describe()
*/
func AVBprintChannelLayout(bp *AVBPrint, nbChannels int, channelLayout uint64) {
	var tmpbp *C.struct_AVBPrint
	if bp != nil {
		tmpbp = bp.ptr
	}
	C.av_bprint_channel_layout(tmpbp, C.int(nbChannels), C.uint64_t(channelLayout))
}

// --- Function av_get_channel_layout_nb_channels ---

// AVGetChannelLayoutNbChannels wraps av_get_channel_layout_nb_channels.
/*
  Return the number of channels in the channel layout.
  @deprecated use AVChannelLayout.nb_channels
*/
func AVGetChannelLayoutNbChannels(channelLayout uint64) (int, error) {
	ret := C.av_get_channel_layout_nb_channels(C.uint64_t(channelLayout))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_get_default_channel_layout ---

// AVGetDefaultChannelLayout wraps av_get_default_channel_layout.
/*
  Return default channel layout for a given number of channels.

  @deprecated use av_channel_layout_default()
*/
func AVGetDefaultChannelLayout(nbChannels int) int64 {
	ret := C.av_get_default_channel_layout(C.int(nbChannels))
	return int64(ret)
}

// --- Function av_get_channel_layout_channel_index ---

// AVGetChannelLayoutChannelIndex wraps av_get_channel_layout_channel_index.
/*
  Get the index of a channel in channel_layout.

  @param channel_layout channel layout bitset
  @param channel a channel layout describing exactly one channel which must be
                 present in channel_layout.

  @return index of channel in channel_layout on success, a negative AVERROR
          on error.

  @deprecated use av_channel_layout_index_from_channel()
*/
func AVGetChannelLayoutChannelIndex(channelLayout uint64, channel uint64) (int, error) {
	ret := C.av_get_channel_layout_channel_index(C.uint64_t(channelLayout), C.uint64_t(channel))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_channel_layout_extract_channel ---

// AVChannelLayoutExtractChannel wraps av_channel_layout_extract_channel.
/*
  Get the channel with the given index in channel_layout.
  @deprecated use av_channel_layout_channel_from_index()
*/
func AVChannelLayoutExtractChannel(channelLayout uint64, index int) uint64 {
	ret := C.av_channel_layout_extract_channel(C.uint64_t(channelLayout), C.int(index))
	return uint64(ret)
}

// --- Function av_get_channel_name ---

// AVGetChannelName wraps av_get_channel_name.
/*
  Get the name of a given channel.

  @return channel name on success, NULL on error.

  @deprecated use av_channel_name()
*/
func AVGetChannelName(channel uint64) *CStr {
	ret := C.av_get_channel_name(C.uint64_t(channel))
	return wrapCStr(ret)
}

// --- Function av_get_channel_description ---

// AVGetChannelDescription wraps av_get_channel_description.
/*
  Get the description of a given channel.

  @param channel  a channel layout with a single channel
  @return  channel description on success, NULL on error
  @deprecated use av_channel_description()
*/
func AVGetChannelDescription(channel uint64) *CStr {
	ret := C.av_get_channel_description(C.uint64_t(channel))
	return wrapCStr(ret)
}

// --- Function av_get_standard_channel_layout ---

// av_get_standard_channel_layout skipped due to layout

// --- Function av_channel_name ---

// AVChannelName wraps av_channel_name.
/*
  Get a human readable string in an abbreviated form describing a given channel.
  This is the inverse function of @ref av_channel_from_string().

  @param buf pre-allocated buffer where to put the generated string
  @param buf_size size in bytes of the buffer.
  @param channel the AVChannel whose name to get
  @return amount of bytes needed to hold the output string, or a negative AVERROR
          on failure. If the returned value is bigger than buf_size, then the
          string was truncated.
*/
func AVChannelName(buf *CStr, bufSize uint64, channel AVChannel) (int, error) {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_channel_name(tmpbuf, C.size_t(bufSize), C.enum_AVChannel(channel))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_channel_name_bprint ---

// AVChannelNameBprint wraps av_channel_name_bprint.
/*
  bprint variant of av_channel_name().

  @note the string will be appended to the bprint buffer.
*/
func AVChannelNameBprint(bp *AVBPrint, channelId AVChannel) {
	var tmpbp *C.struct_AVBPrint
	if bp != nil {
		tmpbp = bp.ptr
	}
	C.av_channel_name_bprint(tmpbp, C.enum_AVChannel(channelId))
}

// --- Function av_channel_description ---

// AVChannelDescription wraps av_channel_description.
/*
  Get a human readable string describing a given channel.

  @param buf pre-allocated buffer where to put the generated string
  @param buf_size size in bytes of the buffer.
  @param channel the AVChannel whose description to get
  @return amount of bytes needed to hold the output string, or a negative AVERROR
          on failure. If the returned value is bigger than buf_size, then the
          string was truncated.
*/
func AVChannelDescription(buf *CStr, bufSize uint64, channel AVChannel) (int, error) {
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_channel_description(tmpbuf, C.size_t(bufSize), C.enum_AVChannel(channel))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_channel_description_bprint ---

// AVChannelDescriptionBprint wraps av_channel_description_bprint.
/*
  bprint variant of av_channel_description().

  @note the string will be appended to the bprint buffer.
*/
func AVChannelDescriptionBprint(bp *AVBPrint, channelId AVChannel) {
	var tmpbp *C.struct_AVBPrint
	if bp != nil {
		tmpbp = bp.ptr
	}
	C.av_channel_description_bprint(tmpbp, C.enum_AVChannel(channelId))
}

// --- Function av_channel_from_string ---

// AVChannelFromString wraps av_channel_from_string.
/*
  This is the inverse function of @ref av_channel_name().

  @return the channel with the given name
          AV_CHAN_NONE when name does not identify a known channel
*/
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
/*
  Initialize a native channel layout from a bitmask indicating which channels
  are present.

  @param channel_layout the layout structure to be initialized
  @param mask bitmask describing the channel layout

  @return 0 on success
          AVERROR(EINVAL) for invalid mask values
*/
func AVChannelLayoutFromMask(channelLayout *AVChannelLayout, mask uint64) (int, error) {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	ret := C.av_channel_layout_from_mask(tmpchannelLayout, C.uint64_t(mask))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_channel_layout_from_string ---

// AVChannelLayoutFromString wraps av_channel_layout_from_string.
/*
  Initialize a channel layout from a given string description.
  The input string can be represented by:
   - the formal channel layout name (returned by av_channel_layout_describe())
   - single or multiple channel names (returned by av_channel_name(), eg. "FL",
     or concatenated with "+", each optionally containing a custom name after
     a "@", eg. "FL@Left+FR@Right+LFE")
   - a decimal or hexadecimal value of a native channel layout (eg. "4" or "0x4")
   - the number of channels with default layout (eg. "4c")
   - the number of unordered channels (eg. "4C" or "4 channels")
   - the ambisonic order followed by optional non-diegetic channels (eg.
     "ambisonic 2+stereo")

  @param channel_layout input channel layout
  @param str string describing the channel layout
  @return 0 channel layout was detected, AVERROR_INVALIDATATA otherwise
*/
func AVChannelLayoutFromString(channelLayout *AVChannelLayout, str *CStr) (int, error) {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	var tmpstr *C.char
	if str != nil {
		tmpstr = str.ptr
	}
	ret := C.av_channel_layout_from_string(tmpchannelLayout, tmpstr)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_channel_layout_default ---

// AVChannelLayoutDefault wraps av_channel_layout_default.
/*
  Get the default channel layout for a given number of channels.

  @param ch_layout the layout structure to be initialized
  @param nb_channels number of channels
*/
func AVChannelLayoutDefault(chLayout *AVChannelLayout, nbChannels int) {
	var tmpchLayout *C.AVChannelLayout
	if chLayout != nil {
		tmpchLayout = chLayout.ptr
	}
	C.av_channel_layout_default(tmpchLayout, C.int(nbChannels))
}

// --- Function av_channel_layout_standard ---

// av_channel_layout_standard skipped due to opaque

// --- Function av_channel_layout_uninit ---

// AVChannelLayoutUninit wraps av_channel_layout_uninit.
/*
  Free any allocated data in the channel layout and reset the channel
  count to 0.

  @param channel_layout the layout structure to be uninitialized
*/
func AVChannelLayoutUninit(channelLayout *AVChannelLayout) {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	C.av_channel_layout_uninit(tmpchannelLayout)
}

// --- Function av_channel_layout_copy ---

// AVChannelLayoutCopy wraps av_channel_layout_copy.
/*
  Make a copy of a channel layout. This differs from just assigning src to dst
  in that it allocates and copies the map for AV_CHANNEL_ORDER_CUSTOM.

  @note the destination channel_layout will be always uninitialized before copy.

  @param dst destination channel layout
  @param src source channel layout
  @return 0 on success, a negative AVERROR on error.
*/
func AVChannelLayoutCopy(dst *AVChannelLayout, src *AVChannelLayout) (int, error) {
	var tmpdst *C.AVChannelLayout
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVChannelLayout
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_channel_layout_copy(tmpdst, tmpsrc)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_channel_layout_describe ---

// AVChannelLayoutDescribe wraps av_channel_layout_describe.
/*
  Get a human-readable string describing the channel layout properties.
  The string will be in the same format that is accepted by
  @ref av_channel_layout_from_string(), allowing to rebuild the same
  channel layout, except for opaque pointers.

  @param channel_layout channel layout to be described
  @param buf pre-allocated buffer where to put the generated string
  @param buf_size size in bytes of the buffer.
  @return amount of bytes needed to hold the output string, or a negative AVERROR
          on failure. If the returned value is bigger than buf_size, then the
          string was truncated.
*/
func AVChannelLayoutDescribe(channelLayout *AVChannelLayout, buf *CStr, bufSize uint64) (int, error) {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	var tmpbuf *C.char
	if buf != nil {
		tmpbuf = buf.ptr
	}
	ret := C.av_channel_layout_describe(tmpchannelLayout, tmpbuf, C.size_t(bufSize))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_channel_layout_describe_bprint ---

// AVChannelLayoutDescribeBprint wraps av_channel_layout_describe_bprint.
/*
  bprint variant of av_channel_layout_describe().

  @note the string will be appended to the bprint buffer.
  @return 0 on success, or a negative AVERROR value on failure.
*/
func AVChannelLayoutDescribeBprint(channelLayout *AVChannelLayout, bp *AVBPrint) (int, error) {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	var tmpbp *C.struct_AVBPrint
	if bp != nil {
		tmpbp = bp.ptr
	}
	ret := C.av_channel_layout_describe_bprint(tmpchannelLayout, tmpbp)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_channel_layout_channel_from_index ---

// AVChannelLayoutChannelFromIndex wraps av_channel_layout_channel_from_index.
/*
  Get the channel with the given index in a channel layout.

  @param channel_layout input channel layout
  @param idx index of the channel
  @return channel with the index idx in channel_layout on success or
          AV_CHAN_NONE on failure (if idx is not valid or the channel order is
          unspecified)
*/
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
/*
  Get the index of a given channel in a channel layout. In case multiple
  channels are found, only the first match will be returned.

  @param channel_layout input channel layout
  @param channel the channel whose index to obtain
  @return index of channel in channel_layout on success or a negative number if
          channel is not present in channel_layout.
*/
func AVChannelLayoutIndexFromChannel(channelLayout *AVChannelLayout, channel AVChannel) (int, error) {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	ret := C.av_channel_layout_index_from_channel(tmpchannelLayout, C.enum_AVChannel(channel))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_channel_layout_index_from_string ---

// AVChannelLayoutIndexFromString wraps av_channel_layout_index_from_string.
/*
  Get the index in a channel layout of a channel described by the given string.
  In case multiple channels are found, only the first match will be returned.

  This function accepts channel names in the same format as
  @ref av_channel_from_string().

  @param channel_layout input channel layout
  @param name string describing the channel whose index to obtain
  @return a channel index described by the given string, or a negative AVERROR
          value.
*/
func AVChannelLayoutIndexFromString(channelLayout *AVChannelLayout, name *CStr) (int, error) {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_channel_layout_index_from_string(tmpchannelLayout, tmpname)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_channel_layout_channel_from_string ---

// AVChannelLayoutChannelFromString wraps av_channel_layout_channel_from_string.
/*
  Get a channel described by the given string.

  This function accepts channel names in the same format as
  @ref av_channel_from_string().

  @param channel_layout input channel layout
  @param name string describing the channel to obtain
  @return a channel described by the given string in channel_layout on success
          or AV_CHAN_NONE on failure (if the string is not valid or the channel
          order is unspecified)
*/
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
/*
  Find out what channels from a given set are present in a channel layout,
  without regard for their positions.

  @param channel_layout input channel layout
  @param mask a combination of AV_CH_* representing a set of channels
  @return a bitfield representing all the channels from mask that are present
          in channel_layout
*/
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
/*
  Check whether a channel layout is valid, i.e. can possibly describe audio
  data.

  @param channel_layout input channel layout
  @return 1 if channel_layout is valid, 0 otherwise.
*/
func AVChannelLayoutCheck(channelLayout *AVChannelLayout) (int, error) {
	var tmpchannelLayout *C.AVChannelLayout
	if channelLayout != nil {
		tmpchannelLayout = channelLayout.ptr
	}
	ret := C.av_channel_layout_check(tmpchannelLayout)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_channel_layout_compare ---

// AVChannelLayoutCompare wraps av_channel_layout_compare.
/*
  Check whether two channel layouts are semantically the same, i.e. the same
  channels are present on the same positions in both.

  If one of the channel layouts is AV_CHANNEL_ORDER_UNSPEC, while the other is
  not, they are considered to be unequal. If both are AV_CHANNEL_ORDER_UNSPEC,
  they are considered equal iff the channel counts are the same in both.

  @param chl input channel layout
  @param chl1 input channel layout
  @return 0 if chl and chl1 are equal, 1 if they are not equal. A negative
          AVERROR code if one or both are invalid.
*/
func AVChannelLayoutCompare(chl *AVChannelLayout, chl1 *AVChannelLayout) (int, error) {
	var tmpchl *C.AVChannelLayout
	if chl != nil {
		tmpchl = chl.ptr
	}
	var tmpchl1 *C.AVChannelLayout
	if chl1 != nil {
		tmpchl1 = chl1.ptr
	}
	ret := C.av_channel_layout_compare(tmpchl, tmpchl1)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_dict_get ---

// AVDictGet wraps av_dict_get.
/*
  Get a dictionary entry with matching key.

  The returned entry key or value must not be changed, or it will
  cause undefined behavior.

  @param prev  Set to the previous matching element to find the next.
               If set to NULL the first matching element is returned.
  @param key   Matching key
  @param flags A collection of AV_DICT_* flags controlling how the
               entry is retrieved

  @return      Found entry or NULL in case no matching entry was found in the dictionary
*/
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
/*
  Iterate over a dictionary

  Iterates through all entries in the dictionary.

  @warning The returned AVDictionaryEntry key/value must not be changed.

  @warning As av_dict_set() invalidates all previous entries returned
  by this function, it must not be called while iterating over the dict.

  Typical usage:
  @code
  const AVDictionaryEntry *e = NULL;
  while ((e = av_dict_iterate(m, e))) {
      // ...
  }
  @endcode

  @param m     The dictionary to iterate over
  @param prev  Pointer to the previous AVDictionaryEntry, NULL initially

  @retval AVDictionaryEntry* The next element in the dictionary
  @retval NULL               No more elements in the dictionary
*/
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
/*
  Get number of entries in dictionary.

  @param m dictionary
  @return  number of entries in dictionary
*/
func AVDictCount(m *AVDictionary) (int, error) {
	var tmpm *C.AVDictionary
	if m != nil {
		tmpm = m.ptr
	}
	ret := C.av_dict_count(tmpm)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_dict_set ---

// AVDictSet wraps av_dict_set.
/*
  Set the given entry in *pm, overwriting an existing entry.

  Note: If AV_DICT_DONT_STRDUP_KEY or AV_DICT_DONT_STRDUP_VAL is set,
  these arguments will be freed on error.

  @warning Adding a new entry to a dictionary invalidates all existing entries
  previously returned with av_dict_get() or av_dict_iterate().

  @param pm        Pointer to a pointer to a dictionary struct. If *pm is NULL
                   a dictionary struct is allocated and put in *pm.
  @param key       Entry key to add to *pm (will either be av_strduped or added as a new key depending on flags)
  @param value     Entry value to add to *pm (will be av_strduped or added as a new key depending on flags).
                   Passing a NULL value will cause an existing entry to be deleted.

  @return          >= 0 on success otherwise an error code <0
*/
func AVDictSet(pm **AVDictionary, key *CStr, value *CStr, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_dict_set_int ---

// AVDictSetInt wraps av_dict_set_int.
/*
  Convenience wrapper for av_dict_set() that converts the value to a string
  and stores it.

  Note: If ::AV_DICT_DONT_STRDUP_KEY is set, key will be freed on error.
*/
func AVDictSetInt(pm **AVDictionary, key *CStr, value int64, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_dict_parse_string ---

// AVDictParseString wraps av_dict_parse_string.
/*
  Parse the key/value pairs list and add the parsed entries to a dictionary.

  In case of failure, all the successfully set entries are stored in
  *pm. You may need to manually free the created dictionary.

  @param key_val_sep  A 0-terminated list of characters used to separate
                      key from value
  @param pairs_sep    A 0-terminated list of characters used to separate
                      two pairs from each other
  @param flags        Flags to use when adding to the dictionary.
                      ::AV_DICT_DONT_STRDUP_KEY and ::AV_DICT_DONT_STRDUP_VAL
                      are ignored since the key/value tokens will always
                      be duplicated.

  @return             0 on success, negative AVERROR code on failure
*/
func AVDictParseString(pm **AVDictionary, str *CStr, keyValSep *CStr, pairsSep *CStr, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_dict_copy ---

// AVDictCopy wraps av_dict_copy.
/*
  Copy entries from one AVDictionary struct into another.

  @note Metadata is read using the ::AV_DICT_IGNORE_SUFFIX flag

  @param dst   Pointer to a pointer to a AVDictionary struct to copy into. If *dst is NULL,
               this function will allocate a struct for you and put it in *dst
  @param src   Pointer to the source AVDictionary struct to copy items from.
  @param flags Flags to use when setting entries in *dst

  @return 0 on success, negative AVERROR code on failure. If dst was allocated
            by this function, callers should free the associated memory.
*/
func AVDictCopy(dst **AVDictionary, src *AVDictionary, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_dict_free ---

// AVDictFree wraps av_dict_free.
/*
  Free all the memory allocated for an AVDictionary struct
  and all keys and values.
*/
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

// av_dict_get_string skipped due to buffer

// --- Function av_strerror ---

// AVStrerror wraps av_strerror.
/*
  Put a description of the AVERROR code errnum in errbuf.
  In case of failure the global variable errno is set to indicate the
  error. Even in case of failure av_strerror() will print a generic
  error message indicating the errnum provided to errbuf.

  @param errnum      error code to describe
  @param errbuf      buffer to which description is written
  @param errbuf_size the size in bytes of errbuf
  @return 0 on success, a negative value if a description for errnum
  cannot be found
*/
func AVStrerror(errnum int, errbuf *CStr, errbufSize uint64) (int, error) {
	var tmperrbuf *C.char
	if errbuf != nil {
		tmperrbuf = errbuf.ptr
	}
	ret := C.av_strerror(C.int(errnum), tmperrbuf, C.size_t(errbufSize))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_make_error_string ---

// AVMakeErrorString wraps av_make_error_string.
/*
  Fill the provided buffer with a string containing an error string
  corresponding to the AVERROR code errnum.

  @param errbuf         a buffer
  @param errbuf_size    size in bytes of errbuf
  @param errnum         error code to describe
  @return the buffer in input, filled with the error description
  @see av_strerror()
*/
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
/*
  Allocate an AVFrame and set its fields to default values.  The resulting
  struct must be freed using av_frame_free().

  @return An AVFrame filled with default values or NULL on failure.

  @note this only allocates the AVFrame itself, not the data buffers. Those
  must be allocated through other means, e.g. with av_frame_get_buffer() or
  manually.
*/
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
/*
  Free the frame and any dynamically allocated objects in it,
  e.g. extended_data. If the frame is reference counted, it will be
  unreferenced first.

  @param frame frame to be freed. The pointer will be set to NULL.
*/
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
/*
  Set up a new reference to the data described by the source frame.

  Copy frame properties from src to dst and create a new reference for each
  AVBufferRef from src.

  If src is not reference counted, new buffers are allocated and the data is
  copied.

  @warning: dst MUST have been either unreferenced with av_frame_unref(dst),
            or newly allocated with av_frame_alloc() before calling this
            function, or undefined behavior will occur.

  @return 0 on success, a negative AVERROR on error
*/
func AVFrameRef(dst *AVFrame, src *AVFrame) (int, error) {
	var tmpdst *C.AVFrame
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_frame_ref(tmpdst, tmpsrc)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_frame_clone ---

// AVFrameClone wraps av_frame_clone.
/*
  Create a new frame that references the same data as src.

  This is a shortcut for av_frame_alloc()+av_frame_ref().

  @return newly created AVFrame on success, NULL on error.
*/
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
//
//	Unreference all the buffers referenced by frame and reset the frame fields.
func AVFrameUnref(frame *AVFrame) {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	C.av_frame_unref(tmpframe)
}

// --- Function av_frame_move_ref ---

// AVFrameMoveRef wraps av_frame_move_ref.
/*
  Move everything contained in src to dst and reset src.

  @warning: dst is not unreferenced, but directly overwritten without reading
            or deallocating its contents. Call av_frame_unref(dst) manually
            before calling this function to ensure that no memory is leaked.
*/
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
/*
  Allocate new buffer(s) for audio or video data.

  The following fields must be set on frame before calling this function:
  - format (pixel format for video, sample format for audio)
  - width and height for video
  - nb_samples and ch_layout for audio

  This function will fill AVFrame.data and AVFrame.buf arrays and, if
  necessary, allocate and fill AVFrame.extended_data and AVFrame.extended_buf.
  For planar formats, one buffer will be allocated for each plane.

  @warning: if frame already has been allocated, calling this function will
            leak memory. In addition, undefined behavior can occur in certain
            cases.

  @param frame frame in which to store the new buffers.
  @param align Required buffer size alignment. If equal to 0, alignment will be
               chosen automatically for the current CPU. It is highly
               recommended to pass 0 here unless you know what you are doing.

  @return 0 on success, a negative AVERROR on error.
*/
func AVFrameGetBuffer(frame *AVFrame, align int) (int, error) {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_frame_get_buffer(tmpframe, C.int(align))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_frame_is_writable ---

// AVFrameIsWritable wraps av_frame_is_writable.
/*
  Check if the frame data is writable.

  @return A positive value if the frame data is writable (which is true if and
  only if each of the underlying buffers has only one reference, namely the one
  stored in this frame). Return 0 otherwise.

  If 1 is returned the answer is valid until av_buffer_ref() is called on any
  of the underlying AVBufferRefs (e.g. through av_frame_ref() or directly).

  @see av_frame_make_writable(), av_buffer_is_writable()
*/
func AVFrameIsWritable(frame *AVFrame) (int, error) {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_frame_is_writable(tmpframe)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_frame_make_writable ---

// AVFrameMakeWritable wraps av_frame_make_writable.
/*
  Ensure that the frame data is writable, avoiding data copy if possible.

  Do nothing if the frame is writable, allocate new buffers and copy the data
  if it is not. Non-refcounted frames behave as non-writable, i.e. a copy
  is always made.

  @return 0 on success, a negative AVERROR on error.

  @see av_frame_is_writable(), av_buffer_is_writable(),
  av_buffer_make_writable()
*/
func AVFrameMakeWritable(frame *AVFrame) (int, error) {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_frame_make_writable(tmpframe)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_frame_copy ---

// AVFrameCopy wraps av_frame_copy.
/*
  Copy the frame data from src to dst.

  This function does not allocate anything, dst must be already initialized and
  allocated with the same parameters as src.

  This function only copies the frame data (i.e. the contents of the data /
  extended data arrays), not any other properties.

  @return >= 0 on success, a negative AVERROR on error.
*/
func AVFrameCopy(dst *AVFrame, src *AVFrame) (int, error) {
	var tmpdst *C.AVFrame
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_frame_copy(tmpdst, tmpsrc)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_frame_copy_props ---

// AVFrameCopyProps wraps av_frame_copy_props.
/*
  Copy only "metadata" fields from src to dst.

  Metadata for the purpose of this function are those fields that do not affect
  the data layout in the buffers.  E.g. pts, sample rate (for audio) or sample
  aspect ratio (for video), but not width/height or channel layout.
  Side data is also copied.
*/
func AVFrameCopyProps(dst *AVFrame, src *AVFrame) (int, error) {
	var tmpdst *C.AVFrame
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_frame_copy_props(tmpdst, tmpsrc)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_frame_get_plane_buffer ---

// AVFrameGetPlaneBuffer wraps av_frame_get_plane_buffer.
/*
  Get the buffer reference a given data plane is stored in.

  @param frame the frame to get the plane's buffer from
  @param plane index of the data plane of interest in frame->extended_data.

  @return the buffer reference that contains the plane or NULL if the input
  frame is not valid.
*/
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
/*
  Add a new side data to a frame.

  @param frame a frame to which the side data should be added
  @param type type of the added side data
  @param size size of the side data

  @return newly added side data on success, NULL on error
*/
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
/*
  Add a new side data to a frame from an existing AVBufferRef

  @param frame a frame to which the side data should be added
  @param type  the type of the added side data
  @param buf   an AVBufferRef to add as side data. The ownership of
               the reference is transferred to the frame.

  @return newly added side data on success, NULL on error. On failure
          the frame is unchanged and the AVBufferRef remains owned by
          the caller.
*/
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
/*
  @return a pointer to the side data of a given type on success, NULL if there
  is no side data with such type in this frame.
*/
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
//
//	Remove and free all side data instances of the given type.
func AVFrameRemoveSideData(frame *AVFrame, _type AVFrameSideDataType) {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	C.av_frame_remove_side_data(tmpframe, C.enum_AVFrameSideDataType(_type))
}

// --- Function av_frame_apply_cropping ---

// AVFrameApplyCropping wraps av_frame_apply_cropping.
/*
  Crop the given video AVFrame according to its crop_left/crop_top/crop_right/
  crop_bottom fields. If cropping is successful, the function will adjust the
  data pointers and the width/height fields, and set the crop fields to 0.

  In all cases, the cropping boundaries will be rounded to the inherent
  alignment of the pixel format. In some cases, such as for opaque hwaccel
  formats, the left/top cropping is ignored. The crop fields are set to 0 even
  if the cropping was rounded or ignored.

  @param frame the frame which should be cropped
  @param flags Some combination of AV_FRAME_CROP_* flags, or 0.

  @return >= 0 on success, a negative AVERROR on error. If the cropping fields
  were invalid, AVERROR(ERANGE) is returned, and nothing is changed.
*/
func AVFrameApplyCropping(frame *AVFrame, flags int) (int, error) {
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_frame_apply_cropping(tmpframe, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_frame_side_data_name ---

// AVFrameSideDataName wraps av_frame_side_data_name.
//
//	@return a string identifying the side data type
func AVFrameSideDataName(_type AVFrameSideDataType) *CStr {
	ret := C.av_frame_side_data_name(C.enum_AVFrameSideDataType(_type))
	return wrapCStr(ret)
}

// --- Function av_hwdevice_find_type_by_name ---

// AVHwdeviceFindTypeByName wraps av_hwdevice_find_type_by_name.
/*
  Look up an AVHWDeviceType by name.

  @param name String name of the device type (case-insensitive).
  @return The type from enum AVHWDeviceType, or AV_HWDEVICE_TYPE_NONE if
          not found.
*/
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
/*
  Get the string name of an AVHWDeviceType.

  @param type Type from enum AVHWDeviceType.
  @return Pointer to a static string containing the name, or NULL if the type
          is not valid.
*/
func AVHwdeviceGetTypeName(_type AVHWDeviceType) *CStr {
	ret := C.av_hwdevice_get_type_name(C.enum_AVHWDeviceType(_type))
	return wrapCStr(ret)
}

// --- Function av_hwdevice_iterate_types ---

// AVHwdeviceIterateTypes wraps av_hwdevice_iterate_types.
/*
  Iterate over supported device types.

  @param prev AV_HWDEVICE_TYPE_NONE initially, then the previous type
              returned by this function in subsequent iterations.
  @return The next usable device type from enum AVHWDeviceType, or
          AV_HWDEVICE_TYPE_NONE if there are no more.
*/
func AVHwdeviceIterateTypes(prev AVHWDeviceType) AVHWDeviceType {
	ret := C.av_hwdevice_iterate_types(C.enum_AVHWDeviceType(prev))
	return AVHWDeviceType(ret)
}

// --- Function av_hwdevice_ctx_alloc ---

// AVHwdeviceCtxAlloc wraps av_hwdevice_ctx_alloc.
/*
  Allocate an AVHWDeviceContext for a given hardware type.

  @param type the type of the hardware device to allocate.
  @return a reference to the newly created AVHWDeviceContext on success or NULL
          on failure.
*/
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
/*
  Finalize the device context before use. This function must be called after
  the context is filled with all the required information and before it is
  used in any way.

  @param ref a reference to the AVHWDeviceContext
  @return 0 on success, a negative AVERROR code on failure
*/
func AVHwdeviceCtxInit(ref *AVBufferRef) (int, error) {
	var tmpref *C.AVBufferRef
	if ref != nil {
		tmpref = ref.ptr
	}
	ret := C.av_hwdevice_ctx_init(tmpref)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_hwdevice_ctx_create ---

// AVHwdeviceCtxCreate wraps av_hwdevice_ctx_create.
/*
  Open a device of the specified type and create an AVHWDeviceContext for it.

  This is a convenience function intended to cover the simple cases. Callers
  who need to fine-tune device creation/management should open the device
  manually and then wrap it in an AVHWDeviceContext using
  av_hwdevice_ctx_alloc()/av_hwdevice_ctx_init().

  The returned context is already initialized and ready for use, the caller
  should not call av_hwdevice_ctx_init() on it. The user_opaque/free fields of
  the created AVHWDeviceContext are set by this function and should not be
  touched by the caller.

  @param device_ctx On success, a reference to the newly-created device context
                    will be written here. The reference is owned by the caller
                    and must be released with av_buffer_unref() when no longer
                    needed. On failure, NULL will be written to this pointer.
  @param type The type of the device to create.
  @param device A type-specific string identifying the device to open.
  @param opts A dictionary of additional (type-specific) options to use in
              opening the device. The dictionary remains owned by the caller.
  @param flags currently unused

  @return 0 on success, a negative AVERROR code on failure.
*/
func AVHwdeviceCtxCreate(deviceCtx **AVBufferRef, _type AVHWDeviceType, device *CStr, opts *AVDictionary, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_hwdevice_ctx_create_derived ---

// AVHwdeviceCtxCreateDerived wraps av_hwdevice_ctx_create_derived.
/*
  Create a new device of the specified type from an existing device.

  If the source device is a device of the target type or was originally
  derived from such a device (possibly through one or more intermediate
  devices of other types), then this will return a reference to the
  existing device of the same type as is requested.

  Otherwise, it will attempt to derive a new device from the given source
  device.  If direct derivation to the new type is not implemented, it will
  attempt the same derivation from each ancestor of the source device in
  turn looking for an implemented derivation method.

  @param dst_ctx On success, a reference to the newly-created
                 AVHWDeviceContext.
  @param type    The type of the new device to create.
  @param src_ctx A reference to an existing AVHWDeviceContext which will be
                 used to create the new device.
  @param flags   Currently unused; should be set to zero.
  @return        Zero on success, a negative AVERROR code on failure.
*/
func AVHwdeviceCtxCreateDerived(dstCtx **AVBufferRef, _type AVHWDeviceType, srcCtx *AVBufferRef, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_hwdevice_ctx_create_derived_opts ---

// AVHwdeviceCtxCreateDerivedOpts wraps av_hwdevice_ctx_create_derived_opts.
/*
  Create a new device of the specified type from an existing device.

  This function performs the same action as av_hwdevice_ctx_create_derived,
  however, it is able to set options for the new device to be derived.

  @param dst_ctx On success, a reference to the newly-created
                 AVHWDeviceContext.
  @param type    The type of the new device to create.
  @param src_ctx A reference to an existing AVHWDeviceContext which will be
                 used to create the new device.
  @param options Options for the new device to create, same format as in
                 av_hwdevice_ctx_create.
  @param flags   Currently unused; should be set to zero.
  @return        Zero on success, a negative AVERROR code on failure.
*/
func AVHwdeviceCtxCreateDerivedOpts(dstCtx **AVBufferRef, _type AVHWDeviceType, srcCtx *AVBufferRef, options *AVDictionary, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_hwframe_ctx_alloc ---

// AVHwframeCtxAlloc wraps av_hwframe_ctx_alloc.
/*
  Allocate an AVHWFramesContext tied to a given device context.

  @param device_ctx a reference to a AVHWDeviceContext. This function will make
                    a new reference for internal use, the one passed to the
                    function remains owned by the caller.
  @return a reference to the newly created AVHWFramesContext on success or NULL
          on failure.
*/
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
/*
  Finalize the context before use. This function must be called after the
  context is filled with all the required information and before it is attached
  to any frames.

  @param ref a reference to the AVHWFramesContext
  @return 0 on success, a negative AVERROR code on failure
*/
func AVHwframeCtxInit(ref *AVBufferRef) (int, error) {
	var tmpref *C.AVBufferRef
	if ref != nil {
		tmpref = ref.ptr
	}
	ret := C.av_hwframe_ctx_init(tmpref)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_hwframe_get_buffer ---

// AVHwframeGetBuffer wraps av_hwframe_get_buffer.
/*
  Allocate a new frame attached to the given AVHWFramesContext.

  @param hwframe_ctx a reference to an AVHWFramesContext
  @param frame an empty (freshly allocated or unreffed) frame to be filled with
               newly allocated buffers.
  @param flags currently unused, should be set to zero
  @return 0 on success, a negative AVERROR code on failure
*/
func AVHwframeGetBuffer(hwframeCtx *AVBufferRef, frame *AVFrame, flags int) (int, error) {
	var tmphwframeCtx *C.AVBufferRef
	if hwframeCtx != nil {
		tmphwframeCtx = hwframeCtx.ptr
	}
	var tmpframe *C.AVFrame
	if frame != nil {
		tmpframe = frame.ptr
	}
	ret := C.av_hwframe_get_buffer(tmphwframeCtx, tmpframe, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_hwframe_transfer_data ---

// AVHwframeTransferData wraps av_hwframe_transfer_data.
/*
  Copy data to or from a hw surface. At least one of dst/src must have an
  AVHWFramesContext attached.

  If src has an AVHWFramesContext attached, then the format of dst (if set)
  must use one of the formats returned by av_hwframe_transfer_get_formats(src,
  AV_HWFRAME_TRANSFER_DIRECTION_FROM).
  If dst has an AVHWFramesContext attached, then the format of src must use one
  of the formats returned by av_hwframe_transfer_get_formats(dst,
  AV_HWFRAME_TRANSFER_DIRECTION_TO)

  dst may be "clean" (i.e. with data/buf pointers unset), in which case the
  data buffers will be allocated by this function using av_frame_get_buffer().
  If dst->format is set, then this format will be used, otherwise (when
  dst->format is AV_PIX_FMT_NONE) the first acceptable format will be chosen.

  The two frames must have matching allocated dimensions (i.e. equal to
  AVHWFramesContext.width/height), since not all device types support
  transferring a sub-rectangle of the whole surface. The display dimensions
  (i.e. AVFrame.width/height) may be smaller than the allocated dimensions, but
  also have to be equal for both frames. When the display dimensions are
  smaller than the allocated dimensions, the content of the padding in the
  destination frame is unspecified.

  @param dst the destination frame. dst is not touched on failure.
  @param src the source frame.
  @param flags currently unused, should be set to zero
  @return 0 on success, a negative AVERROR error code on failure.
*/
func AVHwframeTransferData(dst *AVFrame, src *AVFrame, flags int) (int, error) {
	var tmpdst *C.AVFrame
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_hwframe_transfer_data(tmpdst, tmpsrc, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_hwframe_transfer_get_formats ---

// av_hwframe_transfer_get_formats skipped due to formats

// --- Function av_hwdevice_hwconfig_alloc ---

// AVHwdeviceHwconfigAlloc wraps av_hwdevice_hwconfig_alloc.
/*
  Allocate a HW-specific configuration structure for a given HW device.
  After use, the user must free all members as required by the specific
  hardware structure being used, then free the structure itself with
  av_free().

  @param device_ctx a reference to the associated AVHWDeviceContext.
  @return The newly created HW-specific configuration structure on
          success or NULL on failure.
*/
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
/*
  Get the constraints on HW frames given a device and the HW-specific
  configuration to be used with that device.  If no HW-specific
  configuration is provided, returns the maximum possible capabilities
  of the device.

  @param ref a reference to the associated AVHWDeviceContext.
  @param hwconfig a filled HW-specific configuration structure, or NULL
         to return the maximum possible capabilities of the device.
  @return AVHWFramesConstraints structure describing the constraints
          on the device, or NULL if not available.
*/
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
/*
  Free an AVHWFrameConstraints structure.

  @param constraints The (filled or unfilled) AVHWFrameConstraints structure.
*/
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
/*
  Map a hardware frame.

  This has a number of different possible effects, depending on the format
  and origin of the src and dst frames.  On input, src should be a usable
  frame with valid buffers and dst should be blank (typically as just created
  by av_frame_alloc()).  src should have an associated hwframe context, and
  dst may optionally have a format and associated hwframe context.

  If src was created by mapping a frame from the hwframe context of dst,
  then this function undoes the mapping - dst is replaced by a reference to
  the frame that src was originally mapped from.

  If both src and dst have an associated hwframe context, then this function
  attempts to map the src frame from its hardware context to that of dst and
  then fill dst with appropriate data to be usable there.  This will only be
  possible if the hwframe contexts and associated devices are compatible -
  given compatible devices, av_hwframe_ctx_create_derived() can be used to
  create a hwframe context for dst in which mapping should be possible.

  If src has a hwframe context but dst does not, then the src frame is
  mapped to normal memory and should thereafter be usable as a normal frame.
  If the format is set on dst, then the mapping will attempt to create dst
  with that format and fail if it is not possible.  If format is unset (is
  AV_PIX_FMT_NONE) then dst will be mapped with whatever the most appropriate
  format to use is (probably the sw_format of the src hwframe context).

  A return value of AVERROR(ENOSYS) indicates that the mapping is not
  possible with the given arguments and hwframe setup, while other return
  values indicate that it failed somehow.

  On failure, the destination frame will be left blank, except for the
  hw_frames_ctx/format fields thay may have been set by the caller - those will
  be preserved as they were.

  @param dst Destination frame, to contain the mapping.
  @param src Source frame, to be mapped.
  @param flags Some combination of AV_HWFRAME_MAP_* flags.
  @return Zero on success, negative AVERROR code on failure.
*/
func AVHwframeMap(dst *AVFrame, src *AVFrame, flags int) (int, error) {
	var tmpdst *C.AVFrame
	if dst != nil {
		tmpdst = dst.ptr
	}
	var tmpsrc *C.AVFrame
	if src != nil {
		tmpsrc = src.ptr
	}
	ret := C.av_hwframe_map(tmpdst, tmpsrc, C.int(flags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_hwframe_ctx_create_derived ---

// AVHwframeCtxCreateDerived wraps av_hwframe_ctx_create_derived.
/*
  Create and initialise an AVHWFramesContext as a mapping of another existing
  AVHWFramesContext on a different device.

  av_hwframe_ctx_init() should not be called after this.

  @param derived_frame_ctx  On success, a reference to the newly created
                            AVHWFramesContext.
  @param format             The AVPixelFormat for the derived context.
  @param derived_device_ctx A reference to the device to create the new
                            AVHWFramesContext on.
  @param source_frame_ctx   A reference to an existing AVHWFramesContext
                            which will be mapped to the derived context.
  @param flags  Some combination of AV_HWFRAME_MAP_* flags, defining the
                mapping parameters to apply to frames which are allocated
                in the derived device.
  @return       Zero on success, negative AVERROR code on failure.
*/
func AVHwframeCtxCreateDerived(derivedFrameCtx **AVBufferRef, format AVPixelFormat, derivedDeviceCtx *AVBufferRef, sourceFrameCtx *AVBufferRef, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_log ---

// av_log skipped due to variadic arg.

// --- Function av_log_once ---

// av_log_once skipped due to variadic arg.

// --- Function av_vlog ---

// av_vlog skipped due to vl.

// --- Function av_log_get_level ---

// AVLogGetLevel wraps av_log_get_level.
/*
  Get the current log level

  @see lavu_log_constants

  @return Current log level
*/
func AVLogGetLevel() (int, error) {
	ret := C.av_log_get_level()
	return int(ret), WrapErr(int(ret))
}

// --- Function av_log_set_level ---

// AVLogSetLevel wraps av_log_set_level.
/*
  Set the log level

  @see lavu_log_constants

  @param level Logging level
*/
func AVLogSetLevel(level int) {
	C.av_log_set_level(C.int(level))
}

// --- Function av_log_set_callback ---

// av_log_set_callback skipped due to callback.

// --- Function av_log_default_callback ---

// av_log_default_callback skipped due to vl.

// --- Function av_default_item_name ---

// AVDefaultItemName wraps av_default_item_name.
/*
  Return the context name

  @param  ctx The AVClass context

  @return The AVClass class_name
*/
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
func AVLogGetFlags() (int, error) {
	ret := C.av_log_get_flags()
	return int(ret), WrapErr(int(ret))
}

// --- Function av_gcd ---

// AVGcd wraps av_gcd.
/*
  Compute the greatest common divisor of two integer operands.

  @param a Operand
  @param b Operand
  @return GCD of a and b up to sign; if a >= 0 and b >= 0, return value is >= 0;
  if a == 0 and b == 0, returns 0.
*/
func AVGcd(a int64, b int64) int64 {
	ret := C.av_gcd(C.int64_t(a), C.int64_t(b))
	return int64(ret)
}

// --- Function av_rescale ---

// AVRescale wraps av_rescale.
/*
  Rescale a 64-bit integer with rounding to nearest.

  The operation is mathematically equivalent to `a * b / c`, but writing that
  directly can overflow.

  This function is equivalent to av_rescale_rnd() with #AV_ROUND_NEAR_INF.

  @see av_rescale_rnd(), av_rescale_q(), av_rescale_q_rnd()
*/
func AVRescale(a int64, b int64, c int64) int64 {
	ret := C.av_rescale(C.int64_t(a), C.int64_t(b), C.int64_t(c))
	return int64(ret)
}

// --- Function av_rescale_rnd ---

// AVRescaleRnd wraps av_rescale_rnd.
/*
  Rescale a 64-bit integer with specified rounding.

  The operation is mathematically equivalent to `a * b / c`, but writing that
  directly can overflow, and does not support different rounding methods.
  If the result is not representable then INT64_MIN is returned.

  @see av_rescale(), av_rescale_q(), av_rescale_q_rnd()
*/
func AVRescaleRnd(a int64, b int64, c int64, rnd AVRounding) int64 {
	ret := C.av_rescale_rnd(C.int64_t(a), C.int64_t(b), C.int64_t(c), C.enum_AVRounding(rnd))
	return int64(ret)
}

// --- Function av_rescale_q ---

// AVRescaleQ wraps av_rescale_q.
/*
  Rescale a 64-bit integer by 2 rational numbers.

  The operation is mathematically equivalent to `a * bq / cq`.

  This function is equivalent to av_rescale_q_rnd() with #AV_ROUND_NEAR_INF.

  @see av_rescale(), av_rescale_rnd(), av_rescale_q_rnd()
*/
func AVRescaleQ(a int64, bq *AVRational, cq *AVRational) int64 {
	ret := C.av_rescale_q(C.int64_t(a), bq.value, cq.value)
	return int64(ret)
}

// --- Function av_rescale_q_rnd ---

// AVRescaleQRnd wraps av_rescale_q_rnd.
/*
  Rescale a 64-bit integer by 2 rational numbers with specified rounding.

  The operation is mathematically equivalent to `a * bq / cq`.

  @see av_rescale(), av_rescale_rnd(), av_rescale_q()
*/
func AVRescaleQRnd(a int64, bq *AVRational, cq *AVRational, rnd AVRounding) int64 {
	ret := C.av_rescale_q_rnd(C.int64_t(a), bq.value, cq.value, C.enum_AVRounding(rnd))
	return int64(ret)
}

// --- Function av_compare_ts ---

// AVCompareTs wraps av_compare_ts.
/*
  Compare two timestamps each in its own time base.

  @return One of the following values:
          - -1 if `ts_a` is before `ts_b`
          - 1 if `ts_a` is after `ts_b`
          - 0 if they represent the same position

  @warning
  The result of the function is undefined if one of the timestamps is outside
  the `int64_t` range when represented in the other's timebase.
*/
func AVCompareTs(tsA int64, tbA *AVRational, tsB int64, tbB *AVRational) (int, error) {
	ret := C.av_compare_ts(C.int64_t(tsA), tbA.value, C.int64_t(tsB), tbB.value)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_compare_mod ---

// AVCompareMod wraps av_compare_mod.
/*
  Compare the remainders of two integer operands divided by a common divisor.

  In other words, compare the least significant `log2(mod)` bits of integers
  `a` and `b`.

  @code{.c}
  av_compare_mod(0x11, 0x02, 0x10) < 0 // since 0x11 % 0x10  (0x1) < 0x02 % 0x10  (0x2)
  av_compare_mod(0x11, 0x02, 0x20) > 0 // since 0x11 % 0x20 (0x11) > 0x02 % 0x20 (0x02)
  @endcode

  @param a Operand
  @param b Operand
  @param mod Divisor; must be a power of 2
  @return
          - a negative value if `a % mod < b % mod`
          - a positive value if `a % mod > b % mod`
          - zero             if `a % mod == b % mod`
*/
func AVCompareMod(a uint64, b uint64, mod uint64) int64 {
	ret := C.av_compare_mod(C.uint64_t(a), C.uint64_t(b), C.uint64_t(mod))
	return int64(ret)
}

// --- Function av_rescale_delta ---

// av_rescale_delta skipped due to last

// --- Function av_add_stable ---

// AVAddStable wraps av_add_stable.
/*
  Add a value to a timestamp.

  This function guarantees that when the same value is repeatly added that
  no accumulation of rounding errors occurs.

  @param[in] ts     Input timestamp
  @param[in] ts_tb  Input timestamp time base
  @param[in] inc    Value to be added
  @param[in] inc_tb Time base of `inc`
*/
func AVAddStable(tsTb *AVRational, ts int64, incTb *AVRational, inc int64) int64 {
	ret := C.av_add_stable(tsTb.value, C.int64_t(ts), incTb.value, C.int64_t(inc))
	return int64(ret)
}

// --- Function av_malloc ---

// AVMalloc wraps av_malloc.
/*
  Allocate a memory block with alignment suitable for all memory accesses
  (including vectors if available on the CPU).

  @param size Size in bytes for the memory block to be allocated
  @return Pointer to the allocated block, or `NULL` if the block cannot
          be allocated
  @see av_mallocz()
*/
func AVMalloc(size uint64) unsafe.Pointer {
	ret := C.av_malloc(C.size_t(size))
	return ret
}

// --- Function av_mallocz ---

// AVMallocz wraps av_mallocz.
/*
  Allocate a memory block with alignment suitable for all memory accesses
  (including vectors if available on the CPU) and zero all the bytes of the
  block.

  @param size Size in bytes for the memory block to be allocated
  @return Pointer to the allocated block, or `NULL` if it cannot be allocated
  @see av_malloc()
*/
func AVMallocz(size uint64) unsafe.Pointer {
	ret := C.av_mallocz(C.size_t(size))
	return ret
}

// --- Function av_malloc_array ---

// AVMallocArray wraps av_malloc_array.
/*
  Allocate a memory block for an array with av_malloc().

  The allocated memory will have size `size * nmemb` bytes.

  @param nmemb Number of element
  @param size  Size of a single element
  @return Pointer to the allocated block, or `NULL` if the block cannot
          be allocated
  @see av_malloc()
*/
func AVMallocArray(nmemb uint64, size uint64) unsafe.Pointer {
	ret := C.av_malloc_array(C.size_t(nmemb), C.size_t(size))
	return ret
}

// --- Function av_calloc ---

// AVCalloc wraps av_calloc.
/*
  Allocate a memory block for an array with av_mallocz().

  The allocated memory will have size `size * nmemb` bytes.

  @param nmemb Number of elements
  @param size  Size of the single element
  @return Pointer to the allocated block, or `NULL` if the block cannot
          be allocated

  @see av_mallocz()
  @see av_malloc_array()
*/
func AVCalloc(nmemb uint64, size uint64) unsafe.Pointer {
	ret := C.av_calloc(C.size_t(nmemb), C.size_t(size))
	return ret
}

// --- Function av_realloc ---

// AVRealloc wraps av_realloc.
/*
  Allocate, reallocate, or free a block of memory.

  If `ptr` is `NULL` and `size` > 0, allocate a new block. Otherwise, expand or
  shrink that block of memory according to `size`.

  @param ptr  Pointer to a memory block already allocated with
              av_realloc() or `NULL`
  @param size Size in bytes of the memory block to be allocated or
              reallocated

  @return Pointer to a newly-reallocated block or `NULL` if the block
          cannot be reallocated

  @warning Unlike av_malloc(), the returned pointer is not guaranteed to be
           correctly aligned. The returned pointer must be freed after even
           if size is zero.
  @see av_fast_realloc()
  @see av_reallocp()
*/
func AVRealloc(ptr unsafe.Pointer, size uint64) unsafe.Pointer {
	ret := C.av_realloc(ptr, C.size_t(size))
	return ret
}

// --- Function av_reallocp ---

// AVReallocp wraps av_reallocp.
/*
  Allocate, reallocate, or free a block of memory through a pointer to a
  pointer.

  If `*ptr` is `NULL` and `size` > 0, allocate a new block. If `size` is
  zero, free the memory block pointed to by `*ptr`. Otherwise, expand or
  shrink that block of memory according to `size`.

  @param[in,out] ptr  Pointer to a pointer to a memory block already allocated
                      with av_realloc(), or a pointer to `NULL`. The pointer
                      is updated on success, or freed on failure.
  @param[in]     size Size in bytes for the memory block to be allocated or
                      reallocated

  @return Zero on success, an AVERROR error code on failure

  @warning Unlike av_malloc(), the allocated memory is not guaranteed to be
           correctly aligned.
*/
func AVReallocp(ptr unsafe.Pointer, size uint64) (int, error) {
	ret := C.av_reallocp(ptr, C.size_t(size))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_realloc_f ---

// AVReallocF wraps av_realloc_f.
/*
  Allocate, reallocate, or free a block of memory.

  This function does the same thing as av_realloc(), except:
  - It takes two size arguments and allocates `nelem * elsize` bytes,
    after checking the result of the multiplication for integer overflow.
  - It frees the input block in case of failure, thus avoiding the memory
    leak with the classic
    @code{.c}
    buf = realloc(buf);
    if (!buf)
        return -1;
    @endcode
    pattern.
*/
func AVReallocF(ptr unsafe.Pointer, nelem uint64, elsize uint64) unsafe.Pointer {
	ret := C.av_realloc_f(ptr, C.size_t(nelem), C.size_t(elsize))
	return ret
}

// --- Function av_realloc_array ---

// AVReallocArray wraps av_realloc_array.
/*
  Allocate, reallocate, or free an array.

  If `ptr` is `NULL` and `nmemb` > 0, allocate a new block.

  @param ptr   Pointer to a memory block already allocated with
               av_realloc() or `NULL`
  @param nmemb Number of elements in the array
  @param size  Size of the single element of the array

  @return Pointer to a newly-reallocated block or NULL if the block
          cannot be reallocated

  @warning Unlike av_malloc(), the allocated memory is not guaranteed to be
           correctly aligned. The returned pointer must be freed after even if
           nmemb is zero.
  @see av_reallocp_array()
*/
func AVReallocArray(ptr unsafe.Pointer, nmemb uint64, size uint64) unsafe.Pointer {
	ret := C.av_realloc_array(ptr, C.size_t(nmemb), C.size_t(size))
	return ret
}

// --- Function av_reallocp_array ---

// AVReallocpArray wraps av_reallocp_array.
/*
  Allocate, reallocate an array through a pointer to a pointer.

  If `*ptr` is `NULL` and `nmemb` > 0, allocate a new block.

  @param[in,out] ptr   Pointer to a pointer to a memory block already
                       allocated with av_realloc(), or a pointer to `NULL`.
                       The pointer is updated on success, or freed on failure.
  @param[in]     nmemb Number of elements
  @param[in]     size  Size of the single element

  @return Zero on success, an AVERROR error code on failure

  @warning Unlike av_malloc(), the allocated memory is not guaranteed to be
           correctly aligned. *ptr must be freed after even if nmemb is zero.
*/
func AVReallocpArray(ptr unsafe.Pointer, nmemb uint64, size uint64) (int, error) {
	ret := C.av_reallocp_array(ptr, C.size_t(nmemb), C.size_t(size))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_fast_realloc ---

// av_fast_realloc skipped due to size

// --- Function av_fast_malloc ---

// av_fast_malloc skipped due to size

// --- Function av_fast_mallocz ---

// av_fast_mallocz skipped due to size

// --- Function av_free ---

// AVFree wraps av_free.
/*
  Free a memory block which has been allocated with a function of av_malloc()
  or av_realloc() family.

  @param ptr Pointer to the memory block which should be freed.

  @note `ptr = NULL` is explicitly allowed.
  @note It is recommended that you use av_freep() instead, to prevent leaving
        behind dangling pointers.
  @see av_freep()
*/
func AVFree(ptr unsafe.Pointer) {
	C.av_free(ptr)
}

// --- Function av_freep ---

// AVFreep wraps av_freep.
/*
  Free a memory block which has been allocated with a function of av_malloc()
  or av_realloc() family, and set the pointer pointing to it to `NULL`.

  @code{.c}
  uint8_t *buf = av_malloc(16);
  av_free(buf);
  // buf now contains a dangling pointer to freed memory, and accidental
  // dereference of buf will result in a use-after-free, which may be a
  // security risk.

  uint8_t *buf = av_malloc(16);
  av_freep(&buf);
  // buf is now NULL, and accidental dereference will only result in a
  // NULL-pointer dereference.
  @endcode

  @param ptr Pointer to the pointer to the memory block which should be freed
  @note `*ptr = NULL` is safe and leads to no action.
  @see av_free()
*/
func AVFreep(ptr unsafe.Pointer) {
	C.av_freep(ptr)
}

// --- Function av_strdup ---

// AVStrdup wraps av_strdup.
/*
  Duplicate a string.

  @param s String to be duplicated
  @return Pointer to a newly-allocated string containing a
          copy of `s` or `NULL` if the string cannot be allocated
  @see av_strndup()
*/
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
/*
  Duplicate a substring of a string.

  @param s   String to be duplicated
  @param len Maximum length of the resulting string (not counting the
             terminating byte)
  @return Pointer to a newly-allocated string containing a
          substring of `s` or `NULL` if the string cannot be allocated
*/
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
/*
  Duplicate a buffer with av_malloc().

  @param p    Buffer to be duplicated
  @param size Size in bytes of the buffer copied
  @return Pointer to a newly allocated buffer containing a
          copy of `p` or `NULL` if the buffer cannot be allocated
*/
func AVMemdup(p unsafe.Pointer, size uint64) unsafe.Pointer {
	ret := C.av_memdup(p, C.size_t(size))
	return ret
}

// --- Function av_memcpy_backptr ---

// AVMemcpyBackptr wraps av_memcpy_backptr.
/*
  Overlapping memcpy() implementation.

  @param dst  Destination buffer
  @param back Number of bytes back to start copying (i.e. the initial size of
              the overlapping window); must be > 0
  @param cnt  Number of bytes to copy; must be >= 0

  @note `cnt > back` is valid, this will copy the bytes we just copied,
        thus creating a repeating pattern with a period length of `back`.
*/
func AVMemcpyBackptr(dst unsafe.Pointer, back int, cnt int) {
	C.av_memcpy_backptr((*C.uint8_t)(dst), C.int(back), C.int(cnt))
}

// --- Function av_dynarray_add ---

// av_dynarray_add skipped due to nbPtr

// --- Function av_dynarray_add_nofree ---

// av_dynarray_add_nofree skipped due to nbPtr

// --- Function av_dynarray2_add ---

// av_dynarray2_add skipped due to tabPtr

// --- Function av_size_mult ---

// av_size_mult skipped due to r

// --- Function av_max_alloc ---

// AVMaxAlloc wraps av_max_alloc.
/*
  Set the maximum size that may be allocated in one block.

  The value specified with this function is effective for all libavutil's @ref
  lavu_mem_funcs "heap management functions."

  By default, the max value is defined as `INT_MAX`.

  @param max Value to be set as the new maximum size

  @warning Exercise extreme caution when using this function. Don't touch
           this if you do not understand the full consequence of doing so.
*/
func AVMaxAlloc(max uint64) {
	C.av_max_alloc(C.size_t(max))
}

// --- Function av_opt_show2 ---

// AVOptShow2 wraps av_opt_show2.
/*
  Show the obj options.

  @param req_flags requested flags for the options to show. Show only the
  options for which it is opt->flags & req_flags.
  @param rej_flags rejected flags for the options to show. Show only the
  options for which it is !(opt->flags & req_flags).
  @param av_log_obj log context to use for showing the options
*/
func AVOptShow2(obj unsafe.Pointer, avLogObj unsafe.Pointer, reqFlags int, rejFlags int) (int, error) {
	ret := C.av_opt_show2(obj, avLogObj, C.int(reqFlags), C.int(rejFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_defaults ---

// AVOptSetDefaults wraps av_opt_set_defaults.
/*
  Set the values of all AVOption fields to their default values.

  @param s an AVOption-enabled struct (its first member must be a pointer to AVClass)
*/
func AVOptSetDefaults(s unsafe.Pointer) {
	C.av_opt_set_defaults(s)
}

// --- Function av_opt_set_defaults2 ---

// AVOptSetDefaults2 wraps av_opt_set_defaults2.
/*
  Set the values of all AVOption fields to their default values. Only these
  AVOption fields for which (opt->flags & mask) == flags will have their
  default applied to s.

  @param s an AVOption-enabled struct (its first member must be a pointer to AVClass)
  @param mask combination of AV_OPT_FLAG_*
  @param flags combination of AV_OPT_FLAG_*
*/
func AVOptSetDefaults2(s unsafe.Pointer, mask int, flags int) {
	C.av_opt_set_defaults2(s, C.int(mask), C.int(flags))
}

// --- Function av_set_options_string ---

// AVSetOptionsString wraps av_set_options_string.
/*
  Parse the key/value pairs list in opts. For each key/value pair
  found, stores the value in the field in ctx that is named like the
  key. ctx must be an AVClass context, storing is done using
  AVOptions.

  @param opts options string to parse, may be NULL
  @param key_val_sep a 0-terminated list of characters used to
  separate key from value
  @param pairs_sep a 0-terminated list of characters used to separate
  two pairs from each other
  @return the number of successfully set key/value pairs, or a negative
  value corresponding to an AVERROR code in case of error:
  AVERROR(EINVAL) if opts cannot be parsed,
  the error code issued by av_opt_set() if a key/value pair
  cannot be set
*/
func AVSetOptionsString(ctx unsafe.Pointer, opts *CStr, keyValSep *CStr, pairsSep *CStr) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_from_string ---

// av_opt_set_from_string skipped due to shorthand

// --- Function av_opt_free ---

// AVOptFree wraps av_opt_free.
//
//	Free all allocated objects in obj.
func AVOptFree(obj unsafe.Pointer) {
	C.av_opt_free(obj)
}

// --- Function av_opt_flag_is_set ---

// AVOptFlagIsSet wraps av_opt_flag_is_set.
/*
  Check whether a particular flag is set in a flags field.

  @param field_name the name of the flag field option
  @param flag_name the name of the flag to check
  @return non-zero if the flag is set, zero if the flag isn't set,
          isn't of the right type, or the flags field doesn't exist.
*/
func AVOptFlagIsSet(obj unsafe.Pointer, fieldName *CStr, flagName *CStr) (int, error) {
	var tmpfieldName *C.char
	if fieldName != nil {
		tmpfieldName = fieldName.ptr
	}
	var tmpflagName *C.char
	if flagName != nil {
		tmpflagName = flagName.ptr
	}
	ret := C.av_opt_flag_is_set(obj, tmpfieldName, tmpflagName)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_dict ---

// AVOptSetDict wraps av_opt_set_dict.
/*
  Set all the options from a given dictionary on an object.

  @param obj a struct whose first element is a pointer to AVClass
  @param options options to process. This dictionary will be freed and replaced
                 by a new one containing all options not found in obj.
                 Of course this new dictionary needs to be freed by caller
                 with av_dict_free().

  @return 0 on success, a negative AVERROR if some option was found in obj,
          but could not be set.

  @see av_dict_copy()
*/
func AVOptSetDict(obj unsafe.Pointer, options **AVDictionary) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_dict2 ---

// AVOptSetDict2 wraps av_opt_set_dict2.
/*
  Set all the options from a given dictionary on an object.

  @param obj a struct whose first element is a pointer to AVClass
  @param options options to process. This dictionary will be freed and replaced
                 by a new one containing all options not found in obj.
                 Of course this new dictionary needs to be freed by caller
                 with av_dict_free().
  @param search_flags A combination of AV_OPT_SEARCH_*.

  @return 0 on success, a negative AVERROR if some option was found in obj,
          but could not be set.

  @see av_dict_copy()
*/
func AVOptSetDict2(obj unsafe.Pointer, options **AVDictionary, searchFlags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_get_key_value ---

// av_opt_get_key_value skipped due to ropts

// --- Function av_opt_eval_flags ---

// av_opt_eval_flags skipped due to flagsOut

// --- Function av_opt_eval_int ---

// av_opt_eval_int skipped due to intOut

// --- Function av_opt_eval_int64 ---

// av_opt_eval_int64 skipped due to int64Out

// --- Function av_opt_eval_float ---

// av_opt_eval_float skipped due to floatOut

// --- Function av_opt_eval_double ---

// av_opt_eval_double skipped due to doubleOut

// --- Function av_opt_eval_q ---

// av_opt_eval_q skipped due to qOut

// --- Function av_opt_find ---

// AVOptFind wraps av_opt_find.
/*
  Look for an option in an object. Consider only options which
  have all the specified flags set.

  @param[in] obj A pointer to a struct whose first element is a
                 pointer to an AVClass.
                 Alternatively a double pointer to an AVClass, if
                 AV_OPT_SEARCH_FAKE_OBJ search flag is set.
  @param[in] name The name of the option to look for.
  @param[in] unit When searching for named constants, name of the unit
                  it belongs to.
  @param opt_flags Find only options with all the specified flags set (AV_OPT_FLAG).
  @param search_flags A combination of AV_OPT_SEARCH_*.

  @return A pointer to the option found, or NULL if no option
          was found.

  @note Options found with AV_OPT_SEARCH_CHILDREN flag may not be settable
  directly with av_opt_set(). Use special calls which take an options
  AVDictionary (e.g. avformat_open_input()) to set options found with this
  flag.
*/
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

// av_opt_find2 skipped due to targetObj

// --- Function av_opt_next ---

// AVOptNext wraps av_opt_next.
/*
  Iterate over all AVOptions belonging to obj.

  @param obj an AVOptions-enabled struct or a double pointer to an
             AVClass describing it.
  @param prev result of the previous call to av_opt_next() on this object
              or NULL
  @return next AVOption or NULL
*/
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
/*
  Iterate over AVOptions-enabled children of obj.

  @param prev result of a previous call to this function or NULL
  @return next AVOptions-enabled child or NULL
*/
func AVOptChildNext(obj unsafe.Pointer, prev unsafe.Pointer) unsafe.Pointer {
	ret := C.av_opt_child_next(obj, prev)
	return ret
}

// --- Function av_opt_child_class_iterate ---

// av_opt_child_class_iterate skipped due to iter

// --- Function av_opt_set ---

// AVOptSet wraps av_opt_set.
/*
  Those functions set the field of obj with the given name to value.

  @param[in] obj A struct whose first element is a pointer to an AVClass.
  @param[in] name the name of the field to set
  @param[in] val The value to set. In case of av_opt_set() if the field is not
  of a string type, then the given string is parsed.
  SI postfixes and some named scalars are supported.
  If the field is of a numeric type, it has to be a numeric or named
  scalar. Behavior with more than one scalar and +- infix operators
  is undefined.
  If the field is of a flags type, it has to be a sequence of numeric
  scalars or named flags separated by '+' or '-'. Prefixing a flag
  with '+' causes it to be set without affecting the other flags;
  similarly, '-' unsets a flag.
  If the field is of a dictionary type, it has to be a ':' separated list of
  key=value parameters. Values containing ':' special characters must be
  escaped.
  @param search_flags flags passed to av_opt_find2. I.e. if AV_OPT_SEARCH_CHILDREN
  is passed here, then the option may be set on a child of obj.

  @return 0 if the value has been set, or an AVERROR code in case of
  error:
  AVERROR_OPTION_NOT_FOUND if no matching option exists
  AVERROR(ERANGE) if the value is out of range
  AVERROR(EINVAL) if the value is not valid
*/
func AVOptSet(obj unsafe.Pointer, name *CStr, val *CStr, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	var tmpval *C.char
	if val != nil {
		tmpval = val.ptr
	}
	ret := C.av_opt_set(obj, tmpname, tmpval, C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_int ---

// AVOptSetInt wraps av_opt_set_int.
func AVOptSetInt(obj unsafe.Pointer, name *CStr, val int64, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_int(obj, tmpname, C.int64_t(val), C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_double ---

// AVOptSetDouble wraps av_opt_set_double.
func AVOptSetDouble(obj unsafe.Pointer, name *CStr, val float64, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_double(obj, tmpname, C.double(val), C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_q ---

// AVOptSetQ wraps av_opt_set_q.
func AVOptSetQ(obj unsafe.Pointer, name *CStr, val *AVRational, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_q(obj, tmpname, val.value, C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_bin ---

// AVOptSetBin wraps av_opt_set_bin.
func AVOptSetBin(obj unsafe.Pointer, name *CStr, val unsafe.Pointer, size int, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_bin(obj, tmpname, (*C.uint8_t)(val), C.int(size), C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_image_size ---

// AVOptSetImageSize wraps av_opt_set_image_size.
func AVOptSetImageSize(obj unsafe.Pointer, name *CStr, w int, h int, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_image_size(obj, tmpname, C.int(w), C.int(h), C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_pixel_fmt ---

// AVOptSetPixelFmt wraps av_opt_set_pixel_fmt.
func AVOptSetPixelFmt(obj unsafe.Pointer, name *CStr, fmt AVPixelFormat, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_pixel_fmt(obj, tmpname, C.enum_AVPixelFormat(fmt), C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_sample_fmt ---

// AVOptSetSampleFmt wraps av_opt_set_sample_fmt.
func AVOptSetSampleFmt(obj unsafe.Pointer, name *CStr, fmt AVSampleFormat, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_sample_fmt(obj, tmpname, C.enum_AVSampleFormat(fmt), C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_video_rate ---

// AVOptSetVideoRate wraps av_opt_set_video_rate.
func AVOptSetVideoRate(obj unsafe.Pointer, name *CStr, val *AVRational, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_video_rate(obj, tmpname, val.value, C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_channel_layout ---

// AVOptSetChannelLayout wraps av_opt_set_channel_layout.
func AVOptSetChannelLayout(obj unsafe.Pointer, name *CStr, chLayout int64, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_set_channel_layout(obj, tmpname, C.int64_t(chLayout), C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_chlayout ---

// AVOptSetChlayout wraps av_opt_set_chlayout.
func AVOptSetChlayout(obj unsafe.Pointer, name *CStr, layout *AVChannelLayout, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	var tmplayout *C.AVChannelLayout
	if layout != nil {
		tmplayout = layout.ptr
	}
	ret := C.av_opt_set_chlayout(obj, tmpname, tmplayout, C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_set_dict_val ---

// AVOptSetDictVal wraps av_opt_set_dict_val.
/*
  @note Any old dictionary present is discarded and replaced with a copy of the new one. The
  caller still owns val is and responsible for freeing it.
*/
func AVOptSetDictVal(obj unsafe.Pointer, name *CStr, val *AVDictionary, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	var tmpval *C.AVDictionary
	if val != nil {
		tmpval = val.ptr
	}
	ret := C.av_opt_set_dict_val(obj, tmpname, tmpval, C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_get ---

// av_opt_get skipped due to outVal

// --- Function av_opt_get_int ---

// av_opt_get_int skipped due to outVal

// --- Function av_opt_get_double ---

// av_opt_get_double skipped due to outVal

// --- Function av_opt_get_q ---

// av_opt_get_q skipped due to outVal

// --- Function av_opt_get_image_size ---

// av_opt_get_image_size skipped due to wOut

// --- Function av_opt_get_pixel_fmt ---

// av_opt_get_pixel_fmt skipped due to outFmt

// --- Function av_opt_get_sample_fmt ---

// av_opt_get_sample_fmt skipped due to outFmt

// --- Function av_opt_get_video_rate ---

// av_opt_get_video_rate skipped due to outVal

// --- Function av_opt_get_channel_layout ---

// av_opt_get_channel_layout skipped due to chLayout

// --- Function av_opt_get_chlayout ---

// AVOptGetChlayout wraps av_opt_get_chlayout.
func AVOptGetChlayout(obj unsafe.Pointer, name *CStr, searchFlags int, layout *AVChannelLayout) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	var tmplayout *C.AVChannelLayout
	if layout != nil {
		tmplayout = layout.ptr
	}
	ret := C.av_opt_get_chlayout(obj, tmpname, C.int(searchFlags), tmplayout)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_get_dict_val ---

// AVOptGetDictVal wraps av_opt_get_dict_val.
/*
  @param[out] out_val The returned dictionary is a copy of the actual value and must
  be freed with av_dict_free() by the caller
*/
func AVOptGetDictVal(obj unsafe.Pointer, name *CStr, searchFlags int, outVal **AVDictionary) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_ptr ---

// AVOptPtr wraps av_opt_ptr.
/*
  Gets a pointer to the requested field in a struct.
  This function allows accessing a struct even when its fields are moved or
  renamed since the application making the access has been compiled,

  @returns a pointer to the field, it can be cast to the correct type and read
           or written to.
*/
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
//
//	Free an AVOptionRanges struct and set it to NULL.
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
/*
  Get a list of allowed ranges for the given option.

  The returned list may depend on other fields in obj like for example profile.

  @param flags is a bitmask of flags, undefined flags should not be set and should be ignored
               AV_OPT_SEARCH_FAKE_OBJ indicates that the obj is a double pointer to a AVClass instead of a full instance
               AV_OPT_MULTI_COMPONENT_RANGE indicates that function may return more than one component, @see AVOptionRanges

  The result must be freed with av_opt_freep_ranges.

  @return number of compontents returned on success, a negative errro code otherwise
*/
func AVOptQueryRanges(param0 **AVOptionRanges, obj unsafe.Pointer, key *CStr, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_copy ---

// AVOptCopy wraps av_opt_copy.
/*
  Copy options from src object into dest object.

  The underlying AVClass of both src and dest must coincide. The guarantee
  below does not apply if this is not fulfilled.

  Options that require memory allocation (e.g. string or binary) are malloc'ed in dest object.
  Original memory allocated for such options is freed unless both src and dest options points to the same memory.

  Even on error it is guaranteed that allocated options from src and dest
  no longer alias each other afterwards; in particular calling av_opt_free()
  on both src and dest is safe afterwards if dest has been memdup'ed from src.

  @param dest Object to copy from
  @param src  Object to copy into
  @return 0 on success, negative on error
*/
func AVOptCopy(dest unsafe.Pointer, src unsafe.Pointer) (int, error) {
	ret := C.av_opt_copy(dest, src)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_query_ranges_default ---

// AVOptQueryRangesDefault wraps av_opt_query_ranges_default.
/*
  Get a default list of allowed ranges for the given option.

  This list is constructed without using the AVClass.query_ranges() callback
  and can be used as fallback from within the callback.

  @param flags is a bitmask of flags, undefined flags should not be set and should be ignored
               AV_OPT_SEARCH_FAKE_OBJ indicates that the obj is a double pointer to a AVClass instead of a full instance
               AV_OPT_MULTI_COMPONENT_RANGE indicates that function may return more than one component, @see AVOptionRanges

  The result must be freed with av_opt_free_ranges.

  @return number of compontents returned on success, a negative errro code otherwise
*/
func AVOptQueryRangesDefault(param0 **AVOptionRanges, obj unsafe.Pointer, key *CStr, flags int) (int, error) {
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
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_is_set_to_default ---

// AVOptIsSetToDefault wraps av_opt_is_set_to_default.
/*
  Check if given option is set to its default value.

  Options o must belong to the obj. This function must not be called to check child's options state.
  @see av_opt_is_set_to_default_by_name().

  @param obj  AVClass object to check option on
  @param o    option to be checked
  @return     >0 when option is set to its default,
               0 when option is not set its default,
              <0 on error
*/
func AVOptIsSetToDefault(obj unsafe.Pointer, o *AVOption) (int, error) {
	var tmpo *C.AVOption
	if o != nil {
		tmpo = o.ptr
	}
	ret := C.av_opt_is_set_to_default(obj, tmpo)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_is_set_to_default_by_name ---

// AVOptIsSetToDefaultByName wraps av_opt_is_set_to_default_by_name.
/*
  Check if given option is set to its default value.

  @param obj          AVClass object to check option on
  @param name         option name
  @param search_flags combination of AV_OPT_SEARCH_*
  @return             >0 when option is set to its default,
                      0 when option is not set its default,
                      <0 on error
*/
func AVOptIsSetToDefaultByName(obj unsafe.Pointer, name *CStr, searchFlags int) (int, error) {
	var tmpname *C.char
	if name != nil {
		tmpname = name.ptr
	}
	ret := C.av_opt_is_set_to_default_by_name(obj, tmpname, C.int(searchFlags))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_opt_serialize ---

// av_opt_serialize skipped due to buffer

// --- Function av_make_q ---

// AVMakeQ wraps av_make_q.
/*
  Create an AVRational.

  Useful for compilers that do not support compound literals.

  @note The return value is not reduced.
  @see av_reduce()
*/
func AVMakeQ(num int, den int) *AVRational {
	ret := C.av_make_q(C.int(num), C.int(den))
	return &AVRational{value: ret}
}

// --- Function av_cmp_q ---

// AVCmpQ wraps av_cmp_q.
/*
  Compare two rationals.

  @param a First rational
  @param b Second rational

  @return One of the following values:
          - 0 if `a == b`
          - 1 if `a > b`
          - -1 if `a < b`
          - `INT_MIN` if one of the values is of the form `0 / 0`
*/
func AVCmpQ(a *AVRational, b *AVRational) (int, error) {
	ret := C.av_cmp_q(a.value, b.value)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_q2d ---

// AVQ2D wraps av_q2d.
/*
  Convert an AVRational to a `double`.
  @param a AVRational to convert
  @return `a` in floating-point form
  @see av_d2q()
*/
func AVQ2D(a *AVRational) float64 {
	ret := C.av_q2d(a.value)
	return float64(ret)
}

// --- Function av_reduce ---

// av_reduce skipped due to dstNum

// --- Function av_mul_q ---

// AVMulQ wraps av_mul_q.
/*
  Multiply two rationals.
  @param b First rational
  @param c Second rational
  @return b*c
*/
func AVMulQ(b *AVRational, c *AVRational) *AVRational {
	ret := C.av_mul_q(b.value, c.value)
	return &AVRational{value: ret}
}

// --- Function av_div_q ---

// AVDivQ wraps av_div_q.
/*
  Divide one rational by another.
  @param b First rational
  @param c Second rational
  @return b/c
*/
func AVDivQ(b *AVRational, c *AVRational) *AVRational {
	ret := C.av_div_q(b.value, c.value)
	return &AVRational{value: ret}
}

// --- Function av_add_q ---

// AVAddQ wraps av_add_q.
/*
  Add two rationals.
  @param b First rational
  @param c Second rational
  @return b+c
*/
func AVAddQ(b *AVRational, c *AVRational) *AVRational {
	ret := C.av_add_q(b.value, c.value)
	return &AVRational{value: ret}
}

// --- Function av_sub_q ---

// AVSubQ wraps av_sub_q.
/*
  Subtract one rational from another.
  @param b First rational
  @param c Second rational
  @return b-c
*/
func AVSubQ(b *AVRational, c *AVRational) *AVRational {
	ret := C.av_sub_q(b.value, c.value)
	return &AVRational{value: ret}
}

// --- Function av_inv_q ---

// AVInvQ wraps av_inv_q.
/*
  Invert a rational.
  @param q value
  @return 1 / q
*/
func AVInvQ(q *AVRational) *AVRational {
	ret := C.av_inv_q(q.value)
	return &AVRational{value: ret}
}

// --- Function av_d2q ---

// AVD2Q wraps av_d2q.
/*
  Convert a double precision floating point number to a rational.

  In case of infinity, the returned value is expressed as `{1, 0}` or
  `{-1, 0}` depending on the sign.

  @param d   `double` to convert
  @param max Maximum allowed numerator and denominator
  @return `d` in AVRational form
  @see av_q2d()
*/
func AVD2Q(d float64, max int) *AVRational {
	ret := C.av_d2q(C.double(d), C.int(max))
	return &AVRational{value: ret}
}

// --- Function av_nearer_q ---

// AVNearerQ wraps av_nearer_q.
/*
  Find which of the two rationals is closer to another rational.

  @param q     Rational to be compared against
  @param q1    Rational to be tested
  @param q2    Rational to be tested
  @return One of the following values:
          - 1 if `q1` is nearer to `q` than `q2`
          - -1 if `q2` is nearer to `q` than `q1`
          - 0 if they have the same distance
*/
func AVNearerQ(q *AVRational, q1 *AVRational, q2 *AVRational) (int, error) {
	ret := C.av_nearer_q(q.value, q1.value, q2.value)
	return int(ret), WrapErr(int(ret))
}

// --- Function av_find_nearest_q_idx ---

// av_find_nearest_q_idx skipped due to qList

// --- Function av_q2intfloat ---

// AVQ2Intfloat wraps av_q2intfloat.
/*
  Convert an AVRational to a IEEE 32-bit `float` expressed in fixed-point
  format.

  @param q Rational to be converted
  @return Equivalent floating-point value, expressed as an unsigned 32-bit
          integer.
  @note The returned value is platform-indepedant.
*/
func AVQ2Intfloat(q *AVRational) uint32 {
	ret := C.av_q2intfloat(q.value)
	return uint32(ret)
}

// --- Function av_gcd_q ---

// AVGcdQ wraps av_gcd_q.
/*
  Return the best rational so that a and b are multiple of it.
  If the resulting denominator is larger than max_den, return def.
*/
func AVGcdQ(a *AVRational, b *AVRational, maxDen int, def *AVRational) *AVRational {
	ret := C.av_gcd_q(a.value, b.value, C.int(maxDen), def.value)
	return &AVRational{value: ret}
}

// --- Function av_get_sample_fmt_name ---

// AVGetSampleFmtName wraps av_get_sample_fmt_name.
/*
  Return the name of sample_fmt, or NULL if sample_fmt is not
  recognized.
*/
func AVGetSampleFmtName(sampleFmt AVSampleFormat) *CStr {
	ret := C.av_get_sample_fmt_name(C.enum_AVSampleFormat(sampleFmt))
	return wrapCStr(ret)
}

// --- Function av_get_sample_fmt ---

// AVGetSampleFmt wraps av_get_sample_fmt.
/*
  Return a sample format corresponding to name, or AV_SAMPLE_FMT_NONE
  on error.
*/
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
/*
  Return the planar<->packed alternative form of the given sample format, or
  AV_SAMPLE_FMT_NONE on error. If the passed sample_fmt is already in the
  requested planar/packed format, the format returned is the same as the
  input.
*/
func AVGetAltSampleFmt(sampleFmt AVSampleFormat, planar int) AVSampleFormat {
	ret := C.av_get_alt_sample_fmt(C.enum_AVSampleFormat(sampleFmt), C.int(planar))
	return AVSampleFormat(ret)
}

// --- Function av_get_packed_sample_fmt ---

// AVGetPackedSampleFmt wraps av_get_packed_sample_fmt.
/*
  Get the packed alternative form of the given sample format.

  If the passed sample_fmt is already in packed format, the format returned is
  the same as the input.

  @return  the packed alternative form of the given sample format or
  AV_SAMPLE_FMT_NONE on error.
*/
func AVGetPackedSampleFmt(sampleFmt AVSampleFormat) AVSampleFormat {
	ret := C.av_get_packed_sample_fmt(C.enum_AVSampleFormat(sampleFmt))
	return AVSampleFormat(ret)
}

// --- Function av_get_planar_sample_fmt ---

// AVGetPlanarSampleFmt wraps av_get_planar_sample_fmt.
/*
  Get the planar alternative form of the given sample format.

  If the passed sample_fmt is already in planar format, the format returned is
  the same as the input.

  @return  the planar alternative form of the given sample format or
  AV_SAMPLE_FMT_NONE on error.
*/
func AVGetPlanarSampleFmt(sampleFmt AVSampleFormat) AVSampleFormat {
	ret := C.av_get_planar_sample_fmt(C.enum_AVSampleFormat(sampleFmt))
	return AVSampleFormat(ret)
}

// --- Function av_get_sample_fmt_string ---

// AVGetSampleFmtString wraps av_get_sample_fmt_string.
/*
  Generate a string corresponding to the sample format with
  sample_fmt, or a header if sample_fmt is negative.

  @param buf the buffer where to write the string
  @param buf_size the size of buf
  @param sample_fmt the number of the sample format to print the
  corresponding info string, or a negative value to print the
  corresponding header.
  @return the pointer to the filled buffer or NULL if sample_fmt is
  unknown or in case of other errors
*/
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
/*
  Return number of bytes per sample.

  @param sample_fmt the sample format
  @return number of bytes per sample or zero if unknown for the given
  sample format
*/
func AVGetBytesPerSample(sampleFmt AVSampleFormat) (int, error) {
	ret := C.av_get_bytes_per_sample(C.enum_AVSampleFormat(sampleFmt))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_sample_fmt_is_planar ---

// AVSampleFmtIsPlanar wraps av_sample_fmt_is_planar.
/*
  Check if the sample format is planar.

  @param sample_fmt the sample format to inspect
  @return 1 if the sample format is planar, 0 if it is interleaved
*/
func AVSampleFmtIsPlanar(sampleFmt AVSampleFormat) (int, error) {
	ret := C.av_sample_fmt_is_planar(C.enum_AVSampleFormat(sampleFmt))
	return int(ret), WrapErr(int(ret))
}

// --- Function av_samples_get_buffer_size ---

// av_samples_get_buffer_size skipped due to linesize

// --- Function av_samples_fill_arrays ---

// av_samples_fill_arrays skipped due to audioData

// --- Function av_samples_alloc ---

// av_samples_alloc skipped due to audioData

// --- Function av_samples_alloc_array_and_samples ---

// av_samples_alloc_array_and_samples skipped due to audioData

// --- Function av_samples_copy ---

// av_samples_copy skipped due to dst

// --- Function av_samples_set_silence ---

// av_samples_set_silence skipped due to audioData
