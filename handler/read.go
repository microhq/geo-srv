package handler

import (
	"log"

	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"
	"github.com/microhq/geo-srv/dao"
	loc "github.com/microhq/geo-srv/proto/location"

	"golang.org/x/net/context"
)

type Location struct{}

func (l *Location) Read(ctx context.Context, req *loc.ReadRequest, rsp *loc.ReadResponse) error {
	log.Print("Received Location.Read request")

	id := req.Id

	if len(id) == 0 {
		return errors.BadRequest(server.DefaultOptions().Name+".read", "Require Id")
	}

	entity, err := dao.Read(id)
	if err != nil {
		return err
	}

	rsp.Entity = entity.ToProto()

	return nil
}
