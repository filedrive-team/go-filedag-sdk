package model

type MetaData struct {
	// Optional name for pinned data; can be used for lookups later
	Name string `json:"name,omitempty" example:"PreciousData.pdf"`

	// Optional metadata for pin object
	Meta map[string]string `json:"meta,omitempty" example:"app_id:99986338-1113-4706-8302-4420da6158aa"`
}
