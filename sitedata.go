package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/tamboto2000/jsonextract/v2"
	"github.com/tamboto2000/random"
)

// prepare SiteData for GraphQL request
func (fb *Facebook) prepSiteData(std *jsonextract.JSON) url.Values {
	final := make(url.Values)
	cUser, _ := fb.cookies.Load("c_user")
	userID := cUser.(*http.Cookie).Value
	final.Set("av", userID)
	final.Set("__user", userID)
	final.Set("__a", "1")
	final.Set("__req", random.RandStrWithOpt(2, random.Option{IncludeNumber: true}))

	if std.KeyVal["be_one_ahead"].Val.(bool) {
		final.Set("__beoa", "1")
	} else {
		final.Set("__beoa", "0")
	}

	final.Set("__pc", std.KeyVal["pkg_cohort"].Val.(string))
	final.Set("dpr", fmt.Sprintf("%g", std.KeyVal["pr"].Val.(float64)))
	final.Set("__ccg", "EXCELLENT")
	final.Set("__rev", strconv.Itoa(std.KeyVal["server_revision"].Val.(int)))
	final.Set("__hsi", std.KeyVal["hsi"].Val.(string))
	final.Set("__comet_req", "1")
	final.Set("__comet_env", "fb")
	final.Set("fb_dtsg", fb.FbDtsg)
	final.Set("jazoest", fb.Jazoest)
	final.Set("__spin_r", strconv.Itoa(std.KeyVal["__spin_r"].Val.(int)))
	final.Set("__spin_b", std.KeyVal["__spin_b"].Val.(string))
	final.Set("__spin_t", strconv.Itoa(std.KeyVal["__spin_t"].Val.(int)))
	final.Set("fb_api_caller_class", "RelayModern")
	final.Set("server_timestamps", "true")

	return final
}
