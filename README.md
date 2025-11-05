# PDF Tools - Free & Open Source PDF Editor

A powerful, free, and open-source web application for manipulating PDF files. Built with Go and powered by pdfcpu.

## Features

- **Merge PDFs**: Insert all pages from one PDF into another at any position
- **Remove Pages**: Delete specific pages from a PDF with visual page selection
- **PDF Preview**: In-browser PDF viewing before and after operations
- **Privacy-Focused**: No data storage, files deleted immediately after processing
- **No Registration**: Use instantly without creating an account

## Screenshots

### Home Page
Choose between merging PDFs or removing pages with an intuitive interface.

### Merge PDFs
Insert all pages from one PDF into another at a specific position with live preview.

### Remove Pages
Select pages to remove with a visual page selector and instant preview.

## Tech Stack

- **Backend**: Go 1.21+
- **PDF Library**: [pdfcpu](https://github.com/pdfcpu/pdfcpu) v0.11.1
- **Frontend**: Vanilla JavaScript with embedded HTML/CSS
- **Server**: Native Go HTTP server

## Installation

### Prerequisites

- Go 1.21 or higher
- Git

### Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/rajmahavir/PDF-Tools.git
   cd PDF-Tools
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build the application**
   ```bash
   go build -o pdf-tools .
   ```

4. **Run the server**
   ```bash
   ./pdf-tools
   ```

5. **Access the application**
   - Local: http://localhost:8080
   - Network: http://[YOUR-IP]:8080

## Docker Deployment

### Using Docker

```bash
docker build -t pdf-tools .
docker run -p 8080:8080 pdf-tools
```

### Using Docker Compose

```bash
docker-compose up -d
```

The application will be available at http://localhost:8080

## Railway Deployment

Deploy to Railway using automatic Nixpacks detection (recommended) or Docker.

### 🚀 Quick Deploy (Recommended)

**Railway automatically detects and deploys Go applications!**

1. **Connect to Railway**
   - Go to [Railway](https://railway.app)
   - Click "New Project" → "Deploy from GitHub repo"
   - Select your repository
   - Railway auto-detects configuration ✨

2. **Deploy!**
   - No configuration needed
   - Railway uses Nixpacks to build
   - App deploys automatically

### 📋 Deployment Methods

This repository supports **3 deployment methods**:

1. **Nixpacks (Automatic)** - Recommended ⭐
   - Uses: `nixpacks.toml`, `Procfile`
   - No Docker required
   - Fastest and simplest

2. **Dockerfile (Multi-stage)**
   - Uses: `Dockerfile`
   - Optimized production build
   - Smaller image size

3. **Dockerfile.railway (Simple)**
   - Uses: `Dockerfile.railway`
   - Single-stage build
   - Fallback option

### 🔧 Configuration Files

The repository includes complete Railway configuration:
- `nixpacks.toml` - Nixpacks build configuration
- `Procfile` - Start command
- `railway.json` - Railway settings
- `railway.toml` - Alternative Railway config
- `Dockerfile` - Production Docker build
- `Dockerfile.railway` - Simplified Docker build

### 📚 Detailed Guide

**Having deployment issues?** See the complete guide:
📖 **[RAILWAY_DEPLOY.md](./RAILWAY_DEPLOY.md)** - Comprehensive Railway deployment guide with troubleshooting

### ⚡ Railway CLI (Alternative)

```bash
# Install Railway CLI
npm install -g @railway/cli

# Login
railway login

# Initialize project
railway init

# Deploy
railway up
```

### 🔄 Switching Deployment Methods

In Railway Dashboard → Settings → Build:
- **For Nixpacks:** Set Builder = "Nixpacks"
- **For Docker:** Set Builder = "Dockerfile", Path = "Dockerfile"
- **For Simple Docker:** Set Builder = "Dockerfile", Path = "Dockerfile.railway"

## Project Structure

```
PDF-Tool-Go/
├── handlers/          # HTTP request handlers
│   ├── home.go       # Home page handler
│   ├── credits.go    # Credits page handler
│   ├── merge_page.go # Merge PDF page handler
│   ├── merge_handler.go # Merge PDF backend logic
│   ├── remove_page.go # Remove pages page handler
│   ├── remove_handler.go # Remove pages backend logic
│   └── pdfinfo.go    # PDF information handler
├── pdf/              # PDF processing logic
│   └── operations.go # PDF manipulation functions
├── templates/        # HTML templates and components
│   └── templates.go  # Common HTML components
├── main.go           # Application entry point
├── go.mod            # Go module definition
├── go.sum            # Go dependencies checksum
├── Dockerfile        # Docker configuration
└── README.md         # This file
```

## Development

### Building

```bash
go build -o pdf-tools .
```

### Running Tests

```bash
go test ./...
```

### Running in Development Mode

```bash
go run main.go
```

## Configuration

The server runs on port 8080 by default and listens on all network interfaces (0.0.0.0).

### Environment Variables

- `PORT`: Server port (default: 8080)

### File Upload Limits

- Maximum file size: 50 MB per PDF
- Supported format: PDF only

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Home page |
| `/merge` | GET | Merge PDFs page |
| `/remove` | GET | Remove pages page |
| `/credits` | GET | Credits and about page |
| `/pdfinfo` | POST | Get PDF page count and metadata |
| `/merge-pdfs` | POST | Merge two PDF files |
| `/remove-pages` | POST | Remove pages from a PDF |

## Security & Privacy

- **No Data Storage**: Files are processed in temporary directories and deleted immediately
- **No Tracking**: No cookies, analytics, or user data collection
- **No Registration**: No user accounts or authentication required
- **Local Processing**: All PDF operations happen on the server, files are never uploaded to third parties
- **Secure**: Input validation and file type checking on all uploads

## Performance

- Fast PDF processing using native Go libraries
- Efficient memory management with automatic cleanup
- Handles PDFs up to 50MB
- Optimized for concurrent requests

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### How to Contribute

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Guidelines

- Follow Go best practices and conventions
- Add tests for new features
- Update documentation as needed
- Keep commits clean and descriptive

## Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## Troubleshooting

### Build Errors

If you encounter build errors, try:

```bash
go mod tidy
go clean -modcache
go build .
```

### Port Already in Use

If port 8080 is already in use, modify the port in `main.go` or set the `PORT` environment variable.

### Network Access Issues

To access from other devices on your network, ensure:
1. The server is running on 0.0.0.0 (default)
2. Your firewall allows connections on port 8080
3. You're using your computer's IP address, not localhost

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Third-Party Licenses

- **pdfcpu**: Apache License 2.0
- **Go**: BSD 3-Clause License

## Credits

### AI-Assisted Development

This application was developed with assistance from Claude.ai (Anthropic), including:
- Architecture design
- Code generation
- User interface development
- PDF processing implementation

### Technology

- **pdfcpu** by Horst Rutter - PDF processing library
- **Go Team** - Go programming language
- **Open Source Community** - For making projects like this possible

## Acknowledgments

Special thanks to:
- Anthropic for creating Claude.ai
- Horst Rutter for developing pdfcpu
- The Go team for the excellent programming language
- The open source community

## Support

- **Issues**: https://github.com/rajmahavir/PDF-Tools/issues
- **Discussions**: https://github.com/rajmahavir/PDF-Tools/discussions

## Roadmap

Future features being considered:
- [ ] PDF splitting into multiple files
- [ ] PDF rotation
- [ ] PDF encryption/decryption
- [ ] PDF watermarking
- [ ] Batch processing
- [ ] REST API
- [ ] PDF compression
- [ ] Image to PDF conversion
- [ ] PDF metadata editing

## Changelog

### Version 1.0.0 (Initial Release)
- PDF merging functionality
- Page removal functionality
- Web-based interface
- Docker support
- Refactored codebase with modular structure

---

Made with ❤️ using Claude.ai | Free Forever | MIT License
