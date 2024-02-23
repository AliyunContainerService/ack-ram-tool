package oidctoken

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/version"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	sts "github.com/alibabacloud-go/sts-20150401/client"
	"github.com/alibabacloud-go/tea/tea"
)

const (
	defaultSTSEndpoint = "sts.aliyuncs.com"
	defaultSTSProtocol = "https"
)

var defaultExpiryWindow = time.Minute * 10

type Credential struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
	Expiration      time.Time
}

type RoleProvider struct {
	providerArn string
	roleArn     string
	tokenFile   string

	policy          string
	roleSessionName string
	sessionDuration time.Duration

	stsEndpoint string
	stsProtocol string

	expiryWindow time.Duration

	cred        *Credential
	lockForCred sync.RWMutex
}

func NewRoleProvider(providerArn, roleArn, tokenFile, policy, roleSessionName string, sessionDuration time.Duration) (*RoleProvider, error) {
	f, err := os.Open(filepath.Clean(tokenFile))
	if err != nil {
		return nil, err
	}
	defer f.Close() // #nosec G307

	p := &RoleProvider{
		providerArn:     providerArn,
		roleArn:         roleArn,
		tokenFile:       tokenFile,
		policy:          policy,
		roleSessionName: roleSessionName,
		sessionDuration: sessionDuration,
		stsEndpoint:     defaultSTSEndpoint,
		stsProtocol:     defaultSTSProtocol,
		expiryWindow:    defaultExpiryWindow,
		lockForCred:     sync.RWMutex{},
	}
	go func() {
		p.refreshCredLoop(context.TODO())
	}()
	return p, nil
}

func (p *RoleProvider) GetCredential(ctx context.Context) (*Credential, error) {
	cred := p.getCred()

	if cred != nil && !cred.shouldRefresh(p.expiryWindow) {
		return cred.DeepCopy(), nil
	}

	cred, err := p.retrieve(ctx)
	if err != nil {
		return nil, err
	}
	p.setCred(cred)

	return cred, nil
}

func (p *RoleProvider) refreshCredLoop(ctx context.Context) {
	if p.expiryWindow <= 0 {
		return
	}

	ticket := time.NewTicker(time.Minute)
	defer ticket.Stop()

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case <-ticket.C:
			_, err := p.GetCredential(context.TODO())
			if err != nil {
				log.Logger.Errorf("refresh credential failed: %+v", err)
			}
		}
	}
}

func (p *RoleProvider) getCred() *Credential {
	p.lockForCred.RLock()
	defer p.lockForCred.RUnlock()

	cred := p.cred
	return cred
}

func (p *RoleProvider) setCred(cred *Credential) {
	p.lockForCred.Lock()
	defer p.lockForCred.Unlock()

	p.cred = cred.DeepCopy()
}

func (p *RoleProvider) retrieve(ctx context.Context) (*Credential, error) {
	token, err := os.ReadFile(p.tokenFile)
	if err != nil {
		return nil, err
	}
	c, err := AssumeRoleWithOIDCToken(ctx,
		p.providerArn, p.roleArn, string(token), p.stsEndpoint, p.stsProtocol, p.policy,
		p.roleSessionName, p.sessionDuration)
	return c, err
}

func AssumeRoleWithOIDCToken(ctx context.Context, providerArn, roleArn, token, stsEndpoint, stsProtocol, policy, roleSessionName string,
	sessionDuration time.Duration) (*Credential, error) {
	stsClient, err := sts.NewClient(&openapi.Config{
		Endpoint:  tea.String(stsEndpoint),
		Protocol:  tea.String(stsProtocol),
		UserAgent: tea.String(version.UserAgent()),
	})
	if err != nil {
		return nil, err
	}
	sessionName := roleSessionName
	if sessionName == "" {
		sessionName = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	stsReq := &sts.AssumeRoleWithOIDCRequest{
		OIDCProviderArn: tea.String(providerArn),
		RoleArn:         tea.String(roleArn),
		OIDCToken:       tea.String(token),
		RoleSessionName: tea.String(sessionName),
	}
	if policy != "" {
		stsReq.Policy = tea.String(policy)
	}
	if sessionDuration > 0 {
		stsReq.DurationSeconds = tea.Int64(int64(sessionDuration / time.Second))
	}
	r, err := stsClient.AssumeRoleWithOIDC(stsReq)
	if err != nil {
		return nil, err
	}
	body := r.Body
	if body == nil || body.Credentials == nil {
		return nil, fmt.Errorf("invalid body: %q", r.String())
	}
	cred := body.Credentials
	exp, err := time.Parse(time.RFC3339, tea.StringValue(cred.Expiration))
	if err != nil {
		return nil, fmt.Errorf("parse Expiration failed: %+v, body: %q", err, r.String())
	}
	return &Credential{
		AccessKeyId:     tea.StringValue(cred.AccessKeyId),
		AccessKeySecret: tea.StringValue(cred.AccessKeySecret),
		SecurityToken:   tea.StringValue(cred.SecurityToken),
		Expiration:      exp,
	}, nil
}

func (c *Credential) shouldRefresh(expiryWindow time.Duration) bool {
	if c == nil {
		return true
	}
	expiryWindow = expiryWindow + time.Duration(rand.Int63n(int64(time.Minute))) // #nosec G404
	return time.Until(c.Expiration) <= expiryWindow
}

func (c *Credential) DeepCopy() *Credential {
	return &Credential{
		AccessKeyId:     c.AccessKeyId,
		AccessKeySecret: c.AccessKeySecret,
		SecurityToken:   c.SecurityToken,
		Expiration:      c.Expiration,
	}
}
