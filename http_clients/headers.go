package main

import (
	"fmt"
)

func main() {
	url := "https://api.boot.dev/v1/courses_rest_api/learn-http/projects/52fdfc07-2182-454f-963f-5f0f9a621d72"
	apiKey := generateKey()

	oldProject, err := getProjectResponse(apiKey, url)
	if err != nil {
		fmt.Println("Error getting old project:", err)
		return
	}
	fmt.Println("Got old project:")
	fmt.Printf("- title: %s\n", oldProject.Title)
	fmt.Printf("- assignees: %d\n", oldProject.Assignees)
	fmt.Println("--------------------------------")

	newProjectData := Project{
		ID:        "52fdfc07-2182-454f-963f-5f0f9a621d72",
		Title:     "Product Roadmap 2025",
		Completed: false,
		Assignees: 1,
	}

	if err := putProject(apiKey, url, newProjectData); err != nil {
		fmt.Println("Error updating project:", err)
		return
	}
	fmt.Println("Project updated!")
	fmt.Println("---")

	//newApiKey := generateKey()
	newProject, err := getProjectResponse(apiKey, url)
	if err != nil {
		fmt.Println("Error getting new project:", err)
		return
	}
	fmt.Println("Got new project:")
	fmt.Printf("- title: %s\n", newProject.Title)
	fmt.Printf("- assignees: %d\n", newProject.Assignees)
	fmt.Println("--------------------------------")
}
func getUsers(url string) ([]User, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
	    return nil, err
	}
	return users, err
}

type Comment struct {
	Id      string `json:"id"`
	UserId  string `json:"user_id"`
	Comment string `json:"comment"`
}

func createComment(url, apiKey string, commentStruct Comment) (Comment, error) {
    // encode our comment as json
	jsonData, err := json.Marshal(commentStruct)
	if err != nil {
		return Comment{}, err
	}

    // create a new request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return Comment{}, err
	}

    // set request headers
	req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-API-Key", apiKey)

    // create a new client and make the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Comment{}, err
	}
	defer res.Body.Close()

    // decode the json data from the response
	// into a new Comment struct
	var comment Comment
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&comment)
	if err != nil {
		return Comment{}, err
	}

	return comment, nil
}