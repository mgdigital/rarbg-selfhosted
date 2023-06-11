package magnet

import "bytes"

func CreateMagnetLink(hash string, title string, magnet []string) string {
	var link bytes.Buffer
	link.WriteString("magnet:?xt=urn:btih:")
	link.WriteString(hash)
	link.WriteString("&dn=")
	link.WriteString(title)
	for i := range magnet {
		link.WriteString("&tr=")
		link.WriteString(magnet[i])
	}
	return link.String()
}
