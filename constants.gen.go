package ffmpeg

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

// AVInputBufferMinSize wraps AV_INPUT_BUFFER_MIN_SIZE.
const AVInputBufferMinSize int = C.AV_INPUT_BUFFER_MIN_SIZE

// AVCodecFlagUnaligned wraps AV_CODEC_FLAG_UNALIGNED.
const AVCodecFlagUnaligned int = C.AV_CODEC_FLAG_UNALIGNED

// AVCodecFlagQscale wraps AV_CODEC_FLAG_QSCALE.
const AVCodecFlagQscale int = C.AV_CODEC_FLAG_QSCALE

// AVCodecFlag4Mv wraps AV_CODEC_FLAG_4MV.
const AVCodecFlag4Mv int = C.AV_CODEC_FLAG_4MV

// AVCodecFlagOutputCorrupt wraps AV_CODEC_FLAG_OUTPUT_CORRUPT.
const AVCodecFlagOutputCorrupt int = C.AV_CODEC_FLAG_OUTPUT_CORRUPT

// AVCodecFlagQpel wraps AV_CODEC_FLAG_QPEL.
const AVCodecFlagQpel int = C.AV_CODEC_FLAG_QPEL

// AVCodecFlagDropchanged wraps AV_CODEC_FLAG_DROPCHANGED.
const AVCodecFlagDropchanged int = C.AV_CODEC_FLAG_DROPCHANGED

// AVCodecFlagReconFrame wraps AV_CODEC_FLAG_RECON_FRAME.
const AVCodecFlagReconFrame int = C.AV_CODEC_FLAG_RECON_FRAME

// AVCodecFlagCopyOpaque wraps AV_CODEC_FLAG_COPY_OPAQUE.
const AVCodecFlagCopyOpaque int = C.AV_CODEC_FLAG_COPY_OPAQUE

// AVCodecFlagFrameDuration wraps AV_CODEC_FLAG_FRAME_DURATION.
const AVCodecFlagFrameDuration int = C.AV_CODEC_FLAG_FRAME_DURATION

// AVCodecFlagPass1 wraps AV_CODEC_FLAG_PASS1.
const AVCodecFlagPass1 int = C.AV_CODEC_FLAG_PASS1

// AVCodecFlagPass2 wraps AV_CODEC_FLAG_PASS2.
const AVCodecFlagPass2 int = C.AV_CODEC_FLAG_PASS2

// AVCodecFlagLoopFilter wraps AV_CODEC_FLAG_LOOP_FILTER.
const AVCodecFlagLoopFilter int = C.AV_CODEC_FLAG_LOOP_FILTER

// AVCodecFlagGray wraps AV_CODEC_FLAG_GRAY.
const AVCodecFlagGray int = C.AV_CODEC_FLAG_GRAY

// AVCodecFlagPsnr wraps AV_CODEC_FLAG_PSNR.
const AVCodecFlagPsnr int = C.AV_CODEC_FLAG_PSNR

// AVCodecFlagInterlacedDct wraps AV_CODEC_FLAG_INTERLACED_DCT.
const AVCodecFlagInterlacedDct int = C.AV_CODEC_FLAG_INTERLACED_DCT

// AVCodecFlagLowDelay wraps AV_CODEC_FLAG_LOW_DELAY.
const AVCodecFlagLowDelay int = C.AV_CODEC_FLAG_LOW_DELAY

// AVCodecFlagGlobalHeader wraps AV_CODEC_FLAG_GLOBAL_HEADER.
const AVCodecFlagGlobalHeader int = C.AV_CODEC_FLAG_GLOBAL_HEADER

// AVCodecFlagBitexact wraps AV_CODEC_FLAG_BITEXACT.
const AVCodecFlagBitexact int = C.AV_CODEC_FLAG_BITEXACT

// AVCodecFlagAcPred wraps AV_CODEC_FLAG_AC_PRED.
const AVCodecFlagAcPred int = C.AV_CODEC_FLAG_AC_PRED

// AVCodecFlagInterlacedMe wraps AV_CODEC_FLAG_INTERLACED_ME.
const AVCodecFlagInterlacedMe int = C.AV_CODEC_FLAG_INTERLACED_ME

// AVCodecFlagClosedGop wraps AV_CODEC_FLAG_CLOSED_GOP.
const AVCodecFlagClosedGop int = C.AV_CODEC_FLAG_CLOSED_GOP

// AVCodecFlag2Fast wraps AV_CODEC_FLAG2_FAST.
const AVCodecFlag2Fast int = C.AV_CODEC_FLAG2_FAST

// AVCodecFlag2NoOutput wraps AV_CODEC_FLAG2_NO_OUTPUT.
const AVCodecFlag2NoOutput int = C.AV_CODEC_FLAG2_NO_OUTPUT

// AVCodecFlag2LocalHeader wraps AV_CODEC_FLAG2_LOCAL_HEADER.
const AVCodecFlag2LocalHeader int = C.AV_CODEC_FLAG2_LOCAL_HEADER

// AVCodecFlag2Chunks wraps AV_CODEC_FLAG2_CHUNKS.
const AVCodecFlag2Chunks int = C.AV_CODEC_FLAG2_CHUNKS

// AVCodecFlag2IgnoreCrop wraps AV_CODEC_FLAG2_IGNORE_CROP.
const AVCodecFlag2IgnoreCrop int = C.AV_CODEC_FLAG2_IGNORE_CROP

// AVCodecFlag2ShowAll wraps AV_CODEC_FLAG2_SHOW_ALL.
const AVCodecFlag2ShowAll int = C.AV_CODEC_FLAG2_SHOW_ALL

// AVCodecFlag2ExportMvs wraps AV_CODEC_FLAG2_EXPORT_MVS.
const AVCodecFlag2ExportMvs int = C.AV_CODEC_FLAG2_EXPORT_MVS

// AVCodecFlag2SkipManual wraps AV_CODEC_FLAG2_SKIP_MANUAL.
const AVCodecFlag2SkipManual int = C.AV_CODEC_FLAG2_SKIP_MANUAL

// AVCodecFlag2RoFlushNoop wraps AV_CODEC_FLAG2_RO_FLUSH_NOOP.
const AVCodecFlag2RoFlushNoop int = C.AV_CODEC_FLAG2_RO_FLUSH_NOOP

// AVCodecFlag2IccProfiles wraps AV_CODEC_FLAG2_ICC_PROFILES.
const AVCodecFlag2IccProfiles int = C.AV_CODEC_FLAG2_ICC_PROFILES

// AVCodecExportDataMvs wraps AV_CODEC_EXPORT_DATA_MVS.
const AVCodecExportDataMvs int = C.AV_CODEC_EXPORT_DATA_MVS

// AVCodecExportDataPrft wraps AV_CODEC_EXPORT_DATA_PRFT.
const AVCodecExportDataPrft int = C.AV_CODEC_EXPORT_DATA_PRFT

// AVCodecExportDataVideoEncParams wraps AV_CODEC_EXPORT_DATA_VIDEO_ENC_PARAMS.
const AVCodecExportDataVideoEncParams int = C.AV_CODEC_EXPORT_DATA_VIDEO_ENC_PARAMS

// AVCodecExportDataFilmGrain wraps AV_CODEC_EXPORT_DATA_FILM_GRAIN.
const AVCodecExportDataFilmGrain int = C.AV_CODEC_EXPORT_DATA_FILM_GRAIN

// AVGetBufferFlagRef wraps AV_GET_BUFFER_FLAG_REF.
const AVGetBufferFlagRef int = C.AV_GET_BUFFER_FLAG_REF

// AVGetEncodeBufferFlagRef wraps AV_GET_ENCODE_BUFFER_FLAG_REF.
const AVGetEncodeBufferFlagRef int = C.AV_GET_ENCODE_BUFFER_FLAG_REF

// FfCompressionDefault wraps FF_COMPRESSION_DEFAULT.
const FfCompressionDefault int = C.FF_COMPRESSION_DEFAULT

// FfCmpSad wraps FF_CMP_SAD.
const FfCmpSad int = C.FF_CMP_SAD

// FfCmpSse wraps FF_CMP_SSE.
const FfCmpSse int = C.FF_CMP_SSE

// FfCmpSatd wraps FF_CMP_SATD.
const FfCmpSatd int = C.FF_CMP_SATD

// FfCmpDct wraps FF_CMP_DCT.
const FfCmpDct int = C.FF_CMP_DCT

// FfCmpPsnr wraps FF_CMP_PSNR.
const FfCmpPsnr int = C.FF_CMP_PSNR

// FfCmpBit wraps FF_CMP_BIT.
const FfCmpBit int = C.FF_CMP_BIT

// FfCmpRd wraps FF_CMP_RD.
const FfCmpRd int = C.FF_CMP_RD

// FfCmpZero wraps FF_CMP_ZERO.
const FfCmpZero int = C.FF_CMP_ZERO

// FfCmpVsad wraps FF_CMP_VSAD.
const FfCmpVsad int = C.FF_CMP_VSAD

// FfCmpVsse wraps FF_CMP_VSSE.
const FfCmpVsse int = C.FF_CMP_VSSE

// FfCmpNsse wraps FF_CMP_NSSE.
const FfCmpNsse int = C.FF_CMP_NSSE

// FfCmpW53 wraps FF_CMP_W53.
const FfCmpW53 int = C.FF_CMP_W53

// FfCmpW97 wraps FF_CMP_W97.
const FfCmpW97 int = C.FF_CMP_W97

// FfCmpDctmax wraps FF_CMP_DCTMAX.
const FfCmpDctmax int = C.FF_CMP_DCTMAX

// FfCmpDct264 wraps FF_CMP_DCT264.
const FfCmpDct264 int = C.FF_CMP_DCT264

// FfCmpMedianSad wraps FF_CMP_MEDIAN_SAD.
const FfCmpMedianSad int = C.FF_CMP_MEDIAN_SAD

// FfCmpChroma wraps FF_CMP_CHROMA.
const FfCmpChroma int = C.FF_CMP_CHROMA

// SliceFlagCodedOrder wraps SLICE_FLAG_CODED_ORDER.
const SliceFlagCodedOrder int = C.SLICE_FLAG_CODED_ORDER

// SliceFlagAllowField wraps SLICE_FLAG_ALLOW_FIELD.
const SliceFlagAllowField int = C.SLICE_FLAG_ALLOW_FIELD

// SliceFlagAllowPlane wraps SLICE_FLAG_ALLOW_PLANE.
const SliceFlagAllowPlane int = C.SLICE_FLAG_ALLOW_PLANE

// FfMbDecisionSimple wraps FF_MB_DECISION_SIMPLE.
const FfMbDecisionSimple int = C.FF_MB_DECISION_SIMPLE

// FfMbDecisionBits wraps FF_MB_DECISION_BITS.
const FfMbDecisionBits int = C.FF_MB_DECISION_BITS

// FfMbDecisionRd wraps FF_MB_DECISION_RD.
const FfMbDecisionRd int = C.FF_MB_DECISION_RD

// FfBugAutodetect wraps FF_BUG_AUTODETECT.
const FfBugAutodetect int = C.FF_BUG_AUTODETECT

// FfBugXvidIlace wraps FF_BUG_XVID_ILACE.
const FfBugXvidIlace int = C.FF_BUG_XVID_ILACE

// FfBugUmp4 wraps FF_BUG_UMP4.
const FfBugUmp4 int = C.FF_BUG_UMP4

// FfBugNoPadding wraps FF_BUG_NO_PADDING.
const FfBugNoPadding int = C.FF_BUG_NO_PADDING

// FfBugAmv wraps FF_BUG_AMV.
const FfBugAmv int = C.FF_BUG_AMV

// FfBugQpelChroma wraps FF_BUG_QPEL_CHROMA.
const FfBugQpelChroma int = C.FF_BUG_QPEL_CHROMA

// FfBugStdQpel wraps FF_BUG_STD_QPEL.
const FfBugStdQpel int = C.FF_BUG_STD_QPEL

// FfBugQpelChroma2 wraps FF_BUG_QPEL_CHROMA2.
const FfBugQpelChroma2 int = C.FF_BUG_QPEL_CHROMA2

// FfBugDirectBlocksize wraps FF_BUG_DIRECT_BLOCKSIZE.
const FfBugDirectBlocksize int = C.FF_BUG_DIRECT_BLOCKSIZE

// FfBugEdge wraps FF_BUG_EDGE.
const FfBugEdge int = C.FF_BUG_EDGE

// FfBugHpelChroma wraps FF_BUG_HPEL_CHROMA.
const FfBugHpelChroma int = C.FF_BUG_HPEL_CHROMA

