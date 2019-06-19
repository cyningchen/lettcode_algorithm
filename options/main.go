package main

import (
	"fmt"
	"time"
)

// 选项模式
type Options struct {
	EtcdAddr string
	Port     int
	Timeout  time.Duration
}

var (
	OPT *Options
)

type Option func(options *Options)

func InitOptions(opts ...Option) {
	OPT = &Options{}
	for _, opt := range opts {
		opt(OPT)
	}
}

func WithEtcdAddrOption(str string) Option {
	return func(options *Options) {
		options.EtcdAddr = str
	}
}

func WithPortOption(port int) Option {
	return func(options *Options) {
		options.Port = port
	}
}

func WithTimeoutOption(duration time.Duration) Option {
	return func(options *Options) {
		options.Timeout = duration
	}
}

func main() {
	InitOptions(WithEtcdAddrOption("localhost"), WithPortOption(80), WithTimeoutOption(time.Millisecond*5))
	fmt.Println(OPT)
}
