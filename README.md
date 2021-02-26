# Facebook

[![Go Reference](https://pkg.go.dev/badge/github.com/tamboto2000/facebook.svg)](https://pkg.go.dev/github.com/tamboto2000/facebook)

Facebook is a library for scraping Facebook profile data, such as username, full name, profile pict, cover pict, about info, and many more

## Current Features

- Scrape profile basic info
- Get user work and education history
- Get user places lived history
- Get user Contact info
- Get user life events history

## Upcoming Features

- Get user friend list
- Get user posts

## Example

```go
package main

import (
	"encoding/json"
	"os"

	"github.com/tamboto2000/facebook"
)

func main() {
	// get new client
	fb := facebook.New()

	// set facebook login cookie
	fb.SetCookieStr(`your_facebook_cookie`)

	// initiate client
	if err := fb.Init(); err != nil {
		panic(err.Error())
	}

	username := "franklin.tamboto.3"

	// get profile
	profile, err := fb.Profile(username)
	if err != nil {
		panic(err.Error())
	}

	// before getting all data from "About" section, make sure to call Profile.SyncAbout first
	if err := profile.SyncAbout(); err != nil {
		panic(err.Error())
	}

	if err := profile.SyncWorkAndEducation(); err != nil {
		panic(err.Error())
	}

	if err := profile.SyncPlacesLived(); err != nil {
		panic(err.Error())
	}

	if err := profile.SyncContactAndBasicInfo(); err != nil {
		panic(err.Error())
	}

	if err := profile.SyncFamilyAndRelationships(); err != nil {
		panic(err.Error())
	}

	if err := profile.SyncDetails(); err != nil {
		panic(err.Error())
	}

	if err := profile.SyncLifeEvents(); err != nil {
		panic(err.Error())
	}

	// save profile to a file
	f, err := os.Create(username + ".json")
	if err != nil {
		panic(err.Error())
	}

	defer f.Close()

	if err := json.NewEncoder(f).Encode(profile); err != nil {
		panic(err.Error())
	}
}
```

Currently, all methods mentioned above, can not run concurrently

See [documentation](https://pkg.go.dev/github.com/tamboto2000/facebook) for more info

# License

MIT