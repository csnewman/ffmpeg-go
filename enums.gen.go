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

// --- Enum AVSubtitleType ---

// AVSubtitleType wraps AVSubtitleType.
type AVSubtitleType C.enum_AVSubtitleType

const SizeOfAVSubtitleType = C.sizeof_enum_AVSubtitleType

func ToAVSubtitleTypeArray(ptr unsafe.Pointer) *Array[AVSubtitleType] {
	if ptr == nil {
		return nil
	}

	return &Array[AVSubtitleType]{
		elemSize: SizeOfAVSubtitleType,
		loadPtr: func(pointer unsafe.Pointer) AVSubtitleType {
			ptr := (*AVSubtitleType)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVSubtitleType) {
			ptr := (*AVSubtitleType)(pointer)
			*ptr = value
		},
	}
}

func AllocAVSubtitleTypeArray(size uint64) *Array[AVSubtitleType] {
	return ToAVSubtitleTypeArray(AVCalloc(size, SizeOfAVSubtitleType))
}

const (
	// SubtitleNone wraps SUBTITLE_NONE.
	SubtitleNone AVSubtitleType = C.SUBTITLE_NONE
	// SubtitleBitmap wraps SUBTITLE_BITMAP.
	//
	//	A bitmap, pict will be set
	SubtitleBitmap AVSubtitleType = C.SUBTITLE_BITMAP
	// SubtitleText wraps SUBTITLE_TEXT.
	/*
	   Plain text, the text field must be set by the decoder and is
	   authoritative. ass and pict fields may contain approximations.
	*/
	SubtitleText AVSubtitleType = C.SUBTITLE_TEXT
	// SubtitleAss wraps SUBTITLE_ASS.
	/*
	   Formatted text, the ass field must be set by the decoder and is
	   authoritative. pict and text fields may contain approximations.
	*/
	SubtitleAss AVSubtitleType = C.SUBTITLE_ASS
)

// --- Enum AVPictureStructure ---

// AVPictureStructure wraps AVPictureStructure.
type AVPictureStructure C.enum_AVPictureStructure

const SizeOfAVPictureStructure = C.sizeof_enum_AVPictureStructure

func ToAVPictureStructureArray(ptr unsafe.Pointer) *Array[AVPictureStructure] {
	if ptr == nil {
		return nil
	}

	return &Array[AVPictureStructure]{
		elemSize: SizeOfAVPictureStructure,
		loadPtr: func(pointer unsafe.Pointer) AVPictureStructure {
			ptr := (*AVPictureStructure)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVPictureStructure) {
			ptr := (*AVPictureStructure)(pointer)
			*ptr = value
		},
	}
}

func AllocAVPictureStructureArray(size uint64) *Array[AVPictureStructure] {
	return ToAVPictureStructureArray(AVCalloc(size, SizeOfAVPictureStructure))
}

const (
	// AVPictureStructureUnknown wraps AV_PICTURE_STRUCTURE_UNKNOWN.
	//
	//	unknown
	AVPictureStructureUnknown AVPictureStructure = C.AV_PICTURE_STRUCTURE_UNKNOWN
	// AVPictureStructureTopField wraps AV_PICTURE_STRUCTURE_TOP_FIELD.
	//
	//	coded as top field
	AVPictureStructureTopField AVPictureStructure = C.AV_PICTURE_STRUCTURE_TOP_FIELD
	// AVPictureStructureBottomField wraps AV_PICTURE_STRUCTURE_BOTTOM_FIELD.
	//
	//	coded as bottom field
	AVPictureStructureBottomField AVPictureStructure = C.AV_PICTURE_STRUCTURE_BOTTOM_FIELD
	// AVPictureStructureFrame wraps AV_PICTURE_STRUCTURE_FRAME.
	//
	//	coded as frame
	AVPictureStructureFrame AVPictureStructure = C.AV_PICTURE_STRUCTURE_FRAME
)

// --- Enum AVCodecID ---

// AVCodecID wraps AVCodecID.
/*
  Identify the syntax and semantics of the bitstream.
  The principle is roughly:
  Two decoders with the same ID can decode the same streams.
  Two encoders with the same ID can encode compatible streams.
  There may be slight deviations from the principle due to implementation
  details.

  If you add a codec ID to this list, add it so that
  1. no value of an existing codec ID changes (that would break ABI),
  2. it is as close as possible to similar codecs

  After adding new codec IDs, do not forget to add an entry to the codec
  descriptor list and bump libavcodec minor version.
*/
type AVCodecID C.enum_AVCodecID

const SizeOfAVCodecID = C.sizeof_enum_AVCodecID

func ToAVCodecIDArray(ptr unsafe.Pointer) *Array[AVCodecID] {
	if ptr == nil {
		return nil
	}

	return &Array[AVCodecID]{
		elemSize: SizeOfAVCodecID,
		loadPtr: func(pointer unsafe.Pointer) AVCodecID {
			ptr := (*AVCodecID)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVCodecID) {
			ptr := (*AVCodecID)(pointer)
			*ptr = value
		},
	}
}

func AllocAVCodecIDArray(size uint64) *Array[AVCodecID] {
	return ToAVCodecIDArray(AVCalloc(size, SizeOfAVCodecID))
}

