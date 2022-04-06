package data

import "github.com/cunderw/ha-tui/internal/model"

func GetEntities(domain *string) []model.Entity {
	entityMap := make(map[string][]model.Entity)

	entityMap["light"] = []model.Entity{
		{
			Domain: "light",
			ID:     "light.office_carson_desk_1",
			Name:   "Office Carson Desk 1",
			State:  "on",
		},
		{
			Domain: "light",
			ID:     "light.office_carson_desk_2",
			Name:   "Office Carson Desk 2",
			State:  "on",
		},
		{
			Domain: "light",
			ID:     "light.office_ceiling_1",
			Name:   "Office Ceiling 1",
			State:  "off",
		},
		{
			Domain: "light",
			ID:     "light.office_ceiling_2",
			Name:   "Office Ceiling 1",
			State:  "off",
		},
		{
			Domain: "light",
			ID:     "light.office_ceiling_3",
			Name:   "Office Ceiling 3",
			State:  "off",
		},
	}

	entityMap["switch"] = []model.Entity{
		{
			Domain: "switch",
			ID:     "switch.office_carson_fan_lights",
			Name:   "Office Ceiling Fan Lights",
			State:  "one",
		},
		{
			Domain: "switch",
			ID:     "switch.office_bathroom_lights",
			Name:   "Office Bathroom Lights",
			State:  "off",
		},
	}

	return entityMap[*domain]
}
