// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PermissionResponseRead - Reformat PermissionResponse with permission scope
type PermissionResponseRead struct {
	PermissionID string `json:"permissionId"`
	// Describes what actions/endpoints the permission entitles to
	PermissionType PermissionType `json:"permissionType"`
	// Scope of a single permission, e.g. workspace, organization
	Scope   PermissionScope `json:"scope"`
	ScopeID string          `json:"scopeId"`
	// Internal Airbyte user ID
	UserID string `json:"userId"`
}

func (o *PermissionResponseRead) GetPermissionID() string {
	if o == nil {
		return ""
	}
	return o.PermissionID
}

func (o *PermissionResponseRead) GetPermissionType() PermissionType {
	if o == nil {
		return PermissionType("")
	}
	return o.PermissionType
}

func (o *PermissionResponseRead) GetScope() PermissionScope {
	if o == nil {
		return PermissionScope("")
	}
	return o.Scope
}

func (o *PermissionResponseRead) GetScopeID() string {
	if o == nil {
		return ""
	}
	return o.ScopeID
}

func (o *PermissionResponseRead) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}
