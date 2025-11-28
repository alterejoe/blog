package theme

var SHADOWS = map[string]ColorSet{
	"none": {
		Light:      "shadow-none",
		LightHover: "",
		Dark:       "dark:shadow-none",
		DarkHover:  "",
	},
	"sm": {
		Light:      "shadow-sm",
		LightHover: "",
		Dark:       "dark:shadow-sm",
		DarkHover:  "",
	},
	"md": {
		Light:      "shadow",
		LightHover: "",
		Dark:       "dark:shadow-md",
		DarkHover:  "",
	},
	"lg": {
		Light:      "shadow-lg",
		LightHover: "",
		Dark:       "dark:shadow-md",
		DarkHover:  "",
	},
	"xl": {
		Light:      "shadow-xl",
		LightHover: "",
		Dark:       "dark:shadow-lg",
		DarkHover:  "",
	},
	"inner": {
		Light:      "shadow-inner",
		LightHover: "",
		Dark:       "dark:shadow-inner",
		DarkHover:  "",
	},
}

const DEFAULT_SHADOW = "md"

func ShadowClass(s string) string {
	v, ok := SHADOWS[s]
	if !ok {
		v = SHADOWS[DEFAULT_SHADOW]
	}
	return v.Light + " " + v.Dark
}
