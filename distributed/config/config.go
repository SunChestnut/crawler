package config

const (
	ItemSaverPort = 1234
	WorkerPort0   = 9000
	WorkerPort1   = 9001

	ElasticSearchAddr  = "http://192.168.56.10:9200/"
	ElasticSearchIndex = "dating_profile"

	ItemSaverServiceRpc = "ItemSaverService.Save"
	CrawlServiceRpc     = "CrawlService.Process"

	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ProfileParser"
	NilParser     = "NilParser"

	// Qps => Rate limiting
	Qps = 20
)
