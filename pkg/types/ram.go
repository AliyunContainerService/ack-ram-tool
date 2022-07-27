package types

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

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

type RamRole struct {
	RoleName                 string
	RoleId                   string
	Arn                      string
	Description              string
	AssumeRolePolicyDocument *RamPolicyDocument
	MaxSessionDuration       int64
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
	if serviceAccount == "" {
		serviceAccount = "*"
	}
	return RamPolicyStatement{
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

func (p *RamPolicyDocument) JSON() string {
	data, err := json.MarshalIndent(p, " ", " ")
	if err != nil {
		log.Printf("errro: %+v\n", err)
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
		log.Printf("parse statements failed: %+v", err)
		return nil, err
	}
	var policies []RamPolicyStatement
	if err := json.Unmarshal(statementsJson, &policies); err != nil {
		log.Printf("parse statements failed: %+v", err)
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
