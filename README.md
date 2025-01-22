```go
  _________ __                         __   .__               __                
 /   _____//  |________   ____ _____  |  | _|__| ____ _____ _/  |_  ___________ 
 \_____  \\   __\_  __ \_/ __ \\__  \ |  |/ /  |/    \\__  \\   __\/  _ \_  __ \
 /        \|  |  |  | \/\  ___/ / __ \|    <|  |   |  \/ __ \|  | (  <_> )  | \/
/_______  /|__|  |__|    \___  >____  /__|_ \__|___|  (____  /__|  \____/|__|   
        \/                   \/     \/     \/       \/     \/                   
                                                                                               
```

Streakinator is a GitHub automation tool that updates a `.txt` file in your repository with the current date and time every 24 hours. It automates commits and pushes to keep your Boot.Dev study streak alive.


## Features
- Clones the repository to a local directory.
- Updates `data/date.txt` with the current date and time.
- Commits and pushes the changes back to the repository.
- Runs every 24 hours via GitHub Actions.


## Setup Instructions

### **1. Fork This Repository**
- Click **Fork** in the top-right corner to create your own copy.

### **2. Enable GitHub Actions**
- Go to the **Actions** tab in your fork and click **Enable Actions**.

### **3. Add Repository Secrets**
- Go to **Settings > Secrets and variables > Actions**, then add the following secrets:

| Name                  | Required | Description                                    |
|-----------------------|----------|------------------------------------------------|
| `GITHUB_REPOSITORY_URL` | ✅      | URL of your fork (e.g., `https://github.com/<your-username>/<your-repo>.git`) |
| `GITHUB_TOKEN`        | ✅      | A GitHub PAT with `repo` permissions.          |
| `GITHUB_USERNAME`     | ✅      | Your GitHub username.                          |
| `GITHUB_EMAIL`        | ✅      | Your GitHub email address.                     |
| `GITHUB_TIMEZONE`     | ❌      | Timezone (e.g., `Europe/Paris`, defaults to UTC). |

---

## Local Development
To run Streakinator locally:
1. Clone the repository:
   ```bash
   git clone https://github.com/<your-username>/Streakinator.git
   cd Streakinator
2. Create a `.env` file at the root of the project with the required variables:


| Variable name                  | Format/example                                  
|-----------------------|----------|
|GITHUB_REPOSITORY_URL	| https://github.com/\<your-username>/\<your-repo>.git
|GITHUB_TOKEN	|  string
|GITHUB_USERNAME | string
|GITHUB_EMAIL | your-email@example.com
|GITHUB_TIMEZONE| Europe/Paris

3. Run the program:
	 ```bash   
	go run cmd/main.go
