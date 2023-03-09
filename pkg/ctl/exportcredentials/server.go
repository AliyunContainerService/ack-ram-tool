package exportcredentials

import (
	"fmt"
	"net/http"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/log"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
)

func startCredServer(client *openapi.Client) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Logger.Info("handel new request")
		cred, err := getCredentials(client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		output := cred.Format(opt.format)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, output)
	})

	return http.ListenAndServe(opt.serve, mux) // #nosec G114
}
