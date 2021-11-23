package domain

// 質問
type Question struct {
	*Model
	Title        string
	FirstAnswer  string
	SecondAnswer string
	FirstCount   uint64
	SecondCount  uint64
	FirstImgURL  string
	SecondImgURL string
}
