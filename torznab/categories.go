package torznab

var categoryMap = map[string]int{
	"ebooks":             7000,
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
	"tv":                 5000,
	"tv_sd":              5030,
	"tv_uhd":             5040,
}

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

//ebooks
//games_pc_iso
//games_pc_rip
//games_ps3
//games_ps4
//games_xbox360
//movies
//movies_bd_full
//movies_bd_remux
//movies_x264
//movies_x264_3d
//movies_x264_4k
//movies_x264_720
//movies_x265
//movies_x265_4k
//movies_x265_4k_hdr
//movies_xvid
//movies_xvid_720
//music_flac
//music_mp3
//software_pc_iso
//tv
//tv_sd
//tv_uhd
//xxx

//Categories                           Category Name
//
//0000                                 Other                           All of Other
//0010                                 Other/Misc                      Anything that could not get categorized
//0020                                 Other/Hashed                    Anything with a hashed name
//1000                                 Console                         All of console
//1010                                 Console/NDS                     Nintendo DS
//1020                                 Console/PSP                     Sony Playstation Portable
//1030                                 Console/Wii                     Nintendo Wii
//1040                                 Console/Xbox                    Microsoft XBox
//1050                                 Console/Xbox 360                Microsoft XBox 360
//1060                                 Console/Wiiware/VC              Wii homebrew
//1070                                 Console/XBOX 360 DLC            Microsoft XBox 360 Downloadable Content
//1080                                 Console/PS3                     Playstation 3
//1999                                 Console/Other                   Misc Console
//1110                                 Console/3DS                     Nintendo 3DS
//1120                                 Console/PS Vita                 Playstation Vita
//1130                                 Console/WiiU                    Nintento Wii U
//1140                                 Console/Xbox One                Xbox One
//1180                                 Console/PS4                     Playstation 4
//2000                                 Movies                          All of movies
//2010                                 Movies/Foreign                  Non english movies
//2020                                 Movies/Other                    Misc movies
//2030                                 Movies/SD                       Standard definition movies
//2040                                 Movies/HD                       High definition movies (720p+)
//2050                                 Movies/3D                       3D movies
//2060                                 Movies/BluRay                   Full BR movies
//2070                                 Movies/DVD                      Full DVD movies
//2080                                 Movies/WEBDL                    WEB-DL movies
//3000                                 Audio                           All of audio
//3010                                 Audio/MP3                       Mp3 music
//3020                                 Audio/Video                     Music videos
//3030                                 Audio/Audiobook                 Books in audio format
//3040                                 Audio/Lossless                  Lossless music
//3999                                 Audio/Other                     Misc music
//3060                                 Audio/Foreign                   Non english music.
//4000                                 PC                              All of PC
//4010                                 PC/0day                         Apps and games not released in ISO.
//4020                                 PC/ISO                          CD-ROM images/DVD Images
//4030                                 PC/Mac                          OS X apps and games
//4040                                 PC/Phone-Other                  Misc mobile phone software
//4050                                 PC/Games                        PC Games
//4060                                 PC/Phone-IOS                    IOS apps
//4070                                 PC/Phone-Android                Android apps
//5000                                 TV                              All of TV
//5010                                 TV/WEB-DL                       WEB-DL TV
//5020                                 TV/FOREIGN                      FOREIGN TV
//5030                                 TV/SD                           SD TV
//5040                                 TV/HD                           HD TV
//5999                                 TV/OTHER                        Other TV Content
//5060                                 TV/Sport                        Sports
//5070                                 TV/Anime                        Anime
//5080                                 TV/Documentary                  Documentaries
//6000                                 XXX                             All of XXX
//6010                                 XXX/DVD                         Full DVD's
//6020                                 XXX/WMV                         WMV rips
//6030                                 XXX/XviD                        dvdrips
//6040                                 XXX/x264                        HD Porn
//6999                                 XXX/Other                       Misc Porn
//6060                                 XXX/Imageset                    Sets of porn images
//6070                                 XXX/Packs                       Packs of multiple porn videos
//7000                                 Books                           All of Books
//7010                                 Books/Magazines                 Magazines
//7020                                 Books/Ebook                     Ebooks
//7030                                 Books/Comics                    Comics Ebooks
//7040                                 Books/Technical                 Technical books
//7060                                 Books/Foreign                   Non english books
//7999                                 Books/Unknown                   Misc books
//100000+                              Custom                          Specific to a site
