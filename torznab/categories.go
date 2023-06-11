package torznab

var categoryMap = map[string]int{
	"ebooks":             7000,
	"games_pc_iso":       4050,
	"games_pc_rip":       4050,
	"games_ps3":          1080,
	"games_ps4":          1180,
	"games_xbox360":      1050,
	"movies":             2000,
	"movies_bd_full":     2060,
	"movies_bd_remux":    2060,
	"movies_x264":        2000,
	"movies_x264_3d":     2050,
	"movies_x264_4k":     2040,
	"movies_x264_720":    2000,
	"movies_x265":        2000,
	"movies_x265_4k":     2040,
	"movies_x265_4k_hdr": 2040,
	"movies_xvid":        2030,
	"movies_xvid_720":    2000,
	"music_flac":         3040,
	"music_mp3":          3010,
	"software_pc_iso":    4020,
	"tv":                 5000,
	"tv_sd":              5030,
	"tv_uhd":             5040,
	"xxx":                6000,
}

// See https://github.com/nZEDb/nZEDb/blob/0.x/docs/newznab_api_specification.txt#L608

func CategoryToId(category string) int {
	for mapCategory, id := range categoryMap {
		if category == mapCategory {
			return id
		}
	}
	return 0
}

func IdToCategories(id int) []string {
	var categories []string
	for category, mapId := range categoryMap {
		if id == mapId {
			categories = append(categories, category)
		}
	}
	return categories
}
