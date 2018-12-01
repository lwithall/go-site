package main

import (
   "net/http"
   "time"
   "html/template"
   "log"
)


type Welcome struct {
   Title string
   Name string
   Time string
}

func login(w http.ResponseWriter, r *http.Request) {
    welcome := Welcome{"Login", "Anonymous", time.Now().Format(time.Stamp)}

    templates := template.Must(template.ParseFiles("login.html"))
    if err := templates.ExecuteTemplate(w, "login.html", welcome); err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/", login)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

/*func main() {
   welcome := Welcome{"Login", "Anonymous", time.Now().Format(time.Stamp)}

   templates := template.Must(template.ParseFiles("login.html"))

   http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {

      if name := r.FormValue("name"); name != "" {
         welcome.Name = name;
      }
      if err := templates.ExecuteTemplate(w, "login.html", welcome); err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
      }
   })

   //Start the web server, set the port to listen to 8080. Without a path it assumes localhost
   //Print any errors from starting the webserver using fmt
   fmt.Println("Listening");
   fmt.Println(http.ListenAndServe(":8080", nil));
}*/
