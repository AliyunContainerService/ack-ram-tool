package addon

import "github.com/AliyunContainerService/ack-ram-tool/pkg/types"

type Meta interface {
	AddonName() string
	RoleName(clusterId string) string
	RamPolicy() types.RamPolicy
	NameSpace() string
	ServiceAccountName() string
}

var addons = map[string]Meta{}

func registryAddon(addon Meta) {
	addons[addon.AddonName()] = addon
}

func ListAddonNames() []string {
	var names []string
	for _, m := range addons {
		names = append(names, m.AddonName())
	}
	return names
}

func GetAddon(name string) Meta {
	return addons[name]
}
