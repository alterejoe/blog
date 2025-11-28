package classes

import "zotes/shared/theme"

func Button(color, size, radius, border, variant, shadow string) string {
	return theme.ColorClass(color) + " " +
		theme.SizeClass(size) + " " +
		theme.RadiusClass(radius) + " " +
		theme.BorderClass(border) + " " +
		theme.VariantClass(variant) + " " +
		theme.ShadowClass(shadow)
}
