# 篮球小组赛晋级推演服务

项目提供纯后端 HTTP 接口，采用 Go 和本地 SQLite 保存巡演数据，默认文件为 `stage-readiness.db`，可由 `STAGE_DB_PATH` 调整。

初始化：`go run ./cmd/migrate`。启动：`go run .`，健康检查为 `/health`。测试：`go test ./...`。容器运行：`docker build -t stage-readiness . && docker run --rm -p 8080:8080 stage-readiness`。