// FfBugDcClip wraps FF_BUG_DC_CLIP.
const FfBugDcClip int = C.FF_BUG_DC_CLIP

// FfBugMs wraps FF_BUG_MS.
const FfBugMs int = C.FF_BUG_MS

// FfBugTruncated wraps FF_BUG_TRUNCATED.
const FfBugTruncated int = C.FF_BUG_TRUNCATED

// FfBugIedge wraps FF_BUG_IEDGE.
const FfBugIedge int = C.FF_BUG_IEDGE

// FfEcGuessMvs wraps FF_EC_GUESS_MVS.
const FfEcGuessMvs int = C.FF_EC_GUESS_MVS

// FfEcDeblock wraps FF_EC_DEBLOCK.
const FfEcDeblock int = C.FF_EC_DEBLOCK

// FfEcFavorInter wraps FF_EC_FAVOR_INTER.
const FfEcFavorInter int = C.FF_EC_FAVOR_INTER

// FfDebugPictInfo wraps FF_DEBUG_PICT_INFO.
const FfDebugPictInfo int = C.FF_DEBUG_PICT_INFO

// FfDebugRc wraps FF_DEBUG_RC.
const FfDebugRc int = C.FF_DEBUG_RC

// FfDebugBitstream wraps FF_DEBUG_BITSTREAM.
const FfDebugBitstream int = C.FF_DEBUG_BITSTREAM

// FfDebugMbType wraps FF_DEBUG_MB_TYPE.
const FfDebugMbType int = C.FF_DEBUG_MB_TYPE

// FfDebugQp wraps FF_DEBUG_QP.
const FfDebugQp int = C.FF_DEBUG_QP

// FfDebugDctCoeff wraps FF_DEBUG_DCT_COEFF.
const FfDebugDctCoeff int = C.FF_DEBUG_DCT_COEFF

// FfDebugSkip wraps FF_DEBUG_SKIP.
const FfDebugSkip int = C.FF_DEBUG_SKIP

// FfDebugStartcode wraps FF_DEBUG_STARTCODE.
const FfDebugStartcode int = C.FF_DEBUG_STARTCODE

// FfDebugEr wraps FF_DEBUG_ER.
const FfDebugEr int = C.FF_DEBUG_ER

// FfDebugMmco wraps FF_DEBUG_MMCO.
const FfDebugMmco int = C.FF_DEBUG_MMCO

// FfDebugBugs wraps FF_DEBUG_BUGS.
const FfDebugBugs int = C.FF_DEBUG_BUGS

// FfDebugBuffers wraps FF_DEBUG_BUFFERS.
const FfDebugBuffers int = C.FF_DEBUG_BUFFERS

// FfDebugThreads wraps FF_DEBUG_THREADS.
const FfDebugThreads int = C.FF_DEBUG_THREADS

// FfDebugGreenMd wraps FF_DEBUG_GREEN_MD.
const FfDebugGreenMd int = C.FF_DEBUG_GREEN_MD

// FfDebugNomc wraps FF_DEBUG_NOMC.
const FfDebugNomc int = C.FF_DEBUG_NOMC

// FfDctAuto wraps FF_DCT_AUTO.
const FfDctAuto int = C.FF_DCT_AUTO

// FfDctFastint wraps FF_DCT_FASTINT.
const FfDctFastint int = C.FF_DCT_FASTINT

// FfDctInt wraps FF_DCT_INT.
const FfDctInt int = C.FF_DCT_INT

// FfDctMmx wraps FF_DCT_MMX.
const FfDctMmx int = C.FF_DCT_MMX

// FfDctAltivec wraps FF_DCT_ALTIVEC.
const FfDctAltivec int = C.FF_DCT_ALTIVEC

// FfDctFaan wraps FF_DCT_FAAN.
const FfDctFaan int = C.FF_DCT_FAAN

// FfIdctAuto wraps FF_IDCT_AUTO.
const FfIdctAuto int = C.FF_IDCT_AUTO

// FfIdctInt wraps FF_IDCT_INT.
const FfIdctInt int = C.FF_IDCT_INT

// FfIdctSimple wraps FF_IDCT_SIMPLE.
const FfIdctSimple int = C.FF_IDCT_SIMPLE

// FfIdctSimplemmx wraps FF_IDCT_SIMPLEMMX.
const FfIdctSimplemmx int = C.FF_IDCT_SIMPLEMMX

// FfIdctArm wraps FF_IDCT_ARM.
const FfIdctArm int = C.FF_IDCT_ARM

// FfIdctAltivec wraps FF_IDCT_ALTIVEC.
const FfIdctAltivec int = C.FF_IDCT_ALTIVEC

// FfIdctSimplearm wraps FF_IDCT_SIMPLEARM.
const FfIdctSimplearm int = C.FF_IDCT_SIMPLEARM

// FfIdctXvid wraps FF_IDCT_XVID.
const FfIdctXvid int = C.FF_IDCT_XVID

// FfIdctSimplearmv5Te wraps FF_IDCT_SIMPLEARMV5TE.
const FfIdctSimplearmv5Te int = C.FF_IDCT_SIMPLEARMV5TE

// FfIdctSimplearmv6 wraps FF_IDCT_SIMPLEARMV6.
const FfIdctSimplearmv6 int = C.FF_IDCT_SIMPLEARMV6

// FfIdctFaan wraps FF_IDCT_FAAN.
const FfIdctFaan int = C.FF_IDCT_FAAN

// FfIdctSimpleneon wraps FF_IDCT_SIMPLENEON.
const FfIdctSimpleneon int = C.FF_IDCT_SIMPLENEON

// FfIdctNone wraps FF_IDCT_NONE.
const FfIdctNone int = C.FF_IDCT_NONE

// FfIdctSimpleauto wraps FF_IDCT_SIMPLEAUTO.
const FfIdctSimpleauto int = C.FF_IDCT_SIMPLEAUTO

// FfThreadFrame wraps FF_THREAD_FRAME.
const FfThreadFrame int = C.FF_THREAD_FRAME

// FfThreadSlice wraps FF_THREAD_SLICE.
const FfThreadSlice int = C.FF_THREAD_SLICE

// FfProfileUnknown wraps FF_PROFILE_UNKNOWN.
const FfProfileUnknown int = C.FF_PROFILE_UNKNOWN

// FfProfileReserved wraps FF_PROFILE_RESERVED.
const FfProfileReserved int = C.FF_PROFILE_RESERVED

// FfProfileAacMain wraps FF_PROFILE_AAC_MAIN.
const FfProfileAacMain int = C.FF_PROFILE_AAC_MAIN

// FfProfileAacLow wraps FF_PROFILE_AAC_LOW.
const FfProfileAacLow int = C.FF_PROFILE_AAC_LOW

// FfProfileAacSsr wraps FF_PROFILE_AAC_SSR.
const FfProfileAacSsr int = C.FF_PROFILE_AAC_SSR

// FfProfileAacLtp wraps FF_PROFILE_AAC_LTP.
const FfProfileAacLtp int = C.FF_PROFILE_AAC_LTP

// FfProfileAacHe wraps FF_PROFILE_AAC_HE.
const FfProfileAacHe int = C.FF_PROFILE_AAC_HE

// FfProfileAacHeV2 wraps FF_PROFILE_AAC_HE_V2.
const FfProfileAacHeV2 int = C.FF_PROFILE_AAC_HE_V2

// FfProfileAacLd wraps FF_PROFILE_AAC_LD.
const FfProfileAacLd int = C.FF_PROFILE_AAC_LD

// FfProfileAacEld wraps FF_PROFILE_AAC_ELD.
const FfProfileAacEld int = C.FF_PROFILE_AAC_ELD

// FfProfileMpeg2AacLow wraps FF_PROFILE_MPEG2_AAC_LOW.
const FfProfileMpeg2AacLow int = C.FF_PROFILE_MPEG2_AAC_LOW

// FfProfileMpeg2AacHe wraps FF_PROFILE_MPEG2_AAC_HE.
const FfProfileMpeg2AacHe int = C.FF_PROFILE_MPEG2_AAC_HE

// FfProfileDnxhd wraps FF_PROFILE_DNXHD.
const FfProfileDnxhd int = C.FF_PROFILE_DNXHD

// FfProfileDnxhrLb wraps FF_PROFILE_DNXHR_LB.
const FfProfileDnxhrLb int = C.FF_PROFILE_DNXHR_LB

// FfProfileDnxhrSq wraps FF_PROFILE_DNXHR_SQ.
const FfProfileDnxhrSq int = C.FF_PROFILE_DNXHR_SQ

// FfProfileDnxhrHq wraps FF_PROFILE_DNXHR_HQ.
const FfProfileDnxhrHq int = C.FF_PROFILE_DNXHR_HQ

// FfProfileDnxhrHqx wraps FF_PROFILE_DNXHR_HQX.
const FfProfileDnxhrHqx int = C.FF_PROFILE_DNXHR_HQX

// FfProfileDnxhr444 wraps FF_PROFILE_DNXHR_444.
const FfProfileDnxhr444 int = C.FF_PROFILE_DNXHR_444

// FfProfileDts wraps FF_PROFILE_DTS.
const FfProfileDts int = C.FF_PROFILE_DTS

// FfProfileDtsEs wraps FF_PROFILE_DTS_ES.
const FfProfileDtsEs int = C.FF_PROFILE_DTS_ES

// FfProfileDts9624 wraps FF_PROFILE_DTS_96_24.
const FfProfileDts9624 int = C.FF_PROFILE_DTS_96_24

// FfProfileDtsHdHra wraps FF_PROFILE_DTS_HD_HRA.
const FfProfileDtsHdHra int = C.FF_PROFILE_DTS_HD_HRA

// FfProfileDtsHdMa wraps FF_PROFILE_DTS_HD_MA.
const FfProfileDtsHdMa int = C.FF_PROFILE_DTS_HD_MA

// FfProfileDtsExpress wraps FF_PROFILE_DTS_EXPRESS.
const FfProfileDtsExpress int = C.FF_PROFILE_DTS_EXPRESS

// FfProfileMpeg2422 wraps FF_PROFILE_MPEG2_422.
const FfProfileMpeg2422 int = C.FF_PROFILE_MPEG2_422

// FfProfileMpeg2High wraps FF_PROFILE_MPEG2_HIGH.
const FfProfileMpeg2High int = C.FF_PROFILE_MPEG2_HIGH

// FfProfileMpeg2Ss wraps FF_PROFILE_MPEG2_SS.
const FfProfileMpeg2Ss int = C.FF_PROFILE_MPEG2_SS

// FfProfileMpeg2SnrScalable wraps FF_PROFILE_MPEG2_SNR_SCALABLE.
const FfProfileMpeg2SnrScalable int = C.FF_PROFILE_MPEG2_SNR_SCALABLE

// FfProfileMpeg2Main wraps FF_PROFILE_MPEG2_MAIN.
const FfProfileMpeg2Main int = C.FF_PROFILE_MPEG2_MAIN

// FfProfileMpeg2Simple wraps FF_PROFILE_MPEG2_SIMPLE.
const FfProfileMpeg2Simple int = C.FF_PROFILE_MPEG2_SIMPLE

// FfProfileH264Constrained wraps FF_PROFILE_H264_CONSTRAINED.
const FfProfileH264Constrained int = C.FF_PROFILE_H264_CONSTRAINED

// FfProfileH264Intra wraps FF_PROFILE_H264_INTRA.
const FfProfileH264Intra int = C.FF_PROFILE_H264_INTRA

// FfProfileH264Baseline wraps FF_PROFILE_H264_BASELINE.
const FfProfileH264Baseline int = C.FF_PROFILE_H264_BASELINE

// FfProfileH264ConstrainedBaseline wraps FF_PROFILE_H264_CONSTRAINED_BASELINE.
const FfProfileH264ConstrainedBaseline int = C.FF_PROFILE_H264_CONSTRAINED_BASELINE

// FfProfileH264Main wraps FF_PROFILE_H264_MAIN.
const FfProfileH264Main int = C.FF_PROFILE_H264_MAIN

// FfProfileH264Extended wraps FF_PROFILE_H264_EXTENDED.
const FfProfileH264Extended int = C.FF_PROFILE_H264_EXTENDED

// FfProfileH264High wraps FF_PROFILE_H264_HIGH.
const FfProfileH264High int = C.FF_PROFILE_H264_HIGH

// FfProfileH264High10 wraps FF_PROFILE_H264_HIGH_10.
const FfProfileH264High10 int = C.FF_PROFILE_H264_HIGH_10

// FfProfileH264High10Intra wraps FF_PROFILE_H264_HIGH_10_INTRA.
const FfProfileH264High10Intra int = C.FF_PROFILE_H264_HIGH_10_INTRA

