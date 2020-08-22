package main  

import "fmt" 

func main() {


}


auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate)
auth.SetAuthInfo(clientID, secretKey)

url := auth.AuthURL(state)

func redirectHandler(w http.ResponseWriter, r *http.Request) {
      // use the same state string here that you used to generate the URL
      token, err := auth.Token(state, r)
      if err != nil {
            http.Error(w, "Couldn't get token", http.StatusNotFound)
            return
      }
      // create a client using the specified token
      client := auth.NewClient(token)
}
