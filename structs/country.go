package structs

type Country struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Iso2     string `json:"iso2"`
	Iso3     string `json:"iso3"`
	Capital  string `json:"capital"`
	Region   string `json:"region"`
	Language string `json:"language"`
}