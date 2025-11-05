package handlers

import (
	"net/http"
	"pdf-merger/templates"
)

func HandleMergePage(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Merge PDFs - PDF Tools</title>
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
        .upload-section {
            margin-bottom: 25px;
        }
        label {
            display: block;
            margin-bottom: 8px;
            color: #333;
            font-weight: 500;
            font-size: 14px;
        }
        .file-input-wrapper {
            position: relative;
            overflow: hidden;
            display: inline-block;
            width: 100%;
        }
        .file-input-wrapper input[type=file] {
            position: absolute;
            left: -9999px;
        }
        .file-input-label {
            display: flex;
            align-items: center;
            justify-content: center;
            padding: 15px;
            background: #f8f9fa;
            border: 2px dashed #ddd;
            border-radius: 10px;
            cursor: pointer;
            transition: all 0.3s;
        }
        .file-input-label:hover {
            background: #e9ecef;
            border-color: #667eea;
        }
        .file-name {
            margin-top: 8px;
            font-size: 13px;
            color: #666;
            font-style: italic;
        }
        .pdf-info {
            margin-top: 10px;
            padding: 10px;
            background: #f0f7ff;
            border: 1px solid #b3d9ff;
            border-radius: 8px;
            font-size: 13px;
            display: none;
        }
        .pdf-info.visible {
            display: block;
        }
        .pdf-info strong {
            color: #0066cc;
        }
        .pdf-preview {
            margin-top: 10px;
            border: 2px solid #ddd;
            border-radius: 8px;
            overflow: hidden;
            display: none;
            max-height: 200px;
        }
        .pdf-preview.visible {
            display: block;
        }
        .pdf-preview iframe {
            width: 100%;
            height: 200px;
            border: none;
        }
        input[type="number"] {
            width: 100%;
            padding: 12px;
            border: 2px solid #ddd;
            border-radius: 10px;
            font-size: 16px;
            transition: border-color 0.3s;
        }
        input[type="number"]:focus {
            outline: none;
            border-color: #667eea;
        }
        .hint {
            font-size: 12px;
            color: #999;
            margin-top: 5px;
        }
        button {
            width: 100%;
            padding: 15px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border: none;
            border-radius: 10px;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
            transition: transform 0.2s, box-shadow 0.2s;
            margin-top: 10px;
        }
        button:hover {
            transform: translateY(-2px);
            box-shadow: 0 10px 20px rgba(102, 126, 234, 0.4);
        }
        button:active {
            transform: translateY(0);
        }
        button:disabled {
            opacity: 0.6;
            cursor: not-allowed;
            transform: none;
        }
        .error {
            background: #fee;
            border: 1px solid #fcc;
            color: #c33;
            padding: 12px;
            border-radius: 10px;
            margin-top: 15px;
            font-size: 14px;
        }
        .success {
            background: #efe;
            border: 1px solid #cfc;
            color: #3c3;
            padding: 12px;
            border-radius: 10px;
            margin-top: 15px;
            font-size: 14px;
        }
        #result {
            margin-top: 20px;
        }
        #pdfViewer {
            width: 100%;
            height: 600px;
            border: 2px solid #ddd;
            border-radius: 10px;
            margin-top: 15px;
        }
        .download-btn {
            background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
            margin-top: 10px;
        }
        .loading {
            display: none;
            text-align: center;
            margin-top: 20px;
        }
        .spinner {
            border: 3px solid #f3f3f3;
            border-top: 3px solid #667eea;
            border-radius: 50%;
            width: 40px;
            height: 40px;
            animation: spin 1s linear infinite;
            margin: 0 auto;
        }
        @keyframes spin {
            0% { transform: rotate(0deg); }
            100% { transform: rotate(360deg); }
        }
    </style>
