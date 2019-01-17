package factory_method

type Button interface {
	Show() string
}

type WinButton struct{}

func (w *WinButton)Show() string {
	return "WinButton"
}

type MacButton struct{}

func (m *MacButton)Show() string {
	return "MacButton"
}

type ButtonFactory interface {
	CreateButton() Button
}

type WinButtonFactory struct {}

func (wf *WinButtonFactory)CreateButton() Button {
	return &WinButton{}
}

type MacButtonFactory struct {}

func (mf *MacButtonFactory)CreateButton() Button{
	return &MacButton{}
}