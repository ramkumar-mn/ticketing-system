package service

import (
	"errors"
	"math/rand"
)

const MaxSeatsPerSection = 100

var MaxSeatsLimitReached = errors.New("max seats limit reached")
var SeatNotAvailable = errors.New("seat is not available")

// SeatAllocator struct for seat allocation
// manages the seat allocation operations
type SeatAllocator struct {
	occupiedSeatsSectionA map[int32]bool
	occupiedSeatsSectionB map[int32]bool
	sections              [2]string
}

// NewSeatAllocator creates instance of SeatAllocator
func NewSeatAllocator() *SeatAllocator {
	return &SeatAllocator{
		sections:              [2]string{"A", "B"},
		occupiedSeatsSectionA: make(map[int32]bool),
		occupiedSeatsSectionB: make(map[int32]bool),
	}
}

// AllocateSeat allocates a seat for the user
// Find a section with available seat and allocate the seat
// Uses random number generator to allocate the seat and section is picked randomly
// Returns the seat number and section
func (s *SeatAllocator) AllocateSeat() (int32, string, error) {

	// find a section with available seat
	section, err := s.findSection()
	if err != nil {
		return 0, "", err
	}

	// generate a seat number within the boundary of max seats per section
	seat := rand.Int31n(MaxSeatsPerSection)
	if section == "A" {
		// if seat is already occupied, find another seat
		if _, ok := s.occupiedSeatsSectionA[seat]; ok {
			return s.AllocateSeat()
		}
		// if seat is available, allocate the seat
		s.occupiedSeatsSectionA[seat] = true
	} else {
		// if seat is already occupied, find another seat
		if _, ok := s.occupiedSeatsSectionB[seat]; ok {
			return s.AllocateSeat()
		}
		// if seat is available, allocate the seat
		s.occupiedSeatsSectionB[seat] = true
	}
	return seat, section, nil
}

// AllocateSpecificSeat allocates a specific seat for the user
// Returns error if the seat is not available
func (s *SeatAllocator) AllocateSpecificSeat(seat int32, section string) error {
	if !s.isSeatAvailable(seat, section) {
		return SeatNotAvailable
	}
	if section == "A" {
		s.occupiedSeatsSectionA[seat] = true
	} else {
		s.occupiedSeatsSectionB[seat] = true
	}
	return nil
}

// isSeatAvailable checks if the seat is available in the given section
// Returns true if the seat is available, false otherwise in case of seat limit reached or already occupied
func (s *SeatAllocator) isSeatAvailable(seat int32, section string) bool {
	if section == "A" && len(s.occupiedSeatsSectionA) < MaxSeatsPerSection {
		if _, ok := s.occupiedSeatsSectionA[seat]; ok {
			return false
		} else {
			return true
		}
	} else if section == "B" && len(s.occupiedSeatsSectionB) < MaxSeatsPerSection {
		if _, ok := s.occupiedSeatsSectionB[seat]; ok {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

// DeallocateSeat deallocates the seat for the user
// removes the seat from the occupied seats list
func (s *SeatAllocator) DeallocateSeat(seat int32, section string) {
	if section == "A" {
		delete(s.occupiedSeatsSectionA, seat)
	} else {
		delete(s.occupiedSeatsSectionB, seat)
	}
}

// findSection finds a section with available seat
// Returns the section if available, error otherwise in case of seat limit reached
func (s *SeatAllocator) findSection() (string, error) {
	if len(s.occupiedSeatsSectionA) < MaxSeatsPerSection && len(s.occupiedSeatsSectionB) < MaxSeatsPerSection {
		return s.sections[rand.Intn(2)], nil
	} else if len(s.occupiedSeatsSectionA) < MaxSeatsPerSection {
		return "A", nil
	} else if len(s.occupiedSeatsSectionB) < MaxSeatsPerSection {
		return "B", nil
	} else {
		return "", MaxSeatsLimitReached
	}
}
