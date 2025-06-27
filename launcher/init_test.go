package launcher

import "github.com/streamingfast/logging"

var zlog, _ = logging.PackageLogger("launcher", "github.com/rockawayx-labs/firehose-core/launcher")

func init() {
	logging.InstantiateLoggers()
}
