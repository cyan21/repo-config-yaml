package commands

import (
	"errors"
	"io/ioutil"
	"strconv"

	"github.com/jfrog/jfrog-cli-core/v2/artifactory/utils"
	"github.com/jfrog/jfrog-cli-core/v2/common/commands"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-core/v2/utils/config"
	"gopkg.in/yaml.v2"

	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/artifactory/services"

	clientutils "github.com/jfrog/jfrog-client-go/utils"

	repoutils "github.com/cyan21/repo-config-yaml/commands/utils"
)

func GetExportCommand() components.Command {
	return components.Command{
		Name:        "export",
		Description: "Export repository definition into a YAML file",
		Aliases:     []string{"hi"},
		Arguments:   getExportArguments(),
		Flags:       getExportFlags(),
		// EnvVars:     getExportEnvVar(),
		Action: func(c *components.Context) error {
			return ExportCmd(c)
		},
	}
}

func getExportArguments() []components.Argument {
	return []components.Argument{
		{
			Name:        "repoType",
			Description: "{all, local, remote, virtual, federated}",
		},
	}
}

func getExportFlags() []components.Flag {
	return []components.Flag{
		components.BoolFlag{
			Name:         "shout",
			Description:  "Makes output uppercase.",
			DefaultValue: false,
		},
		components.StringFlag{
			Name:         "pkgType",
			Description:  "export only specified package type (maven, ...)",
			DefaultValue: "",
		},
		components.StringFlag{
			Name:         "regex-key",
			Description:  "filter repo name on a regex",
			DefaultValue: "*",
		},
		components.StringFlag{
			Name:         "limit",
			Description:  "how many repository to export, 0 means no limit",
			DefaultValue: "0",
		},
		components.StringFlag{
			Name:         "project",
			Description:  "filter repo name based on JFrog project name",
			DefaultValue: "",
		},
	}
}

////////////////// CUSTOM TYPE

var supportedPkgTypes = []string{
	"maven", "gradle", "ivy", "sbt", "docker", "nuget", "npm", "generic", "helm", "bower", "debian", "pypi", "yum",
	// "composer", "gems", "conan",
	// "vagrant", "gitlfs",  "chef","puppet"
}

var supportedArgs = []string{"all", "local", "remote", "virtual"}

type LocalYamlRepositories struct {
	Repos map[string]repoutils.LocalRepoYAML `yaml:"localRepositories"`
}

type RemoteYamlRepositories struct {
	Repos map[string]repoutils.RemoteRepoYAML `yaml:"remoteRepositories"`
}

type VirtualYamlRepositories struct {
	Repos map[string]repoutils.VirtualRepoYAML `yaml:"virtualRepositories"`
}

////////////////// CUSTOM FUNCTIONS

