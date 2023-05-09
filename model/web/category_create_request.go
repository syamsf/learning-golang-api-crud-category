package web

type CategoryCreateRequest struct {
	Name string `validate:"required,max=20,min=1" json:"name"`
}