const (
	// AVCodecIdNone wraps AV_CODEC_ID_NONE.
	AVCodecIdNone AVCodecID = C.AV_CODEC_ID_NONE
	// AVCodecIdMpeg1Video wraps AV_CODEC_ID_MPEG1VIDEO.
	//
	//	video codecs
	AVCodecIdMpeg1Video AVCodecID = C.AV_CODEC_ID_MPEG1VIDEO
	// AVCodecIdMpeg2Video wraps AV_CODEC_ID_MPEG2VIDEO.
	//
	//	preferred ID for MPEG-1/2 video decoding
	AVCodecIdMpeg2Video AVCodecID = C.AV_CODEC_ID_MPEG2VIDEO
	// AVCodecIdH261 wraps AV_CODEC_ID_H261.
	AVCodecIdH261 AVCodecID = C.AV_CODEC_ID_H261
	// AVCodecIdH263 wraps AV_CODEC_ID_H263.
	AVCodecIdH263 AVCodecID = C.AV_CODEC_ID_H263
	// AVCodecIdRv10 wraps AV_CODEC_ID_RV10.
	AVCodecIdRv10 AVCodecID = C.AV_CODEC_ID_RV10
	// AVCodecIdRv20 wraps AV_CODEC_ID_RV20.
	AVCodecIdRv20 AVCodecID = C.AV_CODEC_ID_RV20
	// AVCodecIdMjpeg wraps AV_CODEC_ID_MJPEG.
	AVCodecIdMjpeg AVCodecID = C.AV_CODEC_ID_MJPEG
	// AVCodecIdMjpegb wraps AV_CODEC_ID_MJPEGB.
	AVCodecIdMjpegb AVCodecID = C.AV_CODEC_ID_MJPEGB
	// AVCodecIdLjpeg wraps AV_CODEC_ID_LJPEG.
	AVCodecIdLjpeg AVCodecID = C.AV_CODEC_ID_LJPEG
	// AVCodecIdSp5X wraps AV_CODEC_ID_SP5X.
	AVCodecIdSp5X AVCodecID = C.AV_CODEC_ID_SP5X
	// AVCodecIdJpegls wraps AV_CODEC_ID_JPEGLS.
	AVCodecIdJpegls AVCodecID = C.AV_CODEC_ID_JPEGLS
	// AVCodecIdMpeg4 wraps AV_CODEC_ID_MPEG4.
	AVCodecIdMpeg4 AVCodecID = C.AV_CODEC_ID_MPEG4
	// AVCodecIdRawvideo wraps AV_CODEC_ID_RAWVIDEO.
	AVCodecIdRawvideo AVCodecID = C.AV_CODEC_ID_RAWVIDEO
	// AVCodecIdMsmpeg4V1 wraps AV_CODEC_ID_MSMPEG4V1.
	AVCodecIdMsmpeg4V1 AVCodecID = C.AV_CODEC_ID_MSMPEG4V1
	// AVCodecIdMsmpeg4V2 wraps AV_CODEC_ID_MSMPEG4V2.
	AVCodecIdMsmpeg4V2 AVCodecID = C.AV_CODEC_ID_MSMPEG4V2
	// AVCodecIdMsmpeg4V3 wraps AV_CODEC_ID_MSMPEG4V3.
	AVCodecIdMsmpeg4V3 AVCodecID = C.AV_CODEC_ID_MSMPEG4V3
	// AVCodecIdWmv1 wraps AV_CODEC_ID_WMV1.
	AVCodecIdWmv1 AVCodecID = C.AV_CODEC_ID_WMV1
	// AVCodecIdWmv2 wraps AV_CODEC_ID_WMV2.
	AVCodecIdWmv2 AVCodecID = C.AV_CODEC_ID_WMV2
	// AVCodecIdH263P wraps AV_CODEC_ID_H263P.
	AVCodecIdH263P AVCodecID = C.AV_CODEC_ID_H263P
	// AVCodecIdH263I wraps AV_CODEC_ID_H263I.
	AVCodecIdH263I AVCodecID = C.AV_CODEC_ID_H263I
	// AVCodecIdFlv1 wraps AV_CODEC_ID_FLV1.
	AVCodecIdFlv1 AVCodecID = C.AV_CODEC_ID_FLV1
	// AVCodecIdSvq1 wraps AV_CODEC_ID_SVQ1.
	AVCodecIdSvq1 AVCodecID = C.AV_CODEC_ID_SVQ1
	// AVCodecIdSvq3 wraps AV_CODEC_ID_SVQ3.
	AVCodecIdSvq3 AVCodecID = C.AV_CODEC_ID_SVQ3
	// AVCodecIdDvvideo wraps AV_CODEC_ID_DVVIDEO.
	AVCodecIdDvvideo AVCodecID = C.AV_CODEC_ID_DVVIDEO
	// AVCodecIdHuffyuv wraps AV_CODEC_ID_HUFFYUV.
	AVCodecIdHuffyuv AVCodecID = C.AV_CODEC_ID_HUFFYUV
	// AVCodecIdCyuv wraps AV_CODEC_ID_CYUV.
	AVCodecIdCyuv AVCodecID = C.AV_CODEC_ID_CYUV
	// AVCodecIdH264 wraps AV_CODEC_ID_H264.
	AVCodecIdH264 AVCodecID = C.AV_CODEC_ID_H264
	// AVCodecIdIndeo3 wraps AV_CODEC_ID_INDEO3.
	AVCodecIdIndeo3 AVCodecID = C.AV_CODEC_ID_INDEO3
	// AVCodecIdVp3 wraps AV_CODEC_ID_VP3.
	AVCodecIdVp3 AVCodecID = C.AV_CODEC_ID_VP3
	// AVCodecIdTheora wraps AV_CODEC_ID_THEORA.
	AVCodecIdTheora AVCodecID = C.AV_CODEC_ID_THEORA
	// AVCodecIdAsv1 wraps AV_CODEC_ID_ASV1.
	AVCodecIdAsv1 AVCodecID = C.AV_CODEC_ID_ASV1
	// AVCodecIdAsv2 wraps AV_CODEC_ID_ASV2.
	AVCodecIdAsv2 AVCodecID = C.AV_CODEC_ID_ASV2
	// AVCodecIdFfv1 wraps AV_CODEC_ID_FFV1.
	AVCodecIdFfv1 AVCodecID = C.AV_CODEC_ID_FFV1
	// AVCodecId4Xm wraps AV_CODEC_ID_4XM.
	AVCodecId4Xm AVCodecID = C.AV_CODEC_ID_4XM
	// AVCodecIdVcr1 wraps AV_CODEC_ID_VCR1.
	AVCodecIdVcr1 AVCodecID = C.AV_CODEC_ID_VCR1
	// AVCodecIdCljr wraps AV_CODEC_ID_CLJR.
	AVCodecIdCljr AVCodecID = C.AV_CODEC_ID_CLJR
	// AVCodecIdMdec wraps AV_CODEC_ID_MDEC.
	AVCodecIdMdec AVCodecID = C.AV_CODEC_ID_MDEC
	// AVCodecIdRoq wraps AV_CODEC_ID_ROQ.
	AVCodecIdRoq AVCodecID = C.AV_CODEC_ID_ROQ
	// AVCodecIdInterplayVideo wraps AV_CODEC_ID_INTERPLAY_VIDEO.
	AVCodecIdInterplayVideo AVCodecID = C.AV_CODEC_ID_INTERPLAY_VIDEO
	// AVCodecIdXanWc3 wraps AV_CODEC_ID_XAN_WC3.
	AVCodecIdXanWc3 AVCodecID = C.AV_CODEC_ID_XAN_WC3
	// AVCodecIdXanWc4 wraps AV_CODEC_ID_XAN_WC4.
	AVCodecIdXanWc4 AVCodecID = C.AV_CODEC_ID_XAN_WC4
	// AVCodecIdRpza wraps AV_CODEC_ID_RPZA.
	AVCodecIdRpza AVCodecID = C.AV_CODEC_ID_RPZA
	// AVCodecIdCinepak wraps AV_CODEC_ID_CINEPAK.
	AVCodecIdCinepak AVCodecID = C.AV_CODEC_ID_CINEPAK
	// AVCodecIdWsVqa wraps AV_CODEC_ID_WS_VQA.
	AVCodecIdWsVqa AVCodecID = C.AV_CODEC_ID_WS_VQA
	// AVCodecIdMsrle wraps AV_CODEC_ID_MSRLE.
	AVCodecIdMsrle AVCodecID = C.AV_CODEC_ID_MSRLE
	// AVCodecIdMsvideo1 wraps AV_CODEC_ID_MSVIDEO1.
	AVCodecIdMsvideo1 AVCodecID = C.AV_CODEC_ID_MSVIDEO1
	// AVCodecIdIdcin wraps AV_CODEC_ID_IDCIN.
	AVCodecIdIdcin AVCodecID = C.AV_CODEC_ID_IDCIN
	// AVCodecId8Bps wraps AV_CODEC_ID_8BPS.
	AVCodecId8Bps AVCodecID = C.AV_CODEC_ID_8BPS
	// AVCodecIdSmc wraps AV_CODEC_ID_SMC.
	AVCodecIdSmc AVCodecID = C.AV_CODEC_ID_SMC
	// AVCodecIdFlic wraps AV_CODEC_ID_FLIC.
	AVCodecIdFlic AVCodecID = C.AV_CODEC_ID_FLIC
	// AVCodecIdTruemotion1 wraps AV_CODEC_ID_TRUEMOTION1.
	AVCodecIdTruemotion1 AVCodecID = C.AV_CODEC_ID_TRUEMOTION1
	// AVCodecIdVmdvideo wraps AV_CODEC_ID_VMDVIDEO.
	AVCodecIdVmdvideo AVCodecID = C.AV_CODEC_ID_VMDVIDEO
	// AVCodecIdMszh wraps AV_CODEC_ID_MSZH.
	AVCodecIdMszh AVCodecID = C.AV_CODEC_ID_MSZH
	// AVCodecIdZlib wraps AV_CODEC_ID_ZLIB.
	AVCodecIdZlib AVCodecID = C.AV_CODEC_ID_ZLIB
	// AVCodecIdQtrle wraps AV_CODEC_ID_QTRLE.
	AVCodecIdQtrle AVCodecID = C.AV_CODEC_ID_QTRLE
	// AVCodecIdTscc wraps AV_CODEC_ID_TSCC.
	AVCodecIdTscc AVCodecID = C.AV_CODEC_ID_TSCC
	// AVCodecIdUlti wraps AV_CODEC_ID_ULTI.
	AVCodecIdUlti AVCodecID = C.AV_CODEC_ID_ULTI
	// AVCodecIdQdraw wraps AV_CODEC_ID_QDRAW.
	AVCodecIdQdraw AVCodecID = C.AV_CODEC_ID_QDRAW
	// AVCodecIdVixl wraps AV_CODEC_ID_VIXL.
	AVCodecIdVixl AVCodecID = C.AV_CODEC_ID_VIXL
	// AVCodecIdQpeg wraps AV_CODEC_ID_QPEG.
	AVCodecIdQpeg AVCodecID = C.AV_CODEC_ID_QPEG
	// AVCodecIdPng wraps AV_CODEC_ID_PNG.
	AVCodecIdPng AVCodecID = C.AV_CODEC_ID_PNG
	// AVCodecIdPpm wraps AV_CODEC_ID_PPM.
	AVCodecIdPpm AVCodecID = C.AV_CODEC_ID_PPM
	// AVCodecIdPbm wraps AV_CODEC_ID_PBM.
	AVCodecIdPbm AVCodecID = C.AV_CODEC_ID_PBM
	// AVCodecIdPgm wraps AV_CODEC_ID_PGM.
	AVCodecIdPgm AVCodecID = C.AV_CODEC_ID_PGM
	// AVCodecIdPgmyuv wraps AV_CODEC_ID_PGMYUV.
	AVCodecIdPgmyuv AVCodecID = C.AV_CODEC_ID_PGMYUV
	// AVCodecIdPam wraps AV_CODEC_ID_PAM.
	AVCodecIdPam AVCodecID = C.AV_CODEC_ID_PAM
	// AVCodecIdFfvhuff wraps AV_CODEC_ID_FFVHUFF.
	AVCodecIdFfvhuff AVCodecID = C.AV_CODEC_ID_FFVHUFF
	// AVCodecIdRv30 wraps AV_CODEC_ID_RV30.
	AVCodecIdRv30 AVCodecID = C.AV_CODEC_ID_RV30
	// AVCodecIdRv40 wraps AV_CODEC_ID_RV40.
	AVCodecIdRv40 AVCodecID = C.AV_CODEC_ID_RV40
	// AVCodecIdVc1 wraps AV_CODEC_ID_VC1.
	AVCodecIdVc1 AVCodecID = C.AV_CODEC_ID_VC1
	// AVCodecIdWmv3 wraps AV_CODEC_ID_WMV3.
	AVCodecIdWmv3 AVCodecID = C.AV_CODEC_ID_WMV3
	// AVCodecIdLoco wraps AV_CODEC_ID_LOCO.
	AVCodecIdLoco AVCodecID = C.AV_CODEC_ID_LOCO
	// AVCodecIdWnv1 wraps AV_CODEC_ID_WNV1.
	AVCodecIdWnv1 AVCodecID = C.AV_CODEC_ID_WNV1
	// AVCodecIdAasc wraps AV_CODEC_ID_AASC.
	AVCodecIdAasc AVCodecID = C.AV_CODEC_ID_AASC
	// AVCodecIdIndeo2 wraps AV_CODEC_ID_INDEO2.
	AVCodecIdIndeo2 AVCodecID = C.AV_CODEC_ID_INDEO2
	// AVCodecIdFraps wraps AV_CODEC_ID_FRAPS.
	AVCodecIdFraps AVCodecID = C.AV_CODEC_ID_FRAPS
	// AVCodecIdTruemotion2 wraps AV_CODEC_ID_TRUEMOTION2.
	AVCodecIdTruemotion2 AVCodecID = C.AV_CODEC_ID_TRUEMOTION2
	// AVCodecIdBmp wraps AV_CODEC_ID_BMP.
	AVCodecIdBmp AVCodecID = C.AV_CODEC_ID_BMP
	// AVCodecIdCscd wraps AV_CODEC_ID_CSCD.
	AVCodecIdCscd AVCodecID = C.AV_CODEC_ID_CSCD
	// AVCodecIdMmvideo wraps AV_CODEC_ID_MMVIDEO.
	AVCodecIdMmvideo AVCodecID = C.AV_CODEC_ID_MMVIDEO
	// AVCodecIdZmbv wraps AV_CODEC_ID_ZMBV.
	AVCodecIdZmbv AVCodecID = C.AV_CODEC_ID_ZMBV
	// AVCodecIdAvs wraps AV_CODEC_ID_AVS.
	AVCodecIdAvs AVCodecID = C.AV_CODEC_ID_AVS
	// AVCodecIdSmackvideo wraps AV_CODEC_ID_SMACKVIDEO.
	AVCodecIdSmackvideo AVCodecID = C.AV_CODEC_ID_SMACKVIDEO
	// AVCodecIdNuv wraps AV_CODEC_ID_NUV.
	AVCodecIdNuv AVCodecID = C.AV_CODEC_ID_NUV
	// AVCodecIdKmvc wraps AV_CODEC_ID_KMVC.
	AVCodecIdKmvc AVCodecID = C.AV_CODEC_ID_KMVC
	// AVCodecIdFlashsv wraps AV_CODEC_ID_FLASHSV.
	AVCodecIdFlashsv AVCodecID = C.AV_CODEC_ID_FLASHSV
	// AVCodecIdCavs wraps AV_CODEC_ID_CAVS.
	AVCodecIdCavs AVCodecID = C.AV_CODEC_ID_CAVS
	// AVCodecIdJpeg2000 wraps AV_CODEC_ID_JPEG2000.
	AVCodecIdJpeg2000 AVCodecID = C.AV_CODEC_ID_JPEG2000
	// AVCodecIdVmnc wraps AV_CODEC_ID_VMNC.
	AVCodecIdVmnc AVCodecID = C.AV_CODEC_ID_VMNC
	// AVCodecIdVp5 wraps AV_CODEC_ID_VP5.
	AVCodecIdVp5 AVCodecID = C.AV_CODEC_ID_VP5
	// AVCodecIdVp6 wraps AV_CODEC_ID_VP6.
	AVCodecIdVp6 AVCodecID = C.AV_CODEC_ID_VP6
	// AVCodecIdVp6F wraps AV_CODEC_ID_VP6F.
	AVCodecIdVp6F AVCodecID = C.AV_CODEC_ID_VP6F
	// AVCodecIdTarga wraps AV_CODEC_ID_TARGA.
	AVCodecIdTarga AVCodecID = C.AV_CODEC_ID_TARGA
	// AVCodecIdDsicinvideo wraps AV_CODEC_ID_DSICINVIDEO.
	AVCodecIdDsicinvideo AVCodecID = C.AV_CODEC_ID_DSICINVIDEO
	// AVCodecIdTiertexseqvideo wraps AV_CODEC_ID_TIERTEXSEQVIDEO.
	AVCodecIdTiertexseqvideo AVCodecID = C.AV_CODEC_ID_TIERTEXSEQVIDEO
	// AVCodecIdTiff wraps AV_CODEC_ID_TIFF.
	AVCodecIdTiff AVCodecID = C.AV_CODEC_ID_TIFF
	// AVCodecIdGif wraps AV_CODEC_ID_GIF.
	AVCodecIdGif AVCodecID = C.AV_CODEC_ID_GIF
	// AVCodecIdDxa wraps AV_CODEC_ID_DXA.
	AVCodecIdDxa AVCodecID = C.AV_CODEC_ID_DXA
	// AVCodecIdDnxhd wraps AV_CODEC_ID_DNXHD.
	AVCodecIdDnxhd AVCodecID = C.AV_CODEC_ID_DNXHD
	// AVCodecIdThp wraps AV_CODEC_ID_THP.
	AVCodecIdThp AVCodecID = C.AV_CODEC_ID_THP
	// AVCodecIdSgi wraps AV_CODEC_ID_SGI.
	AVCodecIdSgi AVCodecID = C.AV_CODEC_ID_SGI
	// AVCodecIdC93 wraps AV_CODEC_ID_C93.
	AVCodecIdC93 AVCodecID = C.AV_CODEC_ID_C93
	// AVCodecIdBethsoftvid wraps AV_CODEC_ID_BETHSOFTVID.
	AVCodecIdBethsoftvid AVCodecID = C.AV_CODEC_ID_BETHSOFTVID
	// AVCodecIdPtx wraps AV_CODEC_ID_PTX.
	AVCodecIdPtx AVCodecID = C.AV_CODEC_ID_PTX
	// AVCodecIdTxd wraps AV_CODEC_ID_TXD.
	AVCodecIdTxd AVCodecID = C.AV_CODEC_ID_TXD
	// AVCodecIdVp6A wraps AV_CODEC_ID_VP6A.
	AVCodecIdVp6A AVCodecID = C.AV_CODEC_ID_VP6A
	// AVCodecIdAmv wraps AV_CODEC_ID_AMV.
	AVCodecIdAmv AVCodecID = C.AV_CODEC_ID_AMV
	// AVCodecIdVb wraps AV_CODEC_ID_VB.
	AVCodecIdVb AVCodecID = C.AV_CODEC_ID_VB
	// AVCodecIdPcx wraps AV_CODEC_ID_PCX.
	AVCodecIdPcx AVCodecID = C.AV_CODEC_ID_PCX
	// AVCodecIdSunrast wraps AV_CODEC_ID_SUNRAST.
	AVCodecIdSunrast AVCodecID = C.AV_CODEC_ID_SUNRAST
	// AVCodecIdIndeo4 wraps AV_CODEC_ID_INDEO4.
	AVCodecIdIndeo4 AVCodecID = C.AV_CODEC_ID_INDEO4
	// AVCodecIdIndeo5 wraps AV_CODEC_ID_INDEO5.
	AVCodecIdIndeo5 AVCodecID = C.AV_CODEC_ID_INDEO5
	// AVCodecIdMimic wraps AV_CODEC_ID_MIMIC.
	AVCodecIdMimic AVCodecID = C.AV_CODEC_ID_MIMIC
	// AVCodecIdRl2 wraps AV_CODEC_ID_RL2.
	AVCodecIdRl2 AVCodecID = C.AV_CODEC_ID_RL2
	// AVCodecIdEscape124 wraps AV_CODEC_ID_ESCAPE124.
	AVCodecIdEscape124 AVCodecID = C.AV_CODEC_ID_ESCAPE124
	// AVCodecIdDirac wraps AV_CODEC_ID_DIRAC.
	AVCodecIdDirac AVCodecID = C.AV_CODEC_ID_DIRAC
	// AVCodecIdBfi wraps AV_CODEC_ID_BFI.
	AVCodecIdBfi AVCodecID = C.AV_CODEC_ID_BFI
	// AVCodecIdCmv wraps AV_CODEC_ID_CMV.
	AVCodecIdCmv AVCodecID = C.AV_CODEC_ID_CMV
	// AVCodecIdMotionpixels wraps AV_CODEC_ID_MOTIONPIXELS.
	AVCodecIdMotionpixels AVCodecID = C.AV_CODEC_ID_MOTIONPIXELS
	// AVCodecIdTgv wraps AV_CODEC_ID_TGV.
	AVCodecIdTgv AVCodecID = C.AV_CODEC_ID_TGV
	// AVCodecIdTgq wraps AV_CODEC_ID_TGQ.
	AVCodecIdTgq AVCodecID = C.AV_CODEC_ID_TGQ
	// AVCodecIdTqi wraps AV_CODEC_ID_TQI.
	AVCodecIdTqi AVCodecID = C.AV_CODEC_ID_TQI
	// AVCodecIdAura wraps AV_CODEC_ID_AURA.
	AVCodecIdAura AVCodecID = C.AV_CODEC_ID_AURA
	// AVCodecIdAura2 wraps AV_CODEC_ID_AURA2.
	AVCodecIdAura2 AVCodecID = C.AV_CODEC_ID_AURA2
	// AVCodecIdV210X wraps AV_CODEC_ID_V210X.
	AVCodecIdV210X AVCodecID = C.AV_CODEC_ID_V210X
	// AVCodecIdTmv wraps AV_CODEC_ID_TMV.
	AVCodecIdTmv AVCodecID = C.AV_CODEC_ID_TMV
	// AVCodecIdV210 wraps AV_CODEC_ID_V210.
	AVCodecIdV210 AVCodecID = C.AV_CODEC_ID_V210
	// AVCodecIdDpx wraps AV_CODEC_ID_DPX.
	AVCodecIdDpx AVCodecID = C.AV_CODEC_ID_DPX
	// AVCodecIdMad wraps AV_CODEC_ID_MAD.
	AVCodecIdMad AVCodecID = C.AV_CODEC_ID_MAD
	// AVCodecIdFrwu wraps AV_CODEC_ID_FRWU.
	AVCodecIdFrwu AVCodecID = C.AV_CODEC_ID_FRWU
	// AVCodecIdFlashsv2 wraps AV_CODEC_ID_FLASHSV2.
	AVCodecIdFlashsv2 AVCodecID = C.AV_CODEC_ID_FLASHSV2
	// AVCodecIdCdgraphics wraps AV_CODEC_ID_CDGRAPHICS.
	AVCodecIdCdgraphics AVCodecID = C.AV_CODEC_ID_CDGRAPHICS
	// AVCodecIdR210 wraps AV_CODEC_ID_R210.
	AVCodecIdR210 AVCodecID = C.AV_CODEC_ID_R210
	// AVCodecIdAnm wraps AV_CODEC_ID_ANM.
	AVCodecIdAnm AVCodecID = C.AV_CODEC_ID_ANM
	// AVCodecIdBinkvideo wraps AV_CODEC_ID_BINKVIDEO.
	AVCodecIdBinkvideo AVCodecID = C.AV_CODEC_ID_BINKVIDEO
	// AVCodecIdIffIlbm wraps AV_CODEC_ID_IFF_ILBM.
	AVCodecIdIffIlbm AVCodecID = C.AV_CODEC_ID_IFF_ILBM
	// AVCodecIdKgv1 wraps AV_CODEC_ID_KGV1.
	AVCodecIdKgv1 AVCodecID = C.AV_CODEC_ID_KGV1
	// AVCodecIdYop wraps AV_CODEC_ID_YOP.
	AVCodecIdYop AVCodecID = C.AV_CODEC_ID_YOP
	// AVCodecIdVp8 wraps AV_CODEC_ID_VP8.
	AVCodecIdVp8 AVCodecID = C.AV_CODEC_ID_VP8
	// AVCodecIdPictor wraps AV_CODEC_ID_PICTOR.
	AVCodecIdPictor AVCodecID = C.AV_CODEC_ID_PICTOR
	// AVCodecIdAnsi wraps AV_CODEC_ID_ANSI.
	AVCodecIdAnsi AVCodecID = C.AV_CODEC_ID_ANSI
	// AVCodecIdA64Multi wraps AV_CODEC_ID_A64_MULTI.
	AVCodecIdA64Multi AVCodecID = C.AV_CODEC_ID_A64_MULTI
	// AVCodecIdA64Multi5 wraps AV_CODEC_ID_A64_MULTI5.
	AVCodecIdA64Multi5 AVCodecID = C.AV_CODEC_ID_A64_MULTI5
	// AVCodecIdR10K wraps AV_CODEC_ID_R10K.
	AVCodecIdR10K AVCodecID = C.AV_CODEC_ID_R10K
	// AVCodecIdMxpeg wraps AV_CODEC_ID_MXPEG.
	AVCodecIdMxpeg AVCodecID = C.AV_CODEC_ID_MXPEG
	// AVCodecIdLagarith wraps AV_CODEC_ID_LAGARITH.
	AVCodecIdLagarith AVCodecID = C.AV_CODEC_ID_LAGARITH
	// AVCodecIdProres wraps AV_CODEC_ID_PRORES.
	AVCodecIdProres AVCodecID = C.AV_CODEC_ID_PRORES
	// AVCodecIdJv wraps AV_CODEC_ID_JV.
	AVCodecIdJv AVCodecID = C.AV_CODEC_ID_JV
	// AVCodecIdDfa wraps AV_CODEC_ID_DFA.
	AVCodecIdDfa AVCodecID = C.AV_CODEC_ID_DFA
	// AVCodecIdWmv3Image wraps AV_CODEC_ID_WMV3IMAGE.
	AVCodecIdWmv3Image AVCodecID = C.AV_CODEC_ID_WMV3IMAGE
	// AVCodecIdVc1Image wraps AV_CODEC_ID_VC1IMAGE.
	AVCodecIdVc1Image AVCodecID = C.AV_CODEC_ID_VC1IMAGE
	// AVCodecIdUtvideo wraps AV_CODEC_ID_UTVIDEO.
	AVCodecIdUtvideo AVCodecID = C.AV_CODEC_ID_UTVIDEO
	// AVCodecIdBmvVideo wraps AV_CODEC_ID_BMV_VIDEO.
	AVCodecIdBmvVideo AVCodecID = C.AV_CODEC_ID_BMV_VIDEO
	// AVCodecIdVble wraps AV_CODEC_ID_VBLE.
	AVCodecIdVble AVCodecID = C.AV_CODEC_ID_VBLE
	// AVCodecIdDxtory wraps AV_CODEC_ID_DXTORY.
	AVCodecIdDxtory AVCodecID = C.AV_CODEC_ID_DXTORY
	// AVCodecIdV410 wraps AV_CODEC_ID_V410.
	AVCodecIdV410 AVCodecID = C.AV_CODEC_ID_V410
	// AVCodecIdXwd wraps AV_CODEC_ID_XWD.
	AVCodecIdXwd AVCodecID = C.AV_CODEC_ID_XWD
	// AVCodecIdCdxl wraps AV_CODEC_ID_CDXL.
	AVCodecIdCdxl AVCodecID = C.AV_CODEC_ID_CDXL
	// AVCodecIdXbm wraps AV_CODEC_ID_XBM.
	AVCodecIdXbm AVCodecID = C.AV_CODEC_ID_XBM
	// AVCodecIdZerocodec wraps AV_CODEC_ID_ZEROCODEC.
	AVCodecIdZerocodec AVCodecID = C.AV_CODEC_ID_ZEROCODEC
	// AVCodecIdMss1 wraps AV_CODEC_ID_MSS1.
	AVCodecIdMss1 AVCodecID = C.AV_CODEC_ID_MSS1
	// AVCodecIdMsa1 wraps AV_CODEC_ID_MSA1.
	AVCodecIdMsa1 AVCodecID = C.AV_CODEC_ID_MSA1
	// AVCodecIdTscc2 wraps AV_CODEC_ID_TSCC2.
	AVCodecIdTscc2 AVCodecID = C.AV_CODEC_ID_TSCC2
	// AVCodecIdMts2 wraps AV_CODEC_ID_MTS2.
	AVCodecIdMts2 AVCodecID = C.AV_CODEC_ID_MTS2
	// AVCodecIdCllc wraps AV_CODEC_ID_CLLC.
	AVCodecIdCllc AVCodecID = C.AV_CODEC_ID_CLLC
	// AVCodecIdMss2 wraps AV_CODEC_ID_MSS2.
	AVCodecIdMss2 AVCodecID = C.AV_CODEC_ID_MSS2
	// AVCodecIdVp9 wraps AV_CODEC_ID_VP9.
	AVCodecIdVp9 AVCodecID = C.AV_CODEC_ID_VP9
	// AVCodecIdAic wraps AV_CODEC_ID_AIC.
	AVCodecIdAic AVCodecID = C.AV_CODEC_ID_AIC
	// AVCodecIdEscape130 wraps AV_CODEC_ID_ESCAPE130.
	AVCodecIdEscape130 AVCodecID = C.AV_CODEC_ID_ESCAPE130
	// AVCodecIdG2M wraps AV_CODEC_ID_G2M.
	AVCodecIdG2M AVCodecID = C.AV_CODEC_ID_G2M
	// AVCodecIdWebp wraps AV_CODEC_ID_WEBP.
	AVCodecIdWebp AVCodecID = C.AV_CODEC_ID_WEBP
	// AVCodecIdHnm4Video wraps AV_CODEC_ID_HNM4_VIDEO.
	AVCodecIdHnm4Video AVCodecID = C.AV_CODEC_ID_HNM4_VIDEO
	// AVCodecIdHevc wraps AV_CODEC_ID_HEVC.
	AVCodecIdHevc AVCodecID = C.AV_CODEC_ID_HEVC
	// AVCodecIdFic wraps AV_CODEC_ID_FIC.
	AVCodecIdFic AVCodecID = C.AV_CODEC_ID_FIC
	// AVCodecIdAliasPix wraps AV_CODEC_ID_ALIAS_PIX.
	AVCodecIdAliasPix AVCodecID = C.AV_CODEC_ID_ALIAS_PIX
	// AVCodecIdBrenderPix wraps AV_CODEC_ID_BRENDER_PIX.
	AVCodecIdBrenderPix AVCodecID = C.AV_CODEC_ID_BRENDER_PIX
	// AVCodecIdPafVideo wraps AV_CODEC_ID_PAF_VIDEO.
	AVCodecIdPafVideo AVCodecID = C.AV_CODEC_ID_PAF_VIDEO
	// AVCodecIdExr wraps AV_CODEC_ID_EXR.
	AVCodecIdExr AVCodecID = C.AV_CODEC_ID_EXR
	// AVCodecIdVp7 wraps AV_CODEC_ID_VP7.
	AVCodecIdVp7 AVCodecID = C.AV_CODEC_ID_VP7
	// AVCodecIdSanm wraps AV_CODEC_ID_SANM.
	AVCodecIdSanm AVCodecID = C.AV_CODEC_ID_SANM
	// AVCodecIdSgirle wraps AV_CODEC_ID_SGIRLE.
	AVCodecIdSgirle AVCodecID = C.AV_CODEC_ID_SGIRLE
	// AVCodecIdMvc1 wraps AV_CODEC_ID_MVC1.
	AVCodecIdMvc1 AVCodecID = C.AV_CODEC_ID_MVC1
	// AVCodecIdMvc2 wraps AV_CODEC_ID_MVC2.
	AVCodecIdMvc2 AVCodecID = C.AV_CODEC_ID_MVC2
	// AVCodecIdHqx wraps AV_CODEC_ID_HQX.
	AVCodecIdHqx AVCodecID = C.AV_CODEC_ID_HQX
	// AVCodecIdTdsc wraps AV_CODEC_ID_TDSC.
	AVCodecIdTdsc AVCodecID = C.AV_CODEC_ID_TDSC
	// AVCodecIdHqHqa wraps AV_CODEC_ID_HQ_HQA.
	AVCodecIdHqHqa AVCodecID = C.AV_CODEC_ID_HQ_HQA
	// AVCodecIdHap wraps AV_CODEC_ID_HAP.
	AVCodecIdHap AVCodecID = C.AV_CODEC_ID_HAP
	// AVCodecIdDds wraps AV_CODEC_ID_DDS.
	AVCodecIdDds AVCodecID = C.AV_CODEC_ID_DDS
	// AVCodecIdDxv wraps AV_CODEC_ID_DXV.
	AVCodecIdDxv AVCodecID = C.AV_CODEC_ID_DXV
	// AVCodecIdScreenpresso wraps AV_CODEC_ID_SCREENPRESSO.
	AVCodecIdScreenpresso AVCodecID = C.AV_CODEC_ID_SCREENPRESSO
	// AVCodecIdRscc wraps AV_CODEC_ID_RSCC.
	AVCodecIdRscc AVCodecID = C.AV_CODEC_ID_RSCC
	// AVCodecIdAvs2 wraps AV_CODEC_ID_AVS2.
	AVCodecIdAvs2 AVCodecID = C.AV_CODEC_ID_AVS2
	// AVCodecIdPgx wraps AV_CODEC_ID_PGX.
	AVCodecIdPgx AVCodecID = C.AV_CODEC_ID_PGX
	// AVCodecIdAvs3 wraps AV_CODEC_ID_AVS3.
	AVCodecIdAvs3 AVCodecID = C.AV_CODEC_ID_AVS3
	// AVCodecIdMsp2 wraps AV_CODEC_ID_MSP2.
	AVCodecIdMsp2 AVCodecID = C.AV_CODEC_ID_MSP2
	// AVCodecIdVvc wraps AV_CODEC_ID_VVC.
	AVCodecIdVvc AVCodecID = C.AV_CODEC_ID_VVC
	// AVCodecIdY41P wraps AV_CODEC_ID_Y41P.
	AVCodecIdY41P AVCodecID = C.AV_CODEC_ID_Y41P
	// AVCodecIdAvrp wraps AV_CODEC_ID_AVRP.
	AVCodecIdAvrp AVCodecID = C.AV_CODEC_ID_AVRP
	// AVCodecId012V wraps AV_CODEC_ID_012V.
	AVCodecId012V AVCodecID = C.AV_CODEC_ID_012V
	// AVCodecIdAvui wraps AV_CODEC_ID_AVUI.
	AVCodecIdAvui AVCodecID = C.AV_CODEC_ID_AVUI
	// AVCodecIdAyuv wraps AV_CODEC_ID_AYUV.
	AVCodecIdAyuv AVCodecID = C.AV_CODEC_ID_AYUV
	// AVCodecIdTargaY216 wraps AV_CODEC_ID_TARGA_Y216.
	AVCodecIdTargaY216 AVCodecID = C.AV_CODEC_ID_TARGA_Y216
	// AVCodecIdV308 wraps AV_CODEC_ID_V308.
	AVCodecIdV308 AVCodecID = C.AV_CODEC_ID_V308
	// AVCodecIdV408 wraps AV_CODEC_ID_V408.
	AVCodecIdV408 AVCodecID = C.AV_CODEC_ID_V408
	// AVCodecIdYuv4 wraps AV_CODEC_ID_YUV4.
	AVCodecIdYuv4 AVCodecID = C.AV_CODEC_ID_YUV4
	// AVCodecIdAvrn wraps AV_CODEC_ID_AVRN.
	AVCodecIdAvrn AVCodecID = C.AV_CODEC_ID_AVRN
	// AVCodecIdCpia wraps AV_CODEC_ID_CPIA.
	AVCodecIdCpia AVCodecID = C.AV_CODEC_ID_CPIA
	// AVCodecIdXface wraps AV_CODEC_ID_XFACE.
	AVCodecIdXface AVCodecID = C.AV_CODEC_ID_XFACE
	// AVCodecIdSnow wraps AV_CODEC_ID_SNOW.
	AVCodecIdSnow AVCodecID = C.AV_CODEC_ID_SNOW
	// AVCodecIdSmvjpeg wraps AV_CODEC_ID_SMVJPEG.
	AVCodecIdSmvjpeg AVCodecID = C.AV_CODEC_ID_SMVJPEG
	// AVCodecIdApng wraps AV_CODEC_ID_APNG.
	AVCodecIdApng AVCodecID = C.AV_CODEC_ID_APNG
	// AVCodecIdDaala wraps AV_CODEC_ID_DAALA.
	AVCodecIdDaala AVCodecID = C.AV_CODEC_ID_DAALA
	// AVCodecIdCfhd wraps AV_CODEC_ID_CFHD.
	AVCodecIdCfhd AVCodecID = C.AV_CODEC_ID_CFHD
	// AVCodecIdTruemotion2Rt wraps AV_CODEC_ID_TRUEMOTION2RT.
	AVCodecIdTruemotion2Rt AVCodecID = C.AV_CODEC_ID_TRUEMOTION2RT
	// AVCodecIdM101 wraps AV_CODEC_ID_M101.
	AVCodecIdM101 AVCodecID = C.AV_CODEC_ID_M101
	// AVCodecIdMagicyuv wraps AV_CODEC_ID_MAGICYUV.
	AVCodecIdMagicyuv AVCodecID = C.AV_CODEC_ID_MAGICYUV
	// AVCodecIdSheervideo wraps AV_CODEC_ID_SHEERVIDEO.
	AVCodecIdSheervideo AVCodecID = C.AV_CODEC_ID_SHEERVIDEO
	// AVCodecIdYlc wraps AV_CODEC_ID_YLC.
	AVCodecIdYlc AVCodecID = C.AV_CODEC_ID_YLC
	// AVCodecIdPsd wraps AV_CODEC_ID_PSD.
	AVCodecIdPsd AVCodecID = C.AV_CODEC_ID_PSD
	// AVCodecIdPixlet wraps AV_CODEC_ID_PIXLET.
	AVCodecIdPixlet AVCodecID = C.AV_CODEC_ID_PIXLET
	// AVCodecIdSpeedhq wraps AV_CODEC_ID_SPEEDHQ.
	AVCodecIdSpeedhq AVCodecID = C.AV_CODEC_ID_SPEEDHQ
	// AVCodecIdFmvc wraps AV_CODEC_ID_FMVC.
	AVCodecIdFmvc AVCodecID = C.AV_CODEC_ID_FMVC
	// AVCodecIdScpr wraps AV_CODEC_ID_SCPR.
	AVCodecIdScpr AVCodecID = C.AV_CODEC_ID_SCPR
	// AVCodecIdClearvideo wraps AV_CODEC_ID_CLEARVIDEO.
	AVCodecIdClearvideo AVCodecID = C.AV_CODEC_ID_CLEARVIDEO
	// AVCodecIdXpm wraps AV_CODEC_ID_XPM.
	AVCodecIdXpm AVCodecID = C.AV_CODEC_ID_XPM
	// AVCodecIdAv1 wraps AV_CODEC_ID_AV1.
	AVCodecIdAv1 AVCodecID = C.AV_CODEC_ID_AV1
	// AVCodecIdBitpacked wraps AV_CODEC_ID_BITPACKED.
	AVCodecIdBitpacked AVCodecID = C.AV_CODEC_ID_BITPACKED
	// AVCodecIdMscc wraps AV_CODEC_ID_MSCC.
	AVCodecIdMscc AVCodecID = C.AV_CODEC_ID_MSCC
	// AVCodecIdSrgc wraps AV_CODEC_ID_SRGC.
	AVCodecIdSrgc AVCodecID = C.AV_CODEC_ID_SRGC
	// AVCodecIdSvg wraps AV_CODEC_ID_SVG.
	AVCodecIdSvg AVCodecID = C.AV_CODEC_ID_SVG
	// AVCodecIdGdv wraps AV_CODEC_ID_GDV.
	AVCodecIdGdv AVCodecID = C.AV_CODEC_ID_GDV
	// AVCodecIdFits wraps AV_CODEC_ID_FITS.
	AVCodecIdFits AVCodecID = C.AV_CODEC_ID_FITS
	// AVCodecIdImm4 wraps AV_CODEC_ID_IMM4.
	AVCodecIdImm4 AVCodecID = C.AV_CODEC_ID_IMM4
	// AVCodecIdProsumer wraps AV_CODEC_ID_PROSUMER.
	AVCodecIdProsumer AVCodecID = C.AV_CODEC_ID_PROSUMER
	// AVCodecIdMwsc wraps AV_CODEC_ID_MWSC.
	AVCodecIdMwsc AVCodecID = C.AV_CODEC_ID_MWSC
	// AVCodecIdWcmv wraps AV_CODEC_ID_WCMV.
	AVCodecIdWcmv AVCodecID = C.AV_CODEC_ID_WCMV
	// AVCodecIdRasc wraps AV_CODEC_ID_RASC.
	AVCodecIdRasc AVCodecID = C.AV_CODEC_ID_RASC
	// AVCodecIdHymt wraps AV_CODEC_ID_HYMT.
	AVCodecIdHymt AVCodecID = C.AV_CODEC_ID_HYMT
	// AVCodecIdArbc wraps AV_CODEC_ID_ARBC.
	AVCodecIdArbc AVCodecID = C.AV_CODEC_ID_ARBC
	// AVCodecIdAgm wraps AV_CODEC_ID_AGM.
	AVCodecIdAgm AVCodecID = C.AV_CODEC_ID_AGM
	// AVCodecIdLscr wraps AV_CODEC_ID_LSCR.
	AVCodecIdLscr AVCodecID = C.AV_CODEC_ID_LSCR
	// AVCodecIdVp4 wraps AV_CODEC_ID_VP4.
	AVCodecIdVp4 AVCodecID = C.AV_CODEC_ID_VP4
	// AVCodecIdImm5 wraps AV_CODEC_ID_IMM5.
	AVCodecIdImm5 AVCodecID = C.AV_CODEC_ID_IMM5
	// AVCodecIdMvdv wraps AV_CODEC_ID_MVDV.
	AVCodecIdMvdv AVCodecID = C.AV_CODEC_ID_MVDV
	// AVCodecIdMvha wraps AV_CODEC_ID_MVHA.
	AVCodecIdMvha AVCodecID = C.AV_CODEC_ID_MVHA
	// AVCodecIdCdtoons wraps AV_CODEC_ID_CDTOONS.
	AVCodecIdCdtoons AVCodecID = C.AV_CODEC_ID_CDTOONS
	// AVCodecIdMv30 wraps AV_CODEC_ID_MV30.
	AVCodecIdMv30 AVCodecID = C.AV_CODEC_ID_MV30
	// AVCodecIdNotchlc wraps AV_CODEC_ID_NOTCHLC.
	AVCodecIdNotchlc AVCodecID = C.AV_CODEC_ID_NOTCHLC
	// AVCodecIdPfm wraps AV_CODEC_ID_PFM.
	AVCodecIdPfm AVCodecID = C.AV_CODEC_ID_PFM
	// AVCodecIdMobiclip wraps AV_CODEC_ID_MOBICLIP.
	AVCodecIdMobiclip AVCodecID = C.AV_CODEC_ID_MOBICLIP
	// AVCodecIdPhotocd wraps AV_CODEC_ID_PHOTOCD.
	AVCodecIdPhotocd AVCodecID = C.AV_CODEC_ID_PHOTOCD
	// AVCodecIdIpu wraps AV_CODEC_ID_IPU.
	AVCodecIdIpu AVCodecID = C.AV_CODEC_ID_IPU
	// AVCodecIdArgo wraps AV_CODEC_ID_ARGO.
	AVCodecIdArgo AVCodecID = C.AV_CODEC_ID_ARGO
	// AVCodecIdCri wraps AV_CODEC_ID_CRI.
	AVCodecIdCri AVCodecID = C.AV_CODEC_ID_CRI
	// AVCodecIdSimbiosisImx wraps AV_CODEC_ID_SIMBIOSIS_IMX.
	AVCodecIdSimbiosisImx AVCodecID = C.AV_CODEC_ID_SIMBIOSIS_IMX
	// AVCodecIdSgaVideo wraps AV_CODEC_ID_SGA_VIDEO.
	AVCodecIdSgaVideo AVCodecID = C.AV_CODEC_ID_SGA_VIDEO
	// AVCodecIdGem wraps AV_CODEC_ID_GEM.
	AVCodecIdGem AVCodecID = C.AV_CODEC_ID_GEM
	// AVCodecIdVbn wraps AV_CODEC_ID_VBN.
	AVCodecIdVbn AVCodecID = C.AV_CODEC_ID_VBN
	// AVCodecIdJpegxl wraps AV_CODEC_ID_JPEGXL.
	AVCodecIdJpegxl AVCodecID = C.AV_CODEC_ID_JPEGXL
	// AVCodecIdQoi wraps AV_CODEC_ID_QOI.
	AVCodecIdQoi AVCodecID = C.AV_CODEC_ID_QOI
	// AVCodecIdPhm wraps AV_CODEC_ID_PHM.
	AVCodecIdPhm AVCodecID = C.AV_CODEC_ID_PHM
	// AVCodecIdRadianceHdr wraps AV_CODEC_ID_RADIANCE_HDR.
	AVCodecIdRadianceHdr AVCodecID = C.AV_CODEC_ID_RADIANCE_HDR
	// AVCodecIdWbmp wraps AV_CODEC_ID_WBMP.
	AVCodecIdWbmp AVCodecID = C.AV_CODEC_ID_WBMP
	// AVCodecIdMedia100 wraps AV_CODEC_ID_MEDIA100.
	AVCodecIdMedia100 AVCodecID = C.AV_CODEC_ID_MEDIA100
	// AVCodecIdVqc wraps AV_CODEC_ID_VQC.
	AVCodecIdVqc AVCodecID = C.AV_CODEC_ID_VQC
	// AVCodecIdFirstAudio wraps AV_CODEC_ID_FIRST_AUDIO.
	//
	//	A dummy id pointing at the start of audio codecs
	AVCodecIdFirstAudio AVCodecID = C.AV_CODEC_ID_FIRST_AUDIO
	// AVCodecIdPcmS16Le wraps AV_CODEC_ID_PCM_S16LE.
	AVCodecIdPcmS16Le AVCodecID = C.AV_CODEC_ID_PCM_S16LE
	// AVCodecIdPcmS16Be wraps AV_CODEC_ID_PCM_S16BE.
	AVCodecIdPcmS16Be AVCodecID = C.AV_CODEC_ID_PCM_S16BE
	// AVCodecIdPcmU16Le wraps AV_CODEC_ID_PCM_U16LE.
	AVCodecIdPcmU16Le AVCodecID = C.AV_CODEC_ID_PCM_U16LE
	// AVCodecIdPcmU16Be wraps AV_CODEC_ID_PCM_U16BE.
	AVCodecIdPcmU16Be AVCodecID = C.AV_CODEC_ID_PCM_U16BE
	// AVCodecIdPcmS8 wraps AV_CODEC_ID_PCM_S8.
	AVCodecIdPcmS8 AVCodecID = C.AV_CODEC_ID_PCM_S8
	// AVCodecIdPcmU8 wraps AV_CODEC_ID_PCM_U8.
	AVCodecIdPcmU8 AVCodecID = C.AV_CODEC_ID_PCM_U8
	// AVCodecIdPcmMulaw wraps AV_CODEC_ID_PCM_MULAW.
	AVCodecIdPcmMulaw AVCodecID = C.AV_CODEC_ID_PCM_MULAW
	// AVCodecIdPcmAlaw wraps AV_CODEC_ID_PCM_ALAW.
	AVCodecIdPcmAlaw AVCodecID = C.AV_CODEC_ID_PCM_ALAW
	// AVCodecIdPcmS32Le wraps AV_CODEC_ID_PCM_S32LE.
	AVCodecIdPcmS32Le AVCodecID = C.AV_CODEC_ID_PCM_S32LE
	// AVCodecIdPcmS32Be wraps AV_CODEC_ID_PCM_S32BE.
	AVCodecIdPcmS32Be AVCodecID = C.AV_CODEC_ID_PCM_S32BE
	// AVCodecIdPcmU32Le wraps AV_CODEC_ID_PCM_U32LE.
	AVCodecIdPcmU32Le AVCodecID = C.AV_CODEC_ID_PCM_U32LE
	// AVCodecIdPcmU32Be wraps AV_CODEC_ID_PCM_U32BE.
	AVCodecIdPcmU32Be AVCodecID = C.AV_CODEC_ID_PCM_U32BE
	// AVCodecIdPcmS24Le wraps AV_CODEC_ID_PCM_S24LE.
	AVCodecIdPcmS24Le AVCodecID = C.AV_CODEC_ID_PCM_S24LE
	// AVCodecIdPcmS24Be wraps AV_CODEC_ID_PCM_S24BE.
	AVCodecIdPcmS24Be AVCodecID = C.AV_CODEC_ID_PCM_S24BE
	// AVCodecIdPcmU24Le wraps AV_CODEC_ID_PCM_U24LE.
	AVCodecIdPcmU24Le AVCodecID = C.AV_CODEC_ID_PCM_U24LE
	// AVCodecIdPcmU24Be wraps AV_CODEC_ID_PCM_U24BE.
	AVCodecIdPcmU24Be AVCodecID = C.AV_CODEC_ID_PCM_U24BE
	// AVCodecIdPcmS24Daud wraps AV_CODEC_ID_PCM_S24DAUD.
	AVCodecIdPcmS24Daud AVCodecID = C.AV_CODEC_ID_PCM_S24DAUD
	// AVCodecIdPcmZork wraps AV_CODEC_ID_PCM_ZORK.
	AVCodecIdPcmZork AVCodecID = C.AV_CODEC_ID_PCM_ZORK
	// AVCodecIdPcmS16LePlanar wraps AV_CODEC_ID_PCM_S16LE_PLANAR.
	AVCodecIdPcmS16LePlanar AVCodecID = C.AV_CODEC_ID_PCM_S16LE_PLANAR
	// AVCodecIdPcmDvd wraps AV_CODEC_ID_PCM_DVD.
	AVCodecIdPcmDvd AVCodecID = C.AV_CODEC_ID_PCM_DVD
	// AVCodecIdPcmF32Be wraps AV_CODEC_ID_PCM_F32BE.
	AVCodecIdPcmF32Be AVCodecID = C.AV_CODEC_ID_PCM_F32BE
	// AVCodecIdPcmF32Le wraps AV_CODEC_ID_PCM_F32LE.
	AVCodecIdPcmF32Le AVCodecID = C.AV_CODEC_ID_PCM_F32LE
	// AVCodecIdPcmF64Be wraps AV_CODEC_ID_PCM_F64BE.
	AVCodecIdPcmF64Be AVCodecID = C.AV_CODEC_ID_PCM_F64BE
	// AVCodecIdPcmF64Le wraps AV_CODEC_ID_PCM_F64LE.
	AVCodecIdPcmF64Le AVCodecID = C.AV_CODEC_ID_PCM_F64LE
	// AVCodecIdPcmBluray wraps AV_CODEC_ID_PCM_BLURAY.
	AVCodecIdPcmBluray AVCodecID = C.AV_CODEC_ID_PCM_BLURAY
	// AVCodecIdPcmLxf wraps AV_CODEC_ID_PCM_LXF.
	AVCodecIdPcmLxf AVCodecID = C.AV_CODEC_ID_PCM_LXF
	// AVCodecIdS302M wraps AV_CODEC_ID_S302M.
	AVCodecIdS302M AVCodecID = C.AV_CODEC_ID_S302M
	// AVCodecIdPcmS8Planar wraps AV_CODEC_ID_PCM_S8_PLANAR.
	AVCodecIdPcmS8Planar AVCodecID = C.AV_CODEC_ID_PCM_S8_PLANAR
	// AVCodecIdPcmS24LePlanar wraps AV_CODEC_ID_PCM_S24LE_PLANAR.
	AVCodecIdPcmS24LePlanar AVCodecID = C.AV_CODEC_ID_PCM_S24LE_PLANAR
	// AVCodecIdPcmS32LePlanar wraps AV_CODEC_ID_PCM_S32LE_PLANAR.
	AVCodecIdPcmS32LePlanar AVCodecID = C.AV_CODEC_ID_PCM_S32LE_PLANAR
	// AVCodecIdPcmS16BePlanar wraps AV_CODEC_ID_PCM_S16BE_PLANAR.
	AVCodecIdPcmS16BePlanar AVCodecID = C.AV_CODEC_ID_PCM_S16BE_PLANAR
	// AVCodecIdPcmS64Le wraps AV_CODEC_ID_PCM_S64LE.
	AVCodecIdPcmS64Le AVCodecID = C.AV_CODEC_ID_PCM_S64LE
	// AVCodecIdPcmS64Be wraps AV_CODEC_ID_PCM_S64BE.
	AVCodecIdPcmS64Be AVCodecID = C.AV_CODEC_ID_PCM_S64BE
	// AVCodecIdPcmF16Le wraps AV_CODEC_ID_PCM_F16LE.
	AVCodecIdPcmF16Le AVCodecID = C.AV_CODEC_ID_PCM_F16LE
	// AVCodecIdPcmF24Le wraps AV_CODEC_ID_PCM_F24LE.
	AVCodecIdPcmF24Le AVCodecID = C.AV_CODEC_ID_PCM_F24LE
	// AVCodecIdPcmVidc wraps AV_CODEC_ID_PCM_VIDC.
	AVCodecIdPcmVidc AVCodecID = C.AV_CODEC_ID_PCM_VIDC
	// AVCodecIdPcmSga wraps AV_CODEC_ID_PCM_SGA.
	AVCodecIdPcmSga AVCodecID = C.AV_CODEC_ID_PCM_SGA
	// AVCodecIdAdpcmImaQt wraps AV_CODEC_ID_ADPCM_IMA_QT.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaQt AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_QT
	// AVCodecIdAdpcmImaWav wraps AV_CODEC_ID_ADPCM_IMA_WAV.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaWav AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_WAV
	// AVCodecIdAdpcmImaDk3 wraps AV_CODEC_ID_ADPCM_IMA_DK3.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaDk3 AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_DK3
	// AVCodecIdAdpcmImaDk4 wraps AV_CODEC_ID_ADPCM_IMA_DK4.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaDk4 AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_DK4
	// AVCodecIdAdpcmImaWs wraps AV_CODEC_ID_ADPCM_IMA_WS.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaWs AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_WS
	// AVCodecIdAdpcmImaSmjpeg wraps AV_CODEC_ID_ADPCM_IMA_SMJPEG.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaSmjpeg AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_SMJPEG
	// AVCodecIdAdpcmMs wraps AV_CODEC_ID_ADPCM_MS.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmMs AVCodecID = C.AV_CODEC_ID_ADPCM_MS
	// AVCodecIdAdpcm4Xm wraps AV_CODEC_ID_ADPCM_4XM.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcm4Xm AVCodecID = C.AV_CODEC_ID_ADPCM_4XM
	// AVCodecIdAdpcmXa wraps AV_CODEC_ID_ADPCM_XA.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmXa AVCodecID = C.AV_CODEC_ID_ADPCM_XA
	// AVCodecIdAdpcmAdx wraps AV_CODEC_ID_ADPCM_ADX.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmAdx AVCodecID = C.AV_CODEC_ID_ADPCM_ADX
	// AVCodecIdAdpcmEa wraps AV_CODEC_ID_ADPCM_EA.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmEa AVCodecID = C.AV_CODEC_ID_ADPCM_EA
	// AVCodecIdAdpcmG726 wraps AV_CODEC_ID_ADPCM_G726.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmG726 AVCodecID = C.AV_CODEC_ID_ADPCM_G726
	// AVCodecIdAdpcmCt wraps AV_CODEC_ID_ADPCM_CT.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmCt AVCodecID = C.AV_CODEC_ID_ADPCM_CT
	// AVCodecIdAdpcmSwf wraps AV_CODEC_ID_ADPCM_SWF.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmSwf AVCodecID = C.AV_CODEC_ID_ADPCM_SWF
	// AVCodecIdAdpcmYamaha wraps AV_CODEC_ID_ADPCM_YAMAHA.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmYamaha AVCodecID = C.AV_CODEC_ID_ADPCM_YAMAHA
	// AVCodecIdAdpcmSbpro4 wraps AV_CODEC_ID_ADPCM_SBPRO_4.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmSbpro4 AVCodecID = C.AV_CODEC_ID_ADPCM_SBPRO_4
	// AVCodecIdAdpcmSbpro3 wraps AV_CODEC_ID_ADPCM_SBPRO_3.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmSbpro3 AVCodecID = C.AV_CODEC_ID_ADPCM_SBPRO_3
	// AVCodecIdAdpcmSbpro2 wraps AV_CODEC_ID_ADPCM_SBPRO_2.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmSbpro2 AVCodecID = C.AV_CODEC_ID_ADPCM_SBPRO_2
	// AVCodecIdAdpcmThp wraps AV_CODEC_ID_ADPCM_THP.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmThp AVCodecID = C.AV_CODEC_ID_ADPCM_THP
	// AVCodecIdAdpcmImaAmv wraps AV_CODEC_ID_ADPCM_IMA_AMV.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaAmv AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_AMV
	// AVCodecIdAdpcmEaR1 wraps AV_CODEC_ID_ADPCM_EA_R1.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmEaR1 AVCodecID = C.AV_CODEC_ID_ADPCM_EA_R1
	// AVCodecIdAdpcmEaR3 wraps AV_CODEC_ID_ADPCM_EA_R3.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmEaR3 AVCodecID = C.AV_CODEC_ID_ADPCM_EA_R3
	// AVCodecIdAdpcmEaR2 wraps AV_CODEC_ID_ADPCM_EA_R2.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmEaR2 AVCodecID = C.AV_CODEC_ID_ADPCM_EA_R2
	// AVCodecIdAdpcmImaEaSead wraps AV_CODEC_ID_ADPCM_IMA_EA_SEAD.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaEaSead AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD
	// AVCodecIdAdpcmImaEaEacs wraps AV_CODEC_ID_ADPCM_IMA_EA_EACS.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaEaEacs AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_EA_EACS
	// AVCodecIdAdpcmEaXas wraps AV_CODEC_ID_ADPCM_EA_XAS.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmEaXas AVCodecID = C.AV_CODEC_ID_ADPCM_EA_XAS
	// AVCodecIdAdpcmEaMaxisXa wraps AV_CODEC_ID_ADPCM_EA_MAXIS_XA.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmEaMaxisXa AVCodecID = C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA
	// AVCodecIdAdpcmImaIss wraps AV_CODEC_ID_ADPCM_IMA_ISS.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaIss AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_ISS
	// AVCodecIdAdpcmG722 wraps AV_CODEC_ID_ADPCM_G722.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmG722 AVCodecID = C.AV_CODEC_ID_ADPCM_G722
	// AVCodecIdAdpcmImaApc wraps AV_CODEC_ID_ADPCM_IMA_APC.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaApc AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_APC
	// AVCodecIdAdpcmVima wraps AV_CODEC_ID_ADPCM_VIMA.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmVima AVCodecID = C.AV_CODEC_ID_ADPCM_VIMA
	// AVCodecIdAdpcmAfc wraps AV_CODEC_ID_ADPCM_AFC.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmAfc AVCodecID = C.AV_CODEC_ID_ADPCM_AFC
	// AVCodecIdAdpcmImaOki wraps AV_CODEC_ID_ADPCM_IMA_OKI.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaOki AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_OKI
	// AVCodecIdAdpcmDtk wraps AV_CODEC_ID_ADPCM_DTK.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmDtk AVCodecID = C.AV_CODEC_ID_ADPCM_DTK
	// AVCodecIdAdpcmImaRad wraps AV_CODEC_ID_ADPCM_IMA_RAD.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaRad AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_RAD
	// AVCodecIdAdpcmG726Le wraps AV_CODEC_ID_ADPCM_G726LE.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmG726Le AVCodecID = C.AV_CODEC_ID_ADPCM_G726LE
	// AVCodecIdAdpcmThpLe wraps AV_CODEC_ID_ADPCM_THP_LE.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmThpLe AVCodecID = C.AV_CODEC_ID_ADPCM_THP_LE
	// AVCodecIdAdpcmPsx wraps AV_CODEC_ID_ADPCM_PSX.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmPsx AVCodecID = C.AV_CODEC_ID_ADPCM_PSX
	// AVCodecIdAdpcmAica wraps AV_CODEC_ID_ADPCM_AICA.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmAica AVCodecID = C.AV_CODEC_ID_ADPCM_AICA
	// AVCodecIdAdpcmImaDat4 wraps AV_CODEC_ID_ADPCM_IMA_DAT4.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaDat4 AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_DAT4
	// AVCodecIdAdpcmMtaf wraps AV_CODEC_ID_ADPCM_MTAF.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmMtaf AVCodecID = C.AV_CODEC_ID_ADPCM_MTAF
	// AVCodecIdAdpcmAgm wraps AV_CODEC_ID_ADPCM_AGM.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmAgm AVCodecID = C.AV_CODEC_ID_ADPCM_AGM
	// AVCodecIdAdpcmArgo wraps AV_CODEC_ID_ADPCM_ARGO.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmArgo AVCodecID = C.AV_CODEC_ID_ADPCM_ARGO
	// AVCodecIdAdpcmImaSsi wraps AV_CODEC_ID_ADPCM_IMA_SSI.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaSsi AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_SSI
	// AVCodecIdAdpcmZork wraps AV_CODEC_ID_ADPCM_ZORK.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmZork AVCodecID = C.AV_CODEC_ID_ADPCM_ZORK
	// AVCodecIdAdpcmImaApm wraps AV_CODEC_ID_ADPCM_IMA_APM.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaApm AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_APM
	// AVCodecIdAdpcmImaAlp wraps AV_CODEC_ID_ADPCM_IMA_ALP.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaAlp AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_ALP
	// AVCodecIdAdpcmImaMtf wraps AV_CODEC_ID_ADPCM_IMA_MTF.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaMtf AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_MTF
	// AVCodecIdAdpcmImaCunning wraps AV_CODEC_ID_ADPCM_IMA_CUNNING.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaCunning AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_CUNNING
	// AVCodecIdAdpcmImaMoflex wraps AV_CODEC_ID_ADPCM_IMA_MOFLEX.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaMoflex AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_MOFLEX
	// AVCodecIdAdpcmImaAcorn wraps AV_CODEC_ID_ADPCM_IMA_ACORN.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmImaAcorn AVCodecID = C.AV_CODEC_ID_ADPCM_IMA_ACORN
	// AVCodecIdAdpcmXmd wraps AV_CODEC_ID_ADPCM_XMD.
	//
	//	various ADPCM codecs
	AVCodecIdAdpcmXmd AVCodecID = C.AV_CODEC_ID_ADPCM_XMD
	// AVCodecIdAmrNb wraps AV_CODEC_ID_AMR_NB.
	//
	//	AMR
	AVCodecIdAmrNb AVCodecID = C.AV_CODEC_ID_AMR_NB
	// AVCodecIdAmrWb wraps AV_CODEC_ID_AMR_WB.
	//
	//	AMR
	AVCodecIdAmrWb AVCodecID = C.AV_CODEC_ID_AMR_WB
	// AVCodecIdRa144 wraps AV_CODEC_ID_RA_144.
	//
	//	RealAudio codecs
	AVCodecIdRa144 AVCodecID = C.AV_CODEC_ID_RA_144
	// AVCodecIdRa288 wraps AV_CODEC_ID_RA_288.
	//
	//	RealAudio codecs
	AVCodecIdRa288 AVCodecID = C.AV_CODEC_ID_RA_288
	// AVCodecIdRoqDpcm wraps AV_CODEC_ID_ROQ_DPCM.
	//
	//	various DPCM codecs
	AVCodecIdRoqDpcm AVCodecID = C.AV_CODEC_ID_ROQ_DPCM
	// AVCodecIdInterplayDpcm wraps AV_CODEC_ID_INTERPLAY_DPCM.
	//
	//	various DPCM codecs
	AVCodecIdInterplayDpcm AVCodecID = C.AV_CODEC_ID_INTERPLAY_DPCM
	// AVCodecIdXanDpcm wraps AV_CODEC_ID_XAN_DPCM.
	//
	//	various DPCM codecs
	AVCodecIdXanDpcm AVCodecID = C.AV_CODEC_ID_XAN_DPCM
	// AVCodecIdSolDpcm wraps AV_CODEC_ID_SOL_DPCM.
	//
	//	various DPCM codecs
	AVCodecIdSolDpcm AVCodecID = C.AV_CODEC_ID_SOL_DPCM
	// AVCodecIdSdx2Dpcm wraps AV_CODEC_ID_SDX2_DPCM.
	//
	//	various DPCM codecs
	AVCodecIdSdx2Dpcm AVCodecID = C.AV_CODEC_ID_SDX2_DPCM
	// AVCodecIdGremlinDpcm wraps AV_CODEC_ID_GREMLIN_DPCM.
	//
	//	various DPCM codecs
	AVCodecIdGremlinDpcm AVCodecID = C.AV_CODEC_ID_GREMLIN_DPCM
	// AVCodecIdDerfDpcm wraps AV_CODEC_ID_DERF_DPCM.
	//
	//	various DPCM codecs
	AVCodecIdDerfDpcm AVCodecID = C.AV_CODEC_ID_DERF_DPCM
	// AVCodecIdWadyDpcm wraps AV_CODEC_ID_WADY_DPCM.
	//
	//	various DPCM codecs
	AVCodecIdWadyDpcm AVCodecID = C.AV_CODEC_ID_WADY_DPCM
	// AVCodecIdCbd2Dpcm wraps AV_CODEC_ID_CBD2_DPCM.
	//
	//	various DPCM codecs
	AVCodecIdCbd2Dpcm AVCodecID = C.AV_CODEC_ID_CBD2_DPCM
	// AVCodecIdMp2 wraps AV_CODEC_ID_MP2.
	//
	//	audio codecs
	AVCodecIdMp2 AVCodecID = C.AV_CODEC_ID_MP2
	// AVCodecIdMp3 wraps AV_CODEC_ID_MP3.
	//
	//	preferred ID for decoding MPEG audio layer 1, 2 or 3
	AVCodecIdMp3 AVCodecID = C.AV_CODEC_ID_MP3
	// AVCodecIdAac wraps AV_CODEC_ID_AAC.
	AVCodecIdAac AVCodecID = C.AV_CODEC_ID_AAC
	// AVCodecIdAc3 wraps AV_CODEC_ID_AC3.
	AVCodecIdAc3 AVCodecID = C.AV_CODEC_ID_AC3
	// AVCodecIdDts wraps AV_CODEC_ID_DTS.
	AVCodecIdDts AVCodecID = C.AV_CODEC_ID_DTS
	// AVCodecIdVorbis wraps AV_CODEC_ID_VORBIS.
	AVCodecIdVorbis AVCodecID = C.AV_CODEC_ID_VORBIS
	// AVCodecIdDvaudio wraps AV_CODEC_ID_DVAUDIO.
	AVCodecIdDvaudio AVCodecID = C.AV_CODEC_ID_DVAUDIO
	// AVCodecIdWmav1 wraps AV_CODEC_ID_WMAV1.
	AVCodecIdWmav1 AVCodecID = C.AV_CODEC_ID_WMAV1
	// AVCodecIdWmav2 wraps AV_CODEC_ID_WMAV2.
	AVCodecIdWmav2 AVCodecID = C.AV_CODEC_ID_WMAV2
	// AVCodecIdMace3 wraps AV_CODEC_ID_MACE3.
	AVCodecIdMace3 AVCodecID = C.AV_CODEC_ID_MACE3
	// AVCodecIdMace6 wraps AV_CODEC_ID_MACE6.
	AVCodecIdMace6 AVCodecID = C.AV_CODEC_ID_MACE6
	// AVCodecIdVmdaudio wraps AV_CODEC_ID_VMDAUDIO.
	AVCodecIdVmdaudio AVCodecID = C.AV_CODEC_ID_VMDAUDIO
	// AVCodecIdFlac wraps AV_CODEC_ID_FLAC.
	AVCodecIdFlac AVCodecID = C.AV_CODEC_ID_FLAC
	// AVCodecIdMp3Adu wraps AV_CODEC_ID_MP3ADU.
	AVCodecIdMp3Adu AVCodecID = C.AV_CODEC_ID_MP3ADU
	// AVCodecIdMp3On4 wraps AV_CODEC_ID_MP3ON4.
	AVCodecIdMp3On4 AVCodecID = C.AV_CODEC_ID_MP3ON4
	// AVCodecIdShorten wraps AV_CODEC_ID_SHORTEN.
	AVCodecIdShorten AVCodecID = C.AV_CODEC_ID_SHORTEN
	// AVCodecIdAlac wraps AV_CODEC_ID_ALAC.
	AVCodecIdAlac AVCodecID = C.AV_CODEC_ID_ALAC
	// AVCodecIdWestwoodSnd1 wraps AV_CODEC_ID_WESTWOOD_SND1.
	AVCodecIdWestwoodSnd1 AVCodecID = C.AV_CODEC_ID_WESTWOOD_SND1
	// AVCodecIdGsm wraps AV_CODEC_ID_GSM.
	//
	//	as in Berlin toast format
	AVCodecIdGsm AVCodecID = C.AV_CODEC_ID_GSM
	// AVCodecIdQdm2 wraps AV_CODEC_ID_QDM2.
	AVCodecIdQdm2 AVCodecID = C.AV_CODEC_ID_QDM2
	// AVCodecIdCook wraps AV_CODEC_ID_COOK.
	AVCodecIdCook AVCodecID = C.AV_CODEC_ID_COOK
	// AVCodecIdTruespeech wraps AV_CODEC_ID_TRUESPEECH.
	AVCodecIdTruespeech AVCodecID = C.AV_CODEC_ID_TRUESPEECH
	// AVCodecIdTta wraps AV_CODEC_ID_TTA.
	AVCodecIdTta AVCodecID = C.AV_CODEC_ID_TTA
	// AVCodecIdSmackaudio wraps AV_CODEC_ID_SMACKAUDIO.
	AVCodecIdSmackaudio AVCodecID = C.AV_CODEC_ID_SMACKAUDIO
	// AVCodecIdQcelp wraps AV_CODEC_ID_QCELP.
	AVCodecIdQcelp AVCodecID = C.AV_CODEC_ID_QCELP
	// AVCodecIdWavpack wraps AV_CODEC_ID_WAVPACK.
	AVCodecIdWavpack AVCodecID = C.AV_CODEC_ID_WAVPACK
	// AVCodecIdDsicinaudio wraps AV_CODEC_ID_DSICINAUDIO.
	AVCodecIdDsicinaudio AVCodecID = C.AV_CODEC_ID_DSICINAUDIO
	// AVCodecIdImc wraps AV_CODEC_ID_IMC.
	AVCodecIdImc AVCodecID = C.AV_CODEC_ID_IMC
	// AVCodecIdMusepack7 wraps AV_CODEC_ID_MUSEPACK7.
	AVCodecIdMusepack7 AVCodecID = C.AV_CODEC_ID_MUSEPACK7
	// AVCodecIdMlp wraps AV_CODEC_ID_MLP.
	AVCodecIdMlp AVCodecID = C.AV_CODEC_ID_MLP
	// AVCodecIdGsmMs wraps AV_CODEC_ID_GSM_MS.
	//
	//	as found in WAV
	AVCodecIdGsmMs AVCodecID = C.AV_CODEC_ID_GSM_MS
	// AVCodecIdAtrac3 wraps AV_CODEC_ID_ATRAC3.
	AVCodecIdAtrac3 AVCodecID = C.AV_CODEC_ID_ATRAC3
	// AVCodecIdApe wraps AV_CODEC_ID_APE.
	AVCodecIdApe AVCodecID = C.AV_CODEC_ID_APE
	// AVCodecIdNellymoser wraps AV_CODEC_ID_NELLYMOSER.
	AVCodecIdNellymoser AVCodecID = C.AV_CODEC_ID_NELLYMOSER
	// AVCodecIdMusepack8 wraps AV_CODEC_ID_MUSEPACK8.
	AVCodecIdMusepack8 AVCodecID = C.AV_CODEC_ID_MUSEPACK8
	// AVCodecIdSpeex wraps AV_CODEC_ID_SPEEX.
	AVCodecIdSpeex AVCodecID = C.AV_CODEC_ID_SPEEX
	// AVCodecIdWmavoice wraps AV_CODEC_ID_WMAVOICE.
	AVCodecIdWmavoice AVCodecID = C.AV_CODEC_ID_WMAVOICE
	// AVCodecIdWmapro wraps AV_CODEC_ID_WMAPRO.
	AVCodecIdWmapro AVCodecID = C.AV_CODEC_ID_WMAPRO
	// AVCodecIdWmalossless wraps AV_CODEC_ID_WMALOSSLESS.
	AVCodecIdWmalossless AVCodecID = C.AV_CODEC_ID_WMALOSSLESS
	// AVCodecIdAtrac3P wraps AV_CODEC_ID_ATRAC3P.
	AVCodecIdAtrac3P AVCodecID = C.AV_CODEC_ID_ATRAC3P
	// AVCodecIdEac3 wraps AV_CODEC_ID_EAC3.
	AVCodecIdEac3 AVCodecID = C.AV_CODEC_ID_EAC3
	// AVCodecIdSipr wraps AV_CODEC_ID_SIPR.
	AVCodecIdSipr AVCodecID = C.AV_CODEC_ID_SIPR
	// AVCodecIdMp1 wraps AV_CODEC_ID_MP1.
	AVCodecIdMp1 AVCodecID = C.AV_CODEC_ID_MP1
	// AVCodecIdTwinvq wraps AV_CODEC_ID_TWINVQ.
	AVCodecIdTwinvq AVCodecID = C.AV_CODEC_ID_TWINVQ
	// AVCodecIdTruehd wraps AV_CODEC_ID_TRUEHD.
	AVCodecIdTruehd AVCodecID = C.AV_CODEC_ID_TRUEHD
	// AVCodecIdMp4Als wraps AV_CODEC_ID_MP4ALS.
	AVCodecIdMp4Als AVCodecID = C.AV_CODEC_ID_MP4ALS
	// AVCodecIdAtrac1 wraps AV_CODEC_ID_ATRAC1.
	AVCodecIdAtrac1 AVCodecID = C.AV_CODEC_ID_ATRAC1
	// AVCodecIdBinkaudioRdft wraps AV_CODEC_ID_BINKAUDIO_RDFT.
	AVCodecIdBinkaudioRdft AVCodecID = C.AV_CODEC_ID_BINKAUDIO_RDFT
	// AVCodecIdBinkaudioDct wraps AV_CODEC_ID_BINKAUDIO_DCT.
	AVCodecIdBinkaudioDct AVCodecID = C.AV_CODEC_ID_BINKAUDIO_DCT
	// AVCodecIdAacLatm wraps AV_CODEC_ID_AAC_LATM.
	AVCodecIdAacLatm AVCodecID = C.AV_CODEC_ID_AAC_LATM
	// AVCodecIdQdmc wraps AV_CODEC_ID_QDMC.
	AVCodecIdQdmc AVCodecID = C.AV_CODEC_ID_QDMC
	// AVCodecIdCelt wraps AV_CODEC_ID_CELT.
	AVCodecIdCelt AVCodecID = C.AV_CODEC_ID_CELT
	// AVCodecIdG7231 wraps AV_CODEC_ID_G723_1.
	AVCodecIdG7231 AVCodecID = C.AV_CODEC_ID_G723_1
	// AVCodecIdG729 wraps AV_CODEC_ID_G729.
	AVCodecIdG729 AVCodecID = C.AV_CODEC_ID_G729
	// AVCodecId8SvxExp wraps AV_CODEC_ID_8SVX_EXP.
	AVCodecId8SvxExp AVCodecID = C.AV_CODEC_ID_8SVX_EXP
	// AVCodecId8SvxFib wraps AV_CODEC_ID_8SVX_FIB.
	AVCodecId8SvxFib AVCodecID = C.AV_CODEC_ID_8SVX_FIB
	// AVCodecIdBmvAudio wraps AV_CODEC_ID_BMV_AUDIO.
	AVCodecIdBmvAudio AVCodecID = C.AV_CODEC_ID_BMV_AUDIO
	// AVCodecIdRalf wraps AV_CODEC_ID_RALF.
	AVCodecIdRalf AVCodecID = C.AV_CODEC_ID_RALF
	// AVCodecIdIac wraps AV_CODEC_ID_IAC.
	AVCodecIdIac AVCodecID = C.AV_CODEC_ID_IAC
	// AVCodecIdIlbc wraps AV_CODEC_ID_ILBC.
	AVCodecIdIlbc AVCodecID = C.AV_CODEC_ID_ILBC
	// AVCodecIdOpus wraps AV_CODEC_ID_OPUS.
	AVCodecIdOpus AVCodecID = C.AV_CODEC_ID_OPUS
	// AVCodecIdComfortNoise wraps AV_CODEC_ID_COMFORT_NOISE.
	AVCodecIdComfortNoise AVCodecID = C.AV_CODEC_ID_COMFORT_NOISE
	// AVCodecIdTak wraps AV_CODEC_ID_TAK.
	AVCodecIdTak AVCodecID = C.AV_CODEC_ID_TAK
	// AVCodecIdMetasound wraps AV_CODEC_ID_METASOUND.
	AVCodecIdMetasound AVCodecID = C.AV_CODEC_ID_METASOUND
	// AVCodecIdPafAudio wraps AV_CODEC_ID_PAF_AUDIO.
	AVCodecIdPafAudio AVCodecID = C.AV_CODEC_ID_PAF_AUDIO
	// AVCodecIdOn2Avc wraps AV_CODEC_ID_ON2AVC.
	AVCodecIdOn2Avc AVCodecID = C.AV_CODEC_ID_ON2AVC
	// AVCodecIdDssSp wraps AV_CODEC_ID_DSS_SP.
	AVCodecIdDssSp AVCodecID = C.AV_CODEC_ID_DSS_SP
	// AVCodecIdCodec2 wraps AV_CODEC_ID_CODEC2.
	AVCodecIdCodec2 AVCodecID = C.AV_CODEC_ID_CODEC2
	// AVCodecIdFfwavesynth wraps AV_CODEC_ID_FFWAVESYNTH.
	AVCodecIdFfwavesynth AVCodecID = C.AV_CODEC_ID_FFWAVESYNTH
	// AVCodecIdSonic wraps AV_CODEC_ID_SONIC.
	AVCodecIdSonic AVCodecID = C.AV_CODEC_ID_SONIC
	// AVCodecIdSonicLs wraps AV_CODEC_ID_SONIC_LS.
	AVCodecIdSonicLs AVCodecID = C.AV_CODEC_ID_SONIC_LS
	// AVCodecIdEvrc wraps AV_CODEC_ID_EVRC.
	AVCodecIdEvrc AVCodecID = C.AV_CODEC_ID_EVRC
	// AVCodecIdSmv wraps AV_CODEC_ID_SMV.
	AVCodecIdSmv AVCodecID = C.AV_CODEC_ID_SMV
	// AVCodecIdDsdLsbf wraps AV_CODEC_ID_DSD_LSBF.
	AVCodecIdDsdLsbf AVCodecID = C.AV_CODEC_ID_DSD_LSBF
	// AVCodecIdDsdMsbf wraps AV_CODEC_ID_DSD_MSBF.
	AVCodecIdDsdMsbf AVCodecID = C.AV_CODEC_ID_DSD_MSBF
	// AVCodecIdDsdLsbfPlanar wraps AV_CODEC_ID_DSD_LSBF_PLANAR.
	AVCodecIdDsdLsbfPlanar AVCodecID = C.AV_CODEC_ID_DSD_LSBF_PLANAR
	// AVCodecIdDsdMsbfPlanar wraps AV_CODEC_ID_DSD_MSBF_PLANAR.
	AVCodecIdDsdMsbfPlanar AVCodecID = C.AV_CODEC_ID_DSD_MSBF_PLANAR
	// AVCodecId4Gv wraps AV_CODEC_ID_4GV.
	AVCodecId4Gv AVCodecID = C.AV_CODEC_ID_4GV
	// AVCodecIdInterplayAcm wraps AV_CODEC_ID_INTERPLAY_ACM.
	AVCodecIdInterplayAcm AVCodecID = C.AV_CODEC_ID_INTERPLAY_ACM
	// AVCodecIdXma1 wraps AV_CODEC_ID_XMA1.
	AVCodecIdXma1 AVCodecID = C.AV_CODEC_ID_XMA1
	// AVCodecIdXma2 wraps AV_CODEC_ID_XMA2.
	AVCodecIdXma2 AVCodecID = C.AV_CODEC_ID_XMA2
	// AVCodecIdDst wraps AV_CODEC_ID_DST.
	AVCodecIdDst AVCodecID = C.AV_CODEC_ID_DST
	// AVCodecIdAtrac3Al wraps AV_CODEC_ID_ATRAC3AL.
	AVCodecIdAtrac3Al AVCodecID = C.AV_CODEC_ID_ATRAC3AL
	// AVCodecIdAtrac3Pal wraps AV_CODEC_ID_ATRAC3PAL.
	AVCodecIdAtrac3Pal AVCodecID = C.AV_CODEC_ID_ATRAC3PAL
	// AVCodecIdDolbyE wraps AV_CODEC_ID_DOLBY_E.
	AVCodecIdDolbyE AVCodecID = C.AV_CODEC_ID_DOLBY_E
	// AVCodecIdAptx wraps AV_CODEC_ID_APTX.
	AVCodecIdAptx AVCodecID = C.AV_CODEC_ID_APTX
	// AVCodecIdAptxHd wraps AV_CODEC_ID_APTX_HD.
	AVCodecIdAptxHd AVCodecID = C.AV_CODEC_ID_APTX_HD
	// AVCodecIdSbc wraps AV_CODEC_ID_SBC.
	AVCodecIdSbc AVCodecID = C.AV_CODEC_ID_SBC
	// AVCodecIdAtrac9 wraps AV_CODEC_ID_ATRAC9.
	AVCodecIdAtrac9 AVCodecID = C.AV_CODEC_ID_ATRAC9
	// AVCodecIdHcom wraps AV_CODEC_ID_HCOM.
	AVCodecIdHcom AVCodecID = C.AV_CODEC_ID_HCOM
	// AVCodecIdAcelpKelvin wraps AV_CODEC_ID_ACELP_KELVIN.
	AVCodecIdAcelpKelvin AVCodecID = C.AV_CODEC_ID_ACELP_KELVIN
	// AVCodecIdMpegh3DAudio wraps AV_CODEC_ID_MPEGH_3D_AUDIO.
	AVCodecIdMpegh3DAudio AVCodecID = C.AV_CODEC_ID_MPEGH_3D_AUDIO
	// AVCodecIdSiren wraps AV_CODEC_ID_SIREN.
	AVCodecIdSiren AVCodecID = C.AV_CODEC_ID_SIREN
	// AVCodecIdHca wraps AV_CODEC_ID_HCA.
	AVCodecIdHca AVCodecID = C.AV_CODEC_ID_HCA
	// AVCodecIdFastaudio wraps AV_CODEC_ID_FASTAUDIO.
	AVCodecIdFastaudio AVCodecID = C.AV_CODEC_ID_FASTAUDIO
	// AVCodecIdMsnsiren wraps AV_CODEC_ID_MSNSIREN.
	AVCodecIdMsnsiren AVCodecID = C.AV_CODEC_ID_MSNSIREN
	// AVCodecIdDfpwm wraps AV_CODEC_ID_DFPWM.
	AVCodecIdDfpwm AVCodecID = C.AV_CODEC_ID_DFPWM
	// AVCodecIdBonk wraps AV_CODEC_ID_BONK.
	AVCodecIdBonk AVCodecID = C.AV_CODEC_ID_BONK
	// AVCodecIdMisc4 wraps AV_CODEC_ID_MISC4.
	AVCodecIdMisc4 AVCodecID = C.AV_CODEC_ID_MISC4
	// AVCodecIdApac wraps AV_CODEC_ID_APAC.
	AVCodecIdApac AVCodecID = C.AV_CODEC_ID_APAC
	// AVCodecIdFtr wraps AV_CODEC_ID_FTR.
	AVCodecIdFtr AVCodecID = C.AV_CODEC_ID_FTR
	// AVCodecIdWavarc wraps AV_CODEC_ID_WAVARC.
	AVCodecIdWavarc AVCodecID = C.AV_CODEC_ID_WAVARC
	// AVCodecIdRka wraps AV_CODEC_ID_RKA.
	AVCodecIdRka AVCodecID = C.AV_CODEC_ID_RKA
	// AVCodecIdFirstSubtitle wraps AV_CODEC_ID_FIRST_SUBTITLE.
	//
	//	A dummy ID pointing at the start of subtitle codecs.
	AVCodecIdFirstSubtitle AVCodecID = C.AV_CODEC_ID_FIRST_SUBTITLE
	// AVCodecIdDvdSubtitle wraps AV_CODEC_ID_DVD_SUBTITLE.
	AVCodecIdDvdSubtitle AVCodecID = C.AV_CODEC_ID_DVD_SUBTITLE
	// AVCodecIdDvbSubtitle wraps AV_CODEC_ID_DVB_SUBTITLE.
	AVCodecIdDvbSubtitle AVCodecID = C.AV_CODEC_ID_DVB_SUBTITLE
	// AVCodecIdText wraps AV_CODEC_ID_TEXT.
	//
	//	raw UTF-8 text
	AVCodecIdText AVCodecID = C.AV_CODEC_ID_TEXT
	// AVCodecIdXsub wraps AV_CODEC_ID_XSUB.
	AVCodecIdXsub AVCodecID = C.AV_CODEC_ID_XSUB
	// AVCodecIdSsa wraps AV_CODEC_ID_SSA.
	AVCodecIdSsa AVCodecID = C.AV_CODEC_ID_SSA
	// AVCodecIdMovText wraps AV_CODEC_ID_MOV_TEXT.
	AVCodecIdMovText AVCodecID = C.AV_CODEC_ID_MOV_TEXT
	// AVCodecIdHdmvPgsSubtitle wraps AV_CODEC_ID_HDMV_PGS_SUBTITLE.
	AVCodecIdHdmvPgsSubtitle AVCodecID = C.AV_CODEC_ID_HDMV_PGS_SUBTITLE
	// AVCodecIdDvbTeletext wraps AV_CODEC_ID_DVB_TELETEXT.
	AVCodecIdDvbTeletext AVCodecID = C.AV_CODEC_ID_DVB_TELETEXT
	// AVCodecIdSrt wraps AV_CODEC_ID_SRT.
	AVCodecIdSrt AVCodecID = C.AV_CODEC_ID_SRT
	// AVCodecIdMicrodvd wraps AV_CODEC_ID_MICRODVD.
	AVCodecIdMicrodvd AVCodecID = C.AV_CODEC_ID_MICRODVD
	// AVCodecIdEia608 wraps AV_CODEC_ID_EIA_608.
	AVCodecIdEia608 AVCodecID = C.AV_CODEC_ID_EIA_608
	// AVCodecIdJacosub wraps AV_CODEC_ID_JACOSUB.
	AVCodecIdJacosub AVCodecID = C.AV_CODEC_ID_JACOSUB
	// AVCodecIdSami wraps AV_CODEC_ID_SAMI.
	AVCodecIdSami AVCodecID = C.AV_CODEC_ID_SAMI
	// AVCodecIdRealtext wraps AV_CODEC_ID_REALTEXT.
	AVCodecIdRealtext AVCodecID = C.AV_CODEC_ID_REALTEXT
	// AVCodecIdStl wraps AV_CODEC_ID_STL.
	AVCodecIdStl AVCodecID = C.AV_CODEC_ID_STL
	// AVCodecIdSubviewer1 wraps AV_CODEC_ID_SUBVIEWER1.
	AVCodecIdSubviewer1 AVCodecID = C.AV_CODEC_ID_SUBVIEWER1
	// AVCodecIdSubviewer wraps AV_CODEC_ID_SUBVIEWER.
	AVCodecIdSubviewer AVCodecID = C.AV_CODEC_ID_SUBVIEWER
	// AVCodecIdSubrip wraps AV_CODEC_ID_SUBRIP.
	AVCodecIdSubrip AVCodecID = C.AV_CODEC_ID_SUBRIP
	// AVCodecIdWebvtt wraps AV_CODEC_ID_WEBVTT.
	AVCodecIdWebvtt AVCodecID = C.AV_CODEC_ID_WEBVTT
	// AVCodecIdMpl2 wraps AV_CODEC_ID_MPL2.
	AVCodecIdMpl2 AVCodecID = C.AV_CODEC_ID_MPL2
	// AVCodecIdVplayer wraps AV_CODEC_ID_VPLAYER.
	AVCodecIdVplayer AVCodecID = C.AV_CODEC_ID_VPLAYER
	// AVCodecIdPjs wraps AV_CODEC_ID_PJS.
	AVCodecIdPjs AVCodecID = C.AV_CODEC_ID_PJS
	// AVCodecIdAss wraps AV_CODEC_ID_ASS.
	AVCodecIdAss AVCodecID = C.AV_CODEC_ID_ASS
	// AVCodecIdHdmvTextSubtitle wraps AV_CODEC_ID_HDMV_TEXT_SUBTITLE.
	AVCodecIdHdmvTextSubtitle AVCodecID = C.AV_CODEC_ID_HDMV_TEXT_SUBTITLE
	// AVCodecIdTtml wraps AV_CODEC_ID_TTML.
	AVCodecIdTtml AVCodecID = C.AV_CODEC_ID_TTML
	// AVCodecIdAribCaption wraps AV_CODEC_ID_ARIB_CAPTION.
	AVCodecIdAribCaption AVCodecID = C.AV_CODEC_ID_ARIB_CAPTION
	// AVCodecIdFirstUnknown wraps AV_CODEC_ID_FIRST_UNKNOWN.
	//
	//	A dummy ID pointing at the start of various fake codecs.
	AVCodecIdFirstUnknown AVCodecID = C.AV_CODEC_ID_FIRST_UNKNOWN
	// AVCodecIdTtf wraps AV_CODEC_ID_TTF.
	AVCodecIdTtf AVCodecID = C.AV_CODEC_ID_TTF
	// AVCodecIdScte35 wraps AV_CODEC_ID_SCTE_35.
	//
	//	Contain timestamp estimated through PCR of program stream.
	AVCodecIdScte35 AVCodecID = C.AV_CODEC_ID_SCTE_35
	// AVCodecIdEpg wraps AV_CODEC_ID_EPG.
	AVCodecIdEpg AVCodecID = C.AV_CODEC_ID_EPG
	// AVCodecIdBintext wraps AV_CODEC_ID_BINTEXT.
	AVCodecIdBintext AVCodecID = C.AV_CODEC_ID_BINTEXT
	// AVCodecIdXbin wraps AV_CODEC_ID_XBIN.
	AVCodecIdXbin AVCodecID = C.AV_CODEC_ID_XBIN
	// AVCodecIdIdf wraps AV_CODEC_ID_IDF.
	AVCodecIdIdf AVCodecID = C.AV_CODEC_ID_IDF
	// AVCodecIdOtf wraps AV_CODEC_ID_OTF.
	AVCodecIdOtf AVCodecID = C.AV_CODEC_ID_OTF
	// AVCodecIdSmpteKlv wraps AV_CODEC_ID_SMPTE_KLV.
	AVCodecIdSmpteKlv AVCodecID = C.AV_CODEC_ID_SMPTE_KLV
	// AVCodecIdDvdNav wraps AV_CODEC_ID_DVD_NAV.
	AVCodecIdDvdNav AVCodecID = C.AV_CODEC_ID_DVD_NAV
	// AVCodecIdTimedId3 wraps AV_CODEC_ID_TIMED_ID3.
	AVCodecIdTimedId3 AVCodecID = C.AV_CODEC_ID_TIMED_ID3
	// AVCodecIdBinData wraps AV_CODEC_ID_BIN_DATA.
	AVCodecIdBinData AVCodecID = C.AV_CODEC_ID_BIN_DATA
	// AVCodecIdProbe wraps AV_CODEC_ID_PROBE.
	//
	//	codec_id is not known (like AV_CODEC_ID_NONE) but lavf should attempt to identify it
	AVCodecIdProbe AVCodecID = C.AV_CODEC_ID_PROBE
	// AVCodecIdMpeg2Ts wraps AV_CODEC_ID_MPEG2TS.
	/*
	   < _FAKE_ codec to indicate a raw MPEG-2 TS
	   stream (only used by libavformat)
	*/
	AVCodecIdMpeg2Ts AVCodecID = C.AV_CODEC_ID_MPEG2TS
	// AVCodecIdMpeg4Systems wraps AV_CODEC_ID_MPEG4SYSTEMS.
	/*
	   < _FAKE_ codec to indicate a MPEG-4 Systems
	   stream (only used by libavformat)
	*/
	AVCodecIdMpeg4Systems AVCodecID = C.AV_CODEC_ID_MPEG4SYSTEMS
	// AVCodecIdFfmetadata wraps AV_CODEC_ID_FFMETADATA.
	//
	//	Dummy codec for streams containing only metadata information.
	AVCodecIdFfmetadata AVCodecID = C.AV_CODEC_ID_FFMETADATA
	// AVCodecIdWrappedAvframe wraps AV_CODEC_ID_WRAPPED_AVFRAME.
	//
	//	Passthrough codec, AVFrames wrapped in AVPacket
	AVCodecIdWrappedAvframe AVCodecID = C.AV_CODEC_ID_WRAPPED_AVFRAME
	// AVCodecIdVnull wraps AV_CODEC_ID_VNULL.
	/*
	   Dummy null video codec, useful mainly for development and debugging.
	   Null encoder/decoder discard all input and never return any output.
	*/
	AVCodecIdVnull AVCodecID = C.AV_CODEC_ID_VNULL
	// AVCodecIdAnull wraps AV_CODEC_ID_ANULL.
	/*
	   Dummy null audio codec, useful mainly for development and debugging.
	   Null encoder/decoder discard all input and never return any output.
	*/
	AVCodecIdAnull AVCodecID = C.AV_CODEC_ID_ANULL
)

