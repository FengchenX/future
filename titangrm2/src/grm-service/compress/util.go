package compress

import "github.com/axgle/mahonia"

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	n, cdata, err := tagCoder.Translate([]byte(srcResult), true)
	if n == 0 || err != nil {
		return src
	}
	return string(cdata)
}
