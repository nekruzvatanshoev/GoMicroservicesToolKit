package web

type Country struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Locale string `json:"locale"`
	CurrencyId string `json:"currency_id"`
	DecimalSeparator string `json:"decimal_separator"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone string `json:"time_zone"`
	GeoInformation GeoInformation `json:"geo_information"`
	States []State `json:"states"`
}


type GeoInformation struct {
	Location GeoLocation `json:"location"`
}

type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}


type State struct {
	Id string `json:"id"`
	Name string `json:"name"`
}



var US Country = Country{
	Id:                 "US",
	Name:               "United States of America",
	Locale:             "en_US",
	CurrencyId:         "USD",
	DecimalSeparator:   ".",
	ThousandsSeparator: ",",
	TimeZone:           "GMT-04:00",
	GeoInformation:     GeoInformation{
		GeoLocation{
			Latitude:  38.8867036,
			Longitude: -77.0081972,
		}},
	States:             []State{
		{
			Id: "US-AL",
			Name: "Alabama",
		},
		{
			Id: "US-AK",
			Name: "Alaska",
		},
		{
			Id: "US-AZ",
			Name: "Arizona",
		},
		{
			Id: "US-AR",
			Name: "Arkansas",
		},
		{
			Id: "US-CA",
			Name: "California",
		},
		{
			Id: "US-CO",
			Name: "Colorado",
		},
		{
			Id: "US-CT",
			Name: "Connecticut",
		},
		{
			Id: "US-DE",
			Name: "Delaware",
		},
		{
			Id: "US-DC",
			Name: "District of Columbia",
		},
		{
			Id: "US-FL",
			Name: "Florida",
		},
		{
			Id: "US-GA",
			Name: "Georgia",
		},
		{
			Id: "US-HI",
			Name: "Hawaii",
		},
		{
			Id: "US-ID",
			Name: "Idaho",
		},
		{
			Id: "US-IL",
			Name: "Illinois",
		},
		{
			Id: "US-IN",
			Name: "Indiana",
		},
		{
			Id: "US-IA",
			Name: "Iowa",
		},
		{
			Id: "US-KS",
			Name: "Kansas",
		},
		{
			Id: "US-KY",
			Name: "Kentucky",
		},
		{
			Id: "US-LA",
			Name: "Louisiana",
		},
		{
			Id: "US-ME",
			Name: "Maine",
		},
		{
			Id: "US-MD",
			Name: "Maryland",
		},
		{
			Id: "US-MA",
			Name: "Massachusetts",
		},
		{
			Id: "US-MI",
			Name: "Michigan",
		},
		{
			Id: "US-MN",
			Name: "Minnesota",
		},
		{
			Id: "US-MS",
			Name: "Mississippi",
		},
		{
			Id: "US-MO",
			Name: "Missouri",
		},
		{
			Id: "US-MT",
			Name: "Montana",
		},
		{
			Id: "US-NE",
			Name: "Nebraska",
		},
		{
			Id: "US-NV",
			Name: "Nevada",
		},
		{
			Id: "US-NH",
			Name: "New Hampshire",
		},
		{
			Id: "US-NJ",
			Name: "New Jersey",
		},
		{
			Id: "US-NM",
			Name: "New Mexico",
		},
		{
			Id: "US-NY",
			Name: "New York",
		},
		{
			Id: "US-NC",
			Name: "North Carolina",
		},
		{
			Id: "US-ND",
			Name: "North Dakota",
		},
		{
			Id: "US-OH",
			Name: "Ohio",
		},
		{
			Id: "US-OK",
			Name: "Oklahoma",
		},
		{
			Id: "US-OR",
			Name: "Oregon",
		},
		{
			Id: "US-PA",
			Name: "Pennsylvania",
		},
		{
			Id: "US-PR",
			Name: "Puerto Rico",
		},
		{
			Id: "US-RI",
			Name: "Rhode Island",
		},
		{
			Id: "US-SC",
			Name: "South Carolina",
		},
		{
			Id: "US-SD",
			Name: "South Dakota",
		},
		{
			Id: "US-TN",
			Name: "Tennessee",
		},
		{
			Id: "US-TX",
			Name: "Texas",
		},
		{
			Id: "US-UT",
			Name: "Utah",
		},
		{
			Id: "US-VT",
			Name: "Vermont",
		},
		{
			Id: "US-VI",
			Name: "Virgin Islands",
		},
		{
			Id: "US-VA",
			Name: "Virginia",
		},
		{
			Id: "US-WA",
			Name: "Washington",
		},
		{
			Id: "US-WV",
			Name: "West Virginia",
		},
		{
			Id: "US-WI",
			Name: "Wisconsin",
		},
		{
			Id: "US-WY",
			Name: "Wyoming",
		},
	},
}
