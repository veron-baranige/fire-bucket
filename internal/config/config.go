package config

import "github.com/spf13/viper"

type (
	envConfig struct {
		ServerPort string `mapstructure:"SERVER_PORT"`

		DbDriver   string `mapstructure:"DB_DRIVER"`
		DbHost     string `mapstructure:"DB_HOST"`
		DbPort     string `mapstructure:"DB_PORT"`
		DbName     string `mapstructure:"DB_NAME"`
		DbUser     string `mapstructure:"DB_USER"`
		DbPassword string `mapstructure:"DB_PASSWORD"`

		FirebaseBucket string `mapstructure:"FIREBASE_STORAGE_BUCKET"`
	}

	firebaseCredentials struct {
		AccountType         string `mapstructure:"FIREBASE_ACCOUNT_TYPE" json:"type"`
		ProjectId           string `mapstructure:"FIREBASE_PROJECT_ID" json:"project_id"`
		PrivateKeyId        string `mapstructure:"FIREBASE_PRIVATE_KEY_ID" json:"private_key_id"`
		PrivateKey          string `mapstructure:"FIREBASE_PRIVATE_KEY" json:"private_key"`
		ClientEmail         string `mapstructure:"FIREBASE_CLIENT_EMAIL" json:"client_email"`
		ClientId            string `mapstructure:"FIREBASE_CLIENT_ID" json:"client_id"`
		AuthUri             string `mapstructure:"FIREBASE_AUTH_URI" json:"auth_uri"`
		TokenUri            string `mapstructure:"FIREBASE_TOKEN_URI" json:"token_uri"`
		AuthProviderCertUrl string `mapstructure:"FIREBASE_AUTH_PROVIDER_CERT_URL" json:"auth_provider_x509_cert_url"`
		ClientCertUrl       string `mapstructure:"FIREBASE_CLIENT_CERT_URL" json:"client_x509_cert_url"`
		UniverseDomain      string `mapstructure:"FIREBASE_UNIVERSE_DOMAIN" json:"universe_domain"`
	}

	ConfigKey string
)

const (
	ServerPort ConfigKey = "SERVER_PORT"

	DbDriver   ConfigKey = "DB_DRIVER"
	DbHost     ConfigKey = "DB_HOST"
	DbPort     ConfigKey = "DB_PORT"
	DbName     ConfigKey = "DB_NAME"
	DbUser     ConfigKey = "DB_USER"
	DbPassword ConfigKey = "DB_PASSWORD"

	FirebaseBucket ConfigKey = "FIREBASE_STORAGE_BUCKET"
	FileUploadDir  ConfigKey = "FILE_UPLOAD_DIR"
)

const (
	fileUploadDir = "uploads"
)

var (
	firebaseConf firebaseCredentials
)

func LoadEnv(configPath string) error {
	viper.SetConfigType("env")
	viper.AddConfigPath(configPath)
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	var env envConfig
	if err := viper.Unmarshal(&env); err != nil {
		return err
	}

	if err := viper.Unmarshal(&firebaseConf); err != nil {
		return err
	}

	// internal configs
	viper.Set(string(FileUploadDir), fileUploadDir)

	return nil
}

func Get(key ConfigKey) string {
	return viper.GetString(string(key))
}

func GetFirebaseCredentials() firebaseCredentials {
	return firebaseConf
}
