package main

import (
	"backend/business"
)

func main() {
	business.PrintLog("========begin task=========")
	business.RunDevice()
	business.PrintLog("========end task=========")
}
