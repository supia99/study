package model

import "fmt"

type Song struct {

	// INDEX
	Index int
	// 曲名
	Name string
	// 作曲者
	Composer string
	// 作詞者
	Lyrist string
	// 編曲者
	Arranger string
	// 歌手
	Singer string
}

type Disk struct {

	// グループ
	Group string
	// シリーズ
	Series string
	// シリーズ内番号
	SeriesNo int
	// 曲
	Songs []Song
}

// TEST用　出力する
func (disk Disk) Print() {
	fmt.Println(
		"\nGroup:", disk.Group,
		"\nSeries:", disk.Series,
		"\nSeriesNo:", disk.SeriesNo,
	)
	for _, song := range disk.Songs {
		fmt.Println(
			"\n index:", song.Index,
			"\n  Name:", song.Name,
			"\n  Composer:", song.Composer,
			"\n  Lyrist:", song.Lyrist,
			"\n  Arranger:", song.Arranger,
			"\n  Singer:", song.Singer,
		)
	}
}

func (disk Disk) PrintExcel() {
	for _, song := range disk.Songs {
		fmt.Println(disk.Series, "| ",
			// song.Index, ", ",
			song.Name, "| ",
			song.Composer, "| ",
			song.Lyrist, "| ",
			song.Arranger, "| ",
			song.Singer)
	}
}
