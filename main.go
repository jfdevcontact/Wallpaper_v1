package main

import(
	"fmt"
	"os/exec"
	
)


func main(){

	//EU DO FUTURO. LEMBRE-SE DE TERMINA ESSE PROJETO

	//DATA DO INICIO DO PROJETO 17/10/2023
	
	md := exec.Command("bash", "-c", "gsettings set org.gnome.desktop.background picture-uri file:////tela.jpeg")

	out,_ := md.Output()
	
	fmt.Println(string(out))


}