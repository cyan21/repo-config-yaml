package utils

import (
	"github.com/jfrog/jfrog-client-go/artifactory/services"
)

////////////////// CUSTOM TYPES

type Enablement struct {
	Enabled bool `yaml:"enabled"`
}

// DUE TO YAML MARSHALLING WE CANT FACTORED THE STRUCTURE WITH SUB STRUCTURE
// THE OUTPUT YAML WILL BE DIFFERENT

type LocalRepoYAML struct {
	ArchiveBrowsingEnabled       bool       `yaml:"archiveBrowsingEnabled"`
	BlackedOut                   bool       `yaml:"blackedOut"`
	BlockPushingSchema1          bool       `yaml:"blockPushingSchema1,omitempty"`
	CalculateYumMetadata         bool       `yaml:"calculateYumMetadata,omitempty"`
	CargoAnonymousAccess         bool       `yaml:"cargoAnonymousAccess,omitempty"`
	ChecksumPolicyType           string     `yaml:"checksumPolicyType,omitempty"`
	DdebSupported                bool       `yaml:"ddebSupported,omitempty"`
	DebianTrivialLayout          bool       `yaml:"debianTrivialLayout,omitempty"`
	Description                  string     `yaml:"description"`
	DockerApiVersion             string     `yaml:"dockerApiVersion,omitempty"`
	DockerTagRetention           int        `yaml:"dockerTagRetention,omitempty"`
	DownloadRedirect             Enablement `yaml:"downloadRedirect"`
	ExcludesPattern              string     `yaml:"excludesPattern"`
	ForceNugetAuthentication     bool       `yaml:"forceNugetAuthentication,omitempty"`
	HandleReleases               bool       `yaml:"handleReleases,omitempty"`
	HandleSnapshots              bool       `yaml:"handleSnapshots,omitempty"`
	IncludesPattern              string     `yaml:"includesPattern"`
	MaxUniqueSnapshots           int        `yaml:"maxUniqueSnapshots,omitempty"`
	MaxUniqueTags                int        `yaml:"maxUniqueTags,omitempty"`
	Notes                        string     `yaml:"notes"`
	PriorityResolution           bool       `yaml:"priorityResolution"`
	PropertySets                 []string   `yaml:"propertySets"`
	RepoLayout                   string     `yaml:"repoLayout"`
	SnapshotVersionBehavior      string     `yaml:"snapshotVersionBehavior,omitempty"`
	SuppressPomConsistencyChecks bool       `yaml:"suppressPomConsistencyChecks,omitempty"`
	Type                         string     `yaml:"type"` // PackageType
	Xray                         Enablement `yaml:"xray"`
	YumRootDepth                 int        `yaml:"yumRootDepth,omitempty"`
}

////////////////// FUNCTIONS

func SetYAMLForLocal(repo *services.LocalRepositoryBaseParams) LocalRepoYAML {
	return LocalRepoYAML{
		ArchiveBrowsingEnabled: *repo.ArchiveBrowsingEnabled,
		BlackedOut:             *repo.BlackedOut,
		Description:            (*repo).Description,
		DownloadRedirect:       Enablement{*repo.DownloadRedirect},
		ExcludesPattern:        (*repo).ExcludesPattern,
		IncludesPattern:        (*repo).IncludesPattern,
		Notes:                  (*repo).Notes,
		PriorityResolution:     *repo.PriorityResolution,
		PropertySets:           (*repo).PropertySets,
		RepoLayout:             (*repo).RepoLayoutRef,
		Type:                   (*repo).PackageType,
		Xray:                   Enablement{*repo.XrayIndex},
	}
}

func SetYAMLforCargoForLocal(repo *services.CargoLocalRepositoryParams) LocalRepoYAML {
	return LocalRepoYAML{
		ArchiveBrowsingEnabled: *repo.ArchiveBrowsingEnabled,
		BlackedOut:             *repo.BlackedOut,
		Description:            (*repo).Description,
		DownloadRedirect:       Enablement{*repo.DownloadRedirect},
		ExcludesPattern:        (*repo).ExcludesPattern,
		IncludesPattern:        (*repo).IncludesPattern,
		Notes:                  (*repo).Notes,
		PriorityResolution:     *repo.PriorityResolution,
		PropertySets:           (*repo).PropertySets,
		RepoLayout:             (*repo).RepoLayoutRef,
		Type:                   (*repo).PackageType,
		Xray:                   Enablement{*repo.XrayIndex},
		// Cargo specific
		CargoAnonymousAccess: *repo.CargoAnonymousAccess,
	}
}

func SetYAMLforDebianForLocal(repo *services.DebianLocalRepositoryParams) LocalRepoYAML {
	return LocalRepoYAML{
		ArchiveBrowsingEnabled: *repo.ArchiveBrowsingEnabled,
		BlackedOut:             *repo.BlackedOut,
		Description:            (*repo).Description,
		DownloadRedirect:       Enablement{*repo.DownloadRedirect},
		ExcludesPattern:        (*repo).ExcludesPattern,
		IncludesPattern:        (*repo).IncludesPattern,
		Notes:                  (*repo).Notes,
		PriorityResolution:     *repo.PriorityResolution,
		PropertySets:           (*repo).PropertySets,
		RepoLayout:             (*repo).RepoLayoutRef,
		Type:                   (*repo).PackageType,
		Xray:                   Enablement{*repo.XrayIndex},
		// Debian specific
		DebianTrivialLayout: *repo.DebianTrivialLayout,
		// Not supported by CLI
		// DdebSupported:       *repo.DdebSupported,
	}
}

