package synologyapi

import "fmt"
import "net/http"
import "log"
import "encoding/json"

// SynologyConnection encapsulates required connection info to specific NAS
type SynologyConnection struct {
	origin string
	token  string
}

type authResponse struct {
	Data    authResponseData `json:"data"`
	Success bool             `json:"success"`
}

type authResponseData struct {
	IsPortalPort bool   `json:"is_portal_port"`
	Sid          string `json:"sid"`
}

type systemUtilizationResponse struct {
	Data    SystemUtilizationData `json:"data"`
	Success bool                  `json:"success"`
}

// SystemUtilizationData Synology response object
type SystemUtilizationData struct {
	CPU struct {
		One5minLoad int    `json:"15min_load"`
		OneMinLoad  int    `json:"1min_load"`
		FiveMinLoad int    `json:"5min_load"`
		Device      string `json:"device"`
		OtherLoad   int    `json:"other_load"`
		SystemLoad  int    `json:"system_load"`
		UserLoad    int    `json:"user_load"`
	} `json:"cpu"`
	Disk struct {
		Disk []struct {
			Device      string `json:"device"`
			DisplayName string `json:"display_name"`
			ReadAccess  int    `json:"read_access"`
			ReadByte    int    `json:"read_byte"`
			Type        string `json:"type"`
			Utilization int    `json:"utilization"`
			WriteAccess int    `json:"write_access"`
			WriteByte   int    `json:"write_byte"`
		} `json:"disk"`
		Total struct {
			Device      string `json:"device"`
			ReadAccess  int    `json:"read_access"`
			ReadByte    int    `json:"read_byte"`
			Utilization int    `json:"utilization"`
			WriteAccess int    `json:"write_access"`
			WriteByte   int    `json:"write_byte"`
		} `json:"total"`
	} `json:"disk"`
	Lun []struct {
		Device          string `json:"device"`
		Path            string `json:"path"`
		QueueExe        int    `json:"queue_exe"`
		QueueQue        int    `json:"queue_que"`
		QueueWbk        int    `json:"queue_wbk"`
		ReadAvgCmdSize  int    `json:"read_avg_cmd_size"`
		ReadAvgLatency  int    `json:"read_avg_latency"`
		ReadBytes       int    `json:"read_bytes"`
		ReadCmdCount    int    `json:"read_cmd_count"`
		RxAvgLatency    int    `json:"rx_avg_latency"`
		TotalCmdCount   int    `json:"total_cmd_count"`
		TotalIoLatency  int    `json:"total_io_latency"`
		TotalIops       int    `json:"total_iops"`
		TotalThroughput int    `json:"total_throughput"`
		TxAvgLatency    int    `json:"tx_avg_latency"`
		Type            string `json:"type"`
		UUID            string `json:"uuid"`
		WriteAvgCmdSize int    `json:"write_avg_cmd_size"`
		WriteAvgLatency int    `json:"write_avg_latency"`
		WriteBytes      int    `json:"write_bytes"`
		WriteCmdCount   int    `json:"write_cmd_count"`
	} `json:"lun"`
	Memory struct {
		AvailReal  int    `json:"avail_real"`
		AvailSwap  int    `json:"avail_swap"`
		Buffer     int    `json:"buffer"`
		Cached     int    `json:"cached"`
		Device     string `json:"device"`
		MemorySize int    `json:"memory_size"`
		RealUsage  int    `json:"real_usage"`
		SiDisk     int    `json:"si_disk"`
		SoDisk     int    `json:"so_disk"`
		SwapUsage  int    `json:"swap_usage"`
		TotalReal  int    `json:"total_real"`
		TotalSwap  int    `json:"total_swap"`
	} `json:"memory"`
	Network []struct {
		Device string `json:"device"`
		Rx     int    `json:"rx"`
		Tx     int    `json:"tx"`
	} `json:"network"`
	Space struct {
		Total struct {
			Device      string `json:"device"`
			ReadAccess  int    `json:"read_access"`
			ReadByte    int    `json:"read_byte"`
			Utilization int    `json:"utilization"`
			WriteAccess int    `json:"write_access"`
			WriteByte   int    `json:"write_byte"`
		} `json:"total"`
		Volume []struct {
			Device      string `json:"device"`
			DisplayName string `json:"display_name"`
			ReadAccess  int    `json:"read_access"`
			ReadByte    int    `json:"read_byte"`
			Utilization int    `json:"utilization"`
			WriteAccess int    `json:"write_access"`
			WriteByte   int    `json:"write_byte"`
		} `json:"volume"`
	} `json:"space"`
	Time int `json:"time"`
}

