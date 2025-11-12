package main

import "log"

func must[T any](in T, err error) T {
	if err != nil {
		log.Panic(err)
	}
	return in
}