func SetYAMLforDockerForLocal(repo *services.DockerLocalRepositoryParams) LocalRepoYAML {
	return LocalRepoYAML{
		ArchiveBrowsingEnabled: *repo.ArchiveBrowsingEnabled,
		BlackedOut:             *repo.BlackedOut,
		Description:            (*repo).Description,
		DownloadRedirect:       Enablement{*repo.DownloadRedirect},
		ExcludesPattern:        (*repo).ExcludesPattern,
		IncludesPattern:        (*repo).IncludesPattern,
		Notes:                  (*repo).Notes,
		PriorityResolution:     *repo.PriorityResolution,
		PropertySets:           (*repo).PropertySets,
		RepoLayout:             (*repo).RepoLayoutRef,
		Type:                   (*repo).PackageType,
		Xray:                   Enablement{*repo.XrayIndex},
		// Docker specific
		BlockPushingSchema1: *repo.BlockPushingSchema1,
		DockerApiVersion:    (*repo).DockerApiVersion,
		DockerTagRetention:  (*repo).DockerTagRetention,
		MaxUniqueTags:       (*repo).MaxUniqueTags,
	}
}

func SetYAMLforJavaForLocal(repo *services.MavenLocalRepositoryParams) LocalRepoYAML {
	return LocalRepoYAML{
		ArchiveBrowsingEnabled: *repo.ArchiveBrowsingEnabled,
		BlackedOut:             *repo.BlackedOut,
		Description:            (*repo).Description,
		DownloadRedirect:       Enablement{*repo.DownloadRedirect},
		ExcludesPattern:        (*repo).ExcludesPattern,
		IncludesPattern:        (*repo).IncludesPattern,
		Notes:                  (*repo).Notes,
		PriorityResolution:     *repo.PriorityResolution,
		PropertySets:           (*repo).PropertySets,
		RepoLayout:             (*repo).RepoLayoutRef,
		Type:                   (*repo).PackageType,
		Xray:                   Enablement{*repo.XrayIndex},
		// Java specific
		ChecksumPolicyType:           (*repo).ChecksumPolicyType,
		HandleReleases:               *repo.HandleReleases,
		HandleSnapshots:              *repo.HandleSnapshots,
		MaxUniqueSnapshots:           (*repo).MaxUniqueSnapshots,
		SnapshotVersionBehavior:      (*repo).SnapshotVersionBehavior,
		SuppressPomConsistencyChecks: *repo.SuppressPomConsistencyChecks,
	}
}

func SetYAMLforNugetForLocal(repo *services.NugetLocalRepositoryParams) LocalRepoYAML {
	return LocalRepoYAML{
		ArchiveBrowsingEnabled: *repo.ArchiveBrowsingEnabled,
		BlackedOut:             *repo.BlackedOut,
		Description:            (*repo).Description,
		DownloadRedirect:       Enablement{*repo.DownloadRedirect},
		ExcludesPattern:        (*repo).ExcludesPattern,
		IncludesPattern:        (*repo).IncludesPattern,
		Notes:                  (*repo).Notes,
		PriorityResolution:     *repo.PriorityResolution,
		PropertySets:           (*repo).PropertySets,
		RepoLayout:             (*repo).RepoLayoutRef,
		Type:                   (*repo).PackageType,
		Xray:                   Enablement{*repo.XrayIndex},
		// Nuget specific
		ForceNugetAuthentication: *repo.ForceNugetAuthentication,
		MaxUniqueSnapshots:       (*repo).MaxUniqueSnapshots,
	}
}

func SetYAMLforRpmForLocal(repo *services.RpmLocalRepositoryParams) LocalRepoYAML {
	return LocalRepoYAML{
		ArchiveBrowsingEnabled: *repo.ArchiveBrowsingEnabled,
		BlackedOut:             *repo.BlackedOut,
		Description:            (*repo).Description,
		DownloadRedirect:       Enablement{*repo.DownloadRedirect},
		ExcludesPattern:        (*repo).ExcludesPattern,
		IncludesPattern:        (*repo).IncludesPattern,
		Notes:                  (*repo).Notes,
		PriorityResolution:     *repo.PriorityResolution,
		PropertySets:           (*repo).PropertySets,
		RepoLayout:             (*repo).RepoLayoutRef,
		Type:                   (*repo).PackageType,
		Xray:                   Enablement{*repo.XrayIndex},
		// RPM specific
		CalculateYumMetadata: *repo.CalculateYumMetadata,
		YumRootDepth:         (*repo).YumRootDepth,
		// Not supported by CLI
		// YumGroupFileNames:    (*repo).YumGroupFileNames,
	}
}
