package kcp

import (
	"io"
	"sync"

	"github.com/whatedcgveg/v2ray-core/common/retry"

	"github.com/whatedcgveg/v2ray-core/common"
	"github.com/whatedcgveg/v2ray-core/common/buf"
)

type SegmentWriter interface {
	Write(seg Segment) error
}

type SimpleSegmentWriter struct {
	sync.Mutex
	buffer *buf.Buffer
	writer io.Writer
}

func NewSegmentWriter(writer io.Writer) SegmentWriter {
	return &SimpleSegmentWriter{
		writer: writer,
		buffer: buf.New(),
	}
}

func (w *SimpleSegmentWriter) Write(seg Segment) error {
	w.Lock()
	defer w.Unlock()

	common.Must(w.buffer.Reset(seg.Bytes()))
	_, err := w.writer.Write(w.buffer.Bytes())
	return err
}

type RetryableWriter struct {
	writer SegmentWriter
}

func NewRetryableWriter(writer SegmentWriter) SegmentWriter {
	return &RetryableWriter{
		writer: writer,
	}
}

func (w *RetryableWriter) Write(seg Segment) error {
	return retry.Timed(5, 100).On(func() error {
		return w.writer.Write(seg)
	})
}
