package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/BenasB/tess-space-app/back/tess"
	"github.com/BenasB/tess-space-app/back/utils"
)

const latestSector = 97

func (h *ApiHandler) DownloadCCD(w http.ResponseWriter, r *http.Request) {
	sectorStr := r.URL.Query().Get("sector")
	if sectorStr == "" {
		log.Printf("Missing sector parameter")
		http.Error(w, "Missing sector parameter", http.StatusBadRequest)
		return
	}

	sector, err := strconv.Atoi(sectorStr)
	if err != nil {
		http.Error(w, "Sector must be a number", http.StatusBadRequest)
		return
	}

	if sector < 1 || sector > latestSector {
		http.Error(w, fmt.Sprintf("Sector must be between 1 and %d", latestSector), http.StatusBadRequest)
		return
	}

	const cam = 1
	const ccd = 1
	info := tessDataMap[sector]
	resource := fmt.Sprintf("tess%s-s%04d-%d-%d-%s-s_ffic.fits", info.Timestamp, sector, cam, ccd, info.SpacecraftNumber)

	localPath, err := h.MastClient.DownloadSingleFile(resource)
	if err != nil {
		log.Printf("Failed to download file: %v", err)
		http.Error(w, fmt.Sprintf("Failed to download file: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	img, err := tess.ConvertFFIToImage(localPath)
	if err != nil {
		log.Printf("Failed to convert file to image: %v", err)
		http.Error(w, fmt.Sprintf("Failed to convert file to image: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	buffer, err := utils.ExportImageToPngBuffer(img)
	if err != nil {
		log.Printf("Failed to export image to bytes buffer: %v", err)
		http.Error(w, fmt.Sprintf("Failed to export image to bytes buffer: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")

	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Printf("Failed to write image to response: %v", err)
	}

	log.Printf("Image sent successfully: %s", resource)
}

var tessDataMap = map[int]struct {
	Timestamp        string
	SpacecraftNumber string
}{
	1:  {Timestamp: "2018216092942", SpacecraftNumber: "0120"},
	2:  {Timestamp: "2018235212941", SpacecraftNumber: "0121"},
	3:  {Timestamp: "2018282025940", SpacecraftNumber: "0123"},
	4:  {Timestamp: "2018297215939", SpacecraftNumber: "0124"},
	5:  {Timestamp: "2018333035938", SpacecraftNumber: "0125"},
	6:  {Timestamp: "2018353022937", SpacecraftNumber: "0126"},
	7:  {Timestamp: "2019009195936", SpacecraftNumber: "0131"},
	8:  {Timestamp: "2019033212935", SpacecraftNumber: "0136"},
	9:  {Timestamp: "2019059195935", SpacecraftNumber: "0139"},
	10: {Timestamp: "2019089065934", SpacecraftNumber: "0140"},
	11: {Timestamp: "2019114062933", SpacecraftNumber: "0143"},
	12: {Timestamp: "2019145202931", SpacecraftNumber: "0144"},
	13: {Timestamp: "2019172122930", SpacecraftNumber: "0146"},
	14: {Timestamp: "2019209132929", SpacecraftNumber: "0150"},
	15: {Timestamp: "2019239065928", SpacecraftNumber: "0151"},
	16: {Timestamp: "2019256212927", SpacecraftNumber: "0152"},
	17: {Timestamp: "2019297002926", SpacecraftNumber: "0161"},
	18: {Timestamp: "2019312005925", SpacecraftNumber: "0162"},
	19: {Timestamp: "2019355002924", SpacecraftNumber: "0164"},
	20: {Timestamp: "2019358235923", SpacecraftNumber: "0165"},
	21: {Timestamp: "2020022032922", SpacecraftNumber: "0167"},
	22: {Timestamp: "2020050195921", SpacecraftNumber: "0174"},
	23: {Timestamp: "2020084105920", SpacecraftNumber: "0177"},
	24: {Timestamp: "2020107065919", SpacecraftNumber: "0180"},
	25: {Timestamp: "2020145235918", SpacecraftNumber: "0182"},
	26: {Timestamp: "2020161182917", SpacecraftNumber: "0187"},
	27: {Timestamp: "2020209215915", SpacecraftNumber: "0189"},
	28: {Timestamp: "2020226153915", SpacecraftNumber: "0190"},
	29: {Timestamp: "2020242021914", SpacecraftNumber: "0193"},
	30: {Timestamp: "2020285032912", SpacecraftNumber: "0195"},
	31: {Timestamp: "2020302172912", SpacecraftNumber: "0198"},
	32: {Timestamp: "2020327162911", SpacecraftNumber: "0200"},
	33: {Timestamp: "2020361150909", SpacecraftNumber: "0203"},
	34: {Timestamp: "2021019085909", SpacecraftNumber: "0204"},
	35: {Timestamp: "2021048194907", SpacecraftNumber: "0205"},
	36: {Timestamp: "2021085232906", SpacecraftNumber: "0207"},
	37: {Timestamp: "2021099193905", SpacecraftNumber: "0208"},
	38: {Timestamp: "2021143220904", SpacecraftNumber: "0209"},
	39: {Timestamp: "2021158184903", SpacecraftNumber: "0210"},
	40: {Timestamp: "2021182154902", SpacecraftNumber: "0211"},
	41: {Timestamp: "2021205231901", SpacecraftNumber: "0212"},
	42: {Timestamp: "2021237153900", SpacecraftNumber: "0213"},
	43: {Timestamp: "2021259224859", SpacecraftNumber: "0214"},
	44: {Timestamp: "2021285231858", SpacecraftNumber: "0215"},
	45: {Timestamp: "2021311163857", SpacecraftNumber: "0216"},
	46: {Timestamp: "2021339135856", SpacecraftNumber: "0217"},
	47: {Timestamp: "2022001200855", SpacecraftNumber: "0218"},
	48: {Timestamp: "2022029055854", SpacecraftNumber: "0219"},
	49: {Timestamp: "2022057232853", SpacecraftNumber: "0221"},
	50: {Timestamp: "2022085185852", SpacecraftNumber: "0222"},
	51: {Timestamp: "2022114194851", SpacecraftNumber: "0223"},
	52: {Timestamp: "2022139155850", SpacecraftNumber: "0224"},
	53: {Timestamp: "2022164233849", SpacecraftNumber: "0226"},
	54: {Timestamp: "2022190211848", SpacecraftNumber: "0227"},
	55: {Timestamp: "2022217141847", SpacecraftNumber: "0240"},
	56: {Timestamp: "2022246070525", SpacecraftNumber: "0243"},
	57: {Timestamp: "2022280081204", SpacecraftNumber: "0245"},
	58: {Timestamp: "2022329012842", SpacecraftNumber: "0247"},
	59: {Timestamp: "2022330223842", SpacecraftNumber: "0248"},
	60: {Timestamp: "2022361031201", SpacecraftNumber: "0249"},
	61: {Timestamp: "2023020225520", SpacecraftNumber: "0250"},
	62: {Timestamp: "2023054093839", SpacecraftNumber: "0254"},
	63: {Timestamp: "2023081192838", SpacecraftNumber: "0255"},
	64: {Timestamp: "2023100171837", SpacecraftNumber: "0257"},
	65: {Timestamp: "2023134204835", SpacecraftNumber: "0259"},
	66: {Timestamp: "2023160062154", SpacecraftNumber: "0260"},
	67: {Timestamp: "2023194125153", SpacecraftNumber: "0261"},
	68: {Timestamp: "2023212053152", SpacecraftNumber: "0262"},
	69: {Timestamp: "2023253072511", SpacecraftNumber: "0264"},
	70: {Timestamp: "2023281211150", SpacecraftNumber: "0265"},
	71: {Timestamp: "2023306131829", SpacecraftNumber: "0266"},
	72: {Timestamp: "2023322145148", SpacecraftNumber: "0267"},
	73: {Timestamp: "2024002094826", SpacecraftNumber: "0268"},
	74: {Timestamp: "2024011174826", SpacecraftNumber: "0269"},
	75: {Timestamp: "2024031133825", SpacecraftNumber: "0270"},
	76: {Timestamp: "2024060023144", SpacecraftNumber: "0271"},
	77: {Timestamp: "2024113185302", SpacecraftNumber: "0272"},
	78: {Timestamp: "2024132174622", SpacecraftNumber: "0273"},
	79: {Timestamp: "2024150031301", SpacecraftNumber: "0274"},
	80: {Timestamp: "2024181021300", SpacecraftNumber: "0275"},
	81: {Timestamp: "2024207054939", SpacecraftNumber: "0276"},
	82: {Timestamp: "2024244142937", SpacecraftNumber: "0278"},
	83: {Timestamp: "2024255165257", SpacecraftNumber: "0280"},
	84: {Timestamp: "2024276020616", SpacecraftNumber: "0281"},
	85: {Timestamp: "2024303091615", SpacecraftNumber: "0282"},
	86: {Timestamp: "2024338090253", SpacecraftNumber: "0283"},
	87: {Timestamp: "2024356001253", SpacecraftNumber: "0284"},
	88: {Timestamp: "2025020080611", SpacecraftNumber: "0285"},
	89: {Timestamp: "2025056232930", SpacecraftNumber: "0286"},
	90: {Timestamp: "2025092103609", SpacecraftNumber: "0287"},
	91: {Timestamp: "2025099202608", SpacecraftNumber: "0288"},
	92: {Timestamp: "2025132143247", SpacecraftNumber: "0289"},
	93: {Timestamp: "2025156185246", SpacecraftNumber: "0290"},
	94: {Timestamp: "2025181090925", SpacecraftNumber: "0291"},
	95: {Timestamp: "2025212234604", SpacecraftNumber: "0292"},
	96: {Timestamp: "2025234122603", SpacecraftNumber: "0293"},
	97: {Timestamp: "2025258082922", SpacecraftNumber: "0294"},
}
