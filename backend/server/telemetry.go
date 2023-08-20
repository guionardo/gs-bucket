package server

import (
	"github.com/go-chi/telemetry"
)

var BucketMetrics = &GSBucketMetrics{telemetry.NewNamespace("bucket")}

type GSBucketMetrics struct {
	*telemetry.Namespace
}

func (m *GSBucketMetrics) RecordFileCount(count int) {
	m.RecordGauge("file_count", nil, float64(count))
}

func (m *GSBucketMetrics) RecordFileSize(size int64) {
	m.RecordGauge("file_size", nil, float64(size))
}

func (m *GSBucketMetrics) RecordUserCount(count int) {
	m.RecordGauge("user_count", nil, float64(count))
}
