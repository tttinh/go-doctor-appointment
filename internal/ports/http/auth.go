package httpport

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-doctor-appointment/internal/app"
	"github.com/tinhtt/go-doctor-appointment/internal/common/auth"
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

type signinDoctorReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *signinDoctorReq) bind(c *gin.Context, d *app.Doctor) error {
	if err := c.ShouldBindJSON(&r); err != nil {
		return err
	}

	d.Username = r.Username
	d.Password = r.Password
	return nil
}

type signinDoctorRes struct {
	Doctor doctor `json:"doctor"`
	Token  string `json:"token"`
}

func (r *signinDoctorRes) from(d app.Doctor, token string) {
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
		d   app.Doctor
	)
	if err := req.bind(c, &d); err != nil {
		c.AbortWithError(400, err)
		return
	}

	d, err := h.user.SigninDoctor(c, d.Username, d.Password)
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

func (r *signupPatientReq) bind(c *gin.Context, d *app.Patient) error {
	if err := c.ShouldBindJSON(r); err != nil {
		c.AbortWithError(400, err)
		return err
	}

	d.Username = r.Username
	d.Email = r.Email
	d.Password = r.Password
	return nil
}

type signupPatientRes struct {
	patient
}

func (r *signupPatientRes) from(p app.Patient) {
	r.patient.ID = p.ID
	r.patient.Username = p.Username
	r.patient.Email = p.Email
	r.patient.CreatedAt = p.CreatedAt
	r.patient.UpdatedAt = p.UpdatedAt
}

func (h handler) signupPatient(c *gin.Context) {
	var (
		req signupPatientReq
		p   app.Patient
	)
	if err := req.bind(c, &p); err != nil {
		c.AbortWithError(400, err)
		return
	}

	p, err := h.user.SignupPatient(c, p)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	var res signupPatientRes
	res.from(p)
	c.JSON(200, res)
}

type signinPatientReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *signinPatientReq) bind(c *gin.Context, p *app.Patient) error {
	if err := c.ShouldBindJSON(&r); err != nil {
		return err
	}

	p.Username = r.Username
	p.Password = r.Password
	return nil
}

type signinPatientRes struct {
	Patient patient `json:"patient"`
	Token   string  `json:"token"`
}

func (r *signinPatientRes) from(p app.Patient, token string) {
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
		p   app.Patient
	)
	if err := req.bind(c, &p); err != nil {
		c.AbortWithError(400, err)
		return
	}

	p, err := h.user.SigninPatient(c, p.Username, p.Password)
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
