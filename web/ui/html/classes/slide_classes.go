package classes

// Slide-in panels
const (
	SlideLeft       = " w-full absolute left-0 z-50 transition-transform duration-500 ease-in-out flex flex-row"
	SlideLeftClosed = SlideLeft + " -translate-x-full"
	SlideLeftOpen   = SlideLeft + " translate-x-0"

	SlideRight       = " w-full absolute right-0 z-50 transition-transform duration-500 ease-in-out flex flex-row"
	SlideRightClosed = SlideRight + " translate-x-full"
	SlideRightOpen   = SlideRight + " translate-x-0"
)
