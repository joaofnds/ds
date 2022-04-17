package ds

import (
	"dagger.io/dagger"
	"universe.dagger.io/go"
)

dagger.#Plan & {
	client: filesystem: ".": read: {
		contents: dagger.#FS
		include: ["go.mod", "go.sum", "**/*.go"]
	}
	actions: {
		test: go.#Test & {
			source:  client.filesystem.".".read.contents
			package: "./..."
		}
	}
}
