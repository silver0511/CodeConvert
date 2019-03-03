package main

import(
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strings"
	"os"
	"io"
)

func main()  {
	data, err := ioutil.ReadFile("packdef.json")
	if err != nil {
		fmt.Printf("load packdef.json err=%s", err)
		return
	}

	data_json := []byte(data)
	var pack_def_map map[string][]string
	err = json.Unmarshal(data_json, &pack_def_map)

	var sort_array = []string{"base", "sys", "user", "chat", "friend", "group", "offcial", "business"}
	fmt.Println("==========================start convert packdef==========================")
	fmt.Printf("PackDef=%v\n",pack_def_map)
	GenHPackTypeDef(sort_array, pack_def_map)
	GenJSPackTypeDef(sort_array, pack_def_map)
	GenJavaPackTypeDef(sort_array, pack_def_map)
	fmt.Println("==========================end convert packdef==========================")

	data, err = ioutil.ReadFile("errorcode.json")
	if err != nil {
		fmt.Printf("load errorcode.json err=%s", err)
		return
	}

	data_json = []byte(data)
	pack_def_map = nil
	err = json.Unmarshal(data_json, &pack_def_map)
	fmt.Println("==========================start convert errorcode==========================")
	fmt.Printf("ErrorDef=%v\n",pack_def_map)
	GenHErrorCode(sort_array, pack_def_map)
	GenJSErrorCode(sort_array, pack_def_map)
	GenJavaErrorCode(sort_array, pack_def_map)

	GenMErrorDetail(sort_array, pack_def_map)
	GenJSErrorDetail(sort_array, pack_def_map)
	GenJavaErrorDetail(sort_array, pack_def_map)
	fmt.Println("==========================end convert errorcode==========================")
}

func FileIsExists(filename string) (bool) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func GenHPackTypeDef(sort_array []string, pack_def_map map[string][]string)  {
	var msg_pack_def string
	msg_pack_def = ""
	msg_pack_def += "#ifndef __PACKTYPEDEF_H\n#define __PACKTYPEDEF_H\n\n"
	for _, sort_key := range sort_array {
		val_list := pack_def_map[sort_key]
		if nil == val_list {
			continue
		}

		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " start" + "//////////////////////////////////////\n"
		} else {
			msg_pack_def += "\n"
		}
		for _, value := range val_list {
			msg_array := strings.Split(value, "|")
			var def_note string = ""
			var def_content string = ""
			m1_len := 52 - len([]rune(msg_array[0]))
			if m1_len <= 0 {
				m1_len = 1
			}
			def_content += "#define " + msg_array[0]
			for m_i := 0; m_i < m1_len; m_i++  {
				def_content += " "
			}

			def_content +=  msg_array[1] + "\n"
			if len(msg_array) > 2 && msg_array[2] != "" {
				def_note += "//" + msg_array[2] + "\n"
				msg_pack_def += def_note
			}

			msg_pack_def += def_content
		}
		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " end" + "///////////////////////////////////////\n\n"
		} else {
			msg_pack_def += "\n"
		}
	}

	msg_pack_def += "\n\n\n#endif //__PACKTYPEDEF_H"
	var filename = "./ccode/PackTypeDef.h"
	if FileIsExists(filename) {
		os.Remove(filename)
	}

	file_packdef, _ := os.Create(filename)
	io.WriteString(file_packdef, msg_pack_def)
	defer file_packdef.Close()
}

func GenJSPackTypeDef(sort_array []string, pack_def_map map[string][]string)  {
	var msg_pack_def string
	msg_pack_def = ""
	for _, sort_key := range sort_array {
		val_list := pack_def_map[sort_key]
		if nil == val_list {
			continue
		}

		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " start" + "//////////////////////////////////////\n"
		} else {
			msg_pack_def += "\n"
		}
		for _, value := range val_list {
			msg_array := strings.Split(value, "|")
			var def_note string = ""
			var def_content string = ""
			m1_len := 57 - len([]rune(msg_array[0]))
			if m1_len <= 0 {
				m1_len = 1
			}
			def_content += "var " + msg_array[0]
			for m_i := 0; m_i < m1_len; m_i++  {
				def_content += " "
			}

			def_content +=  "= " + msg_array[1] + ";\n"
			if len(msg_array) > 2 && msg_array[2] != "" {
				def_note += "//" + msg_array[2] + "\n"
				msg_pack_def += def_note
			}

			msg_pack_def += def_content
		}
		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " end" + "///////////////////////////////////////\n\n"
		} else {
			msg_pack_def += "\n"
		}
	}

	var filename = "./jscode/PackTypeDef.js"
	if FileIsExists(filename) {
		os.Remove(filename)
	}

	file_packdef, _ := os.Create(filename)
	io.WriteString(file_packdef, msg_pack_def)
	defer file_packdef.Close()
}