// FfProfileH264MultiviewHigh wraps FF_PROFILE_H264_MULTIVIEW_HIGH.
const FfProfileH264MultiviewHigh int = C.FF_PROFILE_H264_MULTIVIEW_HIGH

// FfProfileH264High422 wraps FF_PROFILE_H264_HIGH_422.
const FfProfileH264High422 int = C.FF_PROFILE_H264_HIGH_422

// FfProfileH264High422Intra wraps FF_PROFILE_H264_HIGH_422_INTRA.
const FfProfileH264High422Intra int = C.FF_PROFILE_H264_HIGH_422_INTRA

// FfProfileH264StereoHigh wraps FF_PROFILE_H264_STEREO_HIGH.
const FfProfileH264StereoHigh int = C.FF_PROFILE_H264_STEREO_HIGH

// FfProfileH264High444 wraps FF_PROFILE_H264_HIGH_444.
const FfProfileH264High444 int = C.FF_PROFILE_H264_HIGH_444

// FfProfileH264High444Predictive wraps FF_PROFILE_H264_HIGH_444_PREDICTIVE.
const FfProfileH264High444Predictive int = C.FF_PROFILE_H264_HIGH_444_PREDICTIVE

// FfProfileH264High444Intra wraps FF_PROFILE_H264_HIGH_444_INTRA.
const FfProfileH264High444Intra int = C.FF_PROFILE_H264_HIGH_444_INTRA

// FfProfileH264Cavlc444 wraps FF_PROFILE_H264_CAVLC_444.
const FfProfileH264Cavlc444 int = C.FF_PROFILE_H264_CAVLC_444

// FfProfileVc1Simple wraps FF_PROFILE_VC1_SIMPLE.
const FfProfileVc1Simple int = C.FF_PROFILE_VC1_SIMPLE

// FfProfileVc1Main wraps FF_PROFILE_VC1_MAIN.
const FfProfileVc1Main int = C.FF_PROFILE_VC1_MAIN

// FfProfileVc1Complex wraps FF_PROFILE_VC1_COMPLEX.
const FfProfileVc1Complex int = C.FF_PROFILE_VC1_COMPLEX

// FfProfileVc1Advanced wraps FF_PROFILE_VC1_ADVANCED.
const FfProfileVc1Advanced int = C.FF_PROFILE_VC1_ADVANCED

// FfProfileMpeg4Simple wraps FF_PROFILE_MPEG4_SIMPLE.
const FfProfileMpeg4Simple int = C.FF_PROFILE_MPEG4_SIMPLE

// FfProfileMpeg4SimpleScalable wraps FF_PROFILE_MPEG4_SIMPLE_SCALABLE.
const FfProfileMpeg4SimpleScalable int = C.FF_PROFILE_MPEG4_SIMPLE_SCALABLE

// FfProfileMpeg4Core wraps FF_PROFILE_MPEG4_CORE.
const FfProfileMpeg4Core int = C.FF_PROFILE_MPEG4_CORE

// FfProfileMpeg4Main wraps FF_PROFILE_MPEG4_MAIN.
const FfProfileMpeg4Main int = C.FF_PROFILE_MPEG4_MAIN

// FfProfileMpeg4NBit wraps FF_PROFILE_MPEG4_N_BIT.
const FfProfileMpeg4NBit int = C.FF_PROFILE_MPEG4_N_BIT

// FfProfileMpeg4ScalableTexture wraps FF_PROFILE_MPEG4_SCALABLE_TEXTURE.
const FfProfileMpeg4ScalableTexture int = C.FF_PROFILE_MPEG4_SCALABLE_TEXTURE

// FfProfileMpeg4SimpleFaceAnimation wraps FF_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION.
const FfProfileMpeg4SimpleFaceAnimation int = C.FF_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION

// FfProfileMpeg4BasicAnimatedTexture wraps FF_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE.
const FfProfileMpeg4BasicAnimatedTexture int = C.FF_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE

// FfProfileMpeg4Hybrid wraps FF_PROFILE_MPEG4_HYBRID.
const FfProfileMpeg4Hybrid int = C.FF_PROFILE_MPEG4_HYBRID

// FfProfileMpeg4AdvancedRealTime wraps FF_PROFILE_MPEG4_ADVANCED_REAL_TIME.
const FfProfileMpeg4AdvancedRealTime int = C.FF_PROFILE_MPEG4_ADVANCED_REAL_TIME

// FfProfileMpeg4CoreScalable wraps FF_PROFILE_MPEG4_CORE_SCALABLE.
const FfProfileMpeg4CoreScalable int = C.FF_PROFILE_MPEG4_CORE_SCALABLE

// FfProfileMpeg4AdvancedCoding wraps FF_PROFILE_MPEG4_ADVANCED_CODING.
const FfProfileMpeg4AdvancedCoding int = C.FF_PROFILE_MPEG4_ADVANCED_CODING

// FfProfileMpeg4AdvancedCore wraps FF_PROFILE_MPEG4_ADVANCED_CORE.
const FfProfileMpeg4AdvancedCore int = C.FF_PROFILE_MPEG4_ADVANCED_CORE

// FfProfileMpeg4AdvancedScalableTexture wraps FF_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE.
const FfProfileMpeg4AdvancedScalableTexture int = C.FF_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE

// FfProfileMpeg4SimpleStudio wraps FF_PROFILE_MPEG4_SIMPLE_STUDIO.
const FfProfileMpeg4SimpleStudio int = C.FF_PROFILE_MPEG4_SIMPLE_STUDIO

// FfProfileMpeg4AdvancedSimple wraps FF_PROFILE_MPEG4_ADVANCED_SIMPLE.
const FfProfileMpeg4AdvancedSimple int = C.FF_PROFILE_MPEG4_ADVANCED_SIMPLE

// FfProfileJpeg2000CstreamRestriction0 wraps FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0.
const FfProfileJpeg2000CstreamRestriction0 int = C.FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0

// FfProfileJpeg2000CstreamRestriction1 wraps FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1.
const FfProfileJpeg2000CstreamRestriction1 int = C.FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1

// FfProfileJpeg2000CstreamNoRestriction wraps FF_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION.
const FfProfileJpeg2000CstreamNoRestriction int = C.FF_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION

// FfProfileJpeg2000Dcinema2K wraps FF_PROFILE_JPEG2000_DCINEMA_2K.
const FfProfileJpeg2000Dcinema2K int = C.FF_PROFILE_JPEG2000_DCINEMA_2K

// FfProfileJpeg2000Dcinema4K wraps FF_PROFILE_JPEG2000_DCINEMA_4K.
const FfProfileJpeg2000Dcinema4K int = C.FF_PROFILE_JPEG2000_DCINEMA_4K

// FfProfileVp90 wraps FF_PROFILE_VP9_0.
const FfProfileVp90 int = C.FF_PROFILE_VP9_0

// FfProfileVp91 wraps FF_PROFILE_VP9_1.
const FfProfileVp91 int = C.FF_PROFILE_VP9_1

// FfProfileVp92 wraps FF_PROFILE_VP9_2.
const FfProfileVp92 int = C.FF_PROFILE_VP9_2

// FfProfileVp93 wraps FF_PROFILE_VP9_3.
const FfProfileVp93 int = C.FF_PROFILE_VP9_3

// FfProfileHevcMain wraps FF_PROFILE_HEVC_MAIN.
const FfProfileHevcMain int = C.FF_PROFILE_HEVC_MAIN

// FfProfileHevcMain10 wraps FF_PROFILE_HEVC_MAIN_10.
const FfProfileHevcMain10 int = C.FF_PROFILE_HEVC_MAIN_10

// FfProfileHevcMainStillPicture wraps FF_PROFILE_HEVC_MAIN_STILL_PICTURE.
const FfProfileHevcMainStillPicture int = C.FF_PROFILE_HEVC_MAIN_STILL_PICTURE

// FfProfileHevcRext wraps FF_PROFILE_HEVC_REXT.
const FfProfileHevcRext int = C.FF_PROFILE_HEVC_REXT

// FfProfileVvcMain10 wraps FF_PROFILE_VVC_MAIN_10.
const FfProfileVvcMain10 int = C.FF_PROFILE_VVC_MAIN_10

// FfProfileVvcMain10444 wraps FF_PROFILE_VVC_MAIN_10_444.
const FfProfileVvcMain10444 int = C.FF_PROFILE_VVC_MAIN_10_444

// FfProfileAv1Main wraps FF_PROFILE_AV1_MAIN.
const FfProfileAv1Main int = C.FF_PROFILE_AV1_MAIN

// FfProfileAv1High wraps FF_PROFILE_AV1_HIGH.
const FfProfileAv1High int = C.FF_PROFILE_AV1_HIGH

// FfProfileAv1Professional wraps FF_PROFILE_AV1_PROFESSIONAL.
const FfProfileAv1Professional int = C.FF_PROFILE_AV1_PROFESSIONAL

// FfProfileMjpegHuffmanBaselineDct wraps FF_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT.
const FfProfileMjpegHuffmanBaselineDct int = C.FF_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT

// FfProfileMjpegHuffmanExtendedSequentialDct wraps FF_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT.
const FfProfileMjpegHuffmanExtendedSequentialDct int = C.FF_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT

// FfProfileMjpegHuffmanProgressiveDct wraps FF_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT.
const FfProfileMjpegHuffmanProgressiveDct int = C.FF_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT

// FfProfileMjpegHuffmanLossless wraps FF_PROFILE_MJPEG_HUFFMAN_LOSSLESS.
const FfProfileMjpegHuffmanLossless int = C.FF_PROFILE_MJPEG_HUFFMAN_LOSSLESS

// FfProfileMjpegJpegLs wraps FF_PROFILE_MJPEG_JPEG_LS.
const FfProfileMjpegJpegLs int = C.FF_PROFILE_MJPEG_JPEG_LS

// FfProfileSbcMsbc wraps FF_PROFILE_SBC_MSBC.
const FfProfileSbcMsbc int = C.FF_PROFILE_SBC_MSBC

// FfProfileProresProxy wraps FF_PROFILE_PRORES_PROXY.
const FfProfileProresProxy int = C.FF_PROFILE_PRORES_PROXY

// FfProfileProresLt wraps FF_PROFILE_PRORES_LT.
const FfProfileProresLt int = C.FF_PROFILE_PRORES_LT

// FfProfileProresStandard wraps FF_PROFILE_PRORES_STANDARD.
const FfProfileProresStandard int = C.FF_PROFILE_PRORES_STANDARD

// FfProfileProresHq wraps FF_PROFILE_PRORES_HQ.
const FfProfileProresHq int = C.FF_PROFILE_PRORES_HQ

// FfProfileProres4444 wraps FF_PROFILE_PRORES_4444.
const FfProfileProres4444 int = C.FF_PROFILE_PRORES_4444

// FfProfileProresXq wraps FF_PROFILE_PRORES_XQ.
const FfProfileProresXq int = C.FF_PROFILE_PRORES_XQ

// FfProfileAribProfileA wraps FF_PROFILE_ARIB_PROFILE_A.
const FfProfileAribProfileA int = C.FF_PROFILE_ARIB_PROFILE_A

// FfProfileAribProfileC wraps FF_PROFILE_ARIB_PROFILE_C.
const FfProfileAribProfileC int = C.FF_PROFILE_ARIB_PROFILE_C

// FfProfileKlvaSync wraps FF_PROFILE_KLVA_SYNC.
const FfProfileKlvaSync int = C.FF_PROFILE_KLVA_SYNC

// FfProfileKlvaAsync wraps FF_PROFILE_KLVA_ASYNC.
const FfProfileKlvaAsync int = C.FF_PROFILE_KLVA_ASYNC

// FfLevelUnknown wraps FF_LEVEL_UNKNOWN.
const FfLevelUnknown int = C.FF_LEVEL_UNKNOWN

// FfSubCharencModeDoNothing wraps FF_SUB_CHARENC_MODE_DO_NOTHING.
const FfSubCharencModeDoNothing int = C.FF_SUB_CHARENC_MODE_DO_NOTHING

// FfSubCharencModeAutomatic wraps FF_SUB_CHARENC_MODE_AUTOMATIC.
const FfSubCharencModeAutomatic int = C.FF_SUB_CHARENC_MODE_AUTOMATIC

// FfSubCharencModePreDecoder wraps FF_SUB_CHARENC_MODE_PRE_DECODER.
const FfSubCharencModePreDecoder int = C.FF_SUB_CHARENC_MODE_PRE_DECODER

// FfSubCharencModeIgnore wraps FF_SUB_CHARENC_MODE_IGNORE.
const FfSubCharencModeIgnore int = C.FF_SUB_CHARENC_MODE_IGNORE

