package provider

// Model for datasource_accountsnapshot
type queryASResponseModel struct {
	Data dataASResponseModel `json:"data"`
}

type dataASResponseModel struct {
	AccountSnapshot accountSnapshotASResponseModel `json:"accountSnapshot,omitempty"`
}

type accountSnapshotASResponseModel struct {
	Sites []sitesASResponseModel `json:"sites,omitempty"`
}

type sitesASResponseModel struct {
	Id   string              `json:"id"`
	Info infoASResponseModel `json:"info,omitempty"`
}

type infoASResponseModel struct {
	Name    string                   `json:"name"`
	Type    string                   `json:"type,omitempty"`
	Sockets []socketsASResponseModel `json:"sockets,omitempty"`
}

type socketsASResponseModel struct {
	Serial string `json:"serial"`
}

// Model for datasource_entitylookup
type queryELResponseModel struct {
	Data dataELResponseModel `json:"data"`
}

type dataELResponseModel struct {
	EntityLookup entityLookupELResponseModel `json:"entityLookup,omitempty"`
}

type entityLookupELResponseModel struct {
	Items []itemsELResponseModel `json:"items,omitempty"`
}

type itemsELResponseModel struct {
	Entity       entityELResponseModel `json:"entity"`
	Description  string                `json:"description,omitempty"`
	HelperFields map[string]string     `json:"helperFields,omitempty"`
}

type entityELResponseModel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Model for resource_socketsite
type mutationSSResponseModel struct {
	Data dataSSResponseModel `json:"data"`
}

type dataSSResponseModel struct {
	Site SiteSSResponseModel `json:"site,omitempty"`
}

type SiteSSResponseModel struct {
	AddSocketSite addSocketsiteSSResponseModel `json:"addSocketSite,omitempty"`
}

type addSocketsiteSSResponseModel struct {
	SiteID string `json:"siteId,omitempty"`
}

type socketSite struct {
	Id                 string       `json:"id,omitempty"`
	Name               string       `json:"name,omitempty"`
	Description        string       `json:"description,omitempty"`
	SiteType           string       `json:"siteType,omitempty"`
	NativeNetworkRange string       `json:"nativeNetworkRange,omitempty"`
	ConnectionType     string       `json:"connectionType,omitempty"`
	SiteLocation       siteLocation `json:"siteLocation,omitempty"`
}

type siteLocation struct {
	CountryCode string `json:"countryCode"`
	Timezone    string `json:"timezone"`
	StateCode   string `json:"stateCode,omitempty"`
}

// // Temporary models for Cato Client, to be externalised
// type queryResponseModel struct {
// 	Data dataResponseModel `json:"data"`
// }

// type dataResponseModel struct {
// 	AccountSnapshot interface{} `json:"accountSnapshot,omitempty"`
// 	// AccountSnapshot accountSnapshot `json:"accountSnapshot,omitempty"`
// 	// EntityLookup entityLookup `json:"entityLookup,omitempty"`
// 	Site         siteMutation `json:"accountSnapshot,omitempty"`
// }

// type entityLookup struct {
// 	Items []items `json:"items,omitempty"`
// }

// // type accountSnapshot struct {
// // 	Sites []sites `json:"sites,omitempty"`
// // }

// type siteMutation struct {
// 	AddSocketSite addSocketsite `json:"addSocketSite,omitempty"`
// }

// type addSocketsite struct {
// 	SiteID string `json:"siteId,omitempty"`
// }

// // type sites struct {
// // 	Id   string `json:"id"`
// // 	Info info   `json:"info,omitempty"`
// // }

// // type info struct {
// // 	Name    string    `json:"name"`
// // 	Type    string    `json:"type,omitempty"`
// // 	Sockets []sockets `json:"sockets,omitempty"`
// // }

// type sockets struct {
// 	Serial string `json:"serial"`
// }

// type items struct {
// 	Entity       entity            `json:"entity"`
// 	Description  string            `json:"description,omitempty"`
// 	HelperFields map[string]string `json:"helperFields,omitempty"`
// }

// type entity struct {
// 	Id   string `json:"id"`
// 	Name string `json:"name"`
// 	Type string `json:"type"`
// }
