/*
Infoblox GRID API

OpenAPI specification for Infoblox NIOS WAPI GRID objects

API version: 2.13.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package grid

import (
	"encoding/json"
)

// checks if the GridDnsDtcScheduledBackup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GridDnsDtcScheduledBackup{}

// GridDnsDtcScheduledBackup struct for GridDnsDtcScheduledBackup
type GridDnsDtcScheduledBackup struct {
	// The status of the scheduled backup.
	Status *string `json:"status,omitempty"`
	// The state for scheduled backup or restore operation.
	Execute *string `json:"execute,omitempty"`
	// The scheduled backup operation.
	Operation *string `json:"operation,omitempty"`
	// The destination of the backup files.
	BackupType *string `json:"backup_type,omitempty"`
	// Determines whether the local backup performed before uploading backup to remote storage.
	KeepLocalCopy *bool `json:"keep_local_copy,omitempty"`
	// The frequency of backups.
	BackupFrequency *string `json:"backup_frequency,omitempty"`
	// The day of the week when the backup is performed.
	Weekday *string `json:"weekday,omitempty"`
	// The hour of the day past 12:00 AM the backup is performed.
	HourOfDay *int64 `json:"hour_of_day,omitempty"`
	// The minute of the hour when the backup is performed.
	MinutesPastHour *int64 `json:"minutes_past_hour,omitempty"`
	// The user name on the backup server.
	Username *string `json:"username,omitempty"`
	// The user password on the backup server.
	Password *string `json:"password,omitempty"`
	// The IP address of the backup server.
	BackupServer *string `json:"backup_server,omitempty"`
	// The directory path to the backup file stored on the server.
	Path *string `json:"path,omitempty"`
	// The destination of the restore files.
	RestoreType *string `json:"restore_type,omitempty"`
	// The IP address of the restore server.
	RestoreServer *string `json:"restore_server,omitempty"`
	// The user name on the restore server.
	RestoreUsername *string `json:"restore_username,omitempty"`
	// The password on the restore server.
	RestorePassword *string `json:"restore_password,omitempty"`
	// The directory path to the restored file on the server.
	RestorePath *string `json:"restore_path,omitempty"`
	// Determines whether the restore of the NIOS data is enabled.
	NiosData *bool `json:"nios_data,omitempty"`
	// Determines whether the restore the NetMRI data is enabled.
	DiscoveryData *bool `json:"discovery_data,omitempty"`
	// Determines whether the restore of the Splunk application data is enabled.
	SplunkAppData *bool `json:"splunk_app_data,omitempty"`
	// Determines whether the scheduled backup is enabled.
	Enable *bool `json:"enable,omitempty"`
	// If set, scp backup support based on keys
	UseKeys *bool `json:"use_keys,omitempty"`
	// If set, scp backup support based on keys type
	KeyType *string `json:"key_type,omitempty"`
	// If set, scp backup support to upload keys
	UploadKeys *bool `json:"upload_keys,omitempty"`
	// If set, scp backup support to download keys
	DownloadKeys         *bool `json:"download_keys,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GridDnsDtcScheduledBackup GridDnsDtcScheduledBackup

// NewGridDnsDtcScheduledBackup instantiates a new GridDnsDtcScheduledBackup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridDnsDtcScheduledBackup() *GridDnsDtcScheduledBackup {
	this := GridDnsDtcScheduledBackup{}
	return &this
}

// NewGridDnsDtcScheduledBackupWithDefaults instantiates a new GridDnsDtcScheduledBackup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridDnsDtcScheduledBackupWithDefaults() *GridDnsDtcScheduledBackup {
	this := GridDnsDtcScheduledBackup{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GridDnsDtcScheduledBackup) SetStatus(v string) {
	o.Status = &v
}

// GetExecute returns the Execute field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetExecute() string {
	if o == nil || IsNil(o.Execute) {
		var ret string
		return ret
	}
	return *o.Execute
}

// GetExecuteOk returns a tuple with the Execute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetExecuteOk() (*string, bool) {
	if o == nil || IsNil(o.Execute) {
		return nil, false
	}
	return o.Execute, true
}

// HasExecute returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasExecute() bool {
	if o != nil && !IsNil(o.Execute) {
		return true
	}

	return false
}

// SetExecute gets a reference to the given string and assigns it to the Execute field.
func (o *GridDnsDtcScheduledBackup) SetExecute(v string) {
	o.Execute = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetOperationOk() (*string, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *GridDnsDtcScheduledBackup) SetOperation(v string) {
	o.Operation = &v
}

// GetBackupType returns the BackupType field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetBackupType() string {
	if o == nil || IsNil(o.BackupType) {
		var ret string
		return ret
	}
	return *o.BackupType
}

// GetBackupTypeOk returns a tuple with the BackupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetBackupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BackupType) {
		return nil, false
	}
	return o.BackupType, true
}

// HasBackupType returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasBackupType() bool {
	if o != nil && !IsNil(o.BackupType) {
		return true
	}

	return false
}

// SetBackupType gets a reference to the given string and assigns it to the BackupType field.
func (o *GridDnsDtcScheduledBackup) SetBackupType(v string) {
	o.BackupType = &v
}

// GetKeepLocalCopy returns the KeepLocalCopy field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetKeepLocalCopy() bool {
	if o == nil || IsNil(o.KeepLocalCopy) {
		var ret bool
		return ret
	}
	return *o.KeepLocalCopy
}

// GetKeepLocalCopyOk returns a tuple with the KeepLocalCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetKeepLocalCopyOk() (*bool, bool) {
	if o == nil || IsNil(o.KeepLocalCopy) {
		return nil, false
	}
	return o.KeepLocalCopy, true
}

// HasKeepLocalCopy returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasKeepLocalCopy() bool {
	if o != nil && !IsNil(o.KeepLocalCopy) {
		return true
	}

	return false
}

// SetKeepLocalCopy gets a reference to the given bool and assigns it to the KeepLocalCopy field.
func (o *GridDnsDtcScheduledBackup) SetKeepLocalCopy(v bool) {
	o.KeepLocalCopy = &v
}

// GetBackupFrequency returns the BackupFrequency field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetBackupFrequency() string {
	if o == nil || IsNil(o.BackupFrequency) {
		var ret string
		return ret
	}
	return *o.BackupFrequency
}

// GetBackupFrequencyOk returns a tuple with the BackupFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetBackupFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.BackupFrequency) {
		return nil, false
	}
	return o.BackupFrequency, true
}

// HasBackupFrequency returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasBackupFrequency() bool {
	if o != nil && !IsNil(o.BackupFrequency) {
		return true
	}

	return false
}

// SetBackupFrequency gets a reference to the given string and assigns it to the BackupFrequency field.
func (o *GridDnsDtcScheduledBackup) SetBackupFrequency(v string) {
	o.BackupFrequency = &v
}

// GetWeekday returns the Weekday field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetWeekday() string {
	if o == nil || IsNil(o.Weekday) {
		var ret string
		return ret
	}
	return *o.Weekday
}

// GetWeekdayOk returns a tuple with the Weekday field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetWeekdayOk() (*string, bool) {
	if o == nil || IsNil(o.Weekday) {
		return nil, false
	}
	return o.Weekday, true
}

// HasWeekday returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasWeekday() bool {
	if o != nil && !IsNil(o.Weekday) {
		return true
	}

	return false
}

// SetWeekday gets a reference to the given string and assigns it to the Weekday field.
func (o *GridDnsDtcScheduledBackup) SetWeekday(v string) {
	o.Weekday = &v
}

// GetHourOfDay returns the HourOfDay field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetHourOfDay() int64 {
	if o == nil || IsNil(o.HourOfDay) {
		var ret int64
		return ret
	}
	return *o.HourOfDay
}

// GetHourOfDayOk returns a tuple with the HourOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetHourOfDayOk() (*int64, bool) {
	if o == nil || IsNil(o.HourOfDay) {
		return nil, false
	}
	return o.HourOfDay, true
}

// HasHourOfDay returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasHourOfDay() bool {
	if o != nil && !IsNil(o.HourOfDay) {
		return true
	}

	return false
}

// SetHourOfDay gets a reference to the given int64 and assigns it to the HourOfDay field.
func (o *GridDnsDtcScheduledBackup) SetHourOfDay(v int64) {
	o.HourOfDay = &v
}

// GetMinutesPastHour returns the MinutesPastHour field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetMinutesPastHour() int64 {
	if o == nil || IsNil(o.MinutesPastHour) {
		var ret int64
		return ret
	}
	return *o.MinutesPastHour
}

// GetMinutesPastHourOk returns a tuple with the MinutesPastHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetMinutesPastHourOk() (*int64, bool) {
	if o == nil || IsNil(o.MinutesPastHour) {
		return nil, false
	}
	return o.MinutesPastHour, true
}

// HasMinutesPastHour returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasMinutesPastHour() bool {
	if o != nil && !IsNil(o.MinutesPastHour) {
		return true
	}

	return false
}

// SetMinutesPastHour gets a reference to the given int64 and assigns it to the MinutesPastHour field.
func (o *GridDnsDtcScheduledBackup) SetMinutesPastHour(v int64) {
	o.MinutesPastHour = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GridDnsDtcScheduledBackup) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *GridDnsDtcScheduledBackup) SetPassword(v string) {
	o.Password = &v
}

// GetBackupServer returns the BackupServer field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetBackupServer() string {
	if o == nil || IsNil(o.BackupServer) {
		var ret string
		return ret
	}
	return *o.BackupServer
}

// GetBackupServerOk returns a tuple with the BackupServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetBackupServerOk() (*string, bool) {
	if o == nil || IsNil(o.BackupServer) {
		return nil, false
	}
	return o.BackupServer, true
}

// HasBackupServer returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasBackupServer() bool {
	if o != nil && !IsNil(o.BackupServer) {
		return true
	}

	return false
}

// SetBackupServer gets a reference to the given string and assigns it to the BackupServer field.
func (o *GridDnsDtcScheduledBackup) SetBackupServer(v string) {
	o.BackupServer = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GridDnsDtcScheduledBackup) SetPath(v string) {
	o.Path = &v
}

// GetRestoreType returns the RestoreType field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetRestoreType() string {
	if o == nil || IsNil(o.RestoreType) {
		var ret string
		return ret
	}
	return *o.RestoreType
}

// GetRestoreTypeOk returns a tuple with the RestoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetRestoreTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RestoreType) {
		return nil, false
	}
	return o.RestoreType, true
}

// HasRestoreType returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasRestoreType() bool {
	if o != nil && !IsNil(o.RestoreType) {
		return true
	}

	return false
}

// SetRestoreType gets a reference to the given string and assigns it to the RestoreType field.
func (o *GridDnsDtcScheduledBackup) SetRestoreType(v string) {
	o.RestoreType = &v
}

// GetRestoreServer returns the RestoreServer field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetRestoreServer() string {
	if o == nil || IsNil(o.RestoreServer) {
		var ret string
		return ret
	}
	return *o.RestoreServer
}

// GetRestoreServerOk returns a tuple with the RestoreServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetRestoreServerOk() (*string, bool) {
	if o == nil || IsNil(o.RestoreServer) {
		return nil, false
	}
	return o.RestoreServer, true
}

// HasRestoreServer returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasRestoreServer() bool {
	if o != nil && !IsNil(o.RestoreServer) {
		return true
	}

	return false
}

// SetRestoreServer gets a reference to the given string and assigns it to the RestoreServer field.
func (o *GridDnsDtcScheduledBackup) SetRestoreServer(v string) {
	o.RestoreServer = &v
}

// GetRestoreUsername returns the RestoreUsername field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetRestoreUsername() string {
	if o == nil || IsNil(o.RestoreUsername) {
		var ret string
		return ret
	}
	return *o.RestoreUsername
}

// GetRestoreUsernameOk returns a tuple with the RestoreUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetRestoreUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.RestoreUsername) {
		return nil, false
	}
	return o.RestoreUsername, true
}

// HasRestoreUsername returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasRestoreUsername() bool {
	if o != nil && !IsNil(o.RestoreUsername) {
		return true
	}

	return false
}

// SetRestoreUsername gets a reference to the given string and assigns it to the RestoreUsername field.
func (o *GridDnsDtcScheduledBackup) SetRestoreUsername(v string) {
	o.RestoreUsername = &v
}

// GetRestorePassword returns the RestorePassword field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetRestorePassword() string {
	if o == nil || IsNil(o.RestorePassword) {
		var ret string
		return ret
	}
	return *o.RestorePassword
}

// GetRestorePasswordOk returns a tuple with the RestorePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetRestorePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.RestorePassword) {
		return nil, false
	}
	return o.RestorePassword, true
}

// HasRestorePassword returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasRestorePassword() bool {
	if o != nil && !IsNil(o.RestorePassword) {
		return true
	}

	return false
}

// SetRestorePassword gets a reference to the given string and assigns it to the RestorePassword field.
func (o *GridDnsDtcScheduledBackup) SetRestorePassword(v string) {
	o.RestorePassword = &v
}

// GetRestorePath returns the RestorePath field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetRestorePath() string {
	if o == nil || IsNil(o.RestorePath) {
		var ret string
		return ret
	}
	return *o.RestorePath
}

// GetRestorePathOk returns a tuple with the RestorePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetRestorePathOk() (*string, bool) {
	if o == nil || IsNil(o.RestorePath) {
		return nil, false
	}
	return o.RestorePath, true
}

// HasRestorePath returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasRestorePath() bool {
	if o != nil && !IsNil(o.RestorePath) {
		return true
	}

	return false
}

// SetRestorePath gets a reference to the given string and assigns it to the RestorePath field.
func (o *GridDnsDtcScheduledBackup) SetRestorePath(v string) {
	o.RestorePath = &v
}

// GetNiosData returns the NiosData field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetNiosData() bool {
	if o == nil || IsNil(o.NiosData) {
		var ret bool
		return ret
	}
	return *o.NiosData
}

// GetNiosDataOk returns a tuple with the NiosData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetNiosDataOk() (*bool, bool) {
	if o == nil || IsNil(o.NiosData) {
		return nil, false
	}
	return o.NiosData, true
}

// HasNiosData returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasNiosData() bool {
	if o != nil && !IsNil(o.NiosData) {
		return true
	}

	return false
}

// SetNiosData gets a reference to the given bool and assigns it to the NiosData field.
func (o *GridDnsDtcScheduledBackup) SetNiosData(v bool) {
	o.NiosData = &v
}

// GetDiscoveryData returns the DiscoveryData field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetDiscoveryData() bool {
	if o == nil || IsNil(o.DiscoveryData) {
		var ret bool
		return ret
	}
	return *o.DiscoveryData
}

// GetDiscoveryDataOk returns a tuple with the DiscoveryData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetDiscoveryDataOk() (*bool, bool) {
	if o == nil || IsNil(o.DiscoveryData) {
		return nil, false
	}
	return o.DiscoveryData, true
}

// HasDiscoveryData returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasDiscoveryData() bool {
	if o != nil && !IsNil(o.DiscoveryData) {
		return true
	}

	return false
}

// SetDiscoveryData gets a reference to the given bool and assigns it to the DiscoveryData field.
func (o *GridDnsDtcScheduledBackup) SetDiscoveryData(v bool) {
	o.DiscoveryData = &v
}

// GetSplunkAppData returns the SplunkAppData field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetSplunkAppData() bool {
	if o == nil || IsNil(o.SplunkAppData) {
		var ret bool
		return ret
	}
	return *o.SplunkAppData
}

// GetSplunkAppDataOk returns a tuple with the SplunkAppData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetSplunkAppDataOk() (*bool, bool) {
	if o == nil || IsNil(o.SplunkAppData) {
		return nil, false
	}
	return o.SplunkAppData, true
}

// HasSplunkAppData returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasSplunkAppData() bool {
	if o != nil && !IsNil(o.SplunkAppData) {
		return true
	}

	return false
}

// SetSplunkAppData gets a reference to the given bool and assigns it to the SplunkAppData field.
func (o *GridDnsDtcScheduledBackup) SetSplunkAppData(v bool) {
	o.SplunkAppData = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *GridDnsDtcScheduledBackup) SetEnable(v bool) {
	o.Enable = &v
}

// GetUseKeys returns the UseKeys field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetUseKeys() bool {
	if o == nil || IsNil(o.UseKeys) {
		var ret bool
		return ret
	}
	return *o.UseKeys
}

// GetUseKeysOk returns a tuple with the UseKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetUseKeysOk() (*bool, bool) {
	if o == nil || IsNil(o.UseKeys) {
		return nil, false
	}
	return o.UseKeys, true
}

// HasUseKeys returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasUseKeys() bool {
	if o != nil && !IsNil(o.UseKeys) {
		return true
	}

	return false
}

// SetUseKeys gets a reference to the given bool and assigns it to the UseKeys field.
func (o *GridDnsDtcScheduledBackup) SetUseKeys(v bool) {
	o.UseKeys = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetKeyType() string {
	if o == nil || IsNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetKeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasKeyType() bool {
	if o != nil && !IsNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *GridDnsDtcScheduledBackup) SetKeyType(v string) {
	o.KeyType = &v
}

// GetUploadKeys returns the UploadKeys field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetUploadKeys() bool {
	if o == nil || IsNil(o.UploadKeys) {
		var ret bool
		return ret
	}
	return *o.UploadKeys
}

// GetUploadKeysOk returns a tuple with the UploadKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetUploadKeysOk() (*bool, bool) {
	if o == nil || IsNil(o.UploadKeys) {
		return nil, false
	}
	return o.UploadKeys, true
}

// HasUploadKeys returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasUploadKeys() bool {
	if o != nil && !IsNil(o.UploadKeys) {
		return true
	}

	return false
}

// SetUploadKeys gets a reference to the given bool and assigns it to the UploadKeys field.
func (o *GridDnsDtcScheduledBackup) SetUploadKeys(v bool) {
	o.UploadKeys = &v
}

// GetDownloadKeys returns the DownloadKeys field value if set, zero value otherwise.
func (o *GridDnsDtcScheduledBackup) GetDownloadKeys() bool {
	if o == nil || IsNil(o.DownloadKeys) {
		var ret bool
		return ret
	}
	return *o.DownloadKeys
}

// GetDownloadKeysOk returns a tuple with the DownloadKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridDnsDtcScheduledBackup) GetDownloadKeysOk() (*bool, bool) {
	if o == nil || IsNil(o.DownloadKeys) {
		return nil, false
	}
	return o.DownloadKeys, true
}

// HasDownloadKeys returns a boolean if a field has been set.
func (o *GridDnsDtcScheduledBackup) HasDownloadKeys() bool {
	if o != nil && !IsNil(o.DownloadKeys) {
		return true
	}

	return false
}

// SetDownloadKeys gets a reference to the given bool and assigns it to the DownloadKeys field.
func (o *GridDnsDtcScheduledBackup) SetDownloadKeys(v bool) {
	o.DownloadKeys = &v
}

func (o GridDnsDtcScheduledBackup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GridDnsDtcScheduledBackup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Execute) {
		toSerialize["execute"] = o.Execute
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.BackupType) {
		toSerialize["backup_type"] = o.BackupType
	}
	if !IsNil(o.KeepLocalCopy) {
		toSerialize["keep_local_copy"] = o.KeepLocalCopy
	}
	if !IsNil(o.BackupFrequency) {
		toSerialize["backup_frequency"] = o.BackupFrequency
	}
	if !IsNil(o.Weekday) {
		toSerialize["weekday"] = o.Weekday
	}
	if !IsNil(o.HourOfDay) {
		toSerialize["hour_of_day"] = o.HourOfDay
	}
	if !IsNil(o.MinutesPastHour) {
		toSerialize["minutes_past_hour"] = o.MinutesPastHour
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.BackupServer) {
		toSerialize["backup_server"] = o.BackupServer
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.RestoreType) {
		toSerialize["restore_type"] = o.RestoreType
	}
	if !IsNil(o.RestoreServer) {
		toSerialize["restore_server"] = o.RestoreServer
	}
	if !IsNil(o.RestoreUsername) {
		toSerialize["restore_username"] = o.RestoreUsername
	}
	if !IsNil(o.RestorePassword) {
		toSerialize["restore_password"] = o.RestorePassword
	}
	if !IsNil(o.RestorePath) {
		toSerialize["restore_path"] = o.RestorePath
	}
	if !IsNil(o.NiosData) {
		toSerialize["nios_data"] = o.NiosData
	}
	if !IsNil(o.DiscoveryData) {
		toSerialize["discovery_data"] = o.DiscoveryData
	}
	if !IsNil(o.SplunkAppData) {
		toSerialize["splunk_app_data"] = o.SplunkAppData
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.UseKeys) {
		toSerialize["use_keys"] = o.UseKeys
	}
	if !IsNil(o.KeyType) {
		toSerialize["key_type"] = o.KeyType
	}
	if !IsNil(o.UploadKeys) {
		toSerialize["upload_keys"] = o.UploadKeys
	}
	if !IsNil(o.DownloadKeys) {
		toSerialize["download_keys"] = o.DownloadKeys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GridDnsDtcScheduledBackup) UnmarshalJSON(data []byte) (err error) {
	varGridDnsDtcScheduledBackup := _GridDnsDtcScheduledBackup{}

	err = json.Unmarshal(data, &varGridDnsDtcScheduledBackup)

	if err != nil {
		return err
	}

	*o = GridDnsDtcScheduledBackup(varGridDnsDtcScheduledBackup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "execute")
		delete(additionalProperties, "operation")
		delete(additionalProperties, "backup_type")
		delete(additionalProperties, "keep_local_copy")
		delete(additionalProperties, "backup_frequency")
		delete(additionalProperties, "weekday")
		delete(additionalProperties, "hour_of_day")
		delete(additionalProperties, "minutes_past_hour")
		delete(additionalProperties, "username")
		delete(additionalProperties, "password")
		delete(additionalProperties, "backup_server")
		delete(additionalProperties, "path")
		delete(additionalProperties, "restore_type")
		delete(additionalProperties, "restore_server")
		delete(additionalProperties, "restore_username")
		delete(additionalProperties, "restore_password")
		delete(additionalProperties, "restore_path")
		delete(additionalProperties, "nios_data")
		delete(additionalProperties, "discovery_data")
		delete(additionalProperties, "splunk_app_data")
		delete(additionalProperties, "enable")
		delete(additionalProperties, "use_keys")
		delete(additionalProperties, "key_type")
		delete(additionalProperties, "upload_keys")
		delete(additionalProperties, "download_keys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGridDnsDtcScheduledBackup struct {
	value *GridDnsDtcScheduledBackup
	isSet bool
}

func (v NullableGridDnsDtcScheduledBackup) Get() *GridDnsDtcScheduledBackup {
	return v.value
}

func (v *NullableGridDnsDtcScheduledBackup) Set(val *GridDnsDtcScheduledBackup) {
	v.value = val
	v.isSet = true
}

func (v NullableGridDnsDtcScheduledBackup) IsSet() bool {
	return v.isSet
}

func (v *NullableGridDnsDtcScheduledBackup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridDnsDtcScheduledBackup(val *GridDnsDtcScheduledBackup) *NullableGridDnsDtcScheduledBackup {
	return &NullableGridDnsDtcScheduledBackup{value: val, isSet: true}
}

func (v NullableGridDnsDtcScheduledBackup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridDnsDtcScheduledBackup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
