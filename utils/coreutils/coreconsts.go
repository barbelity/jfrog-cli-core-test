package coreutils

const (

	// Common
	TokenRefreshDisabled        = 0
	TokenRefreshDefaultInterval = 60

	// Home Dir
	JfrogCertsDirName        = "certs"
	JfrogConfigFile          = "jfrog-cli.conf"
	JfrogDependenciesDirName = "dependencies"
	JfrogSecurityDirName     = "security"
	JfrogSecurityConfFile    = "security.yaml"
	JfrogBackupDirName       = "backup"
	JfrogLogsDirName         = "logs"
	JfrogLockDirName         = "lock"

	// Env
	LogLevel        = "JFROG_CLI_LOG_LEVEL"
	ReportUsage     = "JFROG_CLI_REPORT_USAGE"
	HomeDir         = "JFROG_CLI_HOME_DIR"
	DependenciesDir = "JFROG_CLI_DEPENDENCIES_DIR"
	BuildName       = "JFROG_CLI_BUILD_NAME"
	BuildNumber     = "JFROG_CLI_BUILD_NUMBER"
	CI              = "CI"
	// Deprecated:
	JfrogHomeEnv = "JFROG_CLI_HOME"
)
