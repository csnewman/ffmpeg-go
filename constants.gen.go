package ffmpeg

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

// AVInputBufferMinSize wraps AV_INPUT_BUFFER_MIN_SIZE.
const AVInputBufferMinSize = C.AV_INPUT_BUFFER_MIN_SIZE

// AVCodecFlagUnaligned wraps AV_CODEC_FLAG_UNALIGNED.
const AVCodecFlagUnaligned = C.AV_CODEC_FLAG_UNALIGNED

// AVCodecFlagQscale wraps AV_CODEC_FLAG_QSCALE.
const AVCodecFlagQscale = C.AV_CODEC_FLAG_QSCALE

// AVCodecFlag4Mv wraps AV_CODEC_FLAG_4MV.
const AVCodecFlag4Mv = C.AV_CODEC_FLAG_4MV

// AVCodecFlagOutputCorrupt wraps AV_CODEC_FLAG_OUTPUT_CORRUPT.
const AVCodecFlagOutputCorrupt = C.AV_CODEC_FLAG_OUTPUT_CORRUPT

// AVCodecFlagQpel wraps AV_CODEC_FLAG_QPEL.
const AVCodecFlagQpel = C.AV_CODEC_FLAG_QPEL

// AVCodecFlagDropchanged wraps AV_CODEC_FLAG_DROPCHANGED.
const AVCodecFlagDropchanged = C.AV_CODEC_FLAG_DROPCHANGED

// AVCodecFlagReconFrame wraps AV_CODEC_FLAG_RECON_FRAME.
const AVCodecFlagReconFrame = C.AV_CODEC_FLAG_RECON_FRAME

// AVCodecFlagCopyOpaque wraps AV_CODEC_FLAG_COPY_OPAQUE.
const AVCodecFlagCopyOpaque = C.AV_CODEC_FLAG_COPY_OPAQUE

// AVCodecFlagFrameDuration wraps AV_CODEC_FLAG_FRAME_DURATION.
const AVCodecFlagFrameDuration = C.AV_CODEC_FLAG_FRAME_DURATION

// AVCodecFlagPass1 wraps AV_CODEC_FLAG_PASS1.
const AVCodecFlagPass1 = C.AV_CODEC_FLAG_PASS1

// AVCodecFlagPass2 wraps AV_CODEC_FLAG_PASS2.
const AVCodecFlagPass2 = C.AV_CODEC_FLAG_PASS2

// AVCodecFlagLoopFilter wraps AV_CODEC_FLAG_LOOP_FILTER.
const AVCodecFlagLoopFilter = C.AV_CODEC_FLAG_LOOP_FILTER

// AVCodecFlagGray wraps AV_CODEC_FLAG_GRAY.
const AVCodecFlagGray = C.AV_CODEC_FLAG_GRAY

// AVCodecFlagPsnr wraps AV_CODEC_FLAG_PSNR.
const AVCodecFlagPsnr = C.AV_CODEC_FLAG_PSNR

// AVCodecFlagInterlacedDct wraps AV_CODEC_FLAG_INTERLACED_DCT.
const AVCodecFlagInterlacedDct = C.AV_CODEC_FLAG_INTERLACED_DCT

// AVCodecFlagLowDelay wraps AV_CODEC_FLAG_LOW_DELAY.
const AVCodecFlagLowDelay = C.AV_CODEC_FLAG_LOW_DELAY

// AVCodecFlagGlobalHeader wraps AV_CODEC_FLAG_GLOBAL_HEADER.
const AVCodecFlagGlobalHeader = C.AV_CODEC_FLAG_GLOBAL_HEADER

// AVCodecFlagBitexact wraps AV_CODEC_FLAG_BITEXACT.
const AVCodecFlagBitexact = C.AV_CODEC_FLAG_BITEXACT

// AVCodecFlagAcPred wraps AV_CODEC_FLAG_AC_PRED.
const AVCodecFlagAcPred = C.AV_CODEC_FLAG_AC_PRED

// AVCodecFlagInterlacedMe wraps AV_CODEC_FLAG_INTERLACED_ME.
const AVCodecFlagInterlacedMe = C.AV_CODEC_FLAG_INTERLACED_ME

// AVCodecFlagClosedGop wraps AV_CODEC_FLAG_CLOSED_GOP.
const AVCodecFlagClosedGop = C.AV_CODEC_FLAG_CLOSED_GOP

// AVCodecFlag2Fast wraps AV_CODEC_FLAG2_FAST.
const AVCodecFlag2Fast = C.AV_CODEC_FLAG2_FAST

// AVCodecFlag2NoOutput wraps AV_CODEC_FLAG2_NO_OUTPUT.
const AVCodecFlag2NoOutput = C.AV_CODEC_FLAG2_NO_OUTPUT

// AVCodecFlag2LocalHeader wraps AV_CODEC_FLAG2_LOCAL_HEADER.
const AVCodecFlag2LocalHeader = C.AV_CODEC_FLAG2_LOCAL_HEADER

// AVCodecFlag2Chunks wraps AV_CODEC_FLAG2_CHUNKS.
const AVCodecFlag2Chunks = C.AV_CODEC_FLAG2_CHUNKS

// AVCodecFlag2IgnoreCrop wraps AV_CODEC_FLAG2_IGNORE_CROP.
const AVCodecFlag2IgnoreCrop = C.AV_CODEC_FLAG2_IGNORE_CROP

// AVCodecFlag2ShowAll wraps AV_CODEC_FLAG2_SHOW_ALL.
const AVCodecFlag2ShowAll = C.AV_CODEC_FLAG2_SHOW_ALL

// AVCodecFlag2ExportMvs wraps AV_CODEC_FLAG2_EXPORT_MVS.
const AVCodecFlag2ExportMvs = C.AV_CODEC_FLAG2_EXPORT_MVS

// AVCodecFlag2SkipManual wraps AV_CODEC_FLAG2_SKIP_MANUAL.
const AVCodecFlag2SkipManual = C.AV_CODEC_FLAG2_SKIP_MANUAL

// AVCodecFlag2RoFlushNoop wraps AV_CODEC_FLAG2_RO_FLUSH_NOOP.
const AVCodecFlag2RoFlushNoop = C.AV_CODEC_FLAG2_RO_FLUSH_NOOP

// AVCodecFlag2IccProfiles wraps AV_CODEC_FLAG2_ICC_PROFILES.
const AVCodecFlag2IccProfiles = C.AV_CODEC_FLAG2_ICC_PROFILES

// AVCodecExportDataMvs wraps AV_CODEC_EXPORT_DATA_MVS.
const AVCodecExportDataMvs = C.AV_CODEC_EXPORT_DATA_MVS

// AVCodecExportDataPrft wraps AV_CODEC_EXPORT_DATA_PRFT.
const AVCodecExportDataPrft = C.AV_CODEC_EXPORT_DATA_PRFT

// AVCodecExportDataVideoEncParams wraps AV_CODEC_EXPORT_DATA_VIDEO_ENC_PARAMS.
const AVCodecExportDataVideoEncParams = C.AV_CODEC_EXPORT_DATA_VIDEO_ENC_PARAMS

// AVCodecExportDataFilmGrain wraps AV_CODEC_EXPORT_DATA_FILM_GRAIN.
const AVCodecExportDataFilmGrain = C.AV_CODEC_EXPORT_DATA_FILM_GRAIN

// AVGetBufferFlagRef wraps AV_GET_BUFFER_FLAG_REF.
const AVGetBufferFlagRef = C.AV_GET_BUFFER_FLAG_REF

// AVGetEncodeBufferFlagRef wraps AV_GET_ENCODE_BUFFER_FLAG_REF.
const AVGetEncodeBufferFlagRef = C.AV_GET_ENCODE_BUFFER_FLAG_REF

// FFCompressionDefault wraps FF_COMPRESSION_DEFAULT.
const FFCompressionDefault = C.FF_COMPRESSION_DEFAULT

// FFCmpSad wraps FF_CMP_SAD.
const FFCmpSad = C.FF_CMP_SAD

// FFCmpSse wraps FF_CMP_SSE.
const FFCmpSse = C.FF_CMP_SSE

// FFCmpSatd wraps FF_CMP_SATD.
const FFCmpSatd = C.FF_CMP_SATD

// FFCmpDct wraps FF_CMP_DCT.
const FFCmpDct = C.FF_CMP_DCT

// FFCmpPsnr wraps FF_CMP_PSNR.
const FFCmpPsnr = C.FF_CMP_PSNR

// FFCmpBit wraps FF_CMP_BIT.
const FFCmpBit = C.FF_CMP_BIT

// FFCmpRd wraps FF_CMP_RD.
const FFCmpRd = C.FF_CMP_RD

// FFCmpZero wraps FF_CMP_ZERO.
const FFCmpZero = C.FF_CMP_ZERO

// FFCmpVsad wraps FF_CMP_VSAD.
const FFCmpVsad = C.FF_CMP_VSAD

// FFCmpVsse wraps FF_CMP_VSSE.
const FFCmpVsse = C.FF_CMP_VSSE

// FFCmpNsse wraps FF_CMP_NSSE.
const FFCmpNsse = C.FF_CMP_NSSE

// FFCmpW53 wraps FF_CMP_W53.
const FFCmpW53 = C.FF_CMP_W53

// FFCmpW97 wraps FF_CMP_W97.
const FFCmpW97 = C.FF_CMP_W97

// FFCmpDctmax wraps FF_CMP_DCTMAX.
const FFCmpDctmax = C.FF_CMP_DCTMAX

// FFCmpDct264 wraps FF_CMP_DCT264.
const FFCmpDct264 = C.FF_CMP_DCT264

// FFCmpMedianSad wraps FF_CMP_MEDIAN_SAD.
const FFCmpMedianSad = C.FF_CMP_MEDIAN_SAD

// FFCmpChroma wraps FF_CMP_CHROMA.
const FFCmpChroma = C.FF_CMP_CHROMA

// SliceFlagCodedOrder wraps SLICE_FLAG_CODED_ORDER.
const SliceFlagCodedOrder = C.SLICE_FLAG_CODED_ORDER

// SliceFlagAllowField wraps SLICE_FLAG_ALLOW_FIELD.
const SliceFlagAllowField = C.SLICE_FLAG_ALLOW_FIELD

// SliceFlagAllowPlane wraps SLICE_FLAG_ALLOW_PLANE.
const SliceFlagAllowPlane = C.SLICE_FLAG_ALLOW_PLANE

// FFMbDecisionSimple wraps FF_MB_DECISION_SIMPLE.
const FFMbDecisionSimple = C.FF_MB_DECISION_SIMPLE

// FFMbDecisionBits wraps FF_MB_DECISION_BITS.
const FFMbDecisionBits = C.FF_MB_DECISION_BITS

// FFMbDecisionRd wraps FF_MB_DECISION_RD.
const FFMbDecisionRd = C.FF_MB_DECISION_RD

// FFBugAutodetect wraps FF_BUG_AUTODETECT.
const FFBugAutodetect = C.FF_BUG_AUTODETECT

// FFBugXvidIlace wraps FF_BUG_XVID_ILACE.
const FFBugXvidIlace = C.FF_BUG_XVID_ILACE

// FFBugUmp4 wraps FF_BUG_UMP4.
const FFBugUmp4 = C.FF_BUG_UMP4

// FFBugNoPadding wraps FF_BUG_NO_PADDING.
const FFBugNoPadding = C.FF_BUG_NO_PADDING

// FFBugAmv wraps FF_BUG_AMV.
const FFBugAmv = C.FF_BUG_AMV

// FFBugQpelChroma wraps FF_BUG_QPEL_CHROMA.
const FFBugQpelChroma = C.FF_BUG_QPEL_CHROMA

// FFBugStdQpel wraps FF_BUG_STD_QPEL.
const FFBugStdQpel = C.FF_BUG_STD_QPEL

// FFBugQpelChroma2 wraps FF_BUG_QPEL_CHROMA2.
const FFBugQpelChroma2 = C.FF_BUG_QPEL_CHROMA2

// FFBugDirectBlocksize wraps FF_BUG_DIRECT_BLOCKSIZE.
const FFBugDirectBlocksize = C.FF_BUG_DIRECT_BLOCKSIZE

// FFBugEdge wraps FF_BUG_EDGE.
const FFBugEdge = C.FF_BUG_EDGE

// FFBugHpelChroma wraps FF_BUG_HPEL_CHROMA.
const FFBugHpelChroma = C.FF_BUG_HPEL_CHROMA

// FFBugDcClip wraps FF_BUG_DC_CLIP.
const FFBugDcClip = C.FF_BUG_DC_CLIP

// FFBugMs wraps FF_BUG_MS.
const FFBugMs = C.FF_BUG_MS

// FFBugTruncated wraps FF_BUG_TRUNCATED.
const FFBugTruncated = C.FF_BUG_TRUNCATED

// FFBugIedge wraps FF_BUG_IEDGE.
const FFBugIedge = C.FF_BUG_IEDGE

// FFEcGuessMvs wraps FF_EC_GUESS_MVS.
const FFEcGuessMvs = C.FF_EC_GUESS_MVS

// FFEcDeblock wraps FF_EC_DEBLOCK.
const FFEcDeblock = C.FF_EC_DEBLOCK

// FFEcFavorInter wraps FF_EC_FAVOR_INTER.
const FFEcFavorInter = C.FF_EC_FAVOR_INTER

// FFDebugPictInfo wraps FF_DEBUG_PICT_INFO.
const FFDebugPictInfo = C.FF_DEBUG_PICT_INFO

// FFDebugRc wraps FF_DEBUG_RC.
const FFDebugRc = C.FF_DEBUG_RC

// FFDebugBitstream wraps FF_DEBUG_BITSTREAM.
const FFDebugBitstream = C.FF_DEBUG_BITSTREAM

// FFDebugMbType wraps FF_DEBUG_MB_TYPE.
const FFDebugMbType = C.FF_DEBUG_MB_TYPE

// FFDebugQp wraps FF_DEBUG_QP.
const FFDebugQp = C.FF_DEBUG_QP

// FFDebugDctCoeff wraps FF_DEBUG_DCT_COEFF.
const FFDebugDctCoeff = C.FF_DEBUG_DCT_COEFF

// FFDebugSkip wraps FF_DEBUG_SKIP.
const FFDebugSkip = C.FF_DEBUG_SKIP

// FFDebugStartcode wraps FF_DEBUG_STARTCODE.
const FFDebugStartcode = C.FF_DEBUG_STARTCODE

// FFDebugEr wraps FF_DEBUG_ER.
const FFDebugEr = C.FF_DEBUG_ER

// FFDebugMmco wraps FF_DEBUG_MMCO.
const FFDebugMmco = C.FF_DEBUG_MMCO

// FFDebugBugs wraps FF_DEBUG_BUGS.
const FFDebugBugs = C.FF_DEBUG_BUGS

// FFDebugBuffers wraps FF_DEBUG_BUFFERS.
const FFDebugBuffers = C.FF_DEBUG_BUFFERS

// FFDebugThreads wraps FF_DEBUG_THREADS.
const FFDebugThreads = C.FF_DEBUG_THREADS

// FFDebugGreenMd wraps FF_DEBUG_GREEN_MD.
const FFDebugGreenMd = C.FF_DEBUG_GREEN_MD

// FFDebugNomc wraps FF_DEBUG_NOMC.
const FFDebugNomc = C.FF_DEBUG_NOMC

// FFDctAuto wraps FF_DCT_AUTO.
const FFDctAuto = C.FF_DCT_AUTO

// FFDctFastint wraps FF_DCT_FASTINT.
const FFDctFastint = C.FF_DCT_FASTINT

// FFDctInt wraps FF_DCT_INT.
const FFDctInt = C.FF_DCT_INT

// FFDctMmx wraps FF_DCT_MMX.
const FFDctMmx = C.FF_DCT_MMX

// FFDctAltivec wraps FF_DCT_ALTIVEC.
const FFDctAltivec = C.FF_DCT_ALTIVEC

// FFDctFaan wraps FF_DCT_FAAN.
const FFDctFaan = C.FF_DCT_FAAN

// FFIdctAuto wraps FF_IDCT_AUTO.
const FFIdctAuto = C.FF_IDCT_AUTO

// FFIdctInt wraps FF_IDCT_INT.
const FFIdctInt = C.FF_IDCT_INT

// FFIdctSimple wraps FF_IDCT_SIMPLE.
const FFIdctSimple = C.FF_IDCT_SIMPLE

// FFIdctSimplemmx wraps FF_IDCT_SIMPLEMMX.
const FFIdctSimplemmx = C.FF_IDCT_SIMPLEMMX

// FFIdctArm wraps FF_IDCT_ARM.
const FFIdctArm = C.FF_IDCT_ARM

// FFIdctAltivec wraps FF_IDCT_ALTIVEC.
const FFIdctAltivec = C.FF_IDCT_ALTIVEC

// FFIdctSimplearm wraps FF_IDCT_SIMPLEARM.
const FFIdctSimplearm = C.FF_IDCT_SIMPLEARM

// FFIdctXvid wraps FF_IDCT_XVID.
const FFIdctXvid = C.FF_IDCT_XVID

// FFIdctSimplearmv5Te wraps FF_IDCT_SIMPLEARMV5TE.
const FFIdctSimplearmv5Te = C.FF_IDCT_SIMPLEARMV5TE

// FFIdctSimplearmv6 wraps FF_IDCT_SIMPLEARMV6.
const FFIdctSimplearmv6 = C.FF_IDCT_SIMPLEARMV6

// FFIdctFaan wraps FF_IDCT_FAAN.
const FFIdctFaan = C.FF_IDCT_FAAN

// FFIdctSimpleneon wraps FF_IDCT_SIMPLENEON.
const FFIdctSimpleneon = C.FF_IDCT_SIMPLENEON

// FFIdctNone wraps FF_IDCT_NONE.
const FFIdctNone = C.FF_IDCT_NONE

// FFIdctSimpleauto wraps FF_IDCT_SIMPLEAUTO.
const FFIdctSimpleauto = C.FF_IDCT_SIMPLEAUTO

// FFThreadFrame wraps FF_THREAD_FRAME.
const FFThreadFrame = C.FF_THREAD_FRAME

// FFThreadSlice wraps FF_THREAD_SLICE.
const FFThreadSlice = C.FF_THREAD_SLICE

