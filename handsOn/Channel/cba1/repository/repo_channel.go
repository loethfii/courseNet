package repository

import "sync"

var wg sync.WaitGroup

// metode
var cuciChannel = make(chan string)
var potongCHannel = make(chan string)
var kupasChannel = make(chan string)

// simpan
var simpanFrozenChannel = make(chan string)
var SimpanChillerChannel = make(chan string)
var SimpanBoxChannel = make(chan string)
