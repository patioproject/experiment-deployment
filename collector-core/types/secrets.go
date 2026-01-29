package types

type Secrets struct {
	User     string
	Password string
	Token    string
	R2       R2Config
}

type R2Config struct {
	BucketName      string
	AccountId       string
	AccessKeyId     string
	AccessKeySecret string
}
