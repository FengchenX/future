package geoserver

import (
//	"encoding/xml"
)

type StyledLayerDescriptor struct {
	Version    string     `xml:"version,attr"  json:",omitempty"`
	NamedLayer NamedLayer `xml:"NamedLayer"  json:",omitempty"`
}

type NamedLayer struct {
	Name      string     `xml:"Name"  json:",omitempty"`
	UserStyle UserStyles `xml:"UserStyle"  json:",omitempty"`
}

type UserStyle struct {
	Name             string            `xml:"Name"  json:",omitempty"`
	Title            string            `xml:"Title"  json:",omitempty"`
	IsDefault        string            `xml:"IsDefault"  json:",omitempty"`
	Abstract         string            `xml:"Abstract"  json:",omitempty"`
	FeatureTypeStyle *FeatureTypeStyle `xml:"FeatureTypeStyle"  json:",omitempty"`

	Pic string `json:"pic,omitempty"`
}

type UserStyles []*UserStyle

// Len()方法和Swap()方法不用变化
// 获取此 slice 的长度
func (s UserStyles) Len() int { return len(s) }

// 交换数据
func (s UserStyles) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (nl UserStyles) Less(i, j int) bool {
	return nl[i].Name > nl[j].Name
}

type FeatureTypeStyle struct {
	Rule []*Rule `xml:"Rule"  json:",omitempty"`
}

type Rule struct {
	Name                string             `xml:"Name"  json:",omitempty"`
	Title               string             `xml:"Title"  json:",omitempty"`
	MinScaleDenominator float64            `xml:"MinScaleDenominator"  json:",omitempty"`
	MaxScaleDenominator float64            `xml:"MaxScaleDenominator"  json:",omitempty"`
	PolygonSymbolizer   *PolygonSymbolizer `xml:"PolygonSymbolizer"  json:",omitempty"`
	PointSymbolizer     *PointSymbolizer   `xml:"PointSymbolizer"  json:",omitempty"`
	TextSymbolizer      *TextSymbolizer    `xml:"TextSymbolizer"  json:",omitempty"`
	LineSymbolizer      *LineSymbolizer    `xml:"LineSymbolizer"  json:",omitempty"`
	Filter              *Filter            `xml:"Filter"  json:",omitempty"`
}

type Filter struct {
	PropertyIsEqualTo              []*PropertyIsEqualTo              `xml:"PropertyIsEqualTo"  json:",omitempty"`
	PropertyIsNotEqualTo           []*PropertyIsNotEqualTo           `xml:"PropertyIsNotEqualTo"  json:",omitempty"`
	PropertyIsLessThan             []*PropertyIsLessThan             `xml:"PropertyIsLessThan"  json:",omitempty"`
	PropertyIsGreaterThan          []*PropertyIsGreaterThan          `xml:"PropertyIsGreaterThan"  json:",omitempty"`
	PropertyIsLessThanOrEqualTo    []*PropertyIsLessThanOrEqualTo    `xml:"PropertyIsLessThanOrEqualTo"  json:",omitempty"`
	PropertyIsGreaterThanOrEqualTo []*PropertyIsGreaterThanOrEqualTo `xml:"PropertyIsGreaterThanOrEqualTo"  json:",omitempty"`
	PropertyIsLike                 []*PropertyIsLike                 `xml:"PropertyIsLike"  json:",omitempty"`
	PropertyIsBetween              []*PropertyIsBetween              `xml:"PropertyIsBetween"  json:",omitempty"`
	And                            *FilterPro                        `xml:"And"  json:",omitempty"`
	Or                             *FilterPro                        `xml:"Or"  json:",omitempty"`
}

type FilterPro struct {
	PropertyIsEqualTo              []*PropertyIsEqualTo              `xml:"PropertyIsEqualTo"  json:",omitempty"`
	PropertyIsNotEqualTo           []*PropertyIsNotEqualTo           `xml:"PropertyIsNotEqualTo"  json:",omitempty"`
	PropertyIsLessThan             []*PropertyIsLessThan             `xml:"PropertyIsLessThan"  json:",omitempty"`
	PropertyIsGreaterThan          []*PropertyIsGreaterThan          `xml:"PropertyIsGreaterThan"  json:",omitempty"`
	PropertyIsLessThanOrEqualTo    []*PropertyIsLessThanOrEqualTo    `xml:"PropertyIsLessThanOrEqualTo"  json:",omitempty"`
	PropertyIsGreaterThanOrEqualTo []*PropertyIsGreaterThanOrEqualTo `xml:"PropertyIsGreaterThanOrEqualTo"  json:",omitempty"`
	PropertyIsLike                 []*PropertyIsLike                 `xml:"PropertyIsLike"  json:",omitempty"`
	PropertyIsBetween              []*PropertyIsBetween              `xml:"PropertyIsBetween"  json:",omitempty"`
}

