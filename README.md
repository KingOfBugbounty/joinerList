# 🚀 UniqWordlist - Fast Subdomain Joiner

UniqWordlist is a high-performance tool designed to merge a wordlist with a list of subdomains. It takes each word from the wordlist and appends it to every subdomain, generating a comprehensive list of potential subdomains. The results are saved in a `final.txt` file.

---

## 📥 Installation

To install and compile UniqWordlist, follow these steps:

### 🔹 Install Go (if not already installed)
```bash
sudo apt update && sudo apt install -y golang
```

### 🔹 Clone the repository
```bash
git clone https://github.com/KingOfBugbounty/uniqwordlist.git
cd uniqwordlist
```

### 🔹 Build the binary
```bash
go build uniqwordlist.go
```

---

## 🛠 Usage

Once compiled, you can use UniqWordlist with the following command:

```bash
./uniqwordlist -wordlist common.txt -subdomains dod.txt
```

🔹 `-wordlist`: Path to the wordlist file (e.g., `common.txt`).  
🔹 `-subdomains`: Path to the subdomains file (e.g., `dod.txt`).

### 📌 Example
#### **Input Files**

🔹 `common.txt`
```
admin
webmail
portal
```

🔹 `dod.txt`
```
army.mil
defense.gov
```

#### **Output (`final.txt`)**
```
admin.army.mil
webmail.army.mil
portal.army.mil
admin.defense.gov
webmail.defense.gov
portal.defense.gov
```

After execution, all generated subdomains will be saved in `final.txt`. 🎯

---

## 📜 License
This project is open-source and available under the MIT License.

---

## 💡 Contributing
Pull requests are welcome! Feel free to submit issues or suggestions.

Happy hacking! 🛠️🔥