// --- Enum AVFieldOrder ---

// AVFieldOrder wraps AVFieldOrder.
type AVFieldOrder C.enum_AVFieldOrder

const SizeOfAVFieldOrder = C.sizeof_enum_AVFieldOrder

func ToAVFieldOrderArray(ptr unsafe.Pointer) *Array[AVFieldOrder] {
	if ptr == nil {
		return nil
	}

	return &Array[AVFieldOrder]{
		elemSize: SizeOfAVFieldOrder,
		loadPtr: func(pointer unsafe.Pointer) AVFieldOrder {
			ptr := (*AVFieldOrder)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVFieldOrder) {
			ptr := (*AVFieldOrder)(pointer)
			*ptr = value
		},
	}
}

func AllocAVFieldOrderArray(size uint64) *Array[AVFieldOrder] {
	return ToAVFieldOrderArray(AVCalloc(size, SizeOfAVFieldOrder))
}

const (
	// AVFieldUnknown wraps AV_FIELD_UNKNOWN.
	AVFieldUnknown AVFieldOrder = C.AV_FIELD_UNKNOWN
	// AVFieldProgressive wraps AV_FIELD_PROGRESSIVE.
	AVFieldProgressive AVFieldOrder = C.AV_FIELD_PROGRESSIVE
	// AVFieldTt wraps AV_FIELD_TT.
	//
	//	Top coded_first, top displayed first
	AVFieldTt AVFieldOrder = C.AV_FIELD_TT
	// AVFieldBb wraps AV_FIELD_BB.
	//
	//	Bottom coded first, bottom displayed first
	AVFieldBb AVFieldOrder = C.AV_FIELD_BB
	// AVFieldTb wraps AV_FIELD_TB.
	//
	//	Top coded first, bottom displayed first
	AVFieldTb AVFieldOrder = C.AV_FIELD_TB
	// AVFieldBt wraps AV_FIELD_BT.
	//
	//	Bottom coded first, top displayed first
	AVFieldBt AVFieldOrder = C.AV_FIELD_BT
)

// --- Enum AVDiscard ---

// AVDiscard wraps AVDiscard.
type AVDiscard C.enum_AVDiscard

const SizeOfAVDiscard = C.sizeof_enum_AVDiscard

func ToAVDiscardArray(ptr unsafe.Pointer) *Array[AVDiscard] {
	if ptr == nil {
		return nil
	}

	return &Array[AVDiscard]{
		elemSize: SizeOfAVDiscard,
		loadPtr: func(pointer unsafe.Pointer) AVDiscard {
			ptr := (*AVDiscard)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVDiscard) {
			ptr := (*AVDiscard)(pointer)
			*ptr = value
		},
	}
}

func AllocAVDiscardArray(size uint64) *Array[AVDiscard] {
	return ToAVDiscardArray(AVCalloc(size, SizeOfAVDiscard))
}

const (
	// AVDiscardNone wraps AVDISCARD_NONE.
	//
	//	discard nothing
	AVDiscardNone AVDiscard = C.AVDISCARD_NONE
	// AVDiscardDefault wraps AVDISCARD_DEFAULT.
	//
	//	discard useless packets like 0 size packets in avi
	AVDiscardDefault AVDiscard = C.AVDISCARD_DEFAULT
	// AVDiscardNonref wraps AVDISCARD_NONREF.
	//
	//	discard all non reference
	AVDiscardNonref AVDiscard = C.AVDISCARD_NONREF
	// AVDiscardBidir wraps AVDISCARD_BIDIR.
	//
	//	discard all bidirectional frames
	AVDiscardBidir AVDiscard = C.AVDISCARD_BIDIR
	// AVDiscardNonintra wraps AVDISCARD_NONINTRA.
	//
	//	discard all non intra frames
	AVDiscardNonintra AVDiscard = C.AVDISCARD_NONINTRA
	// AVDiscardNonkey wraps AVDISCARD_NONKEY.
	//
	//	discard all frames except keyframes
	AVDiscardNonkey AVDiscard = C.AVDISCARD_NONKEY
	// AVDiscardAll wraps AVDISCARD_ALL.
	//
	//	discard all
	AVDiscardAll AVDiscard = C.AVDISCARD_ALL
)

// --- Enum AVAudioServiceType ---

// AVAudioServiceType wraps AVAudioServiceType.
type AVAudioServiceType C.enum_AVAudioServiceType

const SizeOfAVAudioServiceType = C.sizeof_enum_AVAudioServiceType

func ToAVAudioServiceTypeArray(ptr unsafe.Pointer) *Array[AVAudioServiceType] {
	if ptr == nil {
		return nil
	}

	return &Array[AVAudioServiceType]{
		elemSize: SizeOfAVAudioServiceType,
		loadPtr: func(pointer unsafe.Pointer) AVAudioServiceType {
			ptr := (*AVAudioServiceType)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVAudioServiceType) {
			ptr := (*AVAudioServiceType)(pointer)
			*ptr = value
		},
	}
}

func AllocAVAudioServiceTypeArray(size uint64) *Array[AVAudioServiceType] {
	return ToAVAudioServiceTypeArray(AVCalloc(size, SizeOfAVAudioServiceType))
}

const (
	// AVAudioServiceTypeMain wraps AV_AUDIO_SERVICE_TYPE_MAIN.
	AVAudioServiceTypeMain AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_MAIN
	// AVAudioServiceTypeEffects wraps AV_AUDIO_SERVICE_TYPE_EFFECTS.
	AVAudioServiceTypeEffects AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_EFFECTS
	// AVAudioServiceTypeVisuallyImpaired wraps AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED.
	AVAudioServiceTypeVisuallyImpaired AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_VISUALLY_IMPAIRED
	// AVAudioServiceTypeHearingImpaired wraps AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED.
	AVAudioServiceTypeHearingImpaired AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_HEARING_IMPAIRED
	// AVAudioServiceTypeDialogue wraps AV_AUDIO_SERVICE_TYPE_DIALOGUE.
	AVAudioServiceTypeDialogue AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_DIALOGUE
	// AVAudioServiceTypeCommentary wraps AV_AUDIO_SERVICE_TYPE_COMMENTARY.
	AVAudioServiceTypeCommentary AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_COMMENTARY
	// AVAudioServiceTypeEmergency wraps AV_AUDIO_SERVICE_TYPE_EMERGENCY.
	AVAudioServiceTypeEmergency AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_EMERGENCY
	// AVAudioServiceTypeVoiceOver wraps AV_AUDIO_SERVICE_TYPE_VOICE_OVER.
	AVAudioServiceTypeVoiceOver AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_VOICE_OVER
	// AVAudioServiceTypeKaraoke wraps AV_AUDIO_SERVICE_TYPE_KARAOKE.
	AVAudioServiceTypeKaraoke AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_KARAOKE
	// AVAudioServiceTypeNb wraps AV_AUDIO_SERVICE_TYPE_NB.
	//
	//	Not part of ABI
	AVAudioServiceTypeNb AVAudioServiceType = C.AV_AUDIO_SERVICE_TYPE_NB
)