type sharedFolderReponse struct {
	Data    SharedFolderData `json:"data"`
	Success bool             `json:"success"`
}

// SharedFolderData contains specific information on individual Shared Folders
type SharedFolderData struct {
	Shares []struct {
		Desc           string  `json:"desc"`
		IsUsbShare     bool    `json:"is_usb_share"`
		Name           string  `json:"name"`
		QuotaValue     float64 `json:"quota_value"`
		ShareQuotaUsed float64 `json:"share_quota_used"`
		UUID           string  `json:"uuid"`
		VolPath        string  `json:"vol_path"`
	} `json:"shares"`
	Total int `json:"total"`
}

type storageResponse struct {
	Data    StorageData `json:"data"`
	Success bool        `json:"success"`
}

// StorageData contains alot of information about available storage in the Synology
type StorageData struct {
	Disks []struct {
		AdvProgress        string `json:"adv_progress"`
		AdvStatus          string `json:"adv_status"`
		BelowRemainLifeThr bool   `json:"below_remain_life_thr"`
		Container          struct {
			Order                int    `json:"order"`
			Str                  string `json:"str"`
			SupportPwrBtnDisable bool   `json:"supportPwrBtnDisable"`
			Type                 string `json:"type"`
		} `json:"container"`
		Device             string `json:"device"`
		DisableSecera      bool   `json:"disable_secera"`
		DiskType           string `json:"diskType"`
		DiskCode           string `json:"disk_code"`
		EraseTime          int    `json:"erase_time"`
		ExceedBadSectorThr bool   `json:"exceed_bad_sector_thr"`
		Firm               string `json:"firm"`
		HasSystem          bool   `json:"has_system"`
		ID                 string `json:"id"`
		Is4Kn              bool   `json:"is4Kn"`
		IsSsd              bool   `json:"isSsd"`
		IsSynoPartition    bool   `json:"isSynoPartition"`
		IsErasing          bool   `json:"is_erasing"`
		LongName           string `json:"longName"`
		Model              string `json:"model"`
		Name               string `json:"name"`
		NumID              int    `json:"num_id"`
		Order              int    `json:"order"`
		OverviewStatus     string `json:"overview_status"`
		PciSlot            int    `json:"pciSlot"`
		PortType           string `json:"portType"`
		RemainLife         int    `json:"remain_life"`
		Serial             string `json:"serial"`
		SizeTotal          string `json:"size_total"`
		SmartProgress      string `json:"smart_progress"`
		SmartStatus        string `json:"smart_status"`
		SmartTestLimit     int    `json:"smart_test_limit"`
		Status             string `json:"status"`
		Support            bool   `json:"support"`
		Temp               int    `json:"temp"`
		TrayStatus         string `json:"tray_status"`
		Unc                int    `json:"unc"`
		UsedBy             string `json:"used_by"`
		Vendor             string `json:"vendor"`
	} `json:"disks"`
	Env struct {
		Batchtask struct {
			MaxTask    int `json:"max_task"`
			RemainTask int `json:"remain_task"`
		} `json:"batchtask"`
		BayNumber          string        `json:"bay_number"`
		Ebox               []interface{} `json:"ebox"`
		FsActing           bool          `json:"fs_acting"`
		IsSyncSysPartition bool          `json:"isSyncSysPartition"`
		IsSpaceActioning   bool          `json:"is_space_actioning"`
		Isns               struct {
			Address string `json:"address"`
			Enabled bool   `json:"enabled"`
		} `json:"isns"`
		IsnsServer            string `json:"isns_server"`
		MaxFsBytes            string `json:"max_fs_bytes"`
		MaxFsBytesHighEnd     string `json:"max_fs_bytes_high_end"`
		ModelName             string `json:"model_name"`
		RAMEnoughForFsHighEnd bool   `json:"ram_enough_for_fs_high_end"`
		RAMSize               int    `json:"ram_size"`
		RAMSizeRequired       int    `json:"ram_size_required"`
		Showpooltab           bool   `json:"showpooltab"`
		Status                struct {
			SystemCrashed    bool `json:"system_crashed"`
			SystemNeedRepair bool `json:"system_need_repair"`
		} `json:"status"`
		Support struct {
			Ebox      bool `json:"ebox"`
			RaidCross bool `json:"raid_cross"`
			Sysdef    bool `json:"sysdef"`
		} `json:"support"`
		SupportFitFsLimit bool   `json:"support_fit_fs_limit"`
		UniqueKey         string `json:"unique_key"`
	} `json:"env"`
	HotSpareConf struct {
		CrossRepair   bool          `json:"cross_repair"`
		DisableRepair []interface{} `json:"disable_repair"`
	} `json:"hotSpareConf"`
	HotSpares []interface{} `json:"hotSpares"`
	IscsiLuns []struct {
		CanDo struct {
			ConvertShrToPool int  `json:"convert_shr_to_pool"`
			Delete           bool `json:"delete"`
			ExpandByDisk     int  `json:"expand_by_disk"`
			Migrate          struct {
				ToShr2 int `json:"to_shr2"`
			} `json:"migrate"`
			RaidCross bool `json:"raid_cross"`
		} `json:"can_do"`
		ID          string `json:"id"`
		IsActioning bool   `json:"is_actioning"`
		IscsiLun    struct {
			BlkNum        string   `json:"blkNum"`
			DeviceType    string   `json:"device_type"`
			ExtentBased   bool     `json:"extent_based"`
			ExtentSize    string   `json:"extent_size"`
			Lid           int      `json:"lid"`
			Location      string   `json:"location"`
			MappedTargets []int    `json:"mapped_targets"`
			Name          string   `json:"name"`
			Parent        struct{} `json:"parent"`
			RestoredTime  string   `json:"restored_time"`
			Rootpath      string   `json:"rootpath"`
			ScheduledTask []struct {
				General struct {
					Lid         int    `json:"lid"`
					SnapRotate  bool   `json:"snap_rotate"`
					SnapType    string `json:"snap_type"`
					TaskEnabled bool   `json:"task_enabled"`
					TaskName    string `json:"task_name"`
					Tid         int    `json:"tid"`
				} `json:"general"`
				Schedule struct {
					Date                  string      `json:"date"`
					DateType              int         `json:"date_type"`
					Hour                  int         `json:"hour"`
					LastWorkHour          int         `json:"last_work_hour"`
					Min                   int         `json:"min"`
					NextTriggerTime       string      `json:"next_trigger_time"`
					Repeat                int         `json:"repeat"`
					RepeatHour            int         `json:"repeat_hour"`
					RepeatHourStoreConfig interface{} `json:"repeat_hour_store_config"`
					RepeatMin             int         `json:"repeat_min"`
					RepeatMinStoreConfig  interface{} `json:"repeat_min_store_config"`
					WeekName              string      `json:"week_name"`
				} `json:"schedule"`
			} `json:"scheduled_task"`
			Size          string        `json:"size"`
			Snapshots     []interface{} `json:"snapshots"`
			ThinProvision bool          `json:"thin_provision"`
			UsedBy        string        `json:"used_by"`
			UUID          string        `json:"uuid"`
		} `json:"iscsi_lun"`
		NumID    int `json:"num_id"`
		Progress struct {
			Percent string `json:"percent"`
			Step    string `json:"step"`
		} `json:"progress"`
		Status string `json:"status"`
	} `json:"iscsiLuns"`
	IscsiTargets []struct {
		Auth struct {
			MutualUsername string `json:"mutual_username"`
			Type           string `json:"type"`
			Username       string `json:"username"`
		} `json:"auth"`
		DataChksum              bool          `json:"data_chksum"`
		Enabled                 bool          `json:"enabled"`
		HdrChksum               bool          `json:"hdr_chksum"`
		Iqn                     string        `json:"iqn"`
		MappedLogicalUnitNumber []interface{} `json:"mapped_logical_unit_number"`
		MappedLuns              []interface{} `json:"mapped_luns"`
		Masking                 []struct {
			Iqn        string `json:"iqn"`
			Permission string `json:"permission"`
		} `json:"masking"`
		MultiSessions bool          `json:"multi_sessions"`
		Name          string        `json:"name"`
		NumID         int           `json:"num_id"`
		RecvSegBytes  int           `json:"recv_seg_bytes"`
		Remote        []interface{} `json:"remote"`
		SendSegBytes  int           `json:"send_seg_bytes"`
		Status        string        `json:"status"`
		Tid           int           `json:"tid"`
	} `json:"iscsiTargets"`
	Ports        []interface{} `json:"ports"`
	SsdCaches    []interface{} `json:"ssdCaches"`
	StoragePools []interface{} `json:"storagePools"`
	Volumes      []struct {
		AtimeChecked bool   `json:"atime_checked"`
		AtimeOpt     string `json:"atime_opt"`
		CacheStatus  string `json:"cacheStatus"`
		CanDo        struct {
			ConvertShrToPool int  `json:"convert_shr_to_pool"`
			Delete           bool `json:"delete"`
			ExpandByDisk     int  `json:"expand_by_disk"`
			Migrate          struct {
				ToShr2 int `json:"to_shr2"`
			} `json:"migrate"`
			RaidCross bool `json:"raid_cross"`
		} `json:"can_do"`
		Container         string   `json:"container"`
		DeployPath        string   `json:"deploy_path"`
		Desc              string   `json:"desc"`
		DeviceType        string   `json:"device_type"`
		DiskFailureNumber int      `json:"disk_failure_number"`
		Disks             []string `json:"disks"`
		DriveType         int      `json:"drive_type"`
		EppoolUsed        string   `json:"eppool_used"`
		ExistAliveVdsm    bool     `json:"exist_alive_vdsm"`
		FsType            string   `json:"fs_type"`
		ID                string   `json:"id"`
		IsActing          bool     `json:"is_acting"`
		IsActioning       bool     `json:"is_actioning"`
		IsInodeFull       bool     `json:"is_inode_full"`
		IsWritable        bool     `json:"is_writable"`
		LimitedDiskNumber int      `json:"limited_disk_number"`
		MaxFsSize         string   `json:"max_fs_size"`
		MaximalDiskSize   string   `json:"maximal_disk_size"`
		MinimalDiskSize   string   `json:"minimal_disk_size"`
		NumID             int      `json:"num_id"`
		PoolPath          string   `json:"pool_path"`
		Progress          struct {
			Percent string `json:"percent"`
			Step    string `json:"step"`
		} `json:"progress"`
		Raids []struct {
			DesignedDiskCount int `json:"designedDiskCount"`
			Devices           []struct {
				ID     string `json:"id"`
				Slot   int    `json:"slot"`
				Status string `json:"status"`
			} `json:"devices"`
			MinDevSize     string        `json:"minDevSize"`
			NormalDevCount int           `json:"normalDevCount"`
			RaidPath       string        `json:"raidPath"`
			RaidStatus     int           `json:"raidStatus"`
			Spares         []interface{} `json:"spares"`
		} `json:"raids"`
		Size struct {
			FreeInode   string `json:"free_inode"`
			Total       string `json:"total"`
			TotalDevice string `json:"total_device"`
			TotalInode  string `json:"total_inode"`
			Used        string `json:"used"`
		} `json:"size"`
		SpacePath string        `json:"space_path"`
		Spares    []interface{} `json:"spares"`
		SsdTrim   struct {
			Support string `json:"support"`
		} `json:"ssd_trim"`
		Status        string        `json:"status"`
		Suggestions   []interface{} `json:"suggestions"`
		Timebackup    bool          `json:"timebackup"`
		UsedByGluster bool          `json:"used_by_gluster"`
		VolPath       string        `json:"vol_path"`
		VspaceCanDo   struct {
			Drbd struct {
				Resize struct {
					CanDo       bool `json:"can_do"`
					ErrCode     int  `json:"errCode"`
					StopService bool `json:"stopService"`
				} `json:"resize"`
			} `json:"drbd"`
			Flashcache struct {
				Apply struct {
					CanDo       bool `json:"can_do"`
					ErrCode     int  `json:"errCode"`
					StopService bool `json:"stopService"`
				} `json:"apply"`
				Remove struct {
					CanDo       bool `json:"can_do"`
					ErrCode     int  `json:"errCode"`
					StopService bool `json:"stopService"`
				} `json:"remove"`
				Resize struct {
					CanDo       bool `json:"can_do"`
					ErrCode     int  `json:"errCode"`
					StopService bool `json:"stopService"`
				} `json:"resize"`
			} `json:"flashcache"`
			Snapshot struct {
				Resize struct {
					CanDo       bool `json:"can_do"`
					ErrCode     int  `json:"errCode"`
					StopService bool `json:"stopService"`
				} `json:"resize"`
			} `json:"snapshot"`
		} `json:"vspace_can_do"`
	} `json:"volumes"`
}

