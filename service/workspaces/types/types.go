// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Describes a modification to the configuration of Bring Your Own License (BYOL)
// for the specified account.
type AccountModification struct {

	// The IP address range, specified as an IPv4 CIDR block, for the management
	// network interface used for the account.
	DedicatedTenancyManagementCidrRange *string

	// The status of BYOL (whether BYOL is being enabled or disabled).
	DedicatedTenancySupport DedicatedTenancySupportResultEnum

	// The error code that is returned if the configuration of BYOL cannot be modified.
	ErrorCode *string

	// The text of the error message that is returned if the configuration of BYOL
	// cannot be modified.
	ErrorMessage *string

	// The state of the modification to the configuration of BYOL.
	ModificationState DedicatedTenancyModificationStateEnum

	// The timestamp when the modification of the BYOL configuration was started.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Describes an Amazon WorkSpaces client.
type ClientProperties struct {

	// Specifies whether users can cache their credentials on the Amazon WorkSpaces
	// client. When enabled, users can choose to reconnect to their WorkSpaces without
	// re-entering their credentials.
	ReconnectEnabled ReconnectEnum

	noSmithyDocumentSerde
}

// Information about the Amazon WorkSpaces client.
type ClientPropertiesResult struct {

	// Information about the Amazon WorkSpaces client.
	ClientProperties *ClientProperties

	// The resource identifier, in the form of a directory ID.
	ResourceId *string

	noSmithyDocumentSerde
}

// Describes the compute type of the bundle.
type ComputeType struct {

	// The compute type.
	Name Compute

	noSmithyDocumentSerde
}

// Describes a connection alias. Connection aliases are used for cross-Region
// redirection. For more information, see  Cross-Region Redirection for Amazon
// WorkSpaces
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html).
type ConnectionAlias struct {

	// The identifier of the connection alias.
	AliasId *string

	// The association status of the connection alias.
	Associations []ConnectionAliasAssociation

	// The connection string specified for the connection alias. The connection string
	// must be in the form of a fully qualified domain name (FQDN), such as
	// www.example.com.
	ConnectionString *string

	// The identifier of the Amazon Web Services account that owns the connection
	// alias.
	OwnerAccountId *string

	// The current state of the connection alias.
	State ConnectionAliasState

	noSmithyDocumentSerde
}

// Describes a connection alias association that is used for cross-Region
// redirection. For more information, see  Cross-Region Redirection for Amazon
// WorkSpaces
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html).
type ConnectionAliasAssociation struct {

	// The identifier of the Amazon Web Services account that associated the connection
	// alias with a directory.
	AssociatedAccountId *string

	// The association status of the connection alias.
	AssociationStatus AssociationStatus

	// The identifier of the connection alias association. You use the connection
	// identifier in the DNS TXT record when you're configuring your DNS routing
	// policies.
	ConnectionIdentifier *string

	// The identifier of the directory associated with a connection alias.
	ResourceId *string

	noSmithyDocumentSerde
}

// Describes the permissions for a connection alias. Connection aliases are used
// for cross-Region redirection. For more information, see  Cross-Region
// Redirection for Amazon WorkSpaces
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/cross-region-redirection.html).
type ConnectionAliasPermission struct {

	// Indicates whether the specified Amazon Web Services account is allowed to
	// associate the connection alias with a directory.
	//
	// This member is required.
	AllowAssociation *bool

	// The identifier of the Amazon Web Services account that the connection alias is
	// shared with.
	//
	// This member is required.
	SharedAccountId *string

	noSmithyDocumentSerde
}

