package config

const (
	ItemSaverPort = 1234
	WorkerPort0   = 9100 // worker 端口-0
	WorkerPort1   = 9101 // worker 端口-1
	WorkerPort2   = 9102 // worker 端口-2

	ElasticSearchAddr  = "http://127.0.0.1:9200/"
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
