	xdocs "github.com/m3db/m3/src/x/docs"
	"github.com/m3db/m3/src/x/lockfile"
	"github.com/m3db/m3/src/x/mmap"
	"github.com/m3db/m3/src/x/serialize"
		SetMetricsSamplingRate(cfg.Metrics.SampleRate())
		SetWriteBatchPool(writeBatchPool)