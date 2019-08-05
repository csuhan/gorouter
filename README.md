# A simple router based on golang

easy way to use it
```golang
router := New()
router.Get("/user", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("GET"))
})
router.Post("/admin",func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Post"))
})
router.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello,%s", "csuhan")
})
http.ListenAndServe(":8080", router)
```

there is still some bugs,please do not use it.