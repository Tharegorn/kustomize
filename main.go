package main

func main() {
	welcome := Welcome{"GITOPS", time.Now().Format(time.Stamp), os.Getenv("HOSTNAME")}