package main

import (
	"fmt"
	"time"
)

type Subscription struct {
	Id                    int
	CustomerId            int
	MonthlyPriceInDollars int
}

type User struct {
	Id            int
	Name          string
	ActivatedOn   time.Time
	DeactivatedOn time.Time
	CustomerId    int
}

func main() {
	sub := Subscription{
		Id:                    1,
		CustomerId:            1,
		MonthlyPriceInDollars: 4,
	}

	users := []User{
		{
			Id:            1,
			Name:          "test-user",
			ActivatedOn:   time.Time{},
			DeactivatedOn: time.Time{},
			CustomerId:    0,
		},
	}

	_ = BillFor("2019-01", &sub, &users)
}

func BillFor(yearMonth string, activeSubscription *Subscription, users *[]User) float64 {
	var format = "2006-01"
	var totalCost float64

	if activeSubscription == nil {
		return -1.0
	}

	date, err := time.Parse(format, yearMonth)
	if err != nil {
		panic("Cannot parse date")
	}

	// Calculate daily rate for subscription
	diff := LastDayOfMonth(date).Sub(FirstDayOfMonth(date))
	numDays := int(diff.Hours()/24) + 1
	pricePerDay := float64(activeSubscription.MonthlyPriceInDollars) / float64(numDays)
	fmt.Println(pricePerDay)

	// Calculate price for subscription
	for _, user := range *users {
		if user.CustomerId == activeSubscription.CustomerId {
			if user.DeactivatedOn.IsZero() {
				// Todo: sum days from the beginning of the month till now
			} else {
				// Todo: sum days from the beginning of the month until deactivation date
			}
		}
	}

	// Todo: round result to 2 decimals

	return totalCost
}

/*******************
* Helper functions *
*******************/

/*
Takes a time.Time object and returns a time.Time object
which is the first day of that month.

FirstDayOfMonth(time.Date(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb 7
=> time.Date(2019, 2, 1, 0, 0, 0, 0, time.UTC))               // Feb 1
*/
func FirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

/*
Takes a time.Time object and returns a time.Time object
which is the end of the last day of that month.

LastDayOfMonth(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb  7
=> time.Time(2019, 2, 28, 23, 59, 59, 0, time.UTC)           // Feb 28
*/
func LastDayOfMonth(t time.Time) time.Time {
	return FirstDayOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

/*
Takes a time.Time object and returns a time.Time object
which is the next day.

NextDay(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))   // Feb 7
=> time.Time(2019, 2, 8, 0, 0, 0, 0, time.UTC)         // Feb 8

NextDay(time.Time(2019, 2, 28, 0, 0, 0, 0, time.UTC))  // Feb 28
=> time.Time(2019, 3, 1, 0, 0, 0, 0, time.UTC)         // Mar  1
*/
func NextDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}
