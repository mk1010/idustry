package redisclient

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/bsm/redislock"
	"github.com/mk1010/idustry/config"
)

func TestMain(m *testing.M) {
	curEnv := config.CheckEnv()
	// mk
	configFile := fmt.Sprintf("../../conf/industry_identification_center_%s.json", curEnv)

	if err := config.Init(configFile); err != nil {
		panic(err)
	}
	if curEnv != config.ConfInstance.Env {
		panic(errors.New("env error"))
	}
	fmt.Printf("Service running in %s mode\n", curEnv)
	if err := InitRedisClient(); err != nil {
		panic(err)
	}
	m.Run()
}

func TestDistributionLock(t *testing.T) {
	ctx := context.Background()
	// t.Log(Rdb.SetNX(ctx, "lock_a", util.GetUuid(), 10*time.Second))
	locker := redislock.New(Rdb)
	lock, err := locker.Obtain(ctx, "lock_a", 10*time.Second, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(Rdb.Get(ctx, "lock_a"))
	t.Log(lock.Release(ctx))
}
