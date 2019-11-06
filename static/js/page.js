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


$.ajax({url:"docList",success:function(result){

        let titles = Object.keys(result)
        let thisTitle = decodeURI(document.URL.split("/").pop())
        let index = titles.lastIndexOf(thisTitle)
        if (index - 1 >= 0 ){
            $("<span class=\"posts-list-meta\">上一篇： </span>" +
                "<a class=\"posts-list-name\" href=\""+ titles[index - 1]+"\"> "+ titles[index - 1]+"</a>"
            ).prependTo($("#doc-link"))
        }
        if (index + 1 < titles.length ){
            $("<span class=\"posts-list-meta\">下一篇： </span>" +
                "<a class=\"posts-list-name\" href=\""+ titles[index + 1]+"\"> "+ titles[index + 1]+"</a>"
            ).prependTo($("#doc-link"))
        }
}});