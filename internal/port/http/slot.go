package httpport

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-doctor-appointment/internal/app/query"
	"github.com/tinhtt/go-doctor-appointment/internal/domain"
)

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
	slots, err := h.app.Query.ListSlots(c, q)
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
