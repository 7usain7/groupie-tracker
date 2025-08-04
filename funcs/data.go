package funcs

var ExcludeIDs = []int{11, 12, 21, 22, 49}

type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type RelationsData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ArtistDetailPage struct {
	Artist   Artists
	Relation RelationsData
}

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type PageData struct {
	Artist   []Artists
	Location Locations
	// Suggestions []string
}

// Allowed query parameters
var validQueries = map[string]bool{
	"searchQuary": true,
	// "page":        true, // (Not yet implemented)
}

var NotFound = struct {
	Code    string
	Title   string
	Message string
}{
	Code:    "404",
	Title:   "Page not found",
	Message: "Sorry, this page was not found",
}

var BadRequest = struct {
	Code    string
	Title   string
	Message string
}{
	Code:    "400",
	Title:   "Bad Request",
	Message: "Sorry, Bad request.",
}

var MethodNotAllowed = struct {
	Code    string
	Title   string
	Message string
}{
	Code:    "405",
	Title:   "Method Not Allowed",
	Message: "Sorry, This method is not allowed.",
}

var InternalServerError = struct {
	Code    string
	Title   string
	Message string
}{
	Code:    "500",
	Title:   "Internal Error",
	Message: "Zzz, Internal Server Error.",
}
