package main

import (
	"cth-core.xyz/docker-manager/web"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	web.SetupRoute(app)

	// port, err := utils.GetEnv("PORT")

	// if err != nil {
	// 	log.Print("ENV: PORT를 불러올 수 없습니다.")
	// 	os.Exit(1)
	// }

	// if len(port) <= 0 {
	// 	log.Print("웹 서버 실행 오류: PORT 변수가 잘못 입력되었습니다.")
	// 	os.Exit(1)
	// }
	app.Listen(":" + "8081")
}
