package theme

// -----------------------------------------------------------------------------
// CORE COLOR TOKENS
// -----------------------------------------------------------------------------

type ColorSet struct {
	Light      string
	LightHover string
	Dark       string
	DarkHover  string
}

// MASTER COLOR TOKENS — your theme lives HERE and only here.
var COLORS = map[string]ColorSet{
	"primary": {
		Light:      "primary",
		LightHover: "primary-hover",
		Dark:       "primary-dark",
		DarkHover:  "primary-dark-hover",
	},
	"secondary": {
		Light:      "secondary",
		LightHover: "secondary-hover",
		Dark:       "secondary-dark",
		DarkHover:  "secondary-dark-hover",
	},
	"tertiary": {
		Light:      "tertiary",
		LightHover: "tertiary-hover",
		Dark:       "tertiary-dark",
		DarkHover:  "tertiary-dark-hover",
	},
	"accent": {
		Light:      "accent",
		LightHover: "accent-hover",
		Dark:       "accent-dark",
		DarkHover:  "accent-dark-hover",
	},
	"neutral": {
		Light:      "neutral",
		LightHover: "neutral-hover",
		Dark:       "neutral-dark",
		DarkHover:  "neutral-dark-hover",
	},
}

const DEFAULT_COLOR = "neutral"

func getColor(c string) ColorSet {
	v, ok := COLORS[c]
	if !ok {
		return COLORS[DEFAULT_COLOR]
	}
	return v
}

// -----------------------------------------------------------------------------
// CORE COLOR HELPERS
// -----------------------------------------------------------------------------

func Bg(c string) string {
	v := getColor(c)
	return Combine(
		"bg-"+v.Light,
		"hover:bg-"+v.LightHover,
		"peer-checked:bg-"+v.Light,
		"dark:bg-"+v.Dark,
		"dark:hover:bg-"+v.DarkHover,
		"dark:peer-checked:bg-"+v.Dark,
	)
}

func Hover(c string) string {
	v := getColor(c)
	return Combine(
		"hover:bg-"+v.LightHover,
		"dark:hover:bg-"+v.DarkHover,
	)
}

func RingColor(c string) string {
	v := getColor(c)
	return Combine(
		"ring-"+v.Light,
		"dark:ring-"+v.Dark,
	)
}

// Backward-compatible with your old name.
func Ring(c string) string {
	return RingColor(c)
}

func BorderColor(c string) string {
	v := getColor(c)
	return Combine(
		"border-"+v.Light,
		"dark:border-"+v.Dark,
	)
}

// Backward-compatible with your old name (used as "border color").
func Border(c string) string {
	return BorderColor(c)
}

func Checked(c string) string {
	v := getColor(c)
	return Combine(
		"peer-checked:bg-"+v.Light,
		"peer-checked:ring-"+v.Light,
		"dark:peer-checked:bg-"+v.Dark,
		"dark:peer-checked:ring-"+v.Dark,
	)
}

func Text(c string) string {
	v := getColor(c)
	return Combine(
		"text-"+v.Light,
		"dark:text-"+v.Dark,
	)
}

// -----------------------------------------------------------------------------
// BORDER WIDTHS (SIZE) + BORDER CLASS
// -----------------------------------------------------------------------------

// Pure size-only border widths. Color is handled by BorderColor().
var BORDERS = map[string]string{
	"none": "border-0",
	"sm":   "border",
	"md":   "border-2",
	"lg":   "border-4",
}

const DEFAULT_BORDER = "md"

func BorderWidth(size string) string {
	if v, ok := BORDERS[size]; ok {
		return v
	}
	return BORDERS[DEFAULT_BORDER]
}

// Full border: width + theme color.
func BorderClass(size string, color string) string {
	return Combine(
		BorderWidth(size),
		BorderColor(color),
	)
}

// -----------------------------------------------------------------------------
// RING SIZE + RING CLASS
// -----------------------------------------------------------------------------

var RING_SIZE = map[string]string{
	"sm": "ring-1",
	"md": "ring-2",
	"lg": "ring-4",
}

const DEFAULT_RING = "md"

func RingSize(s string) string {
	if v, ok := RING_SIZE[s]; ok {
		return v
	}
	return RING_SIZE[DEFAULT_RING]
}

func RingClass(size string, color string) string {
	return Combine(
		RingSize(size),
		RingColor(color),
	)
}

// -----------------------------------------------------------------------------
// RADIUS
// -----------------------------------------------------------------------------

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

// -----------------------------------------------------------------------------
// SHADOWS
// -----------------------------------------------------------------------------

var SHADOWS = map[string]ColorSet{
	"none": {
		Light: "shadow-none",
		Dark:  "dark:shadow-none",
	},
	"sm": {
		Light: "shadow-sm",
		Dark:  "dark:shadow-sm",
	},
	"md": {
		Light: "shadow",
		Dark:  "dark:shadow-md",
	},
	"lg": {
		Light: "shadow-lg",
		Dark:  "dark:shadow-md",
	},
	"xl": {
		Light: "shadow-xl",
		Dark:  "dark:shadow-lg",
	},
	"inner": {
		Light: "shadow-inner",
		Dark:  "dark:shadow-inner",
	},
}

const DEFAULT_SHADOW = "md"

func ShadowClass(s string) string {
	v, ok := SHADOWS[s]
	if !ok {
		v = SHADOWS[DEFAULT_SHADOW]
	}
	return Combine(
		v.Light,
		v.Dark,
	)
}

// -----------------------------------------------------------------------------
// VARIANTS (STRUCTURAL BEHAVIOR MODES)
// -----------------------------------------------------------------------------

type VariantToken struct {
	Light      string
	LightHover string
	Dark       string
	DarkHover  string
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
		Light:      "bg-transparent",
		LightHover: "hover:bg-slate-200",
		Dark:       "dark:bg-transparent",
		DarkHover:  "dark:hover:bg-neutral-700",
	},
	"quiet": {
		Light:      "opacity-70",
		LightHover: "hover:opacity-100",
		Dark:       "dark:opacity-80",
		DarkHover:  "dark:hover:opacity-100",
	},
}

const DEFAULT_VARIANT = "solid"

func VariantClass(v string) string {
	t, ok := VARIANTS[v]
	if !ok {
		t = VARIANTS[DEFAULT_VARIANT]
	}
	return Combine(
		t.Light,
		t.LightHover,
		t.Dark,
		t.DarkHover,
	)
}

// -----------------------------------------------------------------------------
// SIZES / SPACING / TEXT
// -----------------------------------------------------------------------------

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

// -----------------------------------------------------------------------------
// COMPONENT-SPECIFIC TOKENS (CHECKBOX)
// -----------------------------------------------------------------------------

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

// Optional: size + themed ring + border for checkboxes.
func CheckboxBorder(size string, color string) string {
	return Combine(
		BorderClass(size, color),
		RingClass(size, color),
	)
}

// -----------------------------------------------------------------------------
// CLASS COMBINER
// -----------------------------------------------------------------------------

func Combine(parts ...string) string {
	out := ""
	for _, p := range parts {
		if p == "" {
			continue
		}
		if out != "" {
			out += " "
		}
		out += p
	}
	return out
}
