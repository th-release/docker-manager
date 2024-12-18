package web

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) *fiber.App {
	api := app.Group("/api", Api)

	container := api.Group("/containers", Container) // 컨테이너 컨트롤러
	container.Get("/", GetAllContainer)              // 모든 컨테이너 정보
	container.Get("/:id", GetByIdContainer)          // 특정 컨테이너 정보
	container.Get("/logs/:id", GetByIdContainerLog)  // 특정 컨테이너 로그
	container.Post("/pause/:id", PauseContainer)     // 특정 컨테이너 일시정지
	container.Post("/unPause/:id", UnPauseContainer) // 특정 컨테이너 일시정지 제거
	container.Post("/start/:id", StartContainer)     // 특정 컨테이너 시작
	container.Post("/restart/:id", RestartContainer) // 특정 컨테이너 재시작
	container.Post("/stop/:id", StopContainer)       // 특정 컨테이너 정지
	container.Post("/kill/:id", KillContainer)       // 특정 컨테이너 킬
	container.Post("/rename/:id", RenameContainer)   // 특정 컨테이너 이름 재설정
	container.Post("/remove/:id", RemoveContainer)   // 특정 컨테이너 삭제
	container.Post("/prune/", PruneContainer)        // 안쓰는 컨테이너 삭제
	// container.Post("/update/:id", UpdateContainer)

	network := api.Group("/networks", Network) // 네트워크 컨트롤러
	network.Get("/", GetAllNetwork)            // 모든 네트워크 정보
	network.Get("/:id", GetByIdNetwork)        // 특정 네트워크 정보
	network.Post("/create", CreateNetwork)     // 네트워크 생성
	network.Post("/remove/:id", RemoveNetwork) // 네트워크 삭제
	network.Post("/connect", ConnectNetwork)   // 특정 컨테이너 특정 네트워크로 연결
	network.Post("/disconnect", DisconnectNetwork)
	network.Post("/prune")
	// api.Get("/network/:id")
	return app
}
