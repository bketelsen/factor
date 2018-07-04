package model

import (
	"log"
	"text/template"
)

var modelTypesTpl *template.Template
var modelServerTpl *template.Template
var modelClientTpl *template.Template
var serversTpl *template.Template
var clientsTpl *template.Template

func init() {
	var err error
	modelTypesTpl, err = template.New("modelTypes").Parse(`package models

// this file is generated by 'factor dev' but won't be overwritten. feel free
// to make any modifications you want in here. If you do, make sure to
// check {{.UpperName}}Server and {{.UpperName}}Client because you might need to
// update them too so the thing still builds.

import (
	"context"
	
	"github.com/satori/go.uuid"
)

type Create{{.UpperName}}Req struct {
	Ctx context.Context
	Data {{.UpperName}}
}

type Create{{.UpperName}}Res struct {
	Err error
}

type Get{{.UpperName}}Req struct {
	ID uuid.UUID
}

type Get{{.UpperName}}Res struct {
	Data {{.UpperName}}
}

type Update{{.UpperName}}Req struct {
	ID uuid.UUID
	New {{.UpperName}}
}

type Update{{.UpperName}}Res struct {
	Err error
}

type Delete{{.UpperName}}Req struct {
	ID uuid.UUID
}

type Delete{{.UpperName}}Res struct {
	Err error
}
`)
	if err != nil {
		log.Fatalf("parsing types template: %s", err)
	}

	modelServerTpl, err = template.New("modelServer").Parse(`package models

// this file is generated by 'factor dev' but won't be overwritten. feel free to
// fill in your server implementation here!

import (
	"net/rpc"
)

var {{.LowerName}}Server = New{{.UpperName}}Server()

type {{.UpperName}}Server struct{}

func New{{.UpperName}}Server() {{.UpperName}}Server {
	return &{{.UpperName}}Server{}
}

func (srv *{{.UpperName}}Server) Create(req Create{{.UpperName}}Req, res *Create{{.UpperName}}Res) error {
	// TODO
	return nil
}

func (srv *{{.UpperName}}Server) Get(req Get{{.UpperName}}Req, res *Get{{.UpperName}}Res) error {
	// TODO
	return nil
}

func (srv *{{.UpperName}}Server) Update(req Update{{.UpperName}}Req, res *Update{{.UpperName}}Res) error {
	// TODO
	return nil
}

func (srv *{{.UpperName}}Server) Delete(req Delete{{.UpperName}}Req, res *Delete{{.UpperName}}Res) error {
	// TODO
	return nil
}

func init() {
	rpc.RegisterName("{{.UpperName}}", {{.LowerName}}Server)
}
`)
	if err != nil {
		log.Fatalf("parsing server template: %s", err)
	}

	modelClientTpl, err = template.New("modelClient").Parse(`package models

// this file is generated by 'factor dev' but won't be overwritten. feel free to
// fill in your own client logic here!

import (
	"net/rpc"

	"github.com/satori/go.uuid"
)

type {{.UpperName}}Client struct{
	RPC *rpc.Client
}

func (cl *{{.UpperName}}Client) Get(id uuid.UUID) (*{{.UpperName}}, error) {
	req := Get{{.UpperName}}Req{ID: id}
	var res Get{{.UpperName}}Res
	if err := cl.RPC.Call("{{.UpperName}}.Get", req, &res); err != nil {
		return nil, err
	}
	return &res.Data, nil
}

// TODO: more
`)
	if err != nil {
		log.Fatalf("parsing client template: %s", err)
	}

	serversTpl, err = template.New("servers").Parse(`package models

import (
	"net/rpc"
	"net"
	"fmt"
	"net/http"
)

func StartRPCServer(port int) error {
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return http.Serve(listener, nil)
}
`)
	if err != nil {
		log.Fatalf("parsing servers file: %s", err)
	}
	clientsTpl, err = template.New("clients").Parse(`package models

import (
	"net/rpc"
	"log"
)

var rpcCl *rpc.Client

func init() {
	client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
	if err != nil {
		log.Fatalf("couldn't connect to rpc! (%s)", err)
	}
	rpcCl = client
}

type Client struct {
	{{range $cl := .Clients}}
	{{$cl}}
	{{end}}
}

var RPC *Client = &Client{
	{{range $cl := .Clients}}
	{{$cl}}{RPC: rpcCl},
	{{end}}
}
`)
	if err != nil {
		log.Fatalf("parsing clients file: %s", err)
	}

}
