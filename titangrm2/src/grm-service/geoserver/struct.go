package geoserver

import (
//	"encoding/json"
)

const (
	GeoWorkSpace = "titangrm"
)

type Workspaces struct {
	Workspaces *NameObject `json:"workspace"`
}

type DataStores struct {
	DataStore *DataStore `json:"dataStore"`
}

type DataStore struct {
	Name                 string                `json:"name"`
	ConnectionParameters *ConnectionParameters `json:"connectionParameters"`
}

type ConnectionParameters struct {
	Entrys []*Entry `json:"entry"`
}

type Entry struct {
	Key   string `json:"@key"`
	Value string `json:"$"`
}

type FeatureTypeJson struct {
	FeatureType *FeatureType `json:"featureType"`
}

type FeatureType struct {
	Name              string      `json:"name,omitempty"`
	NativeName        string      `json:"nativeName,omitempty"`
	Srs               string      `json:"srs,omitempty"`
	NativeBoundingBox *NativeBBox `json:"nativeBoundingBox,omitempty"`
	LatLonBoundingBox *WGS84BBox  `json:"latLonBoundingBox,omitempty"`
}

type WGS84BBox struct {
	Minx float64 `json:"minx,omitempty"`
	Maxx float64 `json:"maxx,omitempty"`
	Miny float64 `json:"miny,omitempty"`
	Maxy float64 `json:"maxy,omitempty"`
	//	Crs  *Crs    `json:"crs,omitempty"`
}

type Crs struct {
	Class string `json:"@class,omitempty"`
	Value string `json:"$,omitempty"`
}

type NativeBBox struct {
	Minx float64 `json:"minx,omitempty"`
	Maxx float64 `json:"maxx,omitempty"`
	Miny float64 `json:"miny,omitempty"`
	Maxy float64 `json:"maxy,omitempty"`
	//	Crs  string  `json:"crs,omitempty"`
}

type CoverageStores struct {
	CoverageStore *CoverageStore `json:"coverageStore"`
}

type CoverageStore struct {
	Name       string      `json:"name"`
	Type       string      `json:"type"`
	Url        string      `json:"url"`
	Enabled    bool        `json:"enabled"`
	Workspaces *NameObject `json:"workspace"`
}

type CoverageJson struct {
	Coverage *CoverageInfo `json:"coverage"`
}

type CoverageInfo struct {
	Name              string      `json:"name"`
	NativeName        string      `json:"nativeName"`
	Srs               string      `json:"srs,omitempty"`
	NativeBoundingBox *NativeBBox `json:"nativeBoundingBox,omitempty"`
	LatLonBoundingBox *WGS84BBox  `json:"latLonBoundingBox,omitempty"`
}

type NameObject struct {
	Name string `json:"name"`
}

type Layer struct {
	Name         string        `json:"name,omitempty"`
	DefaultStyle *NameObject   `json:"defaultStyle,omitempty"`
	Resource     *Resource     `json:"resource,omitempty"`
	Styles       *StylesStruct `json:"styles,omitempty"`
}

type Resource struct {
	Class string `json:"@class,omitempty"`
	Name  string `json:"name,omitempty"`
	Href  string `json:"href,omitempty"`
}

type LayerJson struct {
	Layer *Layer `json:"layer"`
}

type Style struct {
	Name string `json:"name,omitempty"`
	Href string `json:"href,omitempty"`
	Pic  string `json:"pic,omitempty"`

	Title string `json:"title,omitempty"`
	Sld   string `json:"sld,omitempty"`
}

type StylesStruct struct {
	Styles []*Style `json:"style"`
}

type StylesJson struct {
	Styles *StylesStruct `json:"styles"`
}
