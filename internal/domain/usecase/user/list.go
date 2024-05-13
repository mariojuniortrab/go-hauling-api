package user_usecase

import protocol_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol"

type filter struct {
	ID string `json:"id"`
}

type ListInputDto struct {
	protocol_usecase.List
	filter
}

type list struct {
}

func NewList() *list {
	return &list{}
}

func (u *list) Execute() {

}