// FfCodecPropertyLossless wraps FF_CODEC_PROPERTY_LOSSLESS.
const FfCodecPropertyLossless int = C.FF_CODEC_PROPERTY_LOSSLESS

// FfCodecPropertyClosedCaptions wraps FF_CODEC_PROPERTY_CLOSED_CAPTIONS.
const FfCodecPropertyClosedCaptions int = C.FF_CODEC_PROPERTY_CLOSED_CAPTIONS

// FfCodecPropertyFilmGrain wraps FF_CODEC_PROPERTY_FILM_GRAIN.
const FfCodecPropertyFilmGrain int = C.FF_CODEC_PROPERTY_FILM_GRAIN

// AVHwaccelCodecCapExperimental wraps AV_HWACCEL_CODEC_CAP_EXPERIMENTAL.
const AVHwaccelCodecCapExperimental int = C.AV_HWACCEL_CODEC_CAP_EXPERIMENTAL

// AVHwaccelFlagIgnoreLevel wraps AV_HWACCEL_FLAG_IGNORE_LEVEL.
const AVHwaccelFlagIgnoreLevel int = C.AV_HWACCEL_FLAG_IGNORE_LEVEL

// AVHwaccelFlagAllowHighDepth wraps AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH.
const AVHwaccelFlagAllowHighDepth int = C.AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH

// AVHwaccelFlagAllowProfileMismatch wraps AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH.
const AVHwaccelFlagAllowProfileMismatch int = C.AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH

// AVHwaccelFlagUnsafeOutput wraps AV_HWACCEL_FLAG_UNSAFE_OUTPUT.
const AVHwaccelFlagUnsafeOutput int = C.AV_HWACCEL_FLAG_UNSAFE_OUTPUT

// AVSubtitleFlagForced wraps AV_SUBTITLE_FLAG_FORCED.
const AVSubtitleFlagForced int = C.AV_SUBTITLE_FLAG_FORCED

// AVParserPtsNb wraps AV_PARSER_PTS_NB.
const AVParserPtsNb int = C.AV_PARSER_PTS_NB

// ParserFlagCompleteFrames wraps PARSER_FLAG_COMPLETE_FRAMES.
const ParserFlagCompleteFrames int = C.PARSER_FLAG_COMPLETE_FRAMES

// ParserFlagOnce wraps PARSER_FLAG_ONCE.
const ParserFlagOnce int = C.PARSER_FLAG_ONCE

// ParserFlagFetchedOffset wraps PARSER_FLAG_FETCHED_OFFSET.
const ParserFlagFetchedOffset int = C.PARSER_FLAG_FETCHED_OFFSET

// ParserFlagUseCodecTs wraps PARSER_FLAG_USE_CODEC_TS.
const ParserFlagUseCodecTs int = C.PARSER_FLAG_USE_CODEC_TS

// AVCodecCapDrawHorizBand wraps AV_CODEC_CAP_DRAW_HORIZ_BAND.
const AVCodecCapDrawHorizBand int = C.AV_CODEC_CAP_DRAW_HORIZ_BAND

// AVCodecCapDr1 wraps AV_CODEC_CAP_DR1.
const AVCodecCapDr1 int = C.AV_CODEC_CAP_DR1

// AVCodecCapDelay wraps AV_CODEC_CAP_DELAY.
const AVCodecCapDelay int = C.AV_CODEC_CAP_DELAY

// AVCodecCapSmallLastFrame wraps AV_CODEC_CAP_SMALL_LAST_FRAME.
const AVCodecCapSmallLastFrame int = C.AV_CODEC_CAP_SMALL_LAST_FRAME

// AVCodecCapSubframes wraps AV_CODEC_CAP_SUBFRAMES.
const AVCodecCapSubframes int = C.AV_CODEC_CAP_SUBFRAMES

// AVCodecCapExperimental wraps AV_CODEC_CAP_EXPERIMENTAL.
const AVCodecCapExperimental int = C.AV_CODEC_CAP_EXPERIMENTAL

// AVCodecCapChannelConf wraps AV_CODEC_CAP_CHANNEL_CONF.
const AVCodecCapChannelConf int = C.AV_CODEC_CAP_CHANNEL_CONF

// AVCodecCapFrameThreads wraps AV_CODEC_CAP_FRAME_THREADS.
const AVCodecCapFrameThreads int = C.AV_CODEC_CAP_FRAME_THREADS

// AVCodecCapSliceThreads wraps AV_CODEC_CAP_SLICE_THREADS.
const AVCodecCapSliceThreads int = C.AV_CODEC_CAP_SLICE_THREADS

// AVCodecCapParamChange wraps AV_CODEC_CAP_PARAM_CHANGE.
const AVCodecCapParamChange int = C.AV_CODEC_CAP_PARAM_CHANGE

// AVCodecCapOtherThreads wraps AV_CODEC_CAP_OTHER_THREADS.
const AVCodecCapOtherThreads int = C.AV_CODEC_CAP_OTHER_THREADS

// AVCodecCapVariableFrameSize wraps AV_CODEC_CAP_VARIABLE_FRAME_SIZE.
const AVCodecCapVariableFrameSize int = C.AV_CODEC_CAP_VARIABLE_FRAME_SIZE

// AVCodecCapAvoidProbing wraps AV_CODEC_CAP_AVOID_PROBING.
const AVCodecCapAvoidProbing int = C.AV_CODEC_CAP_AVOID_PROBING

// AVCodecCapHardware wraps AV_CODEC_CAP_HARDWARE.
const AVCodecCapHardware int = C.AV_CODEC_CAP_HARDWARE

// AVCodecCapHybrid wraps AV_CODEC_CAP_HYBRID.
const AVCodecCapHybrid int = C.AV_CODEC_CAP_HYBRID

// AVCodecCapEncoderReorderedOpaque wraps AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE.
const AVCodecCapEncoderReorderedOpaque int = C.AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE

// AVCodecCapEncoderFlush wraps AV_CODEC_CAP_ENCODER_FLUSH.
const AVCodecCapEncoderFlush int = C.AV_CODEC_CAP_ENCODER_FLUSH

// AVCodecCapEncoderReconFrame wraps AV_CODEC_CAP_ENCODER_RECON_FRAME.
const AVCodecCapEncoderReconFrame int = C.AV_CODEC_CAP_ENCODER_RECON_FRAME

// AVCodecPropIntraOnly wraps AV_CODEC_PROP_INTRA_ONLY.
const AVCodecPropIntraOnly int = C.AV_CODEC_PROP_INTRA_ONLY

// AVCodecPropLossy wraps AV_CODEC_PROP_LOSSY.
const AVCodecPropLossy int = C.AV_CODEC_PROP_LOSSY

// AVCodecPropLossless wraps AV_CODEC_PROP_LOSSLESS.
const AVCodecPropLossless int = C.AV_CODEC_PROP_LOSSLESS

// AVCodecPropReorder wraps AV_CODEC_PROP_REORDER.
const AVCodecPropReorder int = C.AV_CODEC_PROP_REORDER

// AVCodecPropBitmapSub wraps AV_CODEC_PROP_BITMAP_SUB.
const AVCodecPropBitmapSub int = C.AV_CODEC_PROP_BITMAP_SUB

// AVCodecPropTextSub wraps AV_CODEC_PROP_TEXT_SUB.
const AVCodecPropTextSub int = C.AV_CODEC_PROP_TEXT_SUB

// AVCodecIdIffByterun1 wraps AV_CODEC_ID_IFF_BYTERUN1.
const AVCodecIdIffByterun1 int = C.AV_CODEC_ID_IFF_BYTERUN1

// AVCodecIdH265 wraps AV_CODEC_ID_H265.
const AVCodecIdH265 int = C.AV_CODEC_ID_H265

// AVCodecIdH266 wraps AV_CODEC_ID_H266.
const AVCodecIdH266 int = C.AV_CODEC_ID_H266

// AVInputBufferPaddingSize wraps AV_INPUT_BUFFER_PADDING_SIZE.
const AVInputBufferPaddingSize int = C.AV_INPUT_BUFFER_PADDING_SIZE

// AVEfCrccheck wraps AV_EF_CRCCHECK.
const AVEfCrccheck int = C.AV_EF_CRCCHECK

// AVEfBitstream wraps AV_EF_BITSTREAM.
const AVEfBitstream int = C.AV_EF_BITSTREAM

// AVEfBuffer wraps AV_EF_BUFFER.
const AVEfBuffer int = C.AV_EF_BUFFER

// AVEfExplode wraps AV_EF_EXPLODE.
const AVEfExplode int = C.AV_EF_EXPLODE

// AVEfIgnoreErr wraps AV_EF_IGNORE_ERR.
const AVEfIgnoreErr int = C.AV_EF_IGNORE_ERR

// AVEfCareful wraps AV_EF_CAREFUL.
const AVEfCareful int = C.AV_EF_CAREFUL

// AVEfCompliant wraps AV_EF_COMPLIANT.
const AVEfCompliant int = C.AV_EF_COMPLIANT

// AVEfAggressive wraps AV_EF_AGGRESSIVE.
const AVEfAggressive int = C.AV_EF_AGGRESSIVE

// FfComplianceVeryStrict wraps FF_COMPLIANCE_VERY_STRICT.
const FfComplianceVeryStrict int = C.FF_COMPLIANCE_VERY_STRICT

// FfComplianceStrict wraps FF_COMPLIANCE_STRICT.
const FfComplianceStrict int = C.FF_COMPLIANCE_STRICT

// FfComplianceNormal wraps FF_COMPLIANCE_NORMAL.
const FfComplianceNormal int = C.FF_COMPLIANCE_NORMAL

// FfComplianceUnofficial wraps FF_COMPLIANCE_UNOFFICIAL.
const FfComplianceUnofficial int = C.FF_COMPLIANCE_UNOFFICIAL

// FfComplianceExperimental wraps FF_COMPLIANCE_EXPERIMENTAL.
const FfComplianceExperimental int = C.FF_COMPLIANCE_EXPERIMENTAL

// AVPktDataQualityFactor wraps AV_PKT_DATA_QUALITY_FACTOR.
const AVPktDataQualityFactor int = C.AV_PKT_DATA_QUALITY_FACTOR

// AVPktFlagKey wraps AV_PKT_FLAG_KEY.
const AVPktFlagKey int = C.AV_PKT_FLAG_KEY

// AVPktFlagCorrupt wraps AV_PKT_FLAG_CORRUPT.
const AVPktFlagCorrupt int = C.AV_PKT_FLAG_CORRUPT

// AVPktFlagDiscard wraps AV_PKT_FLAG_DISCARD.
const AVPktFlagDiscard int = C.AV_PKT_FLAG_DISCARD

// AVPktFlagTrusted wraps AV_PKT_FLAG_TRUSTED.
const AVPktFlagTrusted int = C.AV_PKT_FLAG_TRUSTED

// AVPktFlagDisposable wraps AV_PKT_FLAG_DISPOSABLE.
const AVPktFlagDisposable int = C.AV_PKT_FLAG_DISPOSABLE

// AVFilterFlagDynamicInputs wraps AVFILTER_FLAG_DYNAMIC_INPUTS.
const AVFilterFlagDynamicInputs int = C.AVFILTER_FLAG_DYNAMIC_INPUTS

// AVFilterFlagDynamicOutputs wraps AVFILTER_FLAG_DYNAMIC_OUTPUTS.
const AVFilterFlagDynamicOutputs int = C.AVFILTER_FLAG_DYNAMIC_OUTPUTS

// AVFilterFlagSliceThreads wraps AVFILTER_FLAG_SLICE_THREADS.
const AVFilterFlagSliceThreads int = C.AVFILTER_FLAG_SLICE_THREADS

// AVFilterFlagMetadataOnly wraps AVFILTER_FLAG_METADATA_ONLY.
const AVFilterFlagMetadataOnly int = C.AVFILTER_FLAG_METADATA_ONLY

// AVFilterFlagSupportTimelineGeneric wraps AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC.
const AVFilterFlagSupportTimelineGeneric int = C.AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC

// AVFilterFlagSupportTimelineInternal wraps AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL.
const AVFilterFlagSupportTimelineInternal int = C.AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL

// AVFilterFlagSupportTimeline wraps AVFILTER_FLAG_SUPPORT_TIMELINE.
const AVFilterFlagSupportTimeline int = C.AVFILTER_FLAG_SUPPORT_TIMELINE

// AVFilterThreadSlice wraps AVFILTER_THREAD_SLICE.
const AVFilterThreadSlice int = C.AVFILTER_THREAD_SLICE

// AVFilterCmdFlagOne wraps AVFILTER_CMD_FLAG_ONE.
const AVFilterCmdFlagOne int = C.AVFILTER_CMD_FLAG_ONE

