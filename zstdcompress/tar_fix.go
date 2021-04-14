// +build !go1.10

package zstdcompress

import "archive/tar"

func setHeaderFormat(header *tar.Header) {
	// no-op
}
