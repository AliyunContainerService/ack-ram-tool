package binding

import (
	"context"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"strconv"
)

func ListAccounts(ctx context.Context, client openapi.RamClientInterface) (map[int64]types.Account, error) {
	accounts := make(map[int64]types.Account, 0)
	users, err := client.ListUsers(ctx)
	if err != nil {
		log.Logger.Errorf("list users failed: %s", err)
		return nil, err
	}
	roles, err := client.ListRoles(ctx)
	if err != nil {
		log.Logger.Errorf("list roles failed: %s", err)
		return nil, err
	}
	for _, u := range users {
		id, _ := strconv.ParseInt(u.Id, 10, 64)
		accounts[id] = types.Account{
			Type: types.AccountTypeUser,
			User: u,
		}
	}
	for _, r := range roles {
		id, _ := strconv.ParseInt(r.RoleId, 10, 64)
		accounts[id] = types.Account{
			Type: types.AccountTypeRole,
			Role: r,
		}
	}

	return accounts, nil
}
