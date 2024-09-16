package tools

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type APIindex struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

type APIlocations struct {
	Index []struct {
		Location []string `json:"locations"`
	} `json:"index"`
}

type APIdates struct {
	Index []struct {
		Dates []string `json:"dates"`
	} `json:"index"`
}

type APIrelations struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type Card struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
	Relation     map[string][]string
}


type PageData struct {
	Cards        []Card
	ArtistNames  map[string]bool
	Members      map[string]bool
	CreationDates map[string]bool
	FirstAlbums  map[string]bool
	Locations    map[string]bool
}
