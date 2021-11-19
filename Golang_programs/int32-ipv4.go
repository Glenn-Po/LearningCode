// int32-ipv4
package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
int32 to IPv4

Take the following IPv4 address: 128.32.10.1

This address has 4 octets where each octet is a single byte (or 8 bits).

1st octet 128 has the binary representation: 10000000
2nd octet 32 has the binary representation: 00100000
3rd octet 10 has the binary representation: 00001010
4th octet 1 has the binary representation: 00000001
So 128.32.10.1 == 10000000.00100000.00001010.00000001

Because the above IP address has 32 bits, we can represent it as the unsigned 32 bit number: 2149583361

Complete the function that takes an unsigned 32 bit number and returns a string representation of its IPv4 address.

sahglie
*/

func main() {
	fmt.Println("=========================")
	fmt.Printf("\n%v =  %v", uint32(2149588938), int32ToIpV4(2149588938)) //2149583361
	addr := "255.72.05.45"
	intFmt := ipV4ToInt32(addr)
	fmt.Printf("\n%v =  %v", addr, intFmt)
	fmt.Printf("\n%v =?=  %v", addr, int32ToIpV4(intFmt))
}

func int32ToIpV4(ip uint32) string {
	var IpV4_Addr string
	cnt := 0
	tmpHolder := ""
	for x := 0; x <= 32; x++ {
		tmpHolder = intToString(ip&1) + tmpHolder
		ip >>= 1
		cnt++
		if cnt == 8 {
			cnt = 0
			IpV4_Addr = intToString(binToDec(tmpHolder)) + "." + IpV4_Addr
			tmpHolder = ""
		}
	}
	return IpV4_Addr[:len(IpV4_Addr)-1]
}

func binToDec(binString string) uint32 {
	var intValue uint32 = 0
	ln := len(binString) - 1
	for pos, val := range binString {
		intValue += (uint32(val) - 48) * (1 << (ln - pos))
	}
	return intValue
}

func intToString(val uint32) string {
	//as easy as this, knowlegde is the key
	//return fmt.Sprintf("%02d", val)
	if val == 0 {
		return "0"
	}
	strNum := ""
	for val != 0 {
		strNum = string(val%10+48) + strNum //'0'+number
		val /= 10
	}
	return strNum
}

func ipV4ToInt32(ipv4Adrr string) (intIp uint32) {
	//split by .
	var hostNetId []string = strings.Split(ipv4Adrr, ".")

	//contatenate binary rep of the numbers
	bitString := ""
	for _, numPart := range hostNetId {
		num, _ := strconv.ParseUint(numPart, 10, 32)
		bitString += fmt.Sprintf("%08b", num)
	}

	//convert the bits string to int32

	intIp = binToDec(bitString)

	return
}
