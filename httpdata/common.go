package httpdata

/*
{
   "touser" : "UserID1|UserID2|UserID3",
   "toparty" : "PartyID1|PartyID2",
   "totag" : "TagID1 | TagID2",
   "msgtype" : "image",
   "agentid" : 1,
   "image" : {
        "media_id" : "MEDIA_ID"
   },
   "safe":0,
   "enable_duplicate_check": 0,
   "duplicate_check_interval": 1800
}
*/

type commonData struct {
	ToUser                 string `json:"touser"`
	ToParty                string `json:"toparty"`
	ToTag                  string `json:"totag"`
	MsgType                string `json:"msgtype"`
	AgentId                int    `json:"agentid"`
	Safe                   int    `json:"safe"`
	EnableDuplicateCheck   int    `json:"enable_duplicate_check"`
	DuplicateCheckInterval int    `json:"cuplicate_check_interval"`
}
