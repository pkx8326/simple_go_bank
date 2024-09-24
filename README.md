# Simnple Go Bank

## NOTE: The code provided in this repository yields the same results as the Python scripts provided here: https://github.com/pkx8326/simple_python_bank. It's the same program written in the Go programming language.

### Overview
In this project, I programed an example of a simple, command-line-based, banking application with Go. The program can performs examples of basic banking transactions, including checking curreent balance and making deposit and withdrawal. The current balance information can be kept locally on file. With each session, the program will read the last updated balance information from the file and continue from that point.

This simple example of a banking program is written in Go to demonstrate the language's modularity, functions, input validation and control flow structures. The code uses infinit for loops and switch-case blocks. Albeit simple, the program is robust and complete and can be used to study the language basic functionalities and structures.

### Program manual
The first time the program is run, the user will be presented with a set of banking transaction menu including:
1. Display the balance
2. Deposit
3. Exit

To perform a transaction (or exit the program), simply input the number preceeding each menu description and press enter. The user can input anything apart from the numbers 1,2, and 3 but will get a notification that the input is invalid and will be asked to input the correct number again. At this time, if the user chooses 1, the displayed balance will be 0 because there's no saved balance data yet. There'll also be a message notifying the user to make the first deposit. Once the user choose to deposit, the program will ask for an amount to deposit. The user can again input anything apart from numbers but the program will notify the user of any invalid input and will keep asking the user again and again until it gets a valid input value. Once the deposit amount is taken, the balance data file is created with the information of the user's current balance.

After finishing the deposit procedure, the program asks if the user still have any more transaction to do. At this point, the user can simply type "Y" or "y" for yes or "N" or "n" for no. Again, the program keeps track of the user input and will take in only the correct input format. If the user inputs "n", the program ends with a goodbye message. Otherwise, the user will be presented with the numbered transaction menu again, this time there are 4 choices as the following list:
1. Display the balance
2. Deposit
3. Withdraw
4. Exit

After depositing some amount, a balance file is created and the program will now read the updated balance information from that file. The file information will be overwritten to update everytime the user deposit or withdraw an amount. The deposit and withdraw menu will not take any number less than or equal to 0. For the deposit menu, the user will be notify if there's an attempt to withdraw any amount bigger than the curreent balance and will be brought back to the main menu after confirming another transaction with 4 transaction menu options.

#### Code structure
The Go code is divided into two main parts being the main program code in the main.go file and the code for different transansaction options in the optn.go file. This code structure is designed to demonstrate Go's modularity and the use of functions.


### Program flow
The program flow is shown as the following numbered list:
1. The program checks for a saved balance data file if no file found, the balance is set to 0
2. The program presents the menu options for the user to choose between different types of transactions (or exit the program)
3. If the user chooses to display the current balance, the read data from the balance file (or 0 in the case that there's no balance file found) will be displayed
4. If the user choses either to deposit or to withdraw an amount, the new balance will be updated according to the saved balance data
5. After each transaction (display balance, deposit, withdraw) the user will be asked if he needs to perform another transaction or to exit the program. If the user chooses to do another transaction, he will be again presented with the main menu, otherwise, the program will end with a goodbye message.


