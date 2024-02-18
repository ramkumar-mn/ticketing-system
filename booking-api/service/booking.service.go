package service

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/ramkumn/ticketing-system/booking"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BookingService struct for booking service
// manages the booking operations
type BookingService struct {
	seatAllocator *SeatAllocator
	data          []*pb.Booking
}

// NewBookingService creates instance of BookingService
func NewBookingService() *BookingService {
	return &BookingService{
		seatAllocator: NewSeatAllocator(),
	}
}

// CreateBooking creates a new booking for the given user
func (b *BookingService) CreateBooking(ctx context.Context, req *pb.BookingRequest) (*pb.BookingResponse, error) {

	// validate request arguments
	if req.GetFrom() == "" || req.GetTo() == "" || req.GetPrice() != 0 || req.GetUser() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid booking request")
	}

	// Allocate seat and section
	seat, section, err := b.seatAllocator.AllocateSeat()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to allocate seat: %v", err)
	}

	booking := &pb.Booking{
		Id:      uuid.NewString(),
		From:    req.GetFrom(),
		To:      req.GetTo(),
		Price:   req.GetPrice(),
		Seat:    seat,
		Section: section,
		User: &pb.User{
			FirstName: req.GetUser().GetFirstName(),
			LastName:  req.GetUser().GetLastName(),
			Email:     req.GetUser().GetEmail(),
		},
	}

	// add booking to in memory datastore
	b.data = append(b.data, booking)
	return transformAsBookingResponse(booking), nil

}

// GetBookingByUser returns the booking for the given user
func (b *BookingService) GetBookingByUser(ctx context.Context, req *pb.GetBookingByUserRequest) (*pb.BookingResponse, error) {
	// find booking entry from data for the given user
	booking, _ := b.getBookingFromDataByUser(ctx, req.GetUser())
	if booking == nil {
		return nil, status.Errorf(codes.NotFound, "booking not found for user %s", req.GetUser().GetEmail())
	}
	return transformAsBookingResponse(booking), nil

}

// GetBookingsBySection returns the list of bookings for the given section
func (b *BookingService) GetBookingsBySection(ctx context.Context, req *pb.GetBookingsBySectionRequest) (*pb.BookingListResponse, error) {
	// find list of bookings from data store for the given section
	var bookings []*pb.BookingResponse
	for _, booking := range b.data {
		if booking.GetSection() == req.GetSection() {
			bookings = append(bookings, transformAsBookingResponse(booking))
		}
	}
	return &pb.BookingListResponse{Bookings: bookings}, nil
}

// RemoveBookingByUser removes the booking and seat allocation for the given user
func (b *BookingService) RemoveBookingByUser(ctx context.Context, req *pb.RemoveBookingByUserRequest) (*pb.RemoveBookingResponse, error) {
	// find booking from data store for the given user
	booking, index := b.getBookingFromDataByUser(ctx, req.GetUser())
	if booking == nil {
		return nil, status.Errorf(codes.NotFound, "booking not found for user %s", req.GetUser().GetEmail())
	}

	// Remove seat allocation for the given user
	b.seatAllocator.DeallocateSeat(b.data[index].GetSeat(), b.data[index].GetSection())

	// remove booking entry from data store
	b.data = append(b.data[:index], b.data[index+1:]...)
	return &pb.RemoveBookingResponse{}, nil
}

// ModifySeatByUser modifies the seat for the given user
func (b *BookingService) ModifySeatByUser(ctx context.Context, req *pb.SeatModificationRequest) (*pb.SeatModificationResponse, error) {

	// find booking from data store for the given user
	booking, _ := b.getBookingFromDataByUser(ctx, req.GetUser())
	if booking == nil {
		return nil, status.Errorf(codes.NotFound, "booking not found for user %s", req.GetUser().GetEmail())
	}

	// validate request arguments
	if booking.GetSection() == req.GetSection() && booking.GetSeat() == req.GetSeat() {
		return nil, status.Errorf(codes.InvalidArgument, "new seat cannot be same as old seat")
	}

	// find by section and seat from data store and return error if seat not available
	allocationErr := b.seatAllocator.AllocateSpecificSeat(req.GetSeat(), req.GetSection())
	if allocationErr != nil {
		return nil, status.Errorf(codes.InvalidArgument, "seat %d in section %s is not available", req.GetSeat(), req.GetSection())
	}

	// update the seat and section for the given user
	booking.Seat = req.GetSeat()
	booking.Section = req.GetSection()
	return &pb.SeatModificationResponse{Seat: booking.GetSeat(), Section: booking.GetSection(), User: booking.GetUser()}, nil

}

// getBookingFromDataByUser returns the booking and index from datastore for the given user (internal helper function)
func (b *BookingService) getBookingFromDataByUser(ctx context.Context, user *pb.User) (*pb.Booking, int) {
	for index, booking := range b.data {
		if booking.GetUser().GetEmail() == user.GetEmail() {
			return booking, index
		}
	}
	return nil, -1
}

// transformAsBookingResponse transforms the booking entry to booking response (internal helper function)
func transformAsBookingResponse(booking *pb.Booking) *pb.BookingResponse {
	return &pb.BookingResponse{
		Id:      booking.GetId(),
		From:    booking.GetFrom(),
		To:      booking.GetTo(),
		Price:   booking.GetPrice(),
		Seat:    booking.GetSeat(),
		Section: booking.GetSection(),
		User:    booking.GetUser(),
	}
}
