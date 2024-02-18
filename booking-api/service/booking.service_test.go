package service

import (
	"context"
	pb "github.com/ramkumn/ticketing-system/booking"
	"testing"
)

func TestBookingService_CreateBookingValidScenario(t *testing.T) {

	req := &pb.BookingRequest{
		From:  "London",
		To:    "France",
		Price: 20,
		User: &pb.User{
			FirstName: "Phil",
			LastName:  "Dunphy",
			Email:     "pd@test.com",
		},
	}

	want := &pb.BookingResponse{
		From:  "London",
		To:    "France",
		Price: 20,
		User: &pb.User{
			FirstName: "Phil",
			LastName:  "Dunphy",
			Email:     "pd@test.com",
		},
	}

	bookingService := NewBookingService()
	got, err := bookingService.CreateBooking(context.TODO(), req)
	if err != nil {
		t.Errorf("CreateBooking() error = %v want nil", err)
		return
	}
	if got.From != want.From || got.To != want.To ||
		got.Price != want.Price || got.User.FirstName != want.User.FirstName ||
		got.User.LastName != want.User.LastName || got.User.Email != want.User.Email {
		t.Errorf("CreateBooking() = %v, want %v", got, want)
	}
}

func TestBookingService_CreateBookingInValidScenario(t *testing.T) {

	req := &pb.BookingRequest{
		From:  "London",
		Price: 20,
		User: &pb.User{
			FirstName: "Phil",
			LastName:  "Dunphy",
			Email:     "pd@test.com",
		},
	}

	bookingService := NewBookingService()
	_, err := bookingService.CreateBooking(context.TODO(), req)
	if err == nil {
		t.Errorf("CreateBooking() err is nil, want non nil")
	}
}

func TestBookingService_GetBookingByUser(t *testing.T) {
	req := &pb.BookingRequest{
		From:  "London",
		To:    "France",
		Price: 20,
		User: &pb.User{
			FirstName: "Phil",
			LastName:  "Dunphy",
			Email:     "pd@test.com",
		},
	}

	bookingService := NewBookingService()
	_, err := bookingService.CreateBooking(context.TODO(), req)
	if err != nil {
		t.Errorf("CreateBooking() error = %v want nil", err)
		return
	}

	reqGetBookingByUser := &pb.GetBookingByUserRequest{User: req.GetUser()}

	want := &pb.BookingResponse{
		From:  "London",
		To:    "France",
		Price: 20,
		User: &pb.User{
			FirstName: "Phil",
			LastName:  "Dunphy",
			Email:     "pd@test.com",
		},
	}

	got, err := bookingService.GetBookingByUser(context.TODO(), reqGetBookingByUser)
	if err != nil {
		t.Errorf("GetBookingByUser() error = %v want nil", err)
		return
	}
	if got.From != want.From || got.To != want.To ||
		got.Price != want.Price || got.User.FirstName != want.User.FirstName ||
		got.User.LastName != want.User.LastName || got.User.Email != want.User.Email {
		t.Errorf("GetBookingByUser() = %v, want %v", got, want)
	}

}

func TestBookingService_GetBookingsBySection(t *testing.T) {
	req := &pb.BookingRequest{
		From:  "London",
		To:    "France",
		Price: 20,
		User: &pb.User{
			FirstName: "Phil",
			LastName:  "Dunphy",
			Email:     "pd@test.com",
		},
	}

	bookingService := NewBookingService()
	booking, err := bookingService.CreateBooking(context.TODO(), req)
	if err != nil {
		t.Errorf("CreateBooking() error = %v want nil", err)
		return
	}

	reqGetBookingsBySection := &pb.GetBookingsBySectionRequest{Section: booking.GetSection()}
	want := &pb.BookingListResponse{
		Bookings: []*pb.BookingResponse{booking},
	}
	got, err := bookingService.GetBookingsBySection(context.TODO(), reqGetBookingsBySection)
	if err != nil {
		t.Errorf("GetBookingsBySection() error = %v want nil", err)
		return
	}
	if len(got.Bookings) != len(want.Bookings) {
		t.Errorf("GetBookingsBySection() = %v, want %v", got, want)
	}

}

func TestBookingService_RemoveBookingByUser(t *testing.T) {
	req := &pb.BookingRequest{
		From:  "London",
		To:    "France",
		Price: 20,
		User: &pb.User{
			FirstName: "Phil",
			LastName:  "Dunphy",
			Email:     "pd@test.com",
		},
	}

	bookingService := NewBookingService()
	_, err := bookingService.CreateBooking(context.TODO(), req)
	if err != nil {
		t.Errorf("CreateBooking() error = %v want nil", err)
		return
	}

	reqRemoveBookingByUser := &pb.RemoveBookingByUserRequest{User: req.GetUser()}
	_, removeErr := bookingService.RemoveBookingByUser(context.TODO(), reqRemoveBookingByUser)
	if removeErr != nil {
		t.Errorf("RemoveBookingByUser() error = %v want nil", removeErr)
		return
	}

}

func TestBookingService_ModifySeatByUser(t *testing.T) {
	req := &pb.BookingRequest{
		From:  "London",
		To:    "France",
		Price: 20,
		User: &pb.User{
			FirstName: "Phil",
			LastName:  "Dunphy",
			Email:     "pd@test.com",
		},
	}

	bookingService := NewBookingService()
	booking, err := bookingService.CreateBooking(context.TODO(), req)
	if err != nil {
		t.Errorf("CreateBooking() error = %v want nil", err)
		return
	}

	reqModifySeatByUser := &pb.SeatModificationRequest{User: req.GetUser(), Seat: 43, Section: booking.GetSection()}
	want := &pb.SeatModificationResponse{
		Seat:    43,
		Section: booking.GetSection(),
		User:    booking.GetUser(),
	}

	got, modifyErr := bookingService.ModifySeatByUser(context.TODO(), reqModifySeatByUser)
	if modifyErr != nil {
		t.Errorf("ModifySeatByUser() error = %v want nil", modifyErr)
		return
	}
	if got.Seat != want.Seat || got.Section != want.Section || got.User.FirstName != want.User.FirstName ||
		got.User.LastName != want.User.LastName || got.User.Email != want.User.Email {
		t.Errorf("ModifySeatByUser() = %v, want %v", got, want)
	}

}
