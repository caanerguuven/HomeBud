package models

import "time"

//GetMonths returns list
func GetMonths() map[int]string {
	return map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
}

//GetYears returns list
func GetYears() map[int]int {
	years := make(map[int]int)

	startYear := time.Time.Year(time.Now())
	for i := 0; i <= 50; i++ {
		years[i] = startYear
		startYear++
	}

	return years
}