func GenJavaPackTypeDef(sort_array []string, pack_def_map map[string][]string)  {
	var msg_pack_def string
	msg_pack_def = ""
	msg_pack_def += "package com.qbao.newim.constdef;\n\n\npublic interface PackTypeDef\n{\n    short\n"
	for _, sort_key := range sort_array {
		val_list := pack_def_map[sort_key]
		if nil == val_list {
			continue
		}

		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " start" + "//////////////////////////////////////\n"
		} else {
			msg_pack_def += "\n"
		}
		for _, value := range val_list {
			msg_array := strings.Split(value, "|")
			var def_note string = ""
			var def_content string = ""
			m1_len := 52 - len([]rune(msg_array[0]))
			if m1_len <= 0 {
				m1_len = 1
			}
			def_content += "    " + msg_array[0]
			for m_i := 0; m_i < m1_len; m_i++  {
				def_content += " "
			}

			def_content +=  "= " + msg_array[1] + ",\n"
			if len(msg_array) > 2 && msg_array[2] != "" {
				def_note += "    //" + msg_array[2] + "\n"
				msg_pack_def += def_note
			}

			msg_pack_def += def_content
		}
		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " end" + "///////////////////////////////////////\n\n"
		} else {
			msg_pack_def += "\n"
		}
	}

	msg_pack_def += "    END_OF_TYPE = 19999;\n}\n"
	var filename = "./javacode/PackTypeDef.java"
	if FileIsExists(filename) {
		os.Remove(filename)
	}

	file_packdef, _ := os.Create(filename)
	io.WriteString(file_packdef, msg_pack_def)
	defer file_packdef.Close()
}

func GenHErrorCode(sort_array []string, pack_def_map map[string][]string) {
	var msg_pack_def string
	msg_pack_def = ""
	msg_pack_def += "#ifndef __ERRORCODEDEF_H_\n#define __ERRORCODEDEF_H_\n\n"
	for _, sort_key := range sort_array {
		val_list := pack_def_map[sort_key]
		if nil == val_list {
			continue
		}

		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " start" + "//////////////////////////////////////\n"
		} else {
			msg_pack_def += "\n"
		}
		for _, value := range val_list {
			msg_array := strings.Split(value, "|")
			var def_note string = ""
			var def_content string = ""
			m1_len := 52 - len([]rune(msg_array[0]))
			if m1_len <= 0 {
				m1_len = 1
			}
			def_content += "#define " + msg_array[0]
			for m_i := 0; m_i < m1_len; m_i++  {
				def_content += " "
			}

			def_content +=  msg_array[1] + "\n"
			if len(msg_array) > 2 && msg_array[2] != "" {
				def_note += "//" + msg_array[2] + "\n"
				msg_pack_def += def_note
			}

			msg_pack_def += def_content
		}
		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " end" + "///////////////////////////////////////\n\n"
		} else {
			msg_pack_def += "\n"
		}
	}

	msg_pack_def += "\n\n\n#endif //__ERRORCODEDEF_H_"
	var filename = "./ccode/ErrorCodeDef.h"
	if FileIsExists(filename) {
		os.Remove(filename)
	}

	file_packdef, _ := os.Create(filename)
	io.WriteString(file_packdef, msg_pack_def)
	defer file_packdef.Close()
	//
}

