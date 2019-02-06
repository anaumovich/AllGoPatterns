package View

import "fmt"

func PrintLoginPage(child ...string) string {
	return fmt.Sprint(`<!DOCTYPE html>
		<html>
			<head>
				<title>Page Title</title>
				<style>
					body, html {
					    height: 100%;
						}
					form {
					    height: 100%;
						display: flex;
						align-items: center;
						justify-content: center;
						flex-direction: column;
						}
				</style>
			</head>
		<body>
			<form action = "http://localhost:8080/">
				
					<p>`, childIterator(child), `</p>
					<input type = "text"> 
					<br>
					<input type = "password">
					<br>
					<button>Войти</button>
				
			</form>
		</body>
		</html>`)
}

func PrintDefaultPage(child ...string) string {
	return fmt.Sprint(`<!DOCTYPE html>
		<html>
			<head>
				<title>Page Title</title>
			</head>
		<body>
					<p>`, childIterator(child), `</p>
		</body>
		</html>`)
}

func PrintStartPage(child ...string) string {
	return fmt.Sprint(`<!DOCTYPE html>
		<html>
			<head>
				<title>Page Title</title>
				<style>
					body, html {
					    height: 100%;
						display: flex;
						align-items: center;
						justify-content: center;
						flex-direction: column;
						}
					table {
						width: 23%;
  						table-layout: fixed;
						}
					button {
						width: 100%;
						}
				</style>
			</head>
		<body>
			
				<p>`, childIterator(child), `</p>
				<table>
						<tr>
							<td><a href="http://localhost:8080/registration"><button>Регистрация </button></a></td>
							<td><a href="http://localhost:8080/login"><button>Войти</button></a></td>
						</tr>
				</table>
			
		</body>
		</html>`)
}

func childIterator(str []string) string {
	newStr := ""
	for i := range str {
		newStr += str[i]
	}

	return newStr
}