// Describes the default values that are used to create WorkSpaces. For more
// information, see Update Directory Details for Your WorkSpaces
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/update-directory-details.html).
type DefaultWorkspaceCreationProperties struct {

	// The identifier of the default security group to apply to WorkSpaces when they
	// are created. For more information, see  Security Groups for Your WorkSpaces
	// (https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces-security-groups.html).
	CustomSecurityGroupId *string

	// The organizational unit (OU) in the directory for the WorkSpace machine
	// accounts.
	DefaultOu *string

	// Specifies whether to automatically assign an Elastic public IP address to
	// WorkSpaces in this directory by default. If enabled, the Elastic public IP
	// address allows outbound internet access from your WorkSpaces when you’re using
	// an internet gateway in the Amazon VPC in which your WorkSpaces are located. If
	// you're using a Network Address Translation (NAT) gateway for outbound internet
	// access from your VPC, or if your WorkSpaces are in public subnets and you
	// manually assign them Elastic IP addresses, you should disable this setting. This
	// setting applies to new WorkSpaces that you launch or to existing WorkSpaces that
	// you rebuild. For more information, see  Configure a VPC for Amazon WorkSpaces
	// (https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces-vpc.html).
	EnableInternetAccess *bool

	// Specifies whether maintenance mode is enabled for WorkSpaces. For more
	// information, see WorkSpace Maintenance
	// (https://docs.aws.amazon.com/workspaces/latest/adminguide/workspace-maintenance.html).
	EnableMaintenanceMode *bool

	// Specifies whether the directory is enabled for Amazon WorkDocs.
	EnableWorkDocs *bool

	// Specifies whether WorkSpace users are local administrators on their WorkSpaces.
	UserEnabledAsLocalAdministrator *bool

	noSmithyDocumentSerde
}

// Describes a WorkSpace that cannot be created.
type FailedCreateWorkspaceRequest struct {

	// The error code that is returned if the WorkSpace cannot be created.
	ErrorCode *string

	// The text of the error message that is returned if the WorkSpace cannot be
	// created.
	ErrorMessage *string

	// Information about the WorkSpace.
	WorkspaceRequest *WorkspaceRequest

	noSmithyDocumentSerde
}

// Describes a WorkSpace that could not be rebooted. (RebootWorkspaces), rebuilt
// (RebuildWorkspaces), restored (RestoreWorkspace), terminated
// (TerminateWorkspaces), started (StartWorkspaces), or stopped (StopWorkspaces).
type FailedWorkspaceChangeRequest struct {

	// The error code that is returned if the WorkSpace cannot be rebooted.
	ErrorCode *string

	// The text of the error message that is returned if the WorkSpace cannot be
	// rebooted.
	ErrorMessage *string

	// The identifier of the WorkSpace.
	WorkspaceId *string

	noSmithyDocumentSerde
}

// Describes the Amazon Web Services accounts that have been granted permission to
// use a shared image. For more information about sharing images, see  Share or
// Unshare a Custom WorkSpaces Image
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/share-custom-image.html).
type ImagePermission struct {

	// The identifier of the Amazon Web Services account that an image has been shared
	// with.
	SharedAccountId *string

	noSmithyDocumentSerde
}

// Describes a rule for an IP access control group.
type IpRuleItem struct {

	// The IP address range, in CIDR notation.
	IpRule *string

	// The description.
	RuleDesc *string

	noSmithyDocumentSerde
}

// Describes a WorkSpace modification.
type ModificationState struct {

	// The resource.
	Resource ModificationResourceEnum

	// The modification state.
	State ModificationStateEnum

	noSmithyDocumentSerde
}

// The operating system that the image is running.
type OperatingSystem struct {

	// The operating system.
	Type OperatingSystemType

	noSmithyDocumentSerde
}

