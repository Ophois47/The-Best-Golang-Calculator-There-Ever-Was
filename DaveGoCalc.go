package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/briandowns/spinner"
)

func add(num1 float32, num2 float32) {
	var sum float32
	sum = num1 + num2
	fmt.Println(num1, "+", num2, "=", sum)
}

func subtract(num1 float32, num2 float32) {
	var remainder float32
	remainder = num1 - num2
	fmt.Println(num1, "-", num2, "=", remainder)
}

func multiply(num1 float32, num2 float32) {
	var product float32
	product = num1 * num2
	fmt.Println(num1, "*", num2, "=", product)
}

func divide(num1 float32, num2 float32) {
	if num1 == 0 || num2 == 0 {
		fmt.Println("You KNOW you can't divide by zero!\nAmateur...")
		playAgain()
	}

	var quotient float32
	quotient = num1 / num2
	fmt.Println(num1, "/", num2, "=", quotient)
}

func celfar(num1 float32) {
	var farenheit float32
	farenheit = (num1/5)*9 + 32

	s := fmt.Sprintf("%.1f", farenheit)
	print("\033[H\033[2J")
	fmt.Println(num1, "degrees Celsius is equal to", s, "degrees Farenheit.")
}

func compoundInterest(principle float64, rate float64, time float64) {
	var compoundInterest float64
	compoundInterest = principle * (math.Pow((1 + rate/100), time))

	c := fmt.Sprintf("The compound interest for $%.2f", principle)
	p := fmt.Sprintf("at %.1f", rate)
	pp := fmt.Sprint("%")
	r := fmt.Sprintf(" for %.1f", time)
	t := fmt.Sprintf("years, is: $%.2f", compoundInterest)

	fmt.Println(c, p, pp, r, t)
	fmt.Println()
}

func bankHeist() {
	clearSc()
	fmt.Println("*~~~Welcome To Your Bank Heist~~~*")
	fmt.Println("This time:")
	rand.Seed(time.Now().UnixNano())
	var isHeistOn = true
	var eludedGuards = rand.Intn(100)
	var playAgain int

	if eludedGuards >= 50 {
		fmt.Printf("Looks like you've managed to make it past the guards.")
		fmt.Println(" Good job, but remember, this is only the first step.")
	} else {
		isHeistOn = false
		fmt.Println("You should plan a better disguise next time. You didn't even get in the door.")
	}

	var openedVault = rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grabbed the loot and ran!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("The vault could not be opened.")
	}

	var leftSafely = rand.Intn(5)

	if isHeistOn == true {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("The Heist failed due to natural causes.")
		case 1:
			isHeistOn = false
			fmt.Println("Suffered a negligent discharge, go back 4 spaces.")
		case 2:
			isHeistOn = false
			fmt.Println("You were shot.")
		case 3:
			isHeistOn = false
			fmt.Println("You were stabbed.")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		var amStolen = 10000 + rand.Intn(1000000)
		var moneyGot string

		moneyGot = fmt.Sprintf("You took $%v", amStolen)
		fmt.Println(moneyGot)
	}

	fmt.Println("Did the Heist Succeed? = ", isHeistOn)
	fmt.Println()
	fmt.Println()
	fmt.Println("(1) Try again? (2) Back to Main Program, or (3) Quit? ")
	fmt.Scan(&playAgain)

	if playAgain == 1 {
		bankHeist()
	} else if playAgain == 2 {
		main()
	} else if playAgain == 3 {
		return
	} else {
		fmt.Println("Only '1', '2', and '3' are accepted inputs...")
		quitCalc()
	}
}

func playAgain() {
	var playAgain int
	fmt.Println()
	fmt.Println("(1) Back to Main Program, or (2) Quit? ")
	fmt.Scan(&playAgain)

	if playAgain == 1 {
		main()
	} else if playAgain == 2 {
		return
	} else {
		fmt.Println("Only '1' and '2' are accepted inputs...")
		quitCalc()
	}
}

