package httpport

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getMondayMidnight(t time.Time) time.Time {
	weekDay := t.Weekday()
	daysToSubstract := int(weekDay) - int(time.Monday)

	// If it is Sunday(0), adjust the number.
	if daysToSubstract < 0 {
		daysToSubstract = 6
	}

	monday := t.AddDate(0, 0, -daysToSubstract)
	return time.Date(
		monday.Year(),
		monday.Month(),
		monday.Day(),
		0,
		0,
		0,
		0,
		monday.Location(),
	)
}

func generateTimeSlots(t time.Time, hours []int, weeks int) []time.Time {
	slots := []time.Time{}
	monday := getMondayMidnight(t)

	for weeks > 0 {
		for i := range 5 {
			for _, h := range hours {
				slot := monday.AddDate(0, 0, i).Add(time.Duration(h) * time.Hour)
				slots = append(slots, slot)
			}
		}
		// move to next Monday.
		monday = monday.AddDate(0, 0, 7)
		weeks -= 1
	}

	return slots
}

func (h handler) generateSlots(c *gin.Context) {
	weeks, _ := strconv.Atoi(c.Query("weeks"))
	if weeks <= 0 {
		weeks = 1
	}

	workingHours := []int{8, 9, 10, 11, 13, 14, 15, 16}
	slots := generateTimeSlots(time.Now(), workingHours, weeks)

	c.JSON(200, slots)
}
