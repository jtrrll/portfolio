// Provides functionality for exploring life experiences.
package xp

import (
	"database/sql"
	"time"
)

// A work experience.
type WorkExperience struct {
	Start time.Time    // The start of the experience.
	End   sql.NullTime // The end of the experience.

	Title       string // The job title.
	Company     string // The employer.
	Description string // A short summary of the experience.
}
