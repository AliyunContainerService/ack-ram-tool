package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/aliyun/credentials-go/credentials"
)

type Credentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}

type CredentialsProvider struct {
	cred credentials.Credential
}

func (c *Credentials) GetAccessKeyID() string {
	return c.AccessKeyId
}

func (c *Credentials) GetAccessKeySecret() string {
	return c.AccessKeySecret
}

func (c *Credentials) GetSecurityToken() string {
	return c.SecurityToken
}

func (p CredentialsProvider) GetCredentials() oss.Credentials {
	id, err := p.cred.GetAccessKeyId()
	if err != nil {
		log.Printf("get access key id failed: %+v", err)
		return &Credentials{}
	}
	secret, err := p.cred.GetAccessKeySecret()
	if err != nil {
		log.Printf("get access key secret failed: %+v", err)
		return &Credentials{}
	}
	token, err := p.cred.GetSecurityToken()
	if err != nil {
		log.Printf("get access security token failed: %+v", err)
		return &Credentials{}
	}

	return &Credentials{
		AccessKeyId:     tea.StringValue(id),
		AccessKeySecret: tea.StringValue(secret),
		SecurityToken:   tea.StringValue(token),
	}
}

func NewClient(endpoint string) (*oss.Client, error) {
	cred, err := credentials.NewCredential(&credentials.Config{
		Type:            tea.String("sts"),
		AccessKeyId:     tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")),
		AccessKeySecret: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")),
		SecurityToken:   tea.String(os.Getenv("ALIBABA_CLOUD_SECURITY_TOKEN")),
	})
	if err != nil {
		return nil, err
	}
	provider := &CredentialsProvider{cred: cred}
	client, err := oss.New(endpoint, "", "", oss.SetCredentialsProvider(provider))
	return client, err
}

func UploadFile(bucket *oss.Bucket, baseObjectDir, filePath string) (string, error) {
	objectPath := fmt.Sprintf("%s/%s", baseObjectDir, path.Base(filePath))

	var err error
	maxN := 5
	for i := 0; i < maxN; i++ {
		err = bucket.PutObjectFromFile(objectPath, filePath)
		if err == nil {
			return objectPath, nil
		}
		log.Printf("%d/%d upload file failed: %s", i+1, maxN, err)
		if i < maxN {
			time.Sleep(time.Second * 3 * time.Duration(i+1))
		}
	}
	if err != nil {
		return "", err
	}
	return objectPath, nil
}

func main() {
	endpoint := flag.String("endpoint", "", "")
	bucketName := flag.String("bucket", "", "")
	objectDir := flag.String("objectdir", "", "")
	flag.Parse()

	filepathList := flag.Args()
	if *endpoint == "" || *bucketName == "" || *objectDir == "" || len(filepathList) == 0 {
		log.Fatalln("missing required arguments")
	}

	client, err := NewClient(*endpoint)
	if err != nil {
		log.Fatalf("init client failed: %s", err)
	}
	bucket, err := client.Bucket(*bucketName)
	if err != nil {
		log.Fatalf("init bucket client failed: %s", err)
	}

	for _, ph := range filepathList {
		_, err = UploadFile(bucket, *objectDir, ph)
		if err != nil {
			log.Fatalf("upload %s failed: %s", ph, err)
		}
		log.Printf("uploaded %s to oss", ph)
	}
}