var sconn SynologyConnection

func performHTTPCall(method string, url string) (response *http.Response) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal("should return error", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("should return error", err)
	}

	return resp
}

// NewConnection initializes Synology Connection
func NewConnection(host string, port string, account string, password string) (conn *SynologyConnection) {
	sconn.origin = fmt.Sprintf("http://%s:%s", host, port)
	sconn.token = getSIDToken(account, password)

	return &sconn
}

// getSIDToken returns a Synology Auth Token, for the given account
func getSIDToken(account string, password string) string {
	authVersion := "6"
	url := fmt.Sprintf("%s/webapi/auth.cgi?api=SYNO.API.Auth&version=%s&method=login&account=%s&passwd=%s&session=Core&format=cookie", sconn.origin, authVersion, account, password)

	resp := performHTTPCall("GET", url)
	defer resp.Body.Close()

	var respData authResponse

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		log.Fatal("Should return err", err)
	}

	return respData.Data.Sid
}

// GetSystemInfo returns Synology system information
func GetSystemInfo(conn *SynologyConnection) *SystemUtilizationData {
	version := "1"
	url := fmt.Sprintf("%s/webapi/entry.cgi?api=SYNO.Core.System.Utilization&version=%s&method=get&_sid=%s", conn.origin, version, conn.token)

	resp := performHTTPCall("GET", url)
	defer resp.Body.Close()

	var respData systemUtilizationResponse

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		log.Fatal("should return err", err)
	}
	return &respData.Data
}

// GetShareInfo returns individual Shared Folder information
func GetShareInfo(conn *SynologyConnection) *SharedFolderData {
	version := "1"
	url := fmt.Sprintf("%s/webapi/entry.cgi?api=SYNO.Core.Share&shareType=all&additional=%%5B%%22share_quota%%22%%5D&method=list&version=%s&_sid=%s", conn.origin, version, conn.token)

	resp := performHTTPCall("GET", url)
	defer resp.Body.Close()

	var respData sharedFolderReponse

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		log.Fatal("should return an err", err)
	}

	return &respData.Data
}

// GetStorageInfo returns alot of information about storage attached to NAS
func GetStorageInfo(conn *SynologyConnection) *StorageData {
	version := "1"
	url := fmt.Sprintf("%s/webapi/entry.cgi?api=SYNO.Storage.CGI.Storage&version=%s&method=load_info&_sid=%s", conn.origin, version, conn.token)

	resp := performHTTPCall("GET", url)
	defer resp.Body.Close()

	var respData storageResponse

	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		log.Fatal("should return an err", err)
	}

	return &respData.Data
}
