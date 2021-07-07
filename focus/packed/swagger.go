package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
