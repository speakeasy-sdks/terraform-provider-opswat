// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UISettings - Configuration of Management Console for this user.
type UISettings struct {
}

// UserRequest - User object in MetaDefender.
type UserRequest struct {
	// Associated apikey with this user
	APIKey *string `json:"api_key,omitempty"`
	// To which User Directories belongs to (LOCAL/System/etc.)
	DirectoryID *int64 `json:"directory_id,omitempty"`
	// Which is the name showed in the Management Console
	DisplayName *string `json:"display_name,omitempty"`
	// User's email address
	Email *string `json:"email,omitempty"`
	// User's full name
	Name *string `json:"name,omitempty"`
	// The user's password
	Password *string `json:"password,omitempty"`
	// A list of roles attached to this user
	Roles []string `json:"roles,omitempty"`
	// Configuration of Management Console for this user.
	UISettings *UISettings `json:"ui_settings,omitempty"`
}

func (o *UserRequest) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *UserRequest) GetDirectoryID() *int64 {
	if o == nil {
		return nil
	}
	return o.DirectoryID
}

func (o *UserRequest) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *UserRequest) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *UserRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UserRequest) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *UserRequest) GetRoles() []string {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *UserRequest) GetUISettings() *UISettings {
	if o == nil {
		return nil
	}
	return o.UISettings
}
