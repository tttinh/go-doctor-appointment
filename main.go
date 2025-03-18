package main

import (
	"github.com/tinhtt/go-doctor-appointment/internal/ports"
)

func main() {
	s := ports.NewHTTP()
	s.ListenAndServe()
}
