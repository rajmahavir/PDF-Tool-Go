# 🔧 Railway Deployment Troubleshooting

## Step-by-Step Fix Guide

### **STEP 1: Share Your Error Message** 📋

Please copy and share the error from Railway:

1. Go to your Railway project dashboard
2. Click on "Deployments" tab
3. Click on the failed deployment
4. Copy the entire error message (especially the first error you see)
5. Share it so I can give you a specific fix

Common error locations:
- **Build Logs**: Look for errors during compilation
- **Deploy Logs**: Look for runtime errors
- **Service Logs**: Look for application errors

---

## **STEP 2: Try the Ultra-Simple Dockerfile** 🚀

The repository now has `Dockerfile.simple` - the absolute simplest configuration.

### In Railway Dashboard:

1. Go to your project
2. Click **Settings** → **Build**
3. Set the following:

```
Builder: Dockerfile
Dockerfile Path: Dockerfile.simple
Root Directory: (leave empty)
```

4. Click **Redeploy**

This uses a simple single-stage Dockerfile with no optimization, just basic build and run.

---

## **STEP 3: Common Errors & Solutions** 🐛

### Error: "go: go.mod file not found"

**Solution:**
```bash
# Make sure go.mod is in the repository root
ls -la go.mod
```

### Error: "unsupported Go version" or "toolchain mismatch"

**Solution:**
The go.mod is already fixed to `go 1.21` (no toolchain). If you still see this:
1. Make sure you pulled the latest changes
2. In Railway, trigger a new deployment (not just redeploy)

### Error: "cannot find package"

**Solution:**
```bash
# Run locally to verify:
go mod download
go build .
```

If it works locally, Railway should work too.

### Error: "port already in use" or "listen tcp: address in use"

**Solution:**
The code already handles this with environment PORT variable. No changes needed.

### Error: "build succeeded but app won't start"

**Check logs:** Railway Dashboard → Deployments → View Logs

Common causes:
- Binary not executable → Fixed in Dockerfile.simple
- PORT not detected → Already handled in main.go
- Missing files → Check Dockerfile COPY command

---

## **STEP 4: Alternative - Remove All Config Files** 🎯

If all else fails, let Railway auto-detect:

### In Railway Dashboard:

1. Settings → Build
2. Set Builder: **Nixpacks** (or leave as Auto)
3. Clear all paths (Dockerfile Path, Root Directory)
4. Redeploy

Railway will auto-detect Go and build it.

---

## **STEP 5: Manual Verification** ✅

Test locally that everything works:

```bash
# 1. Clean build
go clean
go mod tidy

# 2. Build
go build -o pdf-tools .

# 3. Run
./pdf-tools

# 4. Test (in another terminal)
curl http://localhost:8080
```

If this works, Railway should work too.

---

## **STEP 6: Nuclear Option - Start Fresh** ☢️

If nothing works, try creating a new Railway project:

1. **Delete the old Railway project**
2. **Create new project** from GitHub
3. **Select the repository**
4. **Select branch:** `claude/check-this-011CUpJ6f4iP3k14PzNMfmsJ`
5. Railway will auto-detect and deploy

---

## **Configuration Files Available** 📁

Your repository has **4 Dockerfile options**:

1. **`Dockerfile.simple`** ← **TRY THIS FIRST** ⭐
   - Ultra-simple, no optimizations
   - Single stage
   - Most likely to work

2. **`Dockerfile.railway`**
   - Simple single-stage
   - No multi-stage complexity

3. **`Dockerfile`**
   - Multi-stage optimized
   - Production-ready

4. **Nixpacks** (no Dockerfile)
   - Uses `nixpacks.toml` and `Procfile`
   - Railway's auto-detection

---

## **What Railway Needs** 📋

Railway needs these files (all present):
- ✅ `go.mod` - Go module definition
- ✅ `go.sum` - Dependency checksums
- ✅ `main.go` - Application entry point
- ✅ One of: Dockerfile, nixpacks.toml, or nothing (auto-detect)

---

## **Quick Diagnosis Checklist** ✓

Check these in order:

- [ ] go.mod has `go 1.21` (not 1.24)
- [ ] go.mod has no `toolchain` line
- [ ] main.go reads PORT from environment
- [ ] All files are committed and pushed
- [ ] Railway is using the correct branch
- [ ] Railway builder is set correctly

---

## **Get Help** 🆘

To get specific help, share:

1. **Railway error message** (full text or screenshot)
2. **Builder type** (Dockerfile or Nixpacks)
3. **Which Dockerfile** you're using
4. **Build logs** (from Railway dashboard)

---

## **Working Configuration** ✅

This SHOULD work on Railway:

**File: railway.json**
```json
{
  "build": {
    "builder": "DOCKERFILE",
    "dockerfilePath": "Dockerfile.simple"
  }
}
```

**File: Dockerfile.simple**
```dockerfile
FROM golang:1.21
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o pdf-tools .
EXPOSE 8080
CMD ["./pdf-tools"]
```

**File: go.mod**
```
go 1.21
```

This is the absolute minimum configuration. If this doesn't work, there's something else wrong (not with the config).

---

## **Next Steps** 👉

1. **Try Dockerfile.simple** (instructions in STEP 2)
2. **Share error message** if it still fails
3. **Check Railway service status**: https://status.railway.app

Railway sometimes has outages or issues on their end.

---

**Most deployments fail due to:**
1. ❌ Go version mismatch → **FIXED** (go 1.21)
2. ❌ Missing PORT handling → **FIXED** (reads from env)
3. ❌ Complex Dockerfile → **FIXED** (Dockerfile.simple)
4. ❌ Wrong builder → **TRY STEP 2**

Try Dockerfile.simple and share the error if it still fails! 🚀
