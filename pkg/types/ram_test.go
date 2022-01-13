package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeAssumeRolePolicyStatementWithServiceAccount(t *testing.T) {
	policy := MakeAssumeRolePolicyStatementWithServiceAccount(
		"https://oidc-ack.example.com", "acs:ram::12345:oidc-provider/ack-rrsa-cabce", "depart-1", "oss-reader-sa")
	assert.JSONEq(t, policy.JSON(), `
{
      "Action": "sts:AssumeRole",
      "Condition": {
        "StringEquals": {
          "oidc:aud": "sts.aliyuncs.com",
          "oidc:iss": "https://oidc-ack.example.com",
          "oidc:sub": "system:serviceaccount:depart-1:oss-reader-sa"
        }
      },
      "Effect": "Allow",
      "Principal": {
        "Federated": [
          "acs:ram::12345:oidc-provider/ack-rrsa-cabce"
        ]
      }
}
`)
}

func TestAssumeRolePolicyDocument_AppendPolicyIfNotExist(t *testing.T) {
	type args struct {
		policy string
	}
	rrsaState := MakeAssumeRolePolicyStatementWithServiceAccount(
		"https://oidc-ack.example.com", "acs:ram::12345:oidc-provider/ack-rrsa-cabce", "depart-1", "oss-reader-sa")
	tests := []struct {
		name    string
		p       string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "no old policy",
			p:    "",
			args: args{
				policy: rrsaState.JSON(),
			},
			want: `
{
  "Statement": [
		{
			  "Action": "sts:AssumeRole",
			  "Condition": {
				"StringEquals": {
				  "oidc:aud": "sts.aliyuncs.com",
				  "oidc:iss": "https://oidc-ack.example.com",
				  "oidc:sub": "system:serviceaccount:depart-1:oss-reader-sa"
				}
			  },
			  "Effect": "Allow",
			  "Principal": {
				"Federated": [
				  "acs:ram::12345:oidc-provider/ack-rrsa-cabce"
				]
			  }
		}
	],
"Version": "1"
}
`,
			wantErr: false,
		},
		{
			name: "merge old policy v1",
			p: `
{
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "cs.aliyuncs.com"
        ]
      }
    }
  ],
  "Version": "1"
}
`,
			args: args{
				policy: rrsaState.JSON(),
			},
			want: `
{
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "cs.aliyuncs.com"
        ]
      }
    },
		{
			  "Action": "sts:AssumeRole",
			  "Condition": {
				"StringEquals": {
				  "oidc:aud": "sts.aliyuncs.com",
				  "oidc:iss": "https://oidc-ack.example.com",
				  "oidc:sub": "system:serviceaccount:depart-1:oss-reader-sa"
				}
			  },
			  "Effect": "Allow",
			  "Principal": {
				"Federated": [
				  "acs:ram::12345:oidc-provider/ack-rrsa-cabce"
				]
			  }
		}
	],
"Version": "1"
}
`,
			wantErr: false,
		},
		{
			name: "merge old policy v2",
			p: `
{
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "cs.aliyuncs.com"
        ]
      }
    },
	{
			  "Action": "sts:AssumeRole",
			  "Condition": {
				"StringEquals": {
				  "oidc:aud": "sts.aliyuncs.com",
				  "oidc:iss": "https://oidc-ack.example.com/v1",
				  "oidc:sub": "system:serviceaccount:depart-1:oss-reader-sa"
				}
			  },
			  "Effect": "Allow",
			  "Principal": {
				"Federated": [
				  "acs:ram::12345:oidc-provider/ack-rrsa-cabce"
				]
			  }
		}
  ],
  "Version": "1"
}
`,
			args: args{
				policy: rrsaState.JSON(),
			},
			want: `
{
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "cs.aliyuncs.com"
        ]
      }
    },
	{
			  "Action": "sts:AssumeRole",
			  "Condition": {
				"StringEquals": {
				  "oidc:aud": "sts.aliyuncs.com",
				  "oidc:iss": "https://oidc-ack.example.com/v1",
				  "oidc:sub": "system:serviceaccount:depart-1:oss-reader-sa"
				}
			  },
			  "Effect": "Allow",
			  "Principal": {
				"Federated": [
				  "acs:ram::12345:oidc-provider/ack-rrsa-cabce"
				]
			  }
		},
		{
			  "Action": "sts:AssumeRole",
			  "Condition": {
				"StringEquals": {
				  "oidc:aud": "sts.aliyuncs.com",
				  "oidc:iss": "https://oidc-ack.example.com",
				  "oidc:sub": "system:serviceaccount:depart-1:oss-reader-sa"
				}
			  },
			  "Effect": "Allow",
			  "Principal": {
				"Federated": [
				  "acs:ram::12345:oidc-provider/ack-rrsa-cabce"
				]
			  }
		}
	],
"Version": "1"
}
`,
			wantErr: false,
		},
		{
			name: "merge unkown fields",
			p: `
{
	"Foo": "bar",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "cs.aliyuncs.com"
        ],
      "Foobar": "foo"
      }
    }
  ],
  "Version": "1"
}
`,
			args: args{
				policy: rrsaState.JSON(),
			},
			want: `
{
	"Foo": "bar",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "cs.aliyuncs.com"
        ],
      "Foobar": "foo"
      }
    },
		{
			  "Action": "sts:AssumeRole",
			  "Condition": {
				"StringEquals": {
				  "oidc:aud": "sts.aliyuncs.com",
				  "oidc:iss": "https://oidc-ack.example.com",
				  "oidc:sub": "system:serviceaccount:depart-1:oss-reader-sa"
				}
			  },
			  "Effect": "Allow",
			  "Principal": {
				"Federated": [
				  "acs:ram::12345:oidc-provider/ack-rrsa-cabce"
				]
			  }
		}
	],
"Version": "1"
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &AssumeRolePolicyDocument{}
			_ = json.Unmarshal([]byte(tt.p), p)
			policy := &AssumeRolePolicyStatement{}
			_ = json.Unmarshal([]byte(tt.args.policy), policy)

			if err := p.AppendPolicyIfNotExist(*policy); (err != nil) != tt.wantErr {
				t.Errorf("AppendPolicyIfNotExist() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.JSONEq(t, tt.want, p.JSON())
			}
		})
	}
}
