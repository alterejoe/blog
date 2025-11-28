package theme

var CHECKBOX_SIZES = map[string]string{
	"sm": "w-4 h-4 [&>svg]:w-3 [&>svg]:h-3",
	"md": "w-5 h-5 [&>svg]:w-4 [&>svg]:h-4",
	"lg": "w-6 h-6 [&>svg]:w-5 [&>svg]:h-5",
}

func CheckboxSizeClass(s string) string {
	if v, ok := CHECKBOX_SIZES[s]; ok {
		return v
	}
	return CHECKBOX_SIZES["md"]
}