// FFProfileUnknown wraps FF_PROFILE_UNKNOWN.
const FFProfileUnknown = C.FF_PROFILE_UNKNOWN

// FFProfileReserved wraps FF_PROFILE_RESERVED.
const FFProfileReserved = C.FF_PROFILE_RESERVED

// FFProfileAacMain wraps FF_PROFILE_AAC_MAIN.
const FFProfileAacMain = C.FF_PROFILE_AAC_MAIN

// FFProfileAacLow wraps FF_PROFILE_AAC_LOW.
const FFProfileAacLow = C.FF_PROFILE_AAC_LOW

// FFProfileAacSsr wraps FF_PROFILE_AAC_SSR.
const FFProfileAacSsr = C.FF_PROFILE_AAC_SSR

// FFProfileAacLtp wraps FF_PROFILE_AAC_LTP.
const FFProfileAacLtp = C.FF_PROFILE_AAC_LTP

// FFProfileAacHe wraps FF_PROFILE_AAC_HE.
const FFProfileAacHe = C.FF_PROFILE_AAC_HE

// FFProfileAacHeV2 wraps FF_PROFILE_AAC_HE_V2.
const FFProfileAacHeV2 = C.FF_PROFILE_AAC_HE_V2

// FFProfileAacLd wraps FF_PROFILE_AAC_LD.
const FFProfileAacLd = C.FF_PROFILE_AAC_LD

// FFProfileAacEld wraps FF_PROFILE_AAC_ELD.
const FFProfileAacEld = C.FF_PROFILE_AAC_ELD

// FFProfileMpeg2AacLow wraps FF_PROFILE_MPEG2_AAC_LOW.
const FFProfileMpeg2AacLow = C.FF_PROFILE_MPEG2_AAC_LOW

// FFProfileMpeg2AacHe wraps FF_PROFILE_MPEG2_AAC_HE.
const FFProfileMpeg2AacHe = C.FF_PROFILE_MPEG2_AAC_HE

// FFProfileDnxhd wraps FF_PROFILE_DNXHD.
const FFProfileDnxhd = C.FF_PROFILE_DNXHD

// FFProfileDnxhrLb wraps FF_PROFILE_DNXHR_LB.
const FFProfileDnxhrLb = C.FF_PROFILE_DNXHR_LB

// FFProfileDnxhrSq wraps FF_PROFILE_DNXHR_SQ.
const FFProfileDnxhrSq = C.FF_PROFILE_DNXHR_SQ

// FFProfileDnxhrHq wraps FF_PROFILE_DNXHR_HQ.
const FFProfileDnxhrHq = C.FF_PROFILE_DNXHR_HQ

// FFProfileDnxhrHqx wraps FF_PROFILE_DNXHR_HQX.
const FFProfileDnxhrHqx = C.FF_PROFILE_DNXHR_HQX

// FFProfileDnxhr444 wraps FF_PROFILE_DNXHR_444.
const FFProfileDnxhr444 = C.FF_PROFILE_DNXHR_444

// FFProfileDts wraps FF_PROFILE_DTS.
const FFProfileDts = C.FF_PROFILE_DTS

// FFProfileDtsEs wraps FF_PROFILE_DTS_ES.
const FFProfileDtsEs = C.FF_PROFILE_DTS_ES

// FFProfileDts9624 wraps FF_PROFILE_DTS_96_24.
const FFProfileDts9624 = C.FF_PROFILE_DTS_96_24

// FFProfileDtsHdHra wraps FF_PROFILE_DTS_HD_HRA.
const FFProfileDtsHdHra = C.FF_PROFILE_DTS_HD_HRA

// FFProfileDtsHdMa wraps FF_PROFILE_DTS_HD_MA.
const FFProfileDtsHdMa = C.FF_PROFILE_DTS_HD_MA

// FFProfileDtsExpress wraps FF_PROFILE_DTS_EXPRESS.
const FFProfileDtsExpress = C.FF_PROFILE_DTS_EXPRESS

// FFProfileDtsHdMaX wraps FF_PROFILE_DTS_HD_MA_X.
const FFProfileDtsHdMaX = C.FF_PROFILE_DTS_HD_MA_X

// FFProfileDtsHdMaXImax wraps FF_PROFILE_DTS_HD_MA_X_IMAX.
const FFProfileDtsHdMaXImax = C.FF_PROFILE_DTS_HD_MA_X_IMAX

// FFProfileEac3DdpAtmos wraps FF_PROFILE_EAC3_DDP_ATMOS.
const FFProfileEac3DdpAtmos = C.FF_PROFILE_EAC3_DDP_ATMOS

// FFProfileTruehdAtmos wraps FF_PROFILE_TRUEHD_ATMOS.
const FFProfileTruehdAtmos = C.FF_PROFILE_TRUEHD_ATMOS

// FFProfileMpeg2422 wraps FF_PROFILE_MPEG2_422.
const FFProfileMpeg2422 = C.FF_PROFILE_MPEG2_422

// FFProfileMpeg2High wraps FF_PROFILE_MPEG2_HIGH.
const FFProfileMpeg2High = C.FF_PROFILE_MPEG2_HIGH

// FFProfileMpeg2Ss wraps FF_PROFILE_MPEG2_SS.
const FFProfileMpeg2Ss = C.FF_PROFILE_MPEG2_SS

// FFProfileMpeg2SnrScalable wraps FF_PROFILE_MPEG2_SNR_SCALABLE.
const FFProfileMpeg2SnrScalable = C.FF_PROFILE_MPEG2_SNR_SCALABLE

// FFProfileMpeg2Main wraps FF_PROFILE_MPEG2_MAIN.
const FFProfileMpeg2Main = C.FF_PROFILE_MPEG2_MAIN

// FFProfileMpeg2Simple wraps FF_PROFILE_MPEG2_SIMPLE.
const FFProfileMpeg2Simple = C.FF_PROFILE_MPEG2_SIMPLE

// FFProfileH264Constrained wraps FF_PROFILE_H264_CONSTRAINED.
const FFProfileH264Constrained = C.FF_PROFILE_H264_CONSTRAINED

// FFProfileH264Intra wraps FF_PROFILE_H264_INTRA.
const FFProfileH264Intra = C.FF_PROFILE_H264_INTRA

// FFProfileH264Baseline wraps FF_PROFILE_H264_BASELINE.
const FFProfileH264Baseline = C.FF_PROFILE_H264_BASELINE

// FFProfileH264ConstrainedBaseline wraps FF_PROFILE_H264_CONSTRAINED_BASELINE.
const FFProfileH264ConstrainedBaseline = C.FF_PROFILE_H264_CONSTRAINED_BASELINE

// FFProfileH264Main wraps FF_PROFILE_H264_MAIN.
const FFProfileH264Main = C.FF_PROFILE_H264_MAIN

// FFProfileH264Extended wraps FF_PROFILE_H264_EXTENDED.
const FFProfileH264Extended = C.FF_PROFILE_H264_EXTENDED

// FFProfileH264High wraps FF_PROFILE_H264_HIGH.
const FFProfileH264High = C.FF_PROFILE_H264_HIGH

// FFProfileH264High10 wraps FF_PROFILE_H264_HIGH_10.
const FFProfileH264High10 = C.FF_PROFILE_H264_HIGH_10

// FFProfileH264High10Intra wraps FF_PROFILE_H264_HIGH_10_INTRA.
const FFProfileH264High10Intra = C.FF_PROFILE_H264_HIGH_10_INTRA

// FFProfileH264MultiviewHigh wraps FF_PROFILE_H264_MULTIVIEW_HIGH.
const FFProfileH264MultiviewHigh = C.FF_PROFILE_H264_MULTIVIEW_HIGH

// FFProfileH264High422 wraps FF_PROFILE_H264_HIGH_422.
const FFProfileH264High422 = C.FF_PROFILE_H264_HIGH_422

// FFProfileH264High422Intra wraps FF_PROFILE_H264_HIGH_422_INTRA.
const FFProfileH264High422Intra = C.FF_PROFILE_H264_HIGH_422_INTRA

// FFProfileH264StereoHigh wraps FF_PROFILE_H264_STEREO_HIGH.
const FFProfileH264StereoHigh = C.FF_PROFILE_H264_STEREO_HIGH

// FFProfileH264High444 wraps FF_PROFILE_H264_HIGH_444.
const FFProfileH264High444 = C.FF_PROFILE_H264_HIGH_444

// FFProfileH264High444Predictive wraps FF_PROFILE_H264_HIGH_444_PREDICTIVE.
const FFProfileH264High444Predictive = C.FF_PROFILE_H264_HIGH_444_PREDICTIVE

// FFProfileH264High444Intra wraps FF_PROFILE_H264_HIGH_444_INTRA.
const FFProfileH264High444Intra = C.FF_PROFILE_H264_HIGH_444_INTRA

// FFProfileH264Cavlc444 wraps FF_PROFILE_H264_CAVLC_444.
const FFProfileH264Cavlc444 = C.FF_PROFILE_H264_CAVLC_444

// FFProfileVc1Simple wraps FF_PROFILE_VC1_SIMPLE.
const FFProfileVc1Simple = C.FF_PROFILE_VC1_SIMPLE

// FFProfileVc1Main wraps FF_PROFILE_VC1_MAIN.
const FFProfileVc1Main = C.FF_PROFILE_VC1_MAIN

// FFProfileVc1Complex wraps FF_PROFILE_VC1_COMPLEX.
const FFProfileVc1Complex = C.FF_PROFILE_VC1_COMPLEX

// FFProfileVc1Advanced wraps FF_PROFILE_VC1_ADVANCED.
const FFProfileVc1Advanced = C.FF_PROFILE_VC1_ADVANCED

// FFProfileMpeg4Simple wraps FF_PROFILE_MPEG4_SIMPLE.
const FFProfileMpeg4Simple = C.FF_PROFILE_MPEG4_SIMPLE

// FFProfileMpeg4SimpleScalable wraps FF_PROFILE_MPEG4_SIMPLE_SCALABLE.
const FFProfileMpeg4SimpleScalable = C.FF_PROFILE_MPEG4_SIMPLE_SCALABLE

// FFProfileMpeg4Core wraps FF_PROFILE_MPEG4_CORE.
const FFProfileMpeg4Core = C.FF_PROFILE_MPEG4_CORE

// FFProfileMpeg4Main wraps FF_PROFILE_MPEG4_MAIN.
const FFProfileMpeg4Main = C.FF_PROFILE_MPEG4_MAIN

// FFProfileMpeg4NBit wraps FF_PROFILE_MPEG4_N_BIT.
const FFProfileMpeg4NBit = C.FF_PROFILE_MPEG4_N_BIT

// FFProfileMpeg4ScalableTexture wraps FF_PROFILE_MPEG4_SCALABLE_TEXTURE.
const FFProfileMpeg4ScalableTexture = C.FF_PROFILE_MPEG4_SCALABLE_TEXTURE

// FFProfileMpeg4SimpleFaceAnimation wraps FF_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION.
const FFProfileMpeg4SimpleFaceAnimation = C.FF_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION

// FFProfileMpeg4BasicAnimatedTexture wraps FF_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE.
const FFProfileMpeg4BasicAnimatedTexture = C.FF_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE

// FFProfileMpeg4Hybrid wraps FF_PROFILE_MPEG4_HYBRID.
const FFProfileMpeg4Hybrid = C.FF_PROFILE_MPEG4_HYBRID

// FFProfileMpeg4AdvancedRealTime wraps FF_PROFILE_MPEG4_ADVANCED_REAL_TIME.
const FFProfileMpeg4AdvancedRealTime = C.FF_PROFILE_MPEG4_ADVANCED_REAL_TIME

// FFProfileMpeg4CoreScalable wraps FF_PROFILE_MPEG4_CORE_SCALABLE.
const FFProfileMpeg4CoreScalable = C.FF_PROFILE_MPEG4_CORE_SCALABLE

// FFProfileMpeg4AdvancedCoding wraps FF_PROFILE_MPEG4_ADVANCED_CODING.
const FFProfileMpeg4AdvancedCoding = C.FF_PROFILE_MPEG4_ADVANCED_CODING

// FFProfileMpeg4AdvancedCore wraps FF_PROFILE_MPEG4_ADVANCED_CORE.
const FFProfileMpeg4AdvancedCore = C.FF_PROFILE_MPEG4_ADVANCED_CORE

// FFProfileMpeg4AdvancedScalableTexture wraps FF_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE.
const FFProfileMpeg4AdvancedScalableTexture = C.FF_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE

// FFProfileMpeg4SimpleStudio wraps FF_PROFILE_MPEG4_SIMPLE_STUDIO.
const FFProfileMpeg4SimpleStudio = C.FF_PROFILE_MPEG4_SIMPLE_STUDIO

// FFProfileMpeg4AdvancedSimple wraps FF_PROFILE_MPEG4_ADVANCED_SIMPLE.
const FFProfileMpeg4AdvancedSimple = C.FF_PROFILE_MPEG4_ADVANCED_SIMPLE

// FFProfileJpeg2000CstreamRestriction0 wraps FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0.
const FFProfileJpeg2000CstreamRestriction0 = C.FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0

// FFProfileJpeg2000CstreamRestriction1 wraps FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1.
const FFProfileJpeg2000CstreamRestriction1 = C.FF_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1

// FFProfileJpeg2000CstreamNoRestriction wraps FF_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION.
const FFProfileJpeg2000CstreamNoRestriction = C.FF_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION

// FFProfileJpeg2000Dcinema2K wraps FF_PROFILE_JPEG2000_DCINEMA_2K.
const FFProfileJpeg2000Dcinema2K = C.FF_PROFILE_JPEG2000_DCINEMA_2K

// FFProfileJpeg2000Dcinema4K wraps FF_PROFILE_JPEG2000_DCINEMA_4K.
const FFProfileJpeg2000Dcinema4K = C.FF_PROFILE_JPEG2000_DCINEMA_4K

// FFProfileVp90 wraps FF_PROFILE_VP9_0.
const FFProfileVp90 = C.FF_PROFILE_VP9_0

// FFProfileVp91 wraps FF_PROFILE_VP9_1.
const FFProfileVp91 = C.FF_PROFILE_VP9_1

// FFProfileVp92 wraps FF_PROFILE_VP9_2.
const FFProfileVp92 = C.FF_PROFILE_VP9_2

// FFProfileVp93 wraps FF_PROFILE_VP9_3.
const FFProfileVp93 = C.FF_PROFILE_VP9_3

// FFProfileHevcMain wraps FF_PROFILE_HEVC_MAIN.
const FFProfileHevcMain = C.FF_PROFILE_HEVC_MAIN

// FFProfileHevcMain10 wraps FF_PROFILE_HEVC_MAIN_10.
const FFProfileHevcMain10 = C.FF_PROFILE_HEVC_MAIN_10

// FFProfileHevcMainStillPicture wraps FF_PROFILE_HEVC_MAIN_STILL_PICTURE.
const FFProfileHevcMainStillPicture = C.FF_PROFILE_HEVC_MAIN_STILL_PICTURE

// FFProfileHevcRext wraps FF_PROFILE_HEVC_REXT.
const FFProfileHevcRext = C.FF_PROFILE_HEVC_REXT

// FFProfileHevcScc wraps FF_PROFILE_HEVC_SCC.
const FFProfileHevcScc = C.FF_PROFILE_HEVC_SCC

// FFProfileVvcMain10 wraps FF_PROFILE_VVC_MAIN_10.
const FFProfileVvcMain10 = C.FF_PROFILE_VVC_MAIN_10

// FFProfileVvcMain10444 wraps FF_PROFILE_VVC_MAIN_10_444.
const FFProfileVvcMain10444 = C.FF_PROFILE_VVC_MAIN_10_444

// FFProfileAV1Main wraps FF_PROFILE_AV1_MAIN.
const FFProfileAV1Main = C.FF_PROFILE_AV1_MAIN

// FFProfileAV1High wraps FF_PROFILE_AV1_HIGH.
const FFProfileAV1High = C.FF_PROFILE_AV1_HIGH

// FFProfileAV1Professional wraps FF_PROFILE_AV1_PROFESSIONAL.
const FFProfileAV1Professional = C.FF_PROFILE_AV1_PROFESSIONAL

// FFProfileMjpegHuffmanBaselineDct wraps FF_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT.
const FFProfileMjpegHuffmanBaselineDct = C.FF_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT

// FFProfileMjpegHuffmanExtendedSequentialDct wraps FF_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT.
const FFProfileMjpegHuffmanExtendedSequentialDct = C.FF_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT

// FFProfileMjpegHuffmanProgressiveDct wraps FF_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT.
const FFProfileMjpegHuffmanProgressiveDct = C.FF_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT

// FFProfileMjpegHuffmanLossless wraps FF_PROFILE_MJPEG_HUFFMAN_LOSSLESS.
const FFProfileMjpegHuffmanLossless = C.FF_PROFILE_MJPEG_HUFFMAN_LOSSLESS

// FFProfileMjpegJpegLs wraps FF_PROFILE_MJPEG_JPEG_LS.
const FFProfileMjpegJpegLs = C.FF_PROFILE_MJPEG_JPEG_LS

// FFProfileSbcMsbc wraps FF_PROFILE_SBC_MSBC.
const FFProfileSbcMsbc = C.FF_PROFILE_SBC_MSBC

// FFProfileProresProxy wraps FF_PROFILE_PRORES_PROXY.
const FFProfileProresProxy = C.FF_PROFILE_PRORES_PROXY

// FFProfileProresLt wraps FF_PROFILE_PRORES_LT.
const FFProfileProresLt = C.FF_PROFILE_PRORES_LT

// FFProfileProresStandard wraps FF_PROFILE_PRORES_STANDARD.
const FFProfileProresStandard = C.FF_PROFILE_PRORES_STANDARD

// FFProfileProresHq wraps FF_PROFILE_PRORES_HQ.
const FFProfileProresHq = C.FF_PROFILE_PRORES_HQ

// FFProfileProres4444 wraps FF_PROFILE_PRORES_4444.
const FFProfileProres4444 = C.FF_PROFILE_PRORES_4444

// FFProfileProresXq wraps FF_PROFILE_PRORES_XQ.
const FFProfileProresXq = C.FF_PROFILE_PRORES_XQ

// FFProfileAribProfileA wraps FF_PROFILE_ARIB_PROFILE_A.
const FFProfileAribProfileA = C.FF_PROFILE_ARIB_PROFILE_A