// Describes the information used to reboot a WorkSpace.
type RebootRequest struct {

	// The identifier of the WorkSpace.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

// Describes the information used to rebuild a WorkSpace.
type RebuildRequest struct {

	// The identifier of the WorkSpace.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

// Describes the root volume for a WorkSpace bundle.
type RootStorage struct {

	// The size of the root volume.
	Capacity *string

	noSmithyDocumentSerde
}

// Describes the self-service permissions for a directory. For more information,
// see Enable Self-Service WorkSpace Management Capabilities for Your Users
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/enable-user-self-service-workspace-management.html).
type SelfservicePermissions struct {

	// Specifies whether users can change the compute type (bundle) for their
	// WorkSpace.
	ChangeComputeType ReconnectEnum

	// Specifies whether users can increase the volume size of the drives on their
	// WorkSpace.
	IncreaseVolumeSize ReconnectEnum

	// Specifies whether users can rebuild the operating system of a WorkSpace to its
	// original state.
	RebuildWorkspace ReconnectEnum

	// Specifies whether users can restart their WorkSpace.
	RestartWorkspace ReconnectEnum

	// Specifies whether users can switch the running mode of their WorkSpace.
	SwitchRunningMode ReconnectEnum

	noSmithyDocumentSerde
}

// Describes a snapshot.
type Snapshot struct {

	// The time when the snapshot was created.
	SnapshotTime *time.Time

	noSmithyDocumentSerde
}

// Information used to start a WorkSpace.
type StartRequest struct {

	// The identifier of the WorkSpace.
	WorkspaceId *string

	noSmithyDocumentSerde
}

// Describes the information used to stop a WorkSpace.
type StopRequest struct {

	// The identifier of the WorkSpace.
	WorkspaceId *string

	noSmithyDocumentSerde
}

// Describes a tag.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	Value *string

	noSmithyDocumentSerde
}

// Describes the information used to terminate a WorkSpace.
type TerminateRequest struct {

	// The identifier of the WorkSpace.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

// Describes whether a WorkSpace image needs to be updated with the latest drivers
// and other components required by Amazon WorkSpaces. Only Windows 10 WorkSpace
// images can be programmatically updated at this time.
type UpdateResult struct {

	// A description of whether updates for the WorkSpace image are pending or
	// available.
	Description *string

	// Indicates whether updated drivers or other components are available for the
	// specified WorkSpace image.
	UpdateAvailable *bool

	noSmithyDocumentSerde
}

// Describes the user volume for a WorkSpace bundle.
type UserStorage struct {

	// The size of the user volume.
	Capacity *string

	noSmithyDocumentSerde
}

// Describes a WorkSpace.
type Workspace struct {

	// The identifier of the bundle used to create the WorkSpace.
	BundleId *string

	// The name of the WorkSpace, as seen by the operating system. The format of this
	// name varies. For more information, see  Launch a WorkSpace
	// (https://docs.aws.amazon.com/workspaces/latest/adminguide/launch-workspaces-tutorials.html).
	ComputerName *string

	// The identifier of the Directory Service directory for the WorkSpace.
	DirectoryId *string

	// The error code that is returned if the WorkSpace cannot be created.
	ErrorCode *string

	// The text of the error message that is returned if the WorkSpace cannot be
	// created.
	ErrorMessage *string

	// The IP address of the WorkSpace.
	IpAddress *string

	// The modification states of the WorkSpace.
	ModificationStates []ModificationState

	// Indicates whether the data stored on the root volume is encrypted.
	RootVolumeEncryptionEnabled *bool

	// The operational state of the WorkSpace. After a WorkSpace is terminated, the
	// TERMINATED state is returned only briefly before the WorkSpace directory
	// metadata is cleaned up, so this state is rarely returned. To confirm that a
	// WorkSpace is terminated, check for the WorkSpace ID by using  DescribeWorkSpaces
	// (https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeWorkspaces.html).
	// If the WorkSpace ID isn't returned, then the WorkSpace has been successfully
	// terminated.
	State WorkspaceState

	// The identifier of the subnet for the WorkSpace.
	SubnetId *string

	// The user for the WorkSpace.
	UserName *string

	// Indicates whether the data stored on the user volume is encrypted.
	UserVolumeEncryptionEnabled *bool

	// The symmetric KMS key used to encrypt data stored on your WorkSpace. Amazon
	// WorkSpaces does not support asymmetric KMS keys.
	VolumeEncryptionKey *string

	// The identifier of the WorkSpace.
	WorkspaceId *string

	// The properties of the WorkSpace.
	WorkspaceProperties *WorkspaceProperties

	noSmithyDocumentSerde
}

// The device types and operating systems that can be used to access a WorkSpace.
// For more information, see Amazon WorkSpaces Client Network Requirements
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/workspaces-network-requirements.html).
type WorkspaceAccessProperties struct {

	// Indicates whether users can use Android and Android-compatible Chrome OS devices
	// to access their WorkSpaces.
	DeviceTypeAndroid AccessPropertyValue

	// Indicates whether users can use Chromebooks to access their WorkSpaces.
	DeviceTypeChromeOs AccessPropertyValue

	// Indicates whether users can use iOS devices to access their WorkSpaces.
	DeviceTypeIos AccessPropertyValue

	// Indicates whether users can use Linux clients to access their WorkSpaces.
	DeviceTypeLinux AccessPropertyValue

	// Indicates whether users can use macOS clients to access their WorkSpaces.
	DeviceTypeOsx AccessPropertyValue

	// Indicates whether users can access their WorkSpaces through a web browser.
	DeviceTypeWeb AccessPropertyValue

	// Indicates whether users can use Windows clients to access their WorkSpaces.
	DeviceTypeWindows AccessPropertyValue

	// Indicates whether users can use zero client devices to access their WorkSpaces.
	DeviceTypeZeroClient AccessPropertyValue

	noSmithyDocumentSerde
}

