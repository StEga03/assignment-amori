package errorwrapper

import (
	"runtime"
	"strings"
)

// stack represents a stack of program counters.
type stack []uintptr

func callers(pos int) *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[pos:n]
	return &st
}

func (la *errWrapper) getFuncFileLine() (string, string, int) {
	st := la.StackTrace()
	f := st[0]
	pc := uintptr(f) - 1
	fn := runtime.FuncForPC(pc)
	file, line := fn.FileLine(pc)
	return la.funcname(fn.Name()), file, line
}

func (la *errWrapper) funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}

func trimRootPath(file string) string {
	lastBin := strings.LastIndex(file, holder.Config.RepoRoot)
	if (lastBin+len(holder.Config.RepoRoot)) > len(file) || lastBin == -1 {
		return file
	}
	return file[lastBin+len(holder.Config.RepoRoot):]
}

// getMetadata get metadata with string JSON format.
func (la *errWrapper) getMetadata() MetaKV {
	metadata := make(MetaKV)
	for key, val := range la.metadata {
		metadata[key] = val
	}

	metadata[metaKeyDevMsg] = la.devMessage
	metadata[metaKeyErrCode] = la.Code()
	metadata[metaKeyErrID] = la.GetErrorID()
	metadata[metaKeyHTTPStatus] = la.GetHTTPStatus()
	metadata[metaKeyIsRetryable] = la.GetIsRetryable()
	metadata[metaKeyErrLine] = la.GetErrorLine()
	return metadata
}
