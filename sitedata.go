package facebook

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tamboto2000/random"
)

// SiteData is maybe some sort of current session data, required when requesting Facebook GraphQL API
type SiteData struct {
	ServerRevision     int     `json:"server_revision"`
	ClientRevision     int     `json:"client_revision"`
	Tier               string  `json:"tier"`
	PushPhase          string  `json:"push_phase"`
	PkgCohort          string  `json:"pkg_cohort"`
	PR                 float64 `json:"pr"`
	HasteSite          string  `json:"haste_site"`
	BeOneAhead         bool    `json:"be_one_ahead"`
	IROn               bool    `json:"ir_on"`
	IsRTL              bool    `json:"is_rtl"`
	IsComet            bool    `json:"is_comet"`
	IsExperimentalTier bool    `json:"is_experimental_tier"`
	IsJITWarmedUp      bool    `json:"is_jit_warmed_up"`
	Hsi                string  `json:"hsi"`
	SemrHostBucket     string  `json:"semr_host_bucket"`
	Spin               int     `json:"spin"`
	SpinR              int     `json:"__spin_r"`
	SpinB              string  `json:"__spin_b"`
	SpinT              int     `json:"__spin_t"`
	CometEnv           string  `json:"comet_env"`
	Vip                string  `json:"vip"`
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
