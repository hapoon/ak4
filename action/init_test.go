package action

import (
	"fmt"
	"testing"
)

func Test_Config_Load(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		profile string
	}{
		"profile empty": {
			profile: "",
		},
		"profile not empty": {
			profile: "aaa",
		},
	}

	for scenario, test := range tests {
		cfg := Config{}
		if err := cfg.Load(test.profile); err != nil {
			t.Fatal("Config loading failed:", scenario)
		}
		fmt.Printf("config: %+v\n", cfg)
	}
	t.Fatal("success")
}
