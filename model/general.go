package model

// Data holds the actual information plus data for filtering and pagination
type Data struct {
	TotalData    int64
	FilteredData int64
	Data         interface{}
}

// Args is the arguments that can be passed to '/' endpoints
type Args struct {
	Sort   string
	Order  string
	Offset string
	Limit  string
	Search string
}
