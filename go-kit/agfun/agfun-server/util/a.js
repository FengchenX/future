
type Task struct {
    Name         string    `json:"name"`
    DaoId        uint8     `json:"dao_id"`
    AppId        string    `json:"app_id"`
    StartTime    time.Time `json:"start_time"`
    EndTime      time.Time `json:"end_time"`
    AccountsNo   []string  `json:"user_accounts"`
    TaskType     int64     `json:"task_type"`
    TargetAmount float64   `json:"target_amount"`
    RewardAmount float64   `json:"reward_amount"`
    Detail       string    `json:"detail"`
}
