package httpport

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-doctor-appointment/internal/app/command"
	"github.com/tinhtt/go-doctor-appointment/internal/app/query"
	"github.com/tinhtt/go-doctor-appointment/internal/domain"
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

func generateTimeSlots(t time.Time, hours []int) []time.Time {
	slots := []time.Time{}
	monday := getMondayMidnight(t)
	for i := 0; i < 5; i++ {
		for _, h := range hours {
			slot := monday.AddDate(0, 0, i).Add(time.Duration(h) * time.Hour)
			slots = append(slots, slot)
		}
	}

	return slots
}

func (h handler) mockSlots(c *gin.Context) {
	u, _ := userFromContext(c)
	if u.Role != "doctor" {
		c.AbortWithError(403, domain.ErrAccessDenied)
		return
	}

	workingHours := []int{8, 9, 10, 11, 13, 14, 15, 16}
	cmd := command.CreateSlots{
		DoctorID: u.ID,
		Slots:    generateTimeSlots(time.Now(), workingHours),
	}
	err := h.app.Command.CreateSlots.Handle(c, cmd)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
}

type slot struct {
	ID        int       `json:"id"`
	Hour      time.Time `json:"hour"`
	Available bool      `json:"available"`
}

type listSlotsRes struct {
	Slots []slot `json:"slots"`
}

func (r *listSlotsRes) from(slots []domain.Slot) {
	r.Slots = []slot{}
	for _, s := range slots {
		r.Slots = append(r.Slots, slot{
			ID:        s.ID,
			Hour:      s.Hour,
			Available: s.Available,
		})
	}
}

func (h handler) listSlots(c *gin.Context) {
	u, _ := userFromContext(c)
	if u.Role != "doctor" {
		c.AbortWithError(403, domain.ErrAccessDenied)
		return
	}

	q := query.ListSlots{DoctorID: u.ID}
	slots, err := h.app.Query.ListSlots.Handle(c, q)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	var res listSlotsRes
	res.from(slots)
	c.JSON(200, res)
}

func (h handler) addSlots(c *gin.Context) {

}

func (h handler) changeSlotAvailability(c *gin.Context) {

}
