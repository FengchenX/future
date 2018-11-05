package common

import (
	"regexp"
)

const (
	Namespace = "TitanGRM"

	FilePre   = "File:"
	NginxPre  = "Nginx:"
	GisPre    = "Gis:"
	RasterPre = "Raster:"
	OfficePre = "Office:"

	FileUpload   = "file"
	Base64Upload = "base64"

	OmitArg = "omit"

	DataLoading   = "loading"
	DataNormal    = "normal"
	DataObsoleted = "obsoleted"
)

const (
	DBType  = "DB"  // 数据库
	NFSType = "NFS" // 网络文件系统
	DFSType = "DFS" // 分布式文件系统
)

const (
	POSTGRESQL = "PostgreSQL" // pg数据库
	MONGODB    = "MongoDB"    // MongoDB
)

const bitsPerWord = 32 << uint(^uint(0)>>63)

// Implementation-specific size of int and uint in bits.
const BitsPerWord = bitsPerWord // either 32 or 64

// Implementation-specific integer limit values.
const (
	MaxInt  = 1<<(BitsPerWord-1) - 1 // either 1<<31 - 1 or 1<<63 - 1
	MinInt  = -MaxInt - 1            // either -1 << 31 or -1 << 63
	MaxUint = 1<<BitsPerWord - 1     // either 1<<32 - 1 or 1<<64 - 1
)

var (
	// enforceIndexRegex is a regular expression which extracts the enforcement error
	EnforceIndexRegex = regexp.MustCompile(`\((Enforcing job modify index.*)\)`)

	TileBlankPNG = []byte{137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0,
		0, 1, 0, 0, 0, 1, 0, 1, 3, 0, 0, 0, 102, 188, 58, 37, 0, 0, 0, 3, 80, 76, 84, 69, 0, 0, 0, 167, 122, 61,
		218, 0, 0, 0, 1, 116, 82, 78, 83, 0, 64, 230, 216, 102, 0, 0, 0, 31, 73, 68, 65, 84, 104, 222, 237, 193,
		1, 13, 0, 0, 0, 194, 32, 251, 167, 54, 199, 55, 96, 0, 0, 0, 0, 0, 0, 0, 0, 113, 7, 33, 0, 0, 1, 167, 87,
		41, 215, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}
)
