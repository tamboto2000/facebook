package facebook

import (
	"fmt"
	"net/http"
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
func (fb *Facebook) prepSiteData(std *SiteData) map[string]string {
	final := make(map[string]string)
	cUser, _ := fb.cookies.Load("c_user")
	userID := cUser.(*http.Cookie).Value
	final["av"] = userID
	final["__user"] = userID
	final["__a"] = "1"
	final["__req"] = random.RandStrWithOpt(2, random.Option{IncludeNumber: true})

	if std.BeOneAhead {
		final["__beoa"] = "1"
	} else {
		final["__beoa"] = "0"
	}

	final["__pc"] = std.PkgCohort
	final["dpr"] = fmt.Sprintf("%g", std.PR)
	final["__ccg"] = "GOOD"
	final["__rev"] = strconv.Itoa(std.ServerRevision)
	// final["__s"]
	final["__comet_req"] = "1"
	final["__comet_env"] = "fb"
	// final["fb_dtsg"]
	// final["jazoest"]
	final["__spin_r"] = strconv.Itoa(std.SpinR)
	final["__spin_b"] = std.SpinB
	final["__spin_t"] = strconv.Itoa(std.SpinT)
	final["fb_api_caller_class"] = "RelayModern"
	final["server_timestamps"] = "true"

	// fb_api_req_friendly_name
	// variables
	// doc_id

	return final
}
