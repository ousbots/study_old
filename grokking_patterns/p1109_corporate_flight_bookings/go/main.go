package corporate_flight_bookings

// Returns an array of the number of flights booked for each flight. Flight numbers range from 1
// to n and each booking is an interval representing [first flight, last flight, number of seats].
// Each flight in the interval of first <= i <= last has the given number of seats booked.
//
// This solution uses a difference algorithm where the the start of the interval has the number of
// seats added and the next index after the interval has the number of seats subtracted. Then the
// array can be iterated over, adding the previous number to the current one, which fills out the
// flight count array.
func corpFlightBookings(bookings [][]int, n int) []int {
	counts := make([]int, n+1)

	for _, interval := range bookings {
		counts[interval[0]-1] += interval[2]
		counts[interval[1]] -= interval[2]
	}

	for i := range counts {
		if i > 0 {
			counts[i] += counts[i-1]
		}
	}

	return counts[:n]
}

// Returns the number of bookings for each flight. There are 1 to n flights and each booking has
// the values bookings[i][0] = first flights, bookings[i][1] = last flight, bookings[i][2] =
// number of seats booked. The first and last flight represent the inclusive interval of flights
// that have seats booked.
//
// This is the brute force n^2 algorithm, passes all tests, but is pathetically slow.
func corpFlightBookingsSlow(bookings [][]int, n int) []int {
	counts := make([]int, n)

	for _, elem := range bookings {
		for i := elem[0] - 1; i <= elem[1]-1; i++ {
			counts[i] += elem[2]
		}
	}

	return counts
}
