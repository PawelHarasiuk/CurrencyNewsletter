package services

import "fmt"

func SendReport() {
	// Currency data (you can replace these with actual data)
	currency1 := "USD"
	buyingRate1 := 1.1000
	sellingRate1 := 1.1050
	change1 := 0.02

	currency2 := "EUR"
	buyingRate2 := 1.2000
	sellingRate2 := 1.2050
	change2 := -0.01

	currency3 := "GBP"
	buyingRate3 := 1.3000
	sellingRate3 := 1.3050
	change3 := 0.03

	// Report string
	report := fmt.Sprintf(`Daily Currency Rate Report – %s

Currency      | Buying Rate  | Selling Rate | Change (%%)
----------------------------------------------------------
%-13s | %-12.4f | %-12.4f | %.2f
%-13s | %-12.4f | %-12.4f | %.2f
%-13s | %-12.4f | %-12.4f | %.2f

Key Notes:
- Market trends have shown slight fluctuations in exchange rates.
- Predicted to stabilize over the coming days.

Best regards,
[Your Name]
[Your Position]
[Your Company Name]`,
		"2024-10-06", // Date, you can automate this with time.Now()
		currency1, buyingRate1, sellingRate1, change1,
		currency2, buyingRate2, sellingRate2, change2,
		currency3, buyingRate3, sellingRate3, change3)

	// Output the report
	fmt.Println(report)
}