// AVFilterCmdFlagFast wraps AVFILTER_CMD_FLAG_FAST.
const AVFilterCmdFlagFast int = C.AVFILTER_CMD_FLAG_FAST

// AVBuffersinkFlagPeek wraps AV_BUFFERSINK_FLAG_PEEK.
const AVBuffersinkFlagPeek int = C.AV_BUFFERSINK_FLAG_PEEK

// AVBuffersinkFlagNoRequest wraps AV_BUFFERSINK_FLAG_NO_REQUEST.
const AVBuffersinkFlagNoRequest int = C.AV_BUFFERSINK_FLAG_NO_REQUEST

// AVProbeScoreRetry wraps AVPROBE_SCORE_RETRY.
const AVProbeScoreRetry int = C.AVPROBE_SCORE_RETRY

// AVProbeScoreStreamRetry wraps AVPROBE_SCORE_STREAM_RETRY.
const AVProbeScoreStreamRetry int = C.AVPROBE_SCORE_STREAM_RETRY

// AVProbeScoreExtension wraps AVPROBE_SCORE_EXTENSION.
const AVProbeScoreExtension int = C.AVPROBE_SCORE_EXTENSION

// AVProbeScoreMime wraps AVPROBE_SCORE_MIME.
const AVProbeScoreMime int = C.AVPROBE_SCORE_MIME

// AVProbeScoreMax wraps AVPROBE_SCORE_MAX.
const AVProbeScoreMax int = C.AVPROBE_SCORE_MAX

// AVProbePaddingSize wraps AVPROBE_PADDING_SIZE.
const AVProbePaddingSize int = C.AVPROBE_PADDING_SIZE

// AVFmtNofile wraps AVFMT_NOFILE.
const AVFmtNofile int = C.AVFMT_NOFILE

// AVFmtNeednumber wraps AVFMT_NEEDNUMBER.
const AVFmtNeednumber int = C.AVFMT_NEEDNUMBER

// AVFmtExperimental wraps AVFMT_EXPERIMENTAL.
const AVFmtExperimental int = C.AVFMT_EXPERIMENTAL

// AVFmtShowIds wraps AVFMT_SHOW_IDS.
const AVFmtShowIds int = C.AVFMT_SHOW_IDS

// AVFmtGlobalheader wraps AVFMT_GLOBALHEADER.
const AVFmtGlobalheader int = C.AVFMT_GLOBALHEADER

// AVFmtNotimestamps wraps AVFMT_NOTIMESTAMPS.
const AVFmtNotimestamps int = C.AVFMT_NOTIMESTAMPS

// AVFmtGenericIndex wraps AVFMT_GENERIC_INDEX.
const AVFmtGenericIndex int = C.AVFMT_GENERIC_INDEX

// AVFmtTsDiscont wraps AVFMT_TS_DISCONT.
const AVFmtTsDiscont int = C.AVFMT_TS_DISCONT

// AVFmtVariableFps wraps AVFMT_VARIABLE_FPS.
const AVFmtVariableFps int = C.AVFMT_VARIABLE_FPS

// AVFmtNodimensions wraps AVFMT_NODIMENSIONS.
const AVFmtNodimensions int = C.AVFMT_NODIMENSIONS

// AVFmtNostreams wraps AVFMT_NOSTREAMS.
const AVFmtNostreams int = C.AVFMT_NOSTREAMS

// AVFmtNobinsearch wraps AVFMT_NOBINSEARCH.
const AVFmtNobinsearch int = C.AVFMT_NOBINSEARCH

// AVFmtNogensearch wraps AVFMT_NOGENSEARCH.
const AVFmtNogensearch int = C.AVFMT_NOGENSEARCH

// AVFmtNoByteSeek wraps AVFMT_NO_BYTE_SEEK.
const AVFmtNoByteSeek int = C.AVFMT_NO_BYTE_SEEK

// AVFmtAllowFlush wraps AVFMT_ALLOW_FLUSH.
const AVFmtAllowFlush int = C.AVFMT_ALLOW_FLUSH

// AVFmtTsNonstrict wraps AVFMT_TS_NONSTRICT.
const AVFmtTsNonstrict int = C.AVFMT_TS_NONSTRICT

// AVFmtTsNegative wraps AVFMT_TS_NEGATIVE.
const AVFmtTsNegative int = C.AVFMT_TS_NEGATIVE

// AVFmtSeekToPts wraps AVFMT_SEEK_TO_PTS.
const AVFmtSeekToPts int = C.AVFMT_SEEK_TO_PTS

// AVIndexKeyframe wraps AVINDEX_KEYFRAME.
const AVIndexKeyframe int = C.AVINDEX_KEYFRAME

// AVIndexDiscardFrame wraps AVINDEX_DISCARD_FRAME.
const AVIndexDiscardFrame int = C.AVINDEX_DISCARD_FRAME

// AVDispositionDefault wraps AV_DISPOSITION_DEFAULT.
const AVDispositionDefault int = C.AV_DISPOSITION_DEFAULT

// AVDispositionDub wraps AV_DISPOSITION_DUB.
const AVDispositionDub int = C.AV_DISPOSITION_DUB

// AVDispositionOriginal wraps AV_DISPOSITION_ORIGINAL.
const AVDispositionOriginal int = C.AV_DISPOSITION_ORIGINAL

// AVDispositionComment wraps AV_DISPOSITION_COMMENT.
const AVDispositionComment int = C.AV_DISPOSITION_COMMENT

// AVDispositionLyrics wraps AV_DISPOSITION_LYRICS.
const AVDispositionLyrics int = C.AV_DISPOSITION_LYRICS

// AVDispositionKaraoke wraps AV_DISPOSITION_KARAOKE.
const AVDispositionKaraoke int = C.AV_DISPOSITION_KARAOKE

// AVDispositionForced wraps AV_DISPOSITION_FORCED.
const AVDispositionForced int = C.AV_DISPOSITION_FORCED

// AVDispositionHearingImpaired wraps AV_DISPOSITION_HEARING_IMPAIRED.
const AVDispositionHearingImpaired int = C.AV_DISPOSITION_HEARING_IMPAIRED

// AVDispositionVisualImpaired wraps AV_DISPOSITION_VISUAL_IMPAIRED.
const AVDispositionVisualImpaired int = C.AV_DISPOSITION_VISUAL_IMPAIRED

// AVDispositionCleanEffects wraps AV_DISPOSITION_CLEAN_EFFECTS.
const AVDispositionCleanEffects int = C.AV_DISPOSITION_CLEAN_EFFECTS

// AVDispositionAttachedPic wraps AV_DISPOSITION_ATTACHED_PIC.
const AVDispositionAttachedPic int = C.AV_DISPOSITION_ATTACHED_PIC

// AVDispositionTimedThumbnails wraps AV_DISPOSITION_TIMED_THUMBNAILS.
const AVDispositionTimedThumbnails int = C.AV_DISPOSITION_TIMED_THUMBNAILS

// AVDispositionNonDiegetic wraps AV_DISPOSITION_NON_DIEGETIC.
const AVDispositionNonDiegetic int = C.AV_DISPOSITION_NON_DIEGETIC

// AVDispositionCaptions wraps AV_DISPOSITION_CAPTIONS.
const AVDispositionCaptions int = C.AV_DISPOSITION_CAPTIONS

// AVDispositionDescriptions wraps AV_DISPOSITION_DESCRIPTIONS.
const AVDispositionDescriptions int = C.AV_DISPOSITION_DESCRIPTIONS

// AVDispositionMetadata wraps AV_DISPOSITION_METADATA.
const AVDispositionMetadata int = C.AV_DISPOSITION_METADATA

// AVDispositionDependent wraps AV_DISPOSITION_DEPENDENT.
const AVDispositionDependent int = C.AV_DISPOSITION_DEPENDENT

// AVDispositionStillImage wraps AV_DISPOSITION_STILL_IMAGE.
const AVDispositionStillImage int = C.AV_DISPOSITION_STILL_IMAGE

// AVPtsWrapIgnore wraps AV_PTS_WRAP_IGNORE.
const AVPtsWrapIgnore int = C.AV_PTS_WRAP_IGNORE

// AVPtsWrapAddOffset wraps AV_PTS_WRAP_ADD_OFFSET.
const AVPtsWrapAddOffset int = C.AV_PTS_WRAP_ADD_OFFSET

// AVPtsWrapSubOffset wraps AV_PTS_WRAP_SUB_OFFSET.
const AVPtsWrapSubOffset int = C.AV_PTS_WRAP_SUB_OFFSET

// AVStreamEventFlagMetadataUpdated wraps AVSTREAM_EVENT_FLAG_METADATA_UPDATED.
const AVStreamEventFlagMetadataUpdated int = C.AVSTREAM_EVENT_FLAG_METADATA_UPDATED

// AVStreamEventFlagNewPackets wraps AVSTREAM_EVENT_FLAG_NEW_PACKETS.
const AVStreamEventFlagNewPackets int = C.AVSTREAM_EVENT_FLAG_NEW_PACKETS

// AVProgramRunning wraps AV_PROGRAM_RUNNING.
const AVProgramRunning int = C.AV_PROGRAM_RUNNING

// AVFmtctxNoheader wraps AVFMTCTX_NOHEADER.
const AVFmtctxNoheader int = C.AVFMTCTX_NOHEADER

// AVFmtctxUnseekable wraps AVFMTCTX_UNSEEKABLE.
const AVFmtctxUnseekable int = C.AVFMTCTX_UNSEEKABLE

// AVFmtFlagGenpts wraps AVFMT_FLAG_GENPTS.
const AVFmtFlagGenpts int = C.AVFMT_FLAG_GENPTS

// AVFmtFlagIgnidx wraps AVFMT_FLAG_IGNIDX.
const AVFmtFlagIgnidx int = C.AVFMT_FLAG_IGNIDX

// AVFmtFlagNonblock wraps AVFMT_FLAG_NONBLOCK.
const AVFmtFlagNonblock int = C.AVFMT_FLAG_NONBLOCK

// AVFmtFlagIgndts wraps AVFMT_FLAG_IGNDTS.
const AVFmtFlagIgndts int = C.AVFMT_FLAG_IGNDTS

// AVFmtFlagNofillin wraps AVFMT_FLAG_NOFILLIN.
const AVFmtFlagNofillin int = C.AVFMT_FLAG_NOFILLIN

// AVFmtFlagNoparse wraps AVFMT_FLAG_NOPARSE.
const AVFmtFlagNoparse int = C.AVFMT_FLAG_NOPARSE

// AVFmtFlagNobuffer wraps AVFMT_FLAG_NOBUFFER.
const AVFmtFlagNobuffer int = C.AVFMT_FLAG_NOBUFFER

// AVFmtFlagCustomIo wraps AVFMT_FLAG_CUSTOM_IO.
const AVFmtFlagCustomIo int = C.AVFMT_FLAG_CUSTOM_IO

// AVFmtFlagDiscardCorrupt wraps AVFMT_FLAG_DISCARD_CORRUPT.
const AVFmtFlagDiscardCorrupt int = C.AVFMT_FLAG_DISCARD_CORRUPT

// AVFmtFlagFlushPackets wraps AVFMT_FLAG_FLUSH_PACKETS.
const AVFmtFlagFlushPackets int = C.AVFMT_FLAG_FLUSH_PACKETS

// AVFmtFlagBitexact wraps AVFMT_FLAG_BITEXACT.
const AVFmtFlagBitexact int = C.AVFMT_FLAG_BITEXACT

// AVFmtFlagSortDts wraps AVFMT_FLAG_SORT_DTS.
const AVFmtFlagSortDts int = C.AVFMT_FLAG_SORT_DTS

// AVFmtFlagFastSeek wraps AVFMT_FLAG_FAST_SEEK.
const AVFmtFlagFastSeek int = C.AVFMT_FLAG_FAST_SEEK

// AVFmtFlagShortest wraps AVFMT_FLAG_SHORTEST.
const AVFmtFlagShortest int = C.AVFMT_FLAG_SHORTEST

// AVFmtFlagAutoBsf wraps AVFMT_FLAG_AUTO_BSF.
const AVFmtFlagAutoBsf int = C.AVFMT_FLAG_AUTO_BSF

// FfFdebugTs wraps FF_FDEBUG_TS.
const FfFdebugTs int = C.FF_FDEBUG_TS

// AVFmtEventFlagMetadataUpdated wraps AVFMT_EVENT_FLAG_METADATA_UPDATED.
const AVFmtEventFlagMetadataUpdated int = C.AVFMT_EVENT_FLAG_METADATA_UPDATED

