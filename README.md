# 篮球小组赛晋级推演服务

项目提供纯后端 HTTP 接口，采用 Go 和本地 SQLite 保存小组赛数据，默认文件为 `basketball-ranking.db`，可由 `BASKETBALL_DB_PATH` 调整。

初始化：`go run ./cmd/migrate`。启动：`go run .`，健康检查为 `/health`。测试：`go test ./...`。容器运行：`docker build -t basketball-ranking . && docker run --rm -p 8080:8080 basketball-ranking`。
