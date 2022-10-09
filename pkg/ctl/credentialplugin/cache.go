package credentialplugin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const (
	expirationDelta time.Duration = time.Minute * 5
)

var defaultCacheDir = filepath.Join("~", ".kube", "cache", "ack-ram-tool")
var (
	errNoValidCache     = errors.New("no valid cache")
	errNeedRefreshCache = errors.New("need refresh cache")
)

type CredentialCache struct {
	cacheFilePath string
}

func NewCredentialCache(cacheDir string, opts GetCredentialOpts) *CredentialCache {
	return &CredentialCache{
		cacheFilePath: getCacheFilePath(cacheDir, opts),
	}
}

func (c *CredentialCache) GetCredential() (*types.ExecCredential, error) {
	data, err := ioutil.ReadFile(c.cacheFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errNoValidCache
		}
		return nil, err
	}
	var cred types.ExecCredential
	if err := json.Unmarshal(data, &cred); err != nil {
		return nil, err
	}
	remain := time.Until(cred.Status.ExpirationTimestamp.Time)
	if remain <= 0 {
		return nil, errNoValidCache
	} else if remain <= expirationDelta {
		return nil, errNeedRefreshCache
	}
	return &cred, nil
}

func (c *CredentialCache) SaveCredential(cred *types.ExecCredential) error {
	d, err := json.MarshalIndent(cred, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(c.cacheFilePath, d, 0600)
}

func getCacheFilePath(cacheDir string, opts GetCredentialOpts) string {
	addressType := "public"
	if opts.privateIpAddress {
		addressType = "private"
	}
	filename := fmt.Sprintf("%s-%s-exec-auth-credential-%s.json",
		opts.clusterId, addressType, opts.apiVersion)
	return filepath.Join(cacheDir, filename)
}