// AVFmtAvoidNegTsAuto wraps AVFMT_AVOID_NEG_TS_AUTO.
const AVFmtAvoidNegTsAuto int = C.AVFMT_AVOID_NEG_TS_AUTO

// AVFmtAvoidNegTsDisabled wraps AVFMT_AVOID_NEG_TS_DISABLED.
const AVFmtAvoidNegTsDisabled int = C.AVFMT_AVOID_NEG_TS_DISABLED

// AVFmtAvoidNegTsMakeNonNegative wraps AVFMT_AVOID_NEG_TS_MAKE_NON_NEGATIVE.
const AVFmtAvoidNegTsMakeNonNegative int = C.AVFMT_AVOID_NEG_TS_MAKE_NON_NEGATIVE

// AVFmtAvoidNegTsMakeZero wraps AVFMT_AVOID_NEG_TS_MAKE_ZERO.
const AVFmtAvoidNegTsMakeZero int = C.AVFMT_AVOID_NEG_TS_MAKE_ZERO

// AVSeekFlagBackward wraps AVSEEK_FLAG_BACKWARD.
const AVSeekFlagBackward int = C.AVSEEK_FLAG_BACKWARD

// AVSeekFlagByte wraps AVSEEK_FLAG_BYTE.
const AVSeekFlagByte int = C.AVSEEK_FLAG_BYTE

// AVSeekFlagAny wraps AVSEEK_FLAG_ANY.
const AVSeekFlagAny int = C.AVSEEK_FLAG_ANY

// AVSeekFlagFrame wraps AVSEEK_FLAG_FRAME.
const AVSeekFlagFrame int = C.AVSEEK_FLAG_FRAME

// AVStreamInitInWriteHeader wraps AVSTREAM_INIT_IN_WRITE_HEADER.
const AVStreamInitInWriteHeader int = C.AVSTREAM_INIT_IN_WRITE_HEADER

// AVStreamInitInInitOutput wraps AVSTREAM_INIT_IN_INIT_OUTPUT.
const AVStreamInitInInitOutput int = C.AVSTREAM_INIT_IN_INIT_OUTPUT

// AVFrameFilenameFlagsMultiple wraps AV_FRAME_FILENAME_FLAGS_MULTIPLE.
const AVFrameFilenameFlagsMultiple int = C.AV_FRAME_FILENAME_FLAGS_MULTIPLE

// AVIOSeekableNormal wraps AVIO_SEEKABLE_NORMAL.
const AVIOSeekableNormal int = C.AVIO_SEEKABLE_NORMAL

// AVIOSeekableTime wraps AVIO_SEEKABLE_TIME.
const AVIOSeekableTime int = C.AVIO_SEEKABLE_TIME

// AVSeekSize wraps AVSEEK_SIZE.
const AVSeekSize int = C.AVSEEK_SIZE

// AVSeekForce wraps AVSEEK_FORCE.
const AVSeekForce int = C.AVSEEK_FORCE

// AVIOFlagRead wraps AVIO_FLAG_READ.
const AVIOFlagRead int = C.AVIO_FLAG_READ

// AVIOFlagWrite wraps AVIO_FLAG_WRITE.
const AVIOFlagWrite int = C.AVIO_FLAG_WRITE

// AVIOFlagReadWrite wraps AVIO_FLAG_READ_WRITE.
const AVIOFlagReadWrite int = C.AVIO_FLAG_READ_WRITE

// AVIOFlagNonblock wraps AVIO_FLAG_NONBLOCK.
const AVIOFlagNonblock int = C.AVIO_FLAG_NONBLOCK

// AVIOFlagDirect wraps AVIO_FLAG_DIRECT.
const AVIOFlagDirect int = C.AVIO_FLAG_DIRECT

// FfLambdaShift wraps FF_LAMBDA_SHIFT.
const FfLambdaShift int = C.FF_LAMBDA_SHIFT

// FfLambdaScale wraps FF_LAMBDA_SCALE.
const FfLambdaScale int = C.FF_LAMBDA_SCALE

// FfQp2Lambda wraps FF_QP2LAMBDA.
const FfQp2Lambda int = C.FF_QP2LAMBDA

// FfLambdaMax wraps FF_LAMBDA_MAX.
const FfLambdaMax int = C.FF_LAMBDA_MAX

// FfQualityScale wraps FF_QUALITY_SCALE.
const FfQualityScale int = C.FF_QUALITY_SCALE

// AVNoptsValue wraps AV_NOPTS_VALUE.
const AVNoptsValue int = C.AV_NOPTS_VALUE

// AVTimeBase wraps AV_TIME_BASE.
const AVTimeBase int = C.AV_TIME_BASE

// AVFourccMaxStringSize wraps AV_FOURCC_MAX_STRING_SIZE.
const AVFourccMaxStringSize int = C.AV_FOURCC_MAX_STRING_SIZE

// AVBufferFlagReadonly wraps AV_BUFFER_FLAG_READONLY.
const AVBufferFlagReadonly int = C.AV_BUFFER_FLAG_READONLY

// AVChFrontLeft wraps AV_CH_FRONT_LEFT.
const AVChFrontLeft int = C.AV_CH_FRONT_LEFT

// AVChFrontRight wraps AV_CH_FRONT_RIGHT.
const AVChFrontRight int = C.AV_CH_FRONT_RIGHT

// AVChFrontCenter wraps AV_CH_FRONT_CENTER.
const AVChFrontCenter int = C.AV_CH_FRONT_CENTER

// AVChLowFrequency wraps AV_CH_LOW_FREQUENCY.
const AVChLowFrequency int = C.AV_CH_LOW_FREQUENCY

// AVChBackLeft wraps AV_CH_BACK_LEFT.
const AVChBackLeft int = C.AV_CH_BACK_LEFT

// AVChBackRight wraps AV_CH_BACK_RIGHT.
const AVChBackRight int = C.AV_CH_BACK_RIGHT

// AVChFrontLeftOfCenter wraps AV_CH_FRONT_LEFT_OF_CENTER.
const AVChFrontLeftOfCenter int = C.AV_CH_FRONT_LEFT_OF_CENTER

// AVChFrontRightOfCenter wraps AV_CH_FRONT_RIGHT_OF_CENTER.
const AVChFrontRightOfCenter int = C.AV_CH_FRONT_RIGHT_OF_CENTER

// AVChBackCenter wraps AV_CH_BACK_CENTER.
const AVChBackCenter int = C.AV_CH_BACK_CENTER

// AVChSideLeft wraps AV_CH_SIDE_LEFT.
const AVChSideLeft int = C.AV_CH_SIDE_LEFT

// AVChSideRight wraps AV_CH_SIDE_RIGHT.
const AVChSideRight int = C.AV_CH_SIDE_RIGHT

// AVChTopCenter wraps AV_CH_TOP_CENTER.
const AVChTopCenter int = C.AV_CH_TOP_CENTER

// AVChTopFrontLeft wraps AV_CH_TOP_FRONT_LEFT.
const AVChTopFrontLeft int = C.AV_CH_TOP_FRONT_LEFT

// AVChTopFrontCenter wraps AV_CH_TOP_FRONT_CENTER.
const AVChTopFrontCenter int = C.AV_CH_TOP_FRONT_CENTER

// AVChTopFrontRight wraps AV_CH_TOP_FRONT_RIGHT.
const AVChTopFrontRight int = C.AV_CH_TOP_FRONT_RIGHT

// AVChTopBackLeft wraps AV_CH_TOP_BACK_LEFT.
const AVChTopBackLeft int = C.AV_CH_TOP_BACK_LEFT

// AVChTopBackCenter wraps AV_CH_TOP_BACK_CENTER.
const AVChTopBackCenter int = C.AV_CH_TOP_BACK_CENTER

// AVChTopBackRight wraps AV_CH_TOP_BACK_RIGHT.
const AVChTopBackRight int = C.AV_CH_TOP_BACK_RIGHT

// AVChStereoLeft wraps AV_CH_STEREO_LEFT.
const AVChStereoLeft int = C.AV_CH_STEREO_LEFT

// AVChStereoRight wraps AV_CH_STEREO_RIGHT.
const AVChStereoRight int = C.AV_CH_STEREO_RIGHT

// AVChWideLeft wraps AV_CH_WIDE_LEFT.
const AVChWideLeft int = C.AV_CH_WIDE_LEFT

// AVChWideRight wraps AV_CH_WIDE_RIGHT.
const AVChWideRight int = C.AV_CH_WIDE_RIGHT

// AVChSurroundDirectLeft wraps AV_CH_SURROUND_DIRECT_LEFT.
const AVChSurroundDirectLeft int = C.AV_CH_SURROUND_DIRECT_LEFT

// AVChSurroundDirectRight wraps AV_CH_SURROUND_DIRECT_RIGHT.
const AVChSurroundDirectRight int = C.AV_CH_SURROUND_DIRECT_RIGHT

// AVChLowFrequency2 wraps AV_CH_LOW_FREQUENCY_2.
const AVChLowFrequency2 int = C.AV_CH_LOW_FREQUENCY_2

// AVChTopSideLeft wraps AV_CH_TOP_SIDE_LEFT.
const AVChTopSideLeft int = C.AV_CH_TOP_SIDE_LEFT

// AVChTopSideRight wraps AV_CH_TOP_SIDE_RIGHT.
const AVChTopSideRight int = C.AV_CH_TOP_SIDE_RIGHT

// AVChBottomFrontCenter wraps AV_CH_BOTTOM_FRONT_CENTER.
const AVChBottomFrontCenter int = C.AV_CH_BOTTOM_FRONT_CENTER

// AVChBottomFrontLeft wraps AV_CH_BOTTOM_FRONT_LEFT.
const AVChBottomFrontLeft int = C.AV_CH_BOTTOM_FRONT_LEFT

// AVChBottomFrontRight wraps AV_CH_BOTTOM_FRONT_RIGHT.
const AVChBottomFrontRight int = C.AV_CH_BOTTOM_FRONT_RIGHT

// AVChLayoutMono wraps AV_CH_LAYOUT_MONO.
const AVChLayoutMono int = C.AV_CH_LAYOUT_MONO

// AVChLayoutStereo wraps AV_CH_LAYOUT_STEREO.
const AVChLayoutStereo int = C.AV_CH_LAYOUT_STEREO

// AVChLayout2Point1 wraps AV_CH_LAYOUT_2POINT1.
const AVChLayout2Point1 int = C.AV_CH_LAYOUT_2POINT1

// AVChLayout21 wraps AV_CH_LAYOUT_2_1.
const AVChLayout21 int = C.AV_CH_LAYOUT_2_1

// AVChLayoutSurround wraps AV_CH_LAYOUT_SURROUND.
const AVChLayoutSurround int = C.AV_CH_LAYOUT_SURROUND

// AVChLayout3Point1 wraps AV_CH_LAYOUT_3POINT1.
const AVChLayout3Point1 int = C.AV_CH_LAYOUT_3POINT1

// AVChLayout4Point0 wraps AV_CH_LAYOUT_4POINT0.
const AVChLayout4Point0 int = C.AV_CH_LAYOUT_4POINT0

// AVChLayout4Point1 wraps AV_CH_LAYOUT_4POINT1.
const AVChLayout4Point1 int = C.AV_CH_LAYOUT_4POINT1

// AVChLayout22 wraps AV_CH_LAYOUT_2_2.
const AVChLayout22 int = C.AV_CH_LAYOUT_2_2

// AVChLayoutQuad wraps AV_CH_LAYOUT_QUAD.
const AVChLayoutQuad int = C.AV_CH_LAYOUT_QUAD

// AVChLayout5Point0 wraps AV_CH_LAYOUT_5POINT0.
const AVChLayout5Point0 int = C.AV_CH_LAYOUT_5POINT0

// AVChLayout5Point1 wraps AV_CH_LAYOUT_5POINT1.
const AVChLayout5Point1 int = C.AV_CH_LAYOUT_5POINT1

// AVChLayout5Point0Back wraps AV_CH_LAYOUT_5POINT0_BACK.
const AVChLayout5Point0Back int = C.AV_CH_LAYOUT_5POINT0_BACK

// AVChLayout5Point1Back wraps AV_CH_LAYOUT_5POINT1_BACK.
const AVChLayout5Point1Back int = C.AV_CH_LAYOUT_5POINT1_BACK

// AVChLayout6Point0 wraps AV_CH_LAYOUT_6POINT0.
const AVChLayout6Point0 int = C.AV_CH_LAYOUT_6POINT0

// AVChLayout6Point0Front wraps AV_CH_LAYOUT_6POINT0_FRONT.
const AVChLayout6Point0Front int = C.AV_CH_LAYOUT_6POINT0_FRONT

