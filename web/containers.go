package web

import (
	"context"

	"cth-core.xyz/docker-manager/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gofiber/fiber/v2"
)

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