// FFProfileAribProfileC wraps FF_PROFILE_ARIB_PROFILE_C.
const FFProfileAribProfileC = C.FF_PROFILE_ARIB_PROFILE_C

// FFProfileKlvaSync wraps FF_PROFILE_KLVA_SYNC.
const FFProfileKlvaSync = C.FF_PROFILE_KLVA_SYNC

// FFProfileKlvaAsync wraps FF_PROFILE_KLVA_ASYNC.
const FFProfileKlvaAsync = C.FF_PROFILE_KLVA_ASYNC

// FFProfileEvcBaseline wraps FF_PROFILE_EVC_BASELINE.
const FFProfileEvcBaseline = C.FF_PROFILE_EVC_BASELINE

// FFProfileEvcMain wraps FF_PROFILE_EVC_MAIN.
const FFProfileEvcMain = C.FF_PROFILE_EVC_MAIN

// FFLevelUnknown wraps FF_LEVEL_UNKNOWN.
const FFLevelUnknown = C.FF_LEVEL_UNKNOWN

// FFSubCharencModeDoNothing wraps FF_SUB_CHARENC_MODE_DO_NOTHING.
const FFSubCharencModeDoNothing = C.FF_SUB_CHARENC_MODE_DO_NOTHING

// FFSubCharencModeAutomatic wraps FF_SUB_CHARENC_MODE_AUTOMATIC.
const FFSubCharencModeAutomatic = C.FF_SUB_CHARENC_MODE_AUTOMATIC

// FFSubCharencModePreDecoder wraps FF_SUB_CHARENC_MODE_PRE_DECODER.
const FFSubCharencModePreDecoder = C.FF_SUB_CHARENC_MODE_PRE_DECODER

// FFSubCharencModeIgnore wraps FF_SUB_CHARENC_MODE_IGNORE.
const FFSubCharencModeIgnore = C.FF_SUB_CHARENC_MODE_IGNORE

// FFCodecPropertyLossless wraps FF_CODEC_PROPERTY_LOSSLESS.
const FFCodecPropertyLossless = C.FF_CODEC_PROPERTY_LOSSLESS

// FFCodecPropertyClosedCaptions wraps FF_CODEC_PROPERTY_CLOSED_CAPTIONS.
const FFCodecPropertyClosedCaptions = C.FF_CODEC_PROPERTY_CLOSED_CAPTIONS

// FFCodecPropertyFilmGrain wraps FF_CODEC_PROPERTY_FILM_GRAIN.
const FFCodecPropertyFilmGrain = C.FF_CODEC_PROPERTY_FILM_GRAIN

// AVHWAccelCodecCapExperimental wraps AV_HWACCEL_CODEC_CAP_EXPERIMENTAL.
const AVHWAccelCodecCapExperimental = C.AV_HWACCEL_CODEC_CAP_EXPERIMENTAL

// AVHWAccelFlagIgnoreLevel wraps AV_HWACCEL_FLAG_IGNORE_LEVEL.
const AVHWAccelFlagIgnoreLevel = C.AV_HWACCEL_FLAG_IGNORE_LEVEL

// AVHWAccelFlagAllowHighDepth wraps AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH.
const AVHWAccelFlagAllowHighDepth = C.AV_HWACCEL_FLAG_ALLOW_HIGH_DEPTH

// AVHWAccelFlagAllowProfileMismatch wraps AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH.
const AVHWAccelFlagAllowProfileMismatch = C.AV_HWACCEL_FLAG_ALLOW_PROFILE_MISMATCH

// AVHWAccelFlagUnsafeOutput wraps AV_HWACCEL_FLAG_UNSAFE_OUTPUT.
const AVHWAccelFlagUnsafeOutput = C.AV_HWACCEL_FLAG_UNSAFE_OUTPUT

// AVSubtitleFlagForced wraps AV_SUBTITLE_FLAG_FORCED.
const AVSubtitleFlagForced = C.AV_SUBTITLE_FLAG_FORCED

// AVParserPtsNb wraps AV_PARSER_PTS_NB.
const AVParserPtsNb = C.AV_PARSER_PTS_NB

// ParserFlagCompleteFrames wraps PARSER_FLAG_COMPLETE_FRAMES.
const ParserFlagCompleteFrames = C.PARSER_FLAG_COMPLETE_FRAMES

// ParserFlagOnce wraps PARSER_FLAG_ONCE.
const ParserFlagOnce = C.PARSER_FLAG_ONCE

// ParserFlagFetchedOffset wraps PARSER_FLAG_FETCHED_OFFSET.
const ParserFlagFetchedOffset = C.PARSER_FLAG_FETCHED_OFFSET

// ParserFlagUseCodecTs wraps PARSER_FLAG_USE_CODEC_TS.
const ParserFlagUseCodecTs = C.PARSER_FLAG_USE_CODEC_TS

// AVCodecCapDrawHorizBand wraps AV_CODEC_CAP_DRAW_HORIZ_BAND.
const AVCodecCapDrawHorizBand = C.AV_CODEC_CAP_DRAW_HORIZ_BAND

// AVCodecCapDr1 wraps AV_CODEC_CAP_DR1.
const AVCodecCapDr1 = C.AV_CODEC_CAP_DR1

// AVCodecCapDelay wraps AV_CODEC_CAP_DELAY.
const AVCodecCapDelay = C.AV_CODEC_CAP_DELAY

// AVCodecCapSmallLastFrame wraps AV_CODEC_CAP_SMALL_LAST_FRAME.
const AVCodecCapSmallLastFrame = C.AV_CODEC_CAP_SMALL_LAST_FRAME

// AVCodecCapSubframes wraps AV_CODEC_CAP_SUBFRAMES.
const AVCodecCapSubframes = C.AV_CODEC_CAP_SUBFRAMES

// AVCodecCapExperimental wraps AV_CODEC_CAP_EXPERIMENTAL.
const AVCodecCapExperimental = C.AV_CODEC_CAP_EXPERIMENTAL

// AVCodecCapChannelConf wraps AV_CODEC_CAP_CHANNEL_CONF.
const AVCodecCapChannelConf = C.AV_CODEC_CAP_CHANNEL_CONF

// AVCodecCapFrameThreads wraps AV_CODEC_CAP_FRAME_THREADS.
const AVCodecCapFrameThreads = C.AV_CODEC_CAP_FRAME_THREADS

// AVCodecCapSliceThreads wraps AV_CODEC_CAP_SLICE_THREADS.
const AVCodecCapSliceThreads = C.AV_CODEC_CAP_SLICE_THREADS

// AVCodecCapParamChange wraps AV_CODEC_CAP_PARAM_CHANGE.
const AVCodecCapParamChange = C.AV_CODEC_CAP_PARAM_CHANGE

// AVCodecCapOtherThreads wraps AV_CODEC_CAP_OTHER_THREADS.
const AVCodecCapOtherThreads = C.AV_CODEC_CAP_OTHER_THREADS

// AVCodecCapVariableFrameSize wraps AV_CODEC_CAP_VARIABLE_FRAME_SIZE.
const AVCodecCapVariableFrameSize = C.AV_CODEC_CAP_VARIABLE_FRAME_SIZE

// AVCodecCapAVOidProbing wraps AV_CODEC_CAP_AVOID_PROBING.
const AVCodecCapAVOidProbing = C.AV_CODEC_CAP_AVOID_PROBING

// AVCodecCapHardware wraps AV_CODEC_CAP_HARDWARE.
const AVCodecCapHardware = C.AV_CODEC_CAP_HARDWARE

// AVCodecCapHybrid wraps AV_CODEC_CAP_HYBRID.
const AVCodecCapHybrid = C.AV_CODEC_CAP_HYBRID

// AVCodecCapEncoderReorderedOpaque wraps AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE.
const AVCodecCapEncoderReorderedOpaque = C.AV_CODEC_CAP_ENCODER_REORDERED_OPAQUE

// AVCodecCapEncoderFlush wraps AV_CODEC_CAP_ENCODER_FLUSH.
const AVCodecCapEncoderFlush = C.AV_CODEC_CAP_ENCODER_FLUSH

// AVCodecCapEncoderReconFrame wraps AV_CODEC_CAP_ENCODER_RECON_FRAME.
const AVCodecCapEncoderReconFrame = C.AV_CODEC_CAP_ENCODER_RECON_FRAME

// AVCodecHWConfigMethodHWDeviceCtx wraps AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX.
const AVCodecHWConfigMethodHWDeviceCtx = C.AV_CODEC_HW_CONFIG_METHOD_HW_DEVICE_CTX

// AVCodecHWConfigMethodHWFramesCtx wraps AV_CODEC_HW_CONFIG_METHOD_HW_FRAMES_CTX.
const AVCodecHWConfigMethodHWFramesCtx = C.AV_CODEC_HW_CONFIG_METHOD_HW_FRAMES_CTX

// AVCodecHWConfigMethodInternal wraps AV_CODEC_HW_CONFIG_METHOD_INTERNAL.
const AVCodecHWConfigMethodInternal = C.AV_CODEC_HW_CONFIG_METHOD_INTERNAL

// AVCodecHWConfigMethodAdHoc wraps AV_CODEC_HW_CONFIG_METHOD_AD_HOC.
const AVCodecHWConfigMethodAdHoc = C.AV_CODEC_HW_CONFIG_METHOD_AD_HOC

// AVCodecPropIntraOnly wraps AV_CODEC_PROP_INTRA_ONLY.
const AVCodecPropIntraOnly = C.AV_CODEC_PROP_INTRA_ONLY

// AVCodecPropLossy wraps AV_CODEC_PROP_LOSSY.
const AVCodecPropLossy = C.AV_CODEC_PROP_LOSSY

// AVCodecPropLossless wraps AV_CODEC_PROP_LOSSLESS.
const AVCodecPropLossless = C.AV_CODEC_PROP_LOSSLESS

// AVCodecPropReorder wraps AV_CODEC_PROP_REORDER.
const AVCodecPropReorder = C.AV_CODEC_PROP_REORDER

// AVCodecPropFields wraps AV_CODEC_PROP_FIELDS.
const AVCodecPropFields = C.AV_CODEC_PROP_FIELDS

// AVCodecPropBitmapSub wraps AV_CODEC_PROP_BITMAP_SUB.
const AVCodecPropBitmapSub = C.AV_CODEC_PROP_BITMAP_SUB

// AVCodecPropTextSub wraps AV_CODEC_PROP_TEXT_SUB.
const AVCodecPropTextSub = C.AV_CODEC_PROP_TEXT_SUB

// AVCodecIdIffByterun1 wraps AV_CODEC_ID_IFF_BYTERUN1.
const AVCodecIdIffByterun1 = C.AV_CODEC_ID_IFF_BYTERUN1

// AVCodecIdH265 wraps AV_CODEC_ID_H265.
const AVCodecIdH265 = C.AV_CODEC_ID_H265

// AVCodecIdH266 wraps AV_CODEC_ID_H266.
const AVCodecIdH266 = C.AV_CODEC_ID_H266

// AVInputBufferPaddingSize wraps AV_INPUT_BUFFER_PADDING_SIZE.
const AVInputBufferPaddingSize = C.AV_INPUT_BUFFER_PADDING_SIZE

// AVEfCrccheck wraps AV_EF_CRCCHECK.
const AVEfCrccheck = C.AV_EF_CRCCHECK

// AVEfBitstream wraps AV_EF_BITSTREAM.
const AVEfBitstream = C.AV_EF_BITSTREAM

// AVEfBuffer wraps AV_EF_BUFFER.
const AVEfBuffer = C.AV_EF_BUFFER

// AVEfExplode wraps AV_EF_EXPLODE.
const AVEfExplode = C.AV_EF_EXPLODE

// AVEfIgnoreErr wraps AV_EF_IGNORE_ERR.
const AVEfIgnoreErr = C.AV_EF_IGNORE_ERR

// AVEfCareful wraps AV_EF_CAREFUL.
const AVEfCareful = C.AV_EF_CAREFUL

// AVEfCompliant wraps AV_EF_COMPLIANT.
const AVEfCompliant = C.AV_EF_COMPLIANT

// AVEfAggressive wraps AV_EF_AGGRESSIVE.
const AVEfAggressive = C.AV_EF_AGGRESSIVE

// FFComplianceVeryStrict wraps FF_COMPLIANCE_VERY_STRICT.
const FFComplianceVeryStrict = C.FF_COMPLIANCE_VERY_STRICT

// FFComplianceStrict wraps FF_COMPLIANCE_STRICT.
const FFComplianceStrict = C.FF_COMPLIANCE_STRICT

// FFComplianceNormal wraps FF_COMPLIANCE_NORMAL.
const FFComplianceNormal = C.FF_COMPLIANCE_NORMAL

// FFComplianceUnofficial wraps FF_COMPLIANCE_UNOFFICIAL.
const FFComplianceUnofficial = C.FF_COMPLIANCE_UNOFFICIAL

// FFComplianceExperimental wraps FF_COMPLIANCE_EXPERIMENTAL.
const FFComplianceExperimental = C.FF_COMPLIANCE_EXPERIMENTAL

// AVProfileUnknown wraps AV_PROFILE_UNKNOWN.
const AVProfileUnknown = C.AV_PROFILE_UNKNOWN

// AVProfileReserved wraps AV_PROFILE_RESERVED.
const AVProfileReserved = C.AV_PROFILE_RESERVED

// AVProfileAacMain wraps AV_PROFILE_AAC_MAIN.
const AVProfileAacMain = C.AV_PROFILE_AAC_MAIN

// AVProfileAacLow wraps AV_PROFILE_AAC_LOW.
const AVProfileAacLow = C.AV_PROFILE_AAC_LOW

// AVProfileAacSsr wraps AV_PROFILE_AAC_SSR.
const AVProfileAacSsr = C.AV_PROFILE_AAC_SSR

// AVProfileAacLtp wraps AV_PROFILE_AAC_LTP.
const AVProfileAacLtp = C.AV_PROFILE_AAC_LTP

// AVProfileAacHe wraps AV_PROFILE_AAC_HE.
const AVProfileAacHe = C.AV_PROFILE_AAC_HE

// AVProfileAacHeV2 wraps AV_PROFILE_AAC_HE_V2.
const AVProfileAacHeV2 = C.AV_PROFILE_AAC_HE_V2

// AVProfileAacLd wraps AV_PROFILE_AAC_LD.
const AVProfileAacLd = C.AV_PROFILE_AAC_LD

// AVProfileAacEld wraps AV_PROFILE_AAC_ELD.
const AVProfileAacEld = C.AV_PROFILE_AAC_ELD

// AVProfileMpeg2AacLow wraps AV_PROFILE_MPEG2_AAC_LOW.
const AVProfileMpeg2AacLow = C.AV_PROFILE_MPEG2_AAC_LOW

// AVProfileMpeg2AacHe wraps AV_PROFILE_MPEG2_AAC_HE.
const AVProfileMpeg2AacHe = C.AV_PROFILE_MPEG2_AAC_HE

// AVProfileDnxhd wraps AV_PROFILE_DNXHD.
const AVProfileDnxhd = C.AV_PROFILE_DNXHD

// AVProfileDnxhrLb wraps AV_PROFILE_DNXHR_LB.
const AVProfileDnxhrLb = C.AV_PROFILE_DNXHR_LB

// AVProfileDnxhrSq wraps AV_PROFILE_DNXHR_SQ.
const AVProfileDnxhrSq = C.AV_PROFILE_DNXHR_SQ

// AVProfileDnxhrHq wraps AV_PROFILE_DNXHR_HQ.
const AVProfileDnxhrHq = C.AV_PROFILE_DNXHR_HQ

// AVProfileDnxhrHqx wraps AV_PROFILE_DNXHR_HQX.
const AVProfileDnxhrHqx = C.AV_PROFILE_DNXHR_HQX

// AVProfileDnxhr444 wraps AV_PROFILE_DNXHR_444.
const AVProfileDnxhr444 = C.AV_PROFILE_DNXHR_444

// AVProfileDts wraps AV_PROFILE_DTS.
const AVProfileDts = C.AV_PROFILE_DTS

// AVProfileDtsEs wraps AV_PROFILE_DTS_ES.
const AVProfileDtsEs = C.AV_PROFILE_DTS_ES

// AVProfileDts9624 wraps AV_PROFILE_DTS_96_24.
const AVProfileDts9624 = C.AV_PROFILE_DTS_96_24

// AVProfileDtsHdHra wraps AV_PROFILE_DTS_HD_HRA.
const AVProfileDtsHdHra = C.AV_PROFILE_DTS_HD_HRA

// AVProfileDtsHdMa wraps AV_PROFILE_DTS_HD_MA.
const AVProfileDtsHdMa = C.AV_PROFILE_DTS_HD_MA

// AVProfileDtsExpress wraps AV_PROFILE_DTS_EXPRESS.
const AVProfileDtsExpress = C.AV_PROFILE_DTS_EXPRESS

// AVProfileDtsHdMaX wraps AV_PROFILE_DTS_HD_MA_X.
const AVProfileDtsHdMaX = C.AV_PROFILE_DTS_HD_MA_X

// AVProfileDtsHdMaXImax wraps AV_PROFILE_DTS_HD_MA_X_IMAX.
const AVProfileDtsHdMaXImax = C.AV_PROFILE_DTS_HD_MA_X_IMAX

// AVProfileEac3DdpAtmos wraps AV_PROFILE_EAC3_DDP_ATMOS.
const AVProfileEac3DdpAtmos = C.AV_PROFILE_EAC3_DDP_ATMOS

// AVProfileTruehdAtmos wraps AV_PROFILE_TRUEHD_ATMOS.
const AVProfileTruehdAtmos = C.AV_PROFILE_TRUEHD_ATMOS

// AVProfileMpeg2422 wraps AV_PROFILE_MPEG2_422.
const AVProfileMpeg2422 = C.AV_PROFILE_MPEG2_422

// AVProfileMpeg2High wraps AV_PROFILE_MPEG2_HIGH.
const AVProfileMpeg2High = C.AV_PROFILE_MPEG2_HIGH

// AVProfileMpeg2Ss wraps AV_PROFILE_MPEG2_SS.
const AVProfileMpeg2Ss = C.AV_PROFILE_MPEG2_SS

// AVProfileMpeg2SnrScalable wraps AV_PROFILE_MPEG2_SNR_SCALABLE.
const AVProfileMpeg2SnrScalable = C.AV_PROFILE_MPEG2_SNR_SCALABLE

