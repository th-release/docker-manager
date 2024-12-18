package web

type RenameContainerDto struct {
	Name string `json:"name" validate:"required"`
}

type CreateNetworkDto struct {
	Name       string `json:"name" validate:"required"`
	Driver     string `json:"driver" validate:"required,oneof=bridge overlay"`
	Scope      string `json:"scope" validate:"required,oneof=swarm local"`
	EnableIpv6 bool   `json:"enableIpv6" validate:"required"`
	Ingress    bool   `json:"ingress" validate:"required"`
}

type ConnectNetworkDto struct {
	Container string `json:"container" validate:"required"`
	Network   string `json:"network" validate:"required"`
}

type DisconnectNetworkDto struct {
	Container string `json:"container" validate:"required"`
	Network   string `json:"network" validate:"required"`
	Force     bool   `json:"force" validate:"required"`
}
