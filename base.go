package duguang

import "errors"

var (
	ErrNoImg     = errors.New("no image")
	ErrImgRepeat = errors.New("image repeat")
	ErrSize      = errors.New("image base64 size too large")
)

const IMGSIZE = 4 * 1024 * 1024

type Advanced struct {
	// 图像数据：base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px，支持jpg/png/bmp格式，和url参数只能同时存在一个
	Img string `json:"img,omitempty"`
	// 图像url地址：图片完整URL，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px，支持jpg/png/bmp格式，和img参数只能同时存在一个
	URL string `json:"url,omitempty"`
	// 是否需要识别结果中每一行的置信度，默认不需要。 true：需要 false：不需要
	Prob bool `json:"prob"`
	// 是否需要单字识别功能，默认不需要。 true：需要 false：不需要
	CharInfo bool `json:"charInfo"`
	// 是否需要自动旋转功能，默认不需要。 true：需要 false：不需要
	Rotate bool `json:"rotate"`
	// 是否需要表格识别功能，默认不需要。 true：需要 false：不需要
	Table bool `json:"table"`
	// 字块返回顺序，false表示从左往右，从上到下的顺序，true表示从上到下，从左往右的顺序，默认false
	SortPage bool `json:"sortPage"`
}

type Document struct {
	// 图像数据：base64编码，要求base64编码后大小不超过4M，最短边至少15px，最长边最大4096px，支持jpg/png/bmp格式，和url参数只能同时存在一个
	Img string `json:"img,omitempty"`
	// 图像url地址：图片完整URL，URL长度不超过1024字节，URL对应的图片base64编码后大小不超过4M，最短边至少15px，最长边最大4096px，支持jpg/png/bmp格式，和img参数只能同时存在一个
	URL string `json:"url,omitempty"`
	// 是否需要识别结果中每一行的置信度，默认不需要。 true：需要 false：不需要
	Prob bool `json:"prob"`
	// 是否需要单字识别功能，默认不需要。 true：需要 false：不需要
	CharInfo bool `json:"charInfo"`
	// 是否需要自动旋转功能，默认不需要。 true：需要 false：不需要
	Rotate bool `json:"rotate"`
	// 是否需要表格识别功能，默认不需要。 true：需要 false：不需要
	Table bool `json:"table"`
	// 是否需要分页功能，默认不需要。 true：需要 false：不需要
	Page bool `json:"page"`
	// 是否需要分段功能，默认不需要。 true：需要 false：不需要
	Paragraph bool `json:"paragraph"`
	// 是否需要成行功能，默认不需要。 true：需要 false：不需要
	Row bool `json:"row"`
	// 是否需要切边功能，默认不需要。 true：需要 false：不需要
	RemoveBoundary bool `json:"removeBoundary"`
	// 是否需要去印章功能，默认不需要。 true：需要 false：不需要
	NoStamp bool `json:"noStamp"`
	// 字块返回顺序，false表示从左往右，从上到下的顺序，true表示从上到下，从左往右的顺序，默认false
	SortPage bool `json:"sortPage"`
}

type Result struct {
	// 唯一id，用于问题定位
	SID string `json:"sid"`

	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`

	// 算法版本
	PrismVersion string `json:"prism_version"`
	// 识别的文字块的数量，prism_wordsInfo数组大小
	PrismWnum int `json:"prism_wnum"`
	// 范围：0-360,0表示向上，90表示向右，180表示向下，270度表示向左
	Angle int `json:"angle"`
	// 识别的文字的具体内容
	PrismWordsInfo []WordInfo `json:"prism_wordsInfo"`
	// 表格信息，如果不存在表格，则改字段内容为空
	PrismTablesInfo []TableInfo `json:"prism_tablesInfo"`
}

// 识别的文字的具体内容
type WordInfo struct {
	// 文字块
	Word string `json:"word"`
	// 置信度
	Prob int `json:"prob"`
	// 文字块的位置，按照文字块四个角的坐标顺时针排列，分别为左上XY坐标、右上XY坐标、右下XY坐标、左下XY坐标
	Pos []Pos `json:"pos"`
	// 单字信息
	CharInfo []CharInfo `json:"charInfo"`
	// 如果该文字块在表格内则存在该字段，tableId表示表格的id
	TableID int `json:"tableId"`
	// 如果该文字块在表格内则存在该字段，表示表格中单元格的id
	TableCellID int `json:"tableCellId"`
	// 行id
	RowID int `json:"rowId"`
	// 段id
	ParagraphID int `json:"paragraphId"`
	// 页id
	PageID int `json:"pageId"`
}

type Pos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// 单字信息
type CharInfo struct {
	// 单字文字
	Word string `json:"word"`
	// 单字置信度
	Prob int `json:"prob"`
	// 单字左上角横坐标
	X int `json:"x"`
	// 单字左上角纵坐标
	Y int `json:"y"`
	// 单字宽度
	W int `json:"w"`
	// 单字长度
	H int `json:"h"`
}

// 表格信息
type TableInfo struct {
	// 表格id，和prism_wordsInfo信息中的tableId对应
	TableID int `json:"tableId"`
	// 表格中横坐标单元格的数量
	XCellSize int `json:"xCellSize"`
	// 表格中纵坐标单元格的数量
	YCellSize int `json:"yCellSize"`
	// 单元格信息，包含单元格在整个表格中的空间拓扑关系
	CellInfos []CellInfo `json:"cellInfos"`
}

// 单元格信息，包含单元格在整个表格中的空间拓扑关系
type CellInfo struct {
	// 表格中单元格id，和prism_wordsInfo信息中的tableCellId对应
	TableCellID int `json:"tableCellId"`
	// 单元格中的文字
	Word string `json:"word"`
	// xStartCell缩写，表示横轴方向该单元格起始在第几个单元格，第一个单元格值为0
	Xsc int `json:"xsc"`
	// xEndCell缩写，表示横轴方向该单元格结束在第几个单元格，第一个单元格值为0，如果xsc和xec都为0说明该文字在横轴方向占据了一个单元格并且在第一个单元格内
	Xec int `json:"xec"`
	// yStartCell缩写，表示纵轴方向该单元格起始在第几个单元格，第一个单元格值为0
	Ysc int `json:"ysc"`
	// yEndCell缩写，表示纵轴方向该单元格结束在第几个单元格，第一个单元格值为0
	Yec int `json:"yec"`
	// 单元格位置，按照单元格四个角的坐标顺时针排列，分别为左上XY坐标、右上XY坐标、右下XY坐标、左下XY坐标
	Pos []Pos `json:"pos"`
}

// 分页信息
type PageInfo struct {
	// 页id，和prism_wordsInfo信息中的pageId对应
	PageID int `json:"pageId"`
	// 文字内容
	Word string `json:"word"`
}

// 分段信息
type ParagraphInfo struct {
	// 段id，和prism_wordsInfo信息中的paragraphId对应
	ParagraphID int `json:"paragraphId"`
	// 文字内容
	Word string `json:"word"`
}

// 成行信息
type RowsInfo struct {
	// 行id，和prism_wordsInfo信息中的rowId对应
	RowID int `json:"rowId"`
	// 文字内容
	Word string `json:"word"`
}