func quitCalc() {
	var leave string
	fmt.Println()
	fmt.Println()
	fmt.Println("...Type Continue and press ENTER to Continue\nOr type Quit and press ENTER to Quit...")
	fmt.Println()
	fmt.Scan(&leave)
	if leave == "continue" || leave == "CONTINUE" || leave == "Continue" {
		main()
	} else if leave == "quit" || leave == "QUIT" || leave == "Quit" {
		return
	} else {
		print("\033[H\033[2J")
		fmt.Println("You don't read too good, son...")
		fmt.Println()
		return
	}
}

func insertionSort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	fmt.Println("After Sorting:")
	fmt.Print("[")
	for _, val := range arr {
		fmt.Print(" ", val, " ")
	}
	fmt.Print("]")
}

func fibonacci(num int) {
	var a, b int
	for a, b = 0, 1; a <= num; a, b = b, a+b {
		fmt.Printf(" %d", a)
	}
}

func portScan(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)
	if err != nil {
		return false
	}

	defer conn.Close()
	return true
}

func portsScanner() {
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	for i := 1; i < 65535; i++ {
		port := strconv.FormatInt(int64(i), 10)
		fmt.Println()
		msg := fmt.Sprintf("Scanning Port %v...", port)
		fmt.Println(msg)
		s.Color("bgBlack", "bold", "fgHiYellow")
		s.FinalMSG = "Port Closed.\n"
		s.Start()
		conn, err := net.Dial("tcp", "192.168.1.214:"+port)
		if err == nil {
			mess := fmt.Sprintln("***~Port", i, "is OPEN~***")
			s.FinalMSG = mess
			s.Stop()
			conn.Close()
		}

		s.Stop()
	}
}

func isOpen(host string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}

	return false
}

func betterPortScanner() {
	hostname := flag.String("hostname", "192.168.1.214", "hostname to test")
	startPort := flag.Int("start-port", 80, "the port on which the scanning starts")
	endPort := flag.Int("end-port", 100, "the port from which the scanning ends")
	timeout := flag.Duration("timeout", time.Millisecond*200, "timeout")
	flag.Parse()

	ports := []int{}

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, p, *timeout)
			if opened {
				mutex.Lock()
				ports = append(ports, p)
				mutex.Unlock()
			}

			wg.Done()
		}(port)
	}

	wg.Wait()
	fmt.Printf("Opened Ports: %v\n", ports)
}

func multTable() {
	var n int
	print("\033[H\033[2J")
	fmt.Print("Enter an Integer: ")
	fmt.Scan(&n)
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(n, " X ", i, " = ", n*i)
		i++
	}
}

func avgArray() {
	var num [100]int
	var sum, temp int
	print("\033[H\033[2J")
	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&temp)
	for i := 0; i < temp; i++ {
		fmt.Print("Enter an element: ")
		fmt.Scan(&num[i])
		sum += num[i]
	}

	avg := (float64(sum)) / (float64(temp))
	goodAvg := float64(avg)
	fmt.Printf("The average amount of those %d number(s), is %.1f", temp, goodAvg)
}

func standardDeviation() {
	var num [10]float64
	var sum, mean, sd float64
	fmt.Printf("***Enter 10 Elements***")
	fmt.Println()
	for i := 1; i <= 10; i++ {
		fmt.Printf("Enter Element %d: ", i)
		fmt.Scan(&num[i-1])
		sum += num[i-1]
	}

	mean = sum / 10

	for j := 0; j < 10; j++ {
		sd += math.Pow(num[j]-mean, 2)
	}

	sd = math.Sqrt(sd / 10)

	fmt.Println("The Standard Deviation is: ", sd)
}

func cuboidDraw(drawX, drawY, drawZ int) {
	fmt.Printf("Cuboid %d %d %d:\n", drawX, drawY, drawZ)
	cubeLine(drawY+1, drawX, 0, "+-")
	for i := 1; i <= drawY; i++ {
		cubeLine(drawY-i+1, drawX, i-1, "/ |")
	}

	cubeLine(0, drawX, drawY, "+-|")
	for i := 4*drawZ - drawY - 2; i > 0; i-- {
		cubeLine(0, drawX, drawY, "| |")
	}

	cubeLine(0, drawX, drawY, "| +")
	for i := 1; i <= drawY; i++ {
		cubeLine(0, drawX, drawY-i, "| /")
	}

	cubeLine(0, drawX, 0, "+-\n")
}

