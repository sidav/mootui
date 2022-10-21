package game

const (
	PSPECIAL_ULTRA_POOR = iota
	PSPECIAL_POOR
	PSPECIAL_NORMAL
	PSPECIAL_ARTIFACTS
	PSPECIAL_RICH
	PSPECIAL_ULTRA_RICH
	PSPECIAL_TOTAL_NUMBER
)

const (
	PGROWTH_HOSTILE = iota
	PGROWTH_NORMAL
	PGROWTH_FERTILE
	PGROWTH_GAIA
	PGROWTH_TOTAL_NUMBER
)

func (p *planet) GetGrowthAndSpecialString() string {
	growth := "error"
	switch p.growth {
	case PGROWTH_HOSTILE: growth = "Hostile"
	case PGROWTH_NORMAL: growth = "Normal"
	case PGROWTH_FERTILE: growth = "Fertile"
	case PGROWTH_GAIA: growth = "Gaia"
	}
	special := "error"
	switch p.special {
	case PSPECIAL_ULTRA_POOR: special = "u-poor"
	case PSPECIAL_POOR: special = "poor"
	case PSPECIAL_NORMAL: special = ""
	case PSPECIAL_RICH: special = "rich"
	case PSPECIAL_ULTRA_RICH: special = "u-rich"
	case PSPECIAL_ARTIFACTS: special = "artifacts"
	}
	if special != "" {
		special = " " + special
	}
	return growth + special
}
