package os

import (
	"regexp"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type OSUsage struct {
	cpu    float64
	memory float64
	disk   float64
}

func GetCpuPercent() float64 {
	percent, _ := cpu.Percent(time.Second, false)
	// fmt.Println("all cpu:", percent)
	return percent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	deviceReg, _ := regexp.Compile("^/dev/")
	deviceMap := map[string]*disk.UsageStat{}
	deviceTotalSize := 0.0
	deviceUsedSize := 0.0

	for part := range parts {
		if deviceMap[parts[part].Device] != nil {
			continue
		}

		if match := deviceReg.FindString(parts[part].Device); match != "" {
			diskInfo, _ := disk.Usage(parts[part].Mountpoint)
			deviceMap[parts[part].Device] = diskInfo
			deviceTotalSize += float64(diskInfo.Total)
			deviceUsedSize += float64(diskInfo.Used)

			// fmt.Println("Mountpoint:", parts[part])
			// fmt.Println(parts[part].Device, diskInfo.Total, diskInfo.Used)
			// fmt.Println("diskInfo:", diskInfo)
		}
	}

	// fmt.Println(deviceUsedSize, deviceTotalSize)
	usedPercent := deviceUsedSize / deviceTotalSize * 100
	return usedPercent
}

func GetOSUsage() *OSUsage {
	cpu := GetCpuPercent()
	memory := GetMemPercent()
	disk := GetDiskPercent()

	return &OSUsage{
		cpu:    cpu,
		memory: memory,
		disk:   disk,
	}
}
