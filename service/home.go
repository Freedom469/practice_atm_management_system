package service

import "fmt"
import "os"
import "practice/utils"

func HomePage(account, pin, userName string ) {
	utils.ClearScreen()
	exit := false
	
	balance := GetUserBalance(account, pin)
	
		for !exit {
			fmt.Println("\n################################################################")
			fmt.Printf("Hello %s!\t\t\t\tBalance: Kshs %d\n", userName, balance)
			fmt.Println("\n\t\tWelcome To Our Atm Management System...")
			fmt.Println("\t\tChoose a service ...")
			fmt.Println(`
			1: Transfer Funds.
			2: Withdraw Funds.
			0: Exit.

------------------------------------------------------------------
					`)
	
			var input int
			fmt.Print("Enter Your Service of Choice here: ")
			_, err := fmt.Scan(&input)
	
			if err != nil {
				fmt.Println(err)
			}
			// fmt.Println("\n################################################################")
	
			switch input {
			case 1:
				utils.ClearScreen()
				TransferFunds(account, &balance)
			case 2:
				utils.ClearScreen()
				WithdrawFunds(account)
			case 0:
				fmt.Println("Exiting ...")
				exit = true
				os.Exit(0)
			}
		}
}