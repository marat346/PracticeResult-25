package cities

type City struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	District   string `json:"district"`
	Population int    `json:"population"`
	Foundation int    `json:"foundation"`
}

type CityRequest struct {
	Name       string `json:"name"`
	Region     string `json:"region"`
	District   string `json:"district"`
	Population int    `json:"population"`
	Foundation int    `json:"foundation"`
}

type SetPopulationRequest struct {
	Population int `json:"population"`
}

type RangeRequest struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
