package utils

import "github.com/jfrog/jfrog-client-go/artifactory/services"

////////////////// CUSTOM TYPES
type RemoteProvider struct {
	Provider    string `yaml:"provider"`
	DownloadUrl string `yaml:"downloadUrl"`
}

type BowerSpecificities struct {
	BowerRegistryUrl string `yaml:"bowerRegistryUrl"`
}

type CocoaPodsSpecificities struct {
	CocoaPodsSpecsRepoUrl string         `yaml:"cocoaPodsSpecsRepoUrl"`
	SpecRepoProvider      RemoteProvider `yaml:"specRepoProvider"`
}

type ComposerSpecificities struct {
	ComposerRegistryUrl string `yaml:"composerRegistryUrl"`
}

type NugetSpecificities struct {
	DownloadContextPath string `yaml:"downloadContextPath"` // default: api/v2/package
	FeedContextPath     string `yaml:"feedContextPath"`     // default: api/v2
}

type VcsSpecificities struct {
	git  RemoteProvider `yaml:"git"`
	Type string         `yaml:"type"` // default git
}

type RemoteRepoYAML struct {
	AllowAnyHostAuth                  bool                   `yaml:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs          int                    `yaml:"assumedOfflinePeriodSecs"` // default : 300
	BlackedOut                        bool                   `yaml:"blackedOut"`
	BlockMismatchingMimeTypes         bool                   `yaml:"blockMismatchingMimeTypes"` // default : true
	Bower                             BowerSpecificities     `yaml:"bower,omitempty"`
	BypassHeadRequests                bool                   `yaml:"bypassHeadRequests"`
	ClientTlsCertificate              string                 `yaml:"clientTlsCertificate"`
	CocoaPods                         CocoaPodsSpecificities `yaml:"cocoaPods,omitempty"`
	Composer                          ComposerSpecificities  `yaml:"composer,omitempty"`
	Description                       string                 `yaml:"description"`
	DockerApiVersion                  string                 `yaml:"dockerApiVersion,omitempty"`
	DownloadRedirect                  Enablement             `yaml:"downloadRedirect"`
	EnableCookieManagement            bool                   `yaml:"enableCookieManagement"`
	EnableTokenAuthentication         bool                   `yaml:"enableTokenAuthentication"`
	ExcludesPattern                   string                 `yaml:"excludesPattern"`
	FetchJarsEagerly                  bool                   `yaml:"fetchJarsEagerly,omitempty"`
	FetchSourcesEagerly               bool                   `yaml:"fetchSourcesEagerly,omitempty"`
	ForceNugetAuthentication          bool                   `yaml:"forceNugetAuthentication,omitempty"`
	HandleReleases                    bool                   `yaml:"handleReleases,omitempty"`
	HandleSnapshots                   bool                   `yaml:"handleSnapshots,omitempty"`
	IncludesPattern                   string                 `yaml:"includesPattern"`
	ListRemoteFolderItems             bool                   `yaml:"listRemoteFolderItems"` // default : true
	MaxUniqueSnapshots                int                    `yaml:"maxUniqueSnapshots,omitempty"`
	MaxUniqueTags                     int                    `yaml:"maxUniqueTags,omitempty"`
	MismatchingMimeTypesOverrideList  string                 `yaml:"mismatchingMimeTypesOverrideList"`
	MissedRetrievalCachePeriodSecs    int                    `yaml:"missedRetrievalCachePeriodSecs"` // default : 1800
	Notes                             string                 `yaml:"notes"`
	Nuget                             NugetSpecificities     `yaml:"nuget,omitempty"`
	Offline                           bool                   `yaml:"offline"`
	Proxy                             string                 `yaml:"proxy"`
	Password                          string                 `yaml:"password"` // remote URL
	PropertySets                      []string               `yaml:"propertySets"`
	QueryParams                       string                 `yaml:"queryParams"`
	RemoteRepoChecksumPolicyType      string                 `yaml:"remoteRepoChecksumPolicyType,omitempty"`
	RepoLayout                        string                 `yaml:"repoLayout"`
	SocketTimeoutMillis               int                    `yaml:"socketTimeoutMillis"` // default : 15000
	SynchronizeProperties             bool                   `yaml:"synchronizeProperties"`
	Type                              string                 `yaml:"type"` // PackageType
	Url                               string                 `yaml:"url"`  // remote URL
	UnusedArtifactsCleanupPeriodHours int                    `yaml:"unusedArtifactsCleanupPeriodHours"`
	Username                          string                 `yaml:"username"`
	Vcs                               VcsSpecificities       `yaml:"vcs,omitempty"`
	Xray                              Enablement             `yaml:"xray,omitempty"`
}

////////////////// FUNCTIONS

func SetYAMLRemote(repo *services.RemoteRepositoryBaseParams) RemoteRepoYAML {
	return RemoteRepoYAML{
		AllowAnyHostAuth:          *repo.AllowAnyHostAuth,
		AssumedOfflinePeriodSecs:  (*repo).AssumedOfflinePeriodSecs,
		BlackedOut:                *repo.BlackedOut,
		BlockMismatchingMimeTypes: *repo.BlockMismatchingMimeTypes,
		BypassHeadRequests:        *repo.BypassHeadRequests,
		ClientTlsCertificate:      (*repo).ClientTlsCertificate,
		Description:               (*repo).Description,
		DownloadRedirect:          Enablement{*repo.DownloadRedirect},
		EnableCookieManagement:    *repo.EnableCookieManagement,
		// EnableTokenAuthentication         : *repo.
		ExcludesPattern: (*repo).ExcludesPattern,
		IncludesPattern: (*repo).IncludesPattern,
		// ListRemoteFolderItems             : *repo.
		MismatchingMimeTypesOverrideList: (*repo).MismatchingMimeTypesOverrideList,
		MissedRetrievalCachePeriodSecs:   (*repo).MissedRetrievalCachePeriodSecs,
		Notes:                            (*repo).Notes,
		Offline:                          *repo.Offline,
		Proxy:                            (*repo).Proxy,
		Password:                         (*repo).Password,
		PropertySets:                     (*repo).PropertySets,
		// QueryParams:                       (*repo).,
		RepoLayout:                        (*repo).RepoLayoutRef,
		SocketTimeoutMillis:               (*repo).SocketTimeoutMillis,
		SynchronizeProperties:             *repo.SynchronizeProperties,
		Url:                               (*repo).Url,
		UnusedArtifactsCleanupPeriodHours: (*repo).UnusedArtifactsCleanupPeriodHours,
		Username:                          (*repo).Username,
		Xray:                              Enablement{*repo.XrayIndex},
	}
}

func SetYAMLForBowerRemote(repo *services.BowerRemoteRepositoryParams) RemoteRepoYAML {
	return RemoteRepoYAML{
		AllowAnyHostAuth:          *repo.AllowAnyHostAuth,
		AssumedOfflinePeriodSecs:  (*repo).AssumedOfflinePeriodSecs,
		BlackedOut:                *repo.BlackedOut,
		BlockMismatchingMimeTypes: *repo.BlockMismatchingMimeTypes,
		BypassHeadRequests:        *repo.BypassHeadRequests,
		ClientTlsCertificate:      (*repo).ClientTlsCertificate,
		Description:               (*repo).Description,
		DownloadRedirect:          Enablement{*repo.DownloadRedirect},
		EnableCookieManagement:    *repo.EnableCookieManagement,
		// EnableTokenAuthentication         : *repo.
		ExcludesPattern: (*repo).ExcludesPattern,
		IncludesPattern: (*repo).IncludesPattern,
		// ListRemoteFolderItems             : *repo.
		MismatchingMimeTypesOverrideList: (*repo).MismatchingMimeTypesOverrideList,
		MissedRetrievalCachePeriodSecs:   (*repo).MissedRetrievalCachePeriodSecs,
		Notes:                            (*repo).Notes,
		Offline:                          *repo.Offline,
		Proxy:                            (*repo).Proxy,
		Password:                         (*repo).Password,
		PropertySets:                     (*repo).PropertySets,
		// QueryParams:                       (*repo).queryParams,
		RepoLayout:                        (*repo).RepoLayoutRef,
		SocketTimeoutMillis:               (*repo).SocketTimeoutMillis,
		SynchronizeProperties:             *repo.SynchronizeProperties,
		Url:                               (*repo).Url,
		UnusedArtifactsCleanupPeriodHours: (*repo).UnusedArtifactsCleanupPeriodHours,
		Username:                          (*repo).Username,
		Xray:                              Enablement{*repo.XrayIndex},

		// Bower specific
		Bower: BowerSpecificities{(*repo).BowerRegistryUrl},
	}
}

func SetYAMLForComposerRemote(repo *services.ComposerRemoteRepositoryParams) RemoteRepoYAML {
	return RemoteRepoYAML{
		AllowAnyHostAuth:          *repo.AllowAnyHostAuth,
		AssumedOfflinePeriodSecs:  (*repo).AssumedOfflinePeriodSecs,
		BlackedOut:                *repo.BlackedOut,
		BlockMismatchingMimeTypes: *repo.BlockMismatchingMimeTypes,
		BypassHeadRequests:        *repo.BypassHeadRequests,
		ClientTlsCertificate:      (*repo).ClientTlsCertificate,
		Description:               (*repo).Description,
		DownloadRedirect:          Enablement{*repo.DownloadRedirect},
		EnableCookieManagement:    *repo.EnableCookieManagement,
		// EnableTokenAuthentication         : *repo.
		ExcludesPattern: (*repo).ExcludesPattern,
		IncludesPattern: (*repo).IncludesPattern,
		// ListRemoteFolderItems             : *repo.
		MismatchingMimeTypesOverrideList: (*repo).MismatchingMimeTypesOverrideList,
		MissedRetrievalCachePeriodSecs:   (*repo).MissedRetrievalCachePeriodSecs,
		Notes:                            (*repo).Notes,
		Offline:                          *repo.Offline,
		Proxy:                            (*repo).Proxy,
		Password:                         (*repo).Password,
		PropertySets:                     (*repo).PropertySets,
		// QueryParams:                       (*repo).queryParams,
		RepoLayout:                        (*repo).RepoLayoutRef,
		SocketTimeoutMillis:               (*repo).SocketTimeoutMillis,
		SynchronizeProperties:             *repo.SynchronizeProperties,
		Url:                               (*repo).Url,
		UnusedArtifactsCleanupPeriodHours: (*repo).UnusedArtifactsCleanupPeriodHours,
		Username:                          (*repo).Username,
		Xray:                              Enablement{*repo.XrayIndex},

		// Composer specific
		// Composer: ComposerSpecificities{(*repo).}
	}
}

func SetYAMLForJavaRemote(repo *services.MavenRemoteRepositoryParams) RemoteRepoYAML {
	return RemoteRepoYAML{
		AllowAnyHostAuth:          *repo.AllowAnyHostAuth,
		AssumedOfflinePeriodSecs:  (*repo).AssumedOfflinePeriodSecs,
		BlackedOut:                *repo.BlackedOut,
		BlockMismatchingMimeTypes: *repo.BlockMismatchingMimeTypes,
		BypassHeadRequests:        *repo.BypassHeadRequests,
		ClientTlsCertificate:      (*repo).ClientTlsCertificate,
		Description:               (*repo).Description,
		DownloadRedirect:          Enablement{*repo.DownloadRedirect},
		EnableCookieManagement:    *repo.EnableCookieManagement,
		// EnableTokenAuthentication         : *repo.enb
		ExcludesPattern: (*repo).ExcludesPattern,
		IncludesPattern: (*repo).IncludesPattern,
		// ListRemoteFolderItems             : *repo.
		MismatchingMimeTypesOverrideList: (*repo).MismatchingMimeTypesOverrideList,
		MissedRetrievalCachePeriodSecs:   (*repo).MissedRetrievalCachePeriodSecs,
		Notes:                            (*repo).Notes,
		Offline:                          *repo.Offline,
		Proxy:                            (*repo).Proxy,
		Password:                         (*repo).Password,
		PropertySets:                     (*repo).PropertySets,
		// QueryParams:                       (*repo).queryParams,
		RepoLayout:                        (*repo).RepoLayoutRef,
		SocketTimeoutMillis:               (*repo).SocketTimeoutMillis,
		SynchronizeProperties:             *repo.SynchronizeProperties,
		Url:                               (*repo).Url,
		UnusedArtifactsCleanupPeriodHours: (*repo).UnusedArtifactsCleanupPeriodHours,
		Username:                          (*repo).Username,
		Xray:                              Enablement{*repo.XrayIndex},

		// Java specific
		FetchJarsEagerly:    *repo.FetchJarsEagerly,
		FetchSourcesEagerly: *repo.FetchSourcesEagerly,
		// HandleReleases                    : *repo.re,
		// HandleSnapshots                   : *repo.Url,
		RemoteRepoChecksumPolicyType: (*repo).RemoteRepoChecksumPolicyType,
	}
}

func SetYAMLForDockerRemote(repo *services.DockerRemoteRepositoryParams) RemoteRepoYAML {
	return RemoteRepoYAML{
		AllowAnyHostAuth:          *repo.AllowAnyHostAuth,
		AssumedOfflinePeriodSecs:  (*repo).AssumedOfflinePeriodSecs,
		BlackedOut:                *repo.BlackedOut,
		BlockMismatchingMimeTypes: *repo.BlockMismatchingMimeTypes,
		BypassHeadRequests:        *repo.BypassHeadRequests,
		ClientTlsCertificate:      (*repo).ClientTlsCertificate,
		Description:               (*repo).Description,
		DownloadRedirect:          Enablement{*repo.DownloadRedirect},
		EnableCookieManagement:    *repo.EnableCookieManagement,
		// EnableTokenAuthentication         : (*repo).
		ExcludesPattern: (*repo).ExcludesPattern,
		IncludesPattern: (*repo).IncludesPattern,
		// ListRemoteFolderItems             : *repo.
		MismatchingMimeTypesOverrideList: (*repo).MismatchingMimeTypesOverrideList,
		MissedRetrievalCachePeriodSecs:   (*repo).MissedRetrievalCachePeriodSecs,
		Notes:                            (*repo).Notes,
		Offline:                          *repo.Offline,
		Proxy:                            (*repo).Proxy,
		Password:                         (*repo).Password,
		PropertySets:                     (*repo).PropertySets,
		// QueryParams:                       (*repo).queryParams,
		RepoLayout:                        (*repo).RepoLayoutRef,
		SocketTimeoutMillis:               (*repo).SocketTimeoutMillis,
		SynchronizeProperties:             *repo.SynchronizeProperties,
		Url:                               (*repo).Url,
		UnusedArtifactsCleanupPeriodHours: (*repo).UnusedArtifactsCleanupPeriodHours,
		Username:                          (*repo).Username,
		Xray:                              Enablement{*repo.XrayIndex},

		// Docker specific
		// DockerApiVersion:
		// MaxUniqueTags:
	}
}

func SetYAMLForNugetRemote(repo *services.NugetRemoteRepositoryParams) RemoteRepoYAML {
	return RemoteRepoYAML{
		AllowAnyHostAuth:          *repo.AllowAnyHostAuth,
		AssumedOfflinePeriodSecs:  (*repo).AssumedOfflinePeriodSecs,
		BlackedOut:                *repo.BlackedOut,
		BlockMismatchingMimeTypes: *repo.BlockMismatchingMimeTypes,
		BypassHeadRequests:        *repo.BypassHeadRequests,
		ClientTlsCertificate:      (*repo).ClientTlsCertificate,
		Description:               (*repo).Description,
		DownloadRedirect:          Enablement{*repo.DownloadRedirect},
		EnableCookieManagement:    *repo.EnableCookieManagement,
		// EnableTokenAuthentication         : (*repo).
		ExcludesPattern: (*repo).ExcludesPattern,
		IncludesPattern: (*repo).IncludesPattern,
		// ListRemoteFolderItems             : *repo.
		MismatchingMimeTypesOverrideList: (*repo).MismatchingMimeTypesOverrideList,
		MissedRetrievalCachePeriodSecs:   (*repo).MissedRetrievalCachePeriodSecs,
		Notes:                            (*repo).Notes,
		Offline:                          *repo.Offline,
		Proxy:                            (*repo).Proxy,
		Password:                         (*repo).Password,
		PropertySets:                     (*repo).PropertySets,
		// QueryParams:                       (*repo).queryParams,
		RepoLayout:                        (*repo).RepoLayoutRef,
		SocketTimeoutMillis:               (*repo).SocketTimeoutMillis,
		SynchronizeProperties:             *repo.SynchronizeProperties,
		Url:                               (*repo).Url,
		UnusedArtifactsCleanupPeriodHours: (*repo).UnusedArtifactsCleanupPeriodHours,
		Username:                          (*repo).Username,
		Xray:                              Enablement{*repo.XrayIndex},

		// Nuget specific
		// ForceNugetAuthentication: *repo.,
		// Nuget: NugetSpecificities{(*repo).},
	}
}
