package main

/*
 typedef enum ProxyType {
   HTTPS,
   SOCKS4,
   SOCKS5,
 } ProxyType;

 typedef enum ResultType {
   RETRY,
   HIT,
   BAD,
 } ResultType;

typedef struct ResultCheck {
    enum ResultType result_type;
    const char *extra_info;
} ResultCheck;
*/
import "C"

func main() {}

//export PLUGIN_NAME
func PLUGIN_NAME() *C.char {
	return C.CString("name_in_go")
}

//export PLUGIN_VERSION
func PLUGIN_VERSION() *C.char {
	return C.CString("0.0.1")
}

//export get_combo_result
func get_combo_result(user *C.char, pass *C.char, proxy *C.char, proxyType C.ProxyType) C.ResultCheck {
	return C.ResultCheck{C.BAD, C.CString("Add what you want to be saved on the file additionally!")}
}
