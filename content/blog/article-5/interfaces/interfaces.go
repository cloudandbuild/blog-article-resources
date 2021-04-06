package interfaces

import "time"

//go:generate go install github.com/golang/mock/mockgen
//go:generate mockgen -destination=${GOPACKAGE}mocks/${GOFILE} -package=${GOPACKAGE}mocks -source=$GOFILE
type IRepository interface {
	Update(data *Data) error
}
type Data struct {
	ID        int
	UpdatedAt time.Time
}
