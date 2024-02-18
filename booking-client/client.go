package main

import (
	"context"
	pb "github.com/ramkumn/ticketing-system/booking"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	c := pb.NewBookingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create Booking for user1
	bookingRequestUser1 := &pb.BookingRequest{
		From:  "London",
		To:    "France",
		Price: 20,
		User: &pb.User{
			FirstName: "Phil",
			LastName:  "Dunphy",
			Email:     "pd@test.com",
		},
	}
	bookingUser1, bookingErr := c.CreateBooking(ctx, bookingRequestUser1)
	if bookingErr != nil {
		log.Fatalf("error calling function SayHello: %v", bookingErr)
	}

	log.Printf("Response from CreateBooking function for user=Phil: %s", bookingUser1)

	// Create Booking for user2
	bookingRequestUser2 := &pb.BookingRequest{
		From:  "London",
		To:    "France",
		Price: 20,
		User: &pb.User{
			FirstName: "Alex",
			LastName:  "Dunphy",
			Email:     "ad@test.com",
		},
	}
	bookingUser2, bookingErr := c.CreateBooking(ctx, bookingRequestUser2)
	if bookingErr != nil {
		log.Fatalf("error calling function SayHello: %v", bookingErr)
	}

	log.Printf("Response from CreateBooking function for user=Alex: %s", bookingUser2)

	// Create Booking for user3
	bookingRequestUser3 := &pb.BookingRequest{
		From:  "London",
		To:    "France",
		Price: 20,
		User: &pb.User{
			FirstName: "Luke",
			LastName:  "Dunphy",
			Email:     "ld@test.com",
		},
	}
	bookingUser3, bookingErr := c.CreateBooking(ctx, bookingRequestUser3)
	if bookingErr != nil {
		log.Fatalf("error calling function SayHello: %v", bookingErr)
	}

	log.Printf("Response from CreateBooking function for user=Luke: %s", bookingUser3)

	// Get Booking for user3
	getBookingByUserRequest := &pb.GetBookingByUserRequest{
		User: &pb.User{
			FirstName: "Luke",
			LastName:  "Dunphy",
			Email:     "ld@test.com",
		},
	}

	getBookingByUser, getBookingErr := c.GetBookingByUser(ctx, getBookingByUserRequest)
	if getBookingErr != nil {
		log.Fatalf("error calling function GetBookingByUser: %v", getBookingErr)
	}
	log.Printf("Response from GetBookingByUser function for user=Luke: %s", getBookingByUser)

	// List bookings for section B
	getBookingsBySectionRequest := &pb.GetBookingsBySectionRequest{
		Section: "B",
	}
	bookingsBySection, bookingsBySectionErr := c.GetBookingsBySection(ctx, getBookingsBySectionRequest)
	if bookingsBySectionErr != nil {
		log.Fatalf("error calling function GetBookingsBySection: %v", bookingsBySectionErr)
	}
	log.Printf("Response from GetBookingsBySection function for section B: %s", bookingsBySection)

	// Modify seat for user2
	modifySeatRequest := &pb.SeatModificationRequest{
		Section: bookingUser2.Section,
		Seat:    43,
		User:    bookingUser2.GetUser(),
	}
	modifiedSeat, modifySeatErr := c.ModifySeatByUser(ctx, modifySeatRequest)
	if modifySeatErr != nil {
		log.Fatalf("error calling function ModifySeatByUser: %v", modifySeatErr)
	}
	log.Printf("Response from ModifySeatByUser function for user=Alex: %s", modifiedSeat)

	// Remove user1 booking
	removeBookingRequest := &pb.RemoveBookingByUserRequest{User: bookingUser1.GetUser()}
	_, removeBookingErr := c.RemoveBookingByUser(ctx, removeBookingRequest)
	if removeBookingErr != nil {
		log.Fatalf("error calling function RemoveBookingByUser: %v", removeBookingErr)
	}
	log.Printf("User booking removed for user=Phil")

	// List bookings for section A
	getBookingsBySectionRequest = &pb.GetBookingsBySectionRequest{
		Section: "A",
	}
	bookingsBySection, bookingsBySectionErr = c.GetBookingsBySection(ctx, getBookingsBySectionRequest)
	if bookingsBySectionErr != nil {
		log.Fatalf("error calling function GetBookingsBySection: %v", bookingsBySectionErr)
	}
	log.Printf("Response from GetBookingsBySection function for section A: %s", bookingsBySection)

	// List bookings for section B
	getBookingsBySectionRequest = &pb.GetBookingsBySectionRequest{
		Section: "B",
	}
	bookingsBySection, bookingsBySectionErr = c.GetBookingsBySection(ctx, getBookingsBySectionRequest)
	if bookingsBySectionErr != nil {
		log.Fatalf("error calling function GetBookingsBySection: %v", bookingsBySectionErr)
	}
	log.Printf("Response from GetBookingsBySection function for section B: %s", bookingsBySection)

}
