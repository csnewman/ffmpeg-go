#include <libavutil/bprint.h>
#include <libavutil/avutil.h>

void ffgLogCallback(void* ctx, int level, void* msg);

void ffg_log_callback(void* avcl, int level, const char* fmt, va_list vl) {
    // Respect log level to save formatting cost
    if (level >= 0) {
        level &= 0xff;
    }

    if (level > av_log_get_level())
        return;

    AVBPrint msg;
    char *msg_buf;

    av_bprint_init(&msg, 0, AV_BPRINT_SIZE_UNLIMITED);
	av_vbprintf(&msg, fmt, vl);
    av_bprint_finalize(&msg, &msg_buf);

    ffgLogCallback(avcl, level, msg_buf);
}

void ffg_set_log() {
    av_log_set_callback(ffg_log_callback);
}

typedef const char* (*itemNameFunc) (void* ctx);

const char* invokeItemNameFunc(itemNameFunc f, void* ctx)
{
    return f(ctx);
}
