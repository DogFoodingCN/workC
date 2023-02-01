package models

type Job struct {
	Base        Base        `json:"base"`
	Requirement Requirement `json:"requirement"`
	Company     Company     `json:"company"`
}

type Base struct {
	Name   string   `json:"name"`
	Area   string   `json:"area"`
	Salary string   `json:"salary"`
	Link   string   `json:"link"`
	Tags   []string `json:"tags"`
}

type Requirement struct {
	Tags []string `json:"tags"`
}

type Company struct {
	Name string   `json:"name"`
	Desc string   `json:"desc"`
	Logo string   `json:"logo"`
	Link string   `json:"link"`
	Tags []string `json:"tags"`
}
