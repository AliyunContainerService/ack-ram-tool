package types

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/utils"
)

const (
	AssumeRolePolicyKeyStatement  = "Statement"
	AssumeRolePolicyKeyVersion    = "Version"
	AssumeRolePolicyVersionLatest = "1"
)

type RamRole struct {
	RoleName                 string
	RoleId                   string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument *AssumeRolePolicyDocument
	MaxSessionDuration       int64
}

type AssumeRolePolicyDocument map[string]interface{}
type AssumeRolePolicyStatement map[string]interface{}

func MakeAssumeRolePolicyDocument(policies []AssumeRolePolicyStatement) AssumeRolePolicyDocument {
	return AssumeRolePolicyDocument{
		AssumeRolePolicyKeyStatement: policies,
		AssumeRolePolicyKeyVersion:   AssumeRolePolicyVersionLatest,
	}
}

func MakeAssumeRolePolicyStatementWithServiceAccount(oidcIssuer, oidcArn, namespace, serviceAccount string) AssumeRolePolicyStatement {
	if serviceAccount == "" {
		serviceAccount = "*"
	}
	return AssumeRolePolicyStatement{
		"Action": "sts:AssumeRole",
		"Condition": map[string]interface{}{
			"StringEquals": map[string]string{
				"oidc:aud": "sts.aliyuncs.com",
				"oidc:iss": oidcIssuer,
				"oidc:sub": fmt.Sprintf("system:serviceaccount:%s:%s", namespace, serviceAccount),
			},
		},
		"Effect": "Allow",
		"Principal": map[string]interface{}{
			"Federated": []string{oidcArn},
		},
	}
}

func (p *AssumeRolePolicyDocument) AppendPolicyIfNotExist(policy AssumeRolePolicyStatement) error {
	if exist, err := p.IncludePolicy(policy); err != nil {
		return err
	} else if exist {
		return nil
	}
	return p.AppendPolicy(policy)
}

func (p *AssumeRolePolicyDocument) AppendPolicy(policy AssumeRolePolicyStatement) error {
	if p == nil || len(*p) == 0 {
		*p = MakeAssumeRolePolicyDocument([]AssumeRolePolicyStatement{})
	}
	if _, ok := (*p)[AssumeRolePolicyKeyStatement]; !ok {
		(*p)[AssumeRolePolicyKeyStatement] = []AssumeRolePolicyStatement{}
	}
	policies, err := p.policies()
	if err != nil {
		return err
	}

	policies = append(policies, policy)
	(*p)[AssumeRolePolicyKeyStatement] = policies
	return nil
}

func (p *AssumeRolePolicyDocument) IncludePolicy(policy AssumeRolePolicyStatement) (bool, error) {
	policies, err := p.policies()
	if err != nil {
		log.Printf("parse statements failed: %+v", err)
		return false, err
	}

	for _, p := range policies {
		if p.Equal(policy) {
			return true, nil
		}
	}

	return false, nil
}

func (p *AssumeRolePolicyDocument) JSON() string {
	data, err := json.MarshalIndent(p, " ", " ")
	if err != nil {
		log.Printf("errro: %+v\n", err)
	}
	return string(data)
}

func (p *AssumeRolePolicyDocument) policies() ([]AssumeRolePolicyStatement, error) {
	if p == nil || len(*p) == 0 {
		return nil, nil
	}
	statements, ok := (*p)[AssumeRolePolicyKeyStatement]
	if !ok {
		return nil, nil
	}
	statementsJson, err := json.Marshal(statements)
	if err != nil {
		log.Printf("parse statements failed: %+v", err)
		return nil, err
	}
	var policies []AssumeRolePolicyStatement
	if err := json.Unmarshal(statementsJson, &policies); err != nil {
		log.Printf("parse statements failed: %+v", err)
		return nil, err
	}
	return policies, nil
}

func (s *AssumeRolePolicyStatement) Equal(another AssumeRolePolicyStatement) bool {
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

func (s *AssumeRolePolicyStatement) JSON() string {
	data, _ := json.Marshal(s)
	return string(data)
}
