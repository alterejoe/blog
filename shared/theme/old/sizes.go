package theme

var SIZES = map[string]string{
	"xs": "px-1.5 py-1 text-xs",
	"sm": "px-2 py-1.5 text-sm",
	"md": "px-3 py-2 text-base",
	"lg": "px-4 py-3 text-lg",
	"xl": "px-5 py-4 text-xl",
}

const DEFAULT_SIZE = "px-3 py-2 text-base"

func SizeClass(s string) string {
	if v, ok := SIZES[s]; ok {
		return v
	}
	return DEFAULT_SIZE
}
