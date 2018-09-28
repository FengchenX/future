//author xinbing
//time 2018/8/27 14:29
package eft_encrypt

var step byte = 13
func byteEncrypt(str string) string{
	originBytes := []byte(str)
	newBytes := make([]byte,0)
	modeNum := int(originBytes[0] - 65)
	if modeNum < 13 {
		modeNum = modeNum + 13
	}
	for i,b := range originBytes {
		if i % modeNum == 0 {
			rb,times := transfer(b)
			newBytes = append(newBytes,rb,times)
		} else {
			newBytes = append(newBytes,b)
		}
	}
	return string(newBytes)
}

//0~9 48 - 57
//A~Z 65 - 90
//a~z 97 - 122
func transfer(b byte) (byte,byte){
	count := 0
	for !normalByte(b) {
		b = b + step
		if b >= 128 {
			b = b - 128
		}
		count ++
	}
	return b, 65 + byte(count)
}

func restore(b byte,times byte) byte {
	count := byte(times) - 65
	for count > 0 {
		b = b - step
		if b < 0 {
			b += 128
		}
		count--
	}
	return b
}

func normalByte(b byte) bool {
	//return (b >= 97 && b <= 122) || (b >= 48 && b <= 57) || (b >=65 && b <= 90)
	return b >= 97 && b <= 122
}

func byteDecrypt(str string) string {
	originBytes := []byte(str)
	newBytes := make([]byte,0)
	modeNum := int(restore(originBytes[0],originBytes[1]) - 65)
	if modeNum < 13 {
		modeNum = modeNum + 13
	}
	count := 0
	for i:=0; i< len(originBytes); {
		if i == (modeNum + 1) * count {
			newByte := restore(originBytes[i],originBytes[i + 1])
			newBytes = append(newBytes, newByte)
			count++
			i += 2
		} else {
			newBytes = append(newBytes,originBytes[i])
			i++
		}
	}
	return string(newBytes)
}
