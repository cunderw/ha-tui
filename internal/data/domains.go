package data

import "github.com/cunderw/ha-tui/internal/model"

func GetDomains() []model.Domain {
	var domains = []model.Domain{
		{
			ID:   "1",
			Name: "light",
		},
		{
			ID:   "2",
			Name: "switch",
		},
		{
			ID:   "3",
			Name: "sensor",
		},
		{
			ID:   "4",
			Name: "battery",
		},
	}

	return domains
}
