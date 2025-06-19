# 🔐 AEShield

**AEShield** is a lightweight AES-256-GCM file encryptor and decryptor written in Go.  
Use it via a command-line interface or run it as a RESTful API for easy integration into web frontends or automation pipelines.

---

## ⚡ Features

- ✅ AES-256-GCM encryption (secure and authenticated)
- ✅ Fast command-line usage
- ✅ REST API for web integration
- ✅ Works cross-platform (Windows, macOS, Linux)
- ✅ No external dependencies

---

## 🚀 Usage

### 🔧 Command-Line Mode

```bash
# Encrypt a file
go run main.go encrypt path/to/file.txt 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef

# Decrypt a file
go run main.go decrypt path/to/file.txt.enc 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
```

## 🧭 Project Planning

Curious about what's next for AEShield?  
Check out the full roadmap and future feature ideas in [PLANNING.md](./PLANNING.md)
