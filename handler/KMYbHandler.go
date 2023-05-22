package handler

import (
	"bufio"
	"bytes"
	"fmt"
	"syscall"
	"unsafe"
)

// 加载 CHSInterfaceYn.dll
var dll = syscall.NewLazyDLL("D:\\医保\\pack\\CHSInterfaceYn.dll")

var fixmedins_code_str = "H53011200683"
var infosyscode_str = "democase1"
var infosyssign_str = "eef17edbd8b411eb81950242ac13000d"
var url_str = "http://ldjk.yn.hsip.gov.cn/eapdomain/callService"

func Init() {
	// 获取 Init 函数地址
	initFunc := dll.NewProc("Init")

	// 准备参数
	fixmedins_code, _ := syscall.UTF16PtrFromString(fixmedins_code_str)
	infosyscode, _ := syscall.UTF16PtrFromString(infosyscode_str)
	infosyssign, _ := syscall.UTF16PtrFromString(infosyssign_str)
	url, _ := syscall.UTF16PtrFromString(url_str)
	var pErrMsg [256]byte

	// 调用 Init 函数
	r, _, err := initFunc.Call(
		uintptr(unsafe.Pointer(fixmedins_code)),
		uintptr(unsafe.Pointer(infosyscode)),
		uintptr(unsafe.Pointer(infosyssign)),
		uintptr(unsafe.Pointer(url)),
		uintptr(unsafe.Pointer(&pErrMsg[0])),
	)

	// 检查调用结果
	if r != 0 {
		fmt.Println("Init failed:", err)
		return
	}
	fmt.Println("Init success")
}

func BusinessHandleW() {
	// 获取 Init 函数地址
	businessHandleW := dll.NewProc("BusinessHandleW")

	inputData_str := "{\"infno\":\"1101\",\"msgid\":\"20210626000000001\",\"mdtrtarea_admvs\":\"530000\",\"insuplc_admdvs\":\"530000\",\"recer_sys_code\":\"mbs\",\"dev_no\":\"\",\"dev_safe_info\":\"\",\"cainfo\":\"\",\"signtype\":\"\",\"infver\":\"1.0\",\"opter_type\":\"1\",\"opter\":\"test\",\"opter_name\":\"test\",\"inf_time\":\"2021-06-26 00:00:00\",\"fixmedins_code\":\"H53011200683\",\"fixmedins_name\":\"云南省第一人民医院\",\"sign_no\":\"\",\"input\":{\"data\":{\"mdtrt_cert_type\":\"02\",\"mdtrt_cert_type3\":\"02\",\"mdtrt_cert_no\":\"530322199206242216\",\"card_sn\":\"\",\"begntime\":\"2021-05-26\",\"psn_cert_type\":\"\",\"certno\":\"\",\"psn_name\":\"\"}}}\n"

	// 准备参数
	fixmedins_code, _ := syscall.UTF16PtrFromString(fixmedins_code_str)
	infosyscode, _ := syscall.UTF16PtrFromString(infosyscode_str)
	infosyssign, _ := syscall.UTF16PtrFromString(infosyssign_str)
	inputData, _ := syscall.UTF16PtrFromString(inputData_str)
	const respDataLen = 1024 * 1
	var outputData [respDataLen]byte
	var pErrMsg [256]byte

	r, _, err := businessHandleW.Call(
		uintptr(unsafe.Pointer(fixmedins_code)),
		uintptr(unsafe.Pointer(infosyscode)),
		uintptr(unsafe.Pointer(infosyssign)),
		uintptr(unsafe.Pointer(inputData)),
		uintptr(unsafe.Pointer(&outputData[0])),
		uintptr(unsafe.Pointer(&pErrMsg[0])),
	)

	// 检查调用结果
	if r != 0 {
		fmt.Println("BusinessHandleW failed:", err)
		return
	}
	// 假设 outputData 包含有效的 UTF-8 编码字符串
	buf := bytes.NewBuffer(outputData[:])
	reader := bufio.NewReader(buf)
	outputString, err := reader.ReadString(0)
	if err != nil {
		// 处理错误
	}
	outputString = outputString[:len(outputString)-1]
	fmt.Println("outputData", outputString)
	fmt.Println("BusinessHandleW success")
}
