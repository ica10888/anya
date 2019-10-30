
(function() {
    hljs.initHighlightingOnLoad();
    document.addEventListener('DOMContentLoaded', (event) => {
        document.querySelectorAll('.highlight pre').forEach((block) => {
            hljs.highlightBlock(block);
        });
    });
})()