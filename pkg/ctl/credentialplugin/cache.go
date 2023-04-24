package credentialplugin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
)

const (
	minExpirationDelta = time.Minute * 10
)

var defaultCacheDir = filepath.Join("~", ".kube", "cache", "ack-ram-tool", "credential-plugin")
var (
	errNoValidCache     = errors.New("no valid cache")
	errNeedRefreshCache = errors.New("need refresh cache")
)

type CredentialCache struct {
	cacheFilePath   string
	expirationDelta time.Duration
}

func NewCredentialCache(cacheDir string, opts GetCredentialOpts) *CredentialCache {
	c := &CredentialCache{
		cacheFilePath: getCacheFilePath(cacheDir, opts),
		//expirationDelta: opts.expirationDelta,
	}
	expirationDelta := time.Duration(int64(float64(opts.temporaryDuration) * 0.2))
	if expirationDelta < minExpirationDelta {
		expirationDelta = minExpirationDelta
	}
	log.Logger.Debugf("will use %s as expirationDelta", expirationDelta)
	c.expirationDelta = expirationDelta
	return c
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
		return nil, errNoValidCache
	}
	remain := time.Until(cred.Status.ExpirationTimestamp.Time)
	if remain <= 0 {
		return nil, errNoValidCache
	} else if remain <= c.expirationDelta {
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
	filename := fmt.Sprintf("%s-exec-auth-credential-%s.json",
		opts.clusterId, opts.apiVersion)
	return filepath.Join(cacheDir, filename)
}