// AVChLayoutHexagonal wraps AV_CH_LAYOUT_HEXAGONAL.
const AVChLayoutHexagonal int = C.AV_CH_LAYOUT_HEXAGONAL

// AVChLayout6Point1 wraps AV_CH_LAYOUT_6POINT1.
const AVChLayout6Point1 int = C.AV_CH_LAYOUT_6POINT1

// AVChLayout6Point1Back wraps AV_CH_LAYOUT_6POINT1_BACK.
const AVChLayout6Point1Back int = C.AV_CH_LAYOUT_6POINT1_BACK

// AVChLayout6Point1Front wraps AV_CH_LAYOUT_6POINT1_FRONT.
const AVChLayout6Point1Front int = C.AV_CH_LAYOUT_6POINT1_FRONT

// AVChLayout7Point0 wraps AV_CH_LAYOUT_7POINT0.
const AVChLayout7Point0 int = C.AV_CH_LAYOUT_7POINT0

// AVChLayout7Point0Front wraps AV_CH_LAYOUT_7POINT0_FRONT.
const AVChLayout7Point0Front int = C.AV_CH_LAYOUT_7POINT0_FRONT

// AVChLayout7Point1 wraps AV_CH_LAYOUT_7POINT1.
const AVChLayout7Point1 int = C.AV_CH_LAYOUT_7POINT1

// AVChLayout7Point1Wide wraps AV_CH_LAYOUT_7POINT1_WIDE.
const AVChLayout7Point1Wide int = C.AV_CH_LAYOUT_7POINT1_WIDE

// AVChLayout7Point1WideBack wraps AV_CH_LAYOUT_7POINT1_WIDE_BACK.
const AVChLayout7Point1WideBack int = C.AV_CH_LAYOUT_7POINT1_WIDE_BACK

// AVChLayout7Point1TopBack wraps AV_CH_LAYOUT_7POINT1_TOP_BACK.
const AVChLayout7Point1TopBack int = C.AV_CH_LAYOUT_7POINT1_TOP_BACK

// AVChLayoutOctagonal wraps AV_CH_LAYOUT_OCTAGONAL.
const AVChLayoutOctagonal int = C.AV_CH_LAYOUT_OCTAGONAL

// AVChLayoutCube wraps AV_CH_LAYOUT_CUBE.
const AVChLayoutCube int = C.AV_CH_LAYOUT_CUBE

// AVChLayoutHexadecagonal wraps AV_CH_LAYOUT_HEXADECAGONAL.
const AVChLayoutHexadecagonal int = C.AV_CH_LAYOUT_HEXADECAGONAL

// AVChLayoutStereoDownmix wraps AV_CH_LAYOUT_STEREO_DOWNMIX.
const AVChLayoutStereoDownmix int = C.AV_CH_LAYOUT_STEREO_DOWNMIX

// AVChLayout22Point2 wraps AV_CH_LAYOUT_22POINT2.
const AVChLayout22Point2 int = C.AV_CH_LAYOUT_22POINT2

// AVDictMatchCase wraps AV_DICT_MATCH_CASE.
const AVDictMatchCase int = C.AV_DICT_MATCH_CASE

// AVDictIgnoreSuffix wraps AV_DICT_IGNORE_SUFFIX.
const AVDictIgnoreSuffix int = C.AV_DICT_IGNORE_SUFFIX

// AVDictDontStrdupKey wraps AV_DICT_DONT_STRDUP_KEY.
const AVDictDontStrdupKey int = C.AV_DICT_DONT_STRDUP_KEY

// AVDictDontStrdupVal wraps AV_DICT_DONT_STRDUP_VAL.
const AVDictDontStrdupVal int = C.AV_DICT_DONT_STRDUP_VAL

// AVDictDontOverwrite wraps AV_DICT_DONT_OVERWRITE.
const AVDictDontOverwrite int = C.AV_DICT_DONT_OVERWRITE

// AVDictAppend wraps AV_DICT_APPEND.
const AVDictAppend int = C.AV_DICT_APPEND

// AVDictMultikey wraps AV_DICT_MULTIKEY.
const AVDictMultikey int = C.AV_DICT_MULTIKEY

// AVNumDataPointers wraps AV_NUM_DATA_POINTERS.
const AVNumDataPointers int = C.AV_NUM_DATA_POINTERS

// AVFrameFlagCorrupt wraps AV_FRAME_FLAG_CORRUPT.
const AVFrameFlagCorrupt int = C.AV_FRAME_FLAG_CORRUPT

// AVFrameFlagDiscard wraps AV_FRAME_FLAG_DISCARD.
const AVFrameFlagDiscard int = C.AV_FRAME_FLAG_DISCARD

// FfDecodeErrorInvalidBitstream wraps FF_DECODE_ERROR_INVALID_BITSTREAM.
const FfDecodeErrorInvalidBitstream int = C.FF_DECODE_ERROR_INVALID_BITSTREAM

// FfDecodeErrorMissingReference wraps FF_DECODE_ERROR_MISSING_REFERENCE.
const FfDecodeErrorMissingReference int = C.FF_DECODE_ERROR_MISSING_REFERENCE

// FfDecodeErrorConcealmentActive wraps FF_DECODE_ERROR_CONCEALMENT_ACTIVE.
const FfDecodeErrorConcealmentActive int = C.FF_DECODE_ERROR_CONCEALMENT_ACTIVE

// FfDecodeErrorDecodeSlices wraps FF_DECODE_ERROR_DECODE_SLICES.
const FfDecodeErrorDecodeSlices int = C.FF_DECODE_ERROR_DECODE_SLICES

// AVLogQuiet wraps AV_LOG_QUIET.
const AVLogQuiet int = C.AV_LOG_QUIET

// AVLogPanic wraps AV_LOG_PANIC.
const AVLogPanic int = C.AV_LOG_PANIC

// AVLogFatal wraps AV_LOG_FATAL.
const AVLogFatal int = C.AV_LOG_FATAL

// AVLogError wraps AV_LOG_ERROR.
const AVLogError int = C.AV_LOG_ERROR

// AVLogWarning wraps AV_LOG_WARNING.
const AVLogWarning int = C.AV_LOG_WARNING

// AVLogInfo wraps AV_LOG_INFO.
const AVLogInfo int = C.AV_LOG_INFO

// AVLogVerbose wraps AV_LOG_VERBOSE.
const AVLogVerbose int = C.AV_LOG_VERBOSE

// AVLogDebug wraps AV_LOG_DEBUG.
const AVLogDebug int = C.AV_LOG_DEBUG

// AVLogTrace wraps AV_LOG_TRACE.
const AVLogTrace int = C.AV_LOG_TRACE

// AVLogMaxOffset wraps AV_LOG_MAX_OFFSET.
const AVLogMaxOffset int = C.AV_LOG_MAX_OFFSET

// AVLogSkipRepeated wraps AV_LOG_SKIP_REPEATED.
const AVLogSkipRepeated int = C.AV_LOG_SKIP_REPEATED

// AVLogPrintLevel wraps AV_LOG_PRINT_LEVEL.
const AVLogPrintLevel int = C.AV_LOG_PRINT_LEVEL

// AVOptFlagEncodingParam wraps AV_OPT_FLAG_ENCODING_PARAM.
const AVOptFlagEncodingParam int = C.AV_OPT_FLAG_ENCODING_PARAM

// AVOptFlagDecodingParam wraps AV_OPT_FLAG_DECODING_PARAM.
const AVOptFlagDecodingParam int = C.AV_OPT_FLAG_DECODING_PARAM

// AVOptFlagAudioParam wraps AV_OPT_FLAG_AUDIO_PARAM.
const AVOptFlagAudioParam int = C.AV_OPT_FLAG_AUDIO_PARAM

// AVOptFlagVideoParam wraps AV_OPT_FLAG_VIDEO_PARAM.
const AVOptFlagVideoParam int = C.AV_OPT_FLAG_VIDEO_PARAM

// AVOptFlagSubtitleParam wraps AV_OPT_FLAG_SUBTITLE_PARAM.
const AVOptFlagSubtitleParam int = C.AV_OPT_FLAG_SUBTITLE_PARAM

// AVOptFlagExport wraps AV_OPT_FLAG_EXPORT.
const AVOptFlagExport int = C.AV_OPT_FLAG_EXPORT

// AVOptFlagReadonly wraps AV_OPT_FLAG_READONLY.
const AVOptFlagReadonly int = C.AV_OPT_FLAG_READONLY

// AVOptFlagBsfParam wraps AV_OPT_FLAG_BSF_PARAM.
const AVOptFlagBsfParam int = C.AV_OPT_FLAG_BSF_PARAM

// AVOptFlagRuntimeParam wraps AV_OPT_FLAG_RUNTIME_PARAM.
const AVOptFlagRuntimeParam int = C.AV_OPT_FLAG_RUNTIME_PARAM

// AVOptFlagFilteringParam wraps AV_OPT_FLAG_FILTERING_PARAM.
const AVOptFlagFilteringParam int = C.AV_OPT_FLAG_FILTERING_PARAM

// AVOptFlagDeprecated wraps AV_OPT_FLAG_DEPRECATED.
const AVOptFlagDeprecated int = C.AV_OPT_FLAG_DEPRECATED

// AVOptFlagChildConsts wraps AV_OPT_FLAG_CHILD_CONSTS.
const AVOptFlagChildConsts int = C.AV_OPT_FLAG_CHILD_CONSTS

// AVOptSearchChildren wraps AV_OPT_SEARCH_CHILDREN.
const AVOptSearchChildren int = C.AV_OPT_SEARCH_CHILDREN

// AVOptSearchFakeObj wraps AV_OPT_SEARCH_FAKE_OBJ.
const AVOptSearchFakeObj int = C.AV_OPT_SEARCH_FAKE_OBJ

// AVOptAllowNull wraps AV_OPT_ALLOW_NULL.
const AVOptAllowNull int = C.AV_OPT_ALLOW_NULL

// AVOptMultiComponentRange wraps AV_OPT_MULTI_COMPONENT_RANGE.
const AVOptMultiComponentRange int = C.AV_OPT_MULTI_COMPONENT_RANGE

// AVOptSerializeSkipDefaults wraps AV_OPT_SERIALIZE_SKIP_DEFAULTS.
const AVOptSerializeSkipDefaults int = C.AV_OPT_SERIALIZE_SKIP_DEFAULTS

// AVOptSerializeOptFlagsExact wraps AV_OPT_SERIALIZE_OPT_FLAGS_EXACT.
const AVOptSerializeOptFlagsExact int = C.AV_OPT_SERIALIZE_OPT_FLAGS_EXACT

// AVPaletteSize wraps AVPALETTE_SIZE.
const AVPaletteSize int = C.AVPALETTE_SIZE

// AVPaletteCount wraps AVPALETTE_COUNT.
const AVPaletteCount int = C.AVPALETTE_COUNT

// AVPixFmtRgb32 wraps AV_PIX_FMT_RGB32.
const AVPixFmtRgb32 int = C.AV_PIX_FMT_RGB32

// AVPixFmtRgb321 wraps AV_PIX_FMT_RGB32_1.
const AVPixFmtRgb321 int = C.AV_PIX_FMT_RGB32_1

// AVPixFmtBgr32 wraps AV_PIX_FMT_BGR32.
const AVPixFmtBgr32 int = C.AV_PIX_FMT_BGR32

// AVPixFmtBgr321 wraps AV_PIX_FMT_BGR32_1.
const AVPixFmtBgr321 int = C.AV_PIX_FMT_BGR32_1

// AVPixFmt0Rgb32 wraps AV_PIX_FMT_0RGB32.
const AVPixFmt0Rgb32 int = C.AV_PIX_FMT_0RGB32

// AVPixFmt0Bgr32 wraps AV_PIX_FMT_0BGR32.
const AVPixFmt0Bgr32 int = C.AV_PIX_FMT_0BGR32

// AVPixFmtGray9 wraps AV_PIX_FMT_GRAY9.
const AVPixFmtGray9 int = C.AV_PIX_FMT_GRAY9

// AVPixFmtGray10 wraps AV_PIX_FMT_GRAY10.
const AVPixFmtGray10 int = C.AV_PIX_FMT_GRAY10

// AVPixFmtGray12 wraps AV_PIX_FMT_GRAY12.
const AVPixFmtGray12 int = C.AV_PIX_FMT_GRAY12

// AVPixFmtGray14 wraps AV_PIX_FMT_GRAY14.
const AVPixFmtGray14 int = C.AV_PIX_FMT_GRAY14

// AVPixFmtGray16 wraps AV_PIX_FMT_GRAY16.
const AVPixFmtGray16 int = C.AV_PIX_FMT_GRAY16

