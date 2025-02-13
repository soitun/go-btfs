package repo

import (
	"errors"

	keystore "github.com/bittorrent/go-btfs/keystore"
	rcmgr "github.com/libp2p/go-libp2p/p2p/host/resource-manager"

	config "github.com/bittorrent/go-btfs-config"
	filestore "github.com/ipfs/go-filestore"
	ma "github.com/multiformats/go-multiaddr"
)

var errTODO = errors.New("TODO: mock repo")

// Mock is not thread-safe
type Mock struct {
	C config.Config
	D Datastore
	K keystore.Keystore
	F *filestore.FileManager

	Used uint64
}

func (m *Mock) Config() (*config.Config, error) {
	return &m.C, nil // FIXME threadsafety
}

func (m *Mock) SetConfig(updated *config.Config) error {
	m.C = *updated // FIXME threadsafety
	return nil
}

func (m *Mock) BackupConfig(prefix string) (string, error) {
	return "", errTODO
}

func (m *Mock) BackUpConfigV2(suffix string) (string, error) {
	return "", errTODO
}

func (m *Mock) SetConfigKey(key string, value interface{}) error {
	return errTODO
}

func (m *Mock) GetConfigKey(key string) (interface{}, error) {
	return nil, errTODO
}

func (m *Mock) Datastore() Datastore { return m.D }

func (m *Mock) GetStorageUsage() (uint64, error) { return m.Used, nil }

func (m *Mock) Close() error { return m.D.Close() }

func (m *Mock) SetAPIAddr(addr ma.Multiaddr) error { return errTODO }

func (m *Mock) Keystore() keystore.Keystore { return m.K }

func (m *Mock) SwarmKey() ([]byte, error) {
	return nil, nil
}

func (m *Mock) FileManager() *filestore.FileManager { return m.F }

func (m *Mock) UserResourceOverrides() (rcmgr.PartialLimitConfig, error) {
	return rcmgr.PartialLimitConfig{}, nil
}
