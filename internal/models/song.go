package models

type Song struct {
	ID          uint   `json:"id"`
	SongName    string `gorm:"not null" json:"songName"`
	Source      string `gorm:"not null" json:"source"`
	Singer      string `gorm:"not null" json:"singer"`
	SingerPhoto string `json:"singerPhoto"`
	Lyrics      string `json:"lyrics"`
}
