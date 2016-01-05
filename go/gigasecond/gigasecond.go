package gigasecond

import "time"

const TestVersion = 2

var gigasec = 1e+09 * time.Second

// AddGigasecond adds one gigasecond to the given times
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasec)
}

// Birthday is well... my birthday :D
var Birthday, _ = time.Parse("2006-Jan-02", "1990-Nov-08")
