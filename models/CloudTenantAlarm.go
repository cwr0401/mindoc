package models

import (
	"time"

	"github.com/mindoc-org/mindoc/conf"
)

type CloudTenantBasicAlarm struct {
	Id                 int                 `orm:"pk;auto;unique;column(id)" json:"id"`
	CloudTenantAccount *CloudTenantAccount `orm:"rel(fk)"`

	// tencentcloud alarm
	/*
			{
		                "Id":2309258,
		                "ProjectId":0,
		                "ProjectName":"默认项目",
		                "Status":1,
		                "AlarmStatus":"OK",
		                "GroupId":1841426,
		                "GroupName":"N-CU-DL-基础监控-次告警",
		                "FirstOccurTime":"2022-11-02T23:06:00+08:00",
		                "Duration":16620,
		                "LastOccurTime":"2022-11-03T03:43:00+08:00",
		                "Content":"磁盘利用率 \u003e=  70 %",
		                "ObjName":"10.205.105.135(武汉事业群-数字图书馆-AP5)",
		                "ObjId":"25931026-9ead-4759-818c-b368daffb775",
		                "ViewName":"cvm_device",
		                "Vpc":"1",
		                "MetricId":20,
		                "MetricName":"disk_usage",
		                "AlarmType":0,
		                "Region":"wh",
		                "Dimensions":"{\"unInstanceId\":\"ins-3zbmw2wt\"}",
		                "NotifyWay":[
		                    "SMS"
		                ],
		                "InstanceGroup":[
		                    {
		                        "InstanceGroupId":2670,
		                        "InstanceGroupName":"武汉-数图N-CU-DL-生产CVM"
		                    }
		                ]
		    }
	*/
	// 该条告警的ID
	AlarmId uint64 `json:"AlarmId,omitempty" orm:"column(alarm_id);unique"`

	// 项目ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectId int64 `json:"ProjectId,omitempty" orm:"column(project_id);null"`

	// 项目名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProjectName string `json:"ProjectName,omitempty" orm:"column(project_name);null"`

	// 告警状态ID，0表示未恢复；1表示已恢复；2,3,5表示数据不足；4表示已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status int64 `json:"Status,omitempty" orm:"column(status);null"`

	// 告警状态，ALARM表示未恢复；OK表示已恢复；NO_DATA表示数据不足；NO_CONF表示已失效
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmStatus string `json:"AlarmStatus,omitempty" orm:"column(alarm_status);null"`

	// 策略组ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupId int64 `json:"GroupId,omitempty" orm:"column(group_id);null"`

	// 策略组名
	// 注意：此字段可能返回 null，表示取不到有效值。
	GroupName string `json:"GroupName,omitempty" orm:"column(group_name);null"`

	// 发生时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	FirstOccurTime string `json:"FirstOccurTime,omitempty" orm:"column(first_occur_Time);null"`

	// 持续时间，单位s
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration int64 `json:"Duration,omitempty" orm:"column(duration);null"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastOccurTime string `json:"LastOccurTime,omitempty" orm:"column(last_occur_time);null"`

	// 告警内容
	// 注意：此字段可能返回 null，表示取不到有效值。
	Content string `json:"Content,omitempty" orm:"column(content);null"`

	// 告警对象
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjName string `json:"ObjName,omitempty" orm:"column(obj_name);null"`

	// 告警对象ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjId string `json:"ObjId,omitempty" orm:"column(obj_id);null"`

	// 策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	ViewName string `json:"ViewName,omitempty" orm:"column(view_name);null"`

	// VPC，只有CVM有
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vpc string `json:"Vpc,omitempty" orm:"column(vpc);null"`

	// 指标ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricId int64 `json:"MetricId,omitempty" orm:"column(metric_id);null"`

	// 指标名
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName string `json:"MetricName,omitempty" orm:"column(metric_name);null"`

	// 告警类型，0表示指标告警，2表示产品事件告警，3表示平台事件告警
	// 注意：此字段可能返回 null，表示取不到有效值。
	AlarmType int64 `json:"AlarmType,omitempty" orm:"column(alarm_type);null"`

	// 地域
	// 注意：此字段可能返回 null，表示取不到有效值。
	Region string `json:"Region,omitempty" orm:"column(region);null"`

	// 告警对象维度信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Dimensions string `json:"Dimensions,omitempty" orm:"column(dimensions);null"`

	// 通知方式
	// 注意：此字段可能返回 null，表示取不到有效值。
	NotifyWay string `json:"NotifyWay,omitempty" orm:"column(notify_way);null"`

	// 所属实例组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	// InstanceGroup []*InstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`

	// CreateTime 创建时间 .
	CreateTime time.Time `orm:"type(datetime);column(create_time);auto_now_add" json:"CreateTime"`

	// 更新时间
	ModifyTime time.Time `orm:"type(datetime);column(modify_time);null;auto_now" json:"ModifyTime"`
}

// TableName 获取对应数据库表名.
func (account *CloudTenantBasicAlarm) TableName() string {
	return "cloud_tenant_basic_alarm"
}

func (account *CloudTenantBasicAlarm) TableEngine() string {
	return "INNODB"
}

func (account *CloudTenantBasicAlarm) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + account.TableName()
}

func NewCloudTenantBasicAlarm() *CloudTenantBasicAlarm {
	return &CloudTenantBasicAlarm{}
}
