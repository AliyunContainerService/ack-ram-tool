package binding

import (
	"context"
	"errors"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"strconv"
)

func ListAccounts(ctx context.Context, client openapi.ClientInterface) (map[int64]types.Account, error) {
	accounts := make(map[int64]types.Account, 0)
	if acc, err := client.GetCallerIdentity(ctx); err != nil {
		log.Logger.Errorf("GetCallerIdentity failed: %s", err)
		return nil, err
	} else {
		id, _ := strconv.ParseInt(acc.RootUId, 10, 64)
		accounts[id] = types.NewRootAccount(id)
	}

	users, err := client.ListUsers(ctx)
	if err != nil {
		log.Logger.Errorf("list users failed: %s", err)
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("list users failed: No users found")
	}

	roles, err := client.ListRoles(ctx)
	if err != nil {
		log.Logger.Errorf("list roles failed: %s", err)
		return nil, err
	}
	if len(roles) == 0 {
		return nil, errors.New("list roles failed: No roles found")
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
