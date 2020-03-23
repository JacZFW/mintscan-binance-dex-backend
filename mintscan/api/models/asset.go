package models

import (
	"encoding/json"
	"time"
)

// ChosenAssetNames are chosen four asset names for the card view on Asset page
var ChosenAssetNames = []string{
	"TUSDB-888",
	"USDSB-1AC",
	"BTCB-1DE",
	"IRIS-D88",
}

// Asset represents asset detail information
type Asset struct {
	CreateTime      interface{} `json:"createTime"`
	UpdateTime      interface{} `json:"updateTime"`
	ID              int         `json:"id"`
	Asset           string      `json:"asset"`
	MappedAsset     string      `json:"mappedAsset"`
	Name            string      `json:"name"`
	AssetImg        string      `json:"assetImg"`
	Supply          float64     `json:"supply"`
	Price           float64     `json:"price"`
	QuoteUnit       string      `json:"quoteUnit"`
	ChangeRange     float64     `json:"changeRange"`
	Owner           string      `json:"owner"`
	Mintable        int         `json:"mintable"`
	Visible         interface{} `json:"visible"`
	Description     string      `json:"description"`
	AssetCreateTime interface{} `json:"assetCreateTime"`
	Transactions    int         `json:"transactions"`
	Holders         int         `json:"holders"`
	OfficialSiteURL string      `json:"officialSiteUrl"`
	ContactEmail    string      `json:"contactEmail"`
	MediaList       []struct {
		MediaName string `json:"mediaName"`
		MediaURL  string `json:"mediaUrl"`
		MediaImg  string `json:"mediaImg"`
	} `json:"mediaList"`
}

// AssetInfo represents asset information list
type AssetInfo struct {
	TotalNum      int             `json:"totalNum"`
	AssetInfoList []AssetInfoList `json:"assetInfoList"`
}

// AssetInfoList wraps asset information list
type AssetInfoList struct {
	CreateTime      json.RawMessage `json:"createTime"`
	UpdateTime      json.RawMessage `json:"updateTime"`
	ID              int32           `json:"id"`
	Asset           string          `json:"asset"`
	MappedAsset     string          `json:"mappedAsset"`
	Name            string          `json:"name"`
	AssetImg        string          `json:"assetImg"`
	Supply          float64         `json:"supply"`
	Price           float64         `json:"price"`
	QuoteUnit       string          `json:"quoteUnit"`
	ChangeRange     float64         `json:"changeRange"`
	Owner           string          `json:"owner"`
	Mintable        int32           `json:"mintable"`
	Visible         json.RawMessage `json:"visible"`
	Description     json.RawMessage `json:"description"`
	AssetCreateTime int64           `json:"assetCreateTime"`
}

// ResultAssetInfo represents asset information list
type ResultAssetInfo struct {
	TotalNum      int                   `json:"totalNum"`
	AssetInfoList []ResultAssetInfoList `json:"assetInfoList"`
}

// ResultAssetInfoList wraps result asset information list
type ResultAssetInfoList struct {
	Asset       string  `json:"asset"`
	MappedAsset string  `json:"mappedAsset"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	QuoteUnit   string  `json:"quoteUnit"`
}

// AssetHolders represents asset holders list
type AssetHolders struct {
	TotalNum       int `json:"totalNum"`
	AddressHolders []struct {
		Address    string      `json:"address"`
		Quantity   float64     `json:"quantity"`
		Percentage float64     `json:"percentage"`
		Tag        interface{} `json:"tag"`
	} `json:"addressHolders"`
}

// ResultAssetsImages represents assets images
type ResultAssetsImages struct {
	TotalNum  int         `json:"totalNum"`
	ImageList []ImageList `json:"imageList"`
}

// ImageList wraps asset list
type ImageList struct {
	Asset      string `json:"asset"`
	Name       string `json:"name"`
	AssetImage string `json:"assetImg"`
}

// AssetTxs represents asset transactions
type AssetTxs struct {
	TxNums  int `json:"txNums"`
	TxArray []struct {
		TxHash        string  `json:"txHash"`
		BlockHeight   int64   `json:"blockHeight"`
		TxType        string  `json:"txType"`
		TimeStamp     int64   `json:"timeStamp"`
		FromAddr      string  `json:"fromAddr"`
		ToAddr        string  `json:"toAddr"`
		Value         float64 `json:"value"`
		TxAsset       string  `json:"txAsset"`
		TxQuoteAsset  string  `json:"txQuoteAsset"`
		TxFee         float64 `json:"txFee"`
		TxAge         int64   `json:"txAge"`
		OrderID       string  `json:"orderId"`
		Data          string  `json:"data,omitempty"`
		Code          int64   `json:"code"`
		Log           string  `json:"log"`
		ConfirmBlocks int64   `json:"confirmBlocks"`
		Memo          string  `json:"memo"`
		Source        int64   `json:"source"`
		HasChildren   int64   `json:"hasChildren"`
	} `json:"txArray"`
}

// ResultAssetTxs represents response data for AssetTxs
type ResultAssetTxs struct {
	TxNums  int            `json:"txNums"`
	TxArray []AssetTxArray `json:"txArray"`
}

// AssetTxArray wraps ResultAssetTxs TxArray
type AssetTxArray struct {
	BlockHeight   int64        `json:"blockHeight"`
	Code          int64        `json:"code"`
	TxHash        string       `json:"txHash"`
	TxType        string       `json:"txType"`
	TxAsset       string       `json:"txAsset"`
	TxQuoteAsset  string       `json:"txQuoteAsset,omitempty"`
	Value         float64      `json:"value"`
	TxFee         float64      `json:"txFee"`
	FromAddr      string       `json:"fromAddr"`
	ToAddr        string       `json:"toAddr,omitempty"`
	TxAge         int64        `json:"txAge"`
	OrderID       string       `json:"orderId,omitempty"`
	Message       *AssetTxData `json:"message,omitempty"`
	Log           string       `json:"log"`
	ConfirmBlocks int64        `json:"confirmBlocks"`
	Memo          string       `json:"memo"`
	Source        int64        `json:"source"`
	HasChildren   int64        `json:"hasChildren,omitempty"`
	Timestamp     int64        `json:"timeStamp"`
}

// AssetTxData wraps ResultAssetTxs Data
type AssetTxData struct {
	OrderData struct {
		Symbol      string `json:"symbol"`
		OrderType   string `json:"orderType"`
		Side        string `json:"side"`
		Price       string `json:"price"`
		Quantity    string `json:"quantity"`
		TimeInForce string `json:"timeInForce"`
		OrderID     string `json:"orderId"`
	} `json:"orderData"`
}

// AssetChartHistory represents asset chart hisotry
type AssetChartHistory struct {
	Name         string   `json:"name"`
	Asset        string   `json:"asset"`
	MappedAsset  string   `json:"mapped_asset"`
	CurrentPrice float64  `json:"current_price"`
	QuoteUnit    string   `json:"quote_unit"`
	ChangeRange  float64  `json:"change_range"`
	Supply       float64  `json:"supply"`
	Marketcap    float64  `json:"marketcap"`
	AssetImage   string   `json:"asset_img"`
	Prices       []Prices `json:"prices"`
}

// Price wraps price
type Prices struct {
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}
