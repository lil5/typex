// MIT license · Daniel T. Gorski · dtg [at] lengo [dot] org · 06/2020

package ts

import (
	"fmt"
	"io"

	typex "github.com/dtgorski/typex/internal"
)

type (
	moduleLayout struct {
		writer io.Writer
		indent int
	}
)

// NewModuleLayout implements the Layout interface.
func NewModuleLayout(w io.Writer) typex.Layout {
	return &moduleLayout{writer: w}
}

func (m *moduleLayout) Enter(path string, _ bool) {
	m.indent++
}

func (m *moduleLayout) Print(line string, _, _ bool) {
	m.write(m.writer, "%s%s\n", "", line)
}

func (m *moduleLayout) Leave(path string, _ bool) {
	m.indent--
}

func (moduleLayout) write(w io.Writer, f string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, f, a...)
}
