package leap

const TestVersion = 1

// IsLeapYear takes a year and returns true if said year is a leap year or
// false if it is not
func IsLeapYear(year int) bool {
	return year%400 == 0 && year%4 == 0 || year%4 == 0 && year%100 != 0
}
