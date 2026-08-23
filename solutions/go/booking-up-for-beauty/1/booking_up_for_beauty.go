package booking

import "time"
import "fmt"
import "strconv"

func parseOrPanic(date string, layout string) time.Time {
    t, err := time.Parse(layout, date)
	if err != nil {
        panic(err)
    }
    
    return t 
}


//Mon Jan 2 15:04:05 MST 2006
// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	return parseOrPanic(date, layout)
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    apptDate := parseOrPanic(date, layout)
    return apptDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
    apptDate := parseOrPanic(date, layout)
    return apptDate.Hour() >= 12 && apptDate.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    apptDate := parseOrPanic(date, layout)

    outputLayout := "Monday, January 2, 2006"
    outputLayoutTime := "15:04"
    return fmt.Sprintf("You have an appointment on %s, at %s.", apptDate.Format(outputLayout), apptDate.Format(outputLayoutTime))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "2006-01-02"
    anniversary := parseOrPanic(strconv.Itoa(time.Now().Year()) + "-09-15", layout)
    return anniversary
}
