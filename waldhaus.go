package main

import (
	
	"fmt"
	"os" // reader in go
)

func main() {

	//fmt.Println("Arguments :", os.Args)
   if len(os.Args)==1{                                         //Bedingung wenn 
		fmt.Println("error: bite geben sie ein argument")
		os.Exit(0)
	}

	text := os.Args[1]
	 
	//*******zeile 17 bis 25  sprachblasen******


		
	for k := 0; k < len(text); k++ {       //for loop
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println(text)

	for k := 0; k < len(text); k++ {
		fmt.Printf("-")
	}
	

	printCow() // Aufrufen von Methode

}
func printCow() {

	cow := `  
	                    \
	                     \ 
                              \
                               \ 
	                        \
	                         \                           [*]
					                   /*****\
					                 /**********\ 
	                  /.\ | |                       / ********** \
	                 /   \| |                      /**************\             
                        /     \ |           __        /****************\
                       /  (O)  \           /**\      / **************** \
                      /_________\        /******\    /___  ********   ___\
                      |         |        \******/      /   ********    \
                      |     [O] |         \****/      /___********** ___\
                      |         |          \**/          /___________\       
                      | []      |           ||                ||    
                      | []      |           ||                ||
||||||||||||||||||||| -----------|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||		           
			       
			   `

	fmt.Println(cow) // ausgabe in Terminal
}