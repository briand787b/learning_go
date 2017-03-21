package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func main() {
	req, _ := http.NewRequest("GET", `https://flash.envdev.mdsiinc.com/v2.2/customers/attc012/inventory/items`, nil)

	req.Header.Add(`Authorization`, `Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsIng1dCI6IlVyS282dDJsWHQ3TEJZb2ZjeExLcnhZaVppZyIsImtpZCI6IlVyS282dDJsWHQ3TEJZb2ZjeExLcnhZaVppZyJ9.eyJpc3MiOiJodHRwczovL2lkc3J2LmVudmRldi5tZHNpaW5jLmNvbS9jb3JlIiwiYXVkIjoiaHR0cHM6Ly9pZHNydi5lbnZkZXYubWRzaWluYy5jb20vY29yZS9yZXNvdXJjZXMiLCJleHAiOjE0ODIxOTU2NTAsIm5iZiI6MTQ4MjE5MjA1MCwiY2xpZW50X2lkIjoiRTc5NzFFQzEtMUFBMC00ODU4LUFDOTItNEZGOTIxRTFFMUQ2Iiwic2NvcGUiOlsiZmxhc2gtcmVhZCIsImZsYXNoLXdyaXRlIl0sInN1YiI6ImdvcmRvbmF0dHVzZXIiLCJhdXRoX3RpbWUiOjE0ODIxOTIwNTAsImlkcCI6Imlkc3J2IiwiZ2l2ZW5fbmFtZSI6IkdvcmRvbiIsImZhbWlseV9uYW1lIjoiQXR0VXNlciIsImVtYWlsIjoiR29yZG9uQXR0VXNlckBtZHNpaW5jLmNvbSIsImlkIjoiMzc5OSIsImFtciI6WyJwYXNzd29yZCJdfQ.N13BmHggOuKpghOtJfpNQsbsqOoWIcDPlsuYu7dM3z5eHJz4ENsogGnpccSTOBsurHGr6UWbCAW1u05JZt4ptnnyH_isk1xVprD4TXCIBiy2Pnxbi2FxBRJPfX90uAPEuSQOtfN0B-NnS58_wEpDD6TjwZ227Ijq_1aaNdJmnyNxi1FG3C4xu_tsLPw7LAOrA9mx0pdUKj4-CvcJJccwSEFWkxeP2Mde9wYxVwuFepcHQY_PWU0HUjV_mj7hmJyG9LK1z62bhFVQ1znLkrlUYb6--By-8CQzv-Xs0MF9kDBdVWKfhQfET2vV4zxP-upELTfQwSpJQVOXMWYGt8mxIJ8WFGv_XU1vRKTM-mP_eIIeU0hnSIyO0CfMXPfP0X1USqaL5pmnPHFFuTuv7UrUFJGFjTN0r9IkeeO45GS_0bZGOZ3AZ8R8cuxyzewNWjz0epIkilounvkwtCn7WaAL1rSQbQQs1ZloOMQtNPn7MVUmbwy2jz7EaI75LB9mjY_VRP1t3E5dRIDNSO8mndaJlKUBEUq-5NgDsrazKkxdlfozgvMNJIEHXslVZ7UtHochf2uItRw4lFb1tXdxJf9L4cwy4vU99rvZ-h1SCYXbi8O0e2YAWlU1IYp1qwGrG9ztomnih91OxeOABOPa3NAjhNDKGG10jv8ckxoJvzjZLmA`)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}


	var receiver interface{}
	json.Unmarshal(res.Body, &receiver)
	fmt.Println(res)
	fmt.Println("*************************************")
	fmt.Println(slice)
}