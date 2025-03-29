package httpport

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-doctor-appointment/internal/app/command"
	"github.com/tinhtt/go-doctor-appointment/internal/app/query"
	"github.com/tinhtt/go-doctor-appointment/internal/common/auth"
	"github.com/tinhtt/go-doctor-appointment/internal/domain"
)

type doctor struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type patient struct {
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

func (r *signupDoctorReq) bind(c *gin.Context, cmd *command.RegisterDoctor) error {
	if err := c.ShouldBindJSON(r); err != nil {
		c.AbortWithError(400, err)
		return err
	}

	cmd.Username = r.Username
	cmd.Email = r.Email
	cmd.Password = r.Password
	return nil
}

func (h handler) signupDoctor(c *gin.Context) {
	var (
		req signupDoctorReq
		cmd command.RegisterDoctor
	)
	if err := req.bind(c, &cmd); err != nil {
		c.AbortWithError(400, err)
		return
	}

	err := h.app.Command.RegisterDoctor(c, cmd)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
}

type signinDoctorReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *signinDoctorReq) bind(c *gin.Context, q *query.LoginDoctor) error {
	if err := c.ShouldBindJSON(&r); err != nil {
		return err
	}

	q.Username = r.Username
	q.Password = r.Password
	return nil
}

type signinDoctorRes struct {
	Doctor doctor `json:"doctor"`
	Token  string `json:"token"`
}

func (r *signinDoctorRes) from(d domain.Doctor, token string) {
	r.Doctor.ID = d.ID
	r.Doctor.Username = d.Username
	r.Doctor.Email = d.Email
	r.Doctor.CreatedAt = d.CreatedAt
	r.Doctor.UpdatedAt = d.UpdatedAt
	r.Token = token
}

func (h handler) signinDoctor(c *gin.Context) {
	var (
		req signinDoctorReq
		q   query.LoginDoctor
	)
	if err := req.bind(c, &q); err != nil {
		c.AbortWithError(400, err)
		return
	}

	d, err := h.app.Query.LoginDoctor(c, q)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	token, err := h.jwt.Generate(auth.User{
		ID:       d.ID,
		Role:     "doctor",
		Username: d.Username,
	})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	var res signinDoctorRes
	res.from(d, token)
	c.JSON(200, res)
}

type signupPatientReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *signupPatientReq) bind(c *gin.Context, cmd *command.RegisterPatient) error {
	if err := c.ShouldBindJSON(r); err != nil {
		c.AbortWithError(400, err)
		return err
	}

	cmd.Username = r.Username
	cmd.Email = r.Email
	cmd.Password = r.Password
	return nil
}

func (h handler) signupPatient(c *gin.Context) {
	var (
		req signupPatientReq
		cmd command.RegisterPatient
	)
	if err := req.bind(c, &cmd); err != nil {
		c.AbortWithError(400, err)
		return
	}

	err := h.app.Command.RegisterPatient(c, cmd)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
}

type signinPatientReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *signinPatientReq) bind(c *gin.Context, q *query.LoginPatient) error {
	if err := c.ShouldBindJSON(&r); err != nil {
		return err
	}

	q.Username = r.Username
	q.Password = r.Password
	return nil
}

type signinPatientRes struct {
	Patient patient `json:"patient"`
	Token   string  `json:"token"`
}

func (r *signinPatientRes) from(p domain.Patient, token string) {
	r.Patient.ID = p.ID
	r.Patient.Username = p.Username
	r.Patient.Email = p.Email
	r.Patient.CreatedAt = p.CreatedAt
	r.Patient.UpdatedAt = p.UpdatedAt
	r.Token = token
}

func (h handler) signinPatient(c *gin.Context) {
	var (
		req signinPatientReq
		q   query.LoginPatient
	)
	if err := req.bind(c, &q); err != nil {
		c.AbortWithError(400, err)
		return
	}

	p, err := h.app.Query.LoginPatient(c, q)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	token, err := h.jwt.Generate(auth.User{
		ID:       p.ID,
		Role:     "patient",
		Username: p.Username,
	})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	var res signinPatientRes
	res.from(p, token)
	c.JSON(200, res)
}
