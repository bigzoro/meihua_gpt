package models

import (
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Divination captures the result of a single Mei Hua Yi Shu casting.
type Divination struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Subject       string    `json:"subject"`
	Method        string    `json:"method"`
	PrimaryName   string    `json:"primaryName"`
	UpperTrigram  string    `json:"upperTrigram"`
	LowerTrigram  string    `json:"lowerTrigram"`
	SecondaryName string    `json:"secondaryName"`
	Commentary    string    `json:"commentary" gorm:"type:text"`
	LinesRaw      string    `json:"-" gorm:"column:lines"`
	ChangesRaw    string    `json:"-" gorm:"column:changing_lines"`
	Lines         []int     `json:"lines" gorm:"-"`
	ChangingLines []int     `json:"changingLines" gorm:"-"`
}

// BeforeSave serialises slice fields ahead of persistence.
func (d *Divination) BeforeSave(tx *gorm.DB) error {
	if len(d.Lines) > 0 {
		parts := make([]string, len(d.Lines))
		for i, v := range d.Lines {
			parts[i] = strconv.Itoa(v)
		}
		d.LinesRaw = strings.Join(parts, ",")
	}

	if len(d.ChangingLines) > 0 {
		parts := make([]string, len(d.ChangingLines))
		for i, v := range d.ChangingLines {
			parts[i] = strconv.Itoa(v)
		}
		d.ChangesRaw = strings.Join(parts, ",")
	}
	return nil
}

// AfterFind hydrates slice fields following database reads.
func (d *Divination) AfterFind(tx *gorm.DB) error {
	if d.LinesRaw != "" {
		d.Lines = parseIntSlice(d.LinesRaw)
	}
	if d.ChangesRaw != "" {
		d.ChangingLines = parseIntSlice(d.ChangesRaw)
	}
	return nil
}

func parseIntSlice(raw string) []int {
	parts := strings.Split(raw, ",")
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		if p == "" {
			continue
		}
		if v, err := strconv.Atoi(p); err == nil {
			result = append(result, v)
		}
	}
	return result
}
