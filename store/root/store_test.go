package root

import (
	"fmt"
	"testing"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/suite"

	coreheader "cosmossdk.io/core/header"
	"cosmossdk.io/log"
	"cosmossdk.io/store/v2"
	"cosmossdk.io/store/v2/commitment"
	"cosmossdk.io/store/v2/commitment/iavl"
	"cosmossdk.io/store/v2/pruning"
	"cosmossdk.io/store/v2/storage"
	"cosmossdk.io/store/v2/storage/sqlite"
)

const (
	testStoreKey = "test_store_key"
)

type RootStoreTestSuite struct {
	suite.Suite

	rootStore store.RootStore
}

func TestStorageTestSuite(t *testing.T) {
	suite.Run(t, &RootStoreTestSuite{})
}

func (s *RootStoreTestSuite) SetupTest() {
	noopLog := log.NewNopLogger()

	sqliteDB, err := sqlite.New(s.T().TempDir())
	s.Require().NoError(err)
	ss := storage.NewStorageStore(sqliteDB)

	tree := iavl.NewIavlTree(dbm.NewMemDB(), noopLog, iavl.DefaultConfig())
	sc, err := commitment.NewCommitStore(map[string]commitment.Tree{testStoreKey: tree}, noopLog)
	s.Require().NoError(err)

	rs, err := New(noopLog, ss, sc, pruning.DefaultOptions(), pruning.DefaultOptions(), nil)
	s.Require().NoError(err)

	s.rootStore = rs
}

func (s *RootStoreTestSuite) TearDownTest() {
	err := s.rootStore.Close()
	s.Require().NoError(err)
}

func (s *RootStoreTestSuite) TestGetStateCommitment() {
	s.Require().Equal(s.rootStore.GetStateCommitment(), s.rootStore.(*Store).stateCommitment)
}

func (s *RootStoreTestSuite) TestGetStateStorage() {
	s.Require().Equal(s.rootStore.GetStateStorage(), s.rootStore.(*Store).stateStore)
}

func (s *RootStoreTestSuite) TestSetInitialVersion() {
	s.Require().NoError(s.rootStore.SetInitialVersion(100))
}

func (s *RootStoreTestSuite) TestSetCommitHeader() {
	h := &coreheader.Info{
		Height:  100,
		Hash:    []byte("foo"),
		ChainID: "test",
	}
	s.rootStore.SetCommitHeader(h)

	s.Require().Equal(h, s.rootStore.(*Store).commitHeader)
}

func (s *RootStoreTestSuite) TestQuery() {
	_, err := s.rootStore.Query("", 1, []byte("foo"), true)
	s.Require().Error(err)

	// write and commit a changeset
	cs := store.NewChangeset()
	cs.Add(testStoreKey, []byte("foo"), []byte("bar"))

	workingHash, err := s.rootStore.WorkingHash(cs)
	s.Require().NoError(err)
	s.Require().NotNil(workingHash)

	commitHash, err := s.rootStore.Commit(cs)
	s.Require().NoError(err)
	s.Require().NotNil(commitHash)
	s.Require().Equal(workingHash, commitHash)

	// ensure the proof is non-nil for the corresponding version
	result, err := s.rootStore.Query(testStoreKey, 1, []byte("foo"), true)
	s.Require().NoError(err)
	s.Require().NotNil(result.Proof.Proof)
	s.Require().Equal([]byte("foo"), result.Proof.Proof.GetExist().Key)
	s.Require().Equal([]byte("bar"), result.Proof.Proof.GetExist().Value)
}

