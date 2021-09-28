package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/artifactory/auth"
	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-client-go/config"
	"github.com/jfrog/jfrog-client-go/utils/log"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

type LocalYamlRepositories struct {
	Repos map[string]LocalRepoYAML `yaml:"localRepositories"`
}

var layout = map[string]string{
	"composer": "composer-default",
	"conan":    "conan-default",
	"debian":   "simple-default",
	"docker":   "simple-default",
	"generic":  "simple-default",
	"gems":     "simple-default",
	"go":       "go-default",
	"gradle":   "maven-2-default",
	"helm":     "simple-default",
	"maven":    "maven-2-default",
	"ivy":      "ivy-default",
	"npm":      "simple-default",
	"nuget":    "nuget-default",
	"pypi":     "simple-default",
	"yum":      "simple-default",
	"sbt":      "sbt-default",
}

func TestCreateDockerRepo(t *testing.T) {

	if !RtInfoAvailable() {
		t.Log("environment variables TEST_RT_URL and TEST_RT_URL have to be set")
		t.FailNow()
	}

	var testLocalRepos LocalYamlRepositories
	testLocalRepos.Repos = make(map[string]LocalRepoYAML)
	myBoolValue := true
	pkgType := "docker"
	repoKey := "test-" + pkgType + "-release-local"
	autoDelete := true

	// prepare dummy repo
	dockerLocalRepo := services.DockerLocalRepositoryParams{}
	dockerLocalRepo.RepositoryBaseParams.Rclass = "LOCAL"
	dockerLocalRepo.RepositoryBaseParams.Key = repoKey
	dockerLocalRepo.RepositoryBaseParams.PackageType = pkgType
	dockerLocalRepo.RepositoryBaseParams.RepoLayoutRef = layout[pkgType]
	// any pointer has to be valued otherwise the test fails
	dockerLocalRepo.ArchiveBrowsingEnabled = &myBoolValue
	dockerLocalRepo.BlackedOut = &myBoolValue
	dockerLocalRepo.DownloadRedirect = &myBoolValue
	dockerLocalRepo.PriorityResolution = &myBoolValue
	dockerLocalRepo.PropertySets = nil
	dockerLocalRepo.XrayIndex = &myBoolValue
	// specific
	dockerLocalRepo.BlockPushingSchema1 = &myBoolValue

	testLocalRepos.Repos[repoKey] = SetYAMLForDockerLocal(&dockerLocalRepo)
	data, _ := yaml.Marshal(testLocalRepos)
	t.Log(string(data))

	result := RunPatchConfig(os.Getenv("TEST_RT_URL"), os.Getenv("TEST_RT_TOKEN"), repoKey, data, autoDelete, t)

	assert.Contains(t, string(result), "successfully")
}

func TestCreateGenericRepo(t *testing.T) {

	if !RtInfoAvailable() {
		t.Log("environment variables TEST_RT_URL and TEST_RT_URL have to be set")
		t.FailNow()
	}

	var testLocalRepos LocalYamlRepositories
	testLocalRepos.Repos = make(map[string]LocalRepoYAML)
	myBoolValue := true
	repoKey := "test-generic-release-local"
	autoDelete := true

	// prepare dummy repo
	gnrLocalRepo := services.LocalRepositoryBaseParams{}
	gnrLocalRepo.RepositoryBaseParams.Rclass = "LOCAL"
	gnrLocalRepo.RepositoryBaseParams.Key = repoKey
	gnrLocalRepo.RepositoryBaseParams.PackageType = "generic"
	gnrLocalRepo.RepositoryBaseParams.RepoLayoutRef = "simple-default"
	// any pointer has to be valued otherwise the test fails
	gnrLocalRepo.ArchiveBrowsingEnabled = &myBoolValue
	gnrLocalRepo.BlackedOut = &myBoolValue
	gnrLocalRepo.DownloadRedirect = &myBoolValue
	gnrLocalRepo.PriorityResolution = &myBoolValue
	gnrLocalRepo.PropertySets = nil
	gnrLocalRepo.XrayIndex = &myBoolValue

	testLocalRepos.Repos[repoKey] = SetYAMLForLocal(&gnrLocalRepo)
	data, _ := yaml.Marshal(testLocalRepos)
	t.Log(string(data))

	result := RunPatchConfig(os.Getenv("TEST_RT_URL"), os.Getenv("TEST_RT_TOKEN"), repoKey, data, autoDelete, t)

	assert.Contains(t, string(result), "successfully")
}

