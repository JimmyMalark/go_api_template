package keys

import (
	"github.com/jimmymalark/go_api_template/internal/pagination"
	"fmt"
)

const UsersPrefix = "users:"

func UsersList(p pagination.Params) string {
	return fmt.Sprintf(
		UsersPrefix+"page:%d:limit:%d",
		p.Page,
		p.Limit,
	)
}
