package crime

import "fmt"

type rawCode struct {
	Code         int64
	Extension    int64
	TypeID       string
	TypeName     string
	CategoryID   string
	CategoryName string
	IsCrime      int64
	IsTraffic    int64
}

type Code struct {
	ID          string `json:"id"`
	Code        int64  `json:"code"`
	Extension   int64  `json:"extension"`
	Description string `json:"description"`
	Category    string `json:"category"`
	IsCrime     bool   `json:"is_crime"`
	IsTraffic   bool   `json:"is_traffic"`
}

func (c Code) String() string {
	return fmt.Sprintf("%s: %s (%s)", c.ID, c.Description, c.Category)
}