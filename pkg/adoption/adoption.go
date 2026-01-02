package adoption

import (
	"context"
	"fmt"
	"time"

	"github.com/lunarhue/libs-go/log"

	"github.com/lunarhue/metallic-flock/pkg/k3s"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/lunarhue/metallic-flock/pkg/proto/adoption/v1"
)

func AdoptNode(listenPort int, controllerIp, computeIp string, role string) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", computeIp, listenPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("Failed to connect to %s: %v", computeIp, err)
		return
	}

	defer conn.Close()

	adoptionToken, err := k3s.CreateJoinToken("metallic-flock-node", 1*time.Minute)
	if err != nil {
		log.Errorf("Failed to create join token: %v", err)
		return
	}

	log.Infof("Generated join token for %s: %s", computeIp, adoptionToken)

	client := pb.NewFlockServiceClient(conn)
	_, err = client.Adopt(context.Background(), &pb.AdoptRequest{
		ClusterToken: adoptionToken,
		ControllerIp: controllerIp,
		Role:         role,
	})

	if err != nil {
		log.Errorf("Failed to adopt %s: %v", computeIp, err)
	} else {
		log.Infof("Successfully sent adoption command to %s", computeIp)
	}
}