type PropertyIsBetween struct {
	PropertyName  string         `xml:"PropertyName"  json:",omitempty"`
	LowerBoundary *LowerBoundary `xml:"LowerBoundary"  json:",omitempty"`
	UpperBoundary *UpperBoundary `xml:"UpperBoundary"  json:",omitempty"`
}

type LowerBoundary struct {
	Literal string `xml:"Literal"  json:",omitempty"`
}

type UpperBoundary struct {
	Literal string `xml:"Literal"  json:",omitempty"`
}

type PropertyIsLike struct {
	PropertyName string `xml:"PropertyName"  json:",omitempty"`
	Literal      string `xml:"Literal"  json:",omitempty"`
	WildCard     string `xml:"wildCard,attr"  json:",omitempty"`
	SingleChar   string `xml:"singleChar,attr"  json:",omitempty"`
	Escape       string `xml:"escape,attr"  json:",omitempty"`
}

type PropertyIsGreaterThanOrEqualTo struct {
	PropertyName string `xml:"PropertyName"  json:",omitempty"`
	Literal      string `xml:"Literal"  json:",omitempty"`
}

type PropertyIsLessThanOrEqualTo struct {
	PropertyName string `xml:"PropertyName"  json:",omitempty"`
	Literal      string `xml:"Literal"  json:",omitempty"`
}

type PropertyIsEqualTo struct {
	PropertyName string `xml:"PropertyName" json:",omitempty"`
	Literal      string `xml:"Literal" json:",omitempty"`
}

type PropertyIsNotEqualTo struct {
	PropertyName string `xml:"PropertyName"  json:",omitempty"`
	Literal      string `xml:"Literal"  json:",omitempty"`
}

type PropertyIsLessThan struct {
	PropertyName string `xml:"PropertyName"  json:",omitempty"`
	Literal      string `xml:"Literal"  json:",omitempty"`
}

type PropertyIsGreaterThan struct {
	PropertyName string `xml:"PropertyName"  json:",omitempty"`
	Literal      string `xml:"Literal"  json:",omitempty"`
}

type LineSymbolizer struct {
	Stroke Stroke `xml:"Stroke"  json:",omitempty"`
}

type TextSymbolizer struct {
	Label          *Label          `xml:"Label"  json:",omitempty"`
	Font           *Font           `xml:"Font"  json:",omitempty"`
	LabelPlacement *LabelPlacement `xml:"LabelPlacement"  json:",omitempty"`
	Fill           *Fill           `xml:"Fill"  json:",omitempty"`
	Halo           *Halo           `xml:"Halo"  json:",omitempty"`
}

type Halo struct {
	Radius float32 `xml:"Radius"  json:",omitempty"`
	Fill   *Fill   `xml:"Fill"  json:",omitempty"`
}

type Label struct {
	PropertyName string `xml:"PropertyName"  json:",omitempty"`
}

type Font struct {
	CssParameter []*CssParameter `xml:"CssParameter"  json:",omitempty"`
}

type LabelPlacement struct {
	PointPlacement *PointPlacement `xml:"PointPlacement"  json:",omitempty"`
}

type PointPlacement struct {
	AnchorPoint *AnchorPoint `xml:"AnchorPoint"  json:",omitempty"`
}

type AnchorPoint struct {
	AnchorPointX float32 `xml:"AnchorPointX"  json:",omitempty"`
	AnchorPointY float32 `xml:"AnchorPointY"  json:",omitempty"`
}

type PointSymbolizer struct {
	Graphic *Graphic `xml:"Graphic"  json:",omitempty"`
}

type Graphic struct {
	Mark     *Mark   `xml:"Mark"  json:",omitempty"`
	Size     string  `xml:"Size"  json:",omitempty"`
	Rotation float32 `xml:"Rotation"  json:",omitempty"`
}

type Mark struct {
	WellKnownName string  `xml:"WellKnownName"  json:",omitempty"`
	Fill          *Fill   `xml:"Fill"  json:",omitempty"`
	Stroke        *Stroke `xml:"Stroke"  json:",omitempty"`
}

type PolygonSymbolizer struct {
	Fill   *Fill   `xml:"Fill"  json:",omitempty"`
	Stroke *Stroke `xml:"Stroke"  json:",omitempty"`
}

type Fill struct {
	CssParameter []*CssParameter `xml:"CssParameter"  json:",omitempty"`
}

type CssParameter struct {
	Name  string `xml:"name,attr"  json:",omitempty"`
	Value string `xml:",chardata"  json:",omitempty"`
}

type Stroke struct {
	CssParameter []*CssParameter `xml:"CssParameter"  json:",omitempty"`
}
