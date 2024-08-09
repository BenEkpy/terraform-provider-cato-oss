package catogo

import (
	"encoding/json"
	"fmt"
)

type AccountSnapshot struct {
	Id        *string        `json:"id,omitempty"`
	Sites     []SiteSnapshot `json:"sites,omitempty"`
	Users     []UserSnapshot `json:"users,omitempty"`
	Timestamp *string        `json:"timestamp,omitempty"`
}

type IPInfo struct {
	Ip          *string  `json:"ip,omitempty"`
	CountryCode *string  `json:"countryCode,omitempty"`
	CountryName *string  `json:"countryName,omitempty"`
	City        *string  `json:"city,omitempty"`
	State       *string  `json:"state,omitempty"`
	Provider    *string  `json:"provider,omitempty"`
	Latitude    *float64 `json:"latitude,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
}

type SocketInfo struct {
	Id                *string `json:"id,omitempty"`
	Serial            *string `json:"serial,omitempty"`
	IsPrimary         *bool   `json:"isPrimary,omitempty"`
	Platform          *string `json:"platform,omitempty"`
	Version           *string `json:"version,omitempty"`
	VersionUpdateTime *string `json:"versionUpdateTime,omitempty"`
}

type IPSecInfo struct {
	IsPrimary  *bool   `json:"isPrimary,omitempty"`
	CatoIP     *string `json:"catoIP,omitempty"`
	RemoteIP   *string `json:"remoteIP,omitempty"`
	IkeVersion *int64  `json:"ikeVersion,omitempty"`
}

type InterfaceInfo struct {
	Id                  string  `json:"id,omitempty"`
	Name                *string `json:"name,omitempty"`
	UpstreamBandwidth   *int64  `json:"upstreamBandwidth,omitempty"`
	DownstreamBandwidth *int64  `json:"downstreamBandwidth,omitempty"`
	DestType            *string `json:"destType,omitempty"`
}

type SiteInfo struct {
	Name         *string         `json:"name,omitempty"`
	Type         *string         `json:"type,omitempty"`
	Description  *string         `json:"description,omitempty"`
	CountryCode  *string         `json:"countryCode,omitempty"`
	Region       *string         `json:"region,omitempty"`
	CountryName  *string         `json:"countryName,omitempty"`
	IsHA         *bool           `json:"isHA,omitempty"`
	ConnType     *string         `json:"connType,omitempty"`
	CreationTime *string         `json:"creationTime,omitempty"`
	Interfaces   []InterfaceInfo `json:"interfaces,omitempty"`
	Sockets      []SocketInfo    `json:"sockets,omitempty"`
	Ipsec        []IPSecInfo     `json:"ipsec,omitempty"`
}

type HaStatus struct {
	Readiness       *string `json:"readiness,omitempty"`
	WanConnectivity *string `json:"wanConnectivity,omitempty"`
	Keepalive       *string `json:"keepalive,omitempty"`
	SocketVersion   *string `json:"socketVersion,omitempty"`
}

type UserInfo struct {
	Name         *string `json:"name,omitempty"`
	Status       *string `json:"status,omitempty"`
	Email        *string `json:"email,omitempty"`
	CreationTime *string `json:"creationTime,omitempty"`
	PhoneNumber  *string `json:"phoneNumber,omitempty"`
	Origin       *string `json:"origin,omitempty"`
	AuthMethod   *string `json:"authMethod,omitempty"`
}

type RecentConnection struct {
	Duration      *int64  `json:"duration,omitempty"`
	InterfaceName *string `json:"interfaceName,omitempty"`
	DeviceName    *string `json:"deviceName,omitempty"`
	LastConnected *string `json:"lastConnected,omitempty"`
	PopName       *string `json:"popName,omitempty"`
	RemoteIP      *string `json:"remoteIP,omitempty"`
	RemoteIPInfo  *IPInfo `json:"remoteIPInfo,omitempty"`
}

type InterfaceLinkState struct {
	Id        *string `json:"id,omitempty"`
	Up        *bool   `json:"up,omitempty"`
	MediaIn   *bool   `json:"mediaIn,omitempty"`
	LinkSpeed *string `json:"linkSpeed,omitempty"`
	Duplex    *string `json:"duplex,omitempty"`
}

type DeviceSnapshot struct {
	Id                  *string              `json:"id,omitempty"`
	Name                *string              `json:"name,omitempty"`
	Identifier          *string              `json:"identifier,omitempty"`
	Connected           *bool                `json:"connected,omitempty"`
	HaRole              *string              `json:"haRole,omitempty"`
	Interfaces          []InterfaceSnapshot  `json:"interfaces,omitempty"`
	LastConnected       *string              `json:"lastConnected,omitempty"`
	LastDuration        *int64               `json:"lastDuration,omitempty"`
	ConnectedSince      *string              `json:"connectedSince,omitempty"`
	LastPopID           *int64               `json:"lastPopID,omitempty"`
	LastPopName         *string              `json:"lastPopName,omitempty"`
	RecentConnections   []RecentConnection   `json:"recentConnections,omitempty"`
	Type                *string              `json:"type,omitempty"`
	SocketInfo          *SocketInfo          `json:"socketInfo,omitempty"`
	InterfacesLinkState []InterfaceLinkState `json:"interfacesLinkState,omitempty"`
	OsType              *string              `json:"osType,omitempty"`
	OsVersion           *string              `json:"osVersion,omitempty"`
	Version             *string              `json:"version,omitempty"`
	VersionNumber       *int64               `json:"versionNumber,omitempty"`
	ReleaseGroup        *string              `json:"releaseGroup,omitempty"`
	MfaExpirationTime   *int64               `json:"mfaExpirationTime,omitempty"`
	MfaCreationTime     *int64               `json:"mfaCreationTime,omitempty"`
	InternalIP          *string              `json:"internalIP,omitempty"`
}

type InterfaceSnapshot struct {
	Connected              *bool              `json:"connected,omitempty"`
	Id                     *string            `json:"id,omitempty"`
	Name                   *string            `json:"name,omitempty"`
	PhysicalPort           *int64             `json:"physicalPort,omitempty"`
	NaturalOrder           *int64             `json:"naturalOrder,omitempty"`
	PopName                *string            `json:"popName,omitempty"`
	PreviousPopID          *int64             `json:"previousPopID,omitempty"`
	PreviousPopName        *string            `json:"previousPopName,omitempty"`
	TunnelConnectionReason *string            `json:"tunnelConnectionReason,omitempty"`
	TunnelUptime           *int64             `json:"tunnelUptime,omitempty"`
	TunnelRemoteIP         *string            `json:"tunnelRemoteIP,omitempty"`
	TunnelRemoteIPInfo     *IPInfo            `json:"tunnelRemoteIPInfo,omitempty"`
	Type                   *string            `json:"type,omitempty"`
	Info                   *InterfaceInfo     `json:"info,omitempty"`
	CellularInterfaceInfo  *CellularInterface `json:"cellularInterfaceInfo,omitempty"`
}

type CellularInterface struct {
	NetworkType         *string `json:"networkType,omitempty"`
	SimSlotId           *int64  `json:"simSlotId,omitempty"`
	ModemStatus         *string `json:"modemStatus,omitempty"`
	IsModemConnected    bool    `json:"isModemConnected,omitempty"`
	Iccid               *string `json:"iccid,omitempty"`
	Imei                *string `json:"imei,omitempty"`
	OperatorName        *string `json:"operatorName,omitempty"`
	IsModemSuspended    bool    `json:"isModemSuspended,omitempty"`
	Apn                 *string `json:"apn,omitempty"`
	ApnSelectionMethod  *string `json:"apnSelectionMethod,omitempty"`
	SignalStrength      *string `json:"signalStrength,omitempty"`
	IsRoamingAllowed    bool    `json:"isRoamingAllowed,omitempty"`
	SimNumber           *string `json:"simNumber,omitempty"`
	DisconnectionReason *string `json:"disconnectionReason,omitempty"`
	IsSimSlot1Detected  bool    `json:"isSimSlot1Detected,omitempty"`
	IsSimSlot2Detected  bool    `json:"isSimSlot2Detected,omitempty"`
}

type UserSnapshot struct {
	Id                 *string            `json:"id,omitempty"`
	ConnectivityStatus *string            `json:"connectivityStatus,omitempty"`
	OperationalStatus  *string            `json:"operationalStatus,omitempty"`
	Name               *string            `json:"name,omitempty"`
	DeviceName         *string            `json:"deviceName,omitempty"`
	Uptime             *int64             `json:"uptime,omitempty"`
	LastConnected      *string            `json:"lastConnected,omitempty"`
	Version            *string            `json:"version,omitempty"`
	VersionNumber      *int64             `json:"versionNumber,omitempty"`
	PopID              *int64             `json:"popID,omitempty"`
	PopName            *string            `json:"popName,omitempty"`
	RemoteIP           *string            `json:"remoteIP,omitempty"`
	RemoteIPInfo       *IPInfo            `json:"remoteIPInfo,omitempty"`
	InternalIP         *string            `json:"internalIP,omitempty"`
	OsType             *string            `json:"osType,omitempty"`
	OsVersion          *string            `json:"osVersion,omitempty"`
	Devices            []DeviceSnapshot   `json:"devices,omitempty"`
	ConnectedInOffice  *bool              `json:"connectedInOffice,omitempty"`
	Info               *UserInfo          `json:"info,omitempty"`
	RecentConnections  []RecentConnection `json:"recentConnections,omitempty"`
}

type SiteSnapshot struct {
	Id                 *string          `json:"id,omitempty"`
	ProtoId            *int             `json:"protoId,omitempty"`
	ConnectivityStatus *string          `json:"connectivityStatus,omitempty"`
	HaStatus           *HaStatus        `json:"haStatus,omitempty"`
	OperationalStatus  *string          `json:"operationalStatus,omitempty"`
	LastConnected      *string          `json:"lastConnected,omitempty"`
	ConnectedSince     *string          `json:"connectedSince,omitempty"`
	PopName            *string          `json:"popName,omitempty"`
	Devices            []DeviceSnapshot `json:"devices,omitempty"`
	Info               *SiteInfo        `json:"info,omitempty"`
	HostCount          *int64           `json:"hostCount,omitempty"`
	AltWanStatus       *string          `json:"altWanStatus,omitempty"`
}

func (c *Client) AccountSnapshotSite() (*[]SiteSnapshot, error) {

	query := graphQLRequest{
		Query: `
		query accountSnapshot($accountId: ID!) {
			accountSnapshot(accountID: $accountId) {
				id
				sites {
				altWanStatus
				connectedSince
				connectivityStatus
				devices {
					connected
					connectedSince
					haRole
					id
					identifier
					interfaces {
					cellularInterfaceInfo {
						apn
						apnSelectionMethod
						disconnectionReason
						iccid
						imei
						isModemConnected
						isModemSuspended
						isRoamingAllowed
						isSimSlot1Detected
						isSimSlot2Detected
						modemStatus
						networkType
						operatorName
						signalStrength
						simNumber
						simSlotId
					}
					connected
					id
					info {
						destType
						downstreamBandwidth
						id
						name
						upstreamBandwidth
					}
					name
					naturalOrder
					physicalPort
					popName
					previousPopID
					previousPopName
					tunnelConnectionReason
					tunnelRemoteIP
					tunnelRemoteIPInfo {
						city
						countryCode
						countryName
						ip
						latitude
						longitude
						provider
						state
					}
					tunnelUptime
					type
					}
					interfacesLinkState {
					duplex
					id
					mediaIn
					linkSpeed
					up
					}
					internalIP
					lastConnected
					lastDuration
					lastPopID
					lastPopName
					mfaCreationTime
					mfaExpirationTime
					name
					osType
					recentConnections {
					deviceName
					duration
					interfaceName
					lastConnected
					popName
					remoteIP
					remoteIPInfo {
						city
						countryCode
						countryName
						ip
						latitude
						longitude
						provider
						state
					}
					}
					osVersion
					releaseGroup
					socketInfo {
					id
					isPrimary
					platform
					serial
					version
					versionUpdateTime
					}
					type
					version
					versionNumber
				}
				haStatus {
					keepalive
					readiness
					socketVersion
					wanConnectivity
				}
				hostCount
				id
				info {
					connType
					countryCode
					countryName
					creationTime
					description
					interfaces {
					destType
					downstreamBandwidth
					id
					name
					upstreamBandwidth
					}
					ipsec {
					catoIP
					ikeVersion
					isPrimary
					remoteIP
					}
					isHA
					name
					region
					sockets {
					id
					isPrimary
					platform
					serial
					version
					versionUpdateTime
					}
					type
				}
				lastConnected
				operationalStatus
				popName
				protoId
				}
				timestamp
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct{ AccountSnapshot AccountSnapshot }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.AccountSnapshot.Sites, nil
}

func (c *Client) AccountSnapshotSiteById(siteId string) (*SiteSnapshot, error) {

	query := graphQLRequest{
		Query: `
		query accountSnapshot($accountId: ID!, $siteId: [ID!]) {
			accountSnapshot(accountID: $accountId) {
				id
				sites(siteIDs: $siteId) {
				altWanStatus
				connectedSince
				connectivityStatus
				devices {
					connected
					connectedSince
					haRole
					id
					identifier
					interfaces {
					cellularInterfaceInfo {
						apn
						apnSelectionMethod
						disconnectionReason
						iccid
						imei
						isModemConnected
						isModemSuspended
						isRoamingAllowed
						isSimSlot1Detected
						isSimSlot2Detected
						modemStatus
						networkType
						operatorName
						signalStrength
						simNumber
						simSlotId
					}
					connected
					id
					info {
						destType
						downstreamBandwidth
						id
						name
						upstreamBandwidth
					}
					name
					naturalOrder
					physicalPort
					popName
					previousPopID
					previousPopName
					tunnelConnectionReason
					tunnelRemoteIP
					tunnelRemoteIPInfo {
						city
						countryCode
						countryName
						ip
						latitude
						longitude
						provider
						state
					}
					tunnelUptime
					type
					}
					interfacesLinkState {
					duplex
					id
					mediaIn
					linkSpeed
					up
					}
					internalIP
					lastConnected
					lastDuration
					lastPopID
					lastPopName
					mfaCreationTime
					mfaExpirationTime
					name
					osType
					recentConnections {
					deviceName
					duration
					interfaceName
					lastConnected
					popName
					remoteIP
					remoteIPInfo {
						city
						countryCode
						countryName
						ip
						latitude
						longitude
						provider
						state
					}
					}
					osVersion
					releaseGroup
					socketInfo {
					id
					isPrimary
					platform
					serial
					version
					versionUpdateTime
					}
					type
					version
					versionNumber
				}
				haStatus {
					keepalive
					readiness
					socketVersion
					wanConnectivity
				}
				hostCount
				id
				info {
					connType
					countryCode
					countryName
					creationTime
					description
					interfaces {
					destType
					downstreamBandwidth
					id
					name
					upstreamBandwidth
					}
					ipsec {
					catoIP
					ikeVersion
					isPrimary
					remoteIP
					}
					isHA
					name
					region
					sockets {
					id
					isPrimary
					platform
					serial
					version
					versionUpdateTime
					}
					type
				}
				lastConnected
				operationalStatus
				popName
				protoId
				}
				timestamp
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"siteId":    []string{siteId},
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct{ AccountSnapshot AccountSnapshot }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.AccountSnapshot.Sites == nil {
		return nil, fmt.Errorf("Site " + siteId + " doesn't exists")
	} else {
		return &response.AccountSnapshot.Sites[0], nil
	}
}

func (c *Client) GetSocketWanInterfacelist(siteId string) ([]InterfaceInfo, error) {

	query := graphQLRequest{
		Query: `
		query accountSnapshot($accountId: ID!, $siteId: [ID!]) {
			accountSnapshot(accountID: $accountId) {
				id
				sites(siteIDs: $siteId) {
				id
				info {
					interfaces {
						destType
						downstreamBandwidth
						id
						name
						upstreamBandwidth
					}
				}
				}
				timestamp
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"siteId":    []string{siteId},
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct{ AccountSnapshot AccountSnapshot }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.AccountSnapshot.Sites == nil {

		return nil, fmt.Errorf("Site " + siteId + " doesn't exists")

	} else {

		var interfaceList []InterfaceInfo

		// retrieve only interface that have DestType as CATO
		for _, item := range response.AccountSnapshot.Sites[0].Info.Interfaces {

			if *item.DestType == "CATO" {
				interfaceList = append(interfaceList, item)
			}
		}

		return interfaceList, nil
	}
}
