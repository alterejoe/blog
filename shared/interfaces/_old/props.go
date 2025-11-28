package interfaces

type CommonProps interface {
	ID() string
	Name() string
}

type ValueProp interface {
	Value() string
}

type CommonPropsFlex interface {
	CommonProps
	Flex() bool
	SideText() string
}

type LinkProps interface {
	Href() string
}

type HxProps interface {
	HXTarget() string
	HXInclude() string
	HXTrigger() string
	HXSwap() string
	HXIndicator() string
	HXVals() string
}
type RadioProps interface {
	CommonProps
	HxGetProps
	Color() string
	Checked() bool
}

type HxGetProps interface {
	CommonProps
	HxProps
	Get() string
}

type HxDeleteProps interface {
	CommonProps
	HxProps
	Delete() string
}

type HxPostProps interface {
	CommonProps
	HxProps
	Post() string
}
type HxConfirmProps interface {
	Confirm() string
}

type FileDownloadProps interface {
	CommonProps
	LinkProps
	Download() string
}

type FileUploadProps interface {
	CommonProps
	HxPostProps
}

type CheckboxProps interface {
	CommonProps
	Checked() bool
	Color() string
}
