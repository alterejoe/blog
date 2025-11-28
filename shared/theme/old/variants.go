package theme

type VariantToken struct {
	Light string
	Dark  string
}

var VARIANTS = map[string]VariantToken{
	"solid": {
		Light: "text-white",
		Dark:  "dark:text-white",
	},
	"outline": {
		Light: "border border-current bg-transparent",
		Dark:  "dark:border-neutral-400 dark:bg-transparent",
	},
	"ghost": {
		Light: "bg-transparent hover:bg-slate-200",
		Dark:  "dark:bg-transparent dark:hover:bg-neutral-700",
	},
	"quiet": {
		Light: "opacity-70 hover:opacity-100",
		Dark:  "dark:opacity-80 dark:hover:opacity-100",
	},
}

const DEFAULT_VARIANT = "solid"

func VariantClass(v string) string {
	t, ok := VARIANTS[v]
	if !ok {
		t = VARIANTS[DEFAULT_VARIANT]
	}
	return t.Light + " " + t.Dark
}
