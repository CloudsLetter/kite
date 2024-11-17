package utilities

func GetUniqueFilename(ext string) string {
	var filename string
	for {
		rdk := RandomString(20)
		filename = rdk + "." + ext
		if !IfileExist(filename) {
			break
		}
	}
	return filename
}
