package lepet

import (
	"golang.org/x/exp/maps"
)

// Configuration type that is used by the core
type LangItem map[string]string
type LangData struct {
	Active string
	List   map[string]LangItem
}

// Instance of the language
var I *LangData // instance

func New() LangData {
	return LangData{}
}

func (a *LangData) SetList(ln string, l LangItem) {
	a.List = map[string]LangItem{ln: l}
}

func (a *LangData) Add(name string) {
	_, ok := a.List[name]
	if !ok {
		a.List[name] = LangItem{}
	}
}

func (a *LangData) AddMsg(code string, message string, opt ...string) {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}
	a.List[lang][code] = message
}

func (a *LangData) AddMsgList(list LangItem, opt ...string) {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}

	maps.Copy(a.List[lang], list)
}

func (a *LangData) Msg(k string, opt ...string) string {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}

	if msg, ok := a.List[lang][k]; !ok {
		return "** warning: usage of unlisted code **"
	} else {
		return msg
	}
}