// Describes a WorkSpace bundle.
type WorkspaceBundle struct {

	// The identifier of the bundle.
	BundleId *string

	// The compute type of the bundle. For more information, see Amazon WorkSpaces
	// Bundles (http://aws.amazon.com/workspaces/details/#Amazon_WorkSpaces_Bundles).
	ComputeType *ComputeType

	// The time when the bundle was created.
	CreationTime *time.Time

	// The description of the bundle.
	Description *string

	// The identifier of the image that was used to create the bundle.
	ImageId *string

	// The last time that the bundle was updated.
	LastUpdatedTime *time.Time

	// The name of the bundle.
	Name *string

	// The owner of the bundle. This is the account identifier of the owner, or AMAZON
	// if the bundle is provided by Amazon Web Services.
	Owner *string

	// The size of the root volume.
	RootStorage *RootStorage

	// The size of the user volume.
	UserStorage *UserStorage

	noSmithyDocumentSerde
}

// Describes the connection status of a WorkSpace.
type WorkspaceConnectionStatus struct {

	// The connection state of the WorkSpace. The connection state is unknown if the
	// WorkSpace is stopped.
	ConnectionState ConnectionState

	// The timestamp of the connection status check.
	ConnectionStateCheckTimestamp *time.Time

	// The timestamp of the last known user connection.
	LastKnownUserConnectionTimestamp *time.Time

	// The identifier of the WorkSpace.
	WorkspaceId *string

	noSmithyDocumentSerde
}

