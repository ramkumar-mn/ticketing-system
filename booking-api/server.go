package main

import (
	pb "github.com/ramkumn/ticketing-system/booking"
	"github.com/ramkumn/ticketing-system/booking-api/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer()

	// Register the service
	bookingService := service.NewBookingService()
	pb.RegisterBookingServiceServer(s, bookingService)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
