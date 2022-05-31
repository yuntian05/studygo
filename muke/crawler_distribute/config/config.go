package config

const (
	// Parser Name
	ParseCity = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile = "ParseProfile"
	NilParser = "NilParser"

	// Service port
	ItemSaverPort = 1234
	WorkerPort0 = 9002

	// Elasticsearch
	ElasticIndex = "dating_profile"

	// RPC Endpoints
	ItemSaveRPC = "ItemSaverService.Save"
	CrawlServiceRPC = "CrawlerService.Process"
)