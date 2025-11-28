package theme

var RADII = map[string]string{
	"none": "rounded-none",
	"sm":   "rounded-sm",
	"md":   "rounded-md",
	"lg":   "rounded-lg",
	"xl":   "rounded-xl",
	"full": "rounded-full",
}

const DEFAULT_RADIUS = "rounded-md"

func RadiusClass(r string) string {
	if v, ok := RADII[r]; ok {
		return v
	}
	return DEFAULT_RADIUS
}
