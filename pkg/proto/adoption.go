package proto

import (
	"context"
	"fmt"

	"github.com/lunarhue/libs-go/log"
	"github.com/lunarhue/metallic-flock/pkg/k3s"
	pb "github.com/lunarhue/metallic-flock/pkg/proto/adoption/v1"
)

type Server struct {
	pb.UnimplementedFlockServiceServer
}

func (s *Server) Adopt(ctx context.Context, req *pb.AdoptRequest) (*pb.AdoptResponse, error) {
	log.Infof("Received ADOPT command. Role: %s, Controller: %s", req.Role, req.ControllerIp)
	log.Infof("Cluster Token: %s", req.ClusterToken)

	k3s.StartAgent(fmt.Sprintf("https://%s:6443", req.ControllerIp), req.ClusterToken)

	return &pb.AdoptResponse{Success: true, Message: "Adoption started"}, nil
}

func (s *Server) Heartbeat(ctx context.Context, req *pb.HeartbeatRequest) (*pb.HeartbeatResponse, error) {
	return &pb.HeartbeatResponse{Reconfigure: false}, nil
}
