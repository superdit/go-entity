package goentity

func HelloEntity() string {
  return "HelloEntity 111"
}

type User struct {
	Id                 string `json:"id" bson:"_id"`
	Fullname           string `json:"fullname" bson:"fullname"`
	Email              string `json:"email" bson:"email"`
	Phone              string `json:"phone" bson:"phone"`
	Password           string `json:"password" bson:"password"`
	Timezone           string `json:"timezone" bson:"timezone"`
	Verified           bool   `json:"verified" bson:"verified"`
	ImgUrl             string `json:"imgUrl" bson:"imgUrl"`
	FreePlan           bool   `json:"freePlan" bson:"freePlan"`
	SubscriptionPlanId string `json:"subscriptionPlanId" bson:"subscriptionPlanId"`
	CreatedAt          int64  `json:"createdAt" bson:"createdAt"`
	CreatedDate        string `json:"createdDate" bson:"createdDate"`
	ModifiedAt         int64  `json:"modifiedAt" bson:"modifiedAt"`
	ModifiedDate       string `json:"modifiedDate" bson:"modifiedDate"`
}
