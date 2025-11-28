package classes

import "zotes/shared/theme"

func Checkbox(color, size, radius, border, shadow string) string {
	return theme.CheckboxColorClass(color) + " " +
		theme.CheckboxSizeClass(size) + " " +
		theme.RadiusClass(radius) + " " +
		theme.CheckboxBorderClass(border) + " " +
		theme.ShadowClass(shadow)
}
