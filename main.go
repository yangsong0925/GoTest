package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	// 加载 CHSInterfaceYn.dll
	dll := syscall.NewLazyDLL("D:\\医保\\pack\\CHSInterfaceYn.dll")

	// 获取 Init 函数地址
	initFunc := dll.NewProc("Init")

	// 准备参数
	fixmedins_code := "fixmedins_code"
	infosyscode := "infosyscode"
	infosyssign := "infosyssign"
	url := "url"
	var pErrMsg [256]byte

	// 调用 Init 函数
	r, _, err := initFunc.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(fixmedins_code))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(infosyscode))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(infosyssign))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(url))),
		uintptr(unsafe.Pointer(&pErrMsg[0])),
	)

	// 检查调用结果
	if r == 0 {
		fmt.Println("Init failed:", err)
		return
	}

	fmt.Println("Init success")
}
