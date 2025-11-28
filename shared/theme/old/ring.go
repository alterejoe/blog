package theme

var RING_SIZE = map[string]string{
	"sm": "ring-1",
	"md": "ring-2",
	"lg": "ring-4",
}

func RingClass(size string, color string) string {
	sz, ok := RING_SIZE[size]
	if !ok {
		sz = RING_SIZE["md"]
	}
	return Combine(
		sz,
		Ring(color),
	)
}
