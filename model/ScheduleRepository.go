package model

import (
	"time"
)

type Filter struct {
	Boolean string
	Path    string
	Op      string
	Value   string
}

type PendingTask struct {
	TableName    struct{} `sql:"talkbank_bots.pending_task"`
	Id           int `sql:"id"`
	ActionId     int
	UserId       int
	Planned      time.Time
	Delivery     *Delivery
	ScheduleTask *ScheduleTask
}

// DB Model for table schedule_task
type ScheduleTask struct {
	TableName    struct{} `sql:"talkbank_bots.schedule_task"`
	Id           int `sql:"id"`
	ActionId     int
	Type         string `sql:"type:talkbank_bots.SCHEDULE_TASK_TYPE"`
	Category     string
	Template     string
	FromDatetime time.Time `sql:",notnull"`
	ToDatetime   time.Time `sql:",notnull"`
	IsActive     bool
	CreatedAt    time.Time `sql:"default:now()"`
	UpdatedAt    time.Time
	NextDatetime time.Time
	LastRun      time.Time
	NextRun      time.Time `sql:",notnull"`
	StartTz      string
	Delivery     *Delivery
	IsRunning    bool
}

// DB Model for table delivery
type Delivery struct {
	TableName           struct{} `sql:"talkbank_bots.delivery"`
	Id                  int `sql:"id"`
	Text                string `sql:",notnull"`
	Title               string `sql:",notnull"`
	Description         string `sql:",notnull"`
	TagId               int `sql:",notnull"`
	CountUsers          int `sql:"countUsers"`
	Sent                int
	LastSending         time.Time `sql:"lastSending"`
	CategoryId          int
	TimeCondition       string `sql:"timeCondition"`
	ParametersCondition string `sql:"parametersCondition"`
	CreatedAt           time.Time `sql:"default:now()"`
	UpdatedAt           time.Time
	UserIds             string
	Filter              []Filter
	ActionHash          string
}

type UserDelivery struct {
	TableName    struct{} `sql:"talkbank_bots.user_delivery"`
	Id           int `sql:"id"`
	UserId       int
	DeliveryId   int
	Status       string
	CreatedAt    time.Time `sql:"default:now()"`
	UpdatedAt    time.Time
	DeliveryHash string
	Delivery     *Delivery
}

type ScheduleRepository struct {
	task    []ScheduleTask
	deliver []Delivery
	pending []PendingTask
}

func NewScheduleRepository() *ScheduleRepository {
	return &ScheduleRepository{}
}

func (scheduleModel *ScheduleRepository) GetTaskModel() []ScheduleTask {
	return scheduleModel.task
}

func (scheduleModel *ScheduleRepository) GetPendingTaskModel() []PendingTask {
	return scheduleModel.pending
}