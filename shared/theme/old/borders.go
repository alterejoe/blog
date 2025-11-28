package theme

var BORDERS = map[string]string{
	"sm": "border-1",
	"md": "border-2",
	"lg": "border-4",
}

func BorderClass(size string, color string) string {
	sz, ok := BORDERS[size]
	if !ok {
		sz = BORDERS["md"]
	}
	return Combine(
		sz,
		Border(color),
	)
}