// --- Enum AVPacketSideDataType ---

// AVPacketSideDataType wraps AVPacketSideDataType.
type AVPacketSideDataType C.enum_AVPacketSideDataType

const SizeOfAVPacketSideDataType = C.sizeof_enum_AVPacketSideDataType

func ToAVPacketSideDataTypeArray(ptr unsafe.Pointer) *Array[AVPacketSideDataType] {
	if ptr == nil {
		return nil
	}

	return &Array[AVPacketSideDataType]{
		elemSize: SizeOfAVPacketSideDataType,
		loadPtr: func(pointer unsafe.Pointer) AVPacketSideDataType {
			ptr := (*AVPacketSideDataType)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVPacketSideDataType) {
			ptr := (*AVPacketSideDataType)(pointer)
			*ptr = value
		},
	}
}

func AllocAVPacketSideDataTypeArray(size uint64) *Array[AVPacketSideDataType] {
	return ToAVPacketSideDataTypeArray(AVCalloc(size, SizeOfAVPacketSideDataType))
}

const (
	// AVPktDataPalette wraps AV_PKT_DATA_PALETTE.
	/*
	   An AV_PKT_DATA_PALETTE side data packet contains exactly AVPALETTE_SIZE
	   bytes worth of palette. This side data signals that a new palette is
	   present.
	*/
	AVPktDataPalette AVPacketSideDataType = C.AV_PKT_DATA_PALETTE
	// AVPktDataNewExtradata wraps AV_PKT_DATA_NEW_EXTRADATA.
	/*
	   The AV_PKT_DATA_NEW_EXTRADATA is used to notify the codec or the format
	   that the extradata buffer was changed and the receiving side should
	   act upon it appropriately. The new extradata is embedded in the side
	   data buffer and should be immediately used for processing the current
	   frame or packet.
	*/
	AVPktDataNewExtradata AVPacketSideDataType = C.AV_PKT_DATA_NEW_EXTRADATA
	// AVPktDataParamChange wraps AV_PKT_DATA_PARAM_CHANGE.
	/*
	   An AV_PKT_DATA_PARAM_CHANGE side data packet is laid out as follows:
	   @code
	   u32le param_flags
	   if (param_flags & AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT)
	       s32le channel_count
	   if (param_flags & AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT)
	       u64le channel_layout
	   if (param_flags & AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE)
	       s32le sample_rate
	   if (param_flags & AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS)
	       s32le width
	       s32le height
	   @endcode
	*/
	AVPktDataParamChange AVPacketSideDataType = C.AV_PKT_DATA_PARAM_CHANGE
	// AVPktDataH263MbInfo wraps AV_PKT_DATA_H263_MB_INFO.
	/*
	   An AV_PKT_DATA_H263_MB_INFO side data packet contains a number of
	   structures with info about macroblocks relevant to splitting the
	   packet into smaller packets on macroblock edges (e.g. as for RFC 2190).
	   That is, it does not necessarily contain info about all macroblocks,
	   as long as the distance between macroblocks in the info is smaller
	   than the target payload size.
	   Each MB info structure is 12 bytes, and is laid out as follows:
	   @code
	   u32le bit offset from the start of the packet
	   u8    current quantizer at the start of the macroblock
	   u8    GOB number
	   u16le macroblock address within the GOB
	   u8    horizontal MV predictor
	   u8    vertical MV predictor
	   u8    horizontal MV predictor for block number 3
	   u8    vertical MV predictor for block number 3
	   @endcode
	*/
	AVPktDataH263MbInfo AVPacketSideDataType = C.AV_PKT_DATA_H263_MB_INFO
	// AVPktDataReplaygain wraps AV_PKT_DATA_REPLAYGAIN.
	/*
	   This side data should be associated with an audio stream and contains
	   ReplayGain information in form of the AVReplayGain struct.
	*/
	AVPktDataReplaygain AVPacketSideDataType = C.AV_PKT_DATA_REPLAYGAIN
	// AVPktDataDisplaymatrix wraps AV_PKT_DATA_DISPLAYMATRIX.
	/*
	   This side data contains a 3x3 transformation matrix describing an affine
	   transformation that needs to be applied to the decoded video frames for
	   correct presentation.

	   See libavutil/display.h for a detailed description of the data.
	*/
	AVPktDataDisplaymatrix AVPacketSideDataType = C.AV_PKT_DATA_DISPLAYMATRIX
	// AVPktDataStereo3D wraps AV_PKT_DATA_STEREO3D.
	/*
	   This side data should be associated with a video stream and contains
	   Stereoscopic 3D information in form of the AVStereo3D struct.
	*/
	AVPktDataStereo3D AVPacketSideDataType = C.AV_PKT_DATA_STEREO3D
	// AVPktDataAudioServiceType wraps AV_PKT_DATA_AUDIO_SERVICE_TYPE.
	/*
	   This side data should be associated with an audio stream and corresponds
	   to enum AVAudioServiceType.
	*/
	AVPktDataAudioServiceType AVPacketSideDataType = C.AV_PKT_DATA_AUDIO_SERVICE_TYPE
	// AVPktDataQualityStats wraps AV_PKT_DATA_QUALITY_STATS.
	/*
	   This side data contains quality related information from the encoder.
	   @code
	   u32le quality factor of the compressed frame. Allowed range is between 1 (good) and FF_LAMBDA_MAX (bad).
	   u8    picture type
	   u8    error count
	   u16   reserved
	   u64le[error count] sum of squared differences between encoder in and output
	   @endcode
	*/
	AVPktDataQualityStats AVPacketSideDataType = C.AV_PKT_DATA_QUALITY_STATS
	// AVPktDataFallbackTrack wraps AV_PKT_DATA_FALLBACK_TRACK.
	/*
	   This side data contains an integer value representing the stream index
	   of a "fallback" track.  A fallback track indicates an alternate
	   track to use when the current track can not be decoded for some reason.
	   e.g. no decoder available for codec.
	*/
	AVPktDataFallbackTrack AVPacketSideDataType = C.AV_PKT_DATA_FALLBACK_TRACK
	// AVPktDataCpbProperties wraps AV_PKT_DATA_CPB_PROPERTIES.
	//
	//	This side data corresponds to the AVCPBProperties struct.
	AVPktDataCpbProperties AVPacketSideDataType = C.AV_PKT_DATA_CPB_PROPERTIES
	// AVPktDataSkipSamples wraps AV_PKT_DATA_SKIP_SAMPLES.
	/*
	   Recommmends skipping the specified number of samples
	   @code
	   u32le number of samples to skip from start of this packet
	   u32le number of samples to skip from end of this packet
	   u8    reason for start skip
	   u8    reason for end   skip (0=padding silence, 1=convergence)
	   @endcode
	*/
	AVPktDataSkipSamples AVPacketSideDataType = C.AV_PKT_DATA_SKIP_SAMPLES
	// AVPktDataJpDualmono wraps AV_PKT_DATA_JP_DUALMONO.
	/*
	   An AV_PKT_DATA_JP_DUALMONO side data packet indicates that
	   the packet may contain "dual mono" audio specific to Japanese DTV
	   and if it is true, recommends only the selected channel to be used.
	   @code
	   u8    selected channels (0=main/left, 1=sub/right, 2=both)
	   @endcode
	*/
	AVPktDataJpDualmono AVPacketSideDataType = C.AV_PKT_DATA_JP_DUALMONO
	// AVPktDataStringsMetadata wraps AV_PKT_DATA_STRINGS_METADATA.
	/*
	   A list of zero terminated key/value strings. There is no end marker for
	   the list, so it is required to rely on the side data size to stop.
	*/
	AVPktDataStringsMetadata AVPacketSideDataType = C.AV_PKT_DATA_STRINGS_METADATA
	// AVPktDataSubtitlePosition wraps AV_PKT_DATA_SUBTITLE_POSITION.
	/*
	   Subtitle event position
	   @code
	   u32le x1
	   u32le y1
	   u32le x2
	   u32le y2
	   @endcode
	*/
	AVPktDataSubtitlePosition AVPacketSideDataType = C.AV_PKT_DATA_SUBTITLE_POSITION
	// AVPktDataMatroskaBlockadditional wraps AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL.
	/*
	   Data found in BlockAdditional element of matroska container. There is
	   no end marker for the data, so it is required to rely on the side data
	   size to recognize the end. 8 byte id (as found in BlockAddId) followed
	   by data.
	*/
	AVPktDataMatroskaBlockadditional AVPacketSideDataType = C.AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL
	// AVPktDataWebvttIdentifier wraps AV_PKT_DATA_WEBVTT_IDENTIFIER.
	//
	//	The optional first identifier line of a WebVTT cue.
	AVPktDataWebvttIdentifier AVPacketSideDataType = C.AV_PKT_DATA_WEBVTT_IDENTIFIER
	// AVPktDataWebvttSettings wraps AV_PKT_DATA_WEBVTT_SETTINGS.
	/*
	   The optional settings (rendering instructions) that immediately
	   follow the timestamp specifier of a WebVTT cue.
	*/
	AVPktDataWebvttSettings AVPacketSideDataType = C.AV_PKT_DATA_WEBVTT_SETTINGS
	// AVPktDataMetadataUpdate wraps AV_PKT_DATA_METADATA_UPDATE.
	/*
	   A list of zero terminated key/value strings. There is no end marker for
	   the list, so it is required to rely on the side data size to stop. This
	   side data includes updated metadata which appeared in the stream.
	*/
	AVPktDataMetadataUpdate AVPacketSideDataType = C.AV_PKT_DATA_METADATA_UPDATE
	// AVPktDataMpegtsStreamId wraps AV_PKT_DATA_MPEGTS_STREAM_ID.
	/*
	   MPEGTS stream ID as uint8_t, this is required to pass the stream ID
	   information from the demuxer to the corresponding muxer.
	*/
	AVPktDataMpegtsStreamId AVPacketSideDataType = C.AV_PKT_DATA_MPEGTS_STREAM_ID
	// AVPktDataMasteringDisplayMetadata wraps AV_PKT_DATA_MASTERING_DISPLAY_METADATA.
	/*
	   Mastering display metadata (based on SMPTE-2086:2014). This metadata
	   should be associated with a video stream and contains data in the form
	   of the AVMasteringDisplayMetadata struct.
	*/
	AVPktDataMasteringDisplayMetadata AVPacketSideDataType = C.AV_PKT_DATA_MASTERING_DISPLAY_METADATA
	// AVPktDataSpherical wraps AV_PKT_DATA_SPHERICAL.
	/*
	   This side data should be associated with a video stream and corresponds
	   to the AVSphericalMapping structure.
	*/
	AVPktDataSpherical AVPacketSideDataType = C.AV_PKT_DATA_SPHERICAL
	// AVPktDataContentLightLevel wraps AV_PKT_DATA_CONTENT_LIGHT_LEVEL.
	/*
	   Content light level (based on CTA-861.3). This metadata should be
	   associated with a video stream and contains data in the form of the
	   AVContentLightMetadata struct.
	*/
	AVPktDataContentLightLevel AVPacketSideDataType = C.AV_PKT_DATA_CONTENT_LIGHT_LEVEL
	// AVPktDataA53Cc wraps AV_PKT_DATA_A53_CC.
	/*
	   ATSC A53 Part 4 Closed Captions. This metadata should be associated with
	   a video stream. A53 CC bitstream is stored as uint8_t in AVPacketSideData.data.
	   The number of bytes of CC data is AVPacketSideData.size.
	*/
	AVPktDataA53Cc AVPacketSideDataType = C.AV_PKT_DATA_A53_CC
	// AVPktDataEncryptionInitInfo wraps AV_PKT_DATA_ENCRYPTION_INIT_INFO.
	/*
	   This side data is encryption initialization data.
	   The format is not part of ABI, use av_encryption_init_info_* methods to
	   access.
	*/
	AVPktDataEncryptionInitInfo AVPacketSideDataType = C.AV_PKT_DATA_ENCRYPTION_INIT_INFO
	// AVPktDataEncryptionInfo wraps AV_PKT_DATA_ENCRYPTION_INFO.
	/*
	   This side data contains encryption info for how to decrypt the packet.
	   The format is not part of ABI, use av_encryption_info_* methods to access.
	*/
	AVPktDataEncryptionInfo AVPacketSideDataType = C.AV_PKT_DATA_ENCRYPTION_INFO
	// AVPktDataAfd wraps AV_PKT_DATA_AFD.
	/*
	   Active Format Description data consisting of a single byte as specified
	   in ETSI TS 101 154 using AVActiveFormatDescription enum.
	*/
	AVPktDataAfd AVPacketSideDataType = C.AV_PKT_DATA_AFD
	// AVPktDataPrft wraps AV_PKT_DATA_PRFT.
	/*
	   Producer Reference Time data corresponding to the AVProducerReferenceTime struct,
	   usually exported by some encoders (on demand through the prft flag set in the
	   AVCodecContext export_side_data field).
	*/
	AVPktDataPrft AVPacketSideDataType = C.AV_PKT_DATA_PRFT
	// AVPktDataIccProfile wraps AV_PKT_DATA_ICC_PROFILE.
	/*
	   ICC profile data consisting of an opaque octet buffer following the
	   format described by ISO 15076-1.
	*/
	AVPktDataIccProfile AVPacketSideDataType = C.AV_PKT_DATA_ICC_PROFILE
	// AVPktDataDoviConf wraps AV_PKT_DATA_DOVI_CONF.
	/*
	   DOVI configuration
	   ref:
	   dolby-vision-bitstreams-within-the-iso-base-media-file-format-v2.1.2, section 2.2
	   dolby-vision-bitstreams-in-mpeg-2-transport-stream-multiplex-v1.2, section 3.3
	   Tags are stored in struct AVDOVIDecoderConfigurationRecord.
	*/
	AVPktDataDoviConf AVPacketSideDataType = C.AV_PKT_DATA_DOVI_CONF
	// AVPktDataS12MTimecode wraps AV_PKT_DATA_S12M_TIMECODE.
	/*
	   Timecode which conforms to SMPTE ST 12-1:2014. The data is an array of 4 uint32_t
	   where the first uint32_t describes how many (1-3) of the other timecodes are used.
	   The timecode format is described in the documentation of av_timecode_get_smpte_from_framenum()
	   function in libavutil/timecode.h.
	*/
	AVPktDataS12MTimecode AVPacketSideDataType = C.AV_PKT_DATA_S12M_TIMECODE
	// AVPktDataDynamicHdr10Plus wraps AV_PKT_DATA_DYNAMIC_HDR10_PLUS.
	/*
	   HDR10+ dynamic metadata associated with a video frame. The metadata is in
	   the form of the AVDynamicHDRPlus struct and contains
	   information for color volume transform - application 4 of
	   SMPTE 2094-40:2016 standard.
	*/
	AVPktDataDynamicHdr10Plus AVPacketSideDataType = C.AV_PKT_DATA_DYNAMIC_HDR10_PLUS
	// AVPktDataNb wraps AV_PKT_DATA_NB.
	/*
	   The number of side data types.
	   This is not part of the public API/ABI in the sense that it may
	   change when new side data types are added.
	   This must stay the last enum value.
	   If its value becomes huge, some code using it
	   needs to be updated as it assumes it to be smaller than other limits.
	*/
	AVPktDataNb AVPacketSideDataType = C.AV_PKT_DATA_NB
)

// --- Enum AVSideDataParamChangeFlags ---

// AVSideDataParamChangeFlags wraps AVSideDataParamChangeFlags.
type AVSideDataParamChangeFlags C.enum_AVSideDataParamChangeFlags

const SizeOfAVSideDataParamChangeFlags = C.sizeof_enum_AVSideDataParamChangeFlags

func ToAVSideDataParamChangeFlagsArray(ptr unsafe.Pointer) *Array[AVSideDataParamChangeFlags] {
	if ptr == nil {
		return nil
	}

	return &Array[AVSideDataParamChangeFlags]{
		elemSize: SizeOfAVSideDataParamChangeFlags,
		loadPtr: func(pointer unsafe.Pointer) AVSideDataParamChangeFlags {
			ptr := (*AVSideDataParamChangeFlags)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVSideDataParamChangeFlags) {
			ptr := (*AVSideDataParamChangeFlags)(pointer)
			*ptr = value
		},
	}
}

func AllocAVSideDataParamChangeFlagsArray(size uint64) *Array[AVSideDataParamChangeFlags] {
	return ToAVSideDataParamChangeFlagsArray(AVCalloc(size, SizeOfAVSideDataParamChangeFlags))
}

const (
	// AVSideDataParamChangeChannelCount wraps AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT.
	//
	//	@deprecated those are not used by any decoder
	AVSideDataParamChangeChannelCount AVSideDataParamChangeFlags = C.AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT
	// AVSideDataParamChangeChannelLayout wraps AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT.
	//
	//	@deprecated those are not used by any decoder
	AVSideDataParamChangeChannelLayout AVSideDataParamChangeFlags = C.AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT
	// AVSideDataParamChangeSampleRate wraps AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE.
	AVSideDataParamChangeSampleRate AVSideDataParamChangeFlags = C.AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE
	// AVSideDataParamChangeDimensions wraps AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS.
	AVSideDataParamChangeDimensions AVSideDataParamChangeFlags = C.AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS
)

// --- Enum AVStreamParseType ---

// AVStreamParseType wraps AVStreamParseType.
type AVStreamParseType C.enum_AVStreamParseType

const SizeOfAVStreamParseType = C.sizeof_enum_AVStreamParseType

func ToAVStreamParseTypeArray(ptr unsafe.Pointer) *Array[AVStreamParseType] {
	if ptr == nil {
		return nil
	}

	return &Array[AVStreamParseType]{
		elemSize: SizeOfAVStreamParseType,
		loadPtr: func(pointer unsafe.Pointer) AVStreamParseType {
			ptr := (*AVStreamParseType)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVStreamParseType) {
			ptr := (*AVStreamParseType)(pointer)
			*ptr = value
		},
	}
}

func AllocAVStreamParseTypeArray(size uint64) *Array[AVStreamParseType] {
	return ToAVStreamParseTypeArray(AVCalloc(size, SizeOfAVStreamParseType))
}

const (
	// AVStreamParseNone wraps AVSTREAM_PARSE_NONE.
	AVStreamParseNone AVStreamParseType = C.AVSTREAM_PARSE_NONE
	// AVStreamParseFull wraps AVSTREAM_PARSE_FULL.
	//
	//	full parsing and repack
	AVStreamParseFull AVStreamParseType = C.AVSTREAM_PARSE_FULL
	// AVStreamParseHeaders wraps AVSTREAM_PARSE_HEADERS.
	//
	//	Only parse headers, do not repack.
	AVStreamParseHeaders AVStreamParseType = C.AVSTREAM_PARSE_HEADERS
	// AVStreamParseTimestamps wraps AVSTREAM_PARSE_TIMESTAMPS.
	//
	//	full parsing and interpolation of timestamps for frames not starting on a packet boundary
	AVStreamParseTimestamps AVStreamParseType = C.AVSTREAM_PARSE_TIMESTAMPS
	// AVStreamParseFullOnce wraps AVSTREAM_PARSE_FULL_ONCE.
	//
	//	full parsing and repack of the first frame only, only implemented for H.264 currently
	AVStreamParseFullOnce AVStreamParseType = C.AVSTREAM_PARSE_FULL_ONCE
	// AVStreamParseFullRaw wraps AVSTREAM_PARSE_FULL_RAW.
	/*
	   < full parsing and repack with timestamp and position generation by parser for raw
	   this assumes that each packet in the file contains no demuxer level headers and
	   just codec level data, otherwise position generation would fail
	*/
	AVStreamParseFullRaw AVStreamParseType = C.AVSTREAM_PARSE_FULL_RAW
)

// --- Enum AVDurationEstimationMethod ---

// AVDurationEstimationMethod wraps AVDurationEstimationMethod.
/*
  The duration of a video can be estimated through various ways, and this enum can be used
  to know how the duration was estimated.
*/
type AVDurationEstimationMethod C.enum_AVDurationEstimationMethod

const SizeOfAVDurationEstimationMethod = C.sizeof_enum_AVDurationEstimationMethod

func ToAVDurationEstimationMethodArray(ptr unsafe.Pointer) *Array[AVDurationEstimationMethod] {
	if ptr == nil {
		return nil
	}

	return &Array[AVDurationEstimationMethod]{
		elemSize: SizeOfAVDurationEstimationMethod,
		loadPtr: func(pointer unsafe.Pointer) AVDurationEstimationMethod {
			ptr := (*AVDurationEstimationMethod)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVDurationEstimationMethod) {
			ptr := (*AVDurationEstimationMethod)(pointer)
			*ptr = value
		},
	}
}

func AllocAVDurationEstimationMethodArray(size uint64) *Array[AVDurationEstimationMethod] {
	return ToAVDurationEstimationMethodArray(AVCalloc(size, SizeOfAVDurationEstimationMethod))
}

const (
	// AVFmtDurationFromPts wraps AVFMT_DURATION_FROM_PTS.
	//
	//	Duration accurately estimated from PTSes
	AVFmtDurationFromPts AVDurationEstimationMethod = C.AVFMT_DURATION_FROM_PTS
	// AVFmtDurationFromStream wraps AVFMT_DURATION_FROM_STREAM.
	//
	//	Duration estimated from a stream with a known duration
	AVFmtDurationFromStream AVDurationEstimationMethod = C.AVFMT_DURATION_FROM_STREAM
	// AVFmtDurationFromBitrate wraps AVFMT_DURATION_FROM_BITRATE.
	//
	//	Duration estimated from bitrate (less accurate)
	AVFmtDurationFromBitrate AVDurationEstimationMethod = C.AVFMT_DURATION_FROM_BITRATE
)

// --- Enum AVTimebaseSource ---

// AVTimebaseSource wraps AVTimebaseSource.
type AVTimebaseSource C.enum_AVTimebaseSource

const SizeOfAVTimebaseSource = C.sizeof_enum_AVTimebaseSource

func ToAVTimebaseSourceArray(ptr unsafe.Pointer) *Array[AVTimebaseSource] {
	if ptr == nil {
		return nil
	}

	return &Array[AVTimebaseSource]{
		elemSize: SizeOfAVTimebaseSource,
		loadPtr: func(pointer unsafe.Pointer) AVTimebaseSource {
			ptr := (*AVTimebaseSource)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVTimebaseSource) {
			ptr := (*AVTimebaseSource)(pointer)
			*ptr = value
		},
	}
}

func AllocAVTimebaseSourceArray(size uint64) *Array[AVTimebaseSource] {
	return ToAVTimebaseSourceArray(AVCalloc(size, SizeOfAVTimebaseSource))
}

const (
	// AVFmtTbcfAuto wraps AVFMT_TBCF_AUTO.
	AVFmtTbcfAuto AVTimebaseSource = C.AVFMT_TBCF_AUTO
	// AVFmtTbcfDecoder wraps AVFMT_TBCF_DECODER.
	AVFmtTbcfDecoder AVTimebaseSource = C.AVFMT_TBCF_DECODER
	// AVFmtTbcfDemuxer wraps AVFMT_TBCF_DEMUXER.
	AVFmtTbcfDemuxer AVTimebaseSource = C.AVFMT_TBCF_DEMUXER
	// AVFmtTbcfRFramerate wraps AVFMT_TBCF_R_FRAMERATE.
	AVFmtTbcfRFramerate AVTimebaseSource = C.AVFMT_TBCF_R_FRAMERATE
)

// --- Enum AVIODirEntryType ---

// AVIODirEntryType wraps AVIODirEntryType.
//
//	Directory entry types.
type AVIODirEntryType C.enum_AVIODirEntryType

const SizeOfAVIODirEntryType = C.sizeof_enum_AVIODirEntryType

func ToAVIODirEntryTypeArray(ptr unsafe.Pointer) *Array[AVIODirEntryType] {
	if ptr == nil {
		return nil
	}

	return &Array[AVIODirEntryType]{
		elemSize: SizeOfAVIODirEntryType,
		loadPtr: func(pointer unsafe.Pointer) AVIODirEntryType {
			ptr := (*AVIODirEntryType)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVIODirEntryType) {
			ptr := (*AVIODirEntryType)(pointer)
			*ptr = value
		},
	}
}

func AllocAVIODirEntryTypeArray(size uint64) *Array[AVIODirEntryType] {
	return ToAVIODirEntryTypeArray(AVCalloc(size, SizeOfAVIODirEntryType))
}

const (
	// AVIOEntryUnknown wraps AVIO_ENTRY_UNKNOWN.
	AVIOEntryUnknown AVIODirEntryType = C.AVIO_ENTRY_UNKNOWN
	// AVIOEntryBlockDevice wraps AVIO_ENTRY_BLOCK_DEVICE.
	AVIOEntryBlockDevice AVIODirEntryType = C.AVIO_ENTRY_BLOCK_DEVICE
	// AVIOEntryCharacterDevice wraps AVIO_ENTRY_CHARACTER_DEVICE.
	AVIOEntryCharacterDevice AVIODirEntryType = C.AVIO_ENTRY_CHARACTER_DEVICE
	// AVIOEntryDirectory wraps AVIO_ENTRY_DIRECTORY.
	AVIOEntryDirectory AVIODirEntryType = C.AVIO_ENTRY_DIRECTORY
	// AVIOEntryNamedPipe wraps AVIO_ENTRY_NAMED_PIPE.
	AVIOEntryNamedPipe AVIODirEntryType = C.AVIO_ENTRY_NAMED_PIPE
	// AVIOEntrySymbolicLink wraps AVIO_ENTRY_SYMBOLIC_LINK.
	AVIOEntrySymbolicLink AVIODirEntryType = C.AVIO_ENTRY_SYMBOLIC_LINK
	// AVIOEntrySocket wraps AVIO_ENTRY_SOCKET.
	AVIOEntrySocket AVIODirEntryType = C.AVIO_ENTRY_SOCKET
	// AVIOEntryFile wraps AVIO_ENTRY_FILE.
	AVIOEntryFile AVIODirEntryType = C.AVIO_ENTRY_FILE
	// AVIOEntryServer wraps AVIO_ENTRY_SERVER.
	AVIOEntryServer AVIODirEntryType = C.AVIO_ENTRY_SERVER
	// AVIOEntryShare wraps AVIO_ENTRY_SHARE.
	AVIOEntryShare AVIODirEntryType = C.AVIO_ENTRY_SHARE
	// AVIOEntryWorkgroup wraps AVIO_ENTRY_WORKGROUP.
	AVIOEntryWorkgroup AVIODirEntryType = C.AVIO_ENTRY_WORKGROUP
)

// --- Enum AVIODataMarkerType ---

// AVIODataMarkerType wraps AVIODataMarkerType.
/*
  Different data types that can be returned via the AVIO
  write_data_type callback.
*/
type AVIODataMarkerType C.enum_AVIODataMarkerType

const SizeOfAVIODataMarkerType = C.sizeof_enum_AVIODataMarkerType

func ToAVIODataMarkerTypeArray(ptr unsafe.Pointer) *Array[AVIODataMarkerType] {
	if ptr == nil {
		return nil
	}

	return &Array[AVIODataMarkerType]{
		elemSize: SizeOfAVIODataMarkerType,
		loadPtr: func(pointer unsafe.Pointer) AVIODataMarkerType {
			ptr := (*AVIODataMarkerType)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVIODataMarkerType) {
			ptr := (*AVIODataMarkerType)(pointer)
			*ptr = value
		},
	}
}

func AllocAVIODataMarkerTypeArray(size uint64) *Array[AVIODataMarkerType] {
	return ToAVIODataMarkerTypeArray(AVCalloc(size, SizeOfAVIODataMarkerType))
}

const (
	// AVIODataMarkerHeader wraps AVIO_DATA_MARKER_HEADER.
	//
	//	Header data; this needs to be present for the stream to be decodeable.
	AVIODataMarkerHeader AVIODataMarkerType = C.AVIO_DATA_MARKER_HEADER
	// AVIODataMarkerSyncPoint wraps AVIO_DATA_MARKER_SYNC_POINT.
	/*
	   A point in the output bytestream where a decoder can start decoding
	   (i.e. a keyframe). A demuxer/decoder given the data flagged with
	   AVIO_DATA_MARKER_HEADER, followed by any AVIO_DATA_MARKER_SYNC_POINT,
	   should give decodeable results.
	*/
	AVIODataMarkerSyncPoint AVIODataMarkerType = C.AVIO_DATA_MARKER_SYNC_POINT
	// AVIODataMarkerBoundaryPoint wraps AVIO_DATA_MARKER_BOUNDARY_POINT.
	/*
	   A point in the output bytestream where a demuxer can start parsing
	   (for non self synchronizing bytestream formats). That is, any
	   non-keyframe packet start point.
	*/
	AVIODataMarkerBoundaryPoint AVIODataMarkerType = C.AVIO_DATA_MARKER_BOUNDARY_POINT
	// AVIODataMarkerUnknown wraps AVIO_DATA_MARKER_UNKNOWN.
	/*
	   This is any, unlabelled data. It can either be a muxer not marking
	   any positions at all, it can be an actual boundary/sync point
	   that the muxer chooses not to mark, or a later part of a packet/fragment
	   that is cut into multiple write callbacks due to limited IO buffer size.
	*/
	AVIODataMarkerUnknown AVIODataMarkerType = C.AVIO_DATA_MARKER_UNKNOWN
	// AVIODataMarkerTrailer wraps AVIO_DATA_MARKER_TRAILER.
	/*
	   Trailer data, which doesn't contain actual content, but only for
	   finalizing the output file.
	*/
	AVIODataMarkerTrailer AVIODataMarkerType = C.AVIO_DATA_MARKER_TRAILER
	// AVIODataMarkerFlushPoint wraps AVIO_DATA_MARKER_FLUSH_POINT.
	/*
	   A point in the output bytestream where the underlying AVIOContext might
	   flush the buffer depending on latency or buffering requirements. Typically
	   means the end of a packet.
	*/
	AVIODataMarkerFlushPoint AVIODataMarkerType = C.AVIO_DATA_MARKER_FLUSH_POINT
)

// --- Enum AVMediaType ---

// AVMediaType wraps AVMediaType.
//
//	@brief Media Type
type AVMediaType C.enum_AVMediaType

const SizeOfAVMediaType = C.sizeof_enum_AVMediaType

func ToAVMediaTypeArray(ptr unsafe.Pointer) *Array[AVMediaType] {
	if ptr == nil {
		return nil
	}

	return &Array[AVMediaType]{
		elemSize: SizeOfAVMediaType,
		loadPtr: func(pointer unsafe.Pointer) AVMediaType {
			ptr := (*AVMediaType)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVMediaType) {
			ptr := (*AVMediaType)(pointer)
			*ptr = value
		},
	}
}

func AllocAVMediaTypeArray(size uint64) *Array[AVMediaType] {
	return ToAVMediaTypeArray(AVCalloc(size, SizeOfAVMediaType))
}

const (
	// AVMediaTypeUnknown wraps AVMEDIA_TYPE_UNKNOWN.
	//
	//	Usually treated as AVMEDIA_TYPE_DATA
	AVMediaTypeUnknown AVMediaType = C.AVMEDIA_TYPE_UNKNOWN
	// AVMediaTypeVideo wraps AVMEDIA_TYPE_VIDEO.
	AVMediaTypeVideo AVMediaType = C.AVMEDIA_TYPE_VIDEO
	// AVMediaTypeAudio wraps AVMEDIA_TYPE_AUDIO.
	AVMediaTypeAudio AVMediaType = C.AVMEDIA_TYPE_AUDIO
	// AVMediaTypeData wraps AVMEDIA_TYPE_DATA.
	//
	//	Opaque data information usually continuous
	AVMediaTypeData AVMediaType = C.AVMEDIA_TYPE_DATA
	// AVMediaTypeSubtitle wraps AVMEDIA_TYPE_SUBTITLE.
	AVMediaTypeSubtitle AVMediaType = C.AVMEDIA_TYPE_SUBTITLE
	// AVMediaTypeAttachment wraps AVMEDIA_TYPE_ATTACHMENT.
	//
	//	Opaque data information usually sparse
	AVMediaTypeAttachment AVMediaType = C.AVMEDIA_TYPE_ATTACHMENT
	// AVMediaTypeNb wraps AVMEDIA_TYPE_NB.
	AVMediaTypeNb AVMediaType = C.AVMEDIA_TYPE_NB
)

// --- Enum AVPictureType ---

// AVPictureType wraps AVPictureType.
type AVPictureType C.enum_AVPictureType

const SizeOfAVPictureType = C.sizeof_enum_AVPictureType

func ToAVPictureTypeArray(ptr unsafe.Pointer) *Array[AVPictureType] {
	if ptr == nil {
		return nil
	}

	return &Array[AVPictureType]{
		elemSize: SizeOfAVPictureType,
		loadPtr: func(pointer unsafe.Pointer) AVPictureType {
			ptr := (*AVPictureType)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVPictureType) {
			ptr := (*AVPictureType)(pointer)
			*ptr = value
		},
	}
}

func AllocAVPictureTypeArray(size uint64) *Array[AVPictureType] {
	return ToAVPictureTypeArray(AVCalloc(size, SizeOfAVPictureType))
}

const (
	// AVPictureTypeNone wraps AV_PICTURE_TYPE_NONE.
	//
	//	Undefined
	AVPictureTypeNone AVPictureType = C.AV_PICTURE_TYPE_NONE
	// AVPictureTypeI wraps AV_PICTURE_TYPE_I.
	//
	//	Intra
	AVPictureTypeI AVPictureType = C.AV_PICTURE_TYPE_I
	// AVPictureTypeP wraps AV_PICTURE_TYPE_P.
	//
	//	Predicted
	AVPictureTypeP AVPictureType = C.AV_PICTURE_TYPE_P
	// AVPictureTypeB wraps AV_PICTURE_TYPE_B.
	//
	//	Bi-dir predicted
	AVPictureTypeB AVPictureType = C.AV_PICTURE_TYPE_B
	// AVPictureTypeS wraps AV_PICTURE_TYPE_S.
	//
	//	S(GMC)-VOP MPEG-4
	AVPictureTypeS AVPictureType = C.AV_PICTURE_TYPE_S
	// AVPictureTypeSi wraps AV_PICTURE_TYPE_SI.
	//
	//	Switching Intra
	AVPictureTypeSi AVPictureType = C.AV_PICTURE_TYPE_SI
	// AVPictureTypeSp wraps AV_PICTURE_TYPE_SP.
	//
	//	Switching Predicted
	AVPictureTypeSp AVPictureType = C.AV_PICTURE_TYPE_SP
	// AVPictureTypeBi wraps AV_PICTURE_TYPE_BI.
	//
	//	BI type
	AVPictureTypeBi AVPictureType = C.AV_PICTURE_TYPE_BI
)

// --- Enum AVChannel ---

// AVChannel wraps AVChannel.
type AVChannel C.enum_AVChannel

const SizeOfAVChannel = C.sizeof_enum_AVChannel

func ToAVChannelArray(ptr unsafe.Pointer) *Array[AVChannel] {
	if ptr == nil {
		return nil
	}

	return &Array[AVChannel]{
		elemSize: SizeOfAVChannel,
		loadPtr: func(pointer unsafe.Pointer) AVChannel {
			ptr := (*AVChannel)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVChannel) {
			ptr := (*AVChannel)(pointer)
			*ptr = value
		},
	}
}

func AllocAVChannelArray(size uint64) *Array[AVChannel] {
	return ToAVChannelArray(AVCalloc(size, SizeOfAVChannel))
}

