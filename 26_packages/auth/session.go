package auth

//here we are exposing this method outside but we want to keep our logic private
//in private function
//you will not be able to access this using .

func extractSession() string {
	return "loggedIn"
}
func GetSession() string {
	return extractSession()
}
