package time

import (
	"log"
	"math"
	"strconv"
	"time"
)

func CountDigitInNumber(number int64) int {
	count := 0

	for number != 0 {
		number /= 10
		count += 1
	}

	return count
}

func UnixToTime(unixtime int64) time.Time {
	c := CountDigitInNumber(unixtime)

	var tm time.Time
	if c == 10 { //nolint:gomnd
		tm = time.Unix(unixtime, 0)
	} else if c == 13 { //nolint:gomnd
		tm = time.Unix(0, unixtime*int64(time.Millisecond))
	}

	return tm
}

func GetSecondsToTimeFromNow(unixtime int64) int {
	t1 := UnixToTime(unixtime)
	t2 := time.Now()

	diff := t1.Sub(t2)

	return int(math.Round(diff.Seconds()))
}

func GetSecondsFromTimeToNow(unixtime int64) int {
	t1 := UnixToTime(unixtime)
	t2 := time.Now()

	diff := t2.Sub(t1)

	return int(math.Round(diff.Seconds()))
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func FormatDuration(d time.Duration) string {
	if d < time.Millisecond {
		return strconv.FormatInt(int64(d), 10) + "nanos"
	}

	return strconv.FormatInt(int64(d)/int64(time.Millisecond), 10) + "ms"
}

func NowUnixtimeMilliseconds() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
