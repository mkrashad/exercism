package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
  	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout,date)
    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	appointment, _ := time.Parse("January 2, 2006 15:04:05", date)
    return time.Now().After(appointment)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	appointment, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := appointment.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointment, _ := time.Parse("1/2/2006 15:04:05", date)
	formatted := appointment.Format("Monday, January 2, 2006, at 15:04")
    return "You have an appointment on " + formatted + "."
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	year := time.Now().UTC().Year() // current year
	month := time.September        // correct month
	day := 15                      // correct day

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
