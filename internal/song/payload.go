package song

type SongCreateApiResponce struct {
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
}

type SongCreateResponce struct {
	SongName  string `json:"songName"`
	GroupName string `json:"groupName"`
	Link      string `json:"link"`
}

type SongAllResponce struct {
	TotalCount int64          `json:"totalCount"`
	Page       int            `json:"page"`
	Size       int            `json:"size"`
	Songs      []SongResponce `json:"songs"`
}

type SongResponce struct {
	ID          uint   `json:"ID"`
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

type SongUpdateRequest struct {
	SongName    string `json:"songName"`
	GroupName   string `json:"groupName"`
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
}
