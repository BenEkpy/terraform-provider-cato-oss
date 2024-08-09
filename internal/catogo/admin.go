package catogo

import (
	"encoding/json"
)

type AddAdminInput struct {
	FirstName            string                 `json:"firstName,omitempty"`
	LastName             string                 `json:"lastName,omitempty"`
	Email                string                 `json:"email,omitempty"`
	PasswordNeverExpires *bool                  `json:"passwordNeverExpires,omitempty"`
	MfaEnabled           *bool                  `json:"mfaEnabled,omitempty"`
	ManagedRoles         []UpdateAdminRoleInput `json:"managedRoles,omitempty"`
	ResellerRoles        []UpdateAdminRoleInput `json:"resellerRoles,omitempty"`
}

type UpdateAdminInput struct {
	FirstName            *string                `json:"firstName,omitempty"`
	LastName             *string                `json:"lastName,omitempty"`
	PasswordNeverExpires *bool                  `json:"passwordNeverExpires,omitempty"`
	MfaEnabled           *bool                  `json:"mfaEnabled,omitempty"`
	ManagedRoles         []UpdateAdminRoleInput `json:"managedRoles,omitempty"`
	ResellerRoles        []UpdateAdminRoleInput `json:"resellerRoles,omitempty"`
}

type RemoveAdminPayload struct {
	AdminID string `json:"adminID"`
}

type GetAdminPayload struct {
	Id                   string      `json:"id,omitempty"`
	FirstName            string      `json:"firstName,omitempty"`
	LastName             string      `json:"lastName,omitempty"`
	Email                string      `json:"email,omitempty"`
	CreationDate         string      `json:"creationDate,omitempty"`
	PasswordNeverExpires bool        `json:"passwordNeverExpires,omitempty"`
	MfaEnabled           bool        `json:"mfaEnabled,omitempty"`
	ManagedRoles         []AdminRole `json:"managedRoles,omitempty"`
	ResellerRoles        []AdminRole `json:"resellerRoles,omitempty"`
}

type EntityInput struct {
	Id   string  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type string  `json:"type,omitempty"`
}

