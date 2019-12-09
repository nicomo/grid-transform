package cmd

type InstituteTarget struct {
	UID          string   `json:"uid"`
	DgraphType   string   `json:"dgraph.type"`
	Name         string   `json:"name"`
	WikipediaURL string   `json:"wikipedia_url,omitempty"`
	Link         string   `json:"link"`
	Location     Location `json:"location"`
	City         string   `json:"city"`
	Country      string   `json:"country"`
	CountryCode  string   `json:"country_code"`
	GeonamesCity int      `json:"geonames_city"`
	LabelFr      string   `json:"label@fr,omitempty"` //TODO: generate proper json based on lang attr
	Status       string   `json:"status"`
	Established  int      `json:"established,omitempty"`
	Children     []Child  `json:"children,omitempty"`
	Xids         []Xids   `json:"xids"`
	Acronyms     []string `json:"acronyms,omitempty"`
	Parents      []Parent `json:"parents,omitempty"`
}
type Location struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
type Child struct {
	UID string `json:"uid"`
}
type Xids struct {
	Source string `json:"source"`
	Xid    string `json:"xid"`
}
type Parent struct {
	UID string `json:"uid"`
}

///////////////// SOURCE STRUCT: USE TO UNMARSHAL //////////////////////

type InstituteSource struct {
	Name          string                 `json:"name"`
	WikipediaURL  string                 `json:"wikipedia_url"`
	Links         []string               `json:"links"`
	Acronyms      []string               `json:"acronyms"`
	Addresses     []Addresses            `json:"addresses"`
	Labels        []interface{}          `json:"labels"`
	ID            string                 `json:"id"`
	Status        string                 `json:"status"`
	Established   int                    `json:"established"`
	Relationships []Relationship         `json:"relationships"`
	ExternalIds   map[string]interface{} `json:"external_ids,omitempty"`
}

type GeonamesCity struct {
	GCID int `json:"id"`
}

type Addresses struct {
	Lat          float64      `json:"lat"`
	Lng          float64      `json:"lng"`
	City         string       `json:"city"`
	Country      string       `json:"country"`
	CountryCode  string       `json:"country_code"`
	GeonamesCity GeonamesCity `json:"geonames_city"`
}

type Relationship struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	ID    string `json:"id"`
}