// AVProfileMpeg2Main wraps AV_PROFILE_MPEG2_MAIN.
const AVProfileMpeg2Main = C.AV_PROFILE_MPEG2_MAIN

// AVProfileMpeg2Simple wraps AV_PROFILE_MPEG2_SIMPLE.
const AVProfileMpeg2Simple = C.AV_PROFILE_MPEG2_SIMPLE

// AVProfileH264Constrained wraps AV_PROFILE_H264_CONSTRAINED.
const AVProfileH264Constrained = C.AV_PROFILE_H264_CONSTRAINED

// AVProfileH264Intra wraps AV_PROFILE_H264_INTRA.
const AVProfileH264Intra = C.AV_PROFILE_H264_INTRA

// AVProfileH264Baseline wraps AV_PROFILE_H264_BASELINE.
const AVProfileH264Baseline = C.AV_PROFILE_H264_BASELINE

// AVProfileH264ConstrainedBaseline wraps AV_PROFILE_H264_CONSTRAINED_BASELINE.
const AVProfileH264ConstrainedBaseline = C.AV_PROFILE_H264_CONSTRAINED_BASELINE

// AVProfileH264Main wraps AV_PROFILE_H264_MAIN.
const AVProfileH264Main = C.AV_PROFILE_H264_MAIN

// AVProfileH264Extended wraps AV_PROFILE_H264_EXTENDED.
const AVProfileH264Extended = C.AV_PROFILE_H264_EXTENDED

// AVProfileH264High wraps AV_PROFILE_H264_HIGH.
const AVProfileH264High = C.AV_PROFILE_H264_HIGH

// AVProfileH264High10 wraps AV_PROFILE_H264_HIGH_10.
const AVProfileH264High10 = C.AV_PROFILE_H264_HIGH_10

// AVProfileH264High10Intra wraps AV_PROFILE_H264_HIGH_10_INTRA.
const AVProfileH264High10Intra = C.AV_PROFILE_H264_HIGH_10_INTRA

// AVProfileH264MultiviewHigh wraps AV_PROFILE_H264_MULTIVIEW_HIGH.
const AVProfileH264MultiviewHigh = C.AV_PROFILE_H264_MULTIVIEW_HIGH

// AVProfileH264High422 wraps AV_PROFILE_H264_HIGH_422.
const AVProfileH264High422 = C.AV_PROFILE_H264_HIGH_422

// AVProfileH264High422Intra wraps AV_PROFILE_H264_HIGH_422_INTRA.
const AVProfileH264High422Intra = C.AV_PROFILE_H264_HIGH_422_INTRA

// AVProfileH264StereoHigh wraps AV_PROFILE_H264_STEREO_HIGH.
const AVProfileH264StereoHigh = C.AV_PROFILE_H264_STEREO_HIGH

// AVProfileH264High444 wraps AV_PROFILE_H264_HIGH_444.
const AVProfileH264High444 = C.AV_PROFILE_H264_HIGH_444

// AVProfileH264High444Predictive wraps AV_PROFILE_H264_HIGH_444_PREDICTIVE.
const AVProfileH264High444Predictive = C.AV_PROFILE_H264_HIGH_444_PREDICTIVE

// AVProfileH264High444Intra wraps AV_PROFILE_H264_HIGH_444_INTRA.
const AVProfileH264High444Intra = C.AV_PROFILE_H264_HIGH_444_INTRA

// AVProfileH264Cavlc444 wraps AV_PROFILE_H264_CAVLC_444.
const AVProfileH264Cavlc444 = C.AV_PROFILE_H264_CAVLC_444

// AVProfileVc1Simple wraps AV_PROFILE_VC1_SIMPLE.
const AVProfileVc1Simple = C.AV_PROFILE_VC1_SIMPLE

// AVProfileVc1Main wraps AV_PROFILE_VC1_MAIN.
const AVProfileVc1Main = C.AV_PROFILE_VC1_MAIN

// AVProfileVc1Complex wraps AV_PROFILE_VC1_COMPLEX.
const AVProfileVc1Complex = C.AV_PROFILE_VC1_COMPLEX

// AVProfileVc1Advanced wraps AV_PROFILE_VC1_ADVANCED.
const AVProfileVc1Advanced = C.AV_PROFILE_VC1_ADVANCED

// AVProfileMpeg4Simple wraps AV_PROFILE_MPEG4_SIMPLE.
const AVProfileMpeg4Simple = C.AV_PROFILE_MPEG4_SIMPLE

// AVProfileMpeg4SimpleScalable wraps AV_PROFILE_MPEG4_SIMPLE_SCALABLE.
const AVProfileMpeg4SimpleScalable = C.AV_PROFILE_MPEG4_SIMPLE_SCALABLE

// AVProfileMpeg4Core wraps AV_PROFILE_MPEG4_CORE.
const AVProfileMpeg4Core = C.AV_PROFILE_MPEG4_CORE

// AVProfileMpeg4Main wraps AV_PROFILE_MPEG4_MAIN.
const AVProfileMpeg4Main = C.AV_PROFILE_MPEG4_MAIN

// AVProfileMpeg4NBit wraps AV_PROFILE_MPEG4_N_BIT.
const AVProfileMpeg4NBit = C.AV_PROFILE_MPEG4_N_BIT

// AVProfileMpeg4ScalableTexture wraps AV_PROFILE_MPEG4_SCALABLE_TEXTURE.
const AVProfileMpeg4ScalableTexture = C.AV_PROFILE_MPEG4_SCALABLE_TEXTURE

// AVProfileMpeg4SimpleFaceAnimation wraps AV_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION.
const AVProfileMpeg4SimpleFaceAnimation = C.AV_PROFILE_MPEG4_SIMPLE_FACE_ANIMATION

// AVProfileMpeg4BasicAnimatedTexture wraps AV_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE.
const AVProfileMpeg4BasicAnimatedTexture = C.AV_PROFILE_MPEG4_BASIC_ANIMATED_TEXTURE

// AVProfileMpeg4Hybrid wraps AV_PROFILE_MPEG4_HYBRID.
const AVProfileMpeg4Hybrid = C.AV_PROFILE_MPEG4_HYBRID

// AVProfileMpeg4AdvancedRealTime wraps AV_PROFILE_MPEG4_ADVANCED_REAL_TIME.
const AVProfileMpeg4AdvancedRealTime = C.AV_PROFILE_MPEG4_ADVANCED_REAL_TIME

// AVProfileMpeg4CoreScalable wraps AV_PROFILE_MPEG4_CORE_SCALABLE.
const AVProfileMpeg4CoreScalable = C.AV_PROFILE_MPEG4_CORE_SCALABLE

// AVProfileMpeg4AdvancedCoding wraps AV_PROFILE_MPEG4_ADVANCED_CODING.
const AVProfileMpeg4AdvancedCoding = C.AV_PROFILE_MPEG4_ADVANCED_CODING

// AVProfileMpeg4AdvancedCore wraps AV_PROFILE_MPEG4_ADVANCED_CORE.
const AVProfileMpeg4AdvancedCore = C.AV_PROFILE_MPEG4_ADVANCED_CORE

// AVProfileMpeg4AdvancedScalableTexture wraps AV_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE.
const AVProfileMpeg4AdvancedScalableTexture = C.AV_PROFILE_MPEG4_ADVANCED_SCALABLE_TEXTURE

// AVProfileMpeg4SimpleStudio wraps AV_PROFILE_MPEG4_SIMPLE_STUDIO.
const AVProfileMpeg4SimpleStudio = C.AV_PROFILE_MPEG4_SIMPLE_STUDIO

// AVProfileMpeg4AdvancedSimple wraps AV_PROFILE_MPEG4_ADVANCED_SIMPLE.
const AVProfileMpeg4AdvancedSimple = C.AV_PROFILE_MPEG4_ADVANCED_SIMPLE

// AVProfileJpeg2000CstreamRestriction0 wraps AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0.
const AVProfileJpeg2000CstreamRestriction0 = C.AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_0

// AVProfileJpeg2000CstreamRestriction1 wraps AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1.
const AVProfileJpeg2000CstreamRestriction1 = C.AV_PROFILE_JPEG2000_CSTREAM_RESTRICTION_1

// AVProfileJpeg2000CstreamNoRestriction wraps AV_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION.
const AVProfileJpeg2000CstreamNoRestriction = C.AV_PROFILE_JPEG2000_CSTREAM_NO_RESTRICTION

// AVProfileJpeg2000Dcinema2K wraps AV_PROFILE_JPEG2000_DCINEMA_2K.
const AVProfileJpeg2000Dcinema2K = C.AV_PROFILE_JPEG2000_DCINEMA_2K

// AVProfileJpeg2000Dcinema4K wraps AV_PROFILE_JPEG2000_DCINEMA_4K.
const AVProfileJpeg2000Dcinema4K = C.AV_PROFILE_JPEG2000_DCINEMA_4K

// AVProfileVp90 wraps AV_PROFILE_VP9_0.
const AVProfileVp90 = C.AV_PROFILE_VP9_0

// AVProfileVp91 wraps AV_PROFILE_VP9_1.
const AVProfileVp91 = C.AV_PROFILE_VP9_1

// AVProfileVp92 wraps AV_PROFILE_VP9_2.
const AVProfileVp92 = C.AV_PROFILE_VP9_2

// AVProfileVp93 wraps AV_PROFILE_VP9_3.
const AVProfileVp93 = C.AV_PROFILE_VP9_3

// AVProfileHevcMain wraps AV_PROFILE_HEVC_MAIN.
const AVProfileHevcMain = C.AV_PROFILE_HEVC_MAIN

// AVProfileHevcMain10 wraps AV_PROFILE_HEVC_MAIN_10.
const AVProfileHevcMain10 = C.AV_PROFILE_HEVC_MAIN_10

// AVProfileHevcMainStillPicture wraps AV_PROFILE_HEVC_MAIN_STILL_PICTURE.
const AVProfileHevcMainStillPicture = C.AV_PROFILE_HEVC_MAIN_STILL_PICTURE

// AVProfileHevcRext wraps AV_PROFILE_HEVC_REXT.
const AVProfileHevcRext = C.AV_PROFILE_HEVC_REXT

// AVProfileHevcScc wraps AV_PROFILE_HEVC_SCC.
const AVProfileHevcScc = C.AV_PROFILE_HEVC_SCC

// AVProfileVvcMain10 wraps AV_PROFILE_VVC_MAIN_10.
const AVProfileVvcMain10 = C.AV_PROFILE_VVC_MAIN_10

// AVProfileVvcMain10444 wraps AV_PROFILE_VVC_MAIN_10_444.
const AVProfileVvcMain10444 = C.AV_PROFILE_VVC_MAIN_10_444

// AVProfileAV1Main wraps AV_PROFILE_AV1_MAIN.
const AVProfileAV1Main = C.AV_PROFILE_AV1_MAIN

// AVProfileAV1High wraps AV_PROFILE_AV1_HIGH.
const AVProfileAV1High = C.AV_PROFILE_AV1_HIGH

// AVProfileAV1Professional wraps AV_PROFILE_AV1_PROFESSIONAL.
const AVProfileAV1Professional = C.AV_PROFILE_AV1_PROFESSIONAL

// AVProfileMjpegHuffmanBaselineDct wraps AV_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT.
const AVProfileMjpegHuffmanBaselineDct = C.AV_PROFILE_MJPEG_HUFFMAN_BASELINE_DCT

// AVProfileMjpegHuffmanExtendedSequentialDct wraps AV_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT.
const AVProfileMjpegHuffmanExtendedSequentialDct = C.AV_PROFILE_MJPEG_HUFFMAN_EXTENDED_SEQUENTIAL_DCT

// AVProfileMjpegHuffmanProgressiveDct wraps AV_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT.
const AVProfileMjpegHuffmanProgressiveDct = C.AV_PROFILE_MJPEG_HUFFMAN_PROGRESSIVE_DCT

// AVProfileMjpegHuffmanLossless wraps AV_PROFILE_MJPEG_HUFFMAN_LOSSLESS.
const AVProfileMjpegHuffmanLossless = C.AV_PROFILE_MJPEG_HUFFMAN_LOSSLESS

// AVProfileMjpegJpegLs wraps AV_PROFILE_MJPEG_JPEG_LS.
const AVProfileMjpegJpegLs = C.AV_PROFILE_MJPEG_JPEG_LS

// AVProfileSbcMsbc wraps AV_PROFILE_SBC_MSBC.
const AVProfileSbcMsbc = C.AV_PROFILE_SBC_MSBC

// AVProfileProresProxy wraps AV_PROFILE_PRORES_PROXY.
const AVProfileProresProxy = C.AV_PROFILE_PRORES_PROXY

// AVProfileProresLt wraps AV_PROFILE_PRORES_LT.
const AVProfileProresLt = C.AV_PROFILE_PRORES_LT

// AVProfileProresStandard wraps AV_PROFILE_PRORES_STANDARD.
const AVProfileProresStandard = C.AV_PROFILE_PRORES_STANDARD

// AVProfileProresHq wraps AV_PROFILE_PRORES_HQ.
const AVProfileProresHq = C.AV_PROFILE_PRORES_HQ

// AVProfileProres4444 wraps AV_PROFILE_PRORES_4444.
const AVProfileProres4444 = C.AV_PROFILE_PRORES_4444

// AVProfileProresXq wraps AV_PROFILE_PRORES_XQ.
const AVProfileProresXq = C.AV_PROFILE_PRORES_XQ

// AVProfileAribProfileA wraps AV_PROFILE_ARIB_PROFILE_A.
const AVProfileAribProfileA = C.AV_PROFILE_ARIB_PROFILE_A

// AVProfileAribProfileC wraps AV_PROFILE_ARIB_PROFILE_C.
const AVProfileAribProfileC = C.AV_PROFILE_ARIB_PROFILE_C

// AVProfileKlvaSync wraps AV_PROFILE_KLVA_SYNC.
const AVProfileKlvaSync = C.AV_PROFILE_KLVA_SYNC

// AVProfileKlvaAsync wraps AV_PROFILE_KLVA_ASYNC.
const AVProfileKlvaAsync = C.AV_PROFILE_KLVA_ASYNC

// AVProfileEvcBaseline wraps AV_PROFILE_EVC_BASELINE.
const AVProfileEvcBaseline = C.AV_PROFILE_EVC_BASELINE

// AVProfileEvcMain wraps AV_PROFILE_EVC_MAIN.
const AVProfileEvcMain = C.AV_PROFILE_EVC_MAIN

// AVLevelUnknown wraps AV_LEVEL_UNKNOWN.
const AVLevelUnknown = C.AV_LEVEL_UNKNOWN

// AVPktDataQualityFactor wraps AV_PKT_DATA_QUALITY_FACTOR.
const AVPktDataQualityFactor = C.AV_PKT_DATA_QUALITY_FACTOR

// AVPktFlagKey wraps AV_PKT_FLAG_KEY.
const AVPktFlagKey = C.AV_PKT_FLAG_KEY

// AVPktFlagCorrupt wraps AV_PKT_FLAG_CORRUPT.
const AVPktFlagCorrupt = C.AV_PKT_FLAG_CORRUPT

// AVPktFlagDiscard wraps AV_PKT_FLAG_DISCARD.
const AVPktFlagDiscard = C.AV_PKT_FLAG_DISCARD

// AVPktFlagTrusted wraps AV_PKT_FLAG_TRUSTED.
const AVPktFlagTrusted = C.AV_PKT_FLAG_TRUSTED

// AVPktFlagDisposable wraps AV_PKT_FLAG_DISPOSABLE.
const AVPktFlagDisposable = C.AV_PKT_FLAG_DISPOSABLE

// LIBAVCodecVersionMinor wraps LIBAVCODEC_VERSION_MINOR.
const LIBAVCodecVersionMinor = C.LIBAVCODEC_VERSION_MINOR

// LIBAVCodecVersionMicro wraps LIBAVCODEC_VERSION_MICRO.
const LIBAVCodecVersionMicro = C.LIBAVCODEC_VERSION_MICRO

// LIBAVCodecVersionInt wraps LIBAVCODEC_VERSION_INT.
const LIBAVCodecVersionInt = C.LIBAVCODEC_VERSION_INT

// LIBAVCodecBuild wraps LIBAVCODEC_BUILD.
const LIBAVCodecBuild = C.LIBAVCODEC_BUILD

// LIBAVCodecVersionMajor wraps LIBAVCODEC_VERSION_MAJOR.
const LIBAVCodecVersionMajor = C.LIBAVCODEC_VERSION_MAJOR

// FFAPIInitPacket wraps FF_API_INIT_PACKET.
const FFAPIInitPacket = C.FF_API_INIT_PACKET

// FFAPIIdctNone wraps FF_API_IDCT_NONE.
const FFAPIIdctNone = C.FF_API_IDCT_NONE

// FFAPISvtav1Opts wraps FF_API_SVTAV1_OPTS.
const FFAPISvtav1Opts = C.FF_API_SVTAV1_OPTS

// FFAPIAyuvCodecid wraps FF_API_AYUV_CODECID.
const FFAPIAyuvCodecid = C.FF_API_AYUV_CODECID

// FFAPIVtOutputCallback wraps FF_API_VT_OUTPUT_CALLBACK.
const FFAPIVtOutputCallback = C.FF_API_VT_OUTPUT_CALLBACK

// FFAPIAVCodecChromaPos wraps FF_API_AVCODEC_CHROMA_POS.
const FFAPIAVCodecChromaPos = C.FF_API_AVCODEC_CHROMA_POS

// FFAPIVtHWAccelContext wraps FF_API_VT_HWACCEL_CONTEXT.
const FFAPIVtHWAccelContext = C.FF_API_VT_HWACCEL_CONTEXT

// FFAPIAVCtxFrameNumber wraps FF_API_AVCTX_FRAME_NUMBER.
const FFAPIAVCtxFrameNumber = C.FF_API_AVCTX_FRAME_NUMBER

// FFAPISliceOffset wraps FF_API_SLICE_OFFSET.
const FFAPISliceOffset = C.FF_API_SLICE_OFFSET

// FFAPISubframes wraps FF_API_SUBFRAMES.
const FFAPISubframes = C.FF_API_SUBFRAMES

// FFAPITicksPerFrame wraps FF_API_TICKS_PER_FRAME.
const FFAPITicksPerFrame = C.FF_API_TICKS_PER_FRAME