func GenJSErrorCode(sort_array []string, pack_def_map map[string][]string) {
	var msg_pack_def string
	msg_pack_def = ""
	for _, sort_key := range sort_array {
		val_list := pack_def_map[sort_key]
		if nil == val_list {
			continue
		}

		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " start" + "//////////////////////////////////////\n"
		} else {
			msg_pack_def += "\n"
		}
		for _, value := range val_list {
			msg_array := strings.Split(value, "|")
			var def_note string = ""
			var def_content string = ""
			m1_len := 57 - len([]rune(msg_array[0]))
			if m1_len <= 0 {
				m1_len = 1
			}
			def_content += "var " + msg_array[0]
			for m_i := 0; m_i < m1_len; m_i++  {
				def_content += " "
			}

			def_content +=  "= " + msg_array[1] + ";\n"
			if len(msg_array) > 2 && msg_array[2] != "" {
				def_note += "//" + msg_array[2] + "\n"
				msg_pack_def += def_note
			}

			msg_pack_def += def_content
		}
		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " end" + "///////////////////////////////////////\n\n"
		} else {
			msg_pack_def += "\n"
		}
	}

	var filename = "./jscode/ErrorCodeDef.js"
	if FileIsExists(filename) {
		os.Remove(filename)
	}

	file_packdef, _ := os.Create(filename)
	io.WriteString(file_packdef, msg_pack_def)
	defer file_packdef.Close()
}

func GenJavaErrorCode(sort_array []string, pack_def_map map[string][]string) {
	var msg_pack_def string
	msg_pack_def = ""
	msg_pack_def += "package com.qbao.newim.constdef;\n\n\npublic interface ErrorCodeDef\n{\n    int\n"
	for _, sort_key := range sort_array {
		val_list := pack_def_map[sort_key]
		if nil == val_list {
			continue
		}

		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " start" + "//////////////////////////////////////\n"
		} else {
			msg_pack_def += "\n"
		}
		for _, value := range val_list {
			msg_array := strings.Split(value, "|")
			var def_note string = ""
			var def_content string = ""
			m1_len := 52 - len([]rune(msg_array[0]))
			if m1_len <= 0 {
				m1_len = 1
			}
			def_content += "    " + msg_array[0]
			for m_i := 0; m_i < m1_len; m_i++  {
				def_content += " "
			}

			def_content +=  "= " + msg_array[1] + ",\n"
			if len(msg_array) > 2 && msg_array[2] != "" {
				def_note += "    //" + msg_array[2] + "\n"
				msg_pack_def += def_note
			}

			msg_pack_def += def_content
		}
		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " end" + "///////////////////////////////////////\n\n"
		} else {
			msg_pack_def += "\n"
		}
	}

	msg_pack_def += "    END_OF_ERROR = -1;\n}\n"
	var filename = "./javacode/ErrorCodeDef.java"
	if FileIsExists(filename) {
		os.Remove(filename)
	}

	file_packdef, _ := os.Create(filename)
	io.WriteString(file_packdef, msg_pack_def)
	defer file_packdef.Close()
}


func GenMErrorDetail(sort_array []string, pack_def_map map[string][]string) {
	var msg_pack_def string
	msg_pack_def = ""
	msg_pack_def += "#import \"ErrorManager.h\"\n\n@implementation ErrorManager\nSingletonImplementation(ErrorManager)\n\n" +
		"- (id)init\n{\n    self = [super init];\n    err_msg_dic = [[NSMutableDictionary alloc]" +
		"initWithCapacity: 65535];\n    [self initErrorDetail];\n    return self;\n}\n\n" +
		"- (void)dealloc\n{\n}\n\n-(NSString*) getErrorDetail:(uint64_t) error_code\n{\n    NSString* err_detail = " +
		"[err_msg_dic objectForKey:[NSNumber numberWithUnsignedLongLong:error_code]];\n    if(err_detail == NULL)\n    {\n        err_detail = " +
		"[NSString stringWithFormat:@\"%@%lld\",@\"未知的错误信息:\", error_code];\n    }\n    return err_detail;\n}\n\n" +
		"-(void) setErrorDetail:(unsigned int) error_code detail:(NSString *)detail\n{\n" +
		"    [err_msg_dic setObject:detail forKey:[NSNumber numberWithUnsignedInt:error_code]];\n}\n\n" +
		"-(void) initErrorDetail\n{\n"
	for _, sort_key := range sort_array {
		val_list := pack_def_map[sort_key]
		if nil == val_list {
			continue
		}

		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " start" + "//////////////////////////////////////\n"
		} else {
			continue
		}
		for _, value := range val_list {
			msg_array := strings.Split(value, "|")
			if len(msg_array) < 3 ||
				msg_array[2] == "" {
				continue
			}

			var def_content string = ""
			def_content += "    [self setErrorDetail:" + msg_array[0]
			def_content +=  " detail:@\"" + msg_array[2] + "\"];\n"
			msg_pack_def += def_content
		}
		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " end" + "///////////////////////////////////////\n\n"
		} else {
			msg_pack_def += "\n"
		}
	}

	msg_pack_def += "}\n\n@end"
	var filename = "./ccode/ErrorManager.m"
	if FileIsExists(filename) {
		os.Remove(filename)
	}

	file_packdef, _ := os.Create(filename)
	io.WriteString(file_packdef, msg_pack_def)
	defer file_packdef.Close()
	//
}

