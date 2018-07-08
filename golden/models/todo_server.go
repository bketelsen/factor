package models

// this file is generated by 'factor dev' but won't be overwritten. feel free to
// fill in your server implementation here!

import (
	"net/rpc"
)

var todoServer = NewTodoServer()

type TodoServer struct{}

func NewTodoServer() *TodoServer {
	return &TodoServer{}
}

func (srv *TodoServer) Create(req CreateTodoReq, res *CreateTodoRes) error {
	// TODO
	return nil
}

func (srv *TodoServer) Get(req GetTodoReq, res *GetTodoRes) error {
	// TODO
	return nil
}

func (srv *TodoServer) Update(req UpdateTodoReq, res *UpdateTodoRes) error {
	// TODO
	return nil
}

func (srv *TodoServer) Delete(req DeleteTodoReq, res *DeleteTodoRes) error {
	// TODO
	return nil
}

func init() {
	rpc.RegisterName("Todo", todoServer)
}
