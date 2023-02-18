//go:build !custom || fsNotify

package register

import (
	_ "hippo-data-acquisition/inputs/plugins/fs_notify"
)
