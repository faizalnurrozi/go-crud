package usecase

import (
	"database/sql"
	"github.com/faizalnurrozi/go-crud/domain/view_models"
	"github.com/faizalnurrozi/go-crud/pkg/jwe"
	"github.com/faizalnurrozi/go-crud/pkg/jwt"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Contract struct {
	ReqID         string
	UserID        string // Credential ID
	RoleID        int
	App           *fiber.App
	DB            *sql.DB
	TX            *sql.Tx
	JweCredential jwe.Credential
	JwtCredential jwt.JwtCredential
	Validate      *validator.Validate
	Translator    ut.Translator
}

const (
	//default limit for pagination
	defaultLimit = 10

	//max limit for pagination
	maxLimit = 50

	//default order by
	defaultOrderBy = "id"

	//default sort
	defaultSort = "asc"

	//default last page for pagination
	defaultLastPage = 0

	//default role id seller
	defaultRoleIDSeller = 3

	//default status for seller
	defaultSellerStatus = "waiting_approval"

	//dafault status condition seller
	defaultSellerStatusCondition = "approve"

	//default delete,count,read by
	defaultFilterBy = `id`

	superAdminRoleID = 1
	adminRoleID      = 2
	sellerRoleID     = 3
	buyerRoleID      = 4

	//kafka topic
	kafkaTopicMailNotification = "mail_notification"
)

var (
	enumType     = []string{"sex-enum", "seller-status-enum"}
	enumMailType = []string{"forgot-password", "generated-password"}
)

func (uc Contract) setPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}
	if order == "" {
		order = defaultOrderBy
	}
	if sort == "" {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc Contract) setPaginationResponse(page, limit, total int) (res view_models.PaginationVm) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	vm := view_models.NewPaginationVm()
	res = vm.Build(view_models.DetailPaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	})

	return res
}
