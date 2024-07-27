package models

type Host struct {
	ID           int64
	HostName     string
	Environments int64
	Type         int64
	Product      int64
	Site         int64
	Provider     int64
	Location     string
}
