package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
)

const (
	RamPolicyKeyStatement  = "Statement"
	RamPolicyKeyVersion    = "Version"
	RamPolicyVersionLatest = "1"
)

var (
	RamPolicyTypeSystem = "System"
	RamPolicyTypeCustom = "Custom"
)

type AccountType string

const (
	AccountTypeRoot AccountType = "Root"
	AccountTypeUser AccountType = "RamUser"
	AccountTypeRole AccountType = "RamRole"
)

type Account struct {
	Type    AccountType
	RootUId string
	User    RamUser
	Role    RamRole

	PrincipalId string
	Arn         string
}

type RamRole struct {
	RoleName                 string
	RoleId                   string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument *RamPolicyDocument
	MaxSessionDuration       int64
	Deleted                  bool
}

type RamUser struct {
	Id          string
	Name        string
	DisplayName string
	Deleted     bool
}

type RamPolicy struct {
	DefaultVersion string
	//UpdateDate      string
	Description     string
	PolicyDocument  *RamPolicyDocument
	AttachmentCount int32
	PolicyName      string
	//CreateDate      string
	PolicyType string
}

type RamRolePolicy struct {
	DefaultVersion string
	Description    string
	PolicyName     string
	//AttachDate     string
	PolicyType string
}

type RamPolicyDocument map[string]interface{}
type RamPolicyStatement map[string]interface{}

func MakeRamPolicyDocument(policies []RamPolicyStatement) RamPolicyDocument {
	return RamPolicyDocument{
		RamPolicyKeyStatement: policies,
		RamPolicyKeyVersion:   RamPolicyVersionLatest,
	}
}

func MakeAssumeRolePolicyStatementWithServiceAccount(oidcIssuer, oidcArn, namespace, serviceAccount string) RamPolicyStatement {
	condition := map[string]interface{}{
		"StringEquals": map[string]string{
			"oidc:aud": "sts.aliyuncs.com",
			"oidc:iss": oidcIssuer,
			"oidc:sub": getRAMSub(namespace, serviceAccount),
		},
	}
	if namespace == "*" || serviceAccount == "*" {
		condition = map[string]interface{}{
			"StringEquals": map[string]string{
				"oidc:aud": "sts.aliyuncs.com",
				"oidc:iss": oidcIssuer,
			},
			"StringLike": map[string]string{
				"oidc:sub": getRAMSub(namespace, serviceAccount),
			},
		}
	}
	return RamPolicyStatement{
		"Action":    "sts:AssumeRole",
		"Condition": condition,
		"Effect":    "Allow",
		"Principal": map[string]interface{}{
			"Federated": []string{oidcArn},
		},
	}
}

func getRAMSub(namespace, serviceAccount string) string {
	return fmt.Sprintf("system:serviceaccount:%s:%s", namespace, serviceAccount)
}

func (p *RamPolicyDocument) AppendPolicyIfNotExist(policy RamPolicyStatement) error {
	if exist, err := p.IncludePolicy(policy); err != nil {
		return err
	} else if exist {
		return nil
	}
	return p.AppendPolicy(policy)
}

func (p *RamPolicyDocument) AppendPolicy(policy RamPolicyStatement) error {
	if p == nil || len(*p) == 0 {
		*p = MakeRamPolicyDocument([]RamPolicyStatement{})
	}
	if _, ok := (*p)[RamPolicyKeyStatement]; !ok {
		(*p)[RamPolicyKeyStatement] = []RamPolicyStatement{}
	}
	policies, err := p.policies()
	if err != nil {
		return err
	}

	policies = append(policies, policy)
	(*p)[RamPolicyKeyStatement] = policies
	return nil
}

func (p *RamPolicyDocument) IncludePolicy(policy RamPolicyStatement) (bool, error) {
	policies, err := p.policies()
	if err != nil {
		log.Logger.Errorf("parse statements failed: %+v", err)
		return false, err
	}

	for _, p := range policies {
		if p.Equal(policy) {
			return true, nil
		}
	}

	return false, nil
}

func (p *RamPolicyDocument) JSON() string {
	data, err := json.MarshalIndent(p, " ", " ")
	if err != nil {
		log.Logger.Errorf("errro: %+v", err)
	}
	return string(data)
}

func (p *RamPolicyDocument) policies() ([]RamPolicyStatement, error) {
	if p == nil || len(*p) == 0 {
		return nil, nil
	}
	statements, ok := (*p)[RamPolicyKeyStatement]
	if !ok {
		return nil, nil
	}
	statementsJson, err := json.Marshal(statements)
	if err != nil {
		log.Logger.Errorf("parse statements failed: %+v", err)
		return nil, err
	}
	var policies []RamPolicyStatement
	if err := json.Unmarshal(statementsJson, &policies); err != nil {
		log.Logger.Errorf("parse statements failed: %+v", err)
		return nil, err
	}
	return policies, nil
}

func (s *RamPolicyStatement) Equal(another RamPolicyStatement) bool {
	if reflect.DeepEqual(s, another) {
		return true
	}
	if utils.JSONEqual(s, another) {
		return true
	}
	if utils.StringInterfaceMapEqual(*s, another) {
		return true
	}
	return false
}

func (s *RamPolicyStatement) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}

func (a Account) Id() string {
	switch a.Type {
	case AccountTypeUser:
		return a.User.Id
	case AccountTypeRole:
		return a.Role.RoleId
	}
	return ""
}
func (a Account) Name() string {
	switch a.Type {
	case AccountTypeUser:
		return a.User.Name
	case AccountTypeRole:
		return a.Role.RoleName
	}
	return ""
}

func (a Account) Deleted() bool {
	switch a.Type {
	case AccountTypeUser:
		return a.User.Deleted
	case AccountTypeRole:
		return a.Role.Deleted
	}
	return false
}

func NewRootAccount(uid int64) Account {
	idStr := fmt.Sprintf("%d", uid)
	return Account{
		RootUId: idStr,
		Type:    AccountTypeRoot,
		User: RamUser{
			Id: idStr,
		},
	}
}

func NewFakeAccount(uid int64) Account {
	idStr := fmt.Sprintf("%d", uid)
	acc := Account{}
	switch {
	case strings.HasPrefix(idStr, "1"), strings.HasPrefix(idStr, "5"):
		acc.Type = AccountTypeRoot
		acc.User = RamUser{
			Id: idStr,
		}
	case strings.HasPrefix(idStr, "2"):
		acc.Type = AccountTypeUser
		acc.User = RamUser{
			Id: idStr,
		}
	case strings.HasPrefix(idStr, "3"):
		acc.Type = AccountTypeRole
		acc.Role = RamRole{
			RoleId: idStr,
		}
	}
	return acc
}

func (a *Account) IdentityType() string {
	switch a.Type {
	case AccountTypeRoot:
		return "Account"
	case AccountTypeUser:
		return "RAMUser"
	case AccountTypeRole:
		return "AssumedRoleUser"
	}
	return ""
}

func (a *Account) MarkDeleted() {
	switch a.Type {
	case AccountTypeUser:
		u := a.User
		u.Deleted = true
		a.User = u
	case AccountTypeRole:
		r := a.Role
		r.Deleted = true
		a.Role = r
	}
}
