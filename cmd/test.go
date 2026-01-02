package cmd

import (
	"encoding/json"

	"github.com/lunarhue/libs-go/log"
	"github.com/lunarhue/metallic-flock/pkg/fingerprint"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Short placeholder",
	Long:  `Long placeholder`,
	Run: func(cmd *cobra.Command, args []string) {
		fingerprint, err := fingerprint.GetFingerprint()
		if err != nil {
			log.Panicf("Fingerprint failed: %v", err)
		}

		jsonResult, err := json.Marshal(fingerprint)
		if err != nil {
			log.Panicf("Failed to marshal json: %v", err)
		}

		//base64Result := base64.StdEncoding.EncodeToString(jsonResult)

		log.Infof("Fingerprint: %s", jsonResult)
	},
}
