package theme

var TEXT = map[string]string{
	"xs":   "text-xs",
	"sm":   "text-sm",
	"md":   "text-base",
	"lg":   "text-lg",
	"xl":   "text-xl",
	"bold": "font-bold",
}

func TextClass(t string) string {
	if v, ok := TEXT[t]; ok {
		return v
	}
	return TEXT["md"]
}