</head>
<body>
    <div class="container">
        <a href="/" class="back-btn">← Back to Home</a>
        <h1>📄 Merge PDFs</h1>
        <p class="subtitle">Insert all pages from one PDF into another at a specific position</p>
        
        <form id="uploadForm" enctype="multipart/form-data">
            <div class="upload-section">
                <label for="pdf1">First PDF (Base document) *</label>
                <div class="file-input-wrapper">
                    <input type="file" id="pdf1" name="pdf1" accept=".pdf" required>
                    <label for="pdf1" class="file-input-label">
                        Choose PDF file
                    </label>
                </div>
                <div id="file1Name" class="file-name"></div>
                <div id="pdf1Info" class="pdf-info"></div>
                <div id="pdf1Preview" class="pdf-preview"></div>
            </div>

            <div class="upload-section">
                <label for="pdf2">Second PDF (To be inserted) *</label>
                <div class="file-input-wrapper">
                    <input type="file" id="pdf2" name="pdf2" accept=".pdf" required>
                    <label for="pdf2" class="file-input-label">
                        Choose PDF file
                    </label>
                </div>
                <div id="file2Name" class="file-name"></div>
                <div id="pdf2Info" class="pdf-info"></div>
                <div id="pdf2Preview" class="pdf-preview"></div>
            </div>

            <div class="upload-section">
                <label for="pageNumber">Insert at page number *</label>
                <input type="number" id="pageNumber" name="pageNumber" min="1" required placeholder="e.g., 3">
                <div class="hint">Pages from PDF 2 will be inserted after this page in PDF 1</div>
            </div>

            <button type="submit">Merge PDFs</button>
        </form>

        <div class="loading" id="loading">
            <div class="spinner"></div>
            <p style="margin-top: 10px; color: #666;">Processing PDFs...</p>
        </div>

        <div id="result"></div>
        ` + templates.GetFooterHTML() + `
    </div>

    <script>
        var pdf1Input = document.getElementById('pdf1');
        var pdf2Input = document.getElementById('pdf2');
        var file1Name = document.getElementById('file1Name');
        var file2Name = document.getElementById('file2Name');
        var pdf1InfoDiv = document.getElementById('pdf1Info');
        var pdf2InfoDiv = document.getElementById('pdf2Info');
        var pdf1PreviewDiv = document.getElementById('pdf1Preview');
        var pdf2PreviewDiv = document.getElementById('pdf2Preview');

        pdf1Input.addEventListener('change', function(e) {
            if (e.target.files.length > 0) {
                var file = e.target.files[0];
                file1Name.textContent = 'Selected: ' + file.name;
                loadPDFInfo(file, pdf1InfoDiv, pdf1PreviewDiv, 1);
            }
        });

        pdf2Input.addEventListener('change', function(e) {
            if (e.target.files.length > 0) {
                var file = e.target.files[0];
                file2Name.textContent = 'Selected: ' + file.name;
                loadPDFInfo(file, pdf2InfoDiv, pdf2PreviewDiv, 2);
            }
        });

        function loadPDFInfo(file, infoDiv, previewDiv, pdfNumber) {
            var formData = new FormData();
            formData.append('pdf', file);

            infoDiv.innerHTML = 'Loading info...';
            infoDiv.classList.add('visible');
            previewDiv.classList.remove('visible');

            fetch('/pdfinfo', {
                method: 'POST',
                body: formData
            }).then(function(response) {
                return response.json();
            }).then(function(data) {
                if (data.error) {
                    infoDiv.innerHTML = 'Error: ' + data.error;
                } else {
                    infoDiv.innerHTML = '<strong>Pages:</strong> ' + data.pageCount + 
                                      ' | <strong>Size:</strong> ' + formatFileSize(file.size);
                    
                    var url = URL.createObjectURL(file);
                    previewDiv.innerHTML = '<iframe src="' + url + '#view=FitH"></iframe>';
                    previewDiv.classList.add('visible');
                }
            }).catch(function(error) {
                infoDiv.innerHTML = 'Error loading PDF info';
            });
        }

        function formatFileSize(bytes) {
            if (bytes < 1024) return bytes + ' B';
            if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB';
            return (bytes / (1024 * 1024)).toFixed(1) + ' MB';
        }

        document.getElementById('uploadForm').addEventListener('submit', function(e) {
            e.preventDefault();
            
            var formData = new FormData(e.target);
            var resultDiv = document.getElementById('result');
            var loadingDiv = document.getElementById('loading');
            var submitBtn = e.target.querySelector('button[type="submit"]');
            
            resultDiv.innerHTML = '';
            loadingDiv.style.display = 'block';
            submitBtn.disabled = true;

            fetch('/merge-pdfs', {
                method: 'POST',
                body: formData
            }).then(function(response) {
                loadingDiv.style.display = 'none';
                submitBtn.disabled = false;

                if (response.ok) {
                    return response.blob().then(function(blob) {
                        var url = URL.createObjectURL(blob);
                        
                        resultDiv.innerHTML = 
                            '<div class="success">PDFs merged successfully!</div>' +
                            '<iframe id="pdfViewer" src="' + url + '"></iframe>' +
                            '<button class="download-btn" onclick="downloadPDF(\'' + url + '\')">Download Merged PDF</button>';
                    });
                } else {
                    return response.text().then(function(error) {
                        resultDiv.innerHTML = '<div class="error">Error: ' + error + '</div>';
                    });
                }
            }).catch(function(error) {
                loadingDiv.style.display = 'none';
                submitBtn.disabled = false;
                resultDiv.innerHTML = '<div class="error">Error: ' + error.message + '</div>';
            });
        });

        function downloadPDF(url) {
            var a = document.createElement('a');
            a.href = url;
            a.download = 'merged_' + Date.now() + '.pdf';
            document.body.appendChild(a);
            a.click();
            document.body.removeChild(a);
        }
    </script>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
