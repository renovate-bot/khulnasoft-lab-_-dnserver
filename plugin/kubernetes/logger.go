package kubernetes

import (
	clog "github.com/khulnasoft-lab/dnserver/plugin/pkg/log"

	"github.com/go-logr/logr"
)

// loggerAdapter is a simple wrapper around DNServer plugin logger made to implement logr.LogSink interface, which is used
// as part of klog library for logging in Kubernetes client. By using this adapter DNServer is able to log messages/errors from
// kubernetes client in a DNServer logging format
type loggerAdapter struct {
	clog.P
}

func (l *loggerAdapter) Init(_ logr.RuntimeInfo) {
}

func (l *loggerAdapter) Enabled(_ int) bool {
	// verbosity is controlled inside klog library, we do not need to do anything here
	return true
}

func (l *loggerAdapter) Info(_ int, msg string, _ ...interface{}) {
	l.P.Info(msg)
}

func (l *loggerAdapter) Error(_ error, msg string, _ ...interface{}) {
	l.P.Error(msg)
}

func (l *loggerAdapter) WithValues(_ ...interface{}) logr.LogSink {
	return l
}

func (l *loggerAdapter) WithName(_ string) logr.LogSink {
	return l
}
