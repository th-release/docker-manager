package web

type RenameContainerDto struct {
	Name string `json:"name" validate:"required"`
}
