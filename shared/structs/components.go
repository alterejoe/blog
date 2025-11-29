package structs

type Radio struct {
	Checked bool
	Hx
	Common
}

type Button struct {
	Common
	Hx
	Disabled bool
}

type Checkbox struct {
	Common
	Hx
	Checked bool
}