const (
	// AVChanNone wraps AV_CHAN_NONE.
	AVChanNone AVChannel = C.AV_CHAN_NONE
	// AVChanFrontLeft wraps AV_CHAN_FRONT_LEFT.
	AVChanFrontLeft AVChannel = C.AV_CHAN_FRONT_LEFT
	// AVChanFrontRight wraps AV_CHAN_FRONT_RIGHT.
	AVChanFrontRight AVChannel = C.AV_CHAN_FRONT_RIGHT
	// AVChanFrontCenter wraps AV_CHAN_FRONT_CENTER.
	AVChanFrontCenter AVChannel = C.AV_CHAN_FRONT_CENTER
	// AVChanLowFrequency wraps AV_CHAN_LOW_FREQUENCY.
	AVChanLowFrequency AVChannel = C.AV_CHAN_LOW_FREQUENCY
	// AVChanBackLeft wraps AV_CHAN_BACK_LEFT.
	AVChanBackLeft AVChannel = C.AV_CHAN_BACK_LEFT
	// AVChanBackRight wraps AV_CHAN_BACK_RIGHT.
	AVChanBackRight AVChannel = C.AV_CHAN_BACK_RIGHT
	// AVChanFrontLeftOfCenter wraps AV_CHAN_FRONT_LEFT_OF_CENTER.
	AVChanFrontLeftOfCenter AVChannel = C.AV_CHAN_FRONT_LEFT_OF_CENTER
	// AVChanFrontRightOfCenter wraps AV_CHAN_FRONT_RIGHT_OF_CENTER.
	AVChanFrontRightOfCenter AVChannel = C.AV_CHAN_FRONT_RIGHT_OF_CENTER
	// AVChanBackCenter wraps AV_CHAN_BACK_CENTER.
	AVChanBackCenter AVChannel = C.AV_CHAN_BACK_CENTER
	// AVChanSideLeft wraps AV_CHAN_SIDE_LEFT.
	AVChanSideLeft AVChannel = C.AV_CHAN_SIDE_LEFT
	// AVChanSideRight wraps AV_CHAN_SIDE_RIGHT.
	AVChanSideRight AVChannel = C.AV_CHAN_SIDE_RIGHT
	// AVChanTopCenter wraps AV_CHAN_TOP_CENTER.
	AVChanTopCenter AVChannel = C.AV_CHAN_TOP_CENTER
	// AVChanTopFrontLeft wraps AV_CHAN_TOP_FRONT_LEFT.
	AVChanTopFrontLeft AVChannel = C.AV_CHAN_TOP_FRONT_LEFT
	// AVChanTopFrontCenter wraps AV_CHAN_TOP_FRONT_CENTER.
	AVChanTopFrontCenter AVChannel = C.AV_CHAN_TOP_FRONT_CENTER
	// AVChanTopFrontRight wraps AV_CHAN_TOP_FRONT_RIGHT.
	AVChanTopFrontRight AVChannel = C.AV_CHAN_TOP_FRONT_RIGHT
	// AVChanTopBackLeft wraps AV_CHAN_TOP_BACK_LEFT.
	AVChanTopBackLeft AVChannel = C.AV_CHAN_TOP_BACK_LEFT
	// AVChanTopBackCenter wraps AV_CHAN_TOP_BACK_CENTER.
	AVChanTopBackCenter AVChannel = C.AV_CHAN_TOP_BACK_CENTER
	// AVChanTopBackRight wraps AV_CHAN_TOP_BACK_RIGHT.
	AVChanTopBackRight AVChannel = C.AV_CHAN_TOP_BACK_RIGHT
	// AVChanStereoLeft wraps AV_CHAN_STEREO_LEFT.
	//
	//	Stereo downmix.
	AVChanStereoLeft AVChannel = C.AV_CHAN_STEREO_LEFT
	// AVChanStereoRight wraps AV_CHAN_STEREO_RIGHT.
	//
	//	See above.
	AVChanStereoRight AVChannel = C.AV_CHAN_STEREO_RIGHT
	// AVChanWideLeft wraps AV_CHAN_WIDE_LEFT.
	//
	//	See above.
	AVChanWideLeft AVChannel = C.AV_CHAN_WIDE_LEFT
	// AVChanWideRight wraps AV_CHAN_WIDE_RIGHT.
	//
	//	See above.
	AVChanWideRight AVChannel = C.AV_CHAN_WIDE_RIGHT
	// AVChanSurroundDirectLeft wraps AV_CHAN_SURROUND_DIRECT_LEFT.
	//
	//	See above.
	AVChanSurroundDirectLeft AVChannel = C.AV_CHAN_SURROUND_DIRECT_LEFT
	// AVChanSurroundDirectRight wraps AV_CHAN_SURROUND_DIRECT_RIGHT.
	//
	//	See above.
	AVChanSurroundDirectRight AVChannel = C.AV_CHAN_SURROUND_DIRECT_RIGHT
	// AVChanLowFrequency2 wraps AV_CHAN_LOW_FREQUENCY_2.
	//
	//	See above.
	AVChanLowFrequency2 AVChannel = C.AV_CHAN_LOW_FREQUENCY_2
	// AVChanTopSideLeft wraps AV_CHAN_TOP_SIDE_LEFT.
	//
	//	See above.
	AVChanTopSideLeft AVChannel = C.AV_CHAN_TOP_SIDE_LEFT
	// AVChanTopSideRight wraps AV_CHAN_TOP_SIDE_RIGHT.
	//
	//	See above.
	AVChanTopSideRight AVChannel = C.AV_CHAN_TOP_SIDE_RIGHT
	// AVChanBottomFrontCenter wraps AV_CHAN_BOTTOM_FRONT_CENTER.
	//
	//	See above.
	AVChanBottomFrontCenter AVChannel = C.AV_CHAN_BOTTOM_FRONT_CENTER
	// AVChanBottomFrontLeft wraps AV_CHAN_BOTTOM_FRONT_LEFT.
	//
	//	See above.
	AVChanBottomFrontLeft AVChannel = C.AV_CHAN_BOTTOM_FRONT_LEFT
	// AVChanBottomFrontRight wraps AV_CHAN_BOTTOM_FRONT_RIGHT.
	//
	//	See above.
	AVChanBottomFrontRight AVChannel = C.AV_CHAN_BOTTOM_FRONT_RIGHT
	// AVChanUnused wraps AV_CHAN_UNUSED.
	//
	//	Channel is empty can be safely skipped.
	AVChanUnused AVChannel = C.AV_CHAN_UNUSED
	// AVChanUnknown wraps AV_CHAN_UNKNOWN.
	//
	//	Channel contains data, but its position is unknown.
	AVChanUnknown AVChannel = C.AV_CHAN_UNKNOWN
	// AVChanAmbisonicBase wraps AV_CHAN_AMBISONIC_BASE.
	/*
	   Range of channels between AV_CHAN_AMBISONIC_BASE and
	   AV_CHAN_AMBISONIC_END represent Ambisonic components using the ACN system.

	   Given a channel id `<i>` between AV_CHAN_AMBISONIC_BASE and
	   AV_CHAN_AMBISONIC_END (inclusive), the ACN index of the channel `<n>` is
	   `<n> = <i> - AV_CHAN_AMBISONIC_BASE`.

	   @note these values are only used for AV_CHANNEL_ORDER_CUSTOM channel
	   orderings, the AV_CHANNEL_ORDER_AMBISONIC ordering orders the channels
	   implicitly by their position in the stream.
	*/
	AVChanAmbisonicBase AVChannel = C.AV_CHAN_AMBISONIC_BASE
	// AVChanAmbisonicEnd wraps AV_CHAN_AMBISONIC_END.
	/*
	   // leave space for 1024 ids, which correspond to maximum order-32 harmonics,
	   // which should be enough for the foreseeable use cases
	*/
	AVChanAmbisonicEnd AVChannel = C.AV_CHAN_AMBISONIC_END
)

// --- Enum AVChannelOrder ---

// AVChannelOrder wraps AVChannelOrder.
type AVChannelOrder C.enum_AVChannelOrder

const SizeOfAVChannelOrder = C.sizeof_enum_AVChannelOrder

func ToAVChannelOrderArray(ptr unsafe.Pointer) *Array[AVChannelOrder] {
	if ptr == nil {
		return nil
	}

	return &Array[AVChannelOrder]{
		elemSize: SizeOfAVChannelOrder,
		loadPtr: func(pointer unsafe.Pointer) AVChannelOrder {
			ptr := (*AVChannelOrder)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVChannelOrder) {
			ptr := (*AVChannelOrder)(pointer)
			*ptr = value
		},
	}
}

func AllocAVChannelOrderArray(size uint64) *Array[AVChannelOrder] {
	return ToAVChannelOrderArray(AVCalloc(size, SizeOfAVChannelOrder))
}

const (
	// AVChannelOrderUnspec wraps AV_CHANNEL_ORDER_UNSPEC.
	/*
	   Only the channel count is specified, without any further information
	   about the channel order.
	*/
	AVChannelOrderUnspec AVChannelOrder = C.AV_CHANNEL_ORDER_UNSPEC
	// AVChannelOrderNative wraps AV_CHANNEL_ORDER_NATIVE.
	/*
	   The native channel order, i.e. the channels are in the same order in
	   which they are defined in the AVChannel enum. This supports up to 63
	   different channels.
	*/
	AVChannelOrderNative AVChannelOrder = C.AV_CHANNEL_ORDER_NATIVE
	// AVChannelOrderCustom wraps AV_CHANNEL_ORDER_CUSTOM.
	/*
	   The channel order does not correspond to any other predefined order and
	   is stored as an explicit map. For example, this could be used to support
	   layouts with 64 or more channels, or with empty/skipped (AV_CHAN_SILENCE)
	   channels at arbitrary positions.
	*/
	AVChannelOrderCustom AVChannelOrder = C.AV_CHANNEL_ORDER_CUSTOM
	// AVChannelOrderAmbisonic wraps AV_CHANNEL_ORDER_AMBISONIC.
	/*
	   The audio is represented as the decomposition of the sound field into
	   spherical harmonics. Each channel corresponds to a single expansion
	   component. Channels are ordered according to ACN (Ambisonic Channel
	   Number).

	   The channel with the index n in the stream contains the spherical
	   harmonic of degree l and order m given by
	   @code{.unparsed}
	     l   = floor(sqrt(n)),
	     m   = n - l * (l + 1).
	   @endcode

	   Conversely given a spherical harmonic of degree l and order m, the
	   corresponding channel index n is given by
	   @code{.unparsed}
	     n = l * (l + 1) + m.
	   @endcode

	   Normalization is assumed to be SN3D (Schmidt Semi-Normalization)
	   as defined in AmbiX format $ 2.1.
	*/
	AVChannelOrderAmbisonic AVChannelOrder = C.AV_CHANNEL_ORDER_AMBISONIC
)

// --- Enum AVMatrixEncoding ---

// AVMatrixEncoding wraps AVMatrixEncoding.
type AVMatrixEncoding C.enum_AVMatrixEncoding

const SizeOfAVMatrixEncoding = C.sizeof_enum_AVMatrixEncoding

func ToAVMatrixEncodingArray(ptr unsafe.Pointer) *Array[AVMatrixEncoding] {
	if ptr == nil {
		return nil
	}

	return &Array[AVMatrixEncoding]{
		elemSize: SizeOfAVMatrixEncoding,
		loadPtr: func(pointer unsafe.Pointer) AVMatrixEncoding {
			ptr := (*AVMatrixEncoding)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVMatrixEncoding) {
			ptr := (*AVMatrixEncoding)(pointer)
			*ptr = value
		},
	}
}

func AllocAVMatrixEncodingArray(size uint64) *Array[AVMatrixEncoding] {
	return ToAVMatrixEncodingArray(AVCalloc(size, SizeOfAVMatrixEncoding))
}

const (
	// AVMatrixEncodingNone wraps AV_MATRIX_ENCODING_NONE.
	AVMatrixEncodingNone AVMatrixEncoding = C.AV_MATRIX_ENCODING_NONE
	// AVMatrixEncodingDolby wraps AV_MATRIX_ENCODING_DOLBY.
	AVMatrixEncodingDolby AVMatrixEncoding = C.AV_MATRIX_ENCODING_DOLBY
	// AVMatrixEncodingDplii wraps AV_MATRIX_ENCODING_DPLII.
	AVMatrixEncodingDplii AVMatrixEncoding = C.AV_MATRIX_ENCODING_DPLII
	// AVMatrixEncodingDpliix wraps AV_MATRIX_ENCODING_DPLIIX.
	AVMatrixEncodingDpliix AVMatrixEncoding = C.AV_MATRIX_ENCODING_DPLIIX
	// AVMatrixEncodingDpliiz wraps AV_MATRIX_ENCODING_DPLIIZ.
	AVMatrixEncodingDpliiz AVMatrixEncoding = C.AV_MATRIX_ENCODING_DPLIIZ
	// AVMatrixEncodingDolbyex wraps AV_MATRIX_ENCODING_DOLBYEX.
	AVMatrixEncodingDolbyex AVMatrixEncoding = C.AV_MATRIX_ENCODING_DOLBYEX
	// AVMatrixEncodingDolbyheadphone wraps AV_MATRIX_ENCODING_DOLBYHEADPHONE.
	AVMatrixEncodingDolbyheadphone AVMatrixEncoding = C.AV_MATRIX_ENCODING_DOLBYHEADPHONE
	// AVMatrixEncodingNb wraps AV_MATRIX_ENCODING_NB.
	AVMatrixEncodingNb AVMatrixEncoding = C.AV_MATRIX_ENCODING_NB
)

// --- Enum AVFrameSideDataType ---

// AVFrameSideDataType wraps AVFrameSideDataType.
//
//	AVFrame is an abstraction for reference-counted raw multimedia data.
type AVFrameSideDataType C.enum_AVFrameSideDataType

const SizeOfAVFrameSideDataType = C.sizeof_enum_AVFrameSideDataType

func ToAVFrameSideDataTypeArray(ptr unsafe.Pointer) *Array[AVFrameSideDataType] {
	if ptr == nil {
		return nil
	}

	return &Array[AVFrameSideDataType]{
		elemSize: SizeOfAVFrameSideDataType,
		loadPtr: func(pointer unsafe.Pointer) AVFrameSideDataType {
			ptr := (*AVFrameSideDataType)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVFrameSideDataType) {
			ptr := (*AVFrameSideDataType)(pointer)
			*ptr = value
		},
	}
}

func AllocAVFrameSideDataTypeArray(size uint64) *Array[AVFrameSideDataType] {
	return ToAVFrameSideDataTypeArray(AVCalloc(size, SizeOfAVFrameSideDataType))
}

const (
	// AVFrameDataPanscan wraps AV_FRAME_DATA_PANSCAN.
	//
	//	The data is the AVPanScan struct defined in libavcodec.
	AVFrameDataPanscan AVFrameSideDataType = C.AV_FRAME_DATA_PANSCAN
	// AVFrameDataA53Cc wraps AV_FRAME_DATA_A53_CC.
	/*
	   ATSC A53 Part 4 Closed Captions.
	   A53 CC bitstream is stored as uint8_t in AVFrameSideData.data.
	   The number of bytes of CC data is AVFrameSideData.size.
	*/
	AVFrameDataA53Cc AVFrameSideDataType = C.AV_FRAME_DATA_A53_CC
	// AVFrameDataStereo3D wraps AV_FRAME_DATA_STEREO3D.
	/*
	   Stereoscopic 3d metadata.
	   The data is the AVStereo3D struct defined in libavutil/stereo3d.h.
	*/
	AVFrameDataStereo3D AVFrameSideDataType = C.AV_FRAME_DATA_STEREO3D
	// AVFrameDataMatrixencoding wraps AV_FRAME_DATA_MATRIXENCODING.
	//
	//	The data is the AVMatrixEncoding enum defined in libavutil/channel_layout.h.
	AVFrameDataMatrixencoding AVFrameSideDataType = C.AV_FRAME_DATA_MATRIXENCODING
	// AVFrameDataDownmixInfo wraps AV_FRAME_DATA_DOWNMIX_INFO.
	/*
	   Metadata relevant to a downmix procedure.
	   The data is the AVDownmixInfo struct defined in libavutil/downmix_info.h.
	*/
	AVFrameDataDownmixInfo AVFrameSideDataType = C.AV_FRAME_DATA_DOWNMIX_INFO
	// AVFrameDataReplaygain wraps AV_FRAME_DATA_REPLAYGAIN.
	//
	//	ReplayGain information in the form of the AVReplayGain struct.
	AVFrameDataReplaygain AVFrameSideDataType = C.AV_FRAME_DATA_REPLAYGAIN
	// AVFrameDataDisplaymatrix wraps AV_FRAME_DATA_DISPLAYMATRIX.
	/*
	   This side data contains a 3x3 transformation matrix describing an affine
	   transformation that needs to be applied to the frame for correct
	   presentation.

	   See libavutil/display.h for a detailed description of the data.
	*/
	AVFrameDataDisplaymatrix AVFrameSideDataType = C.AV_FRAME_DATA_DISPLAYMATRIX
	// AVFrameDataAfd wraps AV_FRAME_DATA_AFD.
	/*
	   Active Format Description data consisting of a single byte as specified
	   in ETSI TS 101 154 using AVActiveFormatDescription enum.
	*/
	AVFrameDataAfd AVFrameSideDataType = C.AV_FRAME_DATA_AFD
	// AVFrameDataMotionVectors wraps AV_FRAME_DATA_MOTION_VECTORS.
	/*
	   Motion vectors exported by some codecs (on demand through the export_mvs
	   flag set in the libavcodec AVCodecContext flags2 option).
	   The data is the AVMotionVector struct defined in
	   libavutil/motion_vector.h.
	*/
	AVFrameDataMotionVectors AVFrameSideDataType = C.AV_FRAME_DATA_MOTION_VECTORS
	// AVFrameDataSkipSamples wraps AV_FRAME_DATA_SKIP_SAMPLES.
	/*
	   Recommmends skipping the specified number of samples. This is exported
	   only if the "skip_manual" AVOption is set in libavcodec.
	   This has the same format as AV_PKT_DATA_SKIP_SAMPLES.
	   @code
	   u32le number of samples to skip from start of this packet
	   u32le number of samples to skip from end of this packet
	   u8    reason for start skip
	   u8    reason for end   skip (0=padding silence, 1=convergence)
	   @endcode
	*/
	AVFrameDataSkipSamples AVFrameSideDataType = C.AV_FRAME_DATA_SKIP_SAMPLES
	// AVFrameDataAudioServiceType wraps AV_FRAME_DATA_AUDIO_SERVICE_TYPE.
	/*
	   This side data must be associated with an audio frame and corresponds to
	   enum AVAudioServiceType defined in avcodec.h.
	*/
	AVFrameDataAudioServiceType AVFrameSideDataType = C.AV_FRAME_DATA_AUDIO_SERVICE_TYPE
	// AVFrameDataMasteringDisplayMetadata wraps AV_FRAME_DATA_MASTERING_DISPLAY_METADATA.
	/*
	   Mastering display metadata associated with a video frame. The payload is
	   an AVMasteringDisplayMetadata type and contains information about the
	   mastering display color volume.
	*/
	AVFrameDataMasteringDisplayMetadata AVFrameSideDataType = C.AV_FRAME_DATA_MASTERING_DISPLAY_METADATA
	// AVFrameDataGopTimecode wraps AV_FRAME_DATA_GOP_TIMECODE.
	/*
	   The GOP timecode in 25 bit timecode format. Data format is 64-bit integer.
	   This is set on the first frame of a GOP that has a temporal reference of 0.
	*/
	AVFrameDataGopTimecode AVFrameSideDataType = C.AV_FRAME_DATA_GOP_TIMECODE
	// AVFrameDataSpherical wraps AV_FRAME_DATA_SPHERICAL.
	/*
	   The data represents the AVSphericalMapping structure defined in
	   libavutil/spherical.h.
	*/
	AVFrameDataSpherical AVFrameSideDataType = C.AV_FRAME_DATA_SPHERICAL
	// AVFrameDataContentLightLevel wraps AV_FRAME_DATA_CONTENT_LIGHT_LEVEL.
	/*
	   Content light level (based on CTA-861.3). This payload contains data in
	   the form of the AVContentLightMetadata struct.
	*/
	AVFrameDataContentLightLevel AVFrameSideDataType = C.AV_FRAME_DATA_CONTENT_LIGHT_LEVEL
	// AVFrameDataIccProfile wraps AV_FRAME_DATA_ICC_PROFILE.
	/*
	   The data contains an ICC profile as an opaque octet buffer following the
	   format described by ISO 15076-1 with an optional name defined in the
	   metadata key entry "name".
	*/
	AVFrameDataIccProfile AVFrameSideDataType = C.AV_FRAME_DATA_ICC_PROFILE
	// AVFrameDataS12MTimecode wraps AV_FRAME_DATA_S12M_TIMECODE.
	/*
	   Timecode which conforms to SMPTE ST 12-1. The data is an array of 4 uint32_t
	   where the first uint32_t describes how many (1-3) of the other timecodes are used.
	   The timecode format is described in the documentation of av_timecode_get_smpte_from_framenum()
	   function in libavutil/timecode.h.
	*/
	AVFrameDataS12MTimecode AVFrameSideDataType = C.AV_FRAME_DATA_S12M_TIMECODE
	// AVFrameDataDynamicHdrPlus wraps AV_FRAME_DATA_DYNAMIC_HDR_PLUS.
	/*
	   HDR dynamic metadata associated with a video frame. The payload is
	   an AVDynamicHDRPlus type and contains information for color
	   volume transform - application 4 of SMPTE 2094-40:2016 standard.
	*/
	AVFrameDataDynamicHdrPlus AVFrameSideDataType = C.AV_FRAME_DATA_DYNAMIC_HDR_PLUS
	// AVFrameDataRegionsOfInterest wraps AV_FRAME_DATA_REGIONS_OF_INTEREST.
	/*
	   Regions Of Interest, the data is an array of AVRegionOfInterest type, the number of
	   array element is implied by AVFrameSideData.size / AVRegionOfInterest.self_size.
	*/
	AVFrameDataRegionsOfInterest AVFrameSideDataType = C.AV_FRAME_DATA_REGIONS_OF_INTEREST
	// AVFrameDataVideoEncParams wraps AV_FRAME_DATA_VIDEO_ENC_PARAMS.
	//
	//	Encoding parameters for a video frame, as described by AVVideoEncParams.
	AVFrameDataVideoEncParams AVFrameSideDataType = C.AV_FRAME_DATA_VIDEO_ENC_PARAMS
	// AVFrameDataSeiUnregistered wraps AV_FRAME_DATA_SEI_UNREGISTERED.
	/*
	   User data unregistered metadata associated with a video frame.
	   This is the H.26[45] UDU SEI message, and shouldn't be used for any other purpose
	   The data is stored as uint8_t in AVFrameSideData.data which is 16 bytes of
	   uuid_iso_iec_11578 followed by AVFrameSideData.size - 16 bytes of user_data_payload_byte.
	*/
	AVFrameDataSeiUnregistered AVFrameSideDataType = C.AV_FRAME_DATA_SEI_UNREGISTERED
	// AVFrameDataFilmGrainParams wraps AV_FRAME_DATA_FILM_GRAIN_PARAMS.
	/*
	   Film grain parameters for a frame, described by AVFilmGrainParams.
	   Must be present for every frame which should have film grain applied.
	*/
	AVFrameDataFilmGrainParams AVFrameSideDataType = C.AV_FRAME_DATA_FILM_GRAIN_PARAMS
	// AVFrameDataDetectionBboxes wraps AV_FRAME_DATA_DETECTION_BBOXES.
	/*
	   Bounding boxes for object detection and classification,
	   as described by AVDetectionBBoxHeader.
	*/
	AVFrameDataDetectionBboxes AVFrameSideDataType = C.AV_FRAME_DATA_DETECTION_BBOXES
	// AVFrameDataDoviRpuBuffer wraps AV_FRAME_DATA_DOVI_RPU_BUFFER.
	/*
	   Dolby Vision RPU raw data, suitable for passing to x265
	   or other libraries. Array of uint8_t, with NAL emulation
	   bytes intact.
	*/
	AVFrameDataDoviRpuBuffer AVFrameSideDataType = C.AV_FRAME_DATA_DOVI_RPU_BUFFER
	// AVFrameDataDoviMetadata wraps AV_FRAME_DATA_DOVI_METADATA.
	/*
	   Parsed Dolby Vision metadata, suitable for passing to a software
	   implementation. The payload is the AVDOVIMetadata struct defined in
	   libavutil/dovi_meta.h.
	*/
	AVFrameDataDoviMetadata AVFrameSideDataType = C.AV_FRAME_DATA_DOVI_METADATA
	// AVFrameDataDynamicHdrVivid wraps AV_FRAME_DATA_DYNAMIC_HDR_VIVID.
	/*
	   HDR Vivid dynamic metadata associated with a video frame. The payload is
	   an AVDynamicHDRVivid type and contains information for color
	   volume transform - CUVA 005.1-2021.
	*/
	AVFrameDataDynamicHdrVivid AVFrameSideDataType = C.AV_FRAME_DATA_DYNAMIC_HDR_VIVID
	// AVFrameDataAmbientViewingEnvironment wraps AV_FRAME_DATA_AMBIENT_VIEWING_ENVIRONMENT.
	//
	//	Ambient viewing environment metadata, as defined by H.274.
	AVFrameDataAmbientViewingEnvironment AVFrameSideDataType = C.AV_FRAME_DATA_AMBIENT_VIEWING_ENVIRONMENT
)

// --- Enum AVActiveFormatDescription ---

// AVActiveFormatDescription wraps AVActiveFormatDescription.
type AVActiveFormatDescription C.enum_AVActiveFormatDescription

const SizeOfAVActiveFormatDescription = C.sizeof_enum_AVActiveFormatDescription

func ToAVActiveFormatDescriptionArray(ptr unsafe.Pointer) *Array[AVActiveFormatDescription] {
	if ptr == nil {
		return nil
	}

	return &Array[AVActiveFormatDescription]{
		elemSize: SizeOfAVActiveFormatDescription,
		loadPtr: func(pointer unsafe.Pointer) AVActiveFormatDescription {
			ptr := (*AVActiveFormatDescription)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVActiveFormatDescription) {
			ptr := (*AVActiveFormatDescription)(pointer)
			*ptr = value
		},
	}
}

func AllocAVActiveFormatDescriptionArray(size uint64) *Array[AVActiveFormatDescription] {
	return ToAVActiveFormatDescriptionArray(AVCalloc(size, SizeOfAVActiveFormatDescription))
}

const (
	// AVAfdSame wraps AV_AFD_SAME.
	AVAfdSame AVActiveFormatDescription = C.AV_AFD_SAME
	// AVAfd43 wraps AV_AFD_4_3.
	AVAfd43 AVActiveFormatDescription = C.AV_AFD_4_3
	// AVAfd169 wraps AV_AFD_16_9.
	AVAfd169 AVActiveFormatDescription = C.AV_AFD_16_9
	// AVAfd149 wraps AV_AFD_14_9.
	AVAfd149 AVActiveFormatDescription = C.AV_AFD_14_9
	// AVAfd43Sp149 wraps AV_AFD_4_3_SP_14_9.
	AVAfd43Sp149 AVActiveFormatDescription = C.AV_AFD_4_3_SP_14_9
	// AVAfd169Sp149 wraps AV_AFD_16_9_SP_14_9.
	AVAfd169Sp149 AVActiveFormatDescription = C.AV_AFD_16_9_SP_14_9
	// AVAfdSp43 wraps AV_AFD_SP_4_3.
	AVAfdSp43 AVActiveFormatDescription = C.AV_AFD_SP_4_3
)

// --- Enum AVHWDeviceType ---

// AVHWDeviceType wraps AVHWDeviceType.
type AVHWDeviceType C.enum_AVHWDeviceType

const SizeOfAVHWDeviceType = C.sizeof_enum_AVHWDeviceType

func ToAVHWDeviceTypeArray(ptr unsafe.Pointer) *Array[AVHWDeviceType] {
	if ptr == nil {
		return nil
	}

	return &Array[AVHWDeviceType]{
		elemSize: SizeOfAVHWDeviceType,
		loadPtr: func(pointer unsafe.Pointer) AVHWDeviceType {
			ptr := (*AVHWDeviceType)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVHWDeviceType) {
			ptr := (*AVHWDeviceType)(pointer)
			*ptr = value
		},
	}
}

func AllocAVHWDeviceTypeArray(size uint64) *Array[AVHWDeviceType] {
	return ToAVHWDeviceTypeArray(AVCalloc(size, SizeOfAVHWDeviceType))
}

const (
	// AVHwdeviceTypeNone wraps AV_HWDEVICE_TYPE_NONE.
	AVHwdeviceTypeNone AVHWDeviceType = C.AV_HWDEVICE_TYPE_NONE
	// AVHwdeviceTypeVdpau wraps AV_HWDEVICE_TYPE_VDPAU.
	AVHwdeviceTypeVdpau AVHWDeviceType = C.AV_HWDEVICE_TYPE_VDPAU
	// AVHwdeviceTypeCuda wraps AV_HWDEVICE_TYPE_CUDA.
	AVHwdeviceTypeCuda AVHWDeviceType = C.AV_HWDEVICE_TYPE_CUDA
	// AVHwdeviceTypeVaapi wraps AV_HWDEVICE_TYPE_VAAPI.
	AVHwdeviceTypeVaapi AVHWDeviceType = C.AV_HWDEVICE_TYPE_VAAPI
	// AVHwdeviceTypeDxva2 wraps AV_HWDEVICE_TYPE_DXVA2.
	AVHwdeviceTypeDxva2 AVHWDeviceType = C.AV_HWDEVICE_TYPE_DXVA2
	// AVHwdeviceTypeQsv wraps AV_HWDEVICE_TYPE_QSV.
	AVHwdeviceTypeQsv AVHWDeviceType = C.AV_HWDEVICE_TYPE_QSV
	// AVHwdeviceTypeVideotoolbox wraps AV_HWDEVICE_TYPE_VIDEOTOOLBOX.
	AVHwdeviceTypeVideotoolbox AVHWDeviceType = C.AV_HWDEVICE_TYPE_VIDEOTOOLBOX
	// AVHwdeviceTypeD3D11Va wraps AV_HWDEVICE_TYPE_D3D11VA.
	AVHwdeviceTypeD3D11Va AVHWDeviceType = C.AV_HWDEVICE_TYPE_D3D11VA
	// AVHwdeviceTypeDrm wraps AV_HWDEVICE_TYPE_DRM.
	AVHwdeviceTypeDrm AVHWDeviceType = C.AV_HWDEVICE_TYPE_DRM
	// AVHwdeviceTypeOpencl wraps AV_HWDEVICE_TYPE_OPENCL.
	AVHwdeviceTypeOpencl AVHWDeviceType = C.AV_HWDEVICE_TYPE_OPENCL
	// AVHwdeviceTypeMediacodec wraps AV_HWDEVICE_TYPE_MEDIACODEC.
	AVHwdeviceTypeMediacodec AVHWDeviceType = C.AV_HWDEVICE_TYPE_MEDIACODEC
	// AVHwdeviceTypeVulkan wraps AV_HWDEVICE_TYPE_VULKAN.
	AVHwdeviceTypeVulkan AVHWDeviceType = C.AV_HWDEVICE_TYPE_VULKAN
)

// --- Enum AVHWFrameTransferDirection ---

// AVHWFrameTransferDirection wraps AVHWFrameTransferDirection.
type AVHWFrameTransferDirection C.enum_AVHWFrameTransferDirection

const SizeOfAVHWFrameTransferDirection = C.sizeof_enum_AVHWFrameTransferDirection

func ToAVHWFrameTransferDirectionArray(ptr unsafe.Pointer) *Array[AVHWFrameTransferDirection] {
	if ptr == nil {
		return nil
	}

	return &Array[AVHWFrameTransferDirection]{
		elemSize: SizeOfAVHWFrameTransferDirection,
		loadPtr: func(pointer unsafe.Pointer) AVHWFrameTransferDirection {
			ptr := (*AVHWFrameTransferDirection)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVHWFrameTransferDirection) {
			ptr := (*AVHWFrameTransferDirection)(pointer)
			*ptr = value
		},
	}
}

func AllocAVHWFrameTransferDirectionArray(size uint64) *Array[AVHWFrameTransferDirection] {
	return ToAVHWFrameTransferDirectionArray(AVCalloc(size, SizeOfAVHWFrameTransferDirection))
}

const (
	// AVHwframeTransferDirectionFrom wraps AV_HWFRAME_TRANSFER_DIRECTION_FROM.
	//
	//	Transfer the data from the queried hw frame.
	AVHwframeTransferDirectionFrom AVHWFrameTransferDirection = C.AV_HWFRAME_TRANSFER_DIRECTION_FROM
	// AVHwframeTransferDirectionTo wraps AV_HWFRAME_TRANSFER_DIRECTION_TO.
	//
	//	Transfer the data to the queried hw frame.
	AVHwframeTransferDirectionTo AVHWFrameTransferDirection = C.AV_HWFRAME_TRANSFER_DIRECTION_TO
)

// --- Enum AVClassCategory ---

// AVClassCategory wraps AVClassCategory.
type AVClassCategory C.AVClassCategory

const SizeOfAVClassCategory = C.sizeof_AVClassCategory

func ToAVClassCategoryArray(ptr unsafe.Pointer) *Array[AVClassCategory] {
	if ptr == nil {
		return nil
	}

	return &Array[AVClassCategory]{
		elemSize: SizeOfAVClassCategory,
		loadPtr: func(pointer unsafe.Pointer) AVClassCategory {
			ptr := (*AVClassCategory)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVClassCategory) {
			ptr := (*AVClassCategory)(pointer)
			*ptr = value
		},
	}
}

func AllocAVClassCategoryArray(size uint64) *Array[AVClassCategory] {
	return ToAVClassCategoryArray(AVCalloc(size, SizeOfAVClassCategory))
}

const (
	// AVClassCategoryNa wraps AV_CLASS_CATEGORY_NA.
	AVClassCategoryNa AVClassCategory = C.AV_CLASS_CATEGORY_NA
	// AVClassCategoryInput wraps AV_CLASS_CATEGORY_INPUT.
	AVClassCategoryInput AVClassCategory = C.AV_CLASS_CATEGORY_INPUT
	// AVClassCategoryOutput wraps AV_CLASS_CATEGORY_OUTPUT.
	AVClassCategoryOutput AVClassCategory = C.AV_CLASS_CATEGORY_OUTPUT
	// AVClassCategoryMuxer wraps AV_CLASS_CATEGORY_MUXER.
	AVClassCategoryMuxer AVClassCategory = C.AV_CLASS_CATEGORY_MUXER
	// AVClassCategoryDemuxer wraps AV_CLASS_CATEGORY_DEMUXER.
	AVClassCategoryDemuxer AVClassCategory = C.AV_CLASS_CATEGORY_DEMUXER
	// AVClassCategoryEncoder wraps AV_CLASS_CATEGORY_ENCODER.
	AVClassCategoryEncoder AVClassCategory = C.AV_CLASS_CATEGORY_ENCODER
	// AVClassCategoryDecoder wraps AV_CLASS_CATEGORY_DECODER.
	AVClassCategoryDecoder AVClassCategory = C.AV_CLASS_CATEGORY_DECODER
	// AVClassCategoryFilter wraps AV_CLASS_CATEGORY_FILTER.
	AVClassCategoryFilter AVClassCategory = C.AV_CLASS_CATEGORY_FILTER
	// AVClassCategoryBitstreamFilter wraps AV_CLASS_CATEGORY_BITSTREAM_FILTER.
	AVClassCategoryBitstreamFilter AVClassCategory = C.AV_CLASS_CATEGORY_BITSTREAM_FILTER
	// AVClassCategorySwscaler wraps AV_CLASS_CATEGORY_SWSCALER.
	AVClassCategorySwscaler AVClassCategory = C.AV_CLASS_CATEGORY_SWSCALER
	// AVClassCategorySwresampler wraps AV_CLASS_CATEGORY_SWRESAMPLER.
	AVClassCategorySwresampler AVClassCategory = C.AV_CLASS_CATEGORY_SWRESAMPLER
	// AVClassCategoryDeviceVideoOutput wraps AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT.
	AVClassCategoryDeviceVideoOutput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT
	// AVClassCategoryDeviceVideoInput wraps AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT.
	AVClassCategoryDeviceVideoInput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT
	// AVClassCategoryDeviceAudioOutput wraps AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT.
	AVClassCategoryDeviceAudioOutput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT
	// AVClassCategoryDeviceAudioInput wraps AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT.
	AVClassCategoryDeviceAudioInput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT
	// AVClassCategoryDeviceOutput wraps AV_CLASS_CATEGORY_DEVICE_OUTPUT.
	AVClassCategoryDeviceOutput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_OUTPUT
	// AVClassCategoryDeviceInput wraps AV_CLASS_CATEGORY_DEVICE_INPUT.
	AVClassCategoryDeviceInput AVClassCategory = C.AV_CLASS_CATEGORY_DEVICE_INPUT
	// AVClassCategoryNb wraps AV_CLASS_CATEGORY_NB.
	//
	//	not part of ABI/API
	AVClassCategoryNb AVClassCategory = C.AV_CLASS_CATEGORY_NB
)

