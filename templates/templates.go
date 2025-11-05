package templates

// GetFooterHTML returns the common footer HTML for all pages
func GetFooterHTML() string {
	return `
	<footer style="margin-top: 40px; padding-top: 20px; border-top: 1px solid #eee; text-align: center; color: #666; font-size: 13px;">
		<div style="margin-bottom: 10px;">
			<span style="display: inline-block; margin: 0 10px;">🤖 Built with <a href="https://claude.ai" target="_blank" style="color: #667eea; text-decoration: none;">Claude.ai</a></span>
			<span style="display: inline-block; margin: 0 10px;">⚡ Powered by <a href="https://github.com/pdfcpu/pdfcpu" target="_blank" style="color: #667eea; text-decoration: none;">pdfcpu</a></span>
		</div>
		<div>
			<a href="/credits" style="color: #667eea; text-decoration: none; margin: 0 10px;">Credits</a>
			<span style="color: #ddd;">|</span>
			<a href="https://github.com/rajmahavir/PDF-Tools" target="_blank" style="color: #667eea; text-decoration: none; margin: 0 10px;">Source Code</a>
			<span style="color: #ddd;">|</span>
			<span style="margin: 0 10px;">MIT License</span>
		</div>
		<div style="margin-top: 10px; font-size: 12px; color: #999;">
			Free & Open Source PDF Tools
		</div>
	</footer>
	`
}

// GetCommonStyles returns common CSS styles used across pages
func GetCommonStyles() string {
	return `
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 20px;
        }
        .container {
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            padding: 40px;
            max-width: 600px;
            width: 100%;
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
            font-size: 28px;
        }
        .subtitle {
            color: #666;
            margin-bottom: 30px;
            font-size: 14px;
        }
	`
}
