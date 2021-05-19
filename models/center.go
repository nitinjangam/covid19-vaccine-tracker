package models

//AllCenters from COWIN API
type AllCenters struct {
	Centers []struct {
		CenterID      int     `json:"center_id"`
		Name          string  `json:"name"`
		NameL         string  `json:"name_l"`
		StateName     string  `json:"state_name"`
		StateNameL    string  `json:"state_name_l"`
		DistrictName  string  `json:"district_name"`
		DistrictNameL string  `json:"district_name_l"`
		BlockName     string  `json:"block_name"`
		BlockNameL    string  `json:"block_name_l"`
		Pincode       int     `json:"pincode"`
		Lat           float64 `json:"lat"`
		Long          float64 `json:"long"`
		From          string  `json:"from"`
		To            string  `json:"to"`
		FeeType       string  `json:"fee_type"`
		VaccineFees   []struct {
			Vaccine string `json:"vaccine"`
			Fee     string `json:"fee"`
		} `json:"vaccine_fees"`
		Sessions []struct {
			SessionID         string   `json:"session_id"`
			Date              string   `json:"date"`
			AvailableCapacity float64  `json:"available_capacity"`
			MinAgeLimit       int      `json:"min_age_limit"`
			Vaccine           string   `json:"vaccine"`
			Slots             []string `json:"slots"`
		} `json:"sessions"`
	} `json:"centers"`
}

//ValidateCenters function validates centers for available slots
func (c *AllCenters) ValidateCenters() AllCenters {
	p := AllCenters{}
	for _, v := range c.Centers {
		for _, v1 := range v.Sessions {
			if v1.AvailableCapacity > 0 {
				p.Centers = append(p.Centers, v)
			}
		}
	}
	return p
}
