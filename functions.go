package godocgen

import (
	"bytes"
	"path"
	"strings"

	"github.com/gosimple/slug"
)

var (
	funcs = map[string]interface{}{
		"base":        path.Base,
		"trim_prefix": strings.TrimPrefix,
		"comment_md":  commentMdFunc,
		"md":          mdFunc,
		"pre":         preFunc,
		"sanitize":    sanitizeFunc,
		"slugify":     slug.Make,
	}
)

func commentMdFunc(comment string) string {
	var buf bytes.Buffer
	ToMD(&buf, comment)
	return buf.String()
}

func mdFunc(text string) string {
	text = strings.Replace(text, "*", "\\*", -1)
	text = strings.Replace(text, "_", "\\_", -1)
	return text
}

func preFunc(text string) string {
	return "``` go\n" + text + "\n```"
}

// sanitizeFunc sanitizes the argument src by replacing newlines with
// blanks, removing extra blanks, and by removing trailing whitespace
// and commas before closing parentheses.
func sanitizeFunc(src string) string {
	buf := make([]byte, len(src))
	j := 0      // buf index
	comma := -1 // comma index if >= 0
	for i := 0; i < len(src); i++ {
		ch := src[i]
		switch ch {
		case '\t', '\n', ' ':
			// ignore whitespace at the beginning, after a blank, or after opening parentheses
			if j == 0 {
				continue
			}
			if p := buf[j-1]; p == ' ' || p == '(' || p == '{' || p == '[' {
				continue
			}
			// replace all whitespace with blanks
			ch = ' '
		case ',':
			comma = j
		case ')', '}', ']':
			// remove any trailing comma
			if comma >= 0 {
				j = comma
			}
			// remove any trailing whitespace
			if j > 0 && buf[j-1] == ' ' {
				j--
			}
		default:
			comma = -1
		}
		buf[j] = ch
		j++
	}
	// remove trailing blank, if any
	if j > 0 && buf[j-1] == ' ' {
		j--
	}
	return string(buf[:j])
}
