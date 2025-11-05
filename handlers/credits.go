package handlers

import (
	"net/http"
)

// HandleCredits serves the credits page
func HandleCredits(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Credits - PDF Tools</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 40px 20px;
        }
        .container {
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            padding: 40px;
            max-width: 800px;
            margin: 0 auto;
        }
        .back-btn {
            display: inline-block;
            color: #667eea;
            text-decoration: none;
            margin-bottom: 20px;
            font-size: 14px;
            font-weight: 500;
        }
        .back-btn:hover {
            text-decoration: underline;
        }
        h1 {
            color: #333;
            margin-bottom: 10px;
            font-size: 32px;
        }
        .subtitle {
            color: #666;
            margin-bottom: 30px;
            font-size: 16px;
        }
        .section {
            margin-bottom: 30px;
            padding: 20px;
            background: #f8f9fa;
            border-radius: 10px;
            border-left: 4px solid #667eea;
        }
        .section h2 {
            color: #333;
            font-size: 20px;
            margin-bottom: 15px;
        }
        .section h3 {
            color: #555;
            font-size: 16px;
            margin-top: 15px;
            margin-bottom: 10px;
        }
        .section p {
            color: #666;
            line-height: 1.6;
            margin-bottom: 10px;
        }
        .credit-item {
            background: white;
            padding: 15px;
            border-radius: 8px;
            margin-bottom: 15px;
        }
        .credit-item strong {
            color: #667eea;
            display: block;
            margin-bottom: 5px;
        }
        .credit-item a {
            color: #667eea;
            text-decoration: none;
        }
        .credit-item a:hover {
            text-decoration: underline;
        }
        .badge {
            display: inline-block;
            padding: 5px 12px;
            background: #e8f0fe;
            color: #1967d2;
            border-radius: 20px;
            font-size: 12px;
            margin: 5px 5px 5px 0;
            font-weight: 500;
        }
        .badge.ai {
            background: #fce8f3;
            color: #c2185b;
        }
        .badge.license {
            background: #e8f5e9;
            color: #2e7d32;
        }
        ul {
            margin-left: 20px;
            color: #666;
            line-height: 1.8;
        }
        .hero {
            text-align: center;
            padding: 20px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border-radius: 10px;
            margin-bottom: 30px;
        }
        .hero h1 {
            color: white;
            font-size: 36px;
        }
        .hero p {
            color: rgba(255,255,255,0.9);
            font-size: 18px;
        }
    </style>
