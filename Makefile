# 1⃣ 运行爬虫的 Mock Server
StartMockService:
	go run mockserver/main.go

# 2⃣
StartItemSaverService:
	go run distributed/persist/server/itemsaver.go

# 3⃣
StartMultiWorkerService_01:
	 go run distributed/worker/server/worker.go -port 9100

StartMultiWorkerService_02:
	go run distributed/worker/server/worker.go -port 9101

StartMultiWorkerService_03:
	go run distributed/worker/server/worker.go -port 9102

# 4⃣
StartMain:
	 go run distributed/main.go

# 5⃣
StartStaticPage:
	go run frontend/starter.go


gen:
	 protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        proto/*.proto