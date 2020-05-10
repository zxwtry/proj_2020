package backend_comm_func

import (
	"os"

	"github.com/zxwtry/proj_2020/go/src/backend/backend_comm_constant"
)

func CheckZxwPcEnv() bool {
	if os.Getenv(backend_comm_constant.SYS_ENV_KEY_ZXWPC) == backend_comm_constant.SYS_ENV_VAL_ZXWPC {
		return true
	}
	return false
}
