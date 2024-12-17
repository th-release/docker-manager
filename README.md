# Docker-Manager

Web Docker 관리 프로젝트

### 지원할 대표 Routes
* /containers
* /networks
* /images
* /volumes

### /containers
	container.Get("/", GetAllContainer)
	container.Get("/:id", GetByIdContainer)
	container.Get("/logs/:id", GetByIdContainerLog)
	container.Post("/pause/:id", PauseContainer)
	container.Post("/unPause/:id", UnPauseContainer)
	container.Post("/start/:id", StartContainer)
	container.Post("/restart/:id", RestartContainer)
	container.Post("/stop/:id", StopContainer)
	container.Post("/kill/:id", KillContainer)
	container.Post("/rename/:id", RenameContainer)
	container.Post("/remove/:id", RemoveContainer)
	container.Post("/prune/", PruneContainer)

### /networks

### /images

### ...