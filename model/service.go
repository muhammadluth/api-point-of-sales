package model

type (
	Properties struct {
		DBHost             string `json:"db_host"`
		DBPort             string `json:"db_port"`
		DBName             string `json:"db_name"`
		DBUser             string `json:"db_user"`
		DBPassword         string `json:"db_password"`
		LogPath            string `json:"log_path"`
		PoolSize           int    `json:"pool_size"`
		PrivateKey         string `json:"private_key"`
		PublicKey          string `json:"public_key"`
		ExpireAccessToken  string `json:"expire_access_token"`
		ExpireRefreshToken string `json:"expire_refresh_token"`
	}
)
