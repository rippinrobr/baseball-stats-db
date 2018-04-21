package retrosheet

type umpires struct {
	HomeUmpID         string `json:"homeUmpId" bson:"homeUmpId" db:"home_ump_id"`
	HomeUmpName       string `json:"homeUmpName" bson:"homeUmpName" db:"home_ump_name"`
	FirstBaseUmpID    string `json:"firstBaseUmpId" bson:"firstBaseUmpId" db:"first_base_ump_id"`
	FirstBaseUmpName  string `json:"firstBaseUmpName" bson:"firstBaseUmpName" db:"first_base_ump_name"`
	SecondBaseUmpID   string `json:"secondBaseUmpId" bson:"secondBaseUmpId" db:"second_base_ump_id"`
	SecondBaseUmpName string `json:"secondBaseUmpName" bson:"secondBaseUmpName" db:"second_base_ump_name"`
	ThirdBaseUmpID    string `json:"thirdBaseUmpId" bson:"thirdBaseUmpId" db:"third_base_ump_id"`
	ThirdBaseUmpName  string `json:"thirdBaseUmpName" bson:"thirdBaseUmpName" db:"third_base_ump_name"`
	LFUmpID           string `json:"lfUmpId" bson:"lfUmpId" db:"lf_ump_id"`
	LFUmpName         string `json:"lfUmpName" bson:"lfUmpName" db:"lf_ump_name"`
	RFUmpID           string `json:"rfUmpId" bson:"rfUmpId" db:"rf_ump_id"`
	RFUmpName         string `json:"rfUmpName" bson:"rfUmpName" db:"rf_ump_name"`
}
