# 🧭 AEShield — Project Planning & Future Ideas

This document outlines the roadmap, future goals, and planned features for AEShield — a Go-based AES-256 file encryptor/decryptor with both CLI and REST API modes.

---

## ✅ Current Status

- AES-256-GCM file encryption and decryption implemented
- Command-line interface (CLI) for local usage
- REST API with `/encrypt` and `/decrypt` endpoints
- Cross-platform support (Windows, macOS, Linux)
- Secure encryption using random nonce and key validation

---

## 🗓️ Short-Term Goals (Next 1–4 months)

### 🛠 CLI Improvements
- Add CLI flags for:
  - Custom output file name
  - Overwrite confirmation
- Add password-to-key conversion (PBKDF2 or Argon2)

### 🌐 REST API Enhancements
- Return JSON responses with status codes
- Implement API key or token-based authentication
- Add request rate limiting
- Add simple logging and error tracking

### 🎨 Frontend (Prototype)
- Build a basic HTML/JS frontend to:
  - Upload a file
  - Enter a 64-char hex key
  - Call `/encrypt` and `/decrypt`
  - Trigger downloads of the resulting file

### 🐳 Docker Support
- Create a `Dockerfile` to containerize the Go server
- Add usage instructions for running AEShield in a container:
  - `docker build -t aeshield .`
  - `docker run -p 8080:8080 aeshield`
- Document Docker usage in README

---

## 🚀 Medium-Term Goals (4–12 months)

- Batch mode for encrypting multiple files or entire folders
- Optional metadata encryption
- Integrate cloud storage support (AWS S3, Google Drive)
- Alternative encryption algorithms (ChaCha20, AES-CBC)
- GUI desktop wrapper (Electron, Wails, or Fyne again)
- Error message localization and better UX
- Integration/unit tests for crypto logic and API

---

## 🌍 Long-Term Vision (12+ months)

- Publish Docker image to Docker Hub
- Release cross-platform binaries
- Add Web UI + backend authentication for secure deployment
- Create mobile app wrappers (iOS/Android) using shared API
- Support for secure file deletion (file shredding)
- Create library version (for import into other Go projects)

---

## 💡 Optional Ideas

- Password manager or secure vault integration
- Keychain/TPM support for secure key storage
- Encrypted archives (ZIP or TAR-style output)
- Encrypted metadata and secure sharing keys (e.g. PGP layer)
- Gi

### ❌ Dropped: Native Desktop GUI (e.g., Fyne)

Most users are expected to interact via a web browser, so a cross-platform desktop GUI (Electron/Fyne/Wails) is no longer planned. The focus will shift entirely to building a responsive, secure Single Page Application (SPA) frontend.

