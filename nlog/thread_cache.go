package nlog

import (
	"context"
	"fmt"
	"github.com/patrickmn/go-cache"
	"net/http"
	"sync/atomic"
	"time"
)

var inited = int32(0)
var CACHE_INSTANCE *cache.Cache

var CacheKeyBase = "run-thread:%d"

type ThreadCache struct {
	TraceId   string
	RequestId string
}

//init cache
func InitCache() {
	if atomic.CompareAndSwapInt32(&inited, 0, 1) {
		CACHE_INSTANCE = cache.New(1*time.Minute, 5*time.Minute)
	}
}

func SetTraceInfoFromContext(ctx context.Context) {
	traceId := ctx.Value(ISTIO_TRACER_ID)
	if traceId == nil {
		traceId = DefaultTraceId()
	}
	requestId := ctx.Value(ISTIO_REQUEST_ID)
	if requestId == nil {
		requestId = DefaultRequestId()
	}

	tcache := &ThreadCache{
		RequestId: requestId.(string),
		TraceId:   traceId.(string),
	}
	SetIstioHeader2ThreadCache(tcache)
}

func SetTraceFromHttpRequest(r *http.Request) {
	traceId := r.Header.Get(ISTIO_TRACER_ID)
	if traceId == "" {
		traceId = DefaultTraceId()
	}
	requestId := r.Header.Get(ISTIO_REQUEST_ID)
	if requestId == "" {
		requestId = DefaultRequestId()
	}
	tcache := &ThreadCache{
		RequestId: requestId,
		TraceId:   traceId,
	}
	SetIstioHeader2ThreadCache(tcache)
}


func SetIstioHeader2ThreadCache(tcache *ThreadCache) {
	cacheKey := GetThreadCacheKey()
	if inited != 1 {
		return
	}
	CACHE_INSTANCE.Set(cacheKey, tcache, time.Minute)
}

func RemoveCache() {
	cacheKey := GetThreadCacheKey()
	CACHE_INSTANCE.Delete(cacheKey)
}

func GetThreadCache() (cacheEntry *ThreadCache, bol bool) {

	if inited != 1 {
		return nil, false
	}

	ca, bol := CACHE_INSTANCE.Get(GetThreadCacheKey())
	if !bol {
		return nil, bol
	}
	entry := ca.(*ThreadCache)
	return entry, true
}

func GetThreadCacheKey() string {
	runTimeId := CurGoroutineID()
	cacheKey := fmt.Sprintf(CacheKeyBase, runTimeId)
	return cacheKey
}
