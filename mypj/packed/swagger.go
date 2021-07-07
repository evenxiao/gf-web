package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GAo1r0bxIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwMj1+ZbCaueBQUEBARo6fn7+PvrBm0KCjC6ZBxwqbEj7dnSp9qRmg5XGlwnff7sJTTZ26Fpa8LFqBWXmo5M6knqKNBVYGD4/z/Am53DrHPCQgcGBoYQBgYGmFMYMJzCjnAK3HaQbmQ1Ad6MTCLMCK8gmwzyCgxsawSReD2GMAq7UyBAgOG/YzfCKCTdrGwgeSYGJoZOBgaGk2DVgAAAAP//Fvx9SWgBAAA="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
