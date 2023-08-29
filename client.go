package vigem

type Client interface {
	NativeHandle() ClientType
}

type ClientImpl struct {
	nativeHandle ClientType
}

func NewClient() *ClientImpl {
	nativeHandle := Alloc()

	if nativeHandle == 0 {
		panic("VigemAllocFailedException")
	}

	_ = Connect(nativeHandle)
	// TODO: Handle error

	return &ClientImpl{
		nativeHandle: nativeHandle,
	}
}

func (c *ClientImpl) Release() {
	Disconnect(c.nativeHandle)
	Free(c.nativeHandle)
}

func (c *ClientImpl) NativeHandle() ClientType {
	return c.nativeHandle
}
