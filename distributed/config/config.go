package config

const (
	ItemSaverPort      = 1234
	ElasticSearchAddr  = "http://192.168.56.10:9200/"
	ElasticSearchIndex = "dating_profile"

	ItemSaverServiceRpc = "ItemSaverService.Save"
	CrawlServiceRpc     = "CrawlService.Process"

	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ProfileParser"
	NilParser     = "NilParser"

	WorkerPort0 = 9000
)
