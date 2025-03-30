package httpport

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-doctor-appointment/internal/app/command"
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

type createSlotsReq struct {
	Slots []time.Time `json:"slots"`
}

func (r *createSlotsReq) bind(c *gin.Context, cmd *command.CreateSlots) error {
	if err := c.ShouldBindJSON(&r); err != nil {
		return err
	}

	cmd.Slots = r.Slots
	return nil
}

func (h handler) addSlots(c *gin.Context) {
	u, _ := userFromContext(c)
	if u.Role != "doctor" {
		c.AbortWithError(403, domain.ErrAccessDenied)
		return
	}

	cmd := command.CreateSlots{DoctorID: u.ID}
	var req createSlotsReq
	if err := req.bind(c, &cmd); err != nil {
		c.AbortWithError(400, err)
		return

	}

	if err := h.app.Command.CreateSlots(c, cmd); err != nil {
		c.AbortWithError(400, err)
		return
	}
}

type changeSlotAvailabilityReq struct {
	Available bool `json:"available"`
}

func (r *changeSlotAvailabilityReq) bind(c *gin.Context, cmd *command.ChangeSlotAvailability) error {
	if err := c.ShouldBindJSON(&r); err != nil {
		return err
	}

	cmd.Available = r.Available
	return nil
}

func (h handler) changeSlotAvailability(c *gin.Context) {
	u, _ := userFromContext(c)
	if u.Role != "doctor" {
		c.AbortWithError(403, domain.ErrAccessDenied)
		return
	}

	slotID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	var (
		req changeSlotAvailabilityReq
		cmd command.ChangeSlotAvailability
	)

	if err := req.bind(c, &cmd); err != nil {
		c.AbortWithError(400, err)
		return
	}
	cmd.SlotID = slotID
	cmd.DoctorID = u.ID
	if err := h.app.Command.ChangeSlotAvailability(c, cmd); err != nil {
		c.AbortWithError(400, err)
		return
	}

}
