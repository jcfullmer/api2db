package main

type npsBody struct {
	Data []struct {
		ID          string `json:"id"`
		URL         string `json:"url"`
		FullName    string `json:"fullName"`
		ParkCode    string `json:"parkCode"`
		Description string `json:"description"`
		Activities  []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"activities"`
		Topics []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"topics"`
		States   string `json:"states"`
		Contacts struct {
			PhoneNumbers []struct {
				PhoneNumber string `json:"phoneNumber"`
				Description string `json:"description"`
				Extension   string `json:"extension"`
				Type        string `json:"type"`
			} `json:"phoneNumbers"`
		} `json:"contacts"`
		EntranceFees   []any `json:"entranceFees"`
		EntrancePasses []any `json:"entrancePasses"`
		Fees           []any `json:"fees"`
		OperatingHours []struct {
			Exceptions    []any  `json:"exceptions"`
			Description   string `json:"description"`
			StandardHours struct {
				Wednesday string `json:"wednesday"`
				Monday    string `json:"monday"`
				Thursday  string `json:"thursday"`
				Sunday    string `json:"sunday"`
				Tuesday   string `json:"tuesday"`
				Friday    string `json:"friday"`
				Saturday  string `json:"saturday"`
			} `json:"standardHours"`
			Name string `json:"name"`
		} `json:"operatingHours"`
		Addresses []struct {
			PostalCode            string `json:"postalCode"`
			City                  string `json:"city"`
			StateCode             string `json:"stateCode"`
			CountryCode           string `json:"countryCode"`
			ProvinceTerritoryCode string `json:"provinceTerritoryCode"`
			Line1                 string `json:"line1"`
			Type                  string `json:"type"`
			Line3                 string `json:"line3"`
			Line2                 string `json:"line2"`
		} `json:"addresses"`
		Images []struct {
			Credit  string `json:"credit"`
			Title   string `json:"title"`
			AltText string `json:"altText"`
			Caption string `json:"caption"`
			URL     string `json:"url"`
		} `json:"images"`
		WeatherInfo string `json:"weatherInfo"`
		Name        string `json:"name"`
		Designation string `json:"designation"`
		Multimedia  []struct {
			Title string `json:"title"`
			ID    string `json:"id"`
			Type  string `json:"type"`
			URL   string `json:"url"`
		} `json:"multimedia"`
		RelevanceScore float64 `json:"relevanceScore"`
	} `json:"data"`
}
