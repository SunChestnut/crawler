# 1⃣ 运行爬虫的 Mock Server
StartMockService:
	go run mockserver/main.go

# 2⃣
StartItemSaverService:
	go run persist/server/main.go -port 8800

# 3⃣
StartWorkerService_01:
	go run worker/server/main.go -port 9100

StartWorkerService_02:
	go run  worker/server/main.go -port 9101

StartWorkerService_03:
	go run  worker/server/main.go -port 9102

# 4⃣
StartMain:
	 go run main.go &

# 5⃣
StartStaticPage:
	go run frontend/starter.go


gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        proto/*.proto


#StartAll:
#	StartItemSaverService
#	StartWorkerService_01
#	StartWorkerService_02
#	StartWorkerService_03
#	StartMain
