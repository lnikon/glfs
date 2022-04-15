package server

import (
	"fmt"
	"time"

	log "github.com/go-kit/log"
)

type LoggingMiddleware struct {
	Next   ComputationAllocationServiceIfc
	Logger log.Logger
}

func (mw LoggingMiddleware) GetAllocation(name string) (computation AllocationDescription, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetComputation",
			"input", fmt.Sprintf("%v", name),
			"output", fmt.Sprintf("%v", computation),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	computation, err = mw.Next.GetAllocation(name)
	if err != nil {
		mw.Logger.Log("Error: ", err.Error())
	}
	return
}

func (mw LoggingMiddleware) GetAllAllocations() (output []AllocationDescription) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetAllComputations",
			"output", fmt.Sprintf("%v", output),
			"took", time.Since(begin),
		)
	}(time.Now())

	output = mw.Next.GetAllAllocations()
	return
}

func (mw LoggingMiddleware) PostAllocation(description AllocationDescription) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "PostComputation",
			"input", fmt.Sprintf("%v", description),
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Next.PostAllocation(description)
	if err != nil {
		mw.Logger.Log("Error: ", err.Error())
	}
	return
}

func (mw LoggingMiddleware) DeleteAllocation(name string) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "DeleteComputation",
			"input", fmt.Sprintf("%v", name),
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Next.DeleteAllocation(name)
	if err != nil {
		mw.Logger.Log("Error: ", err.Error())
	}
	return
}
