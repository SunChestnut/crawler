# 1⃣ 运行爬虫的 Mock Server
startMockService:
	go run mockserver/main.go

# 2⃣
startItemSaverService:
	go run distributed/persist/server/itemsaver.go

# 3⃣
startMultiWorkerService_01:
	 go run distributed/worker/server/worker.go -port 9100

startMultiWorkerService_02:
	go run distributed/worker/server/worker.go -port 9101

startMultiWorkerService_03:
	go run distributed/worker/server/worker.go -port 9102

# 4⃣
startMain:
	 go run distributed/main.go

# 5⃣
startStaticPage:
	go run frontend/starter.go

