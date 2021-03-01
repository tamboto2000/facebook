package facebook

import (
	"net/http"
	"sync"
)

type cookies struct {
	mutex   *sync.Mutex
	cookies []*http.Cookie
}

func newCookies(cs []*http.Cookie) *cookies {
	return &cookies{
		mutex:   new(sync.Mutex),
		cookies: cs,
	}
}

func (cks *cookies) merge(newCs []*http.Cookie) {
	cks.mutex.Lock()
	for _, newC := range newCs {
		found := false
		for i, c := range cks.cookies {
			if newC.Name == c.Name {
				cks.cookies[i] = newC
				found = true

				break
			}
		}

		if !found {
			cks.cookies = append(cks.cookies, newC)
		}
	}

	cks.mutex.Unlock()
}

func (cks *cookies) getAll() []*http.Cookie {
	cks.mutex.Lock()
	c := cks.cookies
	cks.mutex.Unlock()

	return c
}

func (cks *cookies) getByName(name string) *http.Cookie {
	cks.mutex.Lock()
	var c *http.Cookie

	for _, c = range cks.cookies {
		if c.Name == name {
			break
		}
	}

	cks.mutex.Unlock()

	return c
}
