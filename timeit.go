// Simple timing package. Given an arbitrary function; measure total elapsed time, time spent in user/sys space, and cpu ticks.
package timeit

import (
	"fmt"
	"syscall"
	"time"
)

// Time spent:
// - Elapsed : total elapsed time during execution of the function.
// - User : time spent in user space during execution of the function.
// - Sys : time spent in sys space during execution of the function.
// - Ticks : number of cpu ticks that occurred during execution of the function.
type TimeSpent struct {
	Elapsed time.Duration
	User    int64
	Sys     int64
	Ticks   int64
}

func logResult(ts *TimeSpent) string {
	return fmt.Sprintf("Timeit: Elapsed=%s, User=%dus, Sys=%dus, Ticks=%d", ts.Elapsed.String(), ts.User, ts.Sys, ts.Ticks)
}

// Simple wrapper. Returns a string.
func Timeit(cb func()) string {
	cbW := func() interface{} { cb(); return true }
	ts, _ := TimeitT(cbW)
	return logResult(ts)
}

// Simple wrapper. Returns a string.
func Timeit1(cb func() interface{}) string {
	ts, _ := TimeitT(cb)
	return logResult(ts)
}

// Simple wrapper. Returns a string.
func Timeit2(cb func() (interface{}, interface{})) string {
	ts, _, _ := TimeitT2(cb)
	return logResult(ts)
}

// Simple wrapper. Prints a summary string.
func TimeitPrint(cb func()) {
	fmt.Println(Timeit(cb))
}

// Simple wrapper. Prints a summary string.
func Timeit1Print(cb func() interface{}) {
	fmt.Println(Timeit1(cb))
}

// Simple wrapper. Prints a summary string.
func Timeit2Print(cb func() (interface{}, interface{})) {
	fmt.Println(Timeit2(cb))
}

// Measures the pre/post timing information to execute cb(). Returns the results in a TimeSpent struct.
// Used with function of one return value.
func TimeitT(cb func() interface{}) (*TimeSpent, interface{}) {
	cbW := func() (interface{}, interface{}) { r1 := cb(); return r1, true }
	ts, r1, _ := TimeitT2(cbW)
	return ts, r1
}

// Measures the pre/post timing information to execute cb(). Returns the results in a TimeSpent struct.
// Used with functions of two return values.
func TimeitT2(cb func() (interface{}, interface{})) (*TimeSpent, interface{}, interface{}) {
	ts := new(TimeSpent)
	tmsPre := new(syscall.Tms)
	tmsPost := new(syscall.Tms)

	timePre := time.Now()
	ticksPre, _ := syscall.Times(tmsPre)
	r1, r2 := cb()
	ticksPost, _ := syscall.Times(tmsPost)
	timePost := time.Now()

	ts.Elapsed = timePost.Sub(timePre)
	ts.User = tmsPost.Utime - tmsPre.Utime
	ts.Sys = tmsPost.Stime - tmsPre.Stime
	ts.Ticks = int64(ticksPost - ticksPre)

	return ts, r1, r2
}