// FFAPIDropchanged wraps FF_API_DROPCHANGED.
const FFAPIDropchanged = C.FF_API_DROPCHANGED

// FFAPIAVFFT wraps FF_API_AVFFT.
const FFAPIAVFFT = C.FF_API_AVFFT

// FFAPIFFProfileLevel wraps FF_API_FF_PROFILE_LEVEL.
const FFAPIFFProfileLevel = C.FF_API_FF_PROFILE_LEVEL

// FFCodecCrystalHd wraps FF_CODEC_CRYSTAL_HD.
const FFCodecCrystalHd = C.FF_CODEC_CRYSTAL_HD

// LIBAVDeviceVersionMinor wraps LIBAVDEVICE_VERSION_MINOR.
const LIBAVDeviceVersionMinor = C.LIBAVDEVICE_VERSION_MINOR

// LIBAVDeviceVersionMicro wraps LIBAVDEVICE_VERSION_MICRO.
const LIBAVDeviceVersionMicro = C.LIBAVDEVICE_VERSION_MICRO

// LIBAVDeviceVersionInt wraps LIBAVDEVICE_VERSION_INT.
const LIBAVDeviceVersionInt = C.LIBAVDEVICE_VERSION_INT

// LIBAVDeviceBuild wraps LIBAVDEVICE_BUILD.
const LIBAVDeviceBuild = C.LIBAVDEVICE_BUILD

// LIBAVDeviceVersionMajor wraps LIBAVDEVICE_VERSION_MAJOR.
const LIBAVDeviceVersionMajor = C.LIBAVDEVICE_VERSION_MAJOR

// AVFilterFlagDynamicInputs wraps AVFILTER_FLAG_DYNAMIC_INPUTS.
const AVFilterFlagDynamicInputs = C.AVFILTER_FLAG_DYNAMIC_INPUTS

// AVFilterFlagDynamicOutputs wraps AVFILTER_FLAG_DYNAMIC_OUTPUTS.
const AVFilterFlagDynamicOutputs = C.AVFILTER_FLAG_DYNAMIC_OUTPUTS

// AVFilterFlagSliceThreads wraps AVFILTER_FLAG_SLICE_THREADS.
const AVFilterFlagSliceThreads = C.AVFILTER_FLAG_SLICE_THREADS

// AVFilterFlagMetadataOnly wraps AVFILTER_FLAG_METADATA_ONLY.
const AVFilterFlagMetadataOnly = C.AVFILTER_FLAG_METADATA_ONLY

// AVFilterFlagHWDevice wraps AVFILTER_FLAG_HWDEVICE.
const AVFilterFlagHWDevice = C.AVFILTER_FLAG_HWDEVICE

// AVFilterFlagSupportTimelineGeneric wraps AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC.
const AVFilterFlagSupportTimelineGeneric = C.AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC

// AVFilterFlagSupportTimelineInternal wraps AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL.
const AVFilterFlagSupportTimelineInternal = C.AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL

// AVFilterFlagSupportTimeline wraps AVFILTER_FLAG_SUPPORT_TIMELINE.
const AVFilterFlagSupportTimeline = C.AVFILTER_FLAG_SUPPORT_TIMELINE

// AVFilterThreadSlice wraps AVFILTER_THREAD_SLICE.
const AVFilterThreadSlice = C.AVFILTER_THREAD_SLICE

// AVFilterCmdFlagOne wraps AVFILTER_CMD_FLAG_ONE.
const AVFilterCmdFlagOne = C.AVFILTER_CMD_FLAG_ONE

// AVFilterCmdFlagFast wraps AVFILTER_CMD_FLAG_FAST.
const AVFilterCmdFlagFast = C.AVFILTER_CMD_FLAG_FAST

// AVLinkUninit wraps AVLINK_UNINIT.
const AVLinkUninit = C.AVLINK_UNINIT

// AVLinkStartinit wraps AVLINK_STARTINIT.
const AVLinkStartinit = C.AVLINK_STARTINIT

// AVLinkInit wraps AVLINK_INIT.
const AVLinkInit = C.AVLINK_INIT

// AVFilterAutoConvertAll wraps AVFILTER_AUTO_CONVERT_ALL.
const AVFilterAutoConvertAll = C.AVFILTER_AUTO_CONVERT_ALL

// AVFilterAutoConvertNone wraps AVFILTER_AUTO_CONVERT_NONE.
const AVFilterAutoConvertNone = C.AVFILTER_AUTO_CONVERT_NONE

// AVBuffersinkFlagPeek wraps AV_BUFFERSINK_FLAG_PEEK.
const AVBuffersinkFlagPeek = C.AV_BUFFERSINK_FLAG_PEEK

// AVBuffersinkFlagNoRequest wraps AV_BUFFERSINK_FLAG_NO_REQUEST.
const AVBuffersinkFlagNoRequest = C.AV_BUFFERSINK_FLAG_NO_REQUEST

// AVBuffersrcFlagNoCheckFormat wraps AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT.
const AVBuffersrcFlagNoCheckFormat = C.AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT

// AVBuffersrcFlagPush wraps AV_BUFFERSRC_FLAG_PUSH.
const AVBuffersrcFlagPush = C.AV_BUFFERSRC_FLAG_PUSH

// AVBuffersrcFlagKeepRef wraps AV_BUFFERSRC_FLAG_KEEP_REF.
const AVBuffersrcFlagKeepRef = C.AV_BUFFERSRC_FLAG_KEEP_REF

// LIBAVFilterVersionMinor wraps LIBAVFILTER_VERSION_MINOR.
const LIBAVFilterVersionMinor = C.LIBAVFILTER_VERSION_MINOR

// LIBAVFilterVersionMicro wraps LIBAVFILTER_VERSION_MICRO.
const LIBAVFilterVersionMicro = C.LIBAVFILTER_VERSION_MICRO

// LIBAVFilterVersionInt wraps LIBAVFILTER_VERSION_INT.
const LIBAVFilterVersionInt = C.LIBAVFILTER_VERSION_INT

// LIBAVFilterBuild wraps LIBAVFILTER_BUILD.
const LIBAVFilterBuild = C.LIBAVFILTER_BUILD

// LIBAVFilterVersionMajor wraps LIBAVFILTER_VERSION_MAJOR.
const LIBAVFilterVersionMajor = C.LIBAVFILTER_VERSION_MAJOR

// FFAPILIBPlaceboOpts wraps FF_API_LIBPLACEBO_OPTS.
const FFAPILIBPlaceboOpts = C.FF_API_LIBPLACEBO_OPTS

// AVProbeScoreRetry wraps AVPROBE_SCORE_RETRY.
const AVProbeScoreRetry = C.AVPROBE_SCORE_RETRY

// AVProbeScoreStreamRetry wraps AVPROBE_SCORE_STREAM_RETRY.
const AVProbeScoreStreamRetry = C.AVPROBE_SCORE_STREAM_RETRY

// AVProbeScoreExtension wraps AVPROBE_SCORE_EXTENSION.
const AVProbeScoreExtension = C.AVPROBE_SCORE_EXTENSION

// AVProbeScoreMime wraps AVPROBE_SCORE_MIME.
const AVProbeScoreMime = C.AVPROBE_SCORE_MIME

// AVProbeScoreMax wraps AVPROBE_SCORE_MAX.
const AVProbeScoreMax = C.AVPROBE_SCORE_MAX

// AVProbePaddingSize wraps AVPROBE_PADDING_SIZE.
const AVProbePaddingSize = C.AVPROBE_PADDING_SIZE

// AVFmtNofile wraps AVFMT_NOFILE.
const AVFmtNofile = C.AVFMT_NOFILE

// AVFmtNeednumber wraps AVFMT_NEEDNUMBER.
const AVFmtNeednumber = C.AVFMT_NEEDNUMBER

// AVFmtExperimental wraps AVFMT_EXPERIMENTAL.
const AVFmtExperimental = C.AVFMT_EXPERIMENTAL

// AVFmtShowIds wraps AVFMT_SHOW_IDS.
const AVFmtShowIds = C.AVFMT_SHOW_IDS

// AVFmtGlobalheader wraps AVFMT_GLOBALHEADER.
const AVFmtGlobalheader = C.AVFMT_GLOBALHEADER

// AVFmtNotimestamps wraps AVFMT_NOTIMESTAMPS.
const AVFmtNotimestamps = C.AVFMT_NOTIMESTAMPS

// AVFmtGenericIndex wraps AVFMT_GENERIC_INDEX.
const AVFmtGenericIndex = C.AVFMT_GENERIC_INDEX

// AVFmtTsDiscont wraps AVFMT_TS_DISCONT.
const AVFmtTsDiscont = C.AVFMT_TS_DISCONT

// AVFmtVariableFps wraps AVFMT_VARIABLE_FPS.
const AVFmtVariableFps = C.AVFMT_VARIABLE_FPS

// AVFmtNodimensions wraps AVFMT_NODIMENSIONS.
const AVFmtNodimensions = C.AVFMT_NODIMENSIONS

// AVFmtNostreams wraps AVFMT_NOSTREAMS.
const AVFmtNostreams = C.AVFMT_NOSTREAMS

// AVFmtNobinsearch wraps AVFMT_NOBINSEARCH.
const AVFmtNobinsearch = C.AVFMT_NOBINSEARCH

// AVFmtNogensearch wraps AVFMT_NOGENSEARCH.
const AVFmtNogensearch = C.AVFMT_NOGENSEARCH

// AVFmtNoByteSeek wraps AVFMT_NO_BYTE_SEEK.
const AVFmtNoByteSeek = C.AVFMT_NO_BYTE_SEEK

// AVFmtAllowFlush wraps AVFMT_ALLOW_FLUSH.
const AVFmtAllowFlush = C.AVFMT_ALLOW_FLUSH

// AVFmtTsNonstrict wraps AVFMT_TS_NONSTRICT.
const AVFmtTsNonstrict = C.AVFMT_TS_NONSTRICT

// AVFmtTsNegative wraps AVFMT_TS_NEGATIVE.
const AVFmtTsNegative = C.AVFMT_TS_NEGATIVE

// AVFmtSeekToPts wraps AVFMT_SEEK_TO_PTS.
const AVFmtSeekToPts = C.AVFMT_SEEK_TO_PTS

// AVIndexKeyframe wraps AVINDEX_KEYFRAME.
const AVIndexKeyframe = C.AVINDEX_KEYFRAME

// AVIndexDiscardFrame wraps AVINDEX_DISCARD_FRAME.
const AVIndexDiscardFrame = C.AVINDEX_DISCARD_FRAME

// AVDispositionDefault wraps AV_DISPOSITION_DEFAULT.
const AVDispositionDefault = C.AV_DISPOSITION_DEFAULT

// AVDispositionDub wraps AV_DISPOSITION_DUB.
const AVDispositionDub = C.AV_DISPOSITION_DUB

// AVDispositionOriginal wraps AV_DISPOSITION_ORIGINAL.
const AVDispositionOriginal = C.AV_DISPOSITION_ORIGINAL

// AVDispositionComment wraps AV_DISPOSITION_COMMENT.
const AVDispositionComment = C.AV_DISPOSITION_COMMENT

// AVDispositionLyrics wraps AV_DISPOSITION_LYRICS.
const AVDispositionLyrics = C.AV_DISPOSITION_LYRICS

// AVDispositionKaraoke wraps AV_DISPOSITION_KARAOKE.
const AVDispositionKaraoke = C.AV_DISPOSITION_KARAOKE

// AVDispositionForced wraps AV_DISPOSITION_FORCED.
const AVDispositionForced = C.AV_DISPOSITION_FORCED

// AVDispositionHearingImpaired wraps AV_DISPOSITION_HEARING_IMPAIRED.
const AVDispositionHearingImpaired = C.AV_DISPOSITION_HEARING_IMPAIRED

// AVDispositionVisualImpaired wraps AV_DISPOSITION_VISUAL_IMPAIRED.
const AVDispositionVisualImpaired = C.AV_DISPOSITION_VISUAL_IMPAIRED

// AVDispositionCleanEffects wraps AV_DISPOSITION_CLEAN_EFFECTS.
const AVDispositionCleanEffects = C.AV_DISPOSITION_CLEAN_EFFECTS

// AVDispositionAttachedPic wraps AV_DISPOSITION_ATTACHED_PIC.
const AVDispositionAttachedPic = C.AV_DISPOSITION_ATTACHED_PIC

// AVDispositionTimedThumbnails wraps AV_DISPOSITION_TIMED_THUMBNAILS.
const AVDispositionTimedThumbnails = C.AV_DISPOSITION_TIMED_THUMBNAILS

// AVDispositionNonDiegetic wraps AV_DISPOSITION_NON_DIEGETIC.
const AVDispositionNonDiegetic = C.AV_DISPOSITION_NON_DIEGETIC

// AVDispositionCaptions wraps AV_DISPOSITION_CAPTIONS.
const AVDispositionCaptions = C.AV_DISPOSITION_CAPTIONS

// AVDispositionDescriptions wraps AV_DISPOSITION_DESCRIPTIONS.
const AVDispositionDescriptions = C.AV_DISPOSITION_DESCRIPTIONS

// AVDispositionMetadata wraps AV_DISPOSITION_METADATA.
const AVDispositionMetadata = C.AV_DISPOSITION_METADATA

// AVDispositionDependent wraps AV_DISPOSITION_DEPENDENT.
const AVDispositionDependent = C.AV_DISPOSITION_DEPENDENT

// AVDispositionStillImage wraps AV_DISPOSITION_STILL_IMAGE.
const AVDispositionStillImage = C.AV_DISPOSITION_STILL_IMAGE

// AVPtsWrapIgnore wraps AV_PTS_WRAP_IGNORE.
const AVPtsWrapIgnore = C.AV_PTS_WRAP_IGNORE

// AVPtsWrapAddOffset wraps AV_PTS_WRAP_ADD_OFFSET.
const AVPtsWrapAddOffset = C.AV_PTS_WRAP_ADD_OFFSET

// AVPtsWrapSubOffset wraps AV_PTS_WRAP_SUB_OFFSET.
const AVPtsWrapSubOffset = C.AV_PTS_WRAP_SUB_OFFSET

// AVStreamEventFlagMetadataUpdated wraps AVSTREAM_EVENT_FLAG_METADATA_UPDATED.
const AVStreamEventFlagMetadataUpdated = C.AVSTREAM_EVENT_FLAG_METADATA_UPDATED

// AVStreamEventFlagNewPackets wraps AVSTREAM_EVENT_FLAG_NEW_PACKETS.
const AVStreamEventFlagNewPackets = C.AVSTREAM_EVENT_FLAG_NEW_PACKETS

// AVProgramRunning wraps AV_PROGRAM_RUNNING.
const AVProgramRunning = C.AV_PROGRAM_RUNNING

// AVFmtctxNoheader wraps AVFMTCTX_NOHEADER.
const AVFmtctxNoheader = C.AVFMTCTX_NOHEADER

// AVFmtctxUnseekable wraps AVFMTCTX_UNSEEKABLE.
const AVFmtctxUnseekable = C.AVFMTCTX_UNSEEKABLE

// AVFmtFlagGenpts wraps AVFMT_FLAG_GENPTS.
const AVFmtFlagGenpts = C.AVFMT_FLAG_GENPTS

// AVFmtFlagIgnidx wraps AVFMT_FLAG_IGNIDX.
const AVFmtFlagIgnidx = C.AVFMT_FLAG_IGNIDX

// AVFmtFlagNonblock wraps AVFMT_FLAG_NONBLOCK.
const AVFmtFlagNonblock = C.AVFMT_FLAG_NONBLOCK

// AVFmtFlagIgndts wraps AVFMT_FLAG_IGNDTS.
const AVFmtFlagIgndts = C.AVFMT_FLAG_IGNDTS

// AVFmtFlagNofillin wraps AVFMT_FLAG_NOFILLIN.
const AVFmtFlagNofillin = C.AVFMT_FLAG_NOFILLIN

// AVFmtFlagNoparse wraps AVFMT_FLAG_NOPARSE.
const AVFmtFlagNoparse = C.AVFMT_FLAG_NOPARSE

// AVFmtFlagNobuffer wraps AVFMT_FLAG_NOBUFFER.
const AVFmtFlagNobuffer = C.AVFMT_FLAG_NOBUFFER

// AVFmtFlagCustomIO wraps AVFMT_FLAG_CUSTOM_IO.
const AVFmtFlagCustomIO = C.AVFMT_FLAG_CUSTOM_IO

// AVFmtFlagDiscardCorrupt wraps AVFMT_FLAG_DISCARD_CORRUPT.
const AVFmtFlagDiscardCorrupt = C.AVFMT_FLAG_DISCARD_CORRUPT

// AVFmtFlagFlushPackets wraps AVFMT_FLAG_FLUSH_PACKETS.
const AVFmtFlagFlushPackets = C.AVFMT_FLAG_FLUSH_PACKETS

// AVFmtFlagBitexact wraps AVFMT_FLAG_BITEXACT.
const AVFmtFlagBitexact = C.AVFMT_FLAG_BITEXACT

// AVFmtFlagSortDts wraps AVFMT_FLAG_SORT_DTS.
const AVFmtFlagSortDts = C.AVFMT_FLAG_SORT_DTS

// AVFmtFlagFastSeek wraps AVFMT_FLAG_FAST_SEEK.
const AVFmtFlagFastSeek = C.AVFMT_FLAG_FAST_SEEK

// AVFmtFlagShortest wraps AVFMT_FLAG_SHORTEST.
const AVFmtFlagShortest = C.AVFMT_FLAG_SHORTEST

// AVFmtFlagAutoBsf wraps AVFMT_FLAG_AUTO_BSF.
const AVFmtFlagAutoBsf = C.AVFMT_FLAG_AUTO_BSF

// FFFdebugTs wraps FF_FDEBUG_TS.
const FFFdebugTs = C.FF_FDEBUG_TS

// AVFmtEventFlagMetadataUpdated wraps AVFMT_EVENT_FLAG_METADATA_UPDATED.
const AVFmtEventFlagMetadataUpdated = C.AVFMT_EVENT_FLAG_METADATA_UPDATED

// AVFmtAVOidNegTsAuto wraps AVFMT_AVOID_NEG_TS_AUTO.
const AVFmtAVOidNegTsAuto = C.AVFMT_AVOID_NEG_TS_AUTO

