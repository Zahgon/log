package log

import (
	"bytes"
)

func (l *Logger) jsonFormatter(keyvals ...any) { _ = "STUB: not implemented"; return }

func (l *Logger) jsonFormatterRoot(jw *jsonWriter, key, value any) {
	_ = "STUB: not implemented"
	return
}

func (l *Logger) jsonFormatterItem(jw *jsonWriter, key, value any) {
	_ = "STUB: not implemented"
	return
}

func (l *Logger) writeSlogValue(jw *jsonWriter, v slogValue) { _ = "STUB: not implemented"; return }

type jsonWriter struct {
	w *bytes.Buffer
	d int
}

func (w *jsonWriter) start() { _ = "STUB: not implemented"; return }

func (w *jsonWriter) end() { _ = "STUB: not implemented"; return }

func (w *jsonWriter) objectItem(key string, value any) { _ = "STUB: not implemented"; return }

func (w *jsonWriter) objectKey(key string) { _ = "STUB: not implemented"; return }

func (w *jsonWriter) objectValue(value any) { _ = "STUB: not implemented"; return }

func (w *jsonWriter) writeEncoded(v any) error { _ = "STUB: not implemented"; return nil }

// trailing \n added by json.Encode