// Returns the Artifactory Details of the provided server-id, or the default one.
func getRtDetails(c *components.Context) (*config.ServerDetails, error) {
	details, err := commands.GetConfig(c.GetStringFlagValue("server-id"), false)
	if err != nil {
		return nil, err
	}
	if details.ArtifactoryUrl == "" {
		return nil, errors.New("no server-id was found, or the server-id has no Artifactory url.")
	}
	details.ArtifactoryUrl = clientutils.AddTrailingSlashIfNeeded(details.ArtifactoryUrl)
	err = config.CreateInitialRefreshableTokensIfNeeded(details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

// Generate YAML for a LOCAL repo and add it to the local map
func (r LocalYamlRepositories) add(pkgType string, repoName string, servicesManager *artifactory.ArtifactoryServicesManager) {

	switch pkgType {

	case "Maven", "Gradle", "Ivy", "Sbt":
		theRepo := services.MavenLocalRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		// fmt.Printf("%+v\n", theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForJavaLocal(&theRepo)

	case "Cargo":
		theRepo := services.CargoLocalRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForCargoLocal(&theRepo)

	case "Debian":
		theRepo := services.DebianLocalRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForDebianLocal(&theRepo)

	case "Docker":
		theRepo := services.DockerLocalRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForDockerLocal(&theRepo)

	case "Nuget":
		theRepo := services.NugetLocalRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForNugetLocal(&theRepo)

	case "Rpm", "Yum":
		theRepo := services.RpmLocalRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForRpmLocal(&theRepo)

	default:
		theRepo := services.LocalRepositoryBaseParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForLocal(&theRepo)
	}
}

// Generate YAML for a REMOTE repo and add it to the local map
func (r RemoteYamlRepositories) add(pkgType string, repoName string, servicesManager *artifactory.ArtifactoryServicesManager) {

	switch pkgType {

	case "Maven", "Gradle", "Ivy", "Sbt":
		theRepo := services.MavenRemoteRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForJavaRemote(&theRepo)

	case "Docker":
		theRepo := services.DockerRemoteRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForDockerRemote(&theRepo)

	case "Nuget":
		theRepo := services.NugetRemoteRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForNugetRemote(&theRepo)

	// case "Rpm", "Yum":
	// 	theRepo := services.RpmLocalRepositoryParams{}
	// 	(*servicesManager).GetRepository(repoName, &theRepo)
	// 	r.Repos[repoName] = repoutils.SetYAMLforRpm(&theRepo)

	default:
		theRepo := services.RemoteRepositoryBaseParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLRemote(&theRepo)
	}
}

// Generate YAML for a VIRTUAL repo and add it to the local map
func (r VirtualYamlRepositories) add(pkgType string, repoName string, servicesManager *artifactory.ArtifactoryServicesManager) {

	switch pkgType {

	case "Maven", "Gradle", "Ivy", "Sbt":
		theRepo := services.MavenVirtualRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForJavaVirtual(&theRepo)

	case "Docker":
		theRepo := services.DockerVirtualRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForDockerVirtual(&theRepo)

	case "Nuget":
		theRepo := services.NugetVirtualRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForNugetVirtual(&theRepo)

	case "Npm":
		theRepo := services.NpmVirtualRepositoryParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForNpmVirtual(&theRepo)

	default:
		theRepo := services.VirtualRepositoryBaseParams{}
		(*servicesManager).GetRepository(repoName, &theRepo)
		r.Repos[repoName] = repoutils.SetYAMLForVirtual(&theRepo)
	}
}

////////////////// COMMAND FUNCTION

func ExportCmd(c *components.Context) error {

	if len(c.Arguments) != 1 {
		return errors.New("Wrong number of arguments. Expected: 1, " + "Received: " + strconv.Itoa(len(c.Arguments)))
	}

	switch c.Arguments[0] {
	case "all", "local", "remote", "virtual":
	default:
		return errors.New("Only the following arguments are valid : local, remote, virtual")
	}

	var repos *[]services.RepositoryDetails
	var localRepos LocalYamlRepositories
	var remoteRepos RemoteYamlRepositories
	var virtualRepos VirtualYamlRepositories

	var data []byte
	rtDetails, err := getRtDetails(c)

	if err != nil {
		return err
	}

	servicesManager, err := utils.CreateServiceManager(rtDetails, -1, false)

	// get repositories based on RepoType and PackageType
	if c.Arguments[0] == "all" {
		repos, _ = servicesManager.GetAllRepositories()
	} else {
		params := services.NewRepositoriesFilterParams()
		params.RepoType = c.Arguments[0]
		// params.PackageType = "maven"
		if c.GetStringFlagValue("pkgType") != "" {
			params.PackageType = c.GetStringFlagValue("pkgType")
		}
		repos, _ = servicesManager.GetAllRepositoriesFiltered(params)
	}

	// init map
	switch c.Arguments[0] {
	case "all":
		localRepos.Repos = make(map[string]repoutils.LocalRepoYAML)
		remoteRepos.Repos = make(map[string]repoutils.RemoteRepoYAML)
		virtualRepos.Repos = make(map[string]repoutils.VirtualRepoYAML)

	case "local":
		localRepos.Repos = make(map[string]repoutils.LocalRepoYAML)

	case "remote":
		remoteRepos.Repos = make(map[string]repoutils.RemoteRepoYAML)

	case "virtual":
		virtualRepos.Repos = make(map[string]repoutils.VirtualRepoYAML)
	}

	// loop on repositories
	for _, v := range *repos {
		// log.Output("Name :" + v.Key + ", Type: " + v.Type + ", PackageType:" + v.PackageType)

		// set specific info depending on PackageType
		switch v.Type {
		case "LOCAL":
			localRepos.add(v.PackageType, v.Key, &servicesManager)
		case "REMOTE":
			remoteRepos.add(v.PackageType, v.Key, &servicesManager)
		case "VIRTUAL":
			virtualRepos.add(v.PackageType, v.Key, &servicesManager)
		}
	}

	// generate YAML file
	switch c.Arguments[0] {
	case "all":
		data, _ = yaml.Marshal(localRepos)
		_ = ioutil.WriteFile("gen-local-repos.yml", data, 0644)

		data, _ = yaml.Marshal(remoteRepos)
		_ = ioutil.WriteFile("gen-remote-repos.yml", data, 0644)

		data, _ = yaml.Marshal(virtualRepos)
		_ = ioutil.WriteFile("gen-virtual-repos.yml", data, 0644)

	case "local":
		data, _ = yaml.Marshal(localRepos)
		_ = ioutil.WriteFile("gen-local-repos.yml", data, 0644)

	case "remote":
		data, _ = yaml.Marshal(remoteRepos)
		_ = ioutil.WriteFile("gen-remote-repos.yml", data, 0644)

	case "virtual":
		data, _ = yaml.Marshal(virtualRepos)
		_ = ioutil.WriteFile("gen-virtual-repos.yml", data, 0644)

	}

	return nil
}