// --- Enum AVRounding ---

// AVRounding wraps AVRounding.
//
//	Rounding methods.
type AVRounding C.enum_AVRounding

const SizeOfAVRounding = C.sizeof_enum_AVRounding

func ToAVRoundingArray(ptr unsafe.Pointer) *Array[AVRounding] {
	if ptr == nil {
		return nil
	}

	return &Array[AVRounding]{
		elemSize: SizeOfAVRounding,
		loadPtr: func(pointer unsafe.Pointer) AVRounding {
			ptr := (*AVRounding)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVRounding) {
			ptr := (*AVRounding)(pointer)
			*ptr = value
		},
	}
}

func AllocAVRoundingArray(size uint64) *Array[AVRounding] {
	return ToAVRoundingArray(AVCalloc(size, SizeOfAVRounding))
}

const (
	// AVRoundZero wraps AV_ROUND_ZERO.
	//
	//	Round toward zero.
	AVRoundZero AVRounding = C.AV_ROUND_ZERO
	// AVRoundInf wraps AV_ROUND_INF.
	//
	//	Round away from zero.
	AVRoundInf AVRounding = C.AV_ROUND_INF
	// AVRoundDown wraps AV_ROUND_DOWN.
	//
	//	Round toward -infinity.
	AVRoundDown AVRounding = C.AV_ROUND_DOWN
	// AVRoundUp wraps AV_ROUND_UP.
	//
	//	Round toward +infinity.
	AVRoundUp AVRounding = C.AV_ROUND_UP
	// AVRoundNearInf wraps AV_ROUND_NEAR_INF.
	//
	//	Round to nearest and halfway cases away from zero.
	AVRoundNearInf AVRounding = C.AV_ROUND_NEAR_INF
	// AVRoundPassMinmax wraps AV_ROUND_PASS_MINMAX.
	/*
	   Flag telling rescaling functions to pass `INT64_MIN`/`MAX` through
	   unchanged, avoiding special cases for #AV_NOPTS_VALUE.

	   Unlike other values of the enumeration AVRounding, this value is a
	   bitmask that must be used in conjunction with another value of the
	   enumeration through a bitwise OR, in order to set behavior for normal
	   cases.

	   @code{.c}
	   av_rescale_rnd(3, 1, 2, AV_ROUND_UP | AV_ROUND_PASS_MINMAX);
	   // Rescaling 3:
	   //     Calculating 3 * 1 / 2
	   //     3 / 2 is rounded up to 2
	   //     => 2

	   av_rescale_rnd(AV_NOPTS_VALUE, 1, 2, AV_ROUND_UP | AV_ROUND_PASS_MINMAX);
	   // Rescaling AV_NOPTS_VALUE:
	   //     AV_NOPTS_VALUE == INT64_MIN
	   //     AV_NOPTS_VALUE is passed through
	   //     => AV_NOPTS_VALUE
	   @endcode
	*/
	AVRoundPassMinmax AVRounding = C.AV_ROUND_PASS_MINMAX
)

// --- Enum AVOptionType ---

// AVOptionType wraps AVOptionType.
type AVOptionType C.enum_AVOptionType

const SizeOfAVOptionType = C.sizeof_enum_AVOptionType

func ToAVOptionTypeArray(ptr unsafe.Pointer) *Array[AVOptionType] {
	if ptr == nil {
		return nil
	}

	return &Array[AVOptionType]{
		elemSize: SizeOfAVOptionType,
		loadPtr: func(pointer unsafe.Pointer) AVOptionType {
			ptr := (*AVOptionType)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVOptionType) {
			ptr := (*AVOptionType)(pointer)
			*ptr = value
		},
	}
}

func AllocAVOptionTypeArray(size uint64) *Array[AVOptionType] {
	return ToAVOptionTypeArray(AVCalloc(size, SizeOfAVOptionType))
}

const (
	// AVOptTypeFlags wraps AV_OPT_TYPE_FLAGS.
	AVOptTypeFlags AVOptionType = C.AV_OPT_TYPE_FLAGS
	// AVOptTypeInt wraps AV_OPT_TYPE_INT.
	AVOptTypeInt AVOptionType = C.AV_OPT_TYPE_INT
	// AVOptTypeInt64 wraps AV_OPT_TYPE_INT64.
	AVOptTypeInt64 AVOptionType = C.AV_OPT_TYPE_INT64
	// AVOptTypeDouble wraps AV_OPT_TYPE_DOUBLE.
	AVOptTypeDouble AVOptionType = C.AV_OPT_TYPE_DOUBLE
	// AVOptTypeFloat wraps AV_OPT_TYPE_FLOAT.
	AVOptTypeFloat AVOptionType = C.AV_OPT_TYPE_FLOAT
	// AVOptTypeString wraps AV_OPT_TYPE_STRING.
	AVOptTypeString AVOptionType = C.AV_OPT_TYPE_STRING
	// AVOptTypeRational wraps AV_OPT_TYPE_RATIONAL.
	AVOptTypeRational AVOptionType = C.AV_OPT_TYPE_RATIONAL
	// AVOptTypeBinary wraps AV_OPT_TYPE_BINARY.
	//
	//	offset must point to a pointer immediately followed by an int for the length
	AVOptTypeBinary AVOptionType = C.AV_OPT_TYPE_BINARY
	// AVOptTypeDict wraps AV_OPT_TYPE_DICT.
	AVOptTypeDict AVOptionType = C.AV_OPT_TYPE_DICT
	// AVOptTypeUint64 wraps AV_OPT_TYPE_UINT64.
	AVOptTypeUint64 AVOptionType = C.AV_OPT_TYPE_UINT64
	// AVOptTypeConst wraps AV_OPT_TYPE_CONST.
	AVOptTypeConst AVOptionType = C.AV_OPT_TYPE_CONST
	// AVOptTypeImageSize wraps AV_OPT_TYPE_IMAGE_SIZE.
	//
	//	offset must point to two consecutive integers
	AVOptTypeImageSize AVOptionType = C.AV_OPT_TYPE_IMAGE_SIZE
	// AVOptTypePixelFmt wraps AV_OPT_TYPE_PIXEL_FMT.
	AVOptTypePixelFmt AVOptionType = C.AV_OPT_TYPE_PIXEL_FMT
	// AVOptTypeSampleFmt wraps AV_OPT_TYPE_SAMPLE_FMT.
	AVOptTypeSampleFmt AVOptionType = C.AV_OPT_TYPE_SAMPLE_FMT
	// AVOptTypeVideoRate wraps AV_OPT_TYPE_VIDEO_RATE.
	//
	//	offset must point to AVRational
	AVOptTypeVideoRate AVOptionType = C.AV_OPT_TYPE_VIDEO_RATE
	// AVOptTypeDuration wraps AV_OPT_TYPE_DURATION.
	AVOptTypeDuration AVOptionType = C.AV_OPT_TYPE_DURATION
	// AVOptTypeColor wraps AV_OPT_TYPE_COLOR.
	AVOptTypeColor AVOptionType = C.AV_OPT_TYPE_COLOR
	// AVOptTypeChannelLayout wraps AV_OPT_TYPE_CHANNEL_LAYOUT.
	AVOptTypeChannelLayout AVOptionType = C.AV_OPT_TYPE_CHANNEL_LAYOUT
	// AVOptTypeBool wraps AV_OPT_TYPE_BOOL.
	AVOptTypeBool AVOptionType = C.AV_OPT_TYPE_BOOL
	// AVOptTypeChlayout wraps AV_OPT_TYPE_CHLAYOUT.
	AVOptTypeChlayout AVOptionType = C.AV_OPT_TYPE_CHLAYOUT
)

// --- Enum AVPixelFormat ---

// AVPixelFormat wraps AVPixelFormat.
/*
  Pixel format.

  @note
  AV_PIX_FMT_RGB32 is handled in an endian-specific manner. An RGBA
  color is put together as:
   (A << 24) | (R << 16) | (G << 8) | B
  This is stored as BGRA on little-endian CPU architectures and ARGB on
  big-endian CPUs.

  @note
  If the resolution is not a multiple of the chroma subsampling factor
  then the chroma plane resolution must be rounded up.

  @par
  When the pixel format is palettized RGB32 (AV_PIX_FMT_PAL8), the palettized
  image data is stored in AVFrame.data[0]. The palette is transported in
  AVFrame.data[1], is 1024 bytes long (256 4-byte entries) and is
  formatted the same as in AV_PIX_FMT_RGB32 described above (i.e., it is
  also endian-specific). Note also that the individual RGB32 palette
  components stored in AVFrame.data[1] should be in the range 0..255.
  This is important as many custom PAL8 video codecs that were designed
  to run on the IBM VGA graphics adapter use 6-bit palette components.

  @par
  For all the 8 bits per pixel formats, an RGB32 palette is in data[1] like
  for pal8. This palette is filled in automatically by the function
  allocating the picture.
*/
type AVPixelFormat C.enum_AVPixelFormat

const SizeOfAVPixelFormat = C.sizeof_enum_AVPixelFormat

func ToAVPixelFormatArray(ptr unsafe.Pointer) *Array[AVPixelFormat] {
	if ptr == nil {
		return nil
	}

	return &Array[AVPixelFormat]{
		elemSize: SizeOfAVPixelFormat,
		loadPtr: func(pointer unsafe.Pointer) AVPixelFormat {
			ptr := (*AVPixelFormat)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVPixelFormat) {
			ptr := (*AVPixelFormat)(pointer)
			*ptr = value
		},
	}
}

func AllocAVPixelFormatArray(size uint64) *Array[AVPixelFormat] {
	return ToAVPixelFormatArray(AVCalloc(size, SizeOfAVPixelFormat))
}

