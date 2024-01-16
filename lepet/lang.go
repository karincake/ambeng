package lepet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/exp/maps"
)

// Configuration type that is used by the core
type LangConf struct {
	Active   string
	Path     string
	FileName string
}

type langItem map[string]string
type langData struct {
	Active string
	list   map[string]langItem
}

// Instance of the language
var I *langData // instance

func New() langData {
	return langData{}
}

func Init(conf LangConf) error {
	I = &langData{}
	I.Active = conf.Active
	I.list = map[string]langItem{"en": defaultList}

	jsonFile, err := os.Open(fmt.Sprintf("%v/%v/%v", conf.Path, conf.Active, conf.FileName))
	if err != nil {
		return errors.New("failed to load source file. " + err.Error())
	}
	defer jsonFile.Close()

	var myLI langItem = langItem{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &myLI)
	maps.Copy(I.list["en"], myLI)
	return nil
}

func (a *langData) Add(name string) {
	_, ok := a.list[name]
	if !ok {
		a.list[name] = langItem{}
	}
}

func (a *langData) AddMsg(code string, message string, opt ...string) {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}
	a.list[lang][code] = message
}

func (a *langData) AddMsgList(list langItem, opt ...string) {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}

	maps.Copy(a.list[lang], list)
}

func (a *langData) Msg(k string, opt ...string) string {
	lang := a.Active
	if len(opt) > 0 {
		a.Add(opt[1])
		lang = opt[1]
	}

	if msg, ok := a.list[lang][k]; !ok {
		return "** warning: usage of unlisted code **"
	} else {
		return msg
	}
}
