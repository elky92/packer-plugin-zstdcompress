package zstdcompress

import (
	"fmt"
	"os"
)

const BuilderId = "packer.post-processor.zstdcompress"

type Artifact struct {
	Path string
}

func (a *Artifact) BuilderId() string {
	return BuilderId
}

func (*Artifact) Id() string {
	return ""
}

func (a *Artifact) Files() []string {
	return []string{a.Path}
}

func (a *Artifact) String() string {
	return fmt.Sprintf("compressed artifacts in: %s", a.Path)
}

func (*Artifact) State(name string) interface{} {
	return nil
}

func (a *Artifact) Destroy() error {
	return os.Remove(a.Path)
}
