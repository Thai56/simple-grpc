package db

import (
	"context"
	"net"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"time"

	redisHelper "github.com/Thai56/simple-grpc/src/db/redis"
	pb "github.com/Thai56/simple-grpc/src/protos"
	redis "github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

//Database ...
type Database struct {
	listener net.Listener
	server   *grpc.Server
	DbConn   *redis.Client
}

func init() {
	// figure out what filename the current running code is in
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		println("could not get runtime info")
	}
	// make it say <path>/docker-compose.yml
	dc := path.Join(filepath.Dir(filename), "docker-compose.yml")
	// execute docker-compose up
	cmd := exec.Command("docker-compose", "-f", dc, "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// dont use Run, which waits
	err := cmd.Start()
	if err != nil {
		log.Errorf("Failed to start the docker-compose.yml:%v", err)
		return
	}

	log.Info("Successfully started the docker-compose.yml")
	<-time.After(10 * time.Second)
}

//Serve ...
func (d *Database) Serve() (*Database, error) {
	// create the connection that we need for the reddis db
	var err error

	// create a will create a redis instance that we will need to connect to
	d.DbConn = redisHelper.NewRedisClient()

	pong, err := d.DbConn.Ping().Result()

	if err != nil {
		log.Fatalf("Failed to connect to the redis db: %v", err)
		return nil, err
	}

	log.Infof("Created the redis db: %v", pong)

	// create grpc listener
	d.listener, err = net.Listen("tcp", ":5050")

	// create a grpc server
	d.server = grpc.NewServer()
	pb.RegisterDatabaseServer(d.server, d)

	return d, nil

	//serveErr := make(chan error)
	//func() {
	//if err := d.server.Serve(d.listener); err != nil {
	//serveErr <- err
	//}
	////}()
	// always stop
	//defer func() {
	//d.server.GracefulStop()
	//}()

	//select {
	//case <-ctx.Done():
	//if err := ctx.Err(); err != context.Canceled && err != context.DeadlineExceeded {
	//return err
	//}
	//case err := <-serveErr:
	//if err != nil {
	//return err
	//}
	//}
	//return nil
}

// GetCountries ...
func (d *Database) GetCountries(context context.Context, req *pb.GetCountriesRequest) (*pb.GetCountriesReply, error) {
	return &pb.GetCountriesReply{}, nil
}

// create a docker-compose file that will spin up an instance of a reddis db
