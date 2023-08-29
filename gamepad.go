package vigem

type Gamepad interface {
	NativeHandle() TargetType
}

type GamepadBase struct {
	Client       Client
	IsConnected  bool
	VendorId     uint16
	ProductId    uint16
	NativeHandle TargetType
}

func NewGamepad(client Client, nativeHandle TargetType) *GamepadBase {
	return &GamepadBase{
		Client:       client,
		IsConnected:  false,
		NativeHandle: nativeHandle,
	}
}

func (g *GamepadBase) Connect() {
	if g.IsConnected {
		return
	}

	if g.VendorId > 0 && g.ProductId > 0 {
		TargetSetVid(g.NativeHandle, g.VendorId)
		TargetSetPid(g.NativeHandle, g.ProductId)
	}
	TargetAdd(g.Client.NativeHandle(), g.NativeHandle)
	// TODO: Handle error

	g.IsConnected = true
}

func (g *GamepadBase) Disconnect() {
	if !g.IsConnected {
		return
	}

	TargetRemove(g.Client.NativeHandle(), g.NativeHandle)
	// TODO: Handle error

	g.IsConnected = false
}

func (g *GamepadBase) Release() {
	TargetFree(g.NativeHandle)
}
