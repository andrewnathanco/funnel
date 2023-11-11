package shared

type Button struct {
	Text string
	Disabled bool
}

func NewButton(text string, disabled bool) Button {
	return Button{
		Text: text,
		Disabled: disabled,
	}
}
