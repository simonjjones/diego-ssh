package authenticators

import (
	"encoding/json"
	"fmt"

	"code.cloudfoundry.org/bbs"
	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/diego-ssh/proxy"
	"code.cloudfoundry.org/diego-ssh/routes"
	"code.cloudfoundry.org/lager"
	"golang.org/x/crypto/ssh"
)

type permissionsBuilder struct {
	bbsClient             bbs.InternalClient
	useDirectInstanceAddr bool
}

func NewPermissionsBuilder(bbsClient bbs.InternalClient, useDirectInstanceAddr bool) PermissionsBuilder {
	return &permissionsBuilder{
		bbsClient:             bbsClient,
		useDirectInstanceAddr: useDirectInstanceAddr,
	}
}

func (pb *permissionsBuilder) Build(logger lager.Logger, processGuid string, index int, metadata ssh.ConnMetadata) (*ssh.Permissions, error) {
	actual, err := pb.bbsClient.ActualLRPGroupByProcessGuidAndIndex(logger, processGuid, index)
	if err != nil {
		return nil, err
	}

	desired, err := pb.bbsClient.DesiredLRPByProcessGuid(logger, processGuid)
	if err != nil {
		return nil, err
	}

	sshRoute, err := getRoutingInfo(desired)
	if err != nil {
		return nil, err
	}

	logMessage := fmt.Sprintf("Successful remote access by %s", metadata.RemoteAddr().String())

	actualLRP, _, resolveErr := actual.Resolve()
	if resolveErr != nil {
		return nil, resolveErr
	}
	return pb.createPermissions(sshRoute, actualLRP, desired.LogGuid, logMessage, index)
}

func (pb *permissionsBuilder) createPermissions(
	sshRoute *routes.SSHRoute,
	actual *models.ActualLRP,
	logGuid string,
	logMessage string,
	index int,
) (*ssh.Permissions, error) {
	var targetConfig *proxy.TargetConfig

	for _, mapping := range actual.Ports {
		if mapping.ContainerPort == sshRoute.ContainerPort {
			address := actual.Address
			port := mapping.HostPort
			if pb.useDirectInstanceAddr {
				address = actual.InstanceAddress
				port = mapping.ContainerPort
			}
			targetConfig = &proxy.TargetConfig{
				Address:         fmt.Sprintf("%s:%d", address, port),
				HostFingerprint: sshRoute.HostFingerprint,
				User:            sshRoute.User,
				Password:        sshRoute.Password,
				PrivateKey:      sshRoute.PrivateKey,
			}
			break
		}
	}

	if targetConfig == nil {
		return &ssh.Permissions{}, nil
	}

	targetConfigJson, err := json.Marshal(targetConfig)
	if err != nil {
		return nil, err
	}

	logMessageJson, err := json.Marshal(proxy.LogMessage{
		Guid:    logGuid,
		Message: logMessage,
		Index:   index,
	})
	if err != nil {
		return nil, err
	}

	return &ssh.Permissions{
		CriticalOptions: map[string]string{
			"proxy-target-config": string(targetConfigJson),
			"log-message":         string(logMessageJson),
		},
	}, nil
}

func getRoutingInfo(desired *models.DesiredLRP) (*routes.SSHRoute, error) {
	if desired.Routes == nil {
		return nil, RouteNotFoundErr
	}

	rawMessage := (*desired.Routes)[routes.DIEGO_SSH]
	if rawMessage == nil {
		return nil, RouteNotFoundErr
	}

	var sshRoute routes.SSHRoute
	err := json.Unmarshal(*rawMessage, &sshRoute)
	if err != nil {
		return nil, err
	}

	return &sshRoute, nil
}
