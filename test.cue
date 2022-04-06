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
		_source: client.filesystem.".".read.contents
		_image:  go.#Image & {version: "1.18"}

		test: go.#Test & {
			source:  _source
			package: "./..."
			input:   _image.output
		}
	}
}
