package main

import (
	firecore "github.com/rockawayx-labs/firehose-core"
	fhCMD "github.com/rockawayx-labs/firehose-core/cmd"
	info "github.com/rockawayx-labs/firehose-core/firehose/info"
	pbbstream "github.com/streamingfast/bstream/pb/sf/bstream/v1"
)

func main() {
	firecore.UnsafeRunningFromFirecore = true
	firecore.UnsafeAllowExecutableNameToBeEmpty = true

	fhCMD.Main(&firecore.Chain[*pbbstream.Block]{
		ShortName:            "core",
		LongName:             "CORE", //only used to compose cmd title and description
		FullyQualifiedModule: "github.com/rockawayx-labs/firehose-core",
		Version:              version,
		BlockFactory:         func() firecore.Block { return new(pbbstream.Block) },
		ConsoleReaderFactory: firecore.NewConsoleReader,
		InfoResponseFiller:   info.DefaultInfoResponseFiller,
		Tools:                &firecore.ToolsConfig[*pbbstream.Block]{},
	})
}

// Version value, injected via go build `ldflags` at build time, **must** not be removed or inlined
var version = "dev"
