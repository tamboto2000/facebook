package facebook

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/tamboto2000/jsonextract/v3"
	"github.com/tamboto2000/random"
)

// prepare SiteData for GraphQL request
func (fb *Facebook) prepSiteData(std *jsonextract.JSON) url.Values {
	final := make(url.Values)
	cUser := fb.cookies.getByName("c_user")
	userID := cUser.Value
	final.Set("av", userID)
	final.Set("__user", userID)
	final.Set("__a", "1")
	final.Set("__req", random.RandStrWithOpt(2, random.Option{IncludeNumber: true}))

	stdObj := std.Object()
	if stdObj["be_one_ahead"].Boolean() {
		final.Set("__beoa", "1")
	} else {
		final.Set("__beoa", "0")
	}

	final.Set("__pc", stdObj["pkg_cohort"].String())
	final.Set("dpr", fmt.Sprintf("%g", stdObj["pr"].Float()))
	final.Set("__ccg", "EXCELLENT")
	final.Set("__rev", strconv.Itoa(int(stdObj["server_revision"].Integer())))
	final.Set("__hsi", stdObj["hsi"].String())
	final.Set("__comet_req", "1")
	final.Set("__comet_env", "fb")
	final.Set("fb_dtsg", fb.FbDtsg)
	final.Set("jazoest", fb.Jazoest)
	final.Set("__spin_r", strconv.Itoa(int(stdObj["__spin_r"].Integer())))
	final.Set("__spin_b", stdObj["__spin_b"].String())
	final.Set("__spin_t", strconv.Itoa(int(stdObj["__spin_t"].Integer())))
	final.Set("fb_api_caller_class", "RelayModern")
	final.Set("server_timestamps", "true")

	return final
}