type UpdateAccountRoleInput struct {
	Id   string  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type UpdateAdminRoleInput struct {
	Role            UpdateAccountRoleInput `json:"role,omitempty"`
	AllowedEntities []EntityInput          `json:"allowedEntities,omitempty"`
	AllowedAccounts []string               `json:"allowedAccounts,omitempty"`
}

type AddAdminPayload struct {
	AdminID string `json:"adminID,omitempty"`
}

type UpdateAdminPayload struct {
	AdminID string `json:"adminID,omitempty"`
}

type RBACRole struct {
	Id           string  `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	IsPredefined bool    `json:"isPredefined,omitempty"`
}

type AdminRole struct {
	Role            RBACRole `json:"role,omitempty"`
	AllowedEntities []Entity `json:"allowedEntities,omitempty"`
	AllowedAccounts []string `json:"allowedAccounts,omitempty"`
}

type AdminsResult struct {
	Items []Admin `json:"items,omitempty"`
	Total int64   `json:"total,omitempty"`
}

type Admin struct {
	Id                    string      `json:"id,omitempty"`
	Version               string      `json:"version,omitempty"`
	Role                  *string     `json:"role,omitempty"`
	FirstName             *string     `json:"firstName,omitempty"`
	LastName              *string     `json:"lastName,omitempty"`
	Email                 *string     `json:"email,omitempty"`
	CreationDate          *string     `json:"creationDate,omitempty"`
	ModifyDate            *string     `json:"modifyDate,omitempty"`
	Status                *string     `json:"status,omitempty"`
	PasswordNeverExpires  *bool       `json:"passwordNeverExpires,omitempty"`
	MfaEnabled            *bool       `json:"mfaEnabled,omitempty"`
	NativeAccountID       *string     `json:"nativeAccountID,omitempty"`
	AllowedItems          []Entity    `json:"allowedItems,omitempty"`
	PresentUsageAndEvents *bool       `json:"presentUsageAndEvents,omitempty"`
	ManagedRoles          []AdminRole `json:"managedRoles,omitempty"`
	ResellerRoles         []AdminRole `json:"resellerRoles,omitempty"`
}

type AccountRolesResult struct {
	Items []RBACRole `json:"items,omitempty"`
	Total int64      `json:"total,omitempty"`
}

func (c *Client) GetAdmins() (*AdminsResult, error) {

	query := graphQLRequest{
		Query: `
		query GetAdmins($accountId: ID!) {
			admins(accountID: $accountId) {
				items {
				version
				status
				role
				resellerRoles {
					allowedAccounts
					role {
					description
					id
					isPredefined
					name
					}
					allowedEntities {
					id
					name
					type
					}
				}
				presentUsageAndEvents
				passwordNeverExpires
				nativeAccountID
				modifyDate
				mfaEnabled
				managedRoles {
					allowedAccounts
					allowedEntities {
					id
					name
					type
					}
					role {
					description
					id
					isPredefined
					name
					}
				}
				lastName
				id
				firstName
				email
				creationDate
				allowedItems {
					id
					name
					type
				}
				}
				total
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

	var response struct{ Admins AdminsResult }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Admins, nil
}

func (c *Client) GetAdmin(adminId string) (*GetAdminPayload, error) {

	query := graphQLRequest{
		Query: `
		query GetAdmin ($accountId: ID!, $adminId: ID!) {
			admin(accountId: $accountId, adminID: $adminId) {
				creationDate
				email
				firstName
				id
				lastName
				managedRoles {
				allowedAccounts
				allowedEntities {
					id
					name
					type
				}
				role {
					description
					id
					isPredefined
					name
				}
				}
				mfaEnabled
				passwordNeverExpires
				resellerRoles {
				allowedAccounts
				role {
					description
					name
					id
					isPredefined
				}
				allowedEntities {
					id
					name
					type
				}
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"adminId":   adminId,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct{ Admin GetAdminPayload }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Admin, nil
}

func (c *Client) GetAccountRoles() (*AccountRolesResult, error) {

	query := graphQLRequest{
		Query: `
		query accountRoles($accountId: ID!) {
			accountRoles(accountID: $accountId) {
				items {
					name
					id
					description
				}
				total
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

	var response struct{ AccountRoles AccountRolesResult }

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.AccountRoles, nil
}

func (c *Client) GetAccountRoleByName(name string) (*RBACRole, error) {

	accountRoles, _ := c.GetAccountRoles()

	accountRole := RBACRole{}

	for _, item := range accountRoles.Items {

		if item.Name == name {
			accountRole = item
		}

	}

	return &accountRole, nil
}

func (c *Client) AddAdmin(input AddAdminInput) (*AddAdminPayload, error) {

	query := graphQLRequest{
		Query: `
		mutation addAdmin($accountId:ID!, $input: AddAdminInput!) {
			admin(accountId:$accountId) {
				addAdmin(input:$input) {
					adminID
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"input":     input,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Admin struct{ AddAdmin AddAdminPayload }
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Admin.AddAdmin, nil
}

func (c *Client) UpdateAdmin(adminId string, input UpdateAdminInput) (*UpdateAdminPayload, error) {

	query := graphQLRequest{
		Query: `
		mutation updateAdmin($accountId:ID!, $adminId:ID!, $input: UpdateAdminInput!){
			admin(accountId:$accountId) {
				updateAdmin(adminID:$adminId,input:$input) {
					adminID
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"adminId":   adminId,
			"input":     input,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Admin struct{ UpdateAdmin UpdateAdminPayload }
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Admin.UpdateAdmin, nil
}

func (c *Client) RemoveAdmin(adminId string) (*RemoveAdminPayload, error) {

	query := graphQLRequest{
		Query: `
		mutation removeAdmin($accountId:ID!, $adminId:ID!){
			admin(accountId:$accountId) {
				removeAdmin(adminID:$adminId) {
					adminID
				}
			}
		}`,
		Variables: map[string]interface{}{
			"accountId": c.accountId,
			"adminId":   adminId,
		},
	}

	body, err := c.do(query)
	if err != nil {
		return nil, err
	}

	var response struct {
		Admin struct{ RemoveAdmin RemoveAdminPayload }
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Admin.RemoveAdmin, nil
}
