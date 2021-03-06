package configutils

import (
	"regexp"
	"strings"
	"sync"
)

// VariableHolder 变量信息存储类型
type VariableHolder string
type VariableHolders = []interface{}

var variableMapping = map[string][]interface{}{} // source => [holder1, ...]
var variableLocker = sync.RWMutex{}
var regexpNamedVariable = regexp.MustCompile("\\${[\\w.-]+}")

// ParseVariables 分析变量
func ParseVariables(source string, replacer func(varName string) (value string)) string {
	variableLocker.RLock()
	holders, found := variableMapping[source]
	variableLocker.RUnlock()
	if !found {
		holders = ParseHolders(source)
		variableLocker.Lock()
		variableMapping[source] = holders
		variableLocker.Unlock()
	}

	// no variables
	if len(holders) == 0 {
		return source
	}

	// replace
	result := strings.Builder{}
	for _, h := range holders {
		holder, ok := h.(VariableHolder)
		if ok {
			result.WriteString(replacer(string(holder)))
		} else {
			result.Write(h.([]byte))
		}
	}
	return result.String()
}

// ParseVariablesFromHolders 从占位中分析变量
func ParseVariablesFromHolders(holders VariableHolders, replacer func(varName string) (value string)) string {
	// no variables
	if len(holders) == 0 {
		return ""
	}

	// replace
	result := strings.Builder{}
	for _, h := range holders {
		holder, ok := h.(VariableHolder)
		if ok {
			result.WriteString(replacer(string(holder)))
		} else {
			result.Write(h.([]byte))
		}
	}
	return result.String()
}

// ParseHolders 分析占位
func ParseHolders(source string) (holders VariableHolders) {
	indexes := regexpNamedVariable.FindAllStringIndex(source, -1)
	before := 0
	for _, loc := range indexes {
		holders = append(holders, []byte(source[before:loc[0]]))
		holder := source[loc[0]+2 : loc[1]-1]
		holders = append(holders, VariableHolder(holder))
		before = loc[1]
	}
	if before < len(source) {
		holders = append(holders, []byte(source[before:]))
	}
	return holders
}

// HasVariables 判断是否有变量
func HasVariables(source string) bool {
	return regexpNamedVariable.MatchString(source)
}
