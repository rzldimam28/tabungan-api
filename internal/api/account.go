package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rzldimam28/tabungan-api/domain"
	"github.com/rzldimam28/tabungan-api/dto"
	"github.com/rzldimam28/tabungan-api/internal/util"
)

type accountApi struct {
	accountService domain.AccountService
	mutationService domain.MutationService
}

func NewAccount(app *fiber.App, accountService domain.AccountService, mutationService domain.MutationService) {
	h := accountApi{
		accountService: accountService,
		mutationService: mutationService,
	}

	app.Post("/daftar", h.Register)
	app.Post("/tabung", h.TopUp)
	app.Post("/tarik", h.Withdrawal)
	app.Get("/saldo/:no_rekening", h.CheckBalance)
	app.Get("/mutasi/:no_rekening", h.CheckMutations)
}

func (ths *accountApi) Register(ctx *fiber.Ctx) error {
	
	var req dto.RegisterAccountReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Remark: dto.Remark{
				ErrMessage: "Cant parse request body",
			},
		})
	}

	res, err := ths.accountService.Create(ctx.Context(), req)
	if err != nil {
		errCode, errMsg := util.TranslateErrorCategories(err)
		return ctx.Status(errCode).JSON(dto.ErrorResponse{
			Success: false,
			Remark: dto.Remark{
				ErrMessage: errMsg,
			},
		})
	}
	
	return ctx.Status(200).JSON(dto.SuccessResponse{
		Success: true,
		Data: res,
	})
}

func (ths *accountApi) TopUp(ctx *fiber.Ctx) error {

	var req dto.TopUpReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Remark: dto.Remark{
				ErrMessage: "Cant parse request body",
			},
		})
	}

	res, err := ths.accountService.TopUp(ctx.Context(), req)
	if err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Remark: dto.Remark{
				ErrMessage: err.Error(),
			},
		})
	}

	return ctx.Status(200).JSON(dto.SuccessResponse{
		Success: true,
		Data: res,
	})
}

func (ths *accountApi) Withdrawal(ctx *fiber.Ctx) error {

	var req dto.WithdrawalReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Remark: dto.Remark{
				ErrMessage: "Cant parse request body",
			},
		})
	}

	res, err := ths.accountService.Withdrawal(ctx.Context(), req)
	if err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Remark: dto.Remark{
				ErrMessage: err.Error(),
			},
		})
	}

	return ctx.Status(200).JSON(dto.SuccessResponse{
		Success: true,
		Data: res,
	})
}

func (ths *accountApi) CheckBalance(ctx *fiber.Ctx) error {

	noRekening := ctx.Params("no_rekening")

	res, err := ths.accountService.GetByNoRekening(ctx.Context(), noRekening)
	if err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Remark: dto.Remark{
				ErrMessage: err.Error(),
			},
		})
	}

	return ctx.Status(200).JSON(dto.SuccessResponse{
		Success: true,
		Data: res,
	})
}

func (ths *accountApi) CheckMutations(ctx *fiber.Ctx) error {

	noRekening := ctx.Params("no_rekening")

	res, err := ths.mutationService.GetByNoRekening(ctx.Context(), noRekening)
	if err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{
			Success: false,
			Remark: dto.Remark{
				ErrMessage: err.Error(),
			},
		})
	}

	return ctx.Status(200).JSON(dto.SuccessResponse{
		Success: true,
		Data: res,
	})
}