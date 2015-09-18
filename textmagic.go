package textmagic

import (
	"bytes"
	"strconv"
)

func joinIntSlice(v []int) string {
	buf := &bytes.Buffer{}

	for i := range v {
		if i > 0 {
			buf.WriteByte(',')
		}

		buf.WriteString(strconv.Itoa(v[i]))
	}

	return buf.String()
}