// AVPixFmtYa16 wraps AV_PIX_FMT_YA16.
const AVPixFmtYa16 int = C.AV_PIX_FMT_YA16

// AVPixFmtRgb48 wraps AV_PIX_FMT_RGB48.
const AVPixFmtRgb48 int = C.AV_PIX_FMT_RGB48

// AVPixFmtRgb565 wraps AV_PIX_FMT_RGB565.
const AVPixFmtRgb565 int = C.AV_PIX_FMT_RGB565

// AVPixFmtRgb555 wraps AV_PIX_FMT_RGB555.
const AVPixFmtRgb555 int = C.AV_PIX_FMT_RGB555

// AVPixFmtRgb444 wraps AV_PIX_FMT_RGB444.
const AVPixFmtRgb444 int = C.AV_PIX_FMT_RGB444

// AVPixFmtRgba64 wraps AV_PIX_FMT_RGBA64.
const AVPixFmtRgba64 int = C.AV_PIX_FMT_RGBA64

// AVPixFmtBgr48 wraps AV_PIX_FMT_BGR48.
const AVPixFmtBgr48 int = C.AV_PIX_FMT_BGR48

// AVPixFmtBgr565 wraps AV_PIX_FMT_BGR565.
const AVPixFmtBgr565 int = C.AV_PIX_FMT_BGR565

// AVPixFmtBgr555 wraps AV_PIX_FMT_BGR555.
const AVPixFmtBgr555 int = C.AV_PIX_FMT_BGR555

// AVPixFmtBgr444 wraps AV_PIX_FMT_BGR444.
const AVPixFmtBgr444 int = C.AV_PIX_FMT_BGR444

// AVPixFmtBgra64 wraps AV_PIX_FMT_BGRA64.
const AVPixFmtBgra64 int = C.AV_PIX_FMT_BGRA64

// AVPixFmtYuv420P9 wraps AV_PIX_FMT_YUV420P9.
const AVPixFmtYuv420P9 int = C.AV_PIX_FMT_YUV420P9

// AVPixFmtYuv422P9 wraps AV_PIX_FMT_YUV422P9.
const AVPixFmtYuv422P9 int = C.AV_PIX_FMT_YUV422P9

// AVPixFmtYuv444P9 wraps AV_PIX_FMT_YUV444P9.
const AVPixFmtYuv444P9 int = C.AV_PIX_FMT_YUV444P9

// AVPixFmtYuv420P10 wraps AV_PIX_FMT_YUV420P10.
const AVPixFmtYuv420P10 int = C.AV_PIX_FMT_YUV420P10

// AVPixFmtYuv422P10 wraps AV_PIX_FMT_YUV422P10.
const AVPixFmtYuv422P10 int = C.AV_PIX_FMT_YUV422P10

// AVPixFmtYuv440P10 wraps AV_PIX_FMT_YUV440P10.
const AVPixFmtYuv440P10 int = C.AV_PIX_FMT_YUV440P10

// AVPixFmtYuv444P10 wraps AV_PIX_FMT_YUV444P10.
const AVPixFmtYuv444P10 int = C.AV_PIX_FMT_YUV444P10

// AVPixFmtYuv420P12 wraps AV_PIX_FMT_YUV420P12.
const AVPixFmtYuv420P12 int = C.AV_PIX_FMT_YUV420P12

// AVPixFmtYuv422P12 wraps AV_PIX_FMT_YUV422P12.
const AVPixFmtYuv422P12 int = C.AV_PIX_FMT_YUV422P12

// AVPixFmtYuv440P12 wraps AV_PIX_FMT_YUV440P12.
const AVPixFmtYuv440P12 int = C.AV_PIX_FMT_YUV440P12

// AVPixFmtYuv444P12 wraps AV_PIX_FMT_YUV444P12.
const AVPixFmtYuv444P12 int = C.AV_PIX_FMT_YUV444P12

// AVPixFmtYuv420P14 wraps AV_PIX_FMT_YUV420P14.
const AVPixFmtYuv420P14 int = C.AV_PIX_FMT_YUV420P14

// AVPixFmtYuv422P14 wraps AV_PIX_FMT_YUV422P14.
const AVPixFmtYuv422P14 int = C.AV_PIX_FMT_YUV422P14

// AVPixFmtYuv444P14 wraps AV_PIX_FMT_YUV444P14.
const AVPixFmtYuv444P14 int = C.AV_PIX_FMT_YUV444P14

// AVPixFmtYuv420P16 wraps AV_PIX_FMT_YUV420P16.
const AVPixFmtYuv420P16 int = C.AV_PIX_FMT_YUV420P16

// AVPixFmtYuv422P16 wraps AV_PIX_FMT_YUV422P16.
const AVPixFmtYuv422P16 int = C.AV_PIX_FMT_YUV422P16

// AVPixFmtYuv444P16 wraps AV_PIX_FMT_YUV444P16.
const AVPixFmtYuv444P16 int = C.AV_PIX_FMT_YUV444P16

// AVPixFmtGbrp9 wraps AV_PIX_FMT_GBRP9.
const AVPixFmtGbrp9 int = C.AV_PIX_FMT_GBRP9

// AVPixFmtGbrp10 wraps AV_PIX_FMT_GBRP10.
const AVPixFmtGbrp10 int = C.AV_PIX_FMT_GBRP10

// AVPixFmtGbrp12 wraps AV_PIX_FMT_GBRP12.
const AVPixFmtGbrp12 int = C.AV_PIX_FMT_GBRP12

// AVPixFmtGbrp14 wraps AV_PIX_FMT_GBRP14.
const AVPixFmtGbrp14 int = C.AV_PIX_FMT_GBRP14

// AVPixFmtGbrp16 wraps AV_PIX_FMT_GBRP16.
const AVPixFmtGbrp16 int = C.AV_PIX_FMT_GBRP16

// AVPixFmtGbrap10 wraps AV_PIX_FMT_GBRAP10.
const AVPixFmtGbrap10 int = C.AV_PIX_FMT_GBRAP10

// AVPixFmtGbrap12 wraps AV_PIX_FMT_GBRAP12.
const AVPixFmtGbrap12 int = C.AV_PIX_FMT_GBRAP12

// AVPixFmtGbrap16 wraps AV_PIX_FMT_GBRAP16.
const AVPixFmtGbrap16 int = C.AV_PIX_FMT_GBRAP16

// AVPixFmtBayerBggr16 wraps AV_PIX_FMT_BAYER_BGGR16.
const AVPixFmtBayerBggr16 int = C.AV_PIX_FMT_BAYER_BGGR16

// AVPixFmtBayerRggb16 wraps AV_PIX_FMT_BAYER_RGGB16.
const AVPixFmtBayerRggb16 int = C.AV_PIX_FMT_BAYER_RGGB16

// AVPixFmtBayerGbrg16 wraps AV_PIX_FMT_BAYER_GBRG16.
const AVPixFmtBayerGbrg16 int = C.AV_PIX_FMT_BAYER_GBRG16

// AVPixFmtBayerGrbg16 wraps AV_PIX_FMT_BAYER_GRBG16.
const AVPixFmtBayerGrbg16 int = C.AV_PIX_FMT_BAYER_GRBG16

// AVPixFmtGbrpf32 wraps AV_PIX_FMT_GBRPF32.
const AVPixFmtGbrpf32 int = C.AV_PIX_FMT_GBRPF32

// AVPixFmtGbrapf32 wraps AV_PIX_FMT_GBRAPF32.
const AVPixFmtGbrapf32 int = C.AV_PIX_FMT_GBRAPF32

// AVPixFmtGrayf32 wraps AV_PIX_FMT_GRAYF32.
const AVPixFmtGrayf32 int = C.AV_PIX_FMT_GRAYF32

// AVPixFmtYuva420P9 wraps AV_PIX_FMT_YUVA420P9.
const AVPixFmtYuva420P9 int = C.AV_PIX_FMT_YUVA420P9

// AVPixFmtYuva422P9 wraps AV_PIX_FMT_YUVA422P9.
const AVPixFmtYuva422P9 int = C.AV_PIX_FMT_YUVA422P9

// AVPixFmtYuva444P9 wraps AV_PIX_FMT_YUVA444P9.
const AVPixFmtYuva444P9 int = C.AV_PIX_FMT_YUVA444P9

// AVPixFmtYuva420P10 wraps AV_PIX_FMT_YUVA420P10.
const AVPixFmtYuva420P10 int = C.AV_PIX_FMT_YUVA420P10

// AVPixFmtYuva422P10 wraps AV_PIX_FMT_YUVA422P10.
const AVPixFmtYuva422P10 int = C.AV_PIX_FMT_YUVA422P10

// AVPixFmtYuva444P10 wraps AV_PIX_FMT_YUVA444P10.
const AVPixFmtYuva444P10 int = C.AV_PIX_FMT_YUVA444P10

// AVPixFmtYuva422P12 wraps AV_PIX_FMT_YUVA422P12.
const AVPixFmtYuva422P12 int = C.AV_PIX_FMT_YUVA422P12

// AVPixFmtYuva444P12 wraps AV_PIX_FMT_YUVA444P12.
const AVPixFmtYuva444P12 int = C.AV_PIX_FMT_YUVA444P12

// AVPixFmtYuva420P16 wraps AV_PIX_FMT_YUVA420P16.
const AVPixFmtYuva420P16 int = C.AV_PIX_FMT_YUVA420P16

// AVPixFmtYuva422P16 wraps AV_PIX_FMT_YUVA422P16.
const AVPixFmtYuva422P16 int = C.AV_PIX_FMT_YUVA422P16

// AVPixFmtYuva444P16 wraps AV_PIX_FMT_YUVA444P16.
const AVPixFmtYuva444P16 int = C.AV_PIX_FMT_YUVA444P16

// AVPixFmtXyz12 wraps AV_PIX_FMT_XYZ12.
const AVPixFmtXyz12 int = C.AV_PIX_FMT_XYZ12

// AVPixFmtNv20 wraps AV_PIX_FMT_NV20.
const AVPixFmtNv20 int = C.AV_PIX_FMT_NV20

// AVPixFmtAyuv64 wraps AV_PIX_FMT_AYUV64.
const AVPixFmtAyuv64 int = C.AV_PIX_FMT_AYUV64

// AVPixFmtP010 wraps AV_PIX_FMT_P010.
const AVPixFmtP010 int = C.AV_PIX_FMT_P010

// AVPixFmtP012 wraps AV_PIX_FMT_P012.
const AVPixFmtP012 int = C.AV_PIX_FMT_P012

// AVPixFmtP016 wraps AV_PIX_FMT_P016.
const AVPixFmtP016 int = C.AV_PIX_FMT_P016

// AVPixFmtY210 wraps AV_PIX_FMT_Y210.
const AVPixFmtY210 int = C.AV_PIX_FMT_Y210

// AVPixFmtY212 wraps AV_PIX_FMT_Y212.
const AVPixFmtY212 int = C.AV_PIX_FMT_Y212

// AVPixFmtXv30 wraps AV_PIX_FMT_XV30.
const AVPixFmtXv30 int = C.AV_PIX_FMT_XV30

// AVPixFmtXv36 wraps AV_PIX_FMT_XV36.
const AVPixFmtXv36 int = C.AV_PIX_FMT_XV36

// AVPixFmtX2Rgb10 wraps AV_PIX_FMT_X2RGB10.
const AVPixFmtX2Rgb10 int = C.AV_PIX_FMT_X2RGB10

// AVPixFmtX2Bgr10 wraps AV_PIX_FMT_X2BGR10.
const AVPixFmtX2Bgr10 int = C.AV_PIX_FMT_X2BGR10

// AVPixFmtP210 wraps AV_PIX_FMT_P210.
const AVPixFmtP210 int = C.AV_PIX_FMT_P210

// AVPixFmtP410 wraps AV_PIX_FMT_P410.
const AVPixFmtP410 int = C.AV_PIX_FMT_P410

// AVPixFmtP216 wraps AV_PIX_FMT_P216.
const AVPixFmtP216 int = C.AV_PIX_FMT_P216

// AVPixFmtP416 wraps AV_PIX_FMT_P416.
const AVPixFmtP416 int = C.AV_PIX_FMT_P416

// AVPixFmtRgbaf16 wraps AV_PIX_FMT_RGBAF16.
const AVPixFmtRgbaf16 int = C.AV_PIX_FMT_RGBAF16

// AVPixFmtRgbf32 wraps AV_PIX_FMT_RGBF32.
const AVPixFmtRgbf32 int = C.AV_PIX_FMT_RGBF32

// AVPixFmtRgbaf32 wraps AV_PIX_FMT_RGBAF32.
const AVPixFmtRgbaf32 int = C.AV_PIX_FMT_RGBAF32
