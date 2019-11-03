(function() {
    hljs.initHighlightingOnLoad();
    document.addEventListener('DOMContentLoaded', (event) => {
        document.querySelectorAll('.highlight pre').forEach((block) => {
            let odiv = block.parentElement.classList
            odiv.forEach((s) => {
                if(s.match("highlight-.*")){
                    let langs = s.replace("highlight-","")
                    hljs.configure({languages: [langs]});
                    hljs.highlightBlock(block);
                }})

        });
    });
})()