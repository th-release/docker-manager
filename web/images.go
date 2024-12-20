package web

import (
	"context"

	"cth-core.xyz/docker-manager/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/gofiber/fiber/v2"
)

func GetAllImage(c *fiber.Ctx) error {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	images, err := cli.ImageList(context.Background(), image.ListOptions{
		All: true,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Image Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(utils.ArrayBasicResponse[image.Summary]{
		Success: true,
		Data:    images,
	})
}

func GetByIdImage(c *fiber.Ctx) error {
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

	images, _, err := cli.ImageInspectWithRaw(context.Background(), id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Image Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[types.ImageInspect]{
		Success: true,
		Data:    images,
	})
}

func GetHistoryById(c *fiber.Ctx) error {
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

	images, err := cli.ImageHistory(context.Background(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Image Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(utils.ArrayBasicResponse[image.HistoryResponseItem]{
		Success: true,
		Data:    images,
	})
}
