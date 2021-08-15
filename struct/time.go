package main

import "time"

//TIME
type Time struct {
	sec int64
	nsec int32

}

func Now()Time  {
	sec,nsec := Now()
	return Time{sec: sec + unixToInternal, nsec,Local}
}
func main()  {

}