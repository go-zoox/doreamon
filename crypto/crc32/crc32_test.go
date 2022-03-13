package crc32

import (
	"testing"

	"github.com/go-zoox/doreamon/test"
)

func Test_CRC32_Signed(t *testing.T) {
	type d struct {
		expected int32
		i        string
	}

	data := []d{
		{
			244476679,
			"43741.0:8.176:43742.5:1.247:43740.5:7.830:43743.5:18.560:43740.0:48.006:43744.5:0.783:43739.5:2.376:43745.0:1.620:43739.0:17.646:43747.5:2.479:43738.5:2.592:43748.0:1.691:43738.0:3.548:43749.5:0.750:43737.5:3.186:43750.5:0.756:43737.0:15.433:43751.5:0.837:43736.5:0.516:43752.0:0.508:43736.0:26.938:43752.5:48.021:43735.0:1.654:43753.5:10.098:43734.0:7.416:43754.0:3.078:43733.0:48.040:43754.5:1.947:43732.0:48.017:43755.0:1.053:43731.0:5.275:43755.5:21.654:43730.5:8.073:43756.0:4.266:43730.0:48.001:43756.5:4.606:43729.5:0.524:43757.0:12.323:43729.0:9.332:43758.0:3.996:43728.0:48.017:43758.5:0.648:43727.5:9.715:43759.0:0.594:43727.0:48.038:43759.5:4.860:43726.5:19.791:43760.5:3.148:43726.0:43.435:43761.0:2.376",
		},
		{
			-1911312001,
			"3349.93:0.0029:3351.39:4.7664:3349.92:0.3030:3351.40:4.9973:3349.90:1.0099:3351.42:1.6352:3349.89:5.9996:3351.43:0.0602:3349.88:1.1110:3351.45:30.6329:3349.86:7.5818:3351.46:0.5050:3349.83:0.0055:3351.47:2.7394:3349.81:0.4471:3351.48:0.1608:3349.80:0.0046:3351.53:2.2281:3349.78:0.7272:3351.54:1.0829:3349.73:1.3147:3351.55:2.0316:3349.72:0.0059:3351.62:0.0301:3349.70:0.5929:3351.63:2.8280:3349.66:0.4124:3351.65:2.8566:3349.65:0.3562:3351.66:0.2024:3349.64:0.7790:3351.68:2.5250:3349.63:0.0787:3351.70:0.1280:3349.62:0.6040:3351.73:0.0036:3349.61:0.7533:3351.74:0.0058:3349.57:3.0300:3351.77:3.6604:3349.56:5.9996:3351.78:0.0808:3349.55:0.0983:3351.80:0.7272:3349.52:0.2395:3351.81:0.0824:3349.51:1.0008:3351.82:0.5050:3349.50:6.2071:3351.83:0.2712",
		},
	}

	ttt := test.TestSuit{
		T: t,
	}
	for _, one := range data {
		ttt.Expect(ChecksumSigned(one.i)).ToEqual(one.expected)
	}
}

func Test_CRC32(t *testing.T) {
	type d struct {
		expected uint32
		i        string
	}

	data := []d{
		{
			244476679,
			"43741.0:8.176:43742.5:1.247:43740.5:7.830:43743.5:18.560:43740.0:48.006:43744.5:0.783:43739.5:2.376:43745.0:1.620:43739.0:17.646:43747.5:2.479:43738.5:2.592:43748.0:1.691:43738.0:3.548:43749.5:0.750:43737.5:3.186:43750.5:0.756:43737.0:15.433:43751.5:0.837:43736.5:0.516:43752.0:0.508:43736.0:26.938:43752.5:48.021:43735.0:1.654:43753.5:10.098:43734.0:7.416:43754.0:3.078:43733.0:48.040:43754.5:1.947:43732.0:48.017:43755.0:1.053:43731.0:5.275:43755.5:21.654:43730.5:8.073:43756.0:4.266:43730.0:48.001:43756.5:4.606:43729.5:0.524:43757.0:12.323:43729.0:9.332:43758.0:3.996:43728.0:48.017:43758.5:0.648:43727.5:9.715:43759.0:0.594:43727.0:48.038:43759.5:4.860:43726.5:19.791:43760.5:3.148:43726.0:43.435:43761.0:2.376",
		},
		{
			2383655295,
			"3349.93:0.0029:3351.39:4.7664:3349.92:0.3030:3351.40:4.9973:3349.90:1.0099:3351.42:1.6352:3349.89:5.9996:3351.43:0.0602:3349.88:1.1110:3351.45:30.6329:3349.86:7.5818:3351.46:0.5050:3349.83:0.0055:3351.47:2.7394:3349.81:0.4471:3351.48:0.1608:3349.80:0.0046:3351.53:2.2281:3349.78:0.7272:3351.54:1.0829:3349.73:1.3147:3351.55:2.0316:3349.72:0.0059:3351.62:0.0301:3349.70:0.5929:3351.63:2.8280:3349.66:0.4124:3351.65:2.8566:3349.65:0.3562:3351.66:0.2024:3349.64:0.7790:3351.68:2.5250:3349.63:0.0787:3351.70:0.1280:3349.62:0.6040:3351.73:0.0036:3349.61:0.7533:3351.74:0.0058:3349.57:3.0300:3351.77:3.6604:3349.56:5.9996:3351.78:0.0808:3349.55:0.0983:3351.80:0.7272:3349.52:0.2395:3351.81:0.0824:3349.51:1.0008:3351.82:0.5050:3349.50:6.2071:3351.83:0.2712",
		},
	}

	ttt := test.TestSuit{T: t}
	for _, one := range data {
		ttt.Expect(Checksum(one.i)).ToEqual(one.expected)
	}
}