func (s *RootStoreTestSuite) TestLoadVersion() {
	// write and commit a few changesets
	for v := 1; v <= 5; v++ {
		val := fmt.Sprintf("val%03d", v) // val001, val002, ..., val005

		cs := store.NewChangeset()
		cs.Add(testStoreKey, []byte("key"), []byte(val))

		workingHash, err := s.rootStore.WorkingHash(cs)
		s.Require().NoError(err)
		s.Require().NotNil(workingHash)

		commitHash, err := s.rootStore.Commit(cs)
		s.Require().NoError(err)
		s.Require().NotNil(commitHash)
		s.Require().Equal(workingHash, commitHash)
	}

	// ensure the latest version is correct
	latest, err := s.rootStore.GetLatestVersion()
	s.Require().NoError(err)
	s.Require().Equal(uint64(5), latest)

	// attempt to load a non-existent version
	err = s.rootStore.LoadVersion(6)
	s.Require().Error(err)

	// attempt to load a previously committed version
	err = s.rootStore.LoadVersion(3)
	s.Require().NoError(err)

	// ensure the latest version is correct
	latest, err = s.rootStore.GetLatestVersion()
	s.Require().NoError(err)
	s.Require().Equal(uint64(3), latest)

	// query state and ensure values returned are based on the loaded version
	_, ro, err := s.rootStore.StateLatest()
	s.Require().NoError(err)

	val, err := ro.Get(testStoreKey, []byte("key"))
	s.Require().NoError(err)
	s.Require().Equal([]byte("val003"), val)

	// attempt to write and commit a few changesets
	for v := 4; v <= 5; v++ {
		val := fmt.Sprintf("overwritten_val%03d", v) // overwritten_val004, overwritten_val005

		cs := store.NewChangeset()
		cs.Add(testStoreKey, []byte("key"), []byte(val))

		workingHash, err := s.rootStore.WorkingHash(cs)
		s.Require().NoError(err)
		s.Require().NotNil(workingHash)

		commitHash, err := s.rootStore.Commit(cs)
		s.Require().NoError(err)
		s.Require().NotNil(commitHash)
		s.Require().Equal(workingHash, commitHash)
	}

	// ensure the latest version is correct
	latest, err = s.rootStore.GetLatestVersion()
	s.Require().NoError(err)
	s.Require().Equal(uint64(5), latest)

	// query state and ensure values returned are based on the loaded version
	_, ro, err = s.rootStore.StateLatest()
	s.Require().NoError(err)

	val, err = ro.Get(testStoreKey, []byte("key"))
	s.Require().NoError(err)
	s.Require().Equal([]byte("overwritten_val005"), val)
}

func (s *RootStoreTestSuite) TestCommit() {
	lv, err := s.rootStore.GetLatestVersion()
	s.Require().NoError(err)
	s.Require().Zero(lv)

	// perform changes
	cs := store.NewChangeset()
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%03d", i) // key000, key001, ..., key099
		val := fmt.Sprintf("val%03d", i) // val000, val001, ..., val099

		cs.Add(testStoreKey, []byte(key), []byte(val))
	}

	// committing w/o calling WorkingHash should error
	_, err = s.rootStore.Commit(cs)
	s.Require().Error(err)

	// execute WorkingHash and Commit
	wHash, err := s.rootStore.WorkingHash(cs)
	s.Require().NoError(err)

	cHash, err := s.rootStore.Commit(cs)
	s.Require().NoError(err)
	s.Require().Equal(wHash, cHash)

	// ensure latest version is updated
	lv, err = s.rootStore.GetLatestVersion()
	s.Require().NoError(err)
	s.Require().Equal(uint64(1), lv)

	// perform reads on the updated root store
	_, ro, err := s.rootStore.StateLatest()
	s.Require().NoError(err)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%03d", i) // key000, key001, ..., key099
		val := fmt.Sprintf("val%03d", i) // val000, val001, ..., val099

		result, err := ro.Get(testStoreKey, []byte(key))
		s.Require().NoError(err)

		s.Require().Equal([]byte(val), result)
	}
}

func (s *RootStoreTestSuite) TestStateAt() {
	// write keys over multiple versions
	for v := uint64(1); v <= 5; v++ {
		// perform changes
		cs := store.NewChangeset()
		for i := 0; i < 100; i++ {
			key := fmt.Sprintf("key%03d", i)         // key000, key001, ..., key099
			val := fmt.Sprintf("val%03d_%03d", i, v) // val000_1, val001_1, ..., val099_1

			cs.Add(testStoreKey, []byte(key), []byte(val))
		}

		// execute WorkingHash and Commit
		wHash, err := s.rootStore.WorkingHash(cs)
		s.Require().NoError(err)

		cHash, err := s.rootStore.Commit(cs)
		s.Require().NoError(err)
		s.Require().Equal(wHash, cHash)
	}

	lv, err := s.rootStore.GetLatestVersion()
	s.Require().NoError(err)
	s.Require().Equal(uint64(5), lv)

	// ensure we can read state correctly at each version
	for v := uint64(1); v <= 5; v++ {
		ro, err := s.rootStore.StateAt(v)
		s.Require().NoError(err)

		for i := 0; i < 100; i++ {
			key := fmt.Sprintf("key%03d", i)         // key000, key001, ..., key099
			val := fmt.Sprintf("val%03d_%03d", i, v) // val000_1, val001_1, ..., val099_1

			result, err := ro.Get(testStoreKey, []byte(key))
			s.Require().NoError(err)
			s.Require().Equal([]byte(val), result)
		}
	}
}
