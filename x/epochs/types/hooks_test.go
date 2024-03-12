package types_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/core/appmodule"
	errors "cosmossdk.io/errors"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/epochs/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type KeeperTestSuite struct {
	suite.Suite
	Ctx         sdk.Context
	environment appmodule.Environment
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(types.StoreKey)
	storeService := runtime.NewKVStoreService(key)
	s.Ctx = testutil.DefaultContext(storetypes.NewKVStoreKey(types.StoreKey), storetypes.NewTransientStoreKey("transient_test"))
	s.environment = runtime.NewEnvironment(storeService, log.NewNopLogger())
}

func dummyAfterEpochEndEvent(epochIdentifier string, epochNumber int64) sdk.Event {
	return sdk.NewEvent(
		"afterEpochEnd",
		sdk.NewAttribute("epochIdentifier", epochIdentifier),
		sdk.NewAttribute("epochNumber", strconv.FormatInt(epochNumber, 10)),
	)
}

func dummyBeforeEpochStartEvent(epochIdentifier string, epochNumber int64) sdk.Event {
	return sdk.NewEvent(
		"beforeEpochStart",
		sdk.NewAttribute("epochIdentifier", epochIdentifier),
		sdk.NewAttribute("epochNumber", strconv.FormatInt(epochNumber, 10)),
	)
}

var dummyErr = errors.New("9", 9, "dummyError")

// dummyEpochHook is a struct satisfying the epoch hook interface,
// that maintains a counter for how many times its been successfully called,
// and a boolean for whether it should panic during its execution.
type dummyEpochHook struct {
	successCounter int
	shouldPanic    bool
	shouldError    bool
}

// GetModuleName implements types.EpochHooks.
func (*dummyEpochHook) GetModuleName() string {
	return "dummy"
}

func (hook *dummyEpochHook) AfterEpochEnd(ctx context.Context, epochIdentifier string, epochNumber int64, env appmodule.Environment) error {
	if hook.shouldPanic {
		panic("dummyEpochHook is panicking")
	}
	if hook.shouldError {
		return dummyErr
	}
	hook.successCounter += 1
	return nil
}

func (hook *dummyEpochHook) BeforeEpochStart(ctx context.Context, epochIdentifier string, epochNumber int64, env appmodule.Environment) error {
	if hook.shouldPanic {
		panic("dummyEpochHook is panicking")
	}
	if hook.shouldError {
		return dummyErr
	}
	hook.successCounter += 1
	return nil
}

func (hook *dummyEpochHook) Clone() *dummyEpochHook {
	newHook := dummyEpochHook{shouldPanic: hook.shouldPanic, successCounter: hook.successCounter, shouldError: hook.shouldError}
	return &newHook
}

var _ types.EpochHooks = &dummyEpochHook{}

func (s *KeeperTestSuite) TestHooksPanicRecovery() {
	panicHook := dummyEpochHook{shouldPanic: true}
	noPanicHook := dummyEpochHook{shouldPanic: false}
	errorHook := dummyEpochHook{shouldError: true}
	noErrorHook := dummyEpochHook{shouldError: false} // same as nopanic
	simpleHooks := []dummyEpochHook{panicHook, noPanicHook, errorHook, noErrorHook}

	tests := []struct {
		hooks                 []dummyEpochHook
		expectedCounterValues []int
		lenEvents             int
	}{
		{[]dummyEpochHook{noPanicHook}, []int{1}, 1},
		{[]dummyEpochHook{panicHook}, []int{0}, 0},
		{[]dummyEpochHook{errorHook}, []int{0}, 0},
		{simpleHooks, []int{0, 1, 0, 1}, 2},
	}

	for tcIndex, tc := range tests {
		for epochActionSelector := 0; epochActionSelector < 2; epochActionSelector++ {
			s.SetupTest()
			hookRefs := []types.EpochHooks{}

			for _, hook := range tc.hooks {
				hookRefs = append(hookRefs, hook.Clone())
			}

			hooks := types.NewMultiEpochHooks(hookRefs...)

			events := func(epochID string, epochNumber int64, dummyEvent func(id string, number int64) sdk.Event) sdk.Events {
				evts := make(sdk.Events, tc.lenEvents)
				for i := 0; i < tc.lenEvents; i++ {
					evts[i] = dummyEvent(epochID, epochNumber)
				}
				return evts
			}

			s.NotPanics(func() {
				if epochActionSelector == 0 {
					err := hooks.BeforeEpochStart(s.Ctx, "id", 0, s.environment)
					s.Require().NoError(err)
					s.Require().Equal(events("id", 0, dummyBeforeEpochStartEvent), s.Ctx.EventManager().Events(),
						"test case index %d, before epoch event check", tcIndex)
				} else if epochActionSelector == 1 {
					err := hooks.AfterEpochEnd(s.Ctx, "id", 0, s.environment)
					s.Require().NoError(err)
					s.Require().Equal(events("id", 0, dummyAfterEpochEndEvent), s.Ctx.EventManager().Events(),
						"test case index %d, after epoch event check", tcIndex)
				}
			})

			for i := 0; i < len(hooks); i++ {
				epochHook := hookRefs[i].(*dummyEpochHook)
				s.Require().Equal(tc.expectedCounterValues[i], epochHook.successCounter, "test case index %d", tcIndex)
			}
		}
	}
}