// Describes the default properties that are used for creating WorkSpaces. For more
// information, see Update Directory Details for Your WorkSpaces
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/update-directory-details.html).
type WorkspaceCreationProperties struct {

	// The identifier of your custom security group.
	CustomSecurityGroupId *string

	// The default organizational unit (OU) for your WorkSpaces directories. This
	// string must be the full Lightweight Directory Access Protocol (LDAP)
	// distinguished name for the target domain and OU. It must be in the form
	// "OU=value,DC=value,DC=value", where value is any string of characters, and the
	// number of domain components (DCs) is two or more. For example,
	// OU=WorkSpaces_machines,DC=machines,DC=example,DC=com.
	//
	// * To avoid errors,
	// certain characters in the distinguished name must be escaped. For more
	// information, see  Distinguished Names
	// (https://docs.microsoft.com/previous-versions/windows/desktop/ldap/distinguished-names)
	// in the Microsoft documentation.
	//
	// * The API doesn't validate whether the OU
	// exists.
	DefaultOu *string

	// Indicates whether internet access is enabled for your WorkSpaces.
	EnableInternetAccess *bool

	// Indicates whether maintenance mode is enabled for your WorkSpaces. For more
	// information, see WorkSpace Maintenance
	// (https://docs.aws.amazon.com/workspaces/latest/adminguide/workspace-maintenance.html).
	EnableMaintenanceMode *bool

	// Indicates whether Amazon WorkDocs is enabled for your WorkSpaces.
	//
	// If WorkDocs
	// is already enabled for a WorkSpaces directory and you disable it, new WorkSpaces
	// launched in the directory will not have WorkDocs enabled. However, WorkDocs
	// remains enabled for any existing WorkSpaces, unless you either disable users'
	// access to WorkDocs or you delete the WorkDocs site. To disable users' access to
	// WorkDocs, see Disabling Users
	// (https://docs.aws.amazon.com/workdocs/latest/adminguide/inactive-user.html) in
	// the Amazon WorkDocs Administration Guide. To delete a WorkDocs site, see
	// Deleting a Site
	// (https://docs.aws.amazon.com/workdocs/latest/adminguide/manage-sites.html) in
	// the Amazon WorkDocs Administration Guide. If you enable WorkDocs on a directory
	// that already has existing WorkSpaces, the existing WorkSpaces and any new
	// WorkSpaces that are launched in the directory will have WorkDocs enabled.
	EnableWorkDocs *bool

	// Indicates whether users are local administrators of their WorkSpaces.
	UserEnabledAsLocalAdministrator *bool

	noSmithyDocumentSerde
}

// Describes a directory that is used with Amazon WorkSpaces.
type WorkspaceDirectory struct {

	// The directory alias.
	Alias *string

	// The user name for the service account.
	CustomerUserName *string

	// The directory identifier.
	DirectoryId *string

	// The name of the directory.
	DirectoryName *string

	// The directory type.
	DirectoryType WorkspaceDirectoryType

	// The IP addresses of the DNS servers for the directory.
	DnsIpAddresses []string

	// The identifier of the IAM role. This is the role that allows Amazon WorkSpaces
	// to make calls to other services, such as Amazon EC2, on your behalf.
	IamRoleId *string

	// The identifiers of the IP access control groups associated with the directory.
	IpGroupIds []string

	// The registration code for the directory. This is the code that users enter in
	// their Amazon WorkSpaces client application to connect to the directory.
	RegistrationCode *string

	// The default self-service permissions for WorkSpaces in the directory.
	SelfservicePermissions *SelfservicePermissions

	// The state of the directory's registration with Amazon WorkSpaces. After a
	// directory is deregistered, the DEREGISTERED state is returned very briefly
	// before the directory metadata is cleaned up, so this state is rarely returned.
	// To confirm that a directory is deregistered, check for the directory ID by using
	// DescribeWorkspaceDirectories
	// (https://docs.aws.amazon.com/workspaces/latest/api/API_DescribeWorkspaceDirectories.html).
	// If the directory ID isn't returned, then the directory has been successfully
	// deregistered.
	State WorkspaceDirectoryState

	// The identifiers of the subnets used with the directory.
	SubnetIds []string

	// Specifies whether the directory is dedicated or shared. To use Bring Your Own
	// License (BYOL), this value must be set to DEDICATED. For more information, see
	// Bring Your Own Windows Desktop Images
	// (https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.html).
	Tenancy Tenancy

	// The devices and operating systems that users can use to access WorkSpaces.
	WorkspaceAccessProperties *WorkspaceAccessProperties

	// The default creation properties for all WorkSpaces in the directory.
	WorkspaceCreationProperties *DefaultWorkspaceCreationProperties

	// The identifier of the security group that is assigned to new WorkSpaces.
	WorkspaceSecurityGroupId *string

	noSmithyDocumentSerde
}

