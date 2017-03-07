package compact

import (
	"fmt"
	"strconv"

	"github.com/bcicen/ctop/cwidgets"
	ui "github.com/gizak/termui"
)

func (row *Compact) SetNet(rx int64, tx int64) {
	label := fmt.Sprintf("%s / %s", cwidgets.ByteFormat(rx), cwidgets.ByteFormat(tx))
	row.Net.Set(label)
}

func (row *Compact) SetCPU(val int) {
	row.Cpu.BarColor = colorScale(val)
	row.Cpu.Label = fmt.Sprintf("%s%%", strconv.Itoa(val))
	if val < 5 {
		val = 5
		row.Cpu.BarColor = ui.ThemeAttr("gauge.bar.bg")
	}
	row.Cpu.Percent = val
}

func (row *Compact) SetMem(val int64, limit int64, percent int) {
	row.Memory.Label = fmt.Sprintf("%s / %s", cwidgets.ByteFormat(val), cwidgets.ByteFormat(limit))
	if percent < 5 {
		percent = 5
		row.Memory.BarColor = ui.ColorBlack
	} else {
		row.Memory.BarColor = ui.ThemeAttr("gauge.bar.bg")
	}
	row.Memory.Percent = percent
}
