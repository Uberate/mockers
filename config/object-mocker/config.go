package main

type ObjectMockerConfig struct {
	// NodeName define this node name, it will identify of cluster mode. We provide some different cluster mode. And all
	// of them can be used together. About the cluster mode, see: BackupNodes, MasterNodes config. If NodeName is emtpy
	// use the hash of start time, config-info(origin file), and a random int64 value.
	//
	// If in a cluster, has two node has same NodeName, the last add node will be rejected. If a cluster join in, all
	// the cluster nodes will be rejected directly.
	NodeName string `json:"node_name" yaml:"node_name"`

	// EnableWebSetting if is true, enable the setting change by web request. Else, it sees as read-only.
	EnableWebSetting bool `json:"enable_web_setting" yaml:"enable_web_setting"`

	// EnableOperateLog will log all operate to InnerInfoPath/operate-log. You can set limit of log.
	EnableOperateLog bool `json:"enable_operate_log" yaml:"enable_operate_log"`

	// InnerInfoPath, default is /_inner_path/, you can change it to any path. And the config info will save to here.
	// This key can't be re-writen.
	InnerInfoPath string `json:"inner_info_path" yaml:"inner_info_path"`

	// BackupNodes define the node is a part of a backup-cluster, and one node only can join one cluster. All the
	// backup-node can auto sync data and lookup other backup nodes. But on node only can join one backup-cluster. When
	// node is running, this param can be changed if config set EnableAutoBackupNodes is true.
	//
	// If BackupNodes set to empty, and EnableAutoBackupNodes is false, the nodes can't join any backup-cluster.
	//
	// And the node can't dynamic join a backup-cluster.
	//
	// About the more info of backup setting, see'./cluster_mod.md'.
	BackupNodes []string `json:"backup_master_nodes" yaml:"backup_master_nodes"`

	BackupMode string

	// EnableAutoBackupNodes will enable node auto syne other nodes' info. It will sync info from backup-cluster. It
	// doesn't affect the data sync. The data sync behavior only be controlled by BackupMode. If EnableAutoBackupNodes
	// is false, this node only sync data from BackupNodes.
	EnableAutoBackupNodes bool

	DistributeMasterNodes []string `json:"master_nodes" yaml:"master_nodes"`
}
