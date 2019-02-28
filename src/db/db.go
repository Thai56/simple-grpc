package db

import (
	"context"
	"net"

	redis "github.com/go-redis/redis"
	redisHelper "github.com/simple-grpc/src/db/redis"
	pb "github.com/simple-grpc/src/protos"
	log "github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

//Database ...
type Database struct {
	listener net.Listener
	server   *grpc.Server
	dbConn   *redis.Client
}

//Serve ...
func (d *Database) Serve(ctx context.Context) error {
	// create the connection that we need for the reddis db
	var err error

	// create a will create a redis instance that we will need to connect to
	d.dbConn = redisHelper.NewRedisClient()

	pong, err := d.dbConn.Ping().Result()

	if err != nil {
		log.Fatalf("Failed to connect to the redis db: %v", err)
		return err
	}

	log.Infof("Created the redis db: %v", pong)

	// create grpc listener
	d.listener, err = net.Listen("tcp", ":5050")

	// create a grpc server
	d.server = grpc.NewServer()
	pb.RegisterDatabaseServer(d.server, d)

	serveErr := make(chan error)
	func() {
		if err := d.server.Serve(d.listener); err != nil {
			serveErr <- err
		}
	}()
	// always stop
	defer func() {
		d.server.GracefulStop()
	}()

	select {
	case <-ctx.Done():
		if err := ctx.Err(); err != context.Canceled && err != context.DeadlineExceeded {
			return err
		}
	case err := <-serveErr:
		if err != nil {
			return err
		}
	}
	return nil
}

// GetCountries ...
func (d *Database) GetCountries(context context.Context, req *pb.GetCountriesRequest) (*pb.GetCountriesReply, error) {
	return &pb.GetCountriesReply{}, nil
}

// create a docker-compose file that will spin up an instance of a reddis db
