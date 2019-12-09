package tx

import "github.com/intel-go/cpuid"

func Supported() bool { return cpuid.HasExtendedFeature(cpuid.RTM) }

const (
	Started       uint32 = 0xffffffff
	AbortExplicit uint32 = 1 << 0
	AbortRetry    uint32 = 1 << 1
	AbortConflict uint32 = 1 << 2
	AbortCapacity uint32 = 1 << 3
	AbortDebug    uint32 = 1 << 4
	AbortNested   uint32 = 1 << 5
)

func Begin() uint32

func Abort()

func End()

func Test() uint8
