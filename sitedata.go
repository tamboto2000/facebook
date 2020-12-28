package facebook

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/tamboto2000/random"
)

// SiteData is maybe some sort of current session data, required when requesting Facebook GraphQL API
type SiteData struct {
	ServerRevision     int     `json:"server_revision,omitempty"`
	ClientRevision     int     `json:"client_revision,omitempty"`
	Tier               string  `json:"tier,omitempty"`
	PushPhase          string  `json:"push_phase,omitempty"`
	PkgCohort          string  `json:"pkg_cohort,omitempty"`
	PR                 float64 `json:"pr,omitempty"`
	HasteSite          string  `json:"haste_site,omitempty"`
	BeOneAhead         bool    `json:"be_one_ahead,omitempty"`
	IROn               bool    `json:"ir_on,omitempty"`
	IsRTL              bool    `json:"is_rtl,omitempty"`
	IsComet            bool    `json:"is_comet,omitempty"`
	IsExperimentalTier bool    `json:"is_experimental_tier,omitempty"`
	IsJITWarmedUp      bool    `json:"is_jit_warmed_up,omitempty"`
	Hsi                string  `json:"hsi,omitempty"`
	SemrHostBucket     string  `json:"semr_host_bucket,omitempty"`
	Spin               int     `json:"spin,omitempty"`
	SpinR              int     `json:"__spin_r,omitempty"`
	SpinB              string  `json:"__spin_b,omitempty"`
	SpinT              int     `json:"__spin_t,omitempty"`
	CometEnv           string  `json:"comet_env,omitempty"`
	Vip                string  `json:"vip,omitempty"`
}

// prepare SiteData for GraphQL request
func (fb *Facebook) prepSiteData(std *SiteData) url.Values {
	final := make(url.Values)
	cUser, _ := fb.cookies.Load("c_user")
	userID := cUser.(*http.Cookie).Value
	final.Set("av", userID)
	final.Set("__user", userID)
	final.Set("__a", "1")
	final.Set("__req", random.RandStrWithOpt(2, random.Option{IncludeNumber: true}))

	if std.BeOneAhead {
		final.Set("__beoa", "1")
	} else {
		final.Set("__beoa", "0")
	}

	final.Set("__pc", std.PkgCohort)
	final.Set("dpr", fmt.Sprintf("%g", std.PR))
	final.Set("__ccg", "EXCELLENT")
	final.Set("__rev", strconv.Itoa(std.ServerRevision))
	final.Set("__hsi", std.Hsi)
	final.Set("__comet_req", "1")
	final.Set("__comet_env", "fb")
	final.Set("fb_dtsg", fb.FbDtsg)
	final.Set("jazoest", fb.Jazoest)
	final.Set("__spin_r", strconv.Itoa(std.SpinR))
	final.Set("__spin_b", std.SpinB)
	final.Set("__spin_t", strconv.Itoa(std.SpinT))
	final.Set("fb_api_caller_class", "RelayModern")
	final.Set("server_timestamps", "true")

	return final
}
