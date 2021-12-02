package main

import (
	"context"
	"fmt"
	"net/http"
)

type (
	bugsnag struct {
		reportChan chan error
	}
)

type ReportFunc func(err error)

func NewBugsnag() *bugsnag {

	return &bugsnag{
		reportChan: make(chan error),
	}
}

func (b *bugsnag) Watch() {
	for {
		select {
		case err := <-b.reportChan:
			fmt.Println("kirim ke bugsnag")
			fmt.Println(err)
		}
	}
}

func (b *bugsnag) Close() error {
	return nil
}

func (b *bugsnag) Report(err error) {
	go func() {
		b.reportChan <- err
	}()

}

func (b *bugsnag) Handler(next http.Handler) http.Handler {

	var fnReport ReportFunc
	fnReport = func(err error) {
		b.Report(err)
	}

	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "bugsnag", fnReport)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

	}

	return http.HandlerFunc(fn)
}
