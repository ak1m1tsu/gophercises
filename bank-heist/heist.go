package bankheist

import (
	"fmt"
	"math/rand"
	"time"
)

func RunBankHeist() {
	rand.Seed(time.Now().UnixNano())

	isHeistOn := true

	if eludedGuards := rand.Intn(100); eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards.")
		fmt.Println("Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	if openedVault := rand.Intn(100); isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("What's the combo to this lock again?")
	}

	if isHeistOn {
		switch leftSafely := rand.Intn(5); leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm... run?")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("When did they start raising dogs in vaults??")
		case 3:
			isHeistOn = false
			fmt.Println("Did I even pack the burlap bags?")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Printf("$%v not bad\n", amtStolen)
	}
}
