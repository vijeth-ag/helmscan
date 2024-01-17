package helmscan

import (
	// "scan/internal/log"

	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/chart/loader"
)

type PolicyCheck func(content string) bool

// helm scan
func Scan(chartsPath string) error {

	type Config struct {
		Policies map[string]bool `yaml:"policies"`
	}

	configPath := "scan-policy.yml"
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file: %v\n", err)
		os.Exit(1)
	}

	var config Config

	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling config file: %v\n", err)
		os.Exit(1)
	}

	if chartsPath == "" {
		log.Println("require charts path")
		return errors.New("chartsPath missing")
	}

	log.Println("SCANNIN CHART", chartsPath)

	chart, err := loader.Load(chartsPath)

	if err != nil {
		log.Println("err", err)
	}

	policies := make(map[string]PolicyCheck)
	for policyName, isEnabled := range config.Policies {
		if isEnabled {
			policies[policyName] = createPolicyCheck(policyName)
		}
	}

	for _, tpl := range chart.Templates {
		fmt.Printf("Checking template: %s\n", tpl.Name)
		content := string(tpl.Data)

		// Check against each policy
		for policyName, policyCheck := range policies {
			if !policyCheck(content) {
				s := fmt.Sprintf("Policy violation in template '%s': %s\n", tpl.Name, policyName)
				return errors.New(s)
			}
		}
	}

	return nil
}

func createPolicyCheck(policyName string) PolicyCheck {
	switch policyName {
	case "shouldNotContainSecret":
		return func(content string) bool {
			return !strings.Contains(content, "secret")
		}
	case "shouldNotContainPassword":
		return func(content string) bool {
			return !strings.Contains(content, "password")
		}
	// Add more policies as needed
	default:
		// Default policy check (always true)
		return func(content string) bool {
			return true
		}
	}
}
