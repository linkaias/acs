package model

type Env struct {
	// 是否调试模式
	Debug bool `mapstructure:"DEBUG" json:"debug" form:"debug"`

	// 数据库类型
	DBType string `mapstructure:"DB_TYPE" json:"db_type" form:"db_type"`

	// MySQL连接配置
	MysqlHost     string `mapstructure:"MYSQL_HOST" json:"mysql_host" form:"mysql_host"`
	MysqlPort     uint   `mapstructure:"MYSQL_PORT" json:"mysql_port" form:"mysql_port"`
	MysqlDatabase string `mapstructure:"MYSQL_DATABASE" json:"mysql_database" form:"mysql_database"`
	MysqlUsername string `mapstructure:"MYSQL_USERNAME" json:"mysql_username" form:"mysql_username"`
	MysqlPassword string `mapstructure:"MYSQL_PASSWORD" json:"mysql_password" form:"mysql_password"`

	// ClickHouse连接配置
	ChHost     string `mapstructure:"CH_HOST" json:"ch_host" form:"ch_host"`
	ChPort     uint   `mapstructure:"CH_PORT" json:"ch_port" form:"ch_port"`
	ChDatabase string `mapstructure:"CH_DATABASE" json:"ch_database" form:"ch_database"`
	ChUsername string `mapstructure:"CH_USERNAME" json:"ch_username" form:"ch_username"`
	ChPassword string `mapstructure:"CH_PASSWORD" json:"ch_password" form:"ch_password"`

	// MongoDB连接配置
	MongoHost     string `mapstructure:"MONGO_HOST" json:"mongo_host" form:"mongo_host"`
	MongoPort     uint   `mapstructure:"MONGO_PORT" json:"mongo_port" form:"mongo_port"`
	MongoDatabase string `mapstructure:"MONGO_DATABASE" json:"mongo_database" form:"mongo_database"`
	MongoUsername string `mapstructure:"MONGO_USERNAME" json:"mongo_username" form:"mongo_username"`
	MongoPassword string `mapstructure:"MONGO_PASSWORD" json:"mongo_password" form:"mongo_password"`

	// 单次批量插入数据个数
	DBBatchInsertNum int `mapstructure:"DB_BATCH_INSERT_NUM" json:"db_batch_insert_num" form:"db_batch_insert_num"`
	// 批量自动插入时间间隔
	DBBatchInsertGap int `mapstructure:"DB_BATCH_INSERT_GAP" json:"db_batch_insert_gap" form:"db_batch_insert_gap"`

	// grpc 运行端口
	GRPCServePort string `mapstructure:"GRPC_SERVE_PORT" json:"grpc_serve_port" form:"grpc_serve_port"`

	// 服务端口
	HttpServePort string `mapstructure:"HTTP_SERVE_PORT" json:"http_serve_port" form:"http_serve_port"`
	// 后台用户登陆信息
	AdminUser     string `mapstructure:"ADMIN_USER" json:"admin_user" form:"admin_user"`
	AdminPassword string `mapstructure:"ADMIN_PASSWORD" json:"admin_password" form:"admin_password"`

	// 登陆密钥
	SecretKey string `mapstructure:"SECRET_KEY" json:"secret_key" form:"secret_key"`

	// http 登陆信息
	HTTPUser     string `mapstructure:"HTTP_USER" json:"http_user" form:"http_user"`
	HTTPPassword string `mapstructure:"HTTP_PASSWORD" json:"http_password" form:"http_password"`

	// grpc 登陆信息
	GRPCUser     string `mapstructure:"GRPC_USER" json:"grpc_user" form:"grpc_user"`
	GRPCPassword string `mapstructure:"GRPC_PASSWORD" json:"grpc_password" form:"grpc_password"`

	// jwt token 过期时间
	JWTExpireTime int `mapstructure:"JWT_EXPIRE_TIME" json:"jwt_expire_time" form:"jwt_expire_time"`

	// jwt token admin 过期时间
	JwtAdminExpireTime int `mapstructure:"JWT_ADMIN_EXPIRE_TIME" json:"jwt_admin_expire_time" form:"jwt_admin_expire_time"`

	// 记录保存天数
	RecordSaveDays int `mapstructure:"RECORD_SAVE_DAYS" json:"record_save_days" form:"record_save_days"`

	// 日志保存路径
	LogFilePath string `mapstructure:"LOG_FILE_PATH" json:"log_file_path" form:"log_file_path"`

	// 日志文件最大保存天数
	LogFileMaxAge int `mapstructure:"LOG_FILE_MAX_AGE" json:"log_file_max_age" form:"log_file_max_age"`

	// 日志文件切割时间间隔(小时)
	LogFileGap int `mapstructure:"LOG_FILE_GAP" json:"log_file_gap" form:"log_file_gap"`

	// 开启后台访问记录
	OpenAdminRecord bool `mapstructure:"OPEN_ADMIN_RECORD" json:"open_admin_record" form:"open_admin_record"`
}
