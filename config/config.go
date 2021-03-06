// Copyright 2016 NetApp, Inc. All Rights Reserved.

package config

import (
	"fmt"
	"time"

	dvp "github.com/netapp/netappdvp/storage_drivers"
	"github.com/netapp/trident/utils"
)

type Protocol string
type AccessMode string
type VolumeType string

const (
	/* Misc. orchestrator constants */
	OrchestratorName                 = "trident"
	orchestratorVersion              = "18.01.0"
	OrchestratorAPIVersion           = "1"
	PersistentStoreBootstrapAttempts = 30
	PersistentStoreBootstrapTimeout  = PersistentStoreBootstrapAttempts * time.Second
	PersistentStoreTimeout           = 10 * time.Second

	/* Protocol constants */
	File                Protocol = "file"
	Block               Protocol = "block"
	ProtocolAny         Protocol = ""
	ProtocolUnsupported Protocol = "unsupported"

	/* Access mode constants */
	ReadWriteOnce AccessMode = "ReadWriteOnce"
	ReadOnlyMany  AccessMode = "ReadOnlyMany"
	ReadWriteMany AccessMode = "ReadWriteMany"
	ModeAny       AccessMode = ""

	/* Volume type constants */
	ONTAP_NFS         VolumeType = "ONTAP_NFS"
	ONTAP_iSCSI       VolumeType = "ONTAP_iSCSI"
	SolidFire_iSCSI   VolumeType = "SolidFire_iSCSI"
	Eseries_iSCSI     VolumeType = "Eseries_iSCSI"
	UnknownVolumeType VolumeType = ""

	/* Driver-related constants */
	DefaultOntapIgroup      = OrchestratorName
	DefaultSolidFireVAG     = OrchestratorName
	DefaultEseriesHostGroup = OrchestratorName
	UnknownDriver           = "UnknownDriver"

	/* REST frontend constants */
	MaxRESTRequestSize = 10240

	/* Kubernetes deployment constants */
	ContainerTrident = "trident-main"
	ContainerEtcd    = "etcd"
)

var (
	validProtocols = map[Protocol]bool{
		File:        true,
		Block:       true,
		ProtocolAny: true,
	}

	// BuildHash is the git hash the binary was built from
	BuildHash = "unknown"

	// BuildType is the type of build: custom, beta or stable
	BuildType = "custom"

	// BuildTypeRev is the revision of the build
	BuildTypeRev = "0"

	// BuildTime is the time the binary was built
	BuildTime = "unknown"

	OrchestratorVersion = utils.MustParseDate(version())

	/* API Server and persistent store variables */
	BaseURL         = "/" + OrchestratorName + "/v" + OrchestratorAPIVersion
	VersionURL      = "/" + OrchestratorName + "/v" + OrchestratorAPIVersion + "/version"
	BackendURL      = "/" + OrchestratorName + "/v" + OrchestratorAPIVersion + "/backend"
	VolumeURL       = "/" + OrchestratorName + "/v" + OrchestratorAPIVersion + "/volume"
	TransactionURL  = "/" + OrchestratorName + "/v" + OrchestratorAPIVersion + "/txn"
	StorageClassURL = "/" + OrchestratorName + "/v" + OrchestratorAPIVersion + "/storageclass"
	StoreURL        = "/" + OrchestratorName + "/store"

	UsingPassthroughStore bool
	DriverContext         dvp.DriverContext
)

func IsValidProtocol(p Protocol) bool {
	_, ok := validProtocols[p]
	return ok
}

func GetValidProtocolNames() []string {
	ret := make([]string, len(validProtocols))
	for key, _ := range validProtocols {
		ret = append(ret, string(key))
	}
	return ret
}

func version() string {

	var version string

	if BuildType != "stable" {
		if BuildType == "custom" {
			version = fmt.Sprintf("%v-%v+%v", orchestratorVersion, BuildType, BuildHash)
		} else {
			version = fmt.Sprintf("%v-%v.%v+%v", orchestratorVersion, BuildType, BuildTypeRev, BuildHash)
		}
	} else {
		version = orchestratorVersion
	}

	return version
}
