package cmd

import "os"

func Execute() error {
	if len(os.Args) >= 3 && os.Args[1] == "sys" && os.Args[2] == "info" {
		return SysInfo()
	}
	// per ora non fa nulla
	return nil
}

//gestito da Wide
//root minimale di gestione error per ora
