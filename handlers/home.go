package handlers

import (
	"net/http"
	"pdf-merger/templates"
)

// HandleHome serves the home page
func HandleHome(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>PDF Tools - Free Online PDF Editor</title>
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
            max-width: 700px;
            width: 100%;
        }
        h1 {
            color: #333;
            margin-bottom: 10px;
            font-size: 32px;
            text-align: center;
        }
        .subtitle {
            color: #666;
            margin-bottom: 40px;
            font-size: 14px;
            text-align: center;
        }
        .options-grid {
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 20px;
            margin-top: 20px;
        }
        .option-card {
            background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%);
            border-radius: 15px;
            padding: 30px 20px;
            text-align: center;
            cursor: pointer;
            transition: all 0.3s;
            border: 2px solid transparent;
        }
        .option-card:hover {
            transform: translateY(-5px);
            box-shadow: 0 10px 30px rgba(0,0,0,0.2);
            border-color: #667eea;
        }
        .option-card.remove {
            background: linear-gradient(135deg, #ffecd2 0%, #fcb69f 100%);
        }
        .option-card.merge {
            background: linear-gradient(135deg, #a8edea 0%, #fed6e3 100%);
        }
        .option-icon {
            font-size: 48px;
            margin-bottom: 15px;
        }
        .option-title {
            font-size: 20px;
            font-weight: 600;
            color: #333;
            margin-bottom: 10px;
        }
        .option-desc {
            font-size: 13px;
            color: #666;
        }
        @media (max-width: 600px) {
            .options-grid {
                grid-template-columns: 1fr;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🔧 PDF Tools</h1>
        <p class="subtitle">Free & Open Source PDF Editor</p>

        <div class="options-grid">
            <div class="option-card remove" onclick="window.location.href='/remove'">
                <div class="option-icon">✂️</div>
                <div class="option-title">Remove Pages</div>
                <div class="option-desc">Delete specific pages from your PDF or split it into parts</div>
            </div>

            <div class="option-card merge" onclick="window.location.href='/merge'">
                <div class="option-icon">📄</div>
                <div class="option-title">Merge PDFs</div>
                <div class="option-desc">Insert all pages from one PDF into another at any position</div>
            </div>
        </div>
        ` + templates.GetFooterHTML() + `
    </div>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