// AVFmtAVOidNegTsDisabled wraps AVFMT_AVOID_NEG_TS_DISABLED.
const AVFmtAVOidNegTsDisabled = C.AVFMT_AVOID_NEG_TS_DISABLED

// AVFmtAVOidNegTsMakeNonNegative wraps AVFMT_AVOID_NEG_TS_MAKE_NON_NEGATIVE.
const AVFmtAVOidNegTsMakeNonNegative = C.AVFMT_AVOID_NEG_TS_MAKE_NON_NEGATIVE

// AVFmtAVOidNegTsMakeZero wraps AVFMT_AVOID_NEG_TS_MAKE_ZERO.
const AVFmtAVOidNegTsMakeZero = C.AVFMT_AVOID_NEG_TS_MAKE_ZERO

// AVSeekFlagBackward wraps AVSEEK_FLAG_BACKWARD.
const AVSeekFlagBackward = C.AVSEEK_FLAG_BACKWARD

// AVSeekFlagByte wraps AVSEEK_FLAG_BYTE.
const AVSeekFlagByte = C.AVSEEK_FLAG_BYTE

// AVSeekFlagAny wraps AVSEEK_FLAG_ANY.
const AVSeekFlagAny = C.AVSEEK_FLAG_ANY

// AVSeekFlagFrame wraps AVSEEK_FLAG_FRAME.
const AVSeekFlagFrame = C.AVSEEK_FLAG_FRAME

// AVStreamInitInWriteHeader wraps AVSTREAM_INIT_IN_WRITE_HEADER.
const AVStreamInitInWriteHeader = C.AVSTREAM_INIT_IN_WRITE_HEADER

// AVStreamInitInInitOutput wraps AVSTREAM_INIT_IN_INIT_OUTPUT.
const AVStreamInitInInitOutput = C.AVSTREAM_INIT_IN_INIT_OUTPUT

// AVFrameFilenameFlagsMultiple wraps AV_FRAME_FILENAME_FLAGS_MULTIPLE.
const AVFrameFilenameFlagsMultiple = C.AV_FRAME_FILENAME_FLAGS_MULTIPLE

// AVIOSeekableNormal wraps AVIO_SEEKABLE_NORMAL.
const AVIOSeekableNormal = C.AVIO_SEEKABLE_NORMAL

// AVIOSeekableTime wraps AVIO_SEEKABLE_TIME.
const AVIOSeekableTime = C.AVIO_SEEKABLE_TIME

// AVSeekSize wraps AVSEEK_SIZE.
const AVSeekSize = C.AVSEEK_SIZE

// AVSeekForce wraps AVSEEK_FORCE.
const AVSeekForce = C.AVSEEK_FORCE

// AVIOFlagRead wraps AVIO_FLAG_READ.
const AVIOFlagRead = C.AVIO_FLAG_READ

// AVIOFlagWrite wraps AVIO_FLAG_WRITE.
const AVIOFlagWrite = C.AVIO_FLAG_WRITE

// AVIOFlagReadWrite wraps AVIO_FLAG_READ_WRITE.
const AVIOFlagReadWrite = C.AVIO_FLAG_READ_WRITE

// AVIOFlagNonblock wraps AVIO_FLAG_NONBLOCK.
const AVIOFlagNonblock = C.AVIO_FLAG_NONBLOCK

// AVIOFlagDirect wraps AVIO_FLAG_DIRECT.
const AVIOFlagDirect = C.AVIO_FLAG_DIRECT

// LIBAVFormatVersionMinor wraps LIBAVFORMAT_VERSION_MINOR.
const LIBAVFormatVersionMinor = C.LIBAVFORMAT_VERSION_MINOR

// LIBAVFormatVersionMicro wraps LIBAVFORMAT_VERSION_MICRO.
const LIBAVFormatVersionMicro = C.LIBAVFORMAT_VERSION_MICRO

// LIBAVFormatVersionInt wraps LIBAVFORMAT_VERSION_INT.
const LIBAVFormatVersionInt = C.LIBAVFORMAT_VERSION_INT

// LIBAVFormatBuild wraps LIBAVFORMAT_BUILD.
const LIBAVFormatBuild = C.LIBAVFORMAT_BUILD

// LIBAVFormatVersionMajor wraps LIBAVFORMAT_VERSION_MAJOR.
const LIBAVFormatVersionMajor = C.LIBAVFORMAT_VERSION_MAJOR

// FFAPIComputePktFields2 wraps FF_API_COMPUTE_PKT_FIELDS2.
const FFAPIComputePktFields2 = C.FF_API_COMPUTE_PKT_FIELDS2

// FFAPIGetEndPts wraps FF_API_GET_END_PTS.
const FFAPIGetEndPts = C.FF_API_GET_END_PTS

// FFAPIAVIODircontext wraps FF_API_AVIODIRCONTEXT.
const FFAPIAVIODircontext = C.FF_API_AVIODIRCONTEXT

// FFAPIAVFormatIOClose wraps FF_API_AVFORMAT_IO_CLOSE.
const FFAPIAVFormatIOClose = C.FF_API_AVFORMAT_IO_CLOSE

// FFAPIAVIOWriteNonconst wraps FF_API_AVIO_WRITE_NONCONST.
const FFAPIAVIOWriteNonconst = C.FF_API_AVIO_WRITE_NONCONST

// FFAPILavfShortest wraps FF_API_LAVF_SHORTEST.
const FFAPILavfShortest = C.FF_API_LAVF_SHORTEST

// FFAPIAllowFlush wraps FF_API_ALLOW_FLUSH.
const FFAPIAllowFlush = C.FF_API_ALLOW_FLUSH

// FFAPIAVStreamSideData wraps FF_API_AVSTREAM_SIDE_DATA.
const FFAPIAVStreamSideData = C.FF_API_AVSTREAM_SIDE_DATA

// FFAPIRFrameRate wraps FF_API_R_FRAME_RATE.
const FFAPIRFrameRate = C.FF_API_R_FRAME_RATE

// FFLambdaShift wraps FF_LAMBDA_SHIFT.
const FFLambdaShift = C.FF_LAMBDA_SHIFT

// FFLambdaScale wraps FF_LAMBDA_SCALE.
const FFLambdaScale = C.FF_LAMBDA_SCALE

// FFQp2Lambda wraps FF_QP2LAMBDA.
const FFQp2Lambda = C.FF_QP2LAMBDA

// FFLambdaMax wraps FF_LAMBDA_MAX.
const FFLambdaMax = C.FF_LAMBDA_MAX

// FFQualityScale wraps FF_QUALITY_SCALE.
const FFQualityScale = C.FF_QUALITY_SCALE

// AVNoptsValue wraps AV_NOPTS_VALUE.
const AVNoptsValue = C.AV_NOPTS_VALUE

// AVTimeBase wraps AV_TIME_BASE.
const AVTimeBase = C.AV_TIME_BASE

// AVFourccMaxStringSize wraps AV_FOURCC_MAX_STRING_SIZE.
const AVFourccMaxStringSize = C.AV_FOURCC_MAX_STRING_SIZE

// AVBufferFlagReadonly wraps AV_BUFFER_FLAG_READONLY.
const AVBufferFlagReadonly = C.AV_BUFFER_FLAG_READONLY

// AVChFrontLeft wraps AV_CH_FRONT_LEFT.
const AVChFrontLeft = C.AV_CH_FRONT_LEFT

// AVChFrontRight wraps AV_CH_FRONT_RIGHT.
const AVChFrontRight = C.AV_CH_FRONT_RIGHT

// AVChFrontCenter wraps AV_CH_FRONT_CENTER.
const AVChFrontCenter = C.AV_CH_FRONT_CENTER

// AVChLowFrequency wraps AV_CH_LOW_FREQUENCY.
const AVChLowFrequency = C.AV_CH_LOW_FREQUENCY

// AVChBackLeft wraps AV_CH_BACK_LEFT.
const AVChBackLeft = C.AV_CH_BACK_LEFT

// AVChBackRight wraps AV_CH_BACK_RIGHT.
const AVChBackRight = C.AV_CH_BACK_RIGHT

// AVChFrontLeftOfCenter wraps AV_CH_FRONT_LEFT_OF_CENTER.
const AVChFrontLeftOfCenter = C.AV_CH_FRONT_LEFT_OF_CENTER

// AVChFrontRightOfCenter wraps AV_CH_FRONT_RIGHT_OF_CENTER.
const AVChFrontRightOfCenter = C.AV_CH_FRONT_RIGHT_OF_CENTER

// AVChBackCenter wraps AV_CH_BACK_CENTER.
const AVChBackCenter = C.AV_CH_BACK_CENTER

// AVChSideLeft wraps AV_CH_SIDE_LEFT.
const AVChSideLeft = C.AV_CH_SIDE_LEFT

// AVChSideRight wraps AV_CH_SIDE_RIGHT.
const AVChSideRight = C.AV_CH_SIDE_RIGHT

// AVChTopCenter wraps AV_CH_TOP_CENTER.
const AVChTopCenter = C.AV_CH_TOP_CENTER

// AVChTopFrontLeft wraps AV_CH_TOP_FRONT_LEFT.
const AVChTopFrontLeft = C.AV_CH_TOP_FRONT_LEFT

// AVChTopFrontCenter wraps AV_CH_TOP_FRONT_CENTER.
const AVChTopFrontCenter = C.AV_CH_TOP_FRONT_CENTER

// AVChTopFrontRight wraps AV_CH_TOP_FRONT_RIGHT.
const AVChTopFrontRight = C.AV_CH_TOP_FRONT_RIGHT

// AVChTopBackLeft wraps AV_CH_TOP_BACK_LEFT.
const AVChTopBackLeft = C.AV_CH_TOP_BACK_LEFT

// AVChTopBackCenter wraps AV_CH_TOP_BACK_CENTER.
const AVChTopBackCenter = C.AV_CH_TOP_BACK_CENTER

// AVChTopBackRight wraps AV_CH_TOP_BACK_RIGHT.
const AVChTopBackRight = C.AV_CH_TOP_BACK_RIGHT

// AVChStereoLeft wraps AV_CH_STEREO_LEFT.
const AVChStereoLeft = C.AV_CH_STEREO_LEFT

// AVChStereoRight wraps AV_CH_STEREO_RIGHT.
const AVChStereoRight = C.AV_CH_STEREO_RIGHT

// AVChWideLeft wraps AV_CH_WIDE_LEFT.
const AVChWideLeft = C.AV_CH_WIDE_LEFT

// AVChWideRight wraps AV_CH_WIDE_RIGHT.
const AVChWideRight = C.AV_CH_WIDE_RIGHT

// AVChSurroundDirectLeft wraps AV_CH_SURROUND_DIRECT_LEFT.
const AVChSurroundDirectLeft = C.AV_CH_SURROUND_DIRECT_LEFT

// AVChSurroundDirectRight wraps AV_CH_SURROUND_DIRECT_RIGHT.
const AVChSurroundDirectRight = C.AV_CH_SURROUND_DIRECT_RIGHT

// AVChLowFrequency2 wraps AV_CH_LOW_FREQUENCY_2.
const AVChLowFrequency2 = C.AV_CH_LOW_FREQUENCY_2

// AVChTopSideLeft wraps AV_CH_TOP_SIDE_LEFT.
const AVChTopSideLeft = C.AV_CH_TOP_SIDE_LEFT

// AVChTopSideRight wraps AV_CH_TOP_SIDE_RIGHT.
const AVChTopSideRight = C.AV_CH_TOP_SIDE_RIGHT

// AVChBottomFrontCenter wraps AV_CH_BOTTOM_FRONT_CENTER.
const AVChBottomFrontCenter = C.AV_CH_BOTTOM_FRONT_CENTER

// AVChBottomFrontLeft wraps AV_CH_BOTTOM_FRONT_LEFT.
const AVChBottomFrontLeft = C.AV_CH_BOTTOM_FRONT_LEFT

// AVChBottomFrontRight wraps AV_CH_BOTTOM_FRONT_RIGHT.
const AVChBottomFrontRight = C.AV_CH_BOTTOM_FRONT_RIGHT

// AVChLayoutMono wraps AV_CH_LAYOUT_MONO.
const AVChLayoutMono = C.AV_CH_LAYOUT_MONO

// AVChLayoutStereo wraps AV_CH_LAYOUT_STEREO.
const AVChLayoutStereo = C.AV_CH_LAYOUT_STEREO

// AVChLayout2Point1 wraps AV_CH_LAYOUT_2POINT1.
const AVChLayout2Point1 = C.AV_CH_LAYOUT_2POINT1

// AVChLayout21 wraps AV_CH_LAYOUT_2_1.
const AVChLayout21 = C.AV_CH_LAYOUT_2_1

// AVChLayoutSurround wraps AV_CH_LAYOUT_SURROUND.
const AVChLayoutSurround = C.AV_CH_LAYOUT_SURROUND

// AVChLayout3Point1 wraps AV_CH_LAYOUT_3POINT1.
const AVChLayout3Point1 = C.AV_CH_LAYOUT_3POINT1

// AVChLayout4Point0 wraps AV_CH_LAYOUT_4POINT0.
const AVChLayout4Point0 = C.AV_CH_LAYOUT_4POINT0

// AVChLayout4Point1 wraps AV_CH_LAYOUT_4POINT1.
const AVChLayout4Point1 = C.AV_CH_LAYOUT_4POINT1

// AVChLayout22 wraps AV_CH_LAYOUT_2_2.
const AVChLayout22 = C.AV_CH_LAYOUT_2_2

// AVChLayoutQuad wraps AV_CH_LAYOUT_QUAD.
const AVChLayoutQuad = C.AV_CH_LAYOUT_QUAD

// AVChLayout5Point0 wraps AV_CH_LAYOUT_5POINT0.
const AVChLayout5Point0 = C.AV_CH_LAYOUT_5POINT0

// AVChLayout5Point1 wraps AV_CH_LAYOUT_5POINT1.
const AVChLayout5Point1 = C.AV_CH_LAYOUT_5POINT1

// AVChLayout5Point0Back wraps AV_CH_LAYOUT_5POINT0_BACK.
const AVChLayout5Point0Back = C.AV_CH_LAYOUT_5POINT0_BACK

// AVChLayout5Point1Back wraps AV_CH_LAYOUT_5POINT1_BACK.
const AVChLayout5Point1Back = C.AV_CH_LAYOUT_5POINT1_BACK

// AVChLayout6Point0 wraps AV_CH_LAYOUT_6POINT0.
const AVChLayout6Point0 = C.AV_CH_LAYOUT_6POINT0

// AVChLayout6Point0Front wraps AV_CH_LAYOUT_6POINT0_FRONT.
const AVChLayout6Point0Front = C.AV_CH_LAYOUT_6POINT0_FRONT

// AVChLayoutHexagonal wraps AV_CH_LAYOUT_HEXAGONAL.
const AVChLayoutHexagonal = C.AV_CH_LAYOUT_HEXAGONAL

// AVChLayout3Point1Point2 wraps AV_CH_LAYOUT_3POINT1POINT2.
const AVChLayout3Point1Point2 = C.AV_CH_LAYOUT_3POINT1POINT2

// AVChLayout6Point1 wraps AV_CH_LAYOUT_6POINT1.
const AVChLayout6Point1 = C.AV_CH_LAYOUT_6POINT1

// AVChLayout6Point1Back wraps AV_CH_LAYOUT_6POINT1_BACK.
const AVChLayout6Point1Back = C.AV_CH_LAYOUT_6POINT1_BACK

// AVChLayout6Point1Front wraps AV_CH_LAYOUT_6POINT1_FRONT.
const AVChLayout6Point1Front = C.AV_CH_LAYOUT_6POINT1_FRONT

// AVChLayout7Point0 wraps AV_CH_LAYOUT_7POINT0.
const AVChLayout7Point0 = C.AV_CH_LAYOUT_7POINT0

// AVChLayout7Point0Front wraps AV_CH_LAYOUT_7POINT0_FRONT.
const AVChLayout7Point0Front = C.AV_CH_LAYOUT_7POINT0_FRONT

// AVChLayout7Point1 wraps AV_CH_LAYOUT_7POINT1.
const AVChLayout7Point1 = C.AV_CH_LAYOUT_7POINT1

// AVChLayout7Point1Wide wraps AV_CH_LAYOUT_7POINT1_WIDE.
const AVChLayout7Point1Wide = C.AV_CH_LAYOUT_7POINT1_WIDE

// AVChLayout7Point1WideBack wraps AV_CH_LAYOUT_7POINT1_WIDE_BACK.
const AVChLayout7Point1WideBack = C.AV_CH_LAYOUT_7POINT1_WIDE_BACK

// AVChLayout5Point1Point2Back wraps AV_CH_LAYOUT_5POINT1POINT2_BACK.
const AVChLayout5Point1Point2Back = C.AV_CH_LAYOUT_5POINT1POINT2_BACK

// AVChLayoutOctagonal wraps AV_CH_LAYOUT_OCTAGONAL.
const AVChLayoutOctagonal = C.AV_CH_LAYOUT_OCTAGONAL

// AVChLayoutCube wraps AV_CH_LAYOUT_CUBE.
const AVChLayoutCube = C.AV_CH_LAYOUT_CUBE

// AVChLayout5Point1Point4Back wraps AV_CH_LAYOUT_5POINT1POINT4_BACK.
const AVChLayout5Point1Point4Back = C.AV_CH_LAYOUT_5POINT1POINT4_BACK

// AVChLayout7Point1Point2 wraps AV_CH_LAYOUT_7POINT1POINT2.
const AVChLayout7Point1Point2 = C.AV_CH_LAYOUT_7POINT1POINT2

// AVChLayout7Point1Point4Back wraps AV_CH_LAYOUT_7POINT1POINT4_BACK.
const AVChLayout7Point1Point4Back = C.AV_CH_LAYOUT_7POINT1POINT4_BACK

// AVChLayoutHexadecagonal wraps AV_CH_LAYOUT_HEXADECAGONAL.
const AVChLayoutHexadecagonal = C.AV_CH_LAYOUT_HEXADECAGONAL

// AVChLayoutStereoDownmix wraps AV_CH_LAYOUT_STEREO_DOWNMIX.
const AVChLayoutStereoDownmix = C.AV_CH_LAYOUT_STEREO_DOWNMIX

// AVChLayout22Point2 wraps AV_CH_LAYOUT_22POINT2.
const AVChLayout22Point2 = C.AV_CH_LAYOUT_22POINT2

// AVChLayout7Point1TopBack wraps AV_CH_LAYOUT_7POINT1_TOP_BACK.
const AVChLayout7Point1TopBack = C.AV_CH_LAYOUT_7POINT1_TOP_BACK

// AVDictMatchCase wraps AV_DICT_MATCH_CASE.
const AVDictMatchCase = C.AV_DICT_MATCH_CASE

// AVDictIgnoreSuffix wraps AV_DICT_IGNORE_SUFFIX.
const AVDictIgnoreSuffix = C.AV_DICT_IGNORE_SUFFIX

// AVDictDontStrdupKey wraps AV_DICT_DONT_STRDUP_KEY.
const AVDictDontStrdupKey = C.AV_DICT_DONT_STRDUP_KEY

// AVDictDontStrdupVal wraps AV_DICT_DONT_STRDUP_VAL.
const AVDictDontStrdupVal = C.AV_DICT_DONT_STRDUP_VAL

// AVDictDontOverwrite wraps AV_DICT_DONT_OVERWRITE.
const AVDictDontOverwrite = C.AV_DICT_DONT_OVERWRITE

// AVDictAppend wraps AV_DICT_APPEND.
const AVDictAppend = C.AV_DICT_APPEND

// AVDictMultikey wraps AV_DICT_MULTIKEY.
const AVDictMultikey = C.AV_DICT_MULTIKEY

// AVErrorBsfNotFoundConst wraps AVERROR_BSF_NOT_FOUND.
const AVErrorBsfNotFoundConst = C.AVERROR_BSF_NOT_FOUND

// AVErrorBugConst wraps AVERROR_BUG.
const AVErrorBugConst = C.AVERROR_BUG

// AVErrorBufferTooSmallConst wraps AVERROR_BUFFER_TOO_SMALL.
const AVErrorBufferTooSmallConst = C.AVERROR_BUFFER_TOO_SMALL

// AVErrorDecoderNotFoundConst wraps AVERROR_DECODER_NOT_FOUND.
const AVErrorDecoderNotFoundConst = C.AVERROR_DECODER_NOT_FOUND

// AVErrorDemuxerNotFoundConst wraps AVERROR_DEMUXER_NOT_FOUND.
const AVErrorDemuxerNotFoundConst = C.AVERROR_DEMUXER_NOT_FOUND

// AVErrorEncoderNotFoundConst wraps AVERROR_ENCODER_NOT_FOUND.
const AVErrorEncoderNotFoundConst = C.AVERROR_ENCODER_NOT_FOUND

// AVErrorEofConst wraps AVERROR_EOF.
const AVErrorEofConst = C.AVERROR_EOF

// AVErrorExitConst wraps AVERROR_EXIT.
const AVErrorExitConst = C.AVERROR_EXIT

// AVErrorExternalConst wraps AVERROR_EXTERNAL.
const AVErrorExternalConst = C.AVERROR_EXTERNAL

// AVErrorFilterNotFoundConst wraps AVERROR_FILTER_NOT_FOUND.
const AVErrorFilterNotFoundConst = C.AVERROR_FILTER_NOT_FOUND

// AVErrorInvaliddataConst wraps AVERROR_INVALIDDATA.
const AVErrorInvaliddataConst = C.AVERROR_INVALIDDATA

// AVErrorMuxerNotFoundConst wraps AVERROR_MUXER_NOT_FOUND.
const AVErrorMuxerNotFoundConst = C.AVERROR_MUXER_NOT_FOUND

// AVErrorOptionNotFoundConst wraps AVERROR_OPTION_NOT_FOUND.
const AVErrorOptionNotFoundConst = C.AVERROR_OPTION_NOT_FOUND

// AVErrorPatchwelcomeConst wraps AVERROR_PATCHWELCOME.
const AVErrorPatchwelcomeConst = C.AVERROR_PATCHWELCOME

// AVErrorProtocolNotFoundConst wraps AVERROR_PROTOCOL_NOT_FOUND.
const AVErrorProtocolNotFoundConst = C.AVERROR_PROTOCOL_NOT_FOUND

// AVErrorStreamNotFoundConst wraps AVERROR_STREAM_NOT_FOUND.
const AVErrorStreamNotFoundConst = C.AVERROR_STREAM_NOT_FOUND

// AVErrorBug2Const wraps AVERROR_BUG2.
const AVErrorBug2Const = C.AVERROR_BUG2

// AVErrorUnknownConst wraps AVERROR_UNKNOWN.
const AVErrorUnknownConst = C.AVERROR_UNKNOWN

// AVErrorExperimentalConst wraps AVERROR_EXPERIMENTAL.
const AVErrorExperimentalConst = C.AVERROR_EXPERIMENTAL

// AVErrorInputChangedConst wraps AVERROR_INPUT_CHANGED.
const AVErrorInputChangedConst = C.AVERROR_INPUT_CHANGED

// AVErrorOutputChangedConst wraps AVERROR_OUTPUT_CHANGED.
const AVErrorOutputChangedConst = C.AVERROR_OUTPUT_CHANGED

// AVErrorHttpBadRequestConst wraps AVERROR_HTTP_BAD_REQUEST.
const AVErrorHttpBadRequestConst = C.AVERROR_HTTP_BAD_REQUEST

// AVErrorHttpUnauthorizedConst wraps AVERROR_HTTP_UNAUTHORIZED.
const AVErrorHttpUnauthorizedConst = C.AVERROR_HTTP_UNAUTHORIZED

// AVErrorHttpForbiddenConst wraps AVERROR_HTTP_FORBIDDEN.
const AVErrorHttpForbiddenConst = C.AVERROR_HTTP_FORBIDDEN

// AVErrorHttpNotFoundConst wraps AVERROR_HTTP_NOT_FOUND.
const AVErrorHttpNotFoundConst = C.AVERROR_HTTP_NOT_FOUND

// AVErrorHttpOther4XxConst wraps AVERROR_HTTP_OTHER_4XX.
const AVErrorHttpOther4XxConst = C.AVERROR_HTTP_OTHER_4XX

// AVErrorHttpServerErrorConst wraps AVERROR_HTTP_SERVER_ERROR.
const AVErrorHttpServerErrorConst = C.AVERROR_HTTP_SERVER_ERROR

// AVErrorMaxStringSize wraps AV_ERROR_MAX_STRING_SIZE.
const AVErrorMaxStringSize = C.AV_ERROR_MAX_STRING_SIZE

// AVNumDataPointers wraps AV_NUM_DATA_POINTERS.
const AVNumDataPointers = C.AV_NUM_DATA_POINTERS

// AVFrameFlagCorrupt wraps AV_FRAME_FLAG_CORRUPT.
const AVFrameFlagCorrupt = C.AV_FRAME_FLAG_CORRUPT

// AVFrameFlagKey wraps AV_FRAME_FLAG_KEY.
const AVFrameFlagKey = C.AV_FRAME_FLAG_KEY

// AVFrameFlagDiscard wraps AV_FRAME_FLAG_DISCARD.
const AVFrameFlagDiscard = C.AV_FRAME_FLAG_DISCARD

// AVFrameFlagInterlaced wraps AV_FRAME_FLAG_INTERLACED.
const AVFrameFlagInterlaced = C.AV_FRAME_FLAG_INTERLACED

// AVFrameFlagTopFieldFirst wraps AV_FRAME_FLAG_TOP_FIELD_FIRST.
const AVFrameFlagTopFieldFirst = C.AV_FRAME_FLAG_TOP_FIELD_FIRST

// FFDecodeErrorInvalidBitstream wraps FF_DECODE_ERROR_INVALID_BITSTREAM.
const FFDecodeErrorInvalidBitstream = C.FF_DECODE_ERROR_INVALID_BITSTREAM

// FFDecodeErrorMissingReference wraps FF_DECODE_ERROR_MISSING_REFERENCE.
const FFDecodeErrorMissingReference = C.FF_DECODE_ERROR_MISSING_REFERENCE

// FFDecodeErrorConcealmentActive wraps FF_DECODE_ERROR_CONCEALMENT_ACTIVE.
const FFDecodeErrorConcealmentActive = C.FF_DECODE_ERROR_CONCEALMENT_ACTIVE

// FFDecodeErrorDecodeSlices wraps FF_DECODE_ERROR_DECODE_SLICES.
const FFDecodeErrorDecodeSlices = C.FF_DECODE_ERROR_DECODE_SLICES

// AVFrameCropUnaligned wraps AV_FRAME_CROP_UNALIGNED.
const AVFrameCropUnaligned = C.AV_FRAME_CROP_UNALIGNED

// AVHWFrameMapRead wraps AV_HWFRAME_MAP_READ.
const AVHWFrameMapRead = C.AV_HWFRAME_MAP_READ

// AVHWFrameMapWrite wraps AV_HWFRAME_MAP_WRITE.
const AVHWFrameMapWrite = C.AV_HWFRAME_MAP_WRITE

// AVHWFrameMapOverwrite wraps AV_HWFRAME_MAP_OVERWRITE.
const AVHWFrameMapOverwrite = C.AV_HWFRAME_MAP_OVERWRITE

// AVHWFrameMapDirect wraps AV_HWFRAME_MAP_DIRECT.
const AVHWFrameMapDirect = C.AV_HWFRAME_MAP_DIRECT

// AVLogQuiet wraps AV_LOG_QUIET.
const AVLogQuiet = C.AV_LOG_QUIET

// AVLogPanic wraps AV_LOG_PANIC.
const AVLogPanic = C.AV_LOG_PANIC

// AVLogFatal wraps AV_LOG_FATAL.
const AVLogFatal = C.AV_LOG_FATAL

// AVLogError wraps AV_LOG_ERROR.
const AVLogError = C.AV_LOG_ERROR

// AVLogWarning wraps AV_LOG_WARNING.
const AVLogWarning = C.AV_LOG_WARNING

// AVLogInfo wraps AV_LOG_INFO.
const AVLogInfo = C.AV_LOG_INFO

// AVLogVerbose wraps AV_LOG_VERBOSE.
const AVLogVerbose = C.AV_LOG_VERBOSE

// AVLogDebug wraps AV_LOG_DEBUG.
const AVLogDebug = C.AV_LOG_DEBUG

// AVLogTrace wraps AV_LOG_TRACE.
const AVLogTrace = C.AV_LOG_TRACE

// AVLogMaxOffset wraps AV_LOG_MAX_OFFSET.
const AVLogMaxOffset = C.AV_LOG_MAX_OFFSET

// AVLogSkipRepeated wraps AV_LOG_SKIP_REPEATED.
const AVLogSkipRepeated = C.AV_LOG_SKIP_REPEATED

// AVLogPrintLevel wraps AV_LOG_PRINT_LEVEL.
const AVLogPrintLevel = C.AV_LOG_PRINT_LEVEL

// MEf wraps M_Ef.
const MEf = C.M_Ef

// MLn2F wraps M_LN2f.
const MLn2F = C.M_LN2f

// MLn10F wraps M_LN10f.
const MLn10F = C.M_LN10f

// MLog210 wraps M_LOG2_10.
const MLog210 = C.M_LOG2_10

// MLog210F wraps M_LOG2_10f.
const MLog210F = C.M_LOG2_10f

// MPhi wraps M_PHI.
const MPhi = C.M_PHI

// MPhif wraps M_PHIf.
const MPhif = C.M_PHIf

// MPif wraps M_PIf.
const MPif = C.M_PIf

// MPi2F wraps M_PI_2f.
const MPi2F = C.M_PI_2f

// MPi4F wraps M_PI_4f.
const MPi4F = C.M_PI_4f

// M1Pif wraps M_1_PIf.
const M1Pif = C.M_1_PIf

// M2Pif wraps M_2_PIf.
const M2Pif = C.M_2_PIf

// M2Sqrtpif wraps M_2_SQRTPIf.
const M2Sqrtpif = C.M_2_SQRTPIf

// MSqrt12F wraps M_SQRT1_2f.
const MSqrt12F = C.M_SQRT1_2f

// MSqrt2F wraps M_SQRT2f.
const MSqrt2F = C.M_SQRT2f

// AVOptFlagEncodingParam wraps AV_OPT_FLAG_ENCODING_PARAM.
const AVOptFlagEncodingParam = C.AV_OPT_FLAG_ENCODING_PARAM

// AVOptFlagDecodingParam wraps AV_OPT_FLAG_DECODING_PARAM.
const AVOptFlagDecodingParam = C.AV_OPT_FLAG_DECODING_PARAM

// AVOptFlagAudioParam wraps AV_OPT_FLAG_AUDIO_PARAM.
const AVOptFlagAudioParam = C.AV_OPT_FLAG_AUDIO_PARAM

// AVOptFlagVideoParam wraps AV_OPT_FLAG_VIDEO_PARAM.
const AVOptFlagVideoParam = C.AV_OPT_FLAG_VIDEO_PARAM

// AVOptFlagSubtitleParam wraps AV_OPT_FLAG_SUBTITLE_PARAM.
const AVOptFlagSubtitleParam = C.AV_OPT_FLAG_SUBTITLE_PARAM

// AVOptFlagExport wraps AV_OPT_FLAG_EXPORT.
const AVOptFlagExport = C.AV_OPT_FLAG_EXPORT

// AVOptFlagReadonly wraps AV_OPT_FLAG_READONLY.
const AVOptFlagReadonly = C.AV_OPT_FLAG_READONLY

// AVOptFlagBsfParam wraps AV_OPT_FLAG_BSF_PARAM.
const AVOptFlagBsfParam = C.AV_OPT_FLAG_BSF_PARAM

// AVOptFlagRuntimeParam wraps AV_OPT_FLAG_RUNTIME_PARAM.
const AVOptFlagRuntimeParam = C.AV_OPT_FLAG_RUNTIME_PARAM

// AVOptFlagFilteringParam wraps AV_OPT_FLAG_FILTERING_PARAM.
const AVOptFlagFilteringParam = C.AV_OPT_FLAG_FILTERING_PARAM

// AVOptFlagDeprecated wraps AV_OPT_FLAG_DEPRECATED.
const AVOptFlagDeprecated = C.AV_OPT_FLAG_DEPRECATED

// AVOptFlagChildConsts wraps AV_OPT_FLAG_CHILD_CONSTS.
const AVOptFlagChildConsts = C.AV_OPT_FLAG_CHILD_CONSTS

// AVOptSearchChildren wraps AV_OPT_SEARCH_CHILDREN.
const AVOptSearchChildren = C.AV_OPT_SEARCH_CHILDREN

// AVOptSearchFakeObj wraps AV_OPT_SEARCH_FAKE_OBJ.
const AVOptSearchFakeObj = C.AV_OPT_SEARCH_FAKE_OBJ

// AVOptAllowNull wraps AV_OPT_ALLOW_NULL.
const AVOptAllowNull = C.AV_OPT_ALLOW_NULL

// AVOptMultiComponentRange wraps AV_OPT_MULTI_COMPONENT_RANGE.
const AVOptMultiComponentRange = C.AV_OPT_MULTI_COMPONENT_RANGE

// AVOptSerializeSkipDefaults wraps AV_OPT_SERIALIZE_SKIP_DEFAULTS.
const AVOptSerializeSkipDefaults = C.AV_OPT_SERIALIZE_SKIP_DEFAULTS

// AVOptSerializeOptFlagsExact wraps AV_OPT_SERIALIZE_OPT_FLAGS_EXACT.
const AVOptSerializeOptFlagsExact = C.AV_OPT_SERIALIZE_OPT_FLAGS_EXACT

// AVOptFlagImplicitKey wraps AV_OPT_FLAG_IMPLICIT_KEY.
const AVOptFlagImplicitKey = C.AV_OPT_FLAG_IMPLICIT_KEY

// AVPaletteSize wraps AVPALETTE_SIZE.
const AVPaletteSize = C.AVPALETTE_SIZE

// AVPaletteCount wraps AVPALETTE_COUNT.
const AVPaletteCount = C.AVPALETTE_COUNT

// AVPixFmtRgb32 wraps AV_PIX_FMT_RGB32.
const AVPixFmtRgb32 = C.AV_PIX_FMT_RGB32

// AVPixFmtRgb321 wraps AV_PIX_FMT_RGB32_1.
const AVPixFmtRgb321 = C.AV_PIX_FMT_RGB32_1

// AVPixFmtBgr32 wraps AV_PIX_FMT_BGR32.
const AVPixFmtBgr32 = C.AV_PIX_FMT_BGR32

// AVPixFmtBgr321 wraps AV_PIX_FMT_BGR32_1.
const AVPixFmtBgr321 = C.AV_PIX_FMT_BGR32_1

// AVPixFmt0Rgb32 wraps AV_PIX_FMT_0RGB32.
const AVPixFmt0Rgb32 = C.AV_PIX_FMT_0RGB32

// AVPixFmt0Bgr32 wraps AV_PIX_FMT_0BGR32.
const AVPixFmt0Bgr32 = C.AV_PIX_FMT_0BGR32

// AVPixFmtGray9 wraps AV_PIX_FMT_GRAY9.
const AVPixFmtGray9 = C.AV_PIX_FMT_GRAY9

// AVPixFmtGray10 wraps AV_PIX_FMT_GRAY10.
const AVPixFmtGray10 = C.AV_PIX_FMT_GRAY10

// AVPixFmtGray12 wraps AV_PIX_FMT_GRAY12.
const AVPixFmtGray12 = C.AV_PIX_FMT_GRAY12

// AVPixFmtGray14 wraps AV_PIX_FMT_GRAY14.
const AVPixFmtGray14 = C.AV_PIX_FMT_GRAY14

// AVPixFmtGray16 wraps AV_PIX_FMT_GRAY16.
const AVPixFmtGray16 = C.AV_PIX_FMT_GRAY16

// AVPixFmtYa16 wraps AV_PIX_FMT_YA16.
const AVPixFmtYa16 = C.AV_PIX_FMT_YA16

// AVPixFmtRgb48 wraps AV_PIX_FMT_RGB48.
const AVPixFmtRgb48 = C.AV_PIX_FMT_RGB48

// AVPixFmtRgb565 wraps AV_PIX_FMT_RGB565.
const AVPixFmtRgb565 = C.AV_PIX_FMT_RGB565

// AVPixFmtRgb555 wraps AV_PIX_FMT_RGB555.
const AVPixFmtRgb555 = C.AV_PIX_FMT_RGB555

