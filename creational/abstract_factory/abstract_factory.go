package abstract_factory


type Button interface {
	Press() string
}

type WinButton struct {}

func (b *WinButton)Press() string {
	return "WinButton Press"
}

type MacButton struct{}

func (b *MacButton)Press() string {
	return "MacButton Press"
}

type Window interface {
	View() string
}

type WinWindow struct{}

func (w *WinWindow)View() string {
	return "WinWindow View"
}

type MacWindow struct{}

func (w *MacWindow)View() string {
	return "MacWindow View"
}

type AbstractFactory interface {
	CreateButton() Button
	CreateWindow() Window
}


type WinFactory struct{}

func (wf *WinFactory)CreateButton() Button {
	return &WinButton{}
}

func (wf *WinFactory)CreateWindow() Window {
	return &WinWindow{}
}

type MacFactory struct{}

func (mf *MacFactory)CreateButton() Button {
	return &MacButton{}
}

func (mf *MacFactory)CreateWindow() Window {
	return &MacWindow{}
}
