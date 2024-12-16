package web

import (
	"bytes"
	"context"
	"io"
	"strings"

	"cth-core.xyz/docker-manager/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ContainerController struct {
}

func GetAllContainer(c *fiber.Ctx) error {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Container Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ArrayBasicResponse[types.Container]{
		Success: true,
		Message: "",
		Data:    containers,
	})
}

func GetByIdContainer(c *fiber.Ctx) error {
	id := c.Params("id")

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	container, err := cli.ContainerInspect(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Container Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[types.ContainerJSON]{
		Success: true,
		Message: "",
		Data:    container,
	})
}

func GetByIdContainerLog(c *fiber.Ctx) error {
	id := c.Params("id")

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	log, err := cli.ContainerLogs(context.Background(), id, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     false,
		Timestamps: true,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Container Connect Error:" + err.Error(),
			Data:    nil,
		})
	}
	defer log.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, log)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Container Logs IO Error:" + err.Error(),
			Data:    nil,
		})
	}

	lines := strings.Split(buf.String(), "\n")
	start := 0
	if len(lines) > 300 {
		start = len(lines) - 300
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[string]{
		Success: true,
		Message: "",
		Data:    strings.Join(lines[start:], "\n"),
	})
}

func PauseContainer(c *fiber.Ctx) error {
	id := c.Params("id")

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	err = cli.ContainerPause(context.Background(), id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Pause Access Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
		Success: true,
		Data:    nil,
	})
}

func UnPauseContainer(c *fiber.Ctx) error {
	id := c.Params("id")

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	err = cli.ContainerUnpause(context.Background(), id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker UnPause Access Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
		Success: true,
		Data:    nil,
	})
}

func StartContainer(c *fiber.Ctx) error {
	id := c.Params("id")

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	err = cli.ContainerStart(context.Background(), id, container.StartOptions{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Start Access Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
		Success: true,
		Data:    nil,
	})
}

func RestartContainer(c *fiber.Ctx) error {
	id := c.Params("id")

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	err = cli.ContainerRestart(context.Background(), id, container.StopOptions{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Restart Access Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
		Success: true,
		Data:    nil,
	})
}

func StopContainer(c *fiber.Ctx) error {
	id := c.Params("id")

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	err = cli.ContainerStop(context.Background(), id, container.StopOptions{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Stop Access Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
		Success: true,
		Data:    nil,
	})
}

func KillContainer(c *fiber.Ctx) error {
	id := c.Params("id")

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	err = cli.ContainerKill(context.Background(), id, "SIGKILL")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Kill Access Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
		Success: true,
		Data:    nil,
	})
}

func RenameContainer(c *fiber.Ctx) error {
	id := c.Params("id")

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	var dto RenameContainerDto

	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(utils.BasicResponse[bool]{
				Success: false,
				Message: "올바른 Body 타입이 아닙니다.",
			})
	}

	var validate = validator.New()

	if err := validate.Struct(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(utils.BasicResponse[bool]{
				Success: false,
				Message: err.Error(),
			})
	}

	err = cli.ContainerRename(context.Background(), id, dto.Name)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Rename Access Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
		Success: true,
		Data:    nil,
	})
}

func RemoveContainer(c *fiber.Ctx) error {
	id := c.Params("id")

	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	err = cli.ContainerRemove(context.Background(), id, container.RemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   true,
		Force:         true,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Remove Access Error:" + err.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
		Success: true,
		Data:    nil,
	})
}

func PruneContainer(c *fiber.Ctx) error {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()
	report, err := cli.ContainersPrune(context.Background(), filters.Args{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Prune Access Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[container.PruneReport]{
		Success: true,
		Data:    report,
	})
}

// func UpdateContainer(c *fiber.Ctx) error {
// 	id := c.Params("id")

// 	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
// 			Success: false,
// 			Message: "Docker Connect Error:" + err.Error(),
// 			Data:    nil,
// 		})
// 	}

// 	defer cli.Close()

// 	cli.ContainerUpdate(context.Background(), id, container.UpdateConfig{
// 		container.Resources{},
// 	})

// 	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
// 		Success: true,
// 		Data:    nil,
// 	})
// }
