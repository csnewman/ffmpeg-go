package ffmpeg

import (
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"unsafe"
)

/*
#include <libavutil/log.h>

void ffg_set_log();

typedef const char* (*itemNameFunc) (void* ctx);

const char* invokeItemNameFunc(itemNameFunc f, void* ctx);
*/
import "C"

type LogCallback func(ctx *LogCtx, level int, msg string)

var activeLogCallback LogCallback //nolint:gochecknoglobals

//export ffgLogCallback
func ffgLogCallback(ctxPtr unsafe.Pointer, level int, msgPtr *C.char) {
	var ctx *LogCtx

	if ctxPtr != nil {
		ctx = &LogCtx{
			ptr: ctxPtr,
		}
	}

	msgWrapper := wrapCStr(msgPtr)
	msg := msgWrapper.String()
	msgWrapper.Free()

	if activeLogCallback != nil {
		activeLogCallback(ctx, level, msg)
	}
}

// AVLogSetCallback wraps av_log_set_callback.
/*
  Set the logging callback

  @note The callback must be thread safe, even if the application does not use
        threads itself as some codecs are multithreaded.

  @see av_log_default_callback

  @param callback A logging function with a compatible signature.
*/
func AVLogSetCallback(cb LogCallback) {
	activeLogCallback = cb

	C.ffg_set_log()
}

// ItemName gets the item_name field.
func (s *AVClass) ItemName() func(pointer unsafe.Pointer) *CStr {
	fp := s.ptr.item_name

	return func(pointer unsafe.Pointer) *CStr {

		value := C.invokeItemNameFunc(fp, pointer)
		return wrapCStr(value)
	}
}

type LogCtx struct {
	ptr unsafe.Pointer
}

func (c *LogCtx) RawPtr() unsafe.Pointer {
	return c.ptr
}

func (c *LogCtx) Class() *AVClass {
	classPtr := *(**C.AVClass)(c.ptr)

	if classPtr == nil {
		return nil
	}

	return &AVClass{
		ptr: classPtr,
	}
}

func SLogAdapter(logger *slog.Logger) LogCallback {
	l := &slogLogger{
		logger: logger,
	}

	return l.callback
}

type slogLogger struct {
	logger *slog.Logger
	mu     sync.Mutex
	sb     strings.Builder
}

func (l *slogLogger) callback(ctx *LogCtx, level int, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for {
		before, after, found := strings.Cut(msg, "\n")
		msg = after

		l.sb.WriteString(before)

		if !found {
			break
		}

		scope := "global"

		if ctx != nil {
			class := ctx.Class()
			scope = class.ItemName()(ctx.RawPtr()).String()
		}

		if level >= 0 {
			level &= 0xff
		}

		printMsg := l.sb.String()
		l.sb.Reset()

		switch level {
		case AVLogQuiet:
		case AVLogPanic:
			panic(fmt.Sprintf("FFmpeg log scope=%v log=%v", scope, printMsg))
		case AVLogFatal, AVLogError:
			l.logger.Error("FFmpeg Log", "scope", scope, "log", printMsg)
		case AVLogWarning:
			l.logger.Warn("FFmpeg Log", "scope", scope, "log", printMsg)
		case AVLogDebug, AVLogTrace:
			l.logger.Debug("FFmpeg Log", "scope", scope, "log", printMsg)
		default:
			l.logger.Info("FFmpeg Log", "scope", scope, "log", printMsg)
		}
	}
}
