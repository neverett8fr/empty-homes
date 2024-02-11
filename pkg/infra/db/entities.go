package db

import (
	"time"
)

type Home struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Postcode        string    `json:"postcode"`
	Street          string    `json:"street"`
	Type            string    `json:"type"` // e.g. is it residential? or public land, etc.
	Bedrooms        int       `json:"bedrooms"`
	Inhabited       bool      `json:"inhabited"`
	Safe            bool      `json:"safe"`
	OwnerName       string    `json:"owner_name"`
	OwnerContact    string    `json:"owner_contact"`
	DateAdded       time.Time `json:"date_added"`
	DateLastChecked time.Time `json:"date_last_checked"`
}
