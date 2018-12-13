package store

import (
	"time"

	bolt "github.com/etcd-io/bbolt"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/bolthelper"
)

const (
	clusterBucket                = "clusters"
	clusterLastContactTimeBucket = "clusters_last_contact"
)

// Store provides storage functionality for alerts.
//go:generate mockgen-wrapper Store
type Store interface {
	GetCluster(id string) (*storage.Cluster, bool, error)
	GetClusters() ([]*storage.Cluster, error)
	CountClusters() (int, error)
	AddCluster(cluster *storage.Cluster) (string, error)
	UpdateCluster(cluster *storage.Cluster) error
	RemoveCluster(id string) error
	UpdateClusterContactTime(id string, t time.Time) error
	UpdateMetadata(id string, metadata *storage.ProviderMetadata) error
}

// New returns a new Store instance using the provided bolt DB instance.
func New(db *bolt.DB) Store {
	bolthelper.RegisterBucketOrPanic(db, clusterBucket)
	bolthelper.RegisterBucketOrPanic(db, clusterLastContactTimeBucket)
	return &storeImpl{
		DB: db,
	}
}
