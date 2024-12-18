package web

import (
	"context"

	"cth-core.xyz/docker-manager/utils"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllNetwork(c *fiber.Ctx) error {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	networks, err := cli.NetworkList(context.Background(), network.ListOptions{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Network Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.ArrayBasicResponse[network.Inspect]{
		Success: true,
		Message: "",
		Data:    networks,
	})
}

func GetByIdNetwork(c *fiber.Ctx) error {
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
	data, err := cli.NetworkInspect(context.Background(), id, network.InspectOptions{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Network Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[network.Inspect]{
		Success: true,
		Data:    data,
	})
}

func ConnectNetwork(c *fiber.Ctx) error {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()
	var dto ConnectNetworkDto

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

	err = cli.NetworkConnect(context.Background(), dto.Network, dto.Container, &network.EndpointSettings{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Network Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
		Success: true,
		Data:    nil,
	})
}

func CreateNetwork(c *fiber.Ctx) error {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()
	var dto CreateNetworkDto

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

	data, err := cli.NetworkCreate(context.Background(), dto.Name, network.CreateOptions{
		Driver:     dto.Driver,
		Scope:      dto.Scope,
		Ingress:    dto.Ingress,
		EnableIPv6: &dto.EnableIpv6,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Network Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[network.CreateResponse]{
		Success: true,
		Data:    data,
	})
}

func RemoveNetwork(c *fiber.Ctx) error {
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

	err = cli.NetworkRemove(context.Background(), id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Network Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
		Success: true,
		Data:    nil,
	})
}

func DisconnectNetwork(c *fiber.Ctx) error {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()
	var dto DisconnectNetworkDto

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

	err = cli.NetworkDisconnect(context.Background(), dto.Network, dto.Container, dto.Force)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Network Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[*bool]{
		Success: true,
		Data:    nil,
	})
}

func PruneNetwork(c *fiber.Ctx) error {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	defer cli.Close()

	report, err := cli.NetworksPrune(context.Background(), filters.Args{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.BasicResponse[*bool]{
			Success: false,
			Message: "Docker Network Connect Error:" + err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.BasicResponse[network.PruneReport]{
		Success: true,
		Data:    report,
	})
}
