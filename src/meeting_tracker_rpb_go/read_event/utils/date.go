package read_event

import "time"

func DateEqual(date1, date2 time.Time) bool {
    y1, m1, d1 := date1.Date()
    y2, m2, d2 := date2.Date()
    return y1 == y2 && m1 == m2 && d1 == d2
}


func TimeNow() time.Time {
	//init the loc
	loc, _ := time.LoadLocation("Europe/Kyiv")
	//set timezone,  
	return time.Now().In(loc)
}