// Describes a WorkSpace image.
type WorkspaceImage struct {

	// The date when the image was created. If the image has been shared, the Amazon
	// Web Services account that the image has been shared with sees the original
	// creation date of the image.
	Created *time.Time

	// The description of the image.
	Description *string

	// The error code that is returned for the image.
	ErrorCode *string

	// The text of the error message that is returned for the image.
	ErrorMessage *string

	// The identifier of the image.
	ImageId *string

	// The name of the image.
	Name *string

	// The operating system that the image is running.
	OperatingSystem *OperatingSystem

	// The identifier of the Amazon Web Services account that owns the image.
	OwnerAccountId *string

	// Specifies whether the image is running on dedicated hardware. When Bring Your
	// Own License (BYOL) is enabled, this value is set to DEDICATED. For more
	// information, see Bring Your Own Windows Desktop Images
	// (https://docs.aws.amazon.com/workspaces/latest/adminguide/byol-windows-images.html).
	RequiredTenancy WorkspaceImageRequiredTenancy

	// The status of the image.
	State WorkspaceImageState

	// The updates (if any) that are available for the specified image.
	Updates *UpdateResult

	noSmithyDocumentSerde
}

// Describes a WorkSpace.
type WorkspaceProperties struct {

	// The compute type. For more information, see Amazon WorkSpaces Bundles
	// (http://aws.amazon.com/workspaces/details/#Amazon_WorkSpaces_Bundles).
	ComputeTypeName Compute

	// The size of the root volume. For important information about how to modify the
	// size of the root and user volumes, see Modify a WorkSpace
	// (https://docs.aws.amazon.com/workspaces/latest/adminguide/modify-workspaces.html).
	RootVolumeSizeGib *int32

	// The running mode. For more information, see Manage the WorkSpace Running Mode
	// (https://docs.aws.amazon.com/workspaces/latest/adminguide/running-mode.html).
	RunningMode RunningMode

	// The time after a user logs off when WorkSpaces are automatically stopped.
	// Configured in 60-minute intervals.
	RunningModeAutoStopTimeoutInMinutes *int32

	// The size of the user storage. For important information about how to modify the
	// size of the root and user volumes, see Modify a WorkSpace
	// (https://docs.aws.amazon.com/workspaces/latest/adminguide/modify-workspaces.html).
	UserVolumeSizeGib *int32

	noSmithyDocumentSerde
}

// Describes the information used to create a WorkSpace.
type WorkspaceRequest struct {

	// The identifier of the bundle for the WorkSpace. You can use
	// DescribeWorkspaceBundles to list the available bundles.
	//
	// This member is required.
	BundleId *string

	// The identifier of the Directory Service directory for the WorkSpace. You can use
	// DescribeWorkspaceDirectories to list the available directories.
	//
	// This member is required.
	DirectoryId *string

	// The user name of the user for the WorkSpace. This user name must exist in the
	// Directory Service directory for the WorkSpace.
	//
	// This member is required.
	UserName *string

	// Indicates whether the data stored on the root volume is encrypted.
	RootVolumeEncryptionEnabled *bool

	// The tags for the WorkSpace.
	Tags []Tag

	// Indicates whether the data stored on the user volume is encrypted.
	UserVolumeEncryptionEnabled *bool

	// The symmetric KMS key used to encrypt data stored on your WorkSpace. Amazon
	// WorkSpaces does not support asymmetric KMS keys.
	VolumeEncryptionKey *string

	// The WorkSpace properties.
	WorkspaceProperties *WorkspaceProperties

	noSmithyDocumentSerde
}

// Describes an IP access control group.
type WorkspacesIpGroup struct {

	// The description of the group.
	GroupDesc *string

	// The identifier of the group.
	GroupId *string

	// The name of the group.
	GroupName *string

	// The rules.
	UserRules []IpRuleItem

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
