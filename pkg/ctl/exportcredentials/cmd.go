package exportcredentials

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/spf13/cobra"
)

type option struct {
	format string
	serve  string
}

type Credentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
	Expiration      string
}

var opt = option{}
var (
	formatAliyunCLIConfigJSON = "aliyun-cli-config-json"
	formatAliyunCLIURIJSON    = "aliyun-cli-uri-json"
	formatECSMetadataJSON     = "ecs-metadata-json"
	formatCredentialFileIni   = "credential-file-ini" // #nosec G101
)
var formats = []string{
	formatAliyunCLIConfigJSON,
	formatAliyunCLIURIJSON,
	formatECSMetadataJSON,
	formatCredentialFileIni,
}

var cmd = &cobra.Command{
	Use:   "export-credentials",
	Short: "Export credentials in various formats",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := ctlcommon.GetClientOrDie()

		if opt.serve == "" {
			output, err := getCredOutput(client)
			ctlcommon.ExitIfError(err)
			fmt.Printf("%s\n", output)
			return
		}

		log.Logger.Warnf("Serving HTTP on %s", opt.serve)
		if err := startCredServer(client); err != http.ErrServerClosed {
			ctlcommon.ExitIfError(err)
		}
	},
}

func startCredServer(client *openapi.Client) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Logger.Info("handel new request")
		output, err := getCredOutput(client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, output)
	})

	return http.ListenAndServe(opt.serve, mux) // #nosec G114
}

// TODO: add cache
func getCredOutput(client *openapi.Client) (string, error) {
	cc := client.Credential()
	ak, err := cc.GetAccessKeyId()
	if err != nil {
		return "", err
	}
	as, err := cc.GetAccessKeySecret()
	if err != nil {
		return "", err
	}
	st, err := cc.GetSecurityToken()
	if err != nil {
		return "", err
	}
	exp := getExpirationWithJitter(time.Now())

	cred := Credentials{
		AccessKeyId:     *ak,
		AccessKeySecret: *as,
		SecurityToken:   *st,
		Expiration:      exp.UTC().Format("2006-01-02T15:04:05Z"),
	}

	output := ""
	switch opt.format {
	case formatCredentialFileIni:
		output = toCredentialFileIni(cred)
	case formatAliyunCLIURIJSON, formatECSMetadataJSON:
		output = toAliyunCLIURIBody(cred)
	default:
		output = toAliyunCLIConfigJSON(cred)
	}
	return output, nil
}

func getExpirationWithJitter(t time.Time) time.Time {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))      // #nosec G404
	jitter := time.Duration(r.Int63n(int64(time.Minute) * 4)) // #nosec G404
	exp := t.Add(time.Minute*8 + jitter)                      // 8 + [0, 4) minutes
	return exp
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)

	cmd.Flags().StringVarP(&opt.format, "format", "f", formatAliyunCLIConfigJSON,
		fmt.Sprintf("The output format to display credentials (%s)",
			strings.Join(formats, " or ")))
	cmd.Flags().StringVarP(&opt.serve, "serve", "s", "",
		"start a server to export credentials")
}
