# Railway Deployment Guide

## 🚀 Quick Deploy

Railway now has **3 deployment methods**. Try them in this order:

---

## Method 1: Nixpacks (RECOMMENDED) ✨

This is the easiest and most reliable method for Railway.

### Configuration Files
- `nixpacks.toml` - Tells Railway how to build
- `Procfile` - Tells Railway how to run
- `railway.json` - Railway settings (uses NIXPACKS)

### Deploy Steps
1. Push your code to GitHub
2. Connect to Railway
3. Railway auto-detects and deploys!

**No Docker required!** Railway's Nixpacks will:
- ✅ Detect Go 1.21
- ✅ Download dependencies
- ✅ Build the binary
- ✅ Run the application

---

## Method 2: Dockerfile (If Nixpacks Fails)

If Nixpacks fails, use the Docker method:

### In Railway Dashboard:
1. Go to Settings → Build
2. Change Builder from "Nixpacks" to "Dockerfile"
3. Set Dockerfile Path: `Dockerfile`
4. Redeploy

### Configuration
- Uses `Dockerfile` (multi-stage build)
- Go 1.21-alpine
- Optimized binary

---

## Method 3: Simple Dockerfile (Last Resort)

If the main Dockerfile fails:

### In Railway Dashboard:
1. Go to Settings → Build
2. Set Builder: "Dockerfile"
3. Set Dockerfile Path: `Dockerfile.railway`
4. Redeploy

### Why This Works
- Single-stage build (simpler)
- No multi-stage complexity
- Fewer moving parts

---

## 🔧 Configuration Files

### `nixpacks.toml` (Method 1)
```toml
[phases.setup]
nixPkgs = ["go_1_21"]

[phases.build]
cmds = ["go mod download", "go build -ldflags='-w -s' -o pdf-tools ."]

[start]
cmd = "./pdf-tools"
```

### `Procfile` (Method 1)
```
web: ./pdf-tools
```

### `railway.json` (All Methods)
```json
{
  "build": {
    "builder": "NIXPACKS"  // or "DOCKERFILE"
  },
  "deploy": {
    "startCommand": "./pdf-tools",
    "restartPolicyType": "ON_FAILURE"
  }
}
```

---

## 🐛 Troubleshooting

### Build Fails with "Go version mismatch"

**Problem:** go.mod and Dockerfile have different Go versions

**Solution:**
```bash
# Verify go.mod has:
go 1.21

# NOT:
go 1.24.0
toolchain go1.24.7
```

### Port Binding Issues

**Problem:** App doesn't respond on Railway URL

**Solution:** The app already reads from `PORT` env var:
```go
port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
```
Railway automatically sets this!

### Build Succeeds but App Crashes

**Check Railway Logs:**
1. Go to your Railway project
2. Click "View Logs"
3. Look for errors

**Common Issues:**
- Missing dependencies → Run `go mod tidy` locally
- Build artifacts → Clear Railway cache and rebuild
- Binary not executable → Fixed in latest Dockerfile

### "Cannot find module"

**Solution:**
```bash
# Run locally:
go mod tidy
git add go.mod go.sum
git commit -m "Update dependencies"
git push
```

---

## ✅ Verification Checklist

Before deploying, verify:

- [ ] `go.mod` has `go 1.21` (no toolchain line)
- [ ] `main.go` reads PORT from environment
- [ ] All files committed and pushed
- [ ] Build works locally: `go build .`
- [ ] Tests pass: `go test ./...`

---

## 🔄 Switching Deployment Methods

### From Nixpacks to Dockerfile:
1. Railway Dashboard → Settings → Build
2. Change Builder: "DOCKERFILE"
3. Dockerfile Path: `Dockerfile`
4. Redeploy

### From Dockerfile to Nixpacks:
1. Railway Dashboard → Settings → Build
2. Change Builder: "NIXPACKS"
3. Clear Dockerfile Path
4. Redeploy

### To Simple Dockerfile:
1. Railway Dashboard → Settings → Build
2. Builder: "DOCKERFILE"
3. Dockerfile Path: `Dockerfile.railway`
4. Redeploy

---

## 📊 Expected Results

### Successful Deployment:

```
✅ Build succeeded
✅ Deploy succeeded
✅ Health check: passing
✅ App URL: https://your-app.railway.app
```

### Build Time:
- Nixpacks: ~2-3 minutes
- Dockerfile: ~3-4 minutes
- Dockerfile.railway: ~2-3 minutes

### Binary Size:
- ~21MB (optimized with -ldflags="-w -s")

---

## 🆘 Still Not Working?

1. **Check Railway Status:** https://status.railway.app
2. **Clear Build Cache:** Settings → Clear Build Cache → Redeploy
3. **Try Different Method:** Switch between Nixpacks/Dockerfile
4. **Check Logs:** View detailed build and runtime logs
5. **Environment Variables:** Verify PORT is set (Railway does this automatically)

---

## 📝 Quick Reference

| Method | Builder | File Used | Complexity |
|--------|---------|-----------|------------|
| Method 1 | Nixpacks | nixpacks.toml | ⭐ Easy |
| Method 2 | Dockerfile | Dockerfile | ⭐⭐ Medium |
| Method 3 | Dockerfile | Dockerfile.railway | ⭐ Easy |

**Start with Method 1 (Nixpacks) - it's the most reliable!**

---

## 🎉 Success!

Once deployed, your app will be available at:
```
https://[your-project-name].railway.app
```

Railway provides:
- ✅ Automatic HTTPS
- ✅ Auto-scaling
- ✅ Environment variables
- ✅ Deployment logs
- ✅ Health monitoring

---

**Need more help?** Check Railway docs: https://docs.railway.app
