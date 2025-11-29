package html

// slide-in panels
const (
	slideleft       = "w-full absolute left-0 z-50 transition-transform duration-300 ease-in-out"
	slideleftclosed = slideleft + " -translate-x-full"
	slideleftopen   = slideleft + " translate-x-0"

	slideright       = "w-full absolute right-0 z-50 transition-transform duration-300 ease-in-out"
	sliderightclosed = slideright + " translate-x-full"
	sliderightopen   = slideright + " translate-x-0"
)