const (
	// AVPixFmtNone wraps AV_PIX_FMT_NONE.
	AVPixFmtNone AVPixelFormat = C.AV_PIX_FMT_NONE
	// AVPixFmtYuv420P wraps AV_PIX_FMT_YUV420P.
	//
	//	planar YUV 4:2:0, 12bpp, (1 Cr & Cb sample per 2x2 Y samples)
	AVPixFmtYuv420P AVPixelFormat = C.AV_PIX_FMT_YUV420P
	// AVPixFmtYuyv422 wraps AV_PIX_FMT_YUYV422.
	//
	//	packed YUV 4:2:2, 16bpp, Y0 Cb Y1 Cr
	AVPixFmtYuyv422 AVPixelFormat = C.AV_PIX_FMT_YUYV422
	// AVPixFmtRgb24 wraps AV_PIX_FMT_RGB24.
	//
	//	packed RGB 8:8:8, 24bpp, RGBRGB...
	AVPixFmtRgb24 AVPixelFormat = C.AV_PIX_FMT_RGB24
	// AVPixFmtBgr24 wraps AV_PIX_FMT_BGR24.
	//
	//	packed RGB 8:8:8, 24bpp, BGRBGR...
	AVPixFmtBgr24 AVPixelFormat = C.AV_PIX_FMT_BGR24
	// AVPixFmtYuv422P wraps AV_PIX_FMT_YUV422P.
	//
	//	planar YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	AVPixFmtYuv422P AVPixelFormat = C.AV_PIX_FMT_YUV422P
	// AVPixFmtYuv444P wraps AV_PIX_FMT_YUV444P.
	//
	//	planar YUV 4:4:4, 24bpp, (1 Cr & Cb sample per 1x1 Y samples)
	AVPixFmtYuv444P AVPixelFormat = C.AV_PIX_FMT_YUV444P
	// AVPixFmtYuv410P wraps AV_PIX_FMT_YUV410P.
	//
	//	planar YUV 4:1:0,  9bpp, (1 Cr & Cb sample per 4x4 Y samples)
	AVPixFmtYuv410P AVPixelFormat = C.AV_PIX_FMT_YUV410P
	// AVPixFmtYuv411P wraps AV_PIX_FMT_YUV411P.
	//
	//	planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples)
	AVPixFmtYuv411P AVPixelFormat = C.AV_PIX_FMT_YUV411P
	// AVPixFmtGray8 wraps AV_PIX_FMT_GRAY8.
	//
	//	Y        ,  8bpp
	AVPixFmtGray8 AVPixelFormat = C.AV_PIX_FMT_GRAY8
	// AVPixFmtMonowhite wraps AV_PIX_FMT_MONOWHITE.
	//
	//	Y        ,  1bpp, 0 is white, 1 is black, in each byte pixels are ordered from the msb to the lsb
	AVPixFmtMonowhite AVPixelFormat = C.AV_PIX_FMT_MONOWHITE
	// AVPixFmtMonoblack wraps AV_PIX_FMT_MONOBLACK.
	//
	//	Y        ,  1bpp, 0 is black, 1 is white, in each byte pixels are ordered from the msb to the lsb
	AVPixFmtMonoblack AVPixelFormat = C.AV_PIX_FMT_MONOBLACK
	// AVPixFmtPal8 wraps AV_PIX_FMT_PAL8.
	//
	//	8 bits with AV_PIX_FMT_RGB32 palette
	AVPixFmtPal8 AVPixelFormat = C.AV_PIX_FMT_PAL8
	// AVPixFmtYuvj420P wraps AV_PIX_FMT_YUVJ420P.
	//
	//	planar YUV 4:2:0, 12bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV420P and setting color_range
	AVPixFmtYuvj420P AVPixelFormat = C.AV_PIX_FMT_YUVJ420P
	// AVPixFmtYuvj422P wraps AV_PIX_FMT_YUVJ422P.
	//
	//	planar YUV 4:2:2, 16bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV422P and setting color_range
	AVPixFmtYuvj422P AVPixelFormat = C.AV_PIX_FMT_YUVJ422P
	// AVPixFmtYuvj444P wraps AV_PIX_FMT_YUVJ444P.
	//
	//	planar YUV 4:4:4, 24bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV444P and setting color_range
	AVPixFmtYuvj444P AVPixelFormat = C.AV_PIX_FMT_YUVJ444P
	// AVPixFmtUyvy422 wraps AV_PIX_FMT_UYVY422.
	//
	//	packed YUV 4:2:2, 16bpp, Cb Y0 Cr Y1
	AVPixFmtUyvy422 AVPixelFormat = C.AV_PIX_FMT_UYVY422
	// AVPixFmtUyyvyy411 wraps AV_PIX_FMT_UYYVYY411.
	//
	//	packed YUV 4:1:1, 12bpp, Cb Y0 Y1 Cr Y2 Y3
	AVPixFmtUyyvyy411 AVPixelFormat = C.AV_PIX_FMT_UYYVYY411
	// AVPixFmtBgr8 wraps AV_PIX_FMT_BGR8.
	//
	//	packed RGB 3:3:2,  8bpp, (msb)2B 3G 3R(lsb)
	AVPixFmtBgr8 AVPixelFormat = C.AV_PIX_FMT_BGR8
	// AVPixFmtBgr4 wraps AV_PIX_FMT_BGR4.
	//
	//	packed RGB 1:2:1 bitstream,  4bpp, (msb)1B 2G 1R(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	AVPixFmtBgr4 AVPixelFormat = C.AV_PIX_FMT_BGR4
	// AVPixFmtBgr4Byte wraps AV_PIX_FMT_BGR4_BYTE.
	//
	//	packed RGB 1:2:1,  8bpp, (msb)1B 2G 1R(lsb)
	AVPixFmtBgr4Byte AVPixelFormat = C.AV_PIX_FMT_BGR4_BYTE
	// AVPixFmtRgb8 wraps AV_PIX_FMT_RGB8.
	//
	//	packed RGB 3:3:2,  8bpp, (msb)2R 3G 3B(lsb)
	AVPixFmtRgb8 AVPixelFormat = C.AV_PIX_FMT_RGB8
	// AVPixFmtRgb4 wraps AV_PIX_FMT_RGB4.
	//
	//	packed RGB 1:2:1 bitstream,  4bpp, (msb)1R 2G 1B(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	AVPixFmtRgb4 AVPixelFormat = C.AV_PIX_FMT_RGB4
	// AVPixFmtRgb4Byte wraps AV_PIX_FMT_RGB4_BYTE.
	//
	//	packed RGB 1:2:1,  8bpp, (msb)1R 2G 1B(lsb)
	AVPixFmtRgb4Byte AVPixelFormat = C.AV_PIX_FMT_RGB4_BYTE
	// AVPixFmtNv12 wraps AV_PIX_FMT_NV12.
	//
	//	planar YUV 4:2:0, 12bpp, 1 plane for Y and 1 plane for the UV components, which are interleaved (first byte U and the following byte V)
	AVPixFmtNv12 AVPixelFormat = C.AV_PIX_FMT_NV12
	// AVPixFmtNv21 wraps AV_PIX_FMT_NV21.
	//
	//	as above, but U and V bytes are swapped
	AVPixFmtNv21 AVPixelFormat = C.AV_PIX_FMT_NV21
	// AVPixFmtArgb wraps AV_PIX_FMT_ARGB.
	//
	//	packed ARGB 8:8:8:8, 32bpp, ARGBARGB...
	AVPixFmtArgb AVPixelFormat = C.AV_PIX_FMT_ARGB
	// AVPixFmtRgba wraps AV_PIX_FMT_RGBA.
	//
	//	packed RGBA 8:8:8:8, 32bpp, RGBARGBA...
	AVPixFmtRgba AVPixelFormat = C.AV_PIX_FMT_RGBA
	// AVPixFmtAbgr wraps AV_PIX_FMT_ABGR.
	//
	//	packed ABGR 8:8:8:8, 32bpp, ABGRABGR...
	AVPixFmtAbgr AVPixelFormat = C.AV_PIX_FMT_ABGR
	// AVPixFmtBgra wraps AV_PIX_FMT_BGRA.
	//
	//	packed BGRA 8:8:8:8, 32bpp, BGRABGRA...
	AVPixFmtBgra AVPixelFormat = C.AV_PIX_FMT_BGRA
	// AVPixFmtGray16Be wraps AV_PIX_FMT_GRAY16BE.
	//
	//	Y        , 16bpp, big-endian
	AVPixFmtGray16Be AVPixelFormat = C.AV_PIX_FMT_GRAY16BE
	// AVPixFmtGray16Le wraps AV_PIX_FMT_GRAY16LE.
	//
	//	Y        , 16bpp, little-endian
	AVPixFmtGray16Le AVPixelFormat = C.AV_PIX_FMT_GRAY16LE
	// AVPixFmtYuv440P wraps AV_PIX_FMT_YUV440P.
	//
	//	planar YUV 4:4:0 (1 Cr & Cb sample per 1x2 Y samples)
	AVPixFmtYuv440P AVPixelFormat = C.AV_PIX_FMT_YUV440P
	// AVPixFmtYuvj440P wraps AV_PIX_FMT_YUVJ440P.
	//
	//	planar YUV 4:4:0 full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV440P and setting color_range
	AVPixFmtYuvj440P AVPixelFormat = C.AV_PIX_FMT_YUVJ440P
	// AVPixFmtYuva420P wraps AV_PIX_FMT_YUVA420P.
	//
	//	planar YUV 4:2:0, 20bpp, (1 Cr & Cb sample per 2x2 Y & A samples)
	AVPixFmtYuva420P AVPixelFormat = C.AV_PIX_FMT_YUVA420P
	// AVPixFmtRgb48Be wraps AV_PIX_FMT_RGB48BE.
	//
	//	packed RGB 16:16:16, 48bpp, 16R, 16G, 16B, the 2-byte value for each R/G/B component is stored as big-endian
	AVPixFmtRgb48Be AVPixelFormat = C.AV_PIX_FMT_RGB48BE
	// AVPixFmtRgb48Le wraps AV_PIX_FMT_RGB48LE.
	//
	//	packed RGB 16:16:16, 48bpp, 16R, 16G, 16B, the 2-byte value for each R/G/B component is stored as little-endian
	AVPixFmtRgb48Le AVPixelFormat = C.AV_PIX_FMT_RGB48LE
	// AVPixFmtRgb565Be wraps AV_PIX_FMT_RGB565BE.
	//
	//	packed RGB 5:6:5, 16bpp, (msb)   5R 6G 5B(lsb), big-endian
	AVPixFmtRgb565Be AVPixelFormat = C.AV_PIX_FMT_RGB565BE
	// AVPixFmtRgb565Le wraps AV_PIX_FMT_RGB565LE.
	//
	//	packed RGB 5:6:5, 16bpp, (msb)   5R 6G 5B(lsb), little-endian
	AVPixFmtRgb565Le AVPixelFormat = C.AV_PIX_FMT_RGB565LE
	// AVPixFmtRgb555Be wraps AV_PIX_FMT_RGB555BE.
	//
	//	packed RGB 5:5:5, 16bpp, (msb)1X 5R 5G 5B(lsb), big-endian   , X=unused/undefined
	AVPixFmtRgb555Be AVPixelFormat = C.AV_PIX_FMT_RGB555BE
	// AVPixFmtRgb555Le wraps AV_PIX_FMT_RGB555LE.
	//
	//	packed RGB 5:5:5, 16bpp, (msb)1X 5R 5G 5B(lsb), little-endian, X=unused/undefined
	AVPixFmtRgb555Le AVPixelFormat = C.AV_PIX_FMT_RGB555LE
	// AVPixFmtBgr565Be wraps AV_PIX_FMT_BGR565BE.
	//
	//	packed BGR 5:6:5, 16bpp, (msb)   5B 6G 5R(lsb), big-endian
	AVPixFmtBgr565Be AVPixelFormat = C.AV_PIX_FMT_BGR565BE
	// AVPixFmtBgr565Le wraps AV_PIX_FMT_BGR565LE.
	//
	//	packed BGR 5:6:5, 16bpp, (msb)   5B 6G 5R(lsb), little-endian
	AVPixFmtBgr565Le AVPixelFormat = C.AV_PIX_FMT_BGR565LE
	// AVPixFmtBgr555Be wraps AV_PIX_FMT_BGR555BE.
	//
	//	packed BGR 5:5:5, 16bpp, (msb)1X 5B 5G 5R(lsb), big-endian   , X=unused/undefined
	AVPixFmtBgr555Be AVPixelFormat = C.AV_PIX_FMT_BGR555BE
	// AVPixFmtBgr555Le wraps AV_PIX_FMT_BGR555LE.
	//
	//	packed BGR 5:5:5, 16bpp, (msb)1X 5B 5G 5R(lsb), little-endian, X=unused/undefined
	AVPixFmtBgr555Le AVPixelFormat = C.AV_PIX_FMT_BGR555LE
	// AVPixFmtVaapi wraps AV_PIX_FMT_VAAPI.
	/*
	   Hardware acceleration through VA-API, data[3] contains a
	   VASurfaceID.
	*/
	AVPixFmtVaapi AVPixelFormat = C.AV_PIX_FMT_VAAPI
	// AVPixFmtYuv420P16Le wraps AV_PIX_FMT_YUV420P16LE.
	//
	//	planar YUV 4:2:0, 24bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AVPixFmtYuv420P16Le AVPixelFormat = C.AV_PIX_FMT_YUV420P16LE
	// AVPixFmtYuv420P16Be wraps AV_PIX_FMT_YUV420P16BE.
	//
	//	planar YUV 4:2:0, 24bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AVPixFmtYuv420P16Be AVPixelFormat = C.AV_PIX_FMT_YUV420P16BE
	// AVPixFmtYuv422P16Le wraps AV_PIX_FMT_YUV422P16LE.
	//
	//	planar YUV 4:2:2, 32bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AVPixFmtYuv422P16Le AVPixelFormat = C.AV_PIX_FMT_YUV422P16LE
	// AVPixFmtYuv422P16Be wraps AV_PIX_FMT_YUV422P16BE.
	//
	//	planar YUV 4:2:2, 32bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AVPixFmtYuv422P16Be AVPixelFormat = C.AV_PIX_FMT_YUV422P16BE
	// AVPixFmtYuv444P16Le wraps AV_PIX_FMT_YUV444P16LE.
	//
	//	planar YUV 4:4:4, 48bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AVPixFmtYuv444P16Le AVPixelFormat = C.AV_PIX_FMT_YUV444P16LE
	// AVPixFmtYuv444P16Be wraps AV_PIX_FMT_YUV444P16BE.
	//
	//	planar YUV 4:4:4, 48bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AVPixFmtYuv444P16Be AVPixelFormat = C.AV_PIX_FMT_YUV444P16BE
	// AVPixFmtDxva2Vld wraps AV_PIX_FMT_DXVA2_VLD.
	//
	//	HW decoding through DXVA2, Picture.data[3] contains a LPDIRECT3DSURFACE9 pointer
	AVPixFmtDxva2Vld AVPixelFormat = C.AV_PIX_FMT_DXVA2_VLD
	// AVPixFmtRgb444Le wraps AV_PIX_FMT_RGB444LE.
	//
	//	packed RGB 4:4:4, 16bpp, (msb)4X 4R 4G 4B(lsb), little-endian, X=unused/undefined
	AVPixFmtRgb444Le AVPixelFormat = C.AV_PIX_FMT_RGB444LE
	// AVPixFmtRgb444Be wraps AV_PIX_FMT_RGB444BE.
	//
	//	packed RGB 4:4:4, 16bpp, (msb)4X 4R 4G 4B(lsb), big-endian,    X=unused/undefined
	AVPixFmtRgb444Be AVPixelFormat = C.AV_PIX_FMT_RGB444BE
	// AVPixFmtBgr444Le wraps AV_PIX_FMT_BGR444LE.
	//
	//	packed BGR 4:4:4, 16bpp, (msb)4X 4B 4G 4R(lsb), little-endian, X=unused/undefined
	AVPixFmtBgr444Le AVPixelFormat = C.AV_PIX_FMT_BGR444LE
	// AVPixFmtBgr444Be wraps AV_PIX_FMT_BGR444BE.
	//
	//	packed BGR 4:4:4, 16bpp, (msb)4X 4B 4G 4R(lsb), big-endian,    X=unused/undefined
	AVPixFmtBgr444Be AVPixelFormat = C.AV_PIX_FMT_BGR444BE
	// AVPixFmtYa8 wraps AV_PIX_FMT_YA8.
	//
	//	8 bits gray, 8 bits alpha
	AVPixFmtYa8 AVPixelFormat = C.AV_PIX_FMT_YA8
	// AVPixFmtY400A wraps AV_PIX_FMT_Y400A.
	//
	//	alias for AV_PIX_FMT_YA8
	AVPixFmtY400A AVPixelFormat = C.AV_PIX_FMT_Y400A
	// AVPixFmtGray8A wraps AV_PIX_FMT_GRAY8A.
	//
	//	alias for AV_PIX_FMT_YA8
	AVPixFmtGray8A AVPixelFormat = C.AV_PIX_FMT_GRAY8A
	// AVPixFmtBgr48Be wraps AV_PIX_FMT_BGR48BE.
	//
	//	packed RGB 16:16:16, 48bpp, 16B, 16G, 16R, the 2-byte value for each R/G/B component is stored as big-endian
	AVPixFmtBgr48Be AVPixelFormat = C.AV_PIX_FMT_BGR48BE
	// AVPixFmtBgr48Le wraps AV_PIX_FMT_BGR48LE.
	//
	//	packed RGB 16:16:16, 48bpp, 16B, 16G, 16R, the 2-byte value for each R/G/B component is stored as little-endian
	AVPixFmtBgr48Le AVPixelFormat = C.AV_PIX_FMT_BGR48LE
	// AVPixFmtYuv420P9Be wraps AV_PIX_FMT_YUV420P9BE.
	//
	//	planar YUV 4:2:0, 13.5bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AVPixFmtYuv420P9Be AVPixelFormat = C.AV_PIX_FMT_YUV420P9BE
	// AVPixFmtYuv420P9Le wraps AV_PIX_FMT_YUV420P9LE.
	//
	//	planar YUV 4:2:0, 13.5bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AVPixFmtYuv420P9Le AVPixelFormat = C.AV_PIX_FMT_YUV420P9LE
	// AVPixFmtYuv420P10Be wraps AV_PIX_FMT_YUV420P10BE.
	//
	//	planar YUV 4:2:0, 15bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AVPixFmtYuv420P10Be AVPixelFormat = C.AV_PIX_FMT_YUV420P10BE
	// AVPixFmtYuv420P10Le wraps AV_PIX_FMT_YUV420P10LE.
	//
	//	planar YUV 4:2:0, 15bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AVPixFmtYuv420P10Le AVPixelFormat = C.AV_PIX_FMT_YUV420P10LE
	// AVPixFmtYuv422P10Be wraps AV_PIX_FMT_YUV422P10BE.
	//
	//	planar YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AVPixFmtYuv422P10Be AVPixelFormat = C.AV_PIX_FMT_YUV422P10BE
	// AVPixFmtYuv422P10Le wraps AV_PIX_FMT_YUV422P10LE.
	//
	//	planar YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AVPixFmtYuv422P10Le AVPixelFormat = C.AV_PIX_FMT_YUV422P10LE
	// AVPixFmtYuv444P9Be wraps AV_PIX_FMT_YUV444P9BE.
	//
	//	planar YUV 4:4:4, 27bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AVPixFmtYuv444P9Be AVPixelFormat = C.AV_PIX_FMT_YUV444P9BE
	// AVPixFmtYuv444P9Le wraps AV_PIX_FMT_YUV444P9LE.
	//
	//	planar YUV 4:4:4, 27bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AVPixFmtYuv444P9Le AVPixelFormat = C.AV_PIX_FMT_YUV444P9LE
	// AVPixFmtYuv444P10Be wraps AV_PIX_FMT_YUV444P10BE.
	//
	//	planar YUV 4:4:4, 30bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AVPixFmtYuv444P10Be AVPixelFormat = C.AV_PIX_FMT_YUV444P10BE
	// AVPixFmtYuv444P10Le wraps AV_PIX_FMT_YUV444P10LE.
	//
	//	planar YUV 4:4:4, 30bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AVPixFmtYuv444P10Le AVPixelFormat = C.AV_PIX_FMT_YUV444P10LE
	// AVPixFmtYuv422P9Be wraps AV_PIX_FMT_YUV422P9BE.
	//
	//	planar YUV 4:2:2, 18bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AVPixFmtYuv422P9Be AVPixelFormat = C.AV_PIX_FMT_YUV422P9BE
	// AVPixFmtYuv422P9Le wraps AV_PIX_FMT_YUV422P9LE.
	//
	//	planar YUV 4:2:2, 18bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AVPixFmtYuv422P9Le AVPixelFormat = C.AV_PIX_FMT_YUV422P9LE
	// AVPixFmtGbrp wraps AV_PIX_FMT_GBRP.
	//
	//	planar GBR 4:4:4 24bpp
	AVPixFmtGbrp AVPixelFormat = C.AV_PIX_FMT_GBRP
	// AVPixFmtGbr24P wraps AV_PIX_FMT_GBR24P.
	//
	//	// alias for #AV_PIX_FMT_GBRP
	AVPixFmtGbr24P AVPixelFormat = C.AV_PIX_FMT_GBR24P
	// AVPixFmtGbrp9Be wraps AV_PIX_FMT_GBRP9BE.
	//
	//	planar GBR 4:4:4 27bpp, big-endian
	AVPixFmtGbrp9Be AVPixelFormat = C.AV_PIX_FMT_GBRP9BE
	// AVPixFmtGbrp9Le wraps AV_PIX_FMT_GBRP9LE.
	//
	//	planar GBR 4:4:4 27bpp, little-endian
	AVPixFmtGbrp9Le AVPixelFormat = C.AV_PIX_FMT_GBRP9LE
	// AVPixFmtGbrp10Be wraps AV_PIX_FMT_GBRP10BE.
	//
	//	planar GBR 4:4:4 30bpp, big-endian
	AVPixFmtGbrp10Be AVPixelFormat = C.AV_PIX_FMT_GBRP10BE
	// AVPixFmtGbrp10Le wraps AV_PIX_FMT_GBRP10LE.
	//
	//	planar GBR 4:4:4 30bpp, little-endian
	AVPixFmtGbrp10Le AVPixelFormat = C.AV_PIX_FMT_GBRP10LE
	// AVPixFmtGbrp16Be wraps AV_PIX_FMT_GBRP16BE.
	//
	//	planar GBR 4:4:4 48bpp, big-endian
	AVPixFmtGbrp16Be AVPixelFormat = C.AV_PIX_FMT_GBRP16BE
	// AVPixFmtGbrp16Le wraps AV_PIX_FMT_GBRP16LE.
	//
	//	planar GBR 4:4:4 48bpp, little-endian
	AVPixFmtGbrp16Le AVPixelFormat = C.AV_PIX_FMT_GBRP16LE
	// AVPixFmtYuva422P wraps AV_PIX_FMT_YUVA422P.
	//
	//	planar YUV 4:2:2 24bpp, (1 Cr & Cb sample per 2x1 Y & A samples)
	AVPixFmtYuva422P AVPixelFormat = C.AV_PIX_FMT_YUVA422P
	// AVPixFmtYuva444P wraps AV_PIX_FMT_YUVA444P.
	//
	//	planar YUV 4:4:4 32bpp, (1 Cr & Cb sample per 1x1 Y & A samples)
	AVPixFmtYuva444P AVPixelFormat = C.AV_PIX_FMT_YUVA444P
	// AVPixFmtYuva420P9Be wraps AV_PIX_FMT_YUVA420P9BE.
	//
	//	planar YUV 4:2:0 22.5bpp, (1 Cr & Cb sample per 2x2 Y & A samples), big-endian
	AVPixFmtYuva420P9Be AVPixelFormat = C.AV_PIX_FMT_YUVA420P9BE
	// AVPixFmtYuva420P9Le wraps AV_PIX_FMT_YUVA420P9LE.
	//
	//	planar YUV 4:2:0 22.5bpp, (1 Cr & Cb sample per 2x2 Y & A samples), little-endian
	AVPixFmtYuva420P9Le AVPixelFormat = C.AV_PIX_FMT_YUVA420P9LE
	// AVPixFmtYuva422P9Be wraps AV_PIX_FMT_YUVA422P9BE.
	//
	//	planar YUV 4:2:2 27bpp, (1 Cr & Cb sample per 2x1 Y & A samples), big-endian
	AVPixFmtYuva422P9Be AVPixelFormat = C.AV_PIX_FMT_YUVA422P9BE
	// AVPixFmtYuva422P9Le wraps AV_PIX_FMT_YUVA422P9LE.
	//
	//	planar YUV 4:2:2 27bpp, (1 Cr & Cb sample per 2x1 Y & A samples), little-endian
	AVPixFmtYuva422P9Le AVPixelFormat = C.AV_PIX_FMT_YUVA422P9LE
	// AVPixFmtYuva444P9Be wraps AV_PIX_FMT_YUVA444P9BE.
	//
	//	planar YUV 4:4:4 36bpp, (1 Cr & Cb sample per 1x1 Y & A samples), big-endian
	AVPixFmtYuva444P9Be AVPixelFormat = C.AV_PIX_FMT_YUVA444P9BE
	// AVPixFmtYuva444P9Le wraps AV_PIX_FMT_YUVA444P9LE.
	//
	//	planar YUV 4:4:4 36bpp, (1 Cr & Cb sample per 1x1 Y & A samples), little-endian
	AVPixFmtYuva444P9Le AVPixelFormat = C.AV_PIX_FMT_YUVA444P9LE
	// AVPixFmtYuva420P10Be wraps AV_PIX_FMT_YUVA420P10BE.
	//
	//	planar YUV 4:2:0 25bpp, (1 Cr & Cb sample per 2x2 Y & A samples, big-endian)
	AVPixFmtYuva420P10Be AVPixelFormat = C.AV_PIX_FMT_YUVA420P10BE
	// AVPixFmtYuva420P10Le wraps AV_PIX_FMT_YUVA420P10LE.
	//
	//	planar YUV 4:2:0 25bpp, (1 Cr & Cb sample per 2x2 Y & A samples, little-endian)
	AVPixFmtYuva420P10Le AVPixelFormat = C.AV_PIX_FMT_YUVA420P10LE
	// AVPixFmtYuva422P10Be wraps AV_PIX_FMT_YUVA422P10BE.
	//
	//	planar YUV 4:2:2 30bpp, (1 Cr & Cb sample per 2x1 Y & A samples, big-endian)
	AVPixFmtYuva422P10Be AVPixelFormat = C.AV_PIX_FMT_YUVA422P10BE
	// AVPixFmtYuva422P10Le wraps AV_PIX_FMT_YUVA422P10LE.
	//
	//	planar YUV 4:2:2 30bpp, (1 Cr & Cb sample per 2x1 Y & A samples, little-endian)
	AVPixFmtYuva422P10Le AVPixelFormat = C.AV_PIX_FMT_YUVA422P10LE
	// AVPixFmtYuva444P10Be wraps AV_PIX_FMT_YUVA444P10BE.
	//
	//	planar YUV 4:4:4 40bpp, (1 Cr & Cb sample per 1x1 Y & A samples, big-endian)
	AVPixFmtYuva444P10Be AVPixelFormat = C.AV_PIX_FMT_YUVA444P10BE
	// AVPixFmtYuva444P10Le wraps AV_PIX_FMT_YUVA444P10LE.
	//
	//	planar YUV 4:4:4 40bpp, (1 Cr & Cb sample per 1x1 Y & A samples, little-endian)
	AVPixFmtYuva444P10Le AVPixelFormat = C.AV_PIX_FMT_YUVA444P10LE
	// AVPixFmtYuva420P16Be wraps AV_PIX_FMT_YUVA420P16BE.
	//
	//	planar YUV 4:2:0 40bpp, (1 Cr & Cb sample per 2x2 Y & A samples, big-endian)
	AVPixFmtYuva420P16Be AVPixelFormat = C.AV_PIX_FMT_YUVA420P16BE
	// AVPixFmtYuva420P16Le wraps AV_PIX_FMT_YUVA420P16LE.
	//
	//	planar YUV 4:2:0 40bpp, (1 Cr & Cb sample per 2x2 Y & A samples, little-endian)
	AVPixFmtYuva420P16Le AVPixelFormat = C.AV_PIX_FMT_YUVA420P16LE
	// AVPixFmtYuva422P16Be wraps AV_PIX_FMT_YUVA422P16BE.
	//
	//	planar YUV 4:2:2 48bpp, (1 Cr & Cb sample per 2x1 Y & A samples, big-endian)
	AVPixFmtYuva422P16Be AVPixelFormat = C.AV_PIX_FMT_YUVA422P16BE
	// AVPixFmtYuva422P16Le wraps AV_PIX_FMT_YUVA422P16LE.
	//
	//	planar YUV 4:2:2 48bpp, (1 Cr & Cb sample per 2x1 Y & A samples, little-endian)
	AVPixFmtYuva422P16Le AVPixelFormat = C.AV_PIX_FMT_YUVA422P16LE
	// AVPixFmtYuva444P16Be wraps AV_PIX_FMT_YUVA444P16BE.
	//
	//	planar YUV 4:4:4 64bpp, (1 Cr & Cb sample per 1x1 Y & A samples, big-endian)
	AVPixFmtYuva444P16Be AVPixelFormat = C.AV_PIX_FMT_YUVA444P16BE
	// AVPixFmtYuva444P16Le wraps AV_PIX_FMT_YUVA444P16LE.
	//
	//	planar YUV 4:4:4 64bpp, (1 Cr & Cb sample per 1x1 Y & A samples, little-endian)
	AVPixFmtYuva444P16Le AVPixelFormat = C.AV_PIX_FMT_YUVA444P16LE
	// AVPixFmtVdpau wraps AV_PIX_FMT_VDPAU.
	//
	//	HW acceleration through VDPAU, Picture.data[3] contains a VdpVideoSurface
	AVPixFmtVdpau AVPixelFormat = C.AV_PIX_FMT_VDPAU
	// AVPixFmtXyz12Le wraps AV_PIX_FMT_XYZ12LE.
	//
	//	packed XYZ 4:4:4, 36 bpp, (msb) 12X, 12Y, 12Z (lsb), the 2-byte value for each X/Y/Z is stored as little-endian, the 4 lower bits are set to 0
	AVPixFmtXyz12Le AVPixelFormat = C.AV_PIX_FMT_XYZ12LE
	// AVPixFmtXyz12Be wraps AV_PIX_FMT_XYZ12BE.
	//
	//	packed XYZ 4:4:4, 36 bpp, (msb) 12X, 12Y, 12Z (lsb), the 2-byte value for each X/Y/Z is stored as big-endian, the 4 lower bits are set to 0
	AVPixFmtXyz12Be AVPixelFormat = C.AV_PIX_FMT_XYZ12BE
	// AVPixFmtNv16 wraps AV_PIX_FMT_NV16.
	//
	//	interleaved chroma YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	AVPixFmtNv16 AVPixelFormat = C.AV_PIX_FMT_NV16
	// AVPixFmtNv20Le wraps AV_PIX_FMT_NV20LE.
	//
	//	interleaved chroma YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AVPixFmtNv20Le AVPixelFormat = C.AV_PIX_FMT_NV20LE
	// AVPixFmtNv20Be wraps AV_PIX_FMT_NV20BE.
	//
	//	interleaved chroma YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AVPixFmtNv20Be AVPixelFormat = C.AV_PIX_FMT_NV20BE
	// AVPixFmtRgba64Be wraps AV_PIX_FMT_RGBA64BE.
	//
	//	packed RGBA 16:16:16:16, 64bpp, 16R, 16G, 16B, 16A, the 2-byte value for each R/G/B/A component is stored as big-endian
	AVPixFmtRgba64Be AVPixelFormat = C.AV_PIX_FMT_RGBA64BE
	// AVPixFmtRgba64Le wraps AV_PIX_FMT_RGBA64LE.
	//
	//	packed RGBA 16:16:16:16, 64bpp, 16R, 16G, 16B, 16A, the 2-byte value for each R/G/B/A component is stored as little-endian
	AVPixFmtRgba64Le AVPixelFormat = C.AV_PIX_FMT_RGBA64LE
	// AVPixFmtBgra64Be wraps AV_PIX_FMT_BGRA64BE.
	//
	//	packed RGBA 16:16:16:16, 64bpp, 16B, 16G, 16R, 16A, the 2-byte value for each R/G/B/A component is stored as big-endian
	AVPixFmtBgra64Be AVPixelFormat = C.AV_PIX_FMT_BGRA64BE
	// AVPixFmtBgra64Le wraps AV_PIX_FMT_BGRA64LE.
	//
	//	packed RGBA 16:16:16:16, 64bpp, 16B, 16G, 16R, 16A, the 2-byte value for each R/G/B/A component is stored as little-endian
	AVPixFmtBgra64Le AVPixelFormat = C.AV_PIX_FMT_BGRA64LE
	// AVPixFmtYvyu422 wraps AV_PIX_FMT_YVYU422.
	//
	//	packed YUV 4:2:2, 16bpp, Y0 Cr Y1 Cb
	AVPixFmtYvyu422 AVPixelFormat = C.AV_PIX_FMT_YVYU422
	// AVPixFmtYa16Be wraps AV_PIX_FMT_YA16BE.
	//
	//	16 bits gray, 16 bits alpha (big-endian)
	AVPixFmtYa16Be AVPixelFormat = C.AV_PIX_FMT_YA16BE
	// AVPixFmtYa16Le wraps AV_PIX_FMT_YA16LE.
	//
	//	16 bits gray, 16 bits alpha (little-endian)
	AVPixFmtYa16Le AVPixelFormat = C.AV_PIX_FMT_YA16LE
	// AVPixFmtGbrap wraps AV_PIX_FMT_GBRAP.
	//
	//	planar GBRA 4:4:4:4 32bpp
	AVPixFmtGbrap AVPixelFormat = C.AV_PIX_FMT_GBRAP
	// AVPixFmtGbrap16Be wraps AV_PIX_FMT_GBRAP16BE.
	//
	//	planar GBRA 4:4:4:4 64bpp, big-endian
	AVPixFmtGbrap16Be AVPixelFormat = C.AV_PIX_FMT_GBRAP16BE
	// AVPixFmtGbrap16Le wraps AV_PIX_FMT_GBRAP16LE.
	//
	//	planar GBRA 4:4:4:4 64bpp, little-endian
	AVPixFmtGbrap16Le AVPixelFormat = C.AV_PIX_FMT_GBRAP16LE
	// AVPixFmtQsv wraps AV_PIX_FMT_QSV.
	/*
	   HW acceleration through QSV, data[3] contains a pointer to the
	   mfxFrameSurface1 structure.

	   Before FFmpeg 5.0:
	   mfxFrameSurface1.Data.MemId contains a pointer when importing
	   the following frames as QSV frames:

	   VAAPI:
	   mfxFrameSurface1.Data.MemId contains a pointer to VASurfaceID

	   DXVA2:
	   mfxFrameSurface1.Data.MemId contains a pointer to IDirect3DSurface9

	   FFmpeg 5.0 and above:
	   mfxFrameSurface1.Data.MemId contains a pointer to the mfxHDLPair
	   structure when importing the following frames as QSV frames:

	   VAAPI:
	   mfxHDLPair.first contains a VASurfaceID pointer.
	   mfxHDLPair.second is always MFX_INFINITE.

	   DXVA2:
	   mfxHDLPair.first contains IDirect3DSurface9 pointer.
	   mfxHDLPair.second is always MFX_INFINITE.

	   D3D11:
	   mfxHDLPair.first contains a ID3D11Texture2D pointer.
	   mfxHDLPair.second contains the texture array index of the frame if the
	   ID3D11Texture2D is an array texture, or always MFX_INFINITE if it is a
	   normal texture.
	*/
	AVPixFmtQsv AVPixelFormat = C.AV_PIX_FMT_QSV
	// AVPixFmtMmal wraps AV_PIX_FMT_MMAL.
	/*
	   HW acceleration though MMAL, data[3] contains a pointer to the
	   MMAL_BUFFER_HEADER_T structure.
	*/
	AVPixFmtMmal AVPixelFormat = C.AV_PIX_FMT_MMAL
	// AVPixFmtD3D11VaVld wraps AV_PIX_FMT_D3D11VA_VLD.
	//
	//	HW decoding through Direct3D11 via old API, Picture.data[3] contains a ID3D11VideoDecoderOutputView pointer
	AVPixFmtD3D11VaVld AVPixelFormat = C.AV_PIX_FMT_D3D11VA_VLD
	// AVPixFmtCuda wraps AV_PIX_FMT_CUDA.
	/*
	   HW acceleration through CUDA. data[i] contain CUdeviceptr pointers
	   exactly as for system memory frames.
	*/
	AVPixFmtCuda AVPixelFormat = C.AV_PIX_FMT_CUDA
	// AVPixFmt0Rgb wraps AV_PIX_FMT_0RGB.
	//
	//	packed RGB 8:8:8, 32bpp, XRGBXRGB...   X=unused/undefined
	AVPixFmt0Rgb AVPixelFormat = C.AV_PIX_FMT_0RGB
	// AVPixFmtRgb0 wraps AV_PIX_FMT_RGB0.
	//
	//	packed RGB 8:8:8, 32bpp, RGBXRGBX...   X=unused/undefined
	AVPixFmtRgb0 AVPixelFormat = C.AV_PIX_FMT_RGB0
	// AVPixFmt0Bgr wraps AV_PIX_FMT_0BGR.
	//
	//	packed BGR 8:8:8, 32bpp, XBGRXBGR...   X=unused/undefined
	AVPixFmt0Bgr AVPixelFormat = C.AV_PIX_FMT_0BGR
	// AVPixFmtBgr0 wraps AV_PIX_FMT_BGR0.
	//
	//	packed BGR 8:8:8, 32bpp, BGRXBGRX...   X=unused/undefined
	AVPixFmtBgr0 AVPixelFormat = C.AV_PIX_FMT_BGR0
	// AVPixFmtYuv420P12Be wraps AV_PIX_FMT_YUV420P12BE.
	//
	//	planar YUV 4:2:0,18bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AVPixFmtYuv420P12Be AVPixelFormat = C.AV_PIX_FMT_YUV420P12BE
	// AVPixFmtYuv420P12Le wraps AV_PIX_FMT_YUV420P12LE.
	//
	//	planar YUV 4:2:0,18bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AVPixFmtYuv420P12Le AVPixelFormat = C.AV_PIX_FMT_YUV420P12LE
	// AVPixFmtYuv420P14Be wraps AV_PIX_FMT_YUV420P14BE.
	//
	//	planar YUV 4:2:0,21bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	AVPixFmtYuv420P14Be AVPixelFormat = C.AV_PIX_FMT_YUV420P14BE
	// AVPixFmtYuv420P14Le wraps AV_PIX_FMT_YUV420P14LE.
	//
	//	planar YUV 4:2:0,21bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	AVPixFmtYuv420P14Le AVPixelFormat = C.AV_PIX_FMT_YUV420P14LE
	// AVPixFmtYuv422P12Be wraps AV_PIX_FMT_YUV422P12BE.
	//
	//	planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AVPixFmtYuv422P12Be AVPixelFormat = C.AV_PIX_FMT_YUV422P12BE
	// AVPixFmtYuv422P12Le wraps AV_PIX_FMT_YUV422P12LE.
	//
	//	planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AVPixFmtYuv422P12Le AVPixelFormat = C.AV_PIX_FMT_YUV422P12LE
	// AVPixFmtYuv422P14Be wraps AV_PIX_FMT_YUV422P14BE.
	//
	//	planar YUV 4:2:2,28bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	AVPixFmtYuv422P14Be AVPixelFormat = C.AV_PIX_FMT_YUV422P14BE
	// AVPixFmtYuv422P14Le wraps AV_PIX_FMT_YUV422P14LE.
	//
	//	planar YUV 4:2:2,28bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	AVPixFmtYuv422P14Le AVPixelFormat = C.AV_PIX_FMT_YUV422P14LE
	// AVPixFmtYuv444P12Be wraps AV_PIX_FMT_YUV444P12BE.
	//
	//	planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AVPixFmtYuv444P12Be AVPixelFormat = C.AV_PIX_FMT_YUV444P12BE
	// AVPixFmtYuv444P12Le wraps AV_PIX_FMT_YUV444P12LE.
	//
	//	planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AVPixFmtYuv444P12Le AVPixelFormat = C.AV_PIX_FMT_YUV444P12LE
	// AVPixFmtYuv444P14Be wraps AV_PIX_FMT_YUV444P14BE.
	//
	//	planar YUV 4:4:4,42bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	AVPixFmtYuv444P14Be AVPixelFormat = C.AV_PIX_FMT_YUV444P14BE
	// AVPixFmtYuv444P14Le wraps AV_PIX_FMT_YUV444P14LE.
	//
	//	planar YUV 4:4:4,42bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	AVPixFmtYuv444P14Le AVPixelFormat = C.AV_PIX_FMT_YUV444P14LE
	// AVPixFmtGbrp12Be wraps AV_PIX_FMT_GBRP12BE.
	//
	//	planar GBR 4:4:4 36bpp, big-endian
	AVPixFmtGbrp12Be AVPixelFormat = C.AV_PIX_FMT_GBRP12BE
	// AVPixFmtGbrp12Le wraps AV_PIX_FMT_GBRP12LE.
	//
	//	planar GBR 4:4:4 36bpp, little-endian
	AVPixFmtGbrp12Le AVPixelFormat = C.AV_PIX_FMT_GBRP12LE
	// AVPixFmtGbrp14Be wraps AV_PIX_FMT_GBRP14BE.
	//
	//	planar GBR 4:4:4 42bpp, big-endian
	AVPixFmtGbrp14Be AVPixelFormat = C.AV_PIX_FMT_GBRP14BE
	// AVPixFmtGbrp14Le wraps AV_PIX_FMT_GBRP14LE.
	//
	//	planar GBR 4:4:4 42bpp, little-endian
	AVPixFmtGbrp14Le AVPixelFormat = C.AV_PIX_FMT_GBRP14LE
	// AVPixFmtYuvj411P wraps AV_PIX_FMT_YUVJ411P.
	//
	//	planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples) full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV411P and setting color_range
	AVPixFmtYuvj411P AVPixelFormat = C.AV_PIX_FMT_YUVJ411P
	// AVPixFmtBayerBggr8 wraps AV_PIX_FMT_BAYER_BGGR8.
	//
	//	bayer, BGBG..(odd line), GRGR..(even line), 8-bit samples
	AVPixFmtBayerBggr8 AVPixelFormat = C.AV_PIX_FMT_BAYER_BGGR8
	// AVPixFmtBayerRggb8 wraps AV_PIX_FMT_BAYER_RGGB8.
	//
	//	bayer, RGRG..(odd line), GBGB..(even line), 8-bit samples
	AVPixFmtBayerRggb8 AVPixelFormat = C.AV_PIX_FMT_BAYER_RGGB8
	// AVPixFmtBayerGbrg8 wraps AV_PIX_FMT_BAYER_GBRG8.
	//
	//	bayer, GBGB..(odd line), RGRG..(even line), 8-bit samples
	AVPixFmtBayerGbrg8 AVPixelFormat = C.AV_PIX_FMT_BAYER_GBRG8
	// AVPixFmtBayerGrbg8 wraps AV_PIX_FMT_BAYER_GRBG8.
	//
	//	bayer, GRGR..(odd line), BGBG..(even line), 8-bit samples
	AVPixFmtBayerGrbg8 AVPixelFormat = C.AV_PIX_FMT_BAYER_GRBG8
	// AVPixFmtBayerBggr16Le wraps AV_PIX_FMT_BAYER_BGGR16LE.
	//
	//	bayer, BGBG..(odd line), GRGR..(even line), 16-bit samples, little-endian
	AVPixFmtBayerBggr16Le AVPixelFormat = C.AV_PIX_FMT_BAYER_BGGR16LE
	// AVPixFmtBayerBggr16Be wraps AV_PIX_FMT_BAYER_BGGR16BE.
	//
	//	bayer, BGBG..(odd line), GRGR..(even line), 16-bit samples, big-endian
	AVPixFmtBayerBggr16Be AVPixelFormat = C.AV_PIX_FMT_BAYER_BGGR16BE
	// AVPixFmtBayerRggb16Le wraps AV_PIX_FMT_BAYER_RGGB16LE.
	//
	//	bayer, RGRG..(odd line), GBGB..(even line), 16-bit samples, little-endian
	AVPixFmtBayerRggb16Le AVPixelFormat = C.AV_PIX_FMT_BAYER_RGGB16LE
	// AVPixFmtBayerRggb16Be wraps AV_PIX_FMT_BAYER_RGGB16BE.
	//
	//	bayer, RGRG..(odd line), GBGB..(even line), 16-bit samples, big-endian
	AVPixFmtBayerRggb16Be AVPixelFormat = C.AV_PIX_FMT_BAYER_RGGB16BE
	// AVPixFmtBayerGbrg16Le wraps AV_PIX_FMT_BAYER_GBRG16LE.
	//
	//	bayer, GBGB..(odd line), RGRG..(even line), 16-bit samples, little-endian
	AVPixFmtBayerGbrg16Le AVPixelFormat = C.AV_PIX_FMT_BAYER_GBRG16LE
	// AVPixFmtBayerGbrg16Be wraps AV_PIX_FMT_BAYER_GBRG16BE.
	//
	//	bayer, GBGB..(odd line), RGRG..(even line), 16-bit samples, big-endian
	AVPixFmtBayerGbrg16Be AVPixelFormat = C.AV_PIX_FMT_BAYER_GBRG16BE
	// AVPixFmtBayerGrbg16Le wraps AV_PIX_FMT_BAYER_GRBG16LE.
	//
	//	bayer, GRGR..(odd line), BGBG..(even line), 16-bit samples, little-endian
	AVPixFmtBayerGrbg16Le AVPixelFormat = C.AV_PIX_FMT_BAYER_GRBG16LE
	// AVPixFmtBayerGrbg16Be wraps AV_PIX_FMT_BAYER_GRBG16BE.
	//
	//	bayer, GRGR..(odd line), BGBG..(even line), 16-bit samples, big-endian
	AVPixFmtBayerGrbg16Be AVPixelFormat = C.AV_PIX_FMT_BAYER_GRBG16BE
	// AVPixFmtXvmc wraps AV_PIX_FMT_XVMC.
	//
	//	XVideo Motion Acceleration via common packet passing
	AVPixFmtXvmc AVPixelFormat = C.AV_PIX_FMT_XVMC
	// AVPixFmtYuv440P10Le wraps AV_PIX_FMT_YUV440P10LE.
	//
	//	planar YUV 4:4:0,20bpp, (1 Cr & Cb sample per 1x2 Y samples), little-endian
	AVPixFmtYuv440P10Le AVPixelFormat = C.AV_PIX_FMT_YUV440P10LE
	// AVPixFmtYuv440P10Be wraps AV_PIX_FMT_YUV440P10BE.
	//
	//	planar YUV 4:4:0,20bpp, (1 Cr & Cb sample per 1x2 Y samples), big-endian
	AVPixFmtYuv440P10Be AVPixelFormat = C.AV_PIX_FMT_YUV440P10BE
	// AVPixFmtYuv440P12Le wraps AV_PIX_FMT_YUV440P12LE.
	//
	//	planar YUV 4:4:0,24bpp, (1 Cr & Cb sample per 1x2 Y samples), little-endian
	AVPixFmtYuv440P12Le AVPixelFormat = C.AV_PIX_FMT_YUV440P12LE
	// AVPixFmtYuv440P12Be wraps AV_PIX_FMT_YUV440P12BE.
	//
	//	planar YUV 4:4:0,24bpp, (1 Cr & Cb sample per 1x2 Y samples), big-endian
	AVPixFmtYuv440P12Be AVPixelFormat = C.AV_PIX_FMT_YUV440P12BE
	// AVPixFmtAyuv64Le wraps AV_PIX_FMT_AYUV64LE.
	//
	//	packed AYUV 4:4:4,64bpp (1 Cr & Cb sample per 1x1 Y & A samples), little-endian
	AVPixFmtAyuv64Le AVPixelFormat = C.AV_PIX_FMT_AYUV64LE
	// AVPixFmtAyuv64Be wraps AV_PIX_FMT_AYUV64BE.
	//
	//	packed AYUV 4:4:4,64bpp (1 Cr & Cb sample per 1x1 Y & A samples), big-endian
	AVPixFmtAyuv64Be AVPixelFormat = C.AV_PIX_FMT_AYUV64BE
	// AVPixFmtVideotoolbox wraps AV_PIX_FMT_VIDEOTOOLBOX.
	//
	//	hardware decoding through Videotoolbox
	AVPixFmtVideotoolbox AVPixelFormat = C.AV_PIX_FMT_VIDEOTOOLBOX
	// AVPixFmtP010Le wraps AV_PIX_FMT_P010LE.
	//
	//	like NV12, with 10bpp per component, data in the high bits, zeros in the low bits, little-endian
	AVPixFmtP010Le AVPixelFormat = C.AV_PIX_FMT_P010LE
	// AVPixFmtP010Be wraps AV_PIX_FMT_P010BE.
	//
	//	like NV12, with 10bpp per component, data in the high bits, zeros in the low bits, big-endian
	AVPixFmtP010Be AVPixelFormat = C.AV_PIX_FMT_P010BE
	// AVPixFmtGbrap12Be wraps AV_PIX_FMT_GBRAP12BE.
	//
	//	planar GBR 4:4:4:4 48bpp, big-endian
	AVPixFmtGbrap12Be AVPixelFormat = C.AV_PIX_FMT_GBRAP12BE
	// AVPixFmtGbrap12Le wraps AV_PIX_FMT_GBRAP12LE.
	//
	//	planar GBR 4:4:4:4 48bpp, little-endian
	AVPixFmtGbrap12Le AVPixelFormat = C.AV_PIX_FMT_GBRAP12LE
	// AVPixFmtGbrap10Be wraps AV_PIX_FMT_GBRAP10BE.
	//
	//	planar GBR 4:4:4:4 40bpp, big-endian
	AVPixFmtGbrap10Be AVPixelFormat = C.AV_PIX_FMT_GBRAP10BE
	// AVPixFmtGbrap10Le wraps AV_PIX_FMT_GBRAP10LE.
	//
	//	planar GBR 4:4:4:4 40bpp, little-endian
	AVPixFmtGbrap10Le AVPixelFormat = C.AV_PIX_FMT_GBRAP10LE
	// AVPixFmtMediacodec wraps AV_PIX_FMT_MEDIACODEC.
	//
	//	hardware decoding through MediaCodec
	AVPixFmtMediacodec AVPixelFormat = C.AV_PIX_FMT_MEDIACODEC
	// AVPixFmtGray12Be wraps AV_PIX_FMT_GRAY12BE.
	//
	//	Y        , 12bpp, big-endian
	AVPixFmtGray12Be AVPixelFormat = C.AV_PIX_FMT_GRAY12BE
	// AVPixFmtGray12Le wraps AV_PIX_FMT_GRAY12LE.
	//
	//	Y        , 12bpp, little-endian
	AVPixFmtGray12Le AVPixelFormat = C.AV_PIX_FMT_GRAY12LE
	// AVPixFmtGray10Be wraps AV_PIX_FMT_GRAY10BE.
	//
	//	Y        , 10bpp, big-endian
	AVPixFmtGray10Be AVPixelFormat = C.AV_PIX_FMT_GRAY10BE
	// AVPixFmtGray10Le wraps AV_PIX_FMT_GRAY10LE.
	//
	//	Y        , 10bpp, little-endian
	AVPixFmtGray10Le AVPixelFormat = C.AV_PIX_FMT_GRAY10LE
	// AVPixFmtP016Le wraps AV_PIX_FMT_P016LE.
	//
	//	like NV12, with 16bpp per component, little-endian
	AVPixFmtP016Le AVPixelFormat = C.AV_PIX_FMT_P016LE
	// AVPixFmtP016Be wraps AV_PIX_FMT_P016BE.
	//
	//	like NV12, with 16bpp per component, big-endian
	AVPixFmtP016Be AVPixelFormat = C.AV_PIX_FMT_P016BE
	// AVPixFmtD3D11 wraps AV_PIX_FMT_D3D11.
	/*
	   Hardware surfaces for Direct3D11.

	   This is preferred over the legacy AV_PIX_FMT_D3D11VA_VLD. The new D3D11
	   hwaccel API and filtering support AV_PIX_FMT_D3D11 only.

	   data[0] contains a ID3D11Texture2D pointer, and data[1] contains the
	   texture array index of the frame as intptr_t if the ID3D11Texture2D is
	   an array texture (or always 0 if it's a normal texture).
	*/
	AVPixFmtD3D11 AVPixelFormat = C.AV_PIX_FMT_D3D11
	// AVPixFmtGray9Be wraps AV_PIX_FMT_GRAY9BE.
	//
	//	Y        , 9bpp, big-endian
	AVPixFmtGray9Be AVPixelFormat = C.AV_PIX_FMT_GRAY9BE
	// AVPixFmtGray9Le wraps AV_PIX_FMT_GRAY9LE.
	//
	//	Y        , 9bpp, little-endian
	AVPixFmtGray9Le AVPixelFormat = C.AV_PIX_FMT_GRAY9LE
	// AVPixFmtGbrpf32Be wraps AV_PIX_FMT_GBRPF32BE.
	//
	//	IEEE-754 single precision planar GBR 4:4:4,     96bpp, big-endian
	AVPixFmtGbrpf32Be AVPixelFormat = C.AV_PIX_FMT_GBRPF32BE
	// AVPixFmtGbrpf32Le wraps AV_PIX_FMT_GBRPF32LE.
	//
	//	IEEE-754 single precision planar GBR 4:4:4,     96bpp, little-endian
	AVPixFmtGbrpf32Le AVPixelFormat = C.AV_PIX_FMT_GBRPF32LE
	// AVPixFmtGbrapf32Be wraps AV_PIX_FMT_GBRAPF32BE.
	//
	//	IEEE-754 single precision planar GBRA 4:4:4:4, 128bpp, big-endian
	AVPixFmtGbrapf32Be AVPixelFormat = C.AV_PIX_FMT_GBRAPF32BE
	// AVPixFmtGbrapf32Le wraps AV_PIX_FMT_GBRAPF32LE.
	//
	//	IEEE-754 single precision planar GBRA 4:4:4:4, 128bpp, little-endian
	AVPixFmtGbrapf32Le AVPixelFormat = C.AV_PIX_FMT_GBRAPF32LE
	// AVPixFmtDrmPrime wraps AV_PIX_FMT_DRM_PRIME.
	/*
	   DRM-managed buffers exposed through PRIME buffer sharing.

	   data[0] points to an AVDRMFrameDescriptor.
	*/
	AVPixFmtDrmPrime AVPixelFormat = C.AV_PIX_FMT_DRM_PRIME
	// AVPixFmtOpencl wraps AV_PIX_FMT_OPENCL.
	/*
	   Hardware surfaces for OpenCL.

	   data[i] contain 2D image objects (typed in C as cl_mem, used
	   in OpenCL as image2d_t) for each plane of the surface.
	*/
	AVPixFmtOpencl AVPixelFormat = C.AV_PIX_FMT_OPENCL
	// AVPixFmtGray14Be wraps AV_PIX_FMT_GRAY14BE.
	//
	//	Y        , 14bpp, big-endian
	AVPixFmtGray14Be AVPixelFormat = C.AV_PIX_FMT_GRAY14BE
	// AVPixFmtGray14Le wraps AV_PIX_FMT_GRAY14LE.
	//
	//	Y        , 14bpp, little-endian
	AVPixFmtGray14Le AVPixelFormat = C.AV_PIX_FMT_GRAY14LE
	// AVPixFmtGrayf32Be wraps AV_PIX_FMT_GRAYF32BE.
	//
	//	IEEE-754 single precision Y, 32bpp, big-endian
	AVPixFmtGrayf32Be AVPixelFormat = C.AV_PIX_FMT_GRAYF32BE
	// AVPixFmtGrayf32Le wraps AV_PIX_FMT_GRAYF32LE.
	//
	//	IEEE-754 single precision Y, 32bpp, little-endian
	AVPixFmtGrayf32Le AVPixelFormat = C.AV_PIX_FMT_GRAYF32LE
	// AVPixFmtYuva422P12Be wraps AV_PIX_FMT_YUVA422P12BE.
	//
	//	planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), 12b alpha, big-endian
	AVPixFmtYuva422P12Be AVPixelFormat = C.AV_PIX_FMT_YUVA422P12BE
	// AVPixFmtYuva422P12Le wraps AV_PIX_FMT_YUVA422P12LE.
	//
	//	planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), 12b alpha, little-endian
	AVPixFmtYuva422P12Le AVPixelFormat = C.AV_PIX_FMT_YUVA422P12LE
	// AVPixFmtYuva444P12Be wraps AV_PIX_FMT_YUVA444P12BE.
	//
	//	planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), 12b alpha, big-endian
	AVPixFmtYuva444P12Be AVPixelFormat = C.AV_PIX_FMT_YUVA444P12BE
	// AVPixFmtYuva444P12Le wraps AV_PIX_FMT_YUVA444P12LE.
	//
	//	planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), 12b alpha, little-endian
	AVPixFmtYuva444P12Le AVPixelFormat = C.AV_PIX_FMT_YUVA444P12LE
	// AVPixFmtNv24 wraps AV_PIX_FMT_NV24.
	//
	//	planar YUV 4:4:4, 24bpp, 1 plane for Y and 1 plane for the UV components, which are interleaved (first byte U and the following byte V)
	AVPixFmtNv24 AVPixelFormat = C.AV_PIX_FMT_NV24
	// AVPixFmtNv42 wraps AV_PIX_FMT_NV42.
	//
	//	as above, but U and V bytes are swapped
	AVPixFmtNv42 AVPixelFormat = C.AV_PIX_FMT_NV42
	// AVPixFmtVulkan wraps AV_PIX_FMT_VULKAN.
	/*
	   Vulkan hardware images.

	   data[0] points to an AVVkFrame
	*/
	AVPixFmtVulkan AVPixelFormat = C.AV_PIX_FMT_VULKAN
	// AVPixFmtY210Be wraps AV_PIX_FMT_Y210BE.
	//
	//	packed YUV 4:2:2 like YUYV422, 20bpp, data in the high bits, big-endian
	AVPixFmtY210Be AVPixelFormat = C.AV_PIX_FMT_Y210BE
	// AVPixFmtY210Le wraps AV_PIX_FMT_Y210LE.
	//
	//	packed YUV 4:2:2 like YUYV422, 20bpp, data in the high bits, little-endian
	AVPixFmtY210Le AVPixelFormat = C.AV_PIX_FMT_Y210LE
	// AVPixFmtX2Rgb10Le wraps AV_PIX_FMT_X2RGB10LE.
	//
	//	packed RGB 10:10:10, 30bpp, (msb)2X 10R 10G 10B(lsb), little-endian, X=unused/undefined
	AVPixFmtX2Rgb10Le AVPixelFormat = C.AV_PIX_FMT_X2RGB10LE
	// AVPixFmtX2Rgb10Be wraps AV_PIX_FMT_X2RGB10BE.
	//
	//	packed RGB 10:10:10, 30bpp, (msb)2X 10R 10G 10B(lsb), big-endian, X=unused/undefined
	AVPixFmtX2Rgb10Be AVPixelFormat = C.AV_PIX_FMT_X2RGB10BE
	// AVPixFmtX2Bgr10Le wraps AV_PIX_FMT_X2BGR10LE.
	//
	//	packed BGR 10:10:10, 30bpp, (msb)2X 10B 10G 10R(lsb), little-endian, X=unused/undefined
	AVPixFmtX2Bgr10Le AVPixelFormat = C.AV_PIX_FMT_X2BGR10LE
	// AVPixFmtX2Bgr10Be wraps AV_PIX_FMT_X2BGR10BE.
	//
	//	packed BGR 10:10:10, 30bpp, (msb)2X 10B 10G 10R(lsb), big-endian, X=unused/undefined
	AVPixFmtX2Bgr10Be AVPixelFormat = C.AV_PIX_FMT_X2BGR10BE
	// AVPixFmtP210Be wraps AV_PIX_FMT_P210BE.
	//
	//	interleaved chroma YUV 4:2:2, 20bpp, data in the high bits, big-endian
	AVPixFmtP210Be AVPixelFormat = C.AV_PIX_FMT_P210BE
	// AVPixFmtP210Le wraps AV_PIX_FMT_P210LE.
	//
	//	interleaved chroma YUV 4:2:2, 20bpp, data in the high bits, little-endian
	AVPixFmtP210Le AVPixelFormat = C.AV_PIX_FMT_P210LE
	// AVPixFmtP410Be wraps AV_PIX_FMT_P410BE.
	//
	//	interleaved chroma YUV 4:4:4, 30bpp, data in the high bits, big-endian
	AVPixFmtP410Be AVPixelFormat = C.AV_PIX_FMT_P410BE
	// AVPixFmtP410Le wraps AV_PIX_FMT_P410LE.
	//
	//	interleaved chroma YUV 4:4:4, 30bpp, data in the high bits, little-endian
	AVPixFmtP410Le AVPixelFormat = C.AV_PIX_FMT_P410LE
	// AVPixFmtP216Be wraps AV_PIX_FMT_P216BE.
	//
	//	interleaved chroma YUV 4:2:2, 32bpp, big-endian
	AVPixFmtP216Be AVPixelFormat = C.AV_PIX_FMT_P216BE
	// AVPixFmtP216Le wraps AV_PIX_FMT_P216LE.
	//
	//	interleaved chroma YUV 4:2:2, 32bpp, little-endian
	AVPixFmtP216Le AVPixelFormat = C.AV_PIX_FMT_P216LE
	// AVPixFmtP416Be wraps AV_PIX_FMT_P416BE.
	//
	//	interleaved chroma YUV 4:4:4, 48bpp, big-endian
	AVPixFmtP416Be AVPixelFormat = C.AV_PIX_FMT_P416BE
	// AVPixFmtP416Le wraps AV_PIX_FMT_P416LE.
	//
	//	interleaved chroma YUV 4:4:4, 48bpp, little-endian
	AVPixFmtP416Le AVPixelFormat = C.AV_PIX_FMT_P416LE
	// AVPixFmtVuya wraps AV_PIX_FMT_VUYA.
	//
	//	packed VUYA 4:4:4, 32bpp, VUYAVUYA...
	AVPixFmtVuya AVPixelFormat = C.AV_PIX_FMT_VUYA
	// AVPixFmtRgbaf16Be wraps AV_PIX_FMT_RGBAF16BE.
	//
	//	IEEE-754 half precision packed RGBA 16:16:16:16, 64bpp, RGBARGBA..., big-endian
	AVPixFmtRgbaf16Be AVPixelFormat = C.AV_PIX_FMT_RGBAF16BE
	// AVPixFmtRgbaf16Le wraps AV_PIX_FMT_RGBAF16LE.
	//
	//	IEEE-754 half precision packed RGBA 16:16:16:16, 64bpp, RGBARGBA..., little-endian
	AVPixFmtRgbaf16Le AVPixelFormat = C.AV_PIX_FMT_RGBAF16LE
	// AVPixFmtVuyx wraps AV_PIX_FMT_VUYX.
	//
	//	packed VUYX 4:4:4, 32bpp, Variant of VUYA where alpha channel is left undefined
	AVPixFmtVuyx AVPixelFormat = C.AV_PIX_FMT_VUYX
	// AVPixFmtP012Le wraps AV_PIX_FMT_P012LE.
	//
	//	like NV12, with 12bpp per component, data in the high bits, zeros in the low bits, little-endian
	AVPixFmtP012Le AVPixelFormat = C.AV_PIX_FMT_P012LE
	// AVPixFmtP012Be wraps AV_PIX_FMT_P012BE.
	//
	//	like NV12, with 12bpp per component, data in the high bits, zeros in the low bits, big-endian
	AVPixFmtP012Be AVPixelFormat = C.AV_PIX_FMT_P012BE
	// AVPixFmtY212Be wraps AV_PIX_FMT_Y212BE.
	//
	//	packed YUV 4:2:2 like YUYV422, 24bpp, data in the high bits, zeros in the low bits, big-endian
	AVPixFmtY212Be AVPixelFormat = C.AV_PIX_FMT_Y212BE
	// AVPixFmtY212Le wraps AV_PIX_FMT_Y212LE.
	//
	//	packed YUV 4:2:2 like YUYV422, 24bpp, data in the high bits, zeros in the low bits, little-endian
	AVPixFmtY212Le AVPixelFormat = C.AV_PIX_FMT_Y212LE
	// AVPixFmtXv30Be wraps AV_PIX_FMT_XV30BE.
	//
	//	packed XVYU 4:4:4, 32bpp, (msb)2X 10V 10Y 10U(lsb), big-endian, variant of Y410 where alpha channel is left undefined
	AVPixFmtXv30Be AVPixelFormat = C.AV_PIX_FMT_XV30BE
	// AVPixFmtXv30Le wraps AV_PIX_FMT_XV30LE.
	//
	//	packed XVYU 4:4:4, 32bpp, (msb)2X 10V 10Y 10U(lsb), little-endian, variant of Y410 where alpha channel is left undefined
	AVPixFmtXv30Le AVPixelFormat = C.AV_PIX_FMT_XV30LE
	// AVPixFmtXv36Be wraps AV_PIX_FMT_XV36BE.
	//
	//	packed XVYU 4:4:4, 48bpp, data in the high bits, zeros in the low bits, big-endian, variant of Y412 where alpha channel is left undefined
	AVPixFmtXv36Be AVPixelFormat = C.AV_PIX_FMT_XV36BE
	// AVPixFmtXv36Le wraps AV_PIX_FMT_XV36LE.
	//
	//	packed XVYU 4:4:4, 48bpp, data in the high bits, zeros in the low bits, little-endian, variant of Y412 where alpha channel is left undefined
	AVPixFmtXv36Le AVPixelFormat = C.AV_PIX_FMT_XV36LE
	// AVPixFmtRgbf32Be wraps AV_PIX_FMT_RGBF32BE.
	//
	//	IEEE-754 single precision packed RGB 32:32:32, 96bpp, RGBRGB..., big-endian
	AVPixFmtRgbf32Be AVPixelFormat = C.AV_PIX_FMT_RGBF32BE
	// AVPixFmtRgbf32Le wraps AV_PIX_FMT_RGBF32LE.
	//
	//	IEEE-754 single precision packed RGB 32:32:32, 96bpp, RGBRGB..., little-endian
	AVPixFmtRgbf32Le AVPixelFormat = C.AV_PIX_FMT_RGBF32LE
	// AVPixFmtRgbaf32Be wraps AV_PIX_FMT_RGBAF32BE.
	//
	//	IEEE-754 single precision packed RGBA 32:32:32:32, 128bpp, RGBARGBA..., big-endian
	AVPixFmtRgbaf32Be AVPixelFormat = C.AV_PIX_FMT_RGBAF32BE
	// AVPixFmtRgbaf32Le wraps AV_PIX_FMT_RGBAF32LE.
	//
	//	IEEE-754 single precision packed RGBA 32:32:32:32, 128bpp, RGBARGBA..., little-endian
	AVPixFmtRgbaf32Le AVPixelFormat = C.AV_PIX_FMT_RGBAF32LE
	// AVPixFmtNb wraps AV_PIX_FMT_NB.
	//
	//	number of pixel formats, DO NOT USE THIS if you want to link with shared libav* because the number of formats might differ between versions
	AVPixFmtNb AVPixelFormat = C.AV_PIX_FMT_NB
)