</head>
<body>
    <div class="container">
        <a href="/" class="back-btn">← Back to Home</a>

        <div class="hero">
            <h1>🔧 PDF Tools</h1>
            <p>Free & Open Source</p>
        </div>

        <div class="section">
            <h2>💡 About This Project</h2>
            <p>
                PDF Tools is a free, open-source web application for merging and editing PDF files.
                Built with modern technology and AI assistance, this tool provides powerful PDF
                manipulation capabilities accessible to everyone, completely free of charge.
            </p>
            <div style="margin-top: 15px;">
                <span class="badge license">MIT License</span>
                <span class="badge">Open Source</span>
                <span class="badge">No Registration Required</span>
                <span class="badge">Privacy Focused</span>
            </div>
        </div>

        <div class="section">
            <h2>🤖 AI-Assisted Development</h2>

            <div class="credit-item">
                <strong>Claude.ai by Anthropic</strong>
                <span class="badge ai">AI Assistant</span>
                <p>
                    This application was developed with significant assistance from Claude.ai,
                    an AI assistant created by Anthropic. The collaboration involved architecture
                    design, code generation, user interface development, and implementation of
                    PDF processing features.
                </p>
                <p>
                    <strong>Model:</strong> Claude Sonnet 4.5<br>
                    <strong>Website:</strong> <a href="https://claude.ai" target="_blank">https://claude.ai</a><br>
                    <strong>Company:</strong> Anthropic PBC
                </p>
            </div>

            <div class="credit-item">
                <strong>Human Developer</strong>
                <p>
                    Project direction, testing, customization, and deployment managed by the human developer.
                    The final implementation represents a collaborative effort between human creativity
                    and AI capabilities.
                </p>
            </div>
        </div>

        <div class="section">
            <h2>🔧 Technology Stack</h2>

            <div class="credit-item">
                <strong>pdfcpu</strong>
                <span class="badge license">Apache License 2.0</span>
                <p>
                    A powerful PDF processing library written in Go. This is the core technology
                    that powers all PDF manipulation features in this application.
                </p>
                <p>
                    <strong>Author:</strong> Horst Rutter<br>
                    <strong>Repository:</strong> <a href="https://github.com/pdfcpu/pdfcpu" target="_blank">github.com/pdfcpu/pdfcpu</a><br>
                    <strong>License:</strong> Apache License 2.0
                </p>
            </div>

            <div class="credit-item">
                <strong>Go Programming Language</strong>
                <span class="badge license">BSD 3-Clause</span>
                <p>
                    The backend server and PDF operations are built using Go (Golang),
                    a fast, reliable, and efficient programming language.
                </p>
                <p>
                    <strong>Website:</strong> <a href="https://golang.org" target="_blank">https://golang.org</a><br>
                    <strong>License:</strong> BSD 3-Clause License
                </p>
            </div>
        </div>

        <div class="section">
            <h2>📄 License Information</h2>
            <h3>This Application (PDF Tools)</h3>
            <p>
                <strong>License:</strong> MIT License<br>
                <strong>Copyright:</strong> © 2025
            </p>
            <p>
                Permission is hereby granted, free of charge, to any person obtaining a copy
                of this software and associated documentation files, to deal in the Software
                without restriction, including without limitation the rights to use, copy,
                modify, merge, publish, distribute, sublicense, and/or sell copies of the Software.
            </p>

            <h3>Third-Party Licenses</h3>
            <ul>
                <li><strong>pdfcpu:</strong> Apache License 2.0 - Requires attribution and license inclusion</li>
                <li><strong>Go:</strong> BSD 3-Clause - Permissive open source license</li>
            </ul>
        </div>

        <div class="section">
            <h2>🔒 Privacy & Security</h2>
            <p><strong>Your privacy is our priority:</strong></p>
            <ul>
                <li>All PDF processing happens on the server temporarily</li>
                <li>Files are automatically deleted immediately after processing</li>
                <li>We do not store, access, or analyze your PDF content</li>
                <li>No user data collection or tracking</li>
                <li>No cookies or analytics</li>
                <li>No registration or login required</li>
            </ul>
        </div>

        <div class="section">
            <h2>⚖️ Disclaimer</h2>
            <p>
                This tool is provided "as is" without warranty of any kind, express or implied.
                The developers and contributors are not responsible for any data loss, corruption,
                or issues arising from the use of this service. Use at your own risk.
            </p>
            <p>
                While we take precautions to ensure file security and privacy, users should not
                upload sensitive or confidential documents to any online service without proper
                risk assessment.
            </p>
        </div>

        <div class="section">
            <h2>📦 Source Code</h2>
            <p>
                This project is open source and available on GitHub. You can view the code,
                report issues, or contribute improvements.
            </p>
            <p>
                <strong>Repository:</strong> <a href="https://github.com/rajmahavir/PDF-Tools" target="_blank">github.com/rajmahavir/PDF-Tools</a>
            </p>
            <p style="margin-top: 15px;">
                <strong>How to Contribute:</strong>
            </p>
            <ul>
                <li>Report bugs or suggest features via GitHub Issues</li>
                <li>Submit pull requests for improvements</li>
                <li>Share the project with others who might find it useful</li>
                <li>Star the repository to show your support</li>
            </ul>
        </div>

        <div class="section">
            <h2>🙏 Acknowledgments</h2>
            <p>Special thanks to:</p>
            <ul>
                <li><strong>Anthropic</strong> - For creating Claude.ai and enabling AI-assisted development</li>
                <li><strong>Horst Rutter</strong> - For developing and maintaining pdfcpu</li>
                <li><strong>Go Team</strong> - For the excellent Go programming language</li>
                <li><strong>Open Source Community</strong> - For making projects like this possible</li>
            </ul>
        </div>

        <div style="text-align: center; margin-top: 40px; padding-top: 20px; border-top: 2px solid #eee;">
            <p style="color: #999; font-size: 14px;">
                Made with ❤️ using Claude.ai<br>
                © 2025 PDF Tools • Free Forever
            </p>
        </div>
    </div>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
