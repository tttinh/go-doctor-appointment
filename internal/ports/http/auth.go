package httpport

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-doctor-appointment/internal/app"
)

type doctor struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type signupDoctorReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *signupDoctorReq) bind(c *gin.Context, d *app.Doctor) error {
	if err := c.ShouldBindJSON(r); err != nil {
		c.AbortWithError(400, err)
		return err
	}

	d.Username = r.Username
	d.Email = r.Email
	d.Password = r.Password
	return nil
}

type signupDoctorRes struct {
	doctor
}

func (r *signupDoctorRes) from(d app.Doctor) {
	r.doctor.ID = d.ID
	r.doctor.Username = d.Username
	r.doctor.Email = d.Email
	r.doctor.CreatedAt = d.CreatedAt
	r.doctor.UpdatedAt = d.UpdatedAt
}

func (h handler) signupDoctor(c *gin.Context) {
	var (
		req signupDoctorReq
		d   app.Doctor
	)
	if err := req.bind(c, &d); err != nil {
		c.AbortWithError(400, err)
		return
	}

	d, err := h.user.SignupDoctor(c, d)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	var res signupDoctorRes
	res.from(d)
	c.JSON(200, res)
}

func (h handler) signinDoctor(c *gin.Context) {

}

func (h handler) signupPatient(c *gin.Context) {

}

func (h handler) signinPatient(c *gin.Context) {

}
