package theme

type ColorSet struct {
	Light      string
	LightHover string
	Dark       string
	DarkHover  string
}

// -----------------------------------------------------------------------------
// MASTER COLOR TOKENS — your theme lives HERE and only here.
// -----------------------------------------------------------------------------

var COLORS = map[string]ColorSet{
	"primary":   {Light: "primary", LightHover: "primary-hover", Dark: "primary-dark", DarkHover: "primary-dark-hover"},
	"secondary": {Light: "secondary", LightHover: "secondary-hover", Dark: "secondary-dark", DarkHover: "secondary-dark-hover"},
	"tertiary":  {Light: "tertiary", LightHover: "tertiary-hover", Dark: "tertiary-dark", DarkHover: "tertiary-dark-hover"},
	"accent":    {Light: "accent", LightHover: "accent-hover", Dark: "accent-dark", DarkHover: "accent-dark-hover"},
	"neutral":   {Light: "neutral", LightHover: "neutral-hover", Dark: "neutral-dark", DarkHover: "neutral-dark-hover"},
}

const DEFAULT_COLOR = "neutral"

func get(c string) ColorSet {
	v, ok := COLORS[c]
	if !ok {
		return COLORS[DEFAULT_COLOR]
	}
	return v
}

func Bg(c string) string {
	v := get(c)
	return "bg-" + v.Light +
		" hover:bg-" + v.LightHover +
		" peer-checked:bg-" + v.Light +
		" dark:bg-" + v.Dark +
		" dark:hover:bg-" + v.DarkHover +
		" dark:peer-checked:bg-" + v.Dark
}

func Hover(c string) string {
	v := get(c)
	return "hover:bg-" + v.LightHover + " dark:hover:bg-" + v.DarkHover
}

func Ring(c string) string {
	v := get(c)
	return "ring-" + v.Light + " dark:ring-" + v.Dark
}

func Border(c string) string {
	v := get(c)
	return "border-" + v.Light + " dark:border-" + v.Dark
}

func Checked(c string) string {
	v := get(c)
	return "peer-checked:bg-" + v.Light +
		" peer-checked:ring-" + v.Light +
		" dark:peer-checked:bg-" + v.Dark +
		" dark:peer-checked:ring-" + v.Dark
}

func Text(c string) string {
	v := get(c)
	return "text-" + v.Light + " dark:text-" + v.Dark
}

func Combine(parts ...string) string {
	out := ""
	for _, p := range parts {
		if p != "" {
			out += p + " "
		}
	}
	return out
}
