package handler

import (
	"course_service/internal/app/entity"
	"course_service/internal/app/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ActHandler struct {
	actUsecase usecase.ActUsecase
}

func NewActHandler(actUsecase usecase.ActUsecase) *ActHandler {
	return &ActHandler{actUsecase: actUsecase}
}

func (h *ActHandler) GetActsByAnswerID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	acts, err := h.actUsecase.GetActsByAnswerID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(acts)
}

func (h *ActHandler) CreateAct(c *fiber.Ctx) error {
	var act entity.Act
	if err := c.BodyParser(&act); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	err := h.actUsecase.CreateAct(&act)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(act)
}

func (h *ActHandler) UpdateAct(c *fiber.Ctx) error {
	var act entity.Act
	if err := c.BodyParser(&act); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	err = h.actUsecase.UpdateAct(&act, uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(act)
}

func (h *ActHandler) DeleteAct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	err = h.actUsecase.DeleteAct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("Act deleted successfully")
}
