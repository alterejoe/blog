package theme

var SPACING = map[string]string{
	"none":  "p-0",
	"tight": "p-1",
	"snug":  "p-2",
	"base":  "p-3",
	"roomy": "p-4",
	"big":   "p-6",
}

func SpacingClass(s string) string {
	if v, ok := SPACING[s]; ok {
		return v
	}
	return SPACING["base"]
}
