package config

// S3Config is s3 storage configuration.
type S3Config struct {
	Endpoint        string `envconfig:"ENDPOINT"          required:"true"`
	AccessKeyID     string `envconfig:"ACCESS_KEY_ID"     required:"true"`
	SecretAccessKey string `envconfig:"SECRET_ACCESS_KEY" required:"true"`
	UseSSL          bool   `envconfig:"USE_SSL"           required:"true"`
	BucketName      string `envconfig:"BUCKET_NAME"       required:"true"`
	Region          string `envconfig:"REGION"            required:"true"`
}
