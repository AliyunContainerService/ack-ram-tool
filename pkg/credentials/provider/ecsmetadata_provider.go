package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	defaultExpiryWindow               = time.Minute * 30
	defaultECSMetadataServerEndpoint  = "http://100.100.100.200"
	defaultECSMetadataTokenTTLSeconds = 3600
	defaultClientTimeout              = time.Second * 30
)

type ECSMetadataProvider struct {
	u *Updater

	endpoint                string
	roleName                string
	metadataToken           string
	metadataTokenTTLSeconds int
	metadataTokenExp        time.Time

	client *http.Client
}

type ECSMetadataProviderOptions struct {
	Endpoint  string
	Timeout   time.Duration
	Transport http.RoundTripper

	RoleName                string
	MetadataTokenTTLSeconds int

	ExpiryWindow  time.Duration
	RefreshPeriod time.Duration
	Logger        Logger
}

func NewECSMetadataProvider(opts ECSMetadataProviderOptions) *ECSMetadataProvider {
	opts.applyDefaults()

	client := &http.Client{
		Transport: opts.Transport,
		Timeout:   opts.Timeout,
	}
	e := &ECSMetadataProvider{
		endpoint:                opts.Endpoint,
		roleName:                opts.RoleName,
		metadataTokenTTLSeconds: opts.MetadataTokenTTLSeconds,
		client:                  client,
	}
	e.u = NewUpdater(e.getCredentials, UpdaterOptions{
		ExpiryWindow:  opts.ExpiryWindow,
		RefreshPeriod: opts.RefreshPeriod,
		Logger:        opts.Logger,
		LogPrefix:     "[ECSMetadataProvider]",
	})
	e.u.Start(context.TODO())

	return e
}

func (e *ECSMetadataProvider) Credentials(ctx context.Context) (*Credentials, error) {
	return e.u.Credentials(ctx)
}

type ecsMetadataStsResponse struct {
	AccessKeyId     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	SecurityToken   string `json:"SecurityToken"`
	Expiration      string `json:"Expiration"`
	LastUpdated     string `json:"LastUpdated"`
	Code            string `json:"Code"`
}

type metadataError struct {
	code    int
	message string
}

func (e *ECSMetadataProvider) getCredentials(ctx context.Context) (*Credentials, error) {
	roleName, err := e.getRoleName(ctx)
	if err != nil {
		if e, ok := err.(*metadataError); ok && e.code == 404 {
			return nil, NewNotEnableError(fmt.Errorf("get role name from ecs metadata failed: %w", err))
		}
	}
	path := fmt.Sprintf("/latest/meta-data/ram/security-credentials/%s", roleName)
	data, err := e.getMedataDataWithToken(ctx, http.MethodGet, path)
	if err != nil {
		return nil, err
	}

	var obj ecsMetadataStsResponse
	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		return nil, fmt.Errorf("parse credentials failed: %w", err)
	}
	if obj.AccessKeySecret == "" {
		return nil, fmt.Errorf("parse credentials got unexpected data: %s",
			strings.ReplaceAll(data, "\n", " "))
	}

	exp, err := time.Parse("2006-01-02T15:04:05Z", obj.Expiration)
	if err != nil {
		return nil, err
	}
	return &Credentials{
		AccessKeyId:     obj.AccessKeyId,
		AccessKeySecret: obj.AccessKeySecret,
		SecurityToken:   obj.SecurityToken,
		Expiration:      exp,
	}, nil
}

func (e *ECSMetadataProvider) getRoleName(ctx context.Context) (string, error) {
	if e.roleName != "" {
		return e.roleName, nil
	}
	name, err := e.getMedataDataWithToken(ctx, http.MethodGet, "/latest/meta-data/ram/security-credentials/")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(name), nil
}

func (e *ECSMetadataProvider) getMedataToken(ctx context.Context) (string, error) {
	if !e.metadataTokenExp.Before(time.Now()) {
		return e.metadataToken, nil
	}

	h := http.Header{}
	h.Set("X-aliyun-ecs-metadata-token-ttl-seconds", fmt.Sprintf("%d", e.metadataTokenTTLSeconds))
	body, err := e.getMedataData(ctx, http.MethodPut, "/latest/api/token", h)
	if err != nil {
		return "", fmt.Errorf("get metadata token failed: %w", err)
	}

	e.metadataToken = strings.TrimSpace(body)
	e.metadataTokenExp = time.Now().Add(time.Duration(float64(e.metadataTokenTTLSeconds)*0.8) * time.Second)

	return body, nil
}

func (e *ECSMetadataProvider) getMedataDataWithToken(ctx context.Context, method, path string) (string, error) {
	token, err := e.getMedataToken(ctx)
	if err != nil {
		if e, ok := err.(*metadataError); !(ok && e.code == 404) {
			return "", err
		}
	}
	h := http.Header{}
	if token != "" {
		h.Set("X-aliyun-ecs-metadata-token", token)
	}
	return e.getMedataData(ctx, method, path, h)
}

func (e *ECSMetadataProvider) getMedataData(ctx context.Context, method, path string, header http.Header) (string, error) {
	url := fmt.Sprintf("%s%s", e.endpoint, path)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", fmt.Errorf("can not init request with url %s: %w", url, err)
	}
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", UserAgent)
	for k, items := range header {
		for _, v := range items {
			req.Header.Add(k, v)
		}
	}

	resp, err := e.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request %s failed: %w", url, err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body failed when request %s: %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", &metadataError{
			code: resp.StatusCode,
			message: fmt.Sprintf("status code %d is not 200 when request %s: %s",
				resp.StatusCode, url, strings.ReplaceAll(string(data), "\n", " ")),
		}
	}
	return string(data), nil
}

func (o *ECSMetadataProviderOptions) applyDefaults() {
	if o.Timeout <= 0 {
		o.Timeout = defaultClientTimeout
	}
	if o.Transport == nil {
		ts := http.DefaultTransport.(*http.Transport).Clone()
		o.Transport = ts
	}
	if o.Endpoint == "" {
		o.Endpoint = defaultECSMetadataServerEndpoint
	} else {
		o.Endpoint = strings.TrimRight(o.Endpoint, "/")
	}
	if o.MetadataTokenTTLSeconds == 0 {
		o.MetadataTokenTTLSeconds = defaultECSMetadataTokenTTLSeconds
	}
	if o.ExpiryWindow == 0 {
		o.ExpiryWindow = defaultExpiryWindow
	}
	if o.Logger == nil {
		o.Logger = defaultLog
	}
}

func (e metadataError) Error() string {
	return e.message
}
