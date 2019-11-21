/*
Copyright 2019 The OpenEBS Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package blockdevice

// BlockDevice is an internal representation of any block device present on the system.
// All data related to that device will be held by this struct
//
// 1. Example blockdevice struct for a partition /dev/sda1
// 		{
// 				UUID:blockdevice-4c25d69f9adc868f61e3d891cf3a5613
// 				NodeAttributes:map[hostname:my-machine]
// 				Path:/dev/sda1
// 				FSInfo:{
// 					FileSystem:ext4
// 					MountPoint:[/home]
// 				}
// 				Parent:/dev/sda
// 				Partitions:[]
// 				Holders:[]
// 				Slaves:[]
// 				Status:{
// 					State:Active
// 					ClaimPhase:Unclaimed
// 				}
// 			}
//
// 2. Example blockdevice struct for a partition that is part of an LVM
// 		{
// 				UUID:blockdevice-4c25d69f9adc868f61e3d891cf3a5613
// 				NodeAttributes:map[hostname:my-machine]
// 				Path:/dev/sda1
// 				FSInfo:{
// 					FileSystem:LVM2_member
// 					MountPoint:[]
// 				}
// 				Parent:/dev/sda
// 				Partitions:[]
// 				Holders:[/dev/dm-0]
// 				Slaves:[]
// 				Status:{
// 					State:Active
// 					ClaimPhase:Unclaimed
// 				}
// 			}
// 3. Example blockdevice struct for an LVM carved from nvme partition
// 		{
// 				UUID:blockdevice-4c25d69f9adc868f61e3d891cf3a5613
// 				NodeAttributes:map[hostname:my-machine]
// 				Path:/dev/dm-0
// 				FSInfo:{
// 					FileSystem:ext4
// 					MountPoint:[]
// 				}
// 				Parent:
// 				Partitions:[]
// 				Holders:[]
// 				Slaves:[/dev/nvme0n1p1 /dev/nvme0n1p2]
// 				Status:{
// 					State:Active
// 					ClaimPhase:Unclaimed
// 				}
// 			}
//
type BlockDevice struct {
	// UUID is the UUID for this BlockDevice generated by NDM
	// It is an md5 hash generated based on the various parameters
	// of the BlockDevice. eg: blockdevice-xxx
	UUID string

	// NodeAttributes contains the details of the node on which
	// the BlockDevice is attached
	NodeAttributes NodeAttribute

	// Path is dev path of a device
	// eg : /dev/sda, /dev/sda1, /dev/dm-0
	Path string

	// FSInfo contains the file system related information of this
	// BlockDevice if it exists
	FSInfo FileSystemInformation

	// DeviceType is the type of the blockdevice. can be sparse/disk/partition etc
	DeviceType string

	// Parent is the parent device of this blockdevice, if it exists.
	// It will always be a single device.
	Parent string

	// Partitions is the list of partitions(again blockdevices) for
	// this blockdevice, if it exists.
	Partitions []string

	// Holders is the list of blockdevices that are held by this blockdevice.
	// eg: sda1 can hold dm-0. Then the list of sda1 will contain dm-0.
	Holders []string

	// Slaves is the list of blockdevices to which this blockdevice is a slave.
	// Slaves are represented in SysFS under the block device layer,
	// so they are dependent on the device
	// eg: dm-0 is a slave to sda1. Then the list of dm-0 will contain sda1
	Slaves []string

	// Status contains the state of the blockdevice
	Status Status
}

// NodeAttribute is the representing the various attributes of the machine
// on which this block device is present
type NodeAttribute map[string]string

const (
	// HostName is the hostname of the system on which this BD is present
	HostName string = "hostname"

	// NodeName is the nodename (may be FQDN) on which this BD is present
	NodeName string = "nodename"

	// ZoneName is the zone in which the system is present.
	//
	// NOTE: Valid only for cloud providers
	ZoneName string = "zone"

	// RegionName is the region in which the system is present.
	//
	// NOTE: Valid only for cloud providers
	RegionName string = "region"
)

const (
	// SparseBlockDeviceType is the sparse blockdevice type
	SparseBlockDeviceType = "sparse"
	// BlockDeviceType is the type for blockdevice.
	BlockDeviceType = "blockdevice"
)

// FileSystemInformation contains the filesystem and mount information of blockdevice, if present
type FileSystemInformation struct {
	// FileSystem is the filesystem present on the blockdevice
	FileSystem string

	// MountPoint is the list of mountpoints at which this blockdevice is mounted
	MountPoint []string
}

// Status is used to represent the status of the blockdevice
type Status struct {
	// State is the state of this BD like Active(Online), Inactive(Offline) or
	// Unknown
	State string

	// ClaimPhase is the phase of this BD when is it is being used by NDM consumers
	ClaimPhase string
}

const (
	// Active means blockdevice is available on the host machine
	Active string = "Active"
	// Inactive means blockdevice is currently not available on the host machine
	Inactive string = "Inactive"
	// Unknown means the state cannot be determined at this point of time
	Unknown string = "Unknown"

	// Claimed means the blockdevice is in use
	Claimed string = "Claimed"
	// Released means the blockdevice is not in use, but cannot be claimed,
	// because of some pending cleanup tasks
	Released string = "Released"
	// Unclaimed means the blockdevice is free and is availbale for
	// claiming
	Unclaimed string = "Unclaimed"
)
