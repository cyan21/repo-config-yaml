package utils

import "github.com/jfrog/jfrog-client-go/artifactory/services"

type CacheConfig struct {
	VirtualRetrievalCachePeriodSecs int `yaml:"virtualRetrievalCachePeriodSecs"` // default 600
}

type NpmSpecificities struct {
	Enabled  bool     `yaml:"enabled"`
	Patterns []string `yaml:"patterns"`
}

type P2Specificities struct {
	Urls []string `yaml:"urls"`
}

type VirtualRepoYAML struct {
	Cache                                CacheConfig      `yaml:"virtualCacheConfig,omitempty"`
	CanRetrieveRemoteArtifacts           bool             `yaml:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	Description                          string           `yaml:"description"`
	DockerApiVersion                     string           `yaml:"dockerApiVersion,omitempty"`
	ExcludesPattern                      string           `yaml:"excludesPattern"`
	ExternalDependencies                 NpmSpecificities `yaml:"externalDependencies,omitempty"`
	ForceNugetAuthentication             bool             `yaml:"forceNugetAuthentication,omitempty"`
	HandleReleases                       bool             `yaml:"handleReleases,omitempty"`
	HandleSnapshots                      bool             `yaml:"handleSnapshots,omitempty"`
	IncludesPattern                      string           `yaml:"includesPattern"`
	KeyPair                              string           `yaml:"keyPair,omitempty"`
	PomRepositoryReferencesCleanupPolicy string           `yaml:"PomRepositoryReferencesCleanupPolicy,omitempty"`
	Repositories                         []string         `yaml:"repositories"`
	Notes                                string           `yaml:"notes"`
	P2                                   P2Specificities  `yaml:"p2,omitempty"`
	RepoLayout                           string           `yaml:"repoLayout"`
	Type                                 string           `yaml:"type"` // PackageType
}

func SetYAMLForVirtual(repo *services.VirtualRepositoryBaseParams) VirtualRepoYAML {
	return VirtualRepoYAML{
		CanRetrieveRemoteArtifacts: *repo.ArtifactoryRequestsCanRetrieveRemoteArtifacts,
		Description:                (*repo).Description,
		ExcludesPattern:            (*repo).ExcludesPattern,
		IncludesPattern:            (*repo).IncludesPattern,
		RepoLayout:                 (*repo).RepoLayoutRef,
		Repositories:               (*repo).Repositories,
		Notes:                      (*repo).Notes,
		Type:                       (*repo).PackageType,
	}
}

func SetYAMLForDockerVirtual(repo *services.DockerVirtualRepositoryParams) VirtualRepoYAML {
	return VirtualRepoYAML{
		CanRetrieveRemoteArtifacts: *repo.ArtifactoryRequestsCanRetrieveRemoteArtifacts,
		Description:                (*repo).Description,
		ExcludesPattern:            (*repo).ExcludesPattern,
		IncludesPattern:            (*repo).IncludesPattern,
		Notes:                      (*repo).Notes,
		Repositories:               (*repo).Repositories,
		RepoLayout:                 (*repo).RepoLayoutRef,
		Type:                       (*repo).PackageType,

		// Docker specificities
		// DockerApiVersion                     : (*repo).
	}
}

func SetYAMLForJavaVirtual(repo *services.MavenVirtualRepositoryParams) VirtualRepoYAML {
	return VirtualRepoYAML{
		CanRetrieveRemoteArtifacts: *repo.ArtifactoryRequestsCanRetrieveRemoteArtifacts,
		Description:                (*repo).Description,
		ExcludesPattern:            (*repo).ExcludesPattern,
		IncludesPattern:            (*repo).IncludesPattern,
		RepoLayout:                 (*repo).RepoLayoutRef,
		Repositories:               (*repo).Repositories,
		Notes:                      (*repo).Notes,
		Type:                       (*repo).PackageType,
		// Java specificities
		// HandleReleases                       *repo
		// HandleSnapshots                      *repo
		// KeyPair                              (*repo).
		// PomRepositoryReferencesCleanupPolicy (*repo).

	}
}

func SetYAMLForNpmVirtual(repo *services.NpmVirtualRepositoryParams) VirtualRepoYAML {
	return VirtualRepoYAML{
		CanRetrieveRemoteArtifacts: *repo.ArtifactoryRequestsCanRetrieveRemoteArtifacts,
		Description:                (*repo).Description,
		ExcludesPattern:            (*repo).ExcludesPattern,
		IncludesPattern:            (*repo).IncludesPattern,
		Notes:                      (*repo).Notes,
		Repositories:               (*repo).Repositories,
		RepoLayout:                 (*repo).RepoLayoutRef,
		Type:                       (*repo).PackageType,

		// NPM specificities
		ExternalDependencies: NpmSpecificities{},
	}
}

func SetYAMLForNugetVirtual(repo *services.NugetVirtualRepositoryParams) VirtualRepoYAML {
	return VirtualRepoYAML{
		CanRetrieveRemoteArtifacts: *repo.ArtifactoryRequestsCanRetrieveRemoteArtifacts,
		Description:                (*repo).Description,
		ExcludesPattern:            (*repo).ExcludesPattern,
		IncludesPattern:            (*repo).IncludesPattern,
		Notes:                      (*repo).Notes,
		Repositories:               (*repo).Repositories,
		RepoLayout:                 (*repo).RepoLayoutRef,
		Type:                       (*repo).PackageType,

		// Nuget specificities
		// ForceNugetAuthentication             : *repo.

	}
}

func SetYAMLForP2Virtual(repo *services.P2VirtualRepositoryParams) VirtualRepoYAML {
	return VirtualRepoYAML{
		CanRetrieveRemoteArtifacts: *repo.ArtifactoryRequestsCanRetrieveRemoteArtifacts,
		Description:                (*repo).Description,
		ExcludesPattern:            (*repo).ExcludesPattern,
		IncludesPattern:            (*repo).IncludesPattern,
		Notes:                      (*repo).Notes,
		Repositories:               (*repo).Repositories,
		RepoLayout:                 (*repo).RepoLayoutRef,
		Type:                       (*repo).PackageType,

		// P2 specificities
		// P2                                   : P2Specificities

	}
}
