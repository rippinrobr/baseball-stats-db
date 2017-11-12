package models


type TeamsFranchises struct {
   Franchid   string `json:"franchid"  csv:"franchID"  db:"franchID"`
   Franchname   string `json:"franchname"  csv:"franchName"  db:"franchName"`
   Active   string `json:"active"  csv:"active"  db:"active"`
   Naassoc   string `json:"naassoc"  csv:"NAassoc"  db:"NAassoc"`
}

func (m *TeamsFranchises) GetTableName() string {
  return "TeamsFranchises"
}

func (m *TeamsFranchises) GetFileName() string {
  return "TeamsFranchises.csv"
}

func (m *TeamsFranchises) GetFilePath() string {
  return "/Users/robertrowe/src/baseballdatabank/core/TeamsFranchises.csv"
}
