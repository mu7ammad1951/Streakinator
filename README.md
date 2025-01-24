```go
  _________ __                         __   .__               __                
 /   _____//  |________   ____ _____  |  | _|__| ____ _____ _/  |_  ___________ 
 \_____  \\   __\_  __ \_/ __ \\__  \ |  |/ /  |/    \\__  \\   __\/  _ \_  __ \
 /        \|  |  |  | \/\  ___/ / __ \|    <|  |   |  \/ __ \|  | (  <_> )  | \/
/_______  /|__|  |__|    \___  >____  /__|_ \__|___|  (____  /__|  \____/|__|   
        \/                   \/     \/     \/       \/     \/                   
                                                                                               
```

![Made with Go and LOVE](https://img.shields.io/badge/Made%20with%20Go%20and-LOVE-blue?logo=go)
[![Go Report Card](https://goreportcard.com/badge/github.com/emield/streakinator)](https://goreportcard.com/report/github.com/emield/streakinator)

Streakinator is a GitHub automation tool that updates a `.txt` file in your repository with the current date and time every 24 hours. It automates commits and pushes to keep your Boot.Dev study streak alive.

> **⚠ Important Note:**  
> If you're new to programming and trying to build the habit of coding regularly, you should **NOT** use this tool!  
>  
> Consistency and personal effort are key to developing programming skills, and relying on automation defeats the purpose of building healthy habits.  
>  
> This tool is not meant for cheating—it's meant for those who already know what they’re doing but are too lazy to worry about keeping their streak up every day.

---

## 🛠 **Features**  
- Clones the repository to a local directory.  
- Updates `data/date.txt` with the current date and time.  
- Commits and pushes the changes back to the repository.  
- Runs every 24 hours via GitHub Actions.  

---

## 📦 **Setup Instructions**  

### **1. 🧸 Make sure your Boot.Dev account is in sync with GitHub**  
- On Boot.Dev -> Settings -> Account Syncing (GitHub).  

### **2. 🍴 Fork This Repository**  
- Click **Fork** in the top-right corner to create your own copy.  

### **3. ⚙️ Enable GitHub Actions**  
- Go to the **Actions** tab in your fork and click **Enable Actions**.  

### **4. 🔑 Add Repository Secrets**  
- Go to **Settings > Secrets and variables > Actions**, then add the following secrets:  

| Name                  | Required | Description                                    |  
|-----------------------|----------|------------------------------------------------|  
| `S_GITHUB_REPOSITORY_URL` | ✅      | URL of your fork (e.g., `https://github.com/<your-username>/Streakinator.git`) |  
| `S_GITHUB_TOKEN`        | ✅      | A GitHub PAT with `repo` permissions.          |  
| `S_GITHUB_USERNAME`     | ✅      | Your GitHub username.                          |  
| `S_GITHUB_EMAIL`        | ✅      | Your GitHub email address.                     |  
| `S_GITHUB_TIMEZONE`     | ❌      | Timezone (e.g., `Europe/Paris`, defaults to UTC). |  

---

## 💻 **Local Development**  

To run Streakinator locally:
1. Clone the repository:
   ```bash
   git clone https://github.com/<your-username>/Streakinator.git
   cd Streakinator
2. Create a `.env` file at the root of the project with the required variables:


| Variable name                  | Format/example                                  
|-----------------------|----------|
|`S_GITHUB_REPOSITORY_URL`	| `https://github.com/<your-username>/Streakinator.git`
|`S_GITHUB_TOKEN`	        | string
|`S_GITHUB_USERNAME`            | string
|`S_GITHUB_EMAIL`               | your-email@example.com
|`S_GITHUB_TIMEZONE`            | Europe/Paris

3. Run the program:
	 ```bash   
	go run cmd/main.go
   cd Streakinator  
