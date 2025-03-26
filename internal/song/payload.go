package song

type SongCreateResponce struct {
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
}

type SongAllResponce struct {
	TotalCount int64          `json:"totalCount"`
	Page       int            `json:"page"`
	Size       int            `json:"size"`
	Songs      []SongResponce `json:"songs"`
}

type SongResponce struct {
	SongName    string `json:"songName"`
	GroupName   string `json:"groupName"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongTextResponce struct {
	TotalCount int64    `json:"totalCount"`
	Page       int      `json:"page"`
	Size       int      `json:"size"`
	Text       []string `json:"text"`
}
