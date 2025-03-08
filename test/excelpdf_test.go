package test

import (
	"fmt"
	"testing"

	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
)

func GenerateExcel() (string, error) {
	f := excelize.NewFile()
	fileName := "output13.xlsx"

	// Menambahkan data ke dalam file Excel
	f.SetCellValue("Sheet1", "A1", "NO : B.1671-MMS-FST/OPD/MIM/01/2025")
	f.SetCellValue("Sheet1", "A2", "Kepada Yth : IBU JUWITA DAN ADMIN PCS (PT. PAS)")
	f.SetCellValue("Sheet1", "A3", "Dari : Merchant Implementation Management")
	f.SetCellValue("Sheet1", "A5", "Lampiran :")
	f.SetCellValue("Sheet1", "A6", "No")
	f.SetCellValue("Sheet1", "B6", "Jumlah")
	f.SetCellValue("Sheet1", "C6", "Macam Yang Dikirim")
	f.SetCellValue("Sheet1", "D6", "Keterangan")

	f.SetColWidth("Sheet1", "A", "A", 8.33)
	f.SetColWidth("Sheet1", "B", "B", 31.22)
	f.SetColWidth("Sheet1", "C", "C", 80.56)
	f.SetColWidth("Sheet1", "D", "D", 81.67)

	// Contoh data isi tabel
	data := [][]any{
		{1, 10, "Barang A", "Keterangan A"},
		{2, 5, "Barang B", "Keterangan B"},
		{3, 15, "Barang C", "Keterangan C"},
	}

	for i, row := range data {
		rowNum := 0
		if i == 0 {
			rowNum = 7
		} else if i == 1 {
			rowNum = 9
		} else if i == 2 {
			rowNum = 11
		}

		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", rowNum), row[0])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", rowNum), row[1])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", rowNum), row[2])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", rowNum), row[3])
	}

	f.MergeCell("Sheet1", "A7", "A8")
	f.MergeCell("Sheet1", "A9", "A10")
	f.MergeCell("Sheet1", "A11", "A12")
	f.MergeCell("Sheet1", "B7", "B8")
	f.MergeCell("Sheet1", "B9", "B10")
	f.MergeCell("Sheet1", "B11", "B12")
	f.MergeCell("Sheet1", "C7", "C8")
	f.MergeCell("Sheet1", "C9", "C10")
	f.MergeCell("Sheet1", "C11", "C12")
	f.MergeCell("Sheet1", "D7", "D8")
	f.MergeCell("Sheet1", "D9", "D10")
	f.MergeCell("Sheet1", "D11", "D12")

	style := excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}

	headerStyle := excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	}

	setingStyle, err := f.NewStyle(&style)
	if err != nil {
		return "", err
	}
	f.SetCellStyle("Sheet1", "A6", "D12", setingStyle)

	boldStyleID, err := f.NewStyle(&headerStyle)
	if err != nil {
		return "", err
	}
	f.SetCellStyle("Sheet1", "A6", "D6", boldStyleID)
	// Simpan file Excel
	if err := f.SaveAs(fileName); err != nil {
		return "", err
	} else {
		fmt.Println("File Excel berhasil dibuat: output.xlsx")
	}

	return fileName, nil
}



// ConvertExcelToPDF membaca data dari file Excel dan menulisnya ke PDF

func ConvertExcelToPDF(excelFilePath, pdfFilePath string) error {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	// **Set Margin Kiri**
	marginLeft := 10.0
	pdf.SetLeftMargin(marginLeft)

	// **Set font untuk judul**
	pdf.SetFont("Times", "B", 14)

	// **Tambahkan Header Dokumen**
	title := "NO : B.1671-MMS-FST/OPD/MIM/01/2025\nKepada Yth : IBU JUWITA DAN ADMIN PCS (PT. PAS)\nDari : Merchant Implementation Management\n\n"
	pdf.SetX(marginLeft)
	pdf.MultiCell(0, 8, title, "", "L", false)

	// **Tambahkan Lampiran**
	pdf.SetFont("Times", "", 12)
	lampiran := "Lampiran :\n"
	pdf.SetX(marginLeft)
	pdf.MultiCell(0, 10, lampiran, "", "L", false)

	// **Header & Data Tabel**
	headers := []string{"NO", "Jumlah", "Macam Yang Dikirim", "Keterangan"}
	data := [][]string{
		{"1", "15 (lima belas)", "Soundbox Beserta Kelengkapannya", "Mohon dapat diimplementasikan/diaktivasi Soundbox pada Merchant sesuai data terlampir"},
		{"2", "1 (satu)", "Surat Izin Kerja", "Ditujukan kepada IBU JUWITA DAN ADMIN PCS (PT. PASIFIK CIPTA SOLUSI)"},
		{"3", "15 (lima belas)", "SIMCARD", "SIM CARD agar disesuaikan dengan kualitas jaringan di masing-masing Merchant"},
	}

	// **Hitung Lebar Kolom Otomatis**
	pageWidth := 297.0 // Lebar A4 Landscape
	totalPadding := 3.0 // Sisa ruang untuk border
	availableWidth := pageWidth - (marginLeft * 2) - totalPadding

	// **Tentukan lebar kolom secara proporsional**
	colWidths := make([]float64, len(headers))
	for i, header := range headers {
		colWidths[i] = pdf.GetStringWidth(header) + 10 // Lebar minimum dari header
	}
	for _, row := range data {
		for i, cell := range row {
			width := pdf.GetStringWidth(cell) + 10
			if width > colWidths[i] {
				colWidths[i] = width
			}
		}
	}

	// **Sesuaikan agar tidak melebihi halaman**
	totalColWidth := 0.0
	for _, w := range colWidths {
		totalColWidth += w
	}

	// Jika lebar melebihi, skala ulang semua kolom
	if totalColWidth > availableWidth {
		scaleFactor := availableWidth / totalColWidth
		for i := range colWidths {
			colWidths[i] *= scaleFactor
		}
	}

	// **Buat Header Tabel**
	pdf.SetFont("Arial", "B", 12)
	pdf.SetX(marginLeft)
	for i, header := range headers {
		pdf.CellFormat(colWidths[i], 10, header, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// **Buat Isi Tabel**
	pdf.SetFont("Arial", "", 10)
	for _, row := range data {
		pdf.SetX(marginLeft)
		for i, cell := range row {
			pdf.CellFormat(colWidths[i], 10, cell, "1", 0, "L", false, 0, "")
		}
		pdf.Ln(-1)
	}

	pdf.Text(marginLeft, pdf.GetY()+80, "Notes :")
	pdf.Text(marginLeft, pdf.GetY()+85, "Jumlah : Aktivasi di kota tersebut yang didapatkan dari Lampiran SIK Soundbox")

	// **Tambahkan Subtitle di Kanan Bawah**
	pdf.SetFont("Times", "B", 13)
	subTitle := "PT. BANK RAKYAT INDONESIA (PERSERO)Tbk.\nFUNDING & RETAIL PAYMENT STRATEGY DIVISION\nMerchant Implementation Management"
	pdf.SetXY(marginLeft+150, pdf.GetY()+20)
	pdf.MultiCell(0, 8, subTitle, "", "C", false)

	// **Simpan PDF**
	err := pdf.OutputFileAndClose(pdfFilePath)
	if err != nil {
		return fmt.Errorf("failed to save PDF file: %v", err)
	}

	return nil

}

func TestConvertPdf(t *testing.T) {
	excelFile := "output13.xlsx"
	err := ConvertExcelToPDF(excelFile, "output13.pdf")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("File PDF berhasil dibuat: output.pdf")
}