func cubeLine(n, drawX, drawY int, cubeDraw string) {
	fmt.Printf("%*s", n+1, cubeDraw[:1])
	for d := 9*drawX - 1; d > 0; d-- {
		fmt.Print(cubeDraw[1:2])
	}

	fmt.Print(cubeDraw[:1])
	fmt.Printf("%*s\n", drawY+1, cubeDraw[2:])
}

func addMatrices() {
	var matrix1 [100][100]int
	var matrix2 [100][100]int
	var sum [100][100]int
	var row, col int
	fmt.Printf("Enter Number of Rows: ")
	fmt.Scan(&row)
	fmt.Printf("Enter Number of Columns: ")
	fmt.Scan(&col)

	fmt.Println()
	fmt.Println("========== Matrix1 =============")
	fmt.Println()
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter element for Matrix1 %d %d :", i+1, j+1)
			fmt.Scan(&matrix1[i][j])
		}
	}

	fmt.Println()
	fmt.Println("========== Matrix2 =============")
	fmt.Println()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("Enter element for Matrix2 %d %d :", i+1, j+1)
			fmt.Scan(&matrix2[i][j])
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sum[i][j] = matrix1[i][j] + matrix2[i][j]
		}
	}

	fmt.Println()
	fmt.Println("========== Sum of Matrix =============")
	fmt.Println()

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ", sum[i][j])
			if j == col-1 {
				fmt.Println("")
			}
		}
	}
}

