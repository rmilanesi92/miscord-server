package data

import "sync"

var DBSet = map[string]string{}
var DBSetMutex = sync.RWMutex{}
