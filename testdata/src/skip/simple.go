package skip

import (
	"io"
)

func BasicWrong(rc io.ReadCloser) {
	rc.Close()
}
