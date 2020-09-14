package _struct

type DailyDesc struct {
	CurrentConfirmedCount int `json:"current_confirmed_count" db:"currentConfirmedCount"`
	ConfirmedCount        int `json:"confirmed_count" db:"confirmedCount"`
	SuspectedCount        int `json:"suspected_count" db:"suspectedCount"`
	CuredCount            int `json:"cured_count" db:"curedCount"`
	DeadCount             int `json:"dead_count" db:"deadCount"`
	SeriousCount          int `json:"serious_count" db:"seriousCount"`
	SuspectedIncr         int `json:"suspected_incr" db:"suspectedIncr"`
	CurrentConfirmedIncr  int `json:"current_confirmed_incr" db:"currentConfirmedIncr"`
	ConfirmedIncr         int `json:"confirmed_incr" db:"confirmedIncr"`
	CuredIncr             int `json:"cured_incr" db:"curedIncr"`
	DeadIncr              int `json:"dead_incr" db:"deadIncr"`
	SeriousIncr           int `json:"serious_incr" db:"seriousIncr"`
}

type DailyNews struct {
	PubDate    int    `json:"pub_date"`
	PubDateStr string `json:"pub_date_str"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	InfoSource string `json:"info_source"`
	SourceURL  string `json:"source_url"`
}

type NcovCity struct {
	Name                  string `json:"name"`
	CurrentConfirmedCount int    `json:"current_confirmed_count" db:"currentConfirmedCount"`
	ConfirmedCount        int    `json:"confirmed_count" db:"confirmedCount"`
	SuspectedCount        int    `json:"suspected_count" db:"suspectedCount"`
	CuredCount            int    `json:"cured_count" db:"curedCount"`
	DeadCount             int    `json:"dead_count" db:"deadCount"`
}

type NcovDistrict struct {
	Locale         string       `json:"locale"`
	Address        string       `json:"address"`
	Lng            string       `json:"lng"`
	Lat            string       `json:"lat"`
	Source         string       `json:"source"`
	Region         int          `json:"region"`
	NcovPublicList []NcovPublic `json:"ncov_public_list"`
}

type NcovPublic struct {
	PersonName string `json:"person_name"`
	Address    string `json:"address"`
	Start      string `json:"start"`
	End        string `json:"end"`
	Date       string `json:"date"`
}

type NcovRide struct {
	PassengerName string `json:"person_name"`
	Date          string `json:"date"`
	Start         string `json:"start"`
	End           string `json:"end"`
	TrainNum      string `json:"train_num"`
	CarriageNum   string `json:"carriage_num"`
	SeatNum       string `json:"seat_num"`
	PosStart      string `json:"pos_start"`
	PosEnd        string `json:"pos_end"`
}

type NcovDistricts struct {
	Ncovs []NcovDistrict `json:"ncovs"`
}

type NcovCom struct {
	DistrictList []NcovDistrict `json:"district_list"`
	RideList     []NcovRide     `json:"ride_list"`
}

type Push struct {
	UserAccount string `json:"user_account"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
