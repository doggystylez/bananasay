package main

import "math/rand"

var bananas = []string{banana1, banana2, banana3}

const (
	banana1 = `
                        ████                                            
                      ████████                                          
                      ██    ██                                          
                      ██    ████                                        
                      ██      ████                                      
                      ██        ████                                    
                      ██          ██                                    
                  ██████████      ████                                  
                ██  ██      ██      ██                                  
                ████████    ██      ██                                  
                ██  ██      ██      ██                                  
                  ██████████    ██  ██                                  
                    ██      ██████  ██                                  
                    ████████▒▒▒▒██  ██                                  
                      ██▒▒▒▒▒▒██    ██                                  
                      ██▒▒████      ██                                  
                    ██████          ██                                  
        ██████      ██            ████  ██████                          
      ██    ████  ████            ██  ████    ██                        
      ██      ██  ██            ████  ██      ██                        
      ██    ████  ██            ██    ████    ██                        
        ██████  ████          ████    ▒▒██████                          
            ▒▒  ██          ████▒▒  ▒▒▒▒                                
            ▒▒▒▒████    ██████  ▒▒▒▒▒▒                                  
          ▒▒▒▒▒▒██████████▒▒▒▒▒▒▒▒▒▒                                    
        ██████                ▒▒██████                                  
      ██      ██████      ██████      ██                                
      ██          ██      ██          ██                                
        ▓▓▓▓████▓▓██      ██▓▓▓▓▓▓████                                  

`

	banana2 = `
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬛⬛⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛🟨🟨⬛⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛🟨🟨⬛⬛⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛🟨🟨🟨⬛⬛⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛🟨🟨🟨🟨⬛⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛🟨🟨🟨🟨⬛⬛⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛🟨🟨🟨🟨🟨⬛⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬛⬛⬛⬛⬛🟨🟨🟨⬛⬛⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬛⬜⬛⬜⬜⬜⬛🟨🟨🟨⬛⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬛⬛⬛⬛⬜⬜⬛🟨🟨🟨⬛⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬛⬜⬛⬜⬜⬜⬛🟨🟨🟨⬛⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬛⬛⬛⬛⬛🟨🟨⬛🟨⬛⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬛🟨🟨🟨⬛⬛⬛🟨⬛⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬛⬛🟫🟫⬛🟨⬛⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛🟫🟫🟫⬛🟨🟨⬛⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛🟫⬛⬛🟨🟨🟨⬛⬛⬛⬛⬜
⬜⬜⬛⬛⬛⬜⬜⬜⬛⬛⬛🟨🟨🟨🟨⬛⬛⬜⬜⬜⬛
⬜⬛⬜⬜⬜⬛⬛⬜⬛🟨🟨🟨🟨🟨🟨⬛⬜⬜⬜⬜⬛
⬜⬛⬜⬜⬜⬜⬛⬜⬛🟨🟨🟨🟨🟨⬛⬛⬛⬜⬜⬜⬛
⬜⬛⬜⬜⬜⬛⬛⬛🟨🟨🟨🟨🟨🟨⬛⬜⬛⬛⬛⬛⬜
⬜⬜⬛⬛⬛⬜⬜⬛🟨🟨🟨🟨🟨⬛⬜⬜⬜⬛⬜⬜⬜
⬜⬜⬜⬜⬛⬜⬛🟨🟨🟨🟨🟨⬛⬜⬜⬛⬛⬛⬜⬜⬜
⬜⬜⬜⬜⬛⬛🟨🟨🟨🟨🟨⬛⬛⬛⬛⬛⬜⬜⬜⬜⬜
⬜⬜⬜⬛⬛⬛⬛🟨🟨⬛⬛⬛⬛⬛⬛⬜⬜⬜⬜⬜⬜
⬜⬛⬛⬛⬛⬜⬛⬛⬛⬛⬜⬜⬜⬛⬛⬛⬛⬜⬜⬜⬜
⬛⬜⬜⬜⬛⬛⬜⬜⬜⬜⬜⬜⬛⬛⬜⬜⬜⬛⬜⬜⬜
⬛⬜⬜⬜⬜⬜⬛⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬛⬜⬜⬜
⬜⬛⬛⬛⬛⬛⬛⬜⬜⬜⬜⬛⬛⬛⬛⬛⬛⬜⬜⬜⬜
`

	banana3 = `
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣤⣼⣼⣶⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⠿⢿⣛⣯⣿⠗⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⣿⣿⣿⣿⣾⣿⢿⣿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣴⣿⡿⢟⣹⡟⢭⣿⢺⣿⣹⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣴⣿⣿⠅⢮⣿⠏⣼⣿⢠⣿⣷⠻⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢀⣠⣾⠟⣾⡏⡘⣾⣿⠣⢼⣿⡆⣿⣿⣅⡛⠷⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢠⣶⢿⡛⢥⣾⡟⠠⢹⣿⣿⡑⢺⣿⡥⠘⣿⣿⣔⠂⠌⠹⢶⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⣰⡿⢃⠆⣩⡾⠣⠜⡰⢻⣿⣿⡜⠼⣿⣧⠃⠄⠻⣿⣦⠐⢀⠂⠙⠷⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢠⡿⡑⡈⢶⠍⢢⠑⡸⢀⢹⣿⣿⡃⠌⢹⣿⡄⠂⠀⠘⢿⣶⣈⠀⢂⠀⠌⠛⣦⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢸⡇⠱⡌⢇⡜⠀⠆⡑⠀⠠⢿⣧⠡⠐⠀⢻⣷⠠⠁⠄⠊⢻⣿⣶⡀⢌⠠⠀⠄⠙⠳⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢽⠌⡱⡘⢆⡂⢩⠐⣈⠁⠠⢿⣿⣂⠁⠀⠌⢿⣷⡀⢂⠐⠀⠙⣷⣿⢦⡈⠡⡈⠤⠁⠌⠛⠷⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⣺⢁⠖⣩⠒⡈⢄⠢⢐⡈⠄⢈⣿⣷⡅⠂⠀⠘⢿⣷⡀⢂⠘⢀⠠⠙⢿⣿⣧⣔⠁⢎⠀⢃⠐⡈⢙⠳⣦⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢼⢢⠘⡥⠓⡌⠀⢆⠡⢀⠂⠠⠘⣿⣷⡾⠀⠀⠌⢻⣿⣦⠈⠄⠃⠌⠠⢉⠻⢿⣿⣦⡘⢂⢂⠡⢂⡐⠠⠙⠛⢷⣤⣀⠀⠀⠀⠀⠀⠀
⢸⡃⢜⢨⠓⢤⠉⢄⠊⠄⢃⠂⠡⠘⣿⣿⣄⠀⠂⠄⠻⢿⣿⣤⡁⢊⠰⠐⡀⠂⠌⠛⠿⣿⣦⣅⡢⢁⠣⠌⡔⢂⡈⢉⠻⢶⣄⠀⠀⠀
⢸⡿⢐⠨⣙⠆⡘⠤⣉⠐⢂⡉⠄⡃⢈⠻⣿⣦⡃⠠⠀⠌⠻⢿⣿⣦⣆⠑⠠⢑⡈⠤⠁⡄⠩⢙⠻⠷⣧⣼⣈⠥⡐⢆⠐⠠⠙⢳⣄⠀
⠀⣿⣀⠃⠿⣾⡔⠰⣀⢃⠂⡌⠰⠠⡁⠄⡉⠻⣿⣤⡁⢊⠡⠀⡝⠻⢿⣿⣶⣆⣈⠓⠰⣀⠣⢈⠰⠐⠠⢉⠛⠻⢷⣬⣜⢠⢃⠎⣯⣷
⠀⢹⣷⡈⠖⡹⣷⡡⢂⠎⠰⡀⢃⠡⠑⠢⢄⠁⠌⠙⠷⣧⣈⠘⢠⠐⠨⢉⠛⠿⢿⣿⣶⣤⣊⡄⢣⠉⠦⠡⣌⠐⡈⢙⠻⣷⣮⣿⣿⡿
⠀⠀⢻⣇⡎⠰⢩⢷⣇⠌⡓⢌⠢⢡⠉⡜⠠⢃⠌⡐⢂⠌⡙⠻⢶⣬⣆⡡⢈⡐⡈⠌⡉⠛⠿⠿⣯⣞⡶⢧⣶⣬⣔⣧⣹⣿⣿⠉⠁⠀
⠀⠀⠀⢻⣧⢃⠆⢬⠹⣷⣌⠢⡑⠦⡑⠤⡁⢌⠰⠡⢌⠒⡠⢃⠤⣈⠉⠛⠷⢶⣥⣂⡔⣁⠂⡐⢈⠙⠻⣶⣾⣼⣬⣿⣿⣿⠟⠁⠀⠀
⠀⠀⠀⠀⠻⣞⣈⠆⡹⠔⡻⢷⣥⣂⠍⡒⢡⠎⡠⢃⠆⡱⢠⠉⡔⢠⠛⠸⢄⠂⠌⢛⢿⣷⣾⣶⢦⡘⢶⣦⣿⣟⡛⠉⠀⠁⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠙⢷⣇⡰⢡⠱⣊⠝⡻⢾⣵⣈⢶⣕⡊⢆⡑⢆⠱⡈⢆⡉⢆⠢⢉⠆⡡⢊⣾⣿⢿⣆⠹⢌⣿⣿⡟⠇⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠙⠷⣦⡑⡈⢎⡑⠪⣍⣿⣿⣾⠿⣶⣭⣆⣧⡉⢆⠱⡈⢆⡁⢆⣡⣿⣿⣿⣫⣶⠿⠛⠛⠁⠁⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⠻⠶⠾⠷⠶⠿⠛⠉⠀⠀⠀⠉⠙⠛⠛⠷⠿⠾⠟⠛⠉⠉
`
)

func banana() string {
	return bananas[rand.Intn(len(bananas))]
}