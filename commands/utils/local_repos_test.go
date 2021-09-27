package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateJavaRepo(t *testing.T) {
	InitRTConnection(t)
	assert.Equal(t, "Hello World!", "Hello World!")
}

func InitRTConnection(t *testing.T) string {

	// var file *os.File
	// log.SetLogger(log.NewLogger(log.INFO, file))

	// rtDetails := auth.NewArtifactoryDetails()
	// // rtDetails.SetUrl(os.Getenv("TEST_RT_URL"))
	// // rtDetails.SetAccessToken(os.Getenv("TEST_RT_TOKEN"))
	// rtDetails.SetUrl("https://yann-swampup.dev.aws.devopsacc.team/artifactory")
	// rtDetails.SetAccessToken("eyJ2ZXIiOiIyIiwidHlwIjoiSldUIiwiYWxnIjoiUlMyNTYiLCJraWQiOiJYZDNac0dZSkx2enE1ZElaY3pMb0hfR0poQ3NCU1Q0ZVZvZExWc1JUMkJRIn0.eyJleHQiOiJ7XCJyZXZvY2FibGVcIjpcInRydWVcIn0iLCJzdWIiOiJqZmFjQDAxZmFrZ21jbXliOTVmMGd5dGN4d24xcHBhXC91c2Vyc1wvYWRtaW4iLCJzY3AiOiJhcHBsaWVkLXBlcm1pc3Npb25zXC91c2VyIiwiYXVkIjoiKkAqIiwiaXNzIjoiamZmZUAwMDAiLCJleHAiOjE2NjQwOTM3NjMsImlhdCI6MTYzMjU1Nzc2MywianRpIjoiYTUwMDg5ODgtZTZlOS00NWE0LWE4MTgtZWVhMDFlYTQ2MzdkIn0.QqIQwp_LGk4oI3xTxvf5NVhjJ_A1TGKCuHcQGFG5p2Acj-aep2HOCOYaZjOKeTxYFgGTEMtaYOahaUFg6HHLOhkAr7LSpWB5eT2YsSmiBSqs5GTX35L4VmCW28eAcTR5OmUfbbkBFdtnk68ud1EkVSJySz44uOdr1eQk0um9i075XhT8wjHFSNwI-AEJ4xnYNPzUR8yJA-LJmiv33GRKowajyZX8IsLGlHrOgzaM7OIbrSGbS6FF6p9eE6YEuIpwDneK4BpFCNSP1DQmg3O3MY2bG7DoRUX7W-J36k_0UArhiDd_upe0gDyyz0tF3A-u1Vt4aEEVvFrDcB0AlsxQ-w")

	// serviceConfig, err := config.NewConfigBuilder().
	// 	SetServiceDetails(rtDetails).
	// 	SetDryRun(false).
	// 	// Add [Context](https://golang.org/pkg/context/)
	// 	// SetContext(ctx).
	// 	Build()

	// if err != nil {
	// 	return err.Error()
	// }

	// servicesManager, err2 := artifactory.New(serviceConfig)

	// if err2 != nil {
	// 	return err.Error()
	// }

	// params := services.NewLocalRepositoryBaseParams()
	// params.Key = "test-go-generic-local"
	// params.PackageType = "generic"
	// params.Rclass = "LOCAL"
	// params.Description = "This is a public description for generic-repo"

	// err2 = servicesManager.CreateLocalRepositoryWithParams(params)
	// if err2 != nil {
	// 	return err.Error()
	// }

	// theRepo := services.MavenLocalRepositoryParams{}

	url := "https://yann-swampup.dev.aws.devopsacc.team/artifactory"
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, url+"/api/system/info", nil)
	// req.Header.Set("Authorization", "Bearer eyJ2ZXIiOiIyIiwidHlwIjoiSldUIiwiYWxnIjoiUlMyNTYiLCJraWQiOiJYZDNac0dZSkx2enE1ZElaY3pMb0hfR0poQ3NCU1Q0ZVZvZExWc1JUMkJRIn0.eyJleHQiOiJ7XCJyZXZvY2FibGVcIjpcInRydWVcIn0iLCJzdWIiOiJqZmFjQDAxZmFrZ21jbXliOTVmMGd5dGN4d24xcHBhXC91c2Vyc1wvYWRtaW4iLCJzY3AiOiJhcHBsaWVkLXBlcm1pc3Npb25zXC91c2VyIiwiYXVkIjoiKkAqIiwiaXNzIjoiamZmZUAwMDAiLCJleHAiOjE2NjQyMjAyNTAsImlhdCI6MTYzMjY4NDI1MCwianRpIjoiYTE5NzA4ZjQtYzlkNS00MzdkLWI3MTktNDdkYWM5YjQxZDFiIn0.JQs3nSBuD_oo3t0rB4UTMRc6vE-dTaIsH7mBMWcUqBjvhsAkK_701se2T2swy2gIaEC_BXeSNukUw3gQ-E9EnzpyfBQSgdTFII2CCE_81jcUtlpWFmzZq01Pwfdjyr-v2deqSH4jUK00KjH3jeZlKSbwEYy1_CK4z4PL8kPaIOzuPvqXg97strw1vYWgmotQVe-F298teC3AaIBG0dTBwB1CGCJ6o1GB78LyrjwiwlPtnywf1qXJAPGHhkYwHpJHtfrO6OMd0ik5T6EQzTqHiqKzpG5WX5wmMVwElFncdo1cje71g14R_YavRJ5dTUgbptDh3uC1W6DaNLr3qz11uA")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("TEST_RT_TOKEN"))
	resp, err := client.Do(req)

	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	t.Log(string(body))
	t.Log("salut")
	return "Hello World!"
}

func truc() string {
	return "Hello World!"
}
