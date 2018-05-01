package models

import (
	"net/http"
	"strings"

	"github.com/serboox/os-cli/configs"
)

// ResServersDetail structure declared response JSON
type ResServersDetail struct {
	Servers []ResServersDetailItem `json:"servers"`
}

// ResServersDetailItem structure declared response JSON part
type ResServersDetailItem struct {
	ID                            string                              `json:"id"`
	OSExtStsTaskState             interface{}                         `json:"OS-EXT-STS:task_state"`
	Addresses                     ResServersDetailAddresses           `json:"addresses"`
	Links                         []ResServersDetailLinks             `json:"links"`
	Image                         interface{}                         `json:"image"`
	OSExtStsVMState               string                              `json:"OS-EXT-STS:vm_state"`
	OSExtSrvAttr                  string                              `json:"OS-EXT-SRV-ATTR:instance_name"`
	OSExtUsg                      string                              `json:"OS-SRV-USG:launched_at"`
	Flavor                        ResServersDetailFlavor              `json:"flavor"`
	SecurityGroups                []ResServersDetailSecurityGroup     `json:"security_groups"`
	UserID                        string                              `json:"user_id"`
	OSDcf                         string                              `json:"OS-DCF:diskConfig"`
	AccessIPv4                    string                              `json:"accessIPv4"`
	AccessIPv6                    string                              `json:"accessIPv6"`
	Progress                      int64                               `json:"progress"`
	OSExtStsPowerStare            int64                               `json:"OS-EXT-STS:power_state"`
	OSExtAz                       string                              `json:"OS-EXT-AZ:availability_zone"`
	MetaData                      interface{}                         `json:"metadata"`
	Status                        string                              `json:"status"`
	Updated                       string                              `json:"updated"`
	HostID                        string                              `json:"hostId"`
	OSExtSrvAttrHost              string                              `json:"OS-EXT-SRV-ATTR:host"`
	OSSrvUsu                      interface{}                         `json:"OS-SRV-USG:terminated_at"`
	KeyName                       interface{}                         `json:"key_name"`
	OSExtSrvAttrHypevisorHostname string                              `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
	Name                          string                              `json:"name"`
	Created                       string                              `json:"created"`
	TenandID                      string                              `json:"tenant_id"`
	OSExtendedVolumes             []ResServersDetailOSExtendedVolumes `json:"os-extended-volumes:volumes_attached"`
	CongigDrive                   string                              `json:"config_drive"`
}

// ResServersDetailAddresses structure declared response JSON part
type ResServersDetailAddresses struct {
	Public []ResServersDetailAddressesPublic `json:"public"`
}

// ResServersDetailAddressesPublic structure declared response JSON part
type ResServersDetailAddressesPublic struct {
	MacAddr  string `json:"OS-EXT-IPS-MAC:mac_addr"`
	Varsion  int64  `json:"version"`
	Addr     string `json:"addr"`
	OSExtIps string `json:"OS-EXT-IPS:type"`
}

// ResServersDetailLinks structure declared response JSON part
type ResServersDetailLinks struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

// ResServersDetailImage structure declared response JSON part
type ResServersDetailImage struct {
	ID    string                       `json:"id"`
	Links []ResServersDetailImageLinks `json:"links"`
}

// ResServersDetailImageLinks structure declared response JSON part
type ResServersDetailImageLinks struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

// ResServersDetailFlavor structure declared response JSON part
type ResServersDetailFlavor struct {
	ID    string                        `json:"id"`
	Links []ResServersDetailFlavorLinks `json:"links"`
}

// ResServersDetailFlavorLinks structure declared response JSON part
type ResServersDetailFlavorLinks struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

// ResServersDetailSecurityGroup structure declared response JSON part
type ResServersDetailSecurityGroup struct {
	Name string `json:"name"`
}

// ResServersDetailOSExtendedVolumes structure declared response JSON part
type ResServersDetailOSExtendedVolumes struct {
	ID string `json:"id"`
}

//Get List Servers
func (res *ResServersDetail) Get(token, url string) (
	resp *http.Response, err error,
) {
	ctx := sendDataCtx{
		methodName: "ResServersDetail.Get",
		urlMethod:  "GET",
		url:        url,
		headers: map[string]string{
			configs.XAuthTokenKey: token,
		},
		res:       res,
		newReader: new(strings.Reader),
	}

	return ctx.Send()
}