func GenJSErrorDetail(sort_array []string, pack_def_map map[string][]string) {
	var msg_pack_def string
	msg_pack_def = ""
	msg_pack_def += "var ErrorDetailDef = {};\n"
	for _, sort_key := range sort_array {
		val_list := pack_def_map[sort_key]
		if nil == val_list {
			continue
		}

		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " start" + "//////////////////////////////////////\n"
		} else {
			continue
		}
		for _, value := range val_list {
			msg_array := strings.Split(value, "|")
			if len(msg_array) < 3 ||
				msg_array[2] == "" {
				continue
			}

			var def_content string = ""
			def_content += "ErrorDetailDef[" + msg_array[0] + "] = \"" + msg_array[2] + "\";\n"
			msg_pack_def += def_content
		}
		if sort_key != "base" {
			msg_pack_def += "//////////////////////////////////////" + sort_key + " end" + "///////////////////////////////////////\n\n"
		} else {
			msg_pack_def += "\n"
		}
	}

	var filename = "./jscode/ErrorDetailDef.js"
	if FileIsExists(filename) {
		os.Remove(filename)
	}

	file_packdef, _ := os.Create(filename)
	io.WriteString(file_packdef, msg_pack_def)
	defer file_packdef.Close()
}

func GenJavaErrorDetail(sort_array []string, pack_def_map map[string][]string) {
	var msg_pack_def string
	msg_pack_def = ""
	msg_pack_def += "package com.qbao.newim.util;\nimport com.qbao.newim.constdef.ErrorCodeDef;\n" +
		"import com.qbao.newim.qbim.R;\nimport java.util.HashMap;\n\npublic class ErrorDetail\n{\n" +
		"    private static HashMap<Integer, String> error_detail_map = new HashMap<>();\n" +
		"    public static String GetErrorDetail(int error_code)\n    {\n" +
		"        if(!error_detail_map.containsKey(error_code))\n        {\n" +
		"            return \"未知的错误信息:\" + error_code;\n        }\n        " +
		"return error_detail_map.get(error_code);\n    }\n\n    public static void Init()\n    {\n"
	for _, sort_key := range sort_array {
		val_list := pack_def_map[sort_key]
		if nil == val_list {
			continue
		}

		if sort_key != "base" {
			msg_pack_def += "        //////////////////////////////////////" + sort_key + " start" + "//////////////////////////////////////\n"
		} else {
			continue
		}
		for _, value := range val_list {
			msg_array := strings.Split(value, "|")
			if len(msg_array) < 3 ||
				msg_array[2] == "" {
				continue
			}

			var def_content string = ""
			def_content += "        error_detail_map.put(ErrorCodeDef." + msg_array[0] + ",\"" + msg_array[2] + "\");\n"
			msg_pack_def += def_content
		}
		if sort_key != "base" {
			msg_pack_def += "        //////////////////////////////////////" + sort_key + " end" + "///////////////////////////////////////\n\n"
		} else {
			msg_pack_def += "\n"
		}
	}

	msg_pack_def += "    }\n}\n"
	var filename = "./javacode/ErrorDetail.java"
	if FileIsExists(filename) {
		os.Remove(filename)
	}

	file_packdef, _ := os.Create(filename)
	io.WriteString(file_packdef, msg_pack_def)
	defer file_packdef.Close()
}