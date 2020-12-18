package facebook

import (
	"encoding/json"
	"reflect"
	"sync"
)

// parser parse JSON bytes to []*raw.Item
type parser struct {
	data [][]byte
	wg   *sync.WaitGroup
	mx   *sync.Mutex
	stop bool
}

func newParser(data [][]byte) *parser {
	return &parser{
		wg:   new(sync.WaitGroup),
		mx:   new(sync.Mutex),
		data: data,
	}
}

func (p *parser) reset() {
	p.wg = new(sync.WaitGroup)
	p.mx = new(sync.Mutex)
	p.stop = false
}

// dest must be pointer to struct or array
func (p *parser) run(handler func(val interface{}) bool, dest interface{}, isObj, isArr bool) {
	for _, frag := range p.data {
		p.wg.Add(1)
		go p.tryToParse(handler, frag, dest, isObj, isArr)
	}

	p.wg.Wait()
}

// dest must be pointer to array or struct
func (p *parser) tryToParse(handler func(val interface{}) bool, data json.RawMessage, dest interface{}, isObj, isArr bool) {
	if p.stop {
		p.wg.Done()
		return
	}

	if isObj {
		arr := make([]json.RawMessage, 0)
		if err := json.Unmarshal(data, &arr); err == nil {
			for _, a := range arr {
				p.wg.Add(1)
				go p.tryToParse(handler, a, dest, isObj, isArr)
			}

			p.wg.Done()
			return
		}

		destVal := reflect.ValueOf(dest).Elem()
		destCpy := reflect.New(destVal.Type()).Interface()
		if err := json.Unmarshal(data, destCpy); err == nil {
			p.mx.Lock()
			if !p.stop {
				p.stop = handler(destCpy)
			}

			p.mx.Unlock()
			p.wg.Done()
			return
		}
	}

	if isArr {
		destVal := reflect.ValueOf(dest).Elem()
		destCpy := reflect.New(destVal.Type()).Interface()
		if err := json.Unmarshal(data, destCpy); err == nil {
			p.mx.Lock()
			if !p.stop {
				p.stop = handler(destCpy)
			}

			p.mx.Unlock()
			p.wg.Done()
			return
		}
	}

	p.wg.Done()
}
