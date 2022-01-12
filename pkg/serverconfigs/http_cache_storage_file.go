package serverconfigs

// HTTPFileCacheStorage 文件缓存存储策略
type HTTPFileCacheStorage struct {
	Dir          string           `yaml:"dir" json:"dir"`                   // 目录
	MemoryPolicy *HTTPCachePolicy `yaml:"memoryPolicy" json:"memoryPolicy"` // 内存二级缓存

	OpenFileCache *OpenFileCacheConfig `yaml:"openFileCache" json:"openFileCache"` // open file cache配置
}

func (this *HTTPFileCacheStorage) Init() error {
	if this.OpenFileCache != nil {
		err := this.OpenFileCache.Init()
		if err != nil {
			return err
		}
	}
	return nil
}