// --- Enum AVColorPrimaries ---

// AVColorPrimaries wraps AVColorPrimaries.
/*
  Chromaticity coordinates of the source primaries.
  These values match the ones defined by ISO/IEC 23091-2_2019 subclause 8.1 and ITU-T H.273.
*/
type AVColorPrimaries C.enum_AVColorPrimaries

const SizeOfAVColorPrimaries = C.sizeof_enum_AVColorPrimaries

func ToAVColorPrimariesArray(ptr unsafe.Pointer) *Array[AVColorPrimaries] {
	if ptr == nil {
		return nil
	}

	return &Array[AVColorPrimaries]{
		elemSize: SizeOfAVColorPrimaries,
		loadPtr: func(pointer unsafe.Pointer) AVColorPrimaries {
			ptr := (*AVColorPrimaries)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVColorPrimaries) {
			ptr := (*AVColorPrimaries)(pointer)
			*ptr = value
		},
	}
}

func AllocAVColorPrimariesArray(size uint64) *Array[AVColorPrimaries] {
	return ToAVColorPrimariesArray(AVCalloc(size, SizeOfAVColorPrimaries))
}

const (
	// AVColPriReserved0 wraps AVCOL_PRI_RESERVED0.
	AVColPriReserved0 AVColorPrimaries = C.AVCOL_PRI_RESERVED0
	// AVColPriBt709 wraps AVCOL_PRI_BT709.
	//
	//	also ITU-R BT1361 / IEC 61966-2-4 / SMPTE RP 177 Annex B
	AVColPriBt709 AVColorPrimaries = C.AVCOL_PRI_BT709
	// AVColPriUnspecified wraps AVCOL_PRI_UNSPECIFIED.
	AVColPriUnspecified AVColorPrimaries = C.AVCOL_PRI_UNSPECIFIED
	// AVColPriReserved wraps AVCOL_PRI_RESERVED.
	AVColPriReserved AVColorPrimaries = C.AVCOL_PRI_RESERVED
	// AVColPriBt470M wraps AVCOL_PRI_BT470M.
	//
	//	also FCC Title 47 Code of Federal Regulations 73.682 (a)(20)
	AVColPriBt470M AVColorPrimaries = C.AVCOL_PRI_BT470M
	// AVColPriBt470Bg wraps AVCOL_PRI_BT470BG.
	//
	//	also ITU-R BT601-6 625 / ITU-R BT1358 625 / ITU-R BT1700 625 PAL & SECAM
	AVColPriBt470Bg AVColorPrimaries = C.AVCOL_PRI_BT470BG
	// AVColPriSmpte170M wraps AVCOL_PRI_SMPTE170M.
	//
	//	also ITU-R BT601-6 525 / ITU-R BT1358 525 / ITU-R BT1700 NTSC
	AVColPriSmpte170M AVColorPrimaries = C.AVCOL_PRI_SMPTE170M
	// AVColPriSmpte240M wraps AVCOL_PRI_SMPTE240M.
	//
	//	identical to above, also called "SMPTE C" even though it uses D65
	AVColPriSmpte240M AVColorPrimaries = C.AVCOL_PRI_SMPTE240M
	// AVColPriFilm wraps AVCOL_PRI_FILM.
	//
	//	colour filters using Illuminant C
	AVColPriFilm AVColorPrimaries = C.AVCOL_PRI_FILM
	// AVColPriBt2020 wraps AVCOL_PRI_BT2020.
	//
	//	ITU-R BT2020
	AVColPriBt2020 AVColorPrimaries = C.AVCOL_PRI_BT2020
	// AVColPriSmpte428 wraps AVCOL_PRI_SMPTE428.
	//
	//	SMPTE ST 428-1 (CIE 1931 XYZ)
	AVColPriSmpte428 AVColorPrimaries = C.AVCOL_PRI_SMPTE428
	// AVColPriSmptest4281 wraps AVCOL_PRI_SMPTEST428_1.
	AVColPriSmptest4281 AVColorPrimaries = C.AVCOL_PRI_SMPTEST428_1
	// AVColPriSmpte431 wraps AVCOL_PRI_SMPTE431.
	//
	//	SMPTE ST 431-2 (2011) / DCI P3
	AVColPriSmpte431 AVColorPrimaries = C.AVCOL_PRI_SMPTE431
	// AVColPriSmpte432 wraps AVCOL_PRI_SMPTE432.
	//
	//	SMPTE ST 432-1 (2010) / P3 D65 / Display P3
	AVColPriSmpte432 AVColorPrimaries = C.AVCOL_PRI_SMPTE432
	// AVColPriEbu3213 wraps AVCOL_PRI_EBU3213.
	//
	//	EBU Tech. 3213-E (nothing there) / one of JEDEC P22 group phosphors
	AVColPriEbu3213 AVColorPrimaries = C.AVCOL_PRI_EBU3213
	// AVColPriJedecP22 wraps AVCOL_PRI_JEDEC_P22.
	AVColPriJedecP22 AVColorPrimaries = C.AVCOL_PRI_JEDEC_P22
	// AVColPriNb wraps AVCOL_PRI_NB.
	//
	//	Not part of ABI
	AVColPriNb AVColorPrimaries = C.AVCOL_PRI_NB
)

// --- Enum AVColorTransferCharacteristic ---

// AVColorTransferCharacteristic wraps AVColorTransferCharacteristic.
/*
  Color Transfer Characteristic.
  These values match the ones defined by ISO/IEC 23091-2_2019 subclause 8.2.
*/
type AVColorTransferCharacteristic C.enum_AVColorTransferCharacteristic

const SizeOfAVColorTransferCharacteristic = C.sizeof_enum_AVColorTransferCharacteristic

func ToAVColorTransferCharacteristicArray(ptr unsafe.Pointer) *Array[AVColorTransferCharacteristic] {
	if ptr == nil {
		return nil
	}

	return &Array[AVColorTransferCharacteristic]{
		elemSize: SizeOfAVColorTransferCharacteristic,
		loadPtr: func(pointer unsafe.Pointer) AVColorTransferCharacteristic {
			ptr := (*AVColorTransferCharacteristic)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVColorTransferCharacteristic) {
			ptr := (*AVColorTransferCharacteristic)(pointer)
			*ptr = value
		},
	}
}

func AllocAVColorTransferCharacteristicArray(size uint64) *Array[AVColorTransferCharacteristic] {
	return ToAVColorTransferCharacteristicArray(AVCalloc(size, SizeOfAVColorTransferCharacteristic))
}

const (
	// AVColTrcReserved0 wraps AVCOL_TRC_RESERVED0.
	AVColTrcReserved0 AVColorTransferCharacteristic = C.AVCOL_TRC_RESERVED0
	// AVColTrcBt709 wraps AVCOL_TRC_BT709.
	//
	//	also ITU-R BT1361
	AVColTrcBt709 AVColorTransferCharacteristic = C.AVCOL_TRC_BT709
	// AVColTrcUnspecified wraps AVCOL_TRC_UNSPECIFIED.
	AVColTrcUnspecified AVColorTransferCharacteristic = C.AVCOL_TRC_UNSPECIFIED
	// AVColTrcReserved wraps AVCOL_TRC_RESERVED.
	AVColTrcReserved AVColorTransferCharacteristic = C.AVCOL_TRC_RESERVED
	// AVColTrcGamma22 wraps AVCOL_TRC_GAMMA22.
	//
	//	also ITU-R BT470M / ITU-R BT1700 625 PAL & SECAM
	AVColTrcGamma22 AVColorTransferCharacteristic = C.AVCOL_TRC_GAMMA22
	// AVColTrcGamma28 wraps AVCOL_TRC_GAMMA28.
	//
	//	also ITU-R BT470BG
	AVColTrcGamma28 AVColorTransferCharacteristic = C.AVCOL_TRC_GAMMA28
	// AVColTrcSmpte170M wraps AVCOL_TRC_SMPTE170M.
	//
	//	also ITU-R BT601-6 525 or 625 / ITU-R BT1358 525 or 625 / ITU-R BT1700 NTSC
	AVColTrcSmpte170M AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE170M
	// AVColTrcSmpte240M wraps AVCOL_TRC_SMPTE240M.
	AVColTrcSmpte240M AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE240M
	// AVColTrcLinear wraps AVCOL_TRC_LINEAR.
	//
	//	"Linear transfer characteristics"
	AVColTrcLinear AVColorTransferCharacteristic = C.AVCOL_TRC_LINEAR
	// AVColTrcLog wraps AVCOL_TRC_LOG.
	//
	//	"Logarithmic transfer characteristic (100:1 range)"
	AVColTrcLog AVColorTransferCharacteristic = C.AVCOL_TRC_LOG
	// AVColTrcLogSqrt wraps AVCOL_TRC_LOG_SQRT.
	//
	//	"Logarithmic transfer characteristic (100 * Sqrt(10) : 1 range)"
	AVColTrcLogSqrt AVColorTransferCharacteristic = C.AVCOL_TRC_LOG_SQRT
	// AVColTrcIec6196624 wraps AVCOL_TRC_IEC61966_2_4.
	//
	//	IEC 61966-2-4
	AVColTrcIec6196624 AVColorTransferCharacteristic = C.AVCOL_TRC_IEC61966_2_4
	// AVColTrcBt1361Ecg wraps AVCOL_TRC_BT1361_ECG.
	//
	//	ITU-R BT1361 Extended Colour Gamut
	AVColTrcBt1361Ecg AVColorTransferCharacteristic = C.AVCOL_TRC_BT1361_ECG
	// AVColTrcIec6196621 wraps AVCOL_TRC_IEC61966_2_1.
	//
	//	IEC 61966-2-1 (sRGB or sYCC)
	AVColTrcIec6196621 AVColorTransferCharacteristic = C.AVCOL_TRC_IEC61966_2_1
	// AVColTrcBt202010 wraps AVCOL_TRC_BT2020_10.
	//
	//	ITU-R BT2020 for 10-bit system
	AVColTrcBt202010 AVColorTransferCharacteristic = C.AVCOL_TRC_BT2020_10
	// AVColTrcBt202012 wraps AVCOL_TRC_BT2020_12.
	//
	//	ITU-R BT2020 for 12-bit system
	AVColTrcBt202012 AVColorTransferCharacteristic = C.AVCOL_TRC_BT2020_12
	// AVColTrcSmpte2084 wraps AVCOL_TRC_SMPTE2084.
	//
	//	SMPTE ST 2084 for 10-, 12-, 14- and 16-bit systems
	AVColTrcSmpte2084 AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE2084
	// AVColTrcSmptest2084 wraps AVCOL_TRC_SMPTEST2084.
	AVColTrcSmptest2084 AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTEST2084
	// AVColTrcSmpte428 wraps AVCOL_TRC_SMPTE428.
	//
	//	SMPTE ST 428-1
	AVColTrcSmpte428 AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTE428
	// AVColTrcSmptest4281 wraps AVCOL_TRC_SMPTEST428_1.
	AVColTrcSmptest4281 AVColorTransferCharacteristic = C.AVCOL_TRC_SMPTEST428_1
	// AVColTrcAribStdB67 wraps AVCOL_TRC_ARIB_STD_B67.
	//
	//	ARIB STD-B67, known as "Hybrid log-gamma"
	AVColTrcAribStdB67 AVColorTransferCharacteristic = C.AVCOL_TRC_ARIB_STD_B67
	// AVColTrcNb wraps AVCOL_TRC_NB.
	//
	//	Not part of ABI
	AVColTrcNb AVColorTransferCharacteristic = C.AVCOL_TRC_NB
)

// --- Enum AVColorSpace ---

// AVColorSpace wraps AVColorSpace.
/*
  YUV colorspace type.
  These values match the ones defined by ISO/IEC 23091-2_2019 subclause 8.3.
*/
type AVColorSpace C.enum_AVColorSpace

const SizeOfAVColorSpace = C.sizeof_enum_AVColorSpace

func ToAVColorSpaceArray(ptr unsafe.Pointer) *Array[AVColorSpace] {
	if ptr == nil {
		return nil
	}

	return &Array[AVColorSpace]{
		elemSize: SizeOfAVColorSpace,
		loadPtr: func(pointer unsafe.Pointer) AVColorSpace {
			ptr := (*AVColorSpace)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVColorSpace) {
			ptr := (*AVColorSpace)(pointer)
			*ptr = value
		},
	}
}

func AllocAVColorSpaceArray(size uint64) *Array[AVColorSpace] {
	return ToAVColorSpaceArray(AVCalloc(size, SizeOfAVColorSpace))
}

const (
	// AVColSpcRgb wraps AVCOL_SPC_RGB.
	//
	//	order of coefficients is actually GBR, also IEC 61966-2-1 (sRGB), YZX and ST 428-1
	AVColSpcRgb AVColorSpace = C.AVCOL_SPC_RGB
	// AVColSpcBt709 wraps AVCOL_SPC_BT709.
	//
	//	also ITU-R BT1361 / IEC 61966-2-4 xvYCC709 / derived in SMPTE RP 177 Annex B
	AVColSpcBt709 AVColorSpace = C.AVCOL_SPC_BT709
	// AVColSpcUnspecified wraps AVCOL_SPC_UNSPECIFIED.
	AVColSpcUnspecified AVColorSpace = C.AVCOL_SPC_UNSPECIFIED
	// AVColSpcReserved wraps AVCOL_SPC_RESERVED.
	//
	//	reserved for future use by ITU-T and ISO/IEC just like 15-255 are
	AVColSpcReserved AVColorSpace = C.AVCOL_SPC_RESERVED
	// AVColSpcFcc wraps AVCOL_SPC_FCC.
	//
	//	FCC Title 47 Code of Federal Regulations 73.682 (a)(20)
	AVColSpcFcc AVColorSpace = C.AVCOL_SPC_FCC
	// AVColSpcBt470Bg wraps AVCOL_SPC_BT470BG.
	//
	//	also ITU-R BT601-6 625 / ITU-R BT1358 625 / ITU-R BT1700 625 PAL & SECAM / IEC 61966-2-4 xvYCC601
	AVColSpcBt470Bg AVColorSpace = C.AVCOL_SPC_BT470BG
	// AVColSpcSmpte170M wraps AVCOL_SPC_SMPTE170M.
	//
	//	also ITU-R BT601-6 525 / ITU-R BT1358 525 / ITU-R BT1700 NTSC / functionally identical to above
	AVColSpcSmpte170M AVColorSpace = C.AVCOL_SPC_SMPTE170M
	// AVColSpcSmpte240M wraps AVCOL_SPC_SMPTE240M.
	//
	//	derived from 170M primaries and D65 white point, 170M is derived from BT470 System M's primaries
	AVColSpcSmpte240M AVColorSpace = C.AVCOL_SPC_SMPTE240M
	// AVColSpcYcgco wraps AVCOL_SPC_YCGCO.
	//
	//	used by Dirac / VC-2 and H.264 FRext, see ITU-T SG16
	AVColSpcYcgco AVColorSpace = C.AVCOL_SPC_YCGCO
	// AVColSpcYcocg wraps AVCOL_SPC_YCOCG.
	AVColSpcYcocg AVColorSpace = C.AVCOL_SPC_YCOCG
	// AVColSpcBt2020Ncl wraps AVCOL_SPC_BT2020_NCL.
	//
	//	ITU-R BT2020 non-constant luminance system
	AVColSpcBt2020Ncl AVColorSpace = C.AVCOL_SPC_BT2020_NCL
	// AVColSpcBt2020Cl wraps AVCOL_SPC_BT2020_CL.
	//
	//	ITU-R BT2020 constant luminance system
	AVColSpcBt2020Cl AVColorSpace = C.AVCOL_SPC_BT2020_CL
	// AVColSpcSmpte2085 wraps AVCOL_SPC_SMPTE2085.
	//
	//	SMPTE 2085, Y'D'zD'x
	AVColSpcSmpte2085 AVColorSpace = C.AVCOL_SPC_SMPTE2085
	// AVColSpcChromaDerivedNcl wraps AVCOL_SPC_CHROMA_DERIVED_NCL.
	//
	//	Chromaticity-derived non-constant luminance system
	AVColSpcChromaDerivedNcl AVColorSpace = C.AVCOL_SPC_CHROMA_DERIVED_NCL
	// AVColSpcChromaDerivedCl wraps AVCOL_SPC_CHROMA_DERIVED_CL.
	//
	//	Chromaticity-derived constant luminance system
	AVColSpcChromaDerivedCl AVColorSpace = C.AVCOL_SPC_CHROMA_DERIVED_CL
	// AVColSpcIctcp wraps AVCOL_SPC_ICTCP.
	//
	//	ITU-R BT.2100-0, ICtCp
	AVColSpcIctcp AVColorSpace = C.AVCOL_SPC_ICTCP
	// AVColSpcNb wraps AVCOL_SPC_NB.
	//
	//	Not part of ABI
	AVColSpcNb AVColorSpace = C.AVCOL_SPC_NB
)

// --- Enum AVColorRange ---

// AVColorRange wraps AVColorRange.
/*
  Visual content value range.

  These values are based on definitions that can be found in multiple
  specifications, such as ITU-T BT.709 (3.4 - Quantization of RGB, luminance
  and colour-difference signals), ITU-T BT.2020 (Table 5 - Digital
  Representation) as well as ITU-T BT.2100 (Table 9 - Digital 10- and 12-bit
  integer representation). At the time of writing, the BT.2100 one is
  recommended, as it also defines the full range representation.

  Common definitions:
    - For RGB and luma planes such as Y in YCbCr and I in ICtCp,
      'E' is the original value in range of 0.0 to 1.0.
    - For chroma planes such as Cb,Cr and Ct,Cp, 'E' is the original
      value in range of -0.5 to 0.5.
    - 'n' is the output bit depth.
    - For additional definitions such as rounding and clipping to valid n
      bit unsigned integer range, please refer to BT.2100 (Table 9).
*/
type AVColorRange C.enum_AVColorRange

const SizeOfAVColorRange = C.sizeof_enum_AVColorRange

func ToAVColorRangeArray(ptr unsafe.Pointer) *Array[AVColorRange] {
	if ptr == nil {
		return nil
	}

	return &Array[AVColorRange]{
		elemSize: SizeOfAVColorRange,
		loadPtr: func(pointer unsafe.Pointer) AVColorRange {
			ptr := (*AVColorRange)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVColorRange) {
			ptr := (*AVColorRange)(pointer)
			*ptr = value
		},
	}
}

func AllocAVColorRangeArray(size uint64) *Array[AVColorRange] {
	return ToAVColorRangeArray(AVCalloc(size, SizeOfAVColorRange))
}

const (
	// AVColRangeUnspecified wraps AVCOL_RANGE_UNSPECIFIED.
	AVColRangeUnspecified AVColorRange = C.AVCOL_RANGE_UNSPECIFIED
	// AVColRangeMpeg wraps AVCOL_RANGE_MPEG.
	/*
	   Narrow or limited range content.

	   - For luma planes:

	         (219 * E + 16) * 2^(n-8)

	     F.ex. the range of 16-235 for 8 bits

	   - For chroma planes:

	         (224 * E + 128) * 2^(n-8)

	     F.ex. the range of 16-240 for 8 bits
	*/
	AVColRangeMpeg AVColorRange = C.AVCOL_RANGE_MPEG
	// AVColRangeJpeg wraps AVCOL_RANGE_JPEG.
	/*
	   Full range content.

	   - For RGB and luma planes:

	         (2^n - 1) * E

	     F.ex. the range of 0-255 for 8 bits

	   - For chroma planes:

	         (2^n - 1) * E + 2^(n - 1)

	     F.ex. the range of 1-255 for 8 bits
	*/
	AVColRangeJpeg AVColorRange = C.AVCOL_RANGE_JPEG
	// AVColRangeNb wraps AVCOL_RANGE_NB.
	//
	//	Not part of ABI
	AVColRangeNb AVColorRange = C.AVCOL_RANGE_NB
)

// --- Enum AVChromaLocation ---

// AVChromaLocation wraps AVChromaLocation.
/*
  Location of chroma samples.

  Illustration showing the location of the first (top left) chroma sample of the
  image, the left shows only luma, the right
  shows the location of the chroma sample, the 2 could be imagined to overlay
  each other but are drawn separately due to limitations of ASCII

                 1st 2nd       1st 2nd horizontal luma sample positions
                  v   v         v   v
                  ______        ______
  *1st luma line > |X   X ...    |3 4 X ...     X are luma samples,
                 |             |1 2           1-6 are possible chroma positions
  *2nd luma line > |X   X ...    |5 6 X ...     0 is undefined/unknown position
*/
type AVChromaLocation C.enum_AVChromaLocation

const SizeOfAVChromaLocation = C.sizeof_enum_AVChromaLocation

func ToAVChromaLocationArray(ptr unsafe.Pointer) *Array[AVChromaLocation] {
	if ptr == nil {
		return nil
	}

	return &Array[AVChromaLocation]{
		elemSize: SizeOfAVChromaLocation,
		loadPtr: func(pointer unsafe.Pointer) AVChromaLocation {
			ptr := (*AVChromaLocation)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVChromaLocation) {
			ptr := (*AVChromaLocation)(pointer)
			*ptr = value
		},
	}
}

func AllocAVChromaLocationArray(size uint64) *Array[AVChromaLocation] {
	return ToAVChromaLocationArray(AVCalloc(size, SizeOfAVChromaLocation))
}

const (
	// AVChromaLocUnspecified wraps AVCHROMA_LOC_UNSPECIFIED.
	AVChromaLocUnspecified AVChromaLocation = C.AVCHROMA_LOC_UNSPECIFIED
	// AVChromaLocLeft wraps AVCHROMA_LOC_LEFT.
	//
	//	MPEG-2/4 4:2:0, H.264 default for 4:2:0
	AVChromaLocLeft AVChromaLocation = C.AVCHROMA_LOC_LEFT
	// AVChromaLocCenter wraps AVCHROMA_LOC_CENTER.
	//
	//	MPEG-1 4:2:0, JPEG 4:2:0, H.263 4:2:0
	AVChromaLocCenter AVChromaLocation = C.AVCHROMA_LOC_CENTER
	// AVChromaLocTopleft wraps AVCHROMA_LOC_TOPLEFT.
	//
	//	ITU-R 601, SMPTE 274M 296M S314M(DV 4:1:1), mpeg2 4:2:2
	AVChromaLocTopleft AVChromaLocation = C.AVCHROMA_LOC_TOPLEFT
	// AVChromaLocTop wraps AVCHROMA_LOC_TOP.
	AVChromaLocTop AVChromaLocation = C.AVCHROMA_LOC_TOP
	// AVChromaLocBottomleft wraps AVCHROMA_LOC_BOTTOMLEFT.
	AVChromaLocBottomleft AVChromaLocation = C.AVCHROMA_LOC_BOTTOMLEFT
	// AVChromaLocBottom wraps AVCHROMA_LOC_BOTTOM.
	AVChromaLocBottom AVChromaLocation = C.AVCHROMA_LOC_BOTTOM
	// AVChromaLocNb wraps AVCHROMA_LOC_NB.
	//
	//	Not part of ABI
	AVChromaLocNb AVChromaLocation = C.AVCHROMA_LOC_NB
)

// --- Enum AVSampleFormat ---

// AVSampleFormat wraps AVSampleFormat.
/*
  Audio sample formats

  - The data described by the sample format is always in native-endian order.
    Sample values can be expressed by native C types, hence the lack of a signed
    24-bit sample format even though it is a common raw audio data format.

  - The floating-point formats are based on full volume being in the range
    [-1.0, 1.0]. Any values outside this range are beyond full volume level.

  - The data layout as used in av_samples_fill_arrays() and elsewhere in FFmpeg
    (such as AVFrame in libavcodec) is as follows:

  @par
  For planar sample formats, each audio channel is in a separate data plane,
  and linesize is the buffer size, in bytes, for a single plane. All data
  planes must be the same size. For packed sample formats, only the first data
  plane is used, and samples for each channel are interleaved. In this case,
  linesize is the buffer size, in bytes, for the 1 plane.
*/
type AVSampleFormat C.enum_AVSampleFormat

const SizeOfAVSampleFormat = C.sizeof_enum_AVSampleFormat

func ToAVSampleFormatArray(ptr unsafe.Pointer) *Array[AVSampleFormat] {
	if ptr == nil {
		return nil
	}

	return &Array[AVSampleFormat]{
		elemSize: SizeOfAVSampleFormat,
		loadPtr: func(pointer unsafe.Pointer) AVSampleFormat {
			ptr := (*AVSampleFormat)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value AVSampleFormat) {
			ptr := (*AVSampleFormat)(pointer)
			*ptr = value
		},
	}
}

func AllocAVSampleFormatArray(size uint64) *Array[AVSampleFormat] {
	return ToAVSampleFormatArray(AVCalloc(size, SizeOfAVSampleFormat))
}

const (
	// AVSampleFmtNone wraps AV_SAMPLE_FMT_NONE.
	AVSampleFmtNone AVSampleFormat = C.AV_SAMPLE_FMT_NONE
	// AVSampleFmtU8 wraps AV_SAMPLE_FMT_U8.
	//
	//	unsigned 8 bits
	AVSampleFmtU8 AVSampleFormat = C.AV_SAMPLE_FMT_U8
	// AVSampleFmtS16 wraps AV_SAMPLE_FMT_S16.
	//
	//	signed 16 bits
	AVSampleFmtS16 AVSampleFormat = C.AV_SAMPLE_FMT_S16
	// AVSampleFmtS32 wraps AV_SAMPLE_FMT_S32.
	//
	//	signed 32 bits
	AVSampleFmtS32 AVSampleFormat = C.AV_SAMPLE_FMT_S32
	// AVSampleFmtFlt wraps AV_SAMPLE_FMT_FLT.
	//
	//	float
	AVSampleFmtFlt AVSampleFormat = C.AV_SAMPLE_FMT_FLT
	// AVSampleFmtDbl wraps AV_SAMPLE_FMT_DBL.
	//
	//	double
	AVSampleFmtDbl AVSampleFormat = C.AV_SAMPLE_FMT_DBL
	// AVSampleFmtU8P wraps AV_SAMPLE_FMT_U8P.
	//
	//	unsigned 8 bits, planar
	AVSampleFmtU8P AVSampleFormat = C.AV_SAMPLE_FMT_U8P
	// AVSampleFmtS16P wraps AV_SAMPLE_FMT_S16P.
	//
	//	signed 16 bits, planar
	AVSampleFmtS16P AVSampleFormat = C.AV_SAMPLE_FMT_S16P
	// AVSampleFmtS32P wraps AV_SAMPLE_FMT_S32P.
	//
	//	signed 32 bits, planar
	AVSampleFmtS32P AVSampleFormat = C.AV_SAMPLE_FMT_S32P
	// AVSampleFmtFltp wraps AV_SAMPLE_FMT_FLTP.
	//
	//	float, planar
	AVSampleFmtFltp AVSampleFormat = C.AV_SAMPLE_FMT_FLTP
	// AVSampleFmtDblp wraps AV_SAMPLE_FMT_DBLP.
	//
	//	double, planar
	AVSampleFmtDblp AVSampleFormat = C.AV_SAMPLE_FMT_DBLP
	// AVSampleFmtS64 wraps AV_SAMPLE_FMT_S64.
	//
	//	signed 64 bits
	AVSampleFmtS64 AVSampleFormat = C.AV_SAMPLE_FMT_S64
	// AVSampleFmtS64P wraps AV_SAMPLE_FMT_S64P.
	//
	//	signed 64 bits, planar
	AVSampleFmtS64P AVSampleFormat = C.AV_SAMPLE_FMT_S64P
	// AVSampleFmtNb wraps AV_SAMPLE_FMT_NB.
	//
	//	Number of sample formats. DO NOT USE if linking dynamically
	AVSampleFmtNb AVSampleFormat = C.AV_SAMPLE_FMT_NB
)