func clearSc() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	clearSc()
	var choice int
	currentTime := time.Now()

	fmt.Println(currentTime.Format("January-02-2006"))
	fmt.Println(currentTime.Format("03:04:05"))
	fmt.Println("Logical Processors Available: ", runtime.NumCPU())
	fmt.Println()
	fmt.Println("******************************************")
	fmt.Println("* Welcome to David's Golang Calculator!! *")
	fmt.Println("******************************************")
	fmt.Println()
	fmt.Println("What Operation to Perform ~")
	fmt.Println("1.Addition")
	fmt.Println("2.Subtraction")
	fmt.Println("3.Multiplication")
	fmt.Println("4.Division")
	fmt.Println("5.Convert Celsius to Farenheit")
	fmt.Println("6.Calculate Compound Interest")
	fmt.Println("7.Try Your Luck at Randomized Bank Heists")
	fmt.Println("8.Insertion Sort")
	fmt.Println("9.Run Fibonacci Sequence")
	fmt.Println("10.Scan A Port")
	fmt.Println("11.Scan All Ports")
	fmt.Println("12.Print Multiplication Table")
	fmt.Println("13.Get The Average Of An Array")
	fmt.Println("14.Get Standard Deviation Of An Array")
	fmt.Println("15.Print Cuboid")
	fmt.Println("16.Add Two Matrices Using Multi-Dimensional Arrays")
	fmt.Println("17.Hopefully A Better Port Scanner")
	fmt.Println("18.Quit")
	fmt.Println()
	fmt.Printf("Input: ")
	fmt.Scanf("%d", &choice)

	switch choice {
	case 1:
		clearSc()
		fmt.Println("Addition!")
		var num1 float32
		fmt.Println("Begin with a value: ")
		fmt.Scan(&num1)
		var num2 float32
		fmt.Println("Add how much to that amount? ")
		fmt.Scan(&num2)

		add(num1, num2)
		playAgain()

	case 2:
		clearSc()
		fmt.Println("Subtraction!")
		var num1 float32
		fmt.Println("Begin with a value: ")
		fmt.Scan(&num1)
		var num2 float32
		fmt.Println("How much to subtract? ")
		fmt.Scan(&num2)

		subtract(num1, num2)
		playAgain()

	case 3:
		clearSc()
		fmt.Println("Multiplication!")
		var num1 float32
		fmt.Println("Begin with a value: ")
		fmt.Scan(&num1)
		var num2 float32
		fmt.Println("Multiply that by how much? ")
		fmt.Scan(&num2)

		multiply(num1, num2)
		playAgain()

	case 4:
		clearSc()
		fmt.Println("Division!")
		var num1 float32
		fmt.Println("Begin with a value: ")
		fmt.Scan(&num1)
		var num2 float32
		fmt.Println("Divide that by how much? ")
		fmt.Scan(&num2)

		divide(num1, num2)
		playAgain()

	case 5:
		clearSc()
		fmt.Println("Convert Celsius to Farenheit!")
		var num1 float32
		fmt.Println("What is the temperature in Celsius? ")
		fmt.Scan(&num1)

		celfar(num1)
		playAgain()

	case 6:
		clearSc()
		var principle float64
		var rate float64
		var time float64
		fmt.Println("Compound Interest Calculator!")
		fmt.Printf("Begin by entering a starting amount: $")
		fmt.Scan(&principle)
		fmt.Printf("Input the interest rate: ")
		fmt.Scan(&rate)
		fmt.Printf("For how long (in years)? ")
		fmt.Scan(&time)

		compoundInterest(principle, rate, time)
		playAgain()

	case 7:
		bankHeist()
		quitCalc()

	case 8:
		clearSc()
		length := 0
		fmt.Println("*~~~Welcome to Insertion Sort~~~*")
		fmt.Printf("Input the amount of values to sort: ")
		fmt.Scan(&length)
		fmt.Printf("Please enter a list of numbers, each seperated by a space: ")
		numbers := make([]int, length)
		for i := 0; i < length; i++ {
			fmt.Scan(&numbers[i])
		}

		fmt.Println(numbers)
		insertionSort(numbers)
		playAgain()

	case 9:
		clearSc()
		var num int
		fmt.Println("Welcome to Fibonacci Time!")
		fmt.Print("How far to Fibonacci to? ")
		fmt.Scan(&num)
		fibonacci(num)
		playAgain()

	case 10:
		clearSc()
		var port int
		fmt.Println("Input Port to Scan: ")
		fmt.Scan(&port)
		fmt.Printf("Scanning Port %v...", port)
		open := portScan("tcp", "localhost", port)
		if open == true {
			fmt.Printf("Port %v is Open.", port)
		} else {
			fmt.Printf("Port %v is Closed.", port)
		}
		playAgain()

	case 11:
		clearSc()
		fmt.Println("Scanning All Ports...")
		fmt.Println("(This Could Take A While)")
		fmt.Println("(Press Ctrl+C to Exit)")
		portsScanner()
		playAgain()

	case 12:
		clearSc()
		fmt.Println("*~~~Multiplication Tables~~~*")
		multTable()
		playAgain()

	case 13:
		clearSc()
		fmt.Println("*~~~Find the Average Amount From an Array~~~*")
		avgArray()
		playAgain()

	case 14:
		clearSc()
		fmt.Println("*~~~Find Standard Deviation of an Array~~~*")
		standardDeviation()
		playAgain()

	case 15:
		clearSc()
		fmt.Println("*~~~Print Cuboid to Screen~~~*")
		fmt.Println("*** Enter 3 Dimensions of Cuboid ***")

		var l, b, h int
		fmt.Printf("Enter Length: ")
		fmt.Scan(&l)
		fmt.Printf("Enter Base: ")
		fmt.Scan(&b)
		fmt.Printf("Enter Height: ")
		fmt.Scan(&h)
		cuboidDraw(l, b, h)
		playAgain()

	case 16:
		clearSc()
		fmt.Println(" *~~~Add Two Matrices Using~~~*")
		fmt.Println("*~~~Multi-Dimensional Arrays~~~*")
		addMatrices()
		playAgain()

	case 17:
		clearSc()
		fmt.Println("*~~~Hopefully This is a Better Port Scanner~~~*")
		betterPortScanner()
		playAgain()

	case 18:
		clearSc()
		return

	default:
		clearSc()
		main()
	}
}
