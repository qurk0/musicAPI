package song

type SongCreateRequest struct {
	SongName  string `json:"songName" validate:"required"`
	GroupName string `json:"groupName" validate:"required"`
}
