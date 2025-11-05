package handlers

import (
	"net/http"
	"pdf-merger/templates"
)

func HandleRemovePage(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Remove Pages - PDF Tools</title>
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
            max-height: 300px;
        }
        .pdf-preview.visible {
            display: block;
        }
        .pdf-preview iframe {
            width: 100%;
            height: 300px;
            border: none;
        }
        .pages-to-remove {
            margin-top: 20px;
            display: none;
        }
        .pages-to-remove.visible {
            display: block;
        }
        .page-selector {
            display: flex;
            flex-wrap: wrap;
            gap: 8px;
            margin-top: 10px;
            max-height: 200px;
            overflow-y: auto;
            padding: 10px;
            background: #f8f9fa;
            border-radius: 8px;
        }
        .page-checkbox {
            display: flex;
            align-items: center;
            gap: 5px;
            padding: 8px 12px;
            background: white;
            border: 2px solid #ddd;
            border-radius: 6px;
            cursor: pointer;
            transition: all 0.2s;
            user-select: none;
        }
        .page-checkbox:hover {
            border-color: #667eea;
        }
        .page-checkbox input[type="checkbox"] {
            cursor: pointer;
        }
        .page-checkbox.selected {
            background: #667eea;
            border-color: #667eea;
            color: white;
        }
        .hint {
            font-size: 12px;
            color: #999;
            margin-top: 5px;
        }
        .action-buttons {
            display: flex;
            gap: 10px;
            margin-top: 20px;
        }
        button {
            flex: 1;
            padding: 15px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border: none;
            border-radius: 10px;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
            transition: transform 0.2s, box-shadow 0.2s;
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
        .select-all-btn {
            background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
            margin-bottom: 10px;
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
        <h1>✂️ Remove Pages</h1>
        <p class="subtitle">Select pages to remove from your PDF</p>
        
        <form id="uploadForm" enctype="multipart/form-data">
            <div class="upload-section">
                <label for="pdf">Upload PDF *</label>
                <div class="file-input-wrapper">
                    <input type="file" id="pdf" name="pdf" accept=".pdf" required>
                    <label for="pdf" class="file-input-label">
                        Choose PDF file
                    </label>
                </div>
                <div id="fileName" class="file-name"></div>
                <div id="pdfInfo" class="pdf-info"></div>
                <div id="pdfPreview" class="pdf-preview"></div>
            </div>

            <div id="pageSelectorSection" class="pages-to-remove">
                <label>Select pages to remove:</label>
                <div class="hint">Click on pages you want to delete from the PDF</div>
                <button type="button" class="select-all-btn" onclick="toggleSelectAll()">Select All</button>
                <div id="pageSelector" class="page-selector"></div>
                <div class="action-buttons">
                    <button type="submit">Remove Selected Pages</button>
                </div>
            </div>
        </form>

        <div class="loading" id="loading">
            <div class="spinner"></div>
            <p style="margin-top: 10px; color: #666;">Processing PDF...</p>
        </div>

        <div id="result"></div>
        ` + templates.GetFooterHTML() + `
    </div>

    <script>
        var pdfInput = document.getElementById('pdf');
        var fileName = document.getElementById('fileName');
        var pdfInfoDiv = document.getElementById('pdfInfo');
        var pdfPreviewDiv = document.getElementById('pdfPreview');
        var pageSelectorSection = document.getElementById('pageSelectorSection');
        var pageSelector = document.getElementById('pageSelector');
        var currentPDFFile = null;
        var totalPages = 0;
        var selectedPages = [];

        pdfInput.addEventListener('change', function(e) {
            if (e.target.files.length > 0) {
                currentPDFFile = e.target.files[0];
                fileName.textContent = 'Selected: ' + currentPDFFile.name;
                loadPDFInfo(currentPDFFile);
            }
        });

        function loadPDFInfo(file) {
            var formData = new FormData();
            formData.append('pdf', file);

            pdfInfoDiv.innerHTML = 'Loading info...';
            pdfInfoDiv.classList.add('visible');
            pdfPreviewDiv.classList.remove('visible');
            pageSelectorSection.classList.remove('visible');

            fetch('/pdfinfo', {
                method: 'POST',
                body: formData
            }).then(function(response) {
                return response.json();
            }).then(function(data) {
                if (data.error) {
                    pdfInfoDiv.innerHTML = 'Error: ' + data.error;
                } else {
                    totalPages = data.pageCount;
                    pdfInfoDiv.innerHTML = '<strong>Pages:</strong> ' + data.pageCount + 
                                          ' | <strong>Size:</strong> ' + formatFileSize(file.size);
                    
                    var url = URL.createObjectURL(file);
                    pdfPreviewDiv.innerHTML = '<iframe src="' + url + '#view=FitH"></iframe>';
                    pdfPreviewDiv.classList.add('visible');
                    
                    createPageSelector(data.pageCount);
                    pageSelectorSection.classList.add('visible');
                }
            }).catch(function(error) {
                pdfInfoDiv.innerHTML = 'Error loading PDF info';
            });
        }

        function createPageSelector(pageCount) {
            pageSelector.innerHTML = '';
            selectedPages = [];
            
            for (var i = 1; i <= pageCount; i++) {
                var pageDiv = document.createElement('div');
                pageDiv.className = 'page-checkbox';
                pageDiv.setAttribute('data-page', i);
                
                var checkbox = document.createElement('input');
                checkbox.type = 'checkbox';
                checkbox.id = 'page-' + i;
                checkbox.value = i;
                
                var label = document.createElement('label');
                label.setAttribute('for', 'page-' + i);
                label.textContent = 'Page ' + i;
                label.style.cursor = 'pointer';
                
                pageDiv.appendChild(checkbox);
                pageDiv.appendChild(label);
                
                pageDiv.addEventListener('click', function(e) {
                    var checkbox = this.querySelector('input[type="checkbox"]');
                    if (e.target !== checkbox) {
                        checkbox.checked = !checkbox.checked;
                    }
                    
                    if (checkbox.checked) {
                        this.classList.add('selected');
                        selectedPages.push(parseInt(checkbox.value));
                    } else {
                        this.classList.remove('selected');
                        var index = selectedPages.indexOf(parseInt(checkbox.value));
                        if (index > -1) {
                            selectedPages.splice(index, 1);
                        }
                    }
                });
                
                pageSelector.appendChild(pageDiv);
            }
        }

        function toggleSelectAll() {
            var checkboxes = pageSelector.querySelectorAll('input[type="checkbox"]');
            var allSelected = selectedPages.length === totalPages;
            
            selectedPages = [];
            checkboxes.forEach(function(checkbox) {
                checkbox.checked = !allSelected;
                var pageDiv = checkbox.parentElement;
                if (!allSelected) {
                    pageDiv.classList.add('selected');
                    selectedPages.push(parseInt(checkbox.value));
                } else {
                    pageDiv.classList.remove('selected');
                }
            });
        }

        function formatFileSize(bytes) {
            if (bytes < 1024) return bytes + ' B';
            if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB';
            return (bytes / (1024 * 1024)).toFixed(1) + ' MB';
        }

        document.getElementById('uploadForm').addEventListener('submit', function(e) {
            e.preventDefault();
            
            if (selectedPages.length === 0) {
                alert('Please select at least one page to remove');
                return;
            }
            
            if (selectedPages.length === totalPages) {
                alert('You cannot remove all pages. At least one page must remain.');
                return;
            }
            
            var formData = new FormData();
            formData.append('pdf', currentPDFFile);
            formData.append('pagesToRemove', selectedPages.sort(function(a, b) { return a - b; }).join(','));
            
            var resultDiv = document.getElementById('result');
            var loadingDiv = document.getElementById('loading');
            var submitBtn = e.target.querySelector('button[type="submit"]');
            
            resultDiv.innerHTML = '';
            loadingDiv.style.display = 'block';
            submitBtn.disabled = true;

            fetch('/remove-pages', {
                method: 'POST',
                body: formData
            }).then(function(response) {
                loadingDiv.style.display = 'none';
                submitBtn.disabled = false;

                if (response.ok) {
                    return response.blob().then(function(blob) {
                        var url = URL.createObjectURL(blob);
                        
                        resultDiv.innerHTML = 
                            '<div class="success">Pages removed successfully! Remaining pages: ' + (totalPages - selectedPages.length) + '</div>' +
                            '<iframe id="pdfViewer" src="' + url + '"></iframe>' +
                            '<button class="download-btn" onclick="downloadPDF(\'' + url + '\')">Download Modified PDF</button>';
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
            a.download = 'modified_' + Date.now() + '.pdf';
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
