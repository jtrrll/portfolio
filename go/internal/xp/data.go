package xp

import (
	"database/sql"
	"time"
)

// Hardcoded experiences. TODO: Implement a database.
var (
	Proof = WorkExperience{
		Start:   time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC),
		Title:   "Full-Stack Software Engineer",
		Company: "Proof",
	}
	Vmware = WorkExperience{
		Start:   time.Date(2023, 1, 9, 0, 0, 0, 0, time.UTC),
		End:     sql.NullTime{Time: time.Date(2023, 8, 29, 0, 0, 0, 0, time.UTC), Valid: true},
		Title:   "Back-End Software Engineer",
		Company: "VMware",
	}
)
