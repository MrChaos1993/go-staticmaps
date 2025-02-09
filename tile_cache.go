package sm

import (
	"os"
	"time"
)

// TileCache provides cache information to the tile fetcher
type TileCache interface {
	// Root path to store cached tiles in with no trailing slash.
	Path() string
	// Permission to set when creating missing cache directories.
	Perm() os.FileMode
	TTL() time.Duration
}

// TileCacheStaticPath provides a static path to the tile fetcher.
type TileCacheStaticPath struct {
	path string
	perm os.FileMode
	ttl  time.Duration
}

// Path to the cache.
func (c *TileCacheStaticPath) Path() string {
	return c.path
}

// Perm instructs the permission to set when creating missing cache directories.
func (c *TileCacheStaticPath) Perm() os.FileMode {
	return c.perm
}

func (c *TileCacheStaticPath) TTL() time.Duration {
	return c.ttl
}

// NewTileCache stores cache files in a static path.
func NewTileCache(rootPath string, perm os.FileMode) *TileCacheStaticPath {
	return &TileCacheStaticPath{
		path: rootPath,
		perm: perm,
	}
}

// NewTileCacheFromUserCache stores cache files in a user-specific cache directory.
func NewTileCacheFromUserCache(perm os.FileMode) *TileCacheStaticPath {
	path, err := os.UserCacheDir()
	if err != nil {
		path += "/go-staticmaps"
	}
	return &TileCacheStaticPath{
		path: path,
		perm: perm,
	}
}

// NewTileCacheWithTTL stores cache files in a static path.
func NewTileCacheWithTTL(rootPath string, perm os.FileMode, ttl time.Duration) *TileCacheStaticPath {
	return &TileCacheStaticPath{
		path: rootPath,
		perm: perm,
		ttl:  ttl,
	}
}