func TestCreateMavenRepo(t *testing.T) {

	if !RtInfoAvailable() {
		t.Log("environment variables TEST_RT_URL and TEST_RT_URL have to be set")
		t.FailNow()
	}

	var testLocalRepos LocalYamlRepositories
	testLocalRepos.Repos = make(map[string]LocalRepoYAML)
	myBoolValue := true
	pkgType := "maven"
	repoKey := "test-" + pkgType + "-release-local"
	autoDelete := true

	// prepare dummy repo
	mvnLocalRepo := services.MavenLocalRepositoryParams{}
	mvnLocalRepo.RepositoryBaseParams.Rclass = "LOCAL"
	mvnLocalRepo.RepositoryBaseParams.Key = repoKey
	mvnLocalRepo.RepositoryBaseParams.PackageType = pkgType
	mvnLocalRepo.RepositoryBaseParams.RepoLayoutRef = layout[pkgType]
	// any pointer has to be valued otherwise the test fails
	mvnLocalRepo.ArchiveBrowsingEnabled = &myBoolValue
	mvnLocalRepo.BlackedOut = &myBoolValue
	mvnLocalRepo.DownloadRedirect = &myBoolValue
	mvnLocalRepo.PriorityResolution = &myBoolValue
	mvnLocalRepo.PropertySets = nil
	mvnLocalRepo.XrayIndex = &myBoolValue
	// specific
	mvnLocalRepo.HandleReleases = &myBoolValue
	mvnLocalRepo.HandleSnapshots = &myBoolValue
	mvnLocalRepo.SuppressPomConsistencyChecks = &myBoolValue

	testLocalRepos.Repos[repoKey] = SetYAMLForJavaLocal(&mvnLocalRepo)
	data, _ := yaml.Marshal(testLocalRepos)
	t.Log(string(data))

	result := RunPatchConfig(os.Getenv("TEST_RT_URL"), os.Getenv("TEST_RT_TOKEN"), repoKey, data, autoDelete, t)

	assert.Contains(t, string(result), "successfully")
}

func TestCreateNugetRepo(t *testing.T) {

	if !RtInfoAvailable() {
		t.Log("environment variables TEST_RT_URL and TEST_RT_URL have to be set")
		t.FailNow()
	}

	var testLocalRepos LocalYamlRepositories
	testLocalRepos.Repos = make(map[string]LocalRepoYAML)
	myBoolValue := true
	pkgType := "nuget"
	repoKey := "test-" + pkgType + "-release-local"
	autoDelete := true

	// prepare dummy repo
	nugetLocalRepo := services.NugetLocalRepositoryParams{}
	nugetLocalRepo.RepositoryBaseParams.Rclass = "LOCAL"
	nugetLocalRepo.RepositoryBaseParams.Key = repoKey
	nugetLocalRepo.RepositoryBaseParams.PackageType = pkgType
	nugetLocalRepo.RepositoryBaseParams.RepoLayoutRef = layout[pkgType]
	// any pointer has to be valued otherwise the test fails
	nugetLocalRepo.ArchiveBrowsingEnabled = &myBoolValue
	nugetLocalRepo.BlackedOut = &myBoolValue
	nugetLocalRepo.DownloadRedirect = &myBoolValue
	nugetLocalRepo.PriorityResolution = &myBoolValue
	nugetLocalRepo.PropertySets = nil
	nugetLocalRepo.XrayIndex = &myBoolValue
	// specific
	nugetLocalRepo.ForceNugetAuthentication = &myBoolValue

	testLocalRepos.Repos[repoKey] = SetYAMLForNugetLocal(&nugetLocalRepo)
	data, _ := yaml.Marshal(testLocalRepos)
	t.Log(string(data))

	result := RunPatchConfig(os.Getenv("TEST_RT_URL"), os.Getenv("TEST_RT_TOKEN"), repoKey, data, autoDelete, t)

	assert.Contains(t, string(result), "successfully")
}

////////////////// HELPER FUNCTIONS

func RtInfoAvailable() bool {

	result := true

	if os.Getenv("TEST_RT_URL") == "" {
		result = false
	} else {
		if os.Getenv("TEST_RT_TOKEN") == "" {
			result = false
		}
	}

	return result
}

func RunPatchConfig(url string, token string, repoKey string, data []byte, autoDelete bool, t *testing.T) string {

	result := ""

	// prepare HTTP request
	client := &http.Client{}
	// req, _ := http.NewRequest(http.MethodGet, url+"/api/system/info", nil)
	req, _ := http.NewRequest(http.MethodPatch, url+"/api/system/configuration", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/yaml")
	req.Header.Set("Authorization", "Bearer "+token)

	// run query and parse it
	resp, err := client.Do(req)

	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		result = string(err.Error())
	} else {
		result = string(body)
	}

	// t.Log(string(body))
	// t.Log(resp.StatusCode)

	// cleanup if successful query
	if autoDelete {
		CleanupTestRepo(url, token, repoKey, t)
	}

	if err != nil {
		t.Log(err.Error())
	}

	return result
}

func CleanupTestRepo(url string, token string, repoKey string, t *testing.T) error {

	var file *os.File
	var servicesManager artifactory.ArtifactoryServicesManager

	// mandatory to init the client
	log.SetLogger(log.NewLogger(log.INFO, file))

	rtDetails := auth.NewArtifactoryDetails()
	rtDetails.SetUrl(url + "/")
	rtDetails.SetAccessToken(token)
	serviceConfig, err := config.NewConfigBuilder().
		SetServiceDetails(rtDetails).
		SetDryRun(false).
		Build()

	if err != nil {
		return err
	}

	servicesManager, err = artifactory.New(serviceConfig)

	if err != nil {
		return err
	}

	t.Log("Deleting repo " + repoKey + " ... ")
	time.Sleep(5 * time.Second)

	err = servicesManager.DeleteRepository(repoKey)

	if err != nil {
		t.Log(err.Error())

		return err
	}
	t.Log("Repo " + repoKey + " was successfully deleted")

	return nil
}