// AVPixFmtRgb444 wraps AV_PIX_FMT_RGB444.
const AVPixFmtRgb444 = C.AV_PIX_FMT_RGB444

// AVPixFmtRgba64 wraps AV_PIX_FMT_RGBA64.
const AVPixFmtRgba64 = C.AV_PIX_FMT_RGBA64

// AVPixFmtBgr48 wraps AV_PIX_FMT_BGR48.
const AVPixFmtBgr48 = C.AV_PIX_FMT_BGR48

// AVPixFmtBgr565 wraps AV_PIX_FMT_BGR565.
const AVPixFmtBgr565 = C.AV_PIX_FMT_BGR565

// AVPixFmtBgr555 wraps AV_PIX_FMT_BGR555.
const AVPixFmtBgr555 = C.AV_PIX_FMT_BGR555

// AVPixFmtBgr444 wraps AV_PIX_FMT_BGR444.
const AVPixFmtBgr444 = C.AV_PIX_FMT_BGR444

// AVPixFmtBgra64 wraps AV_PIX_FMT_BGRA64.
const AVPixFmtBgra64 = C.AV_PIX_FMT_BGRA64

// AVPixFmtYuv420P9 wraps AV_PIX_FMT_YUV420P9.
const AVPixFmtYuv420P9 = C.AV_PIX_FMT_YUV420P9

// AVPixFmtYuv422P9 wraps AV_PIX_FMT_YUV422P9.
const AVPixFmtYuv422P9 = C.AV_PIX_FMT_YUV422P9

// AVPixFmtYuv444P9 wraps AV_PIX_FMT_YUV444P9.
const AVPixFmtYuv444P9 = C.AV_PIX_FMT_YUV444P9

// AVPixFmtYuv420P10 wraps AV_PIX_FMT_YUV420P10.
const AVPixFmtYuv420P10 = C.AV_PIX_FMT_YUV420P10

// AVPixFmtYuv422P10 wraps AV_PIX_FMT_YUV422P10.
const AVPixFmtYuv422P10 = C.AV_PIX_FMT_YUV422P10

// AVPixFmtYuv440P10 wraps AV_PIX_FMT_YUV440P10.
const AVPixFmtYuv440P10 = C.AV_PIX_FMT_YUV440P10

// AVPixFmtYuv444P10 wraps AV_PIX_FMT_YUV444P10.
const AVPixFmtYuv444P10 = C.AV_PIX_FMT_YUV444P10

// AVPixFmtYuv420P12 wraps AV_PIX_FMT_YUV420P12.
const AVPixFmtYuv420P12 = C.AV_PIX_FMT_YUV420P12

// AVPixFmtYuv422P12 wraps AV_PIX_FMT_YUV422P12.
const AVPixFmtYuv422P12 = C.AV_PIX_FMT_YUV422P12

// AVPixFmtYuv440P12 wraps AV_PIX_FMT_YUV440P12.
const AVPixFmtYuv440P12 = C.AV_PIX_FMT_YUV440P12

// AVPixFmtYuv444P12 wraps AV_PIX_FMT_YUV444P12.
const AVPixFmtYuv444P12 = C.AV_PIX_FMT_YUV444P12

// AVPixFmtYuv420P14 wraps AV_PIX_FMT_YUV420P14.
const AVPixFmtYuv420P14 = C.AV_PIX_FMT_YUV420P14

// AVPixFmtYuv422P14 wraps AV_PIX_FMT_YUV422P14.
const AVPixFmtYuv422P14 = C.AV_PIX_FMT_YUV422P14

// AVPixFmtYuv444P14 wraps AV_PIX_FMT_YUV444P14.
const AVPixFmtYuv444P14 = C.AV_PIX_FMT_YUV444P14

// AVPixFmtYuv420P16 wraps AV_PIX_FMT_YUV420P16.
const AVPixFmtYuv420P16 = C.AV_PIX_FMT_YUV420P16

// AVPixFmtYuv422P16 wraps AV_PIX_FMT_YUV422P16.
const AVPixFmtYuv422P16 = C.AV_PIX_FMT_YUV422P16

// AVPixFmtYuv444P16 wraps AV_PIX_FMT_YUV444P16.
const AVPixFmtYuv444P16 = C.AV_PIX_FMT_YUV444P16

// AVPixFmtGbrp9 wraps AV_PIX_FMT_GBRP9.
const AVPixFmtGbrp9 = C.AV_PIX_FMT_GBRP9

// AVPixFmtGbrp10 wraps AV_PIX_FMT_GBRP10.
const AVPixFmtGbrp10 = C.AV_PIX_FMT_GBRP10

// AVPixFmtGbrp12 wraps AV_PIX_FMT_GBRP12.
const AVPixFmtGbrp12 = C.AV_PIX_FMT_GBRP12

// AVPixFmtGbrp14 wraps AV_PIX_FMT_GBRP14.
const AVPixFmtGbrp14 = C.AV_PIX_FMT_GBRP14

// AVPixFmtGbrp16 wraps AV_PIX_FMT_GBRP16.
const AVPixFmtGbrp16 = C.AV_PIX_FMT_GBRP16

// AVPixFmtGbrap10 wraps AV_PIX_FMT_GBRAP10.
const AVPixFmtGbrap10 = C.AV_PIX_FMT_GBRAP10

// AVPixFmtGbrap12 wraps AV_PIX_FMT_GBRAP12.
const AVPixFmtGbrap12 = C.AV_PIX_FMT_GBRAP12

// AVPixFmtGbrap14 wraps AV_PIX_FMT_GBRAP14.
const AVPixFmtGbrap14 = C.AV_PIX_FMT_GBRAP14

// AVPixFmtGbrap16 wraps AV_PIX_FMT_GBRAP16.
const AVPixFmtGbrap16 = C.AV_PIX_FMT_GBRAP16

// AVPixFmtBayerBggr16 wraps AV_PIX_FMT_BAYER_BGGR16.
const AVPixFmtBayerBggr16 = C.AV_PIX_FMT_BAYER_BGGR16

// AVPixFmtBayerRggb16 wraps AV_PIX_FMT_BAYER_RGGB16.
const AVPixFmtBayerRggb16 = C.AV_PIX_FMT_BAYER_RGGB16

// AVPixFmtBayerGbrg16 wraps AV_PIX_FMT_BAYER_GBRG16.
const AVPixFmtBayerGbrg16 = C.AV_PIX_FMT_BAYER_GBRG16

// AVPixFmtBayerGrbg16 wraps AV_PIX_FMT_BAYER_GRBG16.
const AVPixFmtBayerGrbg16 = C.AV_PIX_FMT_BAYER_GRBG16

// AVPixFmtGbrpf32 wraps AV_PIX_FMT_GBRPF32.
const AVPixFmtGbrpf32 = C.AV_PIX_FMT_GBRPF32

// AVPixFmtGbrapf32 wraps AV_PIX_FMT_GBRAPF32.
const AVPixFmtGbrapf32 = C.AV_PIX_FMT_GBRAPF32

// AVPixFmtGrayf32 wraps AV_PIX_FMT_GRAYF32.
const AVPixFmtGrayf32 = C.AV_PIX_FMT_GRAYF32

// AVPixFmtYuva420P9 wraps AV_PIX_FMT_YUVA420P9.
const AVPixFmtYuva420P9 = C.AV_PIX_FMT_YUVA420P9

// AVPixFmtYuva422P9 wraps AV_PIX_FMT_YUVA422P9.
const AVPixFmtYuva422P9 = C.AV_PIX_FMT_YUVA422P9

// AVPixFmtYuva444P9 wraps AV_PIX_FMT_YUVA444P9.
const AVPixFmtYuva444P9 = C.AV_PIX_FMT_YUVA444P9

// AVPixFmtYuva420P10 wraps AV_PIX_FMT_YUVA420P10.
const AVPixFmtYuva420P10 = C.AV_PIX_FMT_YUVA420P10

// AVPixFmtYuva422P10 wraps AV_PIX_FMT_YUVA422P10.
const AVPixFmtYuva422P10 = C.AV_PIX_FMT_YUVA422P10

// AVPixFmtYuva444P10 wraps AV_PIX_FMT_YUVA444P10.
const AVPixFmtYuva444P10 = C.AV_PIX_FMT_YUVA444P10

// AVPixFmtYuva422P12 wraps AV_PIX_FMT_YUVA422P12.
const AVPixFmtYuva422P12 = C.AV_PIX_FMT_YUVA422P12

// AVPixFmtYuva444P12 wraps AV_PIX_FMT_YUVA444P12.
const AVPixFmtYuva444P12 = C.AV_PIX_FMT_YUVA444P12

// AVPixFmtYuva420P16 wraps AV_PIX_FMT_YUVA420P16.
const AVPixFmtYuva420P16 = C.AV_PIX_FMT_YUVA420P16

// AVPixFmtYuva422P16 wraps AV_PIX_FMT_YUVA422P16.
const AVPixFmtYuva422P16 = C.AV_PIX_FMT_YUVA422P16

// AVPixFmtYuva444P16 wraps AV_PIX_FMT_YUVA444P16.
const AVPixFmtYuva444P16 = C.AV_PIX_FMT_YUVA444P16

// AVPixFmtXyz12 wraps AV_PIX_FMT_XYZ12.
const AVPixFmtXyz12 = C.AV_PIX_FMT_XYZ12

// AVPixFmtNv20 wraps AV_PIX_FMT_NV20.
const AVPixFmtNv20 = C.AV_PIX_FMT_NV20

// AVPixFmtAyuv64 wraps AV_PIX_FMT_AYUV64.
const AVPixFmtAyuv64 = C.AV_PIX_FMT_AYUV64

// AVPixFmtP010 wraps AV_PIX_FMT_P010.
const AVPixFmtP010 = C.AV_PIX_FMT_P010

// AVPixFmtP012 wraps AV_PIX_FMT_P012.
const AVPixFmtP012 = C.AV_PIX_FMT_P012

// AVPixFmtP016 wraps AV_PIX_FMT_P016.
const AVPixFmtP016 = C.AV_PIX_FMT_P016

// AVPixFmtY210 wraps AV_PIX_FMT_Y210.
const AVPixFmtY210 = C.AV_PIX_FMT_Y210

// AVPixFmtY212 wraps AV_PIX_FMT_Y212.
const AVPixFmtY212 = C.AV_PIX_FMT_Y212

// AVPixFmtXv30 wraps AV_PIX_FMT_XV30.
const AVPixFmtXv30 = C.AV_PIX_FMT_XV30

// AVPixFmtXv36 wraps AV_PIX_FMT_XV36.
const AVPixFmtXv36 = C.AV_PIX_FMT_XV36

// AVPixFmtX2Rgb10 wraps AV_PIX_FMT_X2RGB10.
const AVPixFmtX2Rgb10 = C.AV_PIX_FMT_X2RGB10

// AVPixFmtX2Bgr10 wraps AV_PIX_FMT_X2BGR10.
const AVPixFmtX2Bgr10 = C.AV_PIX_FMT_X2BGR10

// AVPixFmtP210 wraps AV_PIX_FMT_P210.
const AVPixFmtP210 = C.AV_PIX_FMT_P210

// AVPixFmtP410 wraps AV_PIX_FMT_P410.
const AVPixFmtP410 = C.AV_PIX_FMT_P410

// AVPixFmtP212 wraps AV_PIX_FMT_P212.
const AVPixFmtP212 = C.AV_PIX_FMT_P212

// AVPixFmtP412 wraps AV_PIX_FMT_P412.
const AVPixFmtP412 = C.AV_PIX_FMT_P412

// AVPixFmtP216 wraps AV_PIX_FMT_P216.
const AVPixFmtP216 = C.AV_PIX_FMT_P216

// AVPixFmtP416 wraps AV_PIX_FMT_P416.
const AVPixFmtP416 = C.AV_PIX_FMT_P416

// AVPixFmtRgbaf16 wraps AV_PIX_FMT_RGBAF16.
const AVPixFmtRgbaf16 = C.AV_PIX_FMT_RGBAF16

// AVPixFmtRgbf32 wraps AV_PIX_FMT_RGBF32.
const AVPixFmtRgbf32 = C.AV_PIX_FMT_RGBF32

// AVPixFmtRgbaf32 wraps AV_PIX_FMT_RGBAF32.
const AVPixFmtRgbaf32 = C.AV_PIX_FMT_RGBAF32

// LIBAVUtilVersionMajor wraps LIBAVUTIL_VERSION_MAJOR.
const LIBAVUtilVersionMajor = C.LIBAVUTIL_VERSION_MAJOR

// LIBAVUtilVersionMinor wraps LIBAVUTIL_VERSION_MINOR.
const LIBAVUtilVersionMinor = C.LIBAVUTIL_VERSION_MINOR

// LIBAVUtilVersionMicro wraps LIBAVUTIL_VERSION_MICRO.
const LIBAVUtilVersionMicro = C.LIBAVUTIL_VERSION_MICRO

// LIBAVUtilVersionInt wraps LIBAVUTIL_VERSION_INT.
const LIBAVUtilVersionInt = C.LIBAVUTIL_VERSION_INT

// LIBAVUtilBuild wraps LIBAVUTIL_BUILD.
const LIBAVUtilBuild = C.LIBAVUTIL_BUILD

// FFAPIFifoPeek2 wraps FF_API_FIFO_PEEK2.
const FFAPIFifoPeek2 = C.FF_API_FIFO_PEEK2

// FFAPIFifoOldAPI wraps FF_API_FIFO_OLD_API.
const FFAPIFifoOldAPI = C.FF_API_FIFO_OLD_API

// FFAPIXvmc wraps FF_API_XVMC.
const FFAPIXvmc = C.FF_API_XVMC

// FFAPIOldChannelLayout wraps FF_API_OLD_CHANNEL_LAYOUT.
const FFAPIOldChannelLayout = C.FF_API_OLD_CHANNEL_LAYOUT

// FFAPIAVFopenUtf8 wraps FF_API_AV_FOPEN_UTF8.
const FFAPIAVFopenUtf8 = C.FF_API_AV_FOPEN_UTF8

// FFAPIPktDuration wraps FF_API_PKT_DURATION.
const FFAPIPktDuration = C.FF_API_PKT_DURATION

// FFAPIReorderedOpaque wraps FF_API_REORDERED_OPAQUE.
const FFAPIReorderedOpaque = C.FF_API_REORDERED_OPAQUE

// FFAPIFramePictureNumber wraps FF_API_FRAME_PICTURE_NUMBER.
const FFAPIFramePictureNumber = C.FF_API_FRAME_PICTURE_NUMBER

// FFAPIHdrVividThreeSpline wraps FF_API_HDR_VIVID_THREE_SPLINE.
const FFAPIHdrVividThreeSpline = C.FF_API_HDR_VIVID_THREE_SPLINE

// FFAPIFramePkt wraps FF_API_FRAME_PKT.
const FFAPIFramePkt = C.FF_API_FRAME_PKT

// FFAPIInterlacedFrame wraps FF_API_INTERLACED_FRAME.
const FFAPIInterlacedFrame = C.FF_API_INTERLACED_FRAME

// FFAPIFrameKey wraps FF_API_FRAME_KEY.
const FFAPIFrameKey = C.FF_API_FRAME_KEY

// FFAPIPaletteHasChanged wraps FF_API_PALETTE_HAS_CHANGED.
const FFAPIPaletteHasChanged = C.FF_API_PALETTE_HAS_CHANGED

// FFAPIVulkanContiguousMemory wraps FF_API_VULKAN_CONTIGUOUS_MEMORY.
const FFAPIVulkanContiguousMemory = C.FF_API_VULKAN_CONTIGUOUS_MEMORY

// LIBPostprocVersionMinor wraps LIBPOSTPROC_VERSION_MINOR.
const LIBPostprocVersionMinor = C.LIBPOSTPROC_VERSION_MINOR

// LIBPostprocVersionMicro wraps LIBPOSTPROC_VERSION_MICRO.
const LIBPostprocVersionMicro = C.LIBPOSTPROC_VERSION_MICRO

// LIBPostprocVersionInt wraps LIBPOSTPROC_VERSION_INT.
const LIBPostprocVersionInt = C.LIBPOSTPROC_VERSION_INT

// LIBPostprocBuild wraps LIBPOSTPROC_BUILD.
const LIBPostprocBuild = C.LIBPOSTPROC_BUILD

// LIBPostprocVersionMajor wraps LIBPOSTPROC_VERSION_MAJOR.
const LIBPostprocVersionMajor = C.LIBPOSTPROC_VERSION_MAJOR

// LIBSwresampleVersionMinor wraps LIBSWRESAMPLE_VERSION_MINOR.
const LIBSwresampleVersionMinor = C.LIBSWRESAMPLE_VERSION_MINOR

// LIBSwresampleVersionMicro wraps LIBSWRESAMPLE_VERSION_MICRO.
const LIBSwresampleVersionMicro = C.LIBSWRESAMPLE_VERSION_MICRO

// LIBSwresampleVersionInt wraps LIBSWRESAMPLE_VERSION_INT.
const LIBSwresampleVersionInt = C.LIBSWRESAMPLE_VERSION_INT

// LIBSwresampleBuild wraps LIBSWRESAMPLE_BUILD.
const LIBSwresampleBuild = C.LIBSWRESAMPLE_BUILD

// LIBSwresampleVersionMajor wraps LIBSWRESAMPLE_VERSION_MAJOR.
const LIBSwresampleVersionMajor = C.LIBSWRESAMPLE_VERSION_MAJOR

// LIBSwscaleVersionMinor wraps LIBSWSCALE_VERSION_MINOR.
const LIBSwscaleVersionMinor = C.LIBSWSCALE_VERSION_MINOR

// LIBSwscaleVersionMicro wraps LIBSWSCALE_VERSION_MICRO.
const LIBSwscaleVersionMicro = C.LIBSWSCALE_VERSION_MICRO

// LIBSwscaleVersionInt wraps LIBSWSCALE_VERSION_INT.
const LIBSwscaleVersionInt = C.LIBSWSCALE_VERSION_INT

// LIBSwscaleBuild wraps LIBSWSCALE_BUILD.
const LIBSwscaleBuild = C.LIBSWSCALE_BUILD

// LIBSwscaleVersionMajor wraps LIBSWSCALE_VERSION_MAJOR.
const LIBSwscaleVersionMajor = C.LIBSWSCALE_VERSION_MAJOR